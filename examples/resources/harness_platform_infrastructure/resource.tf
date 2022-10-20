resource "harness_platform_infrastructure" "test" {
  identifier      = "%[1]s"
  name            = "%[2]s"
  org_id          = harness_platform_organization.test.id
  project_id      = harness_platform_project.test.id
  env_id          = harness_platform_environment.test.id
  type            = "KubernetesDirect"
  deployment_type = "Kubernetes"
  yaml            = <<-EOT
			   infrastructureDefinition:
         name: "%[2]s"
         identifier: "%[1]s"
         description: ""
         tags:
           asda: ""
         orgIdentifier: ${harness_platform_organization.test.id}
         projectIdentifier: ${harness_platform_project.test.id}
         environmentRef: ${harness_platform_environment.test.id}
         deploymentType: Kubernetes
         type: KubernetesDirect
         spec:
          connectorRef: account.gfgf
          namespace: asdasdsa
          releaseName: release-<+INFRA_KEY>
          allowSimultaneousDeployments: false
      EOT
}
