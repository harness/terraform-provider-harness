package applicationset_test

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
	// namespace is same as agent id but remove "account." prefix if it exists
	namespace := strings.TrimPrefix(agentId, "account.")
	//namespace = "argocd"
	resourceName := "harness_platform_gitops_applicationset.test"
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceGitopsApplicationsetClusterGenerator(id, accountId, name, agentId, namespace),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "applicationset.0.metadata.0.name", id),
				),
			},

			{
				ResourceName:      resourceName,
				ImportState:       true,
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

		resource "harness_platform_gitops_applicationset" "test1" {
			applicationset {
				metadata {
				  name      = "%[1]s"
				  namespace = "%[5]s"
				}
				spec {
				  go_template = true
				  go_template_options = ["missingkey=error"]

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
			lifecycle {
			  ignore_changes = [
				applicationset.0.spec.0.generator.0.clusters,
				applicationset.0.spec.0.template.0.metadata.0.annotations,
				applicationset.0.spec.0.template.0.metadata.0.labels,
				applicationset.0.spec.0.template.0.metadata.0 .finalizers,
			  ]
			}
		}


		resource "harness_platform_gitops_applicationset" "test2" {
			applicationset {
				metadata {
				  name      = "%[1]s"
				  namespace = "%[5]s"
				}
				spec {
				  go_template = true
				  go_template_options = ["missingkey=error"]

				  generator {
					git {
						repo_url = "https://github.com/argoproj/argocd-example-apps.git"
						path = "helm-guestbook"
						target_revision = "HEAD"
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
				project_id = harness_platform_project.test.id
				org_id = harness_platform_organization.test.id
				agent_id   = "%[4]s"
				upsert     = true

			}
		}

		resource "harness_platform_gitops_applicationset" "test3" {
			applicationset {
				metadata {
				  name      = "%[1]s"
				  namespace = "%[5]s"
				}
				spec {
				  go_template = true
				  go_template_options = ["missingkey=error"]

				  generator {
					list {
						items = ["item1", "item2", "item3"]
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
				project_id = harness_platform_project.test.id
				org_id = harness_platform_organization.test.id
				agent_id   = "%[4]s"
				upsert     = true

			}
		}

		`, id, accountId, name, agentId, namespace)

}

func TestAccResourceGitopsApplicationSet_MatrixGitClusters(t *testing.T) {
	id := strings.ToLower(fmt.Sprintf("%s%s", t.Name(), utils.RandStringBytes(5)))
	id = strings.ReplaceAll(id, "_", "")
	name := id
	agentId := os.Getenv("HARNESS_TEST_GITOPS_AGENT_ID")
	accountId := os.Getenv("HARNESS_ACCOUNT_ID")
	namespace := strings.TrimPrefix(agentId, "account.")
	resourceName := "harness_platform_gitops_applicationset.test"
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceGitopsApplicationsetMatrixGitClusters(id, accountId, name, agentId, namespace),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "applicationset.0.metadata.0.name", id),
					resource.TestCheckResourceAttr(resourceName, "applicationset.0.spec.0.generator.0.matrix.0.generator.0.git.0.repo_url", "https://github.com/mteodor/argocd-example-apps"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateIdFunc: acctest.GitopsAgentProjectLevelResourceImportStateIdFunc(resourceName),
			},
		},
	})
}

func testAccResourceGitopsApplicationsetMatrixGitClusters(id, accountId, name, agentId, namespace string) string {
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
				  go_template_options = ["missingkey=error"]

				  generator {
					matrix {
					  generator {
						git {
						  repo_url = "https://github.com/mteodor/argocd-example-apps"
						  revision = "master"
						  file {
							path = "helm2/app1/*/*/config.json"
						  }
						}
					  }
					  generator {
						clusters {
						  enabled = true
						}
					  }
					}
				  }
				  template {
					metadata {
					  name = "external-dns-{{.path.basename}}-{{.name}}"
					}
					spec {
					  project = "default"
					  source {
						repo_url        = "https://github.com/mteodor/argocd-example-apps"
						target_revision = "master"
						chart           = "nginx"
					  }
					  destination {
						server    = "{{.url}}"
						namespace = "external-dns"
					  }
					  sync_policy {
						automated {
						  prune       = true
						  self_heal   = true
						  allow_empty = true
						}
						sync_options = [
						  "CreateNamespace=true",
						  "PruneLast=true",
						  "Replace=false"
						]
						retry {
						  limit = 5
						  backoff {
							duration     = "5s"
							factor       = 2
							max_duration = "3m"
						  }
						}
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

func TestAccResourceGitopsApplicationSet_SCMProviderGitHub(t *testing.T) {
	id := strings.ToLower(fmt.Sprintf("%s%s", t.Name(), utils.RandStringBytes(5)))
	id = strings.ReplaceAll(id, "_", "")
	name := id
	agentId := os.Getenv("HARNESS_TEST_GITOPS_AGENT_ID")
	accountId := os.Getenv("HARNESS_ACCOUNT_ID")
	namespace := strings.TrimPrefix(agentId, "account.")
	resourceName := "harness_platform_gitops_applicationset.test"
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceGitopsApplicationsetSCMProviderGitHub(id, accountId, name, agentId, namespace),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "applicationset.0.metadata.0.name", id),
					resource.TestCheckResourceAttr(resourceName, "applicationset.0.spec.0.generator.0.scm_provider.0.github.0.organization", "myorg"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateIdFunc: acctest.GitopsAgentProjectLevelResourceImportStateIdFunc(resourceName),
			},
		},
	})
}

func testAccResourceGitopsApplicationsetSCMProviderGitHub(id, accountId, name, agentId, namespace string) string {
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
				  go_template_options = ["missingkey=error"]

				  generator {
					scm_provider {
					  github {
						organization = "myorg"
						api = "https://git.example.com/"
						all_branches = true
						token_ref {
						  secret_name = "github-token"
						  key = "token"
						}
						app_secret_name = "gh-app-repo-creds"
					  }
					}
				  }
				  template {
					metadata {
					  name = "external-dns-{{.path.basename}}-{{.name}}"
					}
					spec {
					  project = "default"
					  source {
						repo_url        = "https://github.com/mteodor/argocd-example-apps"
						target_revision = "master"
						chart           = "nginx"
					  }
					  destination {
						server    = "{{.url}}"
						namespace = "external-dns"
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

func TestAccResourceGitopsApplicationSet_SCMProviderGitLab(t *testing.T) {
	id := strings.ToLower(fmt.Sprintf("%s%s", t.Name(), utils.RandStringBytes(5)))
	id = strings.ReplaceAll(id, "_", "")
	name := id
	agentId := os.Getenv("HARNESS_TEST_GITOPS_AGENT_ID")
	accountId := os.Getenv("HARNESS_ACCOUNT_ID")
	namespace := strings.TrimPrefix(agentId, "account.")
	resourceName := "harness_platform_gitops_applicationset.test"
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceGitopsApplicationsetSCMProviderGitLab(id, accountId, name, agentId, namespace),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "applicationset.0.metadata.0.name", id),
					resource.TestCheckResourceAttr(resourceName, "applicationset.0.spec.0.generator.0.scm_provider.0.gitlab.0.group", "testgroup"),
					resource.TestCheckResourceAttr(resourceName, "applicationset.0.spec.0.generator.0.scm_provider.0.gitlab.0.ca_ref.0.config_map_name", "gitlab-ca"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateIdFunc: acctest.GitopsAgentProjectLevelResourceImportStateIdFunc(resourceName),
			},
		},
	})
}

func testAccResourceGitopsApplicationsetSCMProviderGitLab(id, accountId, name, agentId, namespace string) string {
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
				  go_template_options = ["missingkey=error"]

				  generator {
					scm_provider {
					  gitlab {
						group = "testgroup"
						api = "https://git.example.com/"
						all_branches = true
						token_ref {
						  secret_name = "github-token"
						  key = "token"
						}
						ca_ref {
						  config_map_name = "gitlab-ca"
						  key = "ca"
						}
					  }
					}
				  }
				  template {
					metadata {
					  name = "external-dns-{{.path.basename}}-{{.name}}"
					}
					spec {
					  project = "default"
					  source {
						repo_url        = "https://github.com/mteodor/argocd-example-apps"
						target_revision = "master"
						chart           = "nginx"
					  }
					  destination {
						server    = "{{.url}}"
						namespace = "external-dns"
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

func TestAccResourceGitopsApplicationSet_MatrixSCMProviderClustersWithSelector(t *testing.T) {
	id := strings.ToLower(fmt.Sprintf("%s%s", t.Name(), utils.RandStringBytes(5)))
	id = strings.ReplaceAll(id, "_", "")
	name := id
	agentId := os.Getenv("HARNESS_TEST_GITOPS_AGENT_ID")
	accountId := os.Getenv("HARNESS_ACCOUNT_ID")
	namespace := strings.TrimPrefix(agentId, "account.")
	resourceName := "harness_platform_gitops_applicationset.test"
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceGitopsApplicationsetMatrixSCMProviderClustersWithSelector(id, accountId, name, agentId, namespace),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "applicationset.0.metadata.0.name", id),
					resource.TestCheckResourceAttr(resourceName, "applicationset.0.spec.0.generator.0.matrix.0.generator.0.scm_provider.0.github.0.organization", "myorg"),
					resource.TestCheckResourceAttr(resourceName, "applicationset.0.spec.0.generator.0.matrix.0.generator.1.clusters.0.selector.0.match_labels.staging", "true"),
					resource.TestCheckResourceAttr(resourceName, "applicationset.0.spec.0.generator.0.matrix.0.generator.1.clusters.0.selector.0.match_expressions.0.key", "staging"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateIdFunc: acctest.GitopsAgentProjectLevelResourceImportStateIdFunc(resourceName),
			},
		},
	})
}

func testAccResourceGitopsApplicationsetMatrixSCMProviderClustersWithSelector(id, accountId, name, agentId, namespace string) string {
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
				  go_template_options = ["missingkey=error"]

				  generator {
					matrix {
					  generator {
						scm_provider {
						  github {
							organization = "myorg"
							api = "https://git.example.com/"
							all_branches = true
							token_ref {
							  secret_name = "github-token"
							  key = "token"
							}
							app_secret_name = "gh-app-repo-creds"
						  }
						}
					  }
					  generator {
						clusters {
						  selector {
							match_labels = {
							  staging = "true"
							}
							match_expressions {
							  key      = "staging"
							  operator = "In"
							  values   = ["true"]
							}
						  }
						}
					  }
					}
				  }
				  template {
					metadata {
					  name = "external-dns-{{.path.basename}}-{{.name}}"
					}
					spec {
					  project = "default"
					  source {
						repo_url        = "https://github.com/mteodor/argocd-example-apps"
						target_revision = "master"
						chart           = "nginx"
					  }
					  destination {
						server    = "{{.url}}"
						namespace = "external-dns"
					  }
					  sync_policy {
						automated {
						  prune       = true
						  self_heal   = true
						  allow_empty = true
						}
						sync_options = [
						  "CreateNamespace=true",
						  "PruneLast=true",
						  "Replace=false"
						]
						retry {
						  limit = 5
						  backoff {
							duration     = "5s"
							factor       = 2
							max_duration = "3m"
						  }
						}
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
