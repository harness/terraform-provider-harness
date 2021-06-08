terraform {
  required_providers {
    harness = {
      source  = "micahlmartin/harness"
    }
  }
}

provider "harness" {}


resource "harness_application" "my_app" {
  name = "my_app"
  description = "updated"
}

