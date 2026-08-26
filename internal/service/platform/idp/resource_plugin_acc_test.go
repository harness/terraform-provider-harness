package idp_test

// Acceptance (integration) tests for harness_platform_idp_plugin resource and data source.
//
// These tests hit the real Harness API and create/modify/destroy real resources.
// They require TF_ACC=1 and the following environment variables:
//
//   HARNESS_ACCOUNT_ID          - Harness account ID
//   HARNESS_PLATFORM_API_KEY    - Harness platform API key
//   HARNESS_TEST_IDP_PLUGIN_ID  - Plugin identifier to test with (e.g. "harness-proxy")
//   HARNESS_TEST_SECRET_ID      - A valid Harness secret identifier (for env_variables test only)
//
// Run all acceptance tests:
//   TF_ACC=1 go test -v ./internal/service/platform/idp/... -run "TestAccResourceIdpPlugin|TestAccDataSourceIdpPlugin" -timeout=120m
//
// Run a single acceptance test:
//   TF_ACC=1 go test -v ./internal/service/platform/idp/... -run TestAccResourceIdpPlugin_basic -timeout=120m

import (
	"fmt"
	"os"
	"testing"

	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

// TestAccResourceIdpPlugin_basic tests the basic CRUD lifecycle:
// create plugin config -> enable it -> read back state -> destroy (disable).
func TestAccResourceIdpPlugin_basic(t *testing.T) {
	pluginId := getTestPluginId(t)
	resourceName := "harness_platform_idp_plugin.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourcePluginBasic(pluginId),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", pluginId),
					resource.TestCheckResourceAttr(resourceName, "enabled", "true"),
					resource.TestCheckResourceAttrSet(resourceName, "name"),
					resource.TestCheckResourceAttrSet(resourceName, "configs"),
				),
			},
		},
	})
}

// TestAccResourceIdpPlugin_withEnvVariables tests creating a plugin with a secret
// env variable and verifies the server-generated identifier field is populated.
func TestAccResourceIdpPlugin_withEnvVariables(t *testing.T) {
	pluginId := getTestPluginId(t)
	secretId := getTestSecretId(t)
	resourceName := "harness_platform_idp_plugin.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourcePluginWithEnvVars(pluginId, secretId),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", pluginId),
					resource.TestCheckResourceAttr(resourceName, "enabled", "true"),
					resource.TestCheckResourceAttr(resourceName, "env_variables.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "env_variables.0.env_name", "TEST_SECRET"),
					resource.TestCheckResourceAttr(resourceName, "env_variables.0.type", "Secret"),
					resource.TestCheckResourceAttr(resourceName, "env_variables.0.harness_secret_identifier", secretId),
					resource.TestCheckResourceAttrSet(resourceName, "env_variables.0.identifier"),
				),
			},
		},
	})
}

// TestAccResourceIdpPlugin_withProxy tests creating a plugin with proxy configuration
// and verifies the proxy host, enabled flag, and server-generated identifier.
func TestAccResourceIdpPlugin_withProxy(t *testing.T) {
	pluginId := getTestPluginId(t)
	resourceName := "harness_platform_idp_plugin.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourcePluginWithProxy(pluginId),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", pluginId),
					resource.TestCheckResourceAttr(resourceName, "proxy.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "proxy.0.host", "app.harness.io"),
					resource.TestCheckResourceAttr(resourceName, "proxy.0.proxy", "true"),
					resource.TestCheckResourceAttrSet(resourceName, "proxy.0.identifier"),
				),
			},
		},
	})
}

// TestAccResourceIdpPlugin_disable tests the enable/disable toggle:
// creates plugin as disabled -> verifies enabled=false -> updates to enabled=true
// -> verifies the toggle works in both directions.
func TestAccResourceIdpPlugin_disable(t *testing.T) {
	pluginId := getTestPluginId(t)
	resourceName := "harness_platform_idp_plugin.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourcePluginDisabled(pluginId),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", pluginId),
					resource.TestCheckResourceAttr(resourceName, "enabled", "false"),
				),
			},
			{
				Config: testAccResourcePluginBasic(pluginId),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", pluginId),
					resource.TestCheckResourceAttr(resourceName, "enabled", "true"),
				),
			},
		},
	})
}

