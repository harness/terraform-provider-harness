package gitops_cluster_test

import (
	"fmt"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceGitopsCluster(t *testing.T) {

	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(6))
	agentId := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	accountId := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	name := id
	resourceName := "data.harness_platform_gitops_cluster.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceGitopsCluster(id, name, accountId, agentId),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
				),
			},
		},
	})
}

func testAccDataSourceGitopsCluster(id string, name string, accountId string, agentId string) string {
	return fmt.Sprintf(`
		resource "harness_platform_organization" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
		}

		resource "harness_platform_project" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
			org_id = harness_platform_organization.test.id
			color = "#472848"
		}
		
		data "harness_platform_gitops_cluster" "test" {
			identifier = "%[1]s"
			account_identifier = "%[3]s"
			project_identifier = harness_platform_project.test.id
			org_identifier = harness_platform_project.test.org_id
			agent_identifier = "%[4]s"

 			request {
				upsert = false
				id {
					type = "type123"
					value = "value123"
				}
				cluster {
					server = "server_test"
					name = "server_name_test"
					config {
						username = "test_username"
						password = "test_password"
						bearer_token = "bearer_token_test"
						tls_client_config {
							insecure = false
							server_name = "tsl_server_name_test"
							cert_data = "tls_cert_data_test"
						}
						aws_auth_config {
							cluster_name = "aws_cluster_name_test"
							role_a_r_n = "aws_arn_test"
						}
						exec_provider_config {
							command = "exec_command_test"
							args = ["args", "test"]
							env = {
								"abc":"def"
								"ghi":"jkl"
							}
						}
						cluster_connection_type = "TEST_CONNECTION_TYPE"
					}
					namespaces = ["NS1", "NS2", "TEST1"]
					refresh_requested_at {
						seconds = "1234567890"
						nanos = 98765
					}
					info {
						connection_state {
							status = "TEST_STATUS"
							message = "Random message TEST"
							attempted_at {
								seconds = "1234567890"
								nanos = 98765
							}
						}
						server_version = "1.1.1.2.3.4"
					}
				}
			}
		}
		`, id, name, accountId, agentId)

}
