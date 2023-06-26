package ccm_perspective

import (
	"context"

	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceCCMPerspective() *schema.Resource {
	resource := &schema.Resource{
		Description: "Data source for retrieving a Harness CCM Filter.",

		ReadContext: dataSourceCCMPerspectiveRead,
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

func dataSourceCCMPerspectiveRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	id := d.Get("identifier").(string)

	resp, httpResp, err := c.CloudCostPerspectivesApi.GetPerspective(ctx, c.AccountId, id)

	if err != nil {
		return helpers.HandleReadApiError(err, d, httpResp)
	}

	if resp.Data == nil {
		d.SetId("")
		d.MarkNewResource()
		return nil
	}

	readCCMPerspective(d, resp.Data, false)
	return nil
}
