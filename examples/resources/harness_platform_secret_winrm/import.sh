# Import account level WinRM credential
terraform import harness_platform_secret_winrm.example <winrm_credential_id>

# Import organization level WinRM credential
terraform import harness_platform_secret_winrm.example <org_id>/<winrm_credential_id>

# Import project level WinRM credential
terraform import harness_platform_secret_winrm.example <org_id>/<project_id>/<winrm_credential_id>
