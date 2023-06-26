package ccm_perspective

import (
	"context"
	"net/http"

	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceCCMPerspective() *schema.Resource {
	resource := &schema.Resource{
		Description: "Resource for creating a Harness pipeline.",

		ReadContext:   resourceCCMPerspectiveRead,
		UpdateContext: resourceCCMPerspectiveCreateOrUpdate,
		DeleteContext: resourceCCMPerspectiveDelete,
		CreateContext: resourceCCMPerspectiveCreateOrUpdate,
		Importer:      helpers.MultiLevelResourceImporter,

		Schema: map[string]*schema.Schema{
			"account_id": {
				Description: "Account identifier of the CCM Perspective.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"identifier": {
				Description: "Unique identifier of the resource.",
				Type:        schema.TypeString,
				Computed:    true,
				Optional:    true,
			},
			"clone": {
				Description: "Can clone the CCM Perspective.",
				Type:        schema.TypeBool,
				Required:    true,
			},
			"name": {
				Description: "Name of the CCM Perspective.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"folder_id": {
				Description: "Folder Id of the CCM Perspective.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"view_version": {
				Description: "View version of the CCM Perspective.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"view_time_range": {
				Description: "The time range of the CCM Perspective.",
				Type:        schema.TypeList,
				Computed:    true,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"view_time_range_type": {
							Description: "The time range of the CCM Perspective.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"start_time": {
							Description: "The start time the CCM Perspective.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
						"end_time": {
							Description: "The end time the CCM Perspective.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
					}},
			},
			"view_rules": {
				Description: "The rules of the CCM Perspective.",
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"view_conditions": {
							Description: "The conditions of the CCM Perspective.",
							Type:        schema.TypeList,
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"type": {
										Description: "The view type of the CCM Perspective.",
										Type:        schema.TypeString,
										Optional:    true,
									},
									"view_field": {
										Description: "The view field of the CCM Perspective.",
										Type:        schema.TypeList,
										Optional:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"field_id": {
													Description: "The field id of the CCM Perspective.",
													Type:        schema.TypeString,
													Optional:    true,
												},
												"field_name": {
													Description: "The field name of the CCM Perspective.",
													Type:        schema.TypeString,
													Optional:    true,
												},
												"view_field_identifier": {
													Description: "The view field identifier of the CCM Perspective.",
													Type:        schema.TypeString,
													Optional:    true,
												},
												"identifier_name": {
													Description: "The view field identifier name of the CCM Perspective.",
													Type:        schema.TypeString,
													Optional:    true,
												},
											},
										},
									},
									"view_operator": {
										Description: "The view operator of the CCM Perspective.",
										Type:        schema.TypeString,
										Optional:    true,
									},
									"values": {
										Description: "The rules of the CCM Perspective.",
										Type:        schema.TypeSet,
										Optional:    true,
										Elem:        &schema.Schema{Type: schema.TypeString},
									},
								},
							},
						},
					},
				},
			},
			"data_sources": {
				Description: "The data sources of the CCM Perspective.",
				Type:        schema.TypeSet,
				Optional:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"view_visualization": {
				Description: "The view visualization of the CCM Perspective.",
				Type:        schema.TypeList,
				Computed:    true,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"granularity": {
							Description: "The view granularity of the CCM Perspective.",
							Type:        schema.TypeString,
							Computed:    true,
							Optional:    true,
						},
						"group_by": {
							Description: "The group by view of the CCM Perspective.",
							Type:        schema.TypeList,
							Computed:    true,
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"field_id": {
										Description: "The field id of the CCM Perspective.",
										Type:        schema.TypeString,
										Optional:    true,
									},
									"field_name": {
										Description: "The field name of the CCM Perspective.",
										Type:        schema.TypeString,
										Optional:    true,
									},
									"group_by_identifier": {
										Description: "The view field identifier of the CCM Perspective.",
										Type:        schema.TypeString,
										Optional:    true,
									},
									"identifier_name": {
										Description: "The view field identifier name of the CCM Perspective.",
										Type:        schema.TypeString,
										Optional:    true,
									},
								},
							},
						},
						"chart_type": {
							Description: "The chart type of the CCM Perspective.",
							Type:        schema.TypeString,
							Computed:    true,
							Optional:    true,
						},
					},
				},
			},
			"view_preferences": {
				Description: "The view preferences of the CCM Perspective.",
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"include_others": {
							Description: "Include others flag of the CCM Perspective.",
							Type:        schema.TypeBool,
							Optional:    true,
						},
						"include_unallocated_cost": {
							Description: "Include unallocated cost flag of the CCM Perspective.",
							Type:        schema.TypeBool,
							Optional:    true,
						},
					},
				},
			},
			"view_type": {
				Description: "View type of the CCM Perspective.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"view_state": {
				Description: "View state of the CCM Perspective.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"total_cost": {
				Description: "Total cost of the CCM Perspective.",
				Type:        schema.TypeFloat,
				Optional:    true,
			},
			"created_at": {
				Description: "The time the CCM Perspective created.",
				Type:        schema.TypeInt,
				Computed:    true,
			},
			"last_updated_at": {
				Description: "The time the CCM Perspective updated.",
				Type:        schema.TypeInt,
				Computed:    true,
			},
		},
	}
	return resource
}

func resourceCCMPerspectiveRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	id := d.Get("identifier").(string)
	var clone bool
	if attr, ok := d.GetOk("clone"); ok {
		clone = attr.(bool)
	}

	resp, httpResp, err := c.CloudCostPerspectivesApi.GetPerspective(ctx, c.AccountId, id)

	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	if resp.Data == nil {
		return nil
	}

	readCCMPerspective(d, resp.Data, clone)
	return nil
}

func resourceCCMPerspectiveCreateOrUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	var err error
	var clone bool
	var resp nextgen.ResponseDtoceView
	var httpResp *http.Response
	id := d.Id()
	if attr, ok := d.GetOk("clone"); ok {
		clone = attr.(bool)
	}
	ceView := buildceView(d)

	if id == "" {
		resp, httpResp, err = c.CloudCostPerspectivesApi.CreatePerspective(ctx, *ceView, c.AccountId, clone)
	} else {
		resp, httpResp, err = c.CloudCostPerspectivesApi.UpdatePerspective(ctx, *ceView, c.AccountId)
	}

	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	readCCMPerspective(d, resp.Data, clone)

	return nil
}

func resourceCCMPerspectiveDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	_, httpResp, err := c.CloudCostPerspectivesApi.DeletePerspective(ctx, c.AccountId, d.Id())

	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	return nil
}

func readCCMPerspective(d *schema.ResourceData, ceView *nextgen.CeView, clone bool) diag.Diagnostics {
	d.SetId(ceView.Uuid)
	d.Set("identifier", ceView.Uuid)
	d.Set("name", ceView.Name)
	d.Set("account_id", ceView.AccountId)
	d.Set("view_version", ceView.ViewVersion)
	d.Set("clone", clone)
	if ceView.ViewTimeRange != nil {
		viewTimeRangeList := []interface{}{}
		viewTimeRange := map[string]interface{}{}
		viewTimeRange["view_time_range_type"] = ceView.ViewTimeRange.ViewTimeRangeType
		viewTimeRange["start_time"] = ceView.ViewTimeRange.StartTime
		viewTimeRange["end_time"] = ceView.ViewTimeRange.EndTime
		viewTimeRangeList = append(viewTimeRangeList, viewTimeRange)
		d.Set("view_time_range", viewTimeRangeList)
	}

	d.Set("view_rules", flattenViewRules(ceView.ViewRules))
	d.Set("view_visualization", []interface{}{
		map[string]interface{}{
			"granularity": ceView.ViewVisualization.Granularity,
			"chart_type":  ceView.ViewVisualization.ChartType,
			"group_by":    expandGroupBy(ceView.ViewVisualization.GroupBy),
		},
	})

	d.Set("view_type", ceView.ViewType)
	d.Set("view_state", ceView.ViewState)
	d.Set("total_cost", ceView.TotalCost)
	d.Set("created_at", ceView.CreatedAt)
	d.Set("last_updated_at", ceView.LastUpdatedAt)
	return nil
}

func expandGroupBy(viewField *nextgen.ViewField) []interface{} {
	var result []interface{}
	if viewField != nil {
		result = append(result, map[string]interface{}{
			"field_id":            viewField.FieldId,
			"field_name":          viewField.FieldName,
			"group_by_identifier": viewField.Identifier,
			"identifier_name":     viewField.IdentifierName,
		})
	}

	return result
}

func flattenViewRules(viewRules []nextgen.ViewRule) []interface{} {
	var result []interface{}

	for _, viewRule := range viewRules {
		result = append(result, map[string]interface{}{
			"view_conditions": flattenViewConditions(viewRule.ViewConditions),
		})
	}
	return result
}

func flattenViewConditions(viewConditions []nextgen.ViewCondition) []interface{} {
	var result []interface{}
	for _, viewCondition := range viewConditions {
		result = append(result, map[string]interface{}{
			"type": viewCondition.Type_,
		})
	}
	return result
}

func buildceView(d *schema.ResourceData) *nextgen.CeView {
	var ceviewObject nextgen.CeView

	if attr, ok := d.GetOk("name"); ok {
		ceviewObject.Name = attr.(string)
	}

	if attr, ok := d.GetOk("view_version"); ok {
		ceviewObject.ViewVersion = attr.(string)
	}

	if attr, ok := d.GetOk("view_time_range_type"); ok {
		var ViewTimeRangeObject nextgen.ViewTimeRange
		if attr != nil && len(attr.([]interface{})) > 0 {
			viewtimerangeinobject := attr.([]interface{})[0].(map[string]interface{})
			if viewtimerangeinobject["view_time_range_type"] != nil && len(viewtimerangeinobject["view_time_range_type"].(string)) > 0 {
				ViewTimeRangeObject.ViewTimeRangeType = viewtimerangeinobject["view_time_range_type"].(string)
			}
			if viewtimerangeinobject["start_time"] != nil {
				ViewTimeRangeObject.StartTime = int64(viewtimerangeinobject["start_time"].(int))
			}
			if viewtimerangeinobject["end_time"] != nil {
				ViewTimeRangeObject.EndTime = int64(viewtimerangeinobject["end_time"].(int))
			}
		}
		ceviewObject.ViewTimeRange = &ViewTimeRangeObject
	}

	if attr, ok := d.GetOk("data_sources"); ok {
		var dataSourcesObject []string
		datasourcesinobject := attr.([]interface{})[0].(map[string]interface{})
		for _, v := range datasourcesinobject["data_sources"].([]interface{}) {
			dataSourcesObject = append(dataSourcesObject, v.(string))
		}
		ceviewObject.DataSources = dataSourcesObject
	}

	if attr, ok := d.GetOk("view_type"); ok {
		ceviewObject.ViewType = attr.(string)
	}

	if attr, ok := d.GetOk("view_state"); ok {
		ceviewObject.ViewState = attr.(string)
	}

	if attr, ok := d.GetOk("total_cost"); ok {
		ceviewObject.TotalCost = attr.(float64)
	}

	if attr, ok := d.GetOk("created_at"); ok {
		ceviewObject.CreatedAt = attr.(int64)
	}

	if attr, ok := d.GetOk("last_updated_at"); ok {
		ceviewObject.LastUpdatedAt = attr.(int64)
	}

	return &ceviewObject
}
