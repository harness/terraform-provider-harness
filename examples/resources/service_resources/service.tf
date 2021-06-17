resource "harness_application" "myapp" {
  name = "test-app-for-service"
}

resource "harness_service" "myservice" {
  
  app_id = harness_application.myapp.id
  name = "my-service"
  artifact_type = "DOCKER"
  deployment_type = "KUBERNETES"
}

resource "harness_service_kubernetes" {
  app_id = harness_application.myapp.id
  name = "my-k8s-service"
  helm_version = "V3"
}

resource "harness_service_ami" {
  app_id = harness_application.myapp.id
  name = "ami-service"
}
