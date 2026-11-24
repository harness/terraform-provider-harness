package utils

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// ResourceState is opaque: the resource implementation owns its shape.
type ResourceState any

// ApiOpModel describes how the remote API behaves after a mutating call.
type ApiOpModel string

const (
	// AsyncApiOp means the API returns 2xx but observed state may lag;
	// the framework polls until PollDefinition.ResourceComparator returns true.
	AsyncApiOp ApiOpModel = "async"
	// SyncApiOp means the API response already reflects final state;
	// no polling is performed after the mutation.
	SyncApiOp ApiOpModel = "sync"
)

// SourceOfDesiredState identifies where the desired state values are read from.
type SourceOfDesiredState string

const (
	// SourceTerraformConfig captures field values from the evaluated HCL config
	// (d.Get), not from prior state or remote.
	SourceTerraformConfig SourceOfDesiredState = "terraform_config"
)

// DesiredResourceState bundles the source label with the function that reads the desired state.
// Used by async ops to snapshot what the user wants before the API call, so the poller can
// compare it against observed state on each attempt.
type DesiredResourceState struct {
	// Source identifies where the desired state values are read from.
	Source SourceOfDesiredState
	// ToStateFromSourceFn reads the desired state from the identified source.
	ToStateFromSourceFn GetResourceStateFn
}

// ResourceProviderImplData is opaque shared data passed to every callback in the resource.
// The resource implementation owns its shape and interprets it as needed.
type ResourceProviderImplData any

// ApplyStateMappingsFn is called by the framework after an operation produces observed state.
// The resource implementation writes observed state into Terraform schema state via d.Set.
type ApplyStateMappingsFn func(observed ResourceState, d *schema.ResourceData) error

type (
	// ApiCallerFn performs a mutating operation (create, update, or delete).
	ApiCallerFn func(context.Context, *schema.ResourceData, interface{}) error
	// ReadApiCallerFn fetches the current observed state from the remote API.
	ReadApiCallerFn func(context.Context, *schema.ResourceData, interface{}) (ResourceState, error)

	// GetResourceProviderImplDataFn is the single shared fn evaluated for every callback.
	GetResourceProviderImplDataFn func(context.Context, *schema.ResourceData, interface{}) (ResourceProviderImplData, error)

	// GetResourceStateFn captures the desired state from Terraform config (d.Get only, no remote calls).
	GetResourceStateFn func(*schema.ResourceData) (ResourceState, error)

	// ResourceStateComparatorFn compares desired and observed state; returns true when they have converged.
	ResourceStateComparatorFn func(desired, observed ResourceState) (bool, error)
	// ResourceIDFn derives the Terraform resource ID from schema state after a successful create.
	ResourceIDFn func(*schema.ResourceData) (string, error)
)

// PollDefinition controls convergence polling behaviour.
type PollDefinition struct {
	// InitialWait is how long to wait before the first poll attempt.
	InitialWait time.Duration
	// PollInterval is how long to wait between consecutive poll attempts.
	PollInterval time.Duration
	// MaxAttempts caps the number of poll attempts; 0 means unlimited (use MaxTimeout instead).
	MaxAttempts int
	// EachPollTimeout is the per-attempt context deadline; 0 means no per-attempt limit.
	EachPollTimeout time.Duration
	// MaxTimeout is the overall deadline for convergence; 0 means no overall limit.
	MaxTimeout time.Duration
	// PollFn is the remote API caller invoked on each poll attempt to fetch current observed state.
	PollFn ReadApiCallerFn
	// ResourceComparator compares desired state against the latest observed state;
	// returns true when the resource has converged.
	ResourceComparator ResourceStateComparatorFn
}

// ResourceOpDefinition holds fields common to all mutating operations (Create, Update, Delete).
type ResourceOpDefinition struct {
	// ApiOpModel controls whether the API is async (poll until converged) or sync (no polling).
	ApiOpModel ApiOpModel
	// ApiCallerFn is the caller that performs the mutation — wraps any external call
	// such as a remote API, filesystem write, or secret generation.
	ApiCallerFn ApiCallerFn
	// DesiredState captures where the desired state comes from and how to read it.
	// Required when ApiOpModel is AsyncApiOp.
	DesiredState DesiredResourceState
	// PollDefinition controls convergence polling after the mutating call.
	// Required when ApiOpModel is AsyncApiOp.
	PollDefinition *PollDefinition
	// ApplyStateMappings writes the observed post-op state into Terraform schema state.
	// Called after convergence for AsyncApiOp; optional for SyncApiOp (no observed state produced).
	ApplyStateMappings ApplyStateMappingsFn
}

