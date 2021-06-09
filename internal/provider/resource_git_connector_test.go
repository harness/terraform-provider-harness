package provider

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/micahlmartin/terraform-provider-harness/internal/client"
	"github.com/micahlmartin/terraform-provider-harness/internal/utils"
	"github.com/stretchr/testify/require"
)

// func TestStuff(t *testing.T) {
// 	wtf(nil, nil)
// 	fmt.Println(me)
// }

func TestAccResourceGitConnector(t *testing.T) {

	var (
		name         = fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(12))
		resourceName = "harness_git_connector.test"
		connector    client.GitConnector
	)

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		CheckDestroy:      testAccGitConnectorDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceGitConnector(name, false),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "url", "https://github.com/micahlmartin/harness-demo"),
					resource.TestCheckResourceAttr(resourceName, "branch", "master"),
					resource.TestCheckResourceAttr(resourceName, "webhook_url", ""),
					resource.TestCheckResourceAttr(resourceName, "url_type", client.GitUrlTypes.Repo),
					testAccCheckGitConnectorExists(t, resourceName, &connector),
				),
			},
			{
				Config: testAccResourceGitConnector(name, true),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "url", "https://github.com/micahlmartin/harness-demo"),
					resource.TestCheckResourceAttr(resourceName, "branch", "master"),
					resource.TestCheckResourceAttrSet(resourceName, "webhook_url"),
					resource.TestCheckResourceAttr(resourceName, "url_type", client.GitUrlTypes.Repo),
				),
			},
		},
	})
}

func TestAccResourceGitConnector_invalid_urltype(t *testing.T) {

	var (
		name = fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(12))
	)

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		CheckDestroy:      testAccGitConnectorDestroy,
		Steps: []resource.TestStep{
			{
				Config:      testAccResourceGitConnector_invalid_urltype(name),
				ExpectError: regexp.MustCompile("invalid value badvalue. Must be one of ACCOUNT or REPO"),
			},
		},
	})
}

func testAccCheckGitConnectorExists(t *testing.T, name string, connector *client.GitConnector) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]
		require.True(t, ok)
		require.NotEmpty(t, rs.Primary.ID)

		p := testAccProvider

		client := p.Meta().(*client.ApiClient)

		conn, err := client.Connectors().GetGitConnectorById(rs.Primary.ID)
		require.NoError(t, err)
		require.NotNil(t, conn.CustomCommitDetails)
		require.Len(t, conn.DelegateSelectors, 1)

		return nil
	}
}

func testAccGitConnectorDestroy(s *terraform.State) error {
	p := testAccProvider
	client := p.Meta().(*client.ApiClient)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "harness_git_connector" {
			continue
		}

		// Try to find the resource
		conn, err := client.Connectors().GetGitConnectorById(rs.Primary.ID)
		if err == nil {
			if conn != nil {
				return fmt.Errorf("Found git connector: %s", rs.Primary.ID)
			}
		}
	}

	return nil
}

func testAccResourceGitConnector(name string, generateWebhook bool) string {
	return fmt.Sprintf(`
		resource "harness_encrypted_text" "test" {
			name = "%[1]s"
			value = "foo"
		}

		resource "harness_git_connector" "test" {
			name = "%[1]s"
			url = "https://github.com/micahlmartin/harness-demo"
			branch = "master"
			generate_webhook_url = %[2]t
			password_secret_id = harness_encrypted_text.test.id
			url_type = "REPO"
			username = "someuser"

			commit_details {
				author_email_id = "user@example.com"
				author_name = "some user"
				message = "commit message here"
			}

			delegate_selectors = ["primary"]
		}	
`, name, generateWebhook)
}

func testAccResourceGitConnector_invalid_urltype(name string) string {
	return fmt.Sprintf(`
		resource "harness_encrypted_text" "test" {
			name = "%[1]s"
			value = "foo"
		}

		resource "harness_git_connector" "test" {
			name = "%[1]s"
			url = "https://github.com/micahlmartin/harness-demo"
			branch = "master"
			password_secret_id = harness_encrypted_text.test.id
			url_type = "badvalue"
			username = "someuser"
		}	
`, name)
}

// func testAccResourceEncryptedText_UsageScopes(name string) string {
// 	// nonprod :=
// 	return fmt.Sprintf(`
// 	resource "harness_encrypted_text" "usage_scope_test" {
// 		name = "%s"
// 		value = "someval"

// 		usage_scope {
// 			application_filter_type = "ALL"
// 			environment_filter_type = "PRODUCTION_ENVIRONMENTS"
// 		}

// 		usage_scope {
// 			application_filter_type = "ALL"
// 			environment_filter_type = "NON_PRODUCTION_ENVIRONMENTS"
// 		}
// 	}
// 	`, name)
// }

// func testAccResourceEncryptedText_usageScopes_update(name string) string {
// 	// nonprod :=
// 	return fmt.Sprintf(`
// 	resource "harness_encrypted_text" "usage_scope_test" {
// 		name = "%s"
// 		value = "someval"

// 		usage_scope {
// 			application_filter_type = "ALL"
// 			environment_filter_type = "PRODUCTION_ENVIRONMENTS"
// 		}
// 	}
// 	`, name)
// }

// func testAccResourceEncryptedText_update_secretmanagerid(name string, value string, secretManagerId string) string {
// 	return fmt.Sprintf(`
// 		resource "harness_encrypted_text" "test" {
// 			name = "%s"
// 			value = "%s"
// 			secret_manager_id = "%s"
// 		}
// `, name, value, secretManagerId)
// }

// func TestConversion(t *testing.T) {

// 	i := make([]interface{}, 1)
// 	i = append(i, "test")

// 	var s []string

// 	s = i.([]string)

// 	fmt.Println(s)
// }
