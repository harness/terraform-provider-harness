package repository_certificates_test

// import (
// 	"context"
// 	"fmt"
// 	"os"
// 	"testing"

// 	"github.com/antihax/optional"
// 	hh "github.com/harness/harness-go-sdk/harness/helpers"
// 	"github.com/harness/harness-go-sdk/harness/nextgen"
// 	"github.com/harness/harness-go-sdk/harness/utils"
// 	"github.com/harness/terraform-provider-harness/internal/acctest"
// 	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
// 	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
// )

// func TestAccResourceGitopsRepoCert(t *testing.T) {
// 	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
// 	name := id
// 	agentId := "account.terraformagent1"
// 	accountId := os.Getenv("HARNESS_ACCOUNT_ID")
// 	clusterName := id
// 	resourceName := "harness_platform_gitops_repository_certificates.test"
// 	resource.UnitTest(t, resource.TestCase{
// 		PreCheck:          func() { acctest.TestAccPreCheck(t) },
// 		ProviderFactories: acctest.ProviderFactories,
// 		// CheckDestroy:      testAccResourceGitopsRepositoriesCertificatesDestroy(resourceName),
// 		Steps: []resource.TestStep{
// 			{
// 				Config: testAccResourceGitopsRepositoriesCertificates(id, accountId, name, agentId, clusterName),
// 				Check: resource.ComposeTestCheckFunc(
// 					resource.TestCheckResourceAttr(resourceName, "id", id),
// 					resource.TestCheckResourceAttr(resourceName, "identifier", id),
// 				),
// 			},
// 			{
// 				Config: testAccResourceGitopsRepositoriesCertificates(id, accountId, name, agentId, clusterName),
// 				Check: resource.ComposeTestCheckFunc(
// 					resource.TestCheckResourceAttr(resourceName, "id", id),
// 					resource.TestCheckResourceAttr(resourceName, "identifier", id),
// 				),
// 			},
// 			{
// 				ResourceName:      resourceName,
// 				ImportState:       true,
// 				ImportStateVerify: true,
// 				ImportStateIdFunc: acctest.GitopsAgentResourceImportStateIdFunc(resourceName),
// 			},
// 		},
// 	})

// }

// func testAccGetRepoCert(resourceName string, state *terraform.State) (*nextgen.CertificatesRepositoryCertificateList, error) {
// 	r := acctest.TestAccGetResource(resourceName, state)
// 	c, ctx := acctest.TestAccGetPlatformClientWithContext()
// 	ctx = context.WithValue(ctx, nextgen.ContextAccessToken, hh.EnvVars.BearerToken.Get())
// 	agentIdentifier := r.Primary.Attributes["agent_id"]

// 	resp, _, err := c.RepositoryCertificatesApi.AgentCertificateServiceList(ctx, agentIdentifier, c.AccountId, &nextgen.RepositoryCertificatesApiAgentCertificateServiceListOpts{
// 		OrgIdentifier:     optional.NewString(r.Primary.Attributes["org_id"]),
// 		ProjectIdentifier: optional.NewString(r.Primary.Attributes["project_id"]),
// 	})

// 	if err != nil {
// 		return nil, err
// 	}

// 	return &resp, nil
// }

// func testAccResourceGitopsRepositoriesCertificatesDestroy(resourceName string) resource.TestCheckFunc {
// 	return func(state *terraform.State) error {
// 		repoCert, _ := testAccGetRepoCert(resourceName, state)
// 		if repoCert != nil {
// 			return fmt.Errorf("Found repo cert")
// 		}

// 		return nil
// 	}

// }

// func testAccResourceGitopsRepositoriesCertificates(id string, accountId string, name string, agentId string, clusterName string) string {
// 	return fmt.Sprintf(`
// 		resource "harness_platform_organization" "test" {
// 			identifier = "%[1]s"
// 			name = "%[3]s"
// 		}

// 		resource "harness_platform_project" "test" {
// 			identifier = "%[1]s"
// 			name = "%[3]s"
// 			org_id = harness_platform_organization.test.id
// 		}
// 		resource "harness_platform_gitops_repository_certificates" "test" {
// 			account_id = "%[2]s"
// 			project_id = harness_platform_project.test.id
// 			org_id = harness_platform_organization.test.id
// 			agent_id = "%[4]s"

//  			request {
// 				upsert = true
// 				certificates {

// 				}
// 			}
// 		}
// 		`, id, accountId, name, agentId, clusterName)

// }
