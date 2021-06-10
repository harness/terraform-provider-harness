package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/micahlmartin/terraform-provider-harness/internal/client"
	"github.com/micahlmartin/terraform-provider-harness/internal/utils"
	"github.com/stretchr/testify/require"
)

func TestAccResourceSSHCredential(t *testing.T) {

	name := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(12))
	resourceName := "harness_ssh_credential.test"
	updatedName := fmt.Sprintf("%s-updated", name)

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		CheckDestroy:      testAccSSHCredentialDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceSSHCredential(name, true, client.SSHAuthenticationTypes.SSHAuthentication),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "ssh_authentication.0.port", "22"),
					resource.TestCheckResourceAttr(resourceName, "ssh_authentication.0.username", "testuser"),
					resource.TestCheckResourceAttr(resourceName, "usage_scope.0.environment_filter_type", client.EnvironmentFilterTypes.NonProduction),
					resource.TestCheckResourceAttr(resourceName, "usage_scope.0.application_filter_type", client.ApplicationFilterTypes.All),
					testAccSShCredentialCreation(t, resourceName),
				),
			},
			{
				Config: testAccResourceSSHCredential(updatedName, false, client.SSHAuthenticationTypes.SSHAuthentication),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
					resource.TestCheckResourceAttr(resourceName, "ssh_authentication.0.port", "22"),
					resource.TestCheckResourceAttr(resourceName, "ssh_authentication.0.username", "testuser"),
					resource.TestCheckNoResourceAttr(resourceName, "usage_scope.0"),
					testAccSSHCredentialUpdate(t, resourceName),
				),
			},
		},
	})
}

func testAccGetSSHCredential(resourceName string, state *terraform.State) (*client.SSHCredential, error) {
	r := testAccGetResource(resourceName, state)
	c := testAccGetApiClientFromProvider()
	id := r.Primary.ID

	return c.Secrets().GetSSHCredentialById(id)
}

func testAccSShCredentialCreation(t *testing.T, resourceName string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		cred, err := testAccGetSSHCredential(resourceName, state)
		require.NoError(t, err)
		require.NotNil(t, cred)
		require.NotNil(t, cred.UsageScope)
		require.Len(t, cred.UsageScope.AppEnvScopes, 1)
		require.NotNil(t, cred.SSHAuthentication)
		require.Equal(t, cred.AuthenticationType, client.SSHAuthenticationTypes.SSHAuthentication)

		return nil
	}
}

func testAccSSHCredentialUpdate(t *testing.T, resourceName string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		cred, err := testAccGetSSHCredential(resourceName, state)
		require.NoError(t, err)
		require.NotNil(t, cred)
		require.Nil(t, cred.UsageScope)
		return nil
	}
}

func testAccSSHCredentialDestroy(resourceName string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		cred, _ := testAccGetSSHCredential(resourceName, state)
		if cred != nil {
			return fmt.Errorf("Found ssh credential: %s", cred.Id)
		}

		return nil
	}
}

func testAccSSHAuthenticationInlineSSH(keyFileId string) string {
	return fmt.Sprintf(`
		ssh_authentication {

			port = 22
			username = "testuser"
			
			inline_ssh {
				passphrase_secret_id = harness_encrypted_text.test.id
				ssh_key_file_id = "%[1]s"
			}

		}
`, keyFileId)

}

func testAccResourceSSHCredential(name string, withUsageScope bool, authType string) string {

	var (
		usageScope         string
		authenticationType string
	)

	if withUsageScope {
		usageScope = testAccDefaultUsageScope
	}

	switch authType {
	case client.SSHAuthenticationTypes.SSHAuthentication:
		authenticationType = testAccSSHAuthenticationInlineSSH(testAccSecretFileId)
	}

	return fmt.Sprintf(`
	resource "harness_encrypted_text" "test" {
		name = "%[1]s"
		value = "foo"
	}

	resource "harness_ssh_credential" "test" {
		name = "%[1]s"

		lifecycle {
			ignore_changes = [ssh_authentication]
		}

		%[2]s

		%[3]s
	}
	`, name, authenticationType, usageScope)
}
