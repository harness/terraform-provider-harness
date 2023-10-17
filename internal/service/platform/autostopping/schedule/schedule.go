package schedule

import (
	"context"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type scheduleType string

const (
	timeZoneAttribute     = "time_zone"
	timePeriodAttribute   = "time_period"
	scheduleTypeAttribute = "schedule_type"
	startAttribute        = "start"
	endAttribute          = "end"
	periodicityAttribute  = "periodicity"
	startTimeAttribute    = "start_time"
	endTimeAttribute      = "end_time"
	rulesAttribute        = "rules"
	daysAttribute         = "days"
	nameAttribute         = "name"
	scheduleResTypeASrule = "autostopping_rule"
)

var (
	uptimeSchedule   scheduleType = "uptime"
	downtimeSchedule scheduleType = "downtime"
	dayIndex                      = map[string]time.Weekday{
		"SUN": 0,
		"MON": 1,
		"TUE": 2,
		"WED": 3,
		"THU": 4,
		"FRI": 5,
		"SAT": 6,
	}
)

func dateValidation(i interface{}, p cty.Path) diag.Diagnostics {
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

func timeValidation(i interface{}, p cty.Path) diag.Diagnostics {
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

func daysValidationFunc(i interface{}, p cty.Path) diag.Diagnostics {
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
	unique := map[string]struct{}{}
	for _, p := range parts {
		vp := strings.TrimSpace(p)
		if _, checked := unique[vp]; checked {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  fmt.Sprintf("Day `%s` repeats in days", vp),
			})
			return diags
		}
		unique[vp] = struct{}{}
		match := false
		for vd := range dayIndex {
			match = match || strings.EqualFold(vp, vd)
		}
		if !match {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Valid input is comma separated values of `SUN`, `MON`, `TUE`, `WED`, `THU`, `FRI` and `SAT`. Eg : `MON,TUE,WED,THU,FRI` for Mon through Friday ",
			})
		}
	}
	if len(parts) < 1 || len(parts) > 7 {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "At-least one and at-most seven days can be specified",
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
			nameAttribute: {
				Description: "Name of the schedule",
				Type:        schema.TypeString,
				Required:    true,
			},
			scheduleTypeAttribute: {
				Description:  fmt.Sprintf("Type of the schedule. Valid values are `%s` and `%s`", uptimeSchedule, downtimeSchedule),
				Type:         schema.TypeString,
				Required:     true,
				ExactlyOneOf: []string{string(uptimeSchedule), string(downtimeSchedule)},
			},
			timeZoneAttribute: {
				Description: "Time zone in which schedule needs to be executed",
				Type:        schema.TypeString,
				Required:    true,
			},
			timePeriodAttribute: {
				Description: "Time period in which schedule will be active. If specified along with periodicity, this will act as the boundary of periodicity. Otherwise schedule action will be triggered at `start` time and terminate at `end` time.",
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						startAttribute: {
							Description:      "Time from which schedule will be active. Need to be in YYYY-MM-DD HH:mm:SS format. Eg 2006-01-02 15:04:05",
							Type:             schema.TypeString,
							Required:         true,
							ValidateDiagFunc: dateValidation,
						},
						endAttribute: {
							Description:      "Time until which schedule will be active. Need to be in YYYY-MM-DD HH:mm:SS format. Eg 2006-01-02 15:04:05",
							Type:             schema.TypeString,
							Required:         true,
							ValidateDiagFunc: dateValidation,
						},
					},
				},
			},
			periodicityAttribute: {
				Description: "For defining periodic schedule. Periodic nature will be applicable from the time of creation of schedule, unless specific 'time_period' is specified",
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						daysAttribute: {
							Description:      "Days on which schedule need to be active. Comma separated values of `SUN`, `MON`, `TUE`, `WED`, `THU`, `FRI` and `SAT`. Eg : `MON,TUE,WED,THU,FRI` for Mon through Friday",
							Type:             schema.TypeString,
							Required:         true,
							ValidateDiagFunc: daysValidationFunc,
						},
						startTimeAttribute: {
							Description:      "Starting time of schedule action on the day. Accepted format is HH:MM. Eg : 13:15 for 01:15pm",
							Type:             schema.TypeString,
							Required:         true,
							ValidateDiagFunc: timeValidation,
						},
						endTimeAttribute: {
							Description:      "Ending time of schedule action on the day. Accepted format is HH:MM. Eg : 20:00 for 8pm",
							Type:             schema.TypeString,
							Required:         true,
							ValidateDiagFunc: timeValidation,
						},
					},
				},
			},
			rulesAttribute: {
				Description: "ID of AutoStopping rules on which the schedule applies",
				Required:    true,
				Type:        schema.TypeList,
				MinItems:    1,
				Elem: &schema.Schema{
					Type: schema.TypeFloat,
				},
			},
		},
	}

	return resource
}

func resourceScheduleCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)
	schedule := parseSchedule(d, c.AccountId)
	return createSchedule(ctx, d, meta, schedule)
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

func parseSchedule(d *schema.ResourceData, accountId string) *nextgen.FixedSchedule {
	schedule := &nextgen.FixedSchedule{
		Account: accountId,
		Details: &nextgen.OccurrenceSchedule{},
	}
	if attr, ok := d.GetOk(nameAttribute); ok {
		name, ok := attr.(string)
		if ok {
			schedule.Name = name
		}
	}

	if attr, ok := d.GetOk(timeZoneAttribute); ok {
		timezone, ok := attr.(string)
		if ok {
			schedule.Details.Timezone = timezone
		}
	}

	tSchedule := &nextgen.TimeSchedule{}

	attr, ok := d.GetOk(timePeriodAttribute)
	if ok {
		timePeriodInf, ok := attr.([]interface{})
		if ok && len(timePeriodInf) > 0 {
			tSchedule.Period = &nextgen.TimeSchedulePeriod{}
			timePeriodObj, ok := timePeriodInf[0].(map[string]interface{})
			if ok {
				startInf, ok := timePeriodObj[startAttribute]
				if ok {
					start, ok := startInf.(string)
					if ok {
						tSchedule.Period.Start = start
					}
				}
				endInf, ok := timePeriodObj[endAttribute]
				if ok {
					end, ok := endInf.(string)
					if ok {
						tSchedule.Period.End = end
					}
				}
			}
		}
	}

	attr, ok = d.GetOk(periodicityAttribute)
	if ok {
		periodicInf, ok := attr.([]interface{})
		if ok && len(periodicInf) > 0 {
			tSchedule.Days = &nextgen.TimeScheduleDays{}
			periodicityObj, ok := periodicInf[0].(map[string]interface{})
			if ok {
				days := []float64{}
				daysInf, ok := periodicityObj[daysAttribute]
				if ok {
					daysCsv, ok := daysInf.(string)
					if ok {
						dayParts := strings.Split(daysCsv, ",")
						for _, dp := range dayParts {
							dv := strings.TrimSpace(dp)
							i, ok := dayIndex[strings.ToLower(dv)]
							if ok {
								days = append(days, float64(i))
							}
						}
					}
				}
				sort.Float64s(days)
				tSchedule.Days.Days = days

				startTimeInf, ok := periodicityObj[startTimeAttribute]
				if ok {
					startTimeStr, ok := startTimeInf.(string)
					if ok {
						startTime := parseTimeInDay(startTimeStr)
						tSchedule.Days.StartTime = &startTime
					}
				}

				endTimeInf, ok := periodicityObj[endTimeAttribute]
				if ok {
					endTimeStr, ok := endTimeInf.(string)
					if ok {
						endTime := parseTimeInDay(endTimeStr)
						tSchedule.Days.EndTime = &endTime
					}
				}
			}
		}
	}

	if attr, ok := d.GetOk(scheduleTypeAttribute); ok {
		scheduleType, ok := attr.(string)
		if ok {
			if strings.EqualFold(scheduleType, string(uptimeSchedule)) {
				schedule.Details.Uptime = tSchedule
			}
			if strings.EqualFold(scheduleType, string(downtimeSchedule)) {
				schedule.Details.Downtime = tSchedule
			}
		}
	}
	if attr, ok := d.GetOk(rulesAttribute); ok {
		schedule.Resources = []nextgen.StaticScheduleResource{}
		ruleIDsInf, ok := attr.([]interface{})
		if ok {
			for _, ruleIDInf := range ruleIDsInf {
				ruleID, ok := ruleIDInf.(float64)
				if ok {
					res := nextgen.StaticScheduleResource{
						Id:    fmt.Sprintf("%d", int(ruleID)),
						Type_: scheduleResTypeASrule,
					}
					schedule.Resources = append(schedule.Resources, res)
				}
			}
		}
	}
	return schedule
}

func parseTimeInDay(timeInDayStr string) nextgen.TimeInDay {
	timeParts := strings.Split(strings.TrimSpace(timeInDayStr), ":")
	endTime := nextgen.TimeInDay{}
	if len(timeParts) == 2 {
		endTimeHr, err := strconv.ParseInt(timeParts[0], 10, 64)
		if err == nil {
			endTime.Hour = float64(endTimeHr)
		}
		endTimeMin, err := strconv.ParseInt(timeParts[0], 10, 64)
		if err == nil {
			endTime.Min = float64(endTimeMin)
		}
	}
	return endTime
}

func createSchedule(ctx context.Context, d *schema.ResourceData, meta interface{}, schedule *nextgen.FixedSchedule) diag.Diagnostics {
	diag := diag.Diagnostics{}
	return diag
}
