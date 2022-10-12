resource "harness_platform_environment" "example" {
  identifier  = "envidentifier"
  name        = "envname"
  org_id      = harness_platform_project.test.org_id
  project_id  = harness_platform_project.test.id
  tags        = ["foo:bar", "baz"]
  description = "envdescription"
  type        = "PreProduction"
  yaml        = <<-EOT
  environment:
         name: envname
         identifier: envidentifier
         orgIdentifier: ${harness_platform_project.test.org_id}
         projectIdentifier: ${harness_platform_project.test.id}
         type: PreProduction
         description: envdescription
         tags:
           foo: bar
           baz: ""
         variables:
           - name: envVar1
             type: String
             value: v1
             description: ""
           - name: envVar2
             type: String
             value: v2
             description: ""
         overrides:
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


