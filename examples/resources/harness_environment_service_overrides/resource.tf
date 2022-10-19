resource "harness_environment_service_overrides" "test" {
  identifier = "%[1]s"
  org_id     = harness_platform_organization.test.id
  project_id = harness_platform_project.test.id
  env_id     = harness_platform_environment.test.id
  service_id = harness_platform_service.test.id
  yaml       = <<-EOT
        serviceOverrides:
          environmentRef: ${harness_platform_environment.test.id}
          serviceRef: ${harness_platform_service.test.id}
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
