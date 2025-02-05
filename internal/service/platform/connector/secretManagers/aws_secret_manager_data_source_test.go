package secretManagers_test

import (
	"fmt"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceConnectorAwsSM_inherit(t *testing.T) {

	var (
		name         = fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))
		resourceName = "data.harness_platform_connector_aws_secret_manager.test"
	)

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceConnectorAwsSM_inherit(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", name),
					resource.TestCheckResourceAttr(resourceName, "identifier", name),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "delegate_selectors.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "secret_name_prefix", "test"),
					resource.TestCheckResourceAttr(resourceName, "credentials.0.inherit_from_delegate", "true"),
				),
			},
		},
	})
}

func TestProjectDataSourceConnectorAwsSM_inherit(t *testing.T) {

	var (
		name         = fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))
		resourceName = "data.harness_platform_connector_aws_secret_manager.test"
	)

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testProjectDataSourceConnectorAwsSM_inherit(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", name),
					resource.TestCheckResourceAttr(resourceName, "identifier", name),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "delegate_selectors.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "secret_name_prefix", "test"),
					resource.TestCheckResourceAttr(resourceName, "credentials.0.inherit_from_delegate", "true"),
				),
			},
		},
	})
}

func TestOrgDataSourceConnectorAwsSM_inherit(t *testing.T) {

	var (
		name         = fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))
		resourceName = "data.harness_platform_connector_aws_secret_manager.test"
	)

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testOrgDataSourceConnectorAwsSM_inherit(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", name),
					resource.TestCheckResourceAttr(resourceName, "identifier", name),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "delegate_selectors.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "default", "false"),
					resource.TestCheckResourceAttr(resourceName, "secret_name_prefix", "test"),
					resource.TestCheckResourceAttr(resourceName, "credentials.0.inherit_from_delegate", "true"),
				),
			},
		},
	})
}

func TestAccDataSourceConnectorAwsSM_manual(t *testing.T) {

	var (
		name         = fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))
		resourceName = "data.harness_platform_connector_aws_secret_manager.test"
	)

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {},
		},
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceConnectorAwsSM_manual(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", name),
					resource.TestCheckResourceAttr(resourceName, "identifier", name),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "delegate_selectors.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "secret_name_prefix", "test"),
					resource.TestCheckResourceAttr(resourceName, "use_put_secret", "false"),
				),
			},
		},
	})
}

func TestAccDataSourceConnectorAwsSM_manualWithUsePutSecretTrue(t *testing.T) {

	var (
		name         = fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))
		resourceName = "data.harness_platform_connector_aws_secret_manager.test"
	)

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {},
		},
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceConnectorAwsSM_manual_withUsePutSecret(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", name),
					resource.TestCheckResourceAttr(resourceName, "identifier", name),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "delegate_selectors.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "secret_name_prefix", "test"),
					resource.TestCheckResourceAttr(resourceName, "use_put_secret", "true"),
				),
			},
		},
	})
}

func TestProjectDataSourceConnectorAwsSM_manual(t *testing.T) {

	var (
		name         = fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))
		resourceName = "data.harness_platform_connector_aws_secret_manager.test"
	)

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {},
		},
		Steps: []resource.TestStep{
			{
				Config: testProjectDataSourceConnectorAwsSM_manual(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", name),
					resource.TestCheckResourceAttr(resourceName, "identifier", name),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "delegate_selectors.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "secret_name_prefix", "test"),
				),
			},
		},
	})
}
func TestOrgDataSourceConnectorAwsSM_manual(t *testing.T) {

	var (
		name         = fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))
		resourceName = "data.harness_platform_connector_aws_secret_manager.test"
	)

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {},
		},
		Steps: []resource.TestStep{
			{
				Config: testOrgDataSourceConnectorAwsSM_manual(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", name),
					resource.TestCheckResourceAttr(resourceName, "identifier", name),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "delegate_selectors.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "secret_name_prefix", "test"),
				),
			},
		},
	})
}

