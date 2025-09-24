package fme

import (
	"context"

	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func DataSourceFMESplitDefinition() *schema.Resource {
	return &schema.Resource{
		Description: "Data source for retrieving a FME (Feature Management Engine) split definition.",
		ReadContext: dataSourceFMESplitDefinitionRead,

		Schema: map[string]*schema.Schema{
			"workspace_id": {
				Description:  "Unique identifier of the workspace.",
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsNotEmpty,
			},
			"environment_id": {
				Description:  "Unique identifier of the environment.",
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsNotEmpty,
			},
			"split_name": {
				Description:  "Name of the split.",
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsNotEmpty,
			},
			"treatments": {
				Description: "List of treatments for the split definition.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Description: "Name of the treatment.",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"description": {
							Description: "Description of the treatment.",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"configurations": {
							Description: "List of configurations for the treatment.",
							Type:        schema.TypeList,
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Description: "Name of the configuration.",
										Type:        schema.TypeString,
										Computed:    true,
									},
									"value": {
										Description: "Value of the configuration.",
										Type:        schema.TypeString,
										Computed:    true,
									},
								},
							},
						},
					},
				},
			},
			"rules": {
				Description: "List of rules for the split definition.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"treatment": {
							Description: "Treatment for this rule.",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"size": {
							Description: "Size percentage for this rule.",
							Type:        schema.TypeInt,
							Computed:    true,
						},
					},
				},
			},
			"default_rule": {
				Description: "Default rule for the split definition.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"treatment": {
							Description: "Treatment for the default rule.",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"size": {
							Description: "Size percentage for the default rule.",
							Type:        schema.TypeInt,
							Computed:    true,
						},
					},
				},
			},
			"baseline_treatment": {
				Description: "Baseline treatment for the split definition.",
				Type:        schema.TypeString,
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

func dataSourceFMESplitDefinitionRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	session := meta.(*internal.Session)
	fmeClient := session.FMEClient

	if fmeClient == nil {
		return diag.Errorf("FME client not configured")
	}

	c := fmeClient.(*FMEConfig)
	workspaceID := d.Get("workspace_id").(string)
	environmentID := d.Get("environment_id").(string)
	splitName := d.Get("split_name").(string)

	splitDefinition, err := c.APIClient.SplitDefinitions.Get(workspaceID, environmentID, splitName)
	if err != nil {
		return diag.FromErr(err)
	}

	if splitDefinition == nil {
		return diag.Errorf("split definition for split %s not found in workspace %s environment %s", splitName, workspaceID, environmentID)
	}

	d.SetId(splitName)

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
		d.Set("treatments", treatments)
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
	d.Set("workspace_id", workspaceID)
	d.Set("environment_id", environmentID)
	d.Set("split_name", splitName)

	return nil
}