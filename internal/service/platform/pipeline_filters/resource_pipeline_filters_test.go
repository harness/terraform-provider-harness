package pipeline_filters_test

import (
	"fmt"
	"testing"

	"github.com/antihax/optional"
	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccResourcePipelineFilters(t *testing.T) {

	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	name := id
	updatedName := fmt.Sprintf("%s_updated", name)
	resourceName := "harness_platform_pipeline_filters.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccPipelineFiltersDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourcePipelineFilters(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "type", "PipelineExecution"),
					resource.TestCheckResourceAttr(resourceName, "filter_properties.0.tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "filter_properties.0.filter_type", "PipelineExecution"),
					resource.TestCheckResourceAttr(resourceName, "filter_visibility", "EveryOne"),
				),
			},
			{
				Config: testAccResourcePipelineFilters(id, updatedName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
					resource.TestCheckResourceAttr(resourceName, "type", "PipelineExecution"),
					resource.TestCheckResourceAttr(resourceName, "filter_properties.0.filter_type", "PipelineExecution"),
					resource.TestCheckResourceAttr(resourceName, "filter_properties.0.tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "filter_visibility", "EveryOne"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: acctest.ProjectFilterImportStateIdFunc(resourceName),
			},
		},
	})
}

func TestAccResourcePipelineFiltersOrgLevel(t *testing.T) {

	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	name := id
	updatedName := fmt.Sprintf("%s_updated", name)
	resourceName := "harness_platform_pipeline_filters.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccPipelineFiltersOrgLevelDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourcePipelineFiltersOrgLevel(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "type", "PipelineExecution"),
					resource.TestCheckResourceAttr(resourceName, "filter_properties.0.tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "filter_properties.0.filter_type", "PipelineExecution"),
					resource.TestCheckResourceAttr(resourceName, "filter_visibility", "OnlyCreator"),
				),
			},
			{
				Config: testAccResourcePipelineFiltersOrgLevel(id, updatedName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
					resource.TestCheckResourceAttr(resourceName, "type", "PipelineExecution"),
					resource.TestCheckResourceAttr(resourceName, "filter_properties.0.tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "filter_properties.0.filter_type", "PipelineExecution"),
					resource.TestCheckResourceAttr(resourceName, "filter_visibility", "OnlyCreator"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: acctest.OrgFilterImportStateIdFunc(resourceName),
			},
		},
	})
}

func TestAccResourcePipelineFiltersTags(t *testing.T) {

	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	name := id
	updatedName := fmt.Sprintf("%s_updated", name)
	resourceName := "harness_platform_pipeline_filters.pipelinetags"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccPipelineFiltersDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourcePipelineFiltersWithTags(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "type", "PipelineSetup"),
					resource.TestCheckResourceAttr(resourceName, "filter_properties.0.filter_type", "PipelineSetup"),
					resource.TestCheckResourceAttr(resourceName, "filter_properties.0.pipeline_tags.0.key", "tag1"),
					resource.TestCheckResourceAttr(resourceName, "filter_properties.0.pipeline_tags.0.value", "123"),
					resource.TestCheckResourceAttr(resourceName, "filter_properties.0.name", "pipeline_name"),
					resource.TestCheckResourceAttr(resourceName, "filter_properties.0.description", "pipeline_description"),
					resource.TestCheckResourceAttr(resourceName, "filter_properties.0.pipeline_identifiers.0", "id1"),
				),
			},
			{
				Config: testAccResourcePipelineFiltersWithTags(id, updatedName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
					resource.TestCheckResourceAttr(resourceName, "type", "PipelineSetup"),
					resource.TestCheckResourceAttr(resourceName, "filter_properties.0.filter_type", "PipelineSetup"),
					resource.TestCheckResourceAttr(resourceName, "filter_properties.0.pipeline_tags.0.key", "tag1"),
					resource.TestCheckResourceAttr(resourceName, "filter_properties.0.pipeline_tags.0.value", "123"),
					resource.TestCheckResourceAttr(resourceName, "filter_properties.0.name", "pipeline_name"),
					resource.TestCheckResourceAttr(resourceName, "filter_properties.0.description", "pipeline_description"),
					resource.TestCheckResourceAttr(resourceName, "filter_properties.0.pipeline_identifiers.0", "id1"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: acctest.ProjectFilterImportStateIdFunc(resourceName),
			},
		},
	})
}

func testAccGetResourcePipelineFilters(resourceName string, state *terraform.State) (*nextgen.PipelineFilter, error) {
	r := acctest.TestAccGetResource(resourceName, state)
	c, ctx := acctest.TestAccGetPlatformClientWithContext()
	id := r.Primary.ID
	type_ := r.Primary.Attributes["type"]

	resp, _, err := c.FilterApi.PipelinegetFilter(ctx, c.AccountId, id, type_, &nextgen.FilterApiPipelinegetFilterOpts{
		OrgIdentifier:     buildField(r, "org_id"),
		ProjectIdentifier: buildField(r, "project_id"),
	})

	if err != nil {
		return nil, err
	}

	if resp.Data == nil {
		return nil, nil
	}

	return resp.Data, nil
}

func testAccPipelineFiltersOrgLevelDestroy(resourceName string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		filter, _ := testAccGetResourcePipelineFilters(resourceName, state)
		if filter != nil {
			return fmt.Errorf("Found filter: %s", filter.Identifier)
		}

		return nil
	}
}

func testAccPipelineFiltersDestroy(resourceName string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		filter, _ := testAccGetResourcePipelineFilters(resourceName, state)
		if filter != nil {
			return fmt.Errorf("Found filter: %s", filter.Identifier)
		}

		return nil
	}
}

func buildField(r *terraform.ResourceState, field string) optional.String {
	if attr, ok := r.Primary.Attributes[field]; ok {
		return optional.NewString(attr)
	}
	return optional.EmptyString()
}

func testAccResourcePipelineFilters(id string, name string) string {
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
`, id, name)
}

func testAccResourcePipelineFiltersOrgLevel(id string, name string) string {
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
`, id, name)
}

func testAccResourcePipelineFiltersWithTags(id string, name string) string {
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

		resource "harness_platform_pipeline_filters" "pipelinetags" {
			identifier = "%[1]s"
			name = "%[2]s"
			org_id = harness_platform_project.test.org_id
			project_id = harness_platform_project.test.id
			type = "PipelineSetup"
			filter_properties {
				filter_type = "PipelineSetup"
				name = "pipeline_name"
				description = "pipeline_description"
				pipeline_identifiers = ["id1", "id2"]
				pipeline_tags = [
					{
						"key" = "tag1"
                		"value" = "123"
					},
					{
						"key" = "tag2"
                		"value" = "456"
					},
				]
			}
		}
`, id, name)
}
