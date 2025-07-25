---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "harness_platform_infrastructure Resource - terraform-provider-harness"
subcategory: "Next Gen"
description: |-
  Resource for creating a Harness Infrastructure.
---

# harness_platform_infrastructure (Resource)

Resource for creating a Harness Infrastructure.

## Example Usage

```terraform
resource "harness_platform_infrastructure" "example" {
  identifier      = "identifier"
  name            = "name"
  org_id          = "orgIdentifer"
  project_id      = "projectIdentifier"
  env_id          = "environmentIdentifier"
  type            = "KubernetesDirect"
  deployment_type = "Kubernetes"
  git_details {
    branch_name    = "branchName"
    commit_message = "commitMessage"
    file_path      = "filePath"
    connector_ref  = "connectorRef"
    store_type     = "REMOTE"
    repo_name      = "repoName"
  }
  yaml = <<-EOT
        infrastructureDefinition:
         name: name
         identifier: identifier
         description: ""
         tags:
           asda: ""
         orgIdentifier: orgIdentifer
         projectIdentifier: projectIdentifier
         environmentRef: environmentIdentifier
         deploymentType: Kubernetes
         type: KubernetesDirect
         spec:
          connectorRef: account.gfgf
          namespace: asdasdsa
          releaseName: release-<+INFRA_KEY>
          allowSimultaneousDeployments: false
      EOT
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `env_id` (String) Environment Identifier.
- `identifier` (String) Unique identifier of the resource.
- `name` (String) Name of the resource.

### Optional

- `deployment_type` (String) Infrastructure deployment type. Valid values are Kubernetes, NativeHelm, Ssh, WinRm, ServerlessAwsLambda, AzureWebApp, Custom, ECS.
- `description` (String) Description of the resource.
- `force_delete` (Boolean) Enable this flag for force deletion of infrastructure
- `git_details` (Block List, Max: 1) Contains parameters related to creating an Entity for Git Experience. (see [below for nested schema](#nestedblock--git_details))
- `org_id` (String) Unique identifier of the organization.
- `project_id` (String) Unique identifier of the project.
- `tags` (Set of String) Tags to associate with the resource.
- `type` (String) Type of Infrastructure. Valid values are KubernetesDirect, KubernetesGcp, ServerlessAwsLambda, Pdc, KubernetesAzure, SshWinRmAzure, SshWinRmAws, AzureWebApp, ECS, GitOps, CustomDeployment, TAS, KubernetesRancher, AWS_SAM.
- `yaml` (String) Infrastructure YAML. In YAML, to reference an entity at the organization scope, prefix 'org' to the expression: org.{identifier}. To reference an entity at the account scope, prefix 'account` to the expression: account.{identifier}. For eg, to reference a connector with identifier 'connectorId' at the organization scope in a stage mention it as connectorRef: org.connectorId.

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--git_details"></a>
### Nested Schema for `git_details`

Optional:

- `base_branch` (String) Name of the default branch (this checks out a new branch titled by branch_name).
- `branch` (String) Name of the branch.
- `commit_message` (String) Commit message used for the merge commit.
- `connector_ref` (String) Identifier of the Harness Connector used for CRUD operations on the Entity. To reference a connector at the organization scope, prefix 'org' to the expression: org.{identifier}. To reference a connector at the account scope, prefix 'account` to the expression: account.{identifier}.
- `file_path` (String) File path of the Entity in the repository.
- `import_from_git` (Boolean) import infrastructure from git
- `is_force_import` (Boolean) force import infrastructure from remote even if same file path already exist
- `is_harnesscode_repo` (Boolean) If the gitProvider is HarnessCode
- `is_new_branch` (Boolean) If a new branch creation is requested.
- `last_commit_id` (String) Last commit identifier (for Git Repositories other than Github). To be provided only when updating infrastructure.
- `last_object_id` (String) Last object identifier (for Github). To be provided only when updating infrastructure.
- `load_from_cache` (String) If the Entity is to be fetched from cache
- `load_from_fallback_branch` (Boolean) If the Entity is to be fetched from fallbackBranch
- `parent_entity_connector_ref` (String) Identifier of the Harness Connector used for CRUD operations on the Parent Entity. To reference a connector at the organization scope, prefix 'org' to the expression: org.{identifier}. To reference a connector at the account scope, prefix 'account` to the expression: account.{identifier}.
- `parent_entity_repo_name` (String) Name of the repository where parent entity lies.
- `repo_name` (String) Name of the repository.
- `store_type` (String) Specifies whether the Entity is to be stored in Git or not. Possible values: INLINE, REMOTE.

## Import

Import is supported using the following syntax:

The [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import) can be used, for example:

```shell
# Import account level infrastructure
terraform import harness_platform_infrastructure.example <env_id>/<infrastructure_id>

# Import org level infrastructure
terraform import harness_platform_infrastructure.example <org_id>/<env_id>/<infrastructure_id>

# Import project level infrastructure
terraform import harness_platform_infrastructure.example <org_id>/<project_id>/<env_id>/<infrastructure_id>
```