func TestAccDataSourceConnectorAwsSM_assumerole(t *testing.T) {

	var (
		name         = fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))
		resourceName = "data.harness_platform_connector_aws_secret_manager.test"
	)

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceConnectorAwsSM_assumerole(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", name),
					resource.TestCheckResourceAttr(resourceName, "identifier", name),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "delegate_selectors.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "secret_name_prefix", "test"),
					resource.TestCheckResourceAttr(resourceName, "credentials.0.assume_role.0.role_arn", "somerolearn"),
					resource.TestCheckResourceAttr(resourceName, "credentials.0.assume_role.0.external_id", "externalid"),
					resource.TestCheckResourceAttr(resourceName, "credentials.0.assume_role.0.duration", "900"),
				),
			},
		},
	})
}
func TestProjectDataSourceConnectorAwsSM_assumerole(t *testing.T) {

	var (
		name         = fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))
		resourceName = "data.harness_platform_connector_aws_secret_manager.test"
	)

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testProjectDataSourceConnectorAwsSM_assumerole(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", name),
					resource.TestCheckResourceAttr(resourceName, "identifier", name),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "delegate_selectors.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "secret_name_prefix", "test"),
					resource.TestCheckResourceAttr(resourceName, "credentials.0.assume_role.0.role_arn", "somerolearn"),
					resource.TestCheckResourceAttr(resourceName, "credentials.0.assume_role.0.external_id", "externalid"),
					resource.TestCheckResourceAttr(resourceName, "credentials.0.assume_role.0.duration", "900"),
				),
			},
		},
	})
}
func TestOrgDataSourceConnectorAwsSM_assumerole(t *testing.T) {

	var (
		name         = fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))
		resourceName = "data.harness_platform_connector_aws_secret_manager.test"
	)

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testOrgDataSourceConnectorAwsSM_assumerole(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", name),
					resource.TestCheckResourceAttr(resourceName, "identifier", name),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "delegate_selectors.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "secret_name_prefix", "test"),
					resource.TestCheckResourceAttr(resourceName, "credentials.0.assume_role.0.role_arn", "somerolearn"),
					resource.TestCheckResourceAttr(resourceName, "credentials.0.assume_role.0.external_id", "externalid"),
					resource.TestCheckResourceAttr(resourceName, "credentials.0.assume_role.0.duration", "900"),
				),
			},
		},
	})
}

func TestAccDataSourceConnectorAwsSM_oidc_platform(t *testing.T) {

	var (
		name         = fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))
		resourceName = "data.harness_platform_connector_aws_secret_manager.test"
	)

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceConnectorAwsSM_oidc_platform(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", name),
					resource.TestCheckResourceAttr(resourceName, "identifier", name),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "region", "us-east-1"),
					resource.TestCheckResourceAttr(resourceName, "secret_name_prefix", "test"),
					resource.TestCheckResourceAttr(resourceName, "execute_on_delegate", "false"),
					resource.TestCheckResourceAttr(resourceName, "credentials.0.oidc_authentication.0.iam_role_arn", "arn:aws:iam:testarn"),
				),
			},
		},
	})
}

func TestOrgDataSourceConnectorAwsSM_oidc_platform(t *testing.T) {

	var (
		name         = fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))
		resourceName = "data.harness_platform_connector_aws_secret_manager.test"
	)

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testOrgDataSourceConnectorAwsSM_oidc_platform(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", name),
					resource.TestCheckResourceAttr(resourceName, "identifier", name),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "org_id", name),
					resource.TestCheckResourceAttr(resourceName, "region", "us-east-1"),
					resource.TestCheckResourceAttr(resourceName, "secret_name_prefix", "test"),
					resource.TestCheckResourceAttr(resourceName, "execute_on_delegate", "false"),
					resource.TestCheckResourceAttr(resourceName, "credentials.0.oidc_authentication.0.iam_role_arn", "arn:aws:iam:testarn"),
				),
			},
		},
	})
}

