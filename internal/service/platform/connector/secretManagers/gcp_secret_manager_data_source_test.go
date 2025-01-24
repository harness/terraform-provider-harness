package secretManagers_test

import (
	"fmt"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceConnectorGcpSM_manual(t *testing.T) {

	var (
		name         = fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))
		resourceName = "data.harness_platform_connector_gcp_secret_manager.test"
	)

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {},
		},
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceConnectorGcpSM_manual(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", name),
					resource.TestCheckResourceAttr(resourceName, "identifier", name),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "default", "false"),
					resource.TestCheckResourceAttr(resourceName, "execute_on_delegate", "true"),
					resource.TestCheckResourceAttr(resourceName, "manual.0.secret_key_ref", "account."+name),
					resource.TestCheckResourceAttr(resourceName, "manual.0.delegate_selectors.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "manual.0.delegate_selectors.0", "harness-delegate"),
				),
			},
		},
	})
}

func TestOrgDataSourceConnectorGcpSM_manual(t *testing.T) {
	var (
		name         = fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))
		resourceName = "data.harness_platform_connector_gcp_secret_manager.test"
	)

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {},
		},
		Steps: []resource.TestStep{
			{
				Config: testOrgDataSourceConnectorGcpSM_manual(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", name),
					resource.TestCheckResourceAttr(resourceName, "identifier", name),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "org_id", name),
					resource.TestCheckResourceAttr(resourceName, "default", "false"),
					resource.TestCheckResourceAttr(resourceName, "execute_on_delegate", "true"),
					resource.TestCheckResourceAttr(resourceName, "manual.0.secret_key_ref", "org."+name),
					resource.TestCheckResourceAttr(resourceName, "manual.0.delegate_selectors.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "manual.0.delegate_selectors.0", "harness-delegate"),
				),
			},
		},
	})
}

func TestProjectDataSourceConnectorGcpSM_manual(t *testing.T) {
	var (
		name         = fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))
		resourceName = "data.harness_platform_connector_gcp_secret_manager.test"
	)

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {},
		},
		Steps: []resource.TestStep{
			{
				Config: testProjectDataSourceConnectorGcpSM_manual(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", name),
					resource.TestCheckResourceAttr(resourceName, "identifier", name),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "org_id", name),
					resource.TestCheckResourceAttr(resourceName, "project_id", name),
					resource.TestCheckResourceAttr(resourceName, "default", "false"),
					resource.TestCheckResourceAttr(resourceName, "execute_on_delegate", "true"),
					resource.TestCheckResourceAttr(resourceName, "manual.0.secret_key_ref", name),
					resource.TestCheckResourceAttr(resourceName, "manual.0.delegate_selectors.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "manual.0.delegate_selectors.0", "harness-delegate"),
				),
			},
		},
	})
}

func TestAccDataSourceConnectorGcpSM_inherit(t *testing.T) {
	var (
		name         = fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))
		resourceName = "data.harness_platform_connector_gcp_secret_manager.test"
	)

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {},
		},
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceConnectorGcpSM_inherit(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", name),
					resource.TestCheckResourceAttr(resourceName, "identifier", name),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "default", "false"),
					resource.TestCheckResourceAttr(resourceName, "execute_on_delegate", "true"),
					resource.TestCheckResourceAttr(resourceName, "inherit_from_delegate.0.delegate_selectors.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "inherit_from_delegate.0.delegate_selectors.0", "harness-delegate"),
				),
			},
		},
	})
}

func TestOrgDataSourceConnectorGcpSM_inherit(t *testing.T) {
	var (
		name         = fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))
		resourceName = "data.harness_platform_connector_gcp_secret_manager.test"
	)

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {},
		},
		Steps: []resource.TestStep{
			{
				Config: testOrgDataSourceConnectorGcpSM_inherit(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", name),
					resource.TestCheckResourceAttr(resourceName, "identifier", name),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "org_id", name),
					resource.TestCheckResourceAttr(resourceName, "default", "false"),
					resource.TestCheckResourceAttr(resourceName, "execute_on_delegate", "true"),
					resource.TestCheckResourceAttr(resourceName, "inherit_from_delegate.0.delegate_selectors.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "inherit_from_delegate.0.delegate_selectors.0", "harness-delegate"),
				),
			},
		},
	})
}

