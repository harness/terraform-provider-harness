# v0.38.5 (August 22, 2025)
Release Page : https://github.com/harness/terraform-provider-harness/releases/tag/v0.38.5

# v0.36.2 (March 24, 2025)
Reverted the Go version upgrade.

# v0.36.0 (March 24, 2025)
This version is not recommended for use with the Terraform provider. Please use a different version.

# 0.35.8 (March 07,2025) 

# 0.35.7 (March 07,2025) 

ENHANCEMENTS:

* data-source/harness_platform_gitops_agent_operator_yaml: Added support for retrieving GitOps Agent Operator Yaml ([#1185](https://github.com/harness/terraform-provider-harness/issues/1185))

# 0.35.6 (February 21,2025) 

# 0.35.5 (February 17,2025) 

FEATURES:

* **New Resource:** `resource/harness_platform_infra_variable_set add iacm variable set resource to be used by multiple workspaces` ([#1173](https://github.com/harness/terraform-provider-harness/issues/1173))

ENHANCEMENTS:

* harness_platform_connector_jdbc: adding new ServiceAccount auth support ([#1160](https://github.com/harness/terraform-provider-harness/issues/1160))

# 0.35.4 (February 11,2025) 

# 0.35.3 (February 03,2025) 

# 0.35.2 (January 23,2025) 

# 0.35.1 (January 22,2025) 

# 0.35.0 (January 20,2025) 

# 0.34.9 (January 13,2025) 

# 0.34.8 (January 08,2025) 

# 0.34.7 (January 07,2025) 

# 0.34.6 (December 26,2024) 

# 0.34.5 (December 09,2024) 

# 0.34.4 (December 06,2024) 

FEATURES:

* **New Resource:** `harness_platform_infra_module - added a new resource for iacm module registry
  harness_platform_infra_module - added a new data source for iacm module registry` ([#1099](https://github.com/harness/terraform-provider-harness/issues/1099))

# 0.34.3 (November 08,2024) 

ENHANCEMENTS:

* harness_platform_connector_awscc: Added GOVERNANCE as a feature that can be enabled for the AWS Cloud Cost connector.
harness_platform_connector_azure_cloud_cost: Added GOVERNANCE as a feature that can be enabled for the Azure Cloud Cost connector.
harness_platform_connector_gcp_cloud_cost: Added GOVERNANCE as a feature that can be enabled for the GCP Cloud Cost connector. ([#1091](https://github.com/harness/terraform-provider-harness/issues/1091))

# 0.34.2 (November 05,2024) 

ENHANCEMENTS:

* resource/harness_platform_gitops_applications: add support for multi source applications
resource/harness_platform_gitops_repository: add force delete repo that skips validation against apps using repo ([#1082](https://github.com/harness/terraform-provider-harness/issues/1082))

# 0.34.1 (October 11,2024) 

# 0.34.0 (October 09,2024) 

FEATURES:

* **New Resource:** `harness_cluster_orchestrator - Added Cluster Orchestrator resource in Harness terraform provider` ([#1084](https://github.com/harness/terraform-provider-harness/issues/1084))
* **New Resource:** `harness_governance_rule_set - Added Governance Rule Set resource in Harness terraform provider` ([#1075](https://github.com/harness/terraform-provider-harness/issues/1075))

# 0.33.4 (October 01,2024) 

FEATURES:

* **New Resource:** `harness_governance_rule - Added Governance Rule resource in Harness terraform provider` ([#1068](https://github.com/harness/terraform-provider-harness/issues/1068))

# 0.33.3 (September 13,2024) 

# 0.33.2 (September 12,2024) 

# 0.33.1 (September 11,2024) 

# 0.33.0 (September 11,2024) 

FEATURES:

* **New Resource:** `harness_governance_rule_enforcement - Added Governance Enforcement resource in Harness terraform provider` ([#1052](https://github.com/harness/terraform-provider-harness/issues/1052))

# 0.32.9 (September 10,2024) 

# 0.32.9 (September 10,2024) 

# 0.32.8 (September 05,2024) 

# 0.32.7 (September 03,2024) 

# 0.32.6 (August 29,2024) 

# 0.32.5 (August 27,2024) 

# 0.32.4 (August 27,2024) 

# 0.32.3 (August 09,2024) 

FEATURES:

* **New Resource:** `harness_platform_gitx_webhook: Added support for Gitx Webhook endpoints` ([#1014](https://github.com/harness/terraform-provider-harness/issues/1014))

ENHANCEMENTS:

* harness_platform_policy: - Have added support for remote policies, now users can create remote policies using terraform by providing git references in the script.
                         - Updated the Data Source for policy as rego should not be a mandatory field for GET call. ([#1015](https://github.com/harness/terraform-provider-harness/issues/1015))
* resource/harness_platform_connector_gcp: added support for OIDC Auth in GCP connector. ([#1019](https://github.com/harness/terraform-provider-harness/issues/1019))

# 0.32.2 (July 25,2024) 

ENHANCEMENTS:

* harness_platform_gitops_app_project: Rename the file name ([#1011](https://github.com/harness/terraform-provider-harness/issues/1011))
* harness_platform_gitops_project: Added the resource version gitops project ([#1010](https://github.com/harness/terraform-provider-harness/issues/1010))

# 0.32.1 (July 23,2024) 

ENHANCEMENTS:

* harness_platform_gitops_project: Update the schema of gitops project ([#1009](https://github.com/harness/terraform-provider-harness/issues/1009))
* harness_platform_pipeline: Added update call for git metadata in pipeline update flow.
harness_platform_template:Added update call for git metadata in template update flow. ([#1008](https://github.com/harness/terraform-provider-harness/issues/1008))

# 0.31.9 (July 18,2024) 

FEATURES:

* **New Resource:** `harness_platform_notification_rule: Added support for notification rule for SLO and MonitoredService` ([#1013](https://github.com/harness/terraform-provider-harness/issues/1013))

ENHANCEMENTS:

* resource_monitored_service: Add support for custom change sources in monitored service configuration. ([#1007](https://github.com/harness/terraform-provider-harness/issues/1007))

# 0.31.8 (July 11,2024) 

FEATURES:

* **New Resource:** `harness_platform_connector_jdbc` ([#1002](https://github.com/harness/terraform-provider-harness/issues/1002))
* **New Resource:** `harness_platform_db_instance` ([#1001](https://github.com/harness/terraform-provider-harness/issues/1001))
* **New Resource:** `harness_platform_db_schema` ([#998](https://github.com/harness/terraform-provider-harness/issues/998))
* **New Resource:** `harness_platform_gitops_project: Added the terraform support for creating the gitops project` ([#1000](https://github.com/harness/terraform-provider-harness/issues/1000))
* **New Resource:** `harness_platform_pipeline: Added the branch context when imported the pipeline from non main branch` ([#999](https://github.com/harness/terraform-provider-harness/issues/999))

ENHANCEMENTS:

* harness_platform_connector_gcp_cloud_cost: Added support for governance in feature_enabled field in gcp and azure cloud cost connectors
harness_platform_connector_azure_cloud_cost: Added support for governance in feature_enabled field in gcp and azure cloud cost connectors ([#1006](https://github.com/harness/terraform-provider-harness/issues/1006))
* platform_gitops_applications: Added support for skip repo validation, when set true repo doesnt have to be created before creating application. ([#979](https://github.com/harness/terraform-provider-harness/issues/979))

# 0.31.7 (July 02,2024) 

# 0.31.6 (July 01,2024) 

# 0.31.5 (June 17,2024) 

BUG FIXES:

* resource/platform_repo: Fixes issue for repo import giving 404 even after being successful. ([#978](https://github.com/harness/terraform-provider-harness/issues/978))

# 0.31.4 (June 07,2024) 

# 0.31.3 (May 31,2024) 

# 0.31.2 (May 21,2024) 

# 0.31.1 (May 10,2024) 

# 0.31.0 (May 10,2024) 

# 0.31.0 (May 10,2024) 

ENHANCEMENTS:

* resource/harness_pipeline: Updated the description for tags in pipeline create ([#977](https://github.com/harness/terraform-provider-harness/issues/977))

# 0.30.9 (May 03,2024) 

# 0.30.8 (April 04,2024) 

ENHANCEMENTS:

* resource/harness_platform_variable: updated schema to forceNew for identifier, orgId and projectId ([#963](https://github.com/harness/terraform-provider-harness/issues/963))

# 0.30.7 (March 22,2024) 

ENHANCEMENTS:

* resource/harness_platform_user: added support for update for user groups for user entity ([#957](https://github.com/harness/terraform-provider-harness/issues/957))

# 0.30.6 (March 18,2024) 

ENHANCEMENTS:

* harness_platform_infrastructures: Added supprt for creating/Updating remote Infrastructures and import from git for Infrastructures. ([#931](https://github.com/harness/terraform-provider-harness/issues/931))

BUG FIXES:

* fix resource and data source for environment to gitops cluster mapping ([#954](https://github.com/harness/terraform-provider-harness/issues/954))

# 0.30.5 (March 07,2024) 

FEATURES:

* **New Resource:** `platform_repo` ([#879](https://github.com/harness/terraform-provider-harness/issues/879))
* **New Resource:** `platform_repo_rule_branch` ([#879](https://github.com/harness/terraform-provider-harness/issues/879))

ENHANCEMENTS:

* harness_platform_connector_kubernetes: Added support for ca_cert_ref in kubernetes connector cluster. ([#928](https://github.com/harness/terraform-provider-harness/issues/928))
* harness_platform_environment: Added supprt for creating/Updating remote Environment and import from git for Environment. ([#929](https://github.com/harness/terraform-provider-harness/issues/929))
* harness_platform_service: Added supprt for creating/Updating remote services and import from git for service. ([#927](https://github.com/harness/terraform-provider-harness/issues/927))
* harness_platform_service_overrides_v2: Added supprt for creating/Updating remote Overrides and import from git for Overrides. ([#930](https://github.com/harness/terraform-provider-harness/issues/930))

# 0.30.4 (February 28,2024) 

ENHANCEMENTS:

* harness_platform_gitops_applications: The Path field was required. It has been updated to optional. ([#926](https://github.com/harness/terraform-provider-harness/issues/926))
* resource/harness_platform_connector_aws: added support for OIDC Auth in AWS connector. ([#925](https://github.com/harness/terraform-provider-harness/issues/925))

# 0.30.3 (February 27,2024) 

ENHANCEMENTS:

* platform_connector_github.md - Added documentation for github http anonymous connector
resource.tf - Example of github http anonymous resource for github connector
go.mod - upgraded harness-go-sdk version to v0.3.74
github.go - Added Schema and CRUD for github anonymous http credentials in github connector
github_test.go - Test for schema and CRUD for github anonymous http credentials in github connector
github_data_source.go - Added Schema for github anonymous http credentials in github connector
github_data_source_test.go - Added  test for Schema of github anonymous http  credentials in github connector ([#886](https://github.com/harness/terraform-provider-harness/issues/886))
* resource/harness_platform_project: where project and org is required, new resources are created on updating scope. ([#902](https://github.com/harness/terraform-provider-harness/issues/902))

# 0.30.2 (January 23,2024) 

ENHANCEMENTS:

* resource/harness_platform_project: corrected error handling for EntityNotFoundException ([#872](https://github.com/harness/terraform-provider-harness/issues/872))

# 0.30.1 (January 19,2024) 

ENHANCEMENTS:

* resource/harness_platform_connector_github: added force deletion support for github connector ([#855](https://github.com/harness/terraform-provider-harness/issues/855))
* resource/harness_platform_project: added error handling for EntityNotFoundException ([#858](https://github.com/harness/terraform-provider-harness/issues/858))

# 0.30.0 (January 10,2024) 

FEATURES:

* **New Data Source:** `harness_platform_current_account -  new data resource for account id` ([#785](https://github.com/harness/terraform-provider-harness/issues/785))
* **New Resource:** `resource/harness_platform_connector_pdc: Added Pdc connector resource.` ([#832](https://github.com/harness/terraform-provider-harness/issues/832))

# 0.29.4 (December 12,2023) 

BUG FIXES:

* resource/monitored_service - Added missing examples for monitored service with thresholds and NewRelic health source. ([#814](https://github.com/harness/terraform-provider-harness/issues/814))

# 0.29.3 (December 11,2023) 

ENHANCEMENTS:

* Allow for tags to be included in feature flags ([#781](https://github.com/harness/terraform-provider-harness/issues/781))

BUG FIXES:

* resource/harness_platform_connector_aws - Added support to add aws region field in connector to perform connection test. ([#806](https://github.com/harness/terraform-provider-harness/issues/806))
* resource/harness_platform_template: Fixing the update flow for templates ([#780](https://github.com/harness/terraform-provider-harness/issues/780))

# 0.29.2 (November 28,2023) 

# 0.29.1 (November 28,2023) 

BUG FIXES:

* Monitored service didn't use to honour metric threshold correctly earlier with this fix customer can add metric threshold to all the health sources ([#777](https://github.com/harness/terraform-provider-harness/issues/777))
* harness_platform_ff_api_key -  Fix error type ([#763](https://github.com/harness/terraform-provider-harness/issues/763))

# 0.29.0 (November 16,2023) 

FEATURES:

* **New Resource:** `harness_platform_workspace - added a new resource for iacm workspaces
  harness_platform_workspace - added a new data source for iacm workspaces
  harness_platform_workspace_output - added a new data source for iacm workspace outputs` ([#726](https://github.com/harness/terraform-provider-harness/issues/726))
* **New Resource:** `resource/harness_platform_delegatetoken: Added delegate token resource.` ([#719](https://github.com/harness/terraform-provider-harness/issues/719))

ENHANCEMENTS:

* Allow for multiple environments, with each environment containing its own target group and rules ([#742](https://github.com/harness/terraform-provider-harness/issues/742))
* Days attribute of AutoStopping fixed schedule has been modified to be list of weekdays instead of string for brevity ([#731](https://github.com/harness/terraform-provider-harness/issues/731))
* harness_platform_service_overrides_v2 -  Updated documentation with details of how overrides v2 identifiers are generated. ([#730](https://github.com/harness/terraform-provider-harness/issues/730))
* resource/platform_connector_aws: added force deletion support
resource/platform_connector_azure_cloud_provider: added force deletion support
resources/platform_connector_gcp: added force deletion support
resources/platform_connector_helm: added force deletion support
resources/platform_connector_kubernetes: added force deletion support
resources/platform_connector_oci_helm: added force deletion support
resources/platform_connector_rancher: added force deletion support ([#743](https://github.com/harness/terraform-provider-harness/issues/743))

# 0.28.3 (October 25,2023) 

FEATURES:

* **New Resource:** `Fixed schedule support for AutoStopping` ([#714](https://github.com/harness/terraform-provider-harness/issues/714))

ENHANCEMENTS:

* - Support for import of ALB as AutoStopping Loadbalancer
    - Support for import of AppGateway as AutoStopping Loadbalancer ([#707](https://github.com/harness/terraform-provider-harness/issues/707))
* harness_platform_service_overrides_v2 -  Updated docs to use yaml format for the yaml property of overrides instead of json format. ([#711](https://github.com/harness/terraform-provider-harness/issues/711))

BUG FIXES:

* harness_platform_gitops_app_project_mapping -  Fix field name of argo project. ([#709](https://github.com/harness/terraform-provider-harness/issues/709))

# 0.28.2 (October 09,2023) 

ENHANCEMENTS:

* Added is_namespaced for agent creating ([#704](https://github.com/harness/terraform-provider-harness/issues/704))
* platform_connector_github.md - Added documentation for custom health connector
resource.tf - Example of github_app in credentials for github connector
go.mod - upgraded harness-go-sdk version to v0.3.52
github.go - Added Schema and CRUD for github_app credentials in github connector
github_test.go - Test for schema and CRUD for github_app credentials in github connector
github_data_source.go - Added Schema for github_app credentials in github connector
github_data_source_test.go - Added  test for Schema of github_app credentials in github connector ([#703](https://github.com/harness/terraform-provider-harness/issues/703))
* resource/harness_platform_gitops_agent: add support for new flux operator. ([#674](https://github.com/harness/terraform-provider-harness/issues/674))

# 0.28.1 (September 27,2023) 

ENHANCEMENTS:

* resource/harness_platform_connector_aws_secret_manager: add support for new field "default" which sets the secret manager to be default for the account ([#633](https://github.com/harness/terraform-provider-harness/issues/633))

BUG FIXES:

* harness_platform_feature_flag -  Enable adding target and target groups through features ([#699](https://github.com/harness/terraform-provider-harness/issues/699))
* harness_platform_feature_flag -  Fix updates on targets, target groups and feature flags. ([#700](https://github.com/harness/terraform-provider-harness/issues/700))

# 0.28.0 (September 18,2023) 

FEATURES:

* **New Resource:** `- Support for AutoStopping LoadBalancers - AWS proxy, AWS ALB, Azure gateway, Azure proxy, GCP proxy` ([#698](https://github.com/harness/terraform-provider-harness/issues/698))
* **New Resource:** `- Support for AutoStopping Rules - AWS VM, Azure VM, GCP VM, AWS RDS, AWS ECS
 - This version does not support dry run AutoStopping rule creation, provision for hide progress page and tag based ECS rule creation
 - Does not have support for schedules feature` ([#634](https://github.com/harness/terraform-provider-harness/issues/634))

ENHANCEMENTS:

* resource/harness_platform_connector_service_now: GA FF CDS_SERVICENOW_REFRESH_TOKEN_AUTH. ([#691](https://github.com/harness/terraform-provider-harness/issues/691))

BUG FIXES:

* harness_platform_feature_flag_target_group -  Fix type issue when casting between types for include and exclude rules ([#694](https://github.com/harness/terraform-provider-harness/issues/694))

# 0.27.2 (September 14,2023) 

BUG FIXES:

* harness_platform_connector_aws -  Set value for cross_account_access if not nil. ([#687](https://github.com/harness/terraform-provider-harness/issues/687))

# 0.27.1 (September 13,2023) 

ENHANCEMENTS:

* harness_platform_feature_flag - Add support to add targets with feature flags ([#684](https://github.com/harness/terraform-provider-harness/issues/684))

# 0.27.0 (September 13,2023) 

ENHANCEMENTS:

* harness_platform_template - Added support to import Template entity from git.
harness_platform_input_set - Added support to import InputSet entity from git. ([#682](https://github.com/harness/terraform-provider-harness/issues/682))
* harness_platform_template - Added support to update stable version of a template via terraform ([#675](https://github.com/harness/terraform-provider-harness/issues/675))
* platform_connector_customhealthsource.md Added documentation for custom health connector
resource.tf Example of loki health source with custom health source connector added
provider.go : Added new data source harness_platform_connector_customhealthsource
custom_health_data_source.go Schema and CRUD of custom health source connector
custom_health_data_source_test.go Test for Schema and CRUD of custom health source connector
customhealthsource.go Schema and CRUD of custom health source connector
customhealthsource_test.go Test for Schema and CRUD of custom health source connector ([#677](https://github.com/harness/terraform-provider-harness/issues/677))
* resource/harness_platform_service_overrides_v2: Added Support For Service Overrides V2 ([#673](https://github.com/harness/terraform-provider-harness/issues/673))

BUG FIXES:

* - Removed type support from the service dependencies since we only support inside the dependencyMetadata. ([#672](https://github.com/harness/terraform-provider-harness/issues/672))
* Fixed aws-cc to allow non-billing connector type ([#679](https://github.com/harness/terraform-provider-harness/issues/679))
* harness_platform_file_store_file -  Make file content optional, if the file content is provided use it directly else get the content from file path. ([#681](https://github.com/harness/terraform-provider-harness/issues/681))

# 0.26.0 (September 01,2023) 

FEATURES:

* **New Resource:** `resource_feature_flag_target_group - Added feature flag target group resources to the Harness Terraform Provider.` ([#661](https://github.com/harness/terraform-provider-harness/issues/661))

# 0.25.0 (September 01,2023) 

FEATURES:

* **New Resource:** `harness_platform_gitops_app_project_mapping - GitOps app project mapping for agents resource.` ([#659](https://github.com/harness/terraform-provider-harness/issues/659))
* **New Resource:** `resource_feature_flag_target - Added feature flag target resources to the Harness Terraform Provider.` ([#660](https://github.com/harness/terraform-provider-harness/issues/660))

# 0.24.5 (August 29,2023) 

BUG FIXES:

* Fixed policy-set api to correctly enable/disable policy-sets ([#670](https://github.com/harness/terraform-provider-harness/issues/670))

# 0.24.4 (August 29,2023) 

ENHANCEMENTS:

* data_source_monitored_service_test.go Added tests for multiple healthsources such as Prometheus, Datadog etc.
resource_monitored_service.go Added version field and renamed MonitoredServiceSpec to MonitoredService
resource_monitored_service_test.go renamed MonitoredServiceSpec to MonitoredService
utils.go Deserializer updated with new health sources such as azure, signalFx, loki and sumologic
platform_monitored_service.md Added docs for health sources such as azure, signalFx, loki and sumologic
resource.tf Added examples for all newly added health sources, datadog and prometheus ([#669](https://github.com/harness/terraform-provider-harness/issues/669))
* harness_platform_pipeline - Added support to import pipeline entity from git. ([#643](https://github.com/harness/terraform-provider-harness/issues/643))
* resource/harness_plaform_user: Limit the user creation call to 1 at a time. ([#668](https://github.com/harness/terraform-provider-harness/issues/668))

BUG FIXES:

* Fixed harness_platform_file_store_folder create resource plugin crash, when service account token was used to create ([#665](https://github.com/harness/terraform-provider-harness/issues/665))

# 0.24.3 (August 22,2023) 

BUG FIXES:

* PolicySets list must return sorted ordered list of policysets. ([#658](https://github.com/harness/terraform-provider-harness/issues/658))
* UI update to service environment override resource is not reflected accurately if this resource is created using Terraform. ([#654](https://github.com/harness/terraform-provider-harness/issues/654))

# 0.24.2 (August 17,2023) 

ENHANCEMENTS:

* resource/harness_platform_gitops_agent: add support for getting agent token while creation. ([#653](https://github.com/harness/terraform-provider-harness/issues/653))

BUG FIXES:

* - Deprecated enabled from the monitored service dto to not allow customer to set monitored service as enabled to start with via terraform ([#640](https://github.com/harness/terraform-provider-harness/issues/640))
* Fixed the environment group resource to support org and account level environment groups also. ([#655](https://github.com/harness/terraform-provider-harness/issues/655))

# 0.24.1 (August 10,2023) 

ENHANCEMENTS:

* harness_platform_template - Added support to update stable version of a template from terraform. ([#628](https://github.com/harness/terraform-provider-harness/issues/628))

BUG FIXES:

* resource_manual_freeze : bug fix for deployment freeze while saving expired freeze ([#650](https://github.com/harness/terraform-provider-harness/issues/650))

# 0.24.0 (August 01,2023) 

FEATURES:

* **New Resource:** `harness_platform_file_store_file - Added a File Store file resource.
harness_platform_file_store_folder - Added a File Store folder resource.` ([#629](https://github.com/harness/terraform-provider-harness/issues/629))

ENHANCEMENTS:

* Add more fields for overrides v2 data source ([#636](https://github.com/harness/terraform-provider-harness/issues/636))
* resource/harness_platform_gitops_repository: add support for token update for OCI helm repo with ESO ([#638](https://github.com/harness/terraform-provider-harness/issues/638))

# 0.23.3 (July 25,2023) 

ENHANCEMENTS:

* Upgraded harnes-go-sdk@v0.3.39 ([#630](https://github.com/harness/terraform-provider-harness/issues/630))

# 0.23.2 (July 20,2023) 

ENHANCEMENTS:

* resource/harness_platform_connector_service_now: Enhanced servicenow connector resource to support newer RefreshTokenGrantType authentication beans. ([#614](https://github.com/harness/terraform-provider-harness/issues/614))

BUG FIXES:

* Fix dataSourceTokenRead documentation. ([#618](https://github.com/harness/terraform-provider-harness/issues/618))
* Fixed the example documentation for service overrides ([#623](https://github.com/harness/terraform-provider-harness/issues/623))
* resource/harness_platform_apikey: made tags field as set of strings ([#625](https://github.com/harness/terraform-provider-harness/issues/625))
* resource/harness_platform_token: made tags field as set of strings ([#624](https://github.com/harness/terraform-provider-harness/issues/624))

# 0.23.1 (July 14,2023) 

ENHANCEMENTS:

* resource/harness_platform_infrastructure: added AWS SAM infrastructure type. ([#616](https://github.com/harness/terraform-provider-harness/issues/616))

BUG FIXES:

* - Formatting the sample example for harness_platform_environment resource ([#604](https://github.com/harness/terraform-provider-harness/issues/604))
* Fix for supporting yaml for environment and environment group data source ([#610](https://github.com/harness/terraform-provider-harness/issues/610))

# 0.23.0 (July 05,2023) 

FEATURES:

* **New Resource:** `harness_platform_connector_rancher - Added a Rancher connector to the Harness Terraform provider.` ([#598](https://github.com/harness/terraform-provider-harness/issues/598))

ENHANCEMENTS:

* resource/harness_platform_connector_jira: Enhanced jira connector resource to support newer PersonalAccessToken authentication beans. ([#591](https://github.com/harness/terraform-provider-harness/issues/591))
* resource/harness_platform_gitops_cluster: Added support for IAM role in GitOps ([#578](https://github.com/harness/terraform-provider-harness/issues/578))
* resource/harness_platform_service_overrides_v2: Added Support For Service Overrides V2 ([#579](https://github.com/harness/terraform-provider-harness/issues/579))
* resource/harness_platform_template: Update the description for filter_visibility.
resource/harness_platform_template_filters: Update the description force_delete. ([#600](https://github.com/harness/terraform-provider-harness/issues/600))

BUG FIXES:

* resource/harness_platform_token: fix for returning token value in create API ([#605](https://github.com/harness/terraform-provider-harness/issues/605))

# 0.22.1 (June 15,2023) 

ENHANCEMENTS:

* resource/harness_platform_connector_aws - Added Aws BackOff Strategy Override Support ([#560](https://github.com/harness/terraform-provider-harness/issues/560))
* resource/harness_platform_policyset: correct description for 'severity' for policy.
resource/harness_platform_policy: Enhance example to showcase how to add policy with Rego spanning over multiple lines. ([#569](https://github.com/harness/terraform-provider-harness/issues/569))

BUG FIXES:

* - Removed empty change source and health source restriction from monitored service resource. ([#576](https://github.com/harness/terraform-provider-harness/issues/576))
* Fix for supporting import for account/org infrastructure. ([#577](https://github.com/harness/terraform-provider-harness/issues/577))
* resource/harness_platform_usergroup - ignore the order of users and user_emails when doing CRUD. ([#567](https://github.com/harness/terraform-provider-harness/issues/567))

# 0.22.0 (June 05,2023) 

FEATURES:

* **New Resource:** `platform_token - Added apikey token in Harness terraform provider` ([#556](https://github.com/harness/terraform-provider-harness/issues/556))

ENHANCEMENTS:

* examples/resources/harness_platform_template/ - updated the remote template example ([#566](https://github.com/harness/terraform-provider-harness/issues/566))

BUG FIXES:

* resource/harness_platform_template - deprecated the field description from resource ([#561](https://github.com/harness/terraform-provider-harness/issues/561))

# 0.21.0 (May 25,2023) 

FEATURES:

* **New Resource:** `platform_connector_elasticsearch - Added elasticsearch connector resource in Harness terraform provider` ([#538](https://github.com/harness/terraform-provider-harness/issues/538))

# 0.20.0 (May 16,2023) 

FEATURES:

* **New Resource:** `resource/harness_platform_apikey: Added ApiKey resource.` ([#532](https://github.com/harness/terraform-provider-harness/issues/532))

ENHANCEMENTS:

* resources/platform_role_assignments: Made resource_group_identifier, role_identifier and type under principal schema required.
resources/platform_secret_sshkey: Updated the Behaviour of referencing the secrets at account, project and org Level. Made Key Field Required in SSH credential of type keyReference ([#509](https://github.com/harness/terraform-provider-harness/issues/509))

# 0.19.2 (May 11,2023) 

FEATURES:

* **New Resource:** `platform_connector_tas - Added tas connector resource in Harness terraform provider` ([#523](https://github.com/harness/terraform-provider-harness/issues/523))

ENHANCEMENTS:

* data-source/platform_manual_freeze - Added quarterly recurrence support for manual deployment freeze resource in Harness terraform provider ([#522](https://github.com/harness/terraform-provider-harness/issues/522))
* resource/harness_platform_infrastructure: added force deletion support for infrastructures ([#527](https://github.com/harness/terraform-provider-harness/issues/527))
* resource/harness_platform_triggers: added documentation links ([#501](https://github.com/harness/terraform-provider-harness/issues/501))

# 0.19.1 (May 02,2023) 

ENHANCEMENTS:

* datasource: Make identifier required in connector data source. ([#526](https://github.com/harness/terraform-provider-harness/issues/526))
* platform_connector_terraform_cloud - Improved Documentation ([#520](https://github.com/harness/terraform-provider-harness/issues/520))

BUG FIXES:

* resource/harness_platform_connector_awscc: Fix bug in aws cloud cost connector resource ([#524](https://github.com/harness/terraform-provider-harness/issues/524))

# 0.19.0 (April 26,2023) 

ENHANCEMENTS:

* resource/harness_platform_template: added force deletion support for templates ([#518](https://github.com/harness/terraform-provider-harness/issues/518))

# 0.18.0 (April 25,2023) 

FEATURES:

* **New Resource:** `platform_feature_flag - Added feature flag resources to the Harness Terraform Provider.
platform_ff_api_key - Added FF SDK API key resources to the Harness Terraform provider.` ([#517](https://github.com/harness/terraform-provider-harness/issues/517))

# 0.17.5 (April 20,2023) 

BUG FIXES:

* data-source/harness_platform_infrastructure: Fix bug wrt usages of tags in infrastructure yaml. ([#515](https://github.com/harness/terraform-provider-harness/issues/515))

# 0.17.4 (April 20,2023) 

ENHANCEMENTS:

* resource/harness_platform_connector_github: Added support for secret ref in application and installtion in github app authentication method. ([#508](https://github.com/harness/terraform-provider-harness/issues/508))

# 0.17.3 (April 10,2023) 

BUG FIXES:

* data-source/harness_platform_organization: Fixed the data source to use either name or identifier.
data-source/harness_platform_usergroup: Fixed the data source to use either name or identifier. ([#507](https://github.com/harness/terraform-provider-harness/issues/507))

# 0.17.2 (April 06,2023) 

# 0.17.1 (April 05,2023) 

ENHANCEMENTS:

* resource/harness_platform_gitops_cluster: Added support for Optional Tags in cluster ([#486](https://github.com/harness/terraform-provider-harness/issues/486))
* resource/harness_platform_service: added force deletion support for services
resource/harness_platform_environment: added force deletion support for environments
resource/harness_platform_environment_group: added force deletion support for environment groups ([#491](https://github.com/harness/terraform-provider-harness/issues/491))

BUG FIXES:

* harness_platform_input_set: Fixed import.
harness_platform_triggers: Fixed import. ([#478](https://github.com/harness/terraform-provider-harness/issues/478))
* resource/harness_platform_connector_aws_secret_manager: Fixed the plugin crash issue when api key doent have enough permissions. ([#502](https://github.com/harness/terraform-provider-harness/issues/502))

# 0.17.0 (March 24,2023) 

ENHANCEMENTS:

* resource/harness_platform_policyset: Adding the policyset management provider ([#485](https://github.com/harness/terraform-provider-harness/issues/485))

# 0.16.4 (March 22,2023) 

ENHANCEMENTS:

* resource/harness_platform_environment_service_overrides: Support for organisation and account scoped service overrides ([#479](https://github.com/harness/terraform-provider-harness/issues/479))

# 0.16.3 (March 22,2023) 

ENHANCEMENTS:

* resources: Update yaml fields description. ([#484](https://github.com/harness/terraform-provider-harness/issues/484))

# 0.16.2 (March 21,2023) 

BUG FIXES:

* resource/harness_platform_role_assignments: Allow creation of role_assignments without Identifier and set that identifier coming from upStream when doing a get call. ([#477](https://github.com/harness/terraform-provider-harness/issues/477))
* resource/harness_platform_secret_text: The Value Field was Optional , It has been updated to Required Field. ([#472](https://github.com/harness/terraform-provider-harness/issues/472))

# 0.16.1 (March 15,2023) 

ENHANCEMENTS:

* resource/harness_platform_template: Updated connector ref field's description.
resource/harness_platform_pipeline: Updated connector ref field's description.
resource/harness_platform_connector_kubernetes_cloud_cost: Updated connector ref field's description.
resource/harness_platform_input_set: Updated connector ref field's description. ([#471](https://github.com/harness/terraform-provider-harness/issues/471))

# 0.16.0 (March 14,2023) 

FEATURES:

* **New Resource:** `resource/harness_platform_connector_oci_helm: Added Oci Helm connector resource.` ([#466](https://github.com/harness/terraform-provider-harness/issues/466))

# 0.15.0 (March 10,2023) 

FEATURES:

* **New Resource:** `resource/harness_platform_connector_service_now: Added Service Now connector resource.` ([#465](https://github.com/harness/terraform-provider-harness/issues/465))

# 0.14.15 (March 09,2023) 

BUG FIXES:

* harness_platform_secret_text: Mark new resource when secret deleted from ui . ([#461](https://github.com/harness/terraform-provider-harness/issues/461))
* harness_platform_usergroup: Mark new resource when usergroup deleted from ui. ([#462](https://github.com/harness/terraform-provider-harness/issues/462))

# 0.14.14 (March 09,2023) 

BUG FIXES:

* harness_platform_secret_text: Mark new resource when secret deleted from ui . ([#461](https://github.com/harness/terraform-provider-harness/issues/461))

# 0.14.13 (March 08,2023) 

BUG FIXES:

* harness_platform_template: Fixed import.
harness_platform_pipeline: Fixed import. ([#457](https://github.com/harness/terraform-provider-harness/issues/457))

# 0.14.12 (March 08,2023) 

BUG FIXES:

* resource/harness_platform_connector_jira: Fixed Jira Connector Resource to support newer UsernamePassword authentication beans. Users of Jira Connector need to update their Harness Terraform Provider to this version since it is a breaking change in the API. ([#456](https://github.com/harness/terraform-provider-harness/issues/456))

# 0.14.11 (March 06,2023) 

BUG FIXES:

* resource/harness_platform_connector_helm: Fixed documentation. ([#452](https://github.com/harness/terraform-provider-harness/issues/452))
* resource/harness_platform_organization: Fixed the plugin crash issue during terraform refresh when the api key was invalid. ([#454](https://github.com/harness/terraform-provider-harness/issues/454))

# 0.14.10 (March 01,2023) 

BUG FIXES:

* resource/harness_platform_user: Fixed Bug with user resource. ([#451](https://github.com/harness/terraform-provider-harness/issues/451))

# 0.14.9 (March 01,2023) 

ENHANCEMENTS:

* resource/harness_platform_infrastructure: added support for account and org level. ([#438](https://github.com/harness/terraform-provider-harness/issues/438))

BUG FIXES:

* resource/harness_platform_connector_gcp_secret_manager: Fixed GCP Secret Manager resource. ([#442](https://github.com/harness/terraform-provider-harness/issues/442))
* resource/harness_platform_user: Fixed Bug with user resource. ([#446](https://github.com/harness/terraform-provider-harness/issues/446))

# 0.14.5 (February 23,2023) 

ENHANCEMENTS:

* resource/harness_platform_service: added support for account and org level.
resource/harness_platform_environment: added support for account and org level. ([#432](https://github.com/harness/terraform-provider-harness/issues/432))

BUG FIXES:

* resource/harness_platform_secret_file: Fix secret file resource. ([#437](https://github.com/harness/terraform-provider-harness/issues/437))

# 0.14.4 (February 22,2023) 

BUG FIXES:

* resource/harness_platform_environment_service_overrides: Fix import flow ([#423](https://github.com/harness/terraform-provider-harness/issues/423))
* resource/harness_platform_monitored_service: Fields template_ref and version_label shouldn't be required for harness_platform_monitored_service. ([#430](https://github.com/harness/terraform-provider-harness/issues/430))

# 0.14.3 (February 13,2023) 

ENHANCEMENTS:

* resource/harness_user_group: Update filters field in workflo, enviroments ,pipeline object in user group to be optional. ([#422](https://github.com/harness/terraform-provider-harness/issues/422))

# 0.14.2 (February 07,2023) 

ENHANCEMENTS:

* resource/harness_user_group: Update filters field in deployment object in user group to be optional. ([#418](https://github.com/harness/terraform-provider-harness/issues/418))

BUG FIXES:

* resource/harness_platform_user: Fix user schema. ([#413](https://github.com/harness/terraform-provider-harness/issues/413))

# 0.14.1 (January 31,2023) 

ENHANCEMENTS:

* resource/harness_platform_template: In template api slug field is updated to identifier. Made relevent changes in terraform resource. ([#412](https://github.com/harness/terraform-provider-harness/issues/412))

BUG FIXES:

* resource/harness_platform_usergroup: Fix users and user_emails field in user group schema. ([#411](https://github.com/harness/terraform-provider-harness/issues/411))

# 0.14.0 (January 27,2023) 

FEATURES:

* **New Resource:** `platform_manual_freeze - Added manual deployment freeze resource in Harness terraform provider` ([#355](https://github.com/harness/terraform-provider-harness/issues/355))
* **New Resource:** `resource/harness_platform_user: Resource for creating a Harness User` ([#353](https://github.com/harness/terraform-provider-harness/issues/353))

# 0.13.3 (January 19,2023) 

BUG FIXES:

* resource/harness_platform_pipeline: Update terraform resource to reflect the backend changes in api.
resource/harness_platform_input_set: Update terraform resource to reflect the backend changes in api. ([#396](https://github.com/harness/terraform-provider-harness/issues/396))

# 0.13.2 (January 19,2023) 

# 0.13.1 (January 19,2023) 

# 0.13.0 (January 17,2023) 

FEATURES:

* **New Resource:** `platform_monitored_service - Added monitored service resources to the Harness Terraform Provider.
platform_slo - Added service-level objective (SLO) resources to the Harness Terraform provider.` ([#348](https://github.com/harness/terraform-provider-harness/issues/348))

# 0.12.4 (January 16,2023) 

ENHANCEMENTS:

* resource/harness_platform_environment_service_overrides: Mark new resource if resource is deleted during terraform refresh
resource/harness_platform_organization: Mark new resource if resource is deleted during terraform refresh
resource/harness_platform_role_assignments: Mark new resource if resource is deleted during terraform refresh
resource/harness_platform_variables: Mark new resource if resource is deleted during terraform refresh ([#387](https://github.com/harness/terraform-provider-harness/issues/387))
* resource/harness_platform_input_set: Added gitx support for inputSet resource.
data-source/harness_platform_input_set: Added gitx support for inputSet resource. ([#389](https://github.com/harness/terraform-provider-harness/issues/389))
* resource/harness_platform_service: Updating the documentation.
resource/harness_platform_environment: Updating the documentation. ([#378](https://github.com/harness/terraform-provider-harness/issues/378))

# 0.12.3 (January 04,2023) 

BUG FIXES:

* resource/harness_platform_connector_github: Fix connector delete context. ([#377](https://github.com/harness/terraform-provider-harness/issues/377))

# 0.12.2 (January 03,2023) 

ENHANCEMENTS:

* resource/harness_platform_usergroup: Add example to create user group by adding email. ([#373](https://github.com/harness/terraform-provider-harness/issues/373))

# 0.12.1 (January 03,2023) 

ENHANCEMENTS:

* resource/harness_platform_usergroup:  Allow TF resource to support creating of user-groups by adding User email id. ([#371](https://github.com/harness/terraform-provider-harness/issues/371))

# 0.12.0 (December 23,2022) 

FEATURES:

* **New Resource:** `platform_connector_jenkins - Added jenkins connector resource in Harness terraform provider.` ([#365](https://github.com/harness/terraform-provider-harness/issues/365))

# 0.11.5 (December 14,2022) 

BUG FIXES:

* resource/harness_platform_connector_github: Fix terraform refresh for github connector ([#352](https://github.com/harness/terraform-provider-harness/issues/352))

# 0.11.4 (December 13,2022) 

BUG FIXES:

* resource/harness_yaml_config: Fix yaml config resource ([#349](https://github.com/harness/terraform-provider-harness/issues/349))

# 0.11.3 (December 09,2022) 

FEATURES:

* **New Resource:** `resource/harness_platform_template_filters: Resource for creating a Harness template filters` ([#337](https://github.com/harness/terraform-provider-harness/issues/337))

# 0.11.2 (December 07,2022) 

BUG FIXES:

* resource/harness_platform_environment: Handle case when environment is deleted from somewhere else and refresh fails ([#343](https://github.com/harness/terraform-provider-harness/issues/343))

# 0.11.1 (December 06,2022) 

# 0.11.0 (December 02,2022) 

FEATURES:

* **New Resource:** `resource/harness_infrastructure_definition: Add new infrastructure type CUSTOM for CG` ([#335](https://github.com/harness/terraform-provider-harness/issues/335))

BUG FIXES:

* resource/harness_platform_project: Fix project refresh ([#333](https://github.com/harness/terraform-provider-harness/issues/333))

# 0.10.3 (December 02,2022) 

BUG FIXES:

* data-source/harness_platform_usergroup: Fix usergroup data-source to get account and org level usergroups ([#334](https://github.com/harness/terraform-provider-harness/issues/334))

# 0.10.2 (December 01,2022) 

BUG FIXES:

* resource/harness_platform_infrastructure: Fix infrastructure resource when creating multiple infrastructure in same env ([#330](https://github.com/harness/terraform-provider-harness/issues/330))

# 0.10.1 (November 30,2022) 

BUG FIXES:

* resource/harness_platform_template: Fix template resource ([#329](https://github.com/harness/terraform-provider-harness/issues/329))

# 0.10.0 (November 30,2022) 

FEATURES:

* **New Resource:** `platform_connector_spot - Added spot connector resource in Harness terraform provider` ([#314](https://github.com/harness/terraform-provider-harness/issues/314))
* **New Resource:** `platform_template - Added template resource in Harness terraform provider` ([#326](https://github.com/harness/terraform-provider-harness/issues/326))

ENHANCEMENTS:

* resource/harness_platform_policy: Adding the policy management provider ([#319](https://github.com/harness/terraform-provider-harness/issues/319))

# 0.9.1 (November 24,2022) 

BUG FIXES:

* resource/harness_platform_pipeline: Fix error propagation from api ([#318](https://github.com/harness/terraform-provider-harness/issues/318))

# 0.9.0 (November 24,2022) 

FEATURES:

* **New Resource:** `platform_connector_azure_key_vault - Added the Azure Key Vault connector resource to the Harness Terraform provider.` ([#287](https://github.com/harness/terraform-provider-harness/issues/287))

# 0.8.4 (November 23,2022) 

ENHANCEMENTS:

* data-source/harness_trigger: - Added data source for trigger in first gen ([#309](https://github.com/harness/terraform-provider-harness/issues/309))
* platform_user_group - Added support for "Send email to all users" in user group notification configuration. ([#313](https://github.com/harness/terraform-provider-harness/issues/313))

# 0.8.3 (November 14,2022)

FEATURES:

* **New Resource:** `platform_connector_azure_cloud_cost - Added azure cloud cost connector resource in Harness terraform provider` ([#284](https://github.com/harness/terraform-provider-harness/issues/284))
* **New Resource:** `platform_connector_azure_cloud_provider - Added gcp cloud cost connector resource in Harness terraform provider` ([#285](https://github.com/harness/terraform-provider-harness/issues/285))
* **New Resource:** `platform_connector_kubernetes_cloud_cost - Added kubernetes cloud cost connector resource in Harness terraform provider` ([#286](https://github.com/harness/terraform-provider-harness/issues/286))

ENHANCEMENTS:

* platform_pipeline - Added support for pipelines with new Git Experience. ([#294](https://github.com/harness/terraform-provider-harness/issues/294))

# 0.7.1 (November 4,2022)

FEATURES:

* **New Resource:** `harness_platform_gitops_agent 
harness_platform_gitops_cluster
harness_platform_gitops_applications 
harness_platform_gitops_repository 
harness_platform_gitops_repo_cert` ([#282](https://github.com/harness/terraform-provider-harness/issues/282))

BUG FIXES:

* harness_platform_secret_text : fix value field in secret text resource ([#281](https://github.com/harness/terraform-provider-harness/issues/281))

# 0.7.0 (November 2,2022)

FEATURES:

* **New Resource:** `harness_platform_gitops_agent_yaml` ([#253](https://github.com/harness/terraform-provider-harness/issues/253))
* **New Resource:** `harness_platform_gitops_applications` ([#253](https://github.com/harness/terraform-provider-harness/issues/253))
* **New Resource:** `platform_connector_azure_cloud_provider - Added azure cloud provider connector resource in Harness terraform provider` ([#274](https://github.com/harness/terraform-provider-harness/issues/274))
* **New Resource:** `platform_connector_gcp_secret_manager - Added gcp secret manager resource in Harness terraform provider` ([#254](https://github.com/harness/terraform-provider-harness/issues/254))
* **New Resource:** `platform_filters - Added resource to add filters in Harness through terraform` ([#255](https://github.com/harness/terraform-provider-harness/issues/255))

ENHANCEMENTS:

* harness_platform_gitops_agent ([#253](https://github.com/harness/terraform-provider-harness/issues/253))
* harness_platform_gitops_clusters ([#253](https://github.com/harness/terraform-provider-harness/issues/253))
* harness_platform_gitops_repository ([#253](https://github.com/harness/terraform-provider-harness/issues/253))

# 0.6.11 (November 2,2022)

# 0.6.10 (October 20,2022)

BUG FIXES:

* harness_platform_gitops_agent : Fix subcategory in documentation
harness_platform_gitops_cluster : Fix subcategory in documentation ([#242](https://github.com/harness/terraform-provider-harness/issues/242))

# 0.6.9 (October 20,2022)

FEATURES:

* **New Resource:** `harness_platform_gitops_agent` ([#214](https://github.com/harness/terraform-provider-harness/issues/214))
* **New Resource:** `harness_platform_gitops_cluster` ([#214](https://github.com/harness/terraform-provider-harness/issues/214))
* **New Resource:** `harness_platform_infrastructure` ([#239](https://github.com/harness/terraform-provider-harness/issues/239))

ENHANCEMENTS:

* resource/harness_platform_usergroup:  Added saml and ldap sso provider documentation ([#241](https://github.com/harness/terraform-provider-harness/issues/241))

# 0.6.8 (October 13,2022)

ENHANCEMENTS:

* resource/harness_platform_usergroup:  update doc and added more example ([#233](https://github.com/harness/terraform-provider-harness/issues/233))

BUG FIXES:

* resource/harness_platform_usergroup: Fix user group delete flow ([#238](https://github.com/harness/terraform-provider-harness/issues/238))

# 0.6.7 (October 11,2022)

ENHANCEMENTS:

* resource/harness_platform_environment_clusters_mapping:  update docs for cluster resource ([#232](https://github.com/harness/terraform-provider-harness/issues/232))
* resource/harness_platform_environment_clusters_mapping:  update name for cluster resource ([#231](https://github.com/harness/terraform-provider-harness/issues/231))

BUG FIXES:

* resource/harness_platform_environment: Fix Update tags for the environment ([#234](https://github.com/harness/terraform-provider-harness/issues/234))
* resource/harness_platform_environment_group: Fix bug in update env group ([#229](https://github.com/harness/terraform-provider-harness/issues/229))

# 0.6.4 (October 6,2022)

ENHANCEMENTS:

* resource/harness_platform_environment: added suport foroptional `yaml` field. ([#221](https://github.com/harness/terraform-provider-harness/issues/221))
* resource/harness_platform_service: making `yaml` field as optional. ([#218](https://github.com/harness/terraform-provider-harness/issues/218))

BUG FIXES:

* resource/harness_platform_roles: Fix bug in roles resource ([#222](https://github.com/harness/terraform-provider-harness/issues/222))

# 0.6.3 (October 5,2022)

BUG FIXES:

* resource/harness_platform_service_account: Fix account id field in service account ([#219](https://github.com/harness/terraform-provider-harness/issues/219))

# 0.6.2 (October 5,2022)

BUG FIXES:

* resource/harness_platform_service_account: Fix email field in service account ([#217](https://github.com/harness/terraform-provider-harness/issues/217))

# 0.6.1 (October 3,2022)

ENHANCEMENTS:

* resource/harness_platform_service: added documentation for yaml field. ([#215](https://github.com/harness/terraform-provider-harness/issues/215))

BUG FIXES:

* resource/harness_platform_environment_clusters_mapping: Fix added correct batch cluster API. ([#216](https://github.com/harness/terraform-provider-harness/issues/216))

# 0.6.0 (September 30,2022)

FEATURES:

* **New Resource:** `platform_role_assignments` ([#213](https://github.com/harness/terraform-provider-harness/issues/213))

# 0.5.3 (September 25,2022)

BUG FIXES:

* data-source/harness_platform_connector_prometheus: Fix prometheus connector schema to include missing fields ([#210](https://github.com/harness/terraform-provider-harness/issues/210))
* resource/harness_platform_connector_prometheus: Fix prometheus connector schema to include missing fields ([#210](https://github.com/harness/terraform-provider-harness/issues/210))
* resource/harness_platform_environment: Fixed bug in color field ([#209](https://github.com/harness/terraform-provider-harness/issues/209))

# 0.5.2 (September 21,2022)

# 0.5.1 (September 21,2022)

FEATURES:

* **New Resource:** `platform_cluster` ([#204](https://github.com/harness/terraform-provider-harness/issues/204))

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
