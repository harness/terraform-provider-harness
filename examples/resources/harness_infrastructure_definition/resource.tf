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
