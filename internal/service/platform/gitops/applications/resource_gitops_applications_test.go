package applications_test

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/antihax/optional"
	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceGitopsApplication_HelmApp(t *testing.T) {
	id := strings.ToLower(fmt.Sprintf("%s%s", t.Name(), utils.RandStringBytes(5)))
	id = strings.ReplaceAll(id, "_", "")
	name := id
	agentId := os.Getenv("HARNESS_TEST_GITOPS_AGENT_ID")
	accountId := os.Getenv("HARNESS_ACCOUNT_ID")
	clusterServer := os.Getenv("HARNESS_TEST_GITOPS_CLUSTER_SERVER_APP")
	clusterId := os.Getenv("HARNESS_TEST_GITOPS_CLUSTER_ID")
	repoId := os.Getenv("HARNESS_TEST_GITOPS_REPO_ID")
	clusterName := id
	namespace := "test"
	repo := os.Getenv("HARNESS_TEST_GITOPS_REPO")
	namespaceUpdated := namespace + "_updated"
	resourceName := "harness_platform_gitops_applications.test"
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccResourceGitopsApplicationDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceGitopsApplicationHelm(id, accountId, name, agentId, clusterName, namespace, clusterServer, clusterId, repo, repoId),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
				),
			},
			{
				Config: testAccResourceGitopsApplicationHelm(id, accountId, name, agentId, clusterName, namespaceUpdated, clusterServer, clusterId, repo, repoId),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "application.0.spec.0.destination.0.namespace", namespaceUpdated),
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

func TestAccResourceGitopsApplication_HelmAppMultiSource(t *testing.T) {
	id := strings.ToLower(fmt.Sprintf("%s%s", t.Name(), utils.RandStringBytes(5)))
	id = strings.ReplaceAll(id, "_", "")
	name := id
	agentId := os.Getenv("HARNESS_TEST_GITOPS_AGENT_ID")
	accountId := os.Getenv("HARNESS_ACCOUNT_ID")
	clusterServer := os.Getenv("HARNESS_TEST_GITOPS_CLUSTER_SERVER")
	clusterToken := os.Getenv("HARNESS_TEST_GITOPS_CLUSTER_TOKEN")

	clusterId := id
	repoId := id
	clusterName := id
	namespace := "test"
	repo := os.Getenv("HARNESS_TEST_GITOPS_REPO")
	namespaceUpdated := namespace + "_updated"
	resourceName := "harness_platform_gitops_applications.testmultisource"
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccResourceGitopsApplicationDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceGitopsApplicationHelmMultiSource(id, accountId, name, agentId, clusterName, namespace, clusterServer, clusterId, repo, repoId, clusterToken),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
				),
			},
			{
				Config: testAccResourceGitopsApplicationHelmMultiSource(id, accountId, name, agentId, clusterName, namespaceUpdated, clusterServer, clusterId, repo, repoId, clusterToken),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "application.0.spec.0.destination.0.namespace", namespaceUpdated),
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

func TestAccResourceGitopsApplication_KustomizeApp(t *testing.T) {
	id := strings.ToLower(fmt.Sprintf("%s%s", t.Name(), utils.RandStringBytes(5)))
	id = strings.ReplaceAll(id, "_", "")
	name := id
	agentId := os.Getenv("HARNESS_TEST_GITOPS_AGENT_ID")
	accountId := os.Getenv("HARNESS_ACCOUNT_ID")
	clusterServer := os.Getenv("HARNESS_TEST_GITOPS_CLUSTER_SERVER_APP")
	clusterId := os.Getenv("HARNESS_TEST_GITOPS_CLUSTER_ID")
	repoId := id
	clusterName := id
	namespace := "test"
	repo := os.Getenv("HARNESS_TEST_GITOPS_REPO")
	namespaceUpdated := namespace + "_updated"
	resourceName := "harness_platform_gitops_applications.test"
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccResourceGitopsApplicationDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceGitopsApplicationKustomize(id, accountId, name, agentId, clusterName, namespace, clusterServer, clusterId, repo, repoId),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
				),
			},
			{
				Config: testAccResourceGitopsApplicationKustomize(id, accountId, name, agentId, clusterName, namespaceUpdated, clusterServer, clusterId, repo, repoId),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "application.0.spec.0.destination.0.namespace", namespaceUpdated),
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

