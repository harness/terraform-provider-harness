# Data source to retrieve account level WinRM credential by identifier
data "harness_platform_secret_winrm" "example" {
  identifier = "winrm_credential_id"
}

# Data source to retrieve organization level WinRM credential
data "harness_platform_secret_winrm" "org_example" {
  identifier = "winrm_credential_id"
  org_id     = "org_identifier"
}

# Data source to retrieve project level WinRM credential
data "harness_platform_secret_winrm" "project_example" {
  identifier = "winrm_credential_id"
  org_id     = "org_identifier"
  project_id = "project_identifier"
}

# Using the data source output
output "winrm_port" {
  value = data.harness_platform_secret_winrm.example.port
}

output "winrm_auth_type" {
  value = length(data.harness_platform_secret_winrm.example.ntlm) > 0 ? "NTLM" : "Kerberos"
}
