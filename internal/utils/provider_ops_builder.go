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

// ConsistencyModel describes how the remote API behaves after a mutating call.
type ConsistencyModel string

const (
	// EventuallyConsistent means 2xx is returned but observed state may lag;
	// the framework polls until PollDefinition.ResourceComparator returns true.
	EventuallyConsistent ConsistencyModel = "eventually_consistent"
	// ImmediatelyConsistent means the API response already reflects final state;
	// no polling is performed after the mutation.
	ImmediatelyConsistent ConsistencyModel = "immediately_consistent"
)

// DesiredStateSource controls which snapshot feeds the desired-state capture.
type DesiredStateSource string

const (
	// SourceTerraformConfig captures field values from the evaluated HCL config
	// (d.Get), not from prior state or remote.
	SourceTerraformConfig DesiredStateSource = "terraform_config"
)

// ResourceProviderImplData is opaque shared data passed to every callback in the resource.
// The resource implementation owns its shape and interprets it as needed.
type ResourceProviderImplData interface{}

// ResourceProviderOpType identifies which CRUD operation is invoking a callback.
type ResourceProviderOpType string

const (
	OpTypeCreate ResourceProviderOpType = "create"
	OpTypeRead   ResourceProviderOpType = "read"
	OpTypeUpdate ResourceProviderOpType = "update"
	OpTypeDelete ResourceProviderOpType = "delete"
)

// ObservedStateSource identifies where the observed state originated.
type ObservedStateSource string

const (
	// ObservedFromRemoteAPI means the state was fetched from the remote API.
	ObservedFromRemoteAPI ObservedStateSource = "remote_api"
	// ObservedFromPoll means the state was fetched during a convergence poll.
	ObservedFromPoll ObservedStateSource = "poll"
)

// ApplyStateMappingsFn is called by the framework after each operation that
// produces observed state. The resource implementation is responsible for
// writing observed state into Terraform schema state (d.Set calls).
//
//   - opType: which CRUD operation produced the observed state
//   - observedSource: whether it came from a direct read or a convergence poll
//   - observed: the state returned by the remote API caller
//   - d: the Terraform schema state to write into
type ApplyStateMappingsFn func(opType ResourceProviderOpType, observedSource ObservedStateSource, observed ResourceState, d *schema.ResourceData) error

type (
	// CreateApiCallerFn performs the create mutation against the remote API.
	CreateApiCallerFn func(context.Context, *schema.ResourceData, interface{}) error
	// ReadApiCallerFn fetches the current observed state from the remote API.
	ReadApiCallerFn func(context.Context, *schema.ResourceData, interface{}) (ResourceState, error)
	// UpdateApiCallerFn performs the update mutation against the remote API.
	UpdateApiCallerFn func(context.Context, *schema.ResourceData, interface{}) error
	// DeleteApiCallerFn performs the delete mutation against the remote API.
	DeleteApiCallerFn func(context.Context, *schema.ResourceData, interface{}) error

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
	InitialWait        time.Duration
	PollInterval       time.Duration
	MaxAttempts        int
	EachPollTimeout    time.Duration
	MaxTimeout         time.Duration
	PollFn             ReadApiCallerFn
	ResourceComparator ResourceStateComparatorFn
}

// CreateOpApiDefinition describes the Create operation.
type CreateOpApiDefinition struct {
	ConsistencyModel   ConsistencyModel
	CreateApiCallerFn     CreateApiCallerFn
	DesiredStateSource DesiredStateSource
	GetResourceStateFn GetResourceStateFn
	PollDefinition     *PollDefinition
	GetResourceID      ResourceIDFn
}

// ReadOpApiDefinition describes the Read operation.
type ReadOpApiDefinition struct {
	ReadApiCallerFn ReadApiCallerFn
	// HandleNotFound is called when the remote read returns a not-found signal.
	// Return true to clear the resource ID (resource gone); false to leave state.
	HandleNotFound func(ResourceState) bool
}

// UpdateOpApiDefinition describes the Update operation.
type UpdateOpApiDefinition struct {
	ConsistencyModel   ConsistencyModel
	UpdateApiCallerFn  UpdateApiCallerFn
	DesiredStateSource DesiredStateSource
	GetResourceStateFn GetResourceStateFn
	PollDefinition     *PollDefinition
}

