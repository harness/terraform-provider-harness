resource "harness_platform_input_set" "example" {
  identifier = "identifier"
  name       = "name"
  tags = [
    "foo:bar",
  ]
  org_id      = "org_id"
  project_id  = "project_id"
  pipeline_id = "pipeline_id"
  yaml        = <<-EOT
      inputSet:
        identifier: "identifier"
        name: "name"
        tags:
          foo: "bar"
        orgIdentifier: "org_id"
        projectIdentifier: "project_id"
        pipeline:
          identifier: "pipeline_id"
          variables:
          - name: "key"
            type: "String"
            value: "value"
  EOT
}

# Remote InputSet
resource "harness_platform_input_set" "test" {
  identifier = "identifier"
  name       = "name"
  tags = [
    "foo:bar",
  ]
  org_id      = harness_platform_organization.test.id
  project_id  = harness_platform_project.test.id
  pipeline_id = harness_platform_pipeline.test.id
  git_details {
    branch_name    = "main"
    commit_message = "Commit"
    file_path      = ".harness/file_path.yaml"
    connector_ref  = "account.connector_ref"
    store_type     = "REMOTE"
    repo_name      = "repo_name"
  }
  yaml = <<-EOT
inputSet:
  identifier: "identifier"
  name: "name"
  tags:
    foo: "bar"
  orgIdentifier: "${harness_platform_organization.test.id}"
  projectIdentifier: "${harness_platform_project.test.id}"
  pipeline:
    identifier: "${harness_platform_pipeline.test.id}"
    variables:
    - name: "key"
      type: "String"
      value: "value"
EOT
}
