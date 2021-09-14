resource "harness_cloudprovider_gcp" "example" {
  name               = "example"
  skip_validation    = true
  delegate_selectors = ["gcp"]
}