// DeleteOpApiDefinition describes the Delete operation.
type DeleteOpApiDefinition struct {
	ConsistencyModel   ConsistencyModel
	DeleteApiCallerFn  DeleteApiCallerFn
	DesiredStateSource DesiredStateSource
	GetResourceStateFn GetResourceStateFn
	PollDefinition     *PollDefinition
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

// ResourceProviderDefinition groups schema, the shared settings fn, and all five op definitions.
type ResourceProviderDefinition struct {
	Description                 string
	Schema                      map[string]*schema.Schema
	GetResourceProviderImplData GetResourceProviderImplDataFn
	// ApplyStateMappings is called after each operation that produces observed state.
	// The implementation writes observed state into Terraform schema state via d.Set.
	ApplyStateMappings ApplyStateMappingsFn
	Create             CreateOpApiDefinition
	Read               ReadOpApiDefinition
	Update             UpdateOpApiDefinition
	Delete             DeleteOpApiDefinition
	Import             ImportOpDefinition
}

// NewResourceProvider builds a *schema.Resource with all four CRUD context functions.
func NewResourceProvider(def ResourceProviderDefinition) (*schema.Resource, error) {
	if err := checkResourceProviderDefinition(def); err != nil {
		return nil, err
	}
	get := def.GetResourceProviderImplData
	r := &schema.Resource{
		Description: def.Description,
		Schema:      def.Schema,
		Importer:    def.Import.toResourceImporter(),
	}
	setResourceProviderHandlers(r, def, get)
	return r, nil
}

func setResourceProviderHandlers(r *schema.Resource, def ResourceProviderDefinition, get GetResourceProviderImplDataFn) {
	apply := def.ApplyStateMappings
	if def.Create.CreateApiCallerFn != nil {
		r.CreateContext = setCreateOpHandler(def.Create, get, apply)
	} else {
		r.CreateContext = notImplementedHandler("create")
	}
	if def.Read.ReadApiCallerFn != nil {
		r.ReadContext = setReadOpHandler(def.Read, get, apply)
	} else {
		r.ReadContext = notImplementedHandler("read")
	}
	if def.Update.UpdateApiCallerFn != nil {
		r.UpdateContext = setUpdateOpHandler(def.Update, get, apply)
	} else {
		r.UpdateContext = notImplementedHandler("update")
	}
	if def.Delete.DeleteApiCallerFn != nil {
		r.DeleteContext = setDeleteOpHandler(def.Delete, get, apply)
	} else {
		r.DeleteContext = notImplementedHandler("delete")
	}
}

func notImplementedHandler(op string) func(context.Context, *schema.ResourceData, interface{}) diag.Diagnostics {
	return func(_ context.Context, _ *schema.ResourceData, _ interface{}) diag.Diagnostics {
		return diag.Errorf("%s operation is not implemented for this resource", op)
	}
}

func setCreateOpHandler(op CreateOpApiDefinition, get GetResourceProviderImplDataFn, apply ApplyStateMappingsFn) schema.CreateContextFunc {
	return func(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
		var desired ResourceState
		if op.ConsistencyModel == EventuallyConsistent {
			var err error
			desired, err = op.GetResourceStateFn(d)
			if err != nil {
				return diag.FromErr(fmt.Errorf("capture desired state: %w", err))
			}
		}
		_, err := withSettings(ctx, d, meta, get, func(c context.Context) (struct{}, error) {
			return struct{}{}, op.CreateApiCallerFn(c, d, meta)
		})
		if err != nil {
			return diag.FromErr(fmt.Errorf("create remote resource: %w", err))
		}
		id, err := op.GetResourceID(d)
		if err != nil {
			return diag.FromErr(err)
		}
		d.SetId(id)
		if op.ConsistencyModel == EventuallyConsistent {
			observed, err := pollUntilConverged(ctx, d, meta, desired, *op.PollDefinition, get)
			if err != nil {
				return diag.FromErr(err)
			}
			if apply != nil {
				if err := apply(OpTypeCreate, ObservedFromPoll, observed, d); err != nil {
					return diag.FromErr(err)
				}
			}
		}
		return nil
	}
}

func setReadOpHandler(op ReadOpApiDefinition, get GetResourceProviderImplDataFn, apply ApplyStateMappingsFn) schema.ReadContextFunc {
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
		if apply != nil {
			if err := apply(OpTypeRead, ObservedFromRemoteAPI, observed, d); err != nil {
				return diag.FromErr(err)
			}
		}
		return nil
	}
}

