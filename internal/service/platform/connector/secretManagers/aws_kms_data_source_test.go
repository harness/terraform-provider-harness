package secretManagers_test

import (
	"fmt"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceConnectorAwsKms_inherit(t *testing.T) {

	var (
		name         = fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))
		resourceName = "data.harness_platform_connector_awskms.test"
	)

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {},
		},
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceConnectorAwsKms_inherit(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", name),
					resource.TestCheckResourceAttr(resourceName, "identifier", name),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "arn_ref", "account."+name),
					resource.TestCheckResourceAttr(resourceName, "region", "us-east-1"),
					resource.TestCheckResourceAttr(resourceName, "arn_ref", "account."+name),
					resource.TestCheckResourceAttr(resourceName, "region", "us-east-1"),
					resource.TestCheckResourceAttr(resourceName, "delegate_selectors.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "credentials.0.inherit_from_delegate", "true"),
				),
			},
		},
	})
}

func TestProjectDataSourceConnectorAwsKms_inherit(t *testing.T) {

	var (
		name         = fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))
		resourceName = "data.harness_platform_connector_awskms.test"
	)

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {},
		},
		Steps: []resource.TestStep{
			{
				Config: testProjectDataSourceConnectorAwsKms_inherit(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", name),
					resource.TestCheckResourceAttr(resourceName, "identifier", name),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "arn_ref", name),
					resource.TestCheckResourceAttr(resourceName, "region", "us-east-1"),
					resource.TestCheckResourceAttr(resourceName, "delegate_selectors.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "credentials.0.inherit_from_delegate", "true"),
				),
			},
		},
	})
}

func TestOrgDataSourceConnectorAwsKms_inherit(t *testing.T) {

	var (
		name         = fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))
		resourceName = "data.harness_platform_connector_awskms.test"
	)

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {},
		},
		Steps: []resource.TestStep{
			{
				Config: testOrgDataSourceConnectorAwsKms_inherit(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", name),
					resource.TestCheckResourceAttr(resourceName, "identifier", name),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "arn_ref", "org."+name),
					resource.TestCheckResourceAttr(resourceName, "region", "us-east-1"),
					resource.TestCheckResourceAttr(resourceName, "delegate_selectors.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "credentials.0.inherit_from_delegate", "true"),
				),
			},
		},
	})
}

func TestAccDataSourceConnectorAwsKms_manual_arn_plaintext(t *testing.T) {

	var (
		name         = fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))
		resourceName = "data.harness_platform_connector_awskms.test"
	)

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {},
		},
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceConnectorAwsKms_manual_arn_plaintext(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", name),
					resource.TestCheckResourceAttr(resourceName, "identifier", name),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "arn_plaintext", "plaintext arn"),
					resource.TestCheckResourceAttr(resourceName, "region", "us-east-1"),
					resource.TestCheckResourceAttr(resourceName, "delegate_selectors.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "credentials.0.manual.0.access_key_ref", "account."+name),
					resource.TestCheckResourceAttr(resourceName, "credentials.0.manual.0.secret_key_ref", "account."+name),
				),
			},
		},
	})
}

func TestAccDataSourceConnectorAwsKms_manual(t *testing.T) {

	var (
		name         = fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))
		resourceName = "data.harness_platform_connector_awskms.test"
	)

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {},
		},
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceConnectorAwsKms_manual(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", name),
					resource.TestCheckResourceAttr(resourceName, "identifier", name),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "arn_ref", "account."+name),
					resource.TestCheckResourceAttr(resourceName, "region", "us-east-1"),
					resource.TestCheckResourceAttr(resourceName, "delegate_selectors.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "credentials.0.manual.0.access_key_ref", "account."+name),
					resource.TestCheckResourceAttr(resourceName, "credentials.0.manual.0.secret_key_ref", "account."+name),
				),
			},
		},
	})
}

func TestProjectDataSourceConnectorAwsKms_manual(t *testing.T) {

	var (
		name         = fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))
		resourceName = "data.harness_platform_connector_awskms.test"
	)

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {},
		},
		Steps: []resource.TestStep{
			{
				Config: testProjectDataSourceConnectorAwsKms_manual(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", name),
					resource.TestCheckResourceAttr(resourceName, "identifier", name),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "arn_ref", name),
					resource.TestCheckResourceAttr(resourceName, "region", "us-east-1"),
					resource.TestCheckResourceAttr(resourceName, "delegate_selectors.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "credentials.0.manual.0.access_key_ref", name),
					resource.TestCheckResourceAttr(resourceName, "credentials.0.manual.0.secret_key_ref", name),
				),
			},
		},
	})
}