func TestProjectDataSourceConnectorGcpSM_inherit(t *testing.T) {
	var (
		name         = fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))
		resourceName = "data.harness_platform_connector_gcp_secret_manager.test"
	)

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {},
		},
		Steps: []resource.TestStep{
			{
				Config: testProjectDataSourceConnectorGcpSM_inherit(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", name),
					resource.TestCheckResourceAttr(resourceName, "identifier", name),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "org_id", name),
					resource.TestCheckResourceAttr(resourceName, "project_id", name),
					resource.TestCheckResourceAttr(resourceName, "default", "false"),
					resource.TestCheckResourceAttr(resourceName, "execute_on_delegate", "true"),
					resource.TestCheckResourceAttr(resourceName, "inherit_from_delegate.0.delegate_selectors.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "inherit_from_delegate.0.delegate_selectors.0", "harness-delegate"),
				),
			},
		},
	})
}

func TestAccDataSourceConnectorGcpSM_oidc_platform(t *testing.T) {
	var (
		name         = fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))
		resourceName = "data.harness_platform_connector_gcp_secret_manager.test"
	)

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {},
		},
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceConnectorGcpSM_oidc_platform(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", name),
					resource.TestCheckResourceAttr(resourceName, "identifier", name),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "default", "false"),
					resource.TestCheckResourceAttr(resourceName, "execute_on_delegate", "false"),
					resource.TestCheckResourceAttr(resourceName, "oidc_authentication.0.workload_pool_id", "harness-pool-test"),
					resource.TestCheckResourceAttr(resourceName, "oidc_authentication.0.provider_id", "harness"),
					resource.TestCheckResourceAttr(resourceName, "oidc_authentication.0.gcp_project_id", "1234567"),
					resource.TestCheckResourceAttr(resourceName, "oidc_authentication.0.service_account_email", "harness.sample@iam.gserviceaccount.com"),
				),
			},
		},
	})
}

func TestOrgDataSourceConnectorGcpSM_oidc_platform(t *testing.T) {
	var (
		name         = fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))
		resourceName = "data.harness_platform_connector_gcp_secret_manager.test"
	)

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {},
		},
		Steps: []resource.TestStep{
			{
				Config: testOrgDataSourceConnectorGcpSM_oidc_platform(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", name),
					resource.TestCheckResourceAttr(resourceName, "identifier", name),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "org_id", name),
					resource.TestCheckResourceAttr(resourceName, "default", "false"),
					resource.TestCheckResourceAttr(resourceName, "execute_on_delegate", "false"),
					resource.TestCheckResourceAttr(resourceName, "oidc_authentication.0.workload_pool_id", "harness-pool-test"),
					resource.TestCheckResourceAttr(resourceName, "oidc_authentication.0.provider_id", "harness"),
					resource.TestCheckResourceAttr(resourceName, "oidc_authentication.0.gcp_project_id", "1234567"),
					resource.TestCheckResourceAttr(resourceName, "oidc_authentication.0.service_account_email", "harness.sample@iam.gserviceaccount.com"),
				),
			},
		},
	})
}

func TestProjectDataSourceConnectorGcpSM_oidc_platform(t *testing.T) {
	var (
		name         = fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))
		resourceName = "data.harness_platform_connector_gcp_secret_manager.test"
	)

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {},
		},
		Steps: []resource.TestStep{
			{
				Config: testProjectDataSourceConnectorGcpSM_oidc_platform(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", name),
					resource.TestCheckResourceAttr(resourceName, "identifier", name),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "org_id", name),
					resource.TestCheckResourceAttr(resourceName, "project_id", name),
					resource.TestCheckResourceAttr(resourceName, "default", "false"),
					resource.TestCheckResourceAttr(resourceName, "execute_on_delegate", "false"),
					resource.TestCheckResourceAttr(resourceName, "oidc_authentication.0.workload_pool_id", "harness-pool-test"),
					resource.TestCheckResourceAttr(resourceName, "oidc_authentication.0.provider_id", "harness"),
					resource.TestCheckResourceAttr(resourceName, "oidc_authentication.0.gcp_project_id", "1234567"),
					resource.TestCheckResourceAttr(resourceName, "oidc_authentication.0.service_account_email", "harness.sample@iam.gserviceaccount.com"),
				),
			},
		},
	})
}

