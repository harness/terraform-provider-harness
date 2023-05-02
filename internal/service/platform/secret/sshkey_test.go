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

func TestAccSecretSSHKey_kerberos_keyFilePathProject(t *testing.T) {

	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	name := id
	updatedName := fmt.Sprintf("%s_updated", name)
	resourceName := "harness_platform_secret_sshkey.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccSecretDestroy(resourceName),
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {},
		},
		Steps: []resource.TestStep{
			{
				Config: testAccResourceSecret_sshkey_kerberos_keyFilePathProject(id, name),
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
				Config: testAccResourceSecret_sshkey_kerberos_keyFilePathProject(id, updatedName),
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
				ImportStateIdFunc: acctest.ProjectResourceImportStateIdFunc(resourceName),
				ImportStateVerifyIgnore: []string{
					"port",
				},
			},
		},
	})
}

func TestAccSecretSSHKey_kerberos_keyFilePathOrg(t *testing.T) {

	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	name := id
	updatedName := fmt.Sprintf("%s_updated", name)
	resourceName := "harness_platform_secret_sshkey.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccSecretDestroy(resourceName),
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {},
		},
		Steps: []resource.TestStep{
			{
				Config: testAccResourceSecret_sshkey_kerberos_keyFilePathOrg(id, name),
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
				Config: testAccResourceSecret_sshkey_kerberos_keyFilePathOrg(id, updatedName),
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
				ImportStateIdFunc: acctest.OrgResourceImportStateIdFunc(resourceName),
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

func TestAccSecretSSHKey_kerberos_passwordProject(t *testing.T) {

	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	name := id
	updatedName := fmt.Sprintf("%s_updated", name)
	resourceName := "harness_platform_secret_sshkey.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccSecretDestroy(resourceName),
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {},
		},
		Steps: []resource.TestStep{
			{
				Config: testAccResourceSecret_sshkey_kerberos_passwordProject(id, name),
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
				Config: testAccResourceSecret_sshkey_kerberos_passwordProject(id, updatedName),
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
				ImportStateIdFunc:       acctest.ProjectResourceImportStateIdFunc(resourceName),
				ImportStateVerifyIgnore: []string{
					"port",
				},
			},
		},
	})
}
func TestAccSecretSSHKey_kerberos_passwordOrg(t *testing.T) {

	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	name := id
	updatedName := fmt.Sprintf("%s_updated", name)
	resourceName := "harness_platform_secret_sshkey.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccSecretDestroy(resourceName),
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {},
		},
		Steps: []resource.TestStep{
			{
				Config: testAccResourceSecret_sshkey_kerberos_passwordOrg(id, name),
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
				Config: testAccResourceSecret_sshkey_kerberos_passwordOrg(id, updatedName),
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
				ImportStateIdFunc:       acctest.OrgResourceImportStateIdFunc(resourceName),
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

func TestAccSecretSSHKey_sshkey_sshReferenceCredentialProject(t *testing.T) {

	id := fmt.Sprintf("%s_%s", "TestsshReferenceCredential", utils.RandStringBytes(5))
	name := id
	updatedName := fmt.Sprintf("%s_updated", name)
	resourceName := "harness_platform_secret_sshkey.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccSecretDestroy(resourceName),
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {},
		},
		Steps: []resource.TestStep{
			{
				Config: testAccResourceSecret_sshkey_sshReferenceCredentialProject(id, name),
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
				Config: testAccResourceSecret_sshkey_sshReferenceCredentialProject(id, updatedName),
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
				ImportStateIdFunc:       acctest.ProjectResourceImportStateIdFunc(resourceName),
				ImportStateVerifyIgnore: []string{
					"port",
				},
			},
		},
	})
}

