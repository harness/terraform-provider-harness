# 0.5.0 (September 20,2022)

FEATURES:

* **New Resource:** `platform_environment_group` ([#203](https://github.com/harness/terraform-provider-harness/issues/203))

BUG FIXES:

* data-source/harness_platform_secret_text: Fixed value type field documentation in secret text ([#202](https://github.com/harness/terraform-provider-harness/issues/202))
* resource/harness_platform_secret_text: Fixed value type field documentation in secret text ([#202](https://github.com/harness/terraform-provider-harness/issues/202))

# 0.4.2 (September 14,2022)

BUG FIXES:

* data-source/harness_platform_usergroup: Fix user group to include sso related fields ([#199](https://github.com/harness/terraform-provider-harness/issues/199))
* resource/harness_platform_usergroup: Fix user group to include sso related fields ([#199](https://github.com/harness/terraform-provider-harness/issues/199))

# 0.4.1 (August 31,2022)

ENHANCEMENTS:

* data-source: Added example usage for data sources ([#193](https://github.com/harness/terraform-provider-harness/issues/193))

# 0.4.0 (August 31,2022)

FEATURES:

* **New Resource:** `platform_triggers` ([#192](https://github.com/harness/terraform-provider-harness/issues/192))

# 0.3.4 (August 27,2022)

BUG FIXES:

* resource/harness_platform_connector_artifactory: Fix bug in artifactory connector resource ([#191](https://github.com/harness/terraform-provider-harness/issues/191))

# 0.3.3 (August 23,2022)

BUG FIXES:

* resource/harness_platform_resource_group: Fix resource group empty attribute filter bug ([#182](https://github.com/harness/terraform-provider-harness/issues/182))

# 0.3.2 (August 14,2022)

BUG FIXES:

* data-source: Added First Gen and Next Gen sub category for data source. ([#184](https://github.com/harness/terraform-provider-harness/issues/184))
* resource: Added First Gen and Next Gen sub category for resource. ([#184](https://github.com/harness/terraform-provider-harness/issues/184))

# 0.3.1 (August 2,2022)

BUG FIXES:

* data-source/harness_platform_service_account: Added nextgen sub category for service account resource. ([#177](https://github.com/harness/terraform-provider-harness/issues/177))
* resource/harness_platform_service_account: Added nextgen sub category for service account resource. ([#177](https://github.com/harness/terraform-provider-harness/issues/177))

# 0.3.0 (August 2,2022)

FEATURES:

* **New Resource:** `platform_input_set` ([#174](https://github.com/harness/terraform-provider-harness/issues/174))
* **New Resource:** `platform_service_account` ([#170](https://github.com/harness/terraform-provider-harness/issues/170))

# 0.2.13

FEATURES:

* **New Resource:** `platform_resource_group` ([#168](https://github.com/harness/terraform-provider-harness/issues/168))

BUG FIXES:

* data-source/harness_platform_connector_kubernetes: Add delegate selectors. ([#169](https://github.com/harness/terraform-provider-harness/issues/169))
* resource/harness_platform_connector_kubernetes: Add delegate selectors. ([#169](https://github.com/harness/terraform-provider-harness/issues/169))
* resource/harness_platform_project: Fixed the vanishing color in project resource ([#166](https://github.com/harness/terraform-provider-harness/issues/166))

# 0.2.12

FEATURES:

* **New Resource:** `platform_roles` ([#161](https://github.com/harness/terraform-provider-harness/issues/161))
* **New Resource:** `platform_secret_file` ([#157](https://github.com/harness/terraform-provider-harness/issues/157))
* **New Resource:** `platform_secret_sshkey` ([#159](https://github.com/harness/terraform-provider-harness/issues/159))
* **New Resource:** `platform_secret_text` ([#154](https://github.com/harness/terraform-provider-harness/issues/154))

# 0.2.11

# 0.2.10

ENHANCEMENTS:

* resource/application: Added support for tagging. ([#140](https://github.com/harness/terraform-provider-harness/issues/140))

# 0.2.8

ENHANCEMENTS:

* Upgraded harnes-go-sdk@v0.2.27 ([#139](https://github.com/harness/terraform-provider-harness/issues/139))

BUG FIXES:

* datasource/harness_platform_connector: Fixed lookup by name. ([#139](https://github.com/harness/terraform-provider-harness/issues/139))
* datasource/harness_platform_organization: Fixed lookup by name. ([#139](https://github.com/harness/terraform-provider-harness/issues/139))
* datasource/harness_platform_pipeline: Fixed lookup by name. ([#139](https://github.com/harness/terraform-provider-harness/issues/139))
* datasource/harness_platform_project: Fixed lookup by name. ([#139](https://github.com/harness/terraform-provider-harness/issues/139))

# 0.2.7

BUG FIXES:

* Fixed issue with serializing encrypted text references with environment variable overrides. ([#131](https://github.com/harness/terraform-provider-harness/issues/131))

# 0.2.6

BUG FIXES:

* Fixed issue with serializing encrypted text references with service variables. ([#128](https://github.com/harness/terraform-provider-harness/issues/128))

# 0.2.4

BUG FIXES:

* resource/cloudprovider_kubernetes: Fixes issue causing delegate selectors to not be applied properly. ([#123](https://github.com/harness/terraform-provider-harness/issues/123))

# 0.2.3

FEATURES:

* **New Data Source:** `current_account` ([#119](https://github.com/harness/terraform-provider-harness/issues/119))

ENHANCEMENTS:

* data-source/yaml_config: Changing `path` field forces a new resource to be created. ([#117](https://github.com/harness/terraform-provider-harness/issues/117))

BUG FIXES:

* resource/delegate_approval: Force new resource when `delegate_id` or `approve` fields change. ([#115](https://github.com/harness/terraform-provider-harness/issues/115))

# 0.2.2

FEATURES:

* **New Data Source:** `platform_connector_appdynamics` ([#112](https://github.com/harness/terraform-provider-harness/issues/112))
* **New Data Source:** `platform_connector_artifactory` ([#112](https://github.com/harness/terraform-provider-harness/issues/112))
* **New Data Source:** `platform_connector_aws` ([#112](https://github.com/harness/terraform-provider-harness/issues/112))
* **New Data Source:** `platform_connector_aws_secret_m` ([#112](https://github.com/harness/terraform-provider-harness/issues/112))
* **New Data Source:** `platform_connector_awscc` ([#112](https://github.com/harness/terraform-provider-harness/issues/112))
* **New Data Source:** `platform_connector_awskms` ([#112](https://github.com/harness/terraform-provider-harness/issues/112))
* **New Data Source:** `platform_connector_bitbucket` ([#112](https://github.com/harness/terraform-provider-harness/issues/112))
* **New Data Source:** `platform_connector_datadog` ([#112](https://github.com/harness/terraform-provider-harness/issues/112))
* **New Data Source:** `platform_connector_docker` ([#112](https://github.com/harness/terraform-provider-harness/issues/112))
* **New Data Source:** `platform_connector_dynatrace` ([#112](https://github.com/harness/terraform-provider-harness/issues/112))
* **New Data Source:** `platform_connector_gcp` ([#112](https://github.com/harness/terraform-provider-harness/issues/112))
* **New Data Source:** `platform_connector_git` ([#112](https://github.com/harness/terraform-provider-harness/issues/112))
* **New Data Source:** `platform_connector_github` ([#112](https://github.com/harness/terraform-provider-harness/issues/112))
* **New Data Source:** `platform_connector_gitlab` ([#112](https://github.com/harness/terraform-provider-harness/issues/112))
* **New Data Source:** `platform_connector_helm` ([#112](https://github.com/harness/terraform-provider-harness/issues/112))
* **New Data Source:** `platform_connector_jira` ([#112](https://github.com/harness/terraform-provider-harness/issues/112))
* **New Data Source:** `platform_connector_kubernetes` ([#112](https://github.com/harness/terraform-provider-harness/issues/112))
* **New Data Source:** `platform_connector_nexus` ([#112](https://github.com/harness/terraform-provider-harness/issues/112))
* **New Data Source:** `platform_connector_pagerduty` ([#112](https://github.com/harness/terraform-provider-harness/issues/112))
* **New Data Source:** `platform_connector_prometheus` ([#112](https://github.com/harness/terraform-provider-harness/issues/112))
* **New Data Source:** `platform_connector_splunk` ([#112](https://github.com/harness/terraform-provider-harness/issues/112))
* **New Data Source:** `platform_connector_sumologic` ([#112](https://github.com/harness/terraform-provider-harness/issues/112))
* **New Data Source:** `platform_current_user` ([#112](https://github.com/harness/terraform-provider-harness/issues/112))
* **New Data Source:** `platform_environment` ([#112](https://github.com/harness/terraform-provider-harness/issues/112))
* **New Data Source:** `platform_organization` ([#112](https://github.com/harness/terraform-provider-harness/issues/112))
* **New Data Source:** `platform_pipeline` ([#112](https://github.com/harness/terraform-provider-harness/issues/112))
* **New Data Source:** `platform_project` ([#112](https://github.com/harness/terraform-provider-harness/issues/112))
* **New Data Source:** `platform_service` ([#112](https://github.com/harness/terraform-provider-harness/issues/112))
* **New Resource:** `platform_connector_appdynamics` ([#112](https://github.com/harness/terraform-provider-harness/issues/112))
* **New Resource:** `platform_connector_artifactory` ([#112](https://github.com/harness/terraform-provider-harness/issues/112))
* **New Resource:** `platform_connector_aws` ([#112](https://github.com/harness/terraform-provider-harness/issues/112))
* **New Resource:** `platform_connector_aws_secret_manager` ([#112](https://github.com/harness/terraform-provider-harness/issues/112))
* **New Resource:** `platform_connector_awscc` ([#112](https://github.com/harness/terraform-provider-harness/issues/112))
* **New Resource:** `platform_connector_awskms` ([#112](https://github.com/harness/terraform-provider-harness/issues/112))
* **New Resource:** `platform_connector_bitbucket` ([#112](https://github.com/harness/terraform-provider-harness/issues/112))
* **New Resource:** `platform_connector_datadog` ([#112](https://github.com/harness/terraform-provider-harness/issues/112))
* **New Resource:** `platform_connector_docker` ([#112](https://github.com/harness/terraform-provider-harness/issues/112))
* **New Resource:** `platform_connector_dynatrace` ([#112](https://github.com/harness/terraform-provider-harness/issues/112))
* **New Resource:** `platform_connector_gcp` ([#112](https://github.com/harness/terraform-provider-harness/issues/112))
* **New Resource:** `platform_connector_git` ([#112](https://github.com/harness/terraform-provider-harness/issues/112))
* **New Resource:** `platform_connector_github` ([#112](https://github.com/harness/terraform-provider-harness/issues/112))
* **New Resource:** `platform_connector_gitlab` ([#112](https://github.com/harness/terraform-provider-harness/issues/112))
* **New Resource:** `platform_connector_helm` ([#112](https://github.com/harness/terraform-provider-harness/issues/112))
* **New Resource:** `platform_connector_jira` ([#112](https://github.com/harness/terraform-provider-harness/issues/112))
* **New Resource:** `platform_connector_kubernetes` ([#112](https://github.com/harness/terraform-provider-harness/issues/112))
* **New Resource:** `platform_connector_nexus` ([#112](https://github.com/harness/terraform-provider-harness/issues/112))
* **New Resource:** `platform_connector_pagerduty` ([#112](https://github.com/harness/terraform-provider-harness/issues/112))
* **New Resource:** `platform_connector_prometheus` ([#112](https://github.com/harness/terraform-provider-harness/issues/112))
* **New Resource:** `platform_connector_splunk` ([#112](https://github.com/harness/terraform-provider-harness/issues/112))
* **New Resource:** `platform_connector_sumologic` ([#112](https://github.com/harness/terraform-provider-harness/issues/112))
* **New Resource:** `platform_current_user` ([#112](https://github.com/harness/terraform-provider-harness/issues/112))
* **New Resource:** `platform_environment` ([#112](https://github.com/harness/terraform-provider-harness/issues/112))
* **New Resource:** `platform_organization` ([#112](https://github.com/harness/terraform-provider-harness/issues/112))
* **New Resource:** `platform_pipeline` ([#112](https://github.com/harness/terraform-provider-harness/issues/112))
* **New Resource:** `platform_project` ([#112](https://github.com/harness/terraform-provider-harness/issues/112))
* **New Resource:** `platform_service` ([#112](https://github.com/harness/terraform-provider-harness/issues/112))

# 0.2.1

BUG FIXES:

* Fixes issue deleted account-level yaml config resources ([#106](https://github.com/harness/terraform-provider-harness/issues/106))

# 0.2.0

BREAKING CHANGES:

* Separate Nextgen codebase and Current Gen codebase ([#98](https://github.com/harness/terraform-provider-harness/issues/98))

ENHANCEMENTS:

* Upgraded to golang 1.17 ([#98](https://github.com/harness/terraform-provider-harness/issues/98))

# 0.1.16

# 0.1.15

BUG FIXES:

* Fixes issue with config-as-code secret references. ([#89](https://github.com/harness/terraform-provider-harness/issues/89))

# 0.1.14

ENHANCEMENTS:

* data-source/harness_environment: Replaces `id` field with `environment_id` so `id` field can be marked as computed. ([#81](https://github.com/harness/terraform-provider-harness/issues/81))

# 0.1.13

FEATURES:

* **New Resource:** `user_group_permissions` ([#80](https://github.com/harness/terraform-provider-harness/issues/80))

ENHANCEMENTS:

* data-source/harness_current_user: Change `2fa_enabled` to `is_two_factor_auth_enabled` to support `cdk` usage. ([#75](https://github.com/harness/terraform-provider-harness/issues/75))

BUG FIXES:

* Added configuration for auto generating the changelog ([#78](https://github.com/harness/terraform-provider-harness/issues/78))
* Fixed missing nextgen auth configuration in the provider. ([#76](https://github.com/harness/terraform-provider-harness/issues/76))

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
