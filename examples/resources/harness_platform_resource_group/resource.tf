terraform {
  required_providers {
    harness = {
      source = "harness/harness"
    }
  }
}

resource "harness_platform_resource_group" "test" {
  identifier  = "fromTwF"
  name        = "fromTwF"
  description = "test"
  tags        = ["foo:bar"]

  org_id =  "default"
  project_id = "ResourceGroupTest"
  account_id = "Cdrf1gJ7RQOqMEGWV_5FTA"
  allowed_scope_levels = ["project"]
  included_scopes {
    filter     = "EXCLUDING_CHILD_SCOPES"
    account_id = "Cdrf1gJ7RQOqMEGWV_5FTA"
    org_id =  "default"
    project_id = "ResourceGroupTest"
  }
  resource_filter {
    include_all_resources = false
    resources {
      resource_type = "ENVIRONMENT"
      identifiers = [
          "invalid",
          "ok-bye"
      ]
    }
  }
}
