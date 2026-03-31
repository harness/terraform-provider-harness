#!/bin/bash
# ============================================================================
# Harness Chaos Experiment Import Examples
# ============================================================================
#
# Import Format: org_id/project_id/experiment_identity
#
# ============================================================================

# ----------------------------------------------------------------------------
# Example 1: Import Project-Level Experiment
# ----------------------------------------------------------------------------
terraform import harness_chaos_experiment.basic \
  my_org/my_project/my-chaos-experiment

# ----------------------------------------------------------------------------
# After Import
# ----------------------------------------------------------------------------
# 1. Run `terraform state show harness_chaos_experiment.example` to see imported values
# 2. Run `terraform plan` to check for configuration drift
# 3. Update your .tf file to match the imported state
#
# Required configuration after import:
#   - org_id, project_id
#   - hub_org_id, hub_project_id, hub_identity
#   - template_identity
#   - name, infra_ref
#   - import_type (default: "REFERENCE")
# ----------------------------------------------------------------------------
