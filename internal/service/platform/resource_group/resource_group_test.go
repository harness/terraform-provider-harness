package resource_group_test

import (
	"fmt"
	"os"
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
	accountId := os.Getenv("HARNESS_ACCOUNT_ID")

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccResourceGroupDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceResourceGroup(id, name, accountId),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "account_id", accountId),
					resource.TestCheckResourceAttr(resourceName, "allowed_scope_levels.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
				),
			},
			{
				Config: testAccResourceResourceGroup(id, updatedName, accountId),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
					resource.TestCheckResourceAttr(resourceName, "account_id", accountId),
					resource.TestCheckResourceAttr(resourceName, "allowed_scope_levels.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: acctest.ProjectResourceImportStateIdFunc(resourceName),
			},
		},
	})

}

func TestProjectResourceResourceGroup(t *testing.T) {
	name := t.Name()
	id := fmt.Sprintf("%s_%s", name, utils.RandStringBytes(5))
	updatedName := fmt.Sprintf("%s_updated", name)
	resourceName := "harness_platform_resource_group.test"
	accountId := os.Getenv("HARNESS_ACCOUNT_ID")

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccResourceGroupDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testProjectResourceResourceGroup(id, name, accountId),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "account_id", accountId),
					resource.TestCheckResourceAttr(resourceName, "allowed_scope_levels.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
				),
			},
			{
				Config: testProjectResourceResourceGroup(id, updatedName, accountId),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
					resource.TestCheckResourceAttr(resourceName, "account_id", accountId),
					resource.TestCheckResourceAttr(resourceName, "allowed_scope_levels.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: acctest.ProjectResourceImportStateIdFunc(resourceName),
			},
		},
	})

}