func TestProjectDataSourceConnectorAwsSM_oidc_platform(t *testing.T) {

	var (
		name         = fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))
		resourceName = "data.harness_platform_connector_aws_secret_manager.test"
	)

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testProjectDataSourceConnectorAwsSM_oidc_platform(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", name),
					resource.TestCheckResourceAttr(resourceName, "identifier", name),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "org_id", name),
					resource.TestCheckResourceAttr(resourceName, "project_id", name),
					resource.TestCheckResourceAttr(resourceName, "region", "us-east-1"),
					resource.TestCheckResourceAttr(resourceName, "secret_name_prefix", "test"),
					resource.TestCheckResourceAttr(resourceName, "execute_on_delegate", "false"),
					resource.TestCheckResourceAttr(resourceName, "credentials.0.oidc_authentication.0.iam_role_arn", "arn:aws:iam:testarn"),
				),
			},
		},
	})
}

func TestAccDataSourceConnectorAwsSM_oidc_delegate(t *testing.T) {

	var (
		name         = fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))
		resourceName = "data.harness_platform_connector_aws_secret_manager.test"
	)

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceConnectorAwsSM_oidc_delegate(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", name),
					resource.TestCheckResourceAttr(resourceName, "identifier", name),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "region", "us-east-1"),
					resource.TestCheckResourceAttr(resourceName, "secret_name_prefix", "test"),
					resource.TestCheckResourceAttr(resourceName, "execute_on_delegate", "true"),
					resource.TestCheckResourceAttr(resourceName, "delegate_selectors.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "delegate_selectors.0", "harness-delegate"),
					resource.TestCheckResourceAttr(resourceName, "credentials.0.oidc_authentication.0.iam_role_arn", "arn:aws:iam:testarn"),
				),
			},
		},
	})
}

func TestOrgDataSourceConnectorAwsSM_oidc_delegate(t *testing.T) {

	var (
		name         = fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))
		resourceName = "data.harness_platform_connector_aws_secret_manager.test"
	)

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testOrgDataSourceConnectorAwsSM_oidc_delegate(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", name),
					resource.TestCheckResourceAttr(resourceName, "identifier", name),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "org_id", name),
					resource.TestCheckResourceAttr(resourceName, "region", "us-east-1"),
					resource.TestCheckResourceAttr(resourceName, "secret_name_prefix", "test"),
					resource.TestCheckResourceAttr(resourceName, "execute_on_delegate", "true"),
					resource.TestCheckResourceAttr(resourceName, "delegate_selectors.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "delegate_selectors.0", "harness-delegate"),
					resource.TestCheckResourceAttr(resourceName, "credentials.0.oidc_authentication.0.iam_role_arn", "arn:aws:iam:testarn"),
				),
			},
		},
	})
}

func TestProjectDataSourceConnectorAwsSM_oidc_delegate(t *testing.T) {

	var (
		name         = fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))
		resourceName = "data.harness_platform_connector_aws_secret_manager.test"
	)

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testProjectDataSourceConnectorAwsSM_oidc_delegate(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", name),
					resource.TestCheckResourceAttr(resourceName, "identifier", name),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "org_id", name),
					resource.TestCheckResourceAttr(resourceName, "project_id", name),
					resource.TestCheckResourceAttr(resourceName, "region", "us-east-1"),
					resource.TestCheckResourceAttr(resourceName, "secret_name_prefix", "test"),
					resource.TestCheckResourceAttr(resourceName, "execute_on_delegate", "true"),
					resource.TestCheckResourceAttr(resourceName, "delegate_selectors.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "delegate_selectors.0", "harness-delegate"),
					resource.TestCheckResourceAttr(resourceName, "credentials.0.oidc_authentication.0.iam_role_arn", "arn:aws:iam:testarn"),
				),
			},
		},
	})
}

func testAccDataSourceConnectorAwsSM_inherit(name string) string {
	return fmt.Sprintf(`
		resource "harness_platform_connector_aws_secret_manager" "test" {
			identifier = "%[1]s"
			name = "%[1]s"
			description = "test"
			tags = ["foo:bar"]

			secret_name_prefix = "test"
			region = "us-east-1"
			delegate_selectors = ["harness-delegate"]
			credentials {
				inherit_from_delegate = true
			}
		}

		data "harness_platform_connector_aws_secret_manager" "test" {
			identifier = harness_platform_connector_aws_secret_manager.test.identifier
		}
`, name)
}

