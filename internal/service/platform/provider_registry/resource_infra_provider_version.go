package provider_registry

import (
	"context"
	"fmt"
	"io"
	"log"
	"strings"

	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceInfraProviderVersion() *schema.Resource {
	resource := &schema.Resource{
		Description:   "Resource for managing Terraform/OpenTofu Provider Versions in the IaCM Provider Registry.",
		ReadContext:   resourceInfraProviderVersionRead,
		CreateContext: resourceInfraProviderVersionCreate,
		UpdateContext: resourceInfraProviderVersionUpdate,
		DeleteContext: resourceInfraProviderVersionDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceInfraProviderVersionImport,
		},

		Schema: map[string]*schema.Schema{
			"provider_id": {
				Description: "The ID of the provider this version belongs to.",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
			},
			"version": {
				Description: "Version number (e.g., 1.0.0).",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
			},
			"gpg_key_id": {
				Description: "GPG key ID for signing.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"protocols": {
				Description: "Supported Terraform protocol versions (e.g., ['5.0', '6.0']).",
				Type:        schema.TypeList,
				Required:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
	return resource
}

func resourceInfraProviderVersionRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c, ctx := m.(*internal.Session).GetPlatformClientWithContext(ctx)

	providerId := d.Get("provider_id").(string)
	version := d.Get("version").(string)

	resp, httpRes, err := c.ProviderRegistryApi.ProviderRegistryGetProviderVersion(
		ctx,
		providerId,
		version,
		c.AccountId,
	)
	if err != nil {
		return helpers.HandleApiError(err, d, httpRes)
	}
	readProviderVersion(d, &resp, providerId, version)
	return nil
}

func resourceInfraProviderVersionCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c, ctx := m.(*internal.Session).GetPlatformClientWithContext(ctx)

	providerId := d.Get("provider_id").(string)
	version := d.Get("version").(string)

	createVersion := buildCreateProviderVersionRequestBody(d)
	log.Printf("[DEBUG] Creating provider version %s for provider %s with body %v", version, providerId, createVersion)

	httpRes, err := c.ProviderRegistryApi.ProviderRegistryCreateProviderVersion(
		ctx,
		createVersion,
		c.AccountId,
		providerId,
	)

	if err != nil {
		log.Printf("[ERROR] Failed to create provider version: %v, HTTP Response: %+v", err, httpRes)
		if httpRes != nil && httpRes.Body != nil {
			bodyBytes, _ := io.ReadAll(httpRes.Body)
			log.Printf("[ERROR] Response body: %s", string(bodyBytes))
		}
		return parseError(err, httpRes)
	}

	d.SetId(fmt.Sprintf("%s/%s", providerId, version))
	return resourceInfraProviderVersionRead(ctx, d, m)
}

func resourceInfraProviderVersionDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c, ctx := m.(*internal.Session).GetPlatformClientWithContext(ctx)

	providerId := d.Get("provider_id").(string)
	version := d.Get("version").(string)

	httpRes, err := c.ProviderRegistryApi.ProviderRegistryDeleteProviderVersion(
		ctx,
		providerId,
		version,
		c.AccountId,
	)
	if err != nil {
		// If resource is already gone (404), treat as success
		if httpRes != nil && httpRes.StatusCode == 404 {
			log.Printf("[INFO] Version %s not found, already deleted. Removing from state.", version)
			d.SetId("")
			return nil
		}
		// If we get a 400/409 error indicating the version is published and cannot be deleted
		if httpRes != nil && (httpRes.StatusCode == 400 || httpRes.StatusCode == 409) {
			log.Printf("[WARN] Version %s is published and cannot be deleted independently. It will be deleted when provider %s is deleted. Removing from state.", version, providerId)
			d.SetId("")
			return nil
		}
		return parseError(err, httpRes)
	}
	return nil
}

func resourceInfraProviderVersionUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c, ctx := m.(*internal.Session).GetPlatformClientWithContext(ctx)

	providerId := d.Get("provider_id").(string)
	version := d.Get("version").(string)

	if d.HasChanges("gpg_key_id", "protocols") {
		updateVersion := buildUpdateProviderVersionRequestBody(d)
		log.Printf("[DEBUG] Updating provider version %s for provider %s with body %v", version, providerId, updateVersion)

		httpRes, err := c.ProviderRegistryApi.ProviderRegistryUpdateProviderVersion(
			ctx,
			updateVersion,
			c.AccountId,
			providerId,
			version,
		)
		if err != nil {
			return parseError(err, httpRes)
		}
	}

	return resourceInfraProviderVersionRead(ctx, d, m)
}

func resourceInfraProviderVersionImport(ctx context.Context, d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	parts := strings.Split(d.Id(), "/")
	if len(parts) != 2 {
		return nil, fmt.Errorf("invalid import ID format, expected: provider_id/version")
	}

	d.Set("provider_id", parts[0])
	d.Set("version", parts[1])
	d.SetId(d.Id())

	diags := resourceInfraProviderVersionRead(ctx, d, m)
	if diags.HasError() {
		return nil, fmt.Errorf("failed to read provider version: %v", diags)
	}

	return []*schema.ResourceData{d}, nil
}

func readProviderVersion(d *schema.ResourceData, version *nextgen.GetProviderVersionResponseBody, providerId string, versionNum string) {
	d.SetId(fmt.Sprintf("%s/%s", providerId, versionNum))
	d.Set("provider_id", providerId)
	d.Set("version", versionNum)
	d.Set("gpg_key_id", version.GpgKeyId)
	d.Set("protocols", version.Protocols)
}

func buildCreateProviderVersionRequestBody(d *schema.ResourceData) nextgen.CreateProviderVersionRequestBody {
	version := nextgen.CreateProviderVersionRequestBody{
		Version:  d.Get("version").(string),
		GpgKeyId: d.Get("gpg_key_id").(string),
	}

	if protocols, ok := d.GetOk("protocols"); ok {
		protocolList := protocols.([]interface{})
		version.Protocol = make([]string, len(protocolList))
		for i, p := range protocolList {
			version.Protocol[i] = p.(string)
		}
	}

	return version
}

func buildUpdateProviderVersionRequestBody(d *schema.ResourceData) nextgen.UpdateProviderVersionRequestBody {
	version := nextgen.UpdateProviderVersionRequestBody{
		GpgKeyId: d.Get("gpg_key_id").(string),
	}

	if protocols, ok := d.GetOk("protocols"); ok {
		protocolList := protocols.([]interface{})
		version.Protocol = make([]string, len(protocolList))
		for i, p := range protocolList {
			version.Protocol[i] = p.(string)
		}
	}

	return version
}
