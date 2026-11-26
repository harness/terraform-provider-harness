package agent_test

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	pschema "github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	rschema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5"
)

// A minimal stand-in for the terraform-plugin-testing "echo" provider. Its provider-level `data` accepts an
// ephemeral value, and its `echo` resource copies that value into state so a test can assert on it.
type echoProvider struct{}

type echoProviderModel struct {
	Data types.String `tfsdk:"data"`
}

func (echoProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "echo"
}

func (echoProvider) Schema(_ context.Context, _ provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = pschema.Schema{Attributes: map[string]pschema.Attribute{"data": pschema.StringAttribute{Optional: true}}}
}

func (echoProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var m echoProviderModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &m)...)
	resp.ResourceData = m.Data
}

func (echoProvider) DataSources(context.Context) []func() datasource.DataSource { return nil }
func (echoProvider) Resources(context.Context) []func() resource.Resource {
	return []func() resource.Resource{func() resource.Resource { return &echoResource{} }}
}

type echoResource struct{ data types.String }

func (*echoResource) Metadata(_ context.Context, _ resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "echo"
}

func (*echoResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = rschema.Schema{Attributes: map[string]rschema.Attribute{"id": rschema.StringAttribute{Computed: true}, "data": rschema.StringAttribute{Computed: true}, "trigger": rschema.StringAttribute{Optional: true}}}
}

func (r *echoResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if v, ok := req.ProviderData.(types.String); ok {
		r.data = v
	}
}

func (r *echoResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var trigger types.String
	resp.Diagnostics.Append(req.Plan.GetAttribute(ctx, path.Root("trigger"), &trigger)...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("trigger"), trigger)...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), "echo")...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("data"), r.data)...)
}

func (r *echoResource) Read(context.Context, resource.ReadRequest, *resource.ReadResponse) {}

func (r *echoResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var trigger types.String
	resp.Diagnostics.Append(req.Plan.GetAttribute(ctx, path.Root("trigger"), &trigger)...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("trigger"), trigger)...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), "echo")...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("data"), r.data)...)
}

func (r *echoResource) Delete(context.Context, resource.DeleteRequest, *resource.DeleteResponse) {}

func echoProviderServer() (tfprotov5.ProviderServer, error) {
	return providerserver.NewProtocol5(echoProvider{})(), nil
}
