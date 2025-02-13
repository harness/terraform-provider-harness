package secretManagers_test

import (
	"fmt"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceConnectorGcpKMS_manual(t *testing.T) {
	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	name := id
	updatedName := fmt.Sprintf("%s_updated", name)
	resourceName := "harness_platform_connector_gcp_kms.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {},
		},
		CheckDestroy: testAccConnectorDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceConnectorGcpKMS_manual(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "default", "false"),
					resource.TestCheckResourceAttr(resourceName, "execute_on_delegate", "true"),
					resource.TestCheckResourceAttr(resourceName, "region", "us-west1"),
					resource.TestCheckResourceAttr(resourceName, "gcp_project_id", "1234567"),
					resource.TestCheckResourceAttr(resourceName, "key_ring", "key_ring"),
					resource.TestCheckResourceAttr(resourceName, "key_name", "key_name"),
					resource.TestCheckResourceAttr(resourceName, "manual.0.credentials", "account."+id),
					resource.TestCheckResourceAttr(resourceName, "manual.0.delegate_selectors.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "manual.0.delegate_selectors.0", "harness-delegate"),
				),
			},
			{
				Config: testAccResourceConnectorGcpKMS_manual(id, updatedName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "default", "false"),
					resource.TestCheckResourceAttr(resourceName, "execute_on_delegate", "true"),
					resource.TestCheckResourceAttr(resourceName, "region", "us-west1"),
					resource.TestCheckResourceAttr(resourceName, "gcp_project_id", "1234567"),
					resource.TestCheckResourceAttr(resourceName, "key_ring", "key_ring"),
					resource.TestCheckResourceAttr(resourceName, "key_name", "key_name"),
					resource.TestCheckResourceAttr(resourceName, "manual.0.credentials", "account."+id),
					resource.TestCheckResourceAttr(resourceName, "manual.0.delegate_selectors.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "manual.0.delegate_selectors.0", "harness-delegate"),
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

func TestOrgResourceConnectorGcpKMS_manual(t *testing.T) {
	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	name := id
	updatedName := fmt.Sprintf("%s_updated", name)
	resourceName := "harness_platform_connector_gcp_kms.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {},
		},
		CheckDestroy: testAccConnectorDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testOrgResourceConnectorGcpKMS_manual(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "org_id", id),
					resource.TestCheckResourceAttr(resourceName, "default", "false"),
					resource.TestCheckResourceAttr(resourceName, "execute_on_delegate", "true"),
					resource.TestCheckResourceAttr(resourceName, "region", "us-west1"),
					resource.TestCheckResourceAttr(resourceName, "gcp_project_id", "1234567"),
					resource.TestCheckResourceAttr(resourceName, "key_ring", "key_ring"),
					resource.TestCheckResourceAttr(resourceName, "key_name", "key_name"),
					resource.TestCheckResourceAttr(resourceName, "manual.0.credentials", "org."+id),
					resource.TestCheckResourceAttr(resourceName, "manual.0.delegate_selectors.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "manual.0.delegate_selectors.0", "harness-delegate"),
				),
			},
			{
				Config: testOrgResourceConnectorGcpKMS_manual(id, updatedName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "org_id", id),
					resource.TestCheckResourceAttr(resourceName, "default", "false"),
					resource.TestCheckResourceAttr(resourceName, "execute_on_delegate", "true"),
					resource.TestCheckResourceAttr(resourceName, "region", "us-west1"),
					resource.TestCheckResourceAttr(resourceName, "gcp_project_id", "1234567"),
					resource.TestCheckResourceAttr(resourceName, "key_ring", "key_ring"),
					resource.TestCheckResourceAttr(resourceName, "key_name", "key_name"),
					resource.TestCheckResourceAttr(resourceName, "manual.0.credentials", "org."+id),
					resource.TestCheckResourceAttr(resourceName, "manual.0.delegate_selectors.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "manual.0.delegate_selectors.0", "harness-delegate"),
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

func TestProjectResourceConnectorGcpKMS_manual(t *testing.T) {
	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	name := id
	updatedName := fmt.Sprintf("%s_updated", name)
	resourceName := "harness_platform_connector_gcp_kms.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {},
		},
		CheckDestroy: testAccConnectorDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testProjectResourceConnectorGcpKMS_manual(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "org_id", id),
					resource.TestCheckResourceAttr(resourceName, "project_id", id),
					resource.TestCheckResourceAttr(resourceName, "default", "false"),
					resource.TestCheckResourceAttr(resourceName, "execute_on_delegate", "true"),
					resource.TestCheckResourceAttr(resourceName, "region", "us-west1"),
					resource.TestCheckResourceAttr(resourceName, "gcp_project_id", "1234567"),
					resource.TestCheckResourceAttr(resourceName, "key_ring", "key_ring"),
					resource.TestCheckResourceAttr(resourceName, "key_name", "key_name"),
					resource.TestCheckResourceAttr(resourceName, "manual.0.credentials", id),
					resource.TestCheckResourceAttr(resourceName, "manual.0.delegate_selectors.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "manual.0.delegate_selectors.0", "harness-delegate"),
				),
			},
			{
				Config: testProjectResourceConnectorGcpKMS_manual(id, updatedName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "org_id", id),
					resource.TestCheckResourceAttr(resourceName, "project_id", id),
					resource.TestCheckResourceAttr(resourceName, "default", "false"),
					resource.TestCheckResourceAttr(resourceName, "execute_on_delegate", "true"),
					resource.TestCheckResourceAttr(resourceName, "region", "us-west1"),
					resource.TestCheckResourceAttr(resourceName, "gcp_project_id", "1234567"),
					resource.TestCheckResourceAttr(resourceName, "key_ring", "key_ring"),
					resource.TestCheckResourceAttr(resourceName, "key_name", "key_name"),
					resource.TestCheckResourceAttr(resourceName, "manual.0.credentials", id),
					resource.TestCheckResourceAttr(resourceName, "manual.0.delegate_selectors.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "manual.0.delegate_selectors.0", "harness-delegate"),
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

func TestAccResourceConnectorGcpKMS_oidc_platform(t *testing.T) {
	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	name := id
	updatedName := fmt.Sprintf("%s_updated", name)
	resourceName := "harness_platform_connector_gcp_kms.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {},
		},
		CheckDestroy: testAccConnectorDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceConnectorGcpKMS_oidc_platform(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
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
			{
				Config: testAccResourceConnectorGcpKMS_oidc_platform(id, updatedName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
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
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestOrgResourceConnectorGcpKMS_oidc_platform(t *testing.T) {
	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	name := id
	updatedName := fmt.Sprintf("%s_updated", name)
	resourceName := "harness_platform_connector_gcp_kms.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {},
		},
		CheckDestroy: testAccConnectorDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testOrgResourceConnectorGcpKMS_oidc_platform(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "org_id", id),
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
			{
				Config: testOrgResourceConnectorGcpKMS_oidc_platform(id, updatedName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "org_id", id),
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
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: acctest.OrgResourceImportStateIdFunc(resourceName),
			},
		},
	})
}

func TestProjectResourceConnectorGcpKMS_oidc_platform(t *testing.T) {
	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	name := id
	updatedName := fmt.Sprintf("%s_updated", name)
	resourceName := "harness_platform_connector_gcp_kms.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {},
		},
		CheckDestroy: testAccConnectorDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testProjectResourceConnectorGcpKMS_oidc_platform(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "org_id", id),
					resource.TestCheckResourceAttr(resourceName, "project_id", id),
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
			{
				Config: testProjectResourceConnectorGcpKMS_oidc_platform(id, updatedName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "org_id", id),
					resource.TestCheckResourceAttr(resourceName, "project_id", id),
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
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: acctest.ProjectResourceImportStateIdFunc(resourceName),
			},
		},
	})
}

func TestAccResourceConnectorGcpKMS_oidc_delegate(t *testing.T) {
	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	name := id
	updatedName := fmt.Sprintf("%s_updated", name)
	resourceName := "harness_platform_connector_gcp_kms.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {},
		},
		CheckDestroy: testAccConnectorDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceConnectorGcpKMS_oidc_delegate(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
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
			{
				Config: testAccResourceConnectorGcpKMS_oidc_delegate(id, updatedName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
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
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestOrgResourceConnectorGcpKMS_oidc_delegate(t *testing.T) {
	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	name := id
	updatedName := fmt.Sprintf("%s_updated", name)
	resourceName := "harness_platform_connector_gcp_kms.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {},
		},
		CheckDestroy: testAccConnectorDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testOrgResourceConnectorGcpKMS_oidc_delegate(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "org_id", id),
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
			{
				Config: testOrgResourceConnectorGcpKMS_oidc_delegate(id, updatedName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "org_id", id),
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
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: acctest.OrgResourceImportStateIdFunc(resourceName),
			},
		},
	})
}

func TestProjectResourceConnectorGcpKMS_oidc_delegate(t *testing.T) {
	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	name := id
	updatedName := fmt.Sprintf("%s_updated", name)
	resourceName := "harness_platform_connector_gcp_kms.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {},
		},
		CheckDestroy: testAccConnectorDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testProjectResourceConnectorGcpKMS_oidc_delegate(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "org_id", id),
					resource.TestCheckResourceAttr(resourceName, "project_id", id),
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
			{
				Config: testProjectResourceConnectorGcpKMS_oidc_delegate(id, updatedName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "org_id", id),
					resource.TestCheckResourceAttr(resourceName, "project_id", id),
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
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: acctest.ProjectResourceImportStateIdFunc(resourceName),
			},
		},
	})
}

