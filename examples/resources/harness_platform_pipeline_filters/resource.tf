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

# pipeline execution filter consisiting services (service_identifiers) filter
resource "harness_platform_pipeline_filters" "execution" {
  identifier = "identifier"
  name       = "name"
  org_id     = "org_id"
  project_id = "project_id"
  type       = "PipelineSetup"
  filter_properties {
    name                 = "pipeline_name"
    description          = "pipeline_description"
    pipeline_identifiers = ["id1", "id2"]
    filter_type          = "PipelineExecution"
    module_properties {
      cd {
        deployment_types    = "Kubernetes"
        service_identifiers = ["nginx"]
      }
    }
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
    filter_type = "PipelineSetup"
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

resource "harness_platform_pipeline_filters" "pipelinemoduleproperties" {
  identifier = "identifier"
  name       = "name"
  org_id     = harness_platform_project.test.org_id
  project_id = harness_platform_project.test.id
  type       = "PipelineExecution"
  filter_properties {
    filter_type   = "PipelineExecution"
    pipeline_name = "test"
    pipeline_tags = [
      {
        "key"   = "k1"
        "value" = "v1"
      },
      {
        "key"   = "k2"
        "value" = "v2"
      },
    ]
    module_properties {
      cd {
        service_definition_types = "Kubernetes"
        service_identifiers      = ["K8"]
        environment_identifiers  = ["dev"]
        artifact_display_names   = ["artificatname1"]
      }
    }
  }
}
