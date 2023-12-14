terraform {
  required_providers {
    harness = {
      source = "harness/harness"
    }
  }
}

import {
  to = harness_platform_secret_text.example
  id = "test_tf_secret_10"
}

resource "harness_platform_secret_text" "example" {
  identifier  = "test_tf_secret_10"
}
