package feature_flag

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strings"

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
				ForceNew:    false,
			},
			"description": {
				Description: "Description of the Feature Flag",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
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
			"default_off_variation": {
				Description: "Which of the variations to use when the flag is toggled to off state",
				Type:        schema.TypeString,
				Required:    true,
			},
			"default_on_variation": {
				Description: "Which of the variations to use when the flag is toggled to on state",
				Type:        schema.TypeString,
				Required:    true,
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
			},
			"permanent": {
				Description: "Whether or not the flag is permanent. If it is, it will never be flagged as stale",
				Type:        schema.TypeBool,
				Required:    true,
			},
			"tags": {
				Description: "The tags for the flag",
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"identifier": {
							Description: "The identifier of the tag",
							Type:        schema.TypeString,
							Required:    true,
						},
					},
				},
			},
			"environment": {
				Description: "Environment Identifier",
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"identifier": {
							Description: "Identifier of the Environment",
							Type:        schema.TypeString,
							Required:    true,
						},
						"state": {
							Description: "State of the flag in this environment. Possible values are 'on' and 'off'",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"default_on_variation": {
							Description: "Default variation to be served when flag is 'on'",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"default_off_variation": {
							Description: "Default variation to be served when flag is 'off'",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"add_target_rule": {
							Description: "The targeting rules for the flag",
							Type:        schema.TypeList,
							Optional:    true,
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"variation": {
										Description: "The identifier of the variation. Valid values are `enabled`, `disabled`",
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
								},
							},
						},
					},
				},
			},
			"variation": {
				Description: "The options available for your flag",
				Type:        schema.TypeList,
				Required:    true,
				ForceNew:    false,
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
		},
	}

	return resource
}

type FFQueryParameters struct {
	Identifier     string
	OrganizationId string
	ProjectId      string
}

// TargetRules is the target rules for the feature flag
type TargetRules struct {
	Variation string   `json:"variation"`
	Targets   []string `json:"targets"`
}

// Tag is a tag for the feature flag
type Tag struct {
	Name       string `json:"name"`
	Identifier string `json:"identifier"`
}

type FFCreateOpts struct {
	Identifier          string              `json:"identifier"`
	Name                string              `json:"name"`
	Description         string              `json:"description,omitempty"`
	DefaultOffVariation string              `json:"defaultOffVariation"`
	DefaultOnVariation  string              `json:"defaultOnVariation"`
	Kind                string              `json:"kind"`
	Owner               string              `json:"owner,omitempty"`
	Permanent           bool                `json:"permanent"`
	Project             string              `json:"project"`
	Variations          []nextgen.Variation `json:"variations"`
}

type FFPutOpts struct {
	Identifier          string              `json:"identifier"`
	Name                string              `json:"name"`
	Description         string              `json:"description"`
	DefaultOffVariation string              `json:"defaultOffVariation"`
	DefaultOnVariation  string              `json:"defaultOnVariation"`
	Permanent           bool                `json:"permanent"`
	Variations          []nextgen.Variation `json:"variations"`
	Tags                []Tag               `json:"tags"`
	Environments        []Environment       `json:"environments,omitempty"`
}

type Environment struct {
	Identifier          string        `json:"identifier"`
	DefaultOnVariation  string        `json:"defaultOnVariation"`
	DefaultOffVariation string        `json:"defaultOffVariation"`
	State               string        `json:"state"`
	TargetRules         []TargetRules `json:"rules"`
}

type TFEnvironment struct {
	Identifier          string        `json:"identifier"`
	DefaultOnVariation  string        `json:"default_on_variation"`
	DefaultOffVariation string        `json:"default_off_variation"`
	State               string        `json:"state"`
	TargetRules         []TargetRules `json:"add_target_rule"`
}

func resourceFeatureFlagUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	id := d.Id()
	if id == "" {
		return nil
	}

	qp := buildFFQueryParameters(d)
	opts := buildFFPutOpts(d)

	if opts != nil {
		httpResp, err := c.FeatureFlagsApi.PutFeatureFlag(ctx, id, c.AccountId, qp.OrganizationId, qp.ProjectId, opts)

		if err != nil {
			return HandleCFApiError(err, d, httpResp)
		}
	}

	return resourceFeatureFlagRead(ctx, d, meta)

}

func resourceFeatureFlagRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	id := d.Id()
	if id == "" {
		d.MarkNewResource()
		return nil
	}

	qp := buildFFQueryParameters(d)

	var environmentIdentifiers []string

	for _, env := range getEnvironmentData(d) {
		environmentIdentifiers = append(environmentIdentifiers, env.Identifier)
	}
	if len(environmentIdentifiers) == 0 {
		environmentIdentifiers = []string{""}
	}

	// get flag for each env in a loop
	for _, env := range environmentIdentifiers {
		opts := buildFFReadOpts(d, env)
		resp, httpResp, err := c.FeatureFlagsApi.GetFeatureFlag(ctx, id, c.AccountId, qp.OrganizationId, qp.ProjectId, opts)

		if err != nil {
			return HandleCFApiError(err, d, httpResp)
		}

		readFeatureFlag(d, &resp, qp, env)
	}

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

	var err error
	var httpResp *http.Response

	httpResp, err = c.FeatureFlagsApi.CreateFeatureFlag(ctx, c.AccountId, qp.OrganizationId, opts)

	if err != nil {
		// handle conflict
		if httpResp != nil && httpResp.StatusCode == 409 {
			return diag.Errorf("A feature flag with identifier [%s] orgIdentifier [%s] project [%s] already exists", d.Get("identifier").(string), qp.OrganizationId, qp.ProjectId)
		}
		return HandleCFApiError(err, d, httpResp)
	}

	// make updates for anything that can't be configured with initial create request
	putOpts := buildFFPutOpts(d)

	if opts != nil {
		httpResp, err = c.FeatureFlagsApi.PutFeatureFlag(ctx, id, c.AccountId, qp.OrganizationId, qp.ProjectId, putOpts)

		if err != nil {
			return HandleCFApiError(err, d, httpResp)
		}
	}

	d.SetId(id)

	return resourceFeatureFlagRead(ctx, d, meta)
}

func resourceFeatureFlagDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	id := d.Id()
	if id == "" {
		return nil
	}
	qp := buildFFQueryParameters(d)

	httpResp, err := c.FeatureFlagsApi.DeleteFeatureFlag(ctx, d.Id(), c.AccountId, qp.OrganizationId, qp.ProjectId, &nextgen.FeatureFlagsApiDeleteFeatureFlagOpts{CommitMsg: optional.EmptyString(), ForceDelete: optional.NewBool(true)})
	if err != nil {
		return HandleCFApiError(err, d, httpResp)
	}

	return nil
}

func readFeatureFlag(d *schema.ResourceData, flag *nextgen.Feature, qp *FFQueryParameters, env string) {
	d.SetId(flag.Identifier)
	d.Set("identifier", flag.Identifier)
	d.Set("name", flag.Name)
	d.Set("description", flag.Description)
	d.Set("project_id", flag.Project)
	d.Set("default_on_variation", flag.DefaultOnVariation)
	d.Set("default_off_variation", flag.DefaultOffVariation)
	d.Set("description", flag.Description)
	d.Set("kind", flag.Kind)
	d.Set("permanent", flag.Permanent)
	d.Set("owner", strings.Join(flag.Owner, ","))
	d.Set("org_id", qp.OrganizationId)
	d.Set("tags", expandTags(flag.Tags))
	d.Set("variation", expandVariations(flag.Variations))
	// update environment field
	if flag.EnvProperties != nil && env != "" {
		var targetRules []TargetRules
		for _, rule := range flag.EnvProperties.VariationMap {
			var targets []string
			for _, t := range rule.Targets {
				targets = append(targets, t.Identifier)
			}
			targetRules = append(targetRules, TargetRules{
				Variation: rule.Variation,
				Targets:   targets,
			})
		}
		updatedEnv := Environment{
			Identifier:          flag.EnvProperties.Environment,
			DefaultOnVariation:  flag.EnvProperties.DefaultServe.Variation,
			DefaultOffVariation: flag.EnvProperties.OffVariation,
			State:               string(*flag.EnvProperties.State),
			TargetRules:         targetRules,
		}

		// add environment data as an upsert - check array and if environment already exists replace it - if not add it to array
		environments := getEnvironmentData(d)
		var updated bool
		for i, env := range environments {
			if env.Identifier == updatedEnv.Identifier {
				environments[i] = updatedEnv
				updated = true
			}
		}
		// if env is new then add to array
		if !updated {
			environments = append(environments, updatedEnv)
		}
		d.Set("environment", expandEnvironments(environments))
	}

}

