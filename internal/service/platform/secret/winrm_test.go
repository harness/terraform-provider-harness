package secret_test

import (
	"fmt"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccSecretWinRM_ntlm(t *testing.T) {

	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	name := id
	updatedName := fmt.Sprintf("%s_updated", name)
	resourceName := "harness_platform_secret_winrm.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccSecretDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceSecret_winrm_ntlm(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "port", "5986"),
					resource.TestCheckResourceAttr(resourceName, "ntlm.0.domain", "example.com"),
					resource.TestCheckResourceAttr(resourceName, "ntlm.0.username", "admin"),
					resource.TestCheckResourceAttr(resourceName, "ntlm.0.use_ssl", "true"),
					resource.TestCheckResourceAttr(resourceName, "ntlm.0.skip_cert_check", "false"),
				),
			},
			{
				Config: testAccResourceSecret_winrm_ntlm(id, updatedName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "port", "5986"),
					resource.TestCheckResourceAttr(resourceName, "ntlm.0.domain", "example.com"),
					resource.TestCheckResourceAttr(resourceName, "ntlm.0.username", "admin"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"ntlm.0.password_ref",
				},
			},
		},
	})
}

func TestAccSecretWinRM_ntlm_project(t *testing.T) {

	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	name := id
	updatedName := fmt.Sprintf("%s_updated", name)
	resourceName := "harness_platform_secret_winrm.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccSecretDestroy(resourceName),
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {},
		},
		Steps: []resource.TestStep{
			{
				Config: testAccResourceSecret_winrm_ntlm_project(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "port", "5986"),
					resource.TestCheckResourceAttr(resourceName, "ntlm.0.domain", "example.com"),
					resource.TestCheckResourceAttr(resourceName, "ntlm.0.username", "admin"),
				),
			},
			{
				Config: testAccResourceSecret_winrm_ntlm_project(id, updatedName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "port", "5986"),
					resource.TestCheckResourceAttr(resourceName, "ntlm.0.domain", "example.com"),
					resource.TestCheckResourceAttr(resourceName, "ntlm.0.username", "admin"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: acctest.ProjectResourceImportStateIdFunc(resourceName),
				ImportStateVerifyIgnore: []string{
					"ntlm.0.password_ref",
				},
			},
		},
	})
}

func TestAccSecretWinRM_kerberos_password(t *testing.T) {

	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	name := id
	updatedName := fmt.Sprintf("%s_updated", name)
	resourceName := "harness_platform_secret_winrm.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccSecretDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceSecret_winrm_kerberos_password(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "port", "5986"),
					resource.TestCheckResourceAttr(resourceName, "kerberos.0.principal", "user@EXAMPLE.COM"),
					resource.TestCheckResourceAttr(resourceName, "kerberos.0.realm", "EXAMPLE.COM"),
					resource.TestCheckResourceAttr(resourceName, "kerberos.0.tgt_generation_method", "Password"),
				),
			},
			{
				Config: testAccResourceSecret_winrm_kerberos_password(id, updatedName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "port", "5986"),
					resource.TestCheckResourceAttr(resourceName, "kerberos.0.principal", "user@EXAMPLE.COM"),
					resource.TestCheckResourceAttr(resourceName, "kerberos.0.realm", "EXAMPLE.COM"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"kerberos.0.tgt_password_spec.0.password_ref",
				},
			},
		},
	})
}

