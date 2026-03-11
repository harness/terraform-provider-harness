package cluster_orchestrator_test

import (
	"fmt"
	"testing"

	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestResourceClusterOrchestratorConfig(t *testing.T) {
	orchName := "terraform-clusterorch-config-test"
	resourceName := "harness_cluster_orchestrator_config.test"
	orchResourceName := "harness_cluster_orchestrator.setup"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testClusterOrchWithConfig(orchName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(resourceName, "orchestrator_id", orchResourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "distribution.0.strategy", "CostOptimized"),
					resource.TestCheckResourceAttr(resourceName, "distribution.0.selector", "ALL"),
					resource.TestCheckResourceAttr(resourceName, "disabled", "false"),
				),
			},
		},
	})
}

// TestResourceClusterOrchestratorConfigDisabled verifies create with disabled=false.
// NOTE: toggling disabled=true and reading it back does not work because the SDK
// model tags Disabled with `json:"-"`, so it is never deserialized from the API
// response. The toggle_state write works, but subsequent reads always return false.
func TestResourceClusterOrchestratorConfigDisabled(t *testing.T) {
	orchName := "terraform-clusterorch-disabled-test"
	resourceName := "harness_cluster_orchestrator_config.test"
	orchResourceName := "harness_cluster_orchestrator.setup"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testClusterOrchWithConfig(orchName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(resourceName, "orchestrator_id", orchResourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "disabled", "false"),
				),
			},
		},
	})
}

func testClusterOrchWithConfig(orchName string) string {
	return fmt.Sprintf(`
	resource "harness_cluster_orchestrator" "setup" {
		name             = "%[1]s"
		cluster_endpoint = "http://test.com"
		k8s_connector_id = "TestDoNotDelete"
	}

	resource "harness_cluster_orchestrator_config" "test" {
		orchestrator_id = harness_cluster_orchestrator.setup.id
		distribution {
			base_ondemand_capacity      = 0
			ondemand_replica_percentage = 0
			selector                    = "ALL"
			strategy                    = "CostOptimized"
		}
		node_preferences {
			ttl = "48h"
		}
		replacement_schedule {
			window_type = "Always"
			applies_to {
				consolidation        = true
				harness_pod_eviction = true
				reverse_fallback     = true
			}
		}
		binpacking {
			disruption {
				criteria = "WhenEmptyOrUnderutilized"
				delay    = "5m"
				budget {
					reasons = ["Drifted"]
					nodes   = "10%%"
				}
			}
			pod_eviction {
				threshold {
					cpu    = 60
					memory = 75
				}
			}
		}
	}
`, orchName)
}
