package user_test

import (
	"fmt"
	"strings"
	"testing"

	sdk "github.com/harness/harness-go-sdk"
	"github.com/harness/harness-go-sdk/harness/cd/graphql"
	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
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
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
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
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
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
					acctest.TestAccConfigureProvider()
					c := acctest.TestAccProvider.Meta().(*sdk.Session)

					usr, err := c.CDClient.UserClient.GetUserByEmail(expectedEmail)
					require.NoError(t, err)
					require.NotNil(t, usr)

					err = c.CDClient.UserClient.DeleteUser(usr.Id)
					require.NoError(t, err)
				},
				Config:             testAccResourceUser(expectedName, expectedEmail),
				PlanOnly:           true,
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccResourceUser_WithUserGroups(t *testing.T) {

	expectedName := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))
	expectedEmail := strings.ToLower(fmt.Sprintf("%s_%s@example.com", t.Name(), utils.RandStringBytes(4)))
	updatedName := fmt.Sprintf("%s_updated", expectedName)
	resourceName := "harness_user.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccUserDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceUser_WithUserGroups(expectedName, expectedEmail),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", expectedName),
					testAccUserCreation(t, resourceName, expectedEmail),
					func(s *terraform.State) (err error) {
						userId := s.RootModule().Resources[resourceName].Primary.ID
						groupId := s.RootModule().Resources["harness_user_group.test"].Primary.ID
						acctest.TestAccConfigureProvider()
						c := acctest.TestAccProvider.Meta().(*sdk.Session)

						limit := 100
						offset := 0
						hasMore := true

						for hasMore {
							groups, _, err := c.CDClient.UserClient.ListGroupMembershipByUserId(userId, limit, offset)
							if err != nil {
								return err
							}

							for _, group := range groups {
								if group.Id == groupId {
									return nil
								}
							}

							hasMore = len(groups) == limit
							offset += limit
						}

						return fmt.Errorf("user %s is not a member of group %s", userId, groupId)
					},
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
	r := acctest.TestAccGetResource(resourceName, state)
	c := acctest.TestAccGetApiClientFromProvider()
	email := r.Primary.Attributes["email"]

	return c.CDClient.UserClient.GetUserByEmail(email)
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

func testAccResourceUser_WithUserGroups(name string, email string) string {
	return fmt.Sprintf(`
		resource "harness_user_group" "test" {
			name = "%[1]s"
		}

		resource "harness_user" "test" {
			name = "%[1]s"
			email = "%[2]s"
			group_ids = [harness_user_group.test.id]
		}
`, name, email)
}

func testSweepUsers(r string) error {
	c := acctest.TestAccGetApiClientFromProvider()

	limit := 100
	offset := 0
	hasMore := true

	for hasMore {

		users, _, err := c.CDClient.UserClient.ListUsers(limit, offset)
		if err != nil {
			return err
		}

		for _, user := range users {
			// Only delete users that have an email that starts with 'test'
			if strings.HasPrefix(user.Email, "test") {
				if err = c.CDClient.UserClient.DeleteUser(user.Id); err != nil {
					return err
				}
			}
		}

		hasMore = len(users) == limit
	}

	return nil
}
