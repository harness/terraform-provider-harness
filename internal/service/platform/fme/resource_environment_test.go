package fme_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/harness/terraform-provider-harness/internal/service/platform/fme"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccResourceFMEEnvironment(t *testing.T) {
	workspaceID := os.Getenv("SPLIT_WORKSPACE_ID")
	if workspaceID == "" {
		t.Skip("SPLIT_WORKSPACE_ID environment variable must be set for this test")
	}

	name := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	updatedName := fmt.Sprintf("%s_updated", name)
	resourceName := "harness_fme_environment.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheckFME(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccFMEEnvironmentDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceFMEEnvironment(workspaceID, name, false),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "production", "false"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
				),
			},
			{
				Config: testAccResourceFMEEnvironment(workspaceID, updatedName, true),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
					resource.TestCheckResourceAttr(resourceName, "production", "true"),
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

func TestAccResourceFMEEnvironment_Production(t *testing.T) {
	workspaceID := os.Getenv("SPLIT_WORKSPACE_ID")
	if workspaceID == "" {
		t.Skip("SPLIT_WORKSPACE_ID environment variable must be set for this test")
	}

	name := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	resourceName := "harness_fme_environment.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheckFME(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccFMEEnvironmentDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceFMEEnvironment(workspaceID, name, true),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "production", "true"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
				),
			},
		},
	})
}

func testAccFMEEnvironmentDestroy(resourceName string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		// Get the resource from state
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("resource not found: %s", resourceName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("no environment ID is set")
		}

		// Get session and check if environment still exists
		session := acctest.TestAccGetSession()
		if session.FMEClient == nil {
			return fmt.Errorf("FME client not configured")
		}

		workspaceID := rs.Primary.Attributes["workspace_id"]
		if workspaceID == "" {
			return fmt.Errorf("no workspace ID found in state")
		}

		c := session.FMEClient.(*fme.FMEConfig)
		env, err := c.APIClient.Environments.Get(workspaceID, rs.Primary.ID)
		if err == nil && env != nil {
			return fmt.Errorf("FME environment still exists: %s", rs.Primary.ID)
		}

		return nil
	}
}

func testAccResourceFMEEnvironment(workspaceID, name string, production bool) string {
	return fmt.Sprintf(`
		resource "harness_fme_environment" "test" {
			workspace_id = "%[1]s"
			name         = "%[2]s"
			production   = %[3]t
		}
`, workspaceID, name, production)
}