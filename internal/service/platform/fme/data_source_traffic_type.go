package fme

import (
	"context"

	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func DataSourceFMETrafficType() *schema.Resource {
	return &schema.Resource{
		Description: "Data source for retrieving a FME (Feature Management Engine) traffic type.",
		ReadContext: dataSourceFMETrafficTypeRead,

		Schema: map[string]*schema.Schema{
			"workspace_id": {
				Description:  "Unique identifier of the workspace.",
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.IsUUID,
			},
			"name": {
				Description:  "Name of the traffic type.",
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsNotEmpty,
			},
			"id": {
				Description: "Unique identifier of the traffic type.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"type": {
				Description: "Type of the traffic type.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"display_attribute_id": {
				Description: "Display attribute ID for the traffic type.",
				Type:        schema.TypeString,
				Computed:    true,
			},
		},
	}
}

func dataSourceFMETrafficTypeRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	session := meta.(*internal.Session)
	fmeClient := session.FMEClient

	if fmeClient == nil {
		return diag.Errorf("FME client not configured")
	}

	c := fmeClient.(*FMEConfig)
	workspaceID := d.Get("workspace_id").(string)
	trafficTypeName := d.Get("name").(string)

	trafficType, err := c.APIClient.TrafficTypes.FindByName(workspaceID, trafficTypeName)
	if err != nil {
		return diag.FromErr(err)
	}

	if trafficType == nil {
		return diag.Errorf("traffic type with name %s not found in workspace %s", trafficTypeName, workspaceID)
	}

	d.SetId(*trafficType.ID)
	d.Set("name", trafficType.Name)
	d.Set("type", trafficType.Type)
	d.Set("display_attribute_id", trafficType.DisplayAttributeID)
	d.Set("workspace_id", workspaceID)

	return nil
}