func TestAccDataSourceConnectorGcpSM_oidc_delegate(t *testing.T) {
	var (
		name         = fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))
		resourceName = "data.harness_platform_connector_gcp_secret_manager.test"
	)

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {},
		},
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceConnectorGcpSM_oidc_delegate(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", name),
					resource.TestCheckResourceAttr(resourceName, "identifier", name),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "default", "false"),
					resource.TestCheckResourceAttr(resourceName, "execute_on_delegate", "true"),
					resource.TestCheckResourceAttr(resourceName, "oidc_authentication.0.delegate_selectors.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "oidc_authentication.0.delegate_selectors.0", "harness-delegate"),
					resource.TestCheckResourceAttr(resourceName, "oidc_authentication.0.workload_pool_id", "harness-pool-test"),
					resource.TestCheckResourceAttr(resourceName, "oidc_authentication.0.provider_id", "harness"),
					resource.TestCheckResourceAttr(resourceName, "oidc_authentication.0.gcp_project_id", "1234567"),
					resource.TestCheckResourceAttr(resourceName, "oidc_authentication.0.service_account_email", "harness.sample@iam.gserviceaccount.com"),
				),
			},
		},
	})
}

func TestOrgDataSourceConnectorGcpSM_oidc_delegate(t *testing.T) {
	var (
		name         = fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))
		resourceName = "data.harness_platform_connector_gcp_secret_manager.test"
	)

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {},
		},
		Steps: []resource.TestStep{
			{
				Config: testOrgDataSourceConnectorGcpSM_oidc_delegate(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", name),
					resource.TestCheckResourceAttr(resourceName, "identifier", name),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "org_id", name),
					resource.TestCheckResourceAttr(resourceName, "default", "false"),
					resource.TestCheckResourceAttr(resourceName, "execute_on_delegate", "true"),
					resource.TestCheckResourceAttr(resourceName, "oidc_authentication.0.delegate_selectors.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "oidc_authentication.0.delegate_selectors.0", "harness-delegate"),
					resource.TestCheckResourceAttr(resourceName, "oidc_authentication.0.workload_pool_id", "harness-pool-test"),
					resource.TestCheckResourceAttr(resourceName, "oidc_authentication.0.provider_id", "harness"),
					resource.TestCheckResourceAttr(resourceName, "oidc_authentication.0.gcp_project_id", "1234567"),
					resource.TestCheckResourceAttr(resourceName, "oidc_authentication.0.service_account_email", "harness.sample@iam.gserviceaccount.com"),
				),
			},
		},
	})
}

func TestProjectDataSourceConnectorGcpSM_oidc_delegate(t *testing.T) {
	var (
		name         = fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))
		resourceName = "data.harness_platform_connector_gcp_secret_manager.test"
	)

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {},
		},
		Steps: []resource.TestStep{
			{
				Config: testProjectDataSourceConnectorGcpSM_oidc_delegate(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", name),
					resource.TestCheckResourceAttr(resourceName, "identifier", name),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "org_id", name),
					resource.TestCheckResourceAttr(resourceName, "project_id", name),
					resource.TestCheckResourceAttr(resourceName, "default", "false"),
					resource.TestCheckResourceAttr(resourceName, "execute_on_delegate", "true"),
					resource.TestCheckResourceAttr(resourceName, "oidc_authentication.0.delegate_selectors.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "oidc_authentication.0.delegate_selectors.0", "harness-delegate"),
					resource.TestCheckResourceAttr(resourceName, "oidc_authentication.0.workload_pool_id", "harness-pool-test"),
					resource.TestCheckResourceAttr(resourceName, "oidc_authentication.0.provider_id", "harness"),
					resource.TestCheckResourceAttr(resourceName, "oidc_authentication.0.gcp_project_id", "1234567"),
					resource.TestCheckResourceAttr(resourceName, "oidc_authentication.0.service_account_email", "harness.sample@iam.gserviceaccount.com"),
				),
			},
		},
	})
}

func TestAccDataSourceConnectorGcpSM_manual_default(t *testing.T) {
	var (
		name         = fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))
		resourceName = "data.harness_platform_connector_gcp_secret_manager.test"
	)

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {},
		},
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceConnectorGcpSM_manual_default(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", name),
					resource.TestCheckResourceAttr(resourceName, "identifier", name),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "default", "true"),
					resource.TestCheckResourceAttr(resourceName, "execute_on_delegate", "true"),
					resource.TestCheckResourceAttr(resourceName, "manual.0.secret_key_ref", "account."+name),
					resource.TestCheckResourceAttr(resourceName, "manual.0.delegate_selectors.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "manual.0.delegate_selectors.0", "harness-delegate"),
				),
			},
		},
	})
}

