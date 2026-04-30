resource "harness_platform_iacm_ansible_inventory" "manual" {
  identifier = "my_inventory"
  name       = "my-inventory"
  org_id     = harness_platform_organization.example.id
  project_id = harness_platform_project.example.id
  type       = "manual"
  tags       = ["env:prod"]

  groups {
    identifier = "web"
    name       = "web"
    hosts      = ["web-1.example.com", "web-2.example.com"]
    vars {
      key        = "ansible_user"
      value      = "ubuntu"
      value_type = "string"
    }
  }

  vars {
    key        = "ansible_port"
    value      = "22"
    value_type = "string"
  }
}
