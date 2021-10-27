package provider

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/harness-io/harness-go-sdk/harness/api"
	"github.com/harness-io/harness-go-sdk/harness/api/nextgen"
	"github.com/harness-io/harness-go-sdk/harness/helpers"
	"github.com/harness-io/terraform-provider-harness/internal/service/cd"
	"github.com/harness-io/terraform-provider-harness/internal/service/ng"

	// "github.com/harness-io/terraform-provider-harness/internal/service/cd"
	// "github.com/harness-io/terraform-provider-harness/internal/service/ng"
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
					Description: fmt.Sprintf("The Harness API key. This can also be set using the `%s` environment variable.", helpers.EnvVars.HarnessApiKey.String()),
					Type:        schema.TypeString,
					Optional:    true,
					DefaultFunc: schema.EnvDefaultFunc(helpers.EnvVars.HarnessApiKey.String(), nil),
				},
				"ng_endpoint": {
					Description: fmt.Sprintf("The URL of the Harness nextgen API. The default is `%s`. This can also be set using the `%s` environment variable.", api.DefaultNGApiUrl, helpers.EnvVars.HarnessNGEndpoint.String()),
					Type:        schema.TypeString,
					Required:    true,
					DefaultFunc: schema.EnvDefaultFunc(helpers.EnvVars.HarnessNGEndpoint.String(), api.DefaultNGApiUrl),
				},
				"ng_api_key": {
					Description: fmt.Sprintf("The Harness nextgen API key. This can also be set using the `%s` environment variable.", helpers.EnvVars.HarnessNGApiKey.String()),
					Type:        schema.TypeString,
					Optional:    true,
					DefaultFunc: schema.EnvDefaultFunc(helpers.EnvVars.HarnessNGApiKey.String(), nil),
				},
			},
			DataSourcesMap: map[string]*schema.Resource{
				"harness_application":    cd.DataSourceApplication(),
				"harness_delegate":       cd.DataSourceDelegate(),
				"harness_secret_manager": cd.DataSourceSecretManager(),
				"harness_encrypted_text": cd.DataSourceEncryptedText(),
				"harness_git_connector":  cd.DataSourceGitConnector(),
				"harness_service":        cd.DataSourceService(),
				"harness_environment":    cd.DataSourceEnvironment(),
				"harness_sso_provider":   cd.DataSourceSSOProvider(),
				"harness_user":           cd.DataSourceUser(),
				"harness_user_group":     cd.DataSourceUserGroup(),
				"harness_yaml_config":    cd.DataSourceYamlConfig(),
				"harness_project":        ng.DataSourceProject(),
				"harness_organization":   ng.DataSourceOrganization(),
				"harness_connector":      ng.DataSourceConnector(),
				"harness_current_user":   ng.DataSourceCurrentUser(),
			},
			ResourcesMap: map[string]*schema.Resource{
				"harness_application":               cd.ResourceApplication(),
				"harness_encrypted_text":            cd.ResourceEncryptedText(),
				"harness_git_connector":             cd.ResourceGitConnector(),
				"harness_ssh_credential":            cd.ResourceSSHCredential(),
				"harness_service_kubernetes":        cd.ResourceKubernetesService(),
				"harness_service_ami":               cd.ResourceAMIService(),
				"harness_service_ecs":               cd.ResourceECSService(),
				"harness_service_aws_codedeploy":    cd.ResourceAWSCodeDeployService(),
				"harness_service_aws_lambda":        cd.ResourceAWSLambdaService(),
				"harness_service_tanzu":             cd.ResourcePCFService(),
				"harness_service_helm":              cd.ResourceHelmService(),
				"harness_service_ssh":               cd.ResourceSSHService(),
				"harness_service_winrm":             cd.ResourceWinRMService(),
				"harness_environment":               cd.ResourceEnvironment(),
				"harness_cloudprovider_datacenter":  cd.ResourceCloudProviderDataCenter(),
				"harness_cloudprovider_aws":         cd.ResourceCloudProviderAws(),
				"harness_cloudprovider_azure":       cd.ResourceCloudProviderAzure(),
				"harness_cloudprovider_tanzu":       cd.ResourceCloudProviderTanzu(),
				"harness_cloudprovider_gcp":         cd.ResourceCloudProviderGcp(),
				"harness_cloudprovider_kubernetes":  cd.ResourceCloudProviderK8s(),
				"harness_cloudprovider_spot":        cd.ResourceCloudProviderSpot(),
				"harness_user":                      cd.ResourceUser(),
				"harness_user_group":                cd.ResourceUserGroup(),
				"harness_add_user_to_group":         cd.ResourceAddUserToGroup(),
				"harness_infrastructure_definition": cd.ResourceInfraDefinition(),
				"harness_yaml_config":               cd.ResourceYamlConfig(),
				"harness_application_gitsync":       cd.ResourceApplicationGitSync(),
				"harness_project":                   ng.ResourceProject(),
				"harness_organization":              ng.ResourceOrganization(),
				"harness_connector":                 ng.ResourceConnector(),
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
			Backoff: retryablehttp.DefaultBackoff,
			// CheckRetry: retryablehttp.DefaultRetryPolicy,
			CheckRetry: func(ctx context.Context, resp *http.Response, err error) (bool, error) {
				if resp.StatusCode == http.StatusInternalServerError {
					return false, err
				}
				return retryablehttp.ErrorPropagatedRetryPolicy(ctx, resp, err)
			},
		}

		userAgent := p.UserAgent("terraform-provider-harness", version)

		return &api.Client{
			UserAgent:  userAgent,
			Endpoint:   d.Get("endpoint").(string),
			AccountId:  d.Get("account_id").(string),
			APIKey:     d.Get("api_key").(string),
			HTTPClient: httpClient,
			NGClient: nextgen.NewAPIClient(&nextgen.Configuration{
				BasePath: d.Get("ng_endpoint").(string),
				DefaultHeader: map[string]string{
					helpers.HTTPHeaders.ApiKey.String(): d.Get("ng_api_key").(string),
				},
				UserAgent:  userAgent,
				HTTPClient: httpClient,
			}),
		}, nil
	}
}
