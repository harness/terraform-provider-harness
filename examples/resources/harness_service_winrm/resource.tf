resource "harness_application" "example" {
  name = "example"
}

resource "harness_service_winrm" "example" {
  app_id        = harness_application.example.id
  artifact_type = "IIS_APP"
  name          = "iis-app-winrm-svc"
  description   = "Service for deploying IIS appliactions using winrm."
}
