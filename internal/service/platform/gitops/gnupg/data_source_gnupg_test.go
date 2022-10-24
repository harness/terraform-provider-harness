package gnupg_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceGitopsGnupg(t *testing.T) {

	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(6))
	name := id
	resourceName := "data.harness_platform_gitops_gnupg.test"

	agentId := "account.terraformagent1"
	accountId := os.Getenv("HARNESS_ACCOUNT_ID")

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceGitopsGnupg(id, accountId, name, agentId),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
				),
			},
		},
	})
}

func testAccDataSourceGitopsGnupg(id string, accountId string, name string, agentId string) string {
	return fmt.Sprintf(`
	resource "harness_platform_gitops_gnupg" "test" {
		account_id = "%[2]s"
		agent_id = "%[4]s"

		 request {
			upsert = true
			publickey {
				key_data = "-----BEGIN PGP PUBLIC KEY BLOCK-----\n\nmDMEY1Of9RYJKwYBBAHaRw8BAQdAjaTs6ENz1eyiDA62iKYM8aLFTLugqjyQQ0lK\nzqmB1bu0E3JhaiA8cmFqQGdtYWlsLmNvbT6ImQQTFgoAQRYhBOs34rbWDPJvTFXJ\n7xv7RmYkCDCqBQJjU5/1AhsDBQkDwmcABQsJCAcCAiICBhUKCQgLAgQWAgMBAh4H\nAheAAAoJEBv7RmYkCDCq7h8A/0BtunyvIOw+3xs7RlkulBcUvTIc7Xw9XEE74Akr\nle3oAQCweN3rPoGhwLAyrSj+VShhWeGA72zFU+aDR0RrkrXNB7g4BGNTn/USCisG\nAQQBl1UBBQEBB0DfRuVtj+ZXUZA7NyyeuuPWHmmiaPSYer4G24wTOhV4UQMBCAeI\nfgQYFgoAJhYhBOs34rbWDPJvTFXJ7xv7RmYkCDCqBQJjU5/1AhsMBQkDwmcAAAoJ\nEBv7RmYkCDCq6kkA/R712Ki3y88A6MiF1ajB8w9jPvGqoWaFbt1T0DdACQKWAP47\nIJj8ZykISu4EBnW+c+cYSYUceEXNiAMFL0VixHS6Dg==\n=X5Sv\n-----END PGP PUBLIC KEY BLOCK-----"
			}
		}
	}

	data "harness_platform_gitops_gnupg" "test" {
		account_id = "%[2]s"
	}
`, id, accountId, name, agentId)
}
