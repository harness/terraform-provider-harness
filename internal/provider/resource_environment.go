package provider

import (
	"context"
	"errors"

	"github.com/harness-io/harness-go-sdk/harness/api"
	"github.com/harness-io/harness-go-sdk/harness/api/cac"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceEnvironment() *schema.Resource {
	return &schema.Resource{
		Description:   "Resource for creating an environment",
		CreateContext: resourceEnvironmentCreate,
		ReadContext:   resourceEnvironmentRead,
		UpdateContext: resourceEnvironmentUpdate,
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
			},
			"name": {
				Description: "The name of the environment.",
				Type:        schema.TypeString,
				Required:    true,
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
	}
}

func resourceEnvironmentRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	return nil
}

func resourceEnvironmentCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*api.Client)

	env := cac.NewEntity(cac.ObjectTypes.Environment).(*cac.Environment)
	env.Name = d.Get("name").(string)
	env.EnvironmentType = cac.EnvironmentType(d.Get("type").(string))
	env.ApplicationId = d.Get("app_id").(string)

	if overrides := d.Get("variable_overrides"); overrides != nil {
		env.VariableOverrides = expandVariableOverrides(overrides.(*schema.Set).List())
	}

	newEnv, err := c.ConfigAsCode().UpsertEnvironment(env)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(newEnv.Id)

	return nil
}

func resourceEnvironmentUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*api.Client)

	if d.HasChange("app_id") {
		return diag.FromErr(errors.New("app_id cannot be changed"))
	}

	if d.HasChange("name") {
		return diag.FromErr(errors.New("name cannot be changed"))
	}

	envInput := cac.NewEntity(cac.ObjectTypes.Environment).(*cac.Environment)
	envInput.Name = d.Get("name").(string)
	envInput.EnvironmentType = cac.EnvironmentType(d.Get("type").(string))
	envInput.ApplicationId = d.Get("app_id").(string)

	if overrides := d.Get("variable_overrides"); overrides != nil {
		envInput.VariableOverrides = expandVariableOverrides(overrides.(*schema.Set).List())
	}

	// Update the environment
	newEnv, err := c.ConfigAsCode().UpsertEnvironment(envInput)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(newEnv.Id)

	return nil
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

	app, err := c.Applications().GetApplicationById(appId)
	if err != nil {
		return diag.FromErr(err)
	}

	err = c.ConfigAsCode().DeleteEnvironment(app.Name, envName)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}
