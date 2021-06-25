package provider

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/harness-io/harness-go-sdk/harness/api/graphql"
	"github.com/harness-io/harness-go-sdk/harness/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/stretchr/testify/require"
)

func TestAccResourceSSHCredential_SSHAuthentication(t *testing.T) {

	name := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(12))
	resourceName := "harness_ssh_credential.test"
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		CheckDestroy:      testAccSSHCredentialDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceSSHCredential(name, true, graphql.SSHAuthenticationTypes.SSHAuthentication),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "ssh_authentication.0.port", "22"),
					resource.TestCheckResourceAttr(resourceName, "ssh_authentication.0.username", "testuser"),
					resource.TestCheckResourceAttr(resourceName, "usage_scope.0.environment_filter_type", graphql.EnvironmentFilterTypes.NonProduction),
					resource.TestCheckResourceAttr(resourceName, "usage_scope.0.application_filter_type", graphql.ApplicationFilterTypes.All),
					testAccSShCredentialCreation(t, resourceName, graphql.SSHAuthenticationSchemes.SSH),
				),
			},
		},
	})
}

func TestAccResourceSSHCredential_KerberosAuthentication(t *testing.T) {

	name := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(12))
	resourceName := "harness_ssh_credential.test"
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		CheckDestroy:      testAccSSHCredentialDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceSSHCredential(name, true, graphql.SSHAuthenticationTypes.KerberosAuthentication),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "kerberos_authentication.0.port", "22"),
					resource.TestCheckResourceAttr(resourceName, "kerberos_authentication.0.principal", "testuser"),
					resource.TestCheckResourceAttr(resourceName, "kerberos_authentication.0.realm", "domain.com"),
					resource.TestCheckResourceAttr(resourceName, "usage_scope.0.environment_filter_type", graphql.EnvironmentFilterTypes.NonProduction),
					resource.TestCheckResourceAttr(resourceName, "usage_scope.0.application_filter_type", graphql.ApplicationFilterTypes.All),
					testAccSShCredentialCreation(t, resourceName, graphql.SSHAuthenticationSchemes.Kerberos),
				),
			},
		},
	})
}

func TestAccResourceSSHCredential_Force_Recreate(t *testing.T) {

	name := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(12))
	resourceName := "harness_ssh_credential.test"
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		CheckDestroy:      testAccSSHCredentialDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceSSHCredential(name, true, graphql.SSHAuthenticationTypes.SSHAuthentication),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "ssh_authentication.0.port", "22"),
					resource.TestCheckResourceAttr(resourceName, "ssh_authentication.0.username", "testuser"),
					resource.TestCheckResourceAttr(resourceName, "usage_scope.0.environment_filter_type", graphql.EnvironmentFilterTypes.NonProduction),
					resource.TestCheckResourceAttr(resourceName, "usage_scope.0.application_filter_type", graphql.ApplicationFilterTypes.All),
					testAccSShCredentialCreation(t, resourceName, graphql.SSHAuthenticationSchemes.SSH),
				),
			},
			{
				Config: testAccResourceSSHCredential(name, true, graphql.SSHAuthenticationTypes.KerberosAuthentication),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "kerberos_authentication.0.port", "22"),
					resource.TestCheckResourceAttr(resourceName, "kerberos_authentication.0.principal", "testuser"),
					resource.TestCheckResourceAttr(resourceName, "kerberos_authentication.0.realm", "domain.com"),
					resource.TestCheckResourceAttr(resourceName, "usage_scope.0.environment_filter_type", graphql.EnvironmentFilterTypes.NonProduction),
					resource.TestCheckResourceAttr(resourceName, "usage_scope.0.application_filter_type", graphql.ApplicationFilterTypes.All),
					testAccSShCredentialCreation(t, resourceName, graphql.SSHAuthenticationSchemes.Kerberos),
				),
			},
		},
	})
}

func TestAccResourceSSHCredential_BadAuthenticationMethods(t *testing.T) {

	name := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
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
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccResourceSSHCredentialBadSSHAuthenticationTypes(name),
				ExpectError: regexp.MustCompile("only one of"),
			},
		},
	})
}

func testAccGetSSHCredential(resourceName string, state *terraform.State) (*graphql.SSHCredential, error) {
	r := testAccGetResource(resourceName, state)
	c := testAccGetApiClientFromProvider()
	id := r.Primary.ID

	return c.Secrets().GetSSHCredentialById(id)
}

func testAccSShCredentialCreation(t *testing.T, resourceName string, authenticationScheme string) resource.TestCheckFunc {
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
		resource "harness_encrypted_text" "test" {
			name = "%[1]s"
			value = "foo"

			lifecycle {
				ignore_changes = [secret_manager_id]
			}
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
	}

		%[2]s
	`, name, encryptedTextResource)
}

func testAccResourceSSHCredentialBadAuthenticationMethod(name string) string {

	var (
		encryptedTextResource  = testAccResourceSSHCredentialEncryptedText(name)
		sshAuthentication      = testAccSSHAuthenticationInlineSSH(testAccSecretFileId)
		kerberosAuthentication = testAccSSHAuthenticationKerberosAuthentication()
	)

	return fmt.Sprintf(`
		resource "harness_ssh_credential" "test" {
			name = "%[1]s"

			%[2]s

			%[3]s
		}

		%[4]s
	`, name, sshAuthentication, kerberosAuthentication, encryptedTextResource)
}

func testAccResourceSSHCredential(name string, withUsageScope bool, authType string) string {

	var (
		usageScope            string
		authenticationType    string
		encryptedTextResource = testAccResourceSSHCredentialEncryptedText(name)
	)

	if withUsageScope {
		usageScope = testAccDefaultUsageScope
	}

	switch authType {
	case graphql.SSHAuthenticationTypes.SSHAuthentication:
		authenticationType = testAccSSHAuthenticationInlineSSH(testAccSecretFileId)
	case graphql.SSHAuthenticationTypes.KerberosAuthentication:
		authenticationType = testAccSSHAuthenticationKerberosAuthentication()
	}

	return fmt.Sprintf(`
		resource "harness_ssh_credential" "test" {
			name = "%[1]s"

			%[2]s

			%[3]s
		}

		%[4]s
	`, name, authenticationType, usageScope, encryptedTextResource)
}
