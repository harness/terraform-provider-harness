#!/bin/bash

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;94m'
NC='\033[0m' # No Color

echo -e "${BLUE}=== Starting Harness Resource Setup ===${NC}\n"

# Function to check if a secret exists
check_secret() {
    local secret_name=$1
    terraform state list | grep -q "harness_platform_secret_text.$secret_name"
    return $?
}

# Function to check if a connector exists
check_connector() {
    local connector_name=$1
    terraform state list | grep -q "harness_platform_connector_github.$connector_name"
    return $?
}

# Initialize Terraform
echo -e "${YELLOW}Initializing Terraform...${NC}"
terraform init

# First phase: Check and create secrets
echo -e "\n${BLUE}=== Phase 1: Checking and Creating Secrets ===${NC}"

declare -a secrets=(
    "TF_spot_account_id"
    "TF_spot_api_token"
    "TF_spot_api_token_ref"
    "TF_Nexus_Password"
    "TF_git_bot_token"
    "TF_harness_automation_github_token"
)

for secret in "${secrets[@]}"; do
    echo -e "\n${YELLOW}Checking secret: $secret${NC}"
    if check_secret "$secret"; then
        echo -e "${GREEN}✓ Secret '$secret' already exists${NC}"
    else
        echo -e "${RED}Secret '$secret' not found. Creating...${NC}"
        terraform apply -auto-approve -target="harness_platform_secret_text.$secret"
        if [ $? -eq 0 ]; then
            echo -e "${GREEN}✓ Successfully created secret: $secret${NC}"
        else
            echo -e "${RED}✗ Failed to create secret: $secret${NC}"
        fi
    fi
done

# Second phase: Check and create connectors
echo -e "\n${BLUE}=== Phase 2: Checking and Creating Connectors ===${NC}"

declare -a connectors=(
    "TF_GitX_connector"
    "TF_open_repo_github_connector"
    "TF_Jajoo_github_connector"
    "TF_TerraformResource_git_connector"
    "TF_github_account_level_delegate_connector"
    "TF_github_account_level_connector"
)

for connector in "${connectors[@]}"; do
    echo -e "\n${YELLOW}Checking connector: $connector${NC}"
    if check_connector "$connector"; then
        echo -e "${GREEN}✓ Connector '$connector' already exists${NC}"
    else
        echo -e "${RED}Connector '$connector' not found. Creating...${NC}"
        terraform apply -auto-approve -target="harness_platform_connector_github.$connector"
        if [ $? -eq 0 ]; then
            echo -e "${GREEN}✓ Successfully created connector: $connector${NC}"
        else
            echo -e "${RED}✗ Failed to create connector: $connector${NC}"
        fi
    fi
done

echo -e "\n${GREEN}=== Setup Complete ===${NC}"
