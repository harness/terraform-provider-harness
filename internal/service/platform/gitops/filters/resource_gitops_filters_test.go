package filters_test

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/antihax/optional"
	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccResourceGitOpsFilters(t *testing.T) {
	id := strings.ToLower(fmt.Sprintf("%s%s", t.Name(), utils.RandStringBytes(5)))
	id = strings.ReplaceAll(id, "_", "")
	name := id
	resourceName := "harness_platform_gitops_filters.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccResourceGitOpsFiltersDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceGitOpsFiltersProjectLevel(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "type", "APPLICATION"),
					resource.TestCheckResourceAttr(resourceName, "filter_visibility", "OnlyCreator"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: acctest.GitOpsFilterProjectLevelImportStateIdFunc(resourceName),
			},
		},
	})
}

func testAccResourceGitOpsFiltersDestroy(resourceName string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		resource, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("Resource not found: %s", resourceName)
		}

		if resource.Primary.ID == "" {
			return fmt.Errorf("Resource ID is not set")
		}

		client := acctest.TestAccProvider.Meta().(*internal.Session)
		c, ctx := client.GetPlatformClientWithContext(context.Background())

		orgId := resource.Primary.Attributes["org_id"]
		projectId := resource.Primary.Attributes["project_id"]
		filterId := resource.Primary.ID
		filterType := resource.Primary.Attributes["type"]

		_, _, err := c.GitOpsFiltersApi.FilterServiceGet(ctx, filterId, &nextgen.FiltersApiFilterServiceGetOpts{
			AccountIdentifier: optional.NewString(c.AccountId),
			OrgIdentifier:     optional.NewString(orgId),
			ProjectIdentifier: optional.NewString(projectId),
			FilterType:        optional.NewString(filterType),
		})

		if err == nil {
			return fmt.Errorf("GitOps Filter %s still exists", resource.Primary.ID)
		}

		return nil
	}
}

func testAccResourceGitOpsFiltersProjectLevel(id string, name string) string {
	return fmt.Sprintf(`
		resource "harness_platform_organization" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
		}

		resource "harness_platform_project" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
			org_id = harness_platform_organization.test.id
		}

		resource "harness_platform_gitops_filters" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
			org_id = harness_platform_organization.test.id
			project_id = harness_platform_project.test.id
			type = "APPLICATION"
			filter_properties = jsonencode({
				"healthStatus": ["Suspended"],
				"syncStatus": ["Synced"]
			})
			filter_visibility = "OnlyCreator"
		}
	`, id, name)
}
