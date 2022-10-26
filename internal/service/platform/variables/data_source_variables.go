package variables

import (
	"context"
	"errors"
	"net/http"

	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func DataSourceVariables() *schema.Resource {
	resource := &schema.Resource{
		Description: "Data source for retrieving a Harness Variable.",

		ReadContext: dataSourceVariablesRead,

		Schema: map[string]*schema.Schema{
			"identifier": {
				Description: "Unique identifier of the resource",
				Type:        schema.TypeString,
				Required:    true,
			},
			"name": {
				Description: "Unique identifier of the resource",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"description": {
				Description: "Unique identifier of the resource",
				Type:        schema.TypeString,
				Computed:    true,
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
				Computed:    true,
			},
			"spec": {
				Description: "List of Spce Fields.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"value_type": {
							Description:  "Type of Value of the Variable. For now only FIXED is supported",
							Type:         schema.TypeString,
							Computed:     true,
							ValidateFunc: validation.StringInSlice([]string{"FIXED"}, false),
						},
						"fixed_value": {
							Description: "FixedValue of the variable",
							Type:        schema.TypeString,
							Computed:    true,
						},
					},
				},
			},
		},
	}

	return resource
}

func dataSourceVariablesRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	var variable *nextgen.VariableDto
	var err error
	var httpResp *http.Response

	id := d.Get("identifier").(string)

	if id != "" {
		var resp nextgen.ResponseDtoVariableResponseDto
		resp, httpResp, err = c.VariablesApi.GetVariable(ctx, id, c.AccountId, &nextgen.VariablesApiGetVariableOpts{
			OrgIdentifier:     helpers.BuildField(d, "org_id"),
			ProjectIdentifier: helpers.BuildField(d, "project_id"),
		})
		variable = resp.Data.Variable
	} else {
		return diag.FromErr(errors.New(" identifier  must be specified"))
	}

	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	if variable == nil {
		d.SetId("")
		d.MarkNewResource()
		return nil
	}

	readVariable(d, variable)

	return nil
}
