data "harness_platform_dashboard_folders" "all" {
}

# Or filter by name
data "harness_platform_dashboard_folders" "filtered" {
  name = "my-folder"
}
