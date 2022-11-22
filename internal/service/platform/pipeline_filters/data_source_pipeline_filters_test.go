package pipeline_filters_test

import (
	"fmt"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourcePipelineFilters(t *testing.T) {

	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(6))
	name := id
	resourceName := "data.harness_platform_pipeline_filters.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourcePipelineFilters(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "org_id", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "project_id", id),
					resource.TestCheckResourceAttr(resourceName, "type", "PipelineExecution"),
					resource.TestCheckResourceAttr(resourceName, "filter_properties.0.tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "filter_properties.0.filter_type", "PipelineExecution"),
					resource.TestCheckResourceAttr(resourceName, "filter_visibility", "EveryOne"),
				),
			},
		},
	})
}

func TestAccDataSourcePipelineFiltersOrgLevel(t *testing.T) {

	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(6))
	name := id
	resourceName := "data.harness_platform_pipeline_filters.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourcePipelineFiltersOrgLevel(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "org_id", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "type", "PipelineExecution"),
					resource.TestCheckResourceAttr(resourceName, "filter_properties.0.tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "filter_properties.0.filter_type", "PipelineExecution"),
					resource.TestCheckResourceAttr(resourceName, "filter_visibility", "OnlyCreator"),
				),
			},
		},
	})
}

func testAccDataSourcePipelineFilters(id string, name string) string {
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

	resource "harness_platform_pipeline_filters" "test" {
		identifier = "%[1]s"
		name = "%[2]s"
		org_id = harness_platform_project.test.org_id
		project_id = harness_platform_project.test.id
		type = "PipelineExecution"
		filter_properties {
			 tags = ["foo:bar"]
			 filter_type = "PipelineExecution"
	}
	filter_visibility = "EveryOne"
	}

	data "harness_platform_pipeline_filters" "test" {
			identifier = harness_platform_pipeline_filters.test.identifier
			org_id = harness_platform_pipeline_filters.test.org_id
			project_id = harness_platform_pipeline_filters.test.project_id
			type = harness_platform_pipeline_filters.test.type
		}
`, id, name)
}

func testAccDataSourcePipelineFiltersOrgLevel(id string, name string) string {
	return fmt.Sprintf(`
	resource "harness_platform_organization" "test" {
		identifier = "%[1]s"
		name = "%[2]s"
	}

	resource "harness_platform_pipeline_filters" "test" {
		identifier = "%[1]s"
		name = "%[2]s"
		org_id = harness_platform_organization.test.id
		type = "PipelineExecution"
		filter_properties {
			tags = ["foo:bar"]
			filter_type = "PipelineExecution"
		}
		filter_visibility = "OnlyCreator"
	}

	data "harness_platform_pipeline_filters" "test" {
			identifier = harness_platform_filters.test.identifier
			org_id = harness_platform_filters.test.org_id
			type = harness_platform_filters.test.type
		}
`, id, name)
}
