resource "harness_organization" "test" {
  identifier  = "MyOrg"
  name        = "My Otganization"
  description = "An example organization"
  tags        = ["foo:bar", "baz:qux"]
}