func TestOrgResourceResourceGroup(t *testing.T) {
	name := t.Name()
	id := fmt.Sprintf("%s_%s", name, utils.RandStringBytes(5))
	updatedName := fmt.Sprintf("%s_updated", name)
	resourceName := "harness_platform_resource_group.test"
	accountId := os.Getenv("HARNESS_ACCOUNT_ID")

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccResourceGroupDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testOrgResourceResourceGroup(id, name, accountId),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "account_id", accountId),
					resource.TestCheckResourceAttr(resourceName, "allowed_scope_levels.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
				),
			},
			{
				Config: testOrgResourceResourceGroup(id, updatedName, accountId),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
					resource.TestCheckResourceAttr(resourceName, "account_id", accountId),
					resource.TestCheckResourceAttr(resourceName, "allowed_scope_levels.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
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

func TestAccResourceResourceGroup_emptyAttributeFilter(t *testing.T) {
	name := t.Name()
	id := fmt.Sprintf("%s_%s", name, utils.RandStringBytes(5))
	updatedName := fmt.Sprintf("%s_updated", name)
	resourceName := "harness_platform_resource_group.test"
	accountId := os.Getenv("HARNESS_ACCOUNT_ID")

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccResourceGroupDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceResourceGroupEmptyAttributeFilter(id, name, accountId),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "account_id", accountId),
					resource.TestCheckResourceAttr(resourceName, "allowed_scope_levels.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
				),
			},
			{
				Config: testAccResourceResourceGroupEmptyAttributeFilter(id, updatedName, accountId),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
					resource.TestCheckResourceAttr(resourceName, "account_id", accountId),
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

func TestProjectResourceResourceGroup_emptyAttributeFilter(t *testing.T) {
	name := t.Name()
	id := fmt.Sprintf("%s_%s", name, utils.RandStringBytes(5))
	updatedName := fmt.Sprintf("%s_updated", name)
	resourceName := "harness_platform_resource_group.test"
	accountId := os.Getenv("HARNESS_ACCOUNT_ID")

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccResourceGroupDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testProjectResourceResourceGroupEmptyAttributeFilter(id, name, accountId),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "account_id", accountId),
					resource.TestCheckResourceAttr(resourceName, "allowed_scope_levels.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
				),
			},
			{
				Config: testProjectResourceResourceGroupEmptyAttributeFilter(id, updatedName, accountId),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
					resource.TestCheckResourceAttr(resourceName, "account_id", accountId),
					resource.TestCheckResourceAttr(resourceName, "allowed_scope_levels.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: acctest.ProjectResourceImportStateIdFunc(resourceName),
			},
		},
	})

}
func TestOrgResourceResourceGroup_emptyAttributeFilter(t *testing.T) {
	name := t.Name()
	id := fmt.Sprintf("%s_%s", name, utils.RandStringBytes(5))
	updatedName := fmt.Sprintf("%s_updated", name)
	resourceName := "harness_platform_resource_group.test"
	accountId := os.Getenv("HARNESS_ACCOUNT_ID")

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccResourceGroupDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testOrgResourceResourceGroupEmptyAttributeFilter(id, name, accountId),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "account_id", accountId),
					resource.TestCheckResourceAttr(resourceName, "allowed_scope_levels.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
				),
			},
			{
				Config: testOrgResourceResourceGroupEmptyAttributeFilter(id, updatedName, accountId),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
					resource.TestCheckResourceAttr(resourceName, "account_id", accountId),
					resource.TestCheckResourceAttr(resourceName, "allowed_scope_levels.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
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

func TestAccResourceResourceGroup_TagsAttributeFilter(t *testing.T) {
	name := t.Name()
	id := fmt.Sprintf("%s_%s", name, utils.RandStringBytes(5))
	updatedName := fmt.Sprintf("%s_updated", name)
	resourceName := "harness_platform_resource_group.test"
	accountId := os.Getenv("HARNESS_ACCOUNT_ID")

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccResourceGroupDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceResourceGroupTagsAttributeFilter(id, name, accountId),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "account_id", accountId),
					resource.TestCheckResourceAttr(resourceName, "allowed_scope_levels.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
				),
			},
			{
				Config: testAccResourceResourceGroupTagsAttributeFilter(id, updatedName, accountId),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
					resource.TestCheckResourceAttr(resourceName, "account_id", accountId),
					resource.TestCheckResourceAttr(resourceName, "allowed_scope_levels.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: acctest.ProjectResourceImportStateIdFunc(resourceName),
			},
		},
	})

}
func TestOrgResourceResourceGroup_TagsAttributeFilter(t *testing.T) {
	name := t.Name()
	id := fmt.Sprintf("%s_%s", name, utils.RandStringBytes(5))
	updatedName := fmt.Sprintf("%s_updated", name)
	resourceName := "harness_platform_resource_group.test"
	accountId := os.Getenv("HARNESS_ACCOUNT_ID")

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccResourceGroupDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testOrgResourceResourceGroupTagsAttributeFilter(id, name, accountId),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "account_id", accountId),
					resource.TestCheckResourceAttr(resourceName, "allowed_scope_levels.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
				),
			},
			{
				Config: testOrgResourceResourceGroupTagsAttributeFilter(id, updatedName, accountId),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
					resource.TestCheckResourceAttr(resourceName, "account_id", accountId),
					resource.TestCheckResourceAttr(resourceName, "allowed_scope_levels.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: acctest.ProjectResourceImportStateIdFunc(resourceName),
			},
		},
	})

}

func TestProjectResourceResourceGroup_TagsAttributeFilter(t *testing.T) {
	name := t.Name()
	id := fmt.Sprintf("%s_%s", name, utils.RandStringBytes(5))
	updatedName := fmt.Sprintf("%s_updated", name)
	resourceName := "harness_platform_resource_group.test"
	accountId := os.Getenv("HARNESS_ACCOUNT_ID")

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccResourceGroupDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testProjectResourceResourceGroupTagsAttributeFilter(id, name, accountId),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "account_id", accountId),
					resource.TestCheckResourceAttr(resourceName, "allowed_scope_levels.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
				),
			},
			{
				Config: testProjectResourceResourceGroupTagsAttributeFilter(id, updatedName, accountId),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
					resource.TestCheckResourceAttr(resourceName, "account_id", accountId),
					resource.TestCheckResourceAttr(resourceName, "allowed_scope_levels.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: acctest.ProjectResourceImportStateIdFunc(resourceName),
			},
		},
	})

}

