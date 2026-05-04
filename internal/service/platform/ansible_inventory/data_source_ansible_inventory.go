package ansible_inventory

import (
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceAnsibleInventory() *schema.Resource {
	resource := &schema.Resource{
		Description: "Data source for retrieving a Harness IaCM Ansible Inventory.",

		ReadContext: resourceAnsibleInventoryRead,

		Schema: inventorySchema(true),
	}
	helpers.SetProjectLevelDataSourceSchemaIdentifierRequired(resource.Schema)
	return resource
}
