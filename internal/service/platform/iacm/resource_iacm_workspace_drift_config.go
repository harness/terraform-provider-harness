package iacm

import (
	"context"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/internal"
	resourcebuilder "github.com/harness/terraform-provider-harness/internal/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var workspaceDriftMutex sync.Mutex

type workspaceDriftImplData struct {
	mu *sync.Mutex
}

type workspaceDriftState struct {
	Enabled      bool
	PipelineID   string
	RunEvery     int32
	RunEveryUnit string
	TriggerID    string
}

func ResourceIacmWorkspaceDriftConfig() *schema.Resource {
	implData := workspaceDriftImplData{mu: &workspaceDriftMutex}
	getSettings := func(_ context.Context, _ *schema.ResourceData, _ interface{}) (resourcebuilder.ResourceProviderImplData, error) {
		return implData, nil
	}

	resource, err := resourcebuilder.NewResourceProvider(resourcebuilder.ResourceProviderDefinition{
		Description:                 "Resource for managing IaCM workspace-level drift detection configuration. Enables scheduled drift detection for an individual workspace for a given provisioner.",
		Schema:                      workspaceDriftSchema(),
		GetResourceProviderImplData: getSettings,
		ApplyStateMappings:          applyWorkspaceDriftStateMappings,
		Import:                      newWorkspaceDriftImportOp(),
		Create:                      newWorkspaceDriftCreateOp(),
		Read:                        newWorkspaceDriftReadOp(),
		Update:                      newWorkspaceDriftUpdateOp(),
		Delete:                      newWorkspaceDriftDeleteOp(),
	})
	if err != nil {
		panic(fmt.Sprintf("invalid IaCM workspace drift resource: %v", err))
	}
	return resource
}

func newWorkspaceDriftCreateOp() resourcebuilder.CreateOpApiDefinition {
	return resourcebuilder.CreateOpApiDefinition{
		ConsistencyModel:   resourcebuilder.EventuallyConsistent,
		CreateApiCallerFn:  createWorkspaceDriftImpl,
		DesiredStateSource: resourcebuilder.SourceTerraformConfig,
		GetResourceStateFn: getWorkspaceDriftStateFromTerraformConfig,
		PollDefinition:     workspaceDriftPollDef(compareWorkspaceDriftCreateState),
		GetResourceID:      getWorkspaceDriftResourceID,
	}
}

func newWorkspaceDriftReadOp() resourcebuilder.ReadOpApiDefinition {
	return resourcebuilder.ReadOpApiDefinition{
		ReadApiCallerFn: readWorkspaceDriftImpl,
		HandleNotFound:  workspaceDriftHandleNotFound,
	}
}

func newWorkspaceDriftUpdateOp() resourcebuilder.UpdateOpApiDefinition {
	return resourcebuilder.UpdateOpApiDefinition{
		ConsistencyModel:   resourcebuilder.EventuallyConsistent,
		UpdateApiCallerFn:  createWorkspaceDriftImpl,
		DesiredStateSource: resourcebuilder.SourceTerraformConfig,
		GetResourceStateFn: getWorkspaceDriftStateFromTerraformConfig,
		PollDefinition:     workspaceDriftPollDef(compareWorkspaceDriftUpdateState),
	}
}

func newWorkspaceDriftDeleteOp() resourcebuilder.DeleteOpApiDefinition {
	return resourcebuilder.DeleteOpApiDefinition{
		ConsistencyModel:   resourcebuilder.EventuallyConsistent,
		DeleteApiCallerFn:  deleteWorkspaceDriftImpl,
		DesiredStateSource: resourcebuilder.SourceTerraformConfig,
		GetResourceStateFn: getWorkspaceDriftDeleteDesiredState,
		PollDefinition:     workspaceDriftPollDef(compareWorkspaceDriftDeleteState),
	}
}

func newWorkspaceDriftImportOp() resourcebuilder.ImportOpDefinition {
	return resourcebuilder.ImportOpDefinition{
		ImporterFn: workspaceDriftImporter,
	}
}

func workspaceDriftPollDef(comparator resourcebuilder.ResourceStateComparatorFn) *resourcebuilder.PollDefinition {
	return &resourcebuilder.PollDefinition{
		InitialWait:        30 * time.Millisecond,
		PollInterval:       500 * time.Millisecond,
		MaxAttempts:        36,
		EachPollTimeout:    1500 * time.Millisecond,
		MaxTimeout:         3 * time.Minute,
		PollFn:             readWorkspaceDriftImpl,
		ResourceComparator: comparator,
	}
}

func workspaceDriftSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"org_id": {
			Description: "Organization identifier where the project resides.",
			Type:        schema.TypeString,
			Required:    true,
			ForceNew:    true,
		},
		"project_id": {
			Description: "Project identifier where the workspace resides.",
			Type:        schema.TypeString,
			Required:    true,
			ForceNew:    true,
		},
		"workspace_id": {
			Description: "Workspace identifier to configure drift detection for.",
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
			Description: "Whether workspace-level drift automation is currently enabled.",
			Type:        schema.TypeBool,
			Computed:    true,
		},
		"trigger_id": {
			Description: "Generated trigger identifier assigned by the backend when drift is enabled.",
			Type:        schema.TypeString,
			Computed:    true,
		},
	}
}

func applyWorkspaceDriftStateMappings(opType resourcebuilder.ResourceProviderOpType, observedSource resourcebuilder.ObservedStateSource, observed resourcebuilder.ResourceState, d *schema.ResourceData) error {
	s, ok := observed.(workspaceDriftState)
	if !ok {
		return fmt.Errorf("applyWorkspaceDriftStateMappings: unexpected state type %T", observed)
	}
	for attr, val := range map[string]interface{}{
		"enabled":        s.Enabled,
		"pipeline_id":    s.PipelineID,
		"run_every":      s.RunEvery,
		"run_every_unit": s.RunEveryUnit,
		"trigger_id":     s.TriggerID,
	} {
		if err := d.Set(attr, val); err != nil {
			return fmt.Errorf("set %q: %w", attr, err)
		}
	}
	return nil
}

