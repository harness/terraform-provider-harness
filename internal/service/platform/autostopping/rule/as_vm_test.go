package as_rule_test

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestResourceVMRule(t *testing.T) {
	name := utils.RandStringBytes(5)
	resourceName := "harness_autostopping_rule_vm.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testVMRule(name, true),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "dry_run", "true"),
				),
			},
			{
				Config: testVMRule(name, false),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "dry_run", "false"),
				),
			},
		},
	})
}

func testRuleDestroy(resourceName string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		rule, _ := testGetRule(resourceName, state)
		if rule != nil {
			return fmt.Errorf("Found vm rule: %d", rule.Id)
		}
		return nil
	}
}

func testGetRule(resourceName string, state *terraform.State) (*nextgen.Service, error) {
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

func testVMRule(name string, dryRun bool) string {
	return fmt.Sprintf(`
		resource "harness_autostopping_rule_vm" "test" {
		name = "%[1]s"  
		cloud_connector_id = %q
		idle_time_mins = 10              
		dry_run = %[2]t
		filter {
			vm_ids = [%q]
		  	regions = [%q]
		}
		http {
			proxy_id = %q
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
			proxy_id = %q
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
			rule_id = %d
			delay_in_sec = 5
		}        
	}
`, name, dryRun, cloudConnectorIDVM, vmFilterVMID, vmFilterRegion, proxyIDVM, proxyIDVM, ruleIDDependency)
}

