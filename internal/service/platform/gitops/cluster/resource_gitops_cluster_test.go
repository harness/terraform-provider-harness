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
				ImportStateVerifyIgnore: []string{"request.0.cluster.0.info", "request.0.cluster.0.config.0.bearer_token"},
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
				ImportStateVerifyIgnore: []string{"request.0.cluster.0.info", "request.0.cluster.0.config.0.bearer_token"},
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
				ImportStateVerifyIgnore: []string{"request.0.cluster.0.info", "request.0.cluster.0.config.0.bearer_token"},
				ImportStateIdFunc:       acctest.GitopsAgentProjectLevelResourceImportStateIdFunc(resourceName),
			},
		},
	})
}

func TestAccResourceGitopsClusterIAMProject(t *testing.T) {
	// Project Level with IAM
	id := strings.ToLower(fmt.Sprintf("%s%s", t.Name(), utils.RandStringBytes(5)))
	id = strings.ReplaceAll(id, "_", "")
	accountId := os.Getenv("HARNESS_ACCOUNT_ID")
	name := id
	clusterName := id
	resourceName := "harness_platform_gitops_cluster.test"
	clusterServer := os.Getenv("HARNESS_TEST_AWS_CLUSTER_SERVER")
	roleARN := os.Getenv("HARNESS_TEST_AWS_CLUSTER_ROLE_ARN")
	awsClusterName := os.Getenv("HARNESS_TEST_AWS_CLUSTER_NAME")
	caData := os.Getenv("HARNESS_TEST_AWS_CLUSTER_CA_DATA")
	agentId := os.Getenv("HARNESS_TEST_AWS_GITOPS_AGENT")
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccResourceGitopsClusterDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceGitopsClusterProjectLevelIAM(id, accountId, name, agentId, clusterName, clusterServer, roleARN, awsClusterName, caData),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
				),
			},
			{
				Config: testAccResourceGitopsClusterProjectLevelIAM(id, accountId, name, agentId, clusterName, clusterServer, roleARN, awsClusterName, caData),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"request.0.cluster.0.info", "request.0.cluster.0.config.0.bearer_token"},
				ImportStateIdFunc:       acctest.GitopsAgentProjectLevelResourceImportStateIdFunc(resourceName),
			},
		},
	})

}

func TestAccResourceGitopsClusterExecProviderProject(t *testing.T) {
	// Project Level with Exec Provider (AWS EKS)
	id := strings.ToLower(fmt.Sprintf("%s%s", t.Name(), utils.RandStringBytes(5)))
	id = strings.ReplaceAll(id, "_", "")
	accountId := os.Getenv("HARNESS_ACCOUNT_ID")
	name := id
	clusterName := id
	resourceName := "harness_platform_gitops_cluster.test"

	// Exec Provider specific env vars
	clusterServer := os.Getenv("GITOPS_AWS_EKS_CLUSTER_URL")
	caData := os.Getenv("GITOPS_AWS_EKS_CLUSTER_CA_DATA_BASE64")
	awsAccessKeyId := os.Getenv("GITOPS_AWS_CDPLAY_ACCOUNT_ACCESS_KEY_ID")
	awsSecretAccessKey := os.Getenv("GITOPS_AWS_CDPLAY_ACCOUNT_ACCESS_KEY_SECRET")
	agentId := os.Getenv("HARNESS_TEST_GITOPS_AGENT_ID")

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccResourceGitopsClusterDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceGitopsClusterProjectLevelExecProvider(id, accountId, name, agentId, clusterName, clusterServer, caData, awsAccessKeyId, awsSecretAccessKey),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
				),
			},
			{
				Config: testAccResourceGitopsClusterProjectLevelExecProvider(id, accountId, name, agentId, clusterName, clusterServer, caData, awsAccessKeyId, awsSecretAccessKey),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"request.0.cluster.0.info", "request.0.cluster.0.config.0.exec_provider_config.0.env"},
				ImportStateIdFunc:       acctest.GitopsAgentProjectLevelResourceImportStateIdFunc(resourceName),
			},
		},
	})
}

