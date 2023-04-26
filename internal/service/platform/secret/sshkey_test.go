package secret_test

import (
	"fmt"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccSecretSSHKey_kerberos_keyFilePath(t *testing.T) {

	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	name := id
	updatedName := fmt.Sprintf("%s_updated", name)
	resourceName := "harness_platform_secret_sshkey.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccSecretDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceSecret_sshkey_kerberos_keyFilePath(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "kerberos.0.principal", "principal"),
					resource.TestCheckResourceAttr(resourceName, "kerberos.0.realm", "realm"),
					resource.TestCheckResourceAttr(resourceName, "kerberos.0.tgt_generation_method", "KeyTabFilePath"),
					resource.TestCheckResourceAttr(resourceName, "kerberos.0.tgt_key_tab_file_path_spec.0.key_path", "key_path"),
				),
			},
			{
				Config: testAccResourceSecret_sshkey_kerberos_keyFilePath(id, updatedName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "kerberos.0.principal", "principal"),
					resource.TestCheckResourceAttr(resourceName, "kerberos.0.realm", "realm"),
					resource.TestCheckResourceAttr(resourceName, "kerberos.0.tgt_generation_method", "KeyTabFilePath"),
					resource.TestCheckResourceAttr(resourceName, "kerberos.0.tgt_key_tab_file_path_spec.0.key_path", "key_path"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"port",
				},
			},
		},
	})
}

func TestAccSecretSSHKey_kerberos_password(t *testing.T) {

	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	name := id
	updatedName := fmt.Sprintf("%s_updated", name)
	resourceName := "harness_platform_secret_sshkey.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccSecretDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceSecret_sshkey_kerberos_password(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "kerberos.0.principal", "principal"),
					resource.TestCheckResourceAttr(resourceName, "kerberos.0.realm", "realm"),
					resource.TestCheckResourceAttr(resourceName, "kerberos.0.tgt_generation_method", "Password"),
				),
			},
			{
				Config: testAccResourceSecret_sshkey_kerberos_password(id, updatedName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "kerberos.0.principal", "principal"),
					resource.TestCheckResourceAttr(resourceName, "kerberos.0.realm", "realm"),
					resource.TestCheckResourceAttr(resourceName, "kerberos.0.tgt_generation_method", "Password"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"port",
				},
			},
		},
	})
}

func TestAccSecretSSHKey_sshkey_sshReferenceCredential(t *testing.T) {

	id := fmt.Sprintf("%s_%s", "TestsshReferenceCredential", utils.RandStringBytes(5))
	name := id
	updatedName := fmt.Sprintf("%s_updated", name)
	resourceName := "harness_platform_secret_sshkey.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccSecretDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceSecret_sshkey_sshReferenceCredential(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "ssh.0.credential_type", "KeyReference"),
					resource.TestCheckResourceAttr(resourceName, "ssh.0.sshkey_reference_credential.0.user_name", "user_name"),
				),
			},
			{
				Config: testAccResourceSecret_sshkey_sshReferenceCredential(id, updatedName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "ssh.0.credential_type", "KeyReference"),
					resource.TestCheckResourceAttr(resourceName, "ssh.0.sshkey_reference_credential.0.user_name", "user_name"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"port",
				},
			},
		},
	})
}

func TestAccSecretSSHKey_sshkey_sshPathCredential(t *testing.T) {

	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	name := id
	updatedName := fmt.Sprintf("%s_updated", name)
	resourceName := "harness_platform_secret_sshkey.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccSecretDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceSecret_sshkey_sshPathCredential(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "ssh.0.credential_type", "KeyPath"),
					resource.TestCheckResourceAttr(resourceName, "ssh.0.sshkey_path_credential.0.user_name", "user_name"),
					resource.TestCheckResourceAttr(resourceName, "ssh.0.sshkey_path_credential.0.key_path", "key_path"),
					resource.TestCheckResourceAttr(resourceName, "ssh.0.sshkey_path_credential.0.encrypted_passphrase", "encrypted_passphrase"),
				),
			},
			{
				Config: testAccResourceSecret_sshkey_sshPathCredential(id, updatedName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "ssh.0.credential_type", "KeyPath"),
					resource.TestCheckResourceAttr(resourceName, "ssh.0.sshkey_path_credential.0.user_name", "user_name"),
					resource.TestCheckResourceAttr(resourceName, "ssh.0.sshkey_path_credential.0.key_path", "key_path"),
					resource.TestCheckResourceAttr(resourceName, "ssh.0.sshkey_path_credential.0.encrypted_passphrase", "encrypted_passphrase"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"port",
				},
			},
		},
	})
}

