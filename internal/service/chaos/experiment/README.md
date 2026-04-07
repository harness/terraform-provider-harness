# Chaos Experiment Terraform Resource

## Overview

This resource allows you to create and manage Harness Chaos Experiments from existing Experiment Templates. Experiments are instances launched from templates and bound to specific infrastructure.

## Resource: `harness_chaos_experiment`

### Key Characteristics

- **Created FROM Templates**: Experiments are launched from existing experiment templates
- **Infrastructure Binding**: MUST specify infrastructure at creation (cannot exist without it)
- **Immutable Template Reference**: Cannot change template or hub after creation (ForceNew)
- **Limited Updates**: Most fields are immutable; changes trigger recreation

### Example Usage

#### Basic Experiment

```hcl
resource "harness_chaos_experiment" "example" {
  org_id            = "default"
  project_id        = "chaos_project"
  template_identity = "my-experiment-template"
  hub_identity      = "enterprise-hub"
  name              = "production-chaos-test"
  infra_ref         = "prod-k8s-cluster"
  
  description = "Production chaos experiment"
  tags        = ["production", "critical"]
}
```

#### With Custom Identity

```hcl
resource "harness_chaos_experiment" "custom" {
  org_id            = "default"
  project_id        = "chaos_project"
  template_identity = "network-latency-template"
  hub_identity      = "enterprise-hub"
  name              = "network-latency-test"
  infra_ref         = harness_chaos_infrastructure_v2.k8s.id
  
  identity    = "custom-experiment-id"
  description = "Network latency chaos experiment"
  revision    = "v2"
  
  tags = ["network", "latency", "performance"]
}
```

#### Complete Integration Example

```hcl
# 1. Create Hub
resource "harness_chaos_hub_v2" "enterprise" {
  org_id       = "default"
  project_id   = "chaos_project"
  identity     = "enterprise-hub"
  name         = "Enterprise Hub"
  connector_id = "github-connector"
  repo_branch  = "main"
  repo_name    = "chaos-hub"
}

# 2. Create Experiment Template
resource "harness_chaos_experiment_template" "network_test" {
  org_id       = "default"
  project_id   = "chaos_project"
  hub_identity = harness_chaos_hub_v2.enterprise.identity
  identity     = "network-latency-template"
  name         = "Network Latency Template"
  
  # ... template configuration
}

# 3. Create Infrastructure
resource "harness_chaos_infrastructure_v2" "k8s" {
  org_id         = "default"
  project_id     = "chaos_project"
  environment_id = "prod"
  name           = "Production K8s"
  infra_type     = "KUBERNETESV2"
  # ... infrastructure configuration
}

# 4. Create Experiment from Template
resource "harness_chaos_experiment" "prod_test" {
  org_id            = "default"
  project_id        = "chaos_project"
  template_identity = harness_chaos_experiment_template.network_test.identity
  hub_identity      = harness_chaos_hub_v2.enterprise.identity
  name              = "Production Network Test"
  infra_ref         = harness_chaos_infrastructure_v2.k8s.id
  
  description = "Production network latency experiment"
  tags        = ["production", "network"]
}
```

## Data Source: `harness_chaos_experiment`

### Lookup by Identity

```hcl
data "harness_chaos_experiment" "existing" {
  org_id     = "default"
  project_id = "chaos_project"
  identity   = "my-experiment"
}

output "experiment_id" {
  value = data.harness_chaos_experiment.existing.experiment_id
}
```

### Lookup by Name

```hcl
data "harness_chaos_experiment" "by_name" {
  org_id     = "default"
  project_id = "chaos_project"
  name       = "Production Chaos Test"
}

output "template_used" {
  value = data.harness_chaos_experiment.by_name.template_identity
}
```

## Schema Reference

### Required Arguments

| Argument | Type | Description |
|----------|------|-------------|
| `org_id` | string | Organization identifier (ForceNew) |
| `project_id` | string | Project identifier (ForceNew) |
| `template_identity` | string | Experiment template to launch from (ForceNew) |
| `hub_identity` | string | Hub where template resides (ForceNew) |
| `name` | string | Experiment name |
| `infra_ref` | string | Infrastructure reference/ID (ForceNew) |

