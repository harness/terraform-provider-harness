# Import Project level Chaos Image Registry
terraform import harness_chaos_image_registry.example <org_id>/<project_id>/<environment_id>/<infra_id>

# Import Org level Chaos Image Registry
terraform import harness_chaos_image_registry.example <org_id>/<environment_id>/<infra_id>

# Import Account level Chaos Image Registry
terraform import harness_chaos_image_registry.example <environment_id>/<infra_id>
