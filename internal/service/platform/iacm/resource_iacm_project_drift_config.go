package iacm

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	iacmProjectDriftConfigTimeout      = 30 * time.Second
	iacmProjectDriftConfigPollInterval = 9 * time.Second
)

func ResourceIacmProjectDriftConfig() *schema.Resource {
	return &schema.Resource{
		Description: "Resource for managing IaCM project-level drift detection configuration. " +
			"Enables scheduled drift detection across all workspaces in a project for a given provisioner.",

		ReadContext:   resourceIacmProjectDriftConfigRead,
		CreateContext: resourceIacmProjectDriftConfigCreate,
		UpdateContext: resourceIacmProjectDriftConfigUpdate,
		DeleteContext: resourceIacmProjectDriftConfigDelete,

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(iacmProjectDriftConfigTimeout),
			Update: schema.DefaultTimeout(iacmProjectDriftConfigTimeout),
			Delete: schema.DefaultTimeout(iacmProjectDriftConfigTimeout),
		},

		Importer: &schema.ResourceImporter{
			StateContext: func(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
				// ID format: org_id/project_id/provisioner_type
				parts := strings.SplitN(d.Id(), "/", 3)
				if len(parts) != 3 || parts[0] == "" || parts[1] == "" || parts[2] == "" {
					return nil, fmt.Errorf("invalid import ID %q: expected org_id/project_id/provisioner_type", d.Id())
				}
				if err := d.Set("org_id", parts[0]); err != nil {
					return nil, err
				}
				if err := d.Set("project_id", parts[1]); err != nil {
					return nil, err
				}
				if err := d.Set("provisioner_type", parts[2]); err != nil {
					return nil, err
				}
				return []*schema.ResourceData{d}, nil
			},
		},

		Schema: map[string]*schema.Schema{
			"org_id": {
				Description: "Organization identifier where the project resides.",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
			},
			"project_id": {
				Description: "Project identifier to configure drift detection for.",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
			},
			"provisioner_type": {
				Description: "Provisioner type to enable drift detection for (terraform, opentofu, terragrunt, awscdk).",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
			},
			"pipeline_id": {
				Description: "Pipeline to run for drift detection.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"run_every": {
				Description: "Repeat interval — run drift detection every N units (see run_every_unit).",
				Type:        schema.TypeInt,
				Required:    true,
			},
			"run_every_unit": {
				Description: "Unit for the run_every interval. Valid values: seconds (1–59), minutes (1–59), hours (1–24), days (1–31).",
				Type:        schema.TypeString,
				Required:    true,
			},
			"enabled": {
				Description: "Whether project-level drift automation is currently enabled.",
				Type:        schema.TypeBool,
				Computed:    true,
			},
		},
	}
}

func resourceIacmProjectDriftConfigRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	org := d.Get("org_id").(string)
	project := d.Get("project_id").(string)
	provisioner := d.Get("provisioner_type").(string)

	log.Printf("[DEBUG] iacm drift read: org=%q project=%q provisioner=%q", org, project, provisioner)

	result, httpResp, err := c.IacmProjectDriftConfigApi.GetIacmProjectDriftConfig(
		ctx,
		c.AccountId,
		org,
		project,
		provisioner,
	)
	statusCode := 0
	if httpResp != nil {
		statusCode = httpResp.StatusCode
	}
	log.Printf("[DEBUG] iacm drift read result: status=%d err=%v enabled=%v pipeline=%q run_every=%d unit=%q",
		statusCode, err, result.Enabled, result.PipelineId, result.RunEvery, result.RunEveryUnit)

	if err != nil {
		if httpResp != nil && httpResp.StatusCode == http.StatusNotFound {
			log.Printf("[DEBUG] iacm drift read: not found (404), clearing state for org=%q project=%q provisioner=%q", org, project, provisioner)
			d.SetId("")
			return nil
		}
		return parseError(err, httpResp)
	}

	if !result.Enabled {
		log.Printf("[DEBUG] iacm drift read: enabled=false, clearing state for org=%q project=%q provisioner=%q", org, project, provisioner)
		d.SetId("")
		return nil
	}

	if err := d.Set("enabled", result.Enabled); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("pipeline_id", result.PipelineId); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("run_every", int(result.RunEvery)); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("run_every_unit", result.RunEveryUnit); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceIacmProjectDriftConfigCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return resourceIacmProjectDriftConfigUpsert(ctx, d, meta, d.Timeout(schema.TimeoutCreate))
}

func resourceIacmProjectDriftConfigUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return resourceIacmProjectDriftConfigUpsert(ctx, d, meta, d.Timeout(schema.TimeoutUpdate))
}

