package dbschema_test

import (
	"fmt"
	"testing"

	"github.com/antihax/optional"
	"github.com/harness/harness-go-sdk/harness/dbops"
	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccResourceDBSchema(t *testing.T) {

	name := t.Name()
	id := fmt.Sprintf("%s_%s", name, utils.RandStringBytes(5))
	updatedName := fmt.Sprintf("%s_updated", name)
	resourceName := "harness_platform_db_schema.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccDBSchemaDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceDBSchema(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "org_id", id),
					resource.TestCheckResourceAttr(resourceName, "project_id", id),
				),
			},
			{
				Config: testAccResourceDBSchema(id, updatedName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
					resource.TestCheckResourceAttr(resourceName, "org_id", id),
					resource.TestCheckResourceAttr(resourceName, "project_id", id),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateIdFunc:       acctest.ProjectResourceImportStateIdFunc(resourceName),
				ImportStateVerifyIgnore: []string{"schema_source.#", "schema_source.0.%", "schema_source.0.connector", "schema_source.0.location", "schema_source.0.repo"},
			},
		},
	})
}

func testAccDBSchemaDestroy(resourceName string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		env, _ := testAccGetDBSchema(resourceName, state)
		if env != nil {
			return fmt.Errorf("Found environment: %s", env.Identifier)
		}

		return nil
	}
}

func testAccGetDBSchema(resourceName string, state *terraform.State) (*dbops.DbSchemaOut, error) {
	r := acctest.TestAccGetResource(resourceName, state)
	c, ctx := acctest.TestAccGetDBOpsClientWithContext()
	id := r.Primary.ID
	orgId := r.Primary.Attributes["org_id"]
	projId := r.Primary.Attributes["project_id"]

	resp, _, err := c.DatabaseSchemaApi.V1GetProjDbSchema(ctx, orgId, projId, id, &dbops.DatabaseSchemaApiV1GetProjDbSchemaOpts{
		HarnessAccount: optional.NewString(c.AccountId),
	})

	if err != nil {
		return nil, err
	}

	return &resp, nil
}

func testAccResourceDBSchema(id string, name string) string {
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
				resource "harness_platform_connector_github" "test" {
					identifier  = "%[1]s"
					name        = "%[2]s"
					description = "test"
					tags        = ["foo:bar"]
					org_id = harness_platform_project.test.org_id
					project_id = harness_platform_project.test.id
				  
					url                = "https://github.com/account"
					connection_type    = "Account"
					validation_repo    = "some_repo"
					delegate_selectors = ["harness-delegate"]
					credentials {
					  http {
						anonymous {}
					  }
					}
				  }
        resource "harness_platform_db_schema" "test" {
                        identifier = "%[1]s"
                        org_id = harness_platform_project.test.org_id
						project_id = harness_platform_project.test.id
                        name = "%[2]s"
						service = "s1"
                        tags = ["foo:bar", "bar:foo"]
                        schema_source {
							connector = "%[1]s"
							repo = "TestRepo"
							location = "db/example-changelog.yaml"
                        }
 
        }
        `, id, name)
}