func TestAccDataSourceResourceGroup_InvalidEnvironmentNames(t *testing.T) {
	var (
		name         = fmt.Sprintf("pl_auto_rg_%s", utils.RandStringBytes(5))
		resourceName = "harness_platform_resource_group.test"
		projectId    = "ResourceGroupTest"
		orgId        = "default"
		accountId    = os.Getenv("HARNESS_ACCOUNT_ID")
	)

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceGroup_InvalidEnvironmentNames(name, accountId, orgId, projectId),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", name),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "included_scopes.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "resource_filter.0.include_all_resources", "false"),
				),
			},
		},
	})
}

func testAccResourceGroup_InvalidEnvironmentNames(name string, accountId string, orgId string, projectId string) string {
	return fmt.Sprintf(`
	resource "harness_platform_resource_group" "test" {
		identifier  = "%[1]s"
		name        = "%[1]s"
		description = "test"
		tags        = ["foo:bar"]
	
		org_id      = "%[3]s"
		project_id  = "%[4]s"
		account_id  = "%[2]s"
		allowed_scope_levels = ["project"]
	
		included_scopes {
			filter     = "EXCLUDING_CHILD_SCOPES"
			account_id = "%[2]s"
			org_id     = "%[3]s"
			project_id = "%[4]s"
		}
	
		resource_filter {
			include_all_resources = false
			resources {
				resource_type = "ENVIRONMENT"
				identifiers = [
					"QA",
					"Prod-1"
				]
			}
		}

		lifecycle {
			prevent_destroy = true
		}
	}
`, name, accountId, orgId, projectId)
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

func testAccResourceResourceGroup(id string, name string, accountId string) string {
	return fmt.Sprintf(`
		resource "harness_platform_resource_group" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
			description = "test"
			tags = ["foo:bar"]

			account_id = "%[3]s"
			allowed_scope_levels =["account"]
			included_scopes {
				filter = "EXCLUDING_CHILD_SCOPES"
				account_id = "%[3]s"
				
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
	`, id, name, accountId)
}
func testProjectResourceResourceGroup(id string, name string, accountId string) string {
	return fmt.Sprintf(`
	resource "harness_platform_organization" "test" {
		identifier = "%[1]s"
		name = "%[2]s"
	}

	resource "harness_platform_project" "test" {
		identifier = "%[1]s"
		name = "%[2]s"
		color = "#0063F7"
		org_id = harness_platform_organization.test.identifier
	}
	
		resource "harness_platform_resource_group" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
			description = "test"
			tags = ["foo:bar"]

			account_id = "%[3]s"
			org_id = harness_platform_project.test.org_id
			project_id = harness_platform_project.test.id
			allowed_scope_levels =["project"]
			included_scopes {
				filter = "EXCLUDING_CHILD_SCOPES"
				account_id = "%[3]s"
				org_id = harness_platform_project.test.org_id
				project_id = harness_platform_project.test.id
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
	`, id, name, accountId)
}

func testOrgResourceResourceGroup(id string, name string, accountId string) string {
	return fmt.Sprintf(`
		resource "harness_platform_organization" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
		}
		resource "harness_platform_resource_group" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
			description = "test"
			tags = ["foo:bar"]

			account_id = "%[3]s"
			org_id = harness_platform_organization.test.id
			
			allowed_scope_levels =["organization"]
			included_scopes {
				filter = "EXCLUDING_CHILD_SCOPES"
				account_id = "%[3]s"
				org_id = harness_platform_organization.test.id
				
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
	`, id, name, accountId)
}

func testAccResourceResourceGroupEmptyAttributeFilter(id string, name string, accountId string) string {
	return fmt.Sprintf(`
		resource "harness_platform_resource_group" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
			description = "test"
			tags = ["foo:bar"]

			account_id = "%[3]s"
			allowed_scope_levels =["account"]
			included_scopes {
				filter = "EXCLUDING_CHILD_SCOPES"
				account_id = "%[3]s"
			}
			resource_filter {
				include_all_resources = false
				resources {
					resource_type = "CONNECTOR"
				}
			}
		}
	`, id, name, accountId)
}
func testProjectResourceResourceGroupEmptyAttributeFilter(id string, name string, accountId string) string {
	return fmt.Sprintf(`
	resource "harness_platform_organization" "test" {
		identifier = "%[1]s"
		name = "%[2]s"
	}

	resource "harness_platform_project" "test" {
		identifier = "%[1]s"
		name = "%[2]s"
		color = "#0063F7"
		org_id = harness_platform_organization.test.identifier
	}
	
		resource "harness_platform_resource_group" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
			description = "test"
			tags = ["foo:bar"]

			account_id = "%[3]s"
			org_id = harness_platform_project.test.org_id
			project_id = harness_platform_project.test.id
			allowed_scope_levels =["project"]
			included_scopes {
				filter = "EXCLUDING_CHILD_SCOPES"
				account_id = "%[3]s"
				org_id = harness_platform_project.test.org_id
				project_id = harness_platform_project.test.id
			}
			resource_filter {
				include_all_resources = false
				resources {
					resource_type = "CONNECTOR"
				}
			}
		}
	`, id, name, accountId)
}

func testOrgResourceResourceGroupEmptyAttributeFilter(id string, name string, accountId string) string {
	return fmt.Sprintf(`
		resource "harness_platform_organization" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
		}
	
		resource "harness_platform_resource_group" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
			description = "test"
			tags = ["foo:bar"]

			account_id = "%[3]s"
			org_id = harness_platform_organization.test.identifier
			allowed_scope_levels =["organization"]
			included_scopes {
				filter = "EXCLUDING_CHILD_SCOPES"
				account_id = "%[3]s"
				org_id = harness_platform_organization.test.identifier
			}
			resource_filter {
				include_all_resources = false
				resources {
					resource_type = "CONNECTOR"
				}
			}
		}
	`, id, name, accountId)
}

func testAccResourceResourceGroupTagsAttributeFilter(id string, name string, accountId string) string {
	return fmt.Sprintf(`
		resource "harness_platform_resource_group" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
			description = "test"
			tags = ["foo:bar"]

			account_id = "%[3]s"
			allowed_scope_levels =["account"]
			included_scopes {
				filter = "INCLUDING_CHILD_SCOPES"
				account_id = "%[3]s"
				
			}
			resource_filter {
				include_all_resources = false
				resources {
					resource_type = "PIPELINE"
					attribute_filter {
						attribute_name = "tags"
						attribute_values = ["t1:v1","t2:v2"]
					}
				}
			}
		}
	`, id, name, accountId)
}
func testOrgResourceResourceGroupTagsAttributeFilter(id string, name string, accountId string) string {
	return fmt.Sprintf(`
		resource "harness_platform_organization" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
		}

		resource "harness_platform_resource_group" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
			description = "test"
			tags = ["foo:bar"]

			account_id = "%[3]s"
			org_id = harness_platform_organization.test.identifier
			allowed_scope_levels =["organization"]
			included_scopes {
				filter = "INCLUDING_CHILD_SCOPES"
				account_id = "%[3]s"
				org_id = harness_platform_organization.test.identifier
			}
			resource_filter {
				include_all_resources = false
				resources {
					resource_type = "PIPELINE"
					attribute_filter {
						attribute_name = "tags"
						attribute_values = ["t1:v1","t2:v2"]
					}
				}
			}
		}
	`, id, name, accountId)
}
func testProjectResourceResourceGroupTagsAttributeFilter(id string, name string, accountId string) string {
	return fmt.Sprintf(`
		resource "harness_platform_organization" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
		}
	
		resource "harness_platform_project" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
			color = "#0063F7"
			org_id = harness_platform_organization.test.identifier
		}
		resource "harness_platform_resource_group" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
			description = "test"
			tags = ["foo:bar"]

			account_id = "%[3]s"
			org_id = harness_platform_project.test.org_id
			project_id = harness_platform_project.test.id
			allowed_scope_levels =["project"]
			included_scopes {
				filter = "EXCLUDING_CHILD_SCOPES"
				account_id = "%[3]s"
				org_id = harness_platform_project.test.org_id
				project_id = harness_platform_project.test.id
			}
			resource_filter {
				include_all_resources = false
				resources {
					resource_type = "PIPELINE"
					attribute_filter {
						attribute_name = "tags"
						attribute_values = ["t1:v1","t2:v2"]
					}
				}
			}
		}
	`, id, name, accountId)
}
