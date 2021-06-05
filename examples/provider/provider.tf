terraform {
  required_providers {
    harness = {
      source  = "micahlmartin/harness"
    }
  }
}
provider "harness" {
  api_key = ""
  account_id = ""
  # example configuration here
}

data "harness_application" "foo" {
  id = ""
  # name = "changed"
}

output "app_id" {
  value = data.harness_application.foo.id
}

# output "app_name" {
#   value = data.harness_application.foo.name
# }
