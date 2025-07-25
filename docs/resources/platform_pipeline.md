---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "harness_platform_pipeline Resource - terraform-provider-harness"
subcategory: "Next Gen"
description: |-
  Resource for creating a Harness pipeline.
---

# harness_platform_pipeline (Resource)

Resource for creating a Harness pipeline.

## Example Usage

```terraform
resource "harness_platform_pipeline" "example" {
  identifier = "identifier"
  org_id     = "orgIdentifier"
  project_id = "projectIdentifier"
  name       = "name"
  git_details {
    branch_name    = "branchName"
    commit_message = "commitMessage"
    file_path      = "filePath"
    connector_ref  = "connectorRef"
    store_type     = "REMOTE"
    repo_name      = "repoName"
  }
  tags = {}
  yaml = <<-EOT
      pipeline:
          name: name
          identifier: identifier
          allowStageExecutions: false
          projectIdentifier: projectIdentifier
          orgIdentifier: orgIdentifier
          tags: {}
          stages:
              - stage:
                  name: dep
                  identifier: dep
                  description: ""
                  type: Deployment
                  spec:
                      serviceConfig:
                          serviceRef: service
                          serviceDefinition:
                              type: Kubernetes
                              spec:
                                  variables: []
                      infrastructure:
                          environmentRef: testenv
                          infrastructureDefinition:
                              type: KubernetesDirect
                              spec:
                                  connectorRef: testconf
                                  namespace: test
                                  releaseName: release-<+INFRA_KEY>
                          allowSimultaneousDeployments: false
                      execution:
                          steps:
                              - stepGroup:
                                      name: Canary Deployment
                                      identifier: canaryDepoyment
                                      steps:
                                          - step:
                                              name: Canary Deployment
                                              identifier: canaryDeployment
                                              type: K8sCanaryDeploy
                                              timeout: 10m
                                              spec:
                                                  instanceSelection:
                                                      type: Count
                                                      spec:
                                                          count: 1
                                                  skipDryRun: false
                                          - step:
                                              name: Canary Delete
                                              identifier: canaryDelete
                                              type: K8sCanaryDelete
                                              timeout: 10m
                                              spec: {}
                                      rollbackSteps:
                                          - step:
                                              name: Canary Delete
                                              identifier: rollbackCanaryDelete
                                              type: K8sCanaryDelete
                                              timeout: 10m
                                              spec: {}
                              - stepGroup:
                                      name: Primary Deployment
                                      identifier: primaryDepoyment
                                      steps:
                                          - step:
                                              name: Rolling Deployment
                                              identifier: rollingDeployment
                                              type: K8sRollingDeploy
                                              timeout: 10m
                                              spec:
                                                  skipDryRun: false
                                      rollbackSteps:
                                          - step:
                                              name: Rolling Rollback
                                              identifier: rollingRollback
                                              type: K8sRollingRollback
                                              timeout: 10m
                                              spec: {}
                          rollbackSteps: []
                  tags: {}
                  failureStrategies:
                      - onFailure:
                              errors:
                                  - AllErrors
                              action:
                                  type: StageRollback
  EOT
}

### Importing Pipeline from Git
resource "harness_platform_organization" "test" {
  identifier = "identifier"
  name       = "name"
}
resource "harness_platform_pipeline" "test" {
  identifier      = "gitx"
  org_id          = "default"
  project_id      = "V"
  name            = "gitx"
  import_from_git = true
  git_import_info {
    branch_name   = "main"
    file_path     = ".harness/gitx.yaml"
    connector_ref = "account.DoNotDeleteGithub"
    repo_name     = "open-repo"
  }
  pipeline_import_request {
    pipeline_name        = "gitx"
    pipeline_description = "Pipeline Description"
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `identifier` (String) Unique identifier of the resource.
- `name` (String) Name of the resource.
- `org_id` (String) Unique identifier of the organization.
- `project_id` (String) Unique identifier of the project.

### Optional

- `description` (String) Description of the resource.
- `git_details` (Block List, Max: 1) Contains parameters related to creating an Entity for Git Experience. (see [below for nested schema](#nestedblock--git_details))
- `git_import_info` (Block List, Max: 1) Contains Git Information for importing entities from Git (see [below for nested schema](#nestedblock--git_import_info))
- `import_from_git` (Boolean) Flag to set if importing from Git
- `pipeline_import_request` (Block List, Max: 1) Contains parameters for importing a pipeline (see [below for nested schema](#nestedblock--pipeline_import_request))
- `tags` (Set of String) Tags to associate with the resource. These should match the tag value passed in the YAML; if this parameter is null or not passed, the tags specified in YAML should also be null.
- `template_applied` (Boolean) If true, returns Pipeline YAML with Templates applied on it.
- `template_applied_pipeline_yaml` (String) Pipeline YAML after resolving Templates (returned as a String).
- `yaml` (String) YAML of the pipeline. In YAML, to reference an entity at the organization scope, prefix 'org' to the expression: org.{identifier}. To reference an entity at the account scope, prefix 'account` to the expression: account.{identifier}. For eg, to reference a connector with identifier 'connectorId' at the organization scope in a stage mention it as connectorRef: org.connectorId.

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--git_details"></a>
### Nested Schema for `git_details`

Optional:

- `base_branch` (String) Name of the default branch (this checks out a new branch titled by branch_name).
- `branch_name` (String) Name of the branch.
- `commit_message` (String) Commit message used for the merge commit.
- `connector_ref` (String) Identifier of the Harness Connector used for CRUD operations on the Entity. To reference a connector at the organization scope, prefix 'org' to the expression: org.{identifier}. To reference a connector at the account scope, prefix 'account` to the expression: account.{identifier}.
- `file_path` (String) File path of the Entity in the repository.
- `is_harness_code_repo` (Boolean) If the repo is harness code.
- `last_commit_id` (String) Last commit identifier (for Git Repositories other than Github). To be provided only when updating Pipeline.
- `last_object_id` (String) Last object identifier (for Github). To be provided only when updating Pipeline.
- `repo_name` (String) Name of the repository.
- `store_type` (String) Specifies whether the Entity is to be stored in Git or not. Possible values: INLINE, REMOTE.


<a id="nestedblock--git_import_info"></a>
### Nested Schema for `git_import_info`

Optional:

- `branch_name` (String) Name of the branch.
- `connector_ref` (String) Identifier of the Harness Connector used for importing entity from Git To reference a connector at the organization scope, prefix 'org' to the expression: org.{identifier}. To reference a connector at the account scope, prefix 'account` to the expression: account.{identifier}.
- `file_path` (String) File path of the Entity in the repository.
- `repo_name` (String) Name of the repository.


<a id="nestedblock--pipeline_import_request"></a>
### Nested Schema for `pipeline_import_request`

Optional:

- `pipeline_description` (String) Description of the pipeline.
- `pipeline_name` (String) Name of the pipeline.

## Import

Import is supported using the following syntax:

The [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import) can be used, for example:

```shell
# Import pipeline from default branch
terraform import harness_platform_pipeline.example <org_id>/<project_id>/<pipeline_id>

# Import pipeline from non default branch
terraform import harness_platform_pipeline.example <org_id>/<project_id>/<pipeline_id>/<branch>
```
