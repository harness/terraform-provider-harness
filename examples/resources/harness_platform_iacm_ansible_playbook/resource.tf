resource "harness_platform_iacm_ansible_playbook" "example" {
  identifier           = "my_playbook"
  name                 = "my-playbook"
  org_id               = harness_platform_organization.example.id
  project_id           = harness_platform_project.example.id
  repository           = "https://github.com/org/repo"
  repository_branch    = "main"
  repository_path      = "ansible/site.yml"
  repository_connector = "account.my_github_connector"
  ansible_galaxy       = true
  tags                 = ["env:prod"]

  vars {
    key        = "environment"
    value      = "production"
    value_type = "string"
  }

  env_vars {
    key        = "ANSIBLE_CONFIG"
    value      = "ansible.cfg"
    value_type = "string"
  }
}
