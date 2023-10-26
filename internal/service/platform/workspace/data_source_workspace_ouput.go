package workspace

import (
	"context"

	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceWorkspaceOutput() *schema.Resource {
	resource := &schema.Resource{
		Description: "Data source for retrieving workspace outputs.",

		ReadContext: resourceWorkspaceOutputRead,
		Importer:    helpers.ProjectResourceImporter,

		Schema: map[string]*schema.Schema{
			"identifier": {
				Description: "Identifier of the Workspace.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"org_id": {
				Description: "Organization identifier of the organization the workspace resides in.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"project_id": {
				Description: "Project identifier of the project the workspace resides in.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"outputs": {
				Description: "",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Description: "Name associated with the output.",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"value": {
							Description: "Value of the output.",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"sensitive": {
							Description: "Indicates if the output is sensitive.",
							Type:        schema.TypeBool,
							Computed:    true,
						},
					},
				},
			},
		},
	}
	return resource
}

func resourceWorkspaceOutputRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)
	id := d.Id()
	if id == "" {
		d.MarkNewResource()
	}
	resp, httpResp, err := c.WorkspaceApi.WorkspacesListResources(
		ctx,
		d.Get("org_id").(string),
		d.Get("project_id").(string),
		d.Get("identifier").(string),
		c.AccountId,
	)
	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}
	readOutputs(d, &resp)
	return nil
}

func readOutputs(d *schema.ResourceData, outputsResp *nextgen.IacmResourcw) {
	d.SetId(d.Get("identifier").(string))
	d.Set("org_id", d.Get("org_id").(string))
	d.Set("project_id", d.Get("project_id").(string))
	var outputs []interface{}
	for _, o := range *outputsResp.Outputs {
		outputs = append(outputs, map[string]interface{}{
			"name":      o.Name,
			"value":     o.Value,
			"sensitive": o.Sensitive,
		})
	}
	d.Set("outputs", outputs)
}
