package policy_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/antihax/optional"
	"github.com/harness/harness-go-sdk/harness/policymgmt"
	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

var (
	emptyPolicy = policymgmt.Policy{}
)

func TestAccResourcePolicy(t *testing.T) {
	id := fmt.Sprintf("%s%s", t.Name(), utils.RandStringBytes(5))
	name := id
	description := "TF-testing"
	rego := "#Testing Policy Creation Using TF"
	updatedRego := "#Testing Policy Updation Using TF"

	resourceName := "harness_platform_policy.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccPolicyDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourcePolicy(id, name, description, rego),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", description),
					resource.TestCheckResourceAttr(resourceName, "rego", rego),
				),
			},
			{
				Config: testAccResourcePolicy(id, name, description, updatedRego),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", description),
					resource.TestCheckResourceAttr(resourceName, "rego", updatedRego),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateIdFunc:       acctest.ProjectResourceImportStateIdFunc(resourceName),
				ImportStateVerifyIgnore: []string{"description", "git_base_branch", "git_branch", "git_commit_msg", "git_import", "git_is_new_branch"},
			},
		},
	})
}

func testAccResourcePolicy(id, name, description, rego string) string {
	return fmt.Sprintf(`
		resource "harness_platform_policy" "test" {
			identifier       = "%[1]s"
			name             = "%[2]s"
			description      = "%[3]s"
			rego = "%[4]s"
		}
	`, id, name, description, rego)
}

// TestAccResourcePolicyOrgProject covers policy creation scoped at the org level and at the
// project level (in addition to the existing account-level TestAccResourcePolicy), exercising the
// org_id/project_id plumbing shared with harness_platform_policyset.
func TestAccResourcePolicyOrgProject(t *testing.T) {
	id := fmt.Sprintf("%s%s", t.Name(), utils.RandStringBytes(5))
	rego := "#Testing Policy Creation Using TF"
	updatedRego := "#Testing Policy Updation Using TF"

	orgResourceName := "harness_platform_policy.org"
	projectResourceName := "harness_platform_policy.project"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      resource.ComposeTestCheckFunc(testAccPolicyDestroy(orgResourceName), testAccPolicyDestroy(projectResourceName)),
		Steps: []resource.TestStep{
			{
				Config: testAccResourcePolicyOrgProject(id, rego),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(orgResourceName, "identifier", id),
					resource.TestCheckResourceAttrSet(orgResourceName, "org_id"),
					resource.TestCheckResourceAttr(orgResourceName, "project_id", ""),
					resource.TestCheckResourceAttr(orgResourceName, "rego", rego),

					resource.TestCheckResourceAttr(projectResourceName, "identifier", id),
					resource.TestCheckResourceAttrSet(projectResourceName, "org_id"),
					resource.TestCheckResourceAttrSet(projectResourceName, "project_id"),
					resource.TestCheckResourceAttr(projectResourceName, "rego", rego),
				),
			},
			{
				Config: testAccResourcePolicyOrgProject(id, updatedRego),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(orgResourceName, "rego", updatedRego),
					resource.TestCheckResourceAttr(projectResourceName, "rego", updatedRego),
				),
			},
		},
	})
}

func testAccResourcePolicyOrgProject(id, rego string) string {
	return fmt.Sprintf(`
		resource "harness_platform_organization" "test" {
			identifier = "%[1]s"
			name       = "%[1]s"
		}

		resource "harness_platform_project" "test" {
			identifier = "%[1]s"
			org_id     = harness_platform_organization.test.id
			name       = "%[1]s"
		}

		resource "harness_platform_policy" "org" {
			identifier = "%[1]s"
			name       = "%[1]s"
			org_id     = harness_platform_organization.test.id
			rego       = "%[2]s"
		}

		resource "harness_platform_policy" "project" {
			identifier = "%[1]s"
			name       = "%[1]s"
			org_id     = harness_platform_organization.test.id
			project_id = harness_platform_project.test.id
			rego       = "%[2]s"
		}
	`, id, rego)
}

// TestAccResourcePolicyGit covers creating a policy backed by a git-stored rego file, exercising
// the git_connector_ref/git_repo/git_path/git_branch fields that TestAccResourcePolicy's
// ImportStateVerifyIgnore list otherwise leaves untested.
//
// Requires a pre-existing github connector in the test account; defaults to the shared
// "account.TF_Jajoo_github_connector" / "jajoo_git" fixture used by other acceptance tests in
// this repo. Override via TF_VAR-style constants below if that fixture is unavailable.
func TestAccResourcePolicyGit(t *testing.T) {
	id := fmt.Sprintf("%s%s", t.Name(), utils.RandStringBytes(5))
	resourceName := "harness_platform_policy.git"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccPolicyDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourcePolicyGit(id),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "git_connector_ref", "account.TF_Jajoo_github_connector"),
					resource.TestCheckResourceAttr(resourceName, "git_repo", "jajoo_git"),
					resource.TestCheckResourceAttrSet(resourceName, "git_path"),
				),
			},
		},
	})
}

func testAccResourcePolicyGit(id string) string {
	return fmt.Sprintf(`
		resource "harness_platform_policy" "git" {
			identifier        = "%[1]s"
			name              = "%[1]s"
			rego              = "# git-backed policy %[1]s"
			git_connector_ref = "account.TF_Jajoo_github_connector"
			git_repo          = "jajoo_git"
			git_path          = ".harness/policies/%[1]s.rego"
			git_branch        = "main"
			git_commit_msg    = "Add policy %[1]s via terraform acceptance test"
		}
	`, id)
}

func testAccPolicyDestroy(resourceName string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		policy, _ := testAccGetPolicy(resourceName, state)
		if policy != emptyPolicy {
			return fmt.Errorf("Found policy: %s", policy.Identifier)
		}
		return nil
	}
}

func testAccGetPolicy(resourceName string, state *terraform.State) (policymgmt.Policy, error) {
	r := acctest.TestAccGetApiClientFromProvider()
	c := acctest.TestAccGetPolicyManagementClient()
	localVarOptionals := policymgmt.PoliciesApiPoliciesFindOpts{
		AccountIdentifier: optional.NewString(r.AccountId),
		XApiKey:           optional.NewString(r.PLClient.ApiKey),
	}
	policy, _, err := c.PoliciesApi.PoliciesFind(context.Background(), resourceName, &localVarOptionals)
	if err != nil {
		return emptyPolicy, err
	}
	return policy, nil
}
