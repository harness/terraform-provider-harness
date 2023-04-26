#For account level template
data "harness_platform_template" "example" {
  identifier = "identifier"
  version    = "version"
}

#For org level template
data "harness_platform_template" "example1" {
  identifier = "identifier"
  version    = "version"
  org_id     = "org_id"
}

#For project level template
data "harness_platform_template" "example2" {
  identifier = "identifier"
  version    = "version"
  org_id     = "org_id"
  project_id = "project_id"
}
