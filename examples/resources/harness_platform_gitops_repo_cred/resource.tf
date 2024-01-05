resource "harness_platform_gitops_repo_cred" "test" {
  identifier = "identifier"
  account_id = "account_id"
  agent_id   = "agent_id"
  project_id = "project_id"
  org_id     = "org_id"
  creds {
    type            = "git"
    url             = "git@github.com:yourorg"
    ssh_private_key = "----- BEGIN OPENSSH PRIVATE KEY-----\nXXXXX\nXXXXX\nXXXXX\n-----END OPENSSH PRIVATE KEY -----\n"
  }
  lifecycle {
    ignore_changes = [
      account_id,
      creds.0.ssh_private_key,
    ]
  }
}
