package cluster_test

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"

	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceGitopsCluster(t *testing.T) {

	id := strings.ToLower(fmt.Sprintf("%s%s", t.Name(), utils.RandStringBytes(5)))
	id = strings.ReplaceAll(id, "_", "")
	name := id
	agentId := os.Getenv("HARNESS_TEST_GITOPS_AGENT_ID")
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

func TestAccDataSourceGitopsClusterIAM(t *testing.T) {

	id := strings.ToLower(fmt.Sprintf("%s%s", t.Name(), utils.RandStringBytes(5)))
	id = strings.ReplaceAll(id, "_", "")
	name := id
	agentId := os.Getenv("HARNESS_TEST_GITOPS_AGENT_ID")
	accountId := os.Getenv("HARNESS_ACCOUNT_ID")
	clusterName := id
	clusterServer := os.Getenv("HARNESS_TEST_GITOPS_CLUSTER_SERVER_IAM")
	roleARN := os.Getenv("HARNESS_TEST_GITOPS_CLUSTER_ROLE_ARN")
	resourceName := "data.harness_platform_gitops_cluster.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceGitopsClusterIAM(id, accountId, name, agentId, clusterName, clusterServer, roleARN),
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
				upsert = true
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
					request.0.upsert, request.0.cluster.0.config.0.bearer_token, request.0.cluster.0.info,
				]
			}
		}
		data "harness_platform_gitops_cluster" "test" {
			depends_on = [harness_platform_gitops_cluster.test]
			identifier = harness_platform_gitops_cluster.test.id
			account_id = "%[2]s"
			project_id = harness_platform_project.test.id
			org_id = harness_platform_organization.test.id
			agent_id = "%[4]s"

		}
		`, id, accountId, name, agentId, clusterName)

}

func testAccDataSourceGitopsClusterIAM(id string, accountId string, name string, agentId string, clusterName string, clusterServer string, roleARN string) string {
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
				upsert = true
				cluster {
					name = "%[5]s"
					server = "%[6]s"
					config {
						tls_client_config {
							insecure = true
						}
						cluster_connection_type = "IRSA"
						role_a_r_n = "%[7]s"
					}

				}
			}
			lifecycle {
				ignore_changes = [
					request.0.upsert, request.0.cluster.0.config.0.bearer_token, request.0.cluster.0.info,
				]
			}
		}
		data "harness_platform_gitops_cluster" "test" {
			depends_on = [harness_platform_gitops_cluster.test]
			identifier = harness_platform_gitops_cluster.test.id
			account_id = "%[2]s"
			project_id = harness_platform_project.test.id
			org_id = harness_platform_organization.test.id
			agent_id = "%[4]s"

		}
		`, id, accountId, name, agentId, clusterName, clusterServer, roleARN)

}
