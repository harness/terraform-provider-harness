data "harness_chaos_scan" "example" {
  org_id     = "default"
  project_id = "demo"
  identity   = "cart_deploy_scan"
}

output "scan_status" {
  value = data.harness_chaos_scan.example.status
}

output "scan_risks" {
  value = data.harness_chaos_scan.example.risks
}
