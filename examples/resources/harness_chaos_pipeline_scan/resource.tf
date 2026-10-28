# Scan an existing pipeline execution whose status is Success and that deployed
# with a Kubernetes CD stage. Apply waits until the scan finishes.
resource "harness_chaos_pipeline_scan" "example" {
  org_id     = "default"
  project_id = "demo"
  identity   = "cart_deploy_scan"
  name       = "Cart deploy scan"

  # "ai" is AI Scan. "ruleBased" is Rule-based Scan.
  detection_method = "ruleBased"

  pipeline_identity     = "cart_deploy"
  pipeline_execution_id = "exec-123"
}

output "pipeline_scan_risks" {
  value = harness_chaos_pipeline_scan.example.risks
}
