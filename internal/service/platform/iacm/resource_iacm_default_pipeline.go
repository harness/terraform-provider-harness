package iacm

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceIacmDefaultPipeline() *schema.Resource {
	resource := &schema.Resource{
		Description: "Resource for managing IAC default pipelines",

		ReadContext:   resourceIacmDefaultPipelineRead,
		DeleteContext: resourceIacmDefaultPipelineDelete,
		CreateContext: resourceIacmDefaultPipelineCreate,
		UpdateContext: resourceIacmDefaultPipelineUpdate,
		Importer:      helpers.ProjectResourceImporter,

		Schema: map[string]*schema.Schema{
			"org_id": {
				Description: "Organization identifier of the organization the default pipelines resides in.",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
			},
			"project_id": {
				Description: "Project identifier of the project the default pipelines resides in.",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
			},
			"provisioner_type": {
				Description: "The provisioner associated with this default.",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    false,
			},
			"operation": {
				Description: "The operation associated with this default.",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    false,
			},
			"pipeline": {
				Description: "The pipeline associated with this default.",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    false,
			},
		},
	}

	return resource
}

func resourceIacmDefaultPipelineRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)
	id := d.Id()
	if id == "" {
		d.MarkNewResource()
	}
	defaultPipelines, httpResp, err := c.SettingsApi.SettingsListDefaultPipelines(
		ctx,
		d.Get("org_id").(string),
		d.Get("project_id").(string),
		c.AccountId,
		nil,
	)
	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}
	for _, defaultPipeline := range defaultPipelines {
		if defaultPipeline.Provisioner == d.Get("provisioner_type").(string) &&
			defaultPipeline.Operation == d.Get("operation").(string) {
			d.SetId(fmt.Sprintf("%s-%s", defaultPipeline.Provisioner, defaultPipeline.Operation))
			d.Set("pipeline", defaultPipeline.Pipeline)
			return nil
		}
	}
	if d.Id() == "" {
		return diag.Errorf("resource with provisioner type '%s' and operation '%s' not found", d.Get("provisioner_type"), d.Get("operation"))
	}
	return diag.Errorf("resource with ID %s not found", d.Id())
}

func resourceIacmDefaultPipelineDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)
	id := d.Id()
	if id == "" {
		return nil
	}

	httpResp, err := c.SettingsApi.SettingsDeleteDefaultPipeline(
		ctx,
		nextgen.IacmDeleteDefaultPipelineRequestBody{
			Provisioner: d.Get("provisioner_type").(string),
			Operation:   d.Get("operation").(string),
		},
		c.AccountId,
		d.Get("org_id").(string),
		d.Get("project_id").(string),
	)
	if err != nil {
		return parseError(err, httpResp)
	}

	return nil
}

func resourceIacmDefaultPipelineCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return resourceIacmDefaultPipelineUpdate(ctx, d, meta)
}

func resourceIacmDefaultPipelineUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)
	id := d.Id()
	if id == "" {
		d.MarkNewResource()
	}

	httpResp, err := c.SettingsApi.SettingsUpsertDefaultPipeline(
		ctx,
		nextgen.IacmUpsertDefaultPipelineRequestBody{
			Provisioner: d.Get("provisioner_type").(string),
			Operation:   d.Get("operation").(string),
			Pipeline:    d.Get("pipeline").(string),
		},
		c.AccountId,
		d.Get("org_id").(string),
		d.Get("project_id").(string),
	)
	if err != nil {
		return parseError(err, httpResp)
	}

	resourceIacmDefaultPipelineRead(ctx, d, meta)
	return nil
}

// iacm errors are in a different format from other harness services
// this function parses iacm errors and attempts to return them in way
// that is consistent with the provider.
func parseError(err error, httpResp *http.Response) diag.Diagnostics {
	// copied from helpers/errors.go
	if httpResp != nil && httpResp.StatusCode == 401 {
		return diag.Errorf(httpResp.Status + "\n" + "Hint:\n" +
			"1) Please check if token has expired or is wrong.\n" +
			"2) Harness Provider is misconfigured. For firstgen resources please give the correct api_key and for nextgen resources please give the correct platform_api_key.")
	}
	if httpResp != nil && httpResp.StatusCode == 403 {
		return diag.Errorf(httpResp.Status + "\n" + "Hint:\n" +
			"1) Please check if the token has required permission for this operation.\n" +
			"2) Please check if the token has expired or is wrong.")
	}

	se, ok := err.(nextgen.GenericSwaggerError)
	if !ok {
		diag.FromErr(err)
	}

	iacmErrBody := se.Body()
	iacmErr := nextgen.IacmError{}
	jsonErr := json.Unmarshal(iacmErrBody, &iacmErr)
	if jsonErr != nil {
		return diag.Errorf(err.Error())
	}

	return diag.Errorf(httpResp.Status + "\n" + "Hint:\n" +
		"1) " + iacmErr.Message)
}
