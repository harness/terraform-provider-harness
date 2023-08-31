package provider

import (
	"context"
	"fmt"
	"log"

	"github.com/harness/terraform-provider-harness/internal/service/platform/feature_flag"
	"github.com/harness/terraform-provider-harness/internal/service/platform/feature_flag_target"
	"github.com/harness/terraform-provider-harness/internal/service/platform/ff_api_key"
	"github.com/harness/terraform-provider-harness/internal/service/platform/gitops/agent_yaml"
	"github.com/harness/terraform-provider-harness/internal/service/platform/manual_freeze"
	"github.com/harness/terraform-provider-harness/internal/service/platform/policy"
	"github.com/harness/terraform-provider-harness/internal/service/platform/policyset"
	"github.com/sirupsen/logrus"

	"github.com/harness/harness-go-sdk/harness"
	"github.com/harness/harness-go-sdk/harness/cd"
	"github.com/harness/harness-go-sdk/harness/helpers"
	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/harness-go-sdk/harness/utils"
	openapi_client_nextgen "github.com/harness/harness-openapi-go-client/nextgen"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/harness/terraform-provider-harness/internal/service/cd/account"
	"github.com/harness/terraform-provider-harness/internal/service/cd/application"
	"github.com/harness/terraform-provider-harness/internal/service/cd/cloudprovider"
	cd_connector "github.com/harness/terraform-provider-harness/internal/service/cd/connector"
	"github.com/harness/terraform-provider-harness/internal/service/cd/delegate"
	"github.com/harness/terraform-provider-harness/internal/service/cd/environment"
	"github.com/harness/terraform-provider-harness/internal/service/cd/secrets"
	"github.com/harness/terraform-provider-harness/internal/service/cd/service"
	"github.com/harness/terraform-provider-harness/internal/service/cd/sso"
	cd_trigger "github.com/harness/terraform-provider-harness/internal/service/cd/trigger"
	"github.com/harness/terraform-provider-harness/internal/service/cd/user"
	"github.com/harness/terraform-provider-harness/internal/service/cd/yamlconfig"
	pl_apikey "github.com/harness/terraform-provider-harness/internal/service/platform/api_key"
	"github.com/harness/terraform-provider-harness/internal/service/platform/ccm_filters"
	"github.com/harness/terraform-provider-harness/internal/service/platform/connector"
	pl_current_user "github.com/harness/terraform-provider-harness/internal/service/platform/current_user"
	pl_delegatetoken "github.com/harness/terraform-provider-harness/internal/service/platform/delegate_token"
	pl_environment "github.com/harness/terraform-provider-harness/internal/service/platform/environment"
	pl_environment_clusters_mapping "github.com/harness/terraform-provider-harness/internal/service/platform/environment_clusters_mapping"
	pl_environment_group "github.com/harness/terraform-provider-harness/internal/service/platform/environment_group"
	pl_environment_service_overrides "github.com/harness/terraform-provider-harness/internal/service/platform/environment_service_overrides"
	file_store "github.com/harness/terraform-provider-harness/internal/service/platform/file_store"
	"github.com/harness/terraform-provider-harness/internal/service/platform/filters"
	gitops_agent "github.com/harness/terraform-provider-harness/internal/service/platform/gitops/agent"
	gitops_project_mapping "github.com/harness/terraform-provider-harness/internal/service/platform/gitops/app_project"
	gitops_applications "github.com/harness/terraform-provider-harness/internal/service/platform/gitops/applications"
	gitops_cluster "github.com/harness/terraform-provider-harness/internal/service/platform/gitops/cluster"
	gitops_gnupg "github.com/harness/terraform-provider-harness/internal/service/platform/gitops/gnupg"
	gitops_repository "github.com/harness/terraform-provider-harness/internal/service/platform/gitops/repository"
	gitops_repo_cert "github.com/harness/terraform-provider-harness/internal/service/platform/gitops/repository_certificates"
	gitops_repo_cred "github.com/harness/terraform-provider-harness/internal/service/platform/gitops/repository_credentials"
	pl_infrastructure "github.com/harness/terraform-provider-harness/internal/service/platform/infrastructure"
	"github.com/harness/terraform-provider-harness/internal/service/platform/input_set"
	"github.com/harness/terraform-provider-harness/internal/service/platform/monitored_service"
	"github.com/harness/terraform-provider-harness/internal/service/platform/organization"
	pl_permissions "github.com/harness/terraform-provider-harness/internal/service/platform/permissions"
	"github.com/harness/terraform-provider-harness/internal/service/platform/pipeline"
	"github.com/harness/terraform-provider-harness/internal/service/platform/pipeline_filters"
	"github.com/harness/terraform-provider-harness/internal/service/platform/project"
	"github.com/harness/terraform-provider-harness/internal/service/platform/resource_group"
	"github.com/harness/terraform-provider-harness/internal/service/platform/role_assignments"
	"github.com/harness/terraform-provider-harness/internal/service/platform/roles"
	"github.com/harness/terraform-provider-harness/internal/service/platform/secret"
	pl_service "github.com/harness/terraform-provider-harness/internal/service/platform/service"
	"github.com/harness/terraform-provider-harness/internal/service/platform/service_account"
	pl_service_overrides_v2 "github.com/harness/terraform-provider-harness/internal/service/platform/service_overrides_v2"
	"github.com/harness/terraform-provider-harness/internal/service/platform/slo"
	pl_template "github.com/harness/terraform-provider-harness/internal/service/platform/template"
	"github.com/harness/terraform-provider-harness/internal/service/platform/template_filters"
	pl_token "github.com/harness/terraform-provider-harness/internal/service/platform/token"
	"github.com/harness/terraform-provider-harness/internal/service/platform/triggers"
	pl_user "github.com/harness/terraform-provider-harness/internal/service/platform/user"
	"github.com/harness/terraform-provider-harness/internal/service/platform/usergroup"
	"github.com/harness/terraform-provider-harness/internal/service/platform/variables"

	"github.com/harness/harness-go-sdk/logging"
	openapi_client_logging "github.com/harness/harness-openapi-go-client/logging"
	"github.com/hashicorp/go-cleanhttp"
	"github.com/hashicorp/go-retryablehttp"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func init() {
	// Set descriptions to support markdown syntax, this will be used in document generation
	// and the language server.
	schema.DescriptionKind = schema.StringMarkdown

	// Customize the content of descriptions when output. For example you can add defaults on
	// to the exported descriptions if present.
	// schema.SchemaDescriptionBuilder = func(s *schema.Schema) string {
	// 	desc := s.Description
	// 	if s.Default != nil {
	// 		desc += fmt.Sprintf(" Defaults to `%v`.", s.Default)
	// 	}
	// 	return strings.TrimSpace(desc)
	// }
}

func Provider(version string) func() *schema.Provider {
	return func() *schema.Provider {
		p := &schema.Provider{
			Schema: map[string]*schema.Schema{
				"endpoint": {
					Description: fmt.Sprintf("The URL of the Harness API endpoint. The default is `https://app.harness.io/gateway`. This can also be set using the `%s` environment variable.", helpers.EnvVars.Endpoint.String()),
					Type:        schema.TypeString,
					Required:    true,
					DefaultFunc: schema.EnvDefaultFunc(helpers.EnvVars.Endpoint.String(), utils.BaseUrl),
				},
				"account_id": {
					Description: fmt.Sprintf("The Harness account id. This can also be set using the `%s` environment variable.", helpers.EnvVars.AccountId.String()),
					Type:        schema.TypeString,
					Required:    true,
					DefaultFunc: schema.EnvDefaultFunc(helpers.EnvVars.AccountId.String(), nil),
				},
				"api_key": {
					Description: fmt.Sprintf("The Harness API key. This can also be set using the `%s` environment variable. For more information to create an API key in FirstGen, see https://docs.harness.io/article/smloyragsm-api-keys#create_an_api_key.", helpers.EnvVars.ApiKey.String()),
					Type:        schema.TypeString,
					Optional:    true,
					DefaultFunc: schema.EnvDefaultFunc(helpers.EnvVars.ApiKey.String(), nil),
				},
				"platform_api_key": {
					Description: fmt.Sprintf("The API key for the Harness next gen platform. This can also be set using the `%s` environment variable. For more information to create an API key in NextGen, see https://docs.harness.io/article/tdoad7xrh9-add-and-manage-api-keys.", helpers.EnvVars.PlatformApiKey.String()),
					Type:        schema.TypeString,
					Optional:    true,
					DefaultFunc: schema.EnvDefaultFunc(helpers.EnvVars.PlatformApiKey.String(), nil),
				},
			},
			DataSourcesMap: map[string]*schema.Resource{
				"harness_platform_template":                        pl_template.DataSourceTemplate(),
				"harness_platform_connector_azure_key_vault":       connector.DataSourceConnectorAzureKeyVault(),
				"harness_platform_connector_gcp_cloud_cost":        connector.DataSourceConnectorGCPCloudCost(),
				"harness_platform_connector_kubernetes_cloud_cost": connector.DatasourceConnectorKubernetesCloudCost(),
				"harness_platform_connector_azure_cloud_cost":      connector.DataSourceConnectorAzureCloudCost(),
				"harness_platform_connector_appdynamics":           connector.DatasourceConnectorAppDynamics(),
				"harness_platform_connector_elasticsearch":         connector.DatasourceConnectorElasticSearch(),
				"harness_platform_connector_artifactory":           connector.DatasourceConnectorArtifactory(),
				"harness_platform_connector_aws_secret_manager":    connector.DatasourceConnectorAwsSM(),
				"harness_platform_connector_aws":                   connector.DatasourceConnectorAws(),
				"harness_platform_connector_awscc":                 connector.DatasourceConnectorAwsCC(),
				"harness_platform_connector_awskms":                connector.DatasourceConnectorAwsKms(),
				"harness_platform_connector_bitbucket":             connector.DatasourceConnectorBitbucket(),
				"harness_platform_connector_datadog":               connector.DatasourceConnectorDatadog(),
				"harness_platform_connector_docker":                connector.DatasourceConnectorDocker(),
				"harness_platform_connector_dynatrace":             connector.DatasourceConnectorDynatrace(),
				"harness_platform_connector_gcp":                   connector.DatasourceConnectorGcp(),
				"harness_platform_connector_gcp_secret_manager":    connector.DatasourceConnectorGcpSM(),
				"harness_platform_connector_git":                   connector.DatasourceConnectorGit(),
				"harness_platform_connector_github":                connector.DatasourceConnectorGithub(),
				"harness_platform_connector_gitlab":                connector.DatasourceConnectorGitlab(),
				"harness_platform_connector_helm":                  connector.DatasourceConnectorHelm(),
				"harness_platform_connector_oci_helm":              connector.DatasourceConnectorOciHelm(),
				"harness_platform_connector_jira":                  connector.DatasourceConnectorJira(),
				"harness_platform_connector_jenkins":               connector.DataSourceConnectorJenkins(),
				"harness_platform_connector_kubernetes":            connector.DatasourceConnectorKubernetes(),
				"harness_platform_connector_nexus":                 connector.DatasourceConnectorNexus(),
				"harness_platform_connector_pagerduty":             connector.DatasourceConnectorPagerDuty(),
				"harness_platform_connector_prometheus":            connector.DatasourceConnectorPrometheus(),
				"harness_platform_connector_rancher":               connector.DatasourceConnectorRancher(),
				"harness_platform_connector_splunk":                connector.DatasourceConnectorSplunk(),
				"harness_platform_connector_spot":                  connector.DatasourceConnectorSpot(),
				"harness_platform_connector_terraform_cloud":       connector.DatasourceConnectorTerraformCloud(),
				"harness_platform_connector_sumologic":             connector.DatasourceConnectorSumologic(),
				"harness_platform_current_user":                    pl_current_user.DataSourceCurrentUser(),
				"harness_platform_user":                            pl_user.DataSourceUser(),
				"harness_platform_environment":                     pl_environment.DataSourceEnvironment(),
				"harness_platform_environment_group":               pl_environment_group.DataSourceEnvironmentGroup(),
				"harness_platform_environment_clusters_mapping":    pl_environment_clusters_mapping.DataSourceEnvironmentClustersMapping(),
				"harness_platform_environment_service_overrides":   pl_environment_service_overrides.DataSourceEnvironmentServiceOverrides(),
				"harness_platform_service_overrides_v2":            pl_service_overrides_v2.DataSourceServiceOverrides(),
				"harness_platform_gitops_agent":                    gitops_agent.DataSourceGitopsAgent(),
				"harness_platform_gitops_agent_deploy_yaml":        agent_yaml.DataSourceGitopsAgentDeployYaml(),
				"harness_platform_gitops_applications":             gitops_applications.DataSourceGitopsApplications(),
				"harness_platform_gitops_cluster":                  gitops_cluster.DataSourceGitopsCluster(),
				"harness_platform_gitops_gnupg":                    gitops_gnupg.DataSourceGitopsGnupg(),
				"harness_platform_gitops_app_project_mapping":      gitops_project_mapping.DatasourceGitopsAppProjectMapping(),
				"harness_platform_gitops_repository":               gitops_repository.DataSourceGitopsRepository(),
				"harness_platform_gitops_repo_cert":                gitops_repo_cert.DataSourceGitOpsRepoCert(),
				"harness_platform_gitops_repo_cred":                gitops_repo_cred.DataSourceGitOpsRepoCred(),
				"harness_platform_infrastructure":                  pl_infrastructure.DataSourceInfrastructure(),
				"harness_platform_input_set":                       input_set.DataSourceInputSet(),
				"harness_platform_monitored_service":               monitored_service.DataSourceMonitoredService(),
				"harness_platform_organization":                    organization.DataSourceOrganization(),
				"harness_platform_pipeline":                        pipeline.DataSourcePipeline(),
				"harness_platform_permissions":                     pl_permissions.DataSourcePermissions(),
				"harness_platform_project":                         project.DataSourceProject(),
				"harness_platform_service":                         pl_service.DataSourceService(),
				"harness_platform_usergroup":                       usergroup.DataSourceUserGroup(),
				"harness_platform_secret_text":                     secret.DataSourceSecretText(),
				"harness_platform_secret_file":                     secret.DataSourceSecretFile(),
				"harness_platform_secret_sshkey":                   secret.DataSourceSecretSSHKey(),
				"harness_platform_roles":                           roles.DataSourceRoles(),
				"harness_platform_resource_group":                  resource_group.DataSourceResourceGroup(),
				"harness_platform_service_account":                 service_account.DataSourceServiceAccount(),
				"harness_platform_triggers":                        triggers.DataSourceTriggers(),
				"harness_platform_role_assignments":                role_assignments.DataSourceRoleAssignments(),
				"harness_platform_variables":                       variables.DataSourceVariables(),
				"harness_platform_connector_vault":                 connector.DataSourceConnectorVault(),
				"harness_platform_filters":                         filters.DataSourceFilters(),
				"harness_platform_pipeline_filters":                pipeline_filters.DataSourcePipelineFilters(),
				"harness_platform_ccm_filters":                     ccm_filters.DataSourceCCMFilters(),
				"harness_platform_template_filters":                template_filters.DataSourceTemplateFilters(),
				"harness_application":                              application.DataSourceApplication(),
				"harness_current_account":                          account.DataSourceCurrentAccountConnector(),
				"harness_delegate":                                 delegate.DataSourceDelegate(),
				"harness_delegate_ids":                             delegate.DataSourceDelegateIds(),
				"harness_encrypted_text":                           secrets.DataSourceEncryptedText(),
				"harness_environment":                              environment.DataSourceEnvironment(),
				"harness_git_connector":                            cd_connector.DataSourceGitConnector(),
				"harness_secret_manager":                           secrets.DataSourceSecretManager(),
				"harness_service":                                  service.DataSourceService(),
				"harness_platform_slo":                             slo.DataSourceSloService(),
				"harness_ssh_credential":                           secrets.DataSourceSshCredential(),
				"harness_sso_provider":                             sso.DataSourceSSOProvider(),
				"harness_user_group":                               user.DataSourceUserGroup(),
				"harness_user":                                     user.DataSourceUser(),
				"harness_yaml_config":                              yamlconfig.DataSourceYamlConfig(),
				"harness_platform_connector_azure_cloud_provider":  connector.DataSourceConnectorAzureCloudProvider(),
				"harness_platform_connector_tas":                   connector.DataSourceConnectorTas(),
				"harness_trigger":                                  cd_trigger.DataSourceTrigger(),
				"harness_platform_policy":                          policy.DataSourcePolicy(),
				"harness_platform_policyset":                       policyset.DataSourcePolicyset(),
				"harness_platform_manual_freeze":                   manual_freeze.DataSourceManualFreeze(),
				"harness_platform_connector_service_now":           connector.DataSourceConnectorSerivceNow(),
				"harness_platform_apikey":                          pl_apikey.DataSourceApiKey(),
				"harness_platform_token":                           pl_token.DataSourceToken(),
				"harness_platform_file_store_file":                 file_store.DataSourceFileStoreNodeFile(),
				"harness_platform_file_store_folder":               file_store.DataSourceFileStoreNodeFolder(),
				"harness_platform_delegatetoken":                   pl_delegatetoken.DataSourceDelegateToken(),
			},
			ResourcesMap: map[string]*schema.Resource{
				"harness_platform_template":                        pl_template.ResourceTemplate(),
				"harness_platform_connector_azure_key_vault":       connector.ResourceConnectorAzureKeyVault(),
				"harness_platform_connector_gcp_cloud_cost":        connector.ResourceConnectorGCPCloudCost(),
				"harness_platform_connector_kubernetes_cloud_cost": connector.ResourceConnectorKubernetesCloudCost(),
				"harness_platform_connector_azure_cloud_cost":      connector.ResourceConnectorAzureCloudCost(),
				"harness_platform_connector_appdynamics":           connector.ResourceConnectorAppDynamics(),
				"harness_platform_connector_elasticsearch":         connector.ResourceConnectorElasticSearch(),
				"harness_platform_connector_artifactory":           connector.ResourceConnectorArtifactory(),
				"harness_platform_connector_aws_secret_manager":    connector.ResourceConnectorAwsSM(),
				"harness_platform_connector_aws":                   connector.ResourceConnectorAws(),
				"harness_platform_connector_awscc":                 connector.ResourceConnectorAwsCC(),
				"harness_platform_connector_awskms":                connector.ResourceConnectorAwsKms(),
				"harness_platform_connector_bitbucket":             connector.ResourceConnectorBitbucket(),
				"harness_platform_connector_datadog":               connector.ResourceConnectorDatadog(),
				"harness_platform_connector_docker":                connector.ResourceConnectorDocker(),
				"harness_platform_connector_dynatrace":             connector.ResourceConnectorDynatrace(),
				"harness_platform_connector_gcp":                   connector.ResourceConnectorGcp(),
				"harness_platform_connector_gcp_secret_manager":    connector.ResourceConnectorGCPSecretManager(),
				"harness_platform_connector_git":                   connector.ResourceConnectorGit(),
				"harness_platform_connector_github":                connector.ResourceConnectorGithub(),
				"harness_platform_connector_gitlab":                connector.ResourceConnectorGitlab(),
				"harness_platform_connector_helm":                  connector.ResourceConnectorHelm(),
				"harness_platform_connector_oci_helm":              connector.ResourceConnectorOciHelm(),
				"harness_platform_connector_jira":                  connector.ResourceConnectorJira(),
				"harness_platform_connector_jenkins":               connector.ResourceConnectorJenkins(),
				"harness_platform_connector_kubernetes":            connector.ResourceConnectorK8s(),
				"harness_platform_connector_newrelic":              connector.ResourceConnectorNewRelic(),
				"harness_platform_connector_nexus":                 connector.ResourceConnectorNexus(),
				"harness_platform_connector_pagerduty":             connector.ResourceConnectorPagerDuty(),
				"harness_platform_connector_prometheus":            connector.ResourceConnectorPrometheus(),
				"harness_platform_connector_rancher":               connector.ResourceConnectorK8sRancher(),
				"harness_platform_connector_splunk":                connector.ResourceConnectorSplunk(),
				"harness_platform_connector_spot":                  connector.ResourceConnectorSpot(),
				"harness_platform_connector_terraform_cloud":       connector.ResourceConnectorTerraformCloud(),
				"harness_platform_connector_sumologic":             connector.ResourceConnectorSumologic(),
				"harness_platform_environment":                     pl_environment.ResourceEnvironment(),
				"harness_platform_environment_group":               pl_environment_group.ResourceEnvironmentGroup(),
				"harness_platform_environment_clusters_mapping":    pl_environment_clusters_mapping.ResourceEnvironmentClustersMapping(),
				"harness_platform_environment_service_overrides":   pl_environment_service_overrides.ResourceEnvironmentServiceOverrides(),
				"harness_platform_feature_flag":                    feature_flag.ResourceFeatureFlag(),
				"harness_platform_feature_flag_target":             feature_flag_target.ResourceFeatureFlagTarget(),
				"harness_platform_service_overrides_v2":            pl_service_overrides_v2.ResourceServiceOverrides(),
				"harness_platform_ff_api_key":                      ff_api_key.ResourceFFApiKey(),
				"harness_platform_gitops_agent":                    gitops_agent.ResourceGitopsAgent(),
				"harness_platform_gitops_applications":             gitops_applications.ResourceGitopsApplication(),
				"harness_platform_gitops_cluster":                  gitops_cluster.ResourceGitopsCluster(),
				"harness_platform_gitops_gnupg":                    gitops_gnupg.ResourceGitopsGnupg(),
				"harness_platform_gitops_app_project_mapping":      gitops_project_mapping.ResourceGitopsAppProjectMapping(),
				"harness_platform_gitops_repository":               gitops_repository.ResourceGitopsRepositories(),
				"harness_platform_gitops_repo_cert":                gitops_repo_cert.ResourceGitopsRepoCerts(),
				"harness_platform_gitops_repo_cred":                gitops_repo_cred.ResourceGitopsRepoCred(),
				"harness_platform_infrastructure":                  pl_infrastructure.ResourceInfrastructure(),
				"harness_platform_input_set":                       input_set.ResourceInputSet(),
				"harness_platform_monitored_service":               monitored_service.ResourceMonitoredService(),
				"harness_platform_organization":                    organization.ResourceOrganization(),
				"harness_platform_pipeline":                        pipeline.ResourcePipeline(),
				"harness_platform_project":                         project.ResourceProject(),
				"harness_platform_service":                         pl_service.ResourceService(),
				"harness_platform_user":                            pl_user.ResourceUser(),
				"harness_platform_usergroup":                       usergroup.ResourceUserGroup(),
				"harness_platform_secret_text":                     secret.ResourceSecretText(),
				"harness_platform_secret_file":                     secret.ResourceSecretFile(),
				"harness_platform_secret_sshkey":                   secret.ResourceSecretSSHKey(),
				"harness_platform_roles":                           roles.ResourceRoles(),
				"harness_platform_resource_group":                  resource_group.ResourceResourceGroup(),
				"harness_platform_service_account":                 service_account.ResourceServiceAccount(),
				"harness_platform_triggers":                        triggers.ResourceTriggers(),
				"harness_platform_role_assignments":                role_assignments.ResourceRoleAssignments(),
				"harness_platform_variables":                       variables.ResourceVariables(),
				"harness_platform_connector_vault":                 connector.ResourceConnectorVault(),
				"harness_platform_filters":                         filters.ResourceFilters(),
				"harness_platform_pipeline_filters":                pipeline_filters.ResourcePipelineFilters(),
				"harness_platform_ccm_filters":                     ccm_filters.ResourceCCMFilters(),
				"harness_platform_template_filters":                template_filters.ResourceTemplateFilters(),
				"harness_add_user_to_group":                        user.ResourceAddUserToGroup(),
				"harness_application_gitsync":                      application.ResourceApplicationGitSync(),
				"harness_application":                              application.ResourceApplication(),
				"harness_delegate_approval":                        delegate.ResourceDelegateApproval(),
				"harness_cloudprovider_aws":                        cloudprovider.ResourceCloudProviderAws(),
				"harness_cloudprovider_azure":                      cloudprovider.ResourceCloudProviderAzure(),
				"harness_cloudprovider_datacenter":                 cloudprovider.ResourceCloudProviderDataCenter(),
				"harness_cloudprovider_gcp":                        cloudprovider.ResourceCloudProviderGcp(),
				"harness_cloudprovider_kubernetes":                 cloudprovider.ResourceCloudProviderK8s(),
				"harness_cloudprovider_spot":                       cloudprovider.ResourceCloudProviderSpot(),
				"harness_cloudprovider_tanzu":                      cloudprovider.ResourceCloudProviderTanzu(),
				"harness_encrypted_text":                           secrets.ResourceEncryptedText(),
				"harness_environment":                              environment.ResourceEnvironment(),
				"harness_git_connector":                            cd_connector.ResourceGitConnector(),
				"harness_infrastructure_definition":                environment.ResourceInfraDefinition(),
				"harness_service_ami":                              service.ResourceAMIService(),
				"harness_service_aws_codedeploy":                   service.ResourceAWSCodeDeployService(),
				"harness_service_aws_lambda":                       service.ResourceAWSLambdaService(),
				"harness_service_ecs":                              service.ResourceECSService(),
				"harness_service_helm":                             service.ResourceHelmService(),
				"harness_service_kubernetes":                       service.ResourceKubernetesService(),
				"harness_service_ssh":                              service.ResourceSSHService(),
				"harness_service_tanzu":                            service.ResourcePCFService(),
				"harness_service_winrm":                            service.ResourceWinRMService(),
				"harness_platform_slo":                             slo.ResourceSloService(),
				"harness_ssh_credential":                           secrets.ResourceSSHCredential(),
				"harness_user_group":                               user.ResourceUserGroup(),
				"harness_user_group_permissions":                   user.ResourceUserGroupPermissions(),
				"harness_user":                                     user.ResourceUser(),
				"harness_yaml_config":                              yamlconfig.ResourceYamlConfig(),
				"harness_platform_connector_azure_cloud_provider":  connector.ResourceConnectorAzureCloudProvider(),
				"harness_platform_connector_tas":                   connector.ResourceConnectorTas(),
				"harness_platform_policy":                          policy.ResourcePolicy(),
				"harness_platform_policyset":                       policyset.ResourcePolicyset(),
				"harness_platform_manual_freeze":                   manual_freeze.ResourceManualFreeze(),
				"harness_platform_connector_service_now":           connector.ResourceConnectorServiceNow(),
				"harness_platform_apikey":                          pl_apikey.ResourceApiKey(),
				"harness_platform_token":                           pl_token.ResourceToken(),
				"harness_platform_file_store_file":                 file_store.ResourceFileStoreNodeFile(),
				"harness_platform_file_store_folder":               file_store.ResourceFileStoreNodeFolder(),
				"harness_platform_delegatetoken":                   pl_delegatetoken.ResourceDelegateToken(),
			},
		}

		p.ConfigureContextFunc = configure(version, p)

		return p
	}
}

func getHttpClient(logger *logrus.Logger) *retryablehttp.Client {
	httpClient := retryablehttp.NewClient()
	httpClient.HTTPClient.Transport = logging.NewTransport(harness.SDKName, logger, cleanhttp.DefaultPooledClient().Transport)
	httpClient.RetryMax = 10
	return httpClient
}

func getOpenApiHttpClient(logger *logrus.Logger) *retryablehttp.Client {
	httpClient := retryablehttp.NewClient()
	httpClient.HTTPClient.Transport = openapi_client_logging.NewTransport(harness.SDKName, logger, cleanhttp.DefaultPooledClient().Transport)
	httpClient.RetryMax = 10
	return httpClient
}

func getCDClient(d *schema.ResourceData, version string) *cd.ApiClient {
	cfg := cd.DefaultConfig()
	cfg.AccountId = d.Get("account_id").(string)
	cfg.Endpoint = d.Get("endpoint").(string)
	cfg.APIKey = d.Get("api_key").(string)
	cfg.UserAgent = fmt.Sprintf("terraform-provider-harness-%s", version)
	cfg.HTTPClient = getHttpClient(cfg.Logger)
	cfg.DebugLogging = logging.IsDebugOrHigher(cfg.Logger)

	client, err := cd.NewClient(cfg)

	if err != nil {
		log.Printf("[WARN] error creating CD client: %s", err)
	}

	return client
}

func getPLClient(d *schema.ResourceData, version string) *nextgen.APIClient {
	cfg := nextgen.NewConfiguration()
	client := nextgen.NewAPIClient(&nextgen.Configuration{
		AccountId:    d.Get("account_id").(string),
		BasePath:     d.Get("endpoint").(string),
		ApiKey:       d.Get("platform_api_key").(string),
		UserAgent:    fmt.Sprintf("terraform-provider-harness-platform-%s", version),
		HTTPClient:   getHttpClient(cfg.Logger),
		DebugLogging: logging.IsDebugOrHigher(cfg.Logger),
	})

	return client
}

func getClient(d *schema.ResourceData, version string) *openapi_client_nextgen.APIClient {
	cfg := openapi_client_nextgen.NewConfiguration()
	client := openapi_client_nextgen.NewAPIClient(&openapi_client_nextgen.Configuration{
		AccountId:    d.Get("account_id").(string),
		BasePath:     d.Get("endpoint").(string),
		ApiKey:       d.Get("platform_api_key").(string),
		UserAgent:    fmt.Sprintf("terraform-provider-harness-platform-%s", version),
		HTTPClient:   getOpenApiHttpClient(cfg.Logger),
		DebugLogging: openapi_client_logging.IsDebugOrHigher(cfg.Logger),
	})

	return client
}

// Setup the client for interacting with the Harness API
func configure(version string, p *schema.Provider) func(context.Context, *schema.ResourceData) (interface{}, diag.Diagnostics) {
	return func(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
		return &internal.Session{
			AccountId: d.Get("account_id").(string),
			Endpoint:  d.Get("endpoint").(string),
			CDClient:  getCDClient(d, version),
			PLClient:  getPLClient(d, version),
			Client:    getClient(d, version),
		}, nil
	}
}
