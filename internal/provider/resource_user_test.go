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
	resource.AddTestSweepers("harness_user", &resource.Sweeper{
		Name: "harness_user",
		F:    testSweepUsers,
	})
}

func TestAccResourceUser(t *testing.T) {

	expectedName := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))
	expectedEmail := fmt.Sprintf("%s@example.com", expectedName)
	updatedName := fmt.Sprintf("%s_updated", expectedName)
	resourceName := "harness_user.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		CheckDestroy:      testAccUserDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceUser(expectedName, expectedEmail),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", expectedName),
					testAccUserCreation(t, resourceName, expectedEmail),
				),
			},
			{
				Config: testAccResourceUser(updatedName, expectedEmail),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
					testAccUserCreation(t, resourceName, expectedEmail),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: func(s *terraform.State) (string, error) {
					primary := s.RootModule().Resources[resourceName].Primary
					return primary.Attributes["email"], nil
				},
			},
		},
	})
}

func TestAccResourceUser_DeleteUnderlyingResource(t *testing.T) {

	expectedName := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(12))
	expectedEmail := fmt.Sprintf("%s@example.com", expectedName)
	resourceName := "harness_user.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		CheckDestroy:      testAccUserDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceUser(expectedName, expectedEmail),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", expectedName),
					testAccUserCreation(t, resourceName, expectedEmail),
				),
			},
			{
				PreConfig: func() {
					testAccConfigureProvider()
					c := testAccProvider.Meta().(*api.Client)

					usr, err := c.Users().GetUserByEmail(expectedEmail)
					require.NoError(t, err)
					require.NotNil(t, usr)

					err = c.Users().DeleteUser(usr.Id)
					require.NoError(t, err)
				},
				Config:             testAccResourceUser(expectedName, expectedEmail),
				PlanOnly:           true,
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func testAccUserCreation(t *testing.T, resourceName string, email string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		user, err := testAccGetUser(resourceName, state)
		require.NoError(t, err)
		require.NotNil(t, user)
		require.Equal(t, strings.ToLower(email), user.Email)

		return nil
	}
}

func testAccGetUser(resourceName string, state *terraform.State) (*graphql.User, error) {
	r := testAccGetResource(resourceName, state)
	c := testAccGetApiClientFromProvider()
	email := r.Primary.Attributes["email"]

	return c.Users().GetUserByEmail(email)
}

func testAccUserDestroy(resourceName string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		user, _ := testAccGetUser(resourceName, state)
		if user != nil {
			return fmt.Errorf("Found user: %s", user.Id)
		}

		return nil
	}
}

func testAccResourceUser(name string, email string) string {
	return fmt.Sprintf(`
		resource "harness_user" "test" {
			name = "%[1]s"
			email = "%[2]s"
		}
`, name, email)
}

func testSweepUsers(r string) error {
	c := testAccGetApiClientFromProvider()

	limit := 100
	offset := 0
	hasMore := true

	for hasMore {

		users, _, err := c.Users().ListUsers(limit, offset)
		if err != nil {
			return err
		}

		for _, user := range users {
			// Only delete users that have an email that starts with 'test'
			if strings.HasPrefix(user.Email, "test") {
				if err = c.Users().DeleteUser(user.Id); err != nil {
					return err
				}
			}
		}

		hasMore = len(users) == limit
	}

	return nil
}
