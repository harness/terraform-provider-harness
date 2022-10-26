package variables

import (
	"context"
	"net/http"

	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func ResourceVariables() *schema.Resource {
	resource := &schema.Resource{
		Description: "Resource for creating a Harness Variables.",

		ReadContext:   resourceVariablesRead,
		UpdateContext: resourceVariablesCreateOrUpdate,
		DeleteContext: resourceVariablesDelete,
		CreateContext: resourceVariablesCreateOrUpdate,
		Importer:      helpers.MultiLevelResourceImporter,

		Schema: map[string]*schema.Schema{
			"identifier": {
				Description: "Unique identifier of the resource",
				Type:        schema.TypeString,
				Required:    true,
			},
			"name": {
				Description: "Unique identifier of the resource",
				Type:        schema.TypeString,
				Required:    true,
			},
			"description": {
				Description: "Unique identifier of the resource",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"org_id": {
				Description: "Unique identifier of the resource",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"project_id": {
				Description: "Unique identifier of the resource",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"type": {
				Description: "Type of Variable",
				Type:        schema.TypeString,
				Required:    true,
			},
			"spec": {
				Description: "List of Spec Fields.",
				Type:        schema.TypeList,
				MaxItems:    1,
				Required:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"value_type": {
							Description:  "Type of Value of the Variable. For now only FIXED is supported",
							Type:         schema.TypeString,
							Required:     true,
							ValidateFunc: validation.StringInSlice([]string{"FIXED"}, false),
						},
						"fixed_value": {
							Description: "FixedValue of the variable",
							Type:        schema.TypeString,
							Required:    true,
						},
					},
				},
			},
		},
	}

	return resource
}

func resourceVariablesRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	id := d.Id()
	resp, httpResp, err := c.VariablesApi.GetVariable(ctx, id, c.AccountId, &nextgen.VariablesApiGetVariableOpts{
		OrgIdentifier:     helpers.BuildField(d, "org_id"),
		ProjectIdentifier: helpers.BuildField(d, "project_id"),
	})

	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	if resp.Data == nil {
		d.SetId("")
		d.MarkNewResource()
		return nil
	}

	readVariable(d, resp.Data.Variable)

	return nil
}

func resourceVariablesCreateOrUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	var err error
	var resp nextgen.ResponseDtoVariableResponseDto
	var httpResp *http.Response

	id := d.Id()
	variable := buildVariables(d)

	if id == "" {
		resp, httpResp, err = c.VariablesApi.CreateVariable(ctx, nextgen.VariableRequestDto{Variable: variable}, c.AccountId)
	} else {
		resp, httpResp, err = c.VariablesApi.UpdateVariable(ctx, nextgen.VariableRequestDto{Variable: variable}, c.AccountId)
	}

	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	readVariable(d, resp.Data.Variable)

	return nil
}

func resourceVariablesDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	_, httpResp, err := c.VariablesApi.DeleteVariable(ctx, c.AccountId, d.Id(), &nextgen.VariablesApiDeleteVariableOpts{
		OrgIdentifier:     helpers.BuildField(d, "org_id"),
		ProjectIdentifier: helpers.BuildField(d, "project_id"),
	})

	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	return nil
}

func buildVariables(d *schema.ResourceData) *nextgen.VariableDto {
	variable := &nextgen.VariableDto{
		Spec: &nextgen.StringVariableConfigDto{},
	}

	if attr, ok := d.GetOk("org_id"); ok {
		variable.OrgIdentifier = attr.(string)
	}

	if attr, ok := d.GetOk("project_id"); ok {
		variable.ProjectIdentifier = attr.(string)
	}

	if attr, ok := d.GetOk("description"); ok {
		variable.Description = attr.(string)
	}

	if attr, ok := d.GetOk("name"); ok {
		variable.Name = attr.(string)
	}

	if attr, ok := d.GetOk("identifier"); ok {
		variable.Identifier = attr.(string)
	}

	if attr, ok := d.GetOk("spec"); ok {
		config := attr.([]interface{})[0].(map[string]interface{})
		if attr, ok := config["value_type"]; ok {
			variable.Spec.ValueType = attr.(string)
		}

		if attr, ok := config["fixed_value"]; ok {
			variable.Spec.FixedValue = attr.(string)
		}
	}

	if attr, ok := d.GetOk("type"); ok {
		variable.Type_ = attr.(string)
	}
	return variable
}

func readVariable(d *schema.ResourceData, variable *nextgen.VariableDto) {
	d.SetId(variable.Identifier)
	d.Set("identifier", variable.Identifier)
	d.Set("org_id", variable.OrgIdentifier)
	d.Set("project_id", variable.ProjectIdentifier)
	d.Set("name", variable.Name)
	d.Set("description", variable.Description)
	d.Set("type", variable.Type_)
	d.Set("spec", []interface{}{
		map[string]interface{}{
			"value_type":  variable.Spec.ValueType,
			"fixed_value": variable.Spec.FixedValue,
		},
	})
}
