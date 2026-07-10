package cluster_orchestrator_test

import (
	"fmt"
	"testing"

	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestResourceClusterOrchestrator(t *testing.T) {
	name := "terraform-clusterorch-test"
	resourceName := "harness_cluster_orchestrator.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		// CheckDestroy:      testRuleDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testClusterOrch(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
				),
			},
		},
	})
}

// TestResourceClusterOrchestrator_CCM32336_OutOfBandDeleteRecreates verifies that
// when a Cluster Orchestrator is deleted out-of-band (UI / direct API), the next
// terraform refresh does not error and re-plans a create.
//
// Regression test for CCM-32336 (lwd GET returning HTTP 500 for a deleted entity
// causes terraform plan to fail with "giving up after 11 attempt(s)").
func TestResourceClusterOrchestrator_CCM32336_OutOfBandDeleteRecreates(t *testing.T) {
	name := "terraform-co-ccm32336-test"
	resourceName := "harness_cluster_orchestrator.test"

	var orchIDBefore string

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testClusterOrch(name),
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
				Config: testClusterOrch(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
				),
			},
		},
	})
}

func testClusterOrch(name string) string {
	return fmt.Sprintf(`
	resource "harness_cluster_orchestrator" "test" {
		name = "%s"  
		cluster_endpoint = "http://test.com" 
		k8s_connector_id = "TestDoNotDelete"                    
	}
`, name)
}
