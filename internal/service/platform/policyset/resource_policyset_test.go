package policyset_test

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
	emptyPolicyset = policymgmt.PolicySet{}
)

func TestAccResourcePolicyset(t *testing.T) {
	id := fmt.Sprintf("%s%s", t.Name(), utils.RandStringBytes(5))
	name := id
	resourceName := "harness_platform_policyset.test"
	policyType := "pipeline"
	action := "onrun"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccPolicysetDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourcePolicyset(id, name, action, policyType, false),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "action", action),
					resource.TestCheckResourceAttr(resourceName, "type", policyType),
					resource.TestCheckResourceAttr(resourceName, "enabled", "false"),
				),
			},
			{
				Config: testAccResourcePolicyset(id, name, action, policyType, true),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "action", action),
					resource.TestCheckResourceAttr(resourceName, "type", policyType),
					resource.TestCheckResourceAttr(resourceName, "enabled", "true"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: acctest.AccountLevelResourceImportStateIdFunc(resourceName),
			},
		},
	})
}

func testAccResourcePolicyset(id, name, action, policyType string, enabled bool) string {
	return fmt.Sprintf(`
		resource "harness_platform_policy" "first" {
			identifier = "policyFirst"
			name = "policyFirst"
			rego = "some text"
		}

		resource "harness_platform_policy" "second" {
			identifier = "policySecond"
			name = "policySecond"
			rego = "some text"
		}
		resource "harness_platform_policyset" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
			action = "%[3]s"
			type = "%[4]s"
			enabled = %[5]t
			policies {
				identifier = harness_platform_policy.first.identifier
			  severity = "warning"
			}

			policies {
				identifier = harness_platform_policy.second.identifier
			  severity = "warning"
			}
		}
`, id, name, action, policyType, enabled)
}

func testAccPolicysetDestroy(resourceName string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		policyset, _ := testAccGetPolicyset(resourceName, state)
		if policyset.Identifier != emptyPolicyset.Identifier {
			return fmt.Errorf("Found project: %s", policyset.Identifier)
		}

		return nil
	}
}

func testAccGetPolicyset(resourceName string, state *terraform.State) (policymgmt.PolicySet, error) {
	r := acctest.TestAccGetApiClientFromProvider()
	c := acctest.TestAccGetPolicyManagementClient()
	localVarOptionals := policymgmt.PolicysetsApiPolicysetsFindOpts{
		AccountIdentifier: optional.NewString(r.AccountId),
		XApiKey:           optional.NewString(r.PLClient.ApiKey),
	}
	policyset, _, err := c.PolicysetsApi.PolicysetsFind(context.Background(), resourceName, &localVarOptionals)
	if err != nil {
		return emptyPolicyset, err
	}

	return policyset, nil
}
