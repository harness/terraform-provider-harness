# Creating a Kubernetes infrastructure definition

resource "harness_cloudprovider_kubernetes" "dev" {
  name = "k8s-dev"

  authentication {
    delegate_selectors = ["k8s"]
  }
}

resource "harness_application" "example" {
  name = "example"
}

resource "harness_environment" "dev" {
  name   = "dev"
  app_id = harness_application.example.id
  type   = "NON_PROD"
}

# Creating a infrastructure of type KUBERNETES
resource "harness_infrastructure_definition" "k8s" {
  name                = "k8s-eks-us-east-1"
  app_id              = harness_application.example.id
  env_id              = harness_environment.dev.id
  cloud_provider_type = "KUBERNETES_CLUSTER"
  deployment_type     = "KUBERNETES"

  kubernetes {
    cloud_provider_name = harness_cloudprovider_kubernetes.dev.name
    namespace           = "dev"
    release_name        = "$${service.name}"
  }
}

# Creating a Deployment Template for CUSTOM infrastructure type
resource "harness_yaml_config" "example_yaml" {
  path = "Setup/Template Library/Example Folder/deployment_template.yaml"
  content = <<EOF
harnessApiVersion: '1.0'
type: CUSTOM_DEPLOYMENT_TYPE
fetchInstanceScript: |-
  set -ex
  curl http://$${url}/$${file_name} > $${INSTANCE_OUTPUT_PATH}
hostAttributes:
  hostname: host
hostObjectArrayPath: hosts
variables:
- name: url
- name: file_name
EOF
}

# Creating a infrastructure of type CUSTOM
resource "harness_infrastructure_definition" "custom" {
  name = "custom-infra"
  app_id = harness_application.example.id
  env_id = harness_environment.dev.id
  cloud_provider_type = "CUSTOM"
  deployment_type = "CUSTOM"
  deployment_template_uri = "Example Folder/${harness_yaml_config.example_yaml.name}"

  custom {
    deployment_type_template_version = "1"
    variable {
      name = "url"
      value = "localhost:8081"
    }

    variable {
      name = "file_name"
      value = "instances.json"
    }
  }
}
