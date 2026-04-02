resource "harness_platform_service" "example" {
  identifier  = "identifier"
  name        = "name"
  description = "test"
  org_id      = "org_id"
  project_id  = "project_id"
  git_details {
    branch_name    = "branchName"
    commit_message = "commitMessage"
    file_path      = "filePath"
    connector_ref  = "connectorRef"
    store_type     = "REMOTE"
    repo_name      = "repoName"
  }

  ## SERVICE V2 UPDATE
  ## We now take in a YAML that can define the service definition for a given Service
  ## It isn't mandatory for Service creation 
  ## It is mandatory for Service use in a pipeline

  yaml = <<-EOT
                service:
                  name: name
                  identifier: identifier
                  serviceDefinition:
                    spec:
                      manifests:
                        - manifest:
                            identifier: manifest1
                            type: K8sManifest
                            spec:
                              store:
                                type: Github
                                spec:
                                  connectorRef: <+input>
                                  gitFetchType: Branch
                                  paths:
                                    - files1
                                  repoName: <+input>
                                  branch: master
                              skipResourceVersioning: false
                      configFiles:
                        - configFile:
                            identifier: configFile1
                            spec:
                              store:
                                type: Harness
                                spec:
                                  files:
                                    - <+org.description>
                      variables:
                        - name: var1
                          type: String
                          value: val1
                        - name: var2
                          type: String
                          value: val2
                    type: Kubernetes
                  gitOpsEnabled: false
              EOT
}

### Importing Service from Git
resource "harness_platform_service" "test" {
  identifier      = "id"
  name            = "name"
  org_id          = "org_id"
  project_id      = "project_id"
  import_from_git = "true"
  git_details {
    store_type    = "REMOTE"
    connector_ref = "account.DoNotDeleteGitX"
    repo_name     = "pcf_practice"
    file_path     = ".harness/accountService.yaml"
    branch        = "main"
  }
}


### Google Managed Instance Group Service Example
resource "harness_platform_service" "google_mig_service" {
  identifier  = "google_mig_svc"
  name        = "Google MIG Service"
  description = "Service for Google Managed Instance Group deployments"
  org_id      = "default"
  project_id  = "my_project"

  yaml = <<-EOT
    service:
      name: Google MIG Service
      identifier: google_mig_svc
      description: Service for Google Managed Instance Group deployments
      serviceDefinition:
        type: GoogleManagedInstanceGroup
        spec:
          manifests:
            - manifest:
                identifier: instanceTemplate
                type: GoogleMigInstanceTemplate
                spec:
                  store:
                    type: Github
                    spec:
                      connectorRef: account.github
                      repoName: my-repo
                      branch: main
                      paths:
                        - templates/instance-template.json
            - manifest:
                identifier: migConfig
                type: GoogleMigConfiguration
                spec:
                  store:
                    type: Github
                    spec:
                      connectorRef: account.github
                      repoName: my-repo
                      branch: main
                      paths:
                        - templates/mig-config.json
          artifacts:
            primary:
              primaryArtifactRef: <+input>
              sources:
                - identifier: gceImage
                  type: GceImage
                  spec:
                    connectorRef: account.gcpConnector
                    project: my-gcp-project
          variables:
            - name: targetSize
              type: Number
              value: 2
      gitOpsEnabled: false
  EOT
}
