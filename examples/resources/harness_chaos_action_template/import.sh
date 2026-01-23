# Example 1: Import Project-level Action Template
# Format: org_id/project_id/hub_identity/template_identity
terraform import harness_chaos_action_template.example \
  my_org/my_project/my-chaos-hub/delay-action-template

# Example 2: Import Org-level Action Template  
# Format: org_id/hub_identity/template_identity
terraform import harness_chaos_action_template.org_example \
  my_org/org-chaos-hub/org-script-action

# Example 3: Import Account-level Action Template
# Format: hub_identity/template_identity
terraform import harness_chaos_action_template.account_example \
  account-chaos-hub/account-delay-action
