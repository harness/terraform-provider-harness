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