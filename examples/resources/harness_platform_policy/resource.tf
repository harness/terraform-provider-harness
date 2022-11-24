resource "harness_platform_policy" "test" {
	identifier = harness_platform_policy.test.identifier
	name = harness_platform_policy.test.name
	rego = "package test"
}
