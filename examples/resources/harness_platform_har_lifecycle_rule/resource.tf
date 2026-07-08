# Account-scoped DELETE rule — keep last 10 versions, runs nightly
resource "harness_platform_har_lifecycle_rule" "nightly_cleanup" {
  account_id  = "your-account-id"
  name        = "nightly-cleanup"
  action      = "DELETE"
  description = "Keep last 10 versions of all artifacts"

  apply_to {
    mode = "ALL_IN_SCOPE"
  }

  criteria {
    match = "ALL"
    rules {
      type  = "KEEP_LAST_N"
      value = 10
    }
  }

  schedule {
    expression = "0 2 * * *"
    timezone   = "UTC"
  }
}

# Project-scoped DELETE rule — delete artifacts older than 30 days
resource "harness_platform_har_lifecycle_rule" "age_based_cleanup" {
  account_id = "your-account-id"
  org_id     = "your-org-id"
  project_id = "your-project-id"
  name       = "age-based-cleanup"
  action     = "DELETE"

  apply_to {
    mode = "ALL_IN_SCOPE"
  }

  criteria {
    match = "ALL"
    rules {
      type  = "AGE_BASED"
      value = 30
      unit  = "DAYS"
    }
  }
}

# Org-scoped PROTECT rule — protect images in specific registries matching a tag pattern
resource "harness_platform_har_lifecycle_rule" "protect_prod" {
  account_id   = "your-account-id"
  org_id       = "your-org-id"
  name         = "protect-prod-images"
  action       = "PROTECT"
  package_type = "DOCKER"

  apply_to {
    mode       = "EXPLICIT"
    registries = ["prod-registry", "release-registry"]
  }

  filter_config {
    package_type             = "DOCKER"
    tag_name_allowed_pattern = ["v*", "release-*"]
  }
}

# Account-scoped DELETE rule with multiple criteria (ANY match)
resource "harness_platform_har_lifecycle_rule" "multi_criteria_cleanup" {
  account_id = "your-account-id"
  name       = "multi-criteria-cleanup"
  action     = "DELETE"

  apply_to {
    mode = "ALL_IN_SCOPE"
  }

  criteria {
    match = "ANY"
    rules {
      type  = "KEEP_LAST_N"
      value = 5
    }
    rules {
      type  = "AGE_BASED"
      value = 90
      unit  = "DAYS"
    }
  }
}