func TestAccResourceGitopsApplicationHelmCharts_SkipRepoValidationTrue(t *testing.T) {
	id := strings.ToLower(fmt.Sprintf("%s%s", t.Name(), utils.RandStringBytes(5)))
	id = strings.ReplaceAll(id, "_", "")
	name := id
	agentId := os.Getenv("HARNESS_TEST_GITOPS_AGENT_ID")
	accountId := os.Getenv("HARNESS_ACCOUNT_ID")
	clusterServer := os.Getenv("HARNESS_TEST_GITOPS_CLUSTER_SERVER_APP")
	clusterId := os.Getenv("HARNESS_TEST_GITOPS_CLUSTER_ID")
	clusterName := id
	namespace := "test"
	repo := os.Getenv("HARNESS_TEST_GITOPS_HELM_REPO_URL")
	chart := os.Getenv("HARNESS_TEST_GITOPS_HELM_REPO_CHART")
	namespaceUpdated := namespace + "_updated"
	resourceName := "harness_platform_gitops_applications.test"
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccResourceGitopsApplicationDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceGitopsApplicationHelmSkipRepoValidation(id, accountId, name, agentId, clusterName, namespace, clusterServer, clusterId, repo, chart, true),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
				),
			},
			{
				Config: testAccResourceGitopsApplicationHelmSkipRepoValidation(id, accountId, name, agentId, clusterName, namespaceUpdated, clusterServer, clusterId, repo, chart, true),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "application.0.spec.0.destination.0.namespace", namespaceUpdated),
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

func TestAccResourceGitopsApplicationGit_SkipRepoValidationTrue(t *testing.T) {
	id := strings.ToLower(fmt.Sprintf("%s%s", t.Name(), utils.RandStringBytes(5)))
	id = strings.ReplaceAll(id, "_", "")
	name := id
	agentId := os.Getenv("HARNESS_TEST_GITOPS_AGENT_ID")
	accountId := os.Getenv("HARNESS_ACCOUNT_ID")
	clusterServer := os.Getenv("HARNESS_TEST_GITOPS_CLUSTER_SERVER_APP")
	clusterId := os.Getenv("HARNESS_TEST_GITOPS_CLUSTER_ID")
	clusterName := id
	repoId := id
	namespace := "test"
	repo := os.Getenv("HARNESS_TEST_GITOPS_REPO")
	namespaceUpdated := namespace + "_updated"
	resourceName := "harness_platform_gitops_applications.test"
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccResourceGitopsApplicationDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceGitopsApplicationGitSkipRepoValidation(id, accountId, name, agentId, clusterName, namespace, clusterServer, clusterId, repo, true),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "skip_repo_validation", "true"),
				),
			},
			{
				Config: testAccResourceGitopsApplicationGitSkipRepoValidation(id, accountId, name, agentId, clusterName, namespaceUpdated, clusterServer, clusterId, repo, true),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "application.0.spec.0.destination.0.namespace", namespaceUpdated),
				),
			},
			{
				Config: testAccResourceGitopsApplicationKustomize(id, accountId, name, agentId, clusterName, namespaceUpdated, clusterServer, clusterId, repo, repoId),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "application.0.spec.0.destination.0.namespace", namespaceUpdated),
					resource.TestCheckResourceAttr(resourceName, "skip_repo_validation", "false"),
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

func testAccResourceGitopsApplicationDestroy(resourceName string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		application, _ := testAccGetApplication(resourceName, state)
		if application != nil {
			return fmt.Errorf("Found Application: %s", application.Name)
		}
		return nil
	}
}

