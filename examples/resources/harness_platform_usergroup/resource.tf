terraform {
  required_providers {
    harness = {
      source = "harness/harness"
    }
  }
}

resource "harness_platform_usergroup" "abhiTerraformUGV1" {
  identifier         = "abhiTerraformUGV1"
  name               = "abhiTerraformUGV1"
  # org_id             = "abhishek_test"
  # project_id         = "abhishek_test"
  externally_managed = false
  user_emails        = ["abhishek.das+3@harness.io"]
}

resource "harness_platform_usergroup" "abhiTerraformUGV3" {
  identifier         = "abhiTerraformUGV3"
  name               = "abhiTerraformUGV3"
  # org_id             = "abhishek_test"
  # project_id         = "abhishek_test"
  externally_managed = false
  user_emails        = ["abhishek.das+3@harness.io"]
}