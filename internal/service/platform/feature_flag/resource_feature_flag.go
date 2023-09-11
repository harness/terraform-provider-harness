package feature_flag

import (
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/antihax/optional"
	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceFeatureFlag() *schema.Resource {
	resource := &schema.Resource{
		Description: "Resource for managing Feature Flags.",

		ReadContext:   resourceFeatureFlagRead,
		DeleteContext: resourceFeatureFlagDelete,
		CreateContext: resourceFeatureFlagCreate,
		UpdateContext: resourceFeatureFlagUpdate,
		Importer:      helpers.ProjectResourceImporter,

		Schema: map[string]*schema.Schema{
			"identifier": {
				Description: "Identifier of the Feature Flag",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
			},
			"name": {
				Description: "Name of the Feature Flag",
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
			"archived": {
				Description: "Whether or not the flag is archived",
				Type:        schema.TypeBool,
				Optional:    true,
				ForceNew:    true,
			},
			"default_off_variation": {
				Description: "Which of the variations to use when the flag is toggled to off state",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
			},
			"default_on_variation": {
				Description: "Which of the variations to use when the flag is toggled to on state",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
			},
			"git_details": {
				Type:     schema.TypeSet,
				Optional: true,
				ForceNew: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"commit_msg": {
							Description: "The commit message to use as part of a gitsync operation",
							Type:        schema.TypeString,
							Required:    true,
						},
					},
				},
			},
			"kind": {
				Description: "The type of data the flag represents. Valid values are `boolean`, `int`, `string`, `json`",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
			},
			"owner": {
				Description: "The owner of the flag",
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
			},
			"permanent": {
				Description: "Whether or not the flag is permanent. If it is, it will never be flagged as stale",
				Type:        schema.TypeBool,
				Required:    true,
				ForceNew:    true,
			},
			"variation": {
				Description: "The options available for your flag",
				Type:        schema.TypeList,
				Required:    true,
				ForceNew:    true,
				MinItems:    2,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"identifier": {
							Description: "The identifier of the variation",
							Type:        schema.TypeString,
							Required:    true,
						},
						"description": {
							Description: "The description of the variation",
							Type:        schema.TypeString,
							Required:    true,
						},
						"name": {
							Description: "The user friendly name of the variation",
							Type:        schema.TypeString,
							Required:    true,
						},
						"value": {
							Description: "The value of the variation",
							Type:        schema.TypeString,
							Required:    true,
						},
					},
				},
			},
			"instructions": {
				Description: "The targeting rules for the flag",
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"kind": {
							Description: "The type of targeting rule. Valid values are `removeTargets`, `removeRule`, `addRule`, `addTargets`",
							Type:        schema.TypeString,
							Required:    true,
						},
						"parameters": {
							Description: "Whether or not the targeting rules are enabled",
							Type:        schema.TypeList,
							Required:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"uuid": {
										Description: "The identifier of the parameter",
										Type:        schema.TypeString,
										Optional:    true,
									},
									"ruleID": {
										Description: "The identifier of the rule",
										Type:        schema.TypeString,
										Optional:    true,
									},
									"priority": {
										Description: "The priority of the rule",
										Type:        schema.TypeString,
										Optional:    true,
									},
									"variation": {
										Description: "The identifier of the variation",
										Type:        schema.TypeString,
										Optional:    true,
									},
									"targets": {
										Description: "The targets of the rule",
										Type:        schema.TypeList,
										Optional:    true,
										MinItems:    0,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"clauses": {
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
									"serve": {
										Description: "Whether or not the rule is enabled",
										Type:        schema.TypeList,
										Optional:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"variation": {
													Description: "The identifier of the variation. Valid values are `enabled`, `disabled`",
													Type:        schema.TypeString,
													Optional:    true,
												},
												"distribution": {
													Description: "The distribution of the rule",
													Type:        schema.TypeList,
													Optional:    true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"bucketBy": {
																Description: "The bucketing strategy of the rule",
																Type:        schema.TypeString,
																Optional:    true,
															},
															"variations": {
																Description: "The variations of the rule",
																Type:        schema.TypeList,
																Optional:    true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"variation": {
																			Description: "The identifier of the variation",
																			Type:        schema.TypeString,
																			Optional:    true,
																		},
																		"weight": {
																			Description: "The weight of the variation",
																			Type:        schema.TypeInt,
																			Optional:    true,
																		},
																	},
																},
															},
														},
													},
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	return resource
}

type FFQueryParameters struct {
	Identifier     string
	OrganizationId string
	ProjectId      string
}

// KindMap is a map of the kind to the actual kind
var KindMap = map[string]string{
	"removeTargets": "removeTargetsToVariationTargetMap",
	"removeRule":    "removeRule",
	"addRule":       "addRule",
	"addTargets":    "addTargetsToVariationTargetMap",
}

// Variation ...
type Variation struct {
	Variation string `json:"variation,omitempty"`
	Weight    int    `json:"weight,omitempty"`
}

// Distribution ...
type Distribution struct {
	BucketBy   string      `json:"bucketBy,omitempty"`
	Variations []Variation `json:"variations,omitempty"`
}

// Serve ...
type Serve struct {
	Variation    string         `json:"variation,omitempty"`
	Distribution []Distribution `json:"distribution,omitempty"`
}

// Parameter ...
type Parameter struct {
	RuleID    string           `json:"ruleId,omitempty"`
	Variation string           `json:"variation,omitempty"`
	Targets   []string         `json:"targets,omitempty"`
	Priority  string           `json:"priority,omitempty"`
	Clauses   []nextgen.Clause `json:"clauses,omitempty"`
	Serve     []Serve          `json:"serve,omitempty"`
}

// Instructions ...
type Instructions struct {
	Kind       string      `json:"kind"`
	Parameters []Parameter `json:"parameters"`
}

type FFOpts struct {
	Identifier          string              `json:"identifier"`
	Name                string              `json:"name"`
	Description         string              `json:"description,omitempty"`
	Archived            bool                `json:"archived,omitempty"`
	DefaultOffVariation string              `json:"defaultOffVariation"`
	DefaultOnVariation  string              `json:"defaultOnVariation"`
	GitDetails          nextgen.GitDetails  `json:"gitDetails,omitempty"`
	Kind                string              `json:"kind"`
	Owner               string              `json:"owner,omitempty"`
	Permanent           bool                `json:"permanent"`
	Project             string              `json:"project"`
	Variations          []nextgen.Variation `json:"variations"`
}

// FFPatchOpts is the options for patching a feature flag
type FFPatchOpts struct {
	Identifier          string              `json:"identifier"`
	Name                string              `json:"name"`
	Description         string              `json:"description,omitempty"`
	Archived            bool                `json:"archived,omitempty"`
	DefaultOffVariation string              `json:"defaultOffVariation"`
	DefaultOnVariation  string              `json:"defaultOnVariation"`
	GitDetails          nextgen.GitDetails  `json:"gitDetails,omitempty"`
	Kind                string              `json:"kind"`
	Owner               string              `json:"owner,omitempty"`
	Permanent           bool                `json:"permanent"`
	Project             string              `json:"project"`
	Variations          []nextgen.Variation `json:"variations"`
	Instructions        Instructions        `json:"instructions"`
}

func resourceFeatureFlagUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	id := d.Id()
	if id == "" {
		return nil
	}

	qp := buildFFQueryParameters(d)
	opts := buildFFPatchOpts(d)

	feature, httpResp, err := c.FeatureFlagsApi.PatchFeature(ctx, c.AccountId, qp.OrganizationId, qp.ProjectId, id, opts)

	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	readFeatureFlag(d, &feature, qp)

	return nil
}

func resourceFeatureFlagRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	id := d.Id()
	if id == "" {
		d.MarkNewResource()
		return nil
	}

	qp := buildFFQueryParameters(d)
	opts := buildFFReadOpts(d)

	resp, httpResp, err := c.FeatureFlagsApi.GetFeatureFlag(ctx, id, c.AccountId, qp.OrganizationId, qp.ProjectId, opts)

	if err != nil {
		return helpers.HandleReadApiError(err, d, httpResp)
	}

	readFeatureFlag(d, &resp, qp)

	return nil
}

func resourceFeatureFlagCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	id := d.Id()
	if id == "" {
		id = d.Get("identifier").(string)
		d.MarkNewResource()
	}

	qp := buildFFQueryParameters(d)
	opts := buildFFCreateOpts(d)
	readOpts := buildFFReadOpts(d)

	var err error
	var resp nextgen.Feature
	var httpResp *http.Response

	httpResp, err = c.FeatureFlagsApi.CreateFeatureFlag(ctx, c.AccountId, qp.OrganizationId, opts)

	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	time.Sleep(1 * time.Second)

	resp, httpResp, err = c.FeatureFlagsApi.GetFeatureFlag(ctx, id, c.AccountId, qp.OrganizationId, qp.ProjectId, readOpts)

	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	readFeatureFlag(d, &resp, qp)

	return nil
}

func resourceFeatureFlagDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	id := d.Id()
	if id == "" {
		return nil
	}
	qp := buildFFQueryParameters(d)

	httpResp, err := c.FeatureFlagsApi.DeleteFeatureFlag(ctx, d.Id(), c.AccountId, qp.OrganizationId, qp.ProjectId, &nextgen.FeatureFlagsApiDeleteFeatureFlagOpts{CommitMsg: optional.EmptyString()})
	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	return nil
}

func readFeatureFlag(d *schema.ResourceData, flag *nextgen.Feature, qp *FFQueryParameters) {
	d.SetId(flag.Identifier)
	d.Set("identifier", flag.Identifier)
	d.Set("name", flag.Name)
	d.Set("project_id", flag.Project)
	d.Set("default_on_variation", flag.DefaultOnVariation)
	d.Set("default_off_variation", flag.DefaultOffVariation)
	d.Set("description", flag.Description)
	d.Set("kind", flag.Kind)
	d.Set("permanent", flag.Permanent)
	d.Set("owner", strings.Join(flag.Owner, ","))
	d.Set("org_id", qp.OrganizationId)
	d.Set("variation", expandVariations(flag.Variations))
}

