package input_set

import (
	"context"

	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceInputSet() *schema.Resource {
	resource := &schema.Resource{
		Description: "Resource for creating a Harness Resource Group",

		ReadContext:   resourceInputSetRead,
		UpdateContext: resourceInputSetCreateOrUpdate,
		CreateContext: resourceInputSetCreateOrUpdate,
		DeleteContext: resourceInputSetDelete,
		Importer:      helpers.PipelineResourceImporter,

		Schema: map[string]*schema.Schema{
			"pipeline_id": {
				Description: "Identifier of the pipeline",
				Type:        schema.TypeString,
				Required:    true,
			},
			"yaml": {
				Description: "Input Set YAML",
				Type:        schema.TypeString,
				Required:    true,
			},
		},
	}
	helpers.SetMultiLevelResourceSchema(resource.Schema)

	return resource
}

func resourceInputSetRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	id := d.Get("identifier").(string)

	orgId := d.Get("org_id").(string)

	projectId := d.Get("project_id").(string)

	pipelineId := d.Get("pipeline_id").(string)

	resp, _, err := c.InputSetsApi.GetInputSet(ctx, id, c.AccountId, orgId, projectId, pipelineId, &nextgen.InputSetsApiGetInputSetOpts{})

	if err != nil {
		return helpers.HandleApiError(err, d)
	}

	if resp.Data == nil {
		return nil
	}

	readInputSet(d, resp.Data)

	return nil

}

func resourceInputSetCreateOrUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	var err error
	var resp nextgen.ResponseDtoInputSetResponse

	id := d.Id()
	inputSet := buildInputSet(d)
	orgIdentifier := d.Get("org_id").(string)
	projectIdentifier := d.Get("project_id").(string)
	pipelineIdentifier := d.Get("pipeline_id").(string)

	if id == "" {
		resp, _, err = c.InputSetsApi.PostInputSet(ctx, inputSet.InputSetYaml, c.AccountId, orgIdentifier, projectIdentifier, pipelineIdentifier,
			&nextgen.InputSetsApiPostInputSetOpts{})
	} else {
		resp, _, err = c.InputSetsApi.PutInputSet(ctx, inputSet.InputSetYaml, c.AccountId, orgIdentifier, projectIdentifier, pipelineIdentifier, d.Id(),
			&nextgen.InputSetsApiPutInputSetOpts{})
	}

	if err != nil {
		return helpers.HandleApiError(err, d)
	}

	readInputSet(d, resp.Data)

	return nil
}

func resourceInputSetDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)
	orgIdentifier := helpers.BuildField(d, "org_id").Value()
	projectIdentifier := helpers.BuildField(d, "project_id").Value()
	pipelineIdentifier := helpers.BuildField(d, "pipeline_id").Value()

	_, _, err := c.InputSetsApi.DeleteInputSet(ctx, d.Id(), c.AccountId, orgIdentifier, projectIdentifier, pipelineIdentifier,
		&nextgen.InputSetsApiDeleteInputSetOpts{})

	if err != nil {
		return diag.Errorf(err.(nextgen.GenericSwaggerError).Error())
	}

	return nil
}

func buildInputSet(d *schema.ResourceData) *nextgen.InputSetResponse {
	inputSet := &nextgen.InputSetResponse{}

	if attr, ok := d.GetOk("account_id"); ok {
		inputSet.AccountId = attr.(string)
	}

	if attr, ok := d.GetOk("org_id"); ok {
		inputSet.OrgIdentifier = attr.(string)
	}

	if attr, ok := d.GetOk("project_id"); ok {
		inputSet.ProjectIdentifier = attr.(string)
	}

	if attr, ok := d.GetOk("pipeline_id"); ok {
		inputSet.PipelineIdentifier = attr.(string)
	}

	if attr, ok := d.GetOk("identifier"); ok {
		inputSet.Identifier = attr.(string)
	}

	if attr, ok := d.GetOk("yaml"); ok {
		inputSet.InputSetYaml = attr.(string)
	}

	if attr, ok := d.GetOk("name"); ok {
		inputSet.Name = attr.(string)
	}

	if attr, ok := d.GetOk("description"); ok {
		inputSet.Description = attr.(string)
	}

	if attr := d.Get("tags").(*schema.Set).List(); len(attr) > 0 {
		inputSet.Tags = helpers.ExpandTags(attr)
	}

	return inputSet
}

func readInputSet(d *schema.ResourceData, inputSet *nextgen.InputSetResponse) {
	d.SetId(inputSet.Identifier)
	d.Set("identifier", inputSet.Identifier)
	d.Set("name", inputSet.Name)
	d.Set("description", inputSet.Description)
	d.Set("tags", helpers.FlattenTags(inputSet.Tags))
	d.Set("org_id", inputSet.OrgIdentifier)
	d.Set("project_id", inputSet.ProjectIdentifier)
	d.Set("pipeline_id", inputSet.PipelineIdentifier)
	d.Set("yaml", inputSet.InputSetYaml)
}
