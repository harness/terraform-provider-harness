package agent_test

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/harness/terraform-provider-harness/internal/provider"
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5"
	"github.com/hashicorp/terraform-plugin-mux/tf5muxserver"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func muxedProviderFactories() map[string]func() (tfprotov5.ProviderServer, error) {
	return map[string]func() (tfprotov5.ProviderServer, error){
		"echo": echoProviderServer,
		"harness": func() (tfprotov5.ProviderServer, error) {
			mux, err := tf5muxserver.NewMuxServer(context.Background(),
				provider.Provider("dev")().GRPCProvider,
				providerserver.NewProtocol5(provider.NewFrameworkProvider("dev")()),
			)
			if err != nil {
				return nil, err
			}
			return mux.ProviderServer(), nil
		},
	}
}

// The ephemeral resource must read the token after the agent exists (create) and after a rotation, using
// the real server. hashicorp/echo is used to prove a value arrived without the harness resource storing it.
func TestAccEphemeralGitopsAgentToken_ReadAfterCreateAndRotation(t *testing.T) {
	id := strings.ReplaceAll(fmt.Sprintf("%s%s", "tfeph", utils.RandStringBytes(6)), "_", "")
	var first, second string
	cfg := func(rotation int) string {
		return tokenAccAgentConfig(id, fmt.Sprintf("skip_agent_token_in_state = true\n token_rotation = %d", rotation)) + `
			ephemeral "harness_platform_gitops_agent_token" "t" {
				identifier = harness_platform_gitops_agent.test.identifier
				generation = harness_platform_gitops_agent.test.token_generation
			}
			provider "echo" {
				data = ephemeral.harness_platform_gitops_agent_token.t.agent_token
			}
			resource "echo" "token" {
				trigger = tostring(harness_platform_gitops_agent.test.token_generation)
			}
		`
	}
	resource.UnitTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV5ProviderFactories: muxedProviderFactories(),
		CheckDestroy:             testAccResourceGitopsAgentDestroy(tokenAccResource),
		Steps: []resource.TestStep{
			{
				Config: cfg(0),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(tokenAccResource, "agent_token", ""),
					resource.TestCheckResourceAttrSet("echo.token", "data"),
					func(s *terraform.State) error {
						first = s.RootModule().Resources["echo.token"].Primary.Attributes["data"]
						k, err := serverKey(s)
						if err != nil {
							return err
						}
						if first != k {
							return fmt.Errorf("ephemeral token does not match the server key after create")
						}
						return nil
					},
				),
			},
			{
				Config: cfg(1),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(tokenAccResource, "token_generation", "1"),
					func(s *terraform.State) error {
						second = s.RootModule().Resources["echo.token"].Primary.Attributes["data"]
						k, err := serverKey(s)
						if err != nil {
							return err
						}
						if second == first {
							return fmt.Errorf("ephemeral token was not re-read after rotation")
						}
						if second != k {
							return fmt.Errorf("ephemeral token does not match the rotated server key")
						}
						return nil
					},
				),
			},
		},
	})
}