// CreateOpApiDefinition describes the Create operation.
type CreateOpApiDefinition struct {
	ResourceOpDefinition
	// GetResourceID derives the Terraform resource ID from schema state after a successful create.
	GetResourceID ResourceIDFn
}

// ReadOpApiDefinition describes the Read operation.
type ReadOpApiDefinition struct {
	ReadApiCallerFn ReadApiCallerFn
	// HandleNotFound is called when the remote read returns a not-found signal.
	// Return true to clear the resource ID (resource gone); false to leave state.
	HandleNotFound func(ResourceState) bool
	// ApplyStateMappings writes the observed state into Terraform schema state.
	ApplyStateMappings ApplyStateMappingsFn
}

// UpdateOpApiDefinition describes the Update operation.
type UpdateOpApiDefinition struct {
	ResourceOpDefinition
}

// DeleteOpApiDefinition describes the Delete operation.
type DeleteOpApiDefinition struct {
	ResourceOpDefinition
}

// ImporterFn parses an import ID from d.Id(), sets schema fields, and returns
// the populated ResourceData slice. It does not call a remote API.
type ImporterFn func(context.Context, *schema.ResourceData, interface{}) ([]*schema.ResourceData, error)

// ImportOpDefinition describes the Import operation.
type ImportOpDefinition struct {
	ImporterFn ImporterFn
}

func (op ImportOpDefinition) toResourceImporter() *schema.ResourceImporter {
	if op.ImporterFn == nil {
		return nil
	}
	return &schema.ResourceImporter{StateContext: schema.StateContextFunc(op.ImporterFn)}
}

// ResourceProviderDefinition groups schema, the shared impl data, and all five op definitions.
type ResourceProviderDefinition struct {
	Description      string
	Schema           map[string]*schema.Schema
	ProviderImplData ResourceProviderImplData
	Create           CreateOpApiDefinition
	Read             ReadOpApiDefinition
	Update           UpdateOpApiDefinition
	Delete           DeleteOpApiDefinition
	Import           ImportOpDefinition
}

// NewResourceProvider builds a *schema.Resource with all four CRUD context functions.
func NewResourceProvider(def ResourceProviderDefinition) (*schema.Resource, error) {
	if err := checkResourceProviderDefinition(def); err != nil {
		return nil, err
	}
	implData := def.ProviderImplData
	get := func(_ context.Context, _ *schema.ResourceData, _ interface{}) (ResourceProviderImplData, error) {
		return implData, nil
	}
	r := &schema.Resource{
		Description: def.Description,
		Schema:      def.Schema,
		Importer:    def.Import.toResourceImporter(),
	}
	setResourceProviderHandlers(r, def, get)
	return r, nil
}

func setResourceProviderHandlers(r *schema.Resource, def ResourceProviderDefinition, get GetResourceProviderImplDataFn) {
	if def.Create.ApiCallerFn != nil {
		r.CreateContext = setCreateOpHandler(def.Create, get)
	} else {
		r.CreateContext = notImplementedHandler("create")
	}
	if def.Read.ReadApiCallerFn != nil {
		r.ReadContext = setReadOpHandler(def.Read, get)
	} else {
		r.ReadContext = notImplementedHandler("read")
	}
	if def.Update.ApiCallerFn != nil {
		r.UpdateContext = setUpdateOpHandler(def.Update, get)
	} else {
		r.UpdateContext = notImplementedHandler("update")
	}
	if def.Delete.ApiCallerFn != nil {
		r.DeleteContext = setDeleteOpHandler(def.Delete, get)
	} else {
		r.DeleteContext = notImplementedHandler("delete")
	}
}

func notImplementedHandler(op string) func(context.Context, *schema.ResourceData, interface{}) diag.Diagnostics {
	return func(_ context.Context, _ *schema.ResourceData, _ interface{}) diag.Diagnostics {
		return diag.Errorf("%s operation is not implemented for this resource", op)
	}
}

