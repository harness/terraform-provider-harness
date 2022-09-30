package gitops_cluster_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceGitopsCluster(t *testing.T) {

	id := "terraformct"
	agentId := "terraformtestagent"
	orgId := "gitopstest"
	accountId := os.Getenv("HARNESS_ACCOUNT_ID")
	projectId := "gitopsagent"
	clusterName := id

	resourceName := "data.harness_platform_gitops_cluster.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceGitopsCluster(id, accountId, projectId, orgId, agentId, clusterName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
				),
			},
		},
	})
}

func testAccDataSourceGitopsCluster(clusterId string, accoundId string, projectId string, orgId string, agentId string, clusterName string) string {
	return fmt.Sprintf(`
		data "harness_platform_gitops_cluster" "test" {
			identifier = "%[1]s"
			account_id = "%[2]s"
			project_id = "%[3]s"
			org_id = "%[4]s"
			agent_id = "%[5]s"

 			request {
				upsert = false
				id {
					type = "type123"
					value = "value123"
				}
				cluster {
					server = "server_test"
					name = "%[6]s"
					config {
						username = "test_username"
						password = "test_password"
						tls_client_config {
							insecure = true
						}
						cluster_connection_type = "USERNAME_PASSWORD"
					}
				}
			}
		}
		`, clusterId, accoundId, projectId, orgId, agentId, clusterName)

}
