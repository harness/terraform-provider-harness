resource "harness_cluster_orchestrator_config" "example" {
  orchestrator_id = "orch-cvifpfl9rbg8neldj97g"
  disabled        = false # Set to true to disable the orchestrator
  distribution {
    base_ondemand_capacity      = 2
    ondemand_replica_percentage = 50
    selector                    = "ALL"
    strategy                    = "CostOptimized"
  }
  binpacking {
    pod_eviction {
      threshold {
        cpu    = 60
        memory = 80
      }
    }
    disruption {
      criteria = "WhenEmpty"
      delay    = "10m"
      budget {
        reasons = ["Drifted", "Underutilized", "Empty"]
        nodes   = "20"
      }
      budget {
        reasons = ["Drifted", "Empty"]
        nodes   = "1"
        schedule {
          frequency = "@monthly"
          duration  = "10m"
        }
      }
    }
  }
  node_preferences {
    ttl                       = "Never"
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