func setCreateOpHandler(op CreateOpApiDefinition, get GetResourceProviderImplDataFn) schema.CreateContextFunc {
	return func(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
		var desired ResourceState
		if op.ApiOpModel == AsyncApiOp {
			var err error
			desired, err = op.DesiredState.ToStateFromSourceFn(d)
			if err != nil {
				return diag.FromErr(fmt.Errorf("capture desired state: %w", err))
			}
		}
		_, err := withSettings(ctx, d, meta, get, func(c context.Context) (struct{}, error) {
			return struct{}{}, op.ApiCallerFn(c, d, meta)
		})
		if err != nil {
			return diag.FromErr(fmt.Errorf("create remote resource: %w", err))
		}
		id, err := op.GetResourceID(d)
		if err != nil {
			return diag.FromErr(err)
		}
		d.SetId(id)
		if op.ApiOpModel == AsyncApiOp {
			observed, err := pollUntilConverged(ctx, d, meta, desired, *op.PollDefinition, get)
			if err != nil {
				return diag.FromErr(err)
			}
			if op.ApplyStateMappings != nil {
				if err := op.ApplyStateMappings(observed, d); err != nil {
					return diag.FromErr(err)
				}
			}
		}
		return nil
	}
}

func setReadOpHandler(op ReadOpApiDefinition, get GetResourceProviderImplDataFn) schema.ReadContextFunc {
	return func(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
		observed, err := withSettings(ctx, d, meta, get, func(c context.Context) (ResourceState, error) {
			return op.ReadApiCallerFn(c, d, meta)
		})
		if err != nil {
			return diag.FromErr(fmt.Errorf("read remote resource: %w", err))
		}
		if op.HandleNotFound != nil && op.HandleNotFound(observed) {
			d.SetId("")
			return nil
		}
		if op.ApplyStateMappings != nil {
			if err := op.ApplyStateMappings(observed, d); err != nil {
				return diag.FromErr(err)
			}
		}
		return nil
	}
}

func setUpdateOpHandler(op UpdateOpApiDefinition, get GetResourceProviderImplDataFn) schema.UpdateContextFunc {
	return func(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
		var desired ResourceState
		if op.ApiOpModel == AsyncApiOp {
			var err error
			desired, err = op.DesiredState.ToStateFromSourceFn(d)
			if err != nil {
				return diag.FromErr(fmt.Errorf("capture desired state: %w", err))
			}
		}
		_, err := withSettings(ctx, d, meta, get, func(c context.Context) (struct{}, error) {
			return struct{}{}, op.ApiCallerFn(c, d, meta)
		})
		if err != nil {
			return diag.FromErr(fmt.Errorf("update remote resource: %w", err))
		}
		if op.ApiOpModel == AsyncApiOp {
			observed, err := pollUntilConverged(ctx, d, meta, desired, *op.PollDefinition, get)
			if err != nil {
				return diag.FromErr(err)
			}
			if op.ApplyStateMappings != nil {
				if err := op.ApplyStateMappings(observed, d); err != nil {
					return diag.FromErr(err)
				}
			}
		}
		return nil
	}
}

func setDeleteOpHandler(op DeleteOpApiDefinition, get GetResourceProviderImplDataFn) schema.DeleteContextFunc {
	return func(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
		var desired ResourceState
		if op.ApiOpModel == AsyncApiOp {
			var err error
			desired, err = op.DesiredState.ToStateFromSourceFn(d)
			if err != nil {
				return diag.FromErr(fmt.Errorf("capture desired state: %w", err))
			}
		}
		_, err := withSettings(ctx, d, meta, get, func(c context.Context) (struct{}, error) {
			return struct{}{}, op.ApiCallerFn(c, d, meta)
		})
		if err != nil {
			return diag.FromErr(fmt.Errorf("delete remote resource: %w", err))
		}
		if op.ApiOpModel == AsyncApiOp {
			observed, err := pollUntilConverged(ctx, d, meta, desired, *op.PollDefinition, get)
			if err != nil {
				return diag.FromErr(err)
			}
			if op.ApplyStateMappings != nil {
				if err := op.ApplyStateMappings(observed, d); err != nil {
					return diag.FromErr(err)
				}
			}
		}
		return nil
	}
}

func checkResourceProviderDefinition(def ResourceProviderDefinition) error {
	if len(def.Schema) == 0 {
		return errors.New("schema is required")
	}
	if def.Create.ApiCallerFn != nil {
		if err := validateCreateOp(def.Create); err != nil {
			return fmt.Errorf("create: %w", err)
		}
	}
	if def.Update.ApiCallerFn != nil {
		if err := validateUpdateOp(def.Update); err != nil {
			return fmt.Errorf("update: %w", err)
		}
	}
	if def.Delete.ApiCallerFn != nil {
		if err := validateDeleteOp(def.Delete); err != nil {
			return fmt.Errorf("delete: %w", err)
		}
	}
	return nil
}