func TestOrgDataSourceConnectorAwsKms_manual(t *testing.T) {

	var (
		name         = fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))
		resourceName = "data.harness_platform_connector_awskms.test"
	)

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {},
		},
		Steps: []resource.TestStep{
			{
				Config: testOrgDataSourceConnectorAwsKms_manual(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", name),
					resource.TestCheckResourceAttr(resourceName, "identifier", name),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "arn_ref", "org."+name),
					resource.TestCheckResourceAttr(resourceName, "region", "us-east-1"),
					resource.TestCheckResourceAttr(resourceName, "delegate_selectors.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "credentials.0.manual.0.access_key_ref", "org."+name),
					resource.TestCheckResourceAttr(resourceName, "credentials.0.manual.0.secret_key_ref", "org."+name),
				),
			},
		},
	})
}

func TestAccDataSourceConnectorAwsKms_assumerole(t *testing.T) {

	var (
		name         = fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))
		resourceName = "data.harness_platform_connector_awskms.test"
	)

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {},
		},
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceConnectorAwsKms_assumerole(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", name),
					resource.TestCheckResourceAttr(resourceName, "identifier", name),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "arn_ref", "account."+name),
					resource.TestCheckResourceAttr(resourceName, "region", "us-east-1"),
					resource.TestCheckResourceAttr(resourceName, "delegate_selectors.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "credentials.0.assume_role.0.role_arn", "somerolearn"),
					resource.TestCheckResourceAttr(resourceName, "credentials.0.assume_role.0.external_id", "externalid"),
					resource.TestCheckResourceAttr(resourceName, "credentials.0.assume_role.0.duration", "900"),
				),
			},
		},
	})
}

func TestProjectDataSourceConnectorAwsKms_assumerole(t *testing.T) {

	var (
		name         = fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))
		resourceName = "data.harness_platform_connector_awskms.test"
	)

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {},
		},
		Steps: []resource.TestStep{
			{
				Config: testProjectDataSourceConnectorAwsKms_assumerole(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", name),
					resource.TestCheckResourceAttr(resourceName, "identifier", name),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "arn_ref", name),
					resource.TestCheckResourceAttr(resourceName, "region", "us-east-1"),
					resource.TestCheckResourceAttr(resourceName, "delegate_selectors.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "credentials.0.assume_role.0.role_arn", "somerolearn"),
					resource.TestCheckResourceAttr(resourceName, "credentials.0.assume_role.0.external_id", "externalid"),
					resource.TestCheckResourceAttr(resourceName, "credentials.0.assume_role.0.duration", "900"),
				),
			},
		},
	})
}

func TestOrgDataSourceConnectorAwsKms_assumerole(t *testing.T) {

	var (
		name         = fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))
		resourceName = "data.harness_platform_connector_awskms.test"
	)

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {},
		},
		Steps: []resource.TestStep{
			{
				Config: testOrgDataSourceConnectorAwsKms_assumerole(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", name),
					resource.TestCheckResourceAttr(resourceName, "identifier", name),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "arn_ref", "org."+name),
					resource.TestCheckResourceAttr(resourceName, "region", "us-east-1"),
					resource.TestCheckResourceAttr(resourceName, "delegate_selectors.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "credentials.0.assume_role.0.role_arn", "somerolearn"),
					resource.TestCheckResourceAttr(resourceName, "credentials.0.assume_role.0.external_id", "externalid"),
					resource.TestCheckResourceAttr(resourceName, "credentials.0.assume_role.0.duration", "900"),
				),
			},
		},
	})
}

func TestAccDataSourceConnectorAwsKms_oidc_platform(t *testing.T) {

	var (
		name         = fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))
		resourceName = "data.harness_platform_connector_awskms.test"
	)

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {},
		},
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceConnectorAwsKms_oidc_platform(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", name),
					resource.TestCheckResourceAttr(resourceName, "identifier", name),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "arn_ref", "account."+name),
					resource.TestCheckResourceAttr(resourceName, "region", "us-east-1"),
					resource.TestCheckResourceAttr(resourceName, "execute_on_delegate", "false"),
					resource.TestCheckResourceAttr(resourceName, "credentials.0.oidc_authentication.0.iam_role_arn", "somerolearn"),
				),
			},
		},
	})
}

