package fme

import (
	"context"
	"fmt"
	"strings"

	"github.com/harness/terraform-provider-harness/internal"
	"github.com/harness/terraform-provider-harness/internal/service/platform/fme/api"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func ResourceFMESplitDefinition() *schema.Resource {
	return &schema.Resource{
		Description:   "Resource for creating a FME split definition (feature flag configuration).",
		ReadContext:   resourceFMESplitDefinitionRead,
		CreateContext: resourceFMESplitDefinitionCreate,
		UpdateContext: resourceFMESplitDefinitionUpdate,
		DeleteContext: resourceFMESplitDefinitionDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceFMESplitDefinitionImport,
		},

		Schema: map[string]*schema.Schema{
			"workspace_id": {
				Description:  "ID of the workspace this split definition belongs to.",
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringIsNotEmpty,
			},
			"environment_id": {
				Description:  "ID of the environment this split definition belongs to.",
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringIsNotEmpty,
			},
			"split_name": {
				Description:  "Name of the split this definition belongs to.",
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringIsNotEmpty,
			},
			"treatment": {
				Description: "List of treatments for the split definition.",
				Type:        schema.TypeList,
				Required:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Description:  "Name of the treatment.",
							Type:         schema.TypeString,
							Required:     true,
							ValidateFunc: validation.StringIsNotEmpty,
						},
						"description": {
							Description: "Description of the treatment.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"configurations": {
							Description: "JSON string of configurations for the treatment.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"keys": {
							Description: "Set of keys for the treatment.",
							Type:        schema.TypeSet,
							Optional:    true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"segments": {
							Description: "Set of segments for the treatment.",
							Type:        schema.TypeSet,
							Optional:    true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
			"rule": {
				Description: "List of rules for the split definition.",
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"bucket": {
							Description: "List of bucket configurations for this rule.",
							Type:        schema.TypeList,
							Required:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"treatment": {
										Description:  "Treatment for this bucket.",
										Type:         schema.TypeString,
										Required:     true,
										ValidateFunc: validation.StringIsNotEmpty,
									},
									"size": {
										Description:  "Size percentage for this bucket (0-100).",
										Type:         schema.TypeInt,
										Required:     true,
										ValidateFunc: validation.IntBetween(0, 100),
									},
								},
							},
						},
						"condition": {
							Description: "Condition configuration for this rule.",
							Type:        schema.TypeList,
							Optional:    true,
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"combiner": {
										Description: "Combiner type for the condition (AND/OR).",
										Type:        schema.TypeString,
										Required:    true,
									},
									"matcher": {
										Description: "List of matchers for the condition.",
										Type:        schema.TypeList,
										Required:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"type": {
													Description: "Type of the matcher.",
													Type:        schema.TypeString,
													Required:    true,
												},
												"attribute": {
													Description: "Attribute to match on.",
													Type:        schema.TypeString,
													Required:    true,
												},
												"strings": {
													Description: "List of strings to match.",
													Type:        schema.TypeList,
													Optional:    true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"negate": {
													Description: "Whether to negate the match.",
													Type:        schema.TypeBool,
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
			"default_rule": {
				Description: "Default rule for the split definition.",
				Type:        schema.TypeList,
				Required:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"treatment": {
							Description:  "Treatment for the default rule.",
							Type:         schema.TypeString,
							Required:     true,
							ValidateFunc: validation.StringIsNotEmpty,
						},
						"size": {
							Description:  "Size percentage for the default rule (0-100).",
							Type:         schema.TypeInt,
							Required:     true,
							ValidateFunc: validation.IntBetween(0, 100),
						},
					},
				},
			},
			"baseline_treatment": {
				Description:  "Baseline treatment for the split definition.",
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringIsNotEmpty,
			},
			"default_treatment": {
				Description:  "Default treatment for the split definition.",
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsNotEmpty,
			},
			"traffic_allocation": {
				Description: "Traffic allocation percentage for the split definition.",
				Type:        schema.TypeInt,
				Computed:    true,
			},
			"last_update_time": {
				Description: "Last update time of the split definition.",
				Type:        schema.TypeInt,
				Computed:    true,
			},
			"creation_time": {
				Description: "Creation time of the split definition.",
				Type:        schema.TypeInt,
				Computed:    true,
			},
		},
	}
}

func resourceFMESplitDefinitionCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	session := meta.(*internal.Session)
	fmeClient := session.FMEClient

	if fmeClient == nil {
		return diag.Errorf("FME client not configured")
	}

	c := fmeClient.(*FMEConfig)

	workspaceID := d.Get("workspace_id").(string)
	environmentID := d.Get("environment_id").(string)
	splitName := d.Get("split_name").(string)

	req := buildSplitDefinitionCreateRequest(d)

	_, err := c.APIClient.SplitDefinitions.Create(workspaceID, environmentID, splitName, req)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(splitName)
	return resourceFMESplitDefinitionRead(ctx, d, meta)
}

func resourceFMESplitDefinitionRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	session := meta.(*internal.Session)
	fmeClient := session.FMEClient

	if fmeClient == nil {
		return diag.Errorf("FME client not configured")
	}

	c := fmeClient.(*FMEConfig)
	workspaceID := d.Get("workspace_id").(string)
	environmentID := d.Get("environment_id").(string)
	splitName := d.Id()

	splitDefinition, err := c.APIClient.SplitDefinitions.Get(workspaceID, environmentID, splitName)
	if err != nil {
		return diag.FromErr(err)
	}

	if splitDefinition == nil {
		d.SetId("")
		return nil
	}

	// Set treatments
	if splitDefinition.Treatments != nil {
		treatments := make([]interface{}, len(splitDefinition.Treatments))
		for i, treatment := range splitDefinition.Treatments {
			treatmentMap := map[string]interface{}{}
			if treatment.Name != nil {
				treatmentMap["name"] = *treatment.Name
			}
			if treatment.Description != nil {
				treatmentMap["description"] = *treatment.Description
			}

			// Set configurations as JSON string
			if treatment.Configurations != nil {
				treatmentMap["configurations"] = *treatment.Configurations
			}

			// Set keys
			if treatment.Keys != nil {
				treatmentMap["keys"] = treatment.Keys
			}

			// Set segments
			if treatment.Segments != nil {
				treatmentMap["segments"] = treatment.Segments
			}

			treatments[i] = treatmentMap
		}
		d.Set("treatment", treatments)
	}

	// Set rules
	if splitDefinition.Rules != nil {
		rules := make([]interface{}, len(splitDefinition.Rules))
		for i, rule := range splitDefinition.Rules {
			ruleMap := map[string]interface{}{}

			// Handle buckets
			if rule.Buckets != nil {
				buckets := make([]interface{}, len(rule.Buckets))
				for j, bucket := range rule.Buckets {
					bucketMap := map[string]interface{}{}
					if bucket.Treatment != nil {
						bucketMap["treatment"] = *bucket.Treatment
					}
					if bucket.Size != nil {
						bucketMap["size"] = *bucket.Size
					}
					buckets[j] = bucketMap
				}
				ruleMap["bucket"] = buckets
			}

			// Handle condition
			if rule.Condition != nil {
				conditionMap := map[string]interface{}{}
				if rule.Condition.Combiner != nil {
					conditionMap["combiner"] = *rule.Condition.Combiner
				}

				// Handle matchers
				if rule.Condition.Matchers != nil {
					matchers := make([]interface{}, len(rule.Condition.Matchers))
					for k, matcher := range rule.Condition.Matchers {
						matcherMap := map[string]interface{}{}
						if matcher.Type != nil {
							matcherMap["type"] = *matcher.Type
						}
						if matcher.Attribute != nil {
							matcherMap["attribute"] = *matcher.Attribute
						}
						if matcher.Strings != nil {
							matcherMap["strings"] = matcher.Strings
						}
						if matcher.Negate != nil {
							matcherMap["negate"] = *matcher.Negate
						}
						matchers[k] = matcherMap
					}
					conditionMap["matcher"] = matchers
				}

				ruleMap["condition"] = []interface{}{conditionMap}
			}

			rules[i] = ruleMap
		}
		d.Set("rule", rules)
	}

	// Set default rule
	if splitDefinition.DefaultRule != nil {
		defaultRules := make([]interface{}, len(splitDefinition.DefaultRule))
		for i, rule := range splitDefinition.DefaultRule {
			ruleMap := map[string]interface{}{}
			if rule.Treatment != nil {
				ruleMap["treatment"] = *rule.Treatment
			}
			if rule.Size != nil {
				ruleMap["size"] = *rule.Size
			}
			defaultRules[i] = ruleMap
		}
		d.Set("default_rule", defaultRules)
	}

	d.Set("baseline_treatment", splitDefinition.BaselineTreatment)
	d.Set("last_update_time", splitDefinition.LastUpdateTime)
	d.Set("creation_time", splitDefinition.CreationTime)

	return nil
}

func resourceFMESplitDefinitionUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	session := meta.(*internal.Session)
	fmeClient := session.FMEClient

	if fmeClient == nil {
		return diag.Errorf("FME client not configured")
	}

	c := fmeClient.(*FMEConfig)

	workspaceID := d.Get("workspace_id").(string)
	environmentID := d.Get("environment_id").(string)
	splitName := d.Id()

	req := buildSplitDefinitionUpdateRequest(d)

	_, err := c.APIClient.SplitDefinitions.Update(workspaceID, environmentID, splitName, req)
	if err != nil {
		return diag.FromErr(err)
	}

	return resourceFMESplitDefinitionRead(ctx, d, meta)
}

func resourceFMESplitDefinitionDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	session := meta.(*internal.Session)
	fmeClient := session.FMEClient

	if fmeClient == nil {
		return diag.Errorf("FME client not configured")
	}

	c := fmeClient.(*FMEConfig)
	workspaceID := d.Get("workspace_id").(string)
	environmentID := d.Get("environment_id").(string)
	splitName := d.Id()

	err := c.APIClient.SplitDefinitions.Delete(workspaceID, environmentID, splitName)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func buildSplitDefinitionCreateRequest(d *schema.ResourceData) *api.SplitDefinitionCreateRequest {
	req := &api.SplitDefinitionCreateRequest{}

	// Build treatments
	if treatmentsData, ok := d.GetOk("treatment"); ok {
		treatments := treatmentsData.([]interface{})
		req.Treatments = make([]api.SplitDefinitionTreatment, len(treatments))

		for i, treatmentData := range treatments {
			treatment := treatmentData.(map[string]interface{})
			reqTreatment := api.SplitDefinitionTreatment{}

			if name, ok := treatment["name"].(string); ok {
				reqTreatment.Name = &name
			}
			if description, ok := treatment["description"].(string); ok {
				reqTreatment.Description = &description
			}

			// Handle configurations as JSON string
			if configsJSON, ok := treatment["configurations"].(string); ok && configsJSON != "" {
				reqTreatment.Configurations = &configsJSON
			}

			// Handle keys
			if keysData, ok := treatment["keys"].(*schema.Set); ok {
				keys := make([]string, 0, keysData.Len())
				for _, key := range keysData.List() {
					keys = append(keys, key.(string))
				}
				reqTreatment.Keys = keys
			}

			// Handle segments
			if segmentsData, ok := treatment["segments"].(*schema.Set); ok {
				segments := make([]string, 0, segmentsData.Len())
				for _, segment := range segmentsData.List() {
					segments = append(segments, segment.(string))
				}
				reqTreatment.Segments = segments
			}

			req.Treatments[i] = reqTreatment
		}
	}

	// Build rules with bucket and condition structure
	if rulesData, ok := d.GetOk("rule"); ok {
		rules := rulesData.([]interface{})
		req.Rules = make([]api.SplitDefinitionRule, len(rules))

		for i, ruleData := range rules {
			rule := ruleData.(map[string]interface{})
			reqRule := api.SplitDefinitionRule{}

			// Handle bucket structure
			if bucketsData, ok := rule["bucket"].([]interface{}); ok {
				reqRule.Buckets = make([]api.SplitDefinitionBucket, len(bucketsData))
				for j, bucketData := range bucketsData {
					bucket := bucketData.(map[string]interface{})
					reqBucket := api.SplitDefinitionBucket{}

					if treatment, ok := bucket["treatment"].(string); ok {
						reqBucket.Treatment = &treatment
					}
					if size, ok := bucket["size"].(int); ok {
						reqBucket.Size = &size
					}

					reqRule.Buckets[j] = reqBucket
				}
			}

			// Handle condition structure
			if conditionsData, ok := rule["condition"].([]interface{}); ok && len(conditionsData) > 0 {
				condition := conditionsData[0].(map[string]interface{})
				reqRule.Condition = &api.SplitDefinitionCondition{}

				if combiner, ok := condition["combiner"].(string); ok {
					reqRule.Condition.Combiner = &combiner
				}

				// Handle matchers
				if matchersData, ok := condition["matcher"].([]interface{}); ok {
					reqRule.Condition.Matchers = make([]api.SplitDefinitionMatcher, len(matchersData))
					for j, matcherData := range matchersData {
						matcher := matcherData.(map[string]interface{})
						reqMatcher := api.SplitDefinitionMatcher{}

						if matcherType, ok := matcher["type"].(string); ok {
							reqMatcher.Type = &matcherType
						}
						if attribute, ok := matcher["attribute"].(string); ok {
							reqMatcher.Attribute = &attribute
						}
						if stringsData, ok := matcher["strings"].([]interface{}); ok && len(stringsData) > 0 {
							strings := make([]string, len(stringsData))
							for k, s := range stringsData {
								strings[k] = s.(string)
							}
							reqMatcher.Strings = strings
						}
						if negate, ok := matcher["negate"].(bool); ok && negate {
							// Only include negate if it's true
							reqMatcher.Negate = &negate
						}

						reqRule.Condition.Matchers[j] = reqMatcher
					}
				}
			}

			req.Rules[i] = reqRule
		}
	}

	// Build default rule (as array)
	if defaultRuleData, ok := d.GetOk("default_rule"); ok {
		defaultRules := defaultRuleData.([]interface{})
		req.DefaultRule = make([]api.SplitDefinitionDefaultRule, len(defaultRules))

		for i, defaultRuleItem := range defaultRules {
			defaultRule := defaultRuleItem.(map[string]interface{})
			reqDefaultRule := api.SplitDefinitionDefaultRule{}

			if treatment, ok := defaultRule["treatment"].(string); ok {
				reqDefaultRule.Treatment = &treatment
			}
			if size, ok := defaultRule["size"].(int); ok {
				reqDefaultRule.Size = &size
			}

			req.DefaultRule[i] = reqDefaultRule
		}
	}

	// Set baseline treatment (required field - use default treatment if not specified)
	if baselineTreatment, ok := d.GetOk("baseline_treatment"); ok {
		baselineTreatmentStr := baselineTreatment.(string)
		req.BaselineTreatment = &baselineTreatmentStr
	} else if defaultTreatment, ok := d.GetOk("default_treatment"); ok {
		// Use default treatment as baseline if baseline not specified
		defaultTreatmentStr := defaultTreatment.(string)
		req.BaselineTreatment = &defaultTreatmentStr
	}

	// Set default treatment (required field)
	if defaultTreatment, ok := d.GetOk("default_treatment"); ok {
		defaultTreatmentStr := defaultTreatment.(string)
		req.DefaultTreatment = &defaultTreatmentStr
	}

	// Set traffic allocation (default to 100 if not specified)
	if trafficAllocation, ok := d.GetOk("traffic_allocation"); ok {
		trafficAllocationInt := trafficAllocation.(int)
		req.TrafficAllocation = &trafficAllocationInt
	} else {
		defaultTrafficAllocation := 100
		req.TrafficAllocation = &defaultTrafficAllocation
	}

	return req
}

