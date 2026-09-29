package trigger_pipeline

import (
	"context"
	"encoding/json"
	"sort"
	"time"

	"github.com/antihax/optional"
	pipeline_go_sdk "github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
)

// Non-terminal pipeline execution statuses; any status not in this set is treated as final.
var nonTerminalStatuses = map[string]bool{
	"Running": true, "AsyncWaiting": true, "TaskWaiting": true, "TimedWaiting": true,
	"NotStarted": true, "Discontinuing": true, "Queued": true, "Paused": true,
	"ResourceWaiting": true, "InterventionWaiting": true, "ApprovalWaiting": true,
	"WaitStepRunning": true, "QueuedLicenseLimitReached": true, "QueuedExecutionConcurrencyReached": true,
	"Pausing": true,
}

func ResourceTriggerPipeline() *schema.Resource {
	resource := &schema.Resource{
		Description: "Resource for triggering a Harness pipeline execution with optional runtime inputs. Any change to `org_id`, `project_id`, `pipeline_id`, `module_type`, `input_set_yaml`, or `branch` re-runs the pipeline (recreate). `wait_for_completion` and `poll_interval_seconds` can be changed in place.",

		CreateContext: resourceTriggerPipelineCreate,
		ReadContext:   resourceTriggerPipelineRead,
		UpdateContext: resourceTriggerPipelineUpdate,
		DeleteContext: resourceTriggerPipelineDelete,

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(2 * time.Hour),
			Update: schema.DefaultTimeout(2 * time.Hour),
		},

		Schema: map[string]*schema.Schema{
			"org_id": {
				Description: "Unique identifier of the organization.",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
			},
			"project_id": {
				Description: "Unique identifier of the project.",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
			},
			"pipeline_id": {
				Description: "Identifier of the pipeline to trigger.",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
			},
			"module_type": {
				Description: "Module type of the pipeline, e.g. cd, ci.",
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Default:     "cd",
			},
			"input_set_yaml": {
				Description: "Input Set YAML used to run the pipeline. Only needed if the pipeline has runtime inputs.",
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
			},
			"branch": {
				Description: "Name of the branch, for Git Experience pipelines.",
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
			},
			"wait_for_completion": {
				Description: "Whether to wait for the pipeline execution to reach a final status. When true, `status` and `outputs` reflect the completed execution.",
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
			},
			"poll_interval_seconds": {
				Description: "Interval, in seconds, between polls of the execution status when `wait_for_completion` is true.",
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     10,
			},
			"plan_execution_id": {
				Description: "Identifier of the triggered pipeline execution.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"status": {
				Description: "Status of the pipeline execution.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"outputs": {
				Description: "JSON-encoded map of output variables exported by the pipeline execution, keyed by step/stage identifier. Each value is a list of outcome objects (one per node execution; more than one when a matrix/parallel looping strategy runs the same identifier multiple times). Only populated once the execution has finished.",
				Type:        schema.TypeString,
				Computed:    true,
			},
		},
	}

	return resource
}

func resourceTriggerPipelineCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	orgId := d.Get("org_id").(string)
	projectId := d.Get("project_id").(string)
	pipelineId := d.Get("pipeline_id").(string)
	moduleType := d.Get("module_type").(string)

	opts := &pipeline_go_sdk.ExecuteApiPostPipelineExecuteWithInputSetYamlOpts{}
	if v, ok := d.GetOk("input_set_yaml"); ok {
		opts.Body = optional.NewInterface(v.(string))
	}
	if v, ok := d.GetOk("branch"); ok {
		opts.Branch = optional.NewString(v.(string))
	}

	resp, httpResp, err := c.ExecuteApi.PostPipelineExecuteWithInputSetYaml(ctx, c.AccountId, orgId, projectId, moduleType, pipelineId, opts)
	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}
	if resp.Data == nil || resp.Data.PlanExecution == nil {
		return diag.Errorf("triggering pipeline %s did not return a plan execution", pipelineId)
	}

	d.SetId(resp.Data.PlanExecution.Uuid)
	d.Set("plan_execution_id", resp.Data.PlanExecution.Uuid)
	d.Set("status", resp.Data.PlanExecution.Status)

	if d.Get("wait_for_completion").(bool) {
		if diags := waitAndPopulate(ctx, c, d, orgId, projectId, d.Id(), d.Timeout(schema.TimeoutCreate)); diags != nil {
			return diags
		}
	}

	return nil
}

func resourceTriggerPipelineRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	orgId := d.Get("org_id").(string)
	projectId := d.Get("project_id").(string)

	resp, httpResp, err := c.ExecutionDetailsApi.GetExecutionDetail(ctx, c.AccountId, orgId, projectId, d.Id(), nil)
	if httpResp != nil && httpResp.StatusCode == 404 {
		d.SetId("")
		return nil
	}
	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}
	if resp.Data == nil || resp.Data.PipelineExecutionSummary == nil {
		return diag.Errorf("no execution summary returned for pipeline execution %s", d.Id())
	}

	d.Set("status", resp.Data.PipelineExecutionSummary.Status)

	outputs, err := flattenOutcomes(resp.Data.ExecutionGraph)
	if err != nil {
		return diag.FromErr(err)
	}
	d.Set("outputs", outputs)

	return nil
}

func resourceTriggerPipelineUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	if !d.Get("wait_for_completion").(bool) {
		return resourceTriggerPipelineRead(ctx, d, meta)
	}

	orgId := d.Get("org_id").(string)
	projectId := d.Get("project_id").(string)

	return waitAndPopulate(ctx, c, d, orgId, projectId, d.Id(), d.Timeout(schema.TimeoutUpdate))
}

func resourceTriggerPipelineDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// Harness has no API to delete a pipeline execution; it is a historical record. Just drop it from state.
	return nil
}

// waitAndPopulate polls the execution status at a fixed interval until it reaches a terminal
// state or the timeout elapses, then stores the final status and outputs.
func waitAndPopulate(ctx context.Context, c *pipeline_go_sdk.APIClient, d *schema.ResourceData, orgId string, projectId string, planExecutionId string, timeout time.Duration) diag.Diagnostics {
	pollInterval := time.Duration(d.Get("poll_interval_seconds").(int)) * time.Second
	if pollInterval <= 0 {
		pollInterval = 10 * time.Second
	}
	deadline := time.Now().Add(timeout)

	for {
		resp, httpResp, err := c.ExecutionDetailsApi.GetExecutionDetail(ctx, c.AccountId, orgId, projectId, planExecutionId, nil)
		if err != nil {
			return helpers.HandleApiError(err, d, httpResp)
		}
		if resp.Data == nil || resp.Data.PipelineExecutionSummary == nil {
			return diag.Errorf("no execution summary returned for pipeline execution %s", planExecutionId)
		}

		status := resp.Data.PipelineExecutionSummary.Status
		if !nonTerminalStatuses[status] {
			d.Set("status", status)
			outputs, err := flattenOutcomes(resp.Data.ExecutionGraph)
			if err != nil {
				return diag.FromErr(err)
			}
			d.Set("outputs", outputs)
			return nil
		}
		d.Set("status", status)

		if time.Now().After(deadline) {
			return diag.Errorf("timed out waiting for pipeline execution %s to complete, last status was %s", planExecutionId, status)
		}

		select {
		case <-ctx.Done():
			return diag.FromErr(ctx.Err())
		case <-time.After(pollInterval):
		}
	}
}

func flattenOutcomes(graph *pipeline_go_sdk.ExecutionGraph) (string, error) {
	// NodeMap is keyed by node execution id, not by step/stage identifier: matrix/parallel
	// looping strategies expand a single identifier into multiple nodes, so group by
	// identifier into a list (sorted by node id for stable output) instead of overwriting.
	outcomes := map[string][]map[string]map[string]interface{}{}
	if graph != nil {
		nodeIds := make([]string, 0, len(graph.NodeMap))
		for nodeId := range graph.NodeMap {
			nodeIds = append(nodeIds, nodeId)
		}
		sort.Strings(nodeIds)

		for _, nodeId := range nodeIds {
			node := graph.NodeMap[nodeId]
			if len(node.Outcomes) == 0 {
				continue
			}
			outcomes[node.Identifier] = append(outcomes[node.Identifier], node.Outcomes)
		}
	}

	b, err := json.Marshal(outcomes)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
