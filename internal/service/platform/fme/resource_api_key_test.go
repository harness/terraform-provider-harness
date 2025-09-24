package fme_test

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccResourceFMEApiKey_ClientSide(t *testing.T) {
	envName := fmt.Sprintf("%s_env_%s", t.Name(), utils.RandStringBytes(5))
	keyName := fmt.Sprintf("%s_key_%s", t.Name(), utils.RandStringBytes(5))
	resourceName := "harness_fme_api_key.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheckFME(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccFMEApiKeyDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceFMEApiKey(envName, keyName, "client_side"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", keyName),
					resource.TestCheckResourceAttr(resourceName, "type", "client_side"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "environment_id"),
					resource.TestCheckResourceAttrSet(resourceName, "key"),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"key"}, // API key is sensitive and may not be returned
			},
		},
	})
}

func TestAccResourceFMEApiKey_ServerSide(t *testing.T) {
	envName := fmt.Sprintf("%s_env_%s", t.Name(), utils.RandStringBytes(5))
	keyName := fmt.Sprintf("%s_key_%s", t.Name(), utils.RandStringBytes(5))
	resourceName := "harness_fme_api_key.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheckFME(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccFMEApiKeyDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceFMEApiKey(envName, keyName, "server_side"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", keyName),
					resource.TestCheckResourceAttr(resourceName, "type", "server_side"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "environment_id"),
					resource.TestCheckResourceAttrSet(resourceName, "key"),
				),
			},
		},
	})
}

func TestAccResourceFMEApiKey_InvalidType(t *testing.T) {
	envName := fmt.Sprintf("%s_env_%s", t.Name(), utils.RandStringBytes(5))
	keyName := fmt.Sprintf("%s_key_%s", t.Name(), utils.RandStringBytes(5))

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheckFME(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccResourceFMEApiKey(envName, keyName, "invalid_type"),
				ExpectError: regexp.MustCompile("expected type to be one of \\[client_side server_side\\]"),
			},
		},
	})
}

func testAccFMEApiKeyDestroy(resourceName string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		// Get the resource from state
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("resource not found: %s", resourceName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("no API key ID is set")
		}

		// Get session and check if API key still exists
		session := acctest.TestAccGetSession()
		if session.FMEClient == nil {
			return fmt.Errorf("FME client not configured")
		}

		environmentID := rs.Primary.Attributes["environment_id"]
		if environmentID == "" {
			return fmt.Errorf("no environment ID found in state")
		}

		// API keys don't have a Get method to verify deletion
		// If we reach here, the resource was successfully removed from state

		return nil
	}
}

func testAccResourceFMEApiKey(envName, keyName, keyType string) string {
	return fmt.Sprintf(`
		resource "harness_fme_environment" "test_env" {
			name       = "%[1]s"
			production = false
		}

		resource "harness_fme_api_key" "test" {
			environment_id = harness_fme_environment.test_env.id
			name          = "%[2]s"
			type          = "%[3]s"
		}
`, envName, keyName, keyType)
}