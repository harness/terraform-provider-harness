package triggers

import (
	"context"
	"net/http"
	"time"

	"github.com/antihax/optional"
	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceTriggers() *schema.Resource {
	resource := &schema.Resource{
		Description:   "Resource for creating triggers in Harness.",
		ReadContext:   resourceTriggersRead,
		UpdateContext: resourceTriggersCreateOrUpdate,
		CreateContext: resourceTriggersCreateOrUpdate,
		DeleteContext: resourceTriggersDelete,
		Importer:      helpers.TriggerResourceImporter,

		Schema: map[string]*schema.Schema{
			"target_id": {
				Description: "Identifier of the target pipeline",
				Type:        schema.TypeString,
				Required:    true,
			},
			"ignore_error": {
				Description: "ignore error default false",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"yaml": {
				Description:      "trigger yaml." + helpers.Descriptions.YamlText.String(),
				Type:             schema.TypeString,
				Required:         true,
				DiffSuppressFunc: helpers.YamlDiffSuppressFunction,
			},
			"if_match": {
				Description: "if-Match",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"webhook_url": {
				Description: "webhook url",
				Type:        schema.TypeString,
				Computed:    true,
				Optional:    true,
			},
			"executor_type": {
				Description: "Type of principal that executes the pipeline when this trigger fires. Valid values: USER, SERVICE_ACCOUNT. When set, the v2 trigger API is used and executor_identifier must also be provided. Required when the account setting enforce_executor_identity_for_triggers is enabled.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"executor_identifier": {
				Description: "Identifier of the executor principal. For USER type, this is the user's UUID (obtain from the Harness user list API). For SERVICE_ACCOUNT type, this is the service account identifier string (e.g. \"sa005\").",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"executor_org_id": {
				Description: "Org identifier of the SERVICE_ACCOUNT executor. Set for org-scoped or project-scoped service accounts. Omit for account-level service accounts.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"executor_project_id": {
				Description: "Project identifier of the SERVICE_ACCOUNT executor. Set for project-scoped service accounts. Omit for account-level or org-level service accounts.",
				Type:        schema.TypeString,
				Optional:    true,
			},
		},
	}
	helpers.SetProjectLevelResourceSchema(resource.Schema)

	return resource
}

func resourceTriggersRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	id := d.Id()

	resp, httpResp, err := c.TriggersApi.GetTrigger(ctx, c.AccountId,
		d.Get("org_id").(string),
		d.Get("project_id").(string), d.Get("target_id").(string), id)

	if httpResp.StatusCode == 404 {
		d.SetId("")
		d.MarkNewResource()
		return nil
	}

	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	readTriggers(d, resp.Data)

	return nil
}

func resourceTriggersCreateOrUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	var err error
	var resp nextgen.ResponseDtongTriggerResponse
	var httpResp *http.Response
	id := d.Id()

	if isAddTriggerExecutorInfo(d) {
		return resourceTriggersCreateOrUpdateWithExecutorInfo(ctx, d, c, id)
	}

	if id == "" {
		resp, httpResp, err = c.TriggersApi.CreateTrigger(ctx, d.Get("yaml").(string), c.AccountId,
			d.Get("org_id").(string),
			d.Get("project_id").(string),
			d.Get("target_id").(string), &nextgen.TriggersApiCreateTriggerOpts{
				WithServiceV2: optional.NewBool(true),
			})
		// A FORCED PAUSE TO PREVENT DUPLICATE WEBHOOK CREATION.
		time.Sleep(5 * time.Second)
	} else {
		resp, httpResp, err = c.TriggersApi.UpdateTrigger(ctx, d.Get("yaml").(string), c.AccountId, d.Get("org_id").(string),
			d.Get("project_id").(string),
			d.Get("target_id").(string), id, &nextgen.TriggersApiUpdateTriggerOpts{
				IfMatch: helpers.BuildField(d, "if_match"),
			})
	}

	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	readTriggers(d, resp.Data)

	return nil
}

// isAddTriggerExecutorInfo returns true when executor_type is set in config,
// indicating the v2 API should be used with an executor identity.
func isAddTriggerExecutorInfo(d *schema.ResourceData) bool {
	return d.Get("executor_type").(string) != ""
}

// resourceTriggersCreateOrUpdateWithExecutorInfo handles create/update using the v2
// trigger API, which carries executor identity alongside the trigger YAML.
func resourceTriggersCreateOrUpdateWithExecutorInfo(ctx context.Context, d *schema.ResourceData, c *nextgen.APIClient, id string) diag.Diagnostics {
	var err error
	var resp nextgen.ResponseDtongTriggerResponse
	var httpResp *http.Response

	reqBody := nextgen.TriggerYamlRequest{
		Yaml:         d.Get("yaml").(string),
		ExecutorInfo: buildExecutorInfo(d, c.AccountId),
	}

	if id == "" {
		resp, httpResp, err = c.TriggersApi.CreateTriggerV2(ctx, reqBody, c.AccountId,
			d.Get("org_id").(string),
			d.Get("project_id").(string),
			d.Get("target_id").(string), &nextgen.TriggersApiCreateTriggerV2Opts{
				WithServiceV2: optional.NewBool(true),
			})
		// A FORCED PAUSE TO PREVENT DUPLICATE WEBHOOK CREATION.
		time.Sleep(5 * time.Second)
	} else {
		resp, httpResp, err = c.TriggersApi.UpdateTriggerV2(ctx, reqBody, c.AccountId,
			d.Get("org_id").(string),
			d.Get("project_id").(string),
			d.Get("target_id").(string), id, &nextgen.TriggersApiUpdateTriggerV2Opts{
				IfMatch: helpers.BuildField(d, "if_match"),
			})
	}

	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	readTriggers(d, resp.Data)

	return nil
}

// buildExecutorInfo constructs a TriggerExecutorInfo from the resource data.
// For SERVICE_ACCOUNT type, scope is inferred from which fields are set:
//   - neither executor_org_id nor executor_project_id set → account-level SA
//   - only executor_org_id set                            → org-level SA
//   - both executor_org_id and executor_project_id set    → project-level SA
func buildExecutorInfo(d *schema.ResourceData, accountId string) *nextgen.TriggerExecutorInfo {
	info := &nextgen.TriggerExecutorInfo{
		Type_: d.Get("executor_type").(string),
		Uuid:  d.Get("executor_identifier").(string),
	}

	if info.Type_ == "SERVICE_ACCOUNT" {
		info.AccountIdentifier = accountId
		if v, ok := d.GetOk("executor_org_id"); ok {
			info.OrgIdentifier = v.(string)
		}
		if v, ok := d.GetOk("executor_project_id"); ok {
			info.ProjectIdentifier = v.(string)
		}
	}

	return info
}

func resourceTriggersDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	_, httpResp, err := c.TriggersApi.DeleteTrigger(ctx, c.AccountId, d.Get("org_id").(string), d.Get("project_id").(string), d.Get("target_id").(string), d.Id(), &nextgen.TriggersApiDeleteTriggerOpts{
		IfMatch: helpers.BuildField(d, "if_match"),
	})

	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	return nil
}

func readTriggers(d *schema.ResourceData, trigger *nextgen.NgTriggerResponse) {
	d.SetId(trigger.Identifier)
	d.Set("identifier", trigger.Identifier)
	d.Set("name", trigger.Name)
	d.Set("org_id", trigger.OrgIdentifier)
	d.Set("project_id", trigger.ProjectIdentifier)
	d.Set("target_id", trigger.TargetIdentifier)
	d.Set("yaml", trigger.Yaml)
	d.Set("webhook_url", trigger.WebhookUrl)
	if trigger.ExecutorInfo != nil {
		d.Set("executor_type", trigger.ExecutorInfo.Type_)
		d.Set("executor_identifier", trigger.ExecutorInfo.Uuid)
		d.Set("executor_org_id", trigger.ExecutorInfo.OrgIdentifier)
		d.Set("executor_project_id", trigger.ExecutorInfo.ProjectIdentifier)
	}
}
