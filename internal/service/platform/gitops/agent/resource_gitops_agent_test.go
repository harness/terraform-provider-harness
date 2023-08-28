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

func TestAccResourceGitopsAgent(t *testing.T) {
	// Account Level
	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	id = strings.ReplaceAll(id, "_", "")
	accountId := os.Getenv("HARNESS_ACCOUNT_ID")
	resourceName := "harness_platform_gitops_agent.test"
	agentName := id
	namespace := "terraform-test"
	updatedNamespace := namespace + "-updated"
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccResourceGitopsAgentDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceGitopsAgentAccountLevel(id, accountId, agentName, namespace),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", agentName),
					resource.TestCheckResourceAttrSet(resourceName, "agent_token"),
				),
			},
			{
				Config: testAccResourceGitopsAgentAccountLevel(id, accountId, agentName, updatedNamespace),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "metadata.0.namespace", updatedNamespace),
					resource.TestCheckResourceAttrSet(resourceName, "agent_token"),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"account_id", "agent_token"},
				ImportStateIdFunc:       acctest.ProjectResourceImportStateIdFunc(resourceName),
			},
		},
	})

	//Project level
	id = fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	id = strings.ReplaceAll(id, "_", "")
	resourceName = "harness_platform_gitops_agent.test"
	agentName = id
	namespace = "terraform-test"
	updatedNamespace = namespace + "-updated"
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccResourceGitopsAgentDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceGitopsAgentProjectLevel(id, accountId, agentName, namespace),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", agentName),
					resource.TestCheckResourceAttrSet(resourceName, "agent_token"),
				),
			},
			{
				Config: testAccResourceGitopsAgentProjectLevel(id, accountId, agentName, updatedNamespace),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "metadata.0.namespace", updatedNamespace),
					resource.TestCheckResourceAttrSet(resourceName, "agent_token"),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"account_id", "agent_token"},
				ImportStateIdFunc:       acctest.ProjectResourceImportStateIdFunc(resourceName),
			},
		},
	})
}

// FLAMINGO or FLUX
func TestAccResourceGitopsAgentFlamingo(t *testing.T) {
	// Account Level
	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	id = strings.ReplaceAll(id, "_", "")
	accountId := os.Getenv("HARNESS_ACCOUNT_ID")
	resourceName := "harness_platform_gitops_agent.test"
	agentName := id
	namespace := "terraform-test"
	updatedNamespace := namespace + "-updated"
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccResourceGitopsAgentDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceGitopsFluxAgentAccountLevel(id, accountId, agentName, namespace),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", agentName),
					resource.TestCheckResourceAttrSet(resourceName, "agent_token"),
				),
			},
			{
				Config: testAccResourceGitopsFluxAgentAccountLevel(id, accountId, agentName, updatedNamespace),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "metadata.0.namespace", updatedNamespace),
					resource.TestCheckResourceAttrSet(resourceName, "agent_token"),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"account_id", "agent_token"},
				ImportStateIdFunc:       acctest.ProjectResourceImportStateIdFunc(resourceName),
			},
		},
	})

	//Project level
	id = fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	id = strings.ReplaceAll(id, "_", "")
	resourceName = "harness_platform_gitops_agent.test"
	agentName = id
	namespace = "terraform-test"
	updatedNamespace = namespace + "-updated"
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccResourceGitopsAgentDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceGitopsAgentFluxProjectLevel(id, accountId, agentName, namespace),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", agentName),
					resource.TestCheckResourceAttrSet(resourceName, "agent_token"),
				),
			},
			{
				Config: testAccResourceGitopsAgentFluxProjectLevel(id, accountId, agentName, updatedNamespace),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "metadata.0.namespace", updatedNamespace),
					resource.TestCheckResourceAttrSet(resourceName, "agent_token"),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"account_id", "agent_token"},
				ImportStateIdFunc:       acctest.ProjectResourceImportStateIdFunc(resourceName),
			},
		},
	})
}

