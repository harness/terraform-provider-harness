resource "harness_autostopping_azure_proxy" "test" {
  name               = "name"
  cloud_connector_id = "cloud_connector_id"
  host_name          = "host_name"
  region             = "eastus2"
  resource_group     = "resource_group"
  vpc                = "/subscriptions/subscription_id/resourceGroups/resource_group/providers/Microsoft.Network/virtualNetworks/virtual_network"
  subnet_id          = "/subscriptions/subscription_id/resourceGroups/resource_group/providers/Microsoft.Network/virtualNetworks/virtual_network/subnets/subnet_id"
  security_groups    = ["/subscriptions/subscription_id/resourceGroups/resource_group/providers/Microsoft.Network/networkSecurityGroups/network_security_group"]
  allocate_static_ip = true
  machine_type       = "Standard_D2s_v3"
  keypair            = "keypair"
  api_key            = ""
}

