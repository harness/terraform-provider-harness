package provider_registry

import (
	"context"
	"log"

	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceInfraProviderSigningKey() *schema.Resource {
	resource := &schema.Resource{
		Description:   "Resource for managing GPG Signing Keys for Terraform/OpenTofu Providers in the IaCM Provider Registry.",
		ReadContext:   resourceInfraProviderSigningKeyRead,
		CreateContext: resourceInfraProviderSigningKeyCreate,
		UpdateContext: resourceInfraProviderSigningKeyUpdate,
		DeleteContext: resourceInfraProviderSigningKeyDelete,
		Importer:      helpers.AccountLevelResourceImporter,

		Schema: map[string]*schema.Schema{
			"key_id": {
				Description: "GPG key ID.",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
			},
			"key_name": {
				Description: "GPG key name.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"ascii_armor": {
				Description: "ASCII-armored GPG public key.",
				Type:        schema.TypeString,
				Required:    true,
				Sensitive:   true,
			},
			"user": {
				Description: "User who uploaded the key.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"id": {
				Description: "Unique identifier of the signing key.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"created_at": {
				Description: "Creation timestamp.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"updated_at": {
				Description: "Last updated timestamp.",
				Type:        schema.TypeString,
				Computed:    true,
			},
		},
	}
	return resource
}

func resourceInfraProviderSigningKeyRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c, ctx := m.(*internal.Session).GetPlatformClientWithContext(ctx)
	id := d.Id()
	if id == "" {
		d.MarkNewResource()
		return nil
	}

	// List all signing keys and find the one with matching ID
	resp, httpRes, err := c.ProviderRegistryApi.ProviderRegistryListSigningKeys(
		ctx,
		c.AccountId,
		nil,
	)
	if err != nil {
		return helpers.HandleApiError(err, d, httpRes)
	}

	// Find the signing key by ID
	var signingKey *nextgen.SigningKey
	for _, key := range resp {
		if key.Id == id {
			signingKey = &key
			break
		}
	}

	if signingKey == nil {
		d.SetId("")
		return nil
	}

	readSigningKey(d, signingKey)
	return nil
}

func resourceInfraProviderSigningKeyCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c, ctx := m.(*internal.Session).GetPlatformClientWithContext(ctx)

	createKey := buildUploadSigningKeyRequest(d)
	log.Printf("[DEBUG] Creating signing key with key_id %s ", createKey.KeyId)

	httpRes, err := c.ProviderRegistryApi.ProviderRegistryUploadSigningKey(
		ctx,
		createKey,
		c.AccountId,
	)

	if err != nil {
		log.Printf("[ERROR] Failed to create signing key: %v, HTTP Status: %v", err, httpRes)
		return parseError(err, httpRes)
	}

	listResp, httpRes, err := c.ProviderRegistryApi.ProviderRegistryListSigningKeys(
		ctx,
		c.AccountId,
		nil,
	)
	if err != nil {
		log.Printf("[ERROR] Failed to list signing keys after creation: %v", err)
		return parseError(err, httpRes)
	}

	log.Printf("[DEBUG] Found %d signing keys, searching for key_id %s", len(listResp), createKey.KeyId)

	// Find the newly created key by key_id
	for _, key := range listResp {
		if key.KeyId == createKey.KeyId {
			log.Printf("[DEBUG] Found created signing key with ID: %s", key.Id)
			d.SetId(key.Id)
			readSigningKey(d, &key)
			return nil
		}
	}

	return diag.Errorf("Failed to find created signing key with key_id %s", createKey.KeyId)
}

func resourceInfraProviderSigningKeyDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c, ctx := m.(*internal.Session).GetPlatformClientWithContext(ctx)
	id := d.Id()
	if id == "" {
		return nil
	}

	httpRes, err := c.ProviderRegistryApi.ProviderRegistryDeleteSigningKey(
		ctx,
		id,
		c.AccountId,
	)
	if err != nil {
		return parseError(err, httpRes)
	}
	return nil
}

func resourceInfraProviderSigningKeyUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c, ctx := m.(*internal.Session).GetPlatformClientWithContext(ctx)
	id := d.Id()

	if d.HasChanges("key_name", "ascii_armor", "key_id", "user") {
		updateKey := buildUploadSigningKeyRequest(d)
		log.Printf("[DEBUG] Updating signing key %s", id)

		httpRes, err := c.ProviderRegistryApi.ProviderRegistryUpdateSigningKey(
			ctx,
			updateKey,
			c.AccountId,
			id,
		)
		if err != nil {
			return parseError(err, httpRes)
		}
	}

	return resourceInfraProviderSigningKeyRead(ctx, d, m)
}

func readSigningKey(d *schema.ResourceData, key *nextgen.SigningKey) {
	d.SetId(key.Id)
	d.Set("id", key.Id)
	d.Set("key_id", key.KeyId)
	d.Set("key_name", key.KeyName)
	d.Set("ascii_armor", key.AsciiArmor)
	d.Set("user", key.User)
	d.Set("created_at", key.CreatedAt)
	d.Set("updated_at", key.UpdatedAt)
}

func buildUploadSigningKeyRequest(d *schema.ResourceData) nextgen.UploadSigningKeyRequest {
	request := nextgen.UploadSigningKeyRequest{
		KeyId:      d.Get("key_id").(string),
		KeyName:    d.Get("key_name").(string),
		AsciiArmor: d.Get("ascii_armor").(string),
		User:       d.Get("user").(string),
	}
	return request
}
