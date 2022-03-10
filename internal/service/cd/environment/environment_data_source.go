package environment

import (
	"context"

	"github.com/harness/harness-go-sdk/harness/cd"
	"github.com/harness/harness-go-sdk/harness/cd/cac"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceEnvironment() *schema.Resource {
	return &schema.Resource{
		// This description is used by the documentation generator and the language server.
		Description: "Data source for retrieving a Harness service",

		ReadContext: dataSourceEnvironmentRead,

		Schema: map[string]*schema.Schema{
			"id": {
				Description: "The id of the environment.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"environment_id": {
				Description:   "The id of the environment.",
				Type:          schema.TypeString,
				Optional:      true,
				AtLeastOneOf:  []string{"environment_id", "name"},
				ConflictsWith: []string{"name"},
			},
			"app_id": {
				Description: "The id of the application.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"name": {
				Description:   "The name of the environment.",
				Type:          schema.TypeString,
				Optional:      true,
				AtLeastOneOf:  []string{"environment_id", "name"},
				ConflictsWith: []string{"environment_id"},
			},
			"description": {
				Description: "The description of the environment.",
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
	c := meta.(*cd.ApiClient)

	var env *cac.Environment
	var err error

	appId := d.Get("app_id").(string)

	if id := d.Get("environment_id").(string); id != "" {
		env, err = c.ConfigAsCodeClient.GetEnvironmentById(appId, id)
		if err != nil {
			return diag.FromErr(err)
		}
	} else if name := d.Get("name").(string); name != "" {
		env, err = c.ConfigAsCodeClient.GetEnvironmentByName(appId, name)
		if err != nil {
			return diag.FromErr(err)
		}
	}

	if env == nil {
		return diag.Errorf("environment not found")
	}

	d.SetId(env.Id)
	d.Set("environment_id", env.Id)
	d.Set("app_id", env.ApplicationId)
	d.Set("name", env.Name)
	d.Set("type", env.EnvironmentType)
	d.Set("description", env.Description)

	if overrides := flattenVariableOverrides(env.VariableOverrides); len(overrides) > 0 {
		d.Set("variable_overrides", overrides)
	}

	return nil
}
