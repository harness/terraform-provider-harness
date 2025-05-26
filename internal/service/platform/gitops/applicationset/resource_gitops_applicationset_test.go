package applicationset

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
)

func TestAccResourceGitopsApplicationSet_AllClustersGenerator(t *testing.T) {
	id := strings.ToLower(fmt.Sprintf("%s%s", t.Name(), utils.RandStringBytes(5)))
	id = strings.ReplaceAll(id, "_", "")
	name := id
	agentId := os.Getenv("HARNESS_TEST_GITOPS_AGENT_ID")
	accountId := os.Getenv("HARNESS_ACCOUNT_ID")
	namespace := "test"
	resourceName := "harness_platform_gitops_applicationset.test"
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceGitopsApplicationsetClusterGenerator(id, accountId, name, agentId, namespace),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: acctest.GitopsAgentProjectLevelResourceImportStateIdFunc(resourceName),
			},
		},
	})
}

func testAccResourceGitopsApplicationsetClusterGenerator(id, accountId, name, agentId, namespace string) string {
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

		resource "harness_platform_gitops_applicationset" "test" {
			applicationset {
				metadata {
				  name      = "%[1]s"
				  namespace = "%[5]s"
				}
				spec {
				  go_template = true
				  
				  generator {
					clusters {
						enabled = true
					}
				  }
				  template {
					metadata {
					  name = "{{.name}}-guestbook"
					}
					spec {
					  project = "default"
					  source {
						repo_url        = "https://github.com/argoproj/argocd-example-apps.git"
						path            = "helm-guestbook"
						target_revision = "HEAD"
					  }
					  destination {
						server    = "{{.url}}"
						namespace = "%[5]s"
					  }
					}
				  }
				}
			  }
			project_id = harness_platform_project.test.id
			org_id = harness_platform_organization.test.id
		  	agent_id   = "%[4]s"
		  	upsert     = true
		}
		`, id, accountId, name, agentId, namespace)
}
