package split_test

import (
	"fmt"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

// TestAccDataSourceFMEEnvironment creates a Split environment, then looks it up with the data source.
// New Harness org/project workspaces do not include a default "Production" Split environment.
func TestAccDataSourceFMEEnvironment(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping FME environment acceptance test in short mode")
	}
	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(6))
	envName := "tf" + testAccFMEAlphanum(10)
	res := "harness_fme_environment.created"
	ds := "data.harness_fme_environment.test"
	var envID string

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceFMEEnvironment(id, envName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(ds, "id"),
					resource.TestCheckResourceAttrSet(ds, "environment_id"),
					resource.TestCheckResourceAttrPair(ds, "environment_id", res, "environment_id"),
					resource.TestCheckResourceAttrPair(ds, "id", res, "id"),
					resource.TestCheckResourceAttr(ds, "org_id", id),
					resource.TestCheckResourceAttr(ds, "project_id", id),
					resource.TestCheckResourceAttr(ds, "name", envName),
					testAccFMECaptureAttr(res, "environment_id", &envID),
				),
			},
			{
				Config: testAccFMEHarnessOrgProjectOnly(id),
				Check:  testAccFMEVerifyEnvironmentGone(id, id, envID),
			},
		},
	})
}

func testAccDataSourceFMEEnvironment(id, envName string) string {
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

	resource "harness_fme_environment" "created" {
		org_id      = harness_platform_organization.test.id
		project_id  = harness_platform_project.test.id
		name        = "%[2]s"
		production  = false
	}

	data "harness_fme_environment" "test" {
		org_id      = harness_platform_organization.test.id
		project_id  = harness_platform_project.test.id
		name        = harness_fme_environment.created.name
		depends_on  = [harness_fme_environment.created]
	}
	`, id, envName)
}
