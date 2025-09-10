package image_registry_test

import (
	"fmt"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

// TestAccDataSourceChaosImageRegistry verifies the basic data source functionality for Chaos Image Registry at the account level.
func TestAccDataSourceChaosImageRegistry(t *testing.T) {
	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	rName := id
	dataSourceName := "data.harness_chaos_image_registry.test"
	resourceName := "harness_chaos_image_registry.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceChaosImageRegistryConfig(rName, id, "docker.io", "test-account"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "registry_server", resourceName, "registry_server"),
					resource.TestCheckResourceAttrPair(dataSourceName, "registry_account", resourceName, "registry_account"),
					resource.TestCheckResourceAttrPair(dataSourceName, "is_private", resourceName, "is_private"),
					resource.TestCheckResourceAttrPair(dataSourceName, "is_default", resourceName, "is_default"),
					resource.TestCheckResourceAttrSet(dataSourceName, "id"),
				),
			},
		},
	})
}

// TestAccDataSourceChaosImageRegistry_CheckOverride verifies the check_override attribute in the Chaos Image Registry data source.
func TestAccDataSourceChaosImageRegistry_CheckOverride(t *testing.T) {
	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	rName := id
	dataSourceName := "data.harness_chaos_image_registry.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceChaosImageRegistryCheckOverrideConfig(rName, id, "docker.io", "test-account"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(dataSourceName, "check_override", "true"),
					resource.TestCheckResourceAttrSet(dataSourceName, "override_blocked_by_scope"),
				),
			},
		},
	})
}

// TestAccDataSourceChaosImageRegistry_ProjectLevel verifies the data source at the project level.
func TestAccDataSourceChaosImageRegistry_ProjectLevel(t *testing.T) {
	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	rName := id
	dataSourceName := "data.harness_chaos_image_registry.test"
	resourceName := "harness_chaos_image_registry.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceChaosImageRegistryProjectLevelConfig(rName, id, "docker.io", "test-account"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "registry_server", resourceName, "registry_server"),
					resource.TestCheckResourceAttrPair(dataSourceName, "registry_account", resourceName, "registry_account"),
					resource.TestCheckResourceAttrPair(dataSourceName, "is_private", resourceName, "is_private"),
					resource.TestCheckResourceAttrPair(dataSourceName, "is_default", resourceName, "is_default"),
					resource.TestCheckResourceAttrSet(dataSourceName, "id"),
				),
			},
		},
	})
}

// Terraform Configurations

func testAccDataSourceChaosImageRegistryConfig(name, id, server, account string) string {
	return fmt.Sprintf(`
		resource "harness_chaos_image_registry" "test" {
			registry_server = "%[1]s"
			registry_account = "%[2]s"
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

		data "harness_chaos_image_registry" "test" {
			check_override = false
		}
	`, server, account)
}

func testAccDataSourceChaosImageRegistryCheckOverrideConfig(name, id, server, account string) string {
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

	data "harness_chaos_image_registry" "test" {
		org_id     = harness_platform_organization.test.id
		project_id = harness_platform_project.test.id
		check_override = true
	}
	`, name, id, server, account)
}

func testAccDataSourceChaosImageRegistryProjectLevelConfig(name, id, server, account string) string {
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

	data "harness_chaos_image_registry" "test" {
		org_id     = harness_platform_organization.test.id
		project_id = harness_platform_project.test.id
	}
	`, name, id, server, account)
}
