package chaos_hub_v2_test

import (
	"fmt"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

// TestAccDataSourceChaosHubV2_ByIdentity verifies data source lookup by identity.
func TestAccDataSourceChaosHubV2_ByIdentity(t *testing.T) {
	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	name := id
	resourceName := "data.harness_chaos_hub_v2.test"

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceChaosHubV2ConfigByIdentity(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "identity", id),
					resource.TestCheckResourceAttrSet(resourceName, "hub_id"),
					resource.TestCheckResourceAttrSet(resourceName, "account_id"),
					resource.TestCheckResourceAttr(resourceName, "description", "Test Chaos Hub"),
				),
			},
		},
	})
}

// TestAccDataSourceChaosHubV2_ByName verifies data source lookup by name.
func TestAccDataSourceChaosHubV2_ByName(t *testing.T) {
	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	name := id
	resourceName := "data.harness_chaos_hub_v2.test"

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceChaosHubV2ConfigByName(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "identity", id),
					resource.TestCheckResourceAttrSet(resourceName, "hub_id"),
					resource.TestCheckResourceAttrSet(resourceName, "account_id"),
					resource.TestCheckResourceAttr(resourceName, "description", "Test Chaos Hub"),
				),
			},
		},
	})
}

// Terraform Configurations

func testAccDataSourceChaosHubV2ConfigByIdentity(id, name string) string {
	return fmt.Sprintf(`
		// 1. Create Organization
		resource "harness_platform_organization" "test" {
			identifier = "%[1]s"
			name       = "%[2]s"
		}

		// 2. Create Project
		resource "harness_platform_project" "test" {
			identifier = "%[1]s"
			name       = "%[2]s"
			org_id     = harness_platform_organization.test.id
		}

		// 3. Create GitHub Connector
		resource "harness_platform_connector_github" "test" {
			identifier  = "%[1]s"
			name        = "%[2]s"
			org_id      = harness_platform_organization.test.id
			project_id  = harness_platform_project.test.id
			url         = "https://github.com/harness/harness-chaos-hub"
			connection_type = "Account"
			validation_repo = "harness/harness-chaos-hub"
			
			credentials {
				http {
					username = "test"
					token_ref = "account.test_token"
				}
			}
		}

		// 4. Create Chaos Hub V2
		resource "harness_chaos_hub_v2" "test" {
			org_id        = harness_platform_organization.test.id
			project_id    = harness_platform_project.test.id
			identity      = "%[1]s"
			name          = "%[2]s"
			connector_ref = "project.${harness_platform_connector_github.test.identifier}"
			repo_branch   = "main"
			description   = "Test Chaos Hub"
			tags          = ["test:true"]
		}

		// 5. Data Source - Lookup by Identity
		data "harness_chaos_hub_v2" "test" {
			org_id     = harness_chaos_hub_v2.test.org_id
			project_id = harness_chaos_hub_v2.test.project_id
			identity   = harness_chaos_hub_v2.test.identity
		}
	`, id, name)
}

func testAccDataSourceChaosHubV2ConfigByName(id, name string) string {
	return fmt.Sprintf(`
		// 1. Create Organization
		resource "harness_platform_organization" "test" {
			identifier = "%[1]s"
			name       = "%[2]s"
		}

		// 2. Create Project
		resource "harness_platform_project" "test" {
			identifier = "%[1]s"
			name       = "%[2]s"
			org_id     = harness_platform_organization.test.id
		}

		// 3. Create GitHub Connector
		resource "harness_platform_connector_github" "test" {
			identifier  = "%[1]s"
			name        = "%[2]s"
			org_id      = harness_platform_organization.test.id
			project_id  = harness_platform_project.test.id
			url         = "https://github.com/harness/harness-chaos-hub"
			connection_type = "Account"
			validation_repo = "harness/harness-chaos-hub"
			
			credentials {
				http {
					username = "test"
					token_ref = "account.test_token"
				}
			}
		}

		// 4. Create Chaos Hub V2
		resource "harness_chaos_hub_v2" "test" {
			org_id        = harness_platform_organization.test.id
			project_id    = harness_platform_project.test.id
			identity      = "%[1]s"
			name          = "%[2]s"
			connector_ref = "project.${harness_platform_connector_github.test.identifier}"
			repo_branch   = "main"
			description   = "Test Chaos Hub"
			tags          = ["test:true"]
		}

		// 5. Data Source - Lookup by Name
		data "harness_chaos_hub_v2" "test" {
			org_id     = harness_chaos_hub_v2.test.org_id
			project_id = harness_chaos_hub_v2.test.project_id
			name       = harness_chaos_hub_v2.test.name
		}
	`, id, name)
}
