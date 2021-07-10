package provider

import (
	"context"

	"github.com/harness-io/harness-go-sdk/harness/api"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceEnvironment() *schema.Resource {
	return &schema.Resource{
		// This description is used by the documentation generator and the language server.
		Description: "Data source for retrieving a Harness service",

		ReadContext: dataSourceEnvironmentRead,

		Schema: map[string]*schema.Schema{
			"id": {
				Description: "The id of the environment.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"app_id": {
				Description: "The id of the application.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"name": {
				Description: "The name of the environment.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"type": {
				Description: "The type of the environment. Valid values are `PROD` and `NON_PROD`",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"variable_override": {
				Description: "Override for a service variable",
				Type:        schema.TypeSet,
				Computed:    true,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Description: "The name of the variable",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"service_name": {
							Description: "The name of the service",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"value": {
							Description: "The value of the service variable",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"type": {
							Description: "The type of the service variable. Valid values are `TEXT` and `ENCRYPTED_TEXT`",
							Type:        schema.TypeString,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourceEnvironmentRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*api.Client)

	envId := d.Get("id").(string)
	appId := d.Get("app_id").(string)

	env, err := c.ConfigAsCode().GetEnvironmentById(appId, envId)
	if err != nil {
		return diag.FromErr(err)
	}

	d.Set("name", env.Name)
	d.Set("app_id", env.ApplicationId)
	d.Set("type", env.EnvironmentType)

	if overrides := flattenVariableOverrides(env.VariableOverrides); len(overrides) > 0 {
		d.Set("variable_overrides", overrides)
	}

	d.SetId(env.Id)

	return nil
}
