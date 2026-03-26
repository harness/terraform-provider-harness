package split_test

import (
	"fmt"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

// TestAccDataSourceFMEWorkspace_ByOrgProject verifies lookup by org_id and project_id after creating
// a Harness organization and project (FME workspace is provisioned for the project).
func TestAccDataSourceFMEWorkspace_ByOrgProject(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping FME workspace acceptance test in short mode")
	}
	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(6))
	byProject := "data.harness_fme_workspace.by_project"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceFMEWorkspaceByOrgProject(id),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(byProject, "id"),
					resource.TestCheckResourceAttrSet(byProject, "workspace_id"),
					resource.TestCheckResourceAttr(byProject, "org_id", id),
					resource.TestCheckResourceAttr(byProject, "project_id", id),
					resource.TestCheckResourceAttrSet(byProject, "name"),
					resource.TestCheckResourceAttrSet(byProject, "type"),
				),
			},
		},
	})
}

// TestAccDataSourceFMEWorkspace_ByName verifies lookup by Split workspace name using the name
// returned from the org/project lookup in the same configuration.
func TestAccDataSourceFMEWorkspace_ByName(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping FME workspace acceptance test in short mode")
	}
	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(6))
	byProject := "data.harness_fme_workspace.by_project"
	byName := "data.harness_fme_workspace.by_name"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceFMEWorkspaceByName(id),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(byName, "id", byProject, "id"),
					resource.TestCheckResourceAttrPair(byName, "workspace_id", byProject, "workspace_id"),
					resource.TestCheckResourceAttrPair(byName, "name", byProject, "name"),
					resource.TestCheckResourceAttrPair(byName, "org_id", byProject, "org_id"),
					resource.TestCheckResourceAttrPair(byName, "project_id", byProject, "project_id"),
				),
			},
		},
	})
}

func testAccDataSourceFMEWorkspaceByOrgProject(id string) string {
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

	data "harness_fme_workspace" "by_project" {
		org_id     = harness_platform_organization.test.id
		project_id = harness_platform_project.test.id
	}
	`, id)
}

func testAccDataSourceFMEWorkspaceByName(id string) string {
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

	data "harness_fme_workspace" "by_project" {
		org_id     = harness_platform_organization.test.id
		project_id = harness_platform_project.test.id
	}

	data "harness_fme_workspace" "by_name" {
		name = data.harness_fme_workspace.by_project.name
	}
	`, id)
}
