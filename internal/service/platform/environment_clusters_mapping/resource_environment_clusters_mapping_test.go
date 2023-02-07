package cluster_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/antihax/optional"
	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccResourceEnvironmentClustersMapping(t *testing.T) {
	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	name := id
	agentId := os.Getenv("HARNESS_TEST_GITOPS_AGENT_ID")
	accountId := os.Getenv("HARNESS_ACCOUNT_ID")
	clusterServer := os.Getenv("HARNESS_TEST_GITOPS_CLUSTER_SERVER")
	clusterToken := os.Getenv("HARNESS_TEST_GITOPS_CLUSTER_TOKEN")
	clusterName := id
	resourceName := "harness_platform_environment_clusters_mapping.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccEnvironmentClustersMappingDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceEnvironmentClustersMapping(id, name, accountId, agentId, clusterName, clusterServer, clusterToken),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "org_id", id),
					resource.TestCheckResourceAttr(resourceName, "project_id", id),
				),
			},
			// {
			// 	Config: testAccResourceCluster(id, updatedName),
			// 	Check: resource.ComposeTestCheckFunc(
			// 		resource.TestCheckResourceAttr(resourceName, "id", id),
			// 		resource.TestCheckResourceAttr(resourceName, "name", updatedName),
			// 	),
			// },
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: acctest.ProjectResourceImportStateIdFunc(resourceName),
			},
		},
	})
}

func testAccGetPlatformEnvClustersMapping(resourceName string, state *terraform.State) (*nextgen.ClusterResponse, error) {
	r := acctest.TestAccGetResource(resourceName, state)
	c, ctx := acctest.TestAccGetPlatformClientWithContext()
	id := r.Primary.ID
	orgId := r.Primary.Attributes["org_id"]
	projId := r.Primary.Attributes["project_id"]
	envId := r.Primary.Attributes["env_id"]

	resp, _, err := c.ClustersApi.GetCluster((ctx), id, c.AccountId, envId, &nextgen.ClustersApiGetClusterOpts{
		OrgIdentifier:     optional.NewString(orgId),
		ProjectIdentifier: optional.NewString(projId),
	})

	if err != nil {
		return nil, err
	}

	if resp.Data == nil {
		return nil, nil
	}

	return resp.Data, nil
}

func testAccEnvironmentClustersMappingDestroy(resourceName string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		cl, _ := testAccGetPlatformEnvClustersMapping(resourceName, state)
		if cl != nil {
			return fmt.Errorf("Found Cluster: %s", cl.ClusterRef)
		}

		return nil
	}
}

func testAccResourceEnvironmentClustersMapping(id string, name string, accoundId string, agentId string, clusterName string, clusterServer string, clusterToken string) string {
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
			lifecycle {
				ignore_changes = [
					request.0.upsert, request.0.cluster.0.config.0.bearer_token,
				]
			}
		}

		resource "harness_platform_environment_clusters_mapping" "test" {
			env_id = harness_platform_environment.test.id
			clusters {
				identifier = harness_platform_gitops_cluster.test.id
				scope = "PROJECT"
			}
			org_id = harness_platform_organization.test.id
			project_id = harness_platform_project.test.id
		}
`, id, name, accoundId, agentId, clusterName, clusterServer, clusterToken)
}