func testAccGetAgent(resourceName string, state *terraform.State) (*nextgen.V1Agent, error) {
	r := acctest.TestAccGetResource(resourceName, state)
	c, ctx := acctest.TestAccGetPlatformClientWithContext()
	ctx = context.WithValue(ctx, nextgen.ContextAccessToken, hh.EnvVars.BearerToken.Get())
	agentIdentifier := r.Primary.Attributes["identifier"]

	resp, _, err := c.AgentApi.AgentServiceForServerGet(ctx, agentIdentifier, c.AccountId, &nextgen.AgentsApiAgentServiceForServerGetOpts{
		OrgIdentifier:     optional.NewString(r.Primary.Attributes["org_identifier"]),
		ProjectIdentifier: optional.NewString(r.Primary.Attributes["project_identifier"]),
	})

	if err != nil {
		return nil, err
	}

	if resp.Type_ == nil {
		return nil, nil
	}

	return &resp, nil
}

func testAccResourceGitopsAgentDestroy(resourceName string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		agent, _ := testAccGetAgent(resourceName, state)
		if agent != nil {
			return fmt.Errorf("Found Agent: %s", agent.Identifier)
		}
		return nil
	}

}

func testAccResourceGitopsAgentAccountLevel(agentId string, accountId string, agentName string, namespace string) string {
	return fmt.Sprintf(`
		resource "harness_platform_gitops_agent" "test" {
			identifier = "%[1]s"
			account_id = "%[2]s"
			name = "%[3]s"
			type = "MANAGED_ARGO_PROVIDER"
			metadata {
				namespace = "%[4]s"
        		high_availability = false
    		}
		}
		`, agentId, accountId, agentName, namespace)
}
func testAccResourceGitopsAgentProjectLevel(agentId string, accountId string, agentName string, namespace string) string {
	return fmt.Sprintf(`
		resource "harness_platform_organization" "test" {
			identifier = "%[1]s"
			name = "%[3]s"
		}

		resource "harness_platform_project" "test" {
			identifier = "%[1]s"
			name = "%[3]s"
			org_id = harness_platform_organization.test.id
		}
		resource "harness_platform_gitops_agent" "test" {
			identifier = "%[1]s"
			account_id = "%[2]s"
			project_id = harness_platform_project.test.id
			org_id = harness_platform_organization.test.id
			name = "%[3]s"
			type = "MANAGED_ARGO_PROVIDER"
			metadata {
				namespace = "%[4]s"
        		high_availability = false
    		}
		}
		`, agentId, accountId, agentName, namespace)
}

func testAccResourceGitopsFluxAgentAccountLevel(agentId string, accountId string, agentName string, namespace string) string {
	return fmt.Sprintf(`
		resource "harness_platform_gitops_agent" "test" {
			identifier = "%[1]s"
			account_id = "%[2]s"
			name = "%[3]s"
			type = "MANAGED_ARGO_PROVIDER"
			metadata {
				namespace = "%[4]s"
        		high_availability = false
    		}
			operator = "FLAMINGO"
		}
		`, agentId, accountId, agentName, namespace)
}

func testAccResourceGitopsAgentFluxProjectLevel(agentId string, accountId string, agentName string, namespace string) string {
	return fmt.Sprintf(`
		resource "harness_platform_organization" "test" {
			identifier = "%[1]s"
			name = "%[3]s"
		}

		resource "harness_platform_project" "test" {
			identifier = "%[1]s"
			name = "%[3]s"
			org_id = harness_platform_organization.test.id
		}
		resource "harness_platform_gitops_agent" "test" {
			identifier = "%[1]s"
			account_id = "%[2]s"
			project_id = harness_platform_project.test.id
			org_id = harness_platform_organization.test.id
			name = "%[3]s"
			type = "MANAGED_ARGO_PROVIDER"
			metadata {
				namespace = "%[4]s"
        		high_availability = false
    		}
			operator = "FLAMINGO"	
		}
		`, agentId, accountId, agentName, namespace)
}
