package experiment_test

import (
	"fmt"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceChaosExperiment_byIdentity(t *testing.T) {
	// Use shorter name to stay under 47 char identity limit
	name := fmt.Sprintf("ds_id_%s", utils.RandStringBytes(5))
	resourceName := "data.harness_chaos_experiment.test"

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceChaosExperiment_byIdentity(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "org_id"),
					resource.TestCheckResourceAttrSet(resourceName, "project_id"),
					resource.TestCheckResourceAttrSet(resourceName, "experiment_id"),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttrSet(resourceName, "identity"), // API normalizes identity

					// Chaos-specific fields
					resource.TestCheckResourceAttr(resourceName, "import_type", "REFERENCE"),
					resource.TestCheckResourceAttr(resourceName, "template_details.#", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "template_details.0.identity"),
					resource.TestCheckResourceAttrSet(resourceName, "infra_id"),
					resource.TestCheckResourceAttr(resourceName, "infra_type", "KubernetesV2"),
				),
			},
		},
	})
}

func testAccDataSourceChaosExperiment_byIdentity(name string) string {
	return fmt.Sprintf(`
%s

data "harness_chaos_experiment" "test" {
  org_id     = harness_platform_organization.test.id
  project_id = harness_platform_project.test.id
  # Use experiment_id which contains the normalized identity
  identity   = harness_chaos_experiment.test.experiment_id
}
`, testAccResourceChaosExperiment_basic(name))
}

func TestAccDataSourceChaosExperiment_byName(t *testing.T) {
	// Use shorter name to stay under 47 char identity limit
	name := fmt.Sprintf("ds_nm_%s", utils.RandStringBytes(5))
	resourceName := "data.harness_chaos_experiment.test"

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceChaosExperiment_byName(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "org_id"),
					resource.TestCheckResourceAttrSet(resourceName, "project_id"),
					resource.TestCheckResourceAttrSet(resourceName, "experiment_id"),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttrSet(resourceName, "identity"), // API normalizes identity

					// Chaos-specific fields
					resource.TestCheckResourceAttr(resourceName, "import_type", "REFERENCE"),
					resource.TestCheckResourceAttr(resourceName, "template_details.#", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "template_details.0.identity"),
					resource.TestCheckResourceAttrSet(resourceName, "infra_id"),
					resource.TestCheckResourceAttr(resourceName, "infra_type", "KubernetesV2"),
				),
			},
		},
	})
}

func testAccDataSourceChaosExperiment_byName(name string) string {
	return fmt.Sprintf(`
%s

data "harness_chaos_experiment" "test" {
  org_id     = harness_platform_organization.test.id
  project_id = harness_platform_project.test.id
  name       = harness_chaos_experiment.test.name
}
`, testAccResourceChaosExperiment_basic(name))
}
