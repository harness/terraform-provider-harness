package repo_rule_branch_test

import (
	"fmt"
	"testing"

	"github.com/antihax/optional"
	"github.com/harness/harness-go-sdk/harness/code"
	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/stretchr/testify/require"
)

const (
	resourceName       = "harness_platform_repo_rule_branch.test"
	description        = "example_description"
	updatedDescription = "example_description_updated"
)

func TestProjResourceRepoRule(t *testing.T) {
	identifier := identifier(t.Name())

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testRepoRuleDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testProjResourceRepoRule(identifier, description),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", "rule_"+identifier),
					resource.TestCheckResourceAttr(resourceName, "description", description),
					resource.TestCheckResourceAttr(resourceName, "policies.0.require_pull_request", "true"),
					resource.TestCheckResourceAttr(resourceName, "policies.0.require_status_checks.0", "check-stage1"),
				),
			},
			{
				Config: testProjResourceRepoRule(identifier, updatedDescription),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", "rule_"+identifier),
					resource.TestCheckResourceAttr(resourceName, "description", updatedDescription),
					resource.TestCheckResourceAttr(resourceName, "policies.0.require_pull_request", "true"),
					resource.TestCheckResourceAttr(resourceName, "policies.0.require_status_checks.0", "check-stage1"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: false,
				ImportStateIdFunc: acctest.RepoRuleProjectLevelResourceImportStateIdFunc(resourceName),
			},
		},
	})
}

// TestProjResourceRepoRule_DeleteUnderlyingResource verifies that a branch rule deleted
// outside Terraform is dropped from state on refresh instead of failing with a 404.
func TestProjResourceRepoRule_DeleteUnderlyingResource(t *testing.T) {
	identifier := identifier(t.Name())

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testRepoRuleDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testProjResourceRepoRule(identifier, description),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", "rule_"+identifier),
				),
			},
			{
				PreConfig: func() {
					acctest.TestAccConfigureProvider()
					c, ctx := acctest.TestAccGetCodeClientWithContext()
					_, err := c.RepositoryApi.RuleDelete(
						ctx, c.AccountId, identifier, "rule_"+identifier,
						&code.RepositoryApiRuleDeleteOpts{
							OrgIdentifier:     optional.NewString("org_" + identifier),
							ProjectIdentifier: optional.NewString("proj_" + identifier),
						})
					require.NoError(t, err)
				},
				Config:             testProjResourceRepoRule(identifier, description),
				PlanOnly:           true,
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

// TestProjResourceRepoRule_DeleteParentRepo verifies that a branch rule whose parent
// repo was deleted outside Terraform is dropped from state on refresh instead of
// failing with a 404. This is the destroy-retry scenario from CODE-6342.
func TestProjResourceRepoRule_DeleteParentRepo(t *testing.T) {
	identifier := identifier(t.Name())

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testRepoRuleDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testProjResourceRepoRule(identifier, description),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", "rule_"+identifier),
				),
			},
			{
				PreConfig: func() {
					acctest.TestAccConfigureProvider()
					c, ctx := acctest.TestAccGetCodeClientWithContext()
					_, _, err := c.RepositoryApi.DeleteRepository(
						ctx, c.AccountId, identifier,
						&code.RepositoryApiDeleteRepositoryOpts{
							OrgIdentifier:     optional.NewString("org_" + identifier),
							ProjectIdentifier: optional.NewString("proj_" + identifier),
						})
					require.NoError(t, err)
				},
				Config:             testProjResourceRepoRule(identifier, description),
				PlanOnly:           true,
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func identifier(testName string) string {
	return fmt.Sprintf("%s_%s", testName, utils.RandStringBytes(5))
}

func testProjResourceRepoRule(identifier, description string) string {
	return fmt.Sprintf(`
		resource "harness_platform_organization" "test" {
			identifier = "org_%[1]s"
			name = "org_%[1]s"
		}	

		resource "harness_platform_project" "test" {
			identifier = "proj_%[1]s"
			name = "proj_%[1]s"
			org_id = harness_platform_organization.test.id
		}
		
		resource "harness_platform_repo" "test" {
			identifier  = "%[1]s"
			org_id = harness_platform_organization.test.id
			project_id = harness_platform_project.test.id
			default_branch = "master"
			description = "%[2]s"
			readme   = true
		}

		resource "harness_platform_repo_rule_branch" "test" {
			identifier  = "rule_%[1]s"
			repo_identifier = harness_platform_repo.test.identifier
			org_id = harness_platform_organization.test.id
			project_id = harness_platform_project.test.id
			description = "%[2]s"
			state = "active"
			pattern {
				default_branch = true
			}
			policies {
				require_pull_request = true
				require_status_checks = [
					"check-stage1"
				]
				allow_merge_strategies          = ["merge"]
			}
			bypass  {}
		}
	`, identifier, description,
	)
}

func testFindRepoRule(
	resourceName string,
	state *terraform.State,
) (*code.OpenapiRule, error) {
	r := acctest.TestAccGetResource(resourceName, state)
	c, ctx := acctest.TestAccGetCodeClientWithContext()
	repoIdentifier := r.Primary.Attributes["repo_identifier"]
	ruleId := r.Primary.Attributes["identifier"]

	rule, _, err := c.RepositoryApi.RuleGet(
		ctx, c.AccountId, repoIdentifier, ruleId,
		&code.RepositoryApiRuleGetOpts{
			OrgIdentifier:     optional.NewString(r.Primary.Attributes["org_id"]),
			ProjectIdentifier: optional.NewString(r.Primary.Attributes["project_id"]),
		})
	if err != nil {
		return nil, err
	}

	return &rule, nil
}

func testRepoRuleDestroy(resourceName string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		// The resource may already be gone from state (e.g. deleted out-of-band and
		// dropped on refresh), in which case there is nothing left to look up.
		if _, ok := state.RootModule().Resources[resourceName]; !ok {
			return nil
		}
		rule, _ := testFindRepoRule(resourceName, state)
		if rule != nil {
			return fmt.Errorf("found rule: %s", rule.Identifier)
		}

		return nil
	}
}
