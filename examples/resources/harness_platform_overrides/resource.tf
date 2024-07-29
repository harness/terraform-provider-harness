resource "harness_platform_overrides" "test" {
  org_id     = "orgIdentifier"
  project_id = "projectIdentifier"
  env_id     = "environmentIdentifier"
  service_id = "serviceIdentifier"
  infra_id   = "infraIdentifier"
  type       = "INFRA_SERVICE_OVERRIDE"

  ## Git Details are required in case the overrides are remote
    git_details {
    branch_name    = "branchName"
    commit_message = "commitMessage"
    file_path      = "filePath"
    connector_ref  = "connectorRef"
    store_type     = "REMOTE"
    repo_name      = "repoName"
  }
  yaml       = <<-EOT
    variables:
      - name: var1
        type: String
        value: val1
    configFiles:
      - configFile:
          identifier: sampleConfigFile
          spec:
            store:
              type: Harness
              spec:
                files:
                  - account:/configFile1
    manifests:
      - manifest:
          identifier: sampleManifestFile
          type: Values
          spec:
            store:
              type: Harness
              spec:
                files:
                  - account:/manifestFile1
  EOT
}

### Importing Override from Git
  resource "harness_platform_overrides" "test" {
          org_id     = "orgIdentifier"
          project_id = "projectIdentifier"
          env_id     = "environmentIdentifier"
          service_id = "serviceIdentifier"
          infra_id = "infraIdentifier"
          type       = "INFRA_SERVICE_OVERRIDE"
          import_from_git = "true"
          git_details {
            store_type = "REMOTE"
            connector_ref = "connector_ref"
            repo_name = "repo_name"
            file_path = "file_path"
            branch = "branch"
            }
            
}