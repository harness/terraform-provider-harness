# Every kind of an integration has a mapping. This resource manages it, replacing the whole mapping: omitted fields
# get the type's defaults. Destroying the resource only removes it from the state; on Airbyte types such as
# ServiceNow it also restores the type's default field mappings.
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

  # Links records to entities that already exist instead of registering new ones. source_path is a
  # field of the emitted entity data, not of the source record.
  correlation {
    source_path = "metadata.integration_properties.ServiceNow.name"
    target_path = "name"
  }
}

# Adds a correlation to a PagerDuty kind, linking each service to the catalog entity with the same name. Omitted
# fields, such as field_mapping, keep the type's defaults.
resource "harness_platform_idp_integration_kind_mapping" "pagerduty" {
  integration_id = harness_platform_idp_integration.pagerduty.identifier
  kind           = "services"

  correlation {
    source_path = "metadata.integration_properties.PagerDuty.name"
    target_path = "name"
  }
}

# Custom integrations build entities from the payload with JQ. Don't also set a mapping block for this kind on the
# integration; that block is only sent at create.
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
