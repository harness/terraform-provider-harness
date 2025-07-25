resource "harness_chaos_hub" "example" {
  org_id       = "<org_id>"
  project_id   = "<project_id>"
  name         = "<name>"
  description  = "<description>"
  connector_id = "<connector_id>"
  repo_branch  = "<repo_branch>"
  repo_name    = "<repo_name>"
  is_default   = false

  tags = ["<tag1>", "<tag2>"]
}
