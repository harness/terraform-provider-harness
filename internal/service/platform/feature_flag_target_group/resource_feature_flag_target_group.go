package featureflagtargetgroup

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

// ResourceFeatureFlagTargetGroup ...
func ResourceFeatureFlagTargetGroup() *schema.Resource {
	resource := &schema.Resource{
		Description: "Resource for creating a Harness Feature Flag Target Group.",

		ReadContext:   resourceFeatureFlagTargetGroupRead,
		CreateContext: resourceFeatureFlagTargetCreateOrUpdate,
		UpdateContext: resourceFeatureFlagTargetGroupUpdate,
		DeleteContext: resourceFeatureFlagTargetGroupDelete,
		Importer:      helpers.ProjectResourceImporter,

		Schema: map[string]*schema.Schema{
			"identifier": {
				Description: "The unique identifier of the feature flag target group.",
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
			"project": {
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
				Description: "The name of the feature flag target group.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"included": {
				Description: "The list of rules to include in the feature flag target group.",
				Type:        schema.TypeList,
				Optional:    true,
				MinItems:    0,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"excluded": {
				Description: "The list of rules to include in the feature flag target group.",
				Type:        schema.TypeList,
				Optional:    true,
				MinItems:    0,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"rules": {
				Description: "The list of rules to exclude in the feature flag target group.",
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"attribute": {
							Description: "The attribute to use in the clause.  This can be any target attribute",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"id": {
							Description: "The unique ID for the clause",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"negate": {
							Description: "Is the operation negated?",
							Type:        schema.TypeBool,
							Optional:    true,
						},
						"op": {
							Description: "The type of operation such as equals, starts_with, contains",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"values": {
							Description: "The values that are compared against the operator",
							Type:        schema.TypeList,
							Optional:    true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
		},
	}

	return resource
}

// FFTargetGroupQueryParameters ...
type FFTargetGroupQueryParameters struct {
	Identifier  string `json:"identifier,omitempty"`
	OrgID       string `json:"orgId,omitempty"`
	Project     string `json:"project,omitempty"`
	AcountID    string `json:"accountId,omitempty"`
	Environment string `json:"environment,omitempty"`
}

// FFTargetGroupOpts ...
type FFTargetGroupOpts struct {
	Identifier  string           `json:"identifier,omitempty"`
	Name        string           `json:"name,omitempty"`
	Description string           `json:"description,omitempty"`
	Included    []nextgen.Target `json:"included,omitempty"`
	Excluded    []nextgen.Target `json:"excluded,omitempty"`
	Rules       []nextgen.Clause `json:"rules,omitempty"`
	Tags        []nextgen.Tag    `json:"tags,omitempty"`
}

func resourceFeatureFlagTargetGroupRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	id := d.Id()
	if id == "" {
		d.MarkNewResource()
		return nil
	}

	qp := buildFFTargetGroupQueryParameters(d)

	segment, httpResp, err := c.TargetGroupsApi.GetSegment(ctx, c.AccountId, qp.OrgID, id, qp.Project, qp.Environment)
	if err != nil {
		return helpers.HandleReadApiError(err, d, httpResp)
	}

	readFeatureFlagTargetGroup(d, &segment, qp)

	return nil
}

// resourceFeatureFlagTargetGroupCreate ...
func resourceFeatureFlagTargetCreateOrUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)
	var err error
	var httpResp *http.Response
	segment := buildFFTargetGroupCreate(d)
	qp := buildFFTargetGroupQueryParameters(d)
	id := d.Id()

	if id == "" {
		httpResp, err = c.TargetGroupsApi.CreateSegment(ctx, segment, c.AccountId, qp.OrgID)
	} else {
		segment, httpResp, err = c.TargetGroupsApi.PatchSegment(ctx, c.AccountId, qp.OrgID, qp.Project, qp.Environment, id, buildFFTargetGroupOpts(d))
	}

	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	readFeatureFlagTargetGroup(d, &segment, qp)

	return nil
}

func resourceFeatureFlagTargetGroupUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	id := d.Id()
	if id == "" {
		return nil
	}

	qp := buildFFTargetGroupQueryParameters(d)
	opts := buildFFTargetGroupOpts(d)

	var err error
	var segment nextgen.Segment
	var httpResp *http.Response

	segment, httpResp, err = c.TargetGroupsApi.PatchSegment(ctx, c.AccountId, qp.OrgID, qp.Project, qp.Environment, id, opts)
	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	time.Sleep(1 * time.Second)

	segment, httpResp, err = c.TargetGroupsApi.GetSegment(ctx, id, c.AccountId, qp.OrgID, qp.Project, qp.Environment)
	if err != nil {
		body, _ := io.ReadAll(httpResp.Body)
		return diag.Errorf("readstatus: %s, \nBody:%s", httpResp.Status, body)
	}

	readFeatureFlagTargetGroup(d, &segment, qp)

	return nil
}

func resourceFeatureFlagTargetGroupDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	id := d.Id()
	if id == "" {
		return nil
	}

	qp := buildFFTargetGroupQueryParameters(d)

	httpResp, err := c.TargetGroupsApi.DeleteSegment(ctx, c.AccountId, qp.OrgID, id, qp.Project, qp.Environment)
	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	return nil
}

// readFeatureFlagTargetGroupRule ...
func readFeatureFlagTargetGroup(d *schema.ResourceData, segment *nextgen.Segment, qp *FFTargetGroupQueryParameters) {
	d.SetId(segment.Identifier)
	d.Set("identifier", segment.Identifier)
	d.Set("org_id", qp.OrgID)
	d.Set("project", qp.Project)
	d.Set("account_id", qp.AcountID)
	d.Set("environment", segment.Environment)
	d.Set("name", segment.Name)
	d.Set("included", segment.Included)
	d.Set("excluded", segment.Excluded)
	d.Set("rules", segment.Rules)
	d.Set("tags", segment.Tags)
}

// buildFFTargetGroupQueryParameters ...
func buildFFTargetGroupQueryParameters(d *schema.ResourceData) *FFTargetGroupQueryParameters {
	return &FFTargetGroupQueryParameters{
		Identifier:  d.Get("identifier").(string),
		OrgID:       d.Get("org_id").(string),
		Project:     d.Get("project").(string),
		AcountID:    d.Get("account_id").(string),
		Environment: d.Get("environment").(string),
	}
}

// buildFFTargetGroupCreate ...
func buildFFTargetGroupCreate(d *schema.ResourceData) nextgen.Segment {
	opts := nextgen.Segment{
		Identifier:  d.Get("identifier").(string),
		Environment: d.Get("environment").(string),
		Name:        d.Get("name").(string),
		CreatedAt:   time.Now().Unix(),
	}

	if included, ok := d.GetOk("included"); ok {
		opts.Included = included.([]nextgen.Target)
	}

	if excluded, ok := d.GetOk("excluded"); ok {
		opts.Excluded = excluded.([]nextgen.Target)
	}

	if rules, ok := d.GetOk("rules"); ok {
		opts.Rules = rules.([]nextgen.Clause)
	}

	if tags, ok := d.GetOk("tags"); ok {
		opts.Tags = tags.([]nextgen.Tag)
	}

	return opts
}

// buildFFTargetGroupOpts ...
func buildFFTargetGroupOpts(d *schema.ResourceData) *nextgen.TargetGroupsApiPatchSegmentOpts {
	opts := &FFTargetGroupOpts{
		Identifier:  d.Get("identifier").(string),
		Name:        d.Get("name").(string),
		Description: d.Get("description").(string),
	}

	if included, ok := d.GetOk("included"); ok {
		opts.Included = included.([]nextgen.Target)
	}

	if excluded, ok := d.GetOk("excluded"); ok {
		opts.Excluded = excluded.([]nextgen.Target)
	}

	if rules, ok := d.GetOk("rules"); ok {
		opts.Rules = rules.([]nextgen.Clause)
	}

	if tags, ok := d.GetOk("tags"); ok {
		opts.Tags = tags.([]nextgen.Tag)
	}

	return &nextgen.TargetGroupsApiPatchSegmentOpts{
		Body: optional.NewInterface(opts),
	}
}
