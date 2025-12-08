package fme_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceFMEWorkspace(t *testing.T) {
	workspaceID := os.Getenv("SPLIT_WORKSPACE_ID")
	if workspaceID == "" {
		t.Skip("SPLIT_WORKSPACE_ID environment variable must be set for this test")
	}

	workspaceName := os.Getenv("SPLIT_WORKSPACE_NAME")
	if workspaceName == "" {
		t.Skip("SPLIT_WORKSPACE_NAME environment variable must be set for this test")
	}

	resourceName := "data.harness_fme_workspace.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheckFME(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceFMEWorkspace(workspaceName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", workspaceID),
					resource.TestCheckResourceAttr(resourceName, "name", workspaceName),
					resource.TestCheckResourceAttrSet(resourceName, "type"),
				),
			},
		},
	})
}

func testAccDataSourceFMEWorkspace(workspaceName string) string {
	return fmt.Sprintf(`
		data "harness_fme_workspace" "test" {
			name = "%[1]s"
		}
`, workspaceName)
}
