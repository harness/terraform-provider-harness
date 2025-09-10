package image_registry_test

import (
	"fmt"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccResourceChaosImageRegistry(t *testing.T) {

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
				Config: testAccResourceChaosImageRegistryConfigBasic(rName, id, "docker.io", "test-account"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "registry_server", "docker.io"),
					resource.TestCheckResourceAttr(resourceName, "registry_account", "test-account"),
					resource.TestCheckResourceAttr(resourceName, "is_private", "false"),
					resource.TestCheckResourceAttr(resourceName, "is_default", "false"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccResourceChaosImageRegistry_Update(t *testing.T) {

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
					resource.TestCheckResourceAttr(resourceName, "is_default", "true"),
					resource.TestCheckResourceAttr(resourceName, "secret_name", "test-secret"),
				),
			},
		},
	})
}

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
					resource.TestCheckResourceAttr(resourceName, "custom_images.0.log_watcher", "harness/chaos-log-watcher:1.0.0"),
					resource.TestCheckResourceAttr(resourceName, "custom_images.0.ddcr", "harness/chaos-ddcr:1.0.0"),
					resource.TestCheckResourceAttr(resourceName, "custom_images.0.ddcr_lib", "harness/chaos-ddcr-lib:1.0.0"),
					resource.TestCheckResourceAttr(resourceName, "custom_images.0.ddcr_fault", "harness/chaos-ddcr-fault:1.0.0"),
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
	resource "harness_chaos_image_registry" "test" {
		registry_server   = "%s"
		registry_account  = "%s"
		is_private       = false
		is_default       = false
	}
	`, server, account)
}

func testAccResourceChaosImageRegistryConfigUpdate(name, id, server, account string) string {

	return fmt.Sprintf(`
	resource "harness_chaos_image_registry" "test" {
		registry_server     = "%s"
		registry_account    = "%s"
		is_private         = true
		is_default         = true
		secret_name        = "test-secret"
		is_override_allowed = true
	}
	`, server, account)
}

func testAccResourceChaosImageRegistryWithCustomImages(name, id, server, account string) string {

	return fmt.Sprintf(`
	resource "harness_chaos_image_registry" "test" {
		registry_server   = "%s"
		registry_account  = "%s"
		is_private       = true
		is_default       = false
		use_custom_images = true

		custom_images {
			log_watcher = "harness/chaos-log-watcher:1.0.0"
			ddcr        = "harness/chaos-ddcr:1.0.0"
			ddcr_lib    = "harness/chaos-ddcr-lib:1.0.0"
			ddcr_fault  = "harness/chaos-ddcr-fault:1.0.0"
		}
	}
	`, server, account)
}
