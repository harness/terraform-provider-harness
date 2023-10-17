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

resource "harness_autostopping_azure_gateway" "import_test" {
  name               = "import_test"
  cloud_connector_id = "cloud_connector_id"
  host_name          = "host_name"
  region             = "westus2"
  resource_group     = "test_resource_group"
  app_gateway_id     = "/subscriptions/subscription_id/resourceGroups/test_resource_group/providers/Microsoft.Network/applicationGateways/TestAppGateway"
  certificate_id     = "/subscriptions/subscription_id/resourceGroups/test_resource_group/providers/Microsoft.Network/applicationGateways/TestAppGateway/sslCertificates/certificate_name"
  azure_func_region  = "westus2"
  vpc                = "/subscriptions/subscription_id/resourceGroups/test_resource_group/providers/Microsoft.Network/virtualNetworks/test_resource_group_vnet"
}
