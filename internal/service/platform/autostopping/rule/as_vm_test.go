package as_rule_test

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestResourceVMRule(t *testing.T) {
	name := "terraform-rule-test"
	resourceName := "harness_autostopping_rule_vm.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		//		CheckDestroy:      testVMRuleDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testVMRule(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
				),
			},
			// {
			// 	ResourceName:            resourceName,
			// 	ImportState:             true,
			// 	ImportStateVerify:       true,
			// },
		},
	})
}

func testVMRuleDestroy(resourceName string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		rule, _ := testGetVMRule(resourceName, state)
		if rule != nil {
			return fmt.Errorf("Found vm rule: %d", rule.Id)
		}
		return nil
	}
}

func testGetVMRule(resourceName string, state *terraform.State) (*nextgen.Service, error) {
	r := acctest.TestAccGetResource(resourceName, state)
	c, ctx := acctest.TestAccGetPlatformClientWithContext()
	ruleId, err := strconv.ParseFloat(r.Primary.ID, 64)
	if err != nil {
		return nil, err
	}
	resp, _, err := c.CloudCostAutoStoppingRulesApi.AutoStoppingRuleDetails(ctx, c.AccountId, ruleId, c.AccountId)

	if err != nil {
		return nil, err
	}

	return resp.Response.Service, nil
}

func testVMRule(name string) string {
	return fmt.Sprintf(`
	resource "harness_autostopping_rule_vm" "test" {
		name = "%[1]s"  
		cloud_connector_id = "Azure_SE" 
		idle_time_mins = 10              
		filter {
			vm_ids = ["/subscriptions/e8389fc5-0cb8-44ab-947b-c6cf62552be0/resourceGroups/tkouhsari-autostop-1_group/providers/Microsoft.Compute/virtualMachines/tkouhsari-autostop-3"]
		  regions = ["useast2"]
		}
		http {
			proxy_id = "ap-chdpf8f83v0c1aj69oog"           
			routing {
				source_protocol = "https"
				target_protocol = "https"
				source_port = 443
				target_port = 443
				action = "forward"
			}           
			routing {
				source_protocol = "http"
				target_protocol = "http"
				source_port = 80
				target_port = 80
				action = "forward"
			}
			health {
				protocol = "http"
				port = 80
				path = "/"
				timeout = 30
				status_code_from = 200
				status_code_to = 299
			}
		}
		tcp {
			proxy_id = "ap-chdpf8f83v0c1aj69oog"
			ssh {
				port = 22
			}
			rdp {
				port = 3389
			}               
			forward_rule {
				port = 2233
			}                     
		}
		depends {
			rule_id = 24576
			delay_in_sec = 5
		}        
	}
`, name)
}
