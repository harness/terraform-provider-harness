package repository_credentials_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	hh "github.com/harness/harness-go-sdk/harness/helpers"
	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccResourceGitopsRepoCred(t *testing.T) {
	// ACCOUNT LEVEL
	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	agentId := os.Getenv("HARNESS_TEST_GITOPS_AGENT_ID")
	accountId := os.Getenv("HARNESS_ACCOUNT_ID")
	resourceName := "harness_platform_gitops_repo_cred.test"
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccResourceGitopsRepoCredDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceGitopsRepoCredAccountLevel(id, accountId, agentId),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"creds.0.ssh_private_key"},
				ImportStateIdFunc:       acctest.GitopsAgentAccountLevelResourceImportStateIdFunc(resourceName),
			},
		},
	})

	// PROJECT LEVEL
	//id = fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	//name := id
	//resourceName = "harness_platform_gitops_repo_cred.test"
	//resource.UnitTest(t, resource.TestCase{
	//	PreCheck:          func() { acctest.TestAccPreCheck(t) },
	//	ProviderFactories: acctest.ProviderFactories,
	//	CheckDestroy:      testAccResourceGitopsRepoCredDestroy(resourceName),
	//	Steps: []resource.TestStep{
	//		{
	//			Config: testAccResourceGitopsRepoCredProjectLevel(id, name, accountId, agentId),
	//			Check: resource.ComposeTestCheckFunc(
	//				resource.TestCheckResourceAttr(resourceName, "id", id),
	//			),
	//		},
	//		{
	//			ResourceName:            resourceName,
	//			ImportState:             true,
	//			ImportStateVerify:       true,
	//			ImportStateVerifyIgnore: []string{"creds.0.ssh_private_key"},
	//			ImportStateIdFunc:       acctest.GitopsAgentProjectLevelResourceImportStateIdFunc(resourceName),
	//		},
	//	},
	//})
}

func testAccGetRepoCred(resourceName string, state *terraform.State) (*nextgen.Servicev1RepositoryCredentials, error) {
	r := acctest.TestAccGetResource(resourceName, state)
	c, ctx := acctest.TestAccGetPlatformClientWithContext()
	ctx = context.WithValue(ctx, nextgen.ContextAccessToken, hh.EnvVars.BearerToken.Get())
	agentIdentifier := r.Primary.Attributes["agent_id"]
	identifier := r.Primary.Attributes["identifier"]

	resp, _, err := c.RepositoryCredentialsApi.AgentRepositoryCredentialsServiceGetRepositoryCredentials(ctx, agentIdentifier, identifier, c.AccountId, &nextgen.RepositoryCredentialsApiAgentRepositoryCredentialsServiceGetRepositoryCredentialsOpts{})

	if err != nil {
		return nil, err
	}

	return &resp, nil
}

func testAccResourceGitopsRepoCredDestroy(resourceName string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		repoCred, _ := testAccGetRepoCred(resourceName, state)
		if repoCred != nil {
			return fmt.Errorf("Found repo cred")
		}
		return nil
	}
}

