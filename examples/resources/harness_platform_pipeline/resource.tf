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