func TestAccResourceGitopsClusterSecretExpressions(t *testing.T) {
	id := strings.ToLower(fmt.Sprintf("%s%s", t.Name(), utils.RandStringBytes(5)))
	id = strings.ReplaceAll(id, "_", "")
	accountId := os.Getenv("HARNESS_ACCOUNT_ID")
	agentId := os.Getenv("HARNESS_TEST_GITOPS_AGENT_ID")
	clusterServer := os.Getenv("HARNESS_TEST_GITOPS_CLUSTER_SERVER")
	clusterToken := os.Getenv("HARNESS_TEST_GITOPS_CLUSTER_TOKEN")
	secretIdentifier := os.Getenv("HARNESS_TEST_GITOPS_CLUSTER_SECRET_IDENTIFIER")
	clusterName := id
	resourceName := "harness_platform_gitops_cluster.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccResourceGitopsClusterDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				// Step 1: Create cluster with secret_expressions
				Config: testAccResourceGitopsClusterWithSecretExpressions(id, accountId, agentId, clusterName, clusterServer, clusterToken, secretIdentifier),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "request.0.secret_expressions.bearerToken", "account."+secretIdentifier),
				),
			},
			{
				// Step 2: Remove secret_expressions — should be cleared on the server
				Config: testAccResourceGitopsClusterAccountLevel(id, accountId, id, agentId, clusterName, clusterServer, clusterToken),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckNoResourceAttr(resourceName, "request.0.secret_expressions.bearerToken"),
				),
			},
		},
	})
}

func testAccResourceGitopsClusterWithSecretExpressions(id string, accountId string, agentId string, clusterName string, clusterServer string, clusterToken string, secretIdentifier string) string {
	return fmt.Sprintf(`
		resource "harness_platform_gitops_cluster" "test" {
			identifier = "%[1]s"
			account_id = "%[2]s"
			agent_id   = "%[3]s"

			request {
				upsert = true
				secret_expressions = {
					bearerToken = "account.%[7]s"
				}
				cluster {
					server = "%[5]s"
					name   = "%[4]s"
					config {
						bearer_token            = "%[6]s"
						cluster_connection_type = "SERVICE_ACCOUNT"
						tls_client_config {
							insecure = true
						}
					}
				}
			}
			lifecycle {
				ignore_changes = [
					request.0.upsert, request.0.cluster.0.config.0.bearer_token, request.0.cluster.0.info,
				]
			}
		}
		`, id, accountId, agentId, clusterName, clusterServer, clusterToken, secretIdentifier)
}

func testAccGetCluster(resourceName string, state *terraform.State) (*nextgen.Servicev1Cluster, error) {
	r := acctest.TestAccGetResource(resourceName, state)
	if r == nil {
		return nil, nil
	}
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

func testAccResourceGitopsClusterProjectLevelIAM(id string, accountId string, name string, agentId string, clusterName string, clusterServer string, roleARN string, awsClusterName string, caData string) string {
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
						aws_cluster_name = "%[8]s"
						tls_client_config {
							insecure = false
							ca_data = "%[9]s"
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
		`, id, accountId, name, agentId, clusterName, clusterServer, roleARN, awsClusterName, caData)
}

func testAccResourceGitopsClusterProjectLevelExecProvider(id string, accountId string, name string, agentId string, clusterName string, clusterServer string, caData string, awsAccessKeyId string, awsSecretAccessKey string) string {
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
						exec_provider_config {
							command = "argocd-k8s-auth"
							args = ["aws", "--cluster-name", "argo-management"]
							api_version = "client.authentication.k8s.io/v1beta1"
							env = {
								AWS_REGION = "ap-south-1"
								AWS_ACCESS_KEY_ID = "%[8]s"
								AWS_SECRET_ACCESS_KEY = "%[9]s"
							}
						}
						tls_client_config {
							insecure = false
							ca_data = "%[7]s"
						}
						cluster_connection_type = "EXEC_PROVIDER"
					}
				}
			}
			lifecycle {
				ignore_changes = [
					request.0.upsert, request.0.cluster.0.config.0.exec_provider_config.0.env, request.0.cluster.0.config.0.exec_provider_config.0.args, request.0.cluster.0.info,
				]
			}
		}
		`, id, accountId, name, agentId, clusterName, clusterServer, caData, awsAccessKeyId, awsSecretAccessKey)
}

func TestAccResourceGitopsClusterExecProviderProjectUpdate(t *testing.T) {
	// Project Level with Exec Provider (AWS EKS) - Update Test
	id := strings.ToLower(fmt.Sprintf("%s%s", t.Name(), utils.RandStringBytes(5)))
	id = strings.ReplaceAll(id, "_", "")
	accountId := os.Getenv("HARNESS_ACCOUNT_ID")
	name := id
	clusterName := id
	resourceName := "harness_platform_gitops_cluster.test"

	// Exec Provider specific env vars
	clusterServer := os.Getenv("GITOPS_AWS_EKS_CLUSTER_URL")
	caData := os.Getenv("GITOPS_AWS_EKS_CLUSTER_CA_DATA_BASE64")
	awsAccessKeyId := os.Getenv("GITOPS_AWS_CDPLAY_ACCOUNT_ACCESS_KEY_ID")
	awsSecretAccessKey := os.Getenv("GITOPS_AWS_CDPLAY_ACCOUNT_ACCESS_KEY_SECRET")
	agentId := os.Getenv("HARNESS_TEST_GITOPS_AGENT_ID")

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccResourceGitopsClusterDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				// Step 1: Create the cluster with basic config
				Config: testAccResourceGitopsClusterProjectLevelExecProvider(id, accountId, name, agentId, clusterName, clusterServer, caData, awsAccessKeyId, awsSecretAccessKey),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
				),
			},
			{
				// Step 2: Update the cluster with tags, labels, annotations, and different namespaces
				Config: testAccResourceGitopsClusterProjectLevelExecProviderUpdate(id, accountId, name, agentId, clusterName, clusterServer, caData, awsAccessKeyId, awsSecretAccessKey),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "request.0.tags.#", "2"),
					// Note: labels, annotations, and namespaces are accepted by the Harness update API
					// but are not returned in the GET response, so we only verify tags here.
				),
			},
			{
				// Step 3: Verify import still works after update
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"request.0.cluster.0.info", "request.0.cluster.0.config.0.exec_provider_config.0.env"},
				ImportStateIdFunc:       acctest.GitopsAgentProjectLevelResourceImportStateIdFunc(resourceName),
			},
		},
	})
}

