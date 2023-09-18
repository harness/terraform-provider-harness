package as_rule_test

import (
	"fmt"
	"testing"

	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestResourceECSRule(t *testing.T) {
	name := "terraform-rule-test"
	resourceName := "harness_autostopping_rule_ecs.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		//		CheckDestroy:      testRuleDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testECSRule(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
				),
			},
		},
	})
}

func testECSRule(name string) string {
	return fmt.Sprintf(`
	resource "harness_autostopping_rule_ecs" "test" {
		name = "%[1]s"  
		cloud_connector_id = "Azure_SE" 
		idle_time_mins = 10              
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
`, name)
}
