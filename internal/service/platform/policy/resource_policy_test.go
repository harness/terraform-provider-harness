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
	// To test this please update all the fields with valid values.
	projectID := "OPA_TEST"
	id := fmt.Sprintf("%s%s", t.Name(), utils.RandStringBytes(5))
	name := id
	description := "terratest"
	orgID := "Ng_Pipelines_K8s_Organisations"
	gitConnectorRef := "Sameed_Test"
	gitPath := ".harness/" + id + ".rego"
	gitRepo := "test_sameed"
	gitBranch := "main"
	gitBaseBranch := "main"
	gitIsNewBranch := false
	gitImport := false
	gitCommitMsg := "Trying TF out"
	rego := "some text"
	updatedRego := "some text v2"

	resourceName := "harness_platform_policy.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccPolicyDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourcePolicy(id, name, description, orgID, projectID, gitConnectorRef, gitPath, gitRepo, gitBranch, gitBaseBranch, gitIsNewBranch, gitImport, gitCommitMsg, rego),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", description),
					resource.TestCheckResourceAttr(resourceName, "org_id", orgID),
					resource.TestCheckResourceAttr(resourceName, "project_id", projectID),
					resource.TestCheckResourceAttr(resourceName, "git_connector_ref", gitConnectorRef),
					resource.TestCheckResourceAttr(resourceName, "git_path", gitPath),
					resource.TestCheckResourceAttr(resourceName, "git_repo", gitRepo),
					resource.TestCheckResourceAttr(resourceName, "git_branch", gitBranch),
					resource.TestCheckResourceAttr(resourceName, "git_base_branch", gitBaseBranch),
					resource.TestCheckResourceAttr(resourceName, "git_is_new_branch", fmt.Sprintf("%t", gitIsNewBranch)),
					resource.TestCheckResourceAttr(resourceName, "git_import", fmt.Sprintf("%t", gitImport)),
					resource.TestCheckResourceAttr(resourceName, "git_commit_msg", gitCommitMsg),
					resource.TestCheckResourceAttr(resourceName, "rego", rego),
				),
			},
			{
				Config: testAccResourcePolicy(id, name, description, orgID, projectID, gitConnectorRef, gitPath, gitRepo, gitBranch, gitBaseBranch, gitIsNewBranch, gitImport, gitCommitMsg, updatedRego),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", description),
					resource.TestCheckResourceAttr(resourceName, "org_id", orgID),
					resource.TestCheckResourceAttr(resourceName, "project_id", projectID),
					resource.TestCheckResourceAttr(resourceName, "git_connector_ref", gitConnectorRef),
					resource.TestCheckResourceAttr(resourceName, "git_path", gitPath),
					resource.TestCheckResourceAttr(resourceName, "git_repo", gitRepo),
					resource.TestCheckResourceAttr(resourceName, "git_branch", gitBranch),
					resource.TestCheckResourceAttr(resourceName, "git_base_branch", gitBaseBranch),
					resource.TestCheckResourceAttr(resourceName, "git_is_new_branch", fmt.Sprintf("%t", gitIsNewBranch)),
					resource.TestCheckResourceAttr(resourceName, "git_import", fmt.Sprintf("%t", gitImport)),
					resource.TestCheckResourceAttr(resourceName, "git_commit_msg", gitCommitMsg),
					resource.TestCheckResourceAttr(resourceName, "rego", updatedRego),
				),
			},
		},
	})
}

func testAccResourcePolicy(id, name, description, orgID, projectID, gitConnectorRef, gitPath, gitRepo, gitBranch, gitBaseBranch string, gitIsNewBranch, gitImport bool, gitCommitMsg, rego string) string {
	return fmt.Sprintf(`
		resource "harness_platform_policy" "test" {
			identifier       = "%[1]s"
			name             = "%[2]s"
			description      = "%[3]s"
			org_id           = "%[4]s"
			project_id       = "%[5]s"
			git_connector_ref = "%[6]s"
			git_path         = "%[7]s"
			git_repo         = "%[8]s"
			git_branch       = "%[9]s"
			git_base_branch  = "%[10]s"
			git_is_new_branch = %[11]t
			git_import       = %[12]t
			git_commit_msg   = "%[13]s"
			rego = "%[14]s"
		}
	`, id, name, description, orgID, projectID, gitConnectorRef, gitPath, gitRepo, gitBranch, gitBaseBranch, gitIsNewBranch, gitImport, gitCommitMsg, rego)
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
