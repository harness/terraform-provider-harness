resource "harness_autostopping_aws_proxy" "test" {
  name                              = "name"
  cloud_connector_id                = "cloud_connector_id"
  region                            = "region"
  vpc                               = "vpc"
  security_groups                   = ["sg1", "sg2"]
  machine_type                      = "t2.medium"
  api_key                           = ""
  allocate_static_ip                = true
  delete_cloud_resources_on_destroy = true
}