func testAccGetApplication(resourceName string, state *terraform.State) (*nextgen.Servicev1Application, error) {
	r := acctest.TestAccGetResource(resourceName, state)
	c, ctx := acctest.TestAccGetPlatformClientWithContext()
	agentIdentifier := r.Primary.Attributes["agent_id"]
	orgIdentifier := r.Primary.Attributes["org_id"]
	projectIdentifier := r.Primary.Attributes["project_id"]
	queryName := r.Primary.Attributes["name"]
	repoIdentifier := r.Primary.Attributes["repo_id"]

	resp, _, err := c.ApplicationsApiService.AgentApplicationServiceGet(ctx, agentIdentifier, queryName, c.AccountId, orgIdentifier, projectIdentifier, &nextgen.ApplicationsApiAgentApplicationServiceGetOpts{
		QueryRepo: optional.NewString(repoIdentifier),
	})
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

func testAccResourceGitopsApplicationHelm(id string, accountId string, name string, agentId string, clusterName string, namespace string, clusterServer string, clusterId string, repo string, repoId string) string {
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

		resource "harness_platform_service" "test" {
			identifier = "%[1]s"
			name = "%[3]s"
			org_id = harness_platform_project.test.org_id
			project_id = harness_platform_project.test.id
		}
		resource "harness_platform_environment" "test" {
			identifier = "%[1]s"
			name = "%[3]s"
			org_id = harness_platform_project.test.org_id
			project_id = harness_platform_project.test.id
			tags = ["foo:bar", "baz"]
			type = "PreProduction"
		}

		resource "harness_platform_gitops_repository" "test" {
			identifier = "%[1]s"
			account_id = "%[2]s"
			project_id = harness_platform_project.test.id
			org_id = harness_platform_organization.test.id
			agent_id = "%[4]s"
			repo {
					repo = "https://github.com/harness-apps/hosted-gitops-example-apps"
					name = "%[1]s"
					insecure = true
					connection_type = "HTTPS_ANONYMOUS"
			}
			upsert = true
		}

		resource "harness_platform_gitops_applications" "test" {
			depends_on = [harness_platform_gitops_repository.test]
			application {
				metadata {
					annotations = {}
					labels = {
						"harness.io/serviceRef" = harness_platform_service.test.id
						"harness.io/envRef" = harness_platform_environment.test.id
					}
					name = "%[1]s"
				}
				spec {
					sync_policy {
						sync_options = [
							"PrunePropagationPolicy=undefined",
							"CreateNamespace=false",
							"Validate=false",
							"skipSchemaValidations=false",
							"autoCreateNamespace=false",
							"pruneLast=false",
							"applyOutofSyncOnly=false",
							"Replace=false",
							"retry=false"
						]
					}
					source {
						target_revision = "master"
						repo_url = "%[9]s"
						path = "helm-guestbook"
						
					}
					destination {
						namespace = "%[6]s"
						server = "%[7]s"
					}
				}
			}
			project_id = harness_platform_project.test.id
			org_id = harness_platform_organization.test.id
			account_id = "%[2]s"
			identifier = "%[1]s"
			cluster_id = "%[8]s"
			repo_id = "%[10]s"
			agent_id = "%[4]s"
			name = "%[3]s"
		}
		`, id, accountId, name, agentId, clusterName, namespace, clusterServer, clusterId, repo, repoId)
}

func testAccResourceGitopsApplicationHelmMultiSource(id string, accountId string, name string, agentId string, clusterName string, namespace string, clusterServer string, clusterId string, repo string, repoId string, clusterToken string) string {
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

		resource "harness_platform_service" "test" {
			identifier = "%[1]s"
			name = "%[3]s"
			org_id = harness_platform_project.test.org_id
			project_id = harness_platform_project.test.id
		}
		resource "harness_platform_environment" "test" {
			identifier = "%[1]s"
			name = "%[3]s"
			org_id = harness_platform_project.test.org_id
			project_id = harness_platform_project.test.id
			tags = ["foo:bar", "baz"]
			type = "PreProduction"
		}

		resource "harness_platform_gitops_repository" "test" {
			identifier = "%[1]s"
			account_id = "%[2]s"
			project_id = harness_platform_project.test.id
			org_id = harness_platform_organization.test.id
			agent_id = "%[4]s"
			repo {
					repo = "https://github.com/harness-apps/hosted-gitops-example-apps"
					name = "%[1]s"
					insecure = true
					connection_type = "HTTPS_ANONYMOUS"
			}
			upsert = true
		}

		resource "harness_platform_gitops_cluster" "test" {
			identifier = "%[8]s"
			account_id = "%[2]s"
			agent_id = "%[4]s"
			project_id = harness_platform_project.test.id
			org_id = harness_platform_organization.test.id
			request {
				upsert = true
				cluster {
					server = "%[7]s"
					name = "%[8]s"
					config {
						bearer_token = "%[11]s"
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

		resource "harness_platform_gitops_applications" "testmultisource" {
			depends_on = [harness_platform_gitops_repository.test, harness_platform_gitops_cluster.test, harness_platform_service.test, harness_platform_environment.test, harness_platform_project.test, harness_platform_organization.test]
			application {
				metadata {
					annotations = {}
					labels = {
						"harness.io/serviceRef" = harness_platform_service.test.id
						"harness.io/envRef" = harness_platform_environment.test.id
					}
					name = "%[1]s"
				}
				spec {
					sync_policy {
						sync_options = [
							"PrunePropagationPolicy=undefined",
							"CreateNamespace=false",
							"Validate=false",
							"skipSchemaValidations=false",
							"autoCreateNamespace=false",
							"pruneLast=false",
							"applyOutofSyncOnly=false",
							"Replace=false",
							"retry=false"
						]
					}
					sources {
							repo_url = "%[9]s"
							target_revision = "master"
							ref = "val"
					}
					sources {
							repo_url = "%[9]s"
							target_revision = "master"
							path = "helm-guestbook"
							helm {
							  value_files = [
								"$val/helm-guestbook/values.yaml"
							  ]
							}
					}
					destination {
						namespace = "%[6]s"
						server = "%[7]s"
					}
				}
			}
			project_id = harness_platform_project.test.id
			org_id = harness_platform_organization.test.id
			account_id = "%[2]s"
			name = "%[3]s"
			cluster_id = "%[8]s"
			repo_ids = [
				"%[1]s",
				"%[1]s"
			]
			agent_id = "%[4]s"
		}
		`, id, accountId, name, agentId, clusterName, namespace, clusterServer, clusterId, repo, repoId, clusterToken)
}

func testAccResourceGitopsApplicationHelmSkipRepoValidation(id string, accountId string, name string, agentId string, clusterName string, namespace string, clusterServer string, clusterId string, repoURL string, chart string, skipRepoValidation bool) string {
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

		resource "harness_platform_gitops_applications" "test" {
			application {
				metadata {
					annotations = {}
					name = "%[1]s"
				}
				spec {
					sync_policy {
						sync_options = [
							"PrunePropagationPolicy=undefined",
							"CreateNamespace=false",
							"Validate=false",
							"skipSchemaValidations=false",
							"autoCreateNamespace=false",
							"pruneLast=false",
							"applyOutofSyncOnly=false",
							"Replace=false",
							"retry=false"
						]
					}
					source {
						target_revision = "18.0.1"
						repo_url = "%[9]s"
						chart = "%[10]s"
						
					}
					destination {
						namespace = "%[6]s"
						server = "%[7]s"
					}
				}
			}
			project_id = harness_platform_project.test.id
			org_id = harness_platform_organization.test.id
			account_id = "%[2]s"
			identifier = "%[1]s"
			cluster_id = "%[8]s"
			agent_id = "%[4]s"
			name = "%[3]s"
            skip_repo_validation = %[11]t
		}
		`, id, accountId, name, agentId, clusterName, namespace, clusterServer, clusterId, repoURL, chart, skipRepoValidation)
}

func testAccResourceGitopsApplicationKustomize(id string, accountId string, name string, agentId string, clusterName string, namespace string, clusterServer string, clusterId string, repo string, repoId string) string {
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

		resource "harness_platform_service" "test" {
      		identifier = "%[1]s"
      		name = "%[3]s"
      		org_id = harness_platform_project.test.org_id
      		project_id = harness_platform_project.test.id
    	}
		resource "harness_platform_environment" "test" {
			identifier = "%[1]s"
			name = "%[3]s"
			org_id = harness_platform_project.test.org_id
			project_id = harness_platform_project.test.id
			tags = ["foo:bar", "baz"]
			type = "PreProduction"
  		}
		
		resource "harness_platform_gitops_repository" "test" {
			identifier = "%[1]s"
			account_id = "%[2]s"
			project_id = harness_platform_project.test.id
			org_id = harness_platform_organization.test.id
			agent_id = "%[4]s"
			repo {
					repo = "%[9]s"
        			name = "%[1]s"
        			insecure = true
        			connection_type = "HTTPS_ANONYMOUS"
			}
			upsert = true
		}

		resource "harness_platform_gitops_applications" "test" {
			depends_on = [harness_platform_gitops_repository.test]
			application {
				metadata {
					annotations = {}
					labels = {
						"harness.io/serviceRef" = harness_platform_service.test.id
						"harness.io/envRef" = harness_platform_environment.test.id
					}
					name = "%[1]s"
				}
				spec {
					sync_policy {
						sync_options = [
							"PrunePropagationPolicy=undefined",
							"CreateNamespace=false",
							"Validate=false",
							"skipSchemaValidations=false",
							"autoCreateNamespace=false",
							"pruneLast=false",
							"applyOutofSyncOnly=false",
							"Replace=false",
							"retry=false"
						]
					}
					source {
						target_revision = "master"
						repo_url = "%[9]s"
						path = "kustomize-guestbook"
						kustomize {
							images = [
									"gcr.io/heptio-images/ks-guestbook-demo:0.1"
									]
						}
					}
					destination {
						namespace = "%[6]s"
						server = "%[7]s"
					}
				}
			}
			project_id = harness_platform_project.test.id
			org_id = harness_platform_organization.test.id
			account_id = "%[2]s"
			identifier = "%[1]s"
			cluster_id =  "%[8]s"
			repo_id = "%[10]s"
			agent_id = "%[4]s"
			name = "%[3]s"
		}
		`, id, accountId, name, agentId, clusterName, namespace, clusterServer, clusterId, repo, repoId)
}

func testAccResourceGitopsApplicationGitSkipRepoValidation(id string, accountId string, name string, agentId string, clusterName string, namespace string, clusterServer string, clusterId string, repo string, skipRepoValidation bool) string {
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

		resource "harness_platform_service" "test" {
      		identifier = "%[1]s"
      		name = "%[3]s"
      		org_id = harness_platform_project.test.org_id
      		project_id = harness_platform_project.test.id
    	}
		resource "harness_platform_environment" "test" {
			identifier = "%[1]s"
			name = "%[3]s"
			org_id = harness_platform_project.test.org_id
			project_id = harness_platform_project.test.id
			tags = ["foo:bar", "baz"]
			type = "PreProduction"
  		}
		

		resource "harness_platform_gitops_applications" "test" {
			application {
				metadata {
					annotations = {}
					labels = {
						"harness.io/serviceRef" = harness_platform_service.test.id
						"harness.io/envRef" = harness_platform_environment.test.id
					}
					name = "%[1]s"
				}
				spec {
					sync_policy {
						sync_options = [
							"PrunePropagationPolicy=undefined",
							"CreateNamespace=false",
							"Validate=false",
							"skipSchemaValidations=false",
							"autoCreateNamespace=false",
							"pruneLast=false",
							"applyOutofSyncOnly=false",
							"Replace=false",
							"retry=false"
						]
					}
					source {
						target_revision = "master"
						repo_url = "%[9]s"
						path = "kustomize-guestbook"
						kustomize {
							images = [
									"gcr.io/heptio-images/ks-guestbook-demo:0.1"
									]
						}
					}
					destination {
						namespace = "%[6]s"
						server = "%[7]s"
					}
				}
			}
			project_id = harness_platform_project.test.id
			org_id = harness_platform_organization.test.id
			account_id = "%[2]s"
			identifier = "%[1]s"
			cluster_id =  "%[8]s"
			agent_id = "%[4]s"
			name = "%[3]s"
            skip_repo_validation = %[10]t
		}
		`, id, accountId, name, agentId, clusterName, namespace, clusterServer, clusterId, repo, skipRepoValidation)
}
