package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceUser_Email(t *testing.T) {

	var (
		email        = "micahlmartin+testing@gmail.com"
		id           = "XyhrLMYeSXClAhsLUcyyyw"
		resourceName = "data.harness_user.test"
	)

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceUserByEmail(email),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "email", email),
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttrSet(resourceName, "is_email_verified"),
					resource.TestCheckResourceAttrSet(resourceName, "is_imported_from_identity_provider"),
					resource.TestCheckResourceAttrSet(resourceName, "is_password_expired"),
					resource.TestCheckResourceAttrSet(resourceName, "is_two_factor_auth_enabled"),
					resource.TestCheckResourceAttrSet(resourceName, "is_user_locked"),
				),
			},
		},
	})
}

func TestAccDataSourceUser_Id(t *testing.T) {

	var (
		email        = "micahlmartin+testing@gmail.com"
		id           = "XyhrLMYeSXClAhsLUcyyyw"
		resourceName = "data.harness_user.test"
	)

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceUserById(id),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "email", email),
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttrSet(resourceName, "is_email_verified"),
					resource.TestCheckResourceAttrSet(resourceName, "is_imported_from_identity_provider"),
					resource.TestCheckResourceAttrSet(resourceName, "is_password_expired"),
					resource.TestCheckResourceAttrSet(resourceName, "is_two_factor_auth_enabled"),
					resource.TestCheckResourceAttrSet(resourceName, "is_user_locked"),
				),
			},
		},
	})
}

func testAccDataSourceUserByEmail(email string) string {
	return fmt.Sprintf(`
		data "harness_user" "test" {
			email = "%[1]s"
		}
	`, email)
}

func testAccDataSourceUserById(id string) string {
	return fmt.Sprintf(`
		data "harness_user" "test" {
			id = "%[1]s"
		}
	`, id)
}
