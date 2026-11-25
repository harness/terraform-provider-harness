package agent_test

import (
	"context"
	"fmt"
	"os"
	"regexp"
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

// Create standard agent omitting existing_installation (backward compat)
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
				Config: testAccResourceGitopsAgentAccountLevel(id, accountId, agentName, namespace, "false"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", agentName),
					resource.TestCheckResourceAttrSet(resourceName, "agent_token"),
					resource.TestCheckResourceAttr(resourceName, "metadata.0.existing_installation", "false"),
				),
			},
			{
				Config: testAccResourceGitopsAgentAccountLevel(id, accountId, agentName, updatedNamespace, "false"),
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

func TestAccResourceGitopsAgentNS(t *testing.T) {
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
				Config: testAccResourceGitopsAgentAccountLevel(id, accountId, agentName, namespace, "true"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", agentName),
					resource.TestCheckResourceAttrSet(resourceName, "agent_token"),
					resource.TestCheckResourceAttr(resourceName, "metadata.0.is_namespaced", "true"),
				),
			},
			{
				Config: testAccResourceGitopsAgentAccountLevel(id, accountId, agentName, updatedNamespace, "true"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "metadata.0.namespace", updatedNamespace),
					resource.TestCheckResourceAttr(resourceName, "metadata.0.is_namespaced", "true"),
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

// Create BYOA agent with existing_installation = true
// No perpetual drift on BYOA agent after apply
// Destroy BYOA agent via Terraform
func TestAccResourceGitopsAgentBYOA(t *testing.T) {
	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	id = strings.ReplaceAll(id, "_", "")
	accountId := os.Getenv("HARNESS_ACCOUNT_ID")
	resourceName := "harness_platform_gitops_agent.test"
	agentName := id
	namespace := "terraform-test"
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccResourceGitopsAgentDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceGitopsAgentBYOAAccountLevel(id, accountId, agentName, namespace, true),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", agentName),
					resource.TestCheckResourceAttr(resourceName, "metadata.0.existing_installation", "true"),
					resource.TestCheckResourceAttrSet(resourceName, "agent_token"),
				),
			},
			{
				Config:   testAccResourceGitopsAgentBYOAAccountLevel(id, accountId, agentName, namespace, true),
				PlanOnly: true,
			},
			{
				Config:      testAccResourceGitopsAgentBYOAAccountLevel(id, accountId, agentName, namespace, false),
				ExpectError: regexp.MustCompile(`field 'metadata.existing_installation' cannot be changed after the agent is created`),
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

// Create BYOA agent at org scope
// Create BYOA agent at project scope
func TestAccResourceGitopsAgentBYOAScopes(t *testing.T) {
	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	id = strings.ReplaceAll(id, "_", "")
	accountId := os.Getenv("HARNESS_ACCOUNT_ID")
	orgResourceName := "harness_platform_gitops_agent.org"
	projectResourceName := "harness_platform_gitops_agent.project"
	namespace := "terraform-test"
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy: func(s *terraform.State) error {
			if err := testAccResourceGitopsAgentDestroy(orgResourceName)(s); err != nil {
				return err
			}
			return testAccResourceGitopsAgentDestroy(projectResourceName)(s)
		},
		Steps: []resource.TestStep{
			{
				Config: testAccResourceGitopsAgentBYOAScopes(id, accountId, namespace),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(orgResourceName, "metadata.0.existing_installation", "true"),
					resource.TestCheckResourceAttr(projectResourceName, "metadata.0.existing_installation", "true"),
				),
			},
			{
				Config:   testAccResourceGitopsAgentBYOAScopes(id, accountId, namespace),
				PlanOnly: true,
			},
		},
	})
}