func TestAccSecretWinRM_kerberos_keytab(t *testing.T) {

	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	name := id
	updatedName := fmt.Sprintf("%s_updated", name)
	resourceName := "harness_platform_secret_winrm.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccSecretDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceSecret_winrm_kerberos_keytab(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "port", "5986"),
					resource.TestCheckResourceAttr(resourceName, "kerberos.0.principal", "user@EXAMPLE.COM"),
					resource.TestCheckResourceAttr(resourceName, "kerberos.0.realm", "EXAMPLE.COM"),
					resource.TestCheckResourceAttr(resourceName, "kerberos.0.tgt_generation_method", "KeyTabFilePath"),
					resource.TestCheckResourceAttr(resourceName, "kerberos.0.tgt_key_tab_file_path_spec.0.key_path", "/path/to/keytab"),
					resource.TestCheckResourceAttr(resourceName, "kerberos.0.use_ssl", "true"),
					resource.TestCheckResourceAttr(resourceName, "kerberos.0.skip_cert_check", "false"),
					resource.TestCheckResourceAttr(resourceName, "kerberos.0.use_no_profile", "true"),
				),
			},
			{
				Config: testAccResourceSecret_winrm_kerberos_keytab(id, updatedName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "port", "5986"),
					resource.TestCheckResourceAttr(resourceName, "kerberos.0.principal", "user@EXAMPLE.COM"),
					resource.TestCheckResourceAttr(resourceName, "kerberos.0.realm", "EXAMPLE.COM"),
					resource.TestCheckResourceAttr(resourceName, "kerberos.0.tgt_generation_method", "KeyTabFilePath"),
					resource.TestCheckResourceAttr(resourceName, "kerberos.0.tgt_key_tab_file_path_spec.0.key_path", "/path/to/keytab"),
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

func TestAccSecretWinRM_ntlm_org(t *testing.T) {

	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	name := id
	updatedName := fmt.Sprintf("%s_updated", name)
	resourceName := "harness_platform_secret_winrm.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccSecretDestroy(resourceName),
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {},
		},
		Steps: []resource.TestStep{
			{
				Config: testAccResourceSecret_winrm_ntlm_org(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "port", "5986"),
					resource.TestCheckResourceAttr(resourceName, "ntlm.0.domain", "example.com"),
					resource.TestCheckResourceAttr(resourceName, "ntlm.0.username", "admin"),
					resource.TestCheckResourceAttr(resourceName, "ntlm.0.use_ssl", "true"),
					resource.TestCheckResourceAttr(resourceName, "ntlm.0.skip_cert_check", "false"),
				),
			},
			{
				Config: testAccResourceSecret_winrm_ntlm_org(id, updatedName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "port", "5986"),
					resource.TestCheckResourceAttr(resourceName, "ntlm.0.domain", "example.com"),
					resource.TestCheckResourceAttr(resourceName, "ntlm.0.username", "admin"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: acctest.OrgResourceImportStateIdFunc(resourceName),
				ImportStateVerifyIgnore: []string{
					"ntlm.0.password_ref",
				},
			},
		},
	})
}

func TestAccSecretWinRM_kerberos_password_project(t *testing.T) {

	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	name := id
	updatedName := fmt.Sprintf("%s_updated", name)
	resourceName := "harness_platform_secret_winrm.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccSecretDestroy(resourceName),
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {},
		},
		Steps: []resource.TestStep{
			{
				Config: testAccResourceSecret_winrm_kerberos_password_project(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "port", "5986"),
					resource.TestCheckResourceAttr(resourceName, "kerberos.0.principal", "user@EXAMPLE.COM"),
					resource.TestCheckResourceAttr(resourceName, "kerberos.0.realm", "EXAMPLE.COM"),
					resource.TestCheckResourceAttr(resourceName, "kerberos.0.tgt_generation_method", "Password"),
				),
			},
			{
				Config: testAccResourceSecret_winrm_kerberos_password_project(id, updatedName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "port", "5986"),
					resource.TestCheckResourceAttr(resourceName, "kerberos.0.principal", "user@EXAMPLE.COM"),
					resource.TestCheckResourceAttr(resourceName, "kerberos.0.realm", "EXAMPLE.COM"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: acctest.ProjectResourceImportStateIdFunc(resourceName),
				ImportStateVerifyIgnore: []string{
					"kerberos.0.tgt_password_spec.0.password_ref",
				},
			},
		},
	})
}

func TestAccSecretWinRM_kerberos_password_org(t *testing.T) {

	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	name := id
	updatedName := fmt.Sprintf("%s_updated", name)
	resourceName := "harness_platform_secret_winrm.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccSecretDestroy(resourceName),
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {},
		},
		Steps: []resource.TestStep{
			{
				Config: testAccResourceSecret_winrm_kerberos_password_org(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "port", "5986"),
					resource.TestCheckResourceAttr(resourceName, "kerberos.0.principal", "user@EXAMPLE.COM"),
					resource.TestCheckResourceAttr(resourceName, "kerberos.0.realm", "EXAMPLE.COM"),
					resource.TestCheckResourceAttr(resourceName, "kerberos.0.tgt_generation_method", "Password"),
				),
			},
			{
				Config: testAccResourceSecret_winrm_kerberos_password_org(id, updatedName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "port", "5986"),
					resource.TestCheckResourceAttr(resourceName, "kerberos.0.principal", "user@EXAMPLE.COM"),
					resource.TestCheckResourceAttr(resourceName, "kerberos.0.realm", "EXAMPLE.COM"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: acctest.OrgResourceImportStateIdFunc(resourceName),
				ImportStateVerifyIgnore: []string{
					"kerberos.0.tgt_password_spec.0.password_ref",
				},
			},
		},
	})
}

