package dbinstance_test

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

func TestAccResourceDBInstance(t *testing.T) {

	name := t.Name()
	id := fmt.Sprintf("%s_%s", name, utils.RandStringBytes(5))
	updatedName := fmt.Sprintf("%s_updated", name)
	resourceName := "harness_platform_db_instance.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccDBInstanceDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceDBInstance(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "org_id", id),
					resource.TestCheckResourceAttr(resourceName, "project_id", id),
				),
			},
			{
				Config: testAccResourceDBInstance(id, updatedName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
					resource.TestCheckResourceAttr(resourceName, "org_id", id),
					resource.TestCheckResourceAttr(resourceName, "project_id", id),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: acctest.DBInstanceResourceImportStateIdFunc(resourceName),
			},
		},
	})
}

func TestAccResourceDBInstanceWithLiquibaseProps(t *testing.T) {
	name := t.Name()
	id := fmt.Sprintf("%s_%s", name, utils.RandStringBytes(5))
	updatedName := fmt.Sprintf("%s_updated", name)
	resourceName := "harness_platform_db_instance.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccDBInstanceDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceDBInstanceWithLiquibaseProps(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "org_id", id),
					resource.TestCheckResourceAttr(resourceName, "project_id", id),
					resource.TestCheckResourceAttr(resourceName, "branch", "feature/test"),
					resource.TestCheckResourceAttr(resourceName, "context", "test-context"),
					resource.TestCheckResourceAttr(resourceName, "liquibase_substitute_properties.db_user", "testuser"),
					resource.TestCheckResourceAttr(resourceName, "liquibase_substitute_properties.db_password", "testpass"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "2"),
				),
			},
			{
				Config: testAccResourceDBInstanceWithLiquibaseProps(id, updatedName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
					resource.TestCheckResourceAttr(resourceName, "org_id", id),
					resource.TestCheckResourceAttr(resourceName, "project_id", id),
					resource.TestCheckResourceAttr(resourceName, "branch", "feature/test"),
					resource.TestCheckResourceAttr(resourceName, "context", "test-context"),
					resource.TestCheckResourceAttr(resourceName, "liquibase_substitute_properties.db_user", "testuser"),
					resource.TestCheckResourceAttr(resourceName, "liquibase_substitute_properties.db_password", "testpass"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "2"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: acctest.DBInstanceResourceImportStateIdFunc(resourceName),
			},
		},
	})
}

func testAccDBInstanceDestroy(resourceName string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		env, _ := testAccGetDBInstance(resourceName, state)
		if env != nil {
			return fmt.Errorf("Found environment: %s", env.Identifier)
		}

		return nil
	}
}

func testAccGetDBInstance(resourceName string, state *terraform.State) (*dbops.DbInstanceOut, error) {
	r := acctest.TestAccGetResource(resourceName, state)
	c, ctx := acctest.TestAccGetDBOpsClientWithContext()
	id := r.Primary.ID
	orgId := r.Primary.Attributes["org_id"]
	projId := r.Primary.Attributes["project_id"]
	schema := r.Primary.Attributes["schema"]

	resp, _, err := c.DatabaseInstanceApi.V1GetProjDbSchemaInstance(ctx, orgId, projId, schema, id, &dbops.DatabaseInstanceApiV1GetProjDbSchemaInstanceOpts{
		HarnessAccount: optional.NewString(c.AccountId),
	})

	if err != nil {
		return nil, err
	}

	return &resp, nil
}

func testAccResourceDBInstance(id string, name string) string {
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
			depends_on = [harness_platform_organization.test]

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
			depends_on = [harness_platform_project.test]

		}
        resource "harness_platform_db_schema" "test" {
			identifier = "%[1]s"
			org_id = harness_platform_project.test.org_id
			project_id = harness_platform_project.test.id
			name = "%[2]s"
			tags = ["foo:bar", "bar:foo"]
			schema_source {
				connector = harness_platform_connector_github.test.id
				repo = "TestRepo"
				location = "db/example-changelog.yaml"
			}
			depends_on = [harness_platform_connector_github.test]
        }

		resource "harness_platform_secret_text" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
			description = "test"
			org_id = harness_platform_project.test.org_id
			project_id = harness_platform_project.test.id
			tags = ["foo:bar"]
			secret_manager_identifier = "harnessSecretManager"
			value_type = "Inline"
			value = "secret"
		  }
		
		resource "harness_platform_connector_jdbc" "test" {
			  identifier = "%[1]sjdbc"
			  name = "%[2]sjdbc"
			  description = "test"
              org_id = harness_platform_project.test.org_id
              project_id = harness_platform_project.test.id
			  tags = ["foo:bar"]
			  url = "jdbc:sqlserver://1.2.3;trustServerCertificate=true"
			  delegate_selectors = ["harness-delegate"]
			  credentials {
				username = "admin"
				password_ref = harness_platform_secret_text.test.id
			  }
		}
        resource "harness_platform_db_instance" "test" {
			identifier = "%[1]s"
			org_id = harness_platform_project.test.org_id
			project_id = harness_platform_project.test.id
			name = "%[2]s"
			tags = ["foo:bar", "bar:foo"]
			branch = "main"
			connector = harness_platform_connector_jdbc.test.id
			schema = harness_platform_db_schema.test.id
			depends_on = [harness_platform_db_schema.test]
		}
        `, id, name)
}

func testAccResourceDBInstanceWithLiquibaseProps(id string, name string) string {
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
			depends_on = [harness_platform_organization.test]
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
			depends_on = [harness_platform_project.test]
		}
        resource "harness_platform_db_schema" "test" {
			identifier = "%[1]s"
			org_id = harness_platform_project.test.org_id
			project_id = harness_platform_project.test.id
			name = "%[2]s"
			tags = ["foo:bar", "bar:foo"]
			schema_source {
				connector = harness_platform_connector_github.test.id
				repo = "TestRepo"
				location = "db/example-changelog.yaml"
			}
			depends_on = [harness_platform_connector_github.test]
        }

		resource "harness_platform_secret_text" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
			description = "test"
			org_id = harness_platform_project.test.org_id
			project_id = harness_platform_project.test.id
			tags = ["foo:bar"]
			secret_manager_identifier = "harnessSecretManager"
			value_type = "Inline"
			value = "secret"
		}
		
		resource "harness_platform_connector_jdbc" "test" {
			identifier = "%[1]sjdbc"
			name = "%[2]sjdbc"
			description = "test"
			org_id = harness_platform_project.test.org_id
			project_id = harness_platform_project.test.id
			tags = ["foo:bar"]
			url = "jdbc:sqlserver://1.2.3;trustServerCertificate=true"
			delegate_selectors = ["harness-delegate"]
			credentials {
				username = "admin"
				password_ref = harness_platform_secret_text.test.id
			}
		}
        resource "harness_platform_db_instance" "test" {
			identifier = "%[1]s"
			org_id = harness_platform_project.test.org_id
			project_id = harness_platform_project.test.id
			name = "%[2]s"
			tags = ["foo:bar", "bar:foo"]
			branch = "feature/test"
			connector = harness_platform_connector_jdbc.test.id
			schema = harness_platform_db_schema.test.id
			context = "test-context"
			liquibase_substitute_properties = {
				db_user = "testuser"
				db_password = "testpass"
			}
			depends_on = [harness_platform_db_schema.test]
		}
        `, id, name)
}
