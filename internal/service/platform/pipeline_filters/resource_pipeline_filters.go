package pipeline_filters

import (
	"context"
	"net/http"
	"regexp"

	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func ResourcePipelineFilters() *schema.Resource {
	resource := &schema.Resource{
		Description: "Resource for creating Harness Pipeline Filters.",

		ReadContext:   resourcePipelineFiltersRead,
		UpdateContext: resourcePipelineFiltersCreateOrUpdate,
		DeleteContext: resourcePipelineFiltersDelete,
		CreateContext: resourcePipelineFiltersCreateOrUpdate,
		Importer:      helpers.MultiLevelFilterImporter,

		Schema: map[string]*schema.Schema{
			"identifier": {
				Description: "Unique identifier of the resource.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"name": {
				Description: "Name of the pipeline filters.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"type": {
				Description:  "Type of pipeline filters. Currently supported types are {PipelineSetup, PipelineExecution, Deployment, Template, EnvironmentGroup, Environment}",
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringInSlice([]string{"PipelineSetup", "PipelineExecution", "Deployment", "Template", "EnvironmentGroup", "Environment"}, false),
			},
			"org_id": {
				Description: "Organization Identifier for the Entity.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"project_id": {
				Description: "Project Identifier for the Entity.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"filter_properties": {
				Description: "Properties of the filters entity defined in Harness.",
				Type:        schema.TypeList,
				MaxItems:    1,
				Required:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"filter_type": {
							Description:  "Corresponding Entity of the filters. Currently supported types are {Connector, DelegateProfile, Delegate, PipelineSetup, PipelineExecution, Deployment, Audit, Template, EnvironmentGroup, FileStore, CCMRecommendation, Anomaly, Environment}.",
							Type:         schema.TypeString,
							Required:     true,
							ValidateFunc: validation.StringInSlice([]string{"Connector", "DelegateProfile", "Delegate", "PipelineSetup", "PipelineExecution", "Deployment", "Audit", "Template", "EnvironmentGroup", "FileStore", "CCMRecommendation", "Anomaly", "Environment"}, false),
						},
						"tags": {
							Description: "Tags to associate with the resource. Tags should be in the form `name:value`.",
							Type:        schema.TypeSet,
							Optional:    true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"pipeline_tags": {
							Description: "Tags to associate with the pipeline. tags should be in the form of `{key:key1, value:key1value}`",
							Type:        schema.TypeList,
							Optional:    true,
							Elem: &schema.Schema{
								Type:             schema.TypeMap,
								ValidateDiagFunc: validation.MapKeyMatch(regexp.MustCompile("^key$|^value$"), "Please provide valid pipeline tags. valid values: key and value."),
								Elem: &schema.Schema{
									Type: schema.TypeString,
								},
							},
						},
						"pipeline_identifiers": {
							Description: "Pipeline identifiers to filter on.",
							Type:        schema.TypeList,
							Optional:    true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"name": {
							Description: "Name of the pipeline filter.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"description": {
							Description: "description of the pipline filter.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"module_properties": {
							Description: "module properties of the pipline filter.",
							Type:        schema.TypeMap,
							Optional:    true,
							Elem: &schema.Schema{
								Type: schema.TypeMap,
							},
						},
					},
				},
			},
			"filter_visibility": {
				Description:  "This indicates visibility of filters. By default, everyone can view this filter.",
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validation.StringInSlice([]string{"EveryOne", "OnlyCreator"}, false),
			},
		},
	}

	return resource
}

func resourcePipelineFiltersRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	id := d.Id()

	type_ := d.Get("type").(string)
	resp, httpResp, err := c.FilterApi.PipelinegetFilter(ctx, c.AccountId, id, type_, &nextgen.FilterApiPipelinegetFilterOpts{
		OrgIdentifier:     helpers.BuildField(d, "org_id"),
		ProjectIdentifier: helpers.BuildField(d, "project_id"),
	})

	if err != nil {
		return helpers.HandleReadApiError(err, d, httpResp)
	}

	if resp.Data == nil {
		d.SetId("")
		d.MarkNewResource()
		return nil
	}

	readPipelineFilter(d, resp.Data)

	return nil
}

func resourcePipelineFiltersCreateOrUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	var err error
	var resp nextgen.ResponseDtoPipelineFilter
	var httpResp *http.Response

	id := d.Id()
	filter := buildPipelineFilter(d)

	if id == "" {
		resp, httpResp, err = c.FilterApi.PipelinepostFilter(ctx, *filter, c.AccountId)
	} else {
		resp, httpResp, err = c.FilterApi.PipelineupdateFilter(ctx, *filter, c.AccountId)
	}

	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	readPipelineFilter(d, resp.Data)

	return nil
}

func resourcePipelineFiltersDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	type_ := d.Get("type").(string)

	_, httpResp, err := c.FilterApi.PipelinedeleteFilter(ctx, c.AccountId, d.Id(), type_, &nextgen.FilterApiPipelinedeleteFilterOpts{
		OrgIdentifier:     helpers.BuildField(d, "org_id"),
		ProjectIdentifier: helpers.BuildField(d, "project_id"),
	})

	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	return nil
}

func buildPipelineFilter(d *schema.ResourceData) *nextgen.PipelineFilter {
	filter := &nextgen.PipelineFilter{
		FilterProperties: &nextgen.PipelineFilterProperties{},
	}

	if attr, ok := d.GetOk("org_id"); ok {
		filter.OrgIdentifier = attr.(string)
	}

	if attr, ok := d.GetOk("project_id"); ok {
		filter.ProjectIdentifier = attr.(string)
	}

	if attr, ok := d.GetOk("filter_visibility"); ok {
		filter.FilterVisibility = attr.(string)
	}

	if attr, ok := d.GetOk("name"); ok {
		filter.Name = attr.(string)
	}

	if attr, ok := d.GetOk("identifier"); ok {
		filter.Identifier = attr.(string)
	}

	if attr, ok := d.GetOk("filter_properties"); ok {
		filterProperties := attr.([]interface{})[0].(map[string]interface{})
		if attr, ok := filterProperties["filter_type"]; ok {
			filter.FilterProperties.FilterType = attr.(string)
		}

		if attr := filterProperties["tags"].(*schema.Set).List(); len(attr) > 0 {
			filter.FilterProperties.Tags = helpers.ExpandTags(attr)
		}

		if attr := filterProperties["pipeline_tags"].([]interface{}); len(attr) > 0 {
			var hPipelineTags []nextgen.NgTag
			for _, v := range attr {
				if v != nil {
					var vMap = v.(map[string]interface{})
					key := vMap["key"].(string)
					value := vMap["value"].(string)
					if key != "" && value != "" {
						hPipelineTag := nextgen.NgTag{
							Key:   key,
							Value: value,
						}
						hPipelineTags = append(hPipelineTags, hPipelineTag)
					}
				}
			}
			filter.FilterProperties.PipelineTags = hPipelineTags

		}

		if attr := filterProperties["pipeline_identifiers"].([]interface{}); len(attr) > 0 {
			pipelineIdentifiers := helpers.ExpandField(attr)
			filter.FilterProperties.PipelineIdentifiers = pipelineIdentifiers
		}

		if attr, ok := filterProperties["name"]; ok {
			name := attr.(string)
			filter.FilterProperties.Name = name
		}

		if attr, ok := filterProperties["description"]; ok {
			description := attr.(string)
			filter.FilterProperties.Description = description
		}

		if attr, ok := filterProperties["module_properties"]; ok {
			moduleProperties := attr.(map[string]interface{})
			filter.FilterProperties.ModuleProperties = moduleProperties
		}
	}

	return filter
}

func readPipelineFilter(d *schema.ResourceData, filter *nextgen.PipelineFilter) {
	d.SetId(filter.Identifier)
	d.Set("identifier", filter.Identifier)
	d.Set("org_id", filter.OrgIdentifier)
	d.Set("project_id", filter.ProjectIdentifier)
	d.Set("name", filter.Name)
	d.Set("type", filter.FilterProperties.FilterType)
	d.Set("filter_visibility", filter.FilterVisibility)

	filterProperties := make(map[string]interface{})
	filterProperties["filter_type"] = filter.FilterProperties.FilterType
	filterProperties["tags"] = helpers.FlattenTags(filter.FilterProperties.Tags)
	var pipelineTags []interface{}
	for _, tagV := range filter.FilterProperties.PipelineTags {
		pipelineTag := make(map[string]interface{})
		key := tagV.Key
		value := tagV.Value
		if key != "" && value != "" {
			pipelineTag["key"] = key
			pipelineTag["value"] = value
			pipelineTags = append(pipelineTags, pipelineTag)
		}
	}
	filterProperties["pipeline_tags"] = pipelineTags

	if filter.FilterProperties.Name != "" {
		filterProperties["name"] = filter.FilterProperties.Name
	}
	if filter.FilterProperties.Description != "" {
		filterProperties["description"] = filter.FilterProperties.Description
	}
	if filter.FilterProperties.PipelineIdentifiers != nil && len(filter.FilterProperties.PipelineIdentifiers) > 0 {
		filterProperties["pipeline_identifiers"] = filter.FilterProperties.PipelineIdentifiers
	}
	if filter.FilterProperties.ModuleProperties != nil && len(filter.FilterProperties.ModuleProperties) > 0 {
		filterProperties["module_properties"] = filter.FilterProperties.ModuleProperties
	}

	d.Set("filter_properties", []interface{}{filterProperties})
}
