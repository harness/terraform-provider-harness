package fme

import (
	"context"
	"fmt"
	"strings"

	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/harness/terraform-provider-harness/internal/service/platform/fme/api"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func ResourceFMESegmentEnvironmentAssociation() *schema.Resource {
	return &schema.Resource{
		Description:   "Resource for creating a FME segment environment association.",
		ReadContext:   resourceFMESegmentEnvironmentAssociationRead,
		CreateContext: resourceFMESegmentEnvironmentAssociationCreate,
		UpdateContext: resourceFMESegmentEnvironmentAssociationUpdate,
		DeleteContext: resourceFMESegmentEnvironmentAssociationDelete,
		Importer:      helpers.MultiLevelResourceImporter,

		Schema: map[string]*schema.Schema{
			"workspace_id": {
				Description:  "ID of the workspace this association belongs to.",
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringIsNotEmpty,
			},
			"segment_name": {
				Description:  "Name of the segment.",
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringIsNotEmpty,
			},
			"environment_id": {
				Description:  "ID of the environment.",
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringIsNotEmpty,
			},
			"include_in_segment": {
				Description: "Whether to include this environment in the segment.",
				Type:        schema.TypeBool,
				Required:    true,
			},
			"environment_name": {
				Description: "Name of the environment.",
				Type:        schema.TypeString,
				Computed:    true,
			},
		},
	}
}

func resourceFMESegmentEnvironmentAssociationCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	session := meta.(*internal.Session)
	fmeClient := session.FMEClient

	if fmeClient == nil {
		return diag.Errorf("FME client not configured")
	}

	c := fmeClient.(*FMEConfig)

	workspaceID := d.Get("workspace_id").(string)
	segmentName := d.Get("segment_name").(string)
	environmentID := d.Get("environment_id").(string)

	req := &api.SegmentEnvironmentAssociationCreateRequest{
		IncludeInSegment: d.Get("include_in_segment").(bool),
	}

	_, err := c.APIClient.SegmentEnvironmentAssociations.Create(workspaceID, segmentName, environmentID, req)
	if err != nil {
		return diag.FromErr(err)
	}

	// Use a composite ID to uniquely identify this resource
	d.SetId(fmt.Sprintf("%s:%s:%s", workspaceID, segmentName, environmentID))
	return resourceFMESegmentEnvironmentAssociationRead(ctx, d, meta)
}

func resourceFMESegmentEnvironmentAssociationRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	session := meta.(*internal.Session)
	fmeClient := session.FMEClient

	if fmeClient == nil {
		return diag.Errorf("FME client not configured")
	}

	c := fmeClient.(*FMEConfig)
	workspaceID := d.Get("workspace_id").(string)
	segmentName := d.Get("segment_name").(string)
	environmentID := d.Get("environment_id").(string)

	association, err := c.APIClient.SegmentEnvironmentAssociations.Get(workspaceID, segmentName, environmentID)
	if err != nil {
		// Handle 404 errors gracefully during resource replacement scenarios
		if strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "not found") {
			d.SetId("")
			return nil
		}
		return diag.FromErr(err)
	}

	if association == nil {
		d.SetId("")
		return nil
	}

	d.Set("segment_name", association.SegmentName)
	d.Set("environment_id", association.EnvironmentID)
	d.Set("environment_name", association.EnvironmentName)
	d.Set("include_in_segment", association.IncludeInSegment)

	return nil
}

func resourceFMESegmentEnvironmentAssociationUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	session := meta.(*internal.Session)
	fmeClient := session.FMEClient

	if fmeClient == nil {
		return diag.Errorf("FME client not configured")
	}

	c := fmeClient.(*FMEConfig)

	workspaceID := d.Get("workspace_id").(string)
	segmentName := d.Get("segment_name").(string)
	environmentID := d.Get("environment_id").(string)

	if d.HasChange("include_in_segment") {
		req := &api.SegmentEnvironmentAssociationUpdateRequest{
			IncludeInSegment: d.Get("include_in_segment").(bool),
		}

		_, err := c.APIClient.SegmentEnvironmentAssociations.Update(workspaceID, segmentName, environmentID, req)
		if err != nil {
			return diag.FromErr(err)
		}
	}

	return resourceFMESegmentEnvironmentAssociationRead(ctx, d, meta)
}

func resourceFMESegmentEnvironmentAssociationDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	session := meta.(*internal.Session)
	fmeClient := session.FMEClient

	if fmeClient == nil {
		return diag.Errorf("FME client not configured")
	}

	c := fmeClient.(*FMEConfig)
	workspaceID := d.Get("workspace_id").(string)
	segmentName := d.Get("segment_name").(string)
	environmentID := d.Get("environment_id").(string)

	err := c.APIClient.SegmentEnvironmentAssociations.Delete(workspaceID, segmentName, environmentID)
	if err != nil {
		// Handle 404 errors gracefully - association may have already been removed
		if strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "not found") {
			return nil
		}
		return diag.FromErr(err)
	}

	return nil
}