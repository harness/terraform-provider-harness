package provider

import (
	"fmt"
	"strings"
	"testing"

	"github.com/harness-io/harness-go-sdk/harness/api"
	"github.com/harness-io/harness-go-sdk/harness/api/graphql"
	"github.com/harness-io/harness-go-sdk/harness/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/stretchr/testify/require"
)

func init() {
	resource.AddTestSweepers("harness_git_connector", &resource.Sweeper{
		Name: "harness_git_connector",
		F:    testSweepGitConnectors,
	})
}

func testSweepGitConnectors(r string) error {
	c := testAccGetApiClientFromProvider()

	limit := 100
	offset := 0
	hasMore := true

	for hasMore {

		connectors, _, err := c.Connectors().ListGitConnectors(limit, offset)
		if err != nil {
			return err
		}

		for _, conn := range connectors {
			if strings.HasPrefix(conn.Name, "test_") {
				if err = c.Connectors().DeleteConnector(conn.Id); err != nil {
					return err
				}
			}
		}

		hasMore = len(connectors) == limit
	}

	return nil
}

func TestAccResourceGitConnector(t *testing.T) {

	var (
		name         = fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(12))
		updatedName  = fmt.Sprintf("%s_updated", name)
		resourceName = "harness_git_connector.test"
	)

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		CheckDestroy:      testAccGitConnectorDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceGitConnector_default(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "webhook_url"),
					resource.TestCheckResourceAttr(resourceName, "generate_webhook_url", "true"),
					testAccCheckGitConnectorExists(t, resourceName, name),
				),
			},
			{
				Config: testAccResourceGitConnector_default(updatedName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGitConnectorUpdated(t, resourceName, updatedName),
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

func TestAccResourceGitConnector_DeleteUnderlyingResource(t *testing.T) {

	var (
		name         = fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(12))
		resourceName = "harness_git_connector.test"
	)

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config:                    testAccResourceGitConnector(name, false, false, false),
				PreventPostDestroyRefresh: true,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGitConnectorExists(t, resourceName, name),
				),
			},
			{
				PreConfig: func() {
					testAccConfigureProvider()
					c := testAccProvider.Meta().(*api.Client)
					conn, err := c.Connectors().GetGitConnectorByName(name)
					require.NoError(t, err)
					require.NotNil(t, conn)

					err = c.Connectors().DeleteConnector(conn.Id)
					require.NoError(t, err)
				},
				Config:             testAccResourceGitConnector(name, true, true, true),
				PlanOnly:           true,
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func testAccGetGitConnector(resourceName string, state *terraform.State) (*graphql.GitConnector, error) {
	r := testAccGetResource(resourceName, state)
	c := testAccGetApiClientFromProvider()
	id := r.Primary.ID

	return c.Connectors().GetGitConnectorById(id)
}

func testAccCheckGitConnectorExists(t *testing.T, resourceName string, connectorName string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		conn, err := testAccGetGitConnector(resourceName, state)
		require.NoError(t, err)
		require.Equal(t, connectorName, conn.Name)
		require.Equal(t, "https://github.com/micahlmartin/harness-demo", conn.Url)
		require.Equal(t, "master", conn.Branch)
		require.Equal(t, graphql.GitUrlTypes.Repo, conn.UrlType)
		return nil
	}
}

func testAccCheckGitConnectorUpdated(t *testing.T, resourceName string, connectorName string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		conn, err := testAccGetGitConnector(resourceName, state)
		require.NoError(t, err)
		require.Equal(t, connectorName, conn.Name)
		require.NotNil(t, conn.CustomCommitDetails)
		require.Empty(t, conn.CustomCommitDetails.AuthorEmailId)
		require.Empty(t, conn.CustomCommitDetails.AuthorName)
		require.Empty(t, conn.CustomCommitDetails.CommitMessage)
		require.Nil(t, conn.DelegateSelectors)
		return nil
	}
}

func testAccGitConnectorDestroy(resourceName string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		conn, _ := testAccGetGitConnector(resourceName, state)
		if conn != nil {
			return fmt.Errorf("Found git connector: %s", conn.Id)
		}

		return nil
	}
}

func testAccGetDefaultDelegeteSelectors() string {
	return `
		delegate_selectors = ["primary"]
	`
}

func testAccGetCommitDetails() string {
	return `
		commit_details {
			author_email_id = "user@example.com"
			author_name = "some user"
			message = "commit message here"
		}
	`
}

func testAccResourceGitConnector_default(name string) string {

	return fmt.Sprintf(`
		data "harness_secret_manager" "test" {
			default = true
		}

		resource "harness_encrypted_text" "test" {
			name 							= "%[1]s"
			value 					  = "foo"
			secret_manager_id = data.harness_secret_manager.test.id
		}

		resource "harness_git_connector" "test" {
			name = "%[1]s"
			url = "https://github.com/micahlmartin/harness-demo"
			branch = "master"
			generate_webhook_url = true
			password_secret_id = harness_encrypted_text.test.id
			url_type = "REPO"
			username = "someuser"
		}	
`, name)
}

func testAccResourceGitConnector(name string, generateWebhook bool, withCommitDetails bool, withDelegateSelectors bool) string {

	var (
		delegateSelectors string
		commitDetails     string
	)

	if withDelegateSelectors {
		delegateSelectors = testAccGetDefaultDelegeteSelectors()
	}

	if withCommitDetails {
		commitDetails = testAccGetCommitDetails()
	}

	return fmt.Sprintf(`
		data "harness_secret_manager" "test" {
			default = true
		}

		resource "harness_encrypted_text" "test" {
			name 							= "%[1]s"
			value 					  = "foo"
			secret_manager_id = data.harness_secret_manager.test.id
		}

		resource "harness_git_connector" "test" {
			name = "%[1]s"
			url = "https://github.com/micahlmartin/harness-demo"
			branch = "master"
			generate_webhook_url = %[2]t
			password_secret_id = harness_encrypted_text.test.id
			url_type = "REPO"
			username = "someuser"

			%[3]s

			%[4]s
		}	
`, name, generateWebhook, commitDetails, delegateSelectors)
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
