package fme

import (
	"context"

	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func DataSourceFMEFlagSet() *schema.Resource {
	return &schema.Resource{
		Description: "Data source for retrieving a FME (Feature Management Engine) flag set.",
		ReadContext: dataSourceFMEFlagSetRead,

		Schema: map[string]*schema.Schema{
			"workspace_id": {
				Description:  "Unique identifier of the workspace.",
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.IsUUID,
			},
			"name": {
				Description:  "Name of the flag set.",
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsNotEmpty,
			},
			"id": {
				Description: "Unique identifier of the flag set.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"description": {
				Description: "Description of the flag set.",
				Type:        schema.TypeString,
				Computed:    true,
			},
		},
	}
}

func dataSourceFMEFlagSetRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	session := meta.(*internal.Session)
	fmeClient := session.FMEClient

	if fmeClient == nil {
		return diag.Errorf("FME client not configured")
	}

	c := fmeClient.(*FMEConfig)
	workspaceID := d.Get("workspace_id").(string)
	flagSetName := d.Get("name").(string)

	flagSet, err := c.APIClient.FlagSets.FindByName(workspaceID, flagSetName)
	if err != nil {
		return diag.FromErr(err)
	}

	if flagSet == nil {
		return diag.Errorf("flag set with name %s not found in workspace %s", flagSetName, workspaceID)
	}

	d.SetId(*flagSet.ID)
	d.Set("name", flagSet.Name)
	d.Set("description", flagSet.Description)
	d.Set("workspace_id", workspaceID)

	return nil
}