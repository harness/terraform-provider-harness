data "harness_chaos_load_test" "example" {
  org_id     = "default"
  project_id = "demo"
  identifier = "checkout_k6"
}