func testProjectDataSourceConnectorAwsSM_inherit(name string) string {
	return fmt.Sprintf(`
		resource "harness_platform_organization" "test" {
			identifier = "%[1]s"
			name = "%[1]s"
		}

		resource "harness_platform_project" "test" {
			identifier = "%[1]s"
			name = "%[1]s"
			org_id = harness_platform_organization.test.id
			color = "#472848"
		}

		resource "harness_platform_connector_aws_secret_manager" "test" {
			identifier = "%[1]s"
			name = "%[1]s"
			description = "test"
			tags = ["foo:bar"]
			org_id = harness_platform_organization.test.id
			project_id = harness_platform_project.test.id
			secret_name_prefix = "test"
			region = "us-east-1"
			delegate_selectors = ["harness-delegate"]
			credentials {
				inherit_from_delegate = true
			}
		}

		data "harness_platform_connector_aws_secret_manager" "test" {
			identifier = harness_platform_connector_aws_secret_manager.test.identifier
			org_id = harness_platform_connector_aws_secret_manager.test.org_id
			project_id = harness_platform_connector_aws_secret_manager.test.project_id
		}
`, name)
}

func testOrgDataSourceConnectorAwsSM_inherit(name string) string {
	return fmt.Sprintf(`
		resource "harness_platform_organization" "test" {
			identifier = "%[1]s"
			name = "%[1]s"
		}

		resource "harness_platform_connector_aws_secret_manager" "test" {
			identifier = "%[1]s"
			name = "%[1]s"
			description = "test"
			tags = ["foo:bar"]
			org_id = harness_platform_organization.test.id
			secret_name_prefix = "test"
			region = "us-east-1"
			delegate_selectors = ["harness-delegate"]
			default = false
			credentials {
				inherit_from_delegate = true
			}
		}

		data "harness_platform_connector_aws_secret_manager" "test" {
			identifier = harness_platform_connector_aws_secret_manager.test.identifier
			org_id = harness_platform_connector_aws_secret_manager.test.org_id
		}
`, name)
}

func testAccDataSourceConnectorAwsSM_manual(name string) string {
	return fmt.Sprintf(`
		resource "harness_platform_secret_text" "test" {
			identifier = "%[1]s"
			name = "%[1]s"
			description = "test"
			tags = ["foo:bar"]

			secret_manager_identifier = "harnessSecretManager"
			value_type = "Inline"
			value = "secret"
		}

		resource "harness_platform_connector_aws_secret_manager" "test" {
			identifier = "%[1]s"
			name = "%[1]s"
			description = "test"
			tags = ["foo:bar"]

			secret_name_prefix = "test"
			region = "us-east-1"
			delegate_selectors = ["harness-delegate"]
			credentials {
				manual {
					secret_key_ref = "account.${harness_platform_secret_text.test.id}"
					access_key_ref = "account.${harness_platform_secret_text.test.id}"
				}
			}
			depends_on = [time_sleep.wait_4_seconds]
		}

		resource "time_sleep" "wait_4_seconds" {
			depends_on = [harness_platform_secret_text.test]
			destroy_duration = "4s"
		}

		data "harness_platform_connector_aws_secret_manager" "test" {
			identifier = harness_platform_connector_aws_secret_manager.test.identifier
		}
`, name)
}

func testAccDataSourceConnectorAwsSM_manual_withUsePutSecret(name string) string {
	return fmt.Sprintf(`
		resource "harness_platform_secret_text" "test" {
			identifier = "%[1]s"
			name = "%[1]s"
			description = "test"
			tags = ["foo:bar"]

			secret_manager_identifier = "harnessSecretManager"
			value_type = "Inline"
			value = "secret"
		}

		resource "harness_platform_connector_aws_secret_manager" "test" {
			identifier = "%[1]s"
			name = "%[1]s"
			description = "test"
			tags = ["foo:bar"]

			secret_name_prefix = "test"
			region = "us-east-1"
			delegate_selectors = ["harness-delegate"]
			use_put_secret = true
			credentials {
				manual {
					secret_key_ref = "account.${harness_platform_secret_text.test.id}"
					access_key_ref = "account.${harness_platform_secret_text.test.id}"
				}
			}
			depends_on = [time_sleep.wait_4_seconds]
		}

		resource "time_sleep" "wait_4_seconds" {
			depends_on = [harness_platform_secret_text.test]
			destroy_duration = "4s"
		}

		data "harness_platform_connector_aws_secret_manager" "test" {
			identifier = harness_platform_connector_aws_secret_manager.test.identifier
		}
`, name)
}

