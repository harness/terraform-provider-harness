# Failing ServiceNow and GitHub integrations of an organization.
data "harness_platform_idp_integrations" "example" {
  org_id = "default"
  type   = ["ServiceNow", "GitHub"]
  health = "failing"
}

output "failing" {
  value = data.harness_platform_idp_integrations.example.integrations[*].identifier
}
