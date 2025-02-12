package secretManagers_test

import (
	"fmt"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceConnectorGcpKMS_manual(t *testing.T) {
	var (
		name         = fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))
		resourceName = "data.harness_platform_connector_gcp_kms.test"
	)

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {},
		},
		CheckDestroy: testAccConnectorDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceConnectorGcpKMS_manual(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", name),
					resource.TestCheckResourceAttr(resourceName, "identifier", name),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "default", "false"),
					resource.TestCheckResourceAttr(resourceName, "execute_on_delegate", "true"),
					resource.TestCheckResourceAttr(resourceName, "region", "us-west1"),
					resource.TestCheckResourceAttr(resourceName, "gcp_project_id", "1234567"),
					resource.TestCheckResourceAttr(resourceName, "key_ring", "key_ring"),
					resource.TestCheckResourceAttr(resourceName, "key_name", "key_name"),
					resource.TestCheckResourceAttr(resourceName, "manual.0.credentials", "account."+name),
					resource.TestCheckResourceAttr(resourceName, "manual.0.delegate_selectors.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "manual.0.delegate_selectors.0", "harness-delegate"),
				),
			},
		},
	})
}

func TestOrgDataSourceConnectorGcpKMS_manual(t *testing.T) {
	var (
		name         = fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))
		resourceName = "data.harness_platform_connector_gcp_kms.test"
	)

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {},
		},
		CheckDestroy: testAccConnectorDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testOrgDataSourceConnectorGcpKMS_manual(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", name),
					resource.TestCheckResourceAttr(resourceName, "identifier", name),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "org_id", name),
					resource.TestCheckResourceAttr(resourceName, "default", "false"),
					resource.TestCheckResourceAttr(resourceName, "execute_on_delegate", "true"),
					resource.TestCheckResourceAttr(resourceName, "region", "us-west1"),
					resource.TestCheckResourceAttr(resourceName, "gcp_project_id", "1234567"),
					resource.TestCheckResourceAttr(resourceName, "key_ring", "key_ring"),
					resource.TestCheckResourceAttr(resourceName, "key_name", "key_name"),
					resource.TestCheckResourceAttr(resourceName, "manual.0.credentials", "org."+name),
					resource.TestCheckResourceAttr(resourceName, "manual.0.delegate_selectors.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "manual.0.delegate_selectors.0", "harness-delegate"),
				),
			},
		},
	})
}

func TestProjectDataSourceConnectorGcpKMS_manual(t *testing.T) {
	var (
		name         = fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))
		resourceName = "data.harness_platform_connector_gcp_kms.test"
	)

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {},
		},
		CheckDestroy: testAccConnectorDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testProjectDataSourceConnectorGcpKMS_manual(name),
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
					resource.TestCheckResourceAttr(resourceName, "region", "us-west1"),
					resource.TestCheckResourceAttr(resourceName, "gcp_project_id", "1234567"),
					resource.TestCheckResourceAttr(resourceName, "key_ring", "key_ring"),
					resource.TestCheckResourceAttr(resourceName, "key_name", "key_name"),
					resource.TestCheckResourceAttr(resourceName, "manual.0.credentials", name),
					resource.TestCheckResourceAttr(resourceName, "manual.0.delegate_selectors.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "manual.0.delegate_selectors.0", "harness-delegate"),
				),
			},
		},
	})
}

func TestAccDataSourceConnectorGcpKMS_oidc_platform(t *testing.T) {
	var (
		name         = fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))
		resourceName = "data.harness_platform_connector_gcp_kms.test"
	)

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {},
		},
		CheckDestroy: testAccConnectorDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceConnectorGcpKMS_oidc_platform(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", name),
					resource.TestCheckResourceAttr(resourceName, "identifier", name),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "default", "false"),
					resource.TestCheckResourceAttr(resourceName, "execute_on_delegate", "false"),
					resource.TestCheckResourceAttr(resourceName, "region", "us-west1"),
					resource.TestCheckResourceAttr(resourceName, "gcp_project_id", "1234567"),
					resource.TestCheckResourceAttr(resourceName, "key_ring", "key_ring"),
					resource.TestCheckResourceAttr(resourceName, "key_name", "key_name"),
					resource.TestCheckResourceAttr(resourceName, "oidc_authentication.0.workload_pool_id", "harness-pool-test"),
					resource.TestCheckResourceAttr(resourceName, "oidc_authentication.0.provider_id", "harness"),
					resource.TestCheckResourceAttr(resourceName, "oidc_authentication.0.gcp_project_id", "1234567"),
					resource.TestCheckResourceAttr(resourceName, "oidc_authentication.0.service_account_email", "harness.sample@iam.gserviceaccount.com"),
				),
			},
		},
	})
}

