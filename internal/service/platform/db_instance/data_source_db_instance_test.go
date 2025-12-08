package dbinstance_test

import (
	"fmt"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceDBInstance(t *testing.T) {

	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(6))
	name := id
	resourceName := "data.harness_platform_db_instance.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceDBInstance(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "org_id", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "project_id", id),
				),
			},
		},
	})
}

func testAccDataSourceDBInstance(id string, name string) string {
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
			tags = ["foo:bar", "bar:foo"]
			schema_source {
				connector = harness_platform_connector_github.test.id
				repo = "TestRepo"
				location = "db/example-changelog.yaml"
			}
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
		}
		data "harness_platform_db_instance" "test" {
			identifier = harness_platform_db_instance.test.id
			org_id = harness_platform_organization.test.id
			project_id = harness_platform_project.test.id
			schema = harness_platform_db_schema.test.id
		}
        `, id, name)
}

func TestAccDataSourceDBInstanceWithSubstituteProps(t *testing.T) {
	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(6))
	name := id
	resourceName := "data.harness_platform_db_instance.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceDBInstanceWithSubstituteProps(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "org_id", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "project_id", id),
					resource.TestCheckResourceAttr(resourceName, "branch", "feature/test"),
					resource.TestCheckResourceAttr(resourceName, "context", "test-context"),
					resource.TestCheckResourceAttr(resourceName, "substitute_properties.db_user", "testuser"),
					resource.TestCheckResourceAttr(resourceName, "substitute_properties.db_password", "testpass"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "2"),
				),
			},
		},
	})
}

func testAccDataSourceDBInstanceWithSubstituteProps(id string, name string) string {
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
			tags = ["foo:bar", "bar:foo"]
			schema_source {
				connector = harness_platform_connector_github.test.id
				repo = "TestRepo"
				location = "db/example-changelog.yaml"
			}
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
			substitute_properties = {
				db_user = "testuser"
				db_password = "testpass"
			}
		}
		data "harness_platform_db_instance" "test" {
			identifier = harness_platform_db_instance.test.id
			org_id = harness_platform_organization.test.id
			project_id = harness_platform_project.test.id
			schema = harness_platform_db_schema.test.id
		}
        `, id, name)
}