// type = HOSTED_ARGO_PROVIDER is accepted by the provider
func TestAccResourceGitopsAgentHostedType(t *testing.T) {
	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	id = strings.ReplaceAll(id, "_", "")
	accountId := os.Getenv("HARNESS_ACCOUNT_ID")
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:             testAccResourceGitopsAgentHostedType(id, accountId),
				PlanOnly:           true,
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

// Deleting a live PRIMARY while a SECONDARY sibling still exists must be rejected by the
// provider's Delete guard instead of silently orphaning the DR group.
func TestAccResourceGitopsAgentDrDeleteGuard(t *testing.T) {
	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	id = strings.ReplaceAll(id, "_", "")
	accountId := os.Getenv("HARNESS_ACCOUNT_ID")
	primaryResourceName := "harness_platform_gitops_agent.primary"
	secondaryResourceName := "harness_platform_gitops_agent.secondary"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy: func(s *terraform.State) error {
			if err := testAccResourceGitopsAgentDestroy(primaryResourceName)(s); err != nil {
				return err
			}
			return testAccResourceGitopsAgentDestroy(secondaryResourceName)(s)
		},
		Steps: []resource.TestStep{
			{
				Config: testAccResourceGitopsAgentDrPrimarySecondary(id, accountId),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(primaryResourceName, "disaster_recovery_live_type", "PRIMARY"),
					resource.TestCheckResourceAttr(secondaryResourceName, "disaster_recovery_live_type", "SECONDARY"),
				),
			},
			{
				// Removing the PRIMARY resource from config while its SECONDARY sibling is still
				// live must fail at Delete, not succeed and orphan the DR group.
				Config:      testAccResourceGitopsAgentDrSecondaryOnly(id, accountId),
				ExpectError: regexp.MustCompile(`cannot delete live PRIMARY node`),
			},
			{
				// Restore the dependency edge between primary and secondary so the framework's
				// final teardown destroys the SECONDARY before the PRIMARY, which the guard requires.
				Config: testAccResourceGitopsAgentDrPrimarySecondary(id, accountId),
			},
		},
	})
}

// Changing token_rotation on a DR-enrolled agent must be rejected at plan time: credential
// rotation via RegenerateCredentials is not DR-aware.
func TestAccResourceGitopsAgentDrTokenRotationGuard(t *testing.T) {
	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	id = strings.ReplaceAll(id, "_", "")
	accountId := os.Getenv("HARNESS_ACCOUNT_ID")
	resourceName := "harness_platform_gitops_agent.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccResourceGitopsAgentDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceGitopsAgentDrTokenRotation(id, accountId, 0),
			},
			{
				Config:      testAccResourceGitopsAgentDrTokenRotation(id, accountId, 1),
				ExpectError: regexp.MustCompile(`field 'token_rotation' cannot be changed while 'disaster_recovery_identifier' is set`),
			},
		},
	})
}

// DR siblings share one agent identifier and are distinguished by a per-node disaster_recovery_identifier.
func testAccResourceGitopsAgentDrPrimarySecondary(id string, accountId string) string {
	return fmt.Sprintf(`
		resource "harness_platform_gitops_agent" "primary" {
			identifier = "%[1]s"
			account_id = "%[2]s"
			name = "%[1]s"
			type = "MANAGED_ARGO_PROVIDER"
			metadata {
				namespace = "terraform-test"
				high_availability = false
			}
			operator = "ARGO"
			disaster_recovery_identifier = "%[1]sa"
			disaster_recovery_type = "PRIMARY"
		}

		resource "harness_platform_gitops_agent" "secondary" {
			identifier = harness_platform_gitops_agent.primary.identifier
			account_id = "%[2]s"
			name = harness_platform_gitops_agent.primary.name
			type = harness_platform_gitops_agent.primary.type
			metadata {
				namespace = harness_platform_gitops_agent.primary.metadata[0].namespace
				high_availability = false
			}
			operator = "ARGO"
			disaster_recovery_identifier = "%[1]sb"
			disaster_recovery_type = "SECONDARY"
			depends_on = [harness_platform_gitops_agent.primary]
		}
		`, id, accountId)
}

func testAccResourceGitopsAgentDrSecondaryOnly(id string, accountId string) string {
	return fmt.Sprintf(`
		resource "harness_platform_gitops_agent" "secondary" {
			identifier = "%[1]s"
			account_id = "%[2]s"
			name = "%[1]s"
			type = "MANAGED_ARGO_PROVIDER"
			metadata {
				namespace = "terraform-test"
				high_availability = false
			}
			operator = "ARGO"
			disaster_recovery_identifier = "%[1]sb"
			disaster_recovery_type = "SECONDARY"
		}
		`, id, accountId)
}

func testAccResourceGitopsAgentDrTokenRotation(id string, accountId string, tokenRotation int) string {
	return fmt.Sprintf(`
		resource "harness_platform_gitops_agent" "test" {
			identifier = "%[1]s"
			account_id = "%[2]s"
			name = "%[1]s"
			type = "MANAGED_ARGO_PROVIDER"
			metadata {
				namespace = "terraform-test"
				high_availability = false
			}
			operator = "ARGO"
			disaster_recovery_identifier = "%[1]sa"
			disaster_recovery_type = "PRIMARY"
			token_rotation = %[3]d
		}
		`, id, accountId, tokenRotation)
}

