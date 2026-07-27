#!/usr/bin/env bash
# CCM-32488-class nil-read panic tests (local).
# Non-secret QA defaults match CCM_Pipeline_Terraform_Test (environment=QA).
set -euo pipefail

ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
cd "$ROOT"

if [[ -f "$ROOT/.env" ]]; then
  set -a
  # shellcheck disable=SC1091
  source "$ROOT/.env"
  set +a
fi

export HARNESS_ENDPOINT="${HARNESS_ENDPOINT:-https://qa.harness.io/gateway}"
export HARNESS_ACCOUNT_ID="${HARNESS_ACCOUNT_ID:-JAvSzjfNQ6GPmm5PdhFWZQ}"
export AWS_ACCOUNT_ID="${AWS_ACCOUNT_ID:-132359207506}"
export HARNESS_GCP_PROJECT_ID="${HARNESS_GCP_PROJECT_ID:-durable-circle-282815}"
export TF_LOG="${TF_LOG:-}"

echo "=== CCM-32488-class unit tests (no live API required) ==="
PKGS=(
  ./internal/service/platform/connector/
  ./internal/service/platform/ccm_filters/
  ./internal/service/platform/cluster_orchestrator/
  ./internal/service/platform/autostopping/schedule/
)
FAILED=0
for pkg in "${PKGS[@]}"; do
  echo "--- $pkg ---"
  if ! go test "$pkg" -run TestCCM32488Class -v -count=1; then
    FAILED=1
  fi
done

if [[ -z "${HARNESS_PLATFORM_API_KEY:-}" ]]; then
  echo ""
  echo "HARNESS_PLATFORM_API_KEY not set — skipping live acceptance probes."
  echo "Export it from your other session (or add to terraform-provider-harness/.env) to run:"
  echo "  go test ./internal/service/platform/connector/ -run TestAccResourceConnectorAzureCloudCost -count=1 -v"
  exit "$FAILED"
fi

echo ""
echo "=== Live QA probe: Azure cloud cost connector (standard acceptance) ==="
go test ./internal/service/platform/connector/ \
  -run 'TestAccResourceConnectorAzureCloudCost$' -count=1 -v -timeout 30m

exit "$FAILED"
