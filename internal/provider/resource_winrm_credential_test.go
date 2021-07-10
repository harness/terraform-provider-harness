package provider

import (
	"fmt"
	"strings"

	"github.com/harness-io/harness-go-sdk/harness/api/graphql"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func init() {
	resource.AddTestSweepers("harness_winrm_credential", &resource.Sweeper{
		Name: "harness_winrm_credential",
		F:    testAccResourceWinRMCredentialSweep,
	})
}

func testAccResourceWinRMCredentialSweep(r string) error {
	c := testAccGetApiClientFromProvider()

	creds, err := c.Secrets().ListWinRMCredentials()
	fmt.Println(creds)
	fmt.Println(err)
	if err != nil {
		return fmt.Errorf("error retrieving WinRM credentials: %s", err)
	}

	for _, cred := range creds {
		if strings.HasPrefix(cred.Name, "Test") {
			if err = c.Secrets().DeleteSecret(cred.UUID, graphql.SecretTypes.WinRMCredential); err != nil {
				return err
			}
		}
	}

	return nil
}
