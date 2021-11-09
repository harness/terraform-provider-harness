package environment

import (
	"context"
	"strings"

	"github.com/harness-io/harness-go-sdk/harness/api"
	"github.com/harness-io/harness-go-sdk/harness/cd/cac"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func ResourceEnvironment() *schema.Resource {
	return &schema.Resource{
		Description:   "Resource for creating an environment",
		CreateContext: resourceEnvironmentCreateOrUpdate,
		ReadContext:   resourceEnvironmentRead,
		UpdateContext: resourceEnvironmentCreateOrUpdate,
		DeleteContext: resourceEnvironmentDelete,
		Schema: map[string]*schema.Schema{
			"id": {
				Description: "The id of the environment.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"app_id": {
				Description: "The id of the application.",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
			},
			"name": {
				Description: "The name of the environment.",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
			},
			"description": {
				Description: "The description of the environment.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"type": {
				Description:  "The type of the environment. Valid values are `PROD` and `NON_PROD`",
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringInSlice([]string{"PROD", "NON_PROD"}, false),
			},
			"variable_override": {
				Description: "Override for a service variable",
				Type:        schema.TypeSet,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Description: "The name of the variable",
							Type:        schema.TypeString,
							Required:    true,
						},
						"service_name": {
							Description: "The name of the service",
							Type:        schema.TypeString,
							Required:    true,
						},
						"value": {
							Description: "The value of the service variable",
							Type:        schema.TypeString,
							Required:    true,
						},
						"type": {
							Description:  "The type of the service variable. Valid values are `TEXT` and `ENCRYPTED_TEXT`",
							Type:         schema.TypeString,
							Required:     true,
							ValidateFunc: validation.StringInSlice([]string{"TEXT", "ENCRYPTED_TEXT"}, false),
						},
					},
				},
			},
		},

		Importer: &schema.ResourceImporter{
			State: func(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
				// <app_id>/<env_id>
				parts := strings.Split(d.Id(), "/")
				d.Set("app_id", parts[0])
				d.SetId(parts[1])

				return []*schema.ResourceData{d}, nil
			},
		},
	}
}

func resourceEnvironmentRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*api.Client)

	var env *cac.Environment
	var err error

	envId := d.Get("id").(string)
	appId := d.Get("app_id").(string)

	if env, err = c.CDClient.ConfigAsCodeClient.GetEnvironmentById(appId, envId); err != nil {
		return diag.FromErr(err)
	} else if env == nil {
		d.SetId("")
		d.MarkNewResource()
		return nil
	}

	return readEnvironment(d, env)
}

func readEnvironment(d *schema.ResourceData, env *cac.Environment) diag.Diagnostics {
	d.SetId(env.Id)
	d.Set("app_id", env.ApplicationId)
	d.Set("name", env.Name)
	d.Set("app_id", env.ApplicationId)
	d.Set("type", env.EnvironmentType)
	d.Set("description", env.Description)

	if overrides := flattenVariableOverrides(env.VariableOverrides); len(overrides) > 0 {
		d.Set("variable_override", overrides)
	}

	return nil
}

func resourceEnvironmentCreateOrUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*api.Client)

	appId := d.Get("app_id").(string)
	id := d.Get("id").(string)

	var env *cac.Environment
	var err error

	if d.IsNewResource() {
		env = cac.NewEntity(cac.ObjectTypes.Environment).(*cac.Environment)
	} else {
		if env, err = c.CDClient.ConfigAsCodeClient.GetEnvironmentById(appId, id); err != nil {
			return diag.FromErr(err)
		} else if env == nil {
			d.SetId("")
			d.MarkNewResource()
			return nil
		}
	}

	env.Id = id
	env.ApplicationId = appId
	env.Name = d.Get("name").(string)
	env.EnvironmentType = cac.EnvironmentType(d.Get("type").(string))
	env.ApplicationId = d.Get("app_id").(string)
	env.Description = d.Get("description").(string)

	if overrides := d.Get("variable_override"); overrides != nil {
		env.VariableOverrides = expandVariableOverrides(overrides.(*schema.Set).List())
	}

	newEnv, err := c.CDClient.ConfigAsCodeClient.UpsertEnvironment(env)
	if err != nil {
		return diag.FromErr(err)
	}

	return readEnvironment(d, newEnv)
}

func flattenVariableOverrides(overrides []*cac.VariableOverride) []map[string]interface{} {
	if len(overrides) == 0 {
		return make([]map[string]interface{}, 0)
	}

	var results = make([]map[string]interface{}, len(overrides))

	for i, override := range overrides {
		results[i] = map[string]interface{}{
			"name":         override.Name,
			"service_name": override.ServiceName,
			"value":        override.Value,
			"type":         override.ValueType,
		}
	}

	return results
}

func expandVariableOverrides(d []interface{}) []*cac.VariableOverride {

	if len(d) == 0 {
		return make([]*cac.VariableOverride, 0)
	}

	overrides := make([]*cac.VariableOverride, len(d))

	for i, override := range d {
		data := override.(map[string]interface{})
		overrides[i] = &cac.VariableOverride{
			Name:        data["name"].(string),
			ServiceName: data["service_name"].(string),
			Value:       data["value"].(string),
			ValueType:   cac.VariableValueType(data["type"].(string)),
		}
	}

	return overrides
}

func resourceEnvironmentDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*api.Client)

	envName := d.Get("name").(string)
	appId := d.Get("app_id").(string)

	app, err := c.CDClient.ApplicationClient.GetApplicationById(appId)
	if err != nil {
		return diag.FromErr(err)
	}

	err = c.CDClient.ConfigAsCodeClient.DeleteEnvironment(app.Name, envName)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}
