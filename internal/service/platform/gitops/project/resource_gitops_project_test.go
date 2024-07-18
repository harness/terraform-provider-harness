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
	resourceName := "harness_platform_gitops_project.test"
	agentId := "account.rollouts"
	accountId := "1bvyLackQK-Hapk25-Ry4w"
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		//CheckDestroy:      testAccResourceGitopsRepositoryDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceGitopsProjectAccountLevel(agentId, accountId, "14a3dc9eeee9990deeeju1", "*"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "agent_id", agentId),
				),
			},
			{
				Config: testAccResourceGitopsProjectAccountLevel(agentId, accountId, "14a3dc9eeee9990deeeju1", "roll"),
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
		resource "harness_platform_gitops_project" "test" {
			account_id = "%[1]s"
			agent_id = "%[2]s"
			upsert = true
			project {
				metadata {
					name = "%[3]s"
					namespace = "rollouts"
					finalizers = ["name"]
					generate_name = "%[3]s"
					labels = {
						v1 = "k1"
					}
					annotations = {
						v1 = "k1"
					}
					owner_references {
						name = "t1"
						kind = "t2"
						api_version = "v1"
						uid = "uid"
					}					
					managed_fields {
						manager = "agent"
						operation = "Update"
						time      = {}
						fields_v1 = {}
					}
				}
				spec {
					cluster_resource_whitelist {
						group = "*"
						kind = "*"
					}
					destinations {
						namespace = "rollouts"
						server = "*"
						name = "%[3]s"
					}
					roles {
						name = "read-only"
						description = "Read-only privileges to my-project"
						policies = ["proj:my-project:read-only", "applications", "get", "my-project/*", "allow"]
						jwt_tokens {
							iat = "iat"
							exp = "exp"
							id = "id"
						}
						groups = ["my-oidc-group"]

					}
					sync_windows{
						kind = "allow"
						schedule = "10 1 * * *"
						duration = "1h"
						applications = ["*-prod"]
						namespaces = ["rollouts"]
						clusters = ["in-cluster"]
						manual_sync = "true"
						time_zone = "time_zone"
 					}
					namespace_resource_whitelist{
						group = "*"
						kind = "*"
					}

					cluster_resource_blacklist{
						group = "*"
						kind = "*"
					}
					signature_keys {
						key_id = "*"
					}
					source_repos = ["*"]
				}
			}
		}
	`, accountId, agentId, name, namespace)
}

func testAccResourceGitopsProjectUpdateAccountLevel(agentId string, accountId string, name string, namespace string) string {
	return fmt.Sprintf(`	
		resource "harness_platform_gitops_project" "test" {
			account_id = "%[1]s"
			agent_id = "%[2]s"
			upsert = true
			project {
				metadata {
					generation = "1"
					name = "%[3]s"
					namespace = "rollouts"
				}
				spec {
					cluster_resource_whitelist {
						group = "*"
						kind = "*"
					}
					destinations {
						namespace = "%[4]s"
						server = "*"
					}
					source_repos = ["*"]
				}
			}
		}
	`, accountId, agentId, name, namespace)
}
func TestAccResourceGitopsProjectOrgLevel(t *testing.T) {
	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	id = strings.ReplaceAll(id, "_", "")
	name := id
	resourceName := "harness_platform_gitops_project.test"
	agentId := os.Getenv("HARNESS_TEST_GITOPS_AGENT_ID")
	accountId := os.Getenv("HARNESS_ACCOUNT_ID")
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		//CheckDestroy:      testAccResourceGitopsRepositoryDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceGitopsProjectOrgLevel(id, name, agentId, accountId, "14a3dc9eeee999", "*"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "agent_id", agentId),
				),
			},
			{
				Config: testAccResourceGitopsProjectOrgLevel(id, name, agentId, accountId, "14a3dc9eeee999", "rollouts"),
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

func testAccResourceGitopsProjectOrgLevel(id string, name string, agentId string, accountId string, metadat_name string, namespace string) string {
	return fmt.Sprintf(`
	    resource "harness_platform_organization" "test" {
		    identifier = "%[1]s"
		    name = "%[2]s"
	    }
		resource "harness_platform_gitops_project" "test" {
			account_id = "%[3]s"
			org_id = harness_platform_organization.test.id
			agent_id = "%[4]s"
			upsert = true
			project {
				metadata {
					generation = "1"
					name = "%[5]s"
					namespace = "rollouts"
				}
				spec {
					cluster_resource_whitelist {
						group = "*"
						kind = "*"
					}
					destinations {
						namespace = "%[6]s"
						server = "*"
					}
					source_repos = ["*"]
				}
			}
		}
	`, id, name, accountId, agentId, metadat_name, namespace)
}
