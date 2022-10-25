resource "harness_platform_triggers" "example" {
  identifier = "identifier"
  org_id     = "orgIdentifer"
  project_id = "projectIdentifier"
  name       = "name"
  target_id  = "pipelineIdentifier"
  yaml       = <<-EOT
  trigger:
    name: "name"
    identifier: "identifier"
    enabled: true
    description: ""
    tags: {}
    projectIdentifier: "projectIdentifier"
    orgIdentifier: "orgIdentifer"
    pipelineIdentifier: "pipelineIdentifier"
    source:
      type: "Webhook"
      spec:
        type: "Github"
        spec:
          type: "Push"
          spec:
            connectorRef: "account.TestAccResourceConnectorGithub_Ssh_IZBeG"
            autoAbortPreviousExecutions: false
            payloadConditions:
            - key: "changedFiles"
              operator: "Equals"
              value: "value"
            - key: "targetBranch"
              operator: "Equals"
              value: "value"
            headerConditions: []
            repoName: "repoName"
            actions: []
    inputYaml: "pipeline: {}\n"
    EOT
}
