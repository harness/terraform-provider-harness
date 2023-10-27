package schedule

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceFixedSchedule() *schema.Resource {
	resource := &schema.Resource{
		Description: "Data source for retrieving a fixed schedule for Harness AutoStopping rule",
		ReadContext: resourceScheduleRead,
		Schema: map[string]*schema.Schema{
			"identifier": {
				Description: "Unique identifier of the schedule",
				Type:        schema.TypeFloat,
				Computed:    true,
			},
			nameAttribute: {
				Description: "Name of the schedule",
				Type:        schema.TypeString,
				Computed:    true,
			},
			scheduleTypeAttribute: {
				Description: fmt.Sprintf("Type of the schedule. Valid values are `%s` and `%s`", uptimeSchedule, downtimeSchedule),
				Type:        schema.TypeString,
				Required:    true,
			},
			timeZoneAttribute: {
				Description: "Time zone in which schedule needs to be executed",
				Type:        schema.TypeString,
				Computed:    true,
			},
			startingFromAttribute: {
				Description: "Time from which schedule will be active. Schedule will take immediate effect if starting_from is not specified. Need to be in YYYY-MM-DD HH:mm:SS format. Eg 2006-01-02 15:04:05",
				Type:        schema.TypeString,
				Computed:    true,
			},
			endingOnAttribute: {
				Description: "Time until which schedule will be active. Need to be in YYYY-MM-DD HH:mm:SS format. Eg 2006-01-02 15:04:05",
				Type:        schema.TypeString,
				Computed:    true,
			},
			repetitionAttribute: {
				Description: "For defining periodic schedule. Periodic nature will be applicable from the time of creation of schedule, unless specific 'time_period' is specified",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						daysAttribute: {
							Description: "List of days on which schedule need to be active. Valid values are `SUN`, `MON`, `TUE`, `WED`, `THU`, `FRI` and `SAT`.",
							Type:        schema.TypeList,
							Computed:    true,
							Elem: schema.Schema{
								Type: schema.TypeString,
							},
						},
						startTimeAttribute: {
							Description: "Starting time of schedule action on the day. Accepted format is HH:MM. Eg : 13:15 for 01:15pm",
							Type:        schema.TypeString,
							Computed:    true,
						},
						endTimeAttribute: {
							Description: "Ending time of schedule action on the day. Accepted format is HH:MM. Eg : 20:00 for 8pm",
							Type:        schema.TypeString,
							Computed:    true,
						},
					},
				},
			},
			rulesAttribute: {
				Description: "ID of AutoStopping rules on which the schedule applies",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Schema{
					Type: schema.TypeFloat,
				},
			},
		},
	}

	return resource
}
