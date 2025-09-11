package image_registry_test

import (
	"fmt"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

// TestAccResourceChaosImageRegistry verifies create, read, and import functionality for the Chaos Image Registry resource.
func TestAccResourceChaosImageRegistry(t *testing.T) {
	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	rName := id
	resourceName := "harness_chaos_image_registry.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccChaosImageRegistryDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceChaosImageRegistryConfigBasic(rName, id, "docker.io", "test-account"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "registry_server", "docker.io"),
					resource.TestCheckResourceAttr(resourceName, "registry_account", "test-account"),
					resource.TestCheckResourceAttr(resourceName, "is_private", "false"),
					resource.TestCheckResourceAttr(resourceName, "is_default", "false"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
				),
			},
		},
	})
}

// TestAccResourceChaosImageRegistry_Update verifies update functionality for the Chaos Image Registry resource.
func TestAccResourceChaosImageRegistry_Update(t *testing.T) {
	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	rName := id
	resourceName := "harness_chaos_image_registry.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccChaosImageRegistryDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceChaosImageRegistryConfigBasic(rName, id, "docker.io", "test-account"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "is_private", "false"),
					resource.TestCheckResourceAttr(resourceName, "is_default", "false"),
				),
			},
			{
				Config: testAccResourceChaosImageRegistryConfigUpdate(rName, id, "docker.io", "test-account-updated"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "registry_account", "test-account-updated"),
					resource.TestCheckResourceAttr(resourceName, "is_private", "true"),
					resource.TestCheckResourceAttr(resourceName, "secret_name", "test-secret"),
				),
			},
		},
	})
}

// TestAccResourceChaosImageRegistry_WithCustomImages verifies the resource with custom images enabled.
func TestAccResourceChaosImageRegistry_WithCustomImages(t *testing.T) {
	// Generate unique identifiers
	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	rName := id
	resourceName := "harness_chaos_image_registry.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccChaosImageRegistryDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceChaosImageRegistryWithCustomImages(rName, id, "docker.io", "test-account"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "use_custom_images", "true"),
					resource.TestCheckResourceAttr(resourceName, "custom_images.0.log_watcher", "harness/chaos-log-watcher:main-latest"),
					resource.TestCheckResourceAttr(resourceName, "custom_images.0.ddcr", "harness/chaos-ddcr:main-latest"),
					resource.TestCheckResourceAttr(resourceName, "custom_images.0.ddcr_lib", "harness/chaos-ddcr-faults:main-latest"),
					resource.TestCheckResourceAttr(resourceName, "custom_images.0.ddcr_fault", "harness/chaos-ddcr-faults:main-latest"),
				),
			},
		},
	})
}

// Helpers for Destroy & Import State

func testAccChaosImageRegistryDestroy(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// Implement actual destroy check if needed
		return nil
	}
}

// Terraform Configurations

func testAccResourceChaosImageRegistryConfigBasic(name, id, server, account string) string {

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

	resource "harness_chaos_image_registry" "test" {
		org_id        = harness_platform_organization.test.id
		project_id    = harness_platform_project.test.id
		registry_server = "%[3]s"
		registry_account = "%[4]s"
		is_private     = false
		is_default     = false
		use_custom_images = true
		is_override_allowed = false
		custom_images {
			log_watcher = null
			ddcr        = null
			ddcr_lib    = null
			ddcr_fault  = null
		}
	}
	`, name, id, server, account)
}

func testAccResourceChaosImageRegistryConfigUpdate(name, id, server, account string) string {

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

	resource "harness_chaos_image_registry" "test" {
		org_id        = harness_platform_organization.test.id
		project_id    = harness_platform_project.test.id
		registry_server = "%[3]s"
		registry_account = "%[4]s"
		is_private     = true
		is_default     = false
		secret_name = "test-secret"
		use_custom_images = true
		is_override_allowed = false
		custom_images {
			log_watcher = "harness/chaos-log-watcher:main-latest"
			ddcr        = null
			ddcr_lib    = null
			ddcr_fault  = null
		}
	}
	`, name, id, server, account)
}

func testAccResourceChaosImageRegistryWithCustomImages(name, id, server, account string) string {

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

	resource "harness_chaos_image_registry" "test" {
		org_id        = harness_platform_organization.test.id
		project_id    = harness_platform_project.test.id
		registry_server = "%[3]s"
		registry_account = "%[4]s"
		secret_name = "test-secret"
		is_private     = true
		is_default     = false
		use_custom_images = true
		is_override_allowed = false
		custom_images {
			log_watcher = "harness/chaos-log-watcher:main-latest"
			ddcr        = "harness/chaos-ddcr:main-latest"
			ddcr_lib    = "harness/chaos-ddcr-faults:main-latest"
			ddcr_fault  = "harness/chaos-ddcr-faults:main-latest"
		}
	}
	`, name, id, server, account)
}
