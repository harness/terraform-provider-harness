terraform {
    required_providers {
      harness = {
        source = "terraform.local/local/harness"
        version = "0.2"
      }
    }
  }

  provider "harness" {
    endpoint         = "https://stress.harness.io"
    account_id       = "-k53qRQAQ1O7DBLb9ACnjQ"
    platform_api_key = ""
  }
