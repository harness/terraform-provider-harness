package manual_freeze

import (
	"context"

	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceManualFreeze() *schema.Resource {
	resource := &schema.Resource{
		Description: "DataSource for deployment freeze in harness.",

		ReadContext: dataSourceManualFreezeRead,

		Schema: map[string]*schema.Schema{
			"identifier": {
				Description: "Identifier of the freeze",
				Type:        schema.TypeString,
				Required:    true,
			},
			"description": {
				Description: "Description of the freeze",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"name": {
				Description: "Name of the freeze",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"org_id": {
				Description: "Organization identifier of the freeze",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"project_id": {
				Description: "Project identifier of the freeze",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"tags": {
				Description: "Tags associated with the freeze",
				Type:        schema.TypeSet,
				Computed:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"type": {
				Description: "Type of freeze",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"account_id": {
				Description: "Account Identifier of the freeze",
				Type:        schema.TypeString,
				Required:    true,
			},
			"status": {
				Description: "Status of the freeze",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"scope": {
				Description: "Scope of the freeze",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"yaml": {
				Description: "Yaml of the freeze",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"current_or_upcoming_windows": {
				Description: "Current or upcoming windows",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"start_time": {
							Description: "Start time of the freeze window",
							Type:        schema.TypeInt,
							Computed:    true,
						},
						"end_time": {
							Description: "End time of the freeze window",
							Type:        schema.TypeInt,
							Computed:    true,
						},
					}},
			},
			"freeze_windows": {
				Description: "Freeze windows in the freeze response",
				Type:        schema.TypeSet,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"time_zone": {
							Description: "Time zone of the freeze window",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"start_time": {
							Description: "Start Time of the freeze window",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"duration": {
							Description: "Duration of the freeze window",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"end_time": {
							Description: "End Time of the freeze window",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"recurrence": {
							Description: "Recurrence of the freeze window",
							Type:        schema.TypeList,
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"type": {
										Description: "Type of the recurrence",
										Type:        schema.TypeString,
										Computed:    true,
									},
									"recurrence_spec": {
										Description: "Used to filter resources on their attributes",
										Type:        schema.TypeList,
										Computed:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"until": {
													Description: "Time till which freeze window recurrs",
													Type:        schema.TypeString,
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
		},
	}

	return resource
}

func dataSourceManualFreezeRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	id := d.Get("identifier").(string)

	resp, httpResp, err := c.FreezeCRUDApi.GetFreeze(ctx, c.AccountId, id, &nextgen.FreezeCRUDApiGetFreezeOpts{
		OrgIdentifier:     helpers.BuildField(d, "org_id"),
		ProjectIdentifier: helpers.BuildField(d, "project_id"),
	})

	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	if resp.Data == nil {
		return nil
	}

	readFreezeResponse(d, resp.Data)

	return nil

}
