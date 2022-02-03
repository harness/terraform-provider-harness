package provider

import (
	"context"
	"fmt"

	sdk "github.com/harness-io/harness-go-sdk"
	"github.com/harness-io/harness-go-sdk/harness"
	"github.com/harness-io/harness-go-sdk/harness/helpers"
	"github.com/harness-io/harness-go-sdk/harness/utils"
	"github.com/harness-io/terraform-provider-harness/internal/service/cd/application"
	"github.com/harness-io/terraform-provider-harness/internal/service/cd/cloudprovider"
	cd_connector "github.com/harness-io/terraform-provider-harness/internal/service/cd/connector"
	"github.com/harness-io/terraform-provider-harness/internal/service/cd/delegate"
	"github.com/harness-io/terraform-provider-harness/internal/service/cd/environment"
	"github.com/harness-io/terraform-provider-harness/internal/service/cd/secrets"
	"github.com/harness-io/terraform-provider-harness/internal/service/cd/service"
	"github.com/harness-io/terraform-provider-harness/internal/service/cd/sso"
	"github.com/harness-io/terraform-provider-harness/internal/service/cd/user"
	"github.com/harness-io/terraform-provider-harness/internal/service/cd/yamlconfig"
	"github.com/harness-io/terraform-provider-harness/internal/service/ng"
	"github.com/harness-io/terraform-provider-harness/internal/service/ng/connector"
	"github.com/hashicorp/go-cleanhttp"
	"github.com/hashicorp/go-retryablehttp"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/logging"
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
					Description: fmt.Sprintf("The Harness API key. This can also be set using the `%s` environment variable.", helpers.EnvVars.ApiKey.String()),
					Type:        schema.TypeString,
					Optional:    true,
					DefaultFunc: schema.EnvDefaultFunc(helpers.EnvVars.ApiKey.String(), nil),
				},
				"ng_api_key": {
					Description: fmt.Sprintf("The Harness nextgen API key. This can also be set using the `%s` environment variable.", helpers.EnvVars.NGApiKey.String()),
					Type:        schema.TypeString,
					Optional:    true,
					DefaultFunc: schema.EnvDefaultFunc(helpers.EnvVars.NGApiKey.String(), nil),
				},
			},
			DataSourcesMap: map[string]*schema.Resource{
				"harness_application":    application.DataSourceApplication(),
				"harness_connector":      ng.DataSourceConnector(),
				"harness_current_user":   ng.DataSourceCurrentUser(),
				"harness_delegate":       delegate.DataSourceDelegate(),
				"harness_delegate_ids":   delegate.DataSourceDelegateIds(),
				"harness_encrypted_text": secrets.DataSourceEncryptedText(),
				"harness_environment":    environment.DataSourceEnvironment(),
				"harness_git_connector":  cd_connector.DataSourceGitConnector(),
				"harness_organization":   ng.DataSourceOrganization(),
				"harness_project":        ng.DataSourceProject(),
				"harness_secret_manager": secrets.DataSourceSecretManager(),
				"harness_service":        service.DataSourceService(),
				"harness_ssh_credential": secrets.DataSourceSshCredential(),
				"harness_sso_provider":   sso.DataSourceSSOProvider(),
				"harness_user_group":     user.DataSourceUserGroup(),
				"harness_user":           user.DataSourceUser(),
				"harness_yaml_config":    yamlconfig.DataSourceYamlConfig(),
			},
			ResourcesMap: map[string]*schema.Resource{
				"harness_add_user_to_group":         user.ResourceAddUserToGroup(),
				"harness_application_gitsync":       application.ResourceApplicationGitSync(),
				"harness_application":               application.ResourceApplication(),
				"harness_delegate_approval":         delegate.ResourceDelegateApproval(),
				"harness_cloudprovider_aws":         cloudprovider.ResourceCloudProviderAws(),
				"harness_cloudprovider_azure":       cloudprovider.ResourceCloudProviderAzure(),
				"harness_cloudprovider_datacenter":  cloudprovider.ResourceCloudProviderDataCenter(),
				"harness_cloudprovider_gcp":         cloudprovider.ResourceCloudProviderGcp(),
				"harness_cloudprovider_kubernetes":  cloudprovider.ResourceCloudProviderK8s(),
				"harness_cloudprovider_spot":        cloudprovider.ResourceCloudProviderSpot(),
				"harness_cloudprovider_tanzu":       cloudprovider.ResourceCloudProviderTanzu(),
				"harness_connector":                 connector.ResourceConnector(),
				"harness_encrypted_text":            secrets.ResourceEncryptedText(),
				"harness_environment":               environment.ResourceEnvironment(),
				"harness_git_connector":             cd_connector.ResourceGitConnector(),
				"harness_infrastructure_definition": environment.ResourceInfraDefinition(),
				"harness_organization":              ng.ResourceOrganization(),
				"harness_project":                   ng.ResourceProject(),
				"harness_service_ami":               service.ResourceAMIService(),
				"harness_service_aws_codedeploy":    service.ResourceAWSCodeDeployService(),
				"harness_service_aws_lambda":        service.ResourceAWSLambdaService(),
				"harness_service_ecs":               service.ResourceECSService(),
				"harness_service_helm":              service.ResourceHelmService(),
				"harness_service_kubernetes":        service.ResourceKubernetesService(),
				"harness_service_ssh":               service.ResourceSSHService(),
				"harness_service_tanzu":             service.ResourcePCFService(),
				"harness_service_winrm":             service.ResourceWinRMService(),
				"harness_ssh_credential":            secrets.ResourceSSHCredential(),
				"harness_user_group":                user.ResourceUserGroup(),
				"harness_user":                      user.ResourceUser(),
				"harness_yaml_config":               yamlconfig.ResourceYamlConfig(),
			},
		}

		p.ConfigureContextFunc = configure(version, p)

		return p
	}
}

func getHttpClient() *retryablehttp.Client {
	httpClient := retryablehttp.NewClient()
	httpClient.HTTPClient.Transport = logging.NewTransport(harness.SDKName, cleanhttp.DefaultPooledClient().Transport)
	httpClient.RetryMax = 10
	return httpClient
}

// Setup the client for interacting with the Harness API
func configure(version string, p *schema.Provider) func(context.Context, *schema.ResourceData) (interface{}, diag.Diagnostics) {
	return func(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {

		session, err := sdk.NewSession(&sdk.SessionOptions{
			ApiKey:       d.Get("api_key").(string),
			NGApiKey:     d.Get("ng_api_key").(string),
			AccountId:    d.Get("account_id").(string),
			Endpoint:     d.Get("endpoint").(string),
			DebugLogging: logging.IsDebugOrHigher(),
			HTTPClient:   getHttpClient(),
		})
		if err != nil {
			return nil, diag.FromErr(err)
		}

		return session, nil
	}
}
