# Example 1: Import Project-level Probe Template
# Format: org_id/project_id/hub_identity/template_identity
terraform import harness_chaos_probe_template.example \
  my_org/my_project/my-chaos-hub/http-health-check

# Example 2: Import Org-level Probe Template
# Format: org_id/hub_identity/template_identity
terraform import harness_chaos_probe_template.org_example \
  my_org/org-chaos-hub/org-http-probe

# Example 3: Import Account-level Probe Template
# Format: hub_identity/template_identity
terraform import harness_chaos_probe_template.account_example \
  account-chaos-hub/account-k8s-probe
