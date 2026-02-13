package provider_registry

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceInfraProviderVersionFile() *schema.Resource {
	resource := &schema.Resource{
		Description:   "Resource for uploading files to Terraform/OpenTofu Provider Versions in the IaCM Provider Registry.",
		ReadContext:   resourceInfraProviderVersionFileRead,
		CreateContext: resourceInfraProviderVersionFileCreate,
		DeleteContext: resourceInfraProviderVersionFileDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceInfraProviderVersionFileImport,
		},

		Schema: map[string]*schema.Schema{
			"provider_id": {
				Description: "The ID of the provider.",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
			},
			"version": {
				Description: "Provider version number.",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
			},
			"filename": {
				Description: "Name of the file to upload (e.g., terraform-provider-aws_5.0.0_linux_amd64.zip). If not provided, will be derived from file_path.",
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
			},
			"file_path": {
				Description: "Local path to the file to upload. Required for uploading file content.",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
			},
			"synced": {
				Description: "Indicates if the provider version is synced after file upload.",
				Type:        schema.TypeBool,
				Computed:    true,
			},
		},
	}
	return resource
}

func resourceInfraProviderVersionFileRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c, ctx := m.(*internal.Session).GetPlatformClientWithContext(ctx)

	providerId := d.Get("provider_id").(string)
	version := d.Get("version").(string)
	filename := d.Get("filename").(string)

	// Get the provider version to check if the file exists
	resp, httpRes, err := c.ProviderRegistryApi.ProviderRegistryGetProviderVersion(
		ctx,
		providerId,
		version,
		c.AccountId,
	)
	if err != nil {
		return helpers.HandleApiError(err, d, httpRes)
	}

	// Check if file exists in the provider's metadata
	// Get full provider details to see files
	providerResp, httpRes, err := c.ProviderRegistryApi.ProviderRegistryGetProvider(
		ctx,
		providerId,
		c.AccountId,
	)
	if err != nil {
		return helpers.HandleApiError(err, d, httpRes)
	}

	// Find the version and check if file exists
	fileExists := false
	for _, v := range providerResp.Versions {
		if v.Version == version {
			for _, f := range v.Files {
				if f == filename {
					fileExists = true
					break
				}
			}
			break
		}
	}

	if !fileExists {
		d.SetId("")
		return nil
	}

	d.SetId(fmt.Sprintf("%s/%s/%s", providerId, version, filename))
	d.Set("provider_id", providerId)
	d.Set("version", version)
	d.Set("filename", filename)

	// Set synced from the response if available
	if len(resp.Protocols) > 0 {
		d.Set("synced", true)
	}

	return nil
}

func resourceInfraProviderVersionFileCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c, ctx := m.(*internal.Session).GetPlatformClientWithContext(ctx)

	providerId := d.Get("provider_id").(string)
	version := d.Get("version").(string)
	filename := d.Get("filename").(string)
	filePath := d.Get("file_path").(string)

	// Validate that at least one of filename or file_path is provided
	if filename == "" && filePath == "" {
		return diag.Errorf("Either 'filename' or 'file_path' must be provided")
	}

	// If filename is not provided, derive it from file_path
	if filename == "" && filePath != "" {
		filename = filePath[strings.LastIndex(filePath, "/")+1:]
		if filename == "" {
			return diag.Errorf("Could not derive filename from file_path: %s", filePath)
		}
		log.Printf("[DEBUG] Derived filename '%s' from file_path '%s'", filename, filePath)
		d.Set("filename", filename)
	}

	log.Printf("[DEBUG] Uploading file %s for provider %s version %s", filename, providerId, version)

	// Read file content if file_path is provided
	var fileData []byte
	if filePath != "" {
		log.Printf("[DEBUG] Reading file from path: %s", filePath)

		var err error
		fileData, err = os.ReadFile(filePath)
		if err != nil {
			return diag.Errorf("Failed to read file %s: %v", filePath, err)
		}
		log.Printf("[DEBUG] Read %d bytes from file %s", len(fileData), filePath)
	}

	// Upload file using SDK method
	contentDisposition := fmt.Sprintf("attachment; filename=\"%s\"", filename)

	resp, httpRes, err := c.ProviderRegistryApi.ProviderRegistryPostFiles(
		ctx,
		providerId,
		version,
		c.AccountId,
		contentDisposition,
		fileData,
	)

	if err != nil {
		log.Printf("[ERROR] Failed to upload file: %v", err)
		return parseError(err, httpRes)
	}

	log.Printf("[DEBUG] Successfully uploaded file %s", filename)
	d.Set("synced", resp.Synced)

	d.SetId(fmt.Sprintf("%s/%s/%s", providerId, version, filename))
	return resourceInfraProviderVersionFileRead(ctx, d, m)
}

func resourceInfraProviderVersionFileDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c, ctx := m.(*internal.Session).GetPlatformClientWithContext(ctx)

	providerId := d.Get("provider_id").(string)
	version := d.Get("version").(string)
	filename := d.Get("filename").(string)
	synced := d.Get("synced").(bool)

	// If the file is synced, it cannot be deleted independently
	// It will only be deleted when the provider itself is deleted
	if synced {
		log.Printf("[WARN] File %s is synced and cannot be deleted independently. It will be deleted when provider %s is deleted. Removing from state.", filename, providerId)
		d.SetId("")
		return nil
	}

	httpRes, err := c.ProviderRegistryApi.ProviderRegistryDeleteFile(
		ctx,
		providerId,
		version,
		filename,
		c.AccountId,
	)
	if err != nil {
		// If resource is already gone (404), treat as success
		if httpRes != nil && httpRes.StatusCode == 404 {
			log.Printf("[INFO] File %s not found, already deleted. Removing from state.", filename)
			d.SetId("")
			return nil
		}
		// If we get a 400/409 error indicating the file is synced, handle gracefully
		if httpRes != nil && (httpRes.StatusCode == 400 || httpRes.StatusCode == 409) {
			log.Printf("[WARN] File %s cannot be deleted (likely synced). Removing from state.", filename)
			d.SetId("")
			return nil
		}
		return parseError(err, httpRes)
	}
	return nil
}

func resourceInfraProviderVersionFileImport(ctx context.Context, d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	parts := strings.Split(d.Id(), "/")
	if len(parts) != 3 {
		return nil, fmt.Errorf("invalid import ID format, expected: provider_id/version/filename")
	}

	d.Set("provider_id", parts[0])
	d.Set("version", parts[1])
	d.Set("filename", parts[2])
	d.SetId(d.Id())

	diags := resourceInfraProviderVersionFileRead(ctx, d, m)
	if diags.HasError() {
		return nil, fmt.Errorf("failed to read provider version file: %v", diags)
	}

	return []*schema.ResourceData{d}, nil
}
