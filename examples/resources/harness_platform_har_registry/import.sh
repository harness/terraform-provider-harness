#!/bin/bash
# Import format: <space_ref>/<registry_identifier>
#   space_ref defines the scope: account, org, or project level
#   identifier is the unique registry identifier
#
# Account level: <account_id>/<registry_identifier>
# Org level: <account_id>/<org_id>/<registry_identifier>
# Project level: <account_id>/<org_id>/<project_id>/<registry_identifier>
terraform import harness_platform_har_registry.example <space_ref>/<registry_identifier>

# After import:
#   terraform plan
#   terraform apply
