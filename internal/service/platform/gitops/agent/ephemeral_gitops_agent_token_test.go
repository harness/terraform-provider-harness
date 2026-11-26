package agent

import (
	"context"
	"net/http"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/ephemeral"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"github.com/stretchr/testify/require"
)

func openEphemeral(t *testing.T, f *fakeGitops, id string, generation interface{}) (ephemeral.OpenResponse, string) {
	t.Helper()
	ctx := context.Background()
	e := &ephemeralGitopsAgentToken{session: f.session()}

	var sr ephemeral.SchemaResponse
	e.Schema(ctx, ephemeral.SchemaRequest{}, &sr)
	require.False(t, sr.Diagnostics.HasError())

	typ := sr.Schema.Type().TerraformType(ctx)
	cfg := tfsdk.Config{
		Schema: sr.Schema,
		Raw: tftypes.NewValue(typ, map[string]tftypes.Value{
			"identifier":  tftypes.NewValue(tftypes.String, id),
			"org_id":      tftypes.NewValue(tftypes.String, "default"),
			"project_id":  tftypes.NewValue(tftypes.String, "proj"),
			"generation":  tftypes.NewValue(tftypes.Number, generation),
			"agent_token": tftypes.NewValue(tftypes.String, nil),
		}),
	}
	resp := ephemeral.OpenResponse{Result: tfsdk.EphemeralResultData{Schema: sr.Schema, Raw: tftypes.NewValue(typ, nil)}}
	e.Open(ctx, ephemeral.OpenRequest{Config: cfg}, &resp)

	var token string
	if !resp.Diagnostics.HasError() {
		resp.Diagnostics.Append(resp.Result.GetAttribute(ctx, path.Root("agent_token"), &token)...)
	}
	return resp, token
}

func TestEphemeralOpen_ReturnsTokenFromAgent(t *testing.T) {
	f := newFakeGitops(t)
	resp, token := openEphemeral(t, f, "a1", nil)
	require.False(t, resp.Diagnostics.HasError(), "%v", resp.Diagnostics)
	require.Equal(t, "GETKEY", token)
	require.Equal(t, []string{"true"}, f.requests[0].Query["withCredentials"])
	require.Equal(t, []string{"proj"}, f.requests[0].Query["projectIdentifier"])
}

func TestEphemeralOpen_GenerationIsOrderingOnly(t *testing.T) {
	f := newFakeGitops(t)
	resp, token := openEphemeral(t, f, "a1", 3)
	require.False(t, resp.Diagnostics.HasError(), "%v", resp.Diagnostics)
	require.Equal(t, "GETKEY", token)
	require.Equal(t, 1, f.count(http.MethodGet, "/gitops/api/v1/agents/a1"))
}

func TestEphemeralOpen_ErrorsWhenAgentAlreadyConnected(t *testing.T) {
	f := newFakeGitops(t)
	f.getBody = `{"identifier":"a1"}`
	resp, _ := openEphemeral(t, f, "a1", nil)
	require.True(t, resp.Diagnostics.HasError())
	require.Equal(t, "Agent token not available", resp.Diagnostics[0].Summary())
}

func TestEphemeralOpen_ErrorsWhenAgentMissing(t *testing.T) {
	f := newFakeGitops(t)
	resp, _ := openEphemeral(t, f, "does-not-exist", nil)
	require.True(t, resp.Diagnostics.HasError())
	require.Equal(t, "GitOps agent not found", resp.Diagnostics[0].Summary())
}

func TestEphemeralSchema_GenerationIsOptionalAndTokenIsSensitive(t *testing.T) {
	var sr ephemeral.SchemaResponse
	(&ephemeralGitopsAgentToken{}).Schema(context.Background(), ephemeral.SchemaRequest{}, &sr)
	gen, ok := sr.Schema.Attributes["generation"]
	require.True(t, ok)
	require.True(t, gen.IsOptional())
	tok := sr.Schema.Attributes["agent_token"]
	require.True(t, tok.IsSensitive())
	require.True(t, tok.IsComputed())
}
