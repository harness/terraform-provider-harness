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
		CreateContext: resourceFeatureFlagTargetCreate,
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
				Description: "The name of the feature flag target group.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"included": {
				Description: "A list of targets to include in the target group",
				Type:        schema.TypeList,
				Optional:    true,
				MinItems:    0,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"excluded": {
				Description: "A list of targets to exclude from the target group",
				Type:        schema.TypeList,
				Optional:    true,
				MinItems:    0,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"rule": {
				Description: "The list of rules used to include targets in the target group.",
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"attribute": {
							Description: "The attribute to use in the clause.  This can be any target attribute",
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
	Identifier string           `json:"identifier,omitempty"`
	Name       string           `json:"name,omitempty"`
	Included   []string         `json:"included,omitempty"`
	Excluded   []string         `json:"excluded,omitempty"`
	Rules      []nextgen.Clause `json:"rules,omitempty"`
}

// SegmentRequest ...
type SegmentRequest struct {
	Identifier  string           `json:"identifier,omitempty"`
	Project     string           `json:"project,omitempty"`
	Environment string           `json:"environment,omitempty"`
	Name        string           `json:"name,omitempty"`
	Included    []string         `json:"included,omitempty"`
	Excluded    []string         `json:"excluded,omitempty"`
	Rules       []nextgen.Clause `json:"rules,omitempty"`
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
func resourceFeatureFlagTargetCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)
	var err error
	var httpResp *http.Response
	var segment nextgen.Segment
	segmentRequest := buildSegmentRequest(d)
	qp := buildFFTargetGroupQueryParameters(d)
	id := d.Id()
	if id == "" {
		id = d.Get("identifier").(string)
		d.MarkNewResource()
	}

	httpResp, err = c.TargetGroupsApi.CreateSegment(ctx, segmentRequest, c.AccountId, qp.OrgID)

	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	if httpResp.StatusCode != http.StatusCreated {
		return diag.Errorf("createstatus: %s", httpResp.Status)
	}

	segment, httpResp, err = c.TargetGroupsApi.GetSegment(ctx, c.AccountId, qp.OrgID, id, qp.Project, qp.Environment)
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

	segment, httpResp, err = c.TargetGroupsApi.GetSegment(ctx, c.AccountId, qp.OrgID, id, qp.Project, qp.Environment)
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
}

// buildFFTargetGroupQueryParameters ...
func buildFFTargetGroupQueryParameters(d *schema.ResourceData) *FFTargetGroupQueryParameters {
	return &FFTargetGroupQueryParameters{
		Identifier:  d.Get("identifier").(string),
		OrgID:       d.Get("org_id").(string),
		Project:     d.Get("project_id").(string),
		AcountID:    d.Get("account_id").(string),
		Environment: d.Get("environment").(string),
	}
}

// buildSegmentRequest builds a SegmentRequest from a ResourceData
func buildSegmentRequest(d *schema.ResourceData) *SegmentRequest {
	opts := &SegmentRequest{
		Identifier:  d.Get("identifier").(string),
		Project:     d.Get("project_id").(string),
		Environment: d.Get("environment").(string),
		Name:        d.Get("name").(string),
	}

	if included, ok := d.GetOk("included"); ok {
		var targets []string
		for _, target := range included.([]interface{}) {
			targets = append(targets, target.(string))
		}
		opts.Included = targets
	}

	if excluded, ok := d.GetOk("excluded"); ok {
		var targets []string
		for _, target := range excluded.([]interface{}) {
			targets = append(targets, target.(string))
		}
		opts.Excluded = targets
	}

	if rules, ok := d.GetOk("rule"); ok {
		var rulesList []nextgen.Clause
		for _, rule := range rules.([]interface{}) {
			var values []string
			for _, value := range rule.(map[string]interface{})["values"].([]interface{}) {
				values = append(values, value.(string))
			}
			rule := nextgen.Clause{
				Attribute: rule.(map[string]interface{})["attribute"].(string),
				Negate:    rule.(map[string]interface{})["negate"].(bool),
				Op:        rule.(map[string]interface{})["op"].(string),
				Values:    values,
			}
			rules = append(rulesList, rule)
		}
		opts.Rules = rulesList
	}

	return opts
}

// buildFFTargetGroupOpts ...
func buildFFTargetGroupOpts(d *schema.ResourceData) *nextgen.TargetGroupsApiPatchSegmentOpts {
	opts := &FFTargetGroupOpts{
		Identifier: d.Get("identifier").(string),
		Name:       d.Get("name").(string),
	}

	if included, ok := d.GetOk("included"); ok {
		var targets []string
		for _, target := range included.([]interface{}) {
			targets = append(targets, target.(string))
		}
		opts.Included = targets
	}

	if excluded, ok := d.GetOk("excluded"); ok {
		var targets []string
		for _, target := range excluded.([]interface{}) {
			targets = append(targets, target.(string))
		}
		opts.Excluded = targets
	}

	if rules, ok := d.GetOk("rule"); ok {
		var rulesList []nextgen.Clause
		for _, rule := range rules.([]interface{}) {
			var values []string
			for _, value := range rule.(map[string]interface{})["values"].([]interface{}) {
				values = append(values, value.(string))
			}
			rule := nextgen.Clause{
				Attribute: rule.(map[string]interface{})["attribute"].(string),
				Negate:    rule.(map[string]interface{})["negate"].(bool),
				Op:        rule.(map[string]interface{})["op"].(string),
				Values:    values,
			}
			rules = append(rulesList, rule)
		}
		opts.Rules = rulesList
	}

	return &nextgen.TargetGroupsApiPatchSegmentOpts{
		Body: optional.NewInterface(opts),
	}
}