func testAccResourceGitopsRepoCredAccountLevel(id string, accountId string, agentId string) string {
	return fmt.Sprintf(`
		resource "harness_platform_gitops_repo_cred" "test" {
			identifier = "%[1]s"
			account_id = "%[2]s"
			agent_id = "%[3]s"
			creds {
				type = "git"
				url = "github.us"
				ssh_private_key = "-----BEGIN CERTIFICATE-----\nMIIFljCCA34CCQD9pxFhxWfoKTANBgkqhkiG9w0BAQsFADCBjDELMAkGA1UEBhMC\nSU4xFDASBgNVBAgMC01BSEFSQVNIVFJBMQ0wCwYDVQQHDARQVU5FMRAwDgYDVQQK\nDAdIQVJORVNTMQ8wDQYDVQQLDAZERVZPUFMxETAPBgNVBAMMCGJsYWhibGFoMSIw\nIAYJKoZIhvcNAQkBFhNyYmF2aXNrYXJAZ21haWwuY29tMB4XDTIyMTAyMTIxNDI0\nN1oXDTIzMTAyMTIxNDI0N1owgYwxCzAJBgNVBAYTAklOMRQwEgYDVQQIDAtNQUhB\nUkFTSFRSQTENMAsGA1UEBwwEUFVORTEQMA4GA1UECgwHSEFSTkVTUzEPMA0GA1UE\nCwwGREVWT1BTMREwDwYDVQQDDAhibGFoYmxhaDEiMCAGCSqGSIb3DQEJARYTcmJh\ndmlza2FyQGdtYWlsLmNvbTCCAiIwDQYJKoZIhvcNAQEBBQADggIPADCCAgoCggIB\nAOB+tVTRlSpgM2wZakgsaA7ISs6wVduapibUt1vMdKcfIUL2W5Hs+ZQYogdw2+Uf\neVGfYMFrKzkohdZnQkiqN3wRELPk6lsjl7tfQkpZS87/JIFf8ho+QeixjokvxkLi\n/E4RWZGyqAMSCk25g/5ZqS8cvIKhJaULy2oJicwfnPQFemjOtnOWkSKOYZMpcG1c\nKW4QBS0N7lzbnGNOV6/qBybGsC+IkxNFcAeCa0uq0RR5cUe8o6oVKs7Hl0QQPMOj\nzDhKIsMkYIDM/34Q/hn+APvySn/s7PZUUcHpOsReSpMGuT0r2xJeBDfWmeXrMXD3\nEjp9IqTjwNKNrxYGdES+PH15MJH0I8GioWdrM17OyJnM+k/hvcr6meFgnYueomOU\nlCGORI/Dc4lLnTfEvWUQX5xcZ4E1vE8ju2GK+DeP3tnOCIcfkgXUGKWE9B4S2LzD\nuNImaVYMovJ3uXN450qbcVa0saefrB07O0pnj92m/K2KPjlG4XchsuMufjrB+Nd8\nLxF/Qc2FETw4nvF44TLj2NikO7L/RpeWT0eSyVIeFVw45ZbkbD+hjTWYftuB8AyW\nOzhCgTa88Co5Do7yQY3tSfbmmX+AwHXD7sKKsFM7hoaXAx29tLefkN0XjHqJWQh2\nGmL8BQmmUV6uVd0dFuIAJkJTNQAspU+4ChfKRvnGU85nAgMBAAEwDQYJKoZIhvcN\nAQELBQADggIBAB6oMLK1SHULo2zv+pZOil4eroyTG1QE8CN0SXZEdDK/RnCEh+My\nDNWNILkLmLft2te3nqW7ucVVapoHXPNKOGDtvR9ATAbHCwPAxMLdOmBgDutR6tzA\ncG6IYze4UZt52tu2mFN6DJwiuh5dOqIJ+JX+KTKkCbKOfSNuxhRQYTVtttu6ISAl\n+jiGJ3bGRRKdiLJQb/UanlxUjVfUmpuB/Wdfg0DKJ1RJCumWnBEl1yEu2jORRvj+\nPLGm4XNSP7aXpK8AKJvfBZ2lb4hHkbkALbW+493/ZFQWZ+gCQCS/PKZY8rbEANFO\nzV1WsNk4YgRhNkbYE6Evx0s5uu2z/e6HyTkGnmuwkISHr2lPdkKkEb2km6lPn3bz\nFCi4YxW//c0Ix7Yht92sC/02R+2wG/I9V1ms/1a8p3Pho8pzhGE0gspCE0B2Dwzi\n/j7ntEBgM4HafpBF04DLF2R3dA2TJcVAjWvVKRHSnHo6iLjIoiziALhlyClqPYXg\nR9sbIYkAjh7JtSiYdLecxuMiNWSLTwTyvNzrNKCxoZXbHMx7WecESMFH3sPvBK9P\nIAnTehE1kSQKfkAVkvMeZBtZSbHu+NEOLFbliHpjYF0xfB7NNuuW/nbOWyp47SP6\nJO+/gie1XQokaVZ2PolGzfTzSWW9WSQ9JD24cQpEs23U0lZUvQk/qwrE\n-----END CERTIFICATE-----"
			}
			lifecycle {
				ignore_changes = [
					creds.0.ssh_private_key,
				]
			}
		}
		`, id, accountId, agentId)
}

