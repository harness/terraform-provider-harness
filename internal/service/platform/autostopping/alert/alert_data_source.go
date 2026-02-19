package alert

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceAlert() *schema.Resource {
	return &schema.Resource{
		Description: "Data source for retrieving a Harness AutoStopping alert by ID. Use the id (identifier) returned by the API when the alert was created.",
		ReadContext: resourceAlertRead,
		Schema: map[string]*schema.Schema{
			"identifier": {
				Description: "Unique identifier of the alert (same as id, populated from server).",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"name": {
				Description: "Name of the alert.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"enabled": {
				Description: "Whether the alert is enabled.",
				Type:        schema.TypeBool,
				Computed:    true,
			},
			"recipients": {
				Description: "Notification recipients (email and/or slack).",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"email": {
							Description: "List of email addresses.",
							Type:        schema.TypeList,
							Computed:    true,
							Elem:        &schema.Schema{Type: schema.TypeString},
						},
						"slack": {
							Description: "List of Slack webhook URLs or channel identifiers.",
							Type:        schema.TypeList,
							Computed:    true,
							Elem:        &schema.Schema{Type: schema.TypeString},
						},
					},
				},
			},
			"events": {
				Description: "List of event types that trigger the alert.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"applicable_to_all_rules": {
				Description: "Whether the alert applies to all AutoStopping rules.",
				Type:        schema.TypeBool,
				Computed:    true,
			},
			"rule_id_list": {
				Description: "List of AutoStopping rule IDs the alert applies to.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeInt},
			},
		},
	}
}
