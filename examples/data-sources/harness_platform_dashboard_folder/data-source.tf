# By id
data "harness_platform_dashboard_folder" "by_id" {
  id = "1234"
}

# By name (will perform a list + match internally)
data "harness_platform_dashboard_folder" "by_name" {
  name = "my-folder"
}