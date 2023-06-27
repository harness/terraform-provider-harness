package as_rule

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceVMRule() *schema.Resource {
	resource := &schema.Resource{
		Description: "Data source for retrieving a Harness Variable.",

		ReadContext: resourceVMRuleRead,

		Schema: map[string]*schema.Schema{
			"identifier": {
				Description: "Unique identifier of the resource",
				Type:        schema.TypeFloat,
				Computed:    true,
			},
			"name": {
				Description: "Name of the Variable",
				Type:        schema.TypeString,
				Required:    true,
			},
			"cloud_connector_id": {
				Description: "Description of the entity",
				Type:        schema.TypeString,
				Required:    true,
			},
			"idle_time_mins": {
				Description: "Organization Identifier for the Entity",
				Type:        schema.TypeInt,
				Required:    true,
			},
			"use_spot": {
				Description: "Project Identifier for the Entity",
				Type:        schema.TypeBool,
				Default:     false,
				Optional:    true,
			},
			"custom_domains": {
				Description: "Type of Variable",
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"filter": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"vm_ids": {
							Type:     schema.TypeList,
							Required: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"tags": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"key": {
										Type:     schema.TypeString,
										Required: true,
									},
									"value": {
										Type:     schema.TypeString,
										Required: true,
									},
								},
							},
						},
						"regions": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"zones": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
			"http": {
				Description: "List of Spce Fields.",
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"proxy_id": {
							Description: "Type of Value of the Variable. For now only FIXED is supported",
							Type:        schema.TypeString,
							Required:    true,
						},
						"routing": {
							Description: "FixedValue of the variable",
							Type:        schema.TypeList,
							MinItems:    1,
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"source_protocol": {
										Description: "Type of Value of the Variable. For now only FIXED is supported",
										Type:        schema.TypeString,
										Required:    true,
									},
									"target_protocol": {
										Description: "Type of Value of the Variable. For now only FIXED is supported",
										Type:        schema.TypeString,
										Required:    true,
									},
									"source_port": {
										Description: "Organization Identifier for the Entity",
										Type:        schema.TypeInt,
										Optional:    true,
									},
									"target_port": {
										Description: "Organization Identifier for the Entity",
										Type:        schema.TypeInt,
										Optional:    true,
									},
									"action": {
										Description: "Organization Identifier for the Entity",
										Type:        schema.TypeString,
										Optional:    true,
									},
								},
							},
						},
						"health": {
							Description: "FixedValue of the variable",
							Type:        schema.TypeList,
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"protocol": {
										Description: "Type of Value of the Variable. For now only FIXED is supported",
										Type:        schema.TypeString,
										Required:    true,
									},
									"port": {
										Description: "Type of Value of the Variable. For now only FIXED is supported",
										Type:        schema.TypeInt,
										Required:    true,
									},
									"path": {
										Description: "Organization Identifier for the Entity",
										Type:        schema.TypeString,
										Optional:    true,
									},
									"timeout": {
										Description: "Organization Identifier for the Entity",
										Type:        schema.TypeInt,
										Optional:    true,
									},
									"status_code_from": {
										Description: "Organization Identifier for the Entity",
										Type:        schema.TypeInt,
										Optional:    true,
									},
									"status_code_to": {
										Description: "Organization Identifier for the Entity",
										Type:        schema.TypeInt,
										Optional:    true,
									},
								},
							},
						},
					},
				},
			},
			"tcp": {
				Description: "FixedValue of the variable",
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"proxy_id": {
							Description: "Type of Value of the Variable. For now only FIXED is supported",
							Type:        schema.TypeString,
							Required:    true,
						},
						"ssh": {
							Description: "FixedValue of the variable",
							Type:        schema.TypeList,
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"connect_on": {
										Description: "Type of Value of the Variable. For now only FIXED is supported",
										Type:        schema.TypeInt,
										Optional:    true,
									},
									"port": {
										Description: "Type of Value of the Variable. For now only FIXED is supported",
										Type:        schema.TypeInt,
										Optional:    true,
										Default:     22,
									},
								},
							},
						},
						"rdp": {
							Description: "FixedValue of the variable",
							Type:        schema.TypeList,
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"connect_on": {
										Description: "Type of Value of the Variable. For now only FIXED is supported",
										Type:        schema.TypeInt,
										Optional:    true,
									},
									"port": {
										Description: "Type of Value of the Variable. For now only FIXED is supported",
										Type:        schema.TypeInt,
										Optional:    true,
										Default:     3389,
									},
								},
							},
						},
						"forward_rule": {
							Description: "FixedValue of the variable",
							Type:        schema.TypeList,
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"connect_on": {
										Description: "Type of Value of the Variable. For now only FIXED is supported",
										Type:        schema.TypeInt,
										Optional:    true,
									},
									"port": {
										Description: "Type of Value of the Variable. For now only FIXED is supported",
										Type:        schema.TypeInt,
										Required:    true,
									},
								},
							},
						},
					},
				},
			},
			"depends": {
				Description: "FixedValue of the variable",
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"rule_id": {
							Description: "Type of Value of the Variable. For now only FIXED is supported",
							Type:        schema.TypeInt,
							Required:    true,
						},
						"delay_in_sec": {
							Description: "Organization Identifier for the Entity",
							Type:        schema.TypeInt,
							Optional:    true,
							Default:     5,
						},
					},
				},
			},
		},
	}

	return resource
}
