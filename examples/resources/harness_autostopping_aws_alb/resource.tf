resource "harness_autostopping_aws_alb" "test" {
  name                   = "name"
  cloud_connector_id     = "cloud_connector_id"
  host_name              = "host_name"
  region                 = "region"
  vpc                    = "vpc"
  security_groups        = ["sg1", "sg2"]
  route53_hosted_zone_id = "/hostedzone/zone_id"
}

