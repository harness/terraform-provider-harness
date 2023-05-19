package cluster_test

import (
	"context"
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/antihax/optional"
	hh "github.com/harness/harness-go-sdk/harness/helpers"
	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccResourceGitopsCluster(t *testing.T) {

	// Account Level
	id := strings.ToLower(fmt.Sprintf("%s%s", t.Name(), utils.RandStringBytes(5)))
	id = strings.ReplaceAll(id, "_", "")
	name := id
	agentId := os.Getenv("HARNESS_TEST_GITOPS_AGENT_ID")
	accountId := os.Getenv("HARNESS_ACCOUNT_ID")
	clusterServer := os.Getenv("HARNESS_TEST_GITOPS_CLUSTER_SERVER")
	clusterToken := os.Getenv("HARNESS_TEST_GITOPS_CLUSTER_TOKEN")
	clusterName := id
	resourceName := "harness_platform_gitops_cluster.test"
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccResourceGitopsClusterDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceGitopsClusterAccountLevel(id, accountId, name, agentId, clusterName, clusterServer, clusterToken),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
				),
			},
			{
				Config: testAccResourceGitopsClusterAccountLevel(id, accountId, name, agentId, clusterName, clusterServer, clusterToken),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"request.0.cluster.0.info"},
				ImportStateIdFunc:       acctest.GitopsAgentAccountLevelResourceImportStateIdFunc(resourceName),
			},
		},
	})

	// Account level with Optional Tags
	id = strings.ToLower(fmt.Sprintf("%s%s", t.Name(), utils.RandStringBytes(5)))
	id = strings.ReplaceAll(id, "_", "")
	name = id
	clusterName = id
	resourceName = "harness_platform_gitops_cluster.test2"
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccResourceGitopsClusterDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceGitopsClusterAccountLevelTags(id, accountId, name, agentId, clusterName, clusterServer, clusterToken),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
				),
			},
			{
				Config: testAccResourceGitopsClusterAccountLevelTags(id, accountId, name, agentId, clusterName, clusterServer, clusterToken),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"request.0.cluster.0.info"},
				ImportStateIdFunc:       acctest.GitopsAgentAccountLevelResourceImportStateIdFunc(resourceName),
			},
		},
	})

	// Project Level
	id = strings.ToLower(fmt.Sprintf("%s%s", t.Name(), utils.RandStringBytes(5)))
	id = strings.ReplaceAll(id, "_", "")
	name = id
	clusterName = id
	resourceName = "harness_platform_gitops_cluster.test"
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccResourceGitopsClusterDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceGitopsClusterProjectLevel(id, accountId, name, agentId, clusterName, clusterServer, clusterToken),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
				),
			},
			{
				Config: testAccResourceGitopsClusterProjectLevel(id, accountId, name, agentId, clusterName, clusterServer, clusterToken),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"request.0.cluster.0.info"},
				ImportStateIdFunc:       acctest.GitopsAgentProjectLevelResourceImportStateIdFunc(resourceName),
			},
		},
	})

	// Project Level with IAM
	id = strings.ToLower(fmt.Sprintf("%s%s", t.Name(), utils.RandStringBytes(5)))
	id = strings.ReplaceAll(id, "_", "")
	name = id
	clusterName = id
	resourceName = "harness_platform_gitops_cluster.test"
	clusterServer = os.Getenv("HARNESS_TEST_GITOPS_CLUSTER_SERVER_IAM")
	roleARN := os.Getenv("HARNESS_TEST_GITOPS_CLUSTER_ROLE_ARN")
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccResourceGitopsClusterDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceGitopsClusterProjectLevelIAM(id, accountId, name, agentId, clusterName, clusterServer, roleARN),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
				),
			},
			{
				Config: testAccResourceGitopsClusterProjectLevelIAM(id, accountId, name, agentId, clusterName, clusterServer, roleARN),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"request.0.cluster.0.info"},
				ImportStateIdFunc:       acctest.GitopsAgentProjectLevelResourceImportStateIdFunc(resourceName),
			},
		},
	})
}

func testAccGetCluster(resourceName string, state *terraform.State) (*nextgen.Servicev1Cluster, error) {
	r := acctest.TestAccGetResource(resourceName, state)
	c, ctx := acctest.TestAccGetPlatformClientWithContext()
	ctx = context.WithValue(ctx, nextgen.ContextAccessToken, hh.EnvVars.BearerToken.Get())
	agentIdentifier := r.Primary.Attributes["agent_id"]
	identifier := r.Primary.Attributes["identifier"]

	resp, _, err := c.ClustersApi.AgentClusterServiceGet(ctx, agentIdentifier, identifier, c.AccountId, &nextgen.ClustersApiAgentClusterServiceGetOpts{
		OrgIdentifier:     optional.NewString(r.Primary.Attributes["org_id"]),
		ProjectIdentifier: optional.NewString(r.Primary.Attributes["project_id"]),
	})

	if err != nil {
		return nil, err
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

func testAccResourceGitopsClusterAccountLevel(id string, accountId string, name string, agentId string, clusterName string, clusterServer string, clusterToken string) string {
	return fmt.Sprintf(`
		resource "harness_platform_gitops_cluster" "test" {
			identifier = "%[1]s"
			account_id = "%[2]s"
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
					request.0.upsert, request.0.cluster.0.config.0.bearer_token, request.0.cluster.0.info,
				]
			}
		}
		`, id, accountId, name, agentId, clusterName, clusterServer, clusterToken)
}

func testAccResourceGitopsClusterAccountLevelTags(id string, accountId string, name string, agentId string, clusterName string, clusterServer string, clusterToken string) string {
	return fmt.Sprintf(`
		resource "harness_platform_gitops_cluster" "test2" {
			identifier = "%[1]s"
			account_id = "%[2]s"
			agent_id = "%[4]s"

 			request {
				upsert = true
				tags = [
        	"foo:bar",
    		]
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
					request.0.upsert, request.0.cluster.0.config.0.bearer_token, request.0.cluster.0.info,
				]
			}
		}
		`, id, accountId, name, agentId, clusterName, clusterServer, clusterToken)
}

func testAccResourceGitopsClusterProjectLevel(id string, accountId string, name string, agentId string, clusterName string, clusterServer string, clusterToken string) string {
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
			agent_id = "%[4]s"
			project_id = harness_platform_project.test.id
			org_id = harness_platform_organization.test.id
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
					request.0.upsert, request.0.cluster.0.config.0.bearer_token, request.0.cluster.0.info,
				]
			}
		}
		`, id, accountId, name, agentId, clusterName, clusterServer, clusterToken)
}

func testAccResourceGitopsClusterProjectLevelIAM(id string, accountId string, name string, agentId string, clusterName string, clusterServer string, roleARN string) string {
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
			agent_id = "%[4]s"
			project_id = harness_platform_project.test.id
			org_id = harness_platform_organization.test.id
 			request {
				upsert = true
				cluster {
					server = "%[6]s"
					name = "%[5]s"
					config {
						role_a_r_n = "%[7]s"
						tls_client_config {
							insecure = true
						}
						cluster_connection_type = "IRSA"
					}
				}
			}
			lifecycle {
				ignore_changes = [
					request.0.upsert, request.0.cluster.0.config.0.bearer_token, request.0.cluster.0.info,
				]
			}
		}
		`, id, accountId, name, agentId, clusterName, clusterServer, roleARN)
}