func testAccResourceGitopsClusterProjectLevelExecProviderUpdate(id string, accountId string, name string, agentId string, clusterName string, clusterServer string, caData string, awsAccessKeyId string, awsSecretAccessKey string) string {
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
				tags = ["env:qa", "updated:true"]
				cluster {
					server = "%[6]s"
					name = "%[5]s"
					config {
						exec_provider_config {
							command = "argocd-k8s-auth"
							args = ["aws", "--cluster-name", "argo-management"]
							api_version = "client.authentication.k8s.io/v1beta1"
							env = {
								AWS_REGION = "ap-south-1"
								AWS_ACCESS_KEY_ID = "%[8]s"
								AWS_SECRET_ACCESS_KEY = "%[9]s"
							}
						}
						tls_client_config {
							insecure = false
							ca_data = "%[7]s"
						}
						cluster_connection_type = "EXEC_PROVIDER"
					}
				}
			}
			lifecycle {
				ignore_changes = [
					request.0.upsert, request.0.cluster.0.config.0.exec_provider_config.0.env, request.0.cluster.0.config.0.exec_provider_config.0.args, request.0.cluster.0.info,
				]
			}
		}
		`, id, accountId, name, agentId, clusterName, clusterServer, caData, awsAccessKeyId, awsSecretAccessKey)
}

func TestAccResourceGitopsClusterExecProviderProjectAppSync(t *testing.T) {
	// Project Level with Exec Provider (AWS EKS) - App Sync Test
	id := strings.ToLower(fmt.Sprintf("%s%s", t.Name(), utils.RandStringBytes(5)))
	id = strings.ReplaceAll(id, "_", "")
	accountId := os.Getenv("HARNESS_ACCOUNT_ID")
	name := id
	clusterName := id
	resourceName := "harness_platform_gitops_cluster.test"
	appResourceName := "harness_platform_gitops_applications.test"

	// Exec Provider specific env vars
	clusterServer := os.Getenv("GITOPS_AWS_EKS_CLUSTER_URL")
	caData := os.Getenv("GITOPS_AWS_EKS_CLUSTER_CA_DATA_BASE64")
	awsAccessKeyId := os.Getenv("GITOPS_AWS_CDPLAY_ACCOUNT_ACCESS_KEY_ID")
	awsSecretAccessKey := os.Getenv("GITOPS_AWS_CDPLAY_ACCOUNT_ACCESS_KEY_SECRET")
	agentId := os.Getenv("HARNESS_TEST_GITOPS_AGENT_ID")
	repoUrl := os.Getenv("HARNESS_TEST_GITOPS_REPO")
	if repoUrl == "" {
		repoUrl = "https://github.com/harness-apps/hosted-gitops-example-apps"
	}

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccResourceGitopsClusterDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				// Step 1: Create the cluster and the application
				Config: testAccResourceGitopsClusterProjectLevelExecProviderAppSync(id, accountId, name, agentId, clusterName, clusterServer, caData, awsAccessKeyId, awsSecretAccessKey, repoUrl),
				Check: resource.ComposeTestCheckFunc(
					// Verify Cluster
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					// Verify Application
					resource.TestCheckResourceAttr(appResourceName, "id", id),
					resource.TestCheckResourceAttr(appResourceName, "application.0.spec.0.source.0.repo_url", repoUrl),
					resource.TestCheckResourceAttr(appResourceName, "application.0.spec.0.destination.0.server", clusterServer),
				),
			},
		},
	})
}

func TestAccResourceGitopsClusterTLSProject(t *testing.T) {
	// Project Level with TLS Client Cert
	id := strings.ToLower(fmt.Sprintf("%s%s", t.Name(), utils.RandStringBytes(5)))
	id = strings.ReplaceAll(id, "_", "")
	accountId := os.Getenv("HARNESS_ACCOUNT_ID")
	name := id
	clusterName := id
	resourceName := "harness_platform_gitops_cluster.test"

	// TLS specific env vars
	clusterServer := os.Getenv("HARNESS_TEST_GITOPS_CLUSTER_SERVER")
	caData := os.Getenv("GITOPS_TLS_CLUSTER_CA_DATA_BASE64")
	certData := os.Getenv("GITOPS_TLS_CLUSTER_CERT_DATA_BASE64")
	keyData := os.Getenv("GITOPS_TLS_CLUSTER_KEY_DATA_BASE64")
	agentId := os.Getenv("HARNESS_TEST_GITOPS_AGENT_ID")

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccResourceGitopsClusterDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceGitopsClusterProjectLevelTLS(id, accountId, name, agentId, clusterName, clusterServer, caData, certData, keyData),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
				),
			},
			{
				Config: testAccResourceGitopsClusterProjectLevelTLS(id, accountId, name, agentId, clusterName, clusterServer, caData, certData, keyData),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"request.0.cluster.0.info", "request.0.cluster.0.config.0.tls_client_config.0.cert_data", "request.0.cluster.0.config.0.tls_client_config.0.key_data"},
				ImportStateIdFunc:       acctest.GitopsAgentProjectLevelResourceImportStateIdFunc(resourceName),
			},
		},
	})
}

func TestAccResourceGitopsClusterTLSProjectUpdate(t *testing.T) {
	// Project Level with TLS Client Cert - Update Test
	id := strings.ToLower(fmt.Sprintf("%s%s", t.Name(), utils.RandStringBytes(5)))
	id = strings.ReplaceAll(id, "_", "")
	accountId := os.Getenv("HARNESS_ACCOUNT_ID")
	name := id
	clusterName := id
	resourceName := "harness_platform_gitops_cluster.test"

	// TLS specific env vars
	clusterServer := os.Getenv("HARNESS_TEST_GITOPS_CLUSTER_SERVER")
	caData := os.Getenv("GITOPS_TLS_CLUSTER_CA_DATA_BASE64")
	certData := os.Getenv("GITOPS_TLS_CLUSTER_CERT_DATA_BASE64")
	keyData := os.Getenv("GITOPS_TLS_CLUSTER_KEY_DATA_BASE64")
	agentId := os.Getenv("HARNESS_TEST_GITOPS_AGENT_ID")

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccResourceGitopsClusterDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				// Step 1: Create the cluster with basic config
				Config: testAccResourceGitopsClusterProjectLevelTLS(id, accountId, name, agentId, clusterName, clusterServer, caData, certData, keyData),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
				),
			},
			{
				// Step 2: Update the cluster with tags, labels, annotations, and different namespaces
				Config: testAccResourceGitopsClusterProjectLevelTLSUpdate(id, accountId, name, agentId, clusterName, clusterServer, caData, certData, keyData),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "request.0.tags.#", "2"),
					// Note: labels, annotations, and namespaces are accepted by the Harness update API
					// but are not returned in the GET response, so we only verify tags here.
				),
			},
			{
				// Step 3: Verify import still works after update
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"request.0.cluster.0.info", "request.0.cluster.0.config.0.tls_client_config.0.cert_data", "request.0.cluster.0.config.0.tls_client_config.0.key_data"},
				ImportStateIdFunc:       acctest.GitopsAgentProjectLevelResourceImportStateIdFunc(resourceName),
			},
		},
	})
}

func TestAccResourceGitopsClusterTLSProjectAppSync(t *testing.T) {
	// Project Level with TLS Client Cert - App Sync Test
	id := strings.ToLower(fmt.Sprintf("%s%s", t.Name(), utils.RandStringBytes(5)))
	id = strings.ReplaceAll(id, "_", "")
	accountId := os.Getenv("HARNESS_ACCOUNT_ID")
	name := id
	clusterName := id
	resourceName := "harness_platform_gitops_cluster.test"
	appResourceName := "harness_platform_gitops_applications.test"

	// TLS specific env vars
	clusterServer := os.Getenv("HARNESS_TEST_GITOPS_CLUSTER_SERVER")
	caData := os.Getenv("GITOPS_TLS_CLUSTER_CA_DATA_BASE64")
	certData := os.Getenv("GITOPS_TLS_CLUSTER_CERT_DATA_BASE64")
	keyData := os.Getenv("GITOPS_TLS_CLUSTER_KEY_DATA_BASE64")
	agentId := os.Getenv("HARNESS_TEST_GITOPS_AGENT_ID")
	repoUrl := os.Getenv("HARNESS_TEST_GITOPS_REPO")
	if repoUrl == "" {
		repoUrl = "https://github.com/harness-apps/hosted-gitops-example-apps"
	}

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccResourceGitopsClusterDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				// Step 1: Create the cluster and the application
				Config: testAccResourceGitopsClusterProjectLevelTLSAppSync(id, accountId, name, agentId, clusterName, clusterServer, caData, certData, keyData, repoUrl),
				Check: resource.ComposeTestCheckFunc(
					// Verify Cluster
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					// Verify Application
					resource.TestCheckResourceAttr(appResourceName, "id", id),
					resource.TestCheckResourceAttr(appResourceName, "application.0.spec.0.source.0.repo_url", repoUrl),
					resource.TestCheckResourceAttr(appResourceName, "application.0.spec.0.destination.0.server", clusterServer),
				),
			},
		},
	})
}

func testAccResourceGitopsClusterProjectLevelTLS(id string, accountId string, name string, agentId string, clusterName string, clusterServer string, caData string, certData string, keyData string) string {
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
						tls_client_config {
							insecure = false
							ca_data = "%[7]s"
							cert_data = "%[8]s"
							key_data = "%[9]s"
						}
						cluster_connection_type = "TLS_CLIENT_CERT"
					}
				}
			}
			lifecycle {
				ignore_changes = [
					request.0.upsert, request.0.cluster.0.config.0.tls_client_config.0.cert_data, request.0.cluster.0.config.0.tls_client_config.0.key_data, request.0.cluster.0.info,
				]
			}
		}
		`, id, accountId, name, agentId, clusterName, clusterServer, caData, certData, keyData)
}

