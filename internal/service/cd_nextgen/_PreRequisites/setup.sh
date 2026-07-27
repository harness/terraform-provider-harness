#!/bin/bash

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;94m'
NC='\033[0m' # No Color

echo -e "${BLUE}=== Starting Harness Resource Setup ===${NC}\n"

TF_VARS=(-var="github_token_value=${github_token_value}" -var="harness_automation_github_token=${harness_automation_github_token}")

# Try to import a resource into the ephemeral state for this run.
# If the resource doesn't exist in Harness yet, the import fails and apply will create it.
try_import() {
    local resource_address=$1
    local import_id=$2
    echo -e "${YELLOW}  Importing $resource_address...${NC}"
    import_output=$(terraform import "${TF_VARS[@]}" "$resource_address" "$import_id" 2>&1)
    if [ $? -eq 0 ]; then
        echo -e "${GREEN}  ✓ Imported (will update)${NC}"
    else
        if echo "$import_output" | grep -qi "not found\|does not exist\|404\|cannot be found"; then
            echo -e "${YELLOW}  Not found in Harness (will create)${NC}"
        else
            echo -e "${RED}  Import error: $import_output${NC}"
        fi
    fi
}

apply_target() {
    local resource_address=$1
    terraform apply -auto-approve "${TF_VARS[@]}" -target="$resource_address"
    if [ $? -eq 0 ]; then
        echo -e "${GREEN}✓ $resource_address is up to date${NC}"
    else
        echo -e "${RED}✗ Failed: $resource_address${NC}"
    fi
}

# Initialize Terraform
echo -e "${YELLOW}Initializing Terraform...${NC}"
terraform init

# ── Secrets ──────────────────────────────────────────────────────────────────
echo -e "\n${BLUE}=== Phase 1: Secrets ===${NC}"

declare -a secrets=(
    "TF_spot_account_id"
    "TF_spot_api_token"
    "TF_spot_api_token_ref"
    "TF_Nexus_Password"
    "TF_git_bot_token"
    "TF_harness_automation_github_token"
)

for secret in "${secrets[@]}"; do
    echo -e "\n${YELLOW}Secret: $secret${NC}"
    try_import "harness_platform_secret_text.$secret" "$secret"
    apply_target "harness_platform_secret_text.$secret"
done

# ── Connectors ───────────────────────────────────────────────────────────────
echo -e "\n${BLUE}=== Phase 2: Connectors ===${NC}"

declare -a connectors=(
    "TF_GitX_connector"
    "TF_open_repo_github_connector"
    "TF_Jajoo_github_connector"
    "TF_TerraformResource_git_connector"
    "TF_github_account_level_delegate_connector"
    "TF_github_account_level_connector"
)

for connector in "${connectors[@]}"; do
    echo -e "\n${YELLOW}Connector: $connector${NC}"
    try_import "harness_platform_connector_github.$connector" "$connector"
    apply_target "harness_platform_connector_github.$connector"
done

echo -e "\n${GREEN}=== Setup Complete ===${NC}"
