terraform {
  required_providers {
    harness = {
      source = "harness-io/harness"
    }
  }
}
provider "harness" {
  endpoint   = "https://app.harness.io"
  account_id = "...."
  api_key    = "......"
}

