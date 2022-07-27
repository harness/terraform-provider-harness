package triggers

import (
	"context"

	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceTriggers() *schema.Resource {
	resource := &schema.Resource{
		Description:   "Resource for craeting triggers in Harness.",
		ReadContext:   resourceTriggersRead,
		UpdateContext: resourceTriggersCreateOrUpdate,
		CreateContext: resourceTriggersCreateOrUpdate,
		DeleteContext: resourceTriggersDelete,
		Importer:      helpers.PipelineResourceImporter,

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
				Description: "trigger yaml",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"if_match": {
				Description: "if-Match",
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

	resp, _, err := c.TriggersApi.GetTrigger(ctx, c.AccountId,
		d.Get("org_id").(string),
		d.Get("project_id").(string), d.Get("target_id").(string), id)

	if err != nil {
		return helpers.HandleApiError(err, d)
	}

	readTriggers(d, resp.Data)

	return nil
}

func resourceTriggersCreateOrUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	var err error
	var resp nextgen.ResponseDtongTriggerResponse
	id := d.Id()

	trigger := buildTrigger(d)

	if id == "" {
		resp, _, err = c.TriggersApi.CreateTrigger(ctx, trigger.Yaml, c.AccountId,
			d.Get("org_id").(string),
			d.Get("project_id").(string),
			d.Get("target_id").(string))
	} else {
		resp, _, err = c.TriggersApi.UpdateTrigger(ctx, trigger.Yaml, c.AccountId, d.Get("org_id").(string),
			d.Get("project_id").(string),
			d.Get("target_id").(string), id, &nextgen.TriggersApiUpdateTriggerOpts{
				IfMatch: helpers.BuildField(d, "if_match"),
			})
	}

	if err != nil {
		return helpers.HandleApiError(err, d)
	}

	readTriggers(d, resp.Data)

	return nil
}

func resourceTriggersDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	trigger := buildTrigger(d)

	_, _, err := c.TriggersApi.DeleteTrigger(ctx, c.AccountId, trigger.OrgIdentifier, trigger.ProjectIdentifier, d.Get("target_id").(string), d.Id(), &nextgen.TriggersApiDeleteTriggerOpts{
		IfMatch: helpers.BuildField(d, "if_match"),
	})

	if err != nil {
		return helpers.HandleApiError(err, d)
	}

	return nil
}

func buildTrigger(d *schema.ResourceData) *nextgen.Trigger {
	return &nextgen.Trigger{
		Identifier:        d.Get("identifier").(string),
		Name:              d.Get("name").(string),
		OrgIdentifier:     d.Get("org_id").(string),
		ProjectIdentifier: d.Get("project_id").(string),
		Yaml:              d.Get("yaml").(string),
	}
}

func readTriggers(d *schema.ResourceData, trigger *nextgen.NgTriggerResponse) {
	d.SetId(trigger.Identifier)
	d.Set("identifier", trigger.Identifier)
	d.Set("name", trigger.Name)
	d.Set("org_id", trigger.OrgIdentifier)
	d.Set("project_id", trigger.ProjectIdentifier)
	d.Set("target_id", trigger.TargetIdentifier)
}
