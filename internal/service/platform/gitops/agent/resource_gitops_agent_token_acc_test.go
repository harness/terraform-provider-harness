package agent_test

import (
	"context"
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/antihax/optional"
	hh "github.com/harness/harness-go-sdk/harness/helpers"
	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

const tokenAccResource = "harness_platform_gitops_agent.test"

func tokenAccAgentConfig(id, extra string) string {
	return fmt.Sprintf(`
		resource "harness_platform_gitops_agent" "test" {
			identifier = "%[1]s"
			name       = "%[1]s"
			type       = "MANAGED_ARGO_PROVIDER"
			operator   = "ARGO"
			metadata {
				namespace             = "terraform-test"
				high_availability     = false
				existing_installation = false
			}
			%[2]s
		}
	`, id, extra)
}

// serverKey returns the private key the server currently holds for the agent ("" once the agent connected).
func serverKey(state *terraform.State) (string, error) {
	r := acctest.TestAccGetResource(tokenAccResource, state)
	c, ctx := acctest.TestAccGetPlatformClientWithContext()
	ctx = context.WithValue(ctx, nextgen.ContextAccessToken, hh.EnvVars.BearerToken.Get())
	opts := &nextgen.AgentsApiAgentServiceForServerGetOpts{WithCredentials: optional.NewBool(true)}
	if org := r.Primary.Attributes["org_id"]; org != "" {
		opts.OrgIdentifier = optional.NewString(org)
	}
	if project := r.Primary.Attributes["project_id"]; project != "" {
		opts.ProjectIdentifier = optional.NewString(project)
	}
	agent, _, err := c.AgentApi.AgentServiceForServerGet(ctx, r.Primary.Attributes["identifier"], c.AccountId, opts)
	if err != nil {
		return "", err
	}
	if agent.Credentials == nil {
		return "", nil
	}
	return agent.Credentials.PrivateKey, nil
}

// stateNeverContainsServerKey fails if the key the server holds appears in any attribute of any resource in state.
func stateNeverContainsServerKey() resource.TestCheckFunc {
	return func(state *terraform.State) error {
		key, err := serverKey(state)
		if err != nil {
			return err
		}
		if key == "" {
			return fmt.Errorf("server returned no key; test needs an agent that has never connected")
		}
		for _, m := range state.Modules {
			for name, rs := range m.Resources {
				for attr, v := range rs.Primary.Attributes {
					if strings.Contains(v, key) {
						return fmt.Errorf("agent key found in state: %s.%s", name, attr)
					}
				}
			}
		}
		return nil
	}
}

// captureServerKey stores the current server-side key in *dst without exposing it.
func captureServerKey(dst *string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		k, err := serverKey(state)
		if err != nil {
			return err
		}
		*dst = k
		return nil
	}
}

func captureResourceID(name string, dst *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		*dst = s.RootModule().Resources[name].Primary.ID
		return nil
	}
}

func assertResourceID(name string, want *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if got := s.RootModule().Resources[name].Primary.ID; got != *want {
			return fmt.Errorf("agent was replaced: %s -> %s", *want, got)
		}
		return nil
	}
}

func TestAccResourceGitopsAgent_SkipTokenInState(t *testing.T) {
	id := strings.ReplaceAll(fmt.Sprintf("%s%s", "tfskip", utils.RandStringBytes(6)), "_", "")
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccResourceGitopsAgentDestroy(tokenAccResource),
		Steps: []resource.TestStep{
			{
				Config: tokenAccAgentConfig(id, "skip_agent_token_in_state = true"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(tokenAccResource, "agent_token", ""),
					resource.TestCheckResourceAttr(tokenAccResource, "token_generation", "0"),
					stateNeverContainsServerKey(),
				),
			},
			{
				// Re-plan with no change must be empty and keep the token out of state.
				Config: tokenAccAgentConfig(id, "skip_agent_token_in_state = true"),
				Check:  stateNeverContainsServerKey(),
			},
		},
	})
}

func TestAccResourceGitopsAgent_EnableSkipOnExistingAgent(t *testing.T) {
	id := strings.ReplaceAll(fmt.Sprintf("%s%s", "tfenable", utils.RandStringBytes(6)), "_", "")
	var firstID string
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccResourceGitopsAgentDestroy(tokenAccResource),
		Steps: []resource.TestStep{
			{
				// Default behavior is unchanged: the token is stored in state.
				Config: tokenAccAgentConfig(id, ""),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(tokenAccResource, "agent_token"),
					func(s *terraform.State) error {
						firstID = s.RootModule().Resources[tokenAccResource].Primary.ID
						return nil
					},
				),
			},
			{
				// Turning the flag on scrubs the token from state without replacing the agent.
				Config: tokenAccAgentConfig(id, "skip_agent_token_in_state = true"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(tokenAccResource, "agent_token", ""),
					func(s *terraform.State) error {
						if got := s.RootModule().Resources[tokenAccResource].Primary.ID; got != firstID {
							return fmt.Errorf("agent was replaced: %s -> %s", firstID, got)
						}
						return nil
					},
					stateNeverContainsServerKey(),
				),
			},
		},
	})
}