func TestAccSecretSSHKey_sshkey_sshPassword(t *testing.T) {

	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	name := id
	updatedName := fmt.Sprintf("%s_updated", name)
	resourceName := "harness_platform_secret_sshkey.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccSecretDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceSecret_sshkey_sshPassword(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "ssh.0.credential_type", "Password"),
					resource.TestCheckResourceAttr(resourceName, "ssh.0.ssh_password_credential.0.user_name", "user_name"),
				),
			},
			{
				Config: testAccResourceSecret_sshkey_sshPassword(id, updatedName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "ssh.0.credential_type", "Password"),
					resource.TestCheckResourceAttr(resourceName, "ssh.0.ssh_password_credential.0.user_name", "user_name"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"port",
				},
			},
		},
	})
}

func testAccResourceSecret_sshkey_kerberos_keyFilePath(id string, name string) string {
	return fmt.Sprintf(`
		resource "harness_platform_secret_sshkey" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
			description = "test"
			tags = ["foo:bar"]
			port = 22
			kerberos {
				tgt_key_tab_file_path_spec {
					key_path = "key_path"
				}
				principal = "principal"
				realm = "realm"
				tgt_generation_method = "KeyTabFilePath"
			}
		}
`, id, name)
}

func testAccResourceSecret_sshkey_kerberos_password(id string, name string) string {
	return fmt.Sprintf(`
	resource "harness_platform_secret_file" "test" {
		identifier = "%[1]s_a"
		name = "%[2]s_a"
		description = "test"
		tags = ["foo:bar"]
		file_path = "%[3]s"
		secret_manager_identifier = "harnessSecretManager"
	}
		resource "harness_platform_secret_sshkey" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
			description = "test"
			tags = ["foo:bar"]
			port = 22
			kerberos {
				tgt_password_spec {
					password = harness_platform_secret_file.test.id
				}
				principal = "principal"
				realm = "realm"
				tgt_generation_method = "Password"
			}
		}
`, id, name, getAbsFilePath("../../../acctest/secret_files/secret.txt"))
}

func testAccResourceSecret_sshkey_sshReferenceCredential(id string, name string) string {
	return fmt.Sprintf(`
	resource "harness_platform_secret_file" "test" {
		identifier = "%[1]s_a"
		name = "%[2]s_a"
		description = "test"
		tags = ["foo:bar"]
		file_path = "%[3]s"
		secret_manager_identifier = "harnessSecretManager"
	}
		resource "harness_platform_secret_sshkey" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
			description = "test"
			tags = ["foo:bar"]
			port = 22
			ssh {
				sshkey_reference_credential {
					user_name = "user_name"
					key = harness_platform_secret_file.test.id
					encrypted_passphrase = harness_platform_secret_file.test.id
				}
				credential_type = "KeyReference"
			}
		}
`, id, name, getAbsFilePath("../../../acctest/secret_files/secret.txt"))
}

func testAccResourceSecret_sshkey_sshPathCredential(id string, name string) string {
	return fmt.Sprintf(`
	resource "harness_platform_secret_sshkey" "test" {
		identifier = "%[1]s"
		name = "%[2]s"
		description = "test"
		tags = ["foo:bar"]
		port = 22
		ssh {
			sshkey_path_credential {
				user_name = "user_name"
				key_path = "key_path"
				encrypted_passphrase = "encrypted_passphrase"
			}
			credential_type = "KeyPath"
		}
	}
`, id, name)
}

func testAccResourceSecret_sshkey_sshPassword(id string, name string) string {
	return fmt.Sprintf(`
	resource "harness_platform_secret_file" "test" {
		identifier = "%[1]s_a"
		name = "%[2]s_a"
		description = "test"
		tags = ["foo:bar"]
		file_path = "%[3]s"
		secret_manager_identifier = "harnessSecretManager"
	}
	resource "harness_platform_secret_sshkey" "test" {
		identifier = "%[1]s"
		name = "%[2]s"
		description = "test"
		tags = ["foo:bar"]
		port = 22
		ssh {
			ssh_password_credential {
				user_name = "user_name"
				password = harness_platform_secret_file.test.id
			}
			credential_type = "Password"
		}
	}
`, id, name, getAbsFilePath("../../../acctest/secret_files/secret.txt"))
}
