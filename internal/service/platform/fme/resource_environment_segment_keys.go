package fme

import (
	"context"
	"fmt"
	"strings"

	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func ResourceFMEEnvironmentSegmentKeys() *schema.Resource {
	return &schema.Resource{
		Description:   "Resource for managing FME environment segment keys.",
		ReadContext:   resourceFMEEnvironmentSegmentKeysRead,
		CreateContext: resourceFMEEnvironmentSegmentKeysCreate,
		UpdateContext: resourceFMEEnvironmentSegmentKeysUpdate,
		DeleteContext: resourceFMEEnvironmentSegmentKeysDelete,
		Importer:      helpers.MultiLevelResourceImporter,

		Schema: map[string]*schema.Schema{
			"workspace_id": {
				Description:  "ID of the workspace this segment belongs to.",
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
			"keys": {
				Description: "List of keys for the segment in this environment.",
				Type:        schema.TypeSet,
				Required:    true,
				Elem: &schema.Schema{
					Type:         schema.TypeString,
					ValidateFunc: validation.StringIsNotEmpty,
				},
			},
		},
	}
}

func resourceFMEEnvironmentSegmentKeysCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	session := meta.(*internal.Session)
	fmeClient := session.FMEClient

	if fmeClient == nil {
		return diag.Errorf("FME client not configured")
	}

	c := fmeClient.(*FMEConfig)

	workspaceID := d.Get("workspace_id").(string)
	segmentName := d.Get("segment_name").(string)
	environmentID := d.Get("environment_id").(string)

	keysSet := d.Get("keys").(*schema.Set)
	keys := make([]string, keysSet.Len())
	for i, key := range keysSet.List() {
		keys[i] = key.(string)
	}

	_, err := c.APIClient.Environments.AddSegmentKeys(environmentID, segmentName, true, keys)
	if err != nil {
		return diag.FromErr(err)
	}

	// Use a composite ID to uniquely identify this resource
	d.SetId(fmt.Sprintf("%s:%s:%s", workspaceID, segmentName, environmentID))
	return resourceFMEEnvironmentSegmentKeysRead(ctx, d, meta)
}

func resourceFMEEnvironmentSegmentKeysRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	session := meta.(*internal.Session)
	fmeClient := session.FMEClient

	if fmeClient == nil {
		return diag.Errorf("FME client not configured")
	}

	c := fmeClient.(*FMEConfig)
	segmentName := d.Get("segment_name").(string)
	environmentID := d.Get("environment_id").(string)

	keys, err := c.APIClient.Environments.GetSegmentKeys(environmentID, segmentName)
	if err != nil {
		// Handle 404 errors gracefully - segment may have been deactivated from environment
		if strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "There is no segment") {
			d.SetId("")
			return nil
		}
		return diag.FromErr(err)
	}

	if keys == nil {
		d.SetId("")
		return nil
	}

	d.Set("segment_name", segmentName)
	d.Set("environment_id", environmentID)
	d.Set("keys", keys)

	return nil
}

func resourceFMEEnvironmentSegmentKeysUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	session := meta.(*internal.Session)
	fmeClient := session.FMEClient

	if fmeClient == nil {
		return diag.Errorf("FME client not configured")
	}

	c := fmeClient.(*FMEConfig)

	segmentName := d.Get("segment_name").(string)
	environmentID := d.Get("environment_id").(string)

	if d.HasChange("keys") {
		keysSet := d.Get("keys").(*schema.Set)
		keys := make([]string, keysSet.Len())
		for i, key := range keysSet.List() {
			keys[i] = key.(string)
		}

		_, err := c.APIClient.Environments.AddSegmentKeys(environmentID, segmentName, true, keys)
		if err != nil {
			return diag.FromErr(err)
		}
	}

	return resourceFMEEnvironmentSegmentKeysRead(ctx, d, meta)
}

func resourceFMEEnvironmentSegmentKeysDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	session := meta.(*internal.Session)
	fmeClient := session.FMEClient

	if fmeClient == nil {
		return diag.Errorf("FME client not configured")
	}

	c := fmeClient.(*FMEConfig)
	segmentName := d.Get("segment_name").(string)
	environmentID := d.Get("environment_id").(string)

	// Get current keys to remove them all
	keys, err := c.APIClient.Environments.GetSegmentKeys(environmentID, segmentName)
	if err != nil {
		// Handle 404 errors gracefully - segment may have been deactivated from environment
		if strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "There is no segment") {
			return nil
		}
		return diag.FromErr(err)
	}

	if len(keys) > 0 {
		err = c.APIClient.Environments.RemoveSegmentKeys(environmentID, segmentName, keys)
		if err != nil {
			// Handle 404 errors gracefully - segment may have been deactivated from environment
			if strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "There is no segment") {
				return nil
			}
			return diag.FromErr(err)
		}
	}

	return nil
}