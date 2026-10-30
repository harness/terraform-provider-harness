# Keep the GitOps agent token out of Terraform state (requires Terraform >= 1.11).
#
#   First install:   terraform apply -var bootstrap=true
#   Normal runs:     terraform apply
#   Rotate the key:  terraform apply -var bootstrap=true -var token_rotation=1 -var token_revision=2
#                    then keep passing token_rotation=1 and token_revision=2 on later runs, and
#                    restart the agent pods so they pick up the new Secret.

variable "bootstrap" {
  type    = bool
  default = false
}

variable "token_rotation" {
  type    = number
  default = 0
}

variable "token_revision" {
  type    = number
  default = 1
}

resource "harness_platform_gitops_agent" "example" {
  identifier                = "example"
  name                      = "example"
  org_id                    = "default"
  project_id                = "default_project"
  type                      = "MANAGED_ARGO_PROVIDER"
  operator                  = "ARGO"
  skip_agent_token_in_state = true
  token_rotation            = var.token_rotation

  metadata {
    namespace         = "argocd"
    high_availability = false
  }
}

# Source 1: Harness. Only read while bootstrapping or rotating.
ephemeral "harness_platform_gitops_agent_token" "example" {
  count      = var.bootstrap ? 1 : 0
  identifier = harness_platform_gitops_agent.example.identifier
  org_id     = "default"
  project_id = "default_project"
  generation = harness_platform_gitops_agent.example.token_generation
}

# Source 2: AWS Secrets Manager. Read on every other run.
ephemeral "aws_secretsmanager_secret_version" "agent_token" {
  count     = var.bootstrap ? 0 : 1
  secret_id = aws_secretsmanager_secret.agent_token.id
}

locals {
  agent_token = var.bootstrap ? ephemeral.harness_platform_gitops_agent_token.example[0].agent_token : ephemeral.aws_secretsmanager_secret_version.agent_token[0].secret_string
}

resource "aws_secretsmanager_secret" "agent_token" {
  name = "harness/gitops-agent/example"
}

# Written only while bootstrapping or rotating. Dropping this resource later does not delete the secret value.
resource "aws_secretsmanager_secret_version" "agent_token" {
  count                    = var.bootstrap ? 1 : 0
  secret_id                = aws_secretsmanager_secret.agent_token.id
  secret_string_wo         = local.agent_token
  secret_string_wo_version = var.token_revision
}

# Always keep this resource in the configuration. Removing it, or setting data_wo to null, empties or deletes the Secret.
resource "kubernetes_secret_v1" "agent_token" {
  metadata {
    name      = "gitops-agent-token"
    namespace = "argocd"
  }
  data_wo          = { GITOPS_AGENT_TOKEN = local.agent_token }
  data_wo_revision = var.token_revision
}
