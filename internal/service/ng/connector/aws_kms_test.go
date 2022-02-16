package connector_test

import (
	"fmt"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceConnector_awskms_inherit(t *testing.T) {

	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	name := id
	updatedName := fmt.Sprintf("%s_updated", name)
	resourceName := "harness_connector.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccConnectorDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceConnector_awskms_inherit(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "aws_kms.0.delegate_selectors.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "aws_kms.0.credentials.0.inherit_from_delegate", "true"),
				),
			},
			{
				Config: testAccResourceConnector_awskms_inherit(id, updatedName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "aws_kms.0.delegate_selectors.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "aws_kms.0.credentials.0.inherit_from_delegate", "true"),
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

func TestAccResourceConnector_awskms_manual(t *testing.T) {

	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	name := id
	updatedName := fmt.Sprintf("%s_updated", name)
	resourceName := "harness_connector.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccConnectorDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceConnector_awskms_manual(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "aws_kms.0.delegate_selectors.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "aws_kms.0.credentials.0.manual.0.secret_key_ref", "account.acctest_sumo_access_id"),
					resource.TestCheckResourceAttr(resourceName, "aws_kms.0.credentials.0.manual.0.access_key_ref", "account.acctest_appd_password"),
				),
			},
			{
				Config: testAccResourceConnector_awskms_manual(id, updatedName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "aws_kms.0.delegate_selectors.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "aws_kms.0.credentials.0.manual.0.secret_key_ref", "account.acctest_sumo_access_id"),
					resource.TestCheckResourceAttr(resourceName, "aws_kms.0.credentials.0.manual.0.access_key_ref", "account.acctest_appd_password"),
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

func TestAccResourceConnector_awskms_assumerole(t *testing.T) {

	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	name := id
	updatedName := fmt.Sprintf("%s_updated", name)
	resourceName := "harness_connector.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccConnectorDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceConnector_awskms_assumerole(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "aws_kms.0.delegate_selectors.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "aws_kms.0.credentials.0.assume_role.0.role_arn", "somerolearn"),
					resource.TestCheckResourceAttr(resourceName, "aws_kms.0.credentials.0.assume_role.0.external_id", "externalid"),
					resource.TestCheckResourceAttr(resourceName, "aws_kms.0.credentials.0.assume_role.0.duration", "900"),
				),
			},
			{
				Config: testAccResourceConnector_awskms_assumerole(id, updatedName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "aws_kms.0.delegate_selectors.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "aws_kms.0.credentials.0.assume_role.0.role_arn", "somerolearn"),
					resource.TestCheckResourceAttr(resourceName, "aws_kms.0.credentials.0.assume_role.0.external_id", "externalid"),
					resource.TestCheckResourceAttr(resourceName, "aws_kms.0.credentials.0.assume_role.0.duration", "900"),
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

func testAccResourceConnector_awskms_inherit(id string, name string) string {
	return fmt.Sprintf(`
		resource "harness_connector" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
			description = "test"
			tags = ["foo:bar"]

			aws_kms {
				arn_ref = "account.acctest_sumo_access_id"
				region = "us-east-1"
				delegate_selectors = ["harness-delegate"]
				credentials {
					inherit_from_delegate = true
				}
			}
		}
`, id, name)
}

func testAccResourceConnector_awskms_manual(id string, name string) string {
	return fmt.Sprintf(`
		resource "harness_connector" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
			description = "test"
			tags = ["foo:bar"]

			aws_kms {
				arn_ref = "account.acctest_sumo_access_id"
				region = "us-east-1"
				delegate_selectors = ["harness-delegate"]
				credentials {
					manual {
						secret_key_ref = "account.acctest_sumo_access_id"
						access_key_ref = "account.acctest_appd_password"
					}
				}
			}
		}
`, id, name)
}

func testAccResourceConnector_awskms_assumerole(id string, name string) string {
	return fmt.Sprintf(`
		resource "harness_connector" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
			description = "test"
			tags = ["foo:bar"]

			aws_kms {
				arn_ref = "account.acctest_sumo_access_id"
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
		}
`, id, name)
}
