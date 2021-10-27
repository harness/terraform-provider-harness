package cd

import (
	"context"
	"strings"

	"github.com/harness-io/harness-go-sdk/harness/api"
	"github.com/harness-io/harness-go-sdk/harness/api/cac"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func serviceStateImporter(d *schema.ResourceData, i interface{}) ([]*schema.ResourceData, error) {
	// <app_id>/<svc_id>
	parts := strings.Split(d.Id(), "/")

	d.Set("app_id", parts[0])
	d.SetId(parts[1])

	return []*schema.ResourceData{d}, nil
}

func resourceServiceDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*api.Client)

	id := d.Get("id").(string)
	appId := d.Get("app_id").(string)

	err := c.ConfigAsCode().DeleteService(appId, id)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func commonServiceSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": {
			Description: "Id of the service",
			Type:        schema.TypeString,
			Computed:    true,
		},
		"app_id": {
			Description: "The id of the application the service belongs to",
			Type:        schema.TypeString,
			Required:    true,
			ForceNew:    true,
		},
		"name": {
			Description: "Name of the service",
			Type:        schema.TypeString,
			Required:    true,
		},
		"description": {
			Description: "Description of th service",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"variable": {
			Description: "Variables to be used in the service",
			Type:        schema.TypeSet,
			Optional:    true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"name": {
						Description: "Name of the variable",
						Type:        schema.TypeString,
						Required:    true,
					},
					"value": {
						Description: "Value of the variable",
						Type:        schema.TypeString,
						Required:    true,
					},
					"type": {
						Description: "Type of the variable. Options are 'TEXT' and 'ENCRYPTED_TEXT'",
						Type:        schema.TypeString,
						Required:    true,
					},
				},
			},
		},
	}
}

func flattenServiceVariables(variables []*cac.ServiceVariable) []map[string]interface{} {
	if len(variables) == 0 {
		return make([]map[string]interface{}, 0)
	}

	var results = make([]map[string]interface{}, len(variables))

	for i, v := range variables {
		r := map[string]interface{}{
			"name":  v.Name,
			"value": v.Value,
			"type":  v.ValueType,
		}
		results[i] = r
	}

	return results
}

func expandServiceVariables(d []interface{}) []*cac.ServiceVariable {
	if len(d) == 0 {
		return make([]*cac.ServiceVariable, 0)
	}

	variables := make([]*cac.ServiceVariable, len(d))

	for i, v := range d {
		data := v.(map[string]interface{})
		variables[i] = &cac.ServiceVariable{
			Name:      data["name"].(string),
			Value:     data["value"].(string),
			ValueType: cac.VariableValueType(data["type"].(string)),
		}
	}

	return variables
}
