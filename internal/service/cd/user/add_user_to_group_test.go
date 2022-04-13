package user_test

import (
	"fmt"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/stretchr/testify/require"
)

func TestAccResourceAddUserToGroup(t *testing.T) {

	groupName := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(12))
	resourceName := "harness_add_user_to_group.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccAddUserToGroupDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceAddUserToGroup(groupName, "first"),
				Check: resource.ComposeTestCheckFunc(
					testAccUserInGroup(t, "data.harness_user.test", "harness_user_group.first"),
				),
			},
			{
				Config: testAccResourceAddUserToGroup(groupName, "second"),
				Check: resource.ComposeTestCheckFunc(
					testAccUserInGroup(t, "data.harness_user.test", "harness_user_group.second"),
					testAccUserNotInGroup(t, "data.harness_user.test", "harness_user_group.first"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: func(s *terraform.State) (string, error) {
					primary := s.RootModule().Resources[resourceName].Primary
					user_id := primary.Attributes["user_id"]
					group_id := primary.Attributes["group_id"]
					return fmt.Sprintf("%s/%s", user_id, group_id), nil
				},
			},
		},
	})
}

func testAccUserInGroup(t *testing.T, resourceName string, groupResourceName string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		if ok := testAccGetUserInGroupStatus(t, state, resourceName, groupResourceName); !ok {
			return fmt.Errorf("User not in group")
		}
		return nil
	}
}

func testAccUserNotInGroup(t *testing.T, resourceName string, groupResourceName string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		if ok := testAccGetUserInGroupStatus(t, state, resourceName, groupResourceName); ok {
			return fmt.Errorf("User still in group")
		}
		return nil
	}
}

func testAccGetUserInGroupStatus(t *testing.T, state *terraform.State, resourceName string, groupResourceName string) bool {
	user, err := testAccGetUser(resourceName, state)
	require.NoError(t, err)
	require.NotNil(t, user)

	group, err := testAccGetUserGroup(groupResourceName, state)
	require.NoError(t, err)
	require.NotNil(t, group)

	c := acctest.TestAccGetApiClientFromProvider().CDClient
	ok, err := c.UserClient.IsUserInGroup(user.Id, group.Id)
	require.NoError(t, err)

	return ok
}

func testAccAddUserToGroupDestroy(resourceName string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		r := acctest.TestAccGetResource(resourceName, state)
		groupId := r.Primary.Attributes["group_id"]
		userId := r.Primary.Attributes["user_id"]

		c := acctest.TestAccGetApiClientFromProvider().CDClient
		ok, err := c.UserClient.IsUserInGroup(userId, groupId)
		if err != nil {
			return err
		}

		if ok {
			return fmt.Errorf("User still in group")
		}

		return nil
	}
}

func testAccResourceAddUserToGroup(groupName string, groupResourceName string) string {
	return fmt.Sprintf(`
		data "harness_user" "test" {
			email = "micahlmartin+testing@gmail.com"
		}

		resource "harness_user_group" "first" {
			name = "%[1]s_first"
			description = "test"
		}

		resource "harness_user_group" "second" {
			name = "%[1]s_second"
			description = "test"
		}

		resource "harness_add_user_to_group" "test" {
			group_id = harness_user_group.%[2]s.id
			user_id = data.harness_user.test.id
		}
`, groupName, groupResourceName)
}