func validateCreateOp(op CreateOpApiDefinition) error {
	if op.GetResourceID == nil {
		return errors.New("GetResourceID is required")
	}
	if op.ApiOpModel != AsyncApiOp && op.ApiOpModel != SyncApiOp {
		return errors.New("ApiOpModel must be set")
	}
	if op.ApiOpModel == AsyncApiOp {
		if op.DesiredState.ToStateFromSourceFn == nil {
			return errors.New("DesiredState.ToStateFromSourceFn is required for AsyncApiOp")
		}
		if err := isPollDefinitionMissing(op.PollDefinition); err != nil {
			return err
		}
	}
	return nil
}

func validateUpdateOp(op UpdateOpApiDefinition) error {
	if op.ApiOpModel != AsyncApiOp && op.ApiOpModel != SyncApiOp {
		return errors.New("ApiOpModel must be set")
	}
	if op.ApiOpModel == AsyncApiOp {
		if op.DesiredState.ToStateFromSourceFn == nil {
			return errors.New("DesiredState.ToStateFromSourceFn is required for AsyncApiOp")
		}
		if err := isPollDefinitionMissing(op.PollDefinition); err != nil {
			return err
		}
	}
	return nil
}

func validateDeleteOp(op DeleteOpApiDefinition) error {
	if op.ApiOpModel != AsyncApiOp && op.ApiOpModel != SyncApiOp {
		return errors.New("ApiOpModel must be set")
	}
	if op.ApiOpModel == AsyncApiOp {
		if op.DesiredState.ToStateFromSourceFn == nil {
			return errors.New("DesiredState.ToStateFromSourceFn is required for AsyncApiOp")
		}
		if err := isPollDefinitionMissing(op.PollDefinition); err != nil {
			return err
		}
	}
	return nil
}

func isPollDefinitionMissing(p *PollDefinition) error {
	if p == nil {
		return errors.New("PollDefinition is required for AsyncApiOp")
	}
	if p.PollFn == nil || p.ResourceComparator == nil {
		return errors.New("PollFn and ResourceComparator are required")
	}
	if p.MaxAttempts == 0 && p.MaxTimeout == 0 {
		return errors.New("specify MaxAttempts or MaxTimeout")
	}
	return nil
}

// withSettings invokes fn after fetching the shared impl data for the callback.
func withSettings[T any](ctx context.Context, d *schema.ResourceData, meta interface{}, get GetResourceProviderImplDataFn, fn func(context.Context) (T, error)) (T, error) {
	if _, err := get(ctx, d, meta); err != nil {
		var zero T
		return zero, err
	}
	return fn(ctx)
}

func pollUntilConverged(ctx context.Context, d *schema.ResourceData, meta interface{}, desired ResourceState, p PollDefinition, get GetResourceProviderImplDataFn) (ResourceState, error) {
	if p.MaxTimeout > 0 {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, p.MaxTimeout)
		defer cancel()
	}
	if err := sleepOrCancel(ctx, p.InitialWait); err != nil {
		return nil, err
	}
	for attempt := 1; p.MaxAttempts == 0 || attempt <= p.MaxAttempts; attempt++ {
		attemptCtx := ctx
		cancelAttempt := func() {}
		if p.EachPollTimeout > 0 {
			attemptCtx, cancelAttempt = context.WithTimeout(ctx, p.EachPollTimeout)
		}
		observed, err := withSettings(attemptCtx, d, meta, get, func(c context.Context) (ResourceState, error) {
			return p.PollFn(c, d, meta)
		})
		cancelAttempt()
		if err != nil {
			return nil, fmt.Errorf("poll attempt %d: %w", attempt, err)
		}
		done, err := p.ResourceComparator(desired, observed)
		if err != nil {
			return nil, fmt.Errorf("compare attempt %d: %w", attempt, err)
		}
		if done {
			return observed, nil
		}
		if p.MaxAttempts > 0 && attempt == p.MaxAttempts {
			break
		}
		if err := sleepOrCancel(ctx, p.PollInterval); err != nil {
			return nil, err
		}
	}
	return nil, fmt.Errorf("did not converge within %d attempts", p.MaxAttempts)
}

func sleepOrCancel(ctx context.Context, d time.Duration) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	if d == 0 {
		return nil
	}
	t := time.NewTimer(d)
	defer t.Stop()
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-t.C:
		return nil
	}
}