func TestAccSecretWinRM_kerberos_keytab_project(t *testing.T) {

	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	name := id
	updatedName := fmt.Sprintf("%s_updated", name)
	resourceName := "harness_platform_secret_winrm.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccSecretDestroy(resourceName),
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {},
		},
		Steps: []resource.TestStep{
			{
				Config: testAccResourceSecret_winrm_kerberos_keytab_project(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "port", "5986"),
					resource.TestCheckResourceAttr(resourceName, "kerberos.0.principal", "user@EXAMPLE.COM"),
					resource.TestCheckResourceAttr(resourceName, "kerberos.0.realm", "EXAMPLE.COM"),
					resource.TestCheckResourceAttr(resourceName, "kerberos.0.tgt_generation_method", "KeyTabFilePath"),
					resource.TestCheckResourceAttr(resourceName, "kerberos.0.tgt_key_tab_file_path_spec.0.key_path", "/path/to/keytab"),
				),
			},
			{
				Config: testAccResourceSecret_winrm_kerberos_keytab_project(id, updatedName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "port", "5986"),
					resource.TestCheckResourceAttr(resourceName, "kerberos.0.principal", "user@EXAMPLE.COM"),
					resource.TestCheckResourceAttr(resourceName, "kerberos.0.realm", "EXAMPLE.COM"),
					resource.TestCheckResourceAttr(resourceName, "kerberos.0.tgt_generation_method", "KeyTabFilePath"),
					resource.TestCheckResourceAttr(resourceName, "kerberos.0.tgt_key_tab_file_path_spec.0.key_path", "/path/to/keytab"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: acctest.ProjectResourceImportStateIdFunc(resourceName),
			},
		},
	})
}

func TestAccSecretWinRM_kerberos_keytab_org(t *testing.T) {

	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	name := id
	updatedName := fmt.Sprintf("%s_updated", name)
	resourceName := "harness_platform_secret_winrm.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccSecretDestroy(resourceName),
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {},
		},
		Steps: []resource.TestStep{
			{
				Config: testAccResourceSecret_winrm_kerberos_keytab_org(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "port", "5986"),
					resource.TestCheckResourceAttr(resourceName, "kerberos.0.principal", "user@EXAMPLE.COM"),
					resource.TestCheckResourceAttr(resourceName, "kerberos.0.realm", "EXAMPLE.COM"),
					resource.TestCheckResourceAttr(resourceName, "kerberos.0.tgt_generation_method", "KeyTabFilePath"),
					resource.TestCheckResourceAttr(resourceName, "kerberos.0.tgt_key_tab_file_path_spec.0.key_path", "/path/to/keytab"),
				),
			},
			{
				Config: testAccResourceSecret_winrm_kerberos_keytab_org(id, updatedName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "port", "5986"),
					resource.TestCheckResourceAttr(resourceName, "kerberos.0.principal", "user@EXAMPLE.COM"),
					resource.TestCheckResourceAttr(resourceName, "kerberos.0.realm", "EXAMPLE.COM"),
					resource.TestCheckResourceAttr(resourceName, "kerberos.0.tgt_generation_method", "KeyTabFilePath"),
					resource.TestCheckResourceAttr(resourceName, "kerberos.0.tgt_key_tab_file_path_spec.0.key_path", "/path/to/keytab"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: acctest.OrgResourceImportStateIdFunc(resourceName),
			},
		},
	})
}

func testAccResourceSecret_winrm_ntlm(id string, name string) string {
	return fmt.Sprintf(`
		resource "harness_platform_secret_text" "test_password" {
			identifier = "%[1]s_password"
			name = "%[2]s_password"
			description = "test password"
			tags = ["foo:bar"]
			secret_manager_identifier = "harnessSecretManager"
			value_type = "Inline"
			value = "test_password_value"
		}

		resource "harness_platform_secret_winrm" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
			description = "test"
			tags = ["foo:bar"]
			port = 5986
			ntlm {
				domain = "example.com"
				username = "admin"
				password_ref = "account.${harness_platform_secret_text.test_password.id}"
				use_ssl = true
				skip_cert_check = false
				use_no_profile = true
			}
		}
	`, id, name)
}