func getEnvironmentData(d *schema.ResourceData) []Environment {
	var tfEnvironments []TFEnvironment
	var environments []Environment
	// get environment key
	if envData, ok := d.Get("environment").([]interface{}); ok {
		// marshal to bytes
		jsonData, err := json.Marshal(envData)
		if err != nil {
			log.Printf("Error marshalling environment data: %s", err)
			return environments
		}

		// unmarshal into environment array
		err = json.Unmarshal(jsonData, &tfEnvironments)
		if err != nil {
			log.Printf("Error unmarshalling environment data: %s", err)
			return environments
		}
	}

	for _, tfEnv := range tfEnvironments {
		environments = append(environments, Environment{
			Identifier:          tfEnv.Identifier,
			DefaultOnVariation:  tfEnv.DefaultOnVariation,
			DefaultOffVariation: tfEnv.DefaultOffVariation,
			State:               tfEnv.State,
			TargetRules:         tfEnv.TargetRules,
		})
	}

	return environments
}

func expandEnvironments(environments []Environment) []interface{} {
	var result []interface{}
	for _, env := range environments {
		var targetRules []interface{}
		for _, rule := range env.TargetRules {
			targetRules = append(targetRules, map[string]interface{}{
				"variation": rule.Variation,
				"targets":   rule.Targets,
			})
		}
		result = append(result, map[string]interface{}{
			"identifier":            env.Identifier,
			"state":                 env.State,
			"default_on_variation":  env.DefaultOnVariation,
			"default_off_variation": env.DefaultOffVariation,
			"add_target_rule":       targetRules,
		})
	}

	return result
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

func expandTags(tags []nextgen.Tag) []interface{} {
	var result []interface{}
	for _, tag := range tags {
		result = append(result, map[string]interface{}{
			"identifier": tag.Name,
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
	opts := &FFCreateOpts{
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

func buildFFPutOpts(d *schema.ResourceData) *nextgen.FeatureFlagsApiPutFeatureFlagOpts {
	opts := &FFPutOpts{
		Identifier:          d.Get("identifier").(string),
		Name:                d.Get("name").(string),
		DefaultOffVariation: d.Get("default_off_variation").(string),
		DefaultOnVariation:  d.Get("default_on_variation").(string),
	}

	var description string
	if desc, ok := d.GetOk("description"); ok {
		description = desc.(string)
	}
	opts.Description = description

	if permanent, ok := d.GetOk("permanent"); ok {
		opts.Permanent = permanent.(bool)
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

	tags := []Tag{}
	if tagsData, ok := d.GetOk("tags"); ok {
		tagsData := tagsData.([]interface{})

		for _, tagData := range tagsData {
			tMap := tagData.(map[string]interface{})
			tag := Tag{
				Name:       tMap["identifier"].(string),
				Identifier: tMap["identifier"].(string),
			}
			tags = append(tags, tag)
		}
	}
	opts.Tags = tags

	// add environment level attributes
	opts.Environments = getEnvironmentData(d)

	log.Println("flag update request body")
	jsonData, _ := json.Marshal(opts)
	log.Println(string(jsonData))
	return &nextgen.FeatureFlagsApiPutFeatureFlagOpts{
		Body: optional.NewInterface(opts),
	}
}

func buildFFReadOpts(d *schema.ResourceData, env string) *nextgen.FeatureFlagsApiGetFeatureFlagOpts {

	return &nextgen.FeatureFlagsApiGetFeatureFlagOpts{
		EnvironmentIdentifier: optional.NewString(env),
	}

}

// HandleCFApiError - parses the error as type cfError and returns the error message
// if it can't parse as cf error it falls back to the generic error handling helper
func HandleCFApiError(err error, d *schema.ResourceData, httpResp *http.Response) diag.Diagnostics {
	erro, ok := err.(nextgen.GenericSwaggerError)
	if ok {
		cfError, ok := erro.Model().(nextgen.CfError)
		if ok {
			return diag.Errorf(cfError.Message)
		}
	}
	return helpers.HandleApiError(err, d, httpResp)
}
