package acctest

import (
	"context"
	"os"
	"regexp"
	"strings"
	"testing"

	"github.com/harness/harness-go-sdk/harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/harness/terraform-provider-harness/internal/provider"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/stretchr/testify/require"
)

func TestProbeTemplateIdentity(t *testing.T) {
	valid := regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?$`)
	suffixPattern := regexp.MustCompile(`^[a-z0-9]{5}$`)
	cases := []struct {
		testName string
		prefix   string
	}{
		{"TestAccResourceProbeTemplate_httpProbe_basic", "testaccresourceprobetemplate-httpprobe-basic-"},
		{"TestAccResourceProbeTemplate_httpProbe_update", "testaccresourceprobetemplate-httpprobe-update-"},
		{"TestAccResourceProbeTemplate_cmdProbe_basic", "testaccresourceprobetemplate-cmdprobe-basic-"},
		{"TestAccResourceProbeTemplate_k8sProbe_basic", "testaccresourceprobetemplate-k8sprobe-basic-"},
		{"TestAccResourceProbeTemplate_apmProbe_prometheus", "testaccresourceprobetemplate-apmprobe-prometheus-"},
		{"TestAccResourceProbeTemplate_apmProbe_datadog", "testaccresourceprobetemplate-apmprobe-datadog-"},
		{"TestAccDataSourceProbeTemplate_byIdentity", "testaccdatasourceprobetemplate-byidentity-"},
		{"TestAccDataSourceProbeTemplate_byName", "testaccdatasourceprobetemplate-byname-"},
		{"TestAccDataSourceProbeTemplate_cmdProbe", "testaccdatasourceprobetemplate-cmdprobe-"},
		{"TestAccDataSourceProbeTemplate_k8sProbe", "testaccdatasourceprobetemplate-k8sprobe-"},
		{"TestAccResourceExperimentTemplate_nestedNoDrift", "testaccresourceexperimenttemplate-nestednodrift-"},
	}
	for _, tc := range cases {
		t.Run(tc.testName, func(t *testing.T) {
			id := ProbeTemplateIdentity(tc.testName)
			require.Regexp(t, valid, id)
			require.LessOrEqual(t, len(id), 63)
			require.True(t, strings.HasPrefix(id, tc.prefix), "identity %q", id)
			require.Regexp(t, suffixPattern, strings.TrimPrefix(id, tc.prefix))
			require.NotContains(t, id, "_")
			require.Equal(t, strings.ToLower(id), id)
		})
	}
}

func TestDefaultProvider_default_configuration(t *testing.T) {
	// Setup
	p := provider.Provider("dev")()

	// verify
	// diag := p.Configure(context.TODO(), nil)
	// require.False(t, diag.HasError())
	require.NoError(t, p.InternalValidate())
}

func TestProvider_configure_url(t *testing.T) {

	// Setup provider
	const expectedEndpoint = "http://localhost:8200"
	rc := terraform.NewResourceConfigRaw(map[string]interface{}{"endpoint": expectedEndpoint})
	p := provider.Provider("dev")()
	diags := p.Configure(context.TODO(), rc)

	// Verify
	require.False(t, diags.HasError())
	require.NoError(t, p.InternalValidate())
	c := p.Meta().(*internal.Session)
	require.Equal(t, expectedEndpoint, c.Endpoint)
}

func TestProvider_configure_url_env(t *testing.T) {

	// Setup
	const expectedEndpoint = "http://localhost:8200"

	// Cleanup function
	defer func() {
		os.Unsetenv(helpers.EnvVars.Endpoint.String())
	}()

	// Configure environment
	os.Setenv(helpers.EnvVars.Endpoint.String(), expectedEndpoint)

	// Create provider
	rc := terraform.NewResourceConfigRaw(map[string]interface{}{})
	p := provider.Provider("dev")()
	diags := p.Configure(context.TODO(), rc)

	// Verify
	require.False(t, diags.HasError())
	require.NoError(t, p.InternalValidate())
	c := p.Meta().(*internal.Session)
	require.Equal(t, expectedEndpoint, c.Endpoint)
}
