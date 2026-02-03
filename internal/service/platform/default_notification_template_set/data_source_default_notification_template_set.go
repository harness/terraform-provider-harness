package default_notification_template_set

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func DataDefaultNotificationTemplateSet() *schema.Resource {
	return &schema.Resource{
		Description: "Data source for retrieving a Default Notification Template Set.",
		ReadContext: resourceDefaultNotificationTemplateSetRead,
		Schema: map[string]*schema.Schema{
			"org_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Unique identifier of the organization.",
			},
			"project_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Unique identifier of the project.",
			},
			"org": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Unique identifier of the organization. Deprecated: Use org_id instead.",
				Deprecated:  "This field is deprecated and will be removed in a future release. Please use 'org_id' instead.",
			},
			"project": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Unique identifier of the project. Deprecated: Use project_id instead.",
				Deprecated:  "This field is deprecated and will be removed in a future release. Please use 'project_id' instead.",
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of Default Notification Template Set",
			},
			"identifier": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Identifier of Default Notification Template Set",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Description for Default Notification Template Set",
			},
			"notification_entity": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Type of the entity (e.g. PIPELINE, SERVICE, etc.)",
			},
			"notification_channel_type": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Type of channel (e.g. SLACK, EMAIL, etc.)",
			},
			"event_template_configuration_set": {
				Type:        schema.TypeList,
				Required:    true,
				Description: "Set of event-template configurations",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"notification_events": {
							Type:        schema.TypeList,
							Required:    true,
							Description: "List of notification events like PIPELINE_START",
							Elem:        &schema.Schema{Type: schema.TypeString},
						},
						"template": {
							Type:        schema.TypeList,
							Required:    true,
							MaxItems:    1,
							Description: "Template reference configuration",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"template_ref": {
										Type:     schema.TypeString,
										Required: true,
									},
									"version_label": {
										Type:     schema.TypeString,
										Required: true,
									},
									"variables": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "List of variables passed to the template",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"name": {
													Type:     schema.TypeString,
													Required: true,
												},
												"value": {
													Type:     schema.TypeString,
													Required: true,
												},
												"type": {
													Type:     schema.TypeString,
													Required: true,
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
			"tags": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "Key-value tags",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"created": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Timestamp when the notification rule was created.",
			},
			"last_modified": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Timestamp when the notification rule was last modified.",
			},
		},
	}
}
