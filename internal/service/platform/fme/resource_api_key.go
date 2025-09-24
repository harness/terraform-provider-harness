package fme

import (
	"context"

	"github.com/harness/terraform-provider-harness/internal"
	"github.com/harness/terraform-provider-harness/internal/service/platform/fme/api"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func ResourceFMEApiKey() *schema.Resource {
	return &schema.Resource{
		Description:   "Resource for creating a FME API key.",
		CreateContext: resourceFMEApiKeyCreate,
		ReadContext:   resourceFMEApiKeyRead,
		DeleteContext: resourceFMEApiKeyDelete,

		Schema: map[string]*schema.Schema{
			"id": {
				Description: "Unique identifier of the API key.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"workspace_id": {
				Description:  "ID of the workspace.",
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.IsUUID,
			},
			"environment_id": {
				Description:  "ID of the environment this API key belongs to.",
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringIsNotEmpty,
			},
			"name": {
				Description:  "Name of the API key.",
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringIsNotEmpty,
			},
			"type": {
				Description: "Type of the API key.",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				ValidateFunc: validation.StringInSlice([]string{
					"client_side",
					"server_side",
				}, false),
			},
			"key": {
				Description: "The actual API key value.",
				Type:        schema.TypeString,
				Computed:    true,
				Sensitive:   true,
			},
		},
	}
}

func resourceFMEApiKeyCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	session := meta.(*internal.Session)
	fmeClient := session.FMEClient

	if fmeClient == nil {
		return diag.Errorf("FME client not configured")
	}

	c := fmeClient.(*FMEConfig)

	workspaceID := d.Get("workspace_id").(string)
	environmentID := d.Get("environment_id").(string)
	name := d.Get("name").(string)
	keyType := d.Get("type").(string)

	workspaceType := "workspace"
	environmentType := "environment"
	req := &api.KeyRequest{
		Name:       &name,
		APIKeyType: &keyType,
		Environments: []*api.KeyEnvironmentRequest{
			{
				Type: &environmentType,
				ID:   &environmentID,
			},
		},
		Workspace: &api.KeyWorkspaceRequest{
			Type: &workspaceType,
			ID:   &workspaceID,
		},
	}

	key, err := c.APIClient.ApiKeys.Create(req)
	if err != nil {
		return diag.FromErr(err)
	}

	if key == nil {
		return diag.Errorf("API key creation returned nil result")
	}

	if key.Key == nil {
		return diag.Errorf("API key creation returned no key value")
	}

	d.SetId(*key.Key)
	return nil
}

func resourceFMEApiKeyRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// Note: Cannot refresh API key resource due to API limitations
	// This matches the Split.io provider pattern
	// The key value is stored as the resource ID, so set it in the key field for display
	d.Set("key", d.Id())
	return nil
}

func resourceFMEApiKeyDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	session := meta.(*internal.Session)
	fmeClient := session.FMEClient

	if fmeClient == nil {
		return diag.Errorf("FME client not configured")
	}

	c := fmeClient.(*FMEConfig)

	err := c.APIClient.ApiKeys.Delete(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}