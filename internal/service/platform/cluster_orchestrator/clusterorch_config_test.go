package cluster_orchestrator_test

import (
	"fmt"
	"testing"

	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestResourceClusterOrchestratorConfig(t *testing.T) {
	orchID := "terraform-clusterorch-config-test"
	resourceName := "harness_cluster_orchestrator_config.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testClusterOrchConfig(orchID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "orchID", orchID),
				),
			},
		},
	})
}

func testClusterOrchConfig(orchID string) string {
	return fmt.Sprintf(`
	resource "harness_cluster_orchestrator_config" "test" {
		orchestrator_id = "%s"
		distribution {
			base_ondemand_capacity = 1
            ondemand_replica_percentage = 60
            selector = "ALL"
            strategy = "COST OPTIMIZED"
		}
		binpacking {
            pod_eviction {
				threshold {
					cpu = 60
					memory = 80
				}
			}
			disruption {
				criteria = "EmptyOrUnderUtilized"
                delay = "10m"
                budget {
                	reasons = ["Drift","UnderUtilized","Empty"]
                    nodes = "20"
					schedule {
						frequency = "@daily"
						duration = "10m"
					}
                }
				budget {
                	reasons = ["Drift","Empty"]
                    nodes = "1"
					schedule {
						frequency = "@monthly"
						duration = "10m"
					}
                }
			}
		}
		node_preferences {
			ttl = "1h"
            reverse_fallback_interval = "6h"
		} 
	}
`, orchID)
}
