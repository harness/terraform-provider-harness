# Scan an existing pipeline execution.
resource "harness_chaos_pipeline_scan" "example" {
  org_id     = "default"
  project_id = "demo"
  identity   = "cart_deploy_scan"
  name       = "Cart deploy scan"

  detection_method = "ruleBased"

  pipeline_identity     = "cart_deploy"
  pipeline_execution_id = "exec-123"
}

output "pipeline_scan_risks" {
  value = harness_chaos_pipeline_scan.example.risks
}