func TestAccSecretSSHKey_sshkey_sshReferenceCredentialOrg(t *testing.T) {

	id := fmt.Sprintf("%s_%s", "TestsshReferenceCredential", utils.RandStringBytes(5))
	name := id
	updatedName := fmt.Sprintf("%s_updated", name)
	resourceName := "harness_platform_secret_sshkey.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccSecretDestroy(resourceName),
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {},
		},
		Steps: []resource.TestStep{
			{
				Config: testAccResourceSecret_sshkey_sshReferenceCredentialOrg(id, name),
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
				Config: testAccResourceSecret_sshkey_sshReferenceCredentialOrg(id, updatedName),
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
				ImportStateIdFunc:       acctest.OrgResourceImportStateIdFunc(resourceName),
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
					resource.TestCheckResourceAttr(resourceName, "ssh.0.sshkey_path_credential.0.key_path", "account."+ id+"_a"),
					resource.TestCheckResourceAttr(resourceName, "ssh.0.sshkey_path_credential.0.encrypted_passphrase", "account."+ id+"_a"),
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
					resource.TestCheckResourceAttr(resourceName, "ssh.0.sshkey_path_credential.0.key_path", "account."+ id+"_a"),
					resource.TestCheckResourceAttr(resourceName, "ssh.0.sshkey_path_credential.0.encrypted_passphrase", "account."+ id+"_a"),
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

func TestAccSecretSSHKey_sshkey_sshPathCredentialProject(t *testing.T) {

	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	name := id
	updatedName := fmt.Sprintf("%s_updated", name)
	resourceName := "harness_platform_secret_sshkey.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccSecretDestroy(resourceName),
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {},
		},
		Steps: []resource.TestStep{
			{
				Config: testAccResourceSecret_sshkey_sshPathCredentialProject(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "ssh.0.credential_type", "KeyPath"),
					resource.TestCheckResourceAttr(resourceName, "ssh.0.sshkey_path_credential.0.user_name", "user_name"),
					resource.TestCheckResourceAttr(resourceName, "ssh.0.sshkey_path_credential.0.key_path", id+"_a"),
					resource.TestCheckResourceAttr(resourceName, "ssh.0.sshkey_path_credential.0.encrypted_passphrase", id + "_a"),
				),
			},
			{
				Config: testAccResourceSecret_sshkey_sshPathCredentialProject(id, updatedName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "ssh.0.credential_type", "KeyPath"),
					resource.TestCheckResourceAttr(resourceName, "ssh.0.sshkey_path_credential.0.user_name", "user_name"),
					resource.TestCheckResourceAttr(resourceName, "ssh.0.sshkey_path_credential.0.key_path", id+"_a"),
					resource.TestCheckResourceAttr(resourceName, "ssh.0.sshkey_path_credential.0.encrypted_passphrase", id + "_a"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc:       acctest.ProjectResourceImportStateIdFunc(resourceName),
				ImportStateVerifyIgnore: []string{
					"port",
				},
			},
		},
	})
}

func TestAccSecretSSHKey_sshkey_sshPathCredentialOrg(t *testing.T) {

	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	name := id
	updatedName := fmt.Sprintf("%s_updated", name)
	resourceName := "harness_platform_secret_sshkey.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccSecretDestroy(resourceName),
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {},
		},
		Steps: []resource.TestStep{
			{
				Config: testAccResourceSecret_sshkey_sshPathCredentialOrg(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "ssh.0.credential_type", "KeyPath"),
					resource.TestCheckResourceAttr(resourceName, "ssh.0.sshkey_path_credential.0.user_name", "user_name"),
					resource.TestCheckResourceAttr(resourceName, "ssh.0.sshkey_path_credential.0.key_path", "org."+ id+"_a"),
					resource.TestCheckResourceAttr(resourceName, "ssh.0.sshkey_path_credential.0.encrypted_passphrase", "org."+ id+"_a"),
				),
			},
			{
				Config: testAccResourceSecret_sshkey_sshPathCredentialOrg(id, updatedName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "ssh.0.credential_type", "KeyPath"),
					resource.TestCheckResourceAttr(resourceName, "ssh.0.sshkey_path_credential.0.user_name", "user_name"),
					resource.TestCheckResourceAttr(resourceName, "ssh.0.sshkey_path_credential.0.key_path", "org."+ id+"_a"),
					resource.TestCheckResourceAttr(resourceName, "ssh.0.sshkey_path_credential.0.encrypted_passphrase", "org."+ id+"_a"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc:       acctest.OrgResourceImportStateIdFunc(resourceName),
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

func TestAccSecretSSHKey_sshkey_sshPasswordProject(t *testing.T) {

	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	name := id
	updatedName := fmt.Sprintf("%s_updated", name)
	resourceName := "harness_platform_secret_sshkey.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccSecretDestroy(resourceName),
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {},
		},
		Steps: []resource.TestStep{
			{
				Config: testAccResourceSecret_sshkey_sshPasswordProject(id, name),
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
				Config: testAccResourceSecret_sshkey_sshPasswordProject(id, updatedName),
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
				ImportStateIdFunc:       acctest.ProjectResourceImportStateIdFunc(resourceName),
				ImportStateVerifyIgnore: []string{
					"port",
				},
			},
		},
	})
}

