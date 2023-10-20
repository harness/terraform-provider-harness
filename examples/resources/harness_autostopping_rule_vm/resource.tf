resource "harness_autostopping_rule_vm" "test" {
  name               = "name"
  cloud_connector_id = "cloud_connector_id"
  idle_time_mins     = 10
  filter {
    vm_ids  = ["/subscriptions/subscription_id/resourceGroups/resource_group/providers/Microsoft.Compute/virtualMachines/virtual_machine"]
    regions = ["useast2"]
  }
  http {
    proxy_id = "proxy_id"
    routing {
      source_protocol = "https"
      target_protocol = "https"
      source_port     = 443
      target_port     = 443
      action          = "forward"
    }
    routing {
      source_protocol = "http"
      target_protocol = "http"
      source_port     = 80
      target_port     = 80
      action          = "forward"
    }
    health {
      protocol         = "http"
      port             = 80
      path             = "/"
      timeout          = 30
      status_code_from = 200
      status_code_to   = 299
    }
  }
  tcp {
    proxy_id = "proxy_id"
    ssh {
      port = 22
    }
    rdp {
      port = 3389
    }
    forward_rule {
      port = 2233
    }
  }
  depends {
    rule_id      = 24576
    delay_in_sec = 5
  }
}