func testAccResourceSecret_winrm_ntlm_project(id string, name string) string {
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

		resource "harness_platform_secret_text" "test_password" {
			identifier = "%[1]s_password"
			name = "%[2]s_password"
			description = "test password"
			org_id = harness_platform_organization.test.id
			project_id = harness_platform_project.test.id
			tags = ["foo:bar"]
			secret_manager_identifier = "harnessSecretManager"
			value_type = "Inline"
			value = "test_password_value"
		}

		resource "harness_platform_secret_winrm" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
			description = "test"
			org_id = harness_platform_organization.test.id
			project_id = harness_platform_project.test.id
			tags = ["foo:bar"]
			port = 5986
			ntlm {
				domain = "example.com"
				username = "admin"
				password_ref = harness_platform_secret_text.test_password.id
				use_ssl = true
				skip_cert_check = false
				use_no_profile = true
			}
			depends_on = [time_sleep.wait_3_seconds]
		}

		resource "time_sleep" "wait_3_seconds" {
			depends_on = [harness_platform_project.test]
			create_duration = "3s"
		}
	`, id, name)
}

func testAccResourceSecret_winrm_kerberos_password(id string, name string) string {
	return fmt.Sprintf(`
		resource "harness_platform_secret_text" "test_password" {
			identifier = "%[1]s_password"
			name = "%[2]s_password"
			description = "test password"
			tags = ["foo:bar"]
			secret_manager_identifier = "harnessSecretManager"
			value_type = "Inline"
			value = "test_password_value"
		}

		resource "harness_platform_secret_winrm" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
			description = "test"
			tags = ["foo:bar"]
			port = 5986
			kerberos {
				principal = "user@EXAMPLE.COM"
				realm = "EXAMPLE.COM"
				tgt_generation_method = "Password"
				use_ssl = true
				skip_cert_check = false
				use_no_profile = true
				tgt_password_spec {
					password_ref = "account.${harness_platform_secret_text.test_password.id}"
				}
			}
		}
	`, id, name)
}

func testAccResourceSecret_winrm_kerberos_keytab(id string, name string) string {
	return fmt.Sprintf(`
		resource "harness_platform_secret_winrm" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
			description = "test"
			tags = ["foo:bar"]
			port = 5986
			kerberos {
				principal = "user@EXAMPLE.COM"
				realm = "EXAMPLE.COM"
				tgt_generation_method = "KeyTabFilePath"
				use_ssl = true
				skip_cert_check = false
				use_no_profile = true
				tgt_key_tab_file_path_spec {
					key_path = "/path/to/keytab"
				}
			}
		}
	`, id, name)
}

func testAccResourceSecret_winrm_ntlm_org(id string, name string) string {
	return fmt.Sprintf(`
		resource "harness_platform_organization" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
		}

		resource "harness_platform_secret_text" "test_password" {
			identifier = "%[1]s_password"
			name = "%[2]s_password"
			description = "test password"
			org_id = harness_platform_organization.test.id
			tags = ["foo:bar"]
			secret_manager_identifier = "harnessSecretManager"
			value_type = "Inline"
			value = "test_password_value"
			depends_on = [time_sleep.wait_3_seconds]
		}

		resource "time_sleep" "wait_3_seconds" {
			depends_on = [harness_platform_organization.test]
			create_duration = "3s"
		}

		resource "harness_platform_secret_winrm" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
			description = "test"
			org_id = harness_platform_organization.test.id
			tags = ["foo:bar"]
			port = 5986
			ntlm {
				domain = "example.com"
				username = "admin"
				password_ref = "org.${harness_platform_secret_text.test_password.id}"
				use_ssl = true
				skip_cert_check = false
				use_no_profile = true
			}
			depends_on = [time_sleep.wait_4_seconds]
		}

		resource "time_sleep" "wait_4_seconds" {
			depends_on = [harness_platform_secret_text.test_password]
			create_duration = "4s"
		}
	`, id, name)
}

func testAccResourceSecret_winrm_kerberos_password_project(id string, name string) string {
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

		resource "harness_platform_secret_text" "test_password" {
			identifier = "%[1]s_password"
			name = "%[2]s_password"
			description = "test password"
			org_id = harness_platform_organization.test.id
			project_id = harness_platform_project.test.id
			tags = ["foo:bar"]
			secret_manager_identifier = "harnessSecretManager"
			value_type = "Inline"
			value = "test_password_value"
		}

		resource "harness_platform_secret_winrm" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
			description = "test"
			org_id = harness_platform_organization.test.id
			project_id = harness_platform_project.test.id
			tags = ["foo:bar"]
			port = 5986
			kerberos {
				principal = "user@EXAMPLE.COM"
				realm = "EXAMPLE.COM"
				tgt_generation_method = "Password"
				use_ssl = true
				skip_cert_check = false
				use_no_profile = true
				tgt_password_spec {
					password_ref = harness_platform_secret_text.test_password.id
				}
			}
			depends_on = [time_sleep.wait_3_seconds]
		}

		resource "time_sleep" "wait_3_seconds" {
			depends_on = [harness_platform_project.test]
			create_duration = "3s"
		}
	`, id, name)
}

