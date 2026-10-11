# Webhook URLs of a Custom integration, to configure in the sending system.
data "harness_platform_idp_integration" "example" {
  identifier = "orders_webhook"
}

output "webhook_urls" {
  value = data.harness_platform_idp_integration.example.endpoints[*].url
}
