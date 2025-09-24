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

func ResourceFMESplit() *schema.Resource {
	return &schema.Resource{
		Description:   "Resource for creating a FME split (feature flag).",
		ReadContext:   resourceFMESplitRead,
		CreateContext: resourceFMESplitCreate,
		UpdateContext: resourceFMESplitUpdate,
		DeleteContext: resourceFMESplitDelete,
		Importer:      helpers.MultiLevelResourceImporter,

		Schema: map[string]*schema.Schema{
			"id": {
				Description: "Unique identifier of the split.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"workspace_id": {
				Description:  "ID of the workspace this split belongs to.",
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringIsNotEmpty,
			},
			"name": {
				Description:  "Name of the split.",
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringIsNotEmpty,
			},
			"description": {
				Description: "Description of the split.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"traffic_type_id": {
				Description:  "ID of the traffic type this split belongs to.",
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringIsNotEmpty,
			},
			"creation_time": {
				Description: "Creation time of the split.",
				Type:        schema.TypeInt,
				Computed:    true,
			},
			"rollout_status_timestamp": {
				Description: "Rollout status timestamp.",
				Type:        schema.TypeInt,
				Computed:    true,
			},
		},
	}
}

func resourceFMESplitCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	session := meta.(*internal.Session)
	fmeClient := session.FMEClient

	if fmeClient == nil {
		return diag.Errorf("FME client not configured")
	}

	c := fmeClient.(*FMEConfig)

	workspaceID := d.Get("workspace_id").(string)
	trafficTypeID := d.Get("traffic_type_id").(string)
	req := &api.SplitCreateRequest{
		Name:        d.Get("name").(string),
		Description: d.Get("description").(string),
	}

	split, err := c.APIClient.Splits.Create(workspaceID, trafficTypeID, req)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(*split.Name)
	return resourceFMESplitRead(ctx, d, meta)
}

func resourceFMESplitRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	session := meta.(*internal.Session)
	fmeClient := session.FMEClient

	if fmeClient == nil {
		return diag.Errorf("FME client not configured")
	}

	c := fmeClient.(*FMEConfig)
	workspaceID := d.Get("workspace_id").(string)
	split, err := c.APIClient.Splits.Get(workspaceID, d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	if split == nil {
		d.SetId("")
		return nil
	}

	d.Set("name", split.Name)
	d.Set("description", split.Description)
	d.Set("creation_time", split.CreationTime)
	d.Set("rollout_status_timestamp", split.RolloutStatusTimestamp)

	return nil
}

func resourceFMESplitUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	session := meta.(*internal.Session)
	fmeClient := session.FMEClient

	if fmeClient == nil {
		return diag.Errorf("FME client not configured")
	}

	c := fmeClient.(*FMEConfig)

	if d.HasChange("description") {
		workspaceID := d.Get("workspace_id").(string)
		req := &api.SplitUpdateRequest{
			Description: d.Get("description").(string),
		}

		_, err := c.APIClient.Splits.Update(workspaceID, d.Id(), req)
		if err != nil {
			return diag.FromErr(err)
		}
	}

	return resourceFMESplitRead(ctx, d, meta)
}

func resourceFMESplitDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	session := meta.(*internal.Session)
	fmeClient := session.FMEClient

	if fmeClient == nil {
		return diag.Errorf("FME client not configured")
	}

	c := fmeClient.(*FMEConfig)
	workspaceID := d.Get("workspace_id").(string)
	err := c.APIClient.Splits.Delete(workspaceID, d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}