func buildSplitDefinitionUpdateRequest(d *schema.ResourceData) *api.SplitDefinitionUpdateRequest {
	req := &api.SplitDefinitionUpdateRequest{}

	// Build treatments
	if treatmentsData, ok := d.GetOk("treatment"); ok {
		treatments := treatmentsData.([]interface{})
		req.Treatments = make([]api.SplitDefinitionTreatment, len(treatments))

		for i, treatmentData := range treatments {
			treatment := treatmentData.(map[string]interface{})
			reqTreatment := api.SplitDefinitionTreatment{}

			if name, ok := treatment["name"].(string); ok {
				reqTreatment.Name = &name
			}
			if description, ok := treatment["description"].(string); ok {
				reqTreatment.Description = &description
			}

			// Handle configurations as JSON string
			if configsJSON, ok := treatment["configurations"].(string); ok && configsJSON != "" {
				reqTreatment.Configurations = &configsJSON
			}

			// Handle keys
			if keysData, ok := treatment["keys"].(*schema.Set); ok {
				keys := make([]string, 0, keysData.Len())
				for _, key := range keysData.List() {
					keys = append(keys, key.(string))
				}
				reqTreatment.Keys = keys
			}

			// Handle segments
			if segmentsData, ok := treatment["segments"].(*schema.Set); ok {
				segments := make([]string, 0, segmentsData.Len())
				for _, segment := range segmentsData.List() {
					segments = append(segments, segment.(string))
				}
				reqTreatment.Segments = segments
			}

			req.Treatments[i] = reqTreatment
		}
	}

	// Build rules with bucket and condition structure
	if rulesData, ok := d.GetOk("rule"); ok {
		rules := rulesData.([]interface{})
		req.Rules = make([]api.SplitDefinitionRule, len(rules))

		for i, ruleData := range rules {
			rule := ruleData.(map[string]interface{})
			reqRule := api.SplitDefinitionRule{}

			// Handle bucket structure
			if bucketsData, ok := rule["bucket"].([]interface{}); ok {
				reqRule.Buckets = make([]api.SplitDefinitionBucket, len(bucketsData))
				for j, bucketData := range bucketsData {
					bucket := bucketData.(map[string]interface{})
					reqBucket := api.SplitDefinitionBucket{}

					if treatment, ok := bucket["treatment"].(string); ok {
						reqBucket.Treatment = &treatment
					}
					if size, ok := bucket["size"].(int); ok {
						reqBucket.Size = &size
					}

					reqRule.Buckets[j] = reqBucket
				}
			}

			// Handle condition structure
			if conditionsData, ok := rule["condition"].([]interface{}); ok && len(conditionsData) > 0 {
				condition := conditionsData[0].(map[string]interface{})
				reqRule.Condition = &api.SplitDefinitionCondition{}

				if combiner, ok := condition["combiner"].(string); ok {
					reqRule.Condition.Combiner = &combiner
				}

				// Handle matchers
				if matchersData, ok := condition["matcher"].([]interface{}); ok {
					reqRule.Condition.Matchers = make([]api.SplitDefinitionMatcher, len(matchersData))
					for j, matcherData := range matchersData {
						matcher := matcherData.(map[string]interface{})
						reqMatcher := api.SplitDefinitionMatcher{}

						if matcherType, ok := matcher["type"].(string); ok {
							reqMatcher.Type = &matcherType
						}
						if attribute, ok := matcher["attribute"].(string); ok {
							reqMatcher.Attribute = &attribute
						}
						if stringsData, ok := matcher["strings"].([]interface{}); ok && len(stringsData) > 0 {
							strings := make([]string, len(stringsData))
							for k, s := range stringsData {
								strings[k] = s.(string)
							}
							reqMatcher.Strings = strings
						}
						if negate, ok := matcher["negate"].(bool); ok && negate {
							// Only include negate if it's true
							reqMatcher.Negate = &negate
						}

						reqRule.Condition.Matchers[j] = reqMatcher
					}
				}
			}

			req.Rules[i] = reqRule
		}
	}

	// Build default rule (as array)
	if defaultRuleData, ok := d.GetOk("default_rule"); ok {
		defaultRules := defaultRuleData.([]interface{})
		req.DefaultRule = make([]api.SplitDefinitionDefaultRule, len(defaultRules))

		for i, defaultRuleItem := range defaultRules {
			defaultRule := defaultRuleItem.(map[string]interface{})
			reqDefaultRule := api.SplitDefinitionDefaultRule{}

			if treatment, ok := defaultRule["treatment"].(string); ok {
				reqDefaultRule.Treatment = &treatment
			}
			if size, ok := defaultRule["size"].(int); ok {
				reqDefaultRule.Size = &size
			}

			req.DefaultRule[i] = reqDefaultRule
		}
	}

	// Set baseline treatment (required field - use default treatment if not specified)
	if baselineTreatment, ok := d.GetOk("baseline_treatment"); ok {
		baselineTreatmentStr := baselineTreatment.(string)
		req.BaselineTreatment = &baselineTreatmentStr
	} else if defaultTreatment, ok := d.GetOk("default_treatment"); ok {
		// Use default treatment as baseline if baseline not specified
		defaultTreatmentStr := defaultTreatment.(string)
		req.BaselineTreatment = &defaultTreatmentStr
	}

	// Set default treatment (required field)
	if defaultTreatment, ok := d.GetOk("default_treatment"); ok {
		defaultTreatmentStr := defaultTreatment.(string)
		req.DefaultTreatment = &defaultTreatmentStr
	}

	// Set traffic allocation (default to 100 if not specified)
	if trafficAllocation, ok := d.GetOk("traffic_allocation"); ok {
		trafficAllocationInt := trafficAllocation.(int)
		req.TrafficAllocation = &trafficAllocationInt
	} else {
		defaultTrafficAllocation := 100
		req.TrafficAllocation = &defaultTrafficAllocation
	}

	return req
}
// resourceFMESplitDefinitionImport handles importing of existing split definitions
func resourceFMESplitDefinitionImport(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	// Expected format: workspace_id:split_name:environment_id
	parts := strings.Split(d.Id(), ":")
	if len(parts) != 3 {
		return nil, fmt.Errorf("invalid import ID format, expected workspace_id:split_name:environment_id, got: %s", d.Id())
	}

	workspaceID := parts[0]
	splitName := parts[1]
	environmentID := parts[2]

	// Set the individual fields
	d.Set("workspace_id", workspaceID)
	d.Set("split_name", splitName)
	d.Set("environment_id", environmentID)

	// Set the ID to the split name (matching our create behavior)
	d.SetId(splitName)

	// Trigger a read to populate the rest of the state
	diags := resourceFMESplitDefinitionRead(ctx, d, meta)
	if diags.HasError() {
		return nil, fmt.Errorf("failed to read split definition during import: %v", diags)
	}

	return []*schema.ResourceData{d}, nil
}
