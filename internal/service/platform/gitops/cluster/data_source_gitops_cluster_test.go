package cluster_test

import (
	"fmt"
	"github.com/harness/harness-go-sdk/harness/utils"
	"os"
	"testing"

	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceGitopsCluster(t *testing.T) {

	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	name := id
	agentId := "account.terraformagent1"
	accountId := os.Getenv("HARNESS_ACCOUNT_ID")
	clusterName := id
	resourceName := "data.harness_platform_gitops_cluster.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceGitopsCluster(id, accountId, name, agentId, clusterName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "org_id", id),
					resource.TestCheckResourceAttr(resourceName, "project_id", id),
				),
			},
		},
	})
}

func testAccDataSourceGitopsCluster(id string, accountId string, name string, agentId string, clusterName string) string {
	return fmt.Sprintf(`
		resource "harness_platform_organization" "test" {
			identifier = "%[1]s"
			name = "%[3]s"
		}

		resource "harness_platform_project" "test" {
			identifier = "%[1]s"
			name = "%[3]s"
			org_id = harness_platform_organization.test.id
		}
		resource "harness_platform_gitops_cluster" "test" {
			identifier = "%[1]s"
			account_id = "%[2]s"
			project_id = harness_platform_project.test.id
			org_id = harness_platform_organization.test.id
			agent_id = "%[4]s"

 			request {
				upsert = false
				cluster {
					server = "https://kubernetes.default.svc"
					name = "%[5]s"
					config {
						tls_client_config {
							insecure = true
						}
						cluster_connection_type = "IN_CLUSTER"
					}

				}
			}
			lifecycle {
				ignore_changes = [
					request.0.upsert,
				]
			}
		}
		data "harness_platform_gitops_cluster" "test" {
			identifier = harness_platform_gitops_cluster.test.id
			account_id = "%[2]s"
			project_id = harness_platform_project.test.id
			org_id = harness_platform_organization.test.id
			agent_id = "%[4]s"

		}
		`, id, accountId, name, agentId, clusterName)

}
