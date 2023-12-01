package manual_freeze

import (
	"context"
	"net/http"

	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceManualFreeze() *schema.Resource {
	resource := &schema.Resource{
		Description: "Resource for Manual Deployment Freeze Window.",

		ReadContext:   resourceManualFreezeRead,
		UpdateContext: resourceManualFreezeCreateOrUpdate,
		DeleteContext: resourceManualFreezeDelete,
		CreateContext: resourceManualFreezeCreateOrUpdate,
		Importer:      helpers.MultiLevelResourceImporter,

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
				Description:      "Yaml of the freeze",
				Type:             schema.TypeString,
				Required:         true,
				DiffSuppressFunc: helpers.YamlDiffSuppressFunction,
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
												"value": {
													Description: "Every n months recurrence",
													Type:        schema.TypeInt,
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

func resourceManualFreezeRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	id := d.Get("identifier").(string)

	resp, httpResp, err := c.FreezeCRUDApi.GetFreeze(ctx, c.AccountId, id, &nextgen.FreezeCRUDApiGetFreezeOpts{
		OrgIdentifier:     helpers.BuildField(d, "org_id"),
		ProjectIdentifier: helpers.BuildField(d, "project_id"),
	})

	if httpResp.StatusCode == 404 {
		d.SetId("")
		d.MarkNewResource()
		return nil
	}

	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	if resp.Data == nil {
		return nil
	}

	readFreezeResponse(d, resp.Data)

	return nil

}

func resourceManualFreezeCreateOrUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	var err error
	var resp nextgen.ResponseDtoFreezeResponse
	var httpResp *http.Response
	var yaml string

	id := d.Id()

	if attr, ok := d.GetOk("yaml"); ok {
		yaml = attr.(string)
	}

	if id == "" {
		resp, httpResp, err = c.FreezeCRUDApi.CreateFreeze(ctx, yaml, c.AccountId, &nextgen.FreezeCRUDApiCreateFreezeOpts{
			OrgIdentifier:     helpers.BuildField(d, "org_id"),
			ProjectIdentifier: helpers.BuildField(d, "project_id"),
		})
	} else {
		resp, httpResp, err = c.FreezeCRUDApi.UpdateFreeze(ctx, yaml, c.AccountId, d.Id(), &nextgen.FreezeCRUDApiUpdateFreezeOpts{
			OrgIdentifier:     helpers.BuildField(d, "org_id"),
			ProjectIdentifier: helpers.BuildField(d, "project_id"),
		})
	}

	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	// The create/update methods don't return the upcoming windows in the response, so we need to query for it again.
	respGet, httpResp, err := c.FreezeCRUDApi.GetFreeze(ctx, c.AccountId, resp.Data.Identifier, &nextgen.FreezeCRUDApiGetFreezeOpts{
		OrgIdentifier:     helpers.BuildField(d, "org_id"),
		ProjectIdentifier: helpers.BuildField(d, "project_id"),
	})

	readFreezeResponse(d, respGet.Data)

	return nil
}

func resourceManualFreezeDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	httpResp, err := c.FreezeCRUDApi.DeleteFreeze(ctx, c.AccountId, d.Id(), &nextgen.FreezeCRUDApiDeleteFreezeOpts{
		OrgIdentifier:     helpers.BuildField(d, "org_id"),
		ProjectIdentifier: helpers.BuildField(d, "project_id"),
	})

	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	return nil
}

func readFreezeResponse(d *schema.ResourceData, freezeResponse *nextgen.FreezeDetailedResponse) {
	d.SetId(freezeResponse.Identifier)
	d.Set("identifier", freezeResponse.Identifier)
	d.Set("name", freezeResponse.Name)
	d.Set("org_id", freezeResponse.OrgIdentifier)
	d.Set("project_id", freezeResponse.ProjectIdentifier)
	d.Set("account_id", freezeResponse.AccountId)
	d.Set("type", freezeResponse.Type_)
	d.Set("tags", helpers.FlattenTags(freezeResponse.Tags))
	d.Set("status", freezeResponse.Status)
	d.Set("scope", freezeResponse.FreezeScope)
	d.Set("yaml", freezeResponse.Yaml)
	d.Set("description", freezeResponse.Description)

	if freezeResponse.CurrentOrUpcomingWindow != nil {
		d.Set("current_or_upcoming_windows", []interface{}{
			map[string]interface{}{
				"start_time": freezeResponse.CurrentOrUpcomingWindow.StartTime,
				"end_time":   freezeResponse.CurrentOrUpcomingWindow.EndTime,
			},
		})
	} else {
		d.Set("current_or_upcoming_windows", nil)
	}
	d.Set("freeze_windows", expandFreezeWindows(freezeResponse.Windows))
}

func expandFreezeWindows(freezeWindows []nextgen.FreezeWindow) []interface{} {
	var result []interface{}
	for _, window := range freezeWindows {
		result = append(result, map[string]interface{}{
			"time_zone":  window.TimeZone,
			"start_time": window.StartTime,
			"duration":   window.Duration,
			"end_time":   window.EndTime,
			"recurrence": expandRecurrence(window),
		})
	}
	return result
}

func expandRecurrence(window nextgen.FreezeWindow) []interface{} {
	var result []interface{}
	if window.Recurrence != nil {
		result = append(result, map[string]interface{}{
			"type":            window.Recurrence.Type_,
			"recurrence_spec": expandRecurrenceSpec(*window.Recurrence),
		})
	}
	return result
}

func expandRecurrenceSpec(recurrence nextgen.Recurrence) []interface{} {
	var result []interface{}
	if recurrence.Spec != nil {
		result = append(result, map[string]interface{}{
			"value": recurrence.Spec.Value,
			"until": recurrence.Spec.Until,
		})
	}
	return result
}
