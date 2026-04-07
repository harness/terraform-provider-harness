# Example 1: Lookup by Identity (Project-level)
data "harness_chaos_experiment_template" "by_identity" {
  org_id       = "my_org"
  project_id   = "my_project"
  hub_identity = "my-chaos-hub"
  identity     = "simple-pod-delete-experiment"
}

# Example 2: Lookup by Name (Project-level)
data "harness_chaos_experiment_template" "by_name" {
  org_id       = "my_org"
  project_id   = "my_project"
  hub_identity = "my-chaos-hub"
  name         = "Simple Pod Delete Experiment"
}

# Example 3: Lookup Org-level Template
data "harness_chaos_experiment_template" "org_template" {
  org_id       = "my_org"
  hub_identity = "org-chaos-hub"
  identity     = "org-experiment-template"
}

# Example 4: Lookup Account-level Template
data "harness_chaos_experiment_template" "account_template" {
  hub_identity = "account-chaos-hub"
  identity     = "account-experiment-template"
}

# Outputs
output "experiment_id" {
  description = "The ID of the experiment template"
  value       = data.harness_chaos_experiment_template.by_identity.id
}

output "experiment_identity" {
  description = "The identity of the experiment template"
  value       = data.harness_chaos_experiment_template.by_identity.identity
}

output "experiment_spec" {
  description = "The spec of the experiment template"
  value       = data.harness_chaos_experiment_template.by_identity.spec
}

output "faults_count" {
  description = "Number of faults in the experiment"
  value       = length(data.harness_chaos_experiment_template.by_identity.spec[0].faults)
}

output "probes_count" {
  description = "Number of probes in the experiment"
  value       = length(data.harness_chaos_experiment_template.by_identity.spec[0].probes)
}

output "actions_count" {
  description = "Number of actions in the experiment"
  value       = length(data.harness_chaos_experiment_template.by_identity.spec[0].actions)
}

output "vertices_count" {
  description = "Number of vertices in the workflow"
  value       = length(data.harness_chaos_experiment_template.by_identity.spec[0].vertices)
}