func expandVariations(variations []nextgen.Variation) []interface{} {
	var result []interface{}
	for _, variation := range variations {
		result = append(result, map[string]interface{}{
			"identifier":  variation.Identifier,
			"name":        variation.Name,
			"description": variation.Description,
			"value":       variation.Value,
		})
	}

	return result
}

func buildFFQueryParameters(d *schema.ResourceData) *FFQueryParameters {
	return &FFQueryParameters{
		Identifier:     d.Get("identifier").(string),
		OrganizationId: d.Get("org_id").(string),
		ProjectId:      d.Get("project_id").(string),
	}
}

func buildFFCreateOpts(d *schema.ResourceData) *nextgen.FeatureFlagsApiCreateFeatureFlagOpts {
	opts := &FFOpts{
		Identifier:          d.Get("identifier").(string),
		Name:                d.Get("name").(string),
		DefaultOffVariation: d.Get("default_off_variation").(string),
		DefaultOnVariation:  d.Get("default_on_variation").(string),
		Project:             d.Get("project_id").(string),
		Kind:                d.Get("kind").(string),
	}

	if desc, ok := d.GetOk("description"); ok {
		opts.Description = desc.(string)
	}

	if owner, ok := d.GetOk("owner"); ok {
		opts.Owner = owner.(string)
	}

	if archived, ok := d.GetOk("archived"); ok {
		opts.Archived = archived.(bool)
	}

	var variations []nextgen.Variation
	variationsData := d.Get("variation").([]interface{})
	for _, variationData := range variationsData {
		vMap := variationData.(map[string]interface{})
		variation := nextgen.Variation{
			Identifier:  vMap["identifier"].(string),
			Value:       vMap["value"].(string),
			Name:        vMap["name"].(string),
			Description: vMap["description"].(string),
		}
		variations = append(variations, variation)
	}
	opts.Variations = variations

	return &nextgen.FeatureFlagsApiCreateFeatureFlagOpts{
		Body: optional.NewInterface(opts),
	}
}

