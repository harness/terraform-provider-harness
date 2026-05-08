package cluster_orchestrator_test

import (
	"fmt"
	"testing"

	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

// Blocked on CCM-32498: harness_cluster_orchestrator create panics with
// "Invalid address to set: []string{\"identifier\"}" because setId() writes an
// attribute that the resource schema does not declare. Test cannot reach Step 1
// completion until the provider is fixed.
//
// TestResourceClusterOrchestrator_CCM32336_OutOfBandDeleteRecreates verifies that
// when a Cluster Orchestrator is deleted out-of-band (UI / direct API), the next
// terraform refresh does not error. The orchestrator resource's ReadContext
// re-creates on miss, so the contract here is that the API call path used during
// refresh tolerates a deleted entity instead of failing the terraform plan.
//
// Regression test for CCM-32336 (the bug class is: lwd GET returning HTTP 500
// for a deleted entity causes terraform plan to fail with "giving up after 11
// attempt(s)"). Tracked alongside CCM-32403 for non-rule AutoStopping entities.
func TestResourceClusterOrchestrator_CCM32336_OutOfBandDeleteRecreates(t *testing.T) {
	name := "terraform-co-ccm32336-test"
	resourceName := "harness_cluster_orchestrator.test"

	var orchIDBefore string

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testCCM32336ClusterOrch(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttrWith(resourceName, "id", func(value string) error {
						orchIDBefore = value
						return nil
					}),
				),
			},
			{
				PreConfig: func() {
					c, ctx := acctest.TestAccGetPlatformClientWithContext()
					if _, err := c.CloudCostClusterOrchestratorApi.DeleteClusterOrchestrator(
						ctx, c.AccountId, orchIDBefore,
					); err != nil {
						t.Fatalf("CCM-32336: out-of-band delete failed: %v", err)
					}
				},
				Config: testCCM32336ClusterOrch(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
				),
			},
		},
	})
}

func testCCM32336ClusterOrch(name string) string {
	return fmt.Sprintf(`
	resource "harness_cluster_orchestrator" "test" {
		name             = "%s"
		cluster_endpoint = "http://test.com"
		k8s_connector_id = "TestDoNotDelete"
	}
`, name)
}
