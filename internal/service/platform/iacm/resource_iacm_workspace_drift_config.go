package iacm

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/internal"
	resourcebuilder "github.com/harness/terraform-provider-harness/internal/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type workspaceDriftImplData struct {
	DebugLevel string
}

type workspaceDriftState struct {
	Enabled      bool
	PipelineID   string
	RunEvery     int32
	RunEveryUnit string
	TriggerID    string
}

const (
	ResourceIacmWorkspaceDriftConfigDescription = "Resource for IaCM workspace-level drift detection configuration"
)

func ResourceIacmWorkspaceDriftConfig() *schema.Resource {

	resource, err := resourcebuilder.NewResourceProvider(resourcebuilder.ResourceProviderDefinition{
		Description:      ResourceIacmWorkspaceDriftConfigDescription,
		Schema:           workspaceDriftSchema(),
		ProviderImplData: workspaceDriftImplData{DebugLevel: "info"},
		Import:           newWorkspaceDriftImportOp(),
		Create:           newWorkspaceDriftCreateOp(),
		Read:             newWorkspaceDriftReadOp(),
		Update:           newWorkspaceDriftUpdateOp(),
		Delete:           newWorkspaceDriftDeleteOp(),
	})
	if err != nil {
		panic(fmt.Sprintf("invalid IaCM workspace drift resource: %v", err))
	}
	return resource
}

func newWorkspaceDriftCreateOp() resourcebuilder.CreateOpApiDefinition {
	return resourcebuilder.CreateOpApiDefinition{
		ResourceOpDefinition: resourcebuilder.ResourceOpDefinition{
			ApiOpModel:  resourcebuilder.AsyncApiOp,
			ApiCallerFn: createWorkspaceDriftImpl,
			DesiredState: resourcebuilder.DesiredResourceState{
				Source:              resourcebuilder.SourceTerraformConfig,
				ToStateFromSourceFn: getWorkspaceDriftStateFromTerraformConfigImpl,
			},
			PollDefinition:     workspaceDriftPollDef(compareWorkspaceDriftCreateStateImpl),
			ApplyStateMappings: applyWorkspaceDriftCreateMappingsImpl,
		},
		GetResourceID: getWorkspaceDriftResourceIDImpl,
	}
}

func newWorkspaceDriftReadOp() resourcebuilder.ReadOpApiDefinition {
	return resourcebuilder.ReadOpApiDefinition{
		ReadApiCallerFn:    readWorkspaceDriftImpl,
		HandleNotFoundFn:   workspaceDriftHandleNotFoundImpl,
		ApplyStateMappings: applyWorkspaceDriftReadMappingsImpl,
	}
}

func newWorkspaceDriftUpdateOp() resourcebuilder.UpdateOpApiDefinition {
	return resourcebuilder.UpdateOpApiDefinition{
		ResourceOpDefinition: resourcebuilder.ResourceOpDefinition{
			ApiOpModel:  resourcebuilder.AsyncApiOp,
			ApiCallerFn: createWorkspaceDriftImpl,
			DesiredState: resourcebuilder.DesiredResourceState{
				Source:              resourcebuilder.SourceTerraformConfig,
				ToStateFromSourceFn: getWorkspaceDriftStateFromTerraformConfigImpl,
			},
			PollDefinition:     workspaceDriftPollDef(compareWorkspaceDriftUpdateStateImpl),
			ApplyStateMappings: applyWorkspaceDriftUpdateMappingsImpl,
		},
	}
}

func newWorkspaceDriftDeleteOp() resourcebuilder.DeleteOpApiDefinition {
	return resourcebuilder.DeleteOpApiDefinition{
		ResourceOpDefinition: resourcebuilder.ResourceOpDefinition{
			ApiOpModel:  resourcebuilder.AsyncApiOp,
			ApiCallerFn: deleteWorkspaceDriftImpl,
			DesiredState: resourcebuilder.DesiredResourceState{
				Source:              resourcebuilder.SourceTerraformConfig,
				ToStateFromSourceFn: getWorkspaceDriftDeleteDesiredStateImpl,
			},
			PollDefinition:     workspaceDriftPollDef(compareWorkspaceDriftDeleteStateImpl),
			ApplyStateMappings: applyWorkspaceDriftDeleteMappingsImpl,
		},
	}
}

func newWorkspaceDriftImportOp() resourcebuilder.ImportOpDefinition {
	return resourcebuilder.ImportOpDefinition{
		ImporterFn: workspaceDriftImporterImpl,
	}
}