func resourceIacmProjectDriftConfigUpsert(ctx context.Context, d *schema.ResourceData, meta interface{}, timeout time.Duration) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	org := d.Get("org_id").(string)
	project := d.Get("project_id").(string)
	provisioner := d.Get("provisioner_type").(string)
	pipelineID := d.Get("pipeline_id").(string)
	runEvery := int32(d.Get("run_every").(int))
	runEveryUnit := d.Get("run_every_unit").(string)

	log.Printf("[DEBUG] iacm drift upsert: org=%q project=%q provisioner=%q pipeline=%q run_every=%d unit=%q",
		org, project, provisioner, pipelineID, runEvery, runEveryUnit)

	// Bound total operation time (PUT + polling) by the configured Terraform timeout,
	// while preserving cancellation from the parent context.
	pollCtx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	httpResp, err := c.IacmProjectDriftConfigApi.UpsertIacmProjectDriftConfig(
		pollCtx,
		c.AccountId,
		org,
		project,
		nextgen.IacmProjectDriftConfigRequestBody{
			PipelineId:   pipelineID,
			Provisioner:  provisioner,
			RunEvery:     runEvery,
			RunEveryUnit: runEveryUnit,
		},
	)
	if err != nil {
		return parseError(err, httpResp)
	}
	httpResp.Body.Close()

	// Set ID immediately so partial-create can be recovered on a later apply.
	d.SetId(org + "/" + project + "/" + provisioner)

	// PUT is async (202 Accepted) — poll until all desired fields are reflected.
	if err := waitForIacmProjectDriftEnabled(pollCtx, c, org, project, provisioner, pipelineID, runEvery, runEveryUnit); err != nil {
		return diag.Errorf("drift config request accepted but remote convergence timed out: %s", err)
	}

	log.Printf("[DEBUG] iacm drift upsert: convergence confirmed for org=%q project=%q provisioner=%q", org, project, provisioner)
	return resourceIacmProjectDriftConfigRead(ctx, d, meta)
}

func resourceIacmProjectDriftConfigDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	org := d.Get("org_id").(string)
	project := d.Get("project_id").(string)
	provisioner := d.Get("provisioner_type").(string)

	log.Printf("[DEBUG] iacm drift delete: org=%q project=%q provisioner=%q", org, project, provisioner)

	// Bound total operation time (DELETE + polling) by the configured Terraform timeout,
	// while preserving cancellation from the parent context.
	pollCtx, cancel := context.WithTimeout(ctx, d.Timeout(schema.TimeoutDelete))
	defer cancel()

	httpResp, err := c.IacmProjectDriftConfigApi.DisableIacmProjectDriftConfig(
		pollCtx,
		c.AccountId,
		org,
		project,
		provisioner,
	)
	if err != nil {
		return parseError(err, httpResp)
	}
	httpResp.Body.Close()

	// DELETE is async (202 Accepted) — poll until enabled=false or 404 is confirmed.
	for {
		result, getResp, getErr := c.IacmProjectDriftConfigApi.GetIacmProjectDriftConfig(
			pollCtx, c.AccountId, org, project, provisioner,
		)
		statusCode := 0
		if getResp != nil {
			statusCode = getResp.StatusCode
		}
		log.Printf("[DEBUG] iacm drift delete poller: org=%q project=%q provisioner=%q status=%d err=%v enabled=%v",
			org, project, provisioner, statusCode, getErr, result.Enabled)
		if getErr != nil {
			if getResp != nil && getResp.StatusCode == http.StatusNotFound {
				log.Printf("[DEBUG] iacm drift delete: confirmed deleted (404) for org=%q project=%q provisioner=%q", org, project, provisioner)
				d.SetId("")
				return nil
			}
			if getResp != nil && (getResp.StatusCode == http.StatusUnauthorized || getResp.StatusCode == http.StatusForbidden) {
				return diag.Errorf("permanent error polling drift disable (status %d): %s", getResp.StatusCode, getErr)
			}
		} else if getResp != nil && getResp.StatusCode == http.StatusOK && !result.Enabled {
			log.Printf("[DEBUG] iacm drift delete: confirmed disabled (enabled=false) for org=%q project=%q provisioner=%q", org, project, provisioner)
			d.SetId("")
			return nil
		}
		select {
		case <-pollCtx.Done():
			return diag.Errorf("drift config disable request accepted but enabled=false not confirmed within timeout")
		case <-time.After(iacmProjectDriftConfigPollInterval):
		}
	}
}

// waitForIacmProjectDriftEnabled polls GET until all desired fields match or ctx is done.
// Fails immediately on permanent errors (401/403); retries transient failures.
// The caller is responsible for setting an appropriate deadline on ctx.
func waitForIacmProjectDriftEnabled(
	ctx context.Context,
	c *nextgen.APIClient,
	org, project, provisioner, pipelineID string,
	runEvery int32,
	runEveryUnit string,
) error {
	for {
		result, httpResp, err := c.IacmProjectDriftConfigApi.GetIacmProjectDriftConfig(
			ctx, c.AccountId, org, project, provisioner,
		)
		statusCode := 0
		if httpResp != nil {
			statusCode = httpResp.StatusCode
		}
		log.Printf("[DEBUG] iacm drift waiter: status=%d err=%v enabled=%v pipeline=%q run_every=%d unit=%q",
			statusCode, err, result.Enabled, result.PipelineId, result.RunEvery, result.RunEveryUnit)
		if err != nil {
			if statusCode == http.StatusUnauthorized || statusCode == http.StatusForbidden {
				return fmt.Errorf("permanent error polling drift enable (status %d): %w", statusCode, err)
			}
			// transient error — retry
		} else if statusCode == http.StatusOK &&
			result.Enabled &&
			result.PipelineId == pipelineID &&
			result.RunEvery == runEvery &&
			result.RunEveryUnit == runEveryUnit {
			return nil
		}
		select {
		case <-ctx.Done():
			return fmt.Errorf("desired state not observed: %w", ctx.Err())
		case <-time.After(iacmProjectDriftConfigPollInterval):
		}
	}
}
