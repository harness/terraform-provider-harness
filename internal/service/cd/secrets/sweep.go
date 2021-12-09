package secrets

import (
	"fmt"
	"strings"

	"github.com/harness-io/harness-go-sdk/harness/cd/graphql"
	"github.com/harness-io/terraform-provider-harness/internal/sweep"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func init() {
	resource.AddTestSweepers("harness_encrypted_text", &resource.Sweeper{
		Name: "harness_encrypted_text",
		F:    testSweepHarnessEncryptedText,
		Dependencies: []string{
			"harness_cloudprovider_aws",
			"harness_git_connector",
			"harness_ssh_credential",
			"harness_winrm_credential",
		},
	})

	resource.AddTestSweepers("harness_ssh_credential", &resource.Sweeper{
		Name: "harness_ssh_credential",
		F:    testAccResourceSSHCredentialSweep,
	})
}

func testSweepHarnessEncryptedText(r string) error {
	c := sweep.SweeperClient

	limit := 500
	offset := 0
	hasMore := true

	for hasMore {

		secrets, _, err := c.CDClient.SecretClient.ListEncryptedTextSecrets(limit, offset)

		if err != nil {
			return err
		}

		for _, secret := range secrets {
			if strings.HasPrefix(secret.Name, "Test") {
				if err = c.CDClient.SecretClient.DeleteSecret(secret.UUID, graphql.SecretTypes.EncryptedText); err != nil {
					return err
				}
			}
		}

		hasMore = len(secrets) == limit
	}

	return nil
}

func testAccResourceSSHCredentialSweep(r string) error {
	c := sweep.SweeperClient

	creds, err := c.CDClient.SecretClient.ListSSHCredentials()
	if err != nil {
		return fmt.Errorf("error retrieving SSH credentials: %s", err)
	}

	for _, cred := range creds {
		if strings.HasPrefix(cred.Name, "Test") {
			if err = c.CDClient.SecretClient.DeleteSecret(cred.UUID, graphql.SecretTypes.SSHCredential); err != nil {
				return err
			}
		}
	}

	return nil
}
