# The token is available during the run but is never stored in state or plan files.
# Harness returns it only while the agent has never connected.
ephemeral "harness_platform_gitops_agent_token" "example" {
  identifier = "example-agent"
  org_id     = "default"
  project_id = "example-project"
}

# Reference ephemeral.harness_platform_gitops_agent_token.example.agent_token only from
# ephemeral-aware contexts, such as a write-only argument on another resource or a
# provider configuration block.
