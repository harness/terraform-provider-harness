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

func ResourceFMEFlagSet() *schema.Resource {
	return &schema.Resource{
		Description:   "Resource for creating a FME (Feature Management Engine) flag set.",
		ReadContext:   resourceFMEFlagSetRead,
		CreateContext: resourceFMEFlagSetCreate,
		UpdateContext: resourceFMEFlagSetUpdate,
		DeleteContext: resourceFMEFlagSetDelete,
		Importer:      helpers.MultiLevelResourceImporter,

		Schema: map[string]*schema.Schema{
			"id": {
				Description: "Unique identifier of the flag set.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"workspace_id": {
				Description:  "Unique identifier of the workspace.",
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.IsUUID,
			},
			"name": {
				Description:  "Name of the flag set.",
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsNotEmpty,
			},
			"description": {
				Description: "Description of the flag set.",
				Type:        schema.TypeString,
				Optional:    true,
			},
		},
	}
}

func resourceFMEFlagSetCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	session := meta.(*internal.Session)
	fmeClient := session.FMEClient

	if fmeClient == nil {
		return diag.Errorf("FME client not configured")
	}

	c := fmeClient.(*FMEConfig)
	workspaceID := d.Get("workspace_id").(string)
	name := d.Get("name").(string)
	description := d.Get("description").(string)

	workspaceType := "WORKSPACE"
	workspaceRef := &api.WorkspaceIDRef{
		Type: &workspaceType,
		ID:   &workspaceID,
	}

	req := &api.FlagSetRequest{
		Name:        &name,
		Description: &description,
		Workspace:   workspaceRef,
	}

	flagSet, err := c.APIClient.FlagSets.Create(req)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(*flagSet.ID)

	return resourceFMEFlagSetRead(ctx, d, meta)
}

func resourceFMEFlagSetRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	session := meta.(*internal.Session)
	fmeClient := session.FMEClient

	if fmeClient == nil {
		return diag.Errorf("FME client not configured")
	}

	c := fmeClient.(*FMEConfig)
	flagSetID := d.Id()

	flagSet, err := c.APIClient.FlagSets.Get(flagSetID)
	if err != nil {
		return diag.FromErr(err)
	}

	if flagSet == nil {
		d.SetId("")
		return nil
	}

	d.Set("name", flagSet.Name)
	d.Set("description", flagSet.Description)
	if flagSet.Workspace != nil && flagSet.Workspace.ID != nil {
		d.Set("workspace_id", *flagSet.Workspace.ID)
	}

	return nil
}

func resourceFMEFlagSetUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	session := meta.(*internal.Session)
	fmeClient := session.FMEClient

	if fmeClient == nil {
		return diag.Errorf("FME client not configured")
	}

	c := fmeClient.(*FMEConfig)
	flagSetID := d.Id()
	name := d.Get("name").(string)
	description := d.Get("description").(string)

	req := &api.FlagSetRequest{
		Name:        &name,
		Description: &description,
	}

	_, err := c.APIClient.FlagSets.Update(flagSetID, req)
	if err != nil {
		return diag.FromErr(err)
	}

	return resourceFMEFlagSetRead(ctx, d, meta)
}

func resourceFMEFlagSetDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	session := meta.(*internal.Session)
	fmeClient := session.FMEClient

	if fmeClient == nil {
		return diag.Errorf("FME client not configured")
	}

	c := fmeClient.(*FMEConfig)
	flagSetID := d.Id()

	err := c.APIClient.FlagSets.Delete(flagSetID)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")
	return nil
}