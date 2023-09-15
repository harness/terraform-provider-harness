resource "harness_autostopping_aws_proxy" "test" {
  name                   = "name"
  cloud_connector_id     = "cloud_connector_id"
  host_name              = "host_name"
  region                 = "region"
  vpc                    = "vpc"
  security_groups        = ["sg1", "sg2"]
  route53_hosted_zone_id = "/hostedzone/zone_id"
  machine_type           = "t2.medium"
  api_key                = ""
  allocate_static_ip     = true
}

