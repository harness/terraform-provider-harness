package ng

import (
	"context"
	"strings"

	"github.com/antihax/optional"
	sdk "github.com/harness/harness-go-sdk"
	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/internal/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourcePipeline() *schema.Resource {
	return &schema.Resource{
		Description: utils.GetNextgenDescription("Resource for creating a Harness pipeline."),

		ReadContext:   resourcePipelineRead,
		UpdateContext: resourcePipelineUpdate,
		DeleteContext: resourcePipelineDelete,
		CreateContext: resourcePipelineCreate,
		Importer: &schema.ResourceImporter{
			State: func(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
				// <org_id>/<project_id>
				parts := strings.Split(d.Id(), "/")
				d.Set("org_id", parts[0])
				d.SetId(parts[1])

				return []*schema.ResourceData{d}, nil
			},
		},

		Schema: map[string]*schema.Schema{
			"identifier": {
				Description: "Unique identifier of the pipeline.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"org_id": {
				Description: "Unique identifier of the organization.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"project_id": {
				Description: "Unique identifier of the project.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"pipeline_yaml": {
				Description: "YAML of the pipeline.",
				Type:        schema.TypeString,
				Required:	 true,
			},
		},
	}
}

func readPipeline(d *schema.ResourceData, PmsPipelineResponse *nextgen.PmsPipelineResponse) {
	d.Set("pipeline_yaml", PmsPipelineResponse.YamlPipeline)
}

func resourcePipelineRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*sdk.Session)

	id := d.Id()
	if id == "" {
		id = d.Get("identifier").(string)
	}

	resp, _, err := c.NGClient.PipelinesApi.GetPipeline(ctx, c.AccountId, d.Get("org_id").(string), d.Get("project_id").(string), id, &nextgen.PipelinesApiGetPipelineOpts{})
	if err != nil {
		return diag.Errorf(err.(nextgen.GenericSwaggerError).Error())
	}

	readPipeline(d, resp.Data)

	return nil
}

func resourcePipelineCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*sdk.Session)

	id := d.Id()
	if id == "" {
		id = d.Get("identifier").(string)
	}

	_, _, err := c.NGClient.PipelinesApi.PostPipelineV2(ctx, d.Get("pipeline_yaml").(string), c.AccountId, d.Get("org_id").(string), d.Get("project_id").(string), &nextgen.PipelinesApiPostPipelineV2Opts{})
	if err != nil {
		return diag.Errorf(err.(nextgen.GenericSwaggerError).Error())
	}

	d.SetId(id)

	return nil
}

func resourcePipelineUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*sdk.Session)

	id := d.Id()
	if id == "" {
		id = d.Get("identifier").(string)
	}

	_, _, err := c.NGClient.PipelinesApi.UpdatePipeline(ctx, d.Get("pipeline_yaml").(string), c.AccountId, d.Get("org_id").(string), d.Get("project_id").(string), id, &nextgen.PipelinesApiUpdatePipelineOpts{})
	if err != nil {
		return diag.Errorf(err.(nextgen.GenericSwaggerError).Error())
	}

	return nil
}

func resourcePipelineDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*sdk.Session)

	id := d.Id()
	if id == "" {
		id = d.Get("identifier").(string)
	}

	_, _, err := c.NGClient.PipelinesApi.DeletePipeline(ctx, c.AccountId, d.Get("org_id").(string), d.Get("project_id").(string), id, &nextgen.PipelinesApiDeletePipelineOpts{})
	if err != nil {
		return diag.Errorf(err.(nextgen.GenericSwaggerError).Error())
	}

	d.SetId("")

	return nil
}
