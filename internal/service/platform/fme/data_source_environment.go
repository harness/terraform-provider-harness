package fme

import (
	"context"

	"github.com/harness/terraform-provider-harness/internal"
	"github.com/harness/terraform-provider-harness/internal/service/platform/fme/api"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func DataSourceFMEEnvironment() *schema.Resource {
	return &schema.Resource{
		Description: "Data source for retrieving a FME (Feature Management Engine) environment.",
		ReadContext: dataSourceFMEEnvironmentRead,

		Schema: map[string]*schema.Schema{
			"workspace_id": {
				Description:  "Unique identifier of the workspace.",
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.IsUUID,
			},
			"name": {
				Description:  "Name of the environment.",
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsNotEmpty,
			},
			"id": {
				Description: "Unique identifier of the environment.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"production": {
				Description: "Whether this is a production environment.",
				Type:        schema.TypeBool,
				Computed:    true,
			},
		},
	}
}

func dataSourceFMEEnvironmentRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	session := meta.(*internal.Session)
	fmeClient := session.FMEClient

	if fmeClient == nil {
		return diag.Errorf("FME client not configured")
	}

	c := fmeClient.(*FMEConfig)
	workspaceID := d.Get("workspace_id").(string)
	environmentName := d.Get("name").(string)

	// Get all environments and find the one with the matching name
	environments, err := c.APIClient.Environments.List(workspaceID)
	if err != nil {
		return diag.FromErr(err)
	}

	var environment *api.Environment
	for _, env := range environments {
		if env.Name != nil && *env.Name == environmentName {
			environment = env
			break
		}
	}

	if environment == nil {
		return diag.Errorf("environment with name %s not found in workspace %s", environmentName, workspaceID)
	}

	d.SetId(*environment.ID)
	d.Set("name", environment.Name)
	d.Set("production", environment.Production)

	return nil
}