func TestAccSecretSSHKey_sshkey_sshPasswordOrg(t *testing.T) {

	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	name := id
	updatedName := fmt.Sprintf("%s_updated", name)
	resourceName := "harness_platform_secret_sshkey.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccSecretDestroy(resourceName),
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {},
		},
		Steps: []resource.TestStep{
			{
				Config: testAccResourceSecret_sshkey_sshPasswordOrg(id, name),
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
				Config: testAccResourceSecret_sshkey_sshPasswordOrg(id, updatedName),
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
				ImportStateIdFunc:       acctest.OrgResourceImportStateIdFunc(resourceName),
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
func testAccResourceSecret_sshkey_kerberos_keyFilePathProject(id string, name string) string {
	return fmt.Sprintf(`
	resource "harness_platform_organization" "test" {
    identifier = "%[1]s"
    name = "%[2]s"
  }
  resource "harness_platform_project" "test" {
    identifier = "%[1]s"
    name = "%[2]s"
    org_id = harness_platform_organization.test.id
    color = "#472848"
	}

		resource "harness_platform_secret_sshkey" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
			description = "test"
			org_id = harness_platform_organization.test.id
			project_id = harness_platform_project.test.id
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
			depends_on = [time_sleep.wait_3_seconds]
		}
		resource "time_sleep" "wait_3_seconds" {
			create_duration = "3s"
			depends_on = [harness_platform_project.test]
			
		}
`, id, name)
}
func testAccResourceSecret_sshkey_kerberos_keyFilePathOrg(id string, name string) string {
	return fmt.Sprintf(`
	resource "harness_platform_organization" "test" {
    identifier = "%[1]s"
    name = "%[2]s"
  }

		resource "harness_platform_secret_sshkey" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
			description = "test"
			org_id = harness_platform_organization.test.id
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
			depends_on = [time_sleep.wait_3_seconds]
		}
		resource "time_sleep" "wait_3_seconds" {
			create_duration = "3s"
			depends_on = [harness_platform_organization.test]
			
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
					password = "account.${harness_platform_secret_file.test.id}"
				}
				principal = "principal"
				realm = "realm"
				tgt_generation_method = "Password"
			}
		}
`, id, name, getAbsFilePath("../../../acctest/secret_files/secret.txt"))
}
func testAccResourceSecret_sshkey_kerberos_passwordProject(id string, name string) string {
	return fmt.Sprintf(`
	resource "harness_platform_organization" "test" {
    identifier = "%[1]s"
    name = "%[2]s"
  }
  resource "harness_platform_project" "test" {
    identifier = "%[1]s"
    name = "%[2]s"
    org_id = harness_platform_organization.test.id
    color = "#472848"
	}
	resource "harness_platform_secret_file" "test" {
		identifier = "%[1]s_a"
		name = "%[2]s_a"
		description = "test"
		org_id = harness_platform_organization.test.id
		project_id = harness_platform_project.test.id
		tags = ["foo:bar"]
		file_path = "%[3]s"
		secret_manager_identifier = "harnessSecretManager"
		depends_on = [time_sleep.wait_3_seconds]
	}
		resource "time_sleep" "wait_3_seconds" {
			create_duration = "3s"
			depends_on = [harness_platform_project.test]
			
		}

		resource "harness_platform_secret_sshkey" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
			description = "test"
			org_id = harness_platform_organization.test.id
			project_id = harness_platform_project.test.id
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
			depends_on = [time_sleep.wait_4_seconds]
		}
		resource "time_sleep" "wait_4_seconds" {
			create_duration = "4s"
			depends_on = [harness_platform_secret_file.test]
			
		}
`, id, name, getAbsFilePath("../../../acctest/secret_files/secret.txt"))
}

func testAccResourceSecret_sshkey_kerberos_passwordOrg(id string, name string) string {
	return fmt.Sprintf(`
	resource "harness_platform_organization" "test" {
    identifier = "%[1]s"
    name = "%[2]s"
  }

	resource "harness_platform_secret_file" "test" {
		identifier = "%[1]s_a"
		name = "%[2]s_a"
		description = "test"
		tags = ["foo:bar"]
		org_id = harness_platform_organization.test.id
		file_path = "%[3]s"
		secret_manager_identifier = "harnessSecretManager"
		depends_on = [time_sleep.wait_3_seconds]
	}
		resource "time_sleep" "wait_3_seconds" {
			create_duration = "3s"
			depends_on = [harness_platform_organization.test]
			
		}
		resource "harness_platform_secret_sshkey" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
			description = "test"
			org_id = harness_platform_organization.test.id
			tags = ["foo:bar"]
			port = 22
			kerberos {
				tgt_password_spec {
					password = "org.${harness_platform_secret_file.test.id}"
				}
				principal = "principal"
				realm = "realm"
				tgt_generation_method = "Password"
			}
			depends_on = [time_sleep.wait_4_seconds]
		}
		resource "time_sleep" "wait_4_seconds" {
			create_duration = "4s"
			depends_on = [harness_platform_secret_file.test]
			
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
					key = "account.${harness_platform_secret_file.test.id}"
					encrypted_passphrase = "account.${harness_platform_secret_file.test.id}"
				}
				credential_type = "KeyReference"
			}
		}
`, id, name, getAbsFilePath("../../../acctest/secret_files/secret.txt"))
}

func testAccResourceSecret_sshkey_sshReferenceCredentialProject(id string, name string) string {
	return fmt.Sprintf(`
	resource "harness_platform_organization" "test" {
    identifier = "%[1]s"
    name = "%[2]s"
  }
  resource "harness_platform_project" "test" {
    identifier = "%[1]s"
    name = "%[2]s"
    org_id = harness_platform_organization.test.id
    color = "#472848"
	}
	resource "harness_platform_secret_file" "test" {
		identifier = "%[1]s_a"
		name = "%[2]s_a"
		description = "test"
		org_id = harness_platform_organization.test.id
		project_id = harness_platform_project.test.id
		tags = ["foo:bar"]
		file_path = "%[3]s"
		secret_manager_identifier = "harnessSecretManager"
		depends_on = [time_sleep.wait_3_seconds]
	}
		resource "time_sleep" "wait_3_seconds" {
			create_duration = "3s"
			depends_on = [harness_platform_project.test]
			
		}

		resource "harness_platform_secret_sshkey" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
			description = "test"
			org_id = harness_platform_organization.test.id
			project_id = harness_platform_project.test.id
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
			depends_on = [time_sleep.wait_4_seconds]
		}
			resource "time_sleep" "wait_4_seconds" {
				create_duration = "4s"
				depends_on = [harness_platform_secret_file.test]
				
			}
`, id, name, getAbsFilePath("../../../acctest/secret_files/secret.txt"))
}

func testAccResourceSecret_sshkey_sshReferenceCredentialOrg(id string, name string) string {
	return fmt.Sprintf(`
	resource "harness_platform_organization" "test" {
    identifier = "%[1]s"
    name = "%[2]s"
  }

	resource "harness_platform_secret_file" "test" {
		identifier = "%[1]s_a"
		name = "%[2]s_a"
		description = "test"
		org_id = harness_platform_organization.test.id
		tags = ["foo:bar"]
		file_path = "%[3]s"
		secret_manager_identifier = "harnessSecretManager"
		depends_on = [time_sleep.wait_4_seconds]
	}
		resource "time_sleep" "wait_4_seconds" {
			create_duration = "4s"
			depends_on = [harness_platform_organization.test]	
		}

		resource "harness_platform_secret_sshkey" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
			org_id = harness_platform_organization.test.id
			description = "test"
			tags = ["foo:bar"]
			port = 22
			ssh {
				sshkey_reference_credential {
					user_name = "user_name"
					key = "org.${harness_platform_secret_file.test.id}"
					encrypted_passphrase = "org.${harness_platform_secret_file.test.id}"
				}
				credential_type = "KeyReference"
			}
			depends_on = [time_sleep.wait_3_seconds]
		}
			resource "time_sleep" "wait_3_seconds" {
				create_duration = "3s"
				depends_on = [harness_platform_secret_file.test]	
			}
`, id, name, getAbsFilePath("../../../acctest/secret_files/secret.txt"))
}

func testAccResourceSecret_sshkey_sshPathCredential(id string, name string) string {
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
			sshkey_path_credential {
				user_name = "user_name"
				key_path = "account.${harness_platform_secret_file.test.id}"
				encrypted_passphrase = "account.${harness_platform_secret_file.test.id}"
			}
			credential_type = "KeyPath"
		}
	}
`, id, name,getAbsFilePath("../../../acctest/secret_files/secret.txt"))
}

func testAccResourceSecret_sshkey_sshPathCredentialProject(id string, name string) string {
	return fmt.Sprintf(`
	resource "harness_platform_organization" "test" {
    identifier = "%[1]s"
    name = "%[2]s"
  }
  resource "harness_platform_project" "test" {
    identifier = "%[1]s"
    name = "%[2]s"
    org_id = harness_platform_organization.test.id
    color = "#472848"
	}
	resource "harness_platform_secret_file" "test" {
		identifier = "%[1]s_a"
		name = "%[2]s_a"
		description = "test"
		org_id = harness_platform_organization.test.id
		project_id = harness_platform_project.test.id
		tags = ["foo:bar"]
		file_path = "%[3]s"
		secret_manager_identifier = "harnessSecretManager"
		depends_on = [time_sleep.wait_3_seconds]
	}
		resource "time_sleep" "wait_3_seconds" {
			create_duration = "3s"
			depends_on = [harness_platform_project.test]
			
		}
	resource "harness_platform_secret_sshkey" "test" {
		identifier = "%[1]s"
		name = "%[2]s"
		description = "test"
		org_id = harness_platform_organization.test.id
		project_id = harness_platform_project.test.id
		tags = ["foo:bar"]
		port = 22
		ssh {
			sshkey_path_credential {
				user_name = "user_name"
				key_path = "${harness_platform_secret_file.test.id}"
				encrypted_passphrase = "${harness_platform_secret_file.test.id}"
			}
			credential_type = "KeyPath"
		}
		depends_on = [time_sleep.wait_4_seconds]
	}
		resource "time_sleep" "wait_4_seconds" {
			create_duration = "4s"
			depends_on = [harness_platform_secret_file.test]
			
		}
`, id, name,getAbsFilePath("../../../acctest/secret_files/secret.txt"))
}

func testAccResourceSecret_sshkey_sshPathCredentialOrg(id string, name string) string {
	return fmt.Sprintf(`
	resource "harness_platform_organization" "test" {
    identifier = "%[1]s"
    name = "%[2]s"
  }
	resource "harness_platform_secret_file" "test" {
		identifier = "%[1]s_a"
		name = "%[2]s_a"
		description = "test"
		org_id = harness_platform_organization.test.id
		tags = ["foo:bar"]
		file_path = "%[3]s"
		secret_manager_identifier = "harnessSecretManager"
		depends_on = [time_sleep.wait_3_seconds]
	}
		resource "time_sleep" "wait_3_seconds" {
			create_duration = "3s"
			depends_on = [harness_platform_organization.test]
			
		}
	resource "harness_platform_secret_sshkey" "test" {
		identifier = "%[1]s"
		name = "%[2]s"
		description = "test"
		org_id = harness_platform_organization.test.id
		tags = ["foo:bar"]
		port = 22
		ssh {
			sshkey_path_credential {
				user_name = "user_name"
				key_path = "org.${harness_platform_secret_file.test.id}"
				encrypted_passphrase = "org.${harness_platform_secret_file.test.id}"
			}
			credential_type = "KeyPath"
		}
		depends_on = [time_sleep.wait_4_seconds]
	}
		resource "time_sleep" "wait_4_seconds" {
			create_duration = "4s"
			depends_on = [harness_platform_secret_file.test]
			
		}
`, id, name,getAbsFilePath("../../../acctest/secret_files/secret.txt"))
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
				password = "account.${harness_platform_secret_file.test.id}"
			}
			credential_type = "Password"
		}
	}
`, id, name, getAbsFilePath("../../../acctest/secret_files/secret.txt"))
}

