package gitops_agent_test

import (
	"fmt"
	"testing"

	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceGitopsAgent(t *testing.T) {
	// id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(6))
	orgId := "default"
	projectId := "gitops2"
	agentId := "manavtest111"
	accountId := "px7xd_BFRCi-pfWPYXVjvw"
	resourceName := "data.harness_platform_gitops_agent.test"
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceGitopsAgent(agentId, accountId, projectId, orgId, agentId),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", agentId),
				// resource.TestCheckResourceAttr(resourceName, "orgId", orgId),
				),
			},
		},
	})

}

func testAccDataSourceGitopsAgent(agentId string, accountId string, projectId string, orgId string, agentName string) string {
	return fmt.Sprintf(`
		resource "harness_platform_gitops_agent" "test" {
			identifier = "%[1]s"
			account_identifier = "%[2]s"
			project_identifier = "%[3]s"
			org_identifier = "%[4]s"
			name = "%[5]s"
			type = "CONNECTED_ARGO_PROVIDER"
			metadata {
        namespace = "tf-test"
        high_availability = true
    	}
		}
		`, agentId, accountId, projectId, orgId, agentName)
}
