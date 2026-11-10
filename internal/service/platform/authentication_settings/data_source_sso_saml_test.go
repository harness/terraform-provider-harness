package authentication_settings_test

import (
	"testing"

	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceSSOSAML(t *testing.T) {
	resourceName := "data.harness_platform_sso_saml.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceSSOSAML(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "identifier"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "TF SAML DS Test"),
					resource.TestCheckResourceAttrSet(resourceName, "origin"),
				),
			},
		},
	})
}

func testAccDataSourceSSOSAML() string {
	return `
		resource "harness_platform_sso_saml" "test" {
			display_name = "TF SAML DS Test"
			metadata_xml = file("testdata/idp-metadata.xml")
		}

		data "harness_platform_sso_saml" "test" {
			identifier = harness_platform_sso_saml.test.identifier
		}
	`
}
