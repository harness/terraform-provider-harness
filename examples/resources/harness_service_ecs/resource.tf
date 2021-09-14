resource "harness_application" "example" {
  name = "example"
}

resource "harness_service_ecs" "example" {
  app_id      = harness_application.example.id
  name        = "ecs-example-service"
  description = "Service for deploying AWS ECS tasks."
}
