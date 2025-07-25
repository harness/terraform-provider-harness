---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "harness_platform_template Resource - terraform-provider-harness"
subcategory: "Next Gen"
description: |-
  Resource for creating a Template. Description field is deprecated
---

# harness_platform_template (Resource)

Resource for creating a Template. Description field is deprecated

## Example Usage

```terraform
## Remote Pipeline template
resource "harness_platform_template" "pipeline_template_remote" {
  identifier = "identifier"
  org_id     = harness_platform_project.test.org_id
  project_id = harness_platform_project.test.id
  name       = "name"
  comments   = "comments"
  version    = "ab"
  is_stable  = true
  git_details {
    branch_name    = "main"
    commit_message = "Commit"
    file_path      = "file_path"
    connector_ref  = "account.connector_ref"
    store_type     = "REMOTE"
    repo_name      = "repo_name"
  }
  template_yaml = <<-EOT
template:
  name: "name"
  identifier: "identifier"
  versionLabel: "ab"
  type: Pipeline
  projectIdentifier: ${harness_platform_project.test.id}
  orgIdentifier: ${harness_platform_project.test.org_id}
  tags: {}
  spec:
    stages:
      - stage:
          identifier: dvvdvd
          name: dvvdvd
          description: ""
          type: Deployment
          spec:
            deploymentType: Kubernetes
            service:
              serviceRef: <+input>
              serviceInputs: <+input>
            environment:
              environmentRef: <+input>
              deployToAll: false
              environmentInputs: <+input>
              serviceOverrideInputs: <+input>
              infrastructureDefinitions: <+input>
            execution:
              steps:
                - step:
                    name: Rollout Deployment
                    identifier: rolloutDeployment
                    type: K8sRollingDeploy
                    timeout: 10m
                    spec:
                      skipDryRun: false
                      pruningEnabled: false
              rollbackSteps:
                - step:
                    name: Rollback Rollout Deployment
                    identifier: rollbackRolloutDeployment
                    type: K8sRollingRollback
                    timeout: 10m
                    spec:
                      pruningEnabled: false
          tags: {}
          failureStrategies:
            - onFailure:
                errors:
                  - AllErrors
                action:
                  type: StageRollback

  EOT
}

## Remote Pipeline template to create new branch from existing base branch
resource "harness_platform_template" "pipeline_template_remote" {
  identifier = "identifier"
  org_id     = harness_platform_project.test.org_id
  project_id = harness_platform_project.test.id
  name       = "name"
  comments   = "comments"
  version    = "ab"
  is_stable  = true
  git_details {
    branch_name    = "new_branch"
    commit_message = "Commit"
    file_path      = "file_path"
    connector_ref  = "account.connector_ref"
    store_type     = "REMOTE"
    repo_name      = "repo_name"
    base_branch    = "main"
  }
  template_yaml = <<-EOT
template:
  name: "name"
  identifier: "identifier"
  versionLabel: "ab"
  type: Pipeline
  projectIdentifier: ${harness_platform_project.test.id}
  orgIdentifier: ${harness_platform_project.test.org_id}
  tags: {}
  spec:
    stages:
      - stage:
          identifier: dvvdvd
          name: dvvdvd
          description: ""
          type: Deployment
          spec:
            deploymentType: Kubernetes
            service:
              serviceRef: <+input>
              serviceInputs: <+input>
            environment:
              environmentRef: <+input>
              deployToAll: false
              environmentInputs: <+input>
              serviceOverrideInputs: <+input>
              infrastructureDefinitions: <+input>
            execution:
              steps:
                - step:
                    name: Rollout Deployment
                    identifier: rolloutDeployment
                    type: K8sRollingDeploy
                    timeout: 10m
                    spec:
                      skipDryRun: false
                      pruningEnabled: false
              rollbackSteps:
                - step:
                    name: Rollback Rollout Deployment
                    identifier: rollbackRolloutDeployment
                    type: K8sRollingRollback
                    timeout: 10m
                    spec:
                      pruningEnabled: false
          tags: {}
          failureStrategies:
            - onFailure:
                errors:
                  - AllErrors
                action:
                  type: StageRollback

  EOT
}

## Inline Pipeline template
resource "harness_platform_template" "pipeline_template_inline" {
  identifier    = "identifier"
  org_id        = harness_platform_project.test.org_id
  project_id    = harness_platform_project.test.id
  name          = "name"
  comments      = "comments"
  version       = "ab"
  is_stable     = true
  template_yaml = <<-EOT
template:
  name: "name"
  identifier: "identifier"
  versionLabel: "ab"
  type: Pipeline
  projectIdentifier: ${harness_platform_project.test.id}
  orgIdentifier: ${harness_platform_project.test.org_id}
  tags: {}
  spec:
    stages:
      - stage:
          identifier: dvvdvd
          name: dvvdvd
          description: ""
          type: Deployment
          spec:
            deploymentType: Kubernetes
            service:
              serviceRef: <+input>
              serviceInputs: <+input>
            environment:
              environmentRef: <+input>
              deployToAll: false
              environmentInputs: <+input>
              serviceOverrideInputs: <+input>
              infrastructureDefinitions: <+input>
            execution:
              steps:
                - step:
                    name: Rollout Deployment
                    identifier: rolloutDeployment
                    type: K8sRollingDeploy
                    timeout: 10m
                    spec:
                      skipDryRun: false
                      pruningEnabled: false
              rollbackSteps:
                - step:
                    name: Rollback Rollout Deployment
                    identifier: rollbackRolloutDeployment
                    type: K8sRollingRollback
                    timeout: 10m
                    spec:
                      pruningEnabled: false
          tags: {}
          failureStrategies:
            - onFailure:
                errors:
                  - AllErrors
                action:
                  type: StageRollback
    
  EOT
}

## Inline Step template
resource "harness_platform_template" "step_template_inline" {
  identifier    = "identifier"
  org_id        = harness_platform_project.test.org_id
  project_id    = harness_platform_project.test.id
  name          = "name"
  comments      = "comments"
  version       = "ab"
  is_stable     = true
  template_yaml = <<-EOT
template:
  name: "name"
  identifier: "identifier"
  versionLabel: "ab"
  type: Step
  projectIdentifier: ${harness_platform_project.test.id}
  orgIdentifier: ${harness_platform_project.test.org_id}
  tags: {}
  spec:
    timeout: 10m
    type: ShellScript
    spec:
      shell: Bash
      onDelegate: true
      source:
        type: Inline
        spec:
          script: <+input>
      environmentVariables: []
      outputVariables: []

  EOT
}

## Remote Step template
resource "harness_platform_template" "step_template_remote" {
  identifier = "identifier"
  org_id     = harness_platform_project.test.org_id
  project_id = harness_platform_project.test.id
  name       = "name"
  comments   = "comments"
  version    = "ab"
  is_stable  = true
  git_details {
    branch_name    = "main"
    commit_message = "Commit"
    file_path      = "file_path"
    connector_ref  = "account.connector_ref"
    store_type     = "REMOTE"
    repo_name      = "repo_name"
  }
  template_yaml = <<-EOT
template:
  name: "name"
  identifier: "identifier"
  versionLabel: "ab"
  type: Step
  projectIdentifier: ${harness_platform_project.test.id}
  orgIdentifier: ${harness_platform_project.test.org_id}
  tags: {}
  spec:
    timeout: 10m
    type: ShellScript
    spec:
      shell: Bash
      onDelegate: true
      source:
        type: Inline
        spec:
          script: <+input>
      environmentVariables: []
      outputVariables: []

  EOT
}

## Remote Step template to create new branch from existing branch
resource "harness_platform_template" "step_template_remote" {
  identifier = "identifier"
  org_id     = harness_platform_project.test.org_id
  project_id = harness_platform_project.test.id
  name       = "name"
  comments   = "comments"
  version    = "ab"
  is_stable  = true
  git_details {
    branch_name    = "new_branch"
    commit_message = "Commit"
    file_path      = "file_path"
    connector_ref  = "account.connector_ref"
    store_type     = "REMOTE"
    repo_name      = "repo_name"
    base_branch    = "main"
  }
  template_yaml = <<-EOT
template:
  name: "name"
  identifier: "identifier"
  versionLabel: "ab"
  type: Step
  projectIdentifier: ${harness_platform_project.test.id}
  orgIdentifier: ${harness_platform_project.test.org_id}
  tags: {}
  spec:
    timeout: 10m
    type: ShellScript
    spec:
      shell: Bash
      onDelegate: true
      source:
        type: Inline
        spec:
          script: <+input>
      environmentVariables: []
      outputVariables: []

  EOT
}

## Inline Stage template
resource "harness_platform_template" "stage_template_inline" {
  identifier    = "identifier"
  org_id        = harness_platform_project.test.org_id
  project_id    = harness_platform_project.test.id
  name          = "name"
  comments      = "comments"
  version       = "ab"
  is_stable     = true
  template_yaml = <<-EOT
template:
  name: "name"
  identifier: "identifier"
  versionLabel: "ab"
  type: Stage
  projectIdentifier: ${harness_platform_project.test.id}
  orgIdentifier: ${harness_platform_project.test.org_id}
  tags: {}
  spec:
    type: Deployment
    spec:
      deploymentType: Kubernetes
      service:
        serviceRef: <+input>
        serviceInputs: <+input>
      environment:
        environmentRef: <+input>
        deployToAll: false
        environmentInputs: <+input>
        infrastructureDefinitions: <+input>
      execution:
        steps:
          - step:
              type: ShellScript
              name: Shell Script_1
              identifier: ShellScript_1
              spec:
                shell: Bash
                onDelegate: true
                source:
                  type: Inline
                  spec:
                    script: <+input>
                environmentVariables: []
                outputVariables: []
              timeout: <+input>
        rollbackSteps: []
    failureStrategies:
      - onFailure:
          errors:
            - AllErrors
          action:
            type: StageRollback

  EOT
}

## Remote Stage template
resource "harness_platform_template" "stage_template_remote" {
  identifier = "identifier"
  org_id     = harness_platform_project.test.org_id
  project_id = harness_platform_project.test.id
  name       = "name"
  comments   = "comments"
  version    = "ab"
  is_stable  = true
  git_details {
    branch_name    = "main"
    commit_message = "Commit"
    file_path      = "file_path"
    connector_ref  = "account.connector_ref"
    store_type     = "REMOTE"
    repo_name      = "repo_name"
  }
  template_yaml = <<-EOT
template:
  name: "name"
  identifier: "identifier"
  versionLabel: "ab"
  type: Stage
  projectIdentifier: ${harness_platform_project.test.id}
  orgIdentifier: ${harness_platform_project.test.org_id}
  tags: {}
  spec:
    type: Deployment
    spec:
      deploymentType: Kubernetes
      service:
        serviceRef: <+input>
        serviceInputs: <+input>
      environment:
        environmentRef: <+input>
        deployToAll: false
        environmentInputs: <+input>
        infrastructureDefinitions: <+input>
      execution:
        steps:
          - step:
              type: ShellScript
              name: Shell Script_1
              identifier: ShellScript_1
              spec:
                shell: Bash
                onDelegate: true
                source:
                  type: Inline
                  spec:
                    script: <+input>
                environmentVariables: []
                outputVariables: []
              timeout: <+input>
        rollbackSteps: []
    failureStrategies:
      - onFailure:
          errors:
            - AllErrors
          action:
            type: StageRollback

  EOT
}

## Remote Stage template to create new branch from existing branch
resource "harness_platform_template" "stage_template_remote" {
  identifier = "identifier"
  org_id     = harness_platform_project.test.org_id
  project_id = harness_platform_project.test.id
  name       = "name"
  comments   = "comments"
  version    = "ab"
  is_stable  = true
  git_details {
    branch_name    = "new_branch"
    commit_message = "Commit"
    file_path      = "file_path"
    connector_ref  = "account.connector_ref"
    store_type     = "REMOTE"
    repo_name      = "repo_name"
    base_branch    = "main"
  }
  template_yaml = <<-EOT
template:
  name: "name"
  identifier: "identifier"
  versionLabel: "ab"
  type: Stage
  projectIdentifier: ${harness_platform_project.test.id}
  orgIdentifier: ${harness_platform_project.test.org_id}
  tags: {}
  spec:
    type: Deployment
    spec:
      deploymentType: Kubernetes
      service:
        serviceRef: <+input>
        serviceInputs: <+input>
      environment:
        environmentRef: <+input>
        deployToAll: false
        environmentInputs: <+input>
        infrastructureDefinitions: <+input>
      execution:
        steps:
          - step:
              type: ShellScript
              name: Shell Script_1
              identifier: ShellScript_1
              spec:
                shell: Bash
                onDelegate: true
                source:
                  type: Inline
                  spec:
                    script: <+input>
                environmentVariables: []
                outputVariables: []
              timeout: <+input>
        rollbackSteps: []
    failureStrategies:
      - onFailure:
          errors:
            - AllErrors
          action:
            type: StageRollback

  EOT
}

## Inline StepGroup template
resource "harness_platform_template" "stepgroup_template_inline" {
  identifier    = "identifier"
  org_id        = harness_platform_project.test.org_id
  project_id    = harness_platform_project.test.id
  name          = "name"
  comments      = "comments"
  version       = "ab"
  is_stable     = true
  template_yaml = <<-EOT
template:
  name: "name"
  identifier: "identifier"
  versionLabel: "ab"
  type: StepGroup
  projectIdentifier: ${harness_platform_project.test.id}
  orgIdentifier: ${harness_platform_project.test.org_id}
  tags: {}
  spec:
    stageType: Deployment
    steps:
      - step:
          type: ShellScript
          name: Shell Script_1
          identifier: ShellScript_1
          spec:
            shell: Bash
            onDelegate: true
            source:
              type: Inline
              spec:
                script: <+input>
            environmentVariables: []
            outputVariables: []
          timeout: 10m

  EOT
}

## Remote StepGroup template
resource "harness_platform_template" "stepgroup_template_remote" {
  identifier = "identifier"
  org_id     = harness_platform_project.test.org_id
  project_id = harness_platform_project.test.id
  name       = "name"
  comments   = "comments"
  version    = "ab"
  is_stable  = true
  git_details {
    branch_name    = "main"
    commit_message = "Commit"
    file_path      = "file_path"
    connector_ref  = "account.connector_ref"
    store_type     = "REMOTE"
    repo_name      = "repo_name"
  }
  template_yaml = <<-EOT
template:
  name: "name"
  identifier: "identifier"
  versionLabel: "ab"
  type: StepGroup
  projectIdentifier: ${harness_platform_project.test.id}
  orgIdentifier: ${harness_platform_project.test.org_id}
  tags: {}
  spec:
    stageType: Deployment
    steps:
      - step:
          type: ShellScript
          name: Shell Script_1
          identifier: ShellScript_1
          spec:
            shell: Bash
            onDelegate: true
            source:
              type: Inline
              spec:
                script: <+input>
            environmentVariables: []
            outputVariables: []
          timeout: 10m

  EOT
}

## Remote StepGroup template to create new branch from existing branch
resource "harness_platform_template" "stepgroup_template_remote" {
  identifier = "identifier"
  org_id     = harness_platform_project.test.org_id
  project_id = harness_platform_project.test.id
  name       = "name"
  comments   = "comments"
  version    = "ab"
  is_stable  = true
  git_details {
    branch_name    = "new_branch"
    commit_message = "Commit"
    file_path      = "file_path"
    connector_ref  = "account.connector_ref"
    store_type     = "REMOTE"
    repo_name      = "repo_name"
    base_branch    = "main"
  }
  template_yaml = <<-EOT
template:
  name: "name"
  identifier: "identifier"
  versionLabel: "ab"
  type: StepGroup
  projectIdentifier: ${harness_platform_project.test.id}
  orgIdentifier: ${harness_platform_project.test.org_id}
  tags: {}
  spec:
    stageType: Deployment
    steps:
      - step:
          type: ShellScript
          name: Shell Script_1
          identifier: ShellScript_1
          spec:
            shell: Bash
            onDelegate: true
            source:
              type: Inline
              spec:
                script: <+input>
            environmentVariables: []
            outputVariables: []
          timeout: 10m

  EOT
}

## Inline Monitered Service template
resource "harness_platform_template" "monitered_service_template_inline" {
  identifier    = "identifier"
  org_id        = harness_platform_project.test.org_id
  project_id    = harness_platform_project.test.id
  name          = "name"
  comments      = "comments"
  version       = "ab"
  is_stable     = true
  template_yaml = <<-EOT
template:
  name: "name"
  identifier: "identifier"
  versionLabel: "ab"
  type: MonitoredService
  projectIdentifier: ${harness_platform_project.test.id}
  orgIdentifier: ${harness_platform_project.test.org_id}
  tags: {}
  spec:
    serviceRef: <+input>
    environmentRef: <+input>
    type: Application
    sources:
      changeSources:
        - name: Harness CD Next Gen
          identifier: harness_cd_next_gen
          type: HarnessCDNextGen
          enabled: true
          category: Deployment
          spec: {}
      healthSources:
        - name: health
          identifier: health
          type: AppDynamics
          spec:
            applicationName: <+input>
            tierName: <+input>
            metricData:
              Errors: true
              Performance: true
            metricDefinitions: []
            feature: Application Monitoring
            connectorRef: <+input>
            metricPacks:
              - identifier: Errors
              - identifier: Performance

  EOT
}

## Artifact Source template
resource "harness_platform_template" "artifact_source_template" {
  identifier    = "identifier"
  org_id        = harness_platform_project.test.org_id
  project_id    = harness_platform_project.test.id
  name          = "name"
  comments      = "comments"
  version       = "ab"
  is_stable     = true
  template_yaml = <<-EOT
template:
  name: "name"
  identifier: "identifier"
  versionLabel: "ab"
  type: ArtifactSource
  projectIdentifier: ${harness_platform_project.test.id}
  orgIdentifier: ${harness_platform_project.test.org_id}
  tags: {}
  spec:
    type: DockerRegistry
    spec:
      imagePath: library/nginx
      tag: <+input>
      connectorRef: account.Harness_DockerHub

  EOT
}

## Deployment template
resource "harness_platform_template" "deployment_template" {
  identifier    = "identifier"
  org_id        = harness_platform_project.test.org_id
  project_id    = harness_platform_project.test.id
  name          = "name"
  comments      = "comments"
  version       = "ab"
  is_stable     = true
  template_yaml = <<-EOT
template:
  name: "name"
  identifier: "identifier"
  versionLabel: "ab"
  type: CustomDeployment
  projectIdentifier: ${harness_platform_project.test.id}
  orgIdentifier: ${harness_platform_project.test.org_id}
  tags: {}
  spec:
    infrastructure:
      variables:
        - name: kubeConnector
          type: Connector
          value: <+input>
          description: ""
      fetchInstancesScript:
        store:
          type: Inline
          spec:
            content: |
              #
              # Script is expected to query Infrastructure and dump json
              # in $INSTANCE_OUTPUT_PATH file path
              #
              # Harness is expected to initialize ${INSTANCE_OUTPUT_PATH}
              # environment variable - a random unique file path on delegate,
              # so script execution can save the result.
              #
              /opt/harness-delegate/client-tools/kubectl/v1.19.2/kubectl get pods --namespace=harness-delegate-ng -o json > $INSTANCE_OUTPUT_PATH
      instanceAttributes:
        - name: instancename
          jsonPath: metadata.name
          description: ""
      instancesListPath: items
    execution:
      stepTemplateRefs: []

  EOT
}

## Secrets Manager template
resource "harness_platform_template" "secrets_manager_template" {
  identifier    = "identifier"
  org_id        = harness_platform_project.test.org_id
  project_id    = harness_platform_project.test.id
  name          = "name"
  comments      = "comments"
  version       = "ab"
  is_stable     = true
  template_yaml = <<-EOT
template:
  name: "name"
  identifier: "identifier"
  versionLabel: "ab"
  type: SecretManager
  projectIdentifier: ${harness_platform_project.test.id}
  orgIdentifier: ${harness_platform_project.test.org_id}
  tags: {}
  spec:
    executionTarget: {}
    shell: Bash
    onDelegate: true
    source:
      spec:
        script: |-
          curl -o secret.json -X GET https://example.com/<+secretManager.environmentVariables.enginename>/<+secretManager.environmentVariables.path> -H 'X-Vault-Token: <+secrets.getValue("vaultTokenOne")>'
          secret=$(jq -r '.data."<+secretManager.environmentVariables.key>"' secret.json)
        type: Inline
    environmentVariables:
      - name: enginename
        type: String
        value: <+input>
      - name: path
        type: String
        value: <+input>
      - name: key
        type: String
        value: <+input>


  EOT
}


### Creating Multiple Versions of a Template
##Stable version of the Template
resource "harness_platform_template" "template_v1" {
  identifier    = "temp"
  org_id        = harness_platform_project.test.org_id
  name          = "temp"
  comments      = "comments"
  version       = "v1"
  is_stable     = true
  force_delete  = true
  template_yaml = <<-EOT
			template:
      name: "temp"
      identifier: "temp"
      versionLabel: v1
      type: Pipeline
      orgIdentifier: ${harness_platform_organization.test.id}
      tags: {}
      spec:
        stages:
          - stage:
              name: dvvdvd
              identifier: dvvdvd
              description: ""
              type: Deployment
              spec:
                deploymentType: Kubernetes
                service:
                  serviceRef: <+input>
                  serviceInputs: <+input>
                environment:
                  environmentRef: <+input>
                  deployToAll: false
                  environmentInputs: <+input>
                  serviceOverrideInputs: <+input>
                  infrastructureDefinitions: <+input>
                execution:
                  steps:
                    - step:
                        name: Rollout Deployment
                        identifier: rolloutDeployment
                        type: K8sRollingDeploy
                        timeout: 10m
                        spec:
                          skipDryRun: false
                          pruningEnabled: false
                  rollbackSteps:
                    - step:
                        name: Rollback Rollout Deployment
                        identifier: rollbackRolloutDeployment
                        type: K8sRollingRollback
                        timeout: 10m
                        spec:
                          pruningEnabled: false
              tags: {}
              failureStrategies:
                - onFailure:
                    errors:
                      - AllErrors
                    action:
                      type: StageRollback
    
      EOT
}

##Unstable version of the Template
resource "harness_platform_template" "template_v2" {
  identifier    = "temp"
  org_id        = harness_platform_organization.test.id
  name          = "temp"
  comments      = "comments"
  version       = "v2"
  is_stable     = false
  force_delete  = true
  template_yaml = <<-EOT
			template:
      name: "temp"
      identifier: "temp"
      versionLabel: v2
      type: Pipeline
      orgIdentifier: ${harness_platform_organization.test.id}
      tags: {}
      spec:
        stages:
          - stage:
              name: dvvdvd
              identifier: dvvdvd
              description: ""
              type: Deployment
              spec:
                deploymentType: Kubernetes
                service:
                  serviceRef: <+input>
                  serviceInputs: <+input>
                environment:
                  environmentRef: <+input>
                  deployToAll: false
                  environmentInputs: <+input>
                  serviceOverrideInputs: <+input>
                  infrastructureDefinitions: <+input>
                execution:
                  steps:
                    - step:
                        name: Rollout Deployment
                        identifier: rolloutDeployment
                        type: K8sRollingDeploy
                        timeout: 10m
                        spec:
                          skipDryRun: false
                          pruningEnabled: false
                  rollbackSteps:
                    - step:
                        name: Rollback Rollout Deployment
                        identifier: rollbackRolloutDeployment
                        type: K8sRollingRollback
                        timeout: 10m
                        spec:
                          pruningEnabled: false
              tags: {}
              failureStrategies:
                - onFailure:
                    errors:
                      - AllErrors
                    action:
                      type: StageRollback
      EOT
}

##Updating the Stable Version of the Template from v1 to v2.
resource "harness_platform_template" "template_v2" {
  identifier    = "temp"
  org_id        = harness_platform_organization.test.id
  name          = "temp"
  comments      = "comments"
  version       = "v2"
  is_stable     = true
  force_delete  = true
  template_yaml = <<-EOT
			template:
      name: "temp"
      identifier: "temp"
      versionLabel: v2
      type: Pipeline
      orgIdentifier: ${harness_platform_organization.test.id}
      tags: {}
      spec:
        stages:
          - stage:
              name: dvvdvd
              identifier: dvvdvd
              description: ""
              type: Deployment
              spec:
                deploymentType: Kubernetes
                service:
                  serviceRef: <+input>
                  serviceInputs: <+input>
                environment:
                  environmentRef: <+input>
                  deployToAll: false
                  environmentInputs: <+input>
                  serviceOverrideInputs: <+input>
                  infrastructureDefinitions: <+input>
                execution:
                  steps:
                    - step:
                        name: Rollout Deployment
                        identifier: rolloutDeployment
                        type: K8sRollingDeploy
                        timeout: 10m
                        spec:
                          skipDryRun: false
                          pruningEnabled: false
                  rollbackSteps:
                    - step:
                        name: Rollback Rollout Deployment
                        identifier: rollbackRolloutDeployment
                        type: K8sRollingRollback
                        timeout: 10m
                        spec:
                          pruningEnabled: false
              tags: {}
              failureStrategies:
                - onFailure:
                    errors:
                      - AllErrors
                    action:
                      type: StageRollback
      EOT
}

resource "harness_platform_template" "template_v1" {
  identifier    = "temp"
  org_id        = harness_platform_organization.test.id
  name          = "temp"
  comments      = "comments"
  version       = "v1"
  is_stable     = false
  force_delete  = true
  template_yaml = <<-EOT
			template:
      name: "temp"
      identifier: "temp"
      versionLabel: v1
      type: Pipeline
      orgIdentifier: ${harness_platform_organization.test.id}
      tags: {}
      spec:
        stages:
          - stage:
              name: dvvdvd
              identifier: dvvdvd
              description: ""
              type: Deployment
              spec:
                deploymentType: Kubernetes
                service:
                  serviceRef: <+input>
                  serviceInputs: <+input>
                environment:
                  environmentRef: <+input>
                  deployToAll: false
                  environmentInputs: <+input>
                  serviceOverrideInputs: <+input>
                  infrastructureDefinitions: <+input>
                execution:
                  steps:
                    - step:
                        name: Rollout Deployment
                        identifier: rolloutDeployment
                        type: K8sRollingDeploy
                        timeout: 10m
                        spec:
                          skipDryRun: false
                          pruningEnabled: false
                  rollbackSteps:
                    - step:
                        name: Rollback Rollout Deployment
                        identifier: rollbackRolloutDeployment
                        type: K8sRollingRollback
                        timeout: 10m
                        spec:
                          pruningEnabled: false
              tags: {}
              failureStrategies:
                - onFailure:
                    errors:
                      - AllErrors
                    action:
                      type: StageRollback
    
      EOT

  depends_on = [time_sleep.wait_10_seconds]
}

resource "time_sleep" "wait_10_seconds" {
  depends_on       = [harness_platform_template.test2]
  destroy_duration = "10s"
}

##Importing Account Level Templates
resource "harness_platform_template" "test" {
  identifier      = "accounttemplate"
  name            = "accounttemplate"
  version         = "v2"
  is_stable       = false
  import_from_git = true
  git_import_details {
    branch_name   = "main"
    file_path     = ".harness/accounttemplate.yaml"
    connector_ref = "account.DoNotDeleteGithub"
    repo_name     = "open-repo"
  }
  template_import_request {
    template_name        = "accounttemplate"
    template_version     = "v2"
    template_description = ""
  }
}

##Importing Org Level Templates
resource "harness_platform_template" "test" {
  identifier      = "orgtemplate"
  name            = "orgtemplate"
  org_id          = "org"
  version         = "v2"
  is_stable       = false
  import_from_git = true
  git_import_details {
    branch_name   = "main"
    file_path     = ".harness/orgtemplate.yaml"
    connector_ref = "account.DoNotDeleteGithub"
    repo_name     = "open-repo"
  }
  template_import_request {
    template_name        = "orgtemplate"
    template_version     = "v2"
    template_description = ""
  }
}

##Importing Project Level Templates
resource "harness_platform_template" "test" {
  identifier      = "projecttemplate"
  name            = "projecttemplate"
  org_id          = "org"
  project_id      = "project"
  version         = "v2"
  is_stable       = false
  import_from_git = true
  git_import_details {
    branch_name   = "main"
    file_path     = ".harness/projecttemplate.yaml"
    connector_ref = "account.DoNotDeleteGithub"
    repo_name     = "open-repo"
  }
  template_import_request {
    template_name        = "projecttemplate"
    template_version     = "v2"
    template_description = ""
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `identifier` (String) Unique identifier of the resource
- `name` (String) Name of the Variable
- `version` (String) Version Label for Template.

### Optional

- `comments` (String) Specify comment with respect to changes.
- `description` (String, Deprecated) Description of the entity. Description field is deprecated
- `force_delete` (Boolean) Enable this flag for force deletion of template. It will delete the Harness entity even if your pipelines or other entities reference it
- `git_details` (Block List, Max: 1) Contains parameters related to creating an Entity for Git Experience. (see [below for nested schema](#nestedblock--git_details))
- `git_import_details` (Block List, Max: 1) Contains Git Information for importing entities from Git (see [below for nested schema](#nestedblock--git_import_details))
- `import_from_git` (Boolean) Flag to set if importing from Git
- `is_stable` (Boolean) True if given version for template to be set as stable.
- `org_id` (String) Organization Identifier for the Entity
- `project_id` (String) Project Identifier for the Entity
- `tags` (Set of String) Tags to associate with the resource.
- `template_import_request` (Block List, Max: 1) Contains parameters for importing template. (see [below for nested schema](#nestedblock--template_import_request))
- `template_yaml` (String) Yaml for creating new Template. In YAML, to reference an entity at the organization scope, prefix 'org' to the expression: org.{identifier}. To reference an entity at the account scope, prefix 'account` to the expression: account.{identifier}. For eg, to reference a connector with identifier 'connectorId' at the organization scope in a stage mention it as connectorRef: org.connectorId.

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
- `last_commit_id` (String) Last commit identifier (for Git Repositories other than Github). To be provided only when updating Pipeline.
- `last_object_id` (String) Last object identifier (for Github). To be provided only when updating Pipeline.
- `repo_name` (String) Name of the repository.
- `store_type` (String) Specifies whether the Entity is to be stored in Git or not. Possible values: INLINE, REMOTE.


<a id="nestedblock--git_import_details"></a>
### Nested Schema for `git_import_details`

Optional:

- `branch_name` (String) Name of the branch.
- `connector_ref` (String) Identifier of the Harness Connector used for importing entity from Git To reference a connector at the organization scope, prefix 'org' to the expression: org.{identifier}. To reference a connector at the account scope, prefix 'account` to the expression: account.{identifier}.
- `file_path` (String) File path of the Entity in the repository.
- `is_force_import` (Boolean)
- `repo_name` (String) Name of the repository.


<a id="nestedblock--template_import_request"></a>
### Nested Schema for `template_import_request`

Optional:

- `template_description` (String) Description of the template.
- `template_name` (String) Name of the template.
- `template_version` (String) Version of the template.

## Import

Import is supported using the following syntax:

The [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import) can be used, for example:

```shell
# Import account level template
terraform import harness_platform_template.example <template_id>

# Import org level template
terraform import harness_platform_template.example <ord_id>/<template_id>

# Import project level template
terraform import harness_platform_template.example <org_id>/<project_id>/<template_id>
```
