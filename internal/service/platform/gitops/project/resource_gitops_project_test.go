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
	agentId := os.Getenv("HARNESS_TEST_GITOPS_AGENT_ID")
	accountId := os.Getenv("HARNESS_ACCOUNT_ID")
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		//CheckDestroy:      testAccResourceGitopsRepositoryDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceGitopsProjectAccountLevel(agentId, accountId, "14a3dc9eeee999", "*"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "agent_id", agentId),
				),
			},
			{
				Config: testAccResourceGitopsProjectAccountLevel(agentId, accountId, "14a3dc9eeee999", "rollouts"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "agent_id", agentId),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"upsert", "update_mask", "repo.0.type_"},
				ImportStateIdFunc:       acctest.GitopsWebhookImportStateIdFunc(resourceName),
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
