---
subcategory: "Next Gen"
page_title: "GitOps Agent Disaster Recovery"
description: |-
    How to create, promote, and delete GitOps agents configured for Disaster Recovery (DR).
---

# GitOps Agent Disaster Recovery

A Disaster Recovery (DR) group is a set of `harness_platform_gitops_agent` resources that share the
same `disaster_recovery_identifier`. Exactly one node in the group is `PRIMARY` and serves traffic;
any other nodes are `SECONDARY` and stand by until promoted.

- `disaster_recovery_identifier` — shared by every node in the group. Must not contain `:` or `/`.
- `disaster_recovery_type` — this node's role, `PRIMARY` or `SECONDARY`. Required together with
  `disaster_recovery_identifier`.
- `disaster_recovery_live_type` (computed) — the role the server currently reports for this node.
  Only populated by `Read`/`Import`; use this, not `disaster_recovery_type`, to check actual live
  state.
- `uuid` (computed) — shared by every sibling in the same DR group.

The DR node's internal `name` is not a separate Terraform field — it's set automatically to
`disaster_recovery_identifier`.

## Creating a DR PRIMARY agent

```terraform
resource "harness_platform_gitops_agent" "primary" {
  identifier = "my_agent_primary"
  org_id     = "org_id"
  project_id = "project_id"
  name       = "my-agent-primary"
  type       = "MANAGED_ARGO_PROVIDER"

  disaster_recovery_identifier = "my_agent_dr"
  disaster_recovery_type       = "PRIMARY"

  metadata {
    namespace = "argocd"
  }
}
```

## Creating a DR SECONDARY agent

The `SECONDARY` node is a separate resource with its own `identifier`, sharing the same
`disaster_recovery_identifier` as the `PRIMARY`:

```terraform
resource "harness_platform_gitops_agent" "secondary" {
  identifier = "my_agent_secondary"
  org_id     = "org_id"
  project_id = "project_id"
  name       = "my-agent-secondary"
  type       = "MANAGED_ARGO_PROVIDER"

  disaster_recovery_identifier = "my_agent_dr"
  disaster_recovery_type       = "SECONDARY"

  metadata {
    namespace = "argocd"
  }
}
```

A `SECONDARY` can only be created as a brand-new resource — you cannot convert an existing
(already-applied) non-DR agent directly to `SECONDARY`. Converting an existing non-DR agent into a
DR node is only supported going to `PRIMARY`.

## Promoting a SECONDARY to PRIMARY

Promotion is a separate resource, `harness_platform_gitops_agent_dr_promotion`, not a change to the
agent resource itself:

```terraform
resource "harness_platform_gitops_agent_dr_promotion" "promote" {
  org_id        = "org_id"
  project_id    = "project_id"
  agent_id      = "my_agent_secondary"
  dr_identifier = "my_agent_dr"

  disable_sync_windows = false

  # Any change to `triggers` forces replacement and re-runs the promotion.
  triggers = {
    promoted_at = "2024-01-01T00:00:00Z"
  }
}
```

Notes:
- `current_primary` (computed) reports the `dr_identifier` of whichever node the server currently
  considers `PRIMARY`, so you can confirm the promotion took effect.
- Promotion is not reversible via this resource — destroying it is a no-op (it only removes it from
  state). To promote again (e.g. fail back), change `triggers` to force a new promotion.
- After a promotion, the `disaster_recovery_type` you set in each agent's own config does **not**
  change — only `disaster_recovery_live_type` reflects the new live roles. Read `disaster_recovery_live_type`
  (or `current_primary` on the promotion resource) rather than assuming the config value matches reality.

## Deleting agents

- **Non-DR agent**: deleting/destroying the resource deletes it directly, no special handling.
- **DR SECONDARY**: deletes directly, no guard.
- **DR PRIMARY**: before deleting, the provider checks whether this node is still the live `PRIMARY`
  and whether a live `SECONDARY` sibling still exists. If both are true, the delete is blocked:

  ```
  cannot delete live PRIMARY node 'my_agent_primary' while SECONDARY sibling(s) still exist for
  disaster_recovery_identifier 'my_agent_dr' — promote a SECONDARY to PRIMARY first, or destroy the
  SECONDARY resource(s) first
  ```

  To delete a live `PRIMARY` with an active `SECONDARY`, either promote the `SECONDARY` first (so
  the node you're deleting is no longer `PRIMARY`), or destroy the `SECONDARY` resource first.

  This check is check-then-act, not atomic: a `SECONDARY` could in principle be created or promoted
  between the check and the delete call.

## Importing DR agents

Import accepts an optional `:<disaster_recovery_identifier>` suffix to target a specific DR node:

```shell
# Account level
terraform import harness_platform_gitops_agent.example <agent_id>:<disaster_recovery_identifier>

# Org level
terraform import harness_platform_gitops_agent.example <org_id>/<agent_id>:<disaster_recovery_identifier>

# Project level
terraform import harness_platform_gitops_agent.example <org_id>/<project_id>/<agent_id>:<disaster_recovery_identifier>
```

Omit the `:<disaster_recovery_identifier>` suffix for a non-DR agent.

## Other things to know

- **`token_rotation` is not supported on DR agents.** You cannot change `token_rotation` while
  `disaster_recovery_identifier` is set — credential rotation is not currently DR-aware.
- **Treat `disaster_recovery_type` as set-once.** Once a node has a `disaster_recovery_type`, further
  config changes to that field are not applied on subsequent plans/applies. Clearing
  `disaster_recovery_identifier`/`disaster_recovery_type` to turn a DR node back into a regular agent
  is not currently a supported operation — don't rely on it. If you need to decommission a DR node,
  destroy and recreate the resource instead.
- **`disaster_recovery_identifier` changes force replacement** of the agent (moving a node to a
  different DR group isn't an in-place update).
- Use the `harness_platform_gitops_agent_deploy_yaml` data source's `disaster_recovery_identifier`
  argument to fetch the deploy YAML for a specific DR node.
