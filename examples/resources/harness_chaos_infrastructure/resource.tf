
# Sample resource for chaos infrastructure
resource "harness_chaos_infrastructure" "example" {
  identifier = "identifier"
  name      = "name"
  org_id     = "org_id"
  project_id = "project_id"
  environment_id= "env_id"
  namespace= "namespace"
  service_account= "service_acc_name"
}