func TestOrgDataSourceConnectorGcpKMS_oidc_platform(t *testing.T) {
	var (
		name         = fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))
		resourceName = "data.harness_platform_connector_gcp_kms.test"
	)

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {},
		},
		CheckDestroy: testAccConnectorDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testOrgDataSourceConnectorGcpKMS_oidc_platform(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", name),
					resource.TestCheckResourceAttr(resourceName, "identifier", name),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "org_id", name),
					resource.TestCheckResourceAttr(resourceName, "default", "false"),
					resource.TestCheckResourceAttr(resourceName, "execute_on_delegate", "false"),
					resource.TestCheckResourceAttr(resourceName, "region", "us-west1"),
					resource.TestCheckResourceAttr(resourceName, "gcp_project_id", "1234567"),
					resource.TestCheckResourceAttr(resourceName, "key_ring", "key_ring"),
					resource.TestCheckResourceAttr(resourceName, "key_name", "key_name"),
					resource.TestCheckResourceAttr(resourceName, "oidc_authentication.0.workload_pool_id", "harness-pool-test"),
					resource.TestCheckResourceAttr(resourceName, "oidc_authentication.0.provider_id", "harness"),
					resource.TestCheckResourceAttr(resourceName, "oidc_authentication.0.gcp_project_id", "1234567"),
					resource.TestCheckResourceAttr(resourceName, "oidc_authentication.0.service_account_email", "harness.sample@iam.gserviceaccount.com"),
				),
			},
		},
	})
}

func TestProjectDataSourceConnectorGcpKMS_oidc_platform(t *testing.T) {
	var (
		name         = fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))
		resourceName = "data.harness_platform_connector_gcp_kms.test"
	)

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {},
		},
		CheckDestroy: testAccConnectorDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testProjectDataSourceConnectorGcpKMS_oidc_platform(name),
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
					resource.TestCheckResourceAttr(resourceName, "region", "us-west1"),
					resource.TestCheckResourceAttr(resourceName, "gcp_project_id", "1234567"),
					resource.TestCheckResourceAttr(resourceName, "key_ring", "key_ring"),
					resource.TestCheckResourceAttr(resourceName, "key_name", "key_name"),
					resource.TestCheckResourceAttr(resourceName, "oidc_authentication.0.workload_pool_id", "harness-pool-test"),
					resource.TestCheckResourceAttr(resourceName, "oidc_authentication.0.provider_id", "harness"),
					resource.TestCheckResourceAttr(resourceName, "oidc_authentication.0.gcp_project_id", "1234567"),
					resource.TestCheckResourceAttr(resourceName, "oidc_authentication.0.service_account_email", "harness.sample@iam.gserviceaccount.com"),
				),
			},
		},
	})
}

func TestAccDataSourceConnectorGcpKMS_oidc_delegate(t *testing.T) {
	var (
		name         = fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))
		resourceName = "data.harness_platform_connector_gcp_kms.test"
	)

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {},
		},
		CheckDestroy: testAccConnectorDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceConnectorGcpKMS_oidc_delegate(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", name),
					resource.TestCheckResourceAttr(resourceName, "identifier", name),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "default", "false"),
					resource.TestCheckResourceAttr(resourceName, "execute_on_delegate", "true"),
					resource.TestCheckResourceAttr(resourceName, "region", "us-west1"),
					resource.TestCheckResourceAttr(resourceName, "gcp_project_id", "1234567"),
					resource.TestCheckResourceAttr(resourceName, "key_ring", "key_ring"),
					resource.TestCheckResourceAttr(resourceName, "key_name", "key_name"),
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

