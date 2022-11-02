resource "harness_platform_environment_service_overrides" "example" {
  identifier = "identifier"
  org_id     = "orgIdentifier"
  project_id = "projectIdentifier"
  env_id     = "environmentIdentifier"
  service_id = "serviceIdentifier"
  yaml       = <<-EOT
        serviceOverrides:
          environmentRef: environmentIdentifier
          serviceRef: serviceIdentifier
          variables:
           - name: asda
             type: String
             value: asddad
          manifests:
             - manifest:
                 identifier: manifestEnv
                 type: Values
                 spec:
                   store:
                     type: Git
                     spec:
                       connectorRef: <+input>
                       gitFetchType: Branch
                       paths:
                         - file1
                       repoName: <+input>
                       branch: master
          configFiles:
             - configFile:
                 identifier: configFileEnv
                 spec:
                   store:
                     type: Harness
                     spec:
                       files:
                         - account:/Add-ons/svcOverrideTest
                       secretFiles: []
		  EOT
}
