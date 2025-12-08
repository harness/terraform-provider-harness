package fme_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceFMEFlagSet(t *testing.T) {
	workspaceID := os.Getenv("SPLIT_WORKSPACE_ID")
	if workspaceID == "" {
		t.Skip("SPLIT_WORKSPACE_ID environment variable must be set for this test")
	}

	flagSetName := fmt.Sprintf("test-flagset-ds-%s", utils.RandStringBytes(5))
	resourceName := "data.harness_fme_flag_set.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheckFME(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceFMEFlagSet(workspaceID, flagSetName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", flagSetName),
					resource.TestCheckResourceAttr(resourceName, "description", "Data source test flag set"),
					resource.TestCheckResourceAttr(resourceName, "workspace_id", workspaceID),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
				),
			},
		},
	})
}

func testAccDataSourceFMEFlagSet(workspaceID, flagSetName string) string {
	return fmt.Sprintf(`
		resource "harness_fme_flag_set" "test" {
			workspace_id = "%[1]s"
			name         = "%[2]s"
			description  = "Data source test flag set"
		}

		data "harness_fme_flag_set" "test" {
			workspace_id = "%[1]s"
			name         = harness_fme_flag_set.test.name
		}
`, workspaceID, flagSetName)
}
