#!/usr/bin/env bash
# Verify whether QA autostopping API persists explicit opts.dry_run=false on PUT.
# Usage: HARNESS_PLATFORM_API_KEY=... ./scripts/verify_dry_run_api.sh
set -euo pipefail

ACCOUNT_ID="${HARNESS_ACCOUNT_ID:-JAvSzjfNQ6GPmm5PdhFWZQ}"
BASE="${HARNESS_ENDPOINT:-https://qa.harness.io/gateway}"
API_KEY="${HARNESS_PLATFORM_API_KEY:?Set HARNESS_PLATFORM_API_KEY}"

auth=(-H "X-Api-Key: ${API_KEY}" -H "Content-Type: application/json" -H "Accept: application/json")
qs="accountIdentifier=${ACCOUNT_ID}"

jq_dry_run() {
  jq -r '.response.service.opts.dry_run // .response.opts.dry_run // "null"'
}

rand_suffix() {
  date +%s | tail -c 6
}

NAME="api-dryrun-$(rand_suffix)"
PROXY_NAME="api-proxy-$(rand_suffix)"
VM_ID="/subscriptions/20d6a917-99fa-4b1b-9b2e-a3d624e9dcf0/resourceGroups/ccm-terraform-rg/providers/Microsoft.Compute/virtualMachines/DoNotDelete-Terraform-AS-Test-VM-3"

cleanup() {
  if [[ -n "${RULE_ID:-}" ]]; then
    curl -fsS -X DELETE "${auth[@]}" \
      "${BASE}/lw/api/accounts/${ACCOUNT_ID}/autostopping/rules/${RULE_ID}?${qs}" >/dev/null 2>&1 || true
  fi
  if [[ -n "${PROXY_ID:-}" ]]; then
    curl -fsS -X DELETE "${auth[@]}" \
      "${BASE}/lw/api/accounts/${ACCOUNT_ID}/autostopping/loadbalancers?${qs}" \
      -d "{\"ids\":[\"${PROXY_ID}\"],\"with_resources\":true}" >/dev/null 2>&1 || true
  fi
}
trap cleanup EXIT

echo "=== 1) Create Azure autostopping proxy ==="
PROXY_BODY=$(cat <<EOF
{
  "account_id": "${ACCOUNT_ID}",
  "cloud_account_id": "automation_azure_connector",
  "region": "eastus",
  "type": "azure",
  "name": "${PROXY_NAME}",
  "vpc": "/subscriptions/20d6a917-99fa-4b1b-9b2e-a3d624e9dcf0/resourceGroups/ccm-terraform-rg/providers/Microsoft.Network/virtualNetworks/DoNotDelete-Terraform-VNet",
  "kind": "autostopping_proxy",
  "metadata": {
    "security_groups": ["/subscriptions/20d6a917-99fa-4b1b-9b2e-a3d624e9dcf0/resourceGroups/ccm-terraform-rg/providers/Microsoft.Network/networkSecurityGroups/DoNotDelete-Terraform-NSG"],
    "resource_group": "ccm-terraform-rg",
    "subnet_id": "/subscriptions/20d6a917-99fa-4b1b-9b2e-a3d624e9dcf0/resourceGroups/ccm-terraform-rg/providers/Microsoft.Network/virtualNetworks/DoNotDelete-Terraform-VNet/subnets/default",
    "machine_type": "Standard_D2s_v3",
    "api-key": "${API_KEY}",
    "keypair": "DoNotDelete-Terraform-AS-Test-VM_key"
  }
}
EOF
)
PROXY_RESP=$(curl -fsS -X POST "${auth[@]}" \
  "${BASE}/lw/api/accounts/${ACCOUNT_ID}/autostopping/loadbalancers?${qs}" \
  -d "${PROXY_BODY}")
PROXY_ID=$(echo "${PROXY_RESP}" | jq -r '.response.id')
echo "proxy_id=${PROXY_ID}"

echo "=== 2) Wait for proxy created (up to 3m) ==="
for i in $(seq 1 36); do
  STATUS=$(curl -fsS "${auth[@]}" \
    "${BASE}/lw/api/accounts/${ACCOUNT_ID}/autostopping/loadbalancers/${PROXY_ID}?${qs}" | jq -r '.response.status')
  echo "  poll ${i}: status=${STATUS}"
  [[ "${STATUS}" == "created" ]] && break
  [[ "${STATUS}" == "errored" ]] && { echo "proxy errored"; exit 1; }
  sleep 5
done
[[ "${STATUS:-}" == "created" ]] || { echo "proxy not ready"; exit 1; }