func TestAccResourceGitopsAgent_TokenRotation(t *testing.T) {
	id := strings.ReplaceAll(fmt.Sprintf("%s%s", "tfrot", utils.RandStringBytes(6)), "_", "")
	var before, after string
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccResourceGitopsAgentDestroy(tokenAccResource),
		Steps: []resource.TestStep{
			{
				Config: tokenAccAgentConfig(id, "skip_agent_token_in_state = true\n token_rotation = 0"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(tokenAccResource, "token_generation", "0"),
					captureServerKey(&before),
				),
			},
			{
				Config: tokenAccAgentConfig(id, "skip_agent_token_in_state = true\n token_rotation = 1"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(tokenAccResource, "token_generation", "1"),
					resource.TestCheckResourceAttr(tokenAccResource, "agent_token", ""),
					captureServerKey(&after),
					func(*terraform.State) error {
						if before == "" || after == "" {
							return fmt.Errorf("expected a server-side key before and after rotation")
						}
						if before == after {
							return fmt.Errorf("token_rotation did not issue a new key")
						}
						return nil
					},
					stateNeverContainsServerKey(),
				),
			},
			{
				// Same rotation value again: no further regeneration.
				Config: tokenAccAgentConfig(id, "skip_agent_token_in_state = true\n token_rotation = 1"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(tokenAccResource, "token_generation", "1"),
					func(s *terraform.State) error {
						k, err := serverKey(s)
						if err != nil {
							return err
						}
						if k != after {
							return fmt.Errorf("key changed without a token_rotation change")
						}
						return nil
					},
				),
			},
		},
	})
}

// Org/project scoped agents are the case CDS-136012 broke: rotation must reach the right scope on the server.
func TestAccResourceGitopsAgent_TokenRotationProjectLevel(t *testing.T) {
	id := strings.ReplaceAll(fmt.Sprintf("%s%s", "tfrotproj", utils.RandStringBytes(6)), "_", "")
	accountId := os.Getenv("HARNESS_ACCOUNT_ID")
	var before, after string
	cfg := func(extra string) string {
		return fmt.Sprintf(`
		resource "harness_platform_organization" "test" {
			identifier = "%[1]s"
			name       = "%[1]s"
		}
		resource "harness_platform_project" "test" {
			identifier = "%[1]s"
			name       = "%[1]s"
			org_id     = harness_platform_organization.test.id
		}
		resource "harness_platform_gitops_agent" "test" {
			identifier = "%[1]s"
			account_id = "%[2]s"
			name       = "%[1]s"
			project_id = harness_platform_project.test.id
			org_id     = harness_platform_organization.test.id
			type       = "MANAGED_ARGO_PROVIDER"
			operator   = "ARGO"
			metadata {
				namespace             = "terraform-test"
				high_availability     = false
				existing_installation = false
			}
			skip_agent_token_in_state = true
			%[3]s
		}
		`, id, accountId, extra)
	}
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccResourceGitopsAgentDestroy(tokenAccResource),
		Steps: []resource.TestStep{
			{
				Config: cfg("token_rotation = 0"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(tokenAccResource, "token_generation", "0"),
					captureServerKey(&before),
				),
			},
			{
				Config: cfg("token_rotation = 1"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(tokenAccResource, "token_generation", "1"),
					resource.TestCheckResourceAttr(tokenAccResource, "agent_token", ""),
					captureServerKey(&after),
					func(*terraform.State) error {
						if before == "" || after == "" || before == after {
							return fmt.Errorf("project-level rotation did not issue a new server-side key")
						}
						return nil
					},
					stateNeverContainsServerKey(),
				),
			},
		},
	})
}

func TestAccResourceGitopsAgent_TokenRotationStoresKeyWhenSkipIsOff(t *testing.T) {
	id := strings.ReplaceAll(fmt.Sprintf("%s%s", "tfrotstore", utils.RandStringBytes(6)), "_", "")
	var first string
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccResourceGitopsAgentDestroy(tokenAccResource),
		Steps: []resource.TestStep{
			{
				Config: tokenAccAgentConfig(id, "token_rotation = 0"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(tokenAccResource, "agent_token"),
					func(s *terraform.State) error {
						first = s.RootModule().Resources[tokenAccResource].Primary.Attributes["agent_token"]
						return nil
					},
				),
			},
			{
				Config: tokenAccAgentConfig(id, "token_rotation = 1"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(tokenAccResource, "token_generation", "1"),
					func(s *terraform.State) error {
						got := s.RootModule().Resources[tokenAccResource].Primary.Attributes["agent_token"]
						if got == "" || got == first {
							return fmt.Errorf("agent_token in state should be replaced by the new key")
						}
						return nil
					},
				),
			},
		},
	})
}

func TestAccResourceGitopsAgent_ImportWithSkip(t *testing.T) {
	id := strings.ReplaceAll(fmt.Sprintf("%s%s", "tfimp", utils.RandStringBytes(6)), "_", "")
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccResourceGitopsAgentDestroy(tokenAccResource),
		Steps: []resource.TestStep{
			{
				Config: tokenAccAgentConfig(id, "skip_agent_token_in_state = true"),
			},
			{
				ResourceName:            tokenAccResource,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateIdFunc:       acctest.ProjectResourceImportStateIdFunc(tokenAccResource),
				ImportStateVerifyIgnore: []string{"account_id", "agent_token", "skip_agent_token_in_state", "token_rotation", "token_generation"},
			},
		},
	})
}
