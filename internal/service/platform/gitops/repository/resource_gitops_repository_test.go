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

func TestAccResourceGitopsRepository(t *testing.T) {
	// Project level
	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	id = strings.ReplaceAll(id, "_", "")
	name := id
	repo := "https://github.com/willycoll/argocd-example-apps.git"
	repoName := id
	repoNameUpdated := id + "_updated"
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
				Config: testAccResourceGitopsRepositoryProjectLevel(id, name, repo, repoNameUpdated, agentId, accountId),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "repo.0.name", repoNameUpdated),
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
	repoNameUpdated = id + "_updated"
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
				Config: testAccResourceGitopsRepositoryAccountLevel(id, name, repo, repoNameUpdated, agentId, accountId),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "repo.0.name", repoNameUpdated),
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
	repo := "806630305776.dkr.ecr.us-west-1.amazonaws.com"
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
	repo := "806630305776.dkr.ecr.us-west-1.amazonaws.com"
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
	repo := "us.gcr.io/qa-setup/harness-devops"
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
	repo := "us.gcr.io/qa-setup/harness-devops"
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
				Config: testHelmRepoOCIGCRServiceAccount(id, name, repo, repoName, agentId, accountId, awsAccessId, awsAccessKey),
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

func testHelmRepoOCIGCRServiceAccount(id string, name string, repo string, repoName string, agentId string, accountId string, awsAccessId, awsAccessKey string) string {
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
 	 			access_key = "{ \"type\": \"service_account\", \"project_id\": \"gitops-play-338013\", \"private_key_id\": \"5ef370c719dd12674be7be090083caaabd85eca6\", \"private_key\": \"-----BEGIN PRIVATE KEY-----\nMIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQDQXwd7zfou3NDB\nP9TwgQMWyHYpWRJENC7RTpCaQ1KqBElEHPUBzPpv2L/ayn50qpsmtxYTWdQwR7OL\n8qK+8hrYfSvwSTph/PMOtesrZpWrloIG2virgw1y1zeUDP3DKQDemerVrpve189h\nxxTLTriLS0JMsDniE0HHzeeMtxTbi0NcojDufSq3gTZ4TJ80jPtdIURYAwfW4jHA\nOW1nq8EY3VjOkUDHZ++xtHen9FT4OzpcTSkGyovaYZfLUPFDrtZqmcdS5IjFmTpE\nHfIibB+roT0jynhBBPwuBUWtbJEjXg8Gw2hRquSChCMxKkD9PjttyQMjmlhVrLqR\nG1p4vBSDAgMBAAECggEAS8nUsf4oOjVMpI1wCQ4Troy5Fa71CuOkB7M4uzMzdO1c\nLK8PmlkQ2e+PUKgIOKz5A6riF6W7nNfngUZ+VU8/3nAgtCQeXReg3D/kyoNkeuWi\nY5Xvjop7MMMAzxOulPZr/4siNBhvTy1Vm63KbWwziU6VTclnNEhmy6KjzrWkm3ky\ndcYsnGr5eWbQvmSAE38EeSMw6+OoiEmPYk/hoRg85lVouQ72d4FHaNjU3NNCR1Y0\nUEdj8r9K19nZ7ICTxJ6AZp5rc3Z2rG8MrBY2UEhGhFLZmad8hUtyo5Ol93kOwSHW\n+baGXRasBhWN3uZ3PYCh3NzEJeHEVX0HT4FYUF8NQQKBgQD1OMOVAF9XWESPTuyw\nz5/4S+kRWdplXkP9dEacQ21hIwgX5PFWZFHD1VhwSVniIPWao2Fa/QMZK4np+g9d\nnEkgPlFatPsLT0q3/QT59oHEIAIEorOdz0RXxAkU9xo0RfXQWsXFDgcwcKd9yeet\nxbiQO/LscNomM/CcWW63O1l5IwKBgQDZh596QU9Y3z07OfF9pl86X+QIQlEY0nxr\nx2L+JspVXWnIHoVGlODOoP/EmCfS23oJdZZC7TWLSS9GDCsTC4UPcHW5I0cFFT69\n9M0ZvP2P6oCf2Jg7QOX8DIamcv6wI0MQKdUFDW+wtf01hiS/6lwEQL8xFBhw2+xq\njIKdkoOdIQKBgQCN6Z7OURvb6Xor0UoK/O0f/ZZQ80X/mfEQ8cSXVDItn99kLJs6\nGu5yvbnjqZ95zQc1yc1iob+0Rk0W+h8AVpy/KzFbpBcQsX+VQLkri2wHu1pPonT+\nI9/yRsHWvzYMAFzEinOfmYGxl9BmbH1GRIGN/xOTn6+voilh4iO/qHocLwKBgCNy\n7pJFwmCBQME+GBSZ4DrrFYYjCIQ7CPunaoJwX9i5eFucXau650fFBOlMwnCiQ6j2\n+J2/elJQgtuvb/WSkwSJFyYskY5KgAcEtcfT/J5PYNarvWMqmFAS2n6Vjtu1Y2Bm\n8Mf6AJGTlsf6LFL6JjSrOH0PAUyjCkvyyfZTwg8BAoGBAOrOYrOC6zigjC5Kmve3\nORnw318hPOV5oo7a7NpztSwwY1/7xZuOJXLaflZXnYCO1BXY+PosshI1ckcrv6PT\niEr+SQ+mbaaxcFxtJUP6Y4GBI4ayeHnmgkfuVwPEd//rnPD6YA5RRFF/GRI619Hu\nAt9fAayERhb7ipMxMQw6wpbF\n-----END PRIVATE KEY-----\n\", \"client_email\": \"570436793911-compute@developer.gserviceaccount.com\", \"client_id\": \"114979827506890161940\", \"auth_uri\": \"https://accounts.google.com/o/oauth2/auth\", \"token_uri\": \"https://oauth2.googleapis.com/token\", \"auth_provider_x509_cert_url\": \"https://www.googleapis.com/oauth2/v1/certs\", \"client_x509_cert_url\": \"https://www.googleapis.com/robot/v1/metadata/x509/570436793911-compute@eveloper.gserviceaccount.com\", \"universe_domain\": \"googleapis.com\" }"
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
