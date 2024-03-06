# Import account level secret sshkey
terraform import harness_platform_secret_sshkey.example <secret_sshkey_id>

# Import org level secret sshkey
terraform import harness_platform_secret_sshkey.example <ord_id>/<secret_sshkey_id>

# Import project level secret sshkey
terraform import harness_platform_secret_sshkey.example <org_id>/<project_id>/<secret_sshkey_id>
