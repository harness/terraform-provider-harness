package experiment_test

import (
	"testing"

	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceChaosExperiment_basic(t *testing.T) {
	// Skip test - requires full infrastructure setup (hub, template, k8s infrastructure)
	// This test should be run manually with proper test infrastructure in place
	t.Skip("Skipping - requires hub, experiment template, and kubernetes infrastructure")

	resourceName := "harness_chaos_experiment.test"

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceChaosExperiment_basic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", "test-experiment"),
					resource.TestCheckResourceAttr(resourceName, "org_id", "default"),
					resource.TestCheckResourceAttr(resourceName, "project_id", "chaos_terraform_test"),
					resource.TestCheckResourceAttrSet(resourceName, "experiment_id"),
					resource.TestCheckResourceAttrSet(resourceName, "identity"),
					resource.TestCheckResourceAttrSet(resourceName, "infra_id"),
					resource.TestCheckResourceAttr(resourceName, "template_details.#", "1"),
				),
			},
		},
	})
}

func testAccResourceChaosExperiment_basic() string {
	return `
resource "harness_chaos_experiment" "test" {
  org_id            = "default"
  project_id        = "chaos_terraform_test"
  template_identity = "test-exp-template"
  hub_identity      = "enterprise-hub"
  name              = "test-experiment"
  infra_ref         = "test-infra"
  description       = "Test experiment from template"
  
  tags = ["test", "terraform"]
}
`
}