func buildFFPatchOpts(d *schema.ResourceData) *nextgen.FeatureFlagsApiPatchFeatureOpts {
	opts := &FFPatchOpts{
		Identifier:          d.Get("identifier").(string),
		Name:                d.Get("name").(string),
		DefaultOffVariation: d.Get("default_off_variation").(string),
		DefaultOnVariation:  d.Get("default_on_variation").(string),
		Project:             d.Get("project_id").(string),
		Kind:                d.Get("kind").(string),
	}

	if desc, ok := d.GetOk("description"); ok {
		opts.Description = desc.(string)
	}

	if owner, ok := d.GetOk("owner"); ok {
		opts.Owner = owner.(string)
	}

	if archived, ok := d.GetOk("archived"); ok {
		opts.Archived = archived.(bool)
	}

	var variations []nextgen.Variation
	variationsData := d.Get("variation").([]interface{})
	for _, variationData := range variationsData {
		vMap := variationData.(map[string]interface{})
		variation := nextgen.Variation{
			Identifier:  vMap["identifier"].(string),
			Value:       vMap["value"].(string),
			Name:        vMap["name"].(string),
			Description: vMap["description"].(string),
		}
		variations = append(variations, variation)
	}
	opts.Variations = variations

	// TODO: Add instructions
	var instructions []Instructions
	instructionsData := d.Get("instructions").([]interface{})
	for _, instructionData := range instructionsData {
		vMap := instructionData.(map[string]interface{})
		parameters := vMap["parameters"].([]interface{})[0].(map[string]interface{})
		clauses := parameters["clauses"].([]interface{})
		serve := parameters["serve"].([]interface{})

		// get all the variations
		var variations []Variation
		for _, variation := range serve[0].(map[string]interface{})["distribution"].([]interface{})[0].(map[string]interface{})["variations"].([]interface{}) {
			current := Variation{
				Variation: variation.(map[string]interface{})["variation"].(string),
				Weight:    variation.(map[string]interface{})["weight"].(int),
			}
			variations = append(variations, current)
		}

		// build the Instruction
		instruction := Instructions{
			Kind: vMap["kind"].(string),
			Parameters: []Parameter{
				{
					Variation: parameters["variation"].(string),
					Targets:   parameters["targets"].([]string),
					Priority:  parameters["priority"].(string),
					Clauses: []nextgen.Clause{
						{
							Attribute: clauses[0].(map[string]interface{})["attribute"].(string),
							Negate:    clauses[0].(map[string]interface{})["negate"].(bool),
							Op:        clauses[0].(map[string]interface{})["op"].(string),
							Values:    clauses[0].(map[string]interface{})["values"].([]string),
						},
					},
					Serve: []Serve{
						{
							Variation: serve[0].(map[string]interface{})["variation"].(string),
							Distribution: []Distribution{
								{
									BucketBy:   serve[0].(map[string]interface{})["distribution"].([]interface{})[0].(map[string]interface{})["bucketBy"].(string),
									Variations: variations,
								},
							},
						},
					},
				},
			},
		}
		instructions = append(instructions, instruction)
	}

	opts.Instructions = instructions[0]

	// fix kind mapping, since use does need to know the mapping
	if kind, ok := KindMap[opts.Instructions.Kind]; ok {
		opts.Instructions.Kind = kind
	}

	return &nextgen.FeatureFlagsApiPatchFeatureOpts{
		Body:                  optional.NewInterface(opts),
		EnvironmentIdentifier: optional.EmptyString(),
	}
}

func buildFFReadOpts(d *schema.ResourceData) *nextgen.FeatureFlagsApiGetFeatureFlagOpts {

	return &nextgen.FeatureFlagsApiGetFeatureFlagOpts{
		EnvironmentIdentifier: optional.EmptyString(),
	}

}
