terraform {
  required_providers {
    harness = {
      source = "micahlmartin/harness"
    }
  }
}

provider "harness" {}

# data "harness_application" "example" {
#   id = "foo"
# }

data "harness_encrypted_text" "example" {
  name = "somesecret"
}

output "test" {
  value = data.harness_encrypted_text.example.usage_scopes[0].application_id
}
