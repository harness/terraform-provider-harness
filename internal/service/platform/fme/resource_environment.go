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

func ResourceFMEEnvironment() *schema.Resource {
	return &schema.Resource{
		Description:   "Resource for creating a FME (Feature Management Engine) environment.",
		ReadContext:   resourceFMEEnvironmentRead,
		CreateContext: resourceFMEEnvironmentCreate,
		UpdateContext: resourceFMEEnvironmentUpdate,
		DeleteContext: resourceFMEEnvironmentDelete,
		Importer:      helpers.MultiLevelResourceImporter,

		Schema: map[string]*schema.Schema{
			"id": {
				Description: "Unique identifier of the environment.",
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
				Description:  "Name of the environment.",
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsNotEmpty,
			},
			"production": {
				Description: "Whether this is a production environment.",
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
			},
			"remove_environment_from_state_only": {
				Description: "If true, the environment will only be removed from Terraform state and not deleted from the API.",
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
			},
		},
	}
}

func resourceFMEEnvironmentCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	session := meta.(*internal.Session)
	fmeClient := session.FMEClient

	if fmeClient == nil {
		return diag.Errorf("FME client not configured")
	}

	c := fmeClient.(*FMEConfig)

	workspaceID := d.Get("workspace_id").(string)
	name := d.Get("name").(string)
	production := d.Get("production").(bool)

	req := &api.EnvironmentRequest{
		Name:       &name,
		Production: &production,
	}

	env, err := c.APIClient.Environments.Create(workspaceID, req)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(*env.ID)
	return resourceFMEEnvironmentRead(ctx, d, meta)
}

func resourceFMEEnvironmentRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	session := meta.(*internal.Session)
	fmeClient := session.FMEClient

	if fmeClient == nil {
		return diag.Errorf("FME client not configured")
	}

	c := fmeClient.(*FMEConfig)

	workspaceID := d.Get("workspace_id").(string)
	env, err := c.APIClient.Environments.Get(workspaceID, d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	if env == nil {
		d.SetId("")
		return nil
	}

	d.Set("name", env.Name)
	d.Set("production", env.Production)

	return nil
}

func resourceFMEEnvironmentUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	session := meta.(*internal.Session)
	c := session.FMEClient.(*FMEConfig)

	req := &api.EnvironmentRequest{}

	if d.HasChange("name") {
		name := d.Get("name").(string)
		req.Name = &name
	}

	if d.HasChange("production") {
		production := d.Get("production").(bool)
		req.Production = &production
	}

	workspaceID := d.Get("workspace_id").(string)
	_, err := c.APIClient.Environments.Update(workspaceID, d.Id(), req)
	if err != nil {
		return diag.FromErr(err)
	}

	return resourceFMEEnvironmentRead(ctx, d, meta)
}

func resourceFMEEnvironmentDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	session := meta.(*internal.Session)
	c := session.FMEClient.(*FMEConfig)

	// Check if we should only remove from state without calling API delete
	removeFromStateOnly := d.Get("remove_environment_from_state_only").(bool)

	if !removeFromStateOnly {
		workspaceID := d.Get("workspace_id").(string)
		err := c.APIClient.Environments.Delete(workspaceID, d.Id())
		if err != nil {
			return diag.FromErr(err)
		}
	}

	// Always clear the ID from state regardless of the option
	d.SetId("")
	return nil
}