func testAccResourceSecret_sshkey_sshPasswordProject(id string, name string) string {
	return fmt.Sprintf(`
	resource "harness_platform_organization" "test" {
    identifier = "%[1]s"
    name = "%[2]s"
  }
  resource "harness_platform_project" "test" {
    identifier = "%[1]s"
    name = "%[2]s"
    org_id = harness_platform_organization.test.id
    color = "#472848"
	}
	resource "harness_platform_secret_file" "test" {
		identifier = "%[1]s_a"
		name = "%[2]s_a"
		description = "test"
		tags = ["foo:bar"]
		org_id = harness_platform_organization.test.id
		project_id = harness_platform_project.test.id
		file_path = "%[3]s"
		secret_manager_identifier = "harnessSecretManager"
		depends_on = [time_sleep.wait_3_seconds]
	}
		resource "time_sleep" "wait_3_seconds" {
			create_duration = "3s"
			depends_on = [harness_platform_project.test]
			
		}
	resource "harness_platform_secret_sshkey" "test" {
		identifier = "%[1]s"
		name = "%[2]s"
		description = "test"
		org_id = harness_platform_organization.test.id
		project_id = harness_platform_project.test.id
		tags = ["foo:bar"]
		port = 22
		ssh {
			ssh_password_credential {
				user_name = "user_name"
				password = harness_platform_secret_file.test.id
			}
			credential_type = "Password"
		}
		depends_on = [time_sleep.wait_4_seconds]
	}
		resource "time_sleep" "wait_4_seconds" {
			create_duration = "4s"
			depends_on = [harness_platform_secret_file.test]
			
		}
`, id, name, getAbsFilePath("../../../acctest/secret_files/secret.txt"))
}

