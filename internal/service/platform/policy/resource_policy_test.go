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
	projectID := "OPA_TEST"
	id := fmt.Sprintf("%s%s", t.Name(), utils.RandStringBytes(5))
	name := id
	description := "TF-testing"
	orgID := "default"
	rego := "#Testing Policy Creation Using TF"
	updatedRego := "#Testing Policy Updation Using TF@"

	resourceName := "harness_platform_policy.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccPolicyDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourcePolicy(id, name, description, orgID, projectID, rego),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", description),
					resource.TestCheckResourceAttr(resourceName, "org_id", orgID),
					resource.TestCheckResourceAttr(resourceName, "project_id", projectID),
					resource.TestCheckResourceAttr(resourceName, "rego", rego),
				),
			},
			{
				Config: testAccResourcePolicy(id, name, description, orgID, projectID, updatedRego),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", description),
					resource.TestCheckResourceAttr(resourceName, "org_id", orgID),
					resource.TestCheckResourceAttr(resourceName, "project_id", projectID),
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

func testAccResourcePolicy(id, name, description, orgID, projectID, rego string) string {
	return fmt.Sprintf(`
		resource "harness_platform_policy" "test" {
			identifier       = "%[1]s"
			name             = "%[2]s"
			description      = "%[3]s"
			org_id           = "%[4]s"
			project_id       = "%[5]s"
			rego = "%[6]s"
		}
	`, id, name, description, orgID, projectID, rego)
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
