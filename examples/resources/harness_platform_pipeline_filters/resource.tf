resource "harness_platform_pipeline_filters" "test" {
  identifier = "identifier"
  name       = "name"
  org_id     = "org_id"
  project_id = "project_id"
  type       = "PipelineSetup"
  filter_properties {
    name                 = "pipeline_name"
    description          = "pipeline_description"
    pipeline_identifiers = ["id1", "id2"]
    filter_type          = "PipelineSetup"
  }
  filter_visibility = "EveryOne"
}


# pipeline filter with tags
resource "harness_platform_pipeline_filters" "example_with_tags" {
  identifier = "identifier"
  name       = "name"
  org_id     = "org_id"
  project_id = "project_id"
  type       = "PipelineSetup"
  filter_properties {
    filter_type   = "PipelineSetup"
    pipeline_tags = [
      {
        "key"   = "tag1"
        "value" = "123"
      },
      {
        "key"   = "tag2"
        "value" = "456"
      },
    ]
    module_properties {
      cd {
        deployment_types       = "Kubernetes"
        service_names          = ["service1", "service2"]
        environment_names      = ["env1", "env2"]
        artifact_display_names = ["artificatname1", "artifact2"]
      }
      ci {
        build_type = "branch"
        branch     = "branch123"
        repo_names = "repo1234"
      }
    }
  }
}