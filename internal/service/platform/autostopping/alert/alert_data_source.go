package alert

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceAlertRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	id, ok := d.GetOk("identifier")
	if !ok {
		return diag.Errorf("identifier is required")
	}
	return alertReadByID(ctx, d, meta, id.(string))
}

func DataSourceAlert() *schema.Resource {
	return &schema.Resource{
		Description: "Data source for retrieving a Harness AutoStopping alert by ID.",
		ReadContext: dataSourceAlertRead,
		Schema: map[string]*schema.Schema{
			"identifier": {
				Description: "Unique identifier of the alert.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"name": {
				Description: "Name of the alert.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"enabled": {
				Description: "Whether the alert is enabled.",
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
			},
			"recipients": {
				Description: "Notification recipients (email and/or slack).",
				Type:        schema.TypeList,
				Required:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"email": {
							Description: "List of email addresses.",
							Type:        schema.TypeList,
							Optional:    true,
							Elem:        &schema.Schema{Type: schema.TypeString},
						},
						"slack": {
							Description: "List of Slack webhook URLs or channel identifiers.",
							Type:        schema.TypeList,
							Optional:    true,
							Elem:        &schema.Schema{Type: schema.TypeString},
						},
					},
				},
			},
			"events": {
				Description: "List of event types that trigger the alert.",
				Type:        schema.TypeList,
				Required:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"applicable_to_all_rules": {
				Description: "Whether the alert applies to all AutoStopping rules.",
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
			},
			"rule_id_list": {
				Description: "List of AutoStopping rule IDs the alert applies to.",
				Type:        schema.TypeList,
				Optional:    true,
				Elem:        &schema.Schema{Type: schema.TypeInt},
			},
		},
	}
}
