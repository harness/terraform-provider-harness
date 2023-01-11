package role_assignments_test

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
	"github.com/stretchr/testify/require"
)

func TestAccRoleAssignments(t *testing.T) {
	name := t.Name()
	id := fmt.Sprintf("%s_%s", name, utils.RandStringBytes(5))
	resourceName := "harness_platform_role_assignments.test"
	accountId := os.Getenv("HARNESS_ACCOUNT_ID")

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccRoleAssignmentsDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceRoleAssignments(id, name, "false", accountId),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "resource_group_identifier", "_all_project_level_resources"),
					resource.TestCheckResourceAttr(resourceName, "disabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "managed", "false"),
					resource.TestCheckResourceAttr(resourceName, "principal.0.type", "SERVICE_ACCOUNT"),
				),
			},
			{
				Config: testAccResourceRoleAssignments(id, name, "true", accountId),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "resource_group_identifier", "_all_project_level_resources"),
					resource.TestCheckResourceAttr(resourceName, "disabled", "true"),
					resource.TestCheckResourceAttr(resourceName, "managed", "false"),
					resource.TestCheckResourceAttr(resourceName, "principal.0.type", "SERVICE_ACCOUNT"),
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

func TestAccResourceRoleAssignments_DeleteUnderlyingResource(t *testing.T) {
	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	name := id
	resourceName := "harness_platform_role_assignments.test"
	accountId := os.Getenv("HARNESS_ACCOUNT_ID")

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceRoleAssignments(id, name, "false", accountId),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
				),
			},
			{
				PreConfig: func() {
					acctest.TestAccConfigureProvider()
					c, ctx := acctest.TestAccGetPlatformClientWithContext()
					_, _, err := c.RoleAssignmentsApi.DeleteRoleAssignment(ctx, c.AccountId, id, &nextgen.RoleAssignmentsApiDeleteRoleAssignmentOpts{
						OrgIdentifier:     optional.NewString(id),
						ProjectIdentifier: optional.NewString(id),
					})
					require.NoError(t, err)
				},
				Config:             testAccResourceRoleAssignments(id, name, "false", accountId),
				PlanOnly:           true,
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func testAccRoleAssignmentsDestroy(resourceName string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		roleAssignments, _ := testAccGetRoleAssignments(resourceName, state)
		if roleAssignments != nil {
			return fmt.Errorf("Found role assignment: %s", roleAssignments.Identifier)
		}
		return nil
	}
}

func testAccGetRoleAssignments(resourceName string, state *terraform.State) (*nextgen.RoleAssignment, error) {
	r := acctest.TestAccGetResource(resourceName, state)
	c, ctx := acctest.TestAccGetPlatformClientWithContext()
	id := r.Primary.ID

	resp, _, err := c.RoleAssignmentsApi.GetRoleAssignment(ctx, c.AccountId, id, &nextgen.RoleAssignmentsApiGetRoleAssignmentOpts{
		OrgIdentifier:     buildField(r, "org_id"),
		ProjectIdentifier: buildField(r, "project_id"),
	})

	if err != nil {
		return nil, err
	}

	if resp.Data == nil {
		return nil, nil
	}

	return resp.Data.RoleAssignment, nil
}

func buildField(r *terraform.ResourceState, field string) optional.String {
	if attr, ok := r.Primary.Attributes[field]; ok {
		return optional.NewString(attr)
	}
	return optional.EmptyString()
}

func testAccResourceRoleAssignments(id string, name string, disabled string, accountId string) string {
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

	resource "harness_platform_service_account" "test" {
		identifier = "%[1]s"
		name = "%[2]s"
		email = "email@service.harness.io"
		description = "test"
		tags = ["foo:bar"]
		account_id = "%[4]s"
		org_id = harness_platform_project.test.org_id
		project_id = harness_platform_project.test.id
	}

	resource "harness_platform_roles" "test" {
		org_id = harness_platform_project.test.org_id
		project_id = harness_platform_project.test.id
		identifier = "%[1]s"
		name = "%[2]s"
		description = "test"
		tags = ["foo:bar"]
		permissions = ["core_pipeline_edit"]
		allowed_scope_levels = ["project"]
	}

	resource "harness_platform_role_assignments" "test"{
		identifier = "%[1]s"
		org_id = harness_platform_project.test.org_id
		project_id = harness_platform_project.test.id
		resource_group_identifier = "_all_project_level_resources"
		role_identifier = harness_platform_roles.test.id
		principal {
			identifier = harness_platform_service_account.test.id
			type = "SERVICE_ACCOUNT"
		}
		disabled = %[3]s
		managed = false
	}
	`, id, name, disabled, accountId)
}
