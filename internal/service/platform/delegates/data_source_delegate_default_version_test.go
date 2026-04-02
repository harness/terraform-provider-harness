package delegates_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

/*
# For dev testing only do the following
DEV_TEST_PROVIDER_BLOCK="true"

#  Create account, org, project level keys and set the following environment variables before running the tests
export HARNESS_ACCOUNT_ID='xxxxxxxxx'
export LATEST_DELEGATE_ORGANIZATION_ID='xxxxxxxxx'
export LATEST_DELEGATE_PROJECT_ID='xxxxxxxxx'
export HARNESS_PLATFORM_ACCOUNT_LEVEL_API_KEY='xxxxxxxxx'
export HARNESS_PLATFORM_ORGANIZATION_LEVEL_API_KEY='xxxxxxxxx'
export HARNESS_PLATFORM_PROJECT_LEVEL_API_KEY='xxxxxxxxx'

go test -v ./internal/service/platform/delegates/... -run TestAccDataSourceDelegateDefaultVersion
*/
func TestAccDataSourceDelegateDefaultVersion(t *testing.T) {
	t.Setenv("HARNESS_PLATFORM_API_KEY", os.Getenv("HARNESS_PLATFORM_ACCOUNT_LEVEL_API_KEY"))
	resourceName := "data.harness_platform_delegate_default_version.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceDelegateDefaultVersionImplicitAccount(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "version"),
					resource.TestCheckResourceAttrSet(resourceName, "minimal_version"),
					logAttr(t, resourceName, "version"),
					logAttr(t, resourceName, "minimal_version"),
				),
			},
		},
	})
}

func TestAccDataSourceDelegateDefaultVersionOrgLevelImplicitAccount(t *testing.T) {
	t.Setenv("HARNESS_PLATFORM_API_KEY", os.Getenv("HARNESS_PLATFORM_ORGANIZATION_LEVEL_API_KEY"))
	orgId := os.Getenv("LATEST_DELEGATE_ORGANIZATION_ID")
	resourceName := "data.harness_platform_delegate_default_version.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceDelegateDefaultVersionOrgLevelImplicitAccount(orgId),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "org_id", orgId),
					resource.TestCheckResourceAttrSet(resourceName, "version"),
					resource.TestCheckResourceAttrSet(resourceName, "minimal_version"),
					logAttr(t, resourceName, "version"),
					logAttr(t, resourceName, "minimal_version"),
				),
			},
		},
	})
}

func TestAccDataSourceDelegateDefaultVersionProjectLevel(t *testing.T) {
	t.Setenv("HARNESS_PLATFORM_API_KEY", os.Getenv("HARNESS_PLATFORM_PROJECT_LEVEL_API_KEY"))
	orgId := os.Getenv("LATEST_DELEGATE_ORGANIZATION_ID")
	projectId := os.Getenv("LATEST_DELEGATE_PROJECT_ID")
	resourceName := "data.harness_platform_delegate_default_version.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceDelegateDefaultVersionProjectLevel(orgId, projectId),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "org_id", orgId),
					resource.TestCheckResourceAttr(resourceName, "project_id", projectId),
					resource.TestCheckResourceAttrSet(resourceName, "version"),
					resource.TestCheckResourceAttrSet(resourceName, "minimal_version"),
					logAttr(t, resourceName, "version"),
					logAttr(t, resourceName, "minimal_version"),
				),
			},
		},
	})
}

func logAttr(t *testing.T, resourceName, attr string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("resource %q not found in state", resourceName)
		}
		t.Logf("%s.%s = %q", resourceName, attr, rs.Primary.Attributes[attr])
		return nil
	}
}

func devProviderBlock() string {
	if os.Getenv("DEV_TEST_PROVIDER_BLOCK") == "true" {
		return `
terraform {
  required_providers {
    harness = {
      source  = "harness/harness"
      version = "0.4000.2"
    }
  }
}
`
	}
	return ""
}

func testAccDataSourceDelegateDefaultVersionImplicitAccount() string {
	return devProviderBlock() + `data "harness_platform_delegate_default_version" "test" {}`
}

func testAccDataSourceDelegateDefaultVersionOrgLevelImplicitAccount(orgId string) string {
	return devProviderBlock() + fmt.Sprintf(`
data "harness_platform_delegate_default_version" "test" {
  org_id = "%[1]s"
}
`, orgId)
}

func testAccDataSourceDelegateDefaultVersionProjectLevel(orgId, projectId string) string {
	return devProviderBlock() + fmt.Sprintf(`
data "harness_platform_delegate_default_version" "test" {
  org_id     = "%[1]s"
  project_id = "%[2]s"
}
`, orgId, projectId)
}

