package scan

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/harness/harness-go-sdk/harness/chaos"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var scanPollInterval = 5 * time.Second

func ResourceChaosPipelineScan() *schema.Resource {
	return &schema.Resource{
		Description: "Resource for scanning an existing Harness pipeline execution.\n\n" +
			"Set `pipeline_identity` and `pipeline_execution_id`. Create waits until the scan finishes. ERRORED and ABORTED fail the apply. " +
			"`detection_method` is `ai` or `ruleBased`. Later applies update name, description, and tags. " +
			"Change `retry_token` to run the same scan again, or `abort_token` to abort it.",

		CreateContext: resourcePipelineScanCreate,
		ReadContext: func(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
			return readScan(ctx, d, meta, d.Id(), chaos.ScanTypePipelineExecution, true)
		},
		UpdateContext: func(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
			return resourceChaosScanUpdate(ctx, d, meta, chaos.ScanTypePipelineExecution)
		},
		DeleteContext: resourceChaosScanDelete,
		CustomizeDiff: resourcePipelineScanCustomizeDiff,
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(2 * time.Hour),
			Read:   schema.DefaultTimeout(5 * time.Minute),
			Update: schema.DefaultTimeout(2 * time.Hour),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},
		Importer: &schema.ResourceImporter{
			StateContext: func(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
				return importChaosScan(ctx, d, meta, chaos.ScanTypePipelineExecution)
			},
		},

		Schema: resourceChaosPipelineScanSchema(),
	}
}

func ResourceChaosInfrastructureScan() *schema.Resource {
	return &schema.Resource{
		Description: "Resource for scanning an existing Kubernetes chaos infrastructure.\n\n" +
			"Set `infrastructure_id` and `environment_id`, or `agent_identity` and `environment_identity`. " +
			"Create waits until the scan finishes. ERRORED and ABORTED fail the apply. " +
			"`detection_method` is `ai` or `ruleBased`. Change `retry_token` to run the same scan again.",

		CreateContext: resourceInfrastructureScanCreate,
		ReadContext: func(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
			return readScan(ctx, d, meta, d.Id(), chaos.ScanTypeDiscoveryAgent, true)
		},
		UpdateContext: func(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
			return resourceChaosScanUpdate(ctx, d, meta, chaos.ScanTypeDiscoveryAgent)
		},
		DeleteContext: resourceChaosScanDelete,
		CustomizeDiff: resourceInfrastructureScanCustomizeDiff,
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(2 * time.Hour),
			Read:   schema.DefaultTimeout(5 * time.Minute),
			Update: schema.DefaultTimeout(2 * time.Hour),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},
		Importer: &schema.ResourceImporter{
			StateContext: func(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
				return importChaosScan(ctx, d, meta, chaos.ScanTypeDiscoveryAgent)
			},
		},

		Schema: resourceChaosInfrastructureScanSchema(),
	}
}

func resourcePipelineScanCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetChaosClientWithContext(ctx)
	client := chaos.NewScanClient(c)
	scope := scopeFrom(d, c.AccountId)
	req, err := pipelineScanRequest(d)
	if err != nil {
		return diag.FromErr(err)
	}
	log.Printf("[DEBUG] Creating chaos pipeline scan %s for pipeline %s execution %s", req.Identity, req.PipelineIdentity, req.PipelineExecutionID)
	created, httpResp, err := client.CreatePipelineScan(ctx, scope, req)
	if err != nil {
		return helpers.HandleChaosApiError(err, d, httpResp)
	}
	return finishScanCreate(ctx, d, meta, client, scope, created, chaos.ScanTypePipelineExecution)
}

func resourceInfrastructureScanCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetChaosClientWithContext(ctx)
	client := chaos.NewScanClient(c)
	scope := scopeFrom(d, c.AccountId)
	fromInfra, useInfra, err := infrastructureScanRequest(d)
	if err != nil {
		return diag.FromErr(err)
	}
	var created *chaos.Scan
	var httpResp *http.Response
	if useInfra {
		log.Printf("[DEBUG] Creating chaos infrastructure scan %s from infrastructure %s", fromInfra.Identity, fromInfra.InfrastructureID)
		created, httpResp, err = client.CreateInfrastructureScan(ctx, scope, *fromInfra)
	} else {
		req, reqErr := discoveryAgentScanRequest(d)
		if reqErr != nil {
			return diag.FromErr(reqErr)
		}
		log.Printf("[DEBUG] Creating chaos infrastructure scan %s for agent %s", req.Identity, req.Source.DiscoveryAgent.AgentIdentity)
		created, httpResp, err = client.Create(ctx, scope, req)
	}
	if err != nil {
		return helpers.HandleChaosApiError(err, d, httpResp)
	}
	return finishScanCreate(ctx, d, meta, client, scope, created, chaos.ScanTypeDiscoveryAgent)
}

func finishScanCreate(ctx context.Context, d *schema.ResourceData, meta interface{}, client *chaos.ScanClient, scope chaos.Scope, created *chaos.Scan, scanType string) diag.Diagnostics {
	if created == nil || created.Identity == "" {
		return diag.Errorf("create scan returned an empty identity")
	}
	d.SetId(created.Identity)
	waitDiags := waitForTerminalScan(ctx, false, func(ctx context.Context) (*chaos.Scan, *http.Response, error) {
		return client.Get(ctx, scope, created.Identity)
	})
	diags := readScan(ctx, d, meta, created.Identity, scanType, true)
	return append(diags, waitDiags...)
}

func resourceChaosScanUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}, scanType string) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetChaosClientWithContext(ctx)
	client := chaos.NewScanClient(c)
	scope := scopeFrom(d, c.AccountId)
	aborting := d.HasChange("abort_token") && d.Get("abort_token").(string) != ""
	retrying := d.HasChange("retry_token") && d.Get("retry_token").(string) != ""
	if aborting && retrying {
		return diag.Errorf("change abort_token or retry_token in one apply, not both")
	}

	if d.HasChanges("name", "description", "tags") {
		log.Printf("[DEBUG] Updating chaos scan %s", d.Id())
		httpResp, err := client.Update(ctx, scope, d.Id(), scanUpdateRequest(d))
		if err != nil {
			return helpers.HandleChaosApiError(err, d, httpResp)
		}
	}
	if aborting {
		log.Printf("[DEBUG] Aborting chaos scan %s", d.Id())
		httpResp, err := client.Abort(ctx, scope, d.Id())
		if err != nil {
			return helpers.HandleChaosApiError(err, d, httpResp)
		}
	}
	if retrying {
		log.Printf("[DEBUG] Retrying chaos scan %s", d.Id())
		httpResp, err := client.Retry(ctx, scope, d.Id())
		if err != nil {
			return helpers.HandleChaosApiError(err, d, httpResp)
		}
	}
	var waitDiags diag.Diagnostics
	if aborting || retrying {
		waitDiags = waitForTerminalScan(ctx, aborting, func(ctx context.Context) (*chaos.Scan, *http.Response, error) {
			return client.Get(ctx, scope, d.Id())
		})
	}
	return append(readScan(ctx, d, meta, d.Id(), scanType, true), waitDiags...)
}

func resourceChaosScanDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetChaosClientWithContext(ctx)
	log.Printf("[DEBUG] Deleting chaos scan %s", d.Id())
	httpResp, err := chaos.NewScanClient(c).Delete(ctx, scopeFrom(d, c.AccountId), d.Id())
	if err != nil {
		return helpers.HandleChaosReadApiError(err, d, httpResp)
	}
	d.SetId("")
	return nil
}

func importChaosScan(ctx context.Context, d *schema.ResourceData, meta interface{}, scanType string) ([]*schema.ResourceData, error) {
	orgID, projectID, identity, err := parseScanImportID(d.Id())
	if err != nil {
		return nil, err
	}
	d.SetId(identity)
	if err := d.Set("org_id", orgID); err != nil {
		return nil, err
	}
	if err := d.Set("project_id", projectID); err != nil {
		return nil, err
	}
	if err := d.Set("identity", identity); err != nil {
		return nil, err
	}
	c, ctx := meta.(*internal.Session).GetChaosClientWithContext(ctx)
	scan, httpResp, err := chaos.NewScanClient(c).Get(ctx, scopeFrom(d, c.AccountId), identity)
	if err != nil {
		if httpResp != nil {
			return nil, fmt.Errorf("read chaos scan %s: %s: %w", identity, httpResp.Status, err)
		}
		return nil, fmt.Errorf("read chaos scan %s: %w", identity, err)
	}
	if scan.ScanType != scanType {
		return nil, fmt.Errorf("scan %s is a %s scan, not a %s scan", identity, scan.ScanType, scanType)
	}
	diags := readScan(ctx, d, meta, identity, scanType, true)
	if diags.HasError() {
		return nil, fmt.Errorf("error reading chaos scan: %v", diags)
	}
	return []*schema.ResourceData{d}, nil
}

