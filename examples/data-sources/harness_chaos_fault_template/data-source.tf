# Example 1: Lookup Fault Template by Identity (Recommended)
data "harness_chaos_fault_template" "by_identity" {
  org_id       = "my_org"
  project_id   = "my_project"
  hub_identity = "my-chaos-hub"
  identity     = "pod-delete-fault"
}

# Use the fault template data
output "fault_name" {
  value = data.harness_chaos_fault_template.by_identity.name
}

output "fault_category" {
  value = data.harness_chaos_fault_template.by_identity.category
}

output "fault_description" {
  value = data.harness_chaos_fault_template.by_identity.description
}

# Example 2: Lookup Fault Template by Name
data "harness_chaos_fault_template" "by_name" {
  org_id       = "my_org"
  project_id   = "my_project"
  hub_identity = "my-chaos-hub"
  name         = "Pod Delete Fault"
}

# Example 3: Use in Chaos Experiment
resource "harness_chaos_experiment" "example" {
  # ... other configuration ...

  # Reference fault template
  fault {
    name        = data.harness_chaos_fault_template.by_identity.name
    description = data.harness_chaos_fault_template.by_identity.description
    
    # Use fault template's chaos spec
    # ... fault configuration ...
  }
}

# Example 4: Access nested fault template data
output "kubernetes_image" {
  value = try(
    data.harness_chaos_fault_template.by_identity.spec[0].chaos[0].kubernetes[0].image,
    "not configured"
  )
}

output "fault_params" {
  value = try(
    data.harness_chaos_fault_template.by_identity.spec[0].chaos[0].params,
    []
  )
}

# Example 5: Conditional logic based on fault template
locals {
  has_kubernetes_spec = length(try(
    data.harness_chaos_fault_template.by_identity.spec[0].chaos[0].kubernetes,
    []
  )) > 0
  
  has_targets = length(try(
    data.harness_chaos_fault_template.by_identity.spec[0].target,
    []
  )) > 0
}

output "template_features" {
  value = {
    has_kubernetes_spec = local.has_kubernetes_spec
    has_targets         = local.has_targets
    tags                = data.harness_chaos_fault_template.by_identity.tags
  }
}