echo "=== 3) Create VM rule with dry_run=true ==="
CREATE_BODY=$(cat <<EOF
{
  "service": {
    "name": "${NAME}",
    "account_identifier": "${ACCOUNT_ID}",
    "fulfilment": "ondemand",
    "kind": "instance",
    "cloud_account_id": "automation_azure_connector",
    "idle_time_mins": 10,
    "health_check": {"protocol": "http", "path": "/", "port": 80, "timeout": 30, "status_code_from": 200, "status_code_to": 299},
    "routing": {
      "instance": {"filter": {"ids": ["${VM_ID}"], "regions": ["eastus"]}},
      "http": {
        "proxy": {"id": "${PROXY_ID}"},
        "ports": [
          {"protocol": "http", "target_protocol": "http", "port": 80, "target_port": 80, "action": "forward", "routing_rules": [{}]}
        ]
      }
    },
    "opts": {"dry_run": true}
  }
}
EOF
)
CREATE_RESP=$(curl -fsS -X POST "${auth[@]}" \
  "${BASE}/lw/api/accounts/${ACCOUNT_ID}/autostopping/v2/rules?${qs}" \
  -d "${CREATE_BODY}")
RULE_ID=$(echo "${CREATE_RESP}" | jq -r '.response.id')
CREATE_DRY=$(echo "${CREATE_RESP}" | jq -r '.response.opts.dry_run')
echo "rule_id=${RULE_ID} create_dry_run=${CREATE_DRY}"

echo "=== 4) GET after create ==="
GET1=$(curl -fsS "${auth[@]}" \
  "${BASE}/lw/api/accounts/${ACCOUNT_ID}/autostopping/v2/rules/${RULE_ID}?${qs}")
GET1_DRY=$(echo "${GET1}" | jq_dry_run)
echo "get_after_create dry_run=${GET1_DRY}"

echo "=== 5) PUT with explicit dry_run=false ==="
PUT_EXPLICIT=$(cat <<EOF
{
  "service": {
    "name": "${NAME}",
    "account_identifier": "${ACCOUNT_ID}",
    "fulfilment": "ondemand",
    "kind": "instance",
    "cloud_account_id": "automation_azure_connector",
    "idle_time_mins": 10,
    "health_check": {"protocol": "http", "path": "/", "port": 80, "timeout": 30, "status_code_from": 200, "status_code_to": 299},
    "routing": {
      "instance": {"filter": {"ids": ["${VM_ID}"], "regions": ["eastus"]}},
      "http": {
        "proxy": {"id": "${PROXY_ID}"},
        "ports": [
          {"protocol": "http", "target_protocol": "http", "port": 80, "target_port": 80, "action": "forward", "routing_rules": [{}]}
        ]
      }
    },
    "opts": {"dry_run": false}
  }
}
EOF
)
PUT1_RESP=$(curl -fsS -X PUT "${auth[@]}" \
  "${BASE}/lw/api/accounts/${ACCOUNT_ID}/autostopping/v2/rules/${RULE_ID}?${qs}" \
  -d "${PUT_EXPLICIT}")
PUT1_DRY=$(echo "${PUT1_RESP}" | jq -r '.response.opts.dry_run')
echo "put_explicit_response dry_run=${PUT1_DRY}"

sleep 2
echo "=== 6) GET after explicit false PUT ==="
GET2=$(curl -fsS "${auth[@]}" \
  "${BASE}/lw/api/accounts/${ACCOUNT_ID}/autostopping/v2/rules/${RULE_ID}?${qs}")
GET2_DRY=$(echo "${GET2}" | jq_dry_run)
GET2_STATUS=$(echo "${GET2}" | jq -r '.response.service.status')
echo "get_after_explicit_false dry_run=${GET2_DRY} status=${GET2_STATUS}"

echo "=== 7) PUT back to dry_run=true, then PUT with empty opts {} (control) ==="
PUT_TRUE=$(echo "${PUT_EXPLICIT}" | jq '.service.opts = {"dry_run": true}')
curl -fsS -X PUT "${auth[@]}" \
  "${BASE}/lw/api/accounts/${ACCOUNT_ID}/autostopping/v2/rules/${RULE_ID}?${qs}" \
  -d "${PUT_TRUE}" >/dev/null
GET3_DRY=$(curl -fsS "${auth[@]}" \
  "${BASE}/lw/api/accounts/${ACCOUNT_ID}/autostopping/v2/rules/${RULE_ID}?${qs}" | jq_dry_run)
echo "get_after_reset_true dry_run=${GET3_DRY}"

PUT_EMPTY_OPTS=$(echo "${PUT_EXPLICIT}" | jq '.service.opts = {}')
PUT2_RESP=$(curl -fsS -X PUT "${auth[@]}" \
  "${BASE}/lw/api/accounts/${ACCOUNT_ID}/autostopping/v2/rules/${RULE_ID}?${qs}" \
  -d "${PUT_EMPTY_OPTS}")
PUT2_DRY=$(echo "${PUT2_RESP}" | jq -r '.response.opts.dry_run')
echo "put_empty_opts_response dry_run=${PUT2_DRY}"

sleep 2
GET4_DRY=$(curl -fsS "${auth[@]}" \
  "${BASE}/lw/api/accounts/${ACCOUNT_ID}/autostopping/v2/rules/${RULE_ID}?${qs}" | jq_dry_run)
echo "get_after_empty_opts dry_run=${GET4_DRY}"

echo ""
echo "=== SUMMARY ==="
echo "explicit_false_put_persisted=$([[ "${GET2_DRY}" == "false" ]] && echo YES || echo NO)"
echo "empty_opts_left_unchanged=$([[ "${GET4_DRY}" == "true" ]] && echo YES || echo NO)"
