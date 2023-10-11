resource "harness_platform_service_overrides_v2" "test" {
  org_id     = "orgIdentifier"
  project_id = "projectIdentifier"
  env_id     = "environmentIdentifier"
  service_id = "serviceIdentifier"
  infra_id   = "infraIdentifier"
  cluster_id = "clusterIdentifier"
  type       = "INFRA_SERVICE_OVERRIDE"
  yaml       = <<-EOT
    variables:
      - name: var1
        type: String
        value: val1
    configFiles:
      - configFile:
          identifier: sampleConfigFile
          spec:
            store:
              type: Harness
              spec:
                files:
                  - account:/configFile1
    manifests:
      - manifest:
          identifier: sampleManifestFile
          type: Values
          spec:
            store:
              type: Harness
              spec:
                files:
                  - account:/manifestFile1
  EOT
}
