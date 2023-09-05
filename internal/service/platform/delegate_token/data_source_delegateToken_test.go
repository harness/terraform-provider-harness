package delegatetoken_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceDelegateToken(t *testing.T) {
	name := utils.RandStringBytes(5)
	account_id := os.Getenv("HARNESS_ACCOUNT_ID")

	resourceName := "data.harness_platform_delegatetoken.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: tesAccDataSourceDelegateToken(name, account_id),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "token_status", "ACTIVE"),
				),
			},
		},
	})
}

func TestAccDataSourceDelegateTokenOrgLevel(t *testing.T) {
	name := utils.RandStringBytes(5)
	account_id := os.Getenv("HARNESS_ACCOUNT_ID")

	resourceName := "data.harness_platform_delegatetoken.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: tesAccDataSourceDelegateTokenOrgLevel(name, account_id),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "token_status", "ACTIVE"),
					resource.TestCheckResourceAttr(resourceName, "org_id", name),
				),
			},
		},
	})
}

func TestAccDataSourceDelegateTokenProjectLevel(t *testing.T) {
	name := utils.RandStringBytes(5)
	account_id := os.Getenv("HARNESS_ACCOUNT_ID")

	resourceName := "data.harness_platform_delegatetoken.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: tesAccDataSourceDelegateTokenProjectLevel(name, account_id),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "token_status", "ACTIVE"),
					resource.TestCheckResourceAttr(resourceName, "org_id", name),
					resource.TestCheckResourceAttr(resourceName, "project_id", name),
				),
			},
		},
	})
}

func tesAccDataSourceDelegateToken(name string, accountId string) string {
	return fmt.Sprintf(`
		resource "harness_platform_delegatetoken" "test" {
			identifier = "%[1]s"
			name = "%[1]s"
			account_id = "%[2]s"
			token_status = "ACTIVE"
		}

		data "harness_platform_delegatetoken" "test" {
			name = harness_platform_delegatetoken.test.name
			account_id = harness_platform_delegatetoken.test.account_id
		}
	`, name, accountId)
}

func tesAccDataSourceDelegateTokenOrgLevel(name string, accountId string) string {
	return fmt.Sprintf(`
		resource "harness_platform_organization" "test" {
			identifier = "%[1]s"
			name = "%[1]s"
		}

		resource "harness_platform_delegatetoken" "test" {
			identifier = "%[1]s"
			name = "%[1]s"
			account_id = "%[2]s"
			token_status = "ACTIVE"
			org_id = harness_platform_organization.test.id
		}

		data "harness_platform_delegatetoken" "test" {
			name = harness_platform_delegatetoken.test.name
			account_id = harness_platform_delegatetoken.test.account_id
			org_id = harness_platform_delegatetoken.test.org_id
		}
	`, name, accountId)
}

func tesAccDataSourceDelegateTokenProjectLevel(name string, accountId string) string {
	return fmt.Sprintf(`
		resource "harness_platform_organization" "test" {
			identifier = "%[1]s"
			name = "%[1]s"
		}

		resource "harness_platform_project" "test" {
			identifier = "%[1]s"
			name = "%[1]s"
			org_id = harness_platform_organization.test.id
			color = "#472848"
		}

		resource "harness_platform_delegatetoken" "test" {
			identifier = "%[1]s"
			name = "%[1]s"
			account_id = "%[2]s"
			token_status = "ACTIVE"
			org_id = harness_platform_organization.test.id
			project_id = harness_platform_project.test.id
		}

		data "harness_platform_delegatetoken" "test" {
			name = harness_platform_delegatetoken.test.name
			account_id = harness_platform_delegatetoken.test.account_id
			org_id = harness_platform_delegatetoken.test.org_id
			project_id = harness_platform_delegatetoken.test.project_id
		}
	`, name, accountId)
}