func testProjectDataSourceConnectorAwsSM_manual(name string) string {
	return fmt.Sprintf(`
		resource "harness_platform_organization" "test" {
			identifier = "%[1]s"
			name = "%[1]s"
		}

		resource "harness_platform_project" "test" {
			identifier = "%[1]s"
			name = "%[1]s"
			org_id = harness_platform_organization.test.id
			color = "#472848"
		}

		resource "harness_platform_secret_text" "test" {
			identifier = "%[1]s"
			name = "%[1]s"
			description = "test"
			tags = ["foo:bar"]
			org_id = harness_platform_organization.test.id
			project_id = harness_platform_project.test.id
			secret_manager_identifier = "harnessSecretManager"
			value_type = "Inline"
			value = "secret"
			depends_on = [time_sleep.wait_4_seconds]
		}

		resource "time_sleep" "wait_4_seconds" {
			depends_on = [harness_platform_project.test]
			create_duration = "4s"
		}

		resource "harness_platform_connector_aws_secret_manager" "test" {
			identifier = "%[1]s"
			name = "%[1]s"
			description = "test"
			tags = ["foo:bar"]
			org_id = harness_platform_organization.test.id
			project_id = harness_platform_project.test.id
			secret_name_prefix = "test"
			region = "us-east-1"
			delegate_selectors = ["harness-delegate"]
			credentials {
				manual {
					secret_key_ref = "${harness_platform_secret_text.test.id}"
					access_key_ref = "${harness_platform_secret_text.test.id}"
				}
			}
			depends_on = [time_sleep.wait_5_seconds]
		}

		resource "time_sleep" "wait_5_seconds" {
			depends_on = [harness_platform_secret_text.test]
			create_duration = "5s"
		}

		data "harness_platform_connector_aws_secret_manager" "test" {
			identifier = harness_platform_connector_aws_secret_manager.test.identifier
			org_id = harness_platform_connector_aws_secret_manager.test.org_id
			project_id = harness_platform_connector_aws_secret_manager.test.project_id
		}
`, name)
}
func testOrgDataSourceConnectorAwsSM_manual(name string) string {
	return fmt.Sprintf(`
		resource "harness_platform_organization" "test" {
			identifier = "%[1]s"
			name = "%[1]s"
		}

		resource "harness_platform_secret_text" "test" {
			identifier = "%[1]s"
			name = "%[1]s"
			description = "test"
			tags = ["foo:bar"]
			org_id = harness_platform_organization.test.id
			secret_manager_identifier = "harnessSecretManager"
			value_type = "Inline"
			value = "secret"
			depends_on = [time_sleep.wait_3_seconds]
		}

		resource "time_sleep" "wait_3_seconds" {
			depends_on = [harness_platform_organization.test]
			create_duration = "3s"
		}

		resource "harness_platform_connector_aws_secret_manager" "test" {
			identifier = "%[1]s"
			name = "%[1]s"
			description = "test"
			tags = ["foo:bar"]
			org_id = harness_platform_organization.test.id
			secret_name_prefix = "test"
			region = "us-east-1"
			delegate_selectors = ["harness-delegate"]
			credentials {
				manual {
					secret_key_ref = "org.${harness_platform_secret_text.test.id}"
					access_key_ref = "org.${harness_platform_secret_text.test.id}"
				}
			}
			depends_on = [time_sleep.wait_4_seconds]
		}

		resource "time_sleep" "wait_4_seconds" {
			depends_on = [harness_platform_secret_text.test]
			create_duration = "4s"
		}

		data "harness_platform_connector_aws_secret_manager" "test" {
			identifier = harness_platform_connector_aws_secret_manager.test.identifier
			org_id = harness_platform_connector_aws_secret_manager.test.org_id
		}
`, name)
}