func TestProjectDataSourceConnectorAwsKms_oidc_platform(t *testing.T) {

	var (
		name         = fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))
		resourceName = "data.harness_platform_connector_awskms.test"
	)

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {},
		},
		Steps: []resource.TestStep{
			{
				Config: testProjectDataSourceConnectorAwsKms_oidc_platform(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", name),
					resource.TestCheckResourceAttr(resourceName, "identifier", name),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "arn_ref", name),
					resource.TestCheckResourceAttr(resourceName, "region", "us-east-1"),
					resource.TestCheckResourceAttr(resourceName, "execute_on_delegate", "false"),
					resource.TestCheckResourceAttr(resourceName, "credentials.0.oidc_authentication.0.iam_role_arn", "somerolearn"),
				),
			},
		},
	})
}

func TestOrgDataSourceConnectorAwsKms_oidc_platform(t *testing.T) {

	var (
		name         = fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))
		resourceName = "data.harness_platform_connector_awskms.test"
	)

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {},
		},
		Steps: []resource.TestStep{
			{
				Config: testOrgDataSourceConnectorAwsKms_oidc_platform(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", name),
					resource.TestCheckResourceAttr(resourceName, "identifier", name),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "arn_ref", "org."+name),
					resource.TestCheckResourceAttr(resourceName, "region", "us-east-1"),
					resource.TestCheckResourceAttr(resourceName, "execute_on_delegate", "false"),
					resource.TestCheckResourceAttr(resourceName, "credentials.0.oidc_authentication.0.iam_role_arn", "somerolearn"),
				),
			},
		},
	})
}

func TestAccDataSourceConnectorAwsKms_oidc_delegate(t *testing.T) {

	var (
		name         = fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))
		resourceName = "data.harness_platform_connector_awskms.test"
	)

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {},
		},
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceConnectorAwsKms_oidc_delegate(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", name),
					resource.TestCheckResourceAttr(resourceName, "identifier", name),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "arn_ref", "account."+name),
					resource.TestCheckResourceAttr(resourceName, "region", "us-east-1"),
					resource.TestCheckResourceAttr(resourceName, "delegate_selectors.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "credentials.0.oidc_authentication.0.iam_role_arn", "somerolearn"),
				),
			},
		},
	})
}

func TestProjectDataSourceConnectorAwsKms_oidc_delegate(t *testing.T) {

	var (
		name         = fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))
		resourceName = "data.harness_platform_connector_awskms.test"
	)

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {},
		},
		Steps: []resource.TestStep{
			{
				Config: testProjectDataSourceConnectorAwsKms_oidc_delegate(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", name),
					resource.TestCheckResourceAttr(resourceName, "identifier", name),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "arn_ref", name),
					resource.TestCheckResourceAttr(resourceName, "region", "us-east-1"),
					resource.TestCheckResourceAttr(resourceName, "delegate_selectors.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "credentials.0.oidc_authentication.0.iam_role_arn", "somerolearn"),
				),
			},
		},
	})
}

func TestOrgDataSourceConnectorAwsKms_oidc_delegate(t *testing.T) {

	var (
		name         = fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))
		resourceName = "data.harness_platform_connector_awskms.test"
	)

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {},
		},
		Steps: []resource.TestStep{
			{
				Config: testOrgDataSourceConnectorAwsKms_oidc_delegate(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", name),
					resource.TestCheckResourceAttr(resourceName, "identifier", name),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "arn_ref", "org."+name),
					resource.TestCheckResourceAttr(resourceName, "region", "us-east-1"),
					resource.TestCheckResourceAttr(resourceName, "delegate_selectors.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "credentials.0.oidc_authentication.0.iam_role_arn", "somerolearn"),
				),
			},
		},
	})
}

