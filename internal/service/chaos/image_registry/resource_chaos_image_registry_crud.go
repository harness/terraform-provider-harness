package image_registry

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/harness/harness-go-sdk/harness/chaos/graphql/model"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceChaosImageRegistryCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*internal.Session).ChaosClient

	var infraID *string
	if v, ok := d.GetOk("infra_id"); ok {
		infraIDVal := v.(string)
		infraID = &infraIDVal
	}

	req := &model.ImageRegistryRequest{
		RegistryServer:    d.Get("registry_server").(string),
		RegistryAccount:   d.Get("registry_account").(string),
		InfraID:           infraID,
		IsPrivate:         d.Get("is_private").(bool),
		IsDefault:         d.Get("is_default").(bool),
		IsOverrideAllowed: d.Get("is_override_allowed").(bool),
		UseCustomImages:   d.Get("use_custom_images").(bool),
	}

	if v, ok := d.GetOk("secret_name"); ok {
		secretName := v.(string)
		req.SecretName = &secretName
	}

	if v, ok := d.GetOk("custom_images"); ok && len(v.([]interface{})) > 0 {
		customImages := v.([]interface{})[0].(map[string]interface{})
		req.CustomImages = &model.CustomImages{
			LogWatcher: getStringPtr(customImages["log_watcher"]),
			Ddcr:       getStringPtr(customImages["ddcr"]),
			DdcrLib:    getStringPtr(customImages["ddcr_lib"]),
			DdcrFault:  getStringPtr(customImages["ddcr_fault"]),
		}
	}

	identifiers := getIdentifiers(d, c.AccountId)
	_, err := c.ImageRegistryApi.Create(
		ctx,
		identifiers,
		req.RegistryServer,  // string
		req.RegistryAccount, // string
		req.IsPrivate,       // bool
		// Add functional options if needed
		func(r model.ImageRegistryRequest) model.ImageRegistryRequest {
			if req.InfraID != nil {
				r.InfraID = req.InfraID
			}
			r.IsDefault = req.IsDefault
			r.IsOverrideAllowed = req.IsOverrideAllowed
			r.UseCustomImages = req.UseCustomImages
			r.SecretName = req.SecretName
			r.CustomImages = req.CustomImages
			return r
		},
	)

	// If the registry already exists (duplicate key error), log a warning instead of erroring
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key error") {
			log.Printf("[WARN] Chaos image registry already exists: %s", err)
			// Still set the ID and read the state
			d.SetId(generateID(identifiers))
			return resourceChaosImageRegistryRead(ctx, d, meta)
		}
		return diag.Errorf("failed to create image registry: %v", err)
	}

	d.SetId(generateID(identifiers))

	d.SetId(generateID(identifiers))
	return resourceChaosImageRegistryRead(ctx, d, meta)
}

func resourceChaosImageRegistryRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*internal.Session).ChaosClient
	identifiers, err := parseID(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	var infraID *string
	if v, ok := d.GetOk("infra_id"); ok {
		infraIDVal := v.(string)
		infraID = &infraIDVal
	}

	registry, err := c.ImageRegistryApi.Get(ctx, *identifiers, infraID)
	if err != nil {
		return diag.Errorf("failed to read image registry: %v", err)
	}

	d.Set("org_id", identifiers.OrgIdentifier)
	d.Set("project_id", identifiers.ProjectIdentifier)
	d.Set("infra_id", registry.InfraID)
	d.Set("registry_server", registry.RegistryServer)
	d.Set("registry_account", registry.RegistryAccount)
	d.Set("is_private", registry.IsPrivate)
	d.Set("is_default", registry.IsDefault)
	d.Set("is_override_allowed", registry.IsOverrideAllowed)
	d.Set("use_custom_images", registry.UseCustomImages)
	d.Set("created_at", registry.CreatedAt)
	d.Set("updated_at", registry.UpdatedAt)

	if registry.SecretName != nil {
		d.Set("secret_name", *registry.SecretName)
	}

	if registry.CustomImages != nil {
		customImages := map[string]interface{}{}
		if registry.CustomImages.LogWatcher != nil {
			customImages["log_watcher"] = *registry.CustomImages.LogWatcher
		}
		if registry.CustomImages.Ddcr != nil {
			customImages["ddcr"] = *registry.CustomImages.Ddcr
		}
		if registry.CustomImages.DdcrLib != nil {
			customImages["ddcr_lib"] = *registry.CustomImages.DdcrLib
		}
		if registry.CustomImages.DdcrFault != nil {
			customImages["ddcr_fault"] = *registry.CustomImages.DdcrFault
		}
		d.Set("custom_images", []map[string]interface{}{customImages})
	}

	return nil
}

func resourceChaosImageRegistryUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*internal.Session).ChaosClient
	identifiers, err := parseID(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	var infraID *string
	if v, ok := d.GetOk("infra_id"); ok {
		infraIDVal := v.(string)
		infraID = &infraIDVal
	}

	req := &model.ImageRegistryRequest{
		RegistryServer:    d.Get("registry_server").(string),
		RegistryAccount:   d.Get("registry_account").(string),
		InfraID:           infraID,
		IsPrivate:         d.Get("is_private").(bool),
		IsDefault:         d.Get("is_default").(bool),
		IsOverrideAllowed: d.Get("is_override_allowed").(bool),
		UseCustomImages:   d.Get("use_custom_images").(bool),
	}

	if d.HasChange("secret_name") {
		if v, ok := d.GetOk("secret_name"); ok {
			secretName := v.(string)
			req.SecretName = &secretName
		}
	}

	if d.HasChange("custom_images") && d.Get("use_custom_images").(bool) {
		if v, ok := d.GetOk("custom_images"); ok && len(v.([]interface{})) > 0 {
			customImages := v.([]interface{})[0].(map[string]interface{})
			req.CustomImages = &model.CustomImages{
				LogWatcher: getStringPtr(customImages["log_watcher"]),
				Ddcr:       getStringPtr(customImages["ddcr"]),
				DdcrLib:    getStringPtr(customImages["ddcr_lib"]),
				DdcrFault:  getStringPtr(customImages["ddcr_fault"]),
			}
		}
	}

	_, err = c.ImageRegistryApi.Update(
		ctx,
		*identifiers,
		req.InfraID,         // *string
		req.RegistryServer,  // string
		req.RegistryAccount, // string
		req.IsPrivate,       // bool
		// Functional options for additional fields
		func(r model.ImageRegistryRequest) model.ImageRegistryRequest {
			if req.InfraID != nil {
				r.InfraID = req.InfraID
			}
			r.IsDefault = req.IsDefault
			r.IsOverrideAllowed = req.IsOverrideAllowed
			r.UseCustomImages = req.UseCustomImages
			if req.SecretName != nil {
				r.SecretName = req.SecretName
			}
			if req.CustomImages != nil {
				r.CustomImages = req.CustomImages
			}
			return r
		},
	)
	if err != nil {
		return diag.Errorf("failed to update image registry: %v", err)
	}

	return resourceChaosImageRegistryRead(ctx, d, meta)
}

func resourceChaosImageRegistryDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// Nothing to do here as registry configuration is reset using is_default
	d.SetId("")
	return nil
}

// Helper functions
func getIdentifiers(d *schema.ResourceData, accountID string) model.ScopedIdentifiersRequest {
	identifiers := model.ScopedIdentifiersRequest{
		AccountIdentifier: accountID,
	}

	if v, ok := d.GetOk("org_id"); ok {
		orgID := v.(string)
		identifiers.OrgIdentifier = &orgID
	}

	if v, ok := d.GetOk("project_id"); ok {
		projectID := v.(string)
		identifiers.ProjectIdentifier = &projectID
	}

	return identifiers
}

func generateID(identifiers model.ScopedIdentifiersRequest) string {
	parts := []string{identifiers.AccountIdentifier}
	if identifiers.OrgIdentifier != nil {
		parts = append(parts, *identifiers.OrgIdentifier)
		if identifiers.ProjectIdentifier != nil {
			parts = append(parts, *identifiers.ProjectIdentifier)
		}
	}
	return strings.Join(parts, "/")
}

func parseID(id string) (*model.ScopedIdentifiersRequest, error) {
	parts := strings.Split(id, "/")
	if len(parts) < 2 || len(parts) > 4 {
		return nil, fmt.Errorf("invalid ID format, expected account_id/org_id/project_id or account_id/org_id/project_id/registry_account, got: %s", id)
	}

	identifiers := &model.ScopedIdentifiersRequest{
		AccountIdentifier: parts[0],
	}

	identifiers.OrgIdentifier = &parts[1]
	if len(parts) > 2 {
		identifiers.ProjectIdentifier = &parts[2]
	}

	// If we have 4 parts, the last one is the registry account
	if len(parts) == 4 {
		registryAccount := parts[3]
		if registryAccount == "" {
			return nil, fmt.Errorf("registry account cannot be empty in: %s", id)
		}
	}

	return identifiers, nil
}

func getStringPtr(value interface{}) *string {
	if value == nil {
		return nil
	}
	str := value.(string)
	if str == "" {
		return nil
	}
	return &str
}
