package ansible_inventory_test

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

func TestAccResourceAnsibleInventory(t *testing.T) {
	resourceName := "harness_platform_iacm_ansible_inventory.test"
	name := t.Name()
	id := fmt.Sprintf("%s_%s", name, utils.RandStringBytes(5))
	updatedName := fmt.Sprintf("%s_updated", name)

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccResourceAnsibleInventoryDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceAnsibleInventoryManual(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "type", "manual"),
					resource.TestCheckResourceAttr(resourceName, "groups.#", "1"),
				),
			},
			{
				Config: testAccResourceAnsibleInventoryManual(id, updatedName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateIdFunc: acctest.ProjectResourceImportStateIdFunc(resourceName),
			},
		},
	})
}

func testAccResourceAnsibleInventoryDestroy(resourceName string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		inv, _ := testAccGetPlatformAnsibleInventory(resourceName, state)
		if inv != nil {
			return fmt.Errorf("Ansible inventory found: %s", inv.Identifier)
		}
		return nil
	}
}

func testAccGetPlatformAnsibleInventory(resourceName string, state *terraform.State) (*nextgen.ShowInventoryResponse, error) {
	r := acctest.TestAccGetResource(resourceName, state)
	c, ctx := acctest.TestAccGetPlatformClientWithContext()
	id := r.Primary.ID
	org := r.Primary.Attributes["org_id"]
	project := r.Primary.Attributes["project_id"]

	inv, resp, err := c.AnsibleApi.AnsibleShowInventory(ctx, org, project, id, c.AccountId, &nextgen.AnsibleApiAnsibleShowInventoryOpts{
		UseArrays: optional.NewBool(true),
	})
	if err != nil {
		return nil, err
	}
	if resp == nil {
		return nil, nil
	}
	return &inv, nil
}

func testAccResourceAnsibleInventoryManual(id, name string) string {
	return fmt.Sprintf(`
		resource "harness_platform_organization" "test" {
			identifier = "%[1]s"
			name       = "%[2]s"
		}

		resource "harness_platform_project" "test" {
			identifier = "%[1]s"
			name       = "%[2]s"
			org_id     = harness_platform_organization.test.id
		}

		resource "harness_platform_iacm_ansible_inventory" "test" {
			identifier = "%[1]s"
			name       = "%[2]s"
			org_id     = harness_platform_organization.test.id
			project_id = harness_platform_project.test.id
			type       = "manual"

			groups {
				identifier = "web"
				name       = "web"
				hosts      = ["web-1.example.com", "web-2.example.com"]
				vars {
					key        = "ansible_user"
					value      = "ubuntu"
					value_type = "string"
				}
			}

			vars {
				key        = "ansible_port"
				value      = "22"
				value_type = "string"
			}
		}
	`, id, name)
}
