resource "harness_platform_project" "terraform_project" {
	depends_on = [
		harness_platform_organization.terraform_org
	]
	identifier = "terraform_project"
	name = "terraform_project"
	org_id = harness_platform_organization.terraform_org.id
	color = "#472848"
}
