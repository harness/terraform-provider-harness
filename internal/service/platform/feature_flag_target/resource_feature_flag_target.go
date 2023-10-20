package feature_flag_target

import (
	"context"
	"io"
	"net/http"
	"time"

	"github.com/antihax/optional"
	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceFeatureFlagTarget() *schema.Resource {
	resource := &schema.Resource{
		Description: "Resource for managing Feature Flag Targets.",

		ReadContext:   resourceFeatureFlagTargetRead,
		DeleteContext: resourceFeatureFlagTargetDelete,
		CreateContext: resourceFeatureFlagTargetCreateOrUpdate,
		UpdateContext: resourceFeatureFlagTargetCreateOrUpdate,
		Importer:      helpers.ProjectResourceImporter,

		Schema: map[string]*schema.Schema{
			"identifier": {
				Description: "Identifier of the Feature Flag Target",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
			},
			"org_id": {
				Description: "Organization Identifier",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
			},
			"project_id": {
				Description: "Project Identifier",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
			},
			"environment": {
				Description: "Environment Identifier",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
			},
			"account_id": {
				Description: "Account Identifier",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
			},
			"name": {
				Description: "Target Name",
				Type:        schema.TypeString,
				Required:    true,
			},
			"attributes": {
				Description: "Attributes",
				Type:        schema.TypeMap,
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}

	return resource
}

// FFTargetQueryParameters is the query parameters for the feature flag target
type FFTargetQueryParameters struct {
	Identifier     string
	OrganizationID string
	ProjectID      string
	AccountID      string
	Environment    string
}

// FFTargetOpts is the options for the feature flag target
type FFTargetOpts struct {
	Name      string
	Atributes map[string]interface{}
}

// resourceFeatureFlagTargetRead is the read function for the feature flag target
func resourceFeatureFlagTargetRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)
	id := d.Id()

	if id == "" {
		d.MarkNewResource()
		return nil
	}

	qp := buildFFTargetQueryParameters(d)

	resp, httpResp, err := c.TargetsApi.GetTarget(ctx, id, c.AccountId, qp.OrganizationID, qp.ProjectID, qp.Environment)

	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	readFeatureFlagTarget(d, &resp, *qp)

	return nil
}

// resourceFeatureFlagTargetDelete is the delete function for the feature flag target
func resourceFeatureFlagTargetDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	id := d.Id()
	if id == "" {
		return nil
	}

	qp := buildFFTargetQueryParameters(d)

	httpResp, err := c.TargetsApi.DeleteTarget(ctx, id, c.AccountId, qp.OrganizationID, qp.ProjectID, qp.Environment)
	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}
	return nil
}

// resourceFeatureFlagTargetCreateOrUpdate is the create function for the feature flag target
func resourceFeatureFlagTargetCreateOrUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	var err error
	var httpResp *http.Response
	target := buildFFTargetCreate(d)
	qp := buildFFTargetQueryParameters(d)
	id := d.Id()

	if id == "" {
		httpResp, err = c.TargetsApi.CreateTarget(ctx, target, c.AccountId, qp.OrganizationID)
	} else {
		target, httpResp, err = c.TargetsApi.ModifyTarget(ctx, target, c.AccountId, qp.OrganizationID, qp.ProjectID, qp.Environment, id)
	}

	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	readFeatureFlagTarget(d, &target, *qp)

	return nil
}

// resourceFeatureFlagTargetUpdate is the update function for the feature flag target
func resourceFeatureFlagTargetUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	id := d.Id()
	if id == "" {
		id = d.Get("identifier").(string)
		d.MarkNewResource()
	}

	qp := buildFFTargetQueryParameters(d)
	opts := buildFFTargetPatchOpts(d)

	var err error
	var target nextgen.Target
	var httpResp *http.Response

	target, httpResp, err = c.TargetsApi.PatchTarget(ctx, c.AccountId, qp.OrganizationID, qp.ProjectID, qp.Environment, id, opts)
	if err != nil {
		body, _ := io.ReadAll(httpResp.Body)
		return diag.Errorf("readstatus: %s, \nBody:%s", httpResp.Status, body)
	}

	readFeatureFlagTarget(d, &target, *qp)

	return nil
}

// readFeatureFlagTarget is the read function for the feature flag target
func readFeatureFlagTarget(d *schema.ResourceData, flag *nextgen.Target, qp FFTargetQueryParameters) {
	d.SetId(qp.Identifier)
	d.Set("identifier", qp.Identifier)
	d.Set("org_id", qp.OrganizationID)
	d.Set("project_id", qp.ProjectID)
	d.Set("environment", qp.Environment)
	d.Set("account_id", qp.AccountID)
	d.Set("name", flag.Name)
	d.Set("attributes", flag.Attributes)
}

// buildFFTargetQueryParameters is the query parameters for the feature flag target
func buildFFTargetQueryParameters(d *schema.ResourceData) *FFTargetQueryParameters {
	return &FFTargetQueryParameters{
		Identifier:     d.Get("identifier").(string),
		OrganizationID: d.Get("org_id").(string),
		ProjectID:      d.Get("project_id").(string),
		AccountID:      d.Get("account_id").(string),
		Environment:    d.Get("environment").(string),
	}
}

// buildFFTargetCreateOpts
func buildFFTargetCreate(d *schema.ResourceData) nextgen.Target {
	attribute := d.Get("attributes")
	return nextgen.Target{
		Account:     d.Get("account_id").(string),
		Attributes:  &attribute,
		Environment: d.Get("environment").(string),
		Identifier:  d.Get("identifier").(string),
		Org:         d.Get("org_id").(string),
		Name:        d.Get("name").(string),
		Project:     d.Get("project_id").(string),
		CreatedAt:   time.Now().Unix(),
	}
}

// buildFFTargetPatchOpts is the options for the feature flag target
func buildFFTargetPatchOpts(d *schema.ResourceData) *nextgen.TargetsApiPatchTargetOpts {
	opts := &FFTargetOpts{
		Name:      d.Get("name").(string),
		Atributes: d.Get("attributes").(map[string]interface{}),
	}

	return &nextgen.TargetsApiPatchTargetOpts{
		Body: optional.NewInterface(opts),
	}
}