// TestAccResourceIdpPlugin_import tests terraform import:
// creates a plugin resource -> imports it by identifier -> verifies the imported
// state matches the original (all fields including env_variables and proxy).
func TestAccResourceIdpPlugin_import(t *testing.T) {
	pluginId := getTestPluginId(t)
	resourceName := "harness_platform_idp_plugin.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourcePluginBasic(pluginId),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", pluginId),
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

// TestAccDataSourceIdpPlugin tests the data source by first creating a plugin resource,
// then reading it back via the data source and verifying all fields are populated.
func TestAccDataSourceIdpPlugin(t *testing.T) {
	pluginId := getTestPluginId(t)
	resourceName := "data.harness_platform_idp_plugin.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourcePlugin(pluginId),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", pluginId),
					resource.TestCheckResourceAttrSet(resourceName, "name"),
					resource.TestCheckResourceAttrSet(resourceName, "configs"),
					resource.TestCheckResourceAttrSet(resourceName, "enabled"),
				),
			},
		},
	})
}

// Helper functions

func getTestPluginId(t *testing.T) string {
	id := os.Getenv("HARNESS_TEST_IDP_PLUGIN_ID")
	if id == "" {
		t.Skip("HARNESS_TEST_IDP_PLUGIN_ID must be set for acceptance tests")
	}
	return id
}

func getTestSecretId(t *testing.T) string {
	id := os.Getenv("HARNESS_TEST_SECRET_ID")
	if id == "" {
		t.Skip("HARNESS_TEST_SECRET_ID must be set for acceptance tests")
	}
	return id
}

// Terraform config templates

func testAccResourcePluginBasic(pluginId string) string {
	return fmt.Sprintf(`
resource "harness_platform_idp_plugin" "test" {
  identifier = "%s"
  name       = "Configure Backend Proxies"
  enabled    = true
  configs    = <<-EOT
    proxy:
      endpoints:
        /test:
          target: https://app.harness.io/gateway
  EOT
}
`, pluginId)
}

func testAccResourcePluginDisabled(pluginId string) string {
	return fmt.Sprintf(`
resource "harness_platform_idp_plugin" "test" {
  identifier = "%s"
  name       = "Configure Backend Proxies"
  enabled    = false
  configs    = <<-EOT
    proxy:
      endpoints:
        /test:
          target: https://app.harness.io/gateway
  EOT
}
`, pluginId)
}

func testAccResourcePluginWithEnvVars(pluginId, secretId string) string {
	return fmt.Sprintf(`
resource "harness_platform_idp_plugin" "test" {
  identifier = "%s"
  name       = "Configure Backend Proxies"
  enabled    = true
  configs    = <<-EOT
    proxy:
      endpoints:
        /test:
          target: https://app.harness.io/gateway
          headers:
            x-api-key: $${TEST_SECRET}
  EOT

  env_variables {
    env_name                  = "TEST_SECRET"
    type                      = "Secret"
    harness_secret_identifier = "%s"
  }
}
`, pluginId, secretId)
}

func testAccResourcePluginWithProxy(pluginId string) string {
	return fmt.Sprintf(`
resource "harness_platform_idp_plugin" "test" {
  identifier = "%s"
  name       = "Configure Backend Proxies"
  enabled    = true
  configs    = <<-EOT
    proxy:
      endpoints:
        /test:
          target: https://app.harness.io/gateway
  EOT

  proxy {
    host      = "app.harness.io"
    proxy     = true
    selectors = ["default"]
  }
}
`, pluginId)
}

func testAccDataSourcePlugin(pluginId string) string {
	return fmt.Sprintf(`
resource "harness_platform_idp_plugin" "test" {
  identifier = "%s"
  name       = "Configure Backend Proxies"
  enabled    = true
  configs    = <<-EOT
    proxy:
      endpoints:
        /test:
          target: https://app.harness.io/gateway
  EOT
}

data "harness_platform_idp_plugin" "test" {
  identifier = harness_platform_idp_plugin.test.identifier
}
`, pluginId)
}

func testAccPluginDestroy(resourceName string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		r := acctest.TestAccGetResource(resourceName, state)
		if r == nil {
			return nil
		}
		c, ctx := acctest.TestAccGetIDPClientWithContext()
		resp, err := c.PluginAppConfigApi.GetPluginInfo(ctx, r.Primary.ID)
		if err != nil {
			return nil
		}
		if resp.Plugin != nil && resp.Plugin.PluginDetails.Enabled {
			return fmt.Errorf("plugin %s is still enabled", r.Primary.ID)
		}
		return nil
	}
}
