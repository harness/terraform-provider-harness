package sso_test

import (
	"fmt"
	"testing"

	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceProvider_ldap(t *testing.T) {

	var (
		resourceName = "data.harness_sso_provider.test"
	)

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceProvider("ldap-test"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", "ldap-test"),
				),
			},
		},
	})
}

func TestAccDataSourceProvider_saml(t *testing.T) {

	var (
		resourceName = "data.harness_sso_provider.test"
	)

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceProvider("saml-test"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", "saml-test"),
				),
			},
		},
	})
}

func testAccDataSourceProvider(name string) string {
	return fmt.Sprintf(`
		data "harness_sso_provider" "test" {
			name = "%[1]s"
		}
	`, name)
}
