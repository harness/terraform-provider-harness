resource "harness_platform_ff_api_key" "testserverapikey" {
  identifier  = "testserver"
  name        = "TestServer"
  description = "this is a server SDK key"
  org_id      = "test"
  project_id  = "testff"
  env_id      = "testenv"
  expired_at  = 1713729225
  type        = "Server"
}

output "serversdkkey" {
  value     = harness_platform_ff_api_key.testserverapikey.api_key
  sensitive = true
}