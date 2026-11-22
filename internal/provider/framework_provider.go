package provider

import (
	"context"
	"os"

	helpers "github.com/harness/harness-go-sdk/harness/helpers"
	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal"
	gitops_agent "github.com/harness/terraform-provider-harness/internal/service/platform/gitops/agent"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/ephemeral"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// frameworkProvider serves only ephemeral resources, which the SDKv2 provider cannot express.
// It is muxed with the SDKv2 provider in main.go, so its provider schema must stay identical
// to the SDKv2 one. Descriptions are read from the SDKv2 schema to prevent drift, and set as
// MarkdownDescription because init() in provider.go sets schema.DescriptionKind = StringMarkdown.
type frameworkProvider struct {
	version string
}

var (
	_ provider.Provider                       = &frameworkProvider{}
	_ provider.ProviderWithEphemeralResources = &frameworkProvider{}
)

// NewFrameworkProvider returns the terraform-plugin-framework half of the muxed provider.
func NewFrameworkProvider(version string) func() provider.Provider {
	return func() provider.Provider {
		return &frameworkProvider{version: version}
	}
}

func (p *frameworkProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "harness"
	resp.Version = p.version
}

func (p *frameworkProvider) Schema(_ context.Context, _ provider.SchemaRequest, resp *provider.SchemaResponse) {
	sdkSchema := Provider(p.version)().Schema
	attrs := map[string]schema.Attribute{}
	for _, name := range []string{"endpoint", "fme_admin_api_endpoint", "account_id", "api_key", "platform_api_key"} {
		attrs[name] = schema.StringAttribute{
			Optional:            true,
			MarkdownDescription: sdkSchema[name].Description,
		}
	}
	resp.Schema = schema.Schema{Attributes: attrs}
}

type frameworkProviderModel struct {
	Endpoint       types.String `tfsdk:"endpoint"`
	FMEAdminAPI    types.String `tfsdk:"fme_admin_api_endpoint"`
	AccountId      types.String `tfsdk:"account_id"`
	ApiKey         types.String `tfsdk:"api_key"`
	PlatformApiKey types.String `tfsdk:"platform_api_key"`
}

func valueOrEnv(v types.String, env string, def string) string {
	if !v.IsNull() && !v.IsUnknown() && v.ValueString() != "" {
		return v.ValueString()
	}
	if e := os.Getenv(env); e != "" {
		return e
	}
	return def
}

// Configure builds a minimal session (platform client only). The SDKv2 provider already
// validates and warns on the shared provider configuration, so no diagnostics are repeated here.
func (p *frameworkProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var cfg frameworkProviderModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &cfg)...)
	if resp.Diagnostics.HasError() {
		return
	}

	endpoint := valueOrEnv(cfg.Endpoint, helpers.EnvVars.Endpoint.String(), utils.BaseUrl)
	accountId := valueOrEnv(cfg.AccountId, helpers.EnvVars.AccountId.String(), "")
	platformApiKey := valueOrEnv(cfg.PlatformApiKey, helpers.EnvVars.PlatformApiKey.String(), "")

	session := &internal.Session{
		AccountId: accountId,
		Endpoint:  endpoint,
		PLClient:  newPLClient(accountId, endpoint, platformApiKey, p.version),
	}
	resp.EphemeralResourceData = session
}

func (p *frameworkProvider) Resources(_ context.Context) []func() resource.Resource {
	return nil
}

func (p *frameworkProvider) DataSources(_ context.Context) []func() datasource.DataSource {
	return nil
}

func (p *frameworkProvider) EphemeralResources(_ context.Context) []func() ephemeral.EphemeralResource {
	return []func() ephemeral.EphemeralResource{
		gitops_agent.NewEphemeralGitopsAgentToken,
	}
}
