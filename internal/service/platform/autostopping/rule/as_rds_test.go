package as_rule_test

import (
	"fmt"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestResourceRDSRule(t *testing.T) {
	name := utils.RandStringBytes(5)
	resourceName := "harness_autostopping_rule_rds.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testRDSRule(name, true),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "dry_run", "true"),
				),
			},
			{
				Config: testRDSRule(name, false),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "dry_run", "false"),
				),
			},
			{
				Config: testRDSRuleUpdate(name, "15", true),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "idle_time_mins", "15"),
					resource.TestCheckResourceAttr(resourceName, "dry_run", "true"),
				),
			},
			{
				Config: testRDSRuleUpdate(name, "20", false),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "idle_time_mins", "20"),
					resource.TestCheckResourceAttr(resourceName, "dry_run", "false"),
				),
			},
		},
	})
}

func testRDSRule(name string, dryRun bool) string {
	return fmt.Sprintf(`
	resource "harness_autostopping_rule_rds" "test" {
		name = "%[1]s"  
		cloud_connector_id = %q
		idle_time_mins = 10
		dry_run = %[2]t            

		database {
			id = "database_id"
		  	region = "us-east-1"
		}
		tcp {
			proxy_id = %q
			forward_rule {
				port = 2233
			}                     
		}      
	}
`, name, dryRun, cloudConnectorIDRDS, proxyIDRDS)
}

func testRDSRuleUpdate(name string, idleTime string, dryRun bool) string {
	return fmt.Sprintf(`
	resource "harness_autostopping_rule_rds" "test" {
		name = "%[1]s"  
		cloud_connector_id = %q
		idle_time_mins = %[2]s
		dry_run = %[3]t              

		database {
			id = "database_id"
		  	region = "us-east-1"
		}
		tcp {
			proxy_id = %q
			forward_rule {
				port = 2233
			}                     
		}  
	}
`, name, idleTime, dryRun, cloudConnectorIDRDS, proxyIDRDS)
}