func TestOrgDataSourceConnectorGcpKMS_oidc_delegate(t *testing.T) {
	var (
		name         = fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))
		resourceName = "data.harness_platform_connector_gcp_kms.test"
	)

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {},
		},
		CheckDestroy: testAccConnectorDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testOrgDataSourceConnectorGcpKMS_oidc_delegate(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", name),
					resource.TestCheckResourceAttr(resourceName, "identifier", name),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "org_id", name),
					resource.TestCheckResourceAttr(resourceName, "default", "false"),
					resource.TestCheckResourceAttr(resourceName, "execute_on_delegate", "true"),
					resource.TestCheckResourceAttr(resourceName, "region", "us-west1"),
					resource.TestCheckResourceAttr(resourceName, "gcp_project_id", "1234567"),
					resource.TestCheckResourceAttr(resourceName, "key_ring", "key_ring"),
					resource.TestCheckResourceAttr(resourceName, "key_name", "key_name"),
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

func TestProjectDataSourceConnectorGcpKMS_oidc_delegate(t *testing.T) {
	var (
		name         = fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))
		resourceName = "data.harness_platform_connector_gcp_kms.test"
	)

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {},
		},
		CheckDestroy: testAccConnectorDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testProjectDataSourceConnectorGcpKMS_oidc_delegate(name),
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
					resource.TestCheckResourceAttr(resourceName, "region", "us-west1"),
					resource.TestCheckResourceAttr(resourceName, "gcp_project_id", "1234567"),
					resource.TestCheckResourceAttr(resourceName, "key_ring", "key_ring"),
					resource.TestCheckResourceAttr(resourceName, "key_name", "key_name"),
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

func TestAccDataSourceConnectorGcpKMS_manual_default(t *testing.T) {
	var (
		name         = fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))
		resourceName = "data.harness_platform_connector_gcp_kms.test"
	)

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {},
		},
		CheckDestroy: testAccConnectorDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceConnectorGcpKMS_manual_default(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", name),
					resource.TestCheckResourceAttr(resourceName, "identifier", name),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "default", "true"),
					resource.TestCheckResourceAttr(resourceName, "execute_on_delegate", "true"),
					resource.TestCheckResourceAttr(resourceName, "region", "us-west1"),
					resource.TestCheckResourceAttr(resourceName, "gcp_project_id", "1234567"),
					resource.TestCheckResourceAttr(resourceName, "key_ring", "key_ring"),
					resource.TestCheckResourceAttr(resourceName, "key_name", "key_name"),
					resource.TestCheckResourceAttr(resourceName, "manual.0.credentials", "account."+name),
					resource.TestCheckResourceAttr(resourceName, "manual.0.delegate_selectors.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "manual.0.delegate_selectors.0", "harness-delegate"),
				),
			},
		},
	})
}

func testAccDataSourceConnectorGcpKMS_manual(name string) string {
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

		resource "harness_platform_connector_gcp_kms" "test" {
			identifier = "%[1]s"
			name = "%[1]s"
			description = "test"
			tags = ["foo:bar"]

			region     = "us-west1"
			gcp_project_id = "1234567"
			key_ring   = "key_ring"
			key_name   = "key_name"
		
			manual {
   				delegate_selectors = [ "harness-delegate" ]
    			credentials = "account.${harness_platform_secret_text.test.id}"
			}
			depends_on = [time_sleep.wait_4_seconds]
		}

		resource "time_sleep" "wait_4_seconds" {
			depends_on = [harness_platform_secret_text.test]
			destroy_duration = "4s"
		}

		data "harness_platform_connector_gcp_kms" "test" {
			identifier = harness_platform_connector_gcp_kms.test.identifier
		}
`, name)
}

func testOrgDataSourceConnectorGcpKMS_manual(name string) string {
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
			value_type = "Inline"
			value = "secret"
		}

		resource "harness_platform_connector_gcp_kms" "test" {
			identifier = "%[1]s"
			name = "%[1]s"
			description = "test"
			tags = ["foo:bar"]
			org_id = harness_platform_organization.test.id

			region     = "us-west1"
			gcp_project_id = "1234567"
			key_ring   = "key_ring"
			key_name   = "key_name"
		
			manual {
   				delegate_selectors = [ "harness-delegate" ]
    			credentials = "org.${harness_platform_secret_text.test.id}"
			}
			depends_on = [time_sleep.wait_4_seconds]
		}

		resource "time_sleep" "wait_4_seconds" {
			depends_on = [harness_platform_secret_text.test]
			destroy_duration = "4s"
		}

		data "harness_platform_connector_gcp_kms" "test" {
			identifier = harness_platform_connector_gcp_kms.test.identifier
			org_id = harness_platform_organization.test.id
		}
`, name)
}

