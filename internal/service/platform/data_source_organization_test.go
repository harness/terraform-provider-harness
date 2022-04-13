package platform_test

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceOrganization(t *testing.T) {

	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))
	resourceName := "data.harness_platform_organization.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceOrganization(id),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "2"),
				),
			},
		},
	})
}

func TestAccDataSourceOrganization_SearchTerm(t *testing.T) {

	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))
	resourceName := "data.harness_platform_organization.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceOrganization_SearchTerm(id),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "2"),
				),
			},
		},
	})
}

func TestAccDataSourceOrganization_MultipleResults(t *testing.T) {

	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))
	resourceName := "data.harness_platform_organization.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				ExpectError: regexp.MustCompile(`more than one organization was found that matches the search criteria`),
				Config:      testAccDataSourceOrganization_MultipleResults(id, false),
			},
			{
				Config: testAccDataSourceOrganization_MultipleResults(id, true),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "2"),
				),
			},
		},
	})
}

func testAccDataSourceOrganization(id string) string {
	return fmt.Sprintf(`
		resource "harness_platform_organization" "test" {
			identifier = "%[1]s"
			name = "%[1]s"
			description = "test"
			tags = ["foo:bar", "baz:qux"]
		}

		data "harness_platform_organization" "test" {
			identifier = harness_platform_organization.test.identifier
		}
	`, id)
}

func testAccDataSourceOrganization_SearchTerm(id string) string {
	return fmt.Sprintf(`
		resource "harness_platform_organization" "test" {
			identifier = "%[1]s"
			name = "%[1]s"
			description = "test"
			tags = ["foo:bar", "baz:qux"]
		}

		data "harness_platform_organization" "test" {
			search_term = harness_platform_organization.test.name
		}
	`, id)
}

func testAccDataSourceOrganization_MultipleResults(id string, selectFirst bool) string {
	return fmt.Sprintf(`
		resource "harness_platform_organization" "test_1" {
			identifier = "%[1]s_1"
			name = "%[1]s_1 testing"
			description = "test"
			tags = ["foo:bar", "baz:qux"]
		}

		resource "harness_platform_organization" "test_2" {
			identifier = "%[1]s_2"
			name = "%[1]s_2 testing"
			description = "test"
			tags = ["foo:bar", "baz:qux"]
		}

		data "harness_platform_organization" "test" {
			search_term = "testing"
			first_result = %[2]t
		}
	`, id, selectFirst)
}
