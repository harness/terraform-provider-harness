package gnupg_test

import (
	"context"
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/antihax/optional"

	hh "github.com/harness/harness-go-sdk/harness/helpers"
	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccResourceGitopsGnupg(t *testing.T) {
	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	id = strings.ReplaceAll(id, "_", "")
	name := id
	agentId := os.Getenv("HARNESS_TEST_GITOPS_AGENT_ID")
	accountId := os.Getenv("HARNESS_ACCOUNT_ID")
	resourceName := "harness_platform_gitops_gnupg.test"
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccResourceGitopsGnupgDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceGitopsGnupg(id, accountId, name, agentId),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "account_id", accountId),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"org_id", "project_id"},
				ImportStateIdFunc:       acctest.GitopsAgentAccountLevelResourceImportStateIdFunc(resourceName),
			},
		},
	})

}

func testAccGetGnupg(resourceName string, state *terraform.State) (*nextgen.GpgkeysGnuPgPublicKey, error) {
	r := acctest.TestAccGetResource(resourceName, state)
	c, ctx := acctest.TestAccGetPlatformClientWithContext()
	ctx = context.WithValue(ctx, nextgen.ContextAccessToken, hh.EnvVars.BearerToken.Get())
	agentIdentifier := r.Primary.Attributes["agent_id"]

	resp, _, err := c.GnuPGPKeysApi.AgentGPGKeyServiceGet(ctx, agentIdentifier, r.Primary.ID, c.AccountId, &nextgen.GnuPGPKeysApiAgentGPGKeyServiceGetOpts{
		OrgIdentifier:     optional.NewString(r.Primary.Attributes["org_id"]),
		ProjectIdentifier: optional.NewString(r.Primary.Attributes["project_id"]),
	})

	if err != nil {
		return nil, err
	}

	if &resp == nil {
		return nil, nil
	}

	return &resp, nil
}

func testAccResourceGitopsGnupgDestroy(resourceName string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		gnupg, _ := testAccGetGnupg(resourceName, state)
		if gnupg != nil {
			return fmt.Errorf("Found gnupg")
		}

		return nil
	}

}

func testAccResourceGitopsGnupg(id string, accountId string, name string, agentId string) string {
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
			lifecycle {
				ignore_changes = [
					request.0.upsert,
				]
			}
		}
		`, id, accountId, name, agentId)

}
