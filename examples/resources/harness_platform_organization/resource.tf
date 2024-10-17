resource "harness_platform_organization" "this" {
  identifier  = "MyOrg"
  name        = "My Organization"
  description = "An example organization"
  tags        = ["foo:bar", "baz:qux"]
}
