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
