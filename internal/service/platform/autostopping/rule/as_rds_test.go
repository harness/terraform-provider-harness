package as_rule_test

import (
	"fmt"
	"testing"

	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestResourceRDSRule(t *testing.T) {
	name := "terraform-rule-test-rds"
	resourceName := "harness_autostopping_rule_rds.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		// CheckDestroy:      testRuleDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testRDSRule(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
				),
			},
			{
				Config: testRDSRuleUpdate(name, "15"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "idle_time_mins", "15"),
				),
			},
		},
	})
}

func testRDSRule(name string) string {
	return fmt.Sprintf(`
	resource "harness_autostopping_rule_rds" "test" {
		name = "%[1]s"  
		cloud_connector_id = "DoNotDelete_LightwingNonProd" 
		idle_time_mins = 10              
		database {
			id = "database_id"
		  	region = "us-east-1"
		}
		tcp {
			proxy_id = "ap-ciun1635us1fhpjiotfg"             
			forward_rule {
				port = 2233
			}                     
		}      
	}
`, name)
}

func testRDSRuleUpdate(name string, idleTime string) string {
	return fmt.Sprintf(`
	resource "harness_autostopping_rule_rds" "test" {
		name = "%[1]s"  
		cloud_connector_id = "DoNotDelete_LightwingNonProd" 
		idle_time_mins = %[2]s             
		database {
			id = "database_id"
		  	region = "us-east-1"
		}
		tcp {
			proxy_id = "ap-ciun1635us1fhpjiotfg"             
			forward_rule {
				port = 2233
			}                     
		}  
	}
`, name, idleTime)
}
