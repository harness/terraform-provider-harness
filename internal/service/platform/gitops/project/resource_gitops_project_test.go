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
	resourceName := "harness_platform_gitops_app_project.test"
	accountId := os.Getenv("HARNESS_ACCOUNT_ID")
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		//CheckDestroy:      testAccResourceGitopsRepositoryDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceGitopsProjectAccountLevel(agentId, accountId, name, "*"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "agent_id", agentId),
				),
			},
			{
				Config: testAccResourceGitopsProjectAccountLevelUpdate(agentId, accountId, name, "roll"),
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

func testAccResourceGitopsProjectAccountLevel(agentId string, accountId string, name string, namespace string) string {
	return fmt.Sprintf(`
		resource "harness_platform_gitops_app_project" "test" {
			account_id = "%[1]s"
			agent_id = "%[2]s"
			upsert = true
			project {
				metadata {
					name = "%[3]s"
					namespace = "rollouts"
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
	`, accountId, agentId, name, namespace)
}

func testAccResourceGitopsProjectAccountLevelUpdate(agentId string, accountId string, name string, namespace string) string {
	return fmt.Sprintf(`
		resource "harness_platform_gitops_app_project" "test" {
			account_id = "%[1]s"
			agent_id = "%[2]s"
			upsert = true
			project {
				metadata {
					name = "%[3]s"
					namespace = "rollouts"
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
	`, accountId, agentId, name, namespace)
}
