package as_rule_test

import (
	"fmt"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestResourceECSRule(t *testing.T) {
	name := utils.RandStringBytes(5)
	resourceName := "harness_autostopping_rule_ecs.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testECSRule(name, true),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "dry_run", "true"),
				),
			},
			{
				Config: testECSRule(name, false),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "dry_run", "false"),
				),
			},
			{
				Config: testECSRuleUpdate(name, "15", true),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "idle_time_mins", "15"),
					resource.TestCheckResourceAttr(resourceName, "dry_run", "true"),
				),
			},
			{
				Config: testECSRuleUpdate(name, "20", false),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "idle_time_mins", "20"),
					resource.TestCheckResourceAttr(resourceName, "dry_run", "false"),
				),
			},
		},
	})
}

func testECSRule(name string, dryRun bool) string {
	return fmt.Sprintf(`
	resource "harness_autostopping_rule_ecs" "test" {
		name = "%[1]s"  
		cloud_connector_id = %[2]q
		idle_time_mins = 10
		dry_run = %[3]t              

		container {
			cluster = "cluster"
			service = "service"
			region = "us-east-1"
			task_count = 1
		}		
		http {
			proxy_id = %[4]q
		}
		depends {
			rule_id = %[5]d
			delay_in_sec = 5
		}        
	}
`, name, cloudConnectorIDVM, dryRun, proxyIDVM, ruleIDDependency)
}

func testECSRuleUpdate(name string, idleTime string, dryRun bool) string {
	return fmt.Sprintf(`
	resource "harness_autostopping_rule_ecs" "test" {
		name = "%[1]s"  
		cloud_connector_id = %[2]q
		idle_time_mins = %[3]s
		dry_run = %[4]t              

		container {
			cluster = "cluster"
			service = "service"
			region = "us-east-1"
			task_count = 1
		}		
		http {
			proxy_id = %[5]q
		}
		depends {
			rule_id = %[6]d
			delay_in_sec = 5
		}        
	}
`, name, cloudConnectorIDVM, idleTime, dryRun, proxyIDVM, ruleIDDependency)
}
