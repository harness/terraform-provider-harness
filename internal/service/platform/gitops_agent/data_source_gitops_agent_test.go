package gitops_agent_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceGitopsAgent(t *testing.T) {
	orgId := "gitopstest"
	projectId := "gitopsagent"
	agentId := "terraformtestagent"
	accountId := os.Getenv("HARNESS_ACCOUNT_ID")
	resourceName := "data.harness_platform_gitops_agent.test"
	agentName := "testagent"
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceGitopsAgent(agentId, accountId, projectId, orgId, agentId),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", agentName),
				),
			},
		},
	})

}

func testAccDataSourceGitopsAgent(agentId string, accountId string, projectId string, orgId string, agentName string) string {
	return fmt.Sprintf(`
		data "harness_platform_gitops_agent" "test" {
			identifier = "%[1]s"
			account_id = "%[2]s"
			project_id = "%[3]s"
			org_id = "%[4]s"
			name = "%[5]s"
			type = "CONNECTED_ARGO_PROVIDER"
			metadata {
        namespace = "tf-test"
        high_availability = true
    	}
		}
		`, agentId, accountId, projectId, orgId, agentName)
}
