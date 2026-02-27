package as_rule_test

import (
	"fmt"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestResourceScaleGroupRule(t *testing.T) {
	name := utils.RandStringBytes(5)
	resourceName := "harness_autostopping_rule_scale_group.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testScaleGroupRule(name, true),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "dry_run", "true"),
				),
			},
			{
				Config: testScaleGroupRule(name, false),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "dry_run", "false"),
				),
			},
		},
	})
}

func testScaleGroupRule(name string, dryRun bool) string {
	return fmt.Sprintf(`
	resource "harness_autostopping_rule_scale_group" "test" {
		name = "%[1]s"  
		cloud_connector_id = %q
		idle_time_mins = 10              
		dry_run = %[2]t
		scale_group {
			id = %q
			name = %q
			region = %q
			desired = 2
			min = 1
			max = 5
			on_demand = 1
		}       
	}
`, name, dryRun, cloudConnectorIDScaleGroup, scaleGroupARN, scaleGroupName, scaleGroupRegion)
}
