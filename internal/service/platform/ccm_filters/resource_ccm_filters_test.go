package ccm_filters_test

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

func TestAccResourceCCMFilters(t *testing.T) {

	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	name := id
	updatedName := fmt.Sprintf("%s_updated", name)
	resourceName := "harness_platform_ccm_filters.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccCCMFiltersDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceCCMFilters(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "type", "CCMRecommendation"),
					resource.TestCheckResourceAttr(resourceName, "filter_properties.0.tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "filter_properties.0.filter_type", "CCMRecommendation"),
					resource.TestCheckResourceAttr(resourceName, "filter_visibility", "EveryOne"),
				),
			},
			{
				Config: testAccResourceCCMFilters(id, updatedName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
					resource.TestCheckResourceAttr(resourceName, "type", "CCMRecommendation"),
					resource.TestCheckResourceAttr(resourceName, "filter_properties.0.filter_type", "CCMRecommendation"),
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

func TestAccResourceCCMFiltersOrgLevel(t *testing.T) {

	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	name := id
	updatedName := fmt.Sprintf("%s_updated", name)
	resourceName := "harness_platform_ccm_filters.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccCCMFiltersOrgLevelDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceCCMFiltersOrgLevel(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "type", "CCMRecommendation"),
					resource.TestCheckResourceAttr(resourceName, "filter_properties.0.tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "filter_properties.0.filter_type", "CCMRecommendation"),
					resource.TestCheckResourceAttr(resourceName, "filter_visibility", "OnlyCreator"),
				),
			},
			{
				Config: testAccResourceCCMFiltersOrgLevel(id, updatedName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
					resource.TestCheckResourceAttr(resourceName, "type", "CCMRecommendation"),
					resource.TestCheckResourceAttr(resourceName, "filter_properties.0.tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "filter_properties.0.filter_type", "CCMRecommendation"),
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

func testAccGetResourceCCMFilters(resourceName string, state *terraform.State) (*nextgen.Filter, error) {
	r := acctest.TestAccGetResource(resourceName, state)
	c, ctx := acctest.TestAccGetPlatformClientWithContext()
	id := r.Primary.ID
	type_ := r.Primary.Attributes["type"]

	resp, _, err := c.FilterApi.CcmgetFilter(ctx, c.AccountId, id, type_, &nextgen.FilterApiCcmgetFilterOpts{
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

func testAccCCMFiltersOrgLevelDestroy(resourceName string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		filter, _ := testAccGetResourceCCMFilters(resourceName, state)
		if filter != nil {
			return fmt.Errorf("Found filter: %s", filter.Identifier)
		}

		return nil
	}
}

func testAccCCMFiltersDestroy(resourceName string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		filter, _ := testAccGetResourceCCMFilters(resourceName, state)
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

func testAccResourceCCMFilters(id string, name string) string {
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

		resource "harness_platform_ccm_filters" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
			org_id = harness_platform_project.test.org_id
			project_id = harness_platform_project.test.id
			type = "CCMRecommendation"
			filter_properties {
				 tags = ["foo:bar"]
         filter_type = "CCMRecommendation"
    }
    filter_visibility = "EveryOne"
		}
`, id, name)
}

func testAccResourceCCMFiltersOrgLevel(id string, name string) string {
	return fmt.Sprintf(`
		resource "harness_platform_organization" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
		}

		resource "harness_platform_ccm_filters" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
			org_id = harness_platform_organization.test.id
			type = "CCMRecommendation"
			filter_properties {
				tags = ["foo:bar"]
				filter_type = "CCMRecommendation"
			}
			filter_visibility = "OnlyCreator"
		}
`, id, name)
}
