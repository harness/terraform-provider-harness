terraform {
  required_providers {
    harness = {
      version = "0.2"
      source  = "hashicorp.com/micahlmartin/harness"
    }
  }
}
provider "harness" {
  # example configuration here
}