func workspaceDriftPollDef(comparator resourcebuilder.ResourceStateComparatorFn) *resourcebuilder.PollDefinition {
	return &resourcebuilder.PollDefinition{
		InitialWait:          30 * time.Millisecond,
		PollInterval:         500 * time.Millisecond,
		MaxAttempts:          36,
		EachPollTimeout:      1500 * time.Millisecond,
		MaxTimeout:           3 * time.Minute,
		PollFn:               readWorkspaceDriftImpl,
		ResourceComparatorFn: comparator,
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

func applyWorkspaceDriftCreateMappingsImpl(observed resourcebuilder.ResourceState, d *schema.ResourceData) error {
	s, ok := observed.(workspaceDriftState)
	if !ok {
		return fmt.Errorf("applyWorkspaceDriftCreateMappingsImpl: unexpected state type %T", observed)
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

func applyWorkspaceDriftReadMappingsImpl(observed resourcebuilder.ResourceState, d *schema.ResourceData) error {
	return applyWorkspaceDriftCreateMappingsImpl(observed, d)
}

func applyWorkspaceDriftUpdateMappingsImpl(observed resourcebuilder.ResourceState, d *schema.ResourceData) error {
	return applyWorkspaceDriftCreateMappingsImpl(observed, d)
}

func applyWorkspaceDriftDeleteMappingsImpl(observed resourcebuilder.ResourceState, d *schema.ResourceData) error {
	s, ok := observed.(workspaceDriftState)
	if !ok {
		return fmt.Errorf("applyWorkspaceDriftDeleteMappingsImpl: unexpected state type %T", observed)
	}
	return d.Set("enabled", s.Enabled)
}

func workspaceDriftImporterImpl(_ context.Context, d *schema.ResourceData, _ interface{}) ([]*schema.ResourceData, error) {
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

func getWorkspaceDriftStateFromTerraformConfigImpl(d *schema.ResourceData) (resourcebuilder.ResourceState, error) {
	return workspaceDriftState{
		PipelineID:   d.Get("pipeline_id").(string),
		RunEvery:     int32(d.Get("run_every").(int)),
		RunEveryUnit: d.Get("run_every_unit").(string),
	}, nil
}

func getWorkspaceDriftDeleteDesiredStateImpl(_ *schema.ResourceData) (resourcebuilder.ResourceState, error) {
	return workspaceDriftState{Enabled: false}, nil
}

func unmarshalIacmErrMsg(err error) error {
	se, ok := err.(nextgen.GenericSwaggerError)
	if !ok {
		return err
	}
	iacmErr := nextgen.IacmError{}
	if jsonErr := json.Unmarshal(se.Body(), &iacmErr); jsonErr == nil && iacmErr.Message != "" {
		return fmt.Errorf("%s", iacmErr.Message)
	}
	return err
}

func createWorkspaceDriftImpl(ctx context.Context, d *schema.ResourceData, meta interface{}, implData resourcebuilder.ResourceProviderImplData) error {
	impl := implData.(workspaceDriftImplData)
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
	if err != nil {
		err = unmarshalIacmErrMsg(err)
		log.Printf("[%s] createWorkspaceDriftImpl: error: %v", impl.DebugLevel, err)
	}
	return err
}

func deleteWorkspaceDriftImpl(ctx context.Context, d *schema.ResourceData, meta interface{}, implData resourcebuilder.ResourceProviderImplData) error {
	impl := implData.(workspaceDriftImplData)
	client, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)
	response, err := client.IacmWorkspaceDriftConfigApi.DisableIacmWorkspaceDriftConfig(
		ctx, client.AccountId,
		d.Get("org_id").(string), d.Get("project_id").(string),
		d.Get("workspace_id").(string), d.Get("provisioner_type").(string),
	)
	if response != nil && response.Body != nil {
		defer response.Body.Close()
	}
	if err != nil {
		err = unmarshalIacmErrMsg(err)
		log.Printf("[%s] deleteWorkspaceDriftImpl: error: %v", impl.DebugLevel, err)
	}
	return err
}

func readWorkspaceDriftImpl(ctx context.Context, d *schema.ResourceData, meta interface{}, implData resourcebuilder.ResourceProviderImplData) (resourcebuilder.ResourceState, error) {
	impl := implData.(workspaceDriftImplData)
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
		err = unmarshalIacmErrMsg(err)
		log.Printf("[%s] readWorkspaceDriftImpl: error: %v", impl.DebugLevel, err)
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

func workspaceDriftHandleNotFoundImpl(observed resourcebuilder.ResourceState) bool {
	s, ok := observed.(workspaceDriftState)
	return ok && !s.Enabled
}

func compareWorkspaceDriftCreateStateImpl(expected, actual resourcebuilder.ResourceState) (bool, error) {
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

func compareWorkspaceDriftUpdateStateImpl(expected, actual resourcebuilder.ResourceState) (bool, error) {
	return compareWorkspaceDriftCreateStateImpl(expected, actual)
}

func compareWorkspaceDriftDeleteStateImpl(expected, actual resourcebuilder.ResourceState) (bool, error) {
	want, ok := expected.(workspaceDriftState)
	if !ok {
		return false, fmt.Errorf("unexpected desired state %T", expected)
	}
	got, ok := actual.(workspaceDriftState)
	if !ok {
		return false, fmt.Errorf("unexpected observed state %T", actual)
	}
	return got.Enabled == want.Enabled, nil
}

func getWorkspaceDriftResourceIDImpl(d *schema.ResourceData) (string, error) {
	return strings.Join([]string{
		d.Get("org_id").(string), d.Get("project_id").(string),
		d.Get("workspace_id").(string), d.Get("provisioner_type").(string),
	}, "/"), nil
}
