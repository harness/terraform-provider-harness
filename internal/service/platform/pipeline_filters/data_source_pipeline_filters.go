package pipeline_filters

import (
	"context"
	"errors"
	"net/http"
	"regexp"

	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func DataSourcePipelineFilters() *schema.Resource {
	resource := &schema.Resource{
		Description: "Data source for retrieving a Harness Pipeline Filter.",

		ReadContext: dataSourcePipelineFiltersRead,

		Schema: map[string]*schema.Schema{
			"identifier": {
				Description: "Unique identifier of the resource.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"name": {
				Description: "Name of the Filter.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"type": {
				Description:  "Type of filter. Currently supported types are {PipelineSetup, PipelineExecution, Deployment, Template, EnvironmentGroup, Environment}.",
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
				Description: "Properties of the filter entity defined in Harness.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"filter_type": {
							Description: "Corresponding Entity of the filters. Currently supported types are {Connector, DelegateProfile, Delegate, PipelineSetup, PipelineExecution, Deployment, Audit, Template, EnvironmentGroup, FileStore, CCMRecommendation, Anomaly, Environment}.",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"tags": {
							Description: "Tags to associate with the resource. Tags should be in the form `name:value`.",
							Type:        schema.TypeSet,
							Computed:    true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"pipeline_tags": {
							Description: "Tags to associate with the pipeline. tags should be in the form of `{key:key1, value:key1value}`",
							Type:        schema.TypeList,
							Optional:    true,
							Computed:    true,
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
							Computed:    true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"name": {
							Description: "Name of the pipeline filter.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"description": {
							Description: "description of the pipline filter.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"module_properties": {
							Description: "module properties of the pipline filter.",
							Type:        schema.TypeList,
							Optional:    true,
							Computed:    true,
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ci": {
										Type:        schema.TypeList,
										Optional:    true,
										Computed:    true,
										Description: "CI related properties to be filtered on.",
										MaxItems:    1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"build_type": {
													Description: "Build type of the pipeline. Possible values: branch.",
													Type:        schema.TypeString,
													Optional:    true,
													Computed:    true,
												},
												"branch": {
													Description: "Branch which was used while building.",
													Type:        schema.TypeString,
													Optional:    true,
													Computed:    true,
												},
												"tag": {
													Description: "Tags to associate with the CI pipeline resource.",
													Type:        schema.TypeString,
													Optional:    true,
													Computed:    true,
												},
												"repo_names": {
													Description: "name of the repository used in the pipeline.",
													Type:        schema.TypeString,
													Optional:    true,
													Computed:    true,
												},
												"ci_execution_info": {
													Description: "CI execution info for the pipeline.",
													Type:        schema.TypeList,
													Optional:    true,
													MaxItems:    1,
													Computed:    true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"event": {
																Description: "Event for the ci execution, Possible values: pullRequest.",
																Type:        schema.TypeString,
																Optional:    true,
																Computed:    true,
															},
															"pull_request": {
																Description: "The pull request details of the CI pipeline.",
																Type:        schema.TypeList,
																Optional:    true,
																Computed:    true,
																MaxItems:    1,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"source_branch": {
																			Description: "Source branch of the pull request.",
																			Type:        schema.TypeString,
																			Optional:    true,
																			Computed:    true,
																		},
																		"target_branch": {
																			Description: "Target branch of the pull request.",
																			Type:        schema.TypeString,
																			Optional:    true,
																			Computed:    true,
																		},
																	},
																},
															},
														},
													},
												},
											},
										},
									},
									"cd": {
										Description: "CD related properties to be filtered on.",
										Type:        schema.TypeList,
										Optional:    true,
										Computed:    true,
										MaxItems:    1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"deployment_types": {
													Description: "Deployment type of the CD pipeline, eg. Kubernetes",
													Type:        schema.TypeString,
													Optional:    true,
													Computed:    true,
												},
												"service_names": {
													Description: "Service names of the CD pipeline.",
													Type:        schema.TypeSet,
													Optional:    true,
													Computed:    true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"environment_names": {
													Description: "Environment names of the CD pipeline.",
													Type:        schema.TypeSet,
													Optional:    true,
													Computed:    true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"artifact_display_names": {
													Description: "Artifact display names of the CD pipeline.",
													Type:        schema.TypeSet,
													Optional:    true,
													Computed:    true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
			"filter_visibility": {
				Description: "This indicates visibility of filter. By default, everyone can view this filter.",
				Type:        schema.TypeString,
				Computed:    true,
			},
		},
	}

	return resource
}

func dataSourcePipelineFiltersRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	var filter *nextgen.PipelineFilter
	var err error
	var httpResp *http.Response

	id := d.Get("identifier").(string)
	type_ := d.Get("type").(string)

	if id != "" {
		var resp nextgen.ResponseDtoPipelineFilter
		resp, httpResp, err = c.FilterApi.PipelinegetFilter(ctx, c.AccountId, id, type_, &nextgen.FilterApiPipelinegetFilterOpts{
			OrgIdentifier:     helpers.BuildField(d, "org_id"),
			ProjectIdentifier: helpers.BuildField(d, "project_id"),
		})
		filter = resp.Data
	} else {
		return diag.FromErr(errors.New(" identifier  must be specified"))
	}

	if err != nil {
		return helpers.HandleReadApiError(err, d, httpResp)
	}

	if filter == nil {
		d.SetId("")
		d.MarkNewResource()
		return nil
	}

	readPipelineFilter(d, filter)

	return nil
}