func testAccGetAgent(resourceName string, state *terraform.State) (*nextgen.V1Agent, error) {
	r := acctest.TestAccGetResource(resourceName, state)
	c, ctx := acctest.TestAccGetPlatformClientWithContext()
	ctx = context.WithValue(ctx, nextgen.ContextAccessToken, hh.EnvVars.BearerToken.Get())
	agentIdentifier := r.Primary.Attributes["identifier"]

	opts := &nextgen.AgentsApiAgentServiceForServerGetOpts{
		OrgIdentifier:     optional.NewString(r.Primary.Attributes["org_identifier"]),
		ProjectIdentifier: optional.NewString(r.Primary.Attributes["project_identifier"]),
	}
	// Without a DR identifier the server only matches non-DR records, so DR nodes would never be found.
	if drIdentifier := r.Primary.Attributes["disaster_recovery_identifier"]; drIdentifier != "" {
		opts.DrIdentifier = optional.NewString(drIdentifier)
	}
	resp, _, err := c.AgentApi.AgentServiceForServerGet(ctx, agentIdentifier, c.AccountId, opts)

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

func testAccResourceGitopsAgentBYOAScopes(id string, accountId string, namespace string) string {
	return fmt.Sprintf(`
		resource "harness_platform_organization" "test" {
			identifier = "%[1]s"
			name = "%[1]s"
		}

		resource "harness_platform_project" "test" {
			identifier = "%[1]s"
			name = "%[1]s"
			org_id = harness_platform_organization.test.id
		}

		resource "harness_platform_gitops_agent" "org" {
			identifier = "%[1]sorg"
			account_id = "%[2]s"
			org_id = harness_platform_organization.test.id
			name = "%[1]sorg"
			type = "MANAGED_ARGO_PROVIDER"
			metadata {
				namespace = "%[3]s"
				high_availability = false
				existing_installation = true
			}
			operator = "ARGO"
		}

		resource "harness_platform_gitops_agent" "project" {
			identifier = "%[1]sproj"
			account_id = "%[2]s"
			org_id = harness_platform_organization.test.id
			project_id = harness_platform_project.test.id
			name = "%[1]sproj"
			type = "MANAGED_ARGO_PROVIDER"
			metadata {
				namespace = "%[3]s"
				high_availability = false
				existing_installation = true
			}
			operator = "ARGO"
		}
		`, id, accountId, namespace)
}

func testAccResourceGitopsAgentHostedType(agentId string, accountId string) string {
	return fmt.Sprintf(`
		resource "harness_platform_gitops_agent" "test" {
			identifier = "%[1]s"
			account_id = "%[2]s"
			name = "%[1]s"
			type = "HOSTED_ARGO_PROVIDER"
			metadata {
				namespace = "terraform-test"
				high_availability = false
				existing_installation = true
			}
			operator = "ARGO"
		}
		`, agentId, accountId)
}

func testAccResourceGitopsAgentBYOAAccountLevel(agentId string, accountId string, agentName string, namespace string, existingInstallation bool) string {
	return fmt.Sprintf(`
		resource "harness_platform_gitops_agent" "test" {
			identifier = "%[1]s"
			account_id = "%[2]s"
			name = "%[3]s"
			type = "MANAGED_ARGO_PROVIDER"
			metadata {
				namespace = "%[4]s"
        		high_availability = false
				existing_installation = %[5]t
    		}
			operator = "ARGO"		
		}
		`, agentId, accountId, agentName, namespace, existingInstallation)
}

func testAccResourceGitopsAgentAccountLevel(agentId string, accountId string, agentName string, namespace string, isNamespaced string) string {
	return fmt.Sprintf(`
		resource "harness_platform_gitops_agent" "test" {
			identifier = "%[1]s"
			account_id = "%[2]s"
			name = "%[3]s"
			type = "MANAGED_ARGO_PROVIDER"
			metadata {
				namespace = "%[4]s"
        		high_availability = false
				is_namespaced = %[5]s
    		}
			operator = "ARGO"		
		}
		`, agentId, accountId, agentName, namespace, isNamespaced)
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
			operator = "ARGO"		
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