func testAccResourceGitopsClusterProjectLevelTLSUpdate(id string, accountId string, name string, agentId string, clusterName string, clusterServer string, caData string, certData string, keyData string) string {
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
				tags = ["env:qa", "updated:true"]
				cluster {
					server = "%[6]s"
					name = "%[5]s"
					config {
						tls_client_config {
							insecure = false
							ca_data = "%[7]s"
							cert_data = "%[8]s"
							key_data = "%[9]s"
						}
						cluster_connection_type = "TLS_CLIENT_CERT"
					}
					namespaces = ["argocd", "updated-namespace"]
					cluster_resources = false
					labels = {
						"purpose" = "testing"
					}
					annotations = {
						"test-status" = "updated"
					}
				}
			}
			lifecycle {
				ignore_changes = [
					request.0.upsert, request.0.cluster.0.config.0.tls_client_config.0.cert_data, request.0.cluster.0.config.0.tls_client_config.0.key_data, request.0.cluster.0.info,
				]
			}
		}
		`, id, accountId, name, agentId, clusterName, clusterServer, caData, certData, keyData)
}

func testAccResourceGitopsClusterProjectLevelTLSAppSync(id string, accountId string, name string, agentId string, clusterName string, clusterServer string, caData string, certData string, keyData string, repoUrl string) string {
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
						tls_client_config {
							insecure = false
							ca_data = "%[7]s"
							cert_data = "%[8]s"
							key_data = "%[9]s"
						}
						cluster_connection_type = "TLS_CLIENT_CERT"
					}
				}
			}
			lifecycle {
				ignore_changes = [
					request.0.upsert, request.0.cluster.0.config.0.tls_client_config.0.cert_data, request.0.cluster.0.config.0.tls_client_config.0.key_data, request.0.cluster.0.info,
				]
			}
		}
		resource "harness_platform_gitops_applications" "test" {
			depends_on = [harness_platform_gitops_cluster.test]
			identifier = "%[1]s"
			name = "%[3]s"
			account_id = "%[2]s"
			org_id = harness_platform_organization.test.id
			project_id = harness_platform_project.test.id
			agent_id = "%[4]s"
			cluster_id = harness_platform_gitops_cluster.test.id
			
			# Skip repo validation so we don't need to create a repo resource first
			skip_repo_validation = true
			
			application {
				metadata {
					annotations = {}
					labels = {}
					name = "%[3]s"
				}
				spec {
					sync_policy {
						automated {
							prune = true
							self_heal = true
							allow_empty = false
						}
						sync_options = [
							"CreateNamespace=true"
						]
					}
					source {
						target_revision = "master"
						repo_url = "%[10]s"
						path = "helm-guestbook"
					}
					destination {
						namespace = "%[1]s-ns"
						server = "%[6]s"
					}
				}
			}
		}
		`, id, accountId, name, agentId, clusterName, clusterServer, caData, certData, keyData, repoUrl)
}

