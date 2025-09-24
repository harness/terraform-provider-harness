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

func TestAccResourceFMEFlagSet(t *testing.T) {
	workspaceID := os.Getenv("SPLIT_WORKSPACE_ID")
	if workspaceID == "" {
		t.Skip("SPLIT_WORKSPACE_ID environment variable must be set for this test")
	}

	flagSetName := fmt.Sprintf("test-flagset-%s", utils.RandStringBytes(5))
	updatedDescription := "Updated test flag set description"
	resourceName := "harness_fme_flag_set.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheckFME(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccFMEFlagSetDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceFMEFlagSet(workspaceID, flagSetName, "Initial test flag set description"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "workspace_id", workspaceID),
					resource.TestCheckResourceAttr(resourceName, "name", flagSetName),
					resource.TestCheckResourceAttr(resourceName, "description", "Initial test flag set description"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
				),
			},
			{
				Config: testAccResourceFMEFlagSet(workspaceID, flagSetName, updatedDescription),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "workspace_id", workspaceID),
					resource.TestCheckResourceAttr(resourceName, "name", flagSetName),
					resource.TestCheckResourceAttr(resourceName, "description", updatedDescription),
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

func testAccFMEFlagSetDestroy(resourceName string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		// Get the resource from state
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("resource not found: %s", resourceName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("no flag set ID is set")
		}

		// Get session and check if flag set still exists
		session := acctest.TestAccGetSession()
		if session.FMEClient == nil {
			return fmt.Errorf("FME client not configured")
		}

		c := session.FMEClient.(*fme.FMEConfig)
		flagSet, err := c.APIClient.FlagSets.Get(rs.Primary.ID)
		if err == nil && flagSet != nil {
			return fmt.Errorf("FME flag set still exists: %s", rs.Primary.ID)
		}

		return nil
	}
}

func testAccResourceFMEFlagSet(workspaceID, flagSetName, description string) string {
	return fmt.Sprintf(`
		resource "harness_fme_flag_set" "test" {
			workspace_id = "%[1]s"
			name         = "%[2]s"
			description  = "%[3]s"
		}
`, workspaceID, flagSetName, description)
}