func testAccResourceGitopsRepoCredProjectLevel(id string, accountId string, name string, agentId string) string {
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
		resource "harness_platform_gitops_repo_cred" "test" {
			identifier = "%[1]s"
			account_id = "%[2]s"
			agent_id = "%[4]s"
			project_id = harness_platform_project.test.id
			org_id = harness_platform_organization.test.id
			creds {
				type = "git"
				url = "github.net"
				ssh_private_key = "-----BEGIN CERTIFICATE-----\nMIIFljCCA34CCQD9pxFhxWfoKTANBgkqhkiG9w0BAQsFADCBjDELMAkGA1UEBhMC\nSU4xFDASBgNVBAgMC01BSEFSQVNIVFJBMQ0wCwYDVQQHDARQVU5FMRAwDgYDVQQK\nDAdIQVJORVNTMQ8wDQYDVQQLDAZERVZPUFMxETAPBgNVBAMMCGJsYWhibGFoMSIw\nIAYJKoZIhvcNAQkBFhNyYmF2aXNrYXJAZ21haWwuY29tMB4XDTIyMTAyMTIxNDI0\nN1oXDTIzMTAyMTIxNDI0N1owgYwxCzAJBgNVBAYTAklOMRQwEgYDVQQIDAtNQUhB\nUkFTSFRSQTENMAsGA1UEBwwEUFVORTEQMA4GA1UECgwHSEFSTkVTUzEPMA0GA1UE\nCwwGREVWT1BTMREwDwYDVQQDDAhibGFoYmxhaDEiMCAGCSqGSIb3DQEJARYTcmJh\ndmlza2FyQGdtYWlsLmNvbTCCAiIwDQYJKoZIhvcNAQEBBQADggIPADCCAgoCggIB\nAOB+tVTRlSpgM2wZakgsaA7ISs6wVduapibUt1vMdKcfIUL2W5Hs+ZQYogdw2+Uf\neVGfYMFrKzkohdZnQkiqN3wRELPk6lsjl7tfQkpZS87/JIFf8ho+QeixjokvxkLi\n/E4RWZGyqAMSCk25g/5ZqS8cvIKhJaULy2oJicwfnPQFemjOtnOWkSKOYZMpcG1c\nKW4QBS0N7lzbnGNOV6/qBybGsC+IkxNFcAeCa0uq0RR5cUe8o6oVKs7Hl0QQPMOj\nzDhKIsMkYIDM/34Q/hn+APvySn/s7PZUUcHpOsReSpMGuT0r2xJeBDfWmeXrMXD3\nEjp9IqTjwNKNrxYGdES+PH15MJH0I8GioWdrM17OyJnM+k/hvcr6meFgnYueomOU\nlCGORI/Dc4lLnTfEvWUQX5xcZ4E1vE8ju2GK+DeP3tnOCIcfkgXUGKWE9B4S2LzD\nuNImaVYMovJ3uXN450qbcVa0saefrB07O0pnj92m/K2KPjlG4XchsuMufjrB+Nd8\nLxF/Qc2FETw4nvF44TLj2NikO7L/RpeWT0eSyVIeFVw45ZbkbD+hjTWYftuB8AyW\nOzhCgTa88Co5Do7yQY3tSfbmmX+AwHXD7sKKsFM7hoaXAx29tLefkN0XjHqJWQh2\nGmL8BQmmUV6uVd0dFuIAJkJTNQAspU+4ChfKRvnGU85nAgMBAAEwDQYJKoZIhvcN\nAQELBQADggIBAB6oMLK1SHULo2zv+pZOil4eroyTG1QE8CN0SXZEdDK/RnCEh+My\nDNWNILkLmLft2te3nqW7ucVVapoHXPNKOGDtvR9ATAbHCwPAxMLdOmBgDutR6tzA\ncG6IYze4UZt52tu2mFN6DJwiuh5dOqIJ+JX+KTKkCbKOfSNuxhRQYTVtttu6ISAl\n+jiGJ3bGRRKdiLJQb/UanlxUjVfUmpuB/Wdfg0DKJ1RJCumWnBEl1yEu2jORRvj+\nPLGm4XNSP7aXpK8AKJvfBZ2lb4hHkbkALbW+493/ZFQWZ+gCQCS/PKZY8rbEANFO\nzV1WsNk4YgRhNkbYE6Evx0s5uu2z/e6HyTkGnmuwkISHr2lPdkKkEb2km6lPn3bz\nFCi4YxW//c0Ix7Yht92sC/02R+2wG/I9V1ms/1a8p3Pho8pzhGE0gspCE0B2Dwzi\n/j7ntEBgM4HafpBF04DLF2R3dA2TJcVAjWvVKRHSnHo6iLjIoiziALhlyClqPYXg\nR9sbIYkAjh7JtSiYdLecxuMiNWSLTwTyvNzrNKCxoZXbHMx7WecESMFH3sPvBK9P\nIAnTehE1kSQKfkAVkvMeZBtZSbHu+NEOLFbliHpjYF0xfB7NNuuW/nbOWyp47SP6\nJO+/gie1XQokaVZ2PolGzfTzSWW9WSQ9JD24cQpEs23U0lZUvQk/qwrE\n-----END CERTIFICATE-----"
			}
			lifecycle {
				ignore_changes = [
					account_id,
					creds.0.ssh_private_key,
				]
			}
		}
		`, id, accountId, name, agentId)
}