func testAccResourceGitopsClusterProjectLevelExecProviderAppSync(id string, accountId string, name string, agentId string, clusterName string, clusterServer string, caData string, awsAccessKeyId string, awsSecretAccessKey string, repoUrl string) string {
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
						exec_provider_config {
							command = "argocd-k8s-auth"
							args = ["aws", "--cluster-name", "argo-management"]
							api_version = "client.authentication.k8s.io/v1beta1"
							env = {
								AWS_REGION = "ap-south-1"
								AWS_ACCESS_KEY_ID = "%[8]s"
								AWS_SECRET_ACCESS_KEY = "%[9]s"
							}
						}
						tls_client_config {
							insecure = false
							ca_data = "%[7]s"
						}
						cluster_connection_type = "EXEC_PROVIDER"
					}
				}
			}
			lifecycle {
				ignore_changes = [
					request.0.upsert, request.0.cluster.0.config.0.exec_provider_config.0.env, request.0.cluster.0.config.0.exec_provider_config.0.args, request.0.cluster.0.info,
				]
			}
		}
		resource "harness_platform_gitops_applications" "test" {
			depends_on = [harness_platform_gitops_cluster.test]
			identifier = "%[1]s"
			name = "%[3]s"
			account_id = "%[2]s"
			org_id = harness_platform_organization.test.id
			project_id = harness_platform_project.test.id
			agent_id = "%[4]s"
			cluster_id = harness_platform_gitops_cluster.test.id
			
			# Skip repo validation so we don't need to create a repo resource first
			skip_repo_validation = true
			
			application {
				metadata {
					annotations = {}
					labels = {}
					name = "%[3]s"
				}
				spec {
					sync_policy {
						automated {
							prune = true
							self_heal = true
							allow_empty = false
						}
						sync_options = [
							"CreateNamespace=true"
						]
					}
					source {
						target_revision = "master"
						repo_url = "%[10]s"
						path = "helm-guestbook"
					}
					destination {
						namespace = "%[1]s-ns"
						server = "%[6]s"
					}
				}
			}
		}
		`, id, accountId, name, agentId, clusterName, clusterServer, caData, awsAccessKeyId, awsSecretAccessKey, repoUrl)
}

func TestAccResourceGitopsClusterExecProviderProjectInsecure(t *testing.T) {
	// Project Level with Exec Provider (AWS EKS) - Insecure
	id := strings.ToLower(fmt.Sprintf("%s%s", t.Name(), utils.RandStringBytes(5)))
	id = strings.ReplaceAll(id, "_", "")
	accountId := os.Getenv("HARNESS_ACCOUNT_ID")
	name := id
	clusterName := id
	resourceName := "harness_platform_gitops_cluster.test"

	clusterServer := os.Getenv("GITOPS_AWS_EKS_CLUSTER_URL")
	awsAccessKeyId := os.Getenv("GITOPS_AWS_CDPLAY_ACCOUNT_ACCESS_KEY_ID")
	awsSecretAccessKey := os.Getenv("GITOPS_AWS_CDPLAY_ACCOUNT_ACCESS_KEY_SECRET")
	agentId := os.Getenv("HARNESS_TEST_GITOPS_AGENT_ID")

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccResourceGitopsClusterDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceGitopsClusterProjectLevelExecProviderInsecure(id, accountId, name, agentId, clusterName, clusterServer, awsAccessKeyId, awsSecretAccessKey),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
				),
			},
			{
				Config: testAccResourceGitopsClusterProjectLevelExecProviderInsecure(id, accountId, name, agentId, clusterName, clusterServer, awsAccessKeyId, awsSecretAccessKey),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"request.0.cluster.0.info", "request.0.cluster.0.config.0.exec_provider_config.0.env"},
				ImportStateIdFunc:       acctest.GitopsAgentProjectLevelResourceImportStateIdFunc(resourceName),
			},
		},
	})
}

func TestAccResourceGitopsClusterTLSProjectInsecure(t *testing.T) {
	// Project Level with TLS Client Cert - Insecure
	id := strings.ToLower(fmt.Sprintf("%s%s", t.Name(), utils.RandStringBytes(5)))
	id = strings.ReplaceAll(id, "_", "")
	accountId := os.Getenv("HARNESS_ACCOUNT_ID")
	name := id
	clusterName := id
	resourceName := "harness_platform_gitops_cluster.test"

	clusterServer := os.Getenv("HARNESS_TEST_GITOPS_CLUSTER_SERVER")
	certData := os.Getenv("GITOPS_TLS_CLUSTER_CERT_DATA_BASE64")
	keyData := os.Getenv("GITOPS_TLS_CLUSTER_KEY_DATA_BASE64")
	agentId := os.Getenv("HARNESS_TEST_GITOPS_AGENT_ID")

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccResourceGitopsClusterDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceGitopsClusterProjectLevelTLSInsecure(id, accountId, name, agentId, clusterName, clusterServer, certData, keyData),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
				),
			},
			{
				Config: testAccResourceGitopsClusterProjectLevelTLSInsecure(id, accountId, name, agentId, clusterName, clusterServer, certData, keyData),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"request.0.cluster.0.info", "request.0.cluster.0.config.0.tls_client_config.0.cert_data", "request.0.cluster.0.config.0.tls_client_config.0.key_data"},
				ImportStateIdFunc:       acctest.GitopsAgentProjectLevelResourceImportStateIdFunc(resourceName),
			},
		},
	})
}

func testAccResourceGitopsClusterProjectLevelExecProviderInsecure(id string, accountId string, name string, agentId string, clusterName string, clusterServer string, awsAccessKeyId string, awsSecretAccessKey string) string {
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
						exec_provider_config {
							command = "argocd-k8s-auth"
							args = ["aws", "--cluster-name", "argo-management"]
							api_version = "client.authentication.k8s.io/v1beta1"
							env = {
								AWS_REGION = "ap-south-1"
								AWS_ACCESS_KEY_ID = "%[7]s"
								AWS_SECRET_ACCESS_KEY = "%[8]s"
							}
						}
						tls_client_config {
							insecure = true
						}
						cluster_connection_type = "EXEC_PROVIDER"
					}
				}
			}
			lifecycle {
				ignore_changes = [
					request.0.upsert, request.0.cluster.0.config.0.exec_provider_config.0.env, request.0.cluster.0.config.0.exec_provider_config.0.args, request.0.cluster.0.info,
				]
			}
		}
		`, id, accountId, name, agentId, clusterName, clusterServer, awsAccessKeyId, awsSecretAccessKey)
}

