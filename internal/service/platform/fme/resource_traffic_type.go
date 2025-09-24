package fme

import (
	"context"

	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/harness/terraform-provider-harness/internal/service/platform/fme/api"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func ResourceFMETrafficType() *schema.Resource {
	return &schema.Resource{
		Description:   "Resource for creating a FME (Feature Management Engine) traffic type.",
		ReadContext:   resourceFMETrafficTypeRead,
		CreateContext: resourceFMETrafficTypeCreate,
		DeleteContext: resourceFMETrafficTypeDelete,
		Importer:      helpers.MultiLevelResourceImporter,

		Schema: map[string]*schema.Schema{
			"workspace_id": {
				Description:  "Unique identifier of the workspace.",
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.IsUUID,
			},
			"name": {
				Description:  "Name of the traffic type.",
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringIsNotEmpty,
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

func resourceFMETrafficTypeCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	session := meta.(*internal.Session)
	fmeClient := session.FMEClient

	if fmeClient == nil {
		return diag.Errorf("FME client not configured")
	}

	c := fmeClient.(*FMEConfig)
	workspaceID := d.Get("workspace_id").(string)
	name := d.Get("name").(string)

	req := &api.TrafficTypeCreateRequest{
		Name: name,
	}

	trafficType, err := c.APIClient.TrafficTypes.Create(workspaceID, req)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(*trafficType.ID)

	return resourceFMETrafficTypeRead(ctx, d, meta)
}

func resourceFMETrafficTypeRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	session := meta.(*internal.Session)
	fmeClient := session.FMEClient

	if fmeClient == nil {
		return diag.Errorf("FME client not configured")
	}

	c := fmeClient.(*FMEConfig)
	workspaceID := d.Get("workspace_id").(string)
	trafficTypeID := d.Id()

	trafficType, err := c.APIClient.TrafficTypes.FindByID(workspaceID, trafficTypeID)
	if err != nil {
		return diag.FromErr(err)
	}

	if trafficType == nil {
		d.SetId("")
		return nil
	}

	d.Set("name", trafficType.Name)
	d.Set("type", trafficType.Type)
	d.Set("display_attribute_id", trafficType.DisplayAttributeID)
	d.Set("workspace_id", workspaceID)

	return nil
}

func resourceFMETrafficTypeDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	session := meta.(*internal.Session)
	fmeClient := session.FMEClient

	if fmeClient == nil {
		return diag.Errorf("FME client not configured")
	}

	c := fmeClient.(*FMEConfig)
	workspaceID := d.Get("workspace_id").(string)
	trafficTypeID := d.Id()

	err := c.APIClient.TrafficTypes.Delete(workspaceID, trafficTypeID)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")
	return nil
}