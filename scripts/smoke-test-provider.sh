#!/bin/bash
#
# Smoke test: build the provider binary and make Terraform load its schema.
#
# Catches failures that only appear when the real binary is served to Terraform, such as the
# tf5muxserver "Invalid Provider Server Combination" error (differing SDKv2 and framework
# provider schemas) that broke v0.46.0. Needs no Harness credentials and no network access.
#
# Usage: ./scripts/smoke-test-provider.sh
#   TERRAFORM  terraform-compatible binary to use (default: terraform)

set -euo pipefail

TERRAFORM="${TERRAFORM:-terraform}"
REPO_ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
WORK_DIR="$(mktemp -d)"
trap 'rm -rf "${WORK_DIR}"' EXIT

command -v "${TERRAFORM}" >/dev/null || { echo "ERROR: ${TERRAFORM} not found on PATH" >&2; exit 1; }

echo "Building provider..."
mkdir -p "${WORK_DIR}/bin" "${WORK_DIR}/cfg"
(cd "${REPO_ROOT}" && CGO_ENABLED=0 go build -o "${WORK_DIR}/bin/terraform-provider-harness" .)

cat > "${WORK_DIR}/terraformrc" <<EOF
provider_installation {
  dev_overrides {
    "harness/harness" = "${WORK_DIR}/bin"
  }
  direct {}
}
EOF

cat > "${WORK_DIR}/cfg/main.tf" <<'EOF'
terraform {
  required_providers {
    harness = {
      source = "harness/harness"
    }
  }
}
EOF

# dev_overrides skips schema loading during init, so the schema is requested explicitly.
echo "Loading provider schema with ${TERRAFORM}..."
cd "${WORK_DIR}/cfg"
export TF_CLI_CONFIG_FILE="${WORK_DIR}/terraformrc"
schema="$("${TERRAFORM}" providers schema -json)"

# Spot-check one SDKv2 resource and the framework ephemeral resource, so a half-loaded mux fails.
for expected in '"harness_platform_gitops_agent"' '"harness_platform_gitops_agent_token"'; do
  if ! grep -q "${expected}" <<<"${schema}"; then
    echo "ERROR: ${expected} missing from provider schema" >&2
    exit 1
  fi
done

echo "OK: provider binary loads and exposes SDKv2 and ephemeral resources"
