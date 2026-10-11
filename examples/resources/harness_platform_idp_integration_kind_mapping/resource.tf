# Every kind of an integration has a mapping. This resource manages it; destroying the
# resource leaves the mapping as it is.
resource "harness_platform_idp_integration_kind_mapping" "services" {
  integration_id = harness_platform_idp_integration.servicenow.identifier
  org_id         = "default"
  kind           = "cmdb_ci_service"

  display_name_field = "name"
  idp_kind           = "component"
  idp_type           = "service"

  # ServiceNow source fields are top-level fields of the kind's catalog stream.
  field_mapping {
    source_field = "owned_by"
    target_field = "spec.owner"
  }
  field_mapping {
    source_field = "environment"
    target_field = "metadata.environment"
  }

  # Links records to entities that already exist instead of registering new ones.
  correlation {
    source_path = "name"
    target_path = "metadata.identifier"
  }
}

# Custom integrations build entities from the payload with JQ.
resource "harness_platform_idp_integration_kind_mapping" "orders" {
  integration_id = harness_platform_idp_integration.orders.identifier
  kind           = "custom"

  ingestion {
    identifier = ".service.id"
    name       = ".service.name"
    properties = {
      "spec.owner" = ".service.team"
    }
  }
}
