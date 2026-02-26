package cluster_orchestrator_test

import (
	"fmt"
	"testing"

	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

const clusterOrchConfigTestName = "terraform-clusterorch-config-test"
const clusterOrchConfigDisabledTestName = "terraform-clusterorch-disabled-test"

func TestResourceClusterOrchestratorConfig(t *testing.T) {
	resourceName := "harness_cluster_orchestrator_config.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testClusterOrchConfig(clusterOrchConfigTestName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "orchestrator_id"),
					resource.TestCheckResourceAttr(resourceName, "disabled", "false"),
				),
			},
		},
	})
}

func TestResourceClusterOrchestratorConfigDisabled(t *testing.T) {
	resourceName := "harness_cluster_orchestrator_config.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testClusterOrchConfigDisabled(clusterOrchConfigDisabledTestName, false),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "orchestrator_id"),
					resource.TestCheckResourceAttr(resourceName, "disabled", "false"),
				),
			},
			{
				Config: testClusterOrchConfigDisabled(clusterOrchConfigDisabledTestName, true),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "orchestrator_id"),
					resource.TestCheckResourceAttr(resourceName, "disabled", "true"),
				),
			},
			{
				Config: testClusterOrchConfigDisabled(clusterOrchConfigDisabledTestName, false),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "orchestrator_id"),
					resource.TestCheckResourceAttr(resourceName, "disabled", "false"),
				),
			},
		},
	})
}

func testClusterOrchConfig(orchName string) string {
	return fmt.Sprintf(`
	resource "harness_cluster_orchestrator" "orch" {
		name              = "%[1]s"
		cluster_endpoint  = "http://test.com"
		k8s_connector_id  = "TestDoNotDelete"
	}

	resource "harness_cluster_orchestrator_config" "test" {
		orchestrator_id = harness_cluster_orchestrator.orch.id
		disabled        = false
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
		commitment_integration {
			enabled           = true
			master_account_id = "dummyAccountId"
		}
		replacement_schedule {
			window_type = "Custom"
			applies_to {
			  consolidation        = true
			  harness_pod_eviction = true
			  reverse_fallback     = true
			}
			window_details {
			  days       = ["SUN", "WED", "SAT"]
			  time_zone  = "Asia/Calcutta"
			  all_day    = false
			  start_time = "10:30"
			  end_time   = "11:30"
			}
		}
	}
`, orchName)
}

func testClusterOrchConfigDisabled(orchName string, disabled bool) string {
	return fmt.Sprintf(`
	resource "harness_cluster_orchestrator" "orch" {
		name              = "%[1]s"
		cluster_endpoint  = "http://test.com"
		k8s_connector_id  = "TestDoNotDelete"
	}

	resource "harness_cluster_orchestrator_config" "test" {
		orchestrator_id = harness_cluster_orchestrator.orch.id
		disabled        = %t
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
		commitment_integration {
			enabled           = true
			master_account_id = "dummyAccountId"
		}
		replacement_schedule {
			window_type = "Custom"
			applies_to {
			  consolidation        = true
			  harness_pod_eviction = true
			  reverse_fallback     = true
			}
			window_details {
			  days       = ["SUN", "WED", "SAT"]
			  time_zone  = "Asia/Calcutta"
			  all_day    = false
			  start_time = "10:30"
			  end_time   = "11:30"
			}
		}
	}
`, orchName, disabled)
}
