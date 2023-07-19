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
	resourceName := "harness_platform_policy.test"
	rego := "some text"
	updatedRego := "some text v2"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccPolicyDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourcePolicy(id, name, rego),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "rego", rego),
					resource.TestCheckResourceAttr(resourceName, "git_connector_ref", "gitconnector"),
					resource.TestCheckResourceAttr(resourceName, "git_path", "path"),
					resource.TestCheckResourceAttr(resourceName, "git_repo", "harness-core"),
				),
			},
			{
				Config: testAccResourcePolicy(id, name, updatedRego),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "rego", updatedRego),
					resource.TestCheckResourceAttr(resourceName, "git_connector_ref", "gitconnector"),
					resource.TestCheckResourceAttr(resourceName, "git_path", "path"),
					resource.TestCheckResourceAttr(resourceName, "git_repo", "harness-core"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: acctest.OrgResourceImportStateIdFunc(resourceName),
			},
		},
	})
}

func testAccResourcePolicy(id, name, rego string) string {
	return fmt.Sprintf(`
		resource "harness_platform_policy" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
      rego = "%[3]s"
	  git_connector_ref = "gitconnector"
	  git_path = "path"
	  git_repo = "harness-core"
	  git_commit_msg = "hello world"
	  git_import = false
	  git_branch = "develop"
	  git_is_new_branch = false
	  git_base_branch = "main"
		}
`, id, name, rego)
}

func testAccPolicyDestroy(resourceName string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		policy, _ := testAccGetPolicy(resourceName, state)
		if policy != emptyPolicy {
			return fmt.Errorf("Found project: %s", policy.Identifier)
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
