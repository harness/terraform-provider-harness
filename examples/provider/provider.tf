terraform {
  required_providers {
    harness = {
      source  = "micahlmartin/harness"
    }
  }
}
provider "harness" {
  api_key = "VUtoNVl0czdUSFNNQWJjY0czSHJMQTo6R0dERHRQNTNZWXNSWUVBWGFNSDV3ZUdmaFhuSWhLWVNXamY5YzZOMGRlcmZVMVpFM3hmQzdLVENjdFNOVHlUUXRlbnpKMUtQVFpmcGxEekk="
  account_id = "UKh5Yts7THSMAbccG3HrLA"
  # example configuration here
}

data "harness_application" "foo" {
  id = "GEHhvKUCTiiY_MWsUfbRLA"
  # name = "changed"
}

output "app_id" {
  value = data.harness_application.foo.id
}

# output "app_name" {
#   value = data.harness_application.foo.name
# }
