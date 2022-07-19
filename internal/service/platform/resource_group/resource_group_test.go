package resource_group_test

import (
	"fmt"
	"testing"

	"github.com/antihax/optional"
	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccResourceResourceGroup(t *testing.T) {
	name := t.Name()
	id := fmt.Sprintf("%s_%s", name, utils.RandStringBytes(5))
	updatedName := fmt.Sprintf("%s_updated", name)
	resourceName := "harness_platform_resource_group.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccResourceGroupDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceResourceGroup(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "account_id", "UKh5Yts7THSMAbccG3HrLA"),
					resource.TestCheckResourceAttr(resourceName, "allowed_scope_levels.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
				),
			},
			{
				Config: testAccResourceResourceGroup(id, updatedName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
					resource.TestCheckResourceAttr(resourceName, "account_id", "UKh5Yts7THSMAbccG3HrLA"),
					resource.TestCheckResourceAttr(resourceName, "allowed_scope_levels.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})

}

func testAccResourceGroupDestroy(resourceName string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		resourceGroup, _ := testAccGetResourceGroup(resourceName, state)
		if resourceGroup != nil {
			return fmt.Errorf("Found resource group: %s", resourceGroup.Identifier)
		}
		return nil
	}
}

func testAccGetResourceGroup(resourceName string, state *terraform.State) (*nextgen.ResourceGroupV2, error) {
	r := acctest.TestAccGetResource(resourceName, state)
	c, ctx := acctest.TestAccGetPlatformClientWithContext()
	id := r.Primary.ID

	resp, _, err := c.HarnessResourceGroupApi.GetResourceGroupV2(ctx, id, c.AccountId, &nextgen.HarnessResourceGroupApiGetResourceGroupV2Opts{
		OrgIdentifier:     buildField(r, "org_id"),
		ProjectIdentifier: buildField(r, "project_id"),
	})

	if err != nil {
		return nil, err
	}

	if resp.Data == nil {
		return nil, nil
	}

	return resp.Data.ResourceGroup, nil
}

func buildField(r *terraform.ResourceState, field string) optional.String {
	if attr, ok := r.Primary.Attributes[field]; ok {
		return optional.NewString(attr)
	}
	return optional.EmptyString()
}

func testAccResourceResourceGroup(id string, name string) string {
	return fmt.Sprintf(`
		resource "harness_platform_resource_group" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
			description = "test"
			tags = ["foo:bar"]

			account_id = "UKh5Yts7THSMAbccG3HrLA"
			allowed_scope_levels =["account"]
			included_scopes {
				filter = "EXCLUDING_CHILD_SCOPES"
				account_id = "UKh5Yts7THSMAbccG3HrLA"
			}
			resource_filter {
				include_all_resources = false
				resources {
					resource_type = "CONNECTOR"
					attribute_filter {
						attribute_name = "category"
						attribute_values = ["value"]
					}
				}
			}
		}
	`, id, name)
}
