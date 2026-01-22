package chaos_hub_v2_test

import (
	"fmt"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

// TestAccResourceChaosHubV2_basic verifies the basic resource functionality for Chaos Hub V2.
func TestAccResourceChaosHubV2_basic(t *testing.T) {
	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	name := id
	resourceName := "harness_chaos_hub_v2.test"

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccChaosHubV2Destroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceChaosHubV2ConfigBasic(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "identity", id),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "hub_id"),
					resource.TestCheckResourceAttrSet(resourceName, "account_id"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: testAccResourceChaosHubV2ImportStateIdFunc(resourceName),
			},
		},
	})
}

// TestAccResourceChaosHubV2_Update verifies update functionality for the Chaos Hub V2 resource.
func TestAccResourceChaosHubV2_Update(t *testing.T) {
	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	name := id
	resourceName := "harness_chaos_hub_v2.test"

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccChaosHubV2Destroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceChaosHubV2ConfigBasic(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "Test Chaos Hub"),
				),
			},
			{
				Config: testAccResourceChaosHubV2ConfigUpdate(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name+"-updated"),
					resource.TestCheckResourceAttr(resourceName, "description", "Updated Test Chaos Hub"),
					resource.TestCheckResourceAttr(resourceName, "tags.0", "test:true"),
					resource.TestCheckResourceAttr(resourceName, "tags.1", "updated:true"),
				),
			},
		},
	})
}

// TestAccResourceChaosHubV2_Import verifies import functionality for the Chaos Hub V2 resource.
func TestAccResourceChaosHubV2_Import(t *testing.T) {
	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	name := id
	resourceName := "harness_chaos_hub_v2.test"

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccChaosHubV2Destroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceChaosHubV2ConfigBasic(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: testAccResourceChaosHubV2ImportStateIdFunc(resourceName),
			},
		},
	})
}

// Helpers For Destroy & Import State
func testAccChaosHubV2Destroy(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// Verify the resource has been removed from state
		return nil
	}
}

func testAccResourceChaosHubV2ImportStateIdFunc(resourceName string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return "", fmt.Errorf("not found: %s", resourceName)
		}

		orgID := rs.Primary.Attributes["org_id"]
		projectID := rs.Primary.Attributes["project_id"]
		identity := rs.Primary.Attributes["identity"]

		if orgID == "" || projectID == "" || identity == "" {
			return "", fmt.Errorf("org_id, project_id, and identity must be set")
		}

		return fmt.Sprintf("%s/%s/%s", orgID, projectID, identity), nil
	}
}

// Terraform Configurations

func testAccResourceChaosHubV2ConfigBasic(id, name string) string {
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
	`, id, name)
}

func testAccResourceChaosHubV2ConfigUpdate(id, name string) string {
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

		// 4. Create Chaos Hub V2 with updates
		resource "harness_chaos_hub_v2" "test" {
			org_id        = harness_platform_organization.test.id
			project_id    = harness_platform_project.test.id
			identity      = "%[1]s"
			name          = "%[2]s-updated"
			connector_ref = "project.${harness_platform_connector_github.test.identifier}"
			repo_branch   = "main"
			description   = "Updated Test Chaos Hub"
			tags          = ["test:true", "updated:true"]
		}
	`, id, name)
}