func testAccDataSourceConnectorGcpSM_manual(name string) string {
	return fmt.Sprintf(`
		resource "harness_platform_secret_text" "test" {
			identifier = "%[1]s"
			name = "%[1]s"
			description = "test"
			tags = ["foo:bar"]
			secret_manager_identifier = "harnessSecretManager"
			value_type = "Reference"
			value = "secret"
		}

		resource "harness_platform_connector_gcp_secret_manager" "test" {
			identifier = "%[1]s"
			name = "%[1]s"
			description = "test"
			tags = ["foo:bar"]
		
			manual {
   				delegate_selectors = [ "harness-delegate" ]
    			secret_key_ref = "account.${harness_platform_secret_text.test.id}"
			}
			depends_on = [time_sleep.wait_4_seconds]
		}

		resource "time_sleep" "wait_4_seconds" {
			depends_on = [harness_platform_secret_text.test]
			destroy_duration = "4s"
		}

		data "harness_platform_connector_gcp_secret_manager" "test" {
			identifier = harness_platform_connector_gcp_secret_manager.test.identifier
		}
`, name)
}

func testOrgDataSourceConnectorGcpSM_manual(name string) string {
	return fmt.Sprintf(`
		resource "harness_platform_organization" "test" {
			identifier = "%[1]s"
			name = "%[1]s"
		}

		resource "time_sleep" "wait_3_seconds" {
			depends_on = [harness_platform_organization.test]
			create_duration = "3s"
		}

		resource "harness_platform_secret_text" "test" {
			identifier = "%[1]s"
			name = "%[1]s"
			description = "test"
			tags = ["foo:bar"]
			org_id = harness_platform_organization.test.id
			secret_manager_identifier = "harnessSecretManager"
			value_type = "Reference"
			value = "secret"
		}

		resource "harness_platform_connector_gcp_secret_manager" "test" {
			identifier = "%[1]s"
			name = "%[1]s"
			description = "test"
			tags = ["foo:bar"]
			org_id = harness_platform_organization.test.id
		
			manual {
   				delegate_selectors = [ "harness-delegate" ]
    			secret_key_ref = "org.${harness_platform_secret_text.test.id}"
			}
			depends_on = [time_sleep.wait_4_seconds]
		}

		resource "time_sleep" "wait_4_seconds" {
			depends_on = [harness_platform_secret_text.test]
			destroy_duration = "4s"
		}

		data "harness_platform_connector_gcp_secret_manager" "test" {
			identifier = harness_platform_connector_gcp_secret_manager.test.identifier
			org_id = harness_platform_connector_gcp_secret_manager.test.org_id
		}
`, name)
}

func testProjectDataSourceConnectorGcpSM_manual(name string) string {
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

		resource "time_sleep" "wait_3_seconds" {
			depends_on = [harness_platform_project.test]
			create_duration = "3s"
		}

		resource "harness_platform_secret_text" "test" {
			identifier = "%[1]s"
			name = "%[1]s"
			description = "test"
			tags = ["foo:bar"]
			org_id = harness_platform_organization.test.id
			project_id = harness_platform_project.test.id
			secret_manager_identifier = "harnessSecretManager"
			value_type = "Reference"
			value = "secret"
		}

		resource "harness_platform_connector_gcp_secret_manager" "test" {
			identifier = "%[1]s"
			name = "%[1]s"
			description = "test"
			tags = ["foo:bar"]
			org_id = harness_platform_organization.test.id
			project_id = harness_platform_project.test.id
		
			manual {
   				delegate_selectors = [ "harness-delegate" ]
    			secret_key_ref = "${harness_platform_secret_text.test.id}"
			}
			depends_on = [time_sleep.wait_4_seconds]
		}

		resource "time_sleep" "wait_4_seconds" {
			depends_on = [harness_platform_secret_text.test]
			destroy_duration = "4s"
		}

		data "harness_platform_connector_gcp_secret_manager" "test" {
			identifier = harness_platform_connector_gcp_secret_manager.test.identifier
			org_id = harness_platform_connector_gcp_secret_manager.test.org_id
			project_id = harness_platform_connector_gcp_secret_manager.test.project_id
		}
`, name)
}

func testAccDataSourceConnectorGcpSM_inherit(name string) string {
	return fmt.Sprintf(`
		resource "harness_platform_connector_gcp_secret_manager" "test" {
			identifier = "%[1]s"
			name = "%[1]s"
			description = "test"
			tags = ["foo:bar"]
		
			inherit_from_delegate {
   				delegate_selectors = [ "harness-delegate" ]
			}
		}

		data "harness_platform_connector_gcp_secret_manager" "test" {
			identifier = harness_platform_connector_gcp_secret_manager.test.identifier
		}
`, name)
}

func testOrgDataSourceConnectorGcpSM_inherit(name string) string {
	return fmt.Sprintf(`
		resource "harness_platform_organization" "test" {
			identifier = "%[1]s"
			name = "%[1]s"
		}

		resource "time_sleep" "wait_3_seconds" {
			depends_on = [harness_platform_organization.test]
			create_duration = "3s"
		}

		resource "harness_platform_connector_gcp_secret_manager" "test" {
			identifier = "%[1]s"
			name = "%[1]s"
			description = "test"
			tags = ["foo:bar"]
			org_id = harness_platform_organization.test.id
		
			inherit_from_delegate {
   				delegate_selectors = [ "harness-delegate" ]
			}
		}

		data "harness_platform_connector_gcp_secret_manager" "test" {
			identifier = harness_platform_connector_gcp_secret_manager.test.identifier
			org_id = harness_platform_connector_gcp_secret_manager.test.org_id
		}
`, name)
}

func testProjectDataSourceConnectorGcpSM_inherit(name string) string {
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

		resource "time_sleep" "wait_3_seconds" {
			depends_on = [harness_platform_project.test]
			create_duration = "3s"
		}

		resource "harness_platform_connector_gcp_secret_manager" "test" {
			identifier = "%[1]s"
			name = "%[1]s"
			description = "test"
			tags = ["foo:bar"]
			org_id = harness_platform_organization.test.id
			project_id = harness_platform_project.test.id
		
			inherit_from_delegate {
   				delegate_selectors = [ "harness-delegate" ]
			}
		}

		data "harness_platform_connector_gcp_secret_manager" "test" {
			identifier = harness_platform_connector_gcp_secret_manager.test.identifier
			org_id = harness_platform_connector_gcp_secret_manager.test.org_id
			project_id = harness_platform_connector_gcp_secret_manager.test.project_id
		}
`, name)
}

