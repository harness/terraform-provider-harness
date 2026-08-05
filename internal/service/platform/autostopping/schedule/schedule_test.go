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

func TestResourceVMRule(t *testing.T) {
	name := "terraform-schedule-test"
	resourceName := "harness_autostopping_schedule.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testSchedule(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
				),
			},
		},
	})
}

// TestAccResourceSchedule_CCM32336_OutOfBandDeleteRecreates verifies that when an
// AutoStopping fixed schedule is deleted out-of-band (UI / direct API), the
// next refresh treats the GET as "not found" and re-plans a create instead of
// erroring out with "giving up after 11 attempt(s)".
//
// Regression test for CCM-32336. helpers.HandleReadApiError clears state on
// 404 + ENTITY_NOT_FOUND so terraform plan reports a recreate.
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
				Config: testSchedule(name),
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
				Config:             testSchedule(name),
				PlanOnly:           true,
				ExpectNonEmptyPlan: true,
			},
			{
				Config: testSchedule(name),
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

func testSchedule(name string) string {
	return fmt.Sprintf(`
	resource "harness_autostopping_schedule" "test" {
		name          = "%s"
		schedule_type = "uptime"
		time_zone     = "UTC"
		starting_from = "2023-01-02 15:04:05"
		ending_on     = "2024-02-02 15:04:05"

		repeats {
			days      = ["MON", "THU", "TUE"]
			start_time = "09:30"
			end_time   = "16:50"
		}

		rules = [123, 234]
	}`, name)
}
