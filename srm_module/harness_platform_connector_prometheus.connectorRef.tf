resource "harness_platform_connector_prometheus" "connectorRef" {
  depends_on = [
    harness_platform_project.terraform_project
  ]
  identifier  = "connectorRef"
  name        = "connectorRef"
  description = "prometheus"
  tags        = ["foo:bar"]
  org_id = harness_platform_organization.terraform_org.id
  project_id = harness_platform_project.terraform_project.id
  delegate_selectors = ["stress-chi-play-med-ng"]
  url = "http://10.4.2.99:80/health-source/prometheus/"
}
