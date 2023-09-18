// Boolean Flag
resource "harness_platform_feature_flag" "mybooleanflag" {
  org_id     = "test"
  project_id = "testff"

  kind       = "boolean"
  name       = "MY_FEATURE"
  identifier = "MY_FEATURE"
  permanent  = false

  default_on_variation  = "Enabled"
  default_off_variation = "Disabled"

  variation {
    identifier  = "Enabled"
    name        = "Enabled"
    description = "The feature is enabled"
    value       = "true"
  }

  variation {
    identifier  = "Disabled"
    name        = "Disabled"
    description = "The feature is disabled"
    value       = "false"
  }
}


// Multivariate flag
resource "harness_platform_feature_flag" "mymultivariateflag" {
  org_id     = "test"
  project_id = "testff"

  kind       = "int"
  name       = "FREE_TRIAL_DURATION"
  identifier = "FREE_TRIAL_DURATION"
  permanent  = false

  default_on_variation  = "trial7"
  default_off_variation = "trial20"

  variation {
    identifier  = "trial7"
    name        = "7 days trial"
    description = "Free trial period 7 days"
    value       = "7"
  }

  variation {
    identifier  = "trial14"
    name        = "14 days trial"
    description = "Free trial period 14 days"
    value       = "14"
  }

  variation {
    identifier  = "trial20"
    name        = "20 days trial"
    description = "Free trial period 20 days"
    value       = "20"
  }
}

// Assign targets to flag
resource "harness_platform_feature_flag" "mymultivariateflag" {
  org_id     = "test"
  project_id = "testff"

  kind       = "int"
  name       = "FREE_TRIAL_DURATION"
  identifier = "FREE_TRIAL_DURATION"
  permanent  = false

  default_on_variation  = "trial7"
  default_off_variation = "trial20"

  variation {
    identifier  = "trial7"
    name        = "7 days trial"
    description = "Free trial period 7 days"
    value       = "7"
  }

  variation {
    identifier  = "trial14"
    name        = "14 days trial"
    description = "Free trial period 14 days"
    value       = "14"
  }

  variation {
    identifier  = "trial20"
    name        = "20 days trial"
    description = "Free trial period 20 days"
    value       = "20"
  }

  add_target_rule {
    variation = "trial14"
    targets = ["targets1", "targets2"]
  }
}

// Assign target groups to flag
resource "harness_platform_feature_flag" "mymultivariateflag" {
  org_id     = "test"
  project_id = "testff"

  kind       = "int"
  name       = "FREE_TRIAL_DURATION"
  identifier = "FREE_TRIAL_DURATION"
  permanent  = false

  default_on_variation  = "trial7"
  default_off_variation = "trial20"

  variation {
    identifier  = "trial7"
    name        = "7 days trial"
    description = "Free trial period 7 days"
    value       = "7"
  }

  variation {
    identifier  = "trial14"
    name        = "14 days trial"
    description = "Free trial period 14 days"
    value       = "14"
  }

  variation {
    identifier  = "trial20"
    name        = "20 days trial"
    description = "Free trial period 20 days"
    value       = "20"
  }

  add_target_groups_rule {
    group_name = "group_name"
    variation = "trial14"
    distribution = {
      variations = [
        {
            variation = "trial7"
            weight = 30
        },
        {
            variation = "trial14"
            weight = 30
        },
        {
            variation = "trial20"
            weight = 40
        }
      ]
    }
  }
}