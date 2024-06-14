package repository_test

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/antihax/optional"
	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccResourceGitopsRepositoryOrgLevel(t *testing.T) {

	// Org level
	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	id = strings.ReplaceAll(id, "_", "")
	name := id
	repo := "https://github.com/willycoll/argocd-example-apps"
	repoName := id
	agentId := os.Getenv("HARNESS_TEST_GITOPS_AGENT_ID")
	resourceName := "harness_platform_gitops_repository.test"
	accountId := os.Getenv("HARNESS_ACCOUNT_ID")
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		//CheckDestroy:      testAccResourceGitopsRepositoryDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceGitopsRepositoryOrgLevel(id, name, repo, repoName, agentId, accountId),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
				),
			},
			{
				Config: testAccResourceGitopsRepositoryOrgLevel(id, name, repo, repoName, agentId, accountId),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"upsert", "update_mask", "repo.0.type_"},
				ImportStateIdFunc:       acctest.GitopsAgentOrgLevelResourceImportStateIdFunc(resourceName),
			},
		},
	})

}
func TestAccResourceGitopsRepository(t *testing.T) {
	// Project level
	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	id = strings.ReplaceAll(id, "_", "")
	name := id
	repo := "https://github.com/willycoll/argocd-example-apps.git"
	repoName := id
	agentId := os.Getenv("HARNESS_TEST_GITOPS_AGENT_ID")
	resourceName := "harness_platform_gitops_repository.test"
	accountId := os.Getenv("HARNESS_ACCOUNT_ID")
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		//CheckDestroy:      testAccResourceGitopsRepositoryDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceGitopsRepositoryProjectLevel(id, name, repo, repoName, agentId, accountId),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
				),
			},
			{
				Config: testAccResourceGitopsRepositoryProjectLevel(id, name, repo, repoName, agentId, accountId),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"upsert", "update_mask", "repo.0.type_"},
				ImportStateIdFunc:       acctest.GitopsAgentProjectLevelResourceImportStateIdFunc(resourceName),
			},
		},
	})

	// Account level
	id = fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	id = strings.ReplaceAll(id, "_", "")
	name = id
	repo = "https://github.com/willycoll/argocd-example-apps.git"
	repoName = id
	resourceName = "harness_platform_gitops_repository.test"
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		//CheckDestroy:      testAccResourceGitopsRepositoryDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceGitopsRepositoryAccountLevel(id, name, repo, repoName, agentId, accountId),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
				),
			},
			{
				Config: testAccResourceGitopsRepositoryAccountLevel(id, name, repo, repoName, agentId, accountId),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"upsert", "update_mask", "repo.0.type_"},
				ImportStateIdFunc:       acctest.GitopsAgentAccountLevelResourceImportStateIdFunc(resourceName),
			},
		},
	})
}

func testAccGetRepository(resourceName string, state *terraform.State) (*nextgen.Servicev1Repository, error) {
	r := acctest.TestAccGetResource(resourceName, state)
	c, ctx := acctest.TestAccGetPlatformClientWithContext()
	agentIdentifier := r.Primary.Attributes["agent_id"]
	identifier := r.Primary.Attributes["identifier"]
	resp, _, err := c.RepositoriesApiService.AgentRepositoryServiceGet(ctx, agentIdentifier, identifier, c.AccountId, &nextgen.RepositoriesApiAgentRepositoryServiceGetOpts{
		OrgIdentifier:     optional.NewString(r.Primary.Attributes["org_id"]),
		ProjectIdentifier: optional.NewString(r.Primary.Attributes["project_id"]),
		QueryRepo:         optional.NewString(r.Primary.Attributes["query_repo"]),
		QueryForceRefresh: optional.NewBool(r.Primary.Attributes["query_force_refresh"] == "True"),
		QueryProject:      optional.NewString(r.Primary.Attributes["query_project"]),
	})
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

func testAccResourceGitopsRepositoryDestroy(resourceName string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		repo, _ := testAccGetRepository(resourceName, state)
		if repo != nil {
			return fmt.Errorf("Found Repository: %s", repo.Identifier)
		}
		return nil
	}
}

func testAccResourceGitopsRepositoryProjectLevel(id string, name string, repo string, repoName string, agentId string, accountId string) string {
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

		resource "harness_platform_gitops_repository" "test" {
			identifier = "%[1]s"
			account_id = "%[6]s"
			project_id = harness_platform_project.test.id
			org_id = harness_platform_organization.test.id
			agent_id = "%[5]s"
			repo {
					repo = "%[3]s"
        			name = "%[4]s"
        			insecure = true
        			connection_type = "HTTPS_ANONYMOUS"
			}
			upsert = true
			update_mask {
				paths = ["name"]
			}

		}
	`, id, name, repo, repoName, agentId, accountId)
}

func testAccResourceGitopsRepositoryOrgLevel(id string, name string, repo string, repoName string, agentId string, accountId string) string {
	return fmt.Sprintf(`
		resource "harness_platform_organization" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
		}

		resource "harness_platform_gitops_repository" "test" {
			identifier = "%[1]s"
			account_id = "%[6]s"
			org_id = harness_platform_organization.test.id
			agent_id = "%[5]s"
			repo {
					repo = "%[3]s"
        			name = "%[4]s"
        			insecure = true
        			connection_type = "HTTPS_ANONYMOUS"
			}
			upsert = true
			update_mask {
				paths = ["name"]
			}

		}
	`, id, name, repo, repoName, agentId, accountId)
}

func testAccResourceGitopsRepositoryAccountLevel(id string, name string, repo string, repoName string, agentId string, accountId string) string {
	return fmt.Sprintf(`
		resource "harness_platform_gitops_repository" "test" {
			identifier = "%[1]s"
			account_id = "%[6]s"
			agent_id = "%[5]s"
			repo {
					repo = "%[3]s"
        			name = "%[4]s"
        			insecure = true
        			connection_type = "HTTPS_ANONYMOUS"
			}
			upsert = true
			update_mask {
				paths = ["name"]
			}

		}
	`, id, name, repo, repoName, agentId, accountId)
}

func TestHelmRepoOCIECR(t *testing.T) {
	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	id = strings.ReplaceAll(id, "_", "")
	name := id
	repo := os.Getenv("HARNESS_TEST_ECR_REPO")
	agentId := os.Getenv("HARNESS_TEST_GITOPS_AGENT_ID")
	repoName := id
	resourceName := "harness_platform_gitops_repository.test"
	accountId := os.Getenv("HARNESS_ACCOUNT_ID")
	awsAccessId := os.Getenv("HARNESS_TEST_AWS_ACCESS_KEY_ID")
	awsAccessKey := os.Getenv("HARNESS_TEST_AWS_SECRET_ACCESS_KEY")
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testHelmRepoOCIECR(id, name, repo, repoName, agentId, accountId, awsAccessId, awsAccessKey),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "repo.0.name", repoName),
				),
			},
		},
	})

}

func TestHelmRepoOCIECRJwt(t *testing.T) {
	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	id = strings.ReplaceAll(id, "_", "")
	name := id
	repo := os.Getenv("HARNESS_TEST_ECR_REPO")
	agentId := os.Getenv("HARNESS_TEST_GITOPS_AGENT_ID")
	repoName := id
	resourceName := "harness_platform_gitops_repository.test"
	accountId := os.Getenv("HARNESS_ACCOUNT_ID")
	awsAccessId := os.Getenv("HARNESS_TEST_AWS_ACCESS_KEY_ID")
	awsAccessKey := os.Getenv("HARNESS_TEST_AWS_SECRET_ACCESS_KEY")
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testHelmRepoOCIECRJwt(id, name, repo, repoName, agentId, accountId, awsAccessId, awsAccessKey),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "repo.0.name", repoName),
				),
			},
		},
	})

}

func TestHelmRepoOCIGCRWorkload(t *testing.T) {
	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	id = strings.ReplaceAll(id, "_", "")
	name := id
	repo := os.Getenv("HARNESS_TEST_GITOPS_GCR_REPO")
	agentId := os.Getenv("HARNESS_TEST_GITOPS_AGENT_ID")
	repoName := id
	resourceName := "harness_platform_gitops_repository.test"
	accountId := os.Getenv("HARNESS_ACCOUNT_ID")
	awsAccessId := os.Getenv("HARNESS_TEST_AWS_ACCESS_KEY_ID")
	awsAccessKey := os.Getenv("HARNESS_TEST_AWS_SECRET_ACCESS_KEY")
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testHelmRepoOCIGCRWorkload(id, name, repo, repoName, agentId, accountId, awsAccessId, awsAccessKey),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "repo.0.name", repoName),
				),
			},
		},
	})

}

func TestHelmRepoOCIGCRServiceAccount(t *testing.T) {
	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	id = strings.ReplaceAll(id, "_", "")
	name := id
	repo := os.Getenv("HARNESS_TEST_GITOPS_GCR_REPO")
	agentId := os.Getenv("HARNESS_TEST_GITOPS_AGENT_ID")
	gcrAccountKey := os.Getenv("HARNESS_TEST_GITOPS_GCR_ACCOUNT_KEY")
	repoName := id
	resourceName := "harness_platform_gitops_repository.test"
	accountId := os.Getenv("HARNESS_ACCOUNT_ID")
	awsAccessId := os.Getenv("HARNESS_TEST_AWS_ACCESS_KEY_ID")
	awsAccessKey := os.Getenv("HARNESS_TEST_AWS_SECRET_ACCESS_KEY")
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testHelmRepoOCIGCRServiceAccount(id, name, repo, repoName, agentId, accountId, gcrAccountKey, awsAccessId, awsAccessKey),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "repo.0.name", repoName),
				),
			},
		},
	})

}

func testHelmRepoOCIECRJwt(id string, name string, repo string, repoName string, agentId string, accountId string, awsAccessId, awsAccessKey string) string {
	return fmt.Sprintf(`

		resource "harness_platform_gitops_repository" "test" {
			identifier = "%[1]s"
			account_id = "%[6]s"
			agent_id = "%[5]s"
			repo {
					repo = "%[3]s"
        			name = "%[4]s"
                    type_ = "helm"
                    username = "AWS"
                    password = ""
        			enable_oci = true
        			connection_type = "HTTPS"
			}
            refresh_interval = "1m"
            gen_type = "AWS_ECR"
            ecr_gen {
				region = "us-west-1"
                jwt_auth {
					name = "service_name"
					namespace = "service_namespace"
                    audiences = ["service_audience"]
				}
    		}
			upsert = true
		}
		
		data "harness_platform_gitops_repository" "test" {
			depends_on = [harness_platform_gitops_repository.test]	
			identifier = harness_platform_gitops_repository.test.id
			account_id = "%[3]s"
			agent_id = harness_platform_gitops_repository.test.agent_id
		}
	`, id, name, repo, repoName, agentId, accountId, awsAccessId, awsAccessKey)
}

func testHelmRepoOCIECR(id string, name string, repo string, repoName string, agentId string, accountId string, awsAccessId, awsAccessKey string) string {
	return fmt.Sprintf(`
		resource "harness_platform_gitops_repository" "test" {
			identifier = "%[1]s"
			account_id = "%[6]s"
			agent_id = "%[5]s"
			repo {
					repo = "%[3]s"
        			name = "%[4]s"
                    type_ = "helm"
                    username = "AWS"
                    password = ""
        			enable_oci = true
        			connection_type = "HTTPS"
			}
            refresh_interval = "1m"
            gen_type = "AWS_ECR"
            ecr_gen {
        		region = "us-west-1"
            		aws_access_key_id = "%[7]s"
            		aws_secret_access_key = "%[8]s"
        		}
    		}
			upsert = true
		}
		
		data "harness_platform_gitops_repository" "test" {
			depends_on = [harness_platform_gitops_repository.test]	
			identifier = harness_platform_gitops_repository.test.id
			account_id = "%[3]s"
			agent_id = harness_platform_gitops_repository.test.agent_id
		}
	`, id, name, repo, repoName, agentId, accountId, awsAccessId, awsAccessKey)
}

func testHelmRepoOCIGCRWorkload(id string, name string, repo string, repoName string, agentId string, accountId string, awsAccessId, awsAccessKey string) string {
	return fmt.Sprintf(`

		resource "harness_platform_gitops_repository" "test" {
			identifier = "%[1]s"
			account_id = "%[6]s"
			agent_id = "%[5]s"
			repo {
					repo = "%[3]s"
        			name = "%[4]s"
                    type_ = "helm"
                    username = "AWS"
                    password = ""
        			enable_oci = true
        			connection_type = "HTTPS"
			}
            refresh_interval = "1m"
            gen_type = "GOOGLE_GCR"
            gcr_gen {
				project_id = "projectID"
				workload_identity  {
					cluster_location = "cluster_location"
					cluster_name = "cluster_name"
					cluster_project_id = "cluster_project_id"
					service_account_ref  {
						name = "name"
						namespace = "namespace"
					}
				}

    		}
			upsert = true
		}
		
		data "harness_platform_gitops_repository" "test" {
			depends_on = [harness_platform_gitops_repository.test]	
			identifier = harness_platform_gitops_repository.test.id
			account_id = "%[3]s"
			agent_id = harness_platform_gitops_repository.test.agent_id
		}
	`, id, name, repo, repoName, agentId, accountId, awsAccessId, awsAccessKey)
}

func testHelmRepoOCIGCRServiceAccount(id string, name string, repo string, repoName string, agentId string, accountId string, gcrAccountKey, awsAccessId, awsAccessKey string) string {
	return fmt.Sprintf(`
		resource "harness_platform_gitops_repository" "test" {
			identifier = "%[1]s"
			account_id = "%[6]s"
			agent_id = "%[5]s"
			repo {
					repo = "%[3]s"
        			name = "%[4]s"
                    type_ = "helm"
                    username = "AWS"
                    password = ""
        			enable_oci = true
        			connection_type = "HTTPS"
			}
            refresh_interval = "1m"
            gen_type = "GOOGLE_GCR"
            gcr_gen { 
                project_id = "projectID"
 	 			access_key = "%[7]s",
    		}
			upsert = true
		}
		
		data "harness_platform_gitops_repository" "test" {
			depends_on = [harness_platform_gitops_repository.test]	
			identifier = harness_platform_gitops_repository.test.id
			account_id = "%[3]s"
			agent_id = harness_platform_gitops_repository.test.agent_id
		}
	`, id, name, repo, repoName, agentId, accountId, gcrAccountKey, awsAccessId, awsAccessKey)
}
