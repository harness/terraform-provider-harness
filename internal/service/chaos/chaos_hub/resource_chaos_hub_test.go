package chaos_hub_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccResourceChaosHub(t *testing.T) {
	// Check for required environment variables
	accountId := os.Getenv("HARNESS_ACCOUNT_ID")
	if accountId == "" {
		t.Skip("Skipping test because HARNESS_ACCOUNT_ID is not set")
	}

	// Generate unique identifiers
	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	rName := id
	resourceName := "harness_chaos_hub.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccChaosHubDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceChaosHubConfigBasic(rName, id, "main", "https://github.com/harness/chaos-hub.git"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", rName),
					resource.TestCheckResourceAttr(resourceName, "repo_branch", "main"),
					resource.TestCheckResourceAttr(resourceName, "connector_scope", "ACCOUNT"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: testAccResourceChaosHubImportStateIdFunc(resourceName),
			},
		},
	})
}

func TestAccResourceChaosHub_Update(t *testing.T) {
	// Check for required environment variables
	accountId := os.Getenv("HARNESS_ACCOUNT_ID")
	if accountId == "" {
		t.Skip("Skipping test because HARNESS_ACCOUNT_ID is not set")
	}

	// Generate unique identifiers
	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	rName := id
	updatedName := fmt.Sprintf("%s_updated", rName)
	resourceName := "harness_chaos_hub.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccChaosHubDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceChaosHubConfigBasic(rName, id, "main", "https://github.com/harness/chaos-hub.git"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", rName),
				),
			},
			{
				Config: testAccResourceChaosHubConfigUpdate(updatedName, id, "develop", "https://github.com/harness/chaos-hub.git"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
					resource.TestCheckResourceAttr(resourceName, "repo_branch", "develop"),
					resource.TestCheckResourceAttr(resourceName, "description", "Updated test chaos hub"),
					resource.TestCheckResourceAttr(resourceName, "is_default", "true"),
				),
			},
		},
	})
}

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

func testAccResourceChaosHubConfigBasic(name, id, branch, repoUrl string) string {
	// Use the account ID from environment variables
	accountId := os.Getenv("HARNESS_ACCOUNT_ID")
	if accountId == "" {
		accountId = "test" // Default for test cases when not set
	}

	return fmt.Sprintf(`
	resource "harness_platform_connector_github" "test" {
		identifier  = "%[2]s"
		name        = "%[1]s"
		description = "Test connector"
		url         = "https://github.com"
		connection_type = "Account"
		validation_repo = "harness/chaos-hub"
	
		credentials {
			http {
				username  = "test"
				token_ref = "account.do_not_delete_harness_platform_qa_token"
			}
		}
	}

	resource "harness_chaos_hub" "test" {
		name           = "%[1]s"
		connector_id   = harness_platform_connector_github.test.id
		connector_scope = "ACCOUNT"
		repo_branch    = "%[3]s"
		description    = "Test chaos hub"
		tags           = ["test:true", "chaos:true"]
	}
	`, name, id, branch)
}

func testAccResourceChaosHubConfigUpdate(name, id, branch, repoUrl string) string {
	// Use the account ID from environment variables
	accountId := os.Getenv("HARNESS_ACCOUNT_ID")
	if accountId == "" {
		accountId = "test" // Default for test cases when not set
	}

	return fmt.Sprintf(`
	resource "harness_platform_connector_github" "test" {
		identifier  = "%[2]s"
		name        = "%[1]s"
		description = "Test connector"
		url         = "https://github.com"
		connection_type = "Account"
		validation_repo = "harness/chaos-hub"
	
		credentials {
			http {
				username  = "test"
				token_ref = "account.do_not_delete_harness_platform_qa_token"
			}
		}
	}

	resource "harness_chaos_hub" "test" {
		name           = "%[1]s"
		connector_id   = harness_platform_connector_github.test.id
		connector_scope = "ACCOUNT"
		repo_branch    = "%[3]s"
		description    = "Updated test chaos hub"
		is_default     = true
		tags           = ["test:true", "chaos:true", "updated:true"]
	}
	`, name, id, branch)
}
