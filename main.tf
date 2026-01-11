data "harness_platform_workspace" "test_list_all" {
  org_id     = "Platform_Team"
  project_id = "infrastructure"
}



terraform {
  required_providers {
    harness = {
      source = "harness/harness"
    }
  }
}

#Configure the Harness provider for Next Gen resources
provider "harness" {
  endpoint         = "https://app.harness.io/gateway"
  account_id       = "PaWpk1biSb2DvH2qUZEcuQ"
  platform_api_key = "pat.PaWpk1biSb2DvH2qUZEcuQ.696299453c2f12162cc1a0d6.xCQqoHhHeg4w2AJaQtS6"
}
