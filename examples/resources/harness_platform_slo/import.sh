# Import account level SLO
terraform import harness_platform_slo.example <slo_id>

# Import organization level SLO
terraform import harness_platform_slo.example <org_id>/<slo_id>

# Import project level SLO
terraform import harness_platform_slo.example <org_id>/<project_id>/<slo_id>
