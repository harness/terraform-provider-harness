resource "harness_platform_idp_environment_blueprint" "test" {
  identifier  = "identifier"
  version     = "v1.0.0"
  stable      = true
  deprecated  = false
  description = "description"
  yaml        = <<-EOT
		apiVersion: harness.io/v1
		kind: EnvironmentBlueprint
		type: long-lived
		identifier: identifier
		name: name
		owner: group:account/_account_all_users
		metadata:
		  description: description
		spec:
		  entities:
		  - identifier: git
		    backend:
		      type: HarnessCD
		      steps:
		        apply:
		          pipeline: gittest
		          branch: main
		        destroy:
		          pipeline: gittest
		          branch: not-main
		  ownedBy:
		  - group:account/_account_all_users
		EOT
}
