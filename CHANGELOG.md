# 0.2.2 (Unreleased)

FEATURES:

* **New Resource:** `platform_connector_appdynamics
platform_connector_artifactory
platform_connector_aws_secret_manager
platform_connector_aws
platform_connector_awscc
platform_connector_awskms
platform_connector_bitbucket
platform_connector_datadog
platform_connector_docker
platform_connector_dynatrace
platform_connector_gcp
platform_connector_git
platform_connector_github
platform_connector_gitlab
platform_connector_helm
platform_connector_jira
platform_connector_kubernetes
platform_connector_nexus
platform_connector_pagerduty
platform_connector_prometheus
platform_connector_splunk
platform_connector_sumologic
platform_current_user
platform_environment
platform_organization
platform_pipeline
platform_project
platform_service` ([#112](https://github.com/hashicorp/terraform-provider-harness/issues/112))

# 0.2.1

BUG FIXES:

* Fixes issue deleted account-level yaml config resources ([#106](https://github.com/hashicorp/terraform-provider-harness/issues/106))

# 0.2.0

BREAKING CHANGES:

* Separate Nextgen codebase and Current Gen codebase ([#98](https://github.com/hashicorp/terraform-provider-harness/issues/98))

ENHANCEMENTS:

* Upgraded to golang 1.17 ([#98](https://github.com/hashicorp/terraform-provider-harness/issues/98))

# 0.1.16

# 0.1.15

BUG FIXES:

* Fixes issue with config-as-code secret references. ([#89](https://github.com/hashicorp/terraform-provider-harness/issues/89))

# 0.1.14

ENHANCEMENTS:

* data-source/harness_environment: Replaces `id` field with `environment_id` so `id` field can be marked as computed. ([#81](https://github.com/hashicorp/terraform-provider-harness/issues/81))

# 0.1.13

FEATURES:

* **New Resource:** `user_group_permissions` ([#80](https://github.com/hashicorp/terraform-provider-aws/issues/80))

ENHANCEMENTS:

* data-source/harness_current_user: Change `2fa_enabled` to `is_two_factor_auth_enabled` to support `cdk` usage. ([#75](https://github.com/hashicorp/terraform-provider-harness/issues/75))

BUG FIXES:

* Added configuration for auto generating the changelog ([#78](https://github.com/hashicorp/terraform-provider-harness/issues/78))
* Fixed missing nextgen auth configuration in the provider. ([#76](https://github.com/hashicorp/terraform-provider-harness/issues/76))

# 0.1.12

BUG FIXES:

* Fixed missing credentials in provider setup.

# 0.1.11

BUG FIXES:

* Upgraded harness-go-sdk to v0.1.11 to fix authentication configuration bug [71](https://github.com/harness/terraform-provider-harness/issues/71)

## 0.1.10

BREAKING CHANGES:

* Anyone using either on-prem or compliance clusters (anything other than `app.harness.io`) will need to update their provider endpoint configuration. The `/api` should be dropped from the configuration since this is now implied. For example, if you were previously setting this:

```terraform
provider "harness" {
  endpoint = "https://my.domain.com/api"
}
```

It should now be:
```terraform
provider "harness" {
  endpoint = "https://my.domain.com"
}
```

ENHANCEMENTS:

* Upgraded harness-go-sdk to v0.1.10
* Richer debug logging support added
* Refactored and simplified client configuration setup

## 0.1.9

ENHANCEMENTS:

* Upgraded harness-go-sdk to v0.
* **New Resource:** `harness_delegate_ids`
* data-source/delegate: Add support for looking up delegates by `hostname`

BUG FIXES:

* resource/harness_environment: Variable override field `service_name` is now optional. This allows a variable override to apply to all services when being deployed to an environment.
* Fixes delegate not found panic [#64](https://github.com/harness/terraform-provider-harness/issues/64)
