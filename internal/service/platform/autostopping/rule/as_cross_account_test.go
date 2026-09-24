/*
Cross-Account AutoStopping Rule Acceptance Tests
=================================================

These tests verify that AutoStopping rules can reference a proxy (access
point) that lives in a different cloud account than the target resource.
The proxy's cloud_account_id field (proxy_cloud_connector_id in Terraform)
tells the Harness backend which connector owns the proxy.

Supported resource types:

  Cross-account proxy AutoStopping is ONLY supported for VM rules.
  It is NOT available for ECS, RDS, or Scale Group rules.

Test coverage:

  TestResourceVMRuleCrossAccountHTTP  VM rule with cross-account HTTP proxy
  TestResourceVMRuleCrossAccountTCP   VM rule with cross-account TCP proxy

Each test creates the rule (dry_run=true), verifies proxy_cloud_connector_id
survives the round-trip, then updates (dry_run=false) and verifies again.
If any required env var is missing, the test is skipped (not failed).
*/

/*
EC2 VM Cross-Account Tests
===========================

Tests: TestResourceVMRuleCrossAccountHTTP
       TestResourceVMRuleCrossAccountTCP

Environment variables:

  export HARNESS_ACCOUNT_ID="<harness-account-id>"
  export HARNESS_PLATFORM_API_KEY="<harness-api-key>"
  export CROSS_ACCOUNT_TARGET_CONNECTOR_ID="<connector-for-target-account-B>"
  export CROSS_ACCOUNT_TARGET_VM_ID="<i-xxxxx-ec2-instance-in-account-B>"
  export CROSS_ACCOUNT_TARGET_REGION="<aws-region-of-target>"
  export CROSS_ACCOUNT_PROXY_CONNECTOR_ID="<connector-for-proxy-account-A>"
  export CROSS_ACCOUNT_ACCESS_POINT_ID="<ap-xxxxx-proxy-in-account-A>"

Run all cross-account tests:

  TF_ACC=1 go test -v -run TestResourceVMRuleCrossAccount -count=1 \
    ./internal/service/platform/autostopping/rule/ -timeout=30m

Run only HTTP:

  TF_ACC=1 go test -v -run TestResourceVMRuleCrossAccountHTTP -count=1 \
    ./internal/service/platform/autostopping/rule/ -timeout=30m

Run only TCP:

  TF_ACC=1 go test -v -run TestResourceVMRuleCrossAccountTCP -count=1 \
    ./internal/service/platform/autostopping/rule/ -timeout=30m
*/

package as_rule_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

// --- Environment variable constants ---

const (
	crossAccountProxyConnectorIDEnv = "CROSS_ACCOUNT_PROXY_CONNECTOR_ID"
	crossAccountAccessPointIDEnv    = "CROSS_ACCOUNT_ACCESS_POINT_ID"
	crossAccountTargetConnectorEnv  = "CROSS_ACCOUNT_TARGET_CONNECTOR_ID"
	crossAccountTargetRegionEnv     = "CROSS_ACCOUNT_TARGET_REGION"
	crossAccountTargetVMIDEnv       = "CROSS_ACCOUNT_TARGET_VM_ID"
)

type crossAccountVMEnv struct {
	proxyConnectorID  string
	accessPointID     string
	targetConnectorID string
	targetRegion      string
	targetVMID        string
}

func crossAccountVMSetup(t *testing.T) crossAccountVMEnv {
	t.Helper()
	env := crossAccountVMEnv{
		proxyConnectorID:  os.Getenv(crossAccountProxyConnectorIDEnv),
		accessPointID:     os.Getenv(crossAccountAccessPointIDEnv),
		targetConnectorID: os.Getenv(crossAccountTargetConnectorEnv),
		targetRegion:      os.Getenv(crossAccountTargetRegionEnv),
		targetVMID:        os.Getenv(crossAccountTargetVMIDEnv),
	}
	if env.proxyConnectorID == "" || env.accessPointID == "" ||
		env.targetConnectorID == "" || env.targetRegion == "" || env.targetVMID == "" {
		t.Skipf("Skipping: set %s, %s, %s, %s, %s",
			crossAccountProxyConnectorIDEnv, crossAccountAccessPointIDEnv,
			crossAccountTargetConnectorEnv, crossAccountTargetRegionEnv,
			crossAccountTargetVMIDEnv)
	}
	return env
}

// ===========================================================================
// EC2 VM Tests
// ===========================================================================

func TestResourceVMRuleCrossAccountHTTP(t *testing.T) {
	env := crossAccountVMSetup(t)
	name := fmt.Sprintf("terr-xvm-h-%s", randAlnum(5))
	resourceName := "harness_autostopping_rule_vm.xacct"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testVMCrossAccountHTTP(name, env, true),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "cloud_connector_id", env.targetConnectorID),
					resource.TestCheckResourceAttr(resourceName, "http.0.proxy_id", env.accessPointID),
					resource.TestCheckResourceAttr(resourceName, "http.0.proxy_cloud_connector_id", env.proxyConnectorID),
					resource.TestCheckResourceAttr(resourceName, "dry_run", "true"),
				),
			},
			{
				Config: testVMCrossAccountHTTP(name, env, false),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "http.0.proxy_cloud_connector_id", env.proxyConnectorID),
					resource.TestCheckResourceAttr(resourceName, "dry_run", "false"),
				),
			},
		},
	})
}

func TestResourceVMRuleCrossAccountTCP(t *testing.T) {
	env := crossAccountVMSetup(t)
	name := fmt.Sprintf("terr-xvm-t-%s", randAlnum(5))
	resourceName := "harness_autostopping_rule_vm.xacct"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testVMCrossAccountTCP(name, env, true),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "cloud_connector_id", env.targetConnectorID),
					resource.TestCheckResourceAttr(resourceName, "tcp.0.proxy_id", env.accessPointID),
					resource.TestCheckResourceAttr(resourceName, "tcp.0.proxy_cloud_connector_id", env.proxyConnectorID),
					resource.TestCheckResourceAttr(resourceName, "dry_run", "true"),
				),
			},
			{
				Config: testVMCrossAccountTCP(name, env, false),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "tcp.0.proxy_cloud_connector_id", env.proxyConnectorID),
					resource.TestCheckResourceAttr(resourceName, "dry_run", "false"),
				),
			},
		},
	})
}

// ===========================================================================
// Terraform config generators
// ===========================================================================

func testVMCrossAccountHTTP(name string, env crossAccountVMEnv, dryRun bool) string {
	return fmt.Sprintf(`
resource "harness_autostopping_rule_vm" "xacct" {
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
    proxy_id                 = %[6]q
    proxy_cloud_connector_id = %[7]q
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
}
`, name, env.targetConnectorID, dryRun, env.targetVMID, env.targetRegion, env.accessPointID, env.proxyConnectorID)
}

func testVMCrossAccountTCP(name string, env crossAccountVMEnv, dryRun bool) string {
	return fmt.Sprintf(`
resource "harness_autostopping_rule_vm" "xacct" {
  name               = %[1]q
  cloud_connector_id = %[2]q
  idle_time_mins     = 10
  dry_run            = %[3]t

  filter {
    vm_ids  = [%[4]q]
    regions = [%[5]q]
    zones   = []
  }
  tcp {
    proxy_id                 = %[6]q
    proxy_cloud_connector_id = %[7]q
    ssh {
      connect_on = 2222
      port       = 22
    }
  }
}
`, name, env.targetConnectorID, dryRun, env.targetVMID, env.targetRegion, env.accessPointID, env.proxyConnectorID)
}
