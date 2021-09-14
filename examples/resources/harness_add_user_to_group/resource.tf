data "harness_user" "example_user" {
  email = "testuser@example.com"
}

resource "harness_user_group" "admin" {
  name = "admin"
}

resource "harness_add_user_to_group" "example_add_user_to_groups" {
  group_id = harness_user_group.admin.id
  user_id  = data.harness_user.test.id
}
