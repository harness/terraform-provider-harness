package template_filters_test

import (
	"fmt"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceTemplateFilters(t *testing.T) {

	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(6))
	name := id
	resourceName := "data.harness_platform_template_filters.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTemplateFilters(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "org_id", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "project_id", id),
					resource.TestCheckResourceAttr(resourceName, "type", "Template"),
					resource.TestCheckResourceAttr(resourceName, "filter_properties.0.tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "filter_properties.0.filter_type", "Template"),
					resource.TestCheckResourceAttr(resourceName, "filter_visibility", "EveryOne"),
				),
			},
		},
	})
}

func TestAccDataSourceTemplateFiltersOrgLevel(t *testing.T) {

	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(6))
	name := id
	resourceName := "data.harness_platform_template_filters.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTemplateFiltersOrgLevel(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "org_id", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "type", "Template"),
					resource.TestCheckResourceAttr(resourceName, "filter_properties.0.tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "filter_properties.0.filter_type", "Template"),
					resource.TestCheckResourceAttr(resourceName, "filter_visibility", "OnlyCreator"),
				),
			},
		},
	})
}

func testAccDataSourceTemplateFilters(id string, name string) string {
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

	resource "harness_platform_template_filters" "test" {
		identifier = "%[1]s"
		name = "%[2]s"
		org_id = harness_platform_project.test.org_id
		project_id = harness_platform_project.test.id
		type = "Template"
		filter_properties {
			 tags = ["foo:bar"]
			 filter_type = "Template"
	}
	filter_visibility = "EveryOne"
	}

	data "harness_platform_template_filters" "test" {
			identifier = harness_platform_template_filters.test.identifier
			org_id = harness_platform_template_filters.test.org_id
			project_id = harness_platform_template_filters.test.project_id
			type = harness_platform_template_filters.test.type
		}
`, id, name)
}

func testAccDataSourceTemplateFiltersOrgLevel(id string, name string) string {
	return fmt.Sprintf(`
	resource "harness_platform_organization" "test" {
		identifier = "%[1]s"
		name = "%[2]s"
	}

	resource "harness_platform_template_filters" "test" {
		identifier = "%[1]s"
		name = "%[2]s"
		org_id = harness_platform_organization.test.id
		type = "Template"
		filter_properties {
			tags = ["foo:bar"]
			filter_type = "Template"
		}
		filter_visibility = "OnlyCreator"
	}

	data "harness_platform_template_filters" "test" {
			identifier = harness_platform_template_filters.test.identifier
			org_id = harness_platform_template_filters.test.org_id
			type = harness_platform_template_filters.test.type
		}
`, id, name)
}
