package setting_test

import (
	"fmt"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceSetting_ProjectLevel(t *testing.T) {
	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	resourceName := "harness_platform_setting.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceSettingProjectLevel(id, "repo-one", true),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", "default_repo_for_git_experience"),
					resource.TestCheckResourceAttr(resourceName, "org_id", id),
					resource.TestCheckResourceAttr(resourceName, "project_id", id),
					resource.TestCheckResourceAttr(resourceName, "value", "repo-one"),
					resource.TestCheckResourceAttr(resourceName, "allow_overrides", "true"),
					resource.TestCheckResourceAttr(resourceName, "category", "GIT_EXPERIENCE"),
					resource.TestCheckResourceAttr(resourceName, "value_type", "String"),
				),
			},
			{
				Config: testAccResourceSettingProjectLevel(id, "repo-two", true),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "value", "repo-two"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: acctest.ProjectResourceImportStateIdFunc(resourceName),
			},
		},
	})
}

func TestAccResourceSetting_OrgLevel(t *testing.T) {
	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	resourceName := "harness_platform_setting.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceSettingOrgLevel(id, "org-repo", false),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "org_id", id),
					resource.TestCheckResourceAttr(resourceName, "value", "org-repo"),
					resource.TestCheckResourceAttr(resourceName, "allow_overrides", "false"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: acctest.OrgResourceImportStateIdFunc(resourceName),
			},
		},
	})
}

func testAccResourceSettingProjectLevel(id string, value string, allowOverrides bool) string {
	return fmt.Sprintf(`
		resource "harness_platform_organization" "test" {
			identifier = "%[1]s"
			name       = "%[1]s"
		}

		resource "harness_platform_project" "test" {
			identifier = "%[1]s"
			name       = "%[1]s"
			org_id     = harness_platform_organization.test.id
		}

		resource "harness_platform_setting" "test" {
			identifier      = "default_repo_for_git_experience"
			org_id          = harness_platform_organization.test.id
			project_id      = harness_platform_project.test.id
			value           = "%[2]s"
			allow_overrides = %[3]t
		}
	`, id, value, allowOverrides)
}

func testAccResourceSettingOrgLevel(id string, value string, allowOverrides bool) string {
	return fmt.Sprintf(`
		resource "harness_platform_organization" "test" {
			identifier = "%[1]s"
			name       = "%[1]s"
		}

		resource "harness_platform_setting" "test" {
			identifier      = "default_repo_for_git_experience"
			org_id          = harness_platform_organization.test.id
			value           = "%[2]s"
			allow_overrides = %[3]t
		}
	`, id, value, allowOverrides)
}
