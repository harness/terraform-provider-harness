package provider

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/harness-io/harness-go-sdk/harness/api"
	"github.com/harness-io/harness-go-sdk/harness/helpers"
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

func New(version string) func() *schema.Provider {
	return func() *schema.Provider {
		p := &schema.Provider{
			Schema: map[string]*schema.Schema{
				"endpoint": {
					Description: fmt.Sprintf("The URL of the Harness API endpoint. The default is `https://app.harness.io`. This can also be set using the `%s` environment variable.", helpers.EnvVars.HarnessEndpoint.String()),
					Type:        schema.TypeString,
					Required:    true,
					DefaultFunc: schema.EnvDefaultFunc(helpers.EnvVars.HarnessEndpoint.String(), api.DefaultApiUrl),
				},
				"account_id": {
					Description: fmt.Sprintf("The Harness account id. This can also be set using the `%s` environment variable.", helpers.EnvVars.HarnessAccountId.String()),
					Type:        schema.TypeString,
					Required:    true,
					DefaultFunc: schema.EnvDefaultFunc(helpers.EnvVars.HarnessAccountId.String(), nil),
				},
				"api_key": {
					Description: fmt.Sprintf("The Harness api key. This can also be set using the `%s` environment variable.", helpers.EnvVars.HarnessApiKey.String()),
					Type:        schema.TypeString,
					Required:    true,
					DefaultFunc: schema.EnvDefaultFunc(helpers.EnvVars.HarnessApiKey.String(), nil),
				},
			},
			DataSourcesMap: map[string]*schema.Resource{
				"harness_application":    dataSourceApplication(),
				"harness_secret_manager": dataSourceSecretManager(),
				"harness_encrypted_text": dataSourceEncryptedText(),
				"harness_git_connector":  dataSourceGitConnector(),
				"harness_service":        dataSourceService(),
				"harness_environment":    dataSourceEnvironment(),
				"harness_sso_provider":   dataSourceSSOProvider(),
				"harness_user":           dataSourceUser(),
				"harness_user_group":     dataSourceUserGroup(),
				"harness_yaml_config":    dataSourceYamlConfig(),
			},
			ResourcesMap: map[string]*schema.Resource{
				"harness_application":               resourceApplication(),
				"harness_encrypted_text":            resourceEncryptedText(),
				"harness_git_connector":             resourceGitConnector(),
				"harness_ssh_credential":            resourceSSHCredential(),
				"harness_service_kubernetes":        resourceKubernetesService(),
				"harness_service_ami":               resourceAMIService(),
				"harness_service_ecs":               resourceECSService(),
				"harness_service_aws_codedeploy":    resourceAWSCodeDeployService(),
				"harness_service_aws_lambda":        resourceAWSLambdaService(),
				"harness_service_tanzu":             resourcePCFService(),
				"harness_service_helm":              resourceHelmService(),
				"harness_service_ssh":               resourceSSHService(),
				"harness_service_winrm":             resourceWinRMService(),
				"harness_environment":               resourceEnvironment(),
				"harness_cloudprovider_datacenter":  resourceCloudProviderDataCenter(),
				"harness_cloudprovider_aws":         resourceCloudProviderAws(),
				"harness_cloudprovider_azure":       resourceCloudProviderAzure(),
				"harness_cloudprovider_tanzu":       resourceCloudProviderTanzu(),
				"harness_cloudprovider_gcp":         resourceCloudProviderGcp(),
				"harness_cloudprovider_kubernetes":  resourceCloudProviderK8s(),
				"harness_cloudprovider_spot":        resourceCloudProviderSpot(),
				"harness_user":                      resourceUser(),
				"harness_user_group":                resourceUserGroup(),
				"harness_add_user_to_group":         resourceAddUserToGroup(),
				"harness_infrastructure_definition": resourceInfraDefinition(),
				"harness_yaml_config":               resourceYamlConfig(),
				"harness_application_gitsync":       resourceApplicationGitSync(),
			},
		}

		p.ConfigureContextFunc = configure(version, p)

		return p
	}
}

// Setup the client for interacting with the Harness API
func configure(version string, p *schema.Provider) func(context.Context, *schema.ResourceData) (interface{}, diag.Diagnostics) {
	return func(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {

		httpClient := &retryablehttp.Client{
			RetryMax:     125,
			RetryWaitMin: 5 * time.Second,
			RetryWaitMax: 30 * time.Second,
			HTTPClient: &http.Client{
				Timeout: 30 * time.Second,
			},
			Backoff:    retryablehttp.DefaultBackoff,
			CheckRetry: retryablehttp.DefaultRetryPolicy,
		}

		return &api.Client{
			UserAgent:  p.UserAgent("terraform-provider-harness", version),
			Endpoint:   d.Get("endpoint").(string),
			AccountId:  d.Get("account_id").(string),
			APIKey:     d.Get("api_key").(string),
			HTTPClient: httpClient,
		}, nil
	}
}
