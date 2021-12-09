package secrets_test

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/harness-io/harness-go-sdk/harness/api"
	"github.com/harness-io/harness-go-sdk/harness/cd/graphql"
	"github.com/harness-io/harness-go-sdk/harness/utils"
	"github.com/harness-io/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/stretchr/testify/require"
)

func TestAccResourceSSHCredential_SSHAuthentication(t *testing.T) {

	name := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(12))
	resourceName := "harness_ssh_credential.test"
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccSSHCredentialDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceSSHCredential(name, true, graphql.SSHAuthenticationTypes.SSHAuthentication),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "ssh_authentication.0.port", "22"),
					resource.TestCheckResourceAttr(resourceName, "ssh_authentication.0.username", "testuser"),
					resource.TestCheckResourceAttr(resourceName, "usage_scope.0.environment_filter_type", string(graphql.EnvironmentFilterTypes.NonProduction)),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccResourceSSHCredential_SSHAuthentication_DeleteUnderlyingResource(t *testing.T) {

	name := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(12))
	resourceName := "harness_ssh_credential.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceSSHCredential(name, true, graphql.SSHAuthenticationTypes.SSHAuthentication),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
				),
			},
			{
				PreConfig: func() {
					acctest.TestAccConfigureProvider()
					c := acctest.TestAccProvider.Meta().(*api.Client)

					secret, err := c.CDClient.SecretClient.GetSSHCredentialByName(name)
					require.NoError(t, err)
					require.NotNil(t, secret)

					err = c.CDClient.SecretClient.DeleteSecret(secret.Id, secret.SecretType)
					require.NoError(t, err)
				},
				Config:             testAccResourceSSHCredential(name, true, graphql.SSHAuthenticationTypes.SSHAuthentication),
				PlanOnly:           true,
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccResourceSSHCredential_KerberosAuthentication(t *testing.T) {

	name := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(12))
	resourceName := "harness_ssh_credential.test"
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccSSHCredentialDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceSSHCredential(name, true, graphql.SSHAuthenticationTypes.KerberosAuthentication),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "kerberos_authentication.0.port", "22"),
					resource.TestCheckResourceAttr(resourceName, "kerberos_authentication.0.principal", "testuser"),
					resource.TestCheckResourceAttr(resourceName, "kerberos_authentication.0.realm", "domain.com"),
					resource.TestCheckResourceAttr(resourceName, "usage_scope.0.environment_filter_type", string(graphql.EnvironmentFilterTypes.NonProduction)),
					testAccSShCredentialCreation(t, resourceName, graphql.SSHAuthenticationSchemes.Kerberos),
				),
			},
		},
	})
}

func TestAccResourceSSHCredential_BadAuthenticationMethods(t *testing.T) {

	name := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccResourceSSHCredentialBadAuthenticationMethod(name),
				ExpectError: regexp.MustCompile("`kerberos_authentication,ssh_authentication` can be specified"),
			},
		},
	})
}

func TestAccResourceSSHCredential_SSH_BadSSHAuthenticationTypes(t *testing.T) {

	name := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccResourceSSHCredentialBadSSHAuthenticationTypes(name),
				ExpectError: regexp.MustCompile("only one of"),
			},
		},
	})
}

func testAccGetSSHCredential(resourceName string, state *terraform.State) (*graphql.SSHCredential, error) {
	r := acctest.TestAccGetResource(resourceName, state)
	c := acctest.TestAccGetApiClientFromProvider()
	id := r.Primary.ID

	return c.CDClient.SecretClient.GetSSHCredentialById(id)
}

func testAccSShCredentialCreation(t *testing.T, resourceName string, authenticationScheme graphql.SSHAuthenticationScheme) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		cred, err := testAccGetSSHCredential(resourceName, state)
		require.NoError(t, err)
		require.NotNil(t, cred)
		require.NotNil(t, cred.UsageScope)
		require.Len(t, cred.UsageScope.AppEnvScopes, 1)

		if authenticationScheme == graphql.SSHAuthenticationSchemes.Kerberos {
			require.NotNil(t, cred.KerberosAuthentication)
		}

		if authenticationScheme == graphql.SSHAuthenticationSchemes.SSH {
			require.NotNil(t, cred.SSHAuthentication)
		}

		require.NotNil(t, cred.AuthenticationScheme, authenticationScheme)

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

func testAccSSHAuthenticationKerberosAuthentication() string {
	return `
		kerberos_authentication {
			port = 22
			principal = "testuser"
			realm = "domain.com"
		}
`
}

func testAccResourceSSHCredentialEncryptedText(name string) string {
	return fmt.Sprintf(`
		data "harness_secret_manager" "test" {
			default = true
		}

		resource "harness_encrypted_text" "test" {
			name 							= "%[1]s"
			value 						= "foo"
			secret_manager_id = data.harness_secret_manager.test.id
		}
	`, name)
}

func testAccResourceSSHCredentialBadSSHAuthenticationTypes(name string) string {

	var (
		encryptedTextResource = testAccResourceSSHCredentialEncryptedText(name)
	)

	return fmt.Sprintf(`
	resource "harness_ssh_credential" "test" {
		name = "%[1]s"

		ssh_authentication {

			port = 22
			username = "testuser"
			
			inline_ssh {
				passphrase_secret_id = harness_encrypted_text.test.id
				ssh_key_file_id = "%[1]s"
			}

			server_password {
				password_secret_id = harness_encrypted_text.test.id
			}

			ssh_key_file {
				path = "some/path"
				passphrase_secret_id = harness_encrypted_text.test.id
			}
		}

		lifecycle {
			ignore_changes = [
				"ssh_authentication",
				"kerberos_authentication"
			]
		}
	}

		%[2]s
	`, name, encryptedTextResource)
}

func testAccResourceSSHCredentialBadAuthenticationMethod(name string) string {

	var (
		encryptedTextResource  = testAccResourceSSHCredentialEncryptedText(name)
		sshAuthentication      = testAccSSHAuthenticationInlineSSH(acctest.TestAccSecretFileId)
		kerberosAuthentication = testAccSSHAuthenticationKerberosAuthentication()
	)

	return fmt.Sprintf(`
		resource "harness_ssh_credential" "test" {
			name = "%[1]s"

			%[2]s

			%[3]s

			lifecycle {
				ignore_changes = [
					"ssh_authentication",
					"kerberos_authentication"
				]
			}
		}

		%[4]s
	`, name, sshAuthentication, kerberosAuthentication, encryptedTextResource)
}

func testAccResourceSSHCredential(name string, withUsageScope bool, authType graphql.SSHAuthenticationType) string {

	var (
		usageScope            string
		authenticationType    string
		encryptedTextResource = testAccResourceSSHCredentialEncryptedText(name)
	)

	if withUsageScope {
		usageScope = acctest.TestAccDefaultUsageScope
	}

	switch authType {
	case graphql.SSHAuthenticationTypes.SSHAuthentication:
		authenticationType = testAccSSHAuthenticationInlineSSH(acctest.TestAccSecretFileId)
	case graphql.SSHAuthenticationTypes.KerberosAuthentication:
		authenticationType = testAccSSHAuthenticationKerberosAuthentication()
	}

	return fmt.Sprintf(`
		resource "harness_ssh_credential" "test" {
			name = "%[1]s"

			%[2]s

			%[3]s

			lifecycle {
				ignore_changes = [
					"ssh_authentication",
					"kerberos_authentication"
				]
			}
		}

		%[4]s
	`, name, authenticationType, usageScope, encryptedTextResource)
}
