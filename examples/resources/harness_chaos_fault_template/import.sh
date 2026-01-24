# Example 1: Import Project-level Fault Template
# Format: org_id/project_id/hub_identity/template_identity
terraform import harness_chaos_fault_template.example \
  my_org/my_project/my-chaos-hub/pod-delete-fault

# Example 2: Import Org-level Fault Template
# Format: org_id/hub_identity/template_identity
terraform import harness_chaos_fault_template.org_example \
  my_org/org-chaos-hub/org-pod-delete

# Example 3: Import Account-level Fault Template
# Format: hub_identity/template_identity
terraform import harness_chaos_fault_template.account_example \
  account-chaos-hub/account-network-fault
