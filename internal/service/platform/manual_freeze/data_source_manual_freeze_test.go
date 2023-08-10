package manual_freeze_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceResourceGroup(t *testing.T) {
	var (
		name         = fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))
		resourceName = "data.harness_platform_manual_freeze.test"
		accountId    = os.Getenv("HARNESS_ACCOUNT_ID")
	)
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceManualFreeze(name, accountId),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", name),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "account_id", accountId),
					resource.TestCheckResourceAttr(resourceName, "scope", "project"),
					resource.TestCheckResourceAttr(resourceName, "freeze_windows.0.recurrence.0.type", "Daily"),
				),
			},
			{
				Config: testAccDataSourceManualFreezeQuarterly(name, accountId),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", name),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "account_id", accountId),
					resource.TestCheckResourceAttr(resourceName, "scope", "project"),
					resource.TestCheckResourceAttr(resourceName, "freeze_windows.0.recurrence.0.type", "Monthly"),
					resource.TestCheckResourceAttr(resourceName, "freeze_windows.0.recurrence.0.recurrence_spec.0.value", "3"),
					resource.TestCheckResourceAttr(resourceName, "freeze_windows.0.recurrence.0.recurrence_spec.0.until", "2023-12-31 11:59 PM"),
				),
			},
		},
	})
}

func testAccDataSourceManualFreeze(name string, accountId string) string {
	return fmt.Sprintf(`
	resource "harness_platform_organization" "test" {
		identifier = "%[1]s"
		name = "%[1]s"
	}

	resource "harness_platform_project" "test" {
		identifier = "%[1]s"
		name = "%[1]s"
		color = "#0063F7"
		org_id = harness_platform_organization.test.identifier
	}
	
		resource "harness_platform_manual_freeze" "test" {
			identifier = "%[1]s"
			account_id = "%[2]s"
			org_id = harness_platform_project.test.org_id
			project_id = harness_platform_project.test.id
      yaml = <<-EOT
      freeze:
        name: %[1]s
        identifier: %[1]s
        entityConfigs:
          - name: r1
            entities:
              - filterType: All
                type: Org
              - filterType: All
                type: Project
              - filterType: All
                type: Service
              - filterType: All
                type: EnvType
        status: Disabled
        description: hi
        windows:
        - timeZone: Asia/Calcutta
          startTime: 2023-05-03 04:16 PM
          duration: 30m
          recurrence:
            type: Daily
        notificationRules: []
        orgIdentifier: %[1]s
        projectIdentifier: %[1]s
        tags: {}

      EOT
		}
	
		data "harness_platform_manual_freeze" "test" {
			identifier = harness_platform_manual_freeze.test.identifier
			account_id = harness_platform_manual_freeze.test.account_id
			org_id = harness_platform_manual_freeze.test.org_id
      project_id = harness_platform_manual_freeze.test.project_id
		}
	`, name, accountId)
}

func testAccDataSourceManualFreezeQuarterly(name string, accountId string) string {
	return fmt.Sprintf(`
	resource "harness_platform_organization" "test" {
		identifier = "%[1]s"
		name = "%[1]s"
	}

	resource "harness_platform_project" "test" {
		identifier = "%[1]s"
		name = "%[1]s"
		color = "#0063F7"
		org_id = harness_platform_organization.test.identifier
	}
	
		resource "harness_platform_manual_freeze" "test" {
			identifier = "%[1]s"
			account_id = "%[2]s"
			org_id = harness_platform_project.test.org_id
			project_id = harness_platform_project.test.id
      yaml = <<-EOT
      freeze:
        name: %[1]s
        identifier: %[1]s
        entityConfigs:
          - name: r1
            entities:
              - filterType: All
                type: Org
              - filterType: All
                type: Project
              - filterType: All
                type: Service
              - filterType: All
                type: EnvType
        status: Disabled
        description: hi
        windows:
        - timeZone: Asia/Calcutta
          startTime: 2023-05-03 04:16 PM
          duration: 30m
          recurrence:
            type: Monthly
            spec:
              value: 3
              until: 2023-12-31 11:59 PM
        notificationRules: []
        orgIdentifier: %[1]s
        projectIdentifier: %[1]s
        tags: {}

      EOT
		}
	
		data "harness_platform_manual_freeze" "test" {
			identifier = harness_platform_manual_freeze.test.identifier
			account_id = harness_platform_manual_freeze.test.account_id
			org_id = harness_platform_manual_freeze.test.org_id
      project_id = harness_platform_manual_freeze.test.project_id
		}
	`, name, accountId)
}
