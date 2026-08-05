package lifecycle

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceLifecycleRuleSchema(readOnly bool) map[string]*schema.Schema {
	requiredOrComputed := func() (bool, bool) {
		if readOnly {
			return false, true
		}
		return true, false
	}
	_, computed := requiredOrComputed()

	attrRequired := func() *schema.Schema {
		s := &schema.Schema{Type: schema.TypeString}
		if readOnly {
			s.Computed = true
		} else {
			s.Required = true
		}
		return s
	}
	attrOptional := func() *schema.Schema {
		s := &schema.Schema{Type: schema.TypeString, Optional: true}
		if readOnly {
			s.Computed = true
			s.Optional = false
		}
		return s
	}

	_ = computed
	_ = attrRequired

	return map[string]*schema.Schema{
		"account_id": {
			Description: "Account identifier for the lifecycle rule.",
			Type:        schema.TypeString,
			Required:    !readOnly,
			Computed:    readOnly,
		},
		"org_id": {
			Description: "Organization identifier. Required for org-scoped rules.",
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    readOnly,
		},
		"project_id": {
			Description: "Project identifier. Required for project-scoped rules.",
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    readOnly,
		},
		"rule_id": {
			Description: "Unique ID of the lifecycle rule (returned by the API).",
			Type:        schema.TypeString,
			Computed:    true,
		},
		"name": {
			Description: "Name of the lifecycle rule.",
			Type:        schema.TypeString,
			Required:    !readOnly,
			Computed:    readOnly,
		},
		"description": attrOptional(),
		"action": {
			Description: "Action to perform: DELETE or PROTECT.",
			Type:        schema.TypeString,
			Required:    !readOnly,
			Computed:    readOnly,
			ValidateFunc: validation.StringInSlice([]string{
				"DELETE", "PROTECT",
			}, false),
		},
		"package_type": {
			Description: "Package type the rule applies to (e.g. DOCKER, MAVEN, HELM).",
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    readOnly,
		},
		"enabled": {
			Description: "Whether the rule is enabled.",
			Type:        schema.TypeBool,
			Optional:    true,
			Computed:    true,
		},
		"apply_to": {
			Description: "Defines which registries this rule applies to.",
			Type:        schema.TypeList,
			Required:    !readOnly,
			Computed:    readOnly,
			MaxItems:    1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"mode": {
						Description: "Mode: ALL_IN_SCOPE or EXPLICIT.",
						Type:        schema.TypeString,
						Required:    !readOnly,
						Computed:    readOnly,
						ValidateFunc: validation.StringInSlice([]string{
							"ALL_IN_SCOPE", "EXPLICIT",
						}, false),
					},
					"registries": {
						Description: "List of registry identifiers (required when mode=EXPLICIT).",
						Type:        schema.TypeList,
						Optional:    true,
						Computed:    readOnly,
						Elem:        &schema.Schema{Type: schema.TypeString},
					},
				},
			},
		},
		"criteria": {
			Description: "Cleanup criteria for the lifecycle rule.",
			Type:        schema.TypeList,
			Optional:    true,
			Computed:    readOnly,
			MaxItems:    1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"match": {
						Description: "How criteria rules are combined: ALL or ANY.",
						Type:        schema.TypeString,
						Required:    !readOnly,
						Computed:    readOnly,
						ValidateFunc: validation.StringInSlice([]string{
							"ALL", "ANY",
						}, false),
					},
					"rules": {
						Description: "List of individual criteria rules.",
						Type:        schema.TypeList,
						Required:    !readOnly,
						Computed:    readOnly,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"type": {
									Description: "Criteria type: KEEP_LAST_N, AGE_BASED, or UNUSED_FOR.",
									Type:        schema.TypeString,
									Required:    !readOnly,
									Computed:    readOnly,
									ValidateFunc: validation.StringInSlice([]string{
										"KEEP_LAST_N", "AGE_BASED", "UNUSED_FOR",
									}, false),
								},
								"value": {
									Description: "Numeric value for the criteria (e.g. number of versions, number of days).",
									Type:        schema.TypeInt,
									Optional:    true,
									Computed:    readOnly,
								},
								"unit": {
									Description: "Time unit for age/unused-for criteria: DAYS, MONTHS, or YEARS.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    readOnly,
									ValidateFunc: validation.StringInSlice([]string{
										"DAYS", "MONTHS", "YEARS",
									}, false),
								},
							},
						},
					},
				},
			},
		},
		"schedule": {
			Description: "Cron schedule for automatic execution of the rule.",
			Type:        schema.TypeList,
			Optional:    true,
			Computed:    readOnly,
			MaxItems:    1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"expression": {
						Description: "Cron expression (e.g. '0 2 * * *').",
						Type:        schema.TypeString,
						Required:    !readOnly,
						Computed:    readOnly,
					},
					"timezone": {
						Description: "Timezone for the cron schedule (e.g. 'UTC', 'America/New_York').",
						Type:        schema.TypeString,
						Required:    !readOnly,
						Computed:    readOnly,
					},
				},
			},
		},
		"filter_config": {
			Description: "Package-type-specific filter configuration.",
			Type:        schema.TypeList,
			Optional:    true,
			Computed:    readOnly,
			MaxItems:    1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"package_type": {
						Description: "Package type this filter applies to (e.g. DOCKER, MAVEN, HUGGINGFACE).",
						Type:        schema.TypeString,
						Required:    !readOnly,
						Computed:    readOnly,
					},
					"package_name_allowed_pattern": {
						Description: "Glob patterns for package/image names to include.",
						Type:        schema.TypeList,
						Optional:    true,
						Computed:    readOnly,
						Elem:        &schema.Schema{Type: schema.TypeString},
					},
					"version_name_allowed_pattern": {
						Description: "Glob patterns for version/tag names to include.",
						Type:        schema.TypeList,
						Optional:    true,
						Computed:    readOnly,
						Elem:        &schema.Schema{Type: schema.TypeString},
					},
					"tag_name_allowed_pattern": {
						Description: "Glob patterns for Docker tag names to include (DOCKER only).",
						Type:        schema.TypeList,
						Optional:    true,
						Computed:    readOnly,
						Elem:        &schema.Schema{Type: schema.TypeString},
					},
					"group_id_allowed_pattern": {
						Description: "Glob patterns for Maven group IDs to include (MAVEN only).",
						Type:        schema.TypeList,
						Optional:    true,
						Computed:    readOnly,
						Elem:        &schema.Schema{Type: schema.TypeString},
					},
					"model_allowed_pattern": {
						Description: "Glob patterns for HuggingFace model names to include (HUGGINGFACE only).",
						Type:        schema.TypeList,
						Optional:    true,
						Computed:    readOnly,
						Elem:        &schema.Schema{Type: schema.TypeString},
					},
					"dataset_allowed_pattern": {
						Description: "Glob patterns for HuggingFace dataset names to include (HUGGINGFACE only).",
						Type:        schema.TypeList,
						Optional:    true,
						Computed:    readOnly,
						Elem:        &schema.Schema{Type: schema.TypeString},
					},
				},
			},
		},
		// Read-only computed fields
		"created_at": {
			Description: "Timestamp when the rule was created (milliseconds since epoch).",
			Type:        schema.TypeInt,
			Computed:    true,
		},
		"updated_at": {
			Description: "Timestamp when the rule was last updated (milliseconds since epoch).",
			Type:        schema.TypeInt,
			Computed:    true,
		},
		"last_run_at": {
			Description: "Timestamp of the last execution (milliseconds since epoch).",
			Type:        schema.TypeInt,
			Computed:    true,
		},
		"next_run_at": {
			Description: "Timestamp of the next scheduled execution (milliseconds since epoch).",
			Type:        schema.TypeInt,
			Computed:    true,
		},
	}
}