func testAccDataSourceConnectorGcpSM_oidc_platform(name string) string {
	return fmt.Sprintf(`
		resource "harness_platform_connector_gcp_secret_manager" "test" {
			identifier = "%[1]s"
			name = "%[1]s"
			description = "test"
			tags = ["foo:bar"]
		
			execute_on_delegate = false

			oidc_authentication {
   				workload_pool_id = "harness-pool-test"
				provider_id = "harness"
				gcp_project_id = "1234567"
				service_account_email = "harness.sample@iam.gserviceaccount.com"
			}
		}

		data "harness_platform_connector_gcp_secret_manager" "test" {
			identifier = harness_platform_connector_gcp_secret_manager.test.identifier
		}
`, name)
}

func testOrgDataSourceConnectorGcpSM_oidc_platform(name string) string {
	return fmt.Sprintf(`
		resource "harness_platform_organization" "test" {
			identifier = "%[1]s"
			name = "%[1]s"
		}

		resource "time_sleep" "wait_3_seconds" {
			depends_on = [harness_platform_organization.test]
			create_duration = "3s"
		}

		resource "harness_platform_connector_gcp_secret_manager" "test" {
			identifier = "%[1]s"
			name = "%[1]s"
			description = "test"
			tags = ["foo:bar"]
			org_id = harness_platform_organization.test.id

			execute_on_delegate = false
		
			oidc_authentication {
   				workload_pool_id = "harness-pool-test"
				provider_id = "harness"
				gcp_project_id = "1234567"
				service_account_email = "harness.sample@iam.gserviceaccount.com"
			}
		}

		data "harness_platform_connector_gcp_secret_manager" "test" {
			identifier = harness_platform_connector_gcp_secret_manager.test.identifier
			org_id = harness_platform_connector_gcp_secret_manager.test.org_id
		}
`, name)
}

