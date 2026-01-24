package experiment_test

import (
	"testing"

	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceChaosExperiment_byIdentity(t *testing.T) {
	// Skip test - requires existing experiment in the system
	t.Skip("Skipping - requires existing experiment with known identity")

	resourceName := "data.harness_chaos_experiment.test"

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceChaosExperiment_byIdentity(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "org_id", "default"),
					resource.TestCheckResourceAttr(resourceName, "project_id", "chaos_terraform_test"),
					resource.TestCheckResourceAttrSet(resourceName, "experiment_id"),
					resource.TestCheckResourceAttrSet(resourceName, "name"),
					resource.TestCheckResourceAttr(resourceName, "template_details.#", "1"),
				),
			},
		},
	})
}

func testAccDataSourceChaosExperiment_byIdentity() string {
	return `
data "harness_chaos_experiment" "test" {
  org_id     = "default"
  project_id = "chaos_terraform_test"
  identity   = "test-experiment"
}
`
}

func TestAccDataSourceChaosExperiment_byName(t *testing.T) {
	// Skip test - requires existing experiment in the system
	t.Skip("Skipping - requires existing experiment with known name")

	resourceName := "data.harness_chaos_experiment.test"

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceChaosExperiment_byName(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "org_id", "default"),
					resource.TestCheckResourceAttr(resourceName, "project_id", "chaos_terraform_test"),
					resource.TestCheckResourceAttrSet(resourceName, "experiment_id"),
					resource.TestCheckResourceAttrSet(resourceName, "identity"),
				),
			},
		},
	})
}

func testAccDataSourceChaosExperiment_byName() string {
	return `
data "harness_chaos_experiment" "test" {
  org_id     = "default"
  project_id = "chaos_terraform_test"
  name       = "test-experiment"
}
`
}