func testProjectDataSourceConnectorGcpKMS_manual(name string) string {
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

		resource "harness_platform_connector_gcp_kms" "test" {
			identifier = "%[1]s"
			name = "%[1]s"
			description = "test"
			tags = ["foo:bar"]
			org_id = harness_platform_organization.test.id
			project_id = harness_platform_project.test.id

			region     = "us-west1"
			gcp_project_id = "1234567"
			key_ring   = "key_ring"
			key_name   = "key_name"
		
			manual {
   				delegate_selectors = [ "harness-delegate" ]
    			credentials = "${harness_platform_secret_text.test.id}"
			}
			depends_on = [time_sleep.wait_4_seconds]
		}

		resource "time_sleep" "wait_4_seconds" {
			depends_on = [harness_platform_secret_text.test]
			destroy_duration = "4s"
		}

		data "harness_platform_connector_gcp_kms" "test" {
			identifier = harness_platform_connector_gcp_kms.test.identifier
			org_id = harness_platform_organization.test.id
			project_id = harness_platform_project.test.id
		}
`, name)
}

func testAccDataSourceConnectorGcpKMS_oidc_platform(name string) string {
	return fmt.Sprintf(`
		resource "harness_platform_connector_gcp_kms" "test" {
			identifier = "%[1]s"
			name = "%[1]s"
			description = "test"
			tags = ["foo:bar"]

			region     = "us-west1"
			gcp_project_id = "1234567"
			key_ring   = "key_ring"
			key_name   = "key_name"
		
			execute_on_delegate = false

			oidc_authentication {
   				workload_pool_id = "harness-pool-test"
				provider_id = "harness"
				gcp_project_id = "1234567"
				service_account_email = "harness.sample@iam.gserviceaccount.com"
			}
		}

		data "harness_platform_connector_gcp_kms" "test" {
			identifier = harness_platform_connector_gcp_kms.test.identifier
		}
`, name)
}

func testOrgDataSourceConnectorGcpKMS_oidc_platform(name string) string {
	return fmt.Sprintf(`
		resource "harness_platform_organization" "test" {
			identifier = "%[1]s"
			name = "%[1]s"
		}

		resource "time_sleep" "wait_3_seconds" {
			depends_on = [harness_platform_organization.test]
			create_duration = "3s"
		}

		resource "harness_platform_connector_gcp_kms" "test" {
			identifier = "%[1]s"
			name = "%[1]s"
			description = "test"
			tags = ["foo:bar"]
			org_id = harness_platform_organization.test.id

			region     = "us-west1"
			gcp_project_id = "1234567"
			key_ring   = "key_ring"
			key_name   = "key_name"

			execute_on_delegate = false
		
			oidc_authentication {
   				workload_pool_id = "harness-pool-test"
				provider_id = "harness"
				gcp_project_id = "1234567"
				service_account_email = "harness.sample@iam.gserviceaccount.com"
			}
		}

		data "harness_platform_connector_gcp_kms" "test" {
			identifier = harness_platform_connector_gcp_kms.test.identifier
			org_id = harness_platform_organization.test.id
		}
`, name)
}

func testProjectDataSourceConnectorGcpKMS_oidc_platform(name string) string {
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

		resource "harness_platform_connector_gcp_kms" "test" {
			identifier = "%[1]s"
			name = "%[1]s"
			description = "test"
			tags = ["foo:bar"]
			org_id = harness_platform_organization.test.id
			project_id = harness_platform_project.test.id

			region     = "us-west1"
			gcp_project_id = "1234567"
			key_ring   = "key_ring"
			key_name   = "key_name"

			execute_on_delegate = false
		
			oidc_authentication {
   				workload_pool_id = "harness-pool-test"
				provider_id = "harness"
				gcp_project_id = "1234567"
				service_account_email = "harness.sample@iam.gserviceaccount.com"
			}
		}

		data "harness_platform_connector_gcp_kms" "test" {
			identifier = harness_platform_connector_gcp_kms.test.identifier
			org_id = harness_platform_organization.test.id
			project_id = harness_platform_project.test.id
		}
`, name)
}

