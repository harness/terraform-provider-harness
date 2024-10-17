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

func testClusterOrch(name string) string {
	return fmt.Sprintf(`
	resource "harness_cluster_orchestrator" "test" {
		name = "%s"  
		cluster_endpoint = "http://test.com" 
		k8s_connector_id = "TestDoNotDelete"                    
	}
`, name)
}
