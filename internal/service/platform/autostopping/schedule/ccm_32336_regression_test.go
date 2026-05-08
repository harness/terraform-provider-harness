package schedule_test

import (
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

// Blocked on CCM-32403: AutoStopping schedule GET API still returns HTTP 500
// for a deleted schedule (terraform plan: "giving up after 11 attempt(s)").
// Test is parked until the schedule GET returns 404 + ENTITY_NOT_FOUND.
//
// TestAccResourceSchedule_CCM32336_OutOfBandDeleteRecreates verifies that when an
// AutoStopping fixed schedule is deleted out-of-band (UI / direct API), the
// next refresh treats the GET as "not found" and re-plans a create instead of
// erroring out with "giving up after 11 attempt(s)".
//
// Regression test for CCM-32336 (the bug error trace in the ticket explicitly
// referenced /schedules/<id>) and the schedule-variant tracked under CCM-32403.
// The fix changes the schedule GET API to return 404 + ENTITY_NOT_FOUND so that
// helpers.HandleReadApiError clears state and recreates instead of failing the
// terraform plan.
func TestAccResourceSchedule_CCM32336_OutOfBandDeleteRecreates(t *testing.T) {
	suffix := strings.ToLower(utils.RandStringBytes(5))
	name := fmt.Sprintf("terr-c336sched-%s", suffix)
	resourceName := "harness_autostopping_schedule.test"

	var schedIDBefore string

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testCCM32336Schedule(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttrWith(resourceName, "id", func(value string) error {
						schedIDBefore = value
						return nil
					}),
				),
			},
			{
				PreConfig: func() {
					c, ctx := acctest.TestAccGetPlatformClientWithContext()
					sid, err := strconv.ParseFloat(schedIDBefore, 64)
					if err != nil {
						t.Fatalf("invalid captured schedule id %q: %v", schedIDBefore, err)
					}
					if _, _, err := c.CloudCostAutoStoppingFixedSchedulesApi.DeleteAutoStoppingFixedSchedule(
						ctx, c.AccountId, sid, c.AccountId,
					); err != nil {
						t.Fatalf("CCM-32336: out-of-band delete failed: %v", err)
					}
				},
				Config:             testCCM32336Schedule(name),
				PlanOnly:           true,
				ExpectNonEmptyPlan: true,
			},
			{
				Config: testCCM32336Schedule(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttrWith(resourceName, "id", func(value string) error {
						if value == "" || value == schedIDBefore {
							return fmt.Errorf("expected new schedule id after recreate, got %q (before %q)", value, schedIDBefore)
						}
						return nil
					}),
				),
			},
		},
	})
}

func testCCM32336Schedule(name string) string {
	return fmt.Sprintf(`
	resource "harness_autostopping_schedule" "test" {
		name          = "%s"
		schedule_type = "uptime"
		time_zone     = "UTC"
		starting_from = "2023-01-02 15:04:05"
		ending_on     = "2024-02-02 15:04:05"

		repeats {
			days       = ["MON", "THU", "TUE"]
			start_time = "09:30"
			end_time   = "16:50"
		}

		rules = [123, 234]
	}`, name)
}
