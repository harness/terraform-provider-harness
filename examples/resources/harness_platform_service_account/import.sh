# Import account level service account
terraform import harness_platform_service_account.example <service_account_id>

# Import org level service account
terraform import harness_platform_service_account.example <ord_id>/<service_account_id>

# Import project level service account
terraform import harness_platform_service_account.example <org_id>/<project_id>/<service_account_id>
