data "harness_platform_iacm_ansible_playbook" "example" {
  identifier = "my_playbook"
  org_id     = harness_platform_organization.example.id
  project_id = harness_platform_project.example.id
}