func testAccDataSourceConnectorAwsSM_assumerole(name string) string {
	return fmt.Sprintf(`
		resource "harness_platform_connector_aws_secret_manager" "test" {
			identifier = "%[1]s"
			name = "%[1]s"
			description = "test"
			tags = ["foo:bar"]

			secret_name_prefix = "test"
			region = "us-east-1"
			delegate_selectors = ["harness-delegate"]
			credentials {
				assume_role {
					role_arn = "somerolearn"
					external_id = "externalid"
					duration = 900
				}
			}
		}

		data "harness_platform_connector_aws_secret_manager" "test" {
			identifier = harness_platform_connector_aws_secret_manager.test.identifier
		}
`, name)
}

func testProjectDataSourceConnectorAwsSM_assumerole(name string) string {
	return fmt.Sprintf(`
		resource "harness_platform_organization" "test" {
			identifier = "%[1]s"
			name = "%[1]s"
		}

		resource "harness_platform_project" "test" {
			identifier = "%[1]s"
			name = "%[1]s"
			org_id = harness_platform_organization.test.id
			color = "#472848"
		}

		resource "harness_platform_connector_aws_secret_manager" "test" {
			identifier = "%[1]s"
			name = "%[1]s"
			description = "test"
			tags = ["foo:bar"]
			org_id = harness_platform_organization.test.id
			project_id = harness_platform_project.test.id
			secret_name_prefix = "test"
			region = "us-east-1"
			delegate_selectors = ["harness-delegate"]
			credentials {
				assume_role {
					role_arn = "somerolearn"
					external_id = "externalid"
					duration = 900
				}
			}
		}

		data "harness_platform_connector_aws_secret_manager" "test" {
			identifier = harness_platform_connector_aws_secret_manager.test.identifier
			org_id = harness_platform_connector_aws_secret_manager.test.org_id
			project_id = harness_platform_connector_aws_secret_manager.test.project_id
		}
`, name)
}

func testOrgDataSourceConnectorAwsSM_assumerole(name string) string {
	return fmt.Sprintf(`
		resource "harness_platform_organization" "test" {
			identifier = "%[1]s"
			name = "%[1]s"
		}

		resource "harness_platform_connector_aws_secret_manager" "test" {
			identifier = "%[1]s"
			name = "%[1]s"
			description = "test"
			tags = ["foo:bar"]
			org_id = harness_platform_organization.test.id
			secret_name_prefix = "test"
			region = "us-east-1"
			delegate_selectors = ["harness-delegate"]
			credentials {
				assume_role {
					role_arn = "somerolearn"
					external_id = "externalid"
					duration = 900
				}
			}
		}

		data "harness_platform_connector_aws_secret_manager" "test" {
			identifier = harness_platform_connector_aws_secret_manager.test.identifier
			org_id = harness_platform_connector_aws_secret_manager.test.org_id
		}
`, name)
}

func testAccDataSourceConnectorAwsSM_oidc_platform(name string) string {
	return fmt.Sprintf(`
		resource "harness_platform_connector_aws_secret_manager" "test" {
			identifier = "%[1]s"
			name = "%[1]s"
			description = "test"
			tags = ["foo:bar"]

			secret_name_prefix = "test"
			region = "us-east-1"
			execute_on_delegate = false

			credentials {
				oidc_authentication {
					iam_role_arn = "arn:aws:iam:testarn"
				}
			}
		}

		data "harness_platform_connector_aws_secret_manager" "test" {
			identifier = harness_platform_connector_aws_secret_manager.test.identifier
		}
`, name)
}

func testOrgDataSourceConnectorAwsSM_oidc_platform(name string) string {
	return fmt.Sprintf(`
		resource "harness_platform_organization" "test" {
			identifier = "%[1]s"
			name = "%[1]s"
		}

		resource "harness_platform_connector_aws_secret_manager" "test" {
			identifier = "%[1]s"
			name = "%[1]s"
			description = "test"
			tags = ["foo:bar"]
			org_id = harness_platform_organization.test.id

			secret_name_prefix = "test"
			region = "us-east-1"
			execute_on_delegate = false

			credentials {
				oidc_authentication {
					iam_role_arn = "arn:aws:iam:testarn"
				}
			}
		}

		data "harness_platform_connector_aws_secret_manager" "test" {
			identifier = harness_platform_connector_aws_secret_manager.test.identifier
			org_id = harness_platform_connector_aws_secret_manager.test.org_id
		}
`, name)
}

