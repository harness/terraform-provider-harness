package schedule

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type scheduleType string

var (
	uptimeSchedule   scheduleType = "uptime"
	downtimeSchedule scheduleType = "downtime"
)

func dateValidationFunc(i interface{}, p cty.Path) diag.Diagnostics {
	diags := diag.Diagnostics{}
	v, _ := i.(string)
	_, err := time.Parse(time.DateTime, v)
	if err != nil {
		d := diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Should be valid date-time in YYYY-MM-DD HH:mm:SS format. Eg 2006-01-02 15:04:05",
		}
		diags = append(diags, d)
	}
	return diags
}

func timeValidateFunc(i interface{}, p cty.Path) diag.Diagnostics {
	diags := diag.Diagnostics{}
	v, ok := i.(string)
	if !ok {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Non empty value is mandatory",
		})
		return diags
	}
	v = strings.TrimSpace(v)
	parts := strings.Split(v, ":")
	if len(parts) != 2 {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Value should be in HH:MM format. Eg : 13:15 for 01:15pm",
		})
		return diags
	}
	hh, err := strconv.ParseInt(parts[0], 10, 64)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Value should be in HH:MM format. Eg : 13:15 for 01:15pm",
		})
		return diags
	}
	if hh < 0 || hh > 24 {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Hour value should be between 0 and 24",
		})
	}
	mm, err := strconv.ParseInt(parts[0], 10, 64)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Value should be in HH:MM format. Eg : 13:15 for 01:15pm",
		})
		return diags
	}
	if mm < 0 || mm > 59 {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Minute value should be between 0 and 59",
		})
	}
	return diags
}

func ResourceVMRule() *schema.Resource {
	resource := &schema.Resource{
		Description: "Resource for creating a Harness Variables.",

		ReadContext:   resourceScheduleRead,
		UpdateContext: resourceScheduleUpdate,
		DeleteContext: resourceScheduleDelete,
		CreateContext: resourceScheduleCreate,
		Importer:      helpers.MultiLevelResourceImporter,
		Schema: map[string]*schema.Schema{
			"identifier": {
				Description: "Unique identifier of the resource",
				Type:        schema.TypeFloat,
				Computed:    true,
			},
			"name": {
				Description: "Name of the schedule",
				Type:        schema.TypeString,
				Required:    true,
			},
			"schedule_type": {
				Description:  fmt.Sprintf("Type of the schedule. Valid values are `%s` and `%s`", uptimeSchedule, downtimeSchedule),
				Type:         schema.TypeString,
				Required:     true,
				ExactlyOneOf: []string{string(uptimeSchedule), string(downtimeSchedule)},
			},
			"time_zone": {
				Description: "Time zone in which schedule needs to be executed",
				Type:        schema.TypeString,
				Required:    true,
			},
			"time_period": {
				Description: "Time period in which schedule will be active. If specified along with periodicity, this will act as the boundary of periodicity. Otherwise schedule action will be triggered at `start` time and terminate at `end` time.",
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"start": {
							Description:      "Time from which schedule will be active. Need to be in YYYY-MM-DD HH:mm:SS format. Eg 2006-01-02 15:04:05",
							Type:             schema.TypeString,
							Required:         true,
							ValidateDiagFunc: dateValidationFunc,
						},
						"end": {
							Description:      "Time until which schedule will be active. Need to be in YYYY-MM-DD HH:mm:SS format. Eg 2006-01-02 15:04:05",
							Type:             schema.TypeString,
							Required:         true,
							ValidateDiagFunc: dateValidationFunc,
						},
					},
				},
			},
			"periodicity": {
				Description: "For defining periodic schedule. Periodic nature will be applicable from the time of creation of schedule, unless specific 'time_period' is specified",
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"days": {
							Description: "Days on which schedule need to be active. Comma separated values of `SUN`, `MON`, `TUE`, `WED`, `THU`, `FRI` and `SAT`. Eg : `MON,TUE,WED,THU,FRI` for Mon through Friday",
							Type:        schema.TypeString,
							Required:    true,
							ValidateDiagFunc: func(i interface{}, p cty.Path) diag.Diagnostics {
								diags := diag.Diagnostics{}
								v, ok := i.(string)
								if !ok {
									diags = append(diags, diag.Diagnostic{
										Severity: diag.Error,
										Summary:  "Value is mandatory and should be string",
									})
									return diags
								}
								parts := strings.Split(v, ",")
								valids := []string{"SUN", "MON", "TUE", "WED", "THU", "FRI", "SAT"}
								for _, p := range parts {
									vp := strings.TrimSpace(p)
									match := false
									for _, vv := range valids {
										match = match || strings.EqualFold(vp, vv)
									}
									if !match {
										diags = append(diags, diag.Diagnostic{
											Severity: diag.Error,
											Summary:  "Valid input is comma separated values of `SUN`, `MON`, `TUE`, `WED`, `THU`, `FRI` and `SAT`. Eg : `MON,TUE,WED,THU,FRI` for Mon through Friday ",
										})
									}
								}
								if len(valids) < 1 || len(valids) > 7 {
									diags = append(diags, diag.Diagnostic{
										Severity: diag.Error,
										Summary:  "At-least one and at-most seven days can be specified",
									})
								}
								return diags
							},
						},
						"start_time": {
							Description:      "Starting time of schedule action on the day. Accepted format is HH:MM. Eg : 13:15 for 01:15pm",
							Type:             schema.TypeString,
							Required:         true,
							ValidateDiagFunc: timeValidateFunc,
						},
						"end_time": {
							Description:      "Ending time of schedule action on the day. Accepted format is HH:MM. Eg : 20:00 for 8pm",
							Type:             schema.TypeString,
							Required:         true,
							ValidateDiagFunc: timeValidateFunc,
						},
					},
				},
			},
		},
	}

	return resource
}

func resourceScheduleCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return nil
}

func resourceScheduleUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return nil
}

func resourceScheduleDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return nil
}

func resourceScheduleRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return nil
}
