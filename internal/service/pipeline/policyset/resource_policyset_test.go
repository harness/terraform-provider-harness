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

	policyFirstIdentifier := fmt.Sprintf("policyFirst%s", utils.RandStringBytes(5))
	policySecondIdentifier := fmt.Sprintf("policySecond%s", utils.RandStringBytes(5))

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccPolicysetDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourcePolicyset(id, name, action, policyType, policyFirstIdentifier, policySecondIdentifier, false),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "action", action),
					resource.TestCheckResourceAttr(resourceName, "type", policyType),
					resource.TestCheckResourceAttr(resourceName, "enabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "policies.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "policies.0.identifier", policyFirstIdentifier),
					resource.TestCheckResourceAttr(resourceName, "policies.0.severity", "warning"),
					resource.TestCheckResourceAttr(resourceName, "policies.1.identifier", policySecondIdentifier),
					resource.TestCheckResourceAttr(resourceName, "policies.1.severity", "warning"),
				),
			},
			{
				Config: testAccResourcePolicyset(id, name, action, policyType, policyFirstIdentifier, policySecondIdentifier, true),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "action", action),
					resource.TestCheckResourceAttr(resourceName, "type", policyType),
					resource.TestCheckResourceAttr(resourceName, "enabled", "true"),
					// re-applying the same policies block with no changes must not reorder the
					// list, guarding against the API's unstable linked-policy ordering.
					resource.TestCheckResourceAttr(resourceName, "policies.0.identifier", policyFirstIdentifier),
					resource.TestCheckResourceAttr(resourceName, "policies.1.identifier", policySecondIdentifier),
				),
			},
			{
				Config: testAccResourcePolicysetUpdate(id, name, action, policyType, policyFirstIdentifier, policySecondIdentifier, true),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "action", action),
					resource.TestCheckResourceAttr(resourceName, "type", policyType),
					resource.TestCheckResourceAttr(resourceName, "enabled", "true"),
					resource.TestCheckResourceAttr(resourceName, "policy_references.#", "2"),
					resource.TestCheckTypeSetElemNestedAttrs(resourceName, "policy_references.*", map[string]string{
						"identifier": policyFirstIdentifier,
						"severity":   "warning",
					}),
					resource.TestCheckTypeSetElemNestedAttrs(resourceName, "policy_references.*", map[string]string{
						"identifier": policySecondIdentifier,
						"severity":   "warning",
					}),
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

// TestAccResourcePolicysetCrossScope covers PIPE-37472: an org-scoped policyset linking an
// account-scoped policy via a scope-qualified identifier ("account.<id>") must keep that prefix
// on every subsequent read, not just at create time. The SDK test framework's automatic
// post-apply plan check fails if a spurious diff (prefix dropped/reordered) reappears on refresh.
func TestAccResourcePolicysetCrossScope(t *testing.T) {
	id := fmt.Sprintf("%s%s", t.Name(), utils.RandStringBytes(5))
	resourceName := "harness_platform_policyset.test"
	policyIdentifier := fmt.Sprintf("policy%s", utils.RandStringBytes(5))

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccPolicysetDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourcePolicysetCrossScope(id, policyIdentifier),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "policies.0.identifier", "account."+policyIdentifier),
				),
			},
			{
				// re-apply the identical config: must be a no-op, proving the prefix survives read.
				Config: testAccResourcePolicysetCrossScope(id, policyIdentifier),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "policies.0.identifier", "account."+policyIdentifier),
				),
			},
		},
	})
}

func testAccResourcePolicysetCrossScope(id, policyIdentifier string) string {
	return fmt.Sprintf(`
		resource "harness_platform_organization" "test" {
			identifier = "%[1]s"
			name       = "%[1]s"
		}

		resource "harness_platform_policy" "test" {
			identifier = "%[2]s"
			name       = "%[2]s"
			rego       = "some text"
		}

		resource "harness_platform_policyset" "test" {
			identifier = "%[1]s"
			name       = "%[1]s"
			org_id     = harness_platform_organization.test.id
			action     = "onrun"
			type       = "pipeline"
			enabled    = false

			policies {
				identifier = "account.${harness_platform_policy.test.identifier}"
				severity   = "warning"
			}
		}
`, id, policyIdentifier)
}

func testAccResourcePolicyset(id, name, action, policyType string, policyFirstIdentifier, policySecondIdentifier string, enabled bool) string {
	return fmt.Sprintf(`
		resource "harness_platform_policy" "first" {
			identifier = "%[6]s"
			name = "%[6]s"
			rego = "some text"
		}

		resource "harness_platform_policy" "second" {
			identifier = "%[7]s"
			name = "%[7]s"
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
`, id, name, action, policyType, enabled, policyFirstIdentifier, policySecondIdentifier)
}

func testAccResourcePolicysetUpdate(id, name, action, policyType string, policyFirstIdentifier, policySecondIdentifier string, enabled bool) string {
	return fmt.Sprintf(`
		resource "harness_platform_policy" "first" {
			identifier = "%[6]s"
			name = "%[6]s"
			rego = "some text"
		}

		resource "harness_platform_policy" "second" {
			identifier = "%[7]s"
			name = "%[7]s"
			rego = "some text"
		}
		resource "harness_platform_policyset" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
			action = "%[3]s"
			type = "%[4]s"
			enabled = %[5]t
			policy_references {
				identifier = harness_platform_policy.first.identifier
			  	severity = "warning"
			}

			policy_references {
				identifier = harness_platform_policy.second.identifier
				severity = "warning"
			}
		}
`, id, name, action, policyType, enabled, policyFirstIdentifier, policySecondIdentifier)
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
