package fme

import (
	"context"

	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func DataSourceFMEWorkspace() *schema.Resource {
	return &schema.Resource{
		Description: "Data source for retrieving a FME workspace.",
		ReadContext: dataSourceFMEWorkspaceRead,

		Schema: map[string]*schema.Schema{
			"id": {
				Description: "Unique identifier of the workspace.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"name": {
				Description:  "Name of the workspace.",
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsNotEmpty,
			},
			"type": {
				Description: "Type of the workspace.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"display_name": {
				Description: "Display name of the workspace.",
				Type:        schema.TypeString,
				Computed:    true,
			},
		},
	}
}

func dataSourceFMEWorkspaceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	session := meta.(*internal.Session)
	fmeClient := session.FMEClient

	if fmeClient == nil {
		return diag.Errorf("FME client not configured")
	}

	c := fmeClient.(*FMEConfig)
	workspaceName := d.Get("name").(string)
	workspace, err := c.APIClient.Workspaces.FindByName(workspaceName)
	if err != nil {
		return diag.FromErr(err)
	}

	if workspace == nil {
		return diag.Errorf("workspace with name %s not found", workspaceName)
	}

	d.SetId(*workspace.ID)
	d.Set("name", workspace.Name)
	d.Set("type", workspace.Type)
	d.Set("display_name", workspace.DisplayName)

	return nil
}