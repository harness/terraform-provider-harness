resource "harness_platform_ccm_filters" "test" {
  identifier = "identifier"
  name       = "name"
  org_id     = "org_id"
  project_id = "project_id"
  type       = "CCMRecommendation"
  filter_properties {
    tags        = ["foo:bar"]
    filter_type = "CCMRecommendation"
  }
  filter_visibility = "EveryOne"
}
