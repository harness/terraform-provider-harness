package ansible_playbook

import (
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceAnsiblePlaybook() *schema.Resource {
	resource := &schema.Resource{
		Description: "Data source for retrieving a Harness IaCM Ansible Playbook.",

		ReadContext: resourceAnsiblePlaybookRead,

		Schema: playbookSchema(true),
	}
	helpers.SetProjectLevelDataSourceSchemaIdentifierRequired(resource.Schema)
	return resource
}
