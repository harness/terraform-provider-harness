package as_rule_test

import (
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestResourceVMRule(t *testing.T) {
	name := fmt.Sprintf("terr-vm-%s", randAlnum(5))
	proxyName := fmt.Sprintf("terr-az-p-%s", randAlnum(5))
	apiKey := os.Getenv(platformAPIKeyEnv)
	resourceName := "harness_autostopping_rule_vm.test"

	var proxyID string

	resource.UnitTest(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			cleanupStaleRulesForVM(t)
		},
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testVMRule(name, proxyName, apiKey, true),
				Check: resource.ComposeTestCheckFunc(
					extractAttr("harness_autostopping_azure_proxy.test", "identifier", &proxyID),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "dry_run", "true"),
				),
			},
			{
				PreConfig: func() {
					if err := waitForProxyReady(proxyID, 3*time.Minute); err != nil {
						t.Skipf("Skipping: %v", err)
					}
				},
				Config: testVMRule(name, proxyName, apiKey, false),
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

func testVMAzureProxy(proxyName, apiKey string) string {
	return fmt.Sprintf(`
resource "harness_autostopping_azure_proxy" "test" {
  name                 = %[1]q
  cloud_connector_id   = %[2]q
  region               = %[3]q
  resource_group       = %[4]q
  vpc                  = %[5]q
  subnet_id            = %[6]q
  security_groups      = [%[7]q]
  machine_type         = "Standard_D2s_v3"
  keypair              = "DoNotDelete-Terraform-AS-Test-VM_key"
  api_key              = %[8]q
  allocate_static_ip   = false
  delete_cloud_resources_on_destroy = false
}
`, proxyName, cloudConnectorIDVM, vmFilterRegion, azureProxyResourceGroup,
		azureProxyVNet, azureProxySubnet, azureProxyNSG, apiKey)
}

func testVMRule(name, proxyName, apiKey string, dryRun bool) string {
	return testVMAzureProxy(proxyName, apiKey) + fmt.Sprintf(`
resource "harness_autostopping_rule_vm" "test" {
  name               = %[1]q
  cloud_connector_id = %[2]q
  idle_time_mins     = 10
  dry_run            = %[3]t

  filter {
    vm_ids  = [%[4]q]
    regions = [%[5]q]
    zones   = []
  }
  http {
    proxy_id = harness_autostopping_azure_proxy.test.identifier
    routing {
      source_protocol = "https"
      target_protocol = "https"
      source_port     = 443
      target_port     = 443
      action          = "forward"
    }
    routing {
      source_protocol = "http"
      target_protocol = "http"
      source_port     = 80
      target_port     = 80
      action          = "forward"
    }
    health {
      protocol         = "http"
      port             = 80
      path             = "/"
      timeout          = 30
      status_code_from = 200
      status_code_to   = 299
    }
  }
  tcp {
    proxy_id = harness_autostopping_azure_proxy.test.identifier
    ssh {
      connect_on = 22
      port       = 22
    }
    rdp {
      connect_on = 3389
      port       = 3389
    }
    forward_rule {
      connect_on = 2233
      port       = 2233
    }
  }
}
`, name, cloudConnectorIDVM, dryRun, vmFilterVMID, vmFilterRegion)
}