func setUpdateOpHandler(op UpdateOpApiDefinition, get GetResourceProviderImplDataFn, apply ApplyStateMappingsFn) schema.UpdateContextFunc {
	return func(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
		var desired ResourceState
		if op.ConsistencyModel == EventuallyConsistent {
			var err error
			desired, err = op.GetResourceStateFn(d)
			if err != nil {
				return diag.FromErr(fmt.Errorf("capture desired state: %w", err))
			}
		}
		_, err := withSettings(ctx, d, meta, get, func(c context.Context) (struct{}, error) {
			return struct{}{}, op.UpdateApiCallerFn(c, d, meta)
		})
		if err != nil {
			return diag.FromErr(fmt.Errorf("update remote resource: %w", err))
		}
		if op.ConsistencyModel == EventuallyConsistent {
			observed, err := pollUntilConverged(ctx, d, meta, desired, *op.PollDefinition, get)
			if err != nil {
				return diag.FromErr(err)
			}
			if apply != nil {
				if err := apply(OpTypeUpdate, ObservedFromPoll, observed, d); err != nil {
					return diag.FromErr(err)
				}
			}
		}
		return nil
	}
}

func setDeleteOpHandler(op DeleteOpApiDefinition, get GetResourceProviderImplDataFn, apply ApplyStateMappingsFn) schema.DeleteContextFunc {
	return func(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
		var desired ResourceState
		if op.ConsistencyModel == EventuallyConsistent {
			var err error
			desired, err = op.GetResourceStateFn(d)
			if err != nil {
				return diag.FromErr(fmt.Errorf("capture desired state: %w", err))
			}
		}
		_, err := withSettings(ctx, d, meta, get, func(c context.Context) (struct{}, error) {
			return struct{}{}, op.DeleteApiCallerFn(c, d, meta)
		})
		if err != nil {
			return diag.FromErr(fmt.Errorf("delete remote resource: %w", err))
		}
		if op.ConsistencyModel == EventuallyConsistent {
			observed, err := pollUntilConverged(ctx, d, meta, desired, *op.PollDefinition, get)
			if err != nil {
				return diag.FromErr(err)
			}
			if apply != nil {
				if err := apply(OpTypeDelete, ObservedFromPoll, observed, d); err != nil {
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
	if def.GetResourceProviderImplData == nil {
		return errors.New("GetResourceProviderImplData is required")
	}
	if def.Create.CreateApiCallerFn != nil {
		if err := validateCreateOp(def.Create); err != nil {
			return fmt.Errorf("create: %w", err)
		}
	}

	if def.Update.UpdateApiCallerFn != nil {
		if err := validateUpdateOp(def.Update); err != nil {
			return fmt.Errorf("update: %w", err)
		}
	}
	if def.Delete.DeleteApiCallerFn != nil {
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
	if op.ConsistencyModel != EventuallyConsistent && op.ConsistencyModel != ImmediatelyConsistent {
		return errors.New("ConsistencyModel must be set")
	}
	if op.ConsistencyModel == EventuallyConsistent {
		if op.GetResourceStateFn == nil {
			return errors.New("GetResourceStateFn is required for EventuallyConsistent")
		}
		if err := isPollDefinitionMissing(op.PollDefinition); err != nil {
			return err
		}
	}
	return nil
}

func validateUpdateOp(op UpdateOpApiDefinition) error {
	if op.ConsistencyModel != EventuallyConsistent && op.ConsistencyModel != ImmediatelyConsistent {
		return errors.New("ConsistencyModel must be set")
	}
	if op.ConsistencyModel == EventuallyConsistent {
		if op.GetResourceStateFn == nil {
			return errors.New("GetResourceStateFn is required for EventuallyConsistent")
		}
		if err := isPollDefinitionMissing(op.PollDefinition); err != nil {
			return err
		}
	}
	return nil
}

func validateDeleteOp(op DeleteOpApiDefinition) error {
	if op.ConsistencyModel != EventuallyConsistent && op.ConsistencyModel != ImmediatelyConsistent {
		return errors.New("ConsistencyModel must be set")
	}
	if op.ConsistencyModel == EventuallyConsistent {
		if op.GetResourceStateFn == nil {
			return errors.New("GetResourceStateFn is required for EventuallyConsistent")
		}
		if err := isPollDefinitionMissing(op.PollDefinition); err != nil {
			return err
		}
	}
	return nil
}

func isPollDefinitionMissing(p *PollDefinition) error {
	if p == nil {
		return errors.New("PollDefinition is required for EventuallyConsistent")
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
