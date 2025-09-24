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

func ResourceFMESegment() *schema.Resource {
	return &schema.Resource{
		Description:   "Resource for creating a FME (Feature Management Engine) segment.",
		ReadContext:   resourceFMESegmentRead,
		CreateContext: resourceFMESegmentCreate,
		DeleteContext: resourceFMESegmentDelete,
		Importer:      helpers.MultiLevelResourceImporter,

		Schema: map[string]*schema.Schema{
			"workspace_id": {
				Description:  "Unique identifier of the workspace.",
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.IsUUID,
			},
			"traffic_type_id": {
				Description:  "Unique identifier of the traffic type.",
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.IsUUID,
			},
			"name": {
				Description:  "Name of the segment.",
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringIsNotEmpty,
			},
			"description": {
				Description: "Description of the segment.",
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
			},
			"creation_time": {
				Description: "Time when the segment was created.",
				Type:        schema.TypeInt,
				Computed:    true,
			},
			"last_update_time": {
				Description: "Time when the segment was last updated.",
				Type:        schema.TypeInt,
				Computed:    true,
			},
		},
	}
}

func resourceFMESegmentCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	session := meta.(*internal.Session)
	fmeClient := session.FMEClient

	if fmeClient == nil {
		return diag.Errorf("FME client not configured")
	}

	c := fmeClient.(*FMEConfig)
	workspaceID := d.Get("workspace_id").(string)
	trafficTypeID := d.Get("traffic_type_id").(string)
	name := d.Get("name").(string)
	description := d.Get("description").(string)

	req := &api.SegmentCreateRequest{
		Name:        name,
		Description: description,
	}

	segment, err := c.APIClient.Segments.Create(workspaceID, trafficTypeID, req)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(*segment.Name)

	return resourceFMESegmentRead(ctx, d, meta)
}

func resourceFMESegmentRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	session := meta.(*internal.Session)
	fmeClient := session.FMEClient

	if fmeClient == nil {
		return diag.Errorf("FME client not configured")
	}

	c := fmeClient.(*FMEConfig)
	workspaceID := d.Get("workspace_id").(string)
	segmentName := d.Id()

	segment, err := c.APIClient.Segments.Get(workspaceID, segmentName)
	if err != nil {
		return diag.FromErr(err)
	}

	if segment == nil {
		d.SetId("")
		return nil
	}

	d.Set("name", segment.Name)
	d.Set("description", segment.Description)
	d.Set("workspace_id", workspaceID) // Use the workspace_id from config, not from API response
	d.Set("traffic_type_id", segment.TrafficTypeID)
	d.Set("creation_time", segment.CreationTime)
	d.Set("last_update_time", segment.LastUpdateTime)

	return nil
}

func resourceFMESegmentDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	session := meta.(*internal.Session)
	fmeClient := session.FMEClient

	if fmeClient == nil {
		return diag.Errorf("FME client not configured")
	}

	c := fmeClient.(*FMEConfig)
	workspaceID := d.Get("workspace_id").(string)
	segmentName := d.Id()

	err := c.APIClient.Segments.Delete(workspaceID, segmentName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")
	return nil
}