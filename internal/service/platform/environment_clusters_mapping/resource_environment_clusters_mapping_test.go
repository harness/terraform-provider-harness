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

	name := t.Name()
	id := fmt.Sprintf("%s_%s", name, utils.RandStringBytes(5))
	// updatedName := fmt.Sprintf("%s_updated", name)

	// GitOps Utilities
	agentId := "terraformagent"
	accountId := os.Getenv("HARNESS_ACCOUNT_ID")
	clusterName := id
	// clusterId := id
	resourceName := "harness_platform_environment_clusters_mapping.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccEnvironmentClustersMappingDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceEnvironmentClustersMapping(id, name, accountId, agentId, clusterName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "org_id", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
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

func testAccResourceEnvironmentClustersMapping(id string, name string, accoundId string, agentId string, clusterName string) string {
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
`, id, name, accoundId, agentId, clusterName)
}
