package as_rule_test

import (
	"fmt"
	"testing"

	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestResourceScaleGroupRule(t *testing.T) {
	name := "terraform-scale-group-rule-test"
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
		cloud_connector_id = "Lightwing_Non_Prod_5" 
		idle_time_mins = 10              
		dry_run = %[2]t
		scale_group {
			id = "arn:aws:autoscaling:us-east-1:1234:autoScalingGroup:abcd:autoScalingGroupName/demo-asg"
			name = "demo-asg"
			region = "us-east-1"
			desired = 2
			min = 1
			max = 5
			on_demand = 1
		}       
	}
`, name, dryRun)
}
