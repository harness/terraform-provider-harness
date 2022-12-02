package usergroup_test

import (
	"fmt"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceUserGroup(t *testing.T) {

	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(6))
	name := id
	resourceName := "data.harness_platform_usergroup.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceUserGroup(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "org_id", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "project_id", id),
				),
			},
		},
	})
}

func TestAccDataSourceUserGroupAccountLevel(t *testing.T) {

	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(6))
	name := id
	resourceName := "data.harness_platform_usergroup.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceUserGroupAccountLevel(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
				),
			},
		},
	})
}

func TestAccDataSourceUserGroupByName(t *testing.T) {

	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(6))
	name := id
	resourceName := "data.harness_platform_usergroup.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceUserGroupByName(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "org_id", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "project_id", id),
				),
			},
		},
	})
}

func TestAccDataSourceUserGroupByNameAccountLevel(t *testing.T) {

	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(6))
	name := id
	resourceName := "data.harness_platform_usergroup.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceUserGroupByNameAccountLevel(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
				),
			},
		},
	})
}

func testAccDataSourceUserGroupAccountLevel(id string, name string) string {
	return fmt.Sprintf(`
		resource "harness_platform_usergroup" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
		}

		data "harness_platform_usergroup" "test" {
			identifier = harness_platform_usergroup.test.identifier
		}
`, id, name)
}

func testAccDataSourceUserGroup(id string, name string) string {
	return fmt.Sprintf(`
		resource "harness_platform_organization" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
		}

		resource "harness_platform_project" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
			org_id = harness_platform_organization.test.id
			color = "#472848"
		}

		resource "harness_platform_usergroup" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
			org_id = harness_platform_project.test.org_id
			project_id = harness_platform_project.test.id
		}

		data "harness_platform_usergroup" "test" {
			identifier = harness_platform_usergroup.test.identifier
			org_id = harness_platform_usergroup.test.org_id
			project_id = harness_platform_usergroup.test.project_id
		}
`, id, name)
}

func testAccDataSourceUserGroupByNameAccountLevel(id string, name string) string {
	return fmt.Sprintf(`
		resource "harness_platform_usergroup" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
			linked_sso_id = "linked_sso_id"
			externally_managed = false
			users = []
			notification_configs {
				type = "SLACK"
				slack_webhook_url = "https://google.com"
			}
			notification_configs {
				type = "EMAIL"
				group_email = "email@email.com"
				send_email_to_all_users = true
			}
			notification_configs {
				type = "MSTEAMS"
				microsoft_teams_webhook_url = "https://google.com"
			}
			notification_configs {
				type = "PAGERDUTY"
				pager_duty_key = "pagerDutyKey"
			}
			linked_sso_display_name = "linked_sso_display_name"
			sso_group_id = "sso_group_id"
			sso_group_name = "sso_group_name"
			linked_sso_type = "SAML"
			sso_linked = true
		}

		data "harness_platform_usergroup" "test" {
			name = harness_platform_usergroup.test.name
		}
`, id, name)
}

func testAccDataSourceUserGroupByName(id string, name string) string {
	return fmt.Sprintf(`
	resource "harness_platform_organization" "test" {
		identifier = "%[1]s"
		name = "%[2]s"
	}

	resource "harness_platform_project" "test" {
		identifier = "%[1]s"
		name = "%[2]s"
		org_id = harness_platform_organization.test.id
		color = "#472848"
	}
	
		resource "harness_platform_usergroup" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
			org_id = harness_platform_project.test.org_id
			project_id = harness_platform_project.test.id
			linked_sso_id = "linked_sso_id"
			externally_managed = false
			users = []
			notification_configs {
				type = "SLACK"
				slack_webhook_url = "https://google.com"
			}
			notification_configs {
				type = "EMAIL"
				group_email = "email@email.com"
				send_email_to_all_users = true
			}
			notification_configs {
				type = "MSTEAMS"
				microsoft_teams_webhook_url = "https://google.com"
			}
			notification_configs {
				type = "PAGERDUTY"
				pager_duty_key = "pagerDutyKey"
			}
			linked_sso_display_name = "linked_sso_display_name"
			sso_group_id = "sso_group_id"
			sso_group_name = "sso_group_name"
			linked_sso_type = "SAML"
			sso_linked = true
		}

		data "harness_platform_usergroup" "test" {
			name = harness_platform_usergroup.test.name
			org_id = harness_platform_usergroup.test.org_id
			project_id = harness_platform_usergroup.test.project_id
		}
`, id, name)
}
