data "harness_platform_iacm_ansible_inventory" "example" {
  identifier = "my_inventory"
  org_id     = harness_platform_organization.example.id
  project_id = harness_platform_project.example.id
}
