package split_test

import (
	"fmt"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

// TestAccResourceFMEApiKey_basic creates a real server-side API key in Split (destroyed at end of test).
func TestAccResourceFMEApiKey_basic(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping FME API key acceptance test in short mode")
	}
	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(6))
	envName := "tf" + testAccFMEAlphanum(10)
	keyName := "tf_acc_key_" + testAccFMEAlphanum(8)
	res := "harness_fme_api_key.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceFMEApiKey(id, envName, keyName, "server_side"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(res, "name", keyName),
					resource.TestCheckResourceAttr(res, "api_key_type", "server_side"),
					resource.TestCheckResourceAttrSet(res, "key_id"),
					resource.TestCheckResourceAttrSet(res, "api_key"),
				),
			},
			{
				ResourceName:            res,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateIdFunc:       fmeImportStateIDApiKey(res),
				ImportStateVerifyIgnore: []string{"api_key"},
			},
		},
	})
}

func testAccResourceFMEApiKey(id, envName, keyName, apiKeyType string) string {
	return fmt.Sprintf(`
	resource "harness_platform_organization" "test" {
		identifier = "%[1]s"
		name       = "%[1]s"
	}

	resource "harness_platform_project" "test" {
		identifier = "%[1]s"
		org_id     = harness_platform_organization.test.id
		name       = "%[1]s"
	}

	resource "harness_fme_environment" "test" {
		org_id     = harness_platform_organization.test.id
		project_id = harness_platform_project.test.id
		name       = "%[2]s"
	}

	resource "harness_fme_api_key" "test" {
		org_id         = harness_platform_organization.test.id
		project_id     = harness_platform_project.test.id
		name           = "%[3]s"
		api_key_type   = "%[4]s"
		environment_id = harness_fme_environment.test.environment_id
		depends_on     = [harness_fme_environment.test]
	}
	`, id, envName, keyName, apiKeyType)
}

// TestAccResourceFMEApiKey_clientSide exercises client_side key create, import, and destroy.
func TestAccResourceFMEApiKey_clientSide(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping FME API key acceptance test in short mode")
	}
	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(6))
	envName := "tf" + testAccFMEAlphanum(10)
	keyName := "tf_acc_cs_key_" + testAccFMEAlphanum(8)
	res := "harness_fme_api_key.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceFMEApiKey(id, envName, keyName, "client_side"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(res, "name", keyName),
					resource.TestCheckResourceAttr(res, "api_key_type", "client_side"),
					resource.TestCheckResourceAttrSet(res, "key_id"),
					resource.TestCheckResourceAttrSet(res, "api_key"),
				),
			},
			{
				ResourceName:            res,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateIdFunc:       fmeImportStateIDApiKey(res),
				ImportStateVerifyIgnore: []string{"api_key"},
			},
		},
	})
}
