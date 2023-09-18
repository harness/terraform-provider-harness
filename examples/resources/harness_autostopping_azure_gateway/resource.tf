resource "harness_autostopping_azure_gateway" "test" {
  name               = "name"
  cloud_connector_id = "cloud_connector_id"
  host_name          = "host_name"
  region             = "eastus2"
  resource_group     = "resource_group"
  subnet_id          = "/subscriptions/subscription_id/resourceGroups/resource_group/providers/Microsoft.Network/virtualNetworks/virtual_network/subnets/subnet_id"
  vpc                = "/subscriptions/subscription_id/resourceGroups/resource_group/providers/Microsoft.Network/virtualNetworks/virtual_network"
  azure_func_region  = "westus2"
  frontend_ip        = "/subscriptions/subscription_id/resourceGroups/resource_group/providers/Microsoft.Network/publicIPAddresses/publicip"
  sku_size           = "sku2"
}