func testAccResourceGitopsClusterProjectLevelTLSInsecure(id string, accountId string, name string, agentId string, clusterName string, clusterServer string, certData string, keyData string) string {
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
						tls_client_config {
							insecure = true
							cert_data = "%[7]s"
							key_data = "%[8]s"
						}
						cluster_connection_type = "TLS_CLIENT_CERT"
					}
				}
			}
			lifecycle {
				ignore_changes = [
					request.0.upsert, request.0.cluster.0.config.0.tls_client_config.0.cert_data, request.0.cluster.0.config.0.tls_client_config.0.key_data, request.0.cluster.0.info,
				]
			}
		}
		`, id, accountId, name, agentId, clusterName, clusterServer, certData, keyData)
}

/*
// TODO: Enable this test once a proxy setup stage is implemented.
// Currently, Harness API validates the proxy URL immediately and fails if it's unreachable.
func TestAccResourceGitopsClusterExecProviderProjectProxy(t *testing.T) {
	// Project Level with Exec Provider (AWS EKS) - Proxy Test
	id := strings.ToLower(fmt.Sprintf("%s%s", t.Name(), utils.RandStringBytes(5)))
	id = strings.ReplaceAll(id, "_", "")
	accountId := os.Getenv("HARNESS_ACCOUNT_ID")
	name := id
	clusterName := id
	resourceName := "harness_platform_gitops_cluster.test"

	// Exec Provider specific env vars
	clusterServer := os.Getenv("GITOPS_AWS_EKS_CLUSTER_URL")
	caData := os.Getenv("GITOPS_AWS_EKS_CLUSTER_CA_DATA_BASE64")
	awsAccessKeyId := os.Getenv("GITOPS_AWS_CDPLAY_ACCOUNT_ACCESS_KEY_ID")
	awsSecretAccessKey := os.Getenv("GITOPS_AWS_CDPLAY_ACCOUNT_ACCESS_KEY_SECRET")
	agentId := os.Getenv("HARNESS_TEST_GITOPS_AGENT_ID")
	proxyUrl := "http://dummy-proxy.example.com:8080"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccResourceGitopsClusterDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceGitopsClusterProjectLevelExecProviderProxy(id, accountId, name, agentId, clusterName, clusterServer, caData, awsAccessKeyId, awsSecretAccessKey, proxyUrl),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "request.0.cluster.0.config.0.proxy_url", proxyUrl),
				),
			},
		},
	})
}

// TODO: Enable this test once a proxy setup stage is implemented.
// Currently, Harness API validates the proxy URL immediately and fails if it's unreachable.
func TestAccResourceGitopsClusterTLSProjectProxy(t *testing.T) {
	// Project Level with TLS Client Cert - Proxy Test
	id := strings.ToLower(fmt.Sprintf("%s%s", t.Name(), utils.RandStringBytes(5)))
	id = strings.ReplaceAll(id, "_", "")
	accountId := os.Getenv("HARNESS_ACCOUNT_ID")
	name := id
	clusterName := id
	resourceName := "harness_platform_gitops_cluster.test"

	// TLS specific env vars
	clusterServer := os.Getenv("HARNESS_TEST_GITOPS_CLUSTER_SERVER")
	caData := os.Getenv("GITOPS_TLS_CLUSTER_CA_DATA_BASE64")
	certData := os.Getenv("GITOPS_TLS_CLUSTER_CERT_DATA_BASE64")
	keyData := os.Getenv("GITOPS_TLS_CLUSTER_KEY_DATA_BASE64")
	agentId := os.Getenv("HARNESS_TEST_GITOPS_AGENT_ID")
	proxyUrl := "http://dummy-proxy.example.com:8080"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccResourceGitopsClusterDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceGitopsClusterProjectLevelTLSProxy(id, accountId, name, agentId, clusterName, clusterServer, caData, certData, keyData, proxyUrl),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "request.0.cluster.0.config.0.proxy_url", proxyUrl),
				),
			},
		},
	})
}
*/

// HCL Generator for Exec Provider with Proxy
func testAccResourceGitopsClusterProjectLevelExecProviderProxy(id string, accountId string, name string, agentId string, clusterName string, clusterServer string, caData string, awsAccessKeyId string, awsSecretAccessKey string, proxyUrl string) string {
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
						exec_provider_config {
							command = "argocd-k8s-auth"
							args = ["aws", "--cluster-name", "argo-management"]
							api_version = "client.authentication.k8s.io/v1beta1"
							env = {
								AWS_REGION = "ap-south-1"
								AWS_ACCESS_KEY_ID = "%[8]s"
								AWS_SECRET_ACCESS_KEY = "%[9]s"
							}
						}
						tls_client_config {
							insecure = false
							ca_data = "%[7]s"
						}
						cluster_connection_type = "EXEC_PROVIDER"
						proxy_url = "%[10]s"
					}
				}
			}
			lifecycle {
				ignore_changes = [
					request.0.upsert, request.0.cluster.0.config.0.exec_provider_config.0.env, request.0.cluster.0.config.0.exec_provider_config.0.args, request.0.cluster.0.info,
				]
			}
		}
		`, id, accountId, name, agentId, clusterName, clusterServer, caData, awsAccessKeyId, awsSecretAccessKey, proxyUrl)
}

// HCL Generator for TLS with Proxy
func testAccResourceGitopsClusterProjectLevelTLSProxy(id string, accountId string, name string, agentId string, clusterName string, clusterServer string, caData string, certData string, keyData string, proxyUrl string) string {
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
						tls_client_config {
							insecure = false
							ca_data = "%[7]s"
							cert_data = "%[8]s"
							key_data = "%[9]s"
						}
						cluster_connection_type = "TLS_CLIENT_CERT"
						proxy_url = "%[10]s"
					}
				}
			}
			lifecycle {
				ignore_changes = [
					request.0.upsert, request.0.cluster.0.config.0.tls_client_config.0.cert_data, request.0.cluster.0.config.0.tls_client_config.0.key_data, request.0.cluster.0.info,
				]
			}
		}
		`, id, accountId, name, agentId, clusterName, clusterServer, caData, certData, keyData, proxyUrl)
}
