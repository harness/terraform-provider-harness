package project_test

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceGitopsProjectAccLevel(t *testing.T) {
	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	id = strings.ReplaceAll(id, "_", "")
	name := strings.ToLower(id)
	agentId := os.Getenv("HARNESS_TEST_GITOPS_AGENT_ID")
	agentNamespace := os.Getenv("HARNESS_TEST_GITOPS_AGENT_NAMESPACE")
	resourceName := "harness_platform_gitops_app_project.test"
	accountId := os.Getenv("HARNESS_ACCOUNT_ID")
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		//CheckDestroy:      testAccResourceGitopsRepositoryDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceGitopsProjectAccountLevel(agentId, agentNamespace, accountId, name, "*"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "agent_id", agentId),
				),
			},
			{
				Config: testAccResourceGitopsProjectAccountLevelUpdate(agentId, agentNamespace, accountId, name, "roll"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "agent_id", agentId),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"upsert", "update_mask", "repo.0.type_"},
				ImportStateIdFunc:       acctest.GitopsProjectImportStateIdFunc(resourceName),
			},
		},
	})

}

func testAccResourceGitopsProjectAccountLevel(agentId string, agentNamespace string, accountId string, name string, namespace string) string {
	return fmt.Sprintf(`
		resource "harness_platform_gitops_app_project" "test" {
			account_id = "%[1]s"
			agent_id = "%[2]s"
			upsert = true
			project {
				metadata {
					name = "%[3]s"
					namespace = "%[5]s"
					finalizers = ["resources-finalizer.argocd.argoproj.io"]
					labels = {
						v1 = "k1"
					}
					annotations = {
						v1 = "k1"
					}
				}
				spec {
					cluster_resource_whitelist {
						group = "*"
						kind = "Namespace"
					}
					destinations {
						namespace = "guestbook"
						server = "https://kubernetes.default.svc"
						name = "in-cluster"
					}
					roles {
						name = "read-only"
						description = "Read-only privileges to my-project"
						policies = ["p, proj:%[3]s:read-only, applications, get, %[3]s/*, allow"]
						groups = ["my-oidc-group"]
					}
					roles {
						name = "ci-role"
						description = "Sync privileges for guestbook-dev"
						policies = ["p, proj:%[3]s:ci-role, applications, sync, %[3]s/guestbook-dev, allow"]
						jwt_tokens{
							iat = "1535390316"
						}
					}
					sync_windows{
						kind = "allow"
						schedule = "10 1 * * *"
						duration = "1h"
						applications = ["*-prod"]
						manual_sync = "true"
 					}
					 sync_windows{
						kind = "deny"
						schedule = "0 22 * * *"
						duration = "1h"
						namespaces = ["default"]
 					}
					 sync_windows{
						kind = "allow"
						schedule = "0 23 * * *"
						duration = "1h"
						clusters = ["in-cluster", "cluster1"]
 					}
					namespace_resource_blacklist{
						group = "group"
						kind = "ResourceQuota"
					}
					namespace_resource_blacklist{
						group = "group2"
						kind = "LimitRange"
					}
					namespace_resource_blacklist{
						group = "group3"
						kind = "NetworkPolicy"
					}
					namespace_resource_whitelist{
						group = "apps"
						kind = "Deployment"
					}
					namespace_resource_whitelist{
						group = "apps"
						kind = "StatefulSet"
					}
					orphaned_resources {
						warn = "false"
					}
					source_repos = ["*"]
				}
			}
		}
	`, accountId, agentId, name, namespace, agentNamespace)
}

