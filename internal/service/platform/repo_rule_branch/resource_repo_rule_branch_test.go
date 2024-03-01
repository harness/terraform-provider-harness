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
)

const (
	resourceName = "harness_platform_repo_rule_branch.test"
	description  = "example_description"
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
					resource.TestCheckResourceAttr(resourceName, "rules.0.require_pull_request", "true"),
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
			target_patterns {
				default_branch = true
			}
			rules {
				require_pull_request = true
			}
			bypass_list  {}
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
		rule, _ := testFindRepoRule(resourceName, state)
		if rule != nil {
			return fmt.Errorf("Found rule: %s", rule.Identifier)
		}

		return nil
	}
}
