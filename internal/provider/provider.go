package provider

import (
	"context"
	"net/http"
	"time"

	"github.com/harness-io/harness-go-sdk/harness/api"
	"github.com/harness-io/harness-go-sdk/harness/envvar"
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
					Type:        schema.TypeString,
					Optional:    true,
					DefaultFunc: schema.EnvDefaultFunc(envvar.HarnessEndpoint, api.DefaultApiUrl),
				},
				"account_id": {
					Type:        schema.TypeString,
					Optional:    true,
					DefaultFunc: schema.EnvDefaultFunc(envvar.HarnessAccountId, nil),
				},
				"api_key": {
					Type:        schema.TypeString,
					Optional:    true,
					DefaultFunc: schema.EnvDefaultFunc(envvar.HarnessApiKey, nil),
				},
				"bearer_token": {
					Type:        schema.TypeString,
					Optional:    true,
					DefaultFunc: schema.EnvDefaultFunc(envvar.HarnessBearerToken, nil),
				},
			},
			DataSourcesMap: map[string]*schema.Resource{
				"harness_application":    dataSourceApplication(),
				"harness_encrypted_text": dataSourceEncryptedText(),
				"harness_git_connector":  dataSourceGitConnector(),
				"harness_service":        dataSourceService(),
			},
			ResourcesMap: map[string]*schema.Resource{
				"harness_application":            resourceApplication(),
				"harness_encrypted_text":         resourceEncryptedText(),
				"harness_git_connector":          resourceGitConnector(),
				"harness_ssh_credential":         resourceSSHCredential(),
				"harness_service_kubernetes":     resourceKubernetesService(),
				"harness_service_ami":            resourceAMIService(),
				"harness_service_ecs":            resourceECSService(),
				"harness_service_aws_codedeploy": resourceAWSCodeDeployService(),
				"harness_service_aws_lambda":     resourceAWSLambdaService(),
				"harness_service_pcf":            resourcePCFService(),
				"harness_service_helm":           resourceHelmService(),
				"harness_service_ssh":            resourceSSHService(),
				"harness_service_winrm":          resourceWinRMService(),
			},
		}

		p.ConfigureContextFunc = configure(version, p)

		return p
	}
}

type HarnessClient struct {
	GraphQLClient *api.Client
	CaacClient    *api.ConfigAsCodeClient
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
			UserAgent:   p.UserAgent("terraform-provider-harness", version),
			Endpoint:    d.Get("endpoint").(string),
			AccountId:   d.Get("account_id").(string),
			APIKey:      d.Get("api_key").(string),
			BearerToken: d.Get("bearer_token").(string),
			HTTPClient:  httpClient,
		}, nil
	}
}
