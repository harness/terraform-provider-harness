package gitops_cluster_test

import (
	"fmt"
	"testing"

	"github.com/antihax/optional"
	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccResourceGitopsCluster(t *testing.T) {
	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	agentId := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	name := id
	updatedName := fmt.Sprintf("%s_updated", name)
	resourceName := "harness_platform_gitops_cluster.test"
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccResourceGitopsClusterDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceGitopsCluster(id, name, agentId),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
				),
			},
			{
				Config: testAccResourceGitopsCluster(id, updatedName, agentId),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: acctest.ProjectResourceImportStateIdFunc(resourceName),
			},
		},
	})

}

func testAccGetCluster(resourceName string, state *terraform.State) (*nextgen.Servicev1Cluster, error) {
	r := acctest.TestAccGetResource(resourceName, state)
	c, ctx := acctest.TestAccGetPlatformClientWithContext()
	agentIdentifier := r.Primary.Attributes["agent_identifier"]
	identifier := r.Primary.Attributes["identifier"]

	resp, _, err := c.AgentClusterApi.AgentClusterServiceGet(ctx, agentIdentifier, identifier, &nextgen.AgentClusterServiceApiAgentClusterServiceGetOpts{
		OrgIdentifier:     optional.NewString(r.Primary.Attributes["org_identifier"]),
		ProjectIdentifier: optional.NewString(r.Primary.Attributes["project_identifier"]),
		QueryServer:       optional.NewString(r.Primary.Attributes["query.server"]),
		QueryName:         optional.NewString(r.Primary.Attributes["query.name"]),
	})

	if err != nil {
		return nil, err
	}

	if resp.Cluster == nil {
		return nil, nil
	}

	return &resp, nil
}

func testAccResourceGitopsClusterDestroy(resourceName string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		cluster, _ := testAccGetCluster(resourceName, state)
		if cluster != nil {
			return fmt.Errorf("Found cluster: %s", cluster.Identifier)
		}

		return nil
	}

}

func testAccResourceGitopsCluster(id string, name string, agentId string) string {
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
		
		resource "harness_platform_gitops_cluster" "test" {
			identifier = "%[1]s"
			account_identifier = "%[3]s"
			project_identifier = harness_platform_project.test.id
			org_identifier = harness_platform_project.test.org_id
			agent_identifier = "%[3]s"

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
		`, id, name, agentId)

}