func testAccDataSourceConnectorAwsKms_inherit(name string) string {
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

		resource "harness_platform_connector_awskms" "test" {
			identifier = "%[1]s"
			name = "%[1]s"
			description = "test"
			tags = ["foo:bar"]

			arn_ref = "account.${harness_platform_secret_text.test.id}"
			region = "us-east-1"
			delegate_selectors = ["harness-delegate"]
			credentials {
				inherit_from_delegate = true
			}
			depends_on = [time_sleep.wait_4_seconds]
		}

		resource "time_sleep" "wait_4_seconds" {
			depends_on = [harness_platform_secret_text.test]
			destroy_duration = "4s"
		}

		data "harness_platform_connector_awskms" "test" {
			identifier = harness_platform_connector_awskms.test.identifier
		}
`, name)
}

func testProjectDataSourceConnectorAwsKms_inherit(name string) string {
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
		org_id=harness_platform_organization.test.id
		project_id=harness_platform_project.test.id
		secret_manager_identifier = "harnessSecretManager"
		value_type = "Inline"
		value = "secret"
		depends_on = [time_sleep.wait_4_seconds]
	}

	resource "time_sleep" "wait_4_seconds" {
		depends_on = [harness_platform_project.test]
		create_duration = "4s" 
	}

	resource "harness_platform_connector_awskms" "test" {
		identifier = "%[1]s"
		name = "%[1]s"
		description = "test"
		tags = ["foo:bar"]
		org_id=harness_platform_organization.test.id
		project_id=harness_platform_project.test.id
		arn_ref = "${harness_platform_secret_text.test.id}"
		region = "us-east-1"
		delegate_selectors = ["harness-delegate"]
		credentials {
			inherit_from_delegate = true
		}
		depends_on = [time_sleep.wait_3_seconds]
	}

	resource "time_sleep" "wait_3_seconds" {
		depends_on = [harness_platform_secret_text.test]
		create_duration = "3s" 
	}

		data "harness_platform_connector_awskms" "test" {
			identifier = harness_platform_connector_awskms.test.identifier
			org_id = harness_platform_organization.test.id
			project_id = harness_platform_project.test.id
		}
`, name)
}

func testOrgDataSourceConnectorAwsKms_inherit(name string) string {
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
		org_id=harness_platform_organization.test.id
		secret_manager_identifier = "harnessSecretManager"
		value_type = "Inline"
		value = "secret"
		depends_on = [time_sleep.wait_3_seconds]
	}

	resource "time_sleep" "wait_3_seconds" {
		depends_on = [harness_platform_organization.test]
		create_duration = "3s" 
	}

		resource "harness_platform_connector_awskms" "test" {
			identifier = "%[1]s"
			name = "%[1]s"
			description = "test"
			tags = ["foo:bar"]
			org_id=harness_platform_organization.test.id
			arn_ref = "org.${harness_platform_secret_text.test.id}"
			region = "us-east-1"
			delegate_selectors = ["harness-delegate"]
			credentials {
				inherit_from_delegate = true
			}
			depends_on = [time_sleep.wait_4_seconds]
	}

	resource "time_sleep" "wait_4_seconds" {
		depends_on = [harness_platform_secret_text.test]
		create_duration = "4s" 
	}

	data "harness_platform_connector_awskms" "test" {
			identifier = harness_platform_connector_awskms.test.identifier
			org_id = harness_platform_organization.test.id
		}
`, name)
}

func testAccDataSourceConnectorAwsKms_manual(name string) string {
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

		resource "harness_platform_connector_awskms" "test" {
			identifier = "%[1]s"
			name = "%[1]s"
			description = "test"
			tags = ["foo:bar"]

			
			arn_ref = "account.${harness_platform_secret_text.test.id}"
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
			create_duration = "4s"
		}

		data "harness_platform_connector_awskms" "test" {
			identifier = harness_platform_connector_awskms.test.identifier
		}
`, name)
}

