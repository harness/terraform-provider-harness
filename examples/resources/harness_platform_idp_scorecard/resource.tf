resource "harness_platform_idp_scorecard_check" "readme" {
  identifier        = "readme_exists"
  name              = "README exists"
  description       = "Ensure the repository has a README file"
  rule_strategy     = "ALL_OF"
  default_behaviour = "FAIL"

  rules {
    data_source_identifier = "github"
    data_point_identifier  = "isFileExists"
    operator               = "=="
    value                  = "true"

    input_values {
      key   = "filePath"
      value = "README.md"
    }
  }
}

resource "harness_platform_idp_scorecard" "gold" {
  identifier         = "gold_standard"
  name               = "Gold Standard"
  description        = "Baseline production quality scorecard"
  published          = true
  weightage_strategy = "EQUAL_WEIGHTS"

  filter {
    kind = "component"
    type = "service"
  }

  checks {
    identifier = harness_platform_idp_scorecard_check.readme.identifier
    custom     = true
  }
}
