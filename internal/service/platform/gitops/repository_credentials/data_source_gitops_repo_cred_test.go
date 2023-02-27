package repository_credentials_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceGitOpsRepoCred(t *testing.T) {

	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(6))
	name := id
	agentId := os.Getenv("HARNESS_TEST_GITOPS_AGENT_ID")
	accountId := os.Getenv("HARNESS_ACCOUNT_ID")
	clusterName := id
	resourceName := "data.harness_platform_gitops_repo_cred.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceRepoCred(id, accountId, name, agentId, clusterName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "account_id", accountId),
					resource.TestCheckResourceAttr(resourceName, "agent_id", agentId),
				),
			},
		},
	})
}

func testAccDataSourceRepoCred(id string, accountId string, name string, agentId string, clusterName string) string {
	return fmt.Sprintf(`
		resource "harness_platform_gitops_repo_cred" "test" {
			identifier = "%[1]s"
			account_id = "%[2]s"
			agent_id = "%[4]s"
			creds {
				type = "git"
				url = "github.com"
				ssh_private_key = "-----BEGIN CERTIFICATE-----\nMIIFljCCA34CCQD9pxFhxWfoKTANBgkqhkiG9w0BAQsFADCBjDELMAkGA1UEBhMC\nSU4xFDASBgNVBAgMC01BSEFSQVNIVFJBMQ0wCwYDVQQHDARQVU5FMRAwDgYDVQQK\nDAdIQVJORVNTMQ8wDQYDVQQLDAZERVZPUFMxETAPBgNVBAMMCGJsYWhibGFoMSIw\nIAYJKoZIhvcNAQkBFhNyYmF2aXNrYXJAZ21haWwuY29tMB4XDTIyMTAyMTIxNDI0\nN1oXDTIzMTAyMTIxNDI0N1owgYwxCzAJBgNVBAYTAklOMRQwEgYDVQQIDAtNQUhB\nUkFTSFRSQTENMAsGA1UEBwwEUFVORTEQMA4GA1UECgwHSEFSTkVTUzEPMA0GA1UE\nCwwGREVWT1BTMREwDwYDVQQDDAhibGFoYmxhaDEiMCAGCSqGSIb3DQEJARYTcmJh\ndmlza2FyQGdtYWlsLmNvbTCCAiIwDQYJKoZIhvcNAQEBBQADggIPADCCAgoCggIB\nAOB+tVTRlSpgM2wZakgsaA7ISs6wVduapibUt1vMdKcfIUL2W5Hs+ZQYogdw2+Uf\neVGfYMFrKzkohdZnQkiqN3wRELPk6lsjl7tfQkpZS87/JIFf8ho+QeixjokvxkLi\n/E4RWZGyqAMSCk25g/5ZqS8cvIKhJaULy2oJicwfnPQFemjOtnOWkSKOYZMpcG1c\nKW4QBS0N7lzbnGNOV6/qBybGsC+IkxNFcAeCa0uq0RR5cUe8o6oVKs7Hl0QQPMOj\nzDhKIsMkYIDM/34Q/hn+APvySn/s7PZUUcHpOsReSpMGuT0r2xJeBDfWmeXrMXD3\nEjp9IqTjwNKNrxYGdES+PH15MJH0I8GioWdrM17OyJnM+k/hvcr6meFgnYueomOU\nlCGORI/Dc4lLnTfEvWUQX5xcZ4E1vE8ju2GK+DeP3tnOCIcfkgXUGKWE9B4S2LzD\nuNImaVYMovJ3uXN450qbcVa0saefrB07O0pnj92m/K2KPjlG4XchsuMufjrB+Nd8\nLxF/Qc2FETw4nvF44TLj2NikO7L/RpeWT0eSyVIeFVw45ZbkbD+hjTWYftuB8AyW\nOzhCgTa88Co5Do7yQY3tSfbmmX+AwHXD7sKKsFM7hoaXAx29tLefkN0XjHqJWQh2\nGmL8BQmmUV6uVd0dFuIAJkJTNQAspU+4ChfKRvnGU85nAgMBAAEwDQYJKoZIhvcN\nAQELBQADggIBAB6oMLK1SHULo2zv+pZOil4eroyTG1QE8CN0SXZEdDK/RnCEh+My\nDNWNILkLmLft2te3nqW7ucVVapoHXPNKOGDtvR9ATAbHCwPAxMLdOmBgDutR6tzA\ncG6IYze4UZt52tu2mFN6DJwiuh5dOqIJ+JX+KTKkCbKOfSNuxhRQYTVtttu6ISAl\n+jiGJ3bGRRKdiLJQb/UanlxUjVfUmpuB/Wdfg0DKJ1RJCumWnBEl1yEu2jORRvj+\nPLGm4XNSP7aXpK8AKJvfBZ2lb4hHkbkALbW+493/ZFQWZ+gCQCS/PKZY8rbEANFO\nzV1WsNk4YgRhNkbYE6Evx0s5uu2z/e6HyTkGnmuwkISHr2lPdkKkEb2km6lPn3bz\nFCi4YxW//c0Ix7Yht92sC/02R+2wG/I9V1ms/1a8p3Pho8pzhGE0gspCE0B2Dwzi\n/j7ntEBgM4HafpBF04DLF2R3dA2TJcVAjWvVKRHSnHo6iLjIoiziALhlyClqPYXg\nR9sbIYkAjh7JtSiYdLecxuMiNWSLTwTyvNzrNKCxoZXbHMx7WecESMFH3sPvBK9P\nIAnTehE1kSQKfkAVkvMeZBtZSbHu+NEOLFbliHpjYF0xfB7NNuuW/nbOWyp47SP6\nJO+/gie1XQokaVZ2PolGzfTzSWW9WSQ9JD24cQpEs23U0lZUvQk/qwrE\n-----END CERTIFICATE-----"
			}
			lifecycle {
				ignore_changes = [
					account_id,
					creds.0.ssh_private_key,
				]
			}
		}

		data "harness_platform_gitops_repo_cred" "test" {
			depends_on = [harness_platform_gitops_repo_cred.test]
			identifier = harness_platform_gitops_repo_cred.test.id
			account_id = "%[2]s"
			agent_id = "%[4]s"
		}
`, id, accountId, name, agentId, clusterName)
}