func testAccDataSourceConnectorAwsKms_manual_arn_plaintext(name string) string {
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
		resource "harness_platform_connector_awskms" "test" {
			identifier = "%[1]s"
			name = "%[1]s"
			description = "test"
			tags = ["foo:bar"]
			
			arn_plaintext = "plaintext arn"
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
			create_duration = "4s"
		}
		data "harness_platform_connector_awskms" "test" {
			identifier = harness_platform_connector_awskms.test.identifier
		}
`, name)
}

func testProjectDataSourceConnectorAwsKms_manual(name string) string {
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
		org_id=harness_platform_organization.test.id
		project_id=harness_platform_project.test.id
		secret_manager_identifier = "harnessSecretManager"
		value_type = "Inline"
		value = "secret"
		depends_on = [time_sleep.wait_3_seconds]
		}

		resource "time_sleep" "wait_3_seconds" {
			depends_on = [harness_platform_project.test]
			create_duration = "3s"
		}

		resource "harness_platform_connector_awskms" "test" {
			identifier = "%[1]s"
			name = "%[1]s"
			description = "test"
			tags = ["foo:bar"]
			org_id=harness_platform_organization.test.id
			project_id=harness_platform_project.test.id
			arn_ref = "${harness_platform_secret_text.test.id}"
			region = "us-east-1"
			delegate_selectors = ["harness-delegate"]
			credentials {
				manual {
					secret_key_ref = "${harness_platform_secret_text.test.id}"
					access_key_ref = "${harness_platform_secret_text.test.id}"
				}
			}
			depends_on = [time_sleep.wait_4_seconds]
		}

		resource "time_sleep" "wait_4_seconds" {
			depends_on = [harness_platform_secret_text.test]
			create_duration = "4s"
		}

		data "harness_platform_connector_awskms" "test" {
			identifier = harness_platform_connector_awskms.test.identifier
			org_id = harness_platform_organization.test.id
			project_id = harness_platform_project.test.id
		}
`, name)
}

func testOrgDataSourceConnectorAwsKms_manual(name string) string {
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
		org_id=harness_platform_organization.test.id
		
		secret_manager_identifier = "harnessSecretManager"
		value_type = "Inline"
		value = "secret"
		depends_on = [time_sleep.wait_3_seconds]
		}

		resource "time_sleep" "wait_3_seconds" {
			depends_on = [harness_platform_organization.test]
			create_duration = "3s"
		}

		resource "harness_platform_connector_awskms" "test" {
			identifier = "%[1]s"
			name = "%[1]s"
			description = "test"
			tags = ["foo:bar"]
			org_id=harness_platform_organization.test.id
			arn_ref = "org.${harness_platform_secret_text.test.id}"
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

		data "harness_platform_connector_awskms" "test" {
			identifier = harness_platform_connector_awskms.test.identifier
			org_id = harness_platform_organization.test.id
		}
`, name)
}

func testAccDataSourceConnectorAwsKms_assumerole(name string) string {
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

		resource "harness_platform_connector_awskms" "test" {
			identifier = "%[1]s"
			name = "%[1]s"
			description = "test"
			tags = ["foo:bar"]

			arn_ref = "account.${harness_platform_secret_text.test.id}"
			region = "us-east-1"
			delegate_selectors = ["harness-delegate"]
			credentials {
				assume_role {
					role_arn = "somerolearn"
					external_id = "externalid"
					duration = 900
				}
			}
			depends_on = [time_sleep.wait_4_seconds]
		}

		resource "time_sleep" "wait_4_seconds" {
			depends_on = [harness_platform_secret_text.test]
			destroy_duration = "4s"
		}

		data "harness_platform_connector_awskms" "test" {
			identifier = harness_platform_connector_awskms.test.identifier
		}
`, name)
}

func testProjectDataSourceConnectorAwsKms_assumerole(name string) string {
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
		org_id= harness_platform_organization.test.id
		project_id= harness_platform_project.test.id
		secret_manager_identifier = "harnessSecretManager"
		value_type = "Inline"
		value = "secret"
		depends_on = [time_sleep.wait_3_seconds]
		}

		resource "time_sleep" "wait_3_seconds" {
			depends_on = [harness_platform_project.test]
			create_duration = "3s"
		}

		resource "harness_platform_connector_awskms" "test" {
			identifier = "%[1]s"
			name = "%[1]s"
			description = "test"
			tags = ["foo:bar"]
			org_id = harness_platform_organization.test.id
			project_id = harness_platform_project.test.id
			arn_ref = "${harness_platform_secret_text.test.id}"
			region = "us-east-1"
			delegate_selectors = ["harness-delegate"]
			credentials {
				assume_role {
					role_arn = "somerolearn"
					external_id = "externalid"
					duration = 900
				}
			}
			depends_on = [time_sleep.wait_4_seconds]
		}

		resource "time_sleep" "wait_4_seconds" {
			depends_on = [harness_platform_secret_text.test]
			destroy_duration = "4s"
		}

		data "harness_platform_connector_awskms" "test" {
			identifier = harness_platform_connector_awskms.test.identifier
			org_id = harness_platform_organization.test.id
			project_id = harness_platform_project.test.id
		}
