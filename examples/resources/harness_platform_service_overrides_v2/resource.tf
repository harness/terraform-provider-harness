resource "harness_platform_service_overrides_v2" "test" {
  identifier  = "identifier"
  org_id      = "orgIdentifier"
  project_id  = "projectIdentifier"
  env_id      = "environmentIdentifier"
  service_id  = "serviceIdentifier"
  infra_id    = "infraIdentifier"
  cluster_id  = "clusterIdentifier"
  type        = "ENV_SERVICE_OVERRIDE"
  spec        = <<-EOT
    {
      "variables": [
        {
          "name": "v1",
          "type": "String",
          "value": "val1"
        }
      ],
      "configFiles": [
        {
          "configFile": {
            "identifier": "sampleConfigFile",
            "spec": {
              "store": {
                "type": "Harness",
                "spec": {
                  "files": [
                    "/launchTemplate2"
                  ]
                }
              }
            }
          }
        }
      ],
      "manifests": [
        {
          "manifest": {
            "identifier": "sampleManifest",
            "type": "AsgLaunchTemplate",
            "spec": {
              "store": {
                "type": "Harness",
                "spec": {
                  "files": [
                    "/launchTemplate1"
                  ]
                }
              }
            }
          }
        }
      ]
    }
  EOT
}
