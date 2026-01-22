resource "harness_chaos_hub_v2" "project_level" {
  org_id       = "<org_id>"
  project_id   = "<project_id>"
  identity     = "<identity>"
  name         = "<name>"
  description  = "<description>"

  tags = ["<tag1>", "<tag2>"]
}