`, name)
}

func testOrgDataSourceConnectorAwsKms_assumerole(name string) string {
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
		org_id=harness_platform_organization.test.id
		secret_manager_identifier = "harnessSecretManager"
		value_type = "Inline"
		value = "secret"
		depends_on = [time_sleep.wait_3_seconds]
		}

		resource "time_sleep" "wait_3_seconds" {
			depends_on = [harness_platform_organization.test]
			create_duration = "3s"
		}

		resource "harness_platform_connector_awskms" "test" {
			identifier = "%[1]s"
			name = "%[1]s"
			description = "test"
			tags = ["foo:bar"]
			org_id=harness_platform_organization.test.id
			arn_ref = "org.${harness_platform_secret_text.test.id}"
			region = "us-east-1"
			delegate_selectors = ["harness-delegate"]
			credentials {
				assume_role {
					role_arn = "somerolearn"
					external_id = "externalid"
					duration = 900
				}
			}
			depends_on = [time_sleep.wait_4_seconds]
		}

		resource "time_sleep" "wait_4_seconds" {
			depends_on = [harness_platform_secret_text.test]
			create_duration = "4s"
		}

		data "harness_platform_connector_awskms" "test" {
			identifier = harness_platform_connector_awskms.test.identifier
			org_id = harness_platform_organization.test.id
		}
`, name)
}

func testAccDataSourceConnectorAwsKms_oidc_platform(name string) string {
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

		resource "harness_platform_connector_awskms" "test" {
			identifier = "%[1]s"
			name = "%[1]s"
			description = "test"
			tags = ["foo:bar"]

			arn_ref = "account.${harness_platform_secret_text.test.id}"
			region = "us-east-1"
			execute_on_delegate = false
			credentials {
				oidc_authentication {
					iam_role_arn = "somerolearn"
				}
			}
			depends_on = [time_sleep.wait_4_seconds]
		}

		resource "time_sleep" "wait_4_seconds" {
			depends_on = [harness_platform_secret_text.test]
			destroy_duration = "4s"
		}

		data "harness_platform_connector_awskms" "test" {
			identifier = harness_platform_connector_awskms.test.identifier
		}
`, name)
}

func testProjectDataSourceConnectorAwsKms_oidc_platform(name string) string {
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
		org_id= harness_platform_organization.test.id
		project_id= harness_platform_project.test.id
		secret_manager_identifier = "harnessSecretManager"
		value_type = "Inline"
		value = "secret"
		depends_on = [time_sleep.wait_3_seconds]
		}

		resource "time_sleep" "wait_3_seconds" {
			depends_on = [harness_platform_project.test]
			create_duration = "3s"
		}

		resource "harness_platform_connector_awskms" "test" {
			identifier = "%[1]s"
			name = "%[1]s"
			description = "test"
			tags = ["foo:bar"]
			org_id = harness_platform_organization.test.id
			project_id = harness_platform_project.test.id
			arn_ref = "${harness_platform_secret_text.test.id}"
			region = "us-east-1"
			execute_on_delegate = false
			credentials {
				oidc_authentication {
					iam_role_arn = "somerolearn"
				}
			}
			depends_on = [time_sleep.wait_4_seconds]
		}

		resource "time_sleep" "wait_4_seconds" {
			depends_on = [harness_platform_secret_text.test]
			destroy_duration = "4s"
		}
		
		data "harness_platform_connector_awskms" "test" {
			identifier = harness_platform_connector_awskms.test.identifier
			org_id = harness_platform_organization.test.id
			project_id = harness_platform_project.test.id
		}
`, name)
}

