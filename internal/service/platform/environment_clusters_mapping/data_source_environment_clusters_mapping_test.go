package cluster_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceCluster(t *testing.T) {

	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(6))
	name := id

	// GitOps Utilities
	agentId := os.Getenv("HARNESS_TEST_GITOPS_AGENT_ID")
	accountId := os.Getenv("HARNESS_ACCOUNT_ID")
	clusterServer := os.Getenv("HARNESS_TEST_GITOPS_CLUSTER_SERVER")
	clusterToken := os.Getenv("HARNESS_TEST_GITOPS_CLUSTER_TOKEN")
	clusterName := id
	// clusterId := id

	resourceName := "data.harness_platform_environment_clusters_mapping.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceCluster(id, name, accountId, agentId, clusterName, clusterServer, clusterToken),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "org_id", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "project_id", id),
				),
			},
		},
	})
}

func testAccDataSourceCluster(id string, name string, accoundId string, agentId string, clusterName string, clusterServer string, clusterToken string) string {
	return fmt.Sprintf(`
	resource "harness_platform_organization" "test" {
		identifier = "%[1]s"
		name = "%[2]s"
	}

	resource "harness_platform_project" "test" {
		identifier = "%[1]s"
		name = "%[2]s"
		org_id = harness_platform_organization.test.id
	}

	resource "harness_platform_environment" "test" {
		identifier = "%[1]s"
		name = "%[2]s"
		org_id = harness_platform_project.test.org_id
		project_id = harness_platform_project.test.id
		tags = ["foo:bar", "baz"]
		type = "PreProduction"
	}

	resource "harness_platform_gitops_cluster" "test" {
		identifier = "%[1]s"
		account_id = "%[3]s"
		project_id = harness_platform_project.test.id
		org_id = harness_platform_organization.test.id
		agent_id = "%[4]s"

		 request {
			upsert = true
			cluster {
				server = "%[6]s"
				name = "%[5]s"
				config {
					bearer_token = "%[7]s"
					tls_client_config {
						insecure = true
					}
					cluster_connection_type = "SERVICE_ACCOUNT"
				}
			}
		}
	}

	resource "harness_platform_environment_clusters_mapping" "test" {
		identifier = harness_platform_gitops_cluster.test.id
		env_id = harness_platform_environment.test.id
		clusters {
			identifier = harness_platform_gitops_cluster.test.id
			scope = "ACCOUNT"
		}
		org_id = harness_platform_organization.test.id
		project_id = harness_platform_project.test.id
	}

		data "harness_platform_environment_clusters_mapping" "test" {
			identifier = harness_platform_gitops_cluster.test.id
			env_id = harness_platform_environment.test.id
			org_id = harness_platform_organization.test.id
			project_id = harness_platform_project.test.id
		}
`, id, name, accoundId, agentId, clusterName, clusterServer, clusterToken)
}