func workspaceDriftImporter(_ context.Context, d *schema.ResourceData, _ interface{}) ([]*schema.ResourceData, error) {
	parts := strings.SplitN(d.Id(), "/", 4)
	if len(parts) != 4 || parts[0] == "" || parts[1] == "" || parts[2] == "" || parts[3] == "" {
		return nil, fmt.Errorf("invalid import ID %q: expected org_id/project_id/workspace_id/provisioner_type", d.Id())
	}
	if err := d.Set("org_id", parts[0]); err != nil {
		return nil, err
	}
	if err := d.Set("project_id", parts[1]); err != nil {
		return nil, err
	}
	if err := d.Set("workspace_id", parts[2]); err != nil {
		return nil, err
	}
	if err := d.Set("provisioner_type", parts[3]); err != nil {
		return nil, err
	}
	return []*schema.ResourceData{d}, nil
}

func getWorkspaceDriftStateFromTerraformConfig(d *schema.ResourceData) (resourcebuilder.ResourceState, error) {
	return workspaceDriftState{
		PipelineID:   d.Get("pipeline_id").(string),
		RunEvery:     int32(d.Get("run_every").(int)),
		RunEveryUnit: d.Get("run_every_unit").(string),
	}, nil
}

func getWorkspaceDriftDeleteDesiredState(_ *schema.ResourceData) (resourcebuilder.ResourceState, error) {
	return workspaceDriftState{Enabled: false}, nil
}

func createWorkspaceDriftImpl(ctx context.Context, d *schema.ResourceData, meta interface{}) error {
	client, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)
	response, err := client.IacmWorkspaceDriftConfigApi.UpsertIacmWorkspaceDriftConfig(
		ctx, client.AccountId,
		d.Get("org_id").(string), d.Get("project_id").(string), d.Get("workspace_id").(string),
		nextgen.IacmWorkspaceDriftConfigRequestBody{
			PipelineId:   d.Get("pipeline_id").(string),
			Provisioner:  d.Get("provisioner_type").(string),
			RunEvery:     int32(d.Get("run_every").(int)),
			RunEveryUnit: d.Get("run_every_unit").(string),
		},
	)
	if response != nil && response.Body != nil {
		defer response.Body.Close()
	}
	return err
}

func deleteWorkspaceDriftImpl(ctx context.Context, d *schema.ResourceData, meta interface{}) error {
	client, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)
	response, err := client.IacmWorkspaceDriftConfigApi.DisableIacmWorkspaceDriftConfig(
		ctx, client.AccountId,
		d.Get("org_id").(string), d.Get("project_id").(string),
		d.Get("workspace_id").(string), d.Get("provisioner_type").(string),
	)
	if response != nil && response.Body != nil {
		defer response.Body.Close()
	}
	return err
}

// readWorkspaceDriftImpl fetches current state from the remote API.
// Also used as PollFn — no duplicate remote-read adapter needed.
func readWorkspaceDriftImpl(ctx context.Context, d *schema.ResourceData, meta interface{}) (resourcebuilder.ResourceState, error) {
	client, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)
	result, response, err := client.IacmWorkspaceDriftConfigApi.GetIacmWorkspaceDriftConfig(
		ctx, client.AccountId,
		d.Get("org_id").(string), d.Get("project_id").(string),
		d.Get("workspace_id").(string), d.Get("provisioner_type").(string),
	)
	if response != nil && response.Body != nil {
		defer response.Body.Close()
	}
	if err != nil {
		return nil, err
	}
	return workspaceDriftState{
		Enabled:      result.Enabled,
		PipelineID:   result.PipelineId,
		RunEvery:     result.RunEvery,
		RunEveryUnit: result.RunEveryUnit,
		TriggerID:    result.TriggerId,
	}, nil
}

func workspaceDriftHandleNotFound(observed resourcebuilder.ResourceState) bool {
	s, ok := observed.(workspaceDriftState)
	return ok && !s.Enabled
}

func compareWorkspaceDriftCreateState(expected, actual resourcebuilder.ResourceState) (bool, error) {
	want, ok := expected.(workspaceDriftState)
	if !ok {
		return false, fmt.Errorf("unexpected desired state %T", expected)
	}
	got, ok := actual.(workspaceDriftState)
	if !ok {
		return false, fmt.Errorf("unexpected observed state %T", actual)
	}
	return got.Enabled &&
		got.PipelineID == want.PipelineID &&
		got.RunEvery == want.RunEvery &&
		got.RunEveryUnit == want.RunEveryUnit, nil
}

func compareWorkspaceDriftUpdateState(expected, actual resourcebuilder.ResourceState) (bool, error) {
	return compareWorkspaceDriftCreateState(expected, actual)
}

func compareWorkspaceDriftDeleteState(expected, actual resourcebuilder.ResourceState) (bool, error) {
	_, ok := expected.(workspaceDriftState)
	if !ok {
		return false, fmt.Errorf("unexpected desired state %T", expected)
	}
	got, ok := actual.(workspaceDriftState)
	if !ok {
		return false, fmt.Errorf("unexpected observed state %T", actual)
	}
	return !got.Enabled, nil
}

func getWorkspaceDriftResourceID(d *schema.ResourceData) (string, error) {
	return strings.Join([]string{
		d.Get("org_id").(string), d.Get("project_id").(string),
		d.Get("workspace_id").(string), d.Get("provisioner_type").(string),
	}, "/"), nil
}
