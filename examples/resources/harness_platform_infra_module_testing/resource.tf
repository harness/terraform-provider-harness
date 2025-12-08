resource "harness_platform_infra_module" "example" {
  module_id           = 1234
  org                 = "default"
  project             = "project"
  provider_connector  = "account.connector"
  provisioner_type    = "tofu"
  provisioner_version = "1.9.4"
  pipelines = [
    "pipelinea",
    "pipelineb"
  ]
}
