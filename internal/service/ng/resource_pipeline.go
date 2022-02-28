package ng

import (
	"context"
	"strings"

	"github.com/antihax/optional"
	sdk "github.com/harness-io/harness-go-sdk"
	"github.com/harness-io/harness-go-sdk/harness/nextgen"
	"github.com/harness-io/terraform-provider-harness/internal/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// This is the Pipeline Entity details defined in Harness
type Pipeline struct {
	// Organization Identifier for the Entity
	OrgIdentifier string `json:"orgIdentifier,omitempty"`
	// Pipeline Identifier for the Entity
	Identifier string `json:"identifier,omitempty"`
	// Project Identifier for the entity
	ProjectIdentifier string `json:"name,omitempty"`
	// YAML contents of the pipeline
	PipelineYAML string `json:"color,omitempty"`
}


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

func buildPipeline(d *schema.ResourceData) *Pipeline {
	return &Pipeline{
		Identifier:    d.Get("identifier").(string),
		OrgIdentifier: d.Get("org_id").(string),
		ProjectIdentifier: d.Get("project_id").(string),
		PipelineYAML:  d.Get("pipeline_yaml").(string),
	}
}

func readPipeline(d *schema.ResourceData, pipeline *Pipeline) {
	d.SetId(pipeline.Identifier)
	d.Set("identifier", pipeline.Identifier)
	d.Set("org_id", pipeline.OrgIdentifier)
	d.Set("project_id", pipeline.ProjectIdentifier)
	d.Set("pipeline_yaml", pipeline.PipelineYAML)
}

func resourcePipelineRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*sdk.Session)

	pipeline := buildPipeline(d)

	resp, _, err := c.NGClient.PipelinesApi.GetPipeline(ctx, c.AccountId, pipeline.OrgIdentifier, pipeline.ProjectIdentifier, pipeline.Identifier, &nextgen.PipelinesApiGetPipelineOpts{})
	if err != nil {
		return diag.Errorf(err.(nextgen.GenericSwaggerError).Error())
	}

	pipeline.PipelineYAML = resp.Data.YamlPipeline

	readPipeline(d, pipeline)

	return nil
}

func resourcePipelineCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*sdk.Session)

	pipeline := buildPipeline(d)

	_, _, err := c.NGClient.PipelinesApi.PostPipelineV2(ctx, pipeline.PipelineYAML, c.AccountId, pipeline.OrgIdentifier, pipeline.ProjectIdentifier, &nextgen.PipelinesApiPostPipelineV2Opts{})
	if err != nil {
		return diag.Errorf(err.(nextgen.GenericSwaggerError).Error())
	}

	readPipeline(d, pipeline)

	return nil
}

func resourcePipelineUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*sdk.Session)

	pipeline := buildPipeline(d)

	_, _, err := c.NGClient.PipelinesApi.UpdatePipeline(ctx, pipeline.PipelineYAML, c.AccountId, pipeline.OrgIdentifier, pipeline.ProjectIdentifier, pipeline.Identifier, &nextgen.PipelinesApiUpdatePipelineOpts{})
	if err != nil {
		return diag.Errorf(err.(nextgen.GenericSwaggerError).Error())
	}

	return nil
}

func resourcePipelineDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*sdk.Session)

	pipeline := buildPipeline(d)

	_, _, err := c.NGClient.PipelinesApi.DeletePipeline(ctx, c.AccountId, pipeline.OrgIdentifier, pipeline.ProjectIdentifier, pipeline.Identifier, &nextgen.PipelinesApiDeletePipelineOpts{})
	if err != nil {
		return diag.Errorf(err.(nextgen.GenericSwaggerError).Error())
	}

	d.SetId("")

	return nil
}
