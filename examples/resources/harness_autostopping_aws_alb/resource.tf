resource "harness_autostopping_aws_alb" "test" {
  name                              = "name"
  cloud_connector_id                = "cloud_connector_id"
  region                            = "region"
  vpc                               = "vpc"
  security_groups                   = ["sg1", "sg2"]
  delete_cloud_resources_on_destroy = true
}

resource "harness_autostopping_aws_alb" "harness_alb" {
  name                              = "harness_alb"
  cloud_connector_id                = "cloud_connector_id"
  alb_arn                           = "arn:aws:elasticloadbalancing:region:aws_account_id:loadbalancer/app/harness_alb/id"
  region                            = "region"
  vpc                               = "vpc"
  security_groups                   = ["sg-0"]
  delete_cloud_resources_on_destroy = false
}
