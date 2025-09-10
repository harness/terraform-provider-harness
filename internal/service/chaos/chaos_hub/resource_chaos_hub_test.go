package chaos_hub_test

import (
	"fmt"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

// TestAccResourceChaosHub verifies basic create, read, and import functionality for the Chaos Hub resource.
func TestAccResourceChaosHub(t *testing.T) {
	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	rName := id
	resourceName := "harness_chaos_hub.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccChaosHubDestroy(resourceName),
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {},
		},
		Steps: []resource.TestStep{
			{
				Config: testAccResourceChaosHubConfigBasic(rName, id, "master", "https://github.com/litmuschaos/chaos-charts.git"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", rName),
					resource.TestCheckResourceAttr(resourceName, "repo_branch", "master"),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateIdFunc:       acctest.ProjectResourceImportStateIdFunc(resourceName),
				ImportStateVerifyIgnore: []string{"repo_name"},
			},
		},
	})
}

// TestAccResourceChaosHub_Update verifies update functionality for the Chaos Hub resource.
func TestAccResourceChaosHub_Update(t *testing.T) {
	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	rName := id
	updatedName := fmt.Sprintf("%s_updated", rName)
	resourceName := "harness_chaos_hub.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccChaosHubDestroy(resourceName),
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {},
		},
		Steps: []resource.TestStep{
			{
				Config: testAccResourceChaosHubConfigBasic(rName, id, "master", "https://github.com/litmuschaos/chaos-charts.git"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", rName),
				),
			},
			{
				Config: testAccResourceChaosHubConfigUpdate(updatedName, id, "v3.20.x", "https://github.com/litmuschaos/chaos-charts.git"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
					resource.TestCheckResourceAttr(resourceName, "repo_branch", "v3.20.x"),
					resource.TestCheckResourceAttr(resourceName, "description", "Updated test chaos hub"),
					resource.TestCheckResourceAttr(resourceName, "is_default", "false"),
				),
			},
		},
	})
}

// Helpers for Destroy & Import State

func testAccChaosHubDestroy(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// Implement the logic to verify the resource is properly destroyed
		return nil
	}
}

func testAccResourceChaosHubImportStateIdFunc(resourceName string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		primary := s.RootModule().Resources[resourceName].Primary
		return primary.ID, nil
	}
}

// Terraform Configurations

func testAccResourceChaosHubConfigBasic(name, id, branch, repoUrl string) string {

	return fmt.Sprintf(`
	resource "harness_platform_organization" "test" {
		identifier = "%[2]s"
		name       = "%[1]s"
	}

	resource "harness_platform_project" "test" {
		identifier  = "%[2]s"
		name        = "%[1]s"
		org_id      = harness_platform_organization.test.id
		color       = "#0063F7"
		description = "Test project for Chaos Hub"
		tags        = ["foo:bar", "baz:qux"]
	}
	
	resource "harness_platform_secret_text" "test" {
		identifier = "%[2]s"
		name = "%[1]s"
		description = "test"
		tags = ["foo:bar"]

		secret_manager_identifier = "harnessSecretManager"
		value_type = "Inline"
		value = "ghp_dummy_secret"
	}

	resource "harness_platform_connector_github" "test" {
		identifier  = "%[2]s"
		name        = "%[1]s"
		description = "Test connector"
		url         = "%[4]s"
		connection_type = "Repo"
		project_id  = harness_platform_project.test.id
		org_id      = harness_platform_organization.test.id
	
		credentials {
			http {
				username = "admin"
				token_ref = "account.${harness_platform_secret_text.test.id}"
			}
		}
		depends_on = [time_sleep.wait_4_seconds]
	}

	resource "time_sleep" "wait_4_seconds" {
		depends_on = [harness_platform_secret_text.test]
		destroy_duration = "4s"
	}

	resource "harness_chaos_hub" "test" {
		org_id        = harness_platform_organization.test.id
		project_id    = harness_platform_project.test.id
		name          = "%[1]s"
		connector_id  = harness_platform_connector_github.test.id
		connector_scope = "PROJECT"
		repo_branch   = "%[3]s"
		description   = "Test chaos hub in project"
		tags          = ["test:true", "chaos:true"]
	}
	`, name, id, branch, repoUrl)
}

func testAccResourceChaosHubConfigUpdate(name, id, branch, repoUrl string) string {

	return fmt.Sprintf(`
	resource "harness_platform_organization" "test" {
		identifier = "%[2]s"
		name       = "%[1]s"
	}

	resource "harness_platform_project" "test" {
		identifier  = "%[2]s"
		name        = "%[1]s"
		org_id      = harness_platform_organization.test.id
		color       = "#0063F7"
		description = "Test project for Chaos Hub"
		tags        = ["foo:bar", "baz:qux"]
	}
	
	resource "harness_platform_secret_text" "test" {
		identifier = "%[2]s"
		name = "%[1]s"
		description = "test"
		tags = ["foo:bar"]

		secret_manager_identifier = "harnessSecretManager"
		value_type = "Inline"
		value = "ghp_dummy_secret"
	}

	resource "harness_platform_connector_github" "test" {
		identifier  = "%[2]s"
		name        = "%[1]s"
		description = "Test connector"
		url         = "%[4]s"
		connection_type = "Repo"
		project_id  = harness_platform_project.test.id
		org_id      = harness_platform_organization.test.id
	
		credentials {
			http {
				username = "admin"
				token_ref = "account.${harness_platform_secret_text.test.id}"
			}
		}
		depends_on = [time_sleep.wait_4_seconds]
	}

	resource "time_sleep" "wait_4_seconds" {
		depends_on = [harness_platform_secret_text.test]
		destroy_duration = "4s"
	}

	resource "harness_chaos_hub" "test" {
		org_id        = harness_platform_organization.test.id
		project_id    = harness_platform_project.test.id
		name          = "%[1]s"
		connector_id  = harness_platform_connector_github.test.id
		connector_scope = "PROJECT"
		repo_branch   = "%[3]s"
		description   = "Updated test chaos hub"
		tags          = ["test:true", "chaos:true"]
	}
	`, name, id, branch, repoUrl)
}
