package provider_registry

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceInfraProviderVersionPublish() *schema.Resource {
	resource := &schema.Resource{
		Description:   "Resource for publishing Terraform/OpenTofu Provider Versions in the IaCM Provider Registry. Publishing makes the provider version available for use.",
		ReadContext:   resourceInfraProviderVersionPublishRead,
		CreateContext: resourceInfraProviderVersionPublishCreate,
		DeleteContext: resourceInfraProviderVersionPublishDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceInfraProviderVersionPublishImport,
		},

		Schema: map[string]*schema.Schema{
			"provider_id": {
				Description: "The ID of the provider.",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
			},
			"version": {
				Description: "Provider version number to publish.",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
			},
			"published": {
				Description: "Indicates if the provider version is published.",
				Type:        schema.TypeBool,
				Computed:    true,
			},
		},
	}
	return resource
}

func resourceInfraProviderVersionPublishRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c, ctx := m.(*internal.Session).GetPlatformClientWithContext(ctx)

	providerId := d.Get("provider_id").(string)
	version := d.Get("version").(string)

	// Get the provider version to check if it exists
	_, httpRes, err := c.ProviderRegistryApi.ProviderRegistryGetProviderVersion(
		ctx,
		providerId,
		version,
		c.AccountId,
	)
	if err != nil {
		return helpers.HandleApiError(err, d, httpRes)
	}

	// If we can read the version, consider it published
	d.SetId(fmt.Sprintf("%s/%s", providerId, version))
	d.Set("provider_id", providerId)
	d.Set("version", version)
	d.Set("published", true)

	return nil
}

func resourceInfraProviderVersionPublishCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c, ctx := m.(*internal.Session).GetPlatformClientWithContext(ctx)

	providerId := d.Get("provider_id").(string)
	version := d.Get("version").(string)

	log.Printf("[DEBUG] Publishing provider version %s for provider %s", version, providerId)

	httpRes, err := c.ProviderRegistryApi.ProviderRegistryPublishProviderVersion(
		ctx,
		providerId,
		version,
		c.AccountId,
	)

	if err != nil {
		return parseError(err, httpRes)
	}

	d.SetId(fmt.Sprintf("%s/%s", providerId, version))
	d.Set("published", true)

	return resourceInfraProviderVersionPublishRead(ctx, d, m)
}

func resourceInfraProviderVersionPublishDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Publishing is a one-way operation - once published, a version cannot be unpublished
	// On delete, we just remove from state
	log.Printf("[DEBUG] Removing published provider version from state (cannot unpublish)")
	d.SetId("")
	return nil
}

func resourceInfraProviderVersionPublishImport(ctx context.Context, d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	parts := strings.Split(d.Id(), "/")
	if len(parts) != 2 {
		return nil, fmt.Errorf("invalid import ID format, expected: provider_id/version")
	}

	d.Set("provider_id", parts[0])
	d.Set("version", parts[1])
	d.SetId(d.Id())

	diags := resourceInfraProviderVersionPublishRead(ctx, d, m)
	if diags.HasError() {
		return nil, fmt.Errorf("failed to read published provider version: %v", diags)
	}

	return []*schema.ResourceData{d}, nil
}
