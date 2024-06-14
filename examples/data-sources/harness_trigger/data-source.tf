data "harness_trigger" "example_by_name" {
  app_id = "app_id"
  name   = "name"
}

data "harness_trigger" "example_by_id" {
  id = "trigger_id"
}
