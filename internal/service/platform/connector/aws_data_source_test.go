package connector_test

import (
	"fmt"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceConnectorAws(t *testing.T) {

	var (
		name         = fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))
		resourceName = "data.harness_platform_connector_aws.test"
	)

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceConnectorAws(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", name),
					resource.TestCheckResourceAttr(resourceName, "identifier", name),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "inherit_from_delegate.0.delegate_selectors.#", "1"),
				),
			},
		},
	})
}

func TestAccDataSourceConnectorAwsFullJitterBackOff(t *testing.T) {

	var (
		name         = fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))
		resourceName = "data.harness_platform_connector_aws.test"
	)

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceConnectorAwsFullJitterBackOff(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", name),
					resource.TestCheckResourceAttr(resourceName, "identifier", name),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "inherit_from_delegate.0.delegate_selectors.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "full_jitter_backoff_strategy.0.retry_count", "3"),
					resource.TestCheckResourceAttr(resourceName, "full_jitter_backoff_strategy.0.base_delay", "10"),
					resource.TestCheckResourceAttr(resourceName, "full_jitter_backoff_strategy.0.max_backoff_time", "65"),
				),
			},
		},
	})
}

func TestAccDataSourceConnectorAwsEqualJitterBackOff(t *testing.T) {

	var (
		name         = fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))
		resourceName = "data.harness_platform_connector_aws.test"
	)

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceConnectorAwsEqualJitterBackOff(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", name),
					resource.TestCheckResourceAttr(resourceName, "identifier", name),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "inherit_from_delegate.0.delegate_selectors.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "equal_jitter_backoff_strategy.0.retry_count", "3"),
					resource.TestCheckResourceAttr(resourceName, "equal_jitter_backoff_strategy.0.base_delay", "10"),
					resource.TestCheckResourceAttr(resourceName, "equal_jitter_backoff_strategy.0.max_backoff_time", "65"),
				),
			},
		},
	})
}

func TestAccDataSourceConnectorAwsFixedDelayBackOff(t *testing.T) {

	var (
		name         = fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))
		resourceName = "data.harness_platform_connector_aws.test"
	)

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceConnectorAwsFixedDelayBackOff(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", name),
					resource.TestCheckResourceAttr(resourceName, "identifier", name),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "inherit_from_delegate.0.delegate_selectors.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "fixed_delay_backoff_strategy.0.retry_count", "3"),
					resource.TestCheckResourceAttr(resourceName, "fixed_delay_backoff_strategy.0.fixed_backoff", "10"),
				),
			},
		},
	})
}

func testAccDataSourceConnectorAws(name string) string {
	return fmt.Sprintf(`
		resource "harness_platform_connector_aws" "test" {
			identifier = "%[1]s"
			name = "%[1]s"
			description = "test"
			tags = ["foo:bar"]

			inherit_from_delegate {
				delegate_selectors = ["harness-delegate"]
			}
		}

		data "harness_platform_connector_aws" "test" {
			identifier = harness_platform_connector_aws.test.identifier
		}
	`, name)
}

func testAccDataSourceConnectorAwsEqualJitterBackOff(name string) string {
	return fmt.Sprintf(`
		resource "harness_platform_connector_aws" "test" {
			identifier = "%[1]s"
			name = "%[1]s"
			description = "test"
			tags = ["foo:bar"]

			inherit_from_delegate {
				delegate_selectors = ["harness-delegate"]
			}
			equal_jitter_backoff_strategy {
				base_delay = 10
				max_backoff_time = 65
				retry_count = 3
			}
		}

		data "harness_platform_connector_aws" "test" {
			identifier = harness_platform_connector_aws.test.identifier
		}
	`, name)
}

func testAccDataSourceConnectorAwsFullJitterBackOff(name string) string {
	return fmt.Sprintf(`
		resource "harness_platform_connector_aws" "test" {
			identifier = "%[1]s"
			name = "%[1]s"
			description = "test"
			tags = ["foo:bar"]

			inherit_from_delegate {
				delegate_selectors = ["harness-delegate"]
			}
			full_jitter_backoff_strategy {
				base_delay = 10
				max_backoff_time = 65
				retry_count = 3
			}
		}

		data "harness_platform_connector_aws" "test" {
			identifier = harness_platform_connector_aws.test.identifier
		}
	`, name)
}

func testAccDataSourceConnectorAwsFixedDelayBackOff(name string) string {
	return fmt.Sprintf(`
		resource "harness_platform_connector_aws" "test" {
			identifier = "%[1]s"
			name = "%[1]s"
			description = "test"
			tags = ["foo:bar"]

			inherit_from_delegate {
				delegate_selectors = ["harness-delegate"]
			}
			fixed_delay_backoff_strategy {
				fixed_backoff = 10
				retry_count = 3
			}
		}

		data "harness_platform_connector_aws" "test" {
			identifier = harness_platform_connector_aws.test.identifier
		}
	`, name)
}
