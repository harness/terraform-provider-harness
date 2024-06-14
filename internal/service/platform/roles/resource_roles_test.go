package roles_test

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

func TestAccResourceRoles(t *testing.T) {

	name := t.Name()
	id := fmt.Sprintf("%s_%s", name, utils.RandStringBytes(5))
	updatedName := fmt.Sprintf("%s_updated", name)

	resourceName := "harness_platform_roles.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccRolesDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceRoles(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "permissions.0", "core_pipeline_edit"),
					resource.TestCheckResourceAttr(resourceName, "allowed_scope_levels.0", "account"),
				),
			},
			{
				Config: testAccResourceRoles(id, updatedName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "permissions.0", "core_pipeline_edit"),
					resource.TestCheckResourceAttr(resourceName, "allowed_scope_levels.0", "account"),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"org_id", "project_id"},
			},
		},
	})

}

func TestProjectResourceRoles(t *testing.T) {

	name := t.Name()
	id := fmt.Sprintf("%s_%s", name, utils.RandStringBytes(5))
	updatedName := fmt.Sprintf("%s_updated", name)

	resourceName := "harness_platform_roles.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccRolesDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testProjectResourceRoles(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "permissions.0", "core_pipeline_edit"),
					resource.TestCheckResourceAttr(resourceName, "allowed_scope_levels.0", "project"),
				),
			},
			{
				Config: testProjectResourceRoles(id, updatedName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "permissions.0", "core_pipeline_edit"),
					resource.TestCheckResourceAttr(resourceName, "allowed_scope_levels.0", "project"),
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

func TestOrgResourceRoles(t *testing.T) {

	name := t.Name()
	id := fmt.Sprintf("%s_%s", name, utils.RandStringBytes(5))
	updatedName := fmt.Sprintf("%s_updated", name)

	resourceName := "harness_platform_roles.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccRolesDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testOrgResourceRoles(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "permissions.0", "core_pipeline_edit"),
					resource.TestCheckResourceAttr(resourceName, "allowed_scope_levels.0", "organization"),
				),
			},
			{
				Config: testOrgResourceRoles(id, updatedName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "permissions.0", "core_pipeline_edit"),
					resource.TestCheckResourceAttr(resourceName, "allowed_scope_levels.0", "organization"),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateIdFunc:       acctest.OrgResourceImportStateIdFunc(resourceName),
				ImportStateVerifyIgnore: []string{"project_id"},
			},
		},
	})
}

func TestRepoResourceRoles(t *testing.T) {
	name := t.Name()
	id := fmt.Sprintf("%s_%s", name, utils.RandStringBytes(5))
	resourceName := "harness_platform_roles.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccRolesDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testRepoResourceRoles(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "permissions.0", "code_repo_push"),
					resource.TestCheckResourceAttr(resourceName, "permissions.1", "code_repo_view"),
					resource.TestCheckResourceAttr(resourceName, "allowed_scope_levels.0", "project"),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateIdFunc:       acctest.ProjectResourceImportStateIdFunc(resourceName),
			},
		},
	})
}

func testAccRolesDestroy(resourceName string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		roles, _ := testAccGetRoles(resourceName, state)
		if roles != nil {
			return fmt.Errorf("found role: %s", roles.Identifier)
		}
		return nil
	}
}

func testAccGetRoles(resourceName string, state *terraform.State) (*nextgen.Role, error) {
	r := acctest.TestAccGetResource(resourceName, state)
	c, ctx := acctest.TestAccGetPlatformClientWithContext()
	id := r.Primary.ID

	resp, _, err := c.RolesApi.GetRole(ctx, id, &nextgen.RolesApiGetRoleOpts{
		AccountIdentifier: optional.NewString(c.AccountId),
		OrgIdentifier:     buildField(r, "org_id"),
		ProjectIdentifier: buildField(r, "project_id"),
	})

	if err != nil {
		return nil, err
	}

	return resp.Data.Role, nil
}

func buildField(r *terraform.ResourceState, field string) optional.String {
	if attr, ok := r.Primary.Attributes[field]; ok {
		return optional.NewString(attr)
	}
	return optional.EmptyString()
}

func testAccResourceRoles(id string, name string) string {
	return fmt.Sprintf(`
	resource "harness_platform_roles" "test" {
		identifier = "%[1]s"
		name = "%[2]s"
		description = "test"
		tags = ["foo:bar"]
		permissions = ["core_pipeline_edit"]
		allowed_scope_levels = ["account"]
	}
`, id, name)
}

func testProjectResourceRoles(id string, name string) string {
	return fmt.Sprintf(`
	resource "harness_platform_organization" "test" {
		identifier = "%[1]s"
		name = "%[2]s"
	}

	resource "harness_platform_project" "test" {
    identifier = "%[1]s"
    name = "%[2]s"
    org_id = harness_platform_organization.test.id
    color = "#472848"
	}

	resource "harness_platform_roles" "test" {
		org_id = harness_platform_organization.test.id
		project_id=harness_platform_project.test.id
		identifier = "%[1]s"
		name = "%[2]s"
		description = "test"
		tags = ["foo:bar"]
		permissions = ["core_pipeline_edit"]
		allowed_scope_levels = ["project"]
	}
`, id, name)
}

func testOrgResourceRoles(id string, name string) string {
	return fmt.Sprintf(`
	resource "harness_platform_organization" "test" {
		identifier = "%[1]s"
		name = "%[2]s"
	}

	

	resource "harness_platform_roles" "test" {
		org_id = harness_platform_organization.test.id
		
		identifier = "%[1]s"
		name = "%[2]s"
		description = "test"
		tags = ["foo:bar"]
		permissions = ["core_pipeline_edit"]
		allowed_scope_levels = ["organization"]
	}
`, id, name)
}

func testRepoResourceRoles(id string, name string) string {
	return fmt.Sprintf(`
	resource "harness_platform_organization" "test" {
		identifier = "%[1]s"
		name = "%[2]s"
	}

	resource "harness_platform_project" "test" {
    identifier = "%[1]s"
    name = "%[2]s"
    org_id = harness_platform_organization.test.id
    color = "#472848"
	}

	resource "harness_platform_roles" "test" {
		org_id = harness_platform_organization.test.id
		project_id=harness_platform_project.test.id
		identifier = "%[1]s"
		name = "%[2]s"
		description = "test"
		tags = ["foo:bar"]
		permissions = ["code_repo_push", "code_repo_view"]
		allowed_scope_levels = ["project"]
	}
`, id, name)
}
