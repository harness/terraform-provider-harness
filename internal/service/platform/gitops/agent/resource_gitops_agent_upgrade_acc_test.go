package agent_test

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

// lastReleasedProvider is the newest published provider that predates the new attributes.
const lastReleasedProvider = "0.47.5"

func releasedProvider() map[string]resource.ExternalProvider {
	return map[string]resource.ExternalProvider{
		"harness": {VersionConstraint: lastReleasedProvider, Source: "harness/harness"},
	}
}

// Every existing agent configuration is applied with the last released provider and then planned with the
// provider under test. The plan must be empty, and the stored token must still be there.
func TestAccResourceGitopsAgent_UpgradeFromReleasedProvider(t *testing.T) {
	accountId := os.Getenv("HARNESS_ACCOUNT_ID")
	ns := "terraform-test"

	cases := []struct {
		name   string
		config func(id string) string
		// resource that carries the agent token, "" when the scenario has no single such resource
		resourceName string
	}{
		{"account-level", func(id string) string {
			return testAccResourceGitopsAgentAccountLevel(id, accountId, id, ns, "false")
		}, "harness_platform_gitops_agent.test"},
		{"account-level-namespaced", func(id string) string {
			return testAccResourceGitopsAgentAccountLevel(id, accountId, id, ns, "true")
		}, "harness_platform_gitops_agent.test"},
		{"project-level", func(id string) string {
			return testAccResourceGitopsAgentProjectLevel(id, accountId, id, ns)
		}, "harness_platform_gitops_agent.test"},
		{"byoa-account-level", func(id string) string {
			return testAccResourceGitopsAgentBYOAAccountLevel(id, accountId, id, ns, true)
		}, "harness_platform_gitops_agent.test"},
		{"byoa-org-and-project", func(id string) string {
			return testAccResourceGitopsAgentBYOAScopes(id, accountId, ns)
		}, "harness_platform_gitops_agent.org"},
		{"flamingo-account-level", func(id string) string {
			return testAccResourceGitopsFluxAgentAccountLevel(id, accountId, id, ns)
		}, "harness_platform_gitops_agent.test"},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			id := strings.ReplaceAll(fmt.Sprintf("%s%s", "tfup", utils.RandStringBytes(6)), "_", "")
			cfg := tc.config(id)
			resource.UnitTest(t, resource.TestCase{
				PreCheck: func() { acctest.TestAccPreCheck(t) },
				Steps: []resource.TestStep{
					{
						ExternalProviders: releasedProvider(),
						Config:            cfg,
						Check:             resource.TestCheckResourceAttrSet(tc.resourceName, "agent_token"),
					},
					{
						ProviderFactories: acctest.ProviderFactories,
						Config:            cfg,
						PlanOnly:          true,
					},
					{
						ProviderFactories: acctest.ProviderFactories,
						Config:            cfg,
						Check: resource.ComposeTestCheckFunc(
							resource.TestCheckResourceAttrSet(tc.resourceName, "agent_token"),
							resource.TestCheckResourceAttr(tc.resourceName, "token_generation", "0"),
						),
					},
				},
			})
		})
	}
}

// An agent created by the released provider (token stored in state) can opt in to skip_agent_token_in_state
// and then rotate, without being replaced.
func TestAccResourceGitopsAgent_UpgradeThenOptIn(t *testing.T) {
	id := strings.ReplaceAll(fmt.Sprintf("%s%s", "tfupopt", utils.RandStringBytes(6)), "_", "")
	var firstID string
	resource.UnitTest(t, resource.TestCase{
		PreCheck: func() { acctest.TestAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				ExternalProviders: releasedProvider(),
				Config:            tokenAccAgentConfig(id, ""),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(tokenAccResource, "agent_token"),
					captureResourceID(tokenAccResource, &firstID),
				),
			},
			{
				ProviderFactories: acctest.ProviderFactories,
				Config:            tokenAccAgentConfig(id, "skip_agent_token_in_state = true"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(tokenAccResource, "agent_token", ""),
					assertResourceID(tokenAccResource, &firstID),
					stateNeverContainsServerKey(),
				),
			},
			{
				ProviderFactories: acctest.ProviderFactories,
				Config:            tokenAccAgentConfig(id, "skip_agent_token_in_state = true\n token_rotation = 1"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(tokenAccResource, "token_generation", "1"),
					assertResourceID(tokenAccResource, &firstID),
					stateNeverContainsServerKey(),
				),
			},
		},
	})
}
