// K8s Rule
resource "harness_chaos_security_governance_rule" "example" {  
  org_id        = "<org_id>"
  project_id    = "<project_id>"
  name          = "<name>"
  description   = "<description>"
  is_enabled    = true
  condition_ids = ["<condition_id>"]
  user_group_ids = ["_project_all_users"]
  tags          = ["<tag1>", "<tag2>"]

  time_windows {
    time_zone  = "UTC"
    start_time = 1711238400000
    duration   = "24h"
    
    recurrence {
      type  = "Daily"
      until = -1
    }
  }
}

// Linux Rule
resource "harness_chaos_security_governance_rule" "linux_rule" {
  org_id        = "<org_id>"
  project_id    = "<project_id>"
  name          = "<name>"
  description   = "<description>"
  is_enabled    = true
  condition_ids = ["<condition_id>"]
  user_group_ids = ["_project_all_users"]
  tags          = ["<tag1>", "<tag2>"]

  time_windows {
    time_zone  = "UTC"
    start_time = 1711238400000
    duration   = "24h"
    
    recurrence {
      type  = "Daily"
      until = -1
    }
  }
}

// Windows Rule
resource "harness_chaos_security_governance_rule" "windows_rule" {  
  org_id        = "<org_id>"
  project_id    = "<project_id>"
  name          = "<name>"
  description   = "<description>"
  is_enabled    = true
  condition_ids = ["<condition_id>"]
  user_group_ids = ["_project_all_users"]
  tags          = ["<tag1>", "<tag2>"]

  time_windows {
    time_zone  = "UTC"
    start_time = 1711238400000
    duration   = "24h"
    
    recurrence {
      type  = "Daily"
      until = -1
    }
  }
}
