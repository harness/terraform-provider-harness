package as_rule_test

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

// TestAccResourceECSRule_CCM32010_UpdateInPlaceAddsDepends verifies that an existing
// autostopping ECS rule can gain a depends entry without the API rejecting the update
// as a duplicate create (regression for the class of bug fixed under CCM-32010 for
// governance rule sets: "already exists" on PUT when the name is unchanged).
func TestAccResourceECSRule_CCM32010_UpdateInPlaceAddsDepends(t *testing.T) {
	nameFirst := fmt.Sprintf("terr-c10a-%s", randAlnum(5))
	nameSecond := fmt.Sprintf("terr-c10b-%s", randAlnum(5))
	proxyName := fmt.Sprintf("terr-c10p-%s", randAlnum(5))
	apiKey := os.Getenv(platformAPIKeyEnv)

	var proxyID string
	resourceName := "harness_autostopping_rule_ecs.second"
	resourceNameFirst := "harness_autostopping_rule_ecs.first"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testECSCCM32010Step1NoDepends(nameFirst, nameSecond, proxyName, apiKey, true),
				Check: resource.ComposeTestCheckFunc(
					extractAttr("harness_autostopping_aws_proxy.test", "identifier", &proxyID),
					resource.TestCheckResourceAttr(resourceNameFirst, "name", nameFirst),
					resource.TestCheckResourceAttr(resourceName, "name", nameSecond),
				),
			},
			{
				PreConfig: func() {
					if err := waitForProxyReady(proxyID, 90*time.Second); err != nil {
						t.Fatalf("Proxy not ready: %v", err)
					}
				},
				Config: testECSCCM32010Step2WithDepends(nameFirst, nameSecond, proxyName, apiKey, true),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", nameSecond),
					resource.TestCheckResourceAttrPair(resourceName, "depends.0.rule_id", resourceNameFirst, "identifier"),
				),
			},
			{
				Config:   testECSCCM32010Step2WithDepends(nameFirst, nameSecond, proxyName, apiKey, true),
				PlanOnly: true,
			},
		},
	})
}

// TestAccResourceECSRule_CCM32012_DependsOrderNoPerpetualPlan verifies that multiple
// depends blocks do not cause perpetual plan drift when the API returns dependencies
// in a different order than configuration (regression for CCM-32012-style rule_id list
// ordering issues in governance).
func TestAccResourceECSRule_CCM32012_DependsOrderNoPerpetualPlan(t *testing.T) {
	nameA := fmt.Sprintf("terr-c12a-%s", randAlnum(5))
	nameB := fmt.Sprintf("terr-c12b-%s", randAlnum(5))
	nameC := fmt.Sprintf("terr-c12c-%s", randAlnum(5))
	proxyName := fmt.Sprintf("terr-c12p-%s", randAlnum(5))
	apiKey := os.Getenv(platformAPIKeyEnv)

	var proxyID string
	resourceName := "harness_autostopping_rule_ecs.leaf"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testECSCCM32012Stack(nameA, nameB, nameC, proxyName, apiKey, true),
				Check: resource.ComposeTestCheckFunc(
					extractAttr("harness_autostopping_aws_proxy.test", "identifier", &proxyID),
					resource.TestCheckResourceAttr(resourceName, "name", nameC),
				),
			},
			{
				PreConfig: func() {
					if err := waitForProxyReady(proxyID, 90*time.Second); err != nil {
						t.Fatalf("Proxy not ready: %v", err)
					}
				},
				Config: testECSCCM32012Stack(nameA, nameB, nameC, proxyName, apiKey, false),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", nameC),
					resource.TestCheckResourceAttr(resourceName, "depends.#", "2"),
				),
			},
			{
				Config:   testECSCCM32012Stack(nameA, nameB, nameC, proxyName, apiKey, false),
				PlanOnly: true,
			},
		},
	})
}

func testECSCCM32010Step1NoDepends(nameFirst, nameSecond, proxyName, apiKey string, dryRun bool) string {
	return testECSAWSProxy(proxyName, apiKey) + fmt.Sprintf(`
resource "harness_autostopping_rule_ecs" "first" {
  name               = %[1]q
  cloud_connector_id = %[3]q
  idle_time_mins     = 10
  dry_run            = %[4]t

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

resource "harness_autostopping_rule_ecs" "second" {
  name               = %[2]q
  cloud_connector_id = %[3]q
  idle_time_mins     = 10
  dry_run            = %[4]t

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
`, nameFirst, nameSecond, cloudConnectorIDAWS, dryRun)
}

func testECSCCM32010Step2WithDepends(nameFirst, nameSecond, proxyName, apiKey string, dryRun bool) string {
	return testECSAWSProxy(proxyName, apiKey) + fmt.Sprintf(`
resource "harness_autostopping_rule_ecs" "first" {
  name               = %[1]q
  cloud_connector_id = %[3]q
  idle_time_mins     = 10
  dry_run            = %[4]t

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

resource "harness_autostopping_rule_ecs" "second" {
  name               = %[2]q
  cloud_connector_id = %[3]q
  idle_time_mins     = 10
  dry_run            = %[4]t

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
    rule_id      = harness_autostopping_rule_ecs.first.identifier
    delay_in_sec = 5
  }
}
`, nameFirst, nameSecond, cloudConnectorIDAWS, dryRun)
}

func testECSCCM32012Stack(nameA, nameB, nameC, proxyName, apiKey string, dryRun bool) string {
	return testECSAWSProxy(proxyName, apiKey) + fmt.Sprintf(`
resource "harness_autostopping_rule_ecs" "deps_a" {
  name               = %[1]q
  cloud_connector_id = %[4]q
  idle_time_mins     = 10
  dry_run            = %[5]t

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

resource "harness_autostopping_rule_ecs" "deps_b" {
  name               = %[2]q
  cloud_connector_id = %[4]q
  idle_time_mins     = 10
  dry_run            = %[5]t

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

resource "harness_autostopping_rule_ecs" "leaf" {
  name               = %[3]q
  cloud_connector_id = %[4]q
  idle_time_mins     = 10
  dry_run            = %[5]t

  container {
    cluster    = "cluster"
    service    = "service"
    region     = "us-east-1"
    task_count = 1
  }
  http {
    proxy_id = harness_autostopping_aws_proxy.test.identifier
  }

  # Intentionally order depends as deps_b then deps_a so a CCM-32012-style API sort
  # by rule id would reorder relative to many configs and surface as perpetual drift.
  depends {
    rule_id      = harness_autostopping_rule_ecs.deps_b.identifier
    delay_in_sec = 5
  }
  depends {
    rule_id      = harness_autostopping_rule_ecs.deps_a.identifier
    delay_in_sec = 7
  }
}
`, nameA, nameB, nameC, cloudConnectorIDAWS, dryRun)
}
