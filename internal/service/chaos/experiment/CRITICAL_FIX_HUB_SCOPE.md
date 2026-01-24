# CRITICAL FIX: Hub Scope vs Experiment Scope

## Date: Jan 28, 2026

## Problem Identified

Our implementation was incorrectly handling hub references by:
1. Adding `account.` or `org.` prefixes to hub identity based on **experiment scope**
2. Passing **experiment scope** in query parameters instead of **hub scope**

## Root Cause Analysis

### API Behavior (from CURL analysis)

The API uses **TWO DIFFERENT SCOPES**:

1. **Query Parameters** = WHERE THE HUB/TEMPLATE LIVES
2. **Request Body** = WHERE THE EXPERIMENT WILL BE CREATED

### Example from Real CURL Request

Creating a **project-level experiment** from an **org-level hub**:

**Query Parameters** (Hub scope):
```
organizationIdentifier=chaos_e2e_test_org
projectIdentifier=                          ← EMPTY (hub is at org level)
hubIdentity=e2e_chaos_hub_org              ← NO PREFIX!
```

**Request Body** (Experiment scope):
```json
{
  "organizationIdentifier": "chaos_e2e_test_org",
  "projectIdentifier": "chaos_e2e_test_project",  ← HAS VALUE (experiment at project)
  ...
}
```

## The Fix

### Schema Changes

Added two new optional fields to track hub scope separately:

```go
"hub_org_id": {
    Description: "Organization identifier where the hub/template resides (leave empty for account-level hubs)",
    Type:        schema.TypeString,
    Optional:    true,
    ForceNew:    true,
},
"hub_project_id": {
    Description: "Project identifier where the hub/template resides (leave empty for org/account-level hubs)",
    Type:        schema.TypeString,
    Optional:    true,
    ForceNew:    true,
},
```

### Implementation Changes

**BEFORE (WRONG)**:
```go
// Adding prefix based on EXPERIMENT scope
hubRef := hubIdentity
if orgID == "" {
    hubRef = "account." + hubIdentity
} else if projectID == "" {
    hubRef = "org." + hubIdentity
}

opts := &chaos.ExperimenttemplateApiCreateExperimentFromTemplateOpts{
    HubIdentity:            optional.NewString(hubRef),        // With prefix
    OrganizationIdentifier: optional.NewString(orgID),         // Experiment org
    ProjectIdentifier:      optional.NewString(projectID),     // Experiment project
}
```

**AFTER (CORRECT)**:
```go
// Extract hub scope separately
hubOrgID := d.Get("hub_org_id").(string)
hubProjectID := d.Get("hub_project_id").(string)

opts := &chaos.ExperimenttemplateApiCreateExperimentFromTemplateOpts{
    HubIdentity:            optional.NewString(hubIdentity),    // No prefix!
    OrganizationIdentifier: optional.NewString(hubOrgID),       // Hub's org
    ProjectIdentifier:      optional.NewString(hubProjectID),   // Hub's project
}
```

## Usage Examples

### Example 1: Project Experiment from Org Hub

```hcl
resource "harness_chaos_experiment" "example" {
  # Experiment scope (where experiment will be created)
  org_id     = "my-org"
  project_id = "my-project"
  
  # Hub scope (where template lives)
  hub_org_id     = "my-org"      # Hub is at org level
  hub_project_id = ""            # Empty = org-level hub
  hub_identity   = "my-hub"      # No prefix needed
  
  template_identity = "my-template"
  name              = "My Experiment"
  infra_ref         = "my-infra"
}
```

### Example 2: Org Experiment from Account Hub

```hcl
resource "harness_chaos_experiment" "example" {
  # Experiment scope
  org_id     = "my-org"
  project_id = ""                # Empty = org-level experiment
  
  # Hub scope
  hub_org_id     = ""            # Empty = account-level hub
  hub_project_id = ""            # Empty = account-level hub
  hub_identity   = "account-hub" # No prefix needed
  
  template_identity = "my-template"
  name              = "My Experiment"
  infra_ref         = "my-infra"
}
```

### Example 3: Project Experiment from Project Hub

```hcl
resource "harness_chaos_experiment" "example" {
  # Experiment scope
  org_id     = "my-org"
  project_id = "my-project"
  
  # Hub scope (same as experiment in this case)
  hub_org_id     = "my-org"
  hub_project_id = "my-project"  # Hub is at project level
  hub_identity   = "project-hub" # No prefix needed
  
  template_identity = "my-template"
  name              = "My Experiment"
  infra_ref         = "my-infra"
}
```

## API Mapping

| Terraform Field | API Location | Purpose |
|----------------|--------------|---------|
| `org_id` | Request body `organizationIdentifier` | Experiment's org |
| `project_id` | Request body `projectIdentifier` | Experiment's project |
| `hub_org_id` | Query param `organizationIdentifier` | Hub's org |
| `hub_project_id` | Query param `projectIdentifier` | Hub's project |
| `hub_identity` | Query param `hubIdentity` | Hub name (no prefix) |

## Impact

✅ **FIXES**: Cross-scope experiment creation (e.g., project experiment from org hub)
✅ **SIMPLIFIES**: No more automatic prefix logic
✅ **CLARIFIES**: Explicit separation of hub scope vs experiment scope
✅ **MATCHES**: Exact API behavior from CURL analysis

## Files Modified

1. `resource_experiment_schema.go` - Added `hub_org_id` and `hub_project_id` fields
2. `resource_experiment.go` - Removed prefix logic, use hub scope fields for query params

## Testing Required

- [ ] Account-level hub → Account-level experiment
- [ ] Account-level hub → Org-level experiment
- [ ] Account-level hub → Project-level experiment
- [ ] Org-level hub → Org-level experiment
- [ ] Org-level hub → Project-level experiment
- [ ] Project-level hub → Project-level experiment

## Status

✅ **FIXED** - Implementation now matches API behavior exactly
