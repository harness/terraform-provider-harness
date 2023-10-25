package schedule_test

import (
	"fmt"
	"testing"

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

func testSchedule(name string) string {
	return fmt.Sprintf(`
	resource "harness_autostopping_schedule" "test" {
		name = "%s"
		schedule_type = "uptime"
		time_zone = "UTC"    
	
		time_period {
			start = "2023-01-02 15:04:05"
			end = "2024-02-02 15:04:05"
		}
	
		periodicity {
			days = "MON, THU, TUE"
			start_time = "09:30"
			end_time = "16:50"
		}
	
		rules = [ 123, 234 ]
	}`, name)
}