### Optional Arguments

| Argument | Type | Default | Description |
|----------|------|---------|-------------|
| `identity` | string | auto-generated | Custom experiment identifier (ForceNew) |
| `description` | string | - | Experiment description |
| `tags` | set(string) | - | Tags for categorization |
| `revision` | string | "v1" | Template revision to use (ForceNew) |

### Computed Attributes

| Attribute | Type | Description |
|-----------|------|-------------|
| `experiment_id` | string | Full experiment ID |
| `infra_id` | string | Resolved infrastructure ID |
| `infra_type` | string | Infrastructure type (e.g., KubernetesV2) |
| `experiment_type` | string | Type of experiment |
| `is_custom_experiment` | bool | Whether this is a custom experiment |
| `fault_ids` | list(string) | List of fault IDs used |
| `cron_syntax` | string | Cron expression for scheduling |
| `is_cron_enabled` | bool | Whether cron scheduling is enabled |
| `is_single_run_cron_enabled` | bool | Whether single-run cron is enabled |
| `last_executed_at` | int | Last execution timestamp (Unix) |
| `total_experiment_runs` | int | Total number of runs |
| `target_network_map_id` | string | Target network map ID |
| `created_at` | int | Creation timestamp (Unix) |
| `created_by` | string | Creator username |
| `updated_at` | int | Last update timestamp (Unix) |
| `updated_by` | string | Last updater username |
| `template_details` | object | Template details (see below) |

### Template Details Block

| Attribute | Type | Description |
|-----------|------|-------------|
| `identity` | string | Template identity |
| `hub_reference` | string | Hub reference |
| `reference` | string | Full template reference |
| `revision` | string | Template revision used |

## Import

Experiments can be imported using the format: `org_id/project_id/experiment_id`

**Note**: Use the `experiment_id` (the actual ID returned by the API), not the `identity` field.

```bash
terraform import harness_chaos_experiment.example default/chaos_project/my-experiment
```

## API Endpoints Used

- **Create**: `POST /rest/experimenttemplates/{identity}/launch` (CreateExperimentFromTemplate)
- **Read**: `GET /rest/v2/experiments/{experimentId}` (GetChaosV2Experiment)
- **Delete**: `DELETE /rest/v2/experiment/{experimentId}` (DeleteChaosV2Experiment)
- **List**: `GET /rest/v2/experiments` (ListChaosV2Experiment)

## Important Notes

### ForceNew Behavior

The following fields trigger resource recreation when changed:
- `org_id`, `project_id`
- `template_identity`, `hub_identity`
- `infra_ref`
- `identity`, `revision`

### Update Limitations

Experiments created from templates have **limited update support**. Most fields are immutable. Currently, updates are not supported - changes will trigger recreation.

### Infrastructure Dependency

Experiments **MUST** have an infrastructure binding. They cannot exist without infrastructure. Ensure your infrastructure resource is created before the experiment.

### Template Dependency Chain

```
Hub → Templates (Action/Probe/Fault) → Experiment Template → Infrastructure → Experiment
```

## Testing

See test files for comprehensive examples:
- `resource_experiment_test.go` - Resource tests
- `data_source_experiment_test.go` - Data source tests

## SDK Models

- Request: `TypesCreateExperimentFromTemplateRequest`
- Response: `TypesExperimentCreationResponse`
- Read: `ChaosExperimentChaosExperimentRequest`
- List: `TypesListExperimentV2Response`

## Related Resources

- `harness_chaos_experiment_template` - Create experiment templates
- `harness_chaos_infrastructure_v2` - Create infrastructure
- `harness_chaos_hub_v2` - Manage chaos hubs
- `harness_chaos_action_template` - Create action templates
- `harness_chaos_probe_template` - Create probe templates
- `harness_chaos_fault_template` - Create fault templates