func testAccDataSourceConnectorGcpKMS_oidc_delegate(name string) string {
	return fmt.Sprintf(`
		resource "harness_platform_connector_gcp_kms" "test" {
			identifier = "%[1]s"
			name = "%[1]s"
			description = "test"
			tags = ["foo:bar"]

			region     = "us-west1"
			gcp_project_id = "1234567"
			key_ring   = "key_ring"
			key_name   = "key_name"
		
			oidc_authentication {
   				workload_pool_id = "harness-pool-test"
				provider_id = "harness"
				gcp_project_id = "1234567"
				service_account_email = "harness.sample@iam.gserviceaccount.com"
				delegate_selectors = [ "harness-delegate" ]
			}
		}

		data "harness_platform_connector_gcp_kms" "test" {
			identifier = harness_platform_connector_gcp_kms.test.identifier
		}
`, name)
}

func testOrgDataSourceConnectorGcpKMS_oidc_delegate(name string) string {
	return fmt.Sprintf(`
		resource "harness_platform_organization" "test" {
			identifier = "%[1]s"
			name = "%[1]s"
		}

		resource "time_sleep" "wait_3_seconds" {
			depends_on = [harness_platform_organization.test]
			create_duration = "3s"
		}

		resource "harness_platform_connector_gcp_kms" "test" {
			identifier = "%[1]s"
			name = "%[1]s"
			description = "test"
			tags = ["foo:bar"]
			org_id = harness_platform_organization.test.id

			region     = "us-west1"
			gcp_project_id = "1234567"
			key_ring   = "key_ring"
			key_name   = "key_name"
		
			oidc_authentication {
   				workload_pool_id = "harness-pool-test"
				provider_id = "harness"
				gcp_project_id = "1234567"
				service_account_email = "harness.sample@iam.gserviceaccount.com"
				delegate_selectors = [ "harness-delegate" ]
			}
		}

		data "harness_platform_connector_gcp_kms" "test" {
			identifier = harness_platform_connector_gcp_kms.test.identifier
			org_id = harness_platform_organization.test.id
		}
`, name)
}

func testProjectDataSourceConnectorGcpKMS_oidc_delegate(name string) string {
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

		resource "harness_platform_connector_gcp_kms" "test" {
			identifier = "%[1]s"
			name = "%[1]s"
			description = "test"
			tags = ["foo:bar"]
			org_id = harness_platform_organization.test.id
			project_id = harness_platform_project.test.id
			
			region     = "us-west1"
			gcp_project_id = "1234567"
			key_ring   = "key_ring"
			key_name   = "key_name"
		
			oidc_authentication {
   				workload_pool_id = "harness-pool-test"
				provider_id = "harness"
				gcp_project_id = "1234567"
				service_account_email = "harness.sample@iam.gserviceaccount.com"
				delegate_selectors = [ "harness-delegate" ]
			}
		}

		data "harness_platform_connector_gcp_kms" "test" {
			identifier = harness_platform_connector_gcp_kms.test.identifier
			org_id = harness_platform_organization.test.id
			project_id = harness_platform_project.test.id
		}
`, name)
}

func testAccDataSourceConnectorGcpKMS_manual_default(name string) string {
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

		resource "harness_platform_connector_gcp_kms" "test" {
			identifier = "%[1]s"
			name = "%[1]s"
			description = "test"
			tags = ["foo:bar"]

			region     = "us-west1"
			gcp_project_id = "1234567"
			key_ring   = "key_ring"
			key_name   = "key_name"

			default = true
		
			manual {
   				delegate_selectors = [ "harness-delegate" ]
    			credentials = "account.${harness_platform_secret_text.test.id}"
			}
			depends_on = [time_sleep.wait_4_seconds]
		}

		resource "time_sleep" "wait_4_seconds" {
			depends_on = [harness_platform_secret_text.test]
			destroy_duration = "4s"
		}

		data "harness_platform_connector_gcp_kms" "test" {
			identifier = harness_platform_connector_gcp_kms.test.identifier
		}
`, name)
}
