resource "harness_platform_environment_service_overrides" "example" {
  identifier = "%[1]s"
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
                    identifier: manifest1
                    type: K8sManifest
                    spec:
                      store:
                        type: Github
                        spec:
                          connectorRef: <+input>
                          gitFetchType: Branch
                          paths:
                            - files1
                          repoName: <+input>
                          branch: master
                      skipResourceVersioning: false
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