func testProjectDataSourceConnectorAwsSM_oidc_platform(name string) string {
	return fmt.Sprintf(`
		resource "harness_platform_organization" "test" {
			identifier = "%[1]s"
			name = "%[1]s"
		}

		resource "harness_platform_project" "test" {
			identifier = "%[1]s"
			name = "%[1]s"
			org_id = harness_platform_organization.test.id
			color = "#472848"
		}

		resource "harness_platform_connector_aws_secret_manager" "test" {
			identifier = "%[1]s"
			name = "%[1]s"
			description = "test"
			tags = ["foo:bar"]
			org_id = harness_platform_organization.test.id
			project_id = harness_platform_project.test.id

			secret_name_prefix = "test"
			region = "us-east-1"
			execute_on_delegate = false

			credentials {
				oidc_authentication {
					iam_role_arn = "arn:aws:iam:testarn"
				}
			}
		}
			
		data "harness_platform_connector_aws_secret_manager" "test" {
			identifier = harness_platform_connector_aws_secret_manager.test.identifier
			org_id = harness_platform_connector_aws_secret_manager.test.org_id
			project_id = harness_platform_connector_aws_secret_manager.test.project_id
		}
`, name)
}

func testAccDataSourceConnectorAwsSM_oidc_delegate(name string) string {
	return fmt.Sprintf(`
		resource "harness_platform_connector_aws_secret_manager" "test" {
			identifier = "%[1]s"
			name = "%[1]s"
			description = "test"
			tags = ["foo:bar"]

			secret_name_prefix = "test"
			region = "us-east-1"
			delegate_selectors = ["harness-delegate"]

			credentials {
				oidc_authentication {
					iam_role_arn = "arn:aws:iam:testarn"
				}
			}
		}

		data "harness_platform_connector_aws_secret_manager" "test" {
			identifier = harness_platform_connector_aws_secret_manager.test.identifier
		}
`, name)
}

func testOrgDataSourceConnectorAwsSM_oidc_delegate(name string) string {
	return fmt.Sprintf(`
		resource "harness_platform_organization" "test" {
			identifier = "%[1]s"
			name = "%[1]s"
		}

		resource "harness_platform_connector_aws_secret_manager" "test" {
			identifier = "%[1]s"
			name = "%[1]s"
			description = "test"
			tags = ["foo:bar"]
			org_id = harness_platform_organization.test.id

			secret_name_prefix = "test"
			region = "us-east-1"
			delegate_selectors = ["harness-delegate"]

			credentials {
				oidc_authentication {
					iam_role_arn = "arn:aws:iam:testarn"
				}
			}
		}

		data "harness_platform_connector_aws_secret_manager" "test" {
			identifier = harness_platform_connector_aws_secret_manager.test.identifier
			org_id = harness_platform_connector_aws_secret_manager.test.org_id
		}
`, name)
}

func testProjectDataSourceConnectorAwsSM_oidc_delegate(name string) string {
	return fmt.Sprintf(`
		resource "harness_platform_organization" "test" {
			identifier = "%[1]s"
			name = "%[1]s"
		}

		resource "harness_platform_project" "test" {
			identifier = "%[1]s"
			name = "%[1]s"
			org_id = harness_platform_organization.test.id
			color = "#472848"
		}

		resource "harness_platform_connector_aws_secret_manager" "test" {
			identifier = "%[1]s"
			name = "%[1]s"
			description = "test"
			tags = ["foo:bar"]
			org_id = harness_platform_organization.test.id
			project_id = harness_platform_project.test.id

			secret_name_prefix = "test"
			region = "us-east-1"
			delegate_selectors = ["harness-delegate"]

			credentials {
				oidc_authentication {
					iam_role_arn = "arn:aws:iam:testarn"
				}
			}
		}

		data "harness_platform_connector_aws_secret_manager" "test" {
			identifier = harness_platform_connector_aws_secret_manager.test.identifier
			org_id = harness_platform_connector_aws_secret_manager.test.org_id
			project_id = harness_platform_connector_aws_secret_manager.test.project_id
		}
`, name)
}
