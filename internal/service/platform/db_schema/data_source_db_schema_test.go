package dbschema_test

import (
	"fmt"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceDBSchema(t *testing.T) {

	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(6))
	name := id
	resourceName := "data.harness_platform_db_schema.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceDBSchema(id, name),
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

func TestAccDataSourceDBSchemaArtifactory(t *testing.T) {

	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(6))
	name := id
	resourceName := "data.harness_platform_db_schema.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceDBSchemaArtifactory(id, name),
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

func TestAccDataSourceDBSchemaWithChangelogScript(t *testing.T) {
	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(6))
	name := id
	resourceName := "data.harness_platform_db_schema.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceDBSchemaWithChangelogScript(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "org_id", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "project_id", id),
					resource.TestCheckResourceAttr(resourceName, "type", "Script"),
				),
			},
		},
	})
}

func testAccDataSourceDBSchema(id string, name string) string {
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
						connector = "%[1]s"
						repo = "TestRepo"
						location = "db/example-changelog.yaml"
					}
				}
				data "harness_platform_db_schema" "test" {
					identifier = harness_platform_db_schema.test.id
					org_id = harness_platform_db_schema.test.org_id
					project_id = harness_platform_project.test.id
				}
	`, id, name)
}

func testAccDataSourceDBSchemaArtifactory(id string, name string) string {
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
				resource "harness_platform_connector_artifactory" "test" {
					identifier  = "%[1]s"
					name        = "%[2]s"
					org_id = harness_platform_project.test.org_id
					project_id = harness_platform_project.test.id
					description = "test"
  					tags        = ["foo:bar"]
  					url                = "https://artifactory.example.com"
  					delegate_selectors = ["harness-delegate"]
				}
        		resource "harness_platform_db_schema" "test" {
					identifier = "%[1]s"
					org_id = harness_platform_project.test.org_id
					project_id = harness_platform_project.test.id
					name = "%[2]s"
					tags = ["foo:bar", "bar:foo"]
					schema_source {
						connector = "%[1]s"
						repo = "TestRepo"
						location = "db/example-changelog.yaml"
						archive_path = "path/to/archive.zip"
					}
        		}
				data "harness_platform_db_schema" "test" {
					identifier = harness_platform_db_schema.test.id
					org_id = harness_platform_db_schema.test.org_id
					project_id = harness_platform_project.test.id
				}
	`, id, name)
}

func testAccDataSourceDBSchemaWithChangelogScript(id string, name string) string {
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
        		resource "harness_platform_db_schema" "test" {
					identifier = "%[1]s"
					org_id = harness_platform_project.test.org_id
					project_id = harness_platform_project.test.id
					name = "%[2]s"
					tags = ["foo:bar", "bar:foo"]
					type = "Script"
					changelog_script {
						image    = "alpine:latest"
						command  = "echo \\\"hello\\\""
						shell    = "Bash"
						location = "db/example.yaml"
					}
        		}
				data "harness_platform_db_schema" "test" {
					identifier = harness_platform_db_schema.test.id
					org_id = harness_platform_db_schema.test.org_id
					project_id = harness_platform_project.test.id
				}
	`, id, name)
}
