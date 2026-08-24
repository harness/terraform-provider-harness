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
    rule_description       = "Repository has a README"

    input_values {
      key   = "filePath"
      value = "README.md"
    }
  }
}