func testAccResourceSecret_sshkey_sshPasswordOrg(id string, name string) string {
	return fmt.Sprintf(`
	resource "harness_platform_organization" "test" {
    identifier = "%[1]s"
    name = "%[2]s"
  }
	resource "harness_platform_secret_file" "test" {
		identifier = "%[1]s_a"
		name = "%[2]s_a"
		description = "test"
		tags = ["foo:bar"]
		org_id = harness_platform_organization.test.id
		file_path = "%[3]s"
		secret_manager_identifier = "harnessSecretManager"
		depends_on = [time_sleep.wait_3_seconds]
	}
		resource "time_sleep" "wait_3_seconds" {
			create_duration = "3s"
			depends_on = [harness_platform_organization.test]
			
		}
	resource "harness_platform_secret_sshkey" "test" {
		identifier = "%[1]s"
		name = "%[2]s"
		description = "test"
		org_id = harness_platform_organization.test.id
		tags = ["foo:bar"]
		port = 22
		ssh {
			ssh_password_credential {
				user_name = "user_name"
				password = "org.${harness_platform_secret_file.test.id}"
			}
			credential_type = "Password"
		}
		depends_on = [time_sleep.wait_4_seconds]
	}
		resource "time_sleep" "wait_4_seconds" {
			create_duration = "4s"
			depends_on = [harness_platform_secret_file.test]
			
		}
`, id, name, getAbsFilePath("../../../acctest/secret_files/secret.txt"))
}
