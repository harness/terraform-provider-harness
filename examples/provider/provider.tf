terraform {
  required_providers {
    harness = {
      source = "harness/harness"
    }
  }
}

provider "harness" {
  endpoint   = "https://app.harness.io/gateway"
  account_id = "UKh5Yts7THSMAbccG3HrLA"
}

resource "harness_user_group" "terraform" {
  name        = "terraform-group"
  description = "This group is managed by Terraform"
  permissions {
    app_permissions {
      deployment {
        actions = ["READ"]
        filters = ["NON_PRODUCTION_ENVIRONMENTS", "PRODUCTION_ENVIRONMENTS"]
      }
      pipeline {
        actions = ["READ"]
        filters = ["NON_PRODUCTION_ENVIRONMENTS", "PRODUCTION_ENVIRONMENTS"]
      }
    }
  }
}