func testAccResourceSecret_winrm_kerberos_password_org(id string, name string) string {
	return fmt.Sprintf(`
		resource "harness_platform_organization" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
		}

		resource "harness_platform_secret_text" "test_password" {
			identifier = "%[1]s_password"
			name = "%[2]s_password"
			description = "test password"
			org_id = harness_platform_organization.test.id
			tags = ["foo:bar"]
			secret_manager_identifier = "harnessSecretManager"
			value_type = "Inline"
			value = "test_password_value"
			depends_on = [time_sleep.wait_3_seconds]
		}

		resource "time_sleep" "wait_3_seconds" {
			depends_on = [harness_platform_organization.test]
			create_duration = "3s"
		}

		resource "harness_platform_secret_winrm" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
			description = "test"
			org_id = harness_platform_organization.test.id
			tags = ["foo:bar"]
			port = 5986
			kerberos {
				principal = "user@EXAMPLE.COM"
				realm = "EXAMPLE.COM"
				tgt_generation_method = "Password"
				use_ssl = true
				skip_cert_check = false
				use_no_profile = true
				tgt_password_spec {
					password_ref = "org.${harness_platform_secret_text.test_password.id}"
				}
			}
			depends_on = [time_sleep.wait_4_seconds]
		}

		resource "time_sleep" "wait_4_seconds" {
			depends_on = [harness_platform_secret_text.test_password]
			create_duration = "4s"
		}
	`, id, name)
}

func testAccResourceSecret_winrm_kerberos_keytab_project(id string, name string) string {
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

		resource "harness_platform_secret_winrm" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
			description = "test"
			org_id = harness_platform_organization.test.id
			project_id = harness_platform_project.test.id
			tags = ["foo:bar"]
			port = 5986
			kerberos {
				principal = "user@EXAMPLE.COM"
				realm = "EXAMPLE.COM"
				tgt_generation_method = "KeyTabFilePath"
				use_ssl = true
				skip_cert_check = false
				use_no_profile = true
				tgt_key_tab_file_path_spec {
					key_path = "/path/to/keytab"
				}
			}
			depends_on = [time_sleep.wait_3_seconds]
		}

		resource "time_sleep" "wait_3_seconds" {
			depends_on = [harness_platform_project.test]
			create_duration = "3s"
		}
	`, id, name)
}

func testAccResourceSecret_winrm_kerberos_keytab_org(id string, name string) string {
	return fmt.Sprintf(`
		resource "harness_platform_organization" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
		}

		resource "harness_platform_secret_winrm" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
			description = "test"
			org_id = harness_platform_organization.test.id
			tags = ["foo:bar"]
			port = 5986
			kerberos {
				principal = "user@EXAMPLE.COM"
				realm = "EXAMPLE.COM"
				tgt_generation_method = "KeyTabFilePath"
				use_ssl = true
				skip_cert_check = false
				use_no_profile = true
				tgt_key_tab_file_path_spec {
					key_path = "/path/to/keytab"
				}
			}
			depends_on = [time_sleep.wait_3_seconds]
		}

		resource "time_sleep" "wait_3_seconds" {
			depends_on = [harness_platform_organization.test]
			create_duration = "3s"
		}
	`, id, name)
}