func readScan(ctx context.Context, d *schema.ResourceData, meta interface{}, identity, scanType string, keepInfrastructureIDs bool) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetChaosClientWithContext(ctx)
	client := chaos.NewScanClient(c)
	scope := scopeFrom(d, c.AccountId)

	log.Printf("[DEBUG] Reading chaos scan %s", identity)
	scan, httpResp, err := client.Get(ctx, scope, identity)
	if err != nil {
		return helpers.HandleChaosReadApiError(err, d, httpResp)
	}
	if scan.ScanType != scanType {
		return diag.Errorf("scan %s is a %s scan, not a %s scan", identity, scan.ScanType, scanType)
	}

	report, reportResp, reportErr := client.GetReport(ctx, scope, identity)
	var reportDiags diag.Diagnostics
	includeRisks := reportErr == nil
	if reportErr != nil {
		report = nil
		if reportResp == nil || reportResp.StatusCode != http.StatusNotFound {
			reportDiags = diag.Diagnostics{{
				Severity: diag.Warning,
				Summary:  "Chaos scan report was not read",
				Detail:   reportErr.Error() + ". The scan stays in state and the previous risks are kept.",
			}}
		}
	}

	if err := setScanState(d, scan, report, includeRisks, keepInfrastructureIDs); err != nil {
		return diag.FromErr(err)
	}
	d.SetId(scan.Identity)
	return reportDiags
}

func resourcePipelineScanCustomizeDiff(ctx context.Context, d *schema.ResourceDiff, meta interface{}) error {
	return forceNewWhenConfigured(d, "notify_id")
}

func resourceInfrastructureScanCustomizeDiff(ctx context.Context, d *schema.ResourceDiff, meta interface{}) error {
	for _, key := range []string{"agent_identity", "environment_identity", "environment_parent_unique_id"} {
		if err := forceNewWhenConfigured(d, key); err != nil {
			return err
		}
	}
	return nil
}

func forceNewWhenConfigured(d *schema.ResourceDiff, key string) error {
	// A value filled in by the API must not replace the scan.
	if !d.HasChange(key) {
		return nil
	}
	oldValue, newValue := d.GetChange(key)
	oldText, _ := oldValue.(string)
	newText, _ := newValue.(string)
	if oldText == "" || newText == "" || oldText == newText {
		return nil
	}
	return d.ForceNew(key)
}

func waitForTerminalScan(ctx context.Context, acceptAborted bool, get func(context.Context) (*chaos.Scan, *http.Response, error)) diag.Diagnostics {
	var lastStatus string
	for {
		scan, _, err := get(ctx)
		if err != nil {
			return diag.FromErr(err)
		}
		if scan == nil {
			return diag.Errorf("read chaos scan returned an empty body")
		}
		lastStatus = scan.Status
		switch scan.Status {
		case chaos.ScanStatusCompleted:
			return nil
		case chaos.ScanStatusAborted:
			if acceptAborted {
				return nil
			}
			fallthrough
		case chaos.ScanStatusErrored:
			message := scan.ErrorMessage
			if message == "" {
				message = "the scan ended without an error message"
			}
			return diag.Errorf("chaos scan %s is %s: %s", scan.Identity, scan.Status, message)
		}

		timer := time.NewTimer(scanPollInterval)
		select {
		case <-ctx.Done():
			timer.Stop()
			return diag.Errorf("timed out waiting for chaos scan %s to finish; last status %s", scan.Identity, lastStatus)
		case <-timer.C:
		}
	}
}
