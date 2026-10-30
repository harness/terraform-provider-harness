package provider

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5"
	"github.com/hashicorp/terraform-plugin-mux/tf5muxserver"
)

// TestMuxServer_ProviderSchemaAndEphemeral verifies the SDKv2 and framework provider schemas
// are identical (mux refuses to start otherwise) and the ephemeral resource is exposed.
func TestMuxServer_ProviderSchemaAndEphemeral(t *testing.T) {
	ctx := context.Background()

	mux, err := tf5muxserver.NewMuxServer(ctx,
		Provider("test")().GRPCProvider,
		providerserver.NewProtocol5(NewFrameworkProvider("test")()),
	)
	if err != nil {
		t.Fatalf("creating mux server: %s", err)
	}

	server := mux.ProviderServer()
	schemaResp, err := server.GetProviderSchema(ctx, &tfprotov5.GetProviderSchemaRequest{})
	if err != nil {
		t.Fatalf("GetProviderSchema: %s", err)
	}
	for _, d := range schemaResp.Diagnostics {
		t.Errorf("schema diagnostic: %s: %s", d.Summary, d.Detail)
	}
	if _, ok := schemaResp.EphemeralResourceSchemas["harness_platform_gitops_agent_token"]; !ok {
		t.Error("ephemeral resource harness_platform_gitops_agent_token not exposed")
	}
	if _, ok := schemaResp.ResourceSchemas["harness_platform_gitops_agent"]; !ok {
		t.Error("SDKv2 resource harness_platform_gitops_agent missing after mux")
	}
}

// Muxing must not add, drop or alter any SDKv2 resource or data source schema: every consumer of the
// existing provider goes through the mux server after this change.
func TestMuxServer_DoesNotAlterSDKv2Schemas(t *testing.T) {
	ctx := context.Background()

	sdkOnly := Provider("test")().GRPCProvider()
	sdkResp, e := sdkOnly.GetProviderSchema(ctx, &tfprotov5.GetProviderSchemaRequest{})
	if e != nil {
		t.Fatalf("SDK GetProviderSchema: %s", e)
	}

	mux, e := tf5muxserver.NewMuxServer(ctx,
		Provider("test")().GRPCProvider,
		providerserver.NewProtocol5(NewFrameworkProvider("test")()),
	)
	if e != nil {
		t.Fatalf("creating mux server: %s", e)
	}
	muxResp, e := mux.ProviderServer().GetProviderSchema(ctx, &tfprotov5.GetProviderSchemaRequest{})
	if e != nil {
		t.Fatalf("mux GetProviderSchema: %s", e)
	}

	if len(muxResp.ResourceSchemas) != len(sdkResp.ResourceSchemas) {
		t.Errorf("resource count changed: sdk=%d mux=%d", len(sdkResp.ResourceSchemas), len(muxResp.ResourceSchemas))
	}
	if len(muxResp.DataSourceSchemas) != len(sdkResp.DataSourceSchemas) {
		t.Errorf("data source count changed: sdk=%d mux=%d", len(sdkResp.DataSourceSchemas), len(muxResp.DataSourceSchemas))
	}
	for name, want := range sdkResp.ResourceSchemas {
		got, ok := muxResp.ResourceSchemas[name]
		if !ok {
			t.Errorf("resource %s missing after mux", name)
			continue
		}
		if !got.ValueType().Equal(want.ValueType()) {
			t.Errorf("resource %s schema type changed by mux", name)
		}
	}
	for name, want := range sdkResp.DataSourceSchemas {
		got, ok := muxResp.DataSourceSchemas[name]
		if !ok {
			t.Errorf("data source %s missing after mux", name)
			continue
		}
		if !got.ValueType().Equal(want.ValueType()) {
			t.Errorf("data source %s schema type changed by mux", name)
		}
	}
}
