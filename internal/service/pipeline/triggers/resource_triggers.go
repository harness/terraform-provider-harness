package triggers

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/antihax/optional"
	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func ResourceTriggers() *schema.Resource {
	resource := &schema.Resource{
		Description:   "Resource for creating triggers in Harness.",
		ReadContext:   resourceTriggersRead,
		UpdateContext: resourceTriggersCreateOrUpdate,
		CreateContext: resourceTriggersCreateOrUpdate,
		DeleteContext: resourceTriggersDelete,
		Importer:      helpers.TriggerResourceImporter,
		CustomizeDiff: validateExecutorInfoDiff,

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
			"executor_info": {
				Description: "Executor identity for this trigger. When set, the v2 trigger API is used. Required when the account setting enforce_executor_identity_for_triggers is enabled. Optional+Computed: omitting this block does not remove a previously assigned executor from Harness (the server preserves it); use destroy+recreate to obtain a trigger with no executor.",
				Type:        schema.TypeList,
				Optional:    true,
				Computed:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"type": {
							Description:  "Type of principal that executes the pipeline when this trigger fires. Valid values: USER, SERVICE_ACCOUNT.",
							Type:         schema.TypeString,
							Required:     true,
							ValidateFunc: validation.StringInSlice([]string{"USER", "SERVICE_ACCOUNT"}, false),
						},
						"identifier": {
							Description: "Identifier of the executor principal. For USER type, this is the user's UUID (obtain from the Harness user list API). For SERVICE_ACCOUNT type, this is the service account identifier string (e.g. \"sa005\").",
							Type:        schema.TypeString,
							Required:    true,
						},
						"org_id": {
							Description: "Org identifier of the SERVICE_ACCOUNT executor. Set for org-scoped or project-scoped service accounts. Omit for account-level service accounts.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"project_id": {
							Description: "Project identifier of the SERVICE_ACCOUNT executor. Set for project-scoped service accounts. Omit for account-level or org-level service accounts.",
							Type:        schema.TypeString,
							Optional:    true,
						},
					},
				},
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

	if httpResp != nil && httpResp.StatusCode == 404 {
		d.SetId("")
		d.MarkNewResource()
		return nil
	}

	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	// Always reflect the server's executor_info in state. The schema is
	// Optional+Computed: omitting the block in config does not remove the
	// server-side executor (Harness preserves it on update), so ReadContext
	// must always write what the server reports.
	if err := readTriggers(d, resp.Data); err != nil {
		return diag.FromErr(err)
	}

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

	if err := readTriggers(d, resp.Data); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

// validateExecutorInfoDiff enforces scope constraints on the executor_info block:
//   - SERVICE_ACCOUNT: project_id requires org_id to also be set
//   - USER: org_id and project_id must not be set (not meaningful for user executors)
func validateExecutorInfoDiff(_ context.Context, d *schema.ResourceDiff, _ interface{}) error {
	v, ok := d.GetOk("executor_info")
	if !ok {
		return nil
	}
	list, ok := v.([]interface{})
	if !ok || len(list) == 0 {
		return nil
	}
	m, ok := list[0].(map[string]interface{})
	if !ok {
		return nil
	}

	execType, _ := m["type"].(string)
	orgID, _ := m["org_id"].(string)
	projectID, _ := m["project_id"].(string)

	switch execType {
	case "SERVICE_ACCOUNT":
		if projectID != "" && orgID == "" {
			return fmt.Errorf("executor_info: project_id requires org_id to be set for a SERVICE_ACCOUNT executor")
		}
	case "USER":
		if orgID != "" {
			return fmt.Errorf("executor_info: org_id must not be set for a USER executor")
		}
		if projectID != "" {
			return fmt.Errorf("executor_info: project_id must not be set for a USER executor")
		}
	}

	return nil
}

// isAddTriggerExecutorInfo returns true when executor_info is present in config or
// state, indicating the v2 API should be used. Because the schema is Optional+Computed,
// d.GetOk returns true when the value is in state even if omitted from config — this
// ensures a YAML-only update on a trigger with an existing executor routes through v2
// and preserves the executor, matching the Harness server's preserve-on-omission behavior.
//
// Limitation: there is no API path to clear an existing executor. The Harness server
// explicitly preserves an existing executor when executorInfo is omitted or null in
// UpdateTriggerV2 (see TriggerExecutorResolver.handleExecutorOnUpdate). To obtain a
// trigger with no executor, destroy and recreate it.
func isAddTriggerExecutorInfo(d *schema.ResourceData) bool {
	v, ok := d.GetOk("executor_info")
	if !ok {
		return false
	}
	return len(v.([]interface{})) > 0
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

	if err := readTriggers(d, resp.Data); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

// buildExecutorInfo constructs a TriggerExecutorInfo from the executor block.
// For SERVICE_ACCOUNT type, scope is inferred from which fields are set:
//   - neither org_id nor project_id set → account-level SA
//   - only org_id set                   → org-level SA
//   - both org_id and project_id set    → project-level SA
func buildExecutorInfo(d *schema.ResourceData, accountId string) *nextgen.TriggerExecutorInfo {
	v := d.Get("executor_info").([]interface{})
	if len(v) == 0 {
		return nil
	}
	m := v[0].(map[string]interface{})

	info := &nextgen.TriggerExecutorInfo{
		Type_: m["type"].(string),
		Uuid:  m["identifier"].(string),
	}

	if info.Type_ == "SERVICE_ACCOUNT" {
		info.AccountIdentifier = accountId
		if org, ok := m["org_id"].(string); ok && org != "" {
			info.OrgIdentifier = org
		}
		if proj, ok := m["project_id"].(string); ok && proj != "" {
			info.ProjectIdentifier = proj
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

// readTriggers populates d from the API response. executor_info is always written
// from the server response — the schema is Optional+Computed. Harness preserves an
// existing executor when an update omits executorInfo, so ReadContext must always
// reflect what the server reports rather than suppressing it.
func readTriggers(d *schema.ResourceData, trigger *nextgen.NgTriggerResponse) error {
	d.SetId(trigger.Identifier)

	if err := d.Set("identifier", trigger.Identifier); err != nil {
		return fmt.Errorf("error setting identifier: %w", err)
	}
	if err := d.Set("name", trigger.Name); err != nil {
		return fmt.Errorf("error setting name: %w", err)
	}
	if err := d.Set("org_id", trigger.OrgIdentifier); err != nil {
		return fmt.Errorf("error setting org_id: %w", err)
	}
	if err := d.Set("project_id", trigger.ProjectIdentifier); err != nil {
		return fmt.Errorf("error setting project_id: %w", err)
	}
	if err := d.Set("target_id", trigger.TargetIdentifier); err != nil {
		return fmt.Errorf("error setting target_id: %w", err)
	}
	if err := d.Set("yaml", trigger.Yaml); err != nil {
		return fmt.Errorf("error setting yaml: %w", err)
	}
	if err := d.Set("webhook_url", trigger.WebhookUrl); err != nil {
		return fmt.Errorf("error setting webhook_url: %w", err)
	}

	if trigger.ExecutorInfo != nil {
		if err := d.Set("executor_info", []interface{}{
			map[string]interface{}{
				"type":       trigger.ExecutorInfo.Type_,
				"identifier": trigger.ExecutorInfo.Uuid,
				"org_id":     trigger.ExecutorInfo.OrgIdentifier,
				"project_id": trigger.ExecutorInfo.ProjectIdentifier,
			},
		}); err != nil {
			return fmt.Errorf("error setting executor_info: %w", err)
		}
	} else {
		if err := d.Set("executor_info", []interface{}{}); err != nil {
			return fmt.Errorf("error clearing executor_info: %w", err)
		}
	}

	return nil
}
