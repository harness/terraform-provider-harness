package schedule

import (
	"testing"

	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func capturePanic(fn func()) (panicked bool, value interface{}) {
	defer func() {
		if r := recover(); r != nil {
			panicked = true
			value = r
		}
	}()
	fn()
	return panicked, value
}

func TestCCM32488Class_SetScheduleNilUptimeAndDowntime(t *testing.T) {
	d := schema.TestResourceDataRaw(t, ResourceVMRule().Schema, map[string]interface{}{
		"name":          "sched",
		"schedule_type": "uptime",
	})
	panicked, val := capturePanic(func() {
		_ = setSchedule(d, &nextgen.FixedSchedule{
			Id:   1,
			Name: "sched",
			Details: &nextgen.OccurrenceSchedule{
				Uptime:   nil,
				Downtime: nil,
				Timezone: "UTC",
			},
		})
	})
	if panicked {
		t.Errorf("setSchedule panicked (bug confirmed): %v", val)
	} else {
		t.Log("RESULT: PASS (no panic)")
	}
}
