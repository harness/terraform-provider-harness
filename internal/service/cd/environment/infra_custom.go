package environment

import (
	"github.com/harness/harness-go-sdk/harness/cd/cac"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func infraDetailsCustom() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"deployment_type_template_version": {
				Description: "The template version",
				Type:        schema.TypeString,
				Required:    true,
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
					},
				},
			},
		},
	}
}

func expandCustomConfiguration(d []interface{}, infraDef *cac.InfrastructureDefinition) {
	if len(d) == 0 {
		return
	}

	config := d[0].(map[string]interface{})
	details := &cac.InfrastructureCustom{}

	if attr := config["deployment_type_template_version"]; attr != "" {
		details.DeploymentTypeTemplateVersion = attr.(string)
	}
	if vars := config["variable"]; vars != nil {
		details.InfraVariables = expandVariables(vars.(*schema.Set).List())
	}

	infraDef.InfrastructureDetail = details.ToInfrastructureDetail()
}

func flattenCustomConfiguration(d *schema.ResourceData, infraDef *cac.InfrastructureDefinition) []interface{} {
	results := []interface{}{}

	if len(infraDef.InfrastructureDetail) == 0 {
		return results
	}

	detail := infraDef.InfrastructureDetail[0]

	if detail.Type != cac.InfrastructureTypes.Custom {
		return results
	}

	detailConfig := map[string]interface{}{}
	infraDetail := detail.ToCustom()

	detailConfig["deployment_type_template_version"] = infraDetail.DeploymentTypeTemplateVersion
	detailConfig["variable"] = flattenVariables(infraDetail.InfraVariables)

	return append(results, detailConfig)
}

func expandVariables(d []interface{}) []*cac.InfraVariable {
	if len(d) == 0 {
		return make([]*cac.InfraVariable, 0)
	}

	variables := make([]*cac.InfraVariable, len(d))

	for i, v := range d {
		data := v.(map[string]interface{})
		variables[i] = &cac.InfraVariable{
			Name:  data["name"].(string),
			Value: data["value"].(string),
		}
	}

	return variables
}

func flattenVariables(variables []*cac.InfraVariable) []map[string]interface{} {
	if len(variables) == 0 {
		return make([]map[string]interface{}, 0)
	}

	var results = make([]map[string]interface{}, len(variables))

	for i, v := range variables {
		r := map[string]interface{}{
			"name":  v.Name,
			"value": v.Value,
		}
		results[i] = r
	}

	return results
}
