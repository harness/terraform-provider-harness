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
