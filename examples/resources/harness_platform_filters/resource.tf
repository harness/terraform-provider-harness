resource "harness_platform_filters" "test" {
			identifier = "identifier"
			name = "name"
			org_id     = "org_id"
      project_id = "project_id"
			type = "Connector"
			filter_properties {
				 tags = ["foo:bar"]
         filter_type = "Connector"
      }
      filter_visibility = "EveryOne"
		}