func TestAccResourceConnectorGcpKMS_manual_default(t *testing.T) {
	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	name := id
	updatedName := fmt.Sprintf("%s_updated", name)
	resourceName := "harness_platform_connector_gcp_kms.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {},
		},
		CheckDestroy: testAccConnectorDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceConnectorGcpKMS_manual_default(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "default", "true"),
					resource.TestCheckResourceAttr(resourceName, "execute_on_delegate", "true"),
					resource.TestCheckResourceAttr(resourceName, "region", "us-west1"),
					resource.TestCheckResourceAttr(resourceName, "gcp_project_id", "1234567"),
					resource.TestCheckResourceAttr(resourceName, "key_ring", "key_ring"),
					resource.TestCheckResourceAttr(resourceName, "key_name", "key_name"),
					resource.TestCheckResourceAttr(resourceName, "manual.0.credentials", "account."+id),
					resource.TestCheckResourceAttr(resourceName, "manual.0.delegate_selectors.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "manual.0.delegate_selectors.0", "harness-delegate"),
				),
			},
			{
				Config: testAccResourceConnectorGcpKMS_manual_default(id, updatedName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "default", "true"),
					resource.TestCheckResourceAttr(resourceName, "execute_on_delegate", "true"),
					resource.TestCheckResourceAttr(resourceName, "region", "us-west1"),
					resource.TestCheckResourceAttr(resourceName, "gcp_project_id", "1234567"),
					resource.TestCheckResourceAttr(resourceName, "key_ring", "key_ring"),
					resource.TestCheckResourceAttr(resourceName, "key_name", "key_name"),
					resource.TestCheckResourceAttr(resourceName, "manual.0.credentials", "account."+id),
					resource.TestCheckResourceAttr(resourceName, "manual.0.delegate_selectors.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "manual.0.delegate_selectors.0", "harness-delegate"),
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

func testAccResourceConnectorGcpKMS_manual(id string, name string) string {
	return fmt.Sprintf(`
		resource "harness_platform_secret_text" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
			description = "test"
			tags = ["foo:bar"]
			secret_manager_identifier = "harnessSecretManager"
			value_type = "Inline"
			value = "secret"
		}

		resource "harness_platform_connector_gcp_kms" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
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
`, id, name)
}

func testOrgResourceConnectorGcpKMS_manual(id string, name string) string {
	return fmt.Sprintf(`
		resource "harness_platform_organization" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
		}

		resource "time_sleep" "wait_3_seconds" {
			depends_on = [harness_platform_organization.test]
			create_duration = "3s"
		}

		resource "harness_platform_secret_text" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
			description = "test"
			tags = ["foo:bar"]
			org_id = harness_platform_organization.test.id
			secret_manager_identifier = "harnessSecretManager"
			value_type = "Inline"
			value = "secret"
		}

		resource "harness_platform_connector_gcp_kms" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
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
`, id, name)
}

func testProjectResourceConnectorGcpKMS_manual(id string, name string) string {
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

		resource "time_sleep" "wait_3_seconds" {
			depends_on = [harness_platform_project.test]
			create_duration = "3s"
		}

		resource "harness_platform_secret_text" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
			description = "test"
			tags = ["foo:bar"]
			org_id = harness_platform_organization.test.id
			project_id = harness_platform_project.test.id
			secret_manager_identifier = "harnessSecretManager"
			value_type = "Inline"
			value = "secret"
		}

		resource "harness_platform_connector_gcp_kms" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
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
`, id, name)
}

func testAccResourceConnectorGcpKMS_oidc_platform(id string, name string) string {
	return fmt.Sprintf(`
		resource "harness_platform_connector_gcp_kms" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
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
`, id, name)
}

func testOrgResourceConnectorGcpKMS_oidc_platform(id string, name string) string {
	return fmt.Sprintf(`
		resource "harness_platform_organization" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
		}

		resource "time_sleep" "wait_3_seconds" {
			depends_on = [harness_platform_organization.test]
			create_duration = "3s"
		}

		resource "harness_platform_connector_gcp_kms" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
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
`, id, name)
}

func testProjectResourceConnectorGcpKMS_oidc_platform(id string, name string) string {
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

		resource "time_sleep" "wait_3_seconds" {
			depends_on = [harness_platform_project.test]
			create_duration = "3s"
		}

		resource "harness_platform_connector_gcp_kms" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
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
`, id, name)
}

func testAccResourceConnectorGcpKMS_oidc_delegate(id string, name string) string {
	return fmt.Sprintf(`
		resource "harness_platform_connector_gcp_kms" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
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
`, id, name)
}

func testOrgResourceConnectorGcpKMS_oidc_delegate(id string, name string) string {
	return fmt.Sprintf(`
		resource "harness_platform_organization" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
		}

		resource "time_sleep" "wait_3_seconds" {
			depends_on = [harness_platform_organization.test]
			create_duration = "3s"
		}

		resource "harness_platform_connector_gcp_kms" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
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
`, id, name)
}

func testProjectResourceConnectorGcpKMS_oidc_delegate(id string, name string) string {
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

		resource "time_sleep" "wait_3_seconds" {
			depends_on = [harness_platform_project.test]
			create_duration = "3s"
		}

		resource "harness_platform_connector_gcp_kms" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
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
`, id, name)
}

func testAccResourceConnectorGcpKMS_manual_default(id string, name string) string {
	return fmt.Sprintf(`
		resource "harness_platform_secret_text" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
			description = "test"
			tags = ["foo:bar"]
			secret_manager_identifier = "harnessSecretManager"
			value_type = "Inline"
			value = "secret"
		}

		resource "harness_platform_connector_gcp_kms" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
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
`, id, name)
}
