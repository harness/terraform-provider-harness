package image_registry_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceChaosImageRegistry(t *testing.T) {
	// Check for required environment variables
	accountId := os.Getenv("HARNESS_ACCOUNT_ID")
	if accountId == "" {
		t.Skip("Skipping test because HARNESS_ACCOUNT_ID is not set")
	}

	// Generate unique identifiers
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

func TestAccDataSourceChaosImageRegistry_CheckOverride(t *testing.T) {
	// Check for required environment variables
	accountId := os.Getenv("HARNESS_ACCOUNT_ID")
	if accountId == "" {
		t.Skip("Skipping test because HARNESS_ACCOUNT_ID is not set")
	}

	// Generate unique identifiers
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

func TestAccDataSourceChaosImageRegistry_ProjectLevel(t *testing.T) {
	// Check for required environment variables
	accountId := os.Getenv("HARNESS_ACCOUNT_ID")
	if accountId == "" {
		t.Skip("Skipping test because HARNESS_ACCOUNT_ID is not set")
	}

	// Generate unique identifiers
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
					resource.TestCheckResourceAttrPair(dataSourceName, "project_id", resourceName, "project_id"),
					resource.TestCheckResourceAttrPair(dataSourceName, "org_id", resourceName, "org_id"),
					resource.TestCheckResourceAttrSet(dataSourceName, "id"),
				),
			},
		},
	})
}

func testAccDataSourceChaosImageRegistryConfig(name, id, server, account string) string {
	// Use the account ID from environment variables
	accountId := os.Getenv("HARNESS_ACCOUNT_ID")
	if accountId == "" {
		accountId = "test" // Default for test cases when not set
	}

	return fmt.Sprintf(`
	resource "harness_chaos_image_registry" "test" {
		registry_server   = "%s"
		registry_account  = "%s"
		is_private       = false
		is_default       = false
	}

	data "harness_chaos_image_registry" "test" {
		registry_server  = harness_chaos_image_registry.test.registry_server
		registry_account = harness_chaos_image_registry.test.registry_account
	}
	`, server, account)
}

func testAccDataSourceChaosImageRegistryCheckOverrideConfig(name, id, server, account string) string {
	// Use the account ID from environment variables
	accountId := os.Getenv("HARNESS_ACCOUNT_ID")
	if accountId == "" {
		accountId = "test" // Default for test cases when not set
	}

	return fmt.Sprintf(`
	resource "harness_chaos_image_registry" "test" {
		registry_server   = "%s"
		registry_account  = "%s"
		is_private       = true
		is_override_allowed = true
	}

	data "harness_chaos_image_registry" "test" {
		registry_server  = harness_chaos_image_registry.test.registry_server
		registry_account = harness_chaos_image_registry.test.registry_account
		check_override  = true
	}
	`, server, account)
}

func testAccDataSourceChaosImageRegistryProjectLevelConfig(name, id, server, account string) string {
	// Use the account ID from environment variables
	accountId := os.Getenv("HARNESS_ACCOUNT_ID")
	if accountId == "" {
		accountId = "test" // Default for test cases when not set
	}

	return fmt.Sprintf(`
	resource "harness_platform_organization" "test" {
		identifier = "%[2]s"
		name       = "%[1]s"
		account_id = "%[5]s"
	}

	resource "harness_platform_project" "test" {
		identifier  = "%[2]s"
		name        = "%[1]s"
		org_id      = harness_platform_organization.test.id
		account_id  = "%[5]s"
		color       = "#0063F7"
		description = "Test project for Chaos Image Registry"
		tags        = ["foo:bar", "baz:qux"]
	}

	resource "harness_chaos_image_registry" "test" {
		org_id          = harness_platform_organization.test.id
		project_id      = harness_platform_project.test.id
		registry_server = "%[3]s"
		registry_account = "%[4]s"
		is_private     = true
		is_default     = false
	}

	data "harness_chaos_image_registry" "test" {
		org_id         = harness_platform_organization.test.id
		project_id     = harness_platform_project.test.id
		registry_server = harness_chaos_image_registry.test.registry_server
	}
	`, name, id, server, account, accountId)
}