func testProjectDataSourceConnectorGcpSM_oidc_platform(name string) string {
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

		resource "time_sleep" "wait_3_seconds" {
			depends_on = [harness_platform_project.test]
			create_duration = "3s"
		}

		resource "harness_platform_connector_gcp_secret_manager" "test" {
			identifier = "%[1]s"
			name = "%[1]s"
			description = "test"
			tags = ["foo:bar"]
			org_id = harness_platform_organization.test.id
			project_id = harness_platform_project.test.id

			execute_on_delegate = false
		
			oidc_authentication {
   				workload_pool_id = "harness-pool-test"
				provider_id = "harness"
				gcp_project_id = "1234567"
				service_account_email = "harness.sample@iam.gserviceaccount.com"
			}
		}

		data "harness_platform_connector_gcp_secret_manager" "test" {
			identifier = harness_platform_connector_gcp_secret_manager.test.identifier
			org_id = harness_platform_connector_gcp_secret_manager.test.org_id
			project_id = harness_platform_connector_gcp_secret_manager.test.project_id
		}
`, name)
}

func testAccDataSourceConnectorGcpSM_oidc_delegate(name string) string {
	return fmt.Sprintf(`
		resource "harness_platform_connector_gcp_secret_manager" "test" {
			identifier = "%[1]s"
			name = "%[1]s"
			description = "test"
			tags = ["foo:bar"]
		
			oidc_authentication {
   				workload_pool_id = "harness-pool-test"
				provider_id = "harness"
				gcp_project_id = "1234567"
				service_account_email = "harness.sample@iam.gserviceaccount.com"
				delegate_selectors = [ "harness-delegate" ]
			}
		}

		data "harness_platform_connector_gcp_secret_manager" "test" {
			identifier = harness_platform_connector_gcp_secret_manager.test.identifier
		}
`, name)
}

func testOrgDataSourceConnectorGcpSM_oidc_delegate(name string) string {
	return fmt.Sprintf(`
		resource "harness_platform_organization" "test" {
			identifier = "%[1]s"
			name = "%[1]s"
		}

		resource "time_sleep" "wait_3_seconds" {
			depends_on = [harness_platform_organization.test]
			create_duration = "3s"
		}

		resource "harness_platform_connector_gcp_secret_manager" "test" {
			identifier = "%[1]s"
			name = "%[1]s"
			description = "test"
			tags = ["foo:bar"]
			org_id = harness_platform_organization.test.id
		
			oidc_authentication {
   				workload_pool_id = "harness-pool-test"
				provider_id = "harness"
				gcp_project_id = "1234567"
				service_account_email = "harness.sample@iam.gserviceaccount.com"
				delegate_selectors = [ "harness-delegate" ]
			}
		}

		data "harness_platform_connector_gcp_secret_manager" "test" {
			identifier = harness_platform_connector_gcp_secret_manager.test.identifier
			org_id = harness_platform_connector_gcp_secret_manager.test.org_id
		}
`, name)
}

func testProjectDataSourceConnectorGcpSM_oidc_delegate(name string) string {
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

		resource "time_sleep" "wait_3_seconds" {
			depends_on = [harness_platform_project.test]
			create_duration = "3s"
		}

		resource "harness_platform_connector_gcp_secret_manager" "test" {
			identifier = "%[1]s"
			name = "%[1]s"
			description = "test"
			tags = ["foo:bar"]
			org_id = harness_platform_organization.test.id
			project_id = harness_platform_project.test.id
		
			oidc_authentication {
   				workload_pool_id = "harness-pool-test"
				provider_id = "harness"
				gcp_project_id = "1234567"
				service_account_email = "harness.sample@iam.gserviceaccount.com"
				delegate_selectors = [ "harness-delegate" ]
			}
		}

		data "harness_platform_connector_gcp_secret_manager" "test" {
			identifier = harness_platform_connector_gcp_secret_manager.test.identifier
			org_id = harness_platform_connector_gcp_secret_manager.test.org_id
			project_id = harness_platform_connector_gcp_secret_manager.test.project_id
		}
`, name)
}

func testAccDataSourceConnectorGcpSM_manual_default(name string) string {
	return fmt.Sprintf(`
		resource "harness_platform_secret_text" "test" {
			identifier = "%[1]s"
			name = "%[1]s"
			description = "test"
			tags = ["foo:bar"]
			secret_manager_identifier = "harnessSecretManager"
			value_type = "Reference"
			value = "secret"
		}

		resource "harness_platform_connector_gcp_secret_manager" "test" {
			identifier = "%[1]s"
			name = "%[1]s"
			description = "test"
			tags = ["foo:bar"]

			default = true
		
			manual {
   				delegate_selectors = [ "harness-delegate" ]
    			secret_key_ref = "account.${harness_platform_secret_text.test.id}"
			}
			depends_on = [time_sleep.wait_4_seconds]
		}

		resource "time_sleep" "wait_4_seconds" {
			depends_on = [harness_platform_secret_text.test]
			destroy_duration = "4s"
		}
			
		data "harness_platform_connector_gcp_secret_manager" "test" {
			identifier = harness_platform_connector_gcp_secret_manager.test.identifier
		}
`, name)
}
