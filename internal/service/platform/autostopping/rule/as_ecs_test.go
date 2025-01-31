package as_rule_test

import (
	"fmt"
	"testing"

	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestResourceECSRule(t *testing.T) {
	name := "terraform-rule-test-ecs"
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
		cloud_connector_id = "Azure_SE" 
		idle_time_mins = 10
		dry_run = %[2]t              

		container {
			cluster = "cluster"
			service = "service"
			region = "us-east-1"
			task_count = 1
		}		
		http {
			proxy_id = "ap-chdpf8f83v0c1aj69oog"             
		}
		depends {
			rule_id = 24576
			delay_in_sec = 5
		}        
	}
`, name, dryRun)
}

func testECSRuleUpdate(name string, idleTime string, dryRun bool) string {
	return fmt.Sprintf(`
	resource "harness_autostopping_rule_ecs" "test" {
		name = "%[1]s"  
		cloud_connector_id = "Azure_SE" 
		idle_time_mins = %[2]s
		dry_run = %[3]t              

		container {
			cluster = "cluster"
			service = "service"
			region = "us-east-1"
			task_count = 1
		}		
		http {
			proxy_id = "ap-chdpf8f83v0c1aj69oog"             
		}
		depends {
			rule_id = 24576
			delay_in_sec = 5
		}        
	}
`, name, idleTime, dryRun)
}
