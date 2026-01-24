# Example 1: Import Project-level Experiment Template
# Format: org_id/project_id/hub_identity/template_identity
terraform import harness_chaos_experiment_template.simple_fault \
  my_org/my_project/my-chaos-hub/simple-pod-delete-experiment

# Example 2: Import Org-level Experiment Template
# Format: org_id/hub_identity/template_identity
terraform import harness_chaos_experiment_template.org_level \
  my_org/org-chaos-hub/org-experiment-template

# Example 3: Import Account-level Experiment Template
# Format: hub_identity/template_identity
terraform import harness_chaos_experiment_template.account_level \
  account-chaos-hub/account-experiment-template

# Example 4: Import Comprehensive Experiment Template
terraform import harness_chaos_experiment_template.comprehensive \
  my_org/my_project/my-chaos-hub/comprehensive-experiment

# Example 5: Import Multi-Fault Experiment Template
terraform import harness_chaos_experiment_template.multi_fault \
  my_org/my_project/my-chaos-hub/multi-fault-experiment
