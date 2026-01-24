# Example 1: Lookup Probe Template by Identity (Recommended)
data "harness_chaos_probe_template" "by_identity" {
  org_id       = "my_org"
  project_id   = "my_project"
  hub_identity = "my-chaos-hub"
  identity     = "http-health-check"
}

# Use the probe template data
output "probe_name" {
  value = data.harness_chaos_probe_template.by_identity.name
}

output "probe_type" {
  value = data.harness_chaos_probe_template.by_identity.type
}

# Example 2: Lookup Probe Template by Name
data "harness_chaos_probe_template" "by_name" {
  org_id       = "my_org"
  project_id   = "my_project"
  hub_identity = "my-chaos-hub"
  name         = "HTTP Health Check Probe"
}

# Example 3: Use in another resource
resource "harness_chaos_experiment" "example" {
  # ... other configuration ...

  # Reference probe template
  probe {
    name = data.harness_chaos_probe_template.by_identity.name
    type = data.harness_chaos_probe_template.by_identity.type
    # ... probe configuration ...
  }
}
