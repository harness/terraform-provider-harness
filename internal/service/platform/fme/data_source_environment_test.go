package fme_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceFMEEnvironment(t *testing.T) {
	workspaceID := os.Getenv("SPLIT_WORKSPACE_ID")
	if workspaceID == "" {
		t.Skip("SPLIT_WORKSPACE_ID environment variable must be set for this test")
	}

	envName := fmt.Sprintf("test-env-ds-%s", utils.RandStringBytes(5))
	resourceName := "data.harness_fme_environment.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheckFME(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceFMEEnvironment(envName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", envName),
					resource.TestCheckResourceAttr(resourceName, "production", "false"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
				),
			},
		},
	})
}

func testAccDataSourceFMEEnvironment(envName string) string {
	return fmt.Sprintf(`
		resource "harness_fme_environment" "test" {
			name       = "%[1]s"
			production = false
		}

		data "harness_fme_environment" "test" {
			id = harness_fme_environment.test.id
		}
`, envName)
}