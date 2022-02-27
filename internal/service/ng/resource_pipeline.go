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

		//accountIdentifier string, orgIdentifier string, projectIdentifier string, pipelineIdentifier string, localVarOptionals *
		//func (a *PipelinesApiService) PostPipelineV2(ctx context.Context, body string, accountIdentifier string, orgIdentifier string, projectIdentifier string, localVarOptionals *PipelinesApiPostPipelineV2Opts) (ResponseDtoPipelineSaveResponse, *http.Response, error) {
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
				//Optional:    true,
				Computed:    true,
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

//func (a *PipelinesApiService) GetPipeline(ctx context.Context, accountIdentifier string, orgIdentifier string, projectIdentifier string, pipelineIdentifier string, localVarOptionals *PipelinesApiGetPipelineOpts) (ResponseDtopmsPipelineResponse, *http.Response, error) {
func resourcePipelineRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*sdk.Session)

	id := d.Id()
	if id == "" {
		id = d.Get("identifier").(string)
	}

	orgId := d.Get("org_id").(string)
	projectId := d.Get("project_id").(string)

	resp, _, err := c.NGClient.PipelineApi.GetPipeline(ctx, c.AccountId, orgIdentifier, projectIdentifier, id, &nextgen.PipelinesApiGetPipelineOpts{})
	if err != nil {
		return diag.Errorf(err.(nextgen.GenericSwaggerError).Error())
	}

	readPipeline(d, resp.Data.Pipeline)

	return nil
}

//func (a *PipelinesApiService) PostPipelineV2(ctx context.Context, body string, accountIdentifier string, orgIdentifier string, projectIdentifier string, localVarOptionals *PipelinesApiPostPipelineV2Opts) (ResponseDtoPipelineSaveResponse, *http.Response, error) {
func resourcePipelineCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*sdk.Session)

	pipeline := buildPipeline(d)

	pipelineYAML := d.Get("pipeline_yaml").(string)
	orgId := d.Get("org_id").(string)
	projectId := d.Get("project_id").(string)


	resp, _, err := c.NGClient.PipelineApi.PostPipelineV2(ctx, pipelineYAML, c.AccountId, orgId, projectId, &nextgen.PipelinesApiPostPipelineV2Opts{})
	if err != nil {
		return diag.Errorf(err.(nextgen.GenericSwaggerError).Error())
	}

	readPipeline(d, resp.Data.Pipeline)

	return nil
}

//func (a *PipelinesApiService) UpdatePipeline(ctx context.Context, body string, accountIdentifier string, orgIdentifier string, projectIdentifier string, pipelineIdentifier string, localVarOptionals *PipelinesApiUpdatePipelineOpts) (ResponseDtoString, *http.Response, error) {
func resourcePipelineUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*sdk.Session)

	pipeline := buildPipeline(d)

	pipelineYAML := d.Get("pipeline_yaml").(string)
	orgId := d.Get("org_id").(string)
	projectId := d.Get("project_id").(string)

	resp, _, err := c.NGClient.PipelineApi.UpdatePipeline(ctx, pipelineYAML, c.AccountId, orgId, projectId, id, &nextgen.PipelinesApiUpdatePipelineOpts{})
	if err != nil {
		return diag.Errorf(err.(nextgen.GenericSwaggerError).Error())
	}

	readPipeline(d, resp.Data.Pipeline)

	return nil
}

//func (a *PipelinesApiService) DeletePipeline(ctx context.Context, accountIdentifier string, orgIdentifier string, projectIdentifier string, pipelineIdentifier string, localVarOptionals *PipelinesApiDeletePipelineOpts) (ResponseDtoBoolean, *http.Response, error) {
func resourcePipelineDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*sdk.Session)

	id := d.Id()
	if id == "" {
		id = d.Get("identifier").(string)
	}
	orgId := d.Get("org_id").(string)
	projectId := d.Get("project_id").(string)

	//_, _, err := c.NGClient.PipelineApi.DeletePipeline(ctx, c.AccountId, orgId, projectId, id, &nextgen.PipelinesApiDeletePipelineOpts{OrgIdentifier: optional.NewString(d.Get("org_id").(string))})
	_, _, err := c.NGClient.PipelineApi.DeletePipeline(ctx, c.AccountId, orgId, projectId, id, &nextgen.PipelinesApiDeletePipelineOpts{})
	if err != nil {
		return diag.Errorf(err.(nextgen.GenericSwaggerError).Error())
	}

	return nil
}