func testOrgDataSourceConnectorAwsKms_oidc_platform(name string) string {
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
		org_id=harness_platform_organization.test.id
		secret_manager_identifier = "harnessSecretManager"
		value_type = "Inline"
		value = "secret"
		depends_on = [time_sleep.wait_3_seconds]
		}

		resource "time_sleep" "wait_3_seconds" {
			depends_on = [harness_platform_organization.test]
			create_duration = "3s"
		}

		resource "harness_platform_connector_awskms" "test" {
			identifier = "%[1]s"
			name = "%[1]s"
			description = "test"
			tags = ["foo:bar"]
			org_id=harness_platform_organization.test.id
			arn_ref = "org.${harness_platform_secret_text.test.id}"
			region = "us-east-1"
			execute_on_delegate = false
			credentials {
				oidc_authentication {
					iam_role_arn = "somerolearn"
				}
			}
			depends_on = [time_sleep.wait_4_seconds]
		}

		resource "time_sleep" "wait_4_seconds" {
			depends_on = [harness_platform_secret_text.test]
			create_duration = "4s"
		}

		data "harness_platform_connector_awskms" "test" {
			identifier = harness_platform_connector_awskms.test.identifier
			org_id = harness_platform_organization.test.id
		}
`, name)
}

func testAccDataSourceConnectorAwsKms_oidc_delegate(name string) string {
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

		resource "harness_platform_connector_awskms" "test" {
			identifier = "%[1]s"
			name = "%[1]s"
			description = "test"
			tags = ["foo:bar"]

			arn_ref = "account.${harness_platform_secret_text.test.id}"
			region = "us-east-1"
			delegate_selectors = ["harness-delegate"]
			credentials {
				oidc_authentication {
					iam_role_arn = "somerolearn"
				}
			}
			depends_on = [time_sleep.wait_4_seconds]
		}

		resource "time_sleep" "wait_4_seconds" {
			depends_on = [harness_platform_secret_text.test]
			destroy_duration = "4s"
		}

		data "harness_platform_connector_awskms" "test" {
			identifier = harness_platform_connector_awskms.test.identifier
		}
`, name)
}

func testProjectDataSourceConnectorAwsKms_oidc_delegate(name string) string {
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
		org_id= harness_platform_organization.test.id
		project_id= harness_platform_project.test.id
		secret_manager_identifier = "harnessSecretManager"
		value_type = "Inline"
		value = "secret"
		depends_on = [time_sleep.wait_3_seconds]
		}

		resource "time_sleep" "wait_3_seconds" {
			depends_on = [harness_platform_project.test]
			create_duration = "3s"
		}

		resource "harness_platform_connector_awskms" "test" {
			identifier = "%[1]s"
			name = "%[1]s"
			description = "test"
			tags = ["foo:bar"]
			org_id = harness_platform_organization.test.id
			project_id = harness_platform_project.test.id
			arn_ref = "${harness_platform_secret_text.test.id}"
			region = "us-east-1"
			delegate_selectors = ["harness-delegate"]
			credentials {
				oidc_authentication {
					iam_role_arn = "somerolearn"
				}
			}
			depends_on = [time_sleep.wait_4_seconds]
		}

		resource "time_sleep" "wait_4_seconds" {
			depends_on = [harness_platform_secret_text.test]
			destroy_duration = "4s"
		}

		data "harness_platform_connector_awskms" "test" {
			identifier = harness_platform_connector_awskms.test.identifier
			org_id = harness_platform_organization.test.id
			project_id = harness_platform_project.test.id
		}
`, name)
}

func testOrgDataSourceConnectorAwsKms_oidc_delegate(name string) string {
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
		org_id=harness_platform_organization.test.id
		secret_manager_identifier = "harnessSecretManager"
		value_type = "Inline"
		value = "secret"
		depends_on = [time_sleep.wait_3_seconds]
		}

		resource "time_sleep" "wait_3_seconds" {
			depends_on = [harness_platform_organization.test]
			create_duration = "3s"
		}

		resource "harness_platform_connector_awskms" "test" {
			identifier = "%[1]s"
			name = "%[1]s"
			description = "test"
			tags = ["foo:bar"]
			org_id=harness_platform_organization.test.id
			arn_ref = "org.${harness_platform_secret_text.test.id}"
			region = "us-east-1"
			delegate_selectors = ["harness-delegate"]
			credentials {
				oidc_authentication {
					iam_role_arn = "somerolearn"
				}
			}
			depends_on = [time_sleep.wait_4_seconds]
		}

		resource "time_sleep" "wait_4_seconds" {
			depends_on = [harness_platform_secret_text.test]
			create_duration = "4s"
		}

		data "harness_platform_connector_awskms" "test" {
			identifier = harness_platform_connector_awskms.test.identifier
			org_id = harness_platform_organization.test.id
		}
`, name)
}
