package as_rule_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

const platformAPIKeyEnv = "HARNESS_PLATFORM_API_KEY"

func TestResourceECSRule(t *testing.T) {
	// Keep names under 19 chars: "terr-ecs1-"+5=16, "terr-ecs2-"+5=16, "terr-aws-p-"+5=16
	nameFirst := fmt.Sprintf("terr-ecs1-%s", randAlnum(5))
	name := fmt.Sprintf("terr-ecs2-%s", randAlnum(5))
	proxyName := fmt.Sprintf("terr-aws-p-%s", randAlnum(5))
	apiKey := os.Getenv(platformAPIKeyEnv)

	resourceName := "harness_autostopping_rule_ecs.test"
	resourceNameFirst := "harness_autostopping_rule_ecs.first"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testECSRuleFirstOnly(nameFirst, proxyName, apiKey, true),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceNameFirst, "name", nameFirst),
					resource.TestCheckResourceAttr(resourceNameFirst, "dry_run", "true"),
				),
			},
			{
				Config: testECSRuleWithDependency(nameFirst, name, proxyName, apiKey, true),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceNameFirst, "name", nameFirst),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "dry_run", "true"),
				),
			},
			{
				Config: testECSRuleWithDependency(nameFirst, name, proxyName, apiKey, false),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "dry_run", "false"),
				),
			},
			{
				Config: testECSRuleUpdate(nameFirst, name, proxyName, apiKey, "15", true),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "idle_time_mins", "15"),
					resource.TestCheckResourceAttr(resourceName, "dry_run", "true"),
				),
			},
			{
				Config: testECSRuleUpdate(nameFirst, name, proxyName, apiKey, "20", false),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "idle_time_mins", "20"),
					resource.TestCheckResourceAttr(resourceName, "dry_run", "false"),
				),
			},
		},
	})
}

// testECSAWSProxy returns config for AWS proxy (same pattern as TestResourceAWSProxy).
func testECSAWSProxy(proxyName, apiKey string) string {
	return fmt.Sprintf(`
resource "harness_autostopping_aws_proxy" "test" {
  name                 = %[1]q
  cloud_connector_id   = %[2]q
  region               = "us-east-1"
  vpc                  = "vpc-0d47ab08fce6d8cc8"
  security_groups      = ["sg-0a2a6eaa3ad797636"]
  machine_type         = "t2.medium"
  api_key              = %[3]q
  allocate_static_ip   = true
  delete_cloud_resources_on_destroy = false
}
`, proxyName, cloudConnectorIDVM, apiKey)
}

// testECSRuleFirstOnly creates AWS proxy + one ECS rule without depends block.
func testECSRuleFirstOnly(nameFirst, proxyName, apiKey string, dryRun bool) string {
	return testECSAWSProxy(proxyName, apiKey) + fmt.Sprintf(`
resource "harness_autostopping_rule_ecs" "first" {
  name                = %[1]q
  cloud_connector_id  = %[2]q
  idle_time_mins      = 10
  dry_run             = %[3]t

  container {
    cluster    = "cluster"
    service    = "service"
    region     = "us-east-1"
    task_count = 1
  }
  http {
    proxy_id = harness_autostopping_aws_proxy.test.identifier
  }
}
`, nameFirst, cloudConnectorIDVM, dryRun)
}

// testECSRuleWithDependency creates proxy + first ECS rule (no depends) + second rule depending on first.
func testECSRuleWithDependency(nameFirst, name, proxyName, apiKey string, dryRun bool) string {
	return testECSRuleFirstOnly(nameFirst, proxyName, apiKey, dryRun) + fmt.Sprintf(`
resource "harness_autostopping_rule_ecs" "test" {
  name                = %[1]q
  cloud_connector_id  = %[2]q
  idle_time_mins      = 10
  dry_run             = %[3]t

  container {
    cluster    = "cluster"
    service    = "service"
    region     = "us-east-1"
    task_count = 1
  }
  http {
    proxy_id = harness_autostopping_aws_proxy.test.identifier
  }
  depends {
    rule_id     = harness_autostopping_rule_ecs.first.identifier
    delay_in_sec = 5
  }
}
`, name, cloudConnectorIDVM, dryRun)
}

func testECSRuleUpdate(nameFirst, name, proxyName, apiKey, idleTime string, dryRun bool) string {
	return testECSRuleFirstOnly(nameFirst, proxyName, apiKey, dryRun) + fmt.Sprintf(`
resource "harness_autostopping_rule_ecs" "test" {
  name                = %[1]q
  cloud_connector_id  = %[2]q
  idle_time_mins      = %[3]s
  dry_run             = %[4]t

  container {
    cluster    = "cluster"
    service    = "service"
    region     = "us-east-1"
    task_count = 1
  }
  http {
    proxy_id = harness_autostopping_aws_proxy.test.identifier
  }
  depends {
    rule_id     = harness_autostopping_rule_ecs.first.identifier
    delay_in_sec = 5
  }
}
`, name, cloudConnectorIDVM, idleTime, dryRun)
}
