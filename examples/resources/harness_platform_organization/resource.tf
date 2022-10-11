resource "harness_organization" "this" {
  identifier  = "MyOrg"
  name        = "My Otganization"
  description = "An example organization"
  tags        = ["foo:bar", "baz:qux"]
}
