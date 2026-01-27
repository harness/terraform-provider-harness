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

# Google Managed Instance Group Infrastructure Example
resource "harness_platform_infrastructure" "google_mig_example" {
  identifier      = "google_mig_infra"
  name            = "Google MIG Infrastructure"
  org_id          = "default"
  project_id      = "my_project"
  env_id          = "my_environment"
  type            = "GoogleManagedInstanceGroup"
  deployment_type = "GoogleManagedInstanceGroup"
  yaml            = <<-EOT
        infrastructureDefinition:
          name: Google MIG Infrastructure
          identifier: google_mig_infra
          description: "Infrastructure for Google Managed Instance Group deployments"
          orgIdentifier: default
          projectIdentifier: my_project
          environmentRef: my_environment
          deploymentType: GoogleManagedInstanceGroup
          type: GoogleManagedInstanceGroup
          spec:
            connectorRef: account.gcpConnector
            region: us-central1
            zone: us-central1-a
            project: my-gcp-project
      EOT
}
