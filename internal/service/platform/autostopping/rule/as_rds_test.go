package as_rule_test

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestResourceRDSRule(t *testing.T) {
	name := fmt.Sprintf("terr-rds-%s", randAlnum(5)) // "terr-rds-"+5 = 13 chars
	proxyName := fmt.Sprintf("terr-rds-p-%s", randAlnum(5))
	apiKey := os.Getenv(platformAPIKeyEnv)
	var proxyID string
	resourceName := "harness_autostopping_rule_rds.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testRDSRule(name, proxyName, apiKey, true),
				Check: resource.ComposeTestCheckFunc(
					extractAttr("harness_autostopping_aws_proxy.test", "identifier", &proxyID),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "dry_run", "true"),
				),
			},
			{
				PreConfig: func() {
					if err := waitForProxyReady(proxyID, 90*time.Second); err != nil {
						t.Fatalf("Proxy not ready: %v", err)
					}
				},
				Config: testRDSRule(name, proxyName, apiKey, false),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "dry_run", "false"),
				),
			},
			{
				Config: testRDSRuleUpdate(name, proxyName, apiKey, "15", true),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "idle_time_mins", "15"),
					resource.TestCheckResourceAttr(resourceName, "dry_run", "true"),
				),
			},
			{
				Config: testRDSRuleUpdate(name, proxyName, apiKey, "20", false),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "idle_time_mins", "20"),
					resource.TestCheckResourceAttr(resourceName, "dry_run", "false"),
				),
			},
		},
	})
}

func testRDSAWSProxy(proxyName, apiKey string) string {
	return fmt.Sprintf(`
resource "harness_autostopping_aws_proxy" "test" {
  name                 = %[1]q
  cloud_connector_id   = %[2]q
  region               = "us-east-1"
  vpc                  = %[3]q
  security_groups      = [%[4]q]
  machine_type         = "t2.medium"
  api_key              = %[5]q
  allocate_static_ip   = false
  delete_cloud_resources_on_destroy = false
}
`, proxyName, cloudConnectorIDAWS, awsProxyVPC, awsProxySG, apiKey)
}

func testRDSRule(name, proxyName, apiKey string, dryRun bool) string {
	return testRDSAWSProxy(proxyName, apiKey) + fmt.Sprintf(`
resource "harness_autostopping_rule_rds" "test" {
  name               = %[1]q
  cloud_connector_id = %[2]q
  idle_time_mins     = 10
  dry_run            = %[3]t

  database {
    id     = "database_id"
    region = "us-east-1"
  }
  tcp {
    proxy_id = harness_autostopping_aws_proxy.test.identifier
    forward_rule {
      connect_on = 2233
      port       = 2233
    }
  }
}
`, name, cloudConnectorIDAWS, dryRun)
}

func testRDSRuleUpdate(name, proxyName, apiKey, idleTime string, dryRun bool) string {
	return testRDSAWSProxy(proxyName, apiKey) + fmt.Sprintf(`
resource "harness_autostopping_rule_rds" "test" {
  name               = %[1]q
  cloud_connector_id = %[2]q
  idle_time_mins     = %[3]s
  dry_run            = %[4]t

  database {
    id     = "database_id"
    region = "us-east-1"
  }
  tcp {
    proxy_id = harness_autostopping_aws_proxy.test.identifier
    forward_rule {
      connect_on = 2233
      port       = 2233
    }
  }
}
`, name, cloudConnectorIDAWS, idleTime, dryRun)
}
