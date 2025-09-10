package image_registry_test

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

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

func TestAccDataSourceChaosImageRegistry_MissingInfraId(t *testing.T) {
	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	rName := id

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccDataSourceChaosImageRegistryConfig_MissingInfraId(rName, id),
				ExpectError: regexp.MustCompile(`.*infra_id.*required.*`),
			},
		},
	})
}

// Terraform Configurations

func testAccDataSourceChaosImageRegistryConfig(name, id, server, account string) string {
	return fmt.Sprintf(`
resource "harness_chaos_image_registry" "test" {
  registry_server = "%s"
  registry_account = "%s"
  secret_name = "test-secret"
  is_private     = true
  is_default     = false
  use_custom_images = false
  is_override_allowed = false
}

data "harness_chaos_image_registry" "test" {
  infra_id = harness_chaos_image_registry.test.id
}
`, server, account)
}

func testAccDataSourceChaosImageRegistryCheckOverrideConfig(name, id, server, account string) string {
	return fmt.Sprintf(`
resource "harness_chaos_image_registry" "test" {
  registry_server = "%s"
  registry_account = "%s"
  is_default     = false
  secret_name = "test-secret"
  is_private     = true
  use_custom_images = false
  is_override_allowed = false
}

data "harness_chaos_image_registry" "test" {
  check_override = true
}
`, server, account)
}

func testAccDataSourceChaosImageRegistryProjectLevelConfig(name, id, server, account string) string {
	return fmt.Sprintf(`
resource "harness_chaos_image_registry" "test" {
  org_id         = "org_id"
  project_id     = "project_id"
  registry_server = "%s"
  registry_account = "%s"
  secret_name = "test-secret"
  is_private     = true
  is_default     = false
  use_custom_images = false
  is_override_allowed = false
}

data "harness_chaos_image_registry" "test" {
  org_id     = "org_id"
  project_id = "project_id"
}
`, server, account)
}

func testAccDataSourceChaosImageRegistryConfig_MissingInfraId(name, id string) string {
	return fmt.Sprintf(`
data "harness_chaos_image_registry" "test" {
  # infra_id is missing intentionally
}
`)
}