func testAccResourceGitopsProjectAccountLevelUpdate(agentId string, agentNamespace string, accountId string, name string, namespace string) string {
	return fmt.Sprintf(`
		resource "harness_platform_gitops_app_project" "test" {
			account_id = "%[1]s"
			agent_id = "%[2]s"
			upsert = true
			project {
				metadata {
					name = "%[3]s"
					namespace = "%[5]s"
					finalizers = ["resources-finalizer.argocd.argoproj.io"]
					labels = {
						v1 = "k2"
					}
					annotations = {
						v1 = "k1"
					}
				}
				spec {
					cluster_resource_whitelist {
						group = "*"
						kind = "Namespace"
					}
					destinations {
						namespace = "guestbook"
						server = "https://kubernetes.default.svc"
						name = "in-cluster"
					}
					roles {
						name = "read-only"
						description = "Read-only privileges to my-project"
						policies = ["p, proj:%[3]s:read-only, applications, get, %[3]s/*, allow"]
						groups = ["my-oidc-group"]
					}
					roles {
						name = "ci-role"
						description = "Sync privileges for guestbook-dev"
						policies = ["p, proj:%[3]s:ci-role, applications, sync, %[3]s/guestbook-dev, allow"]
						jwt_tokens{
							iat = "1535390316"
						}
					}
					sync_windows{
						kind = "allow"
						schedule = "10 1 * * *"
						duration = "1h"
						applications = ["*-prod"]
						manual_sync = "true"
 					}
					 sync_windows{
						kind = "deny"
						schedule = "0 22 * * *"
						duration = "1h"
						namespaces = ["default"]
 					}
					 sync_windows{
						kind = "allow"
						schedule = "0 23 * * *"
						duration = "1h"
						clusters = ["in-cluster", "cluster1"]
 					}
					namespace_resource_blacklist{
						group = "group"
						kind = "ResourceQuota"
					}
					namespace_resource_blacklist{
						group = "group2"
						kind = "LimitRange"
					}
					namespace_resource_blacklist{
						group = "group3"
						kind = "NetworkPolicy"
					}
					namespace_resource_whitelist{
						group = "apps"
						kind = "Deployment"
					}
					namespace_resource_whitelist{
						group = "apps"
						kind = "StatefulSet"
					}
					orphaned_resources {
						warn = "false"
					}
					source_repos = ["*"]
				}
			}
		}
	`, accountId, agentId, name, namespace, agentNamespace)
}

// TestAccResourceGitopsProjectDestinationServiceAccounts checks that
// destination_service_accounts is stored on create, changed on update, and
// returned on import. It talks to a live GitOps agent and is skipped unless
// TF_ACC=1.
//
//	export HARNESS_ENDPOINT=https://qa.harness.io/gateway
//	export HARNESS_ACCOUNT_ID=...
//	export HARNESS_PLATFORM_API_KEY=...
//	export HARNESS_TEST_GITOPS_AGENT_ID=account.<agent-id>
//	export HARNESS_TEST_GITOPS_AGENT_NAMESPACE=<namespace where the agent is installed>
//	TF_ACC=1 go test ./internal/service/platform/gitops/project/ -count=1 -timeout 20m -run TestAccResourceGitopsProjectDestinationServiceAccounts
func TestAccResourceGitopsProjectDestinationServiceAccounts(t *testing.T) {
	name := strings.ToLower(strings.ReplaceAll("dsa"+utils.RandStringBytes(6), "_", ""))
	agentID := os.Getenv("HARNESS_TEST_GITOPS_AGENT_ID")
	agentNamespace := os.Getenv("HARNESS_TEST_GITOPS_AGENT_NAMESPACE")
	accountID := os.Getenv("HARNESS_ACCOUNT_ID")
	resourceName := "harness_platform_gitops_app_project.test"

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			if agentID == "" || agentNamespace == "" {
				t.Fatal("HARNESS_TEST_GITOPS_AGENT_ID and HARNESS_TEST_GITOPS_AGENT_NAMESPACE must be set")
			}
		},
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccGitopsProjectDestinationServiceAccounts(agentID, agentNamespace, accountID, name, "argocd-deployer"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "project.0.spec.0.destination_service_accounts.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "project.0.spec.0.destination_service_accounts.0.server", "https://kubernetes.default.svc"),
					resource.TestCheckResourceAttr(resourceName, "project.0.spec.0.destination_service_accounts.0.namespace", "myapp-*"),
					resource.TestCheckResourceAttr(resourceName, "project.0.spec.0.destination_service_accounts.0.default_service_account", "argocd-deployer"),
				),
			},
			{
				Config: testAccGitopsProjectDestinationServiceAccounts(agentID, agentNamespace, accountID, name, "argocd-deployer-v2"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "project.0.spec.0.destination_service_accounts.0.default_service_account", "argocd-deployer-v2"),
					resource.TestCheckResourceAttr(resourceName, "project.0.spec.0.destination_service_accounts.0.namespace", "myapp-*"),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"upsert", "update_mask"},
				ImportStateIdFunc:       acctest.GitopsProjectImportStateIdFunc(resourceName),
			},
		},
	})
}

func testAccGitopsProjectDestinationServiceAccounts(agentID, agentNamespace, accountID, name, serviceAccount string) string {
	return fmt.Sprintf(`
		resource "harness_platform_gitops_app_project" "test" {
			account_id = "%[1]s"
			agent_id   = "%[2]s"
			upsert     = true
			project {
				metadata {
					name      = "%[3]s"
					namespace = "%[4]s"
				}
				spec {
					source_repos = ["*"]
					destinations {
						namespace = "myapp-*"
						server    = "https://kubernetes.default.svc"
					}
					destination_service_accounts {
						server                  = "https://kubernetes.default.svc"
						namespace               = "myapp-*"
						default_service_account = "%[5]s"
					}
				}
			}
		}
	`, accountID, agentID, name, agentNamespace, serviceAccount)
}
