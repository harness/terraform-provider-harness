package loadtest

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

var loadTestIdentifierPattern = regexp.MustCompile(`^[A-Za-z0-9_-]+$`)

func resourceChaosLoadTestSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"org_id": {
			Description:  "Organization identifier of the load test. Changing it removes this load test. Recorded runs stay stored under the old unique id and are not listed on a recreated test.",
			Type:         schema.TypeString,
			Required:     true,
			ForceNew:     true,
			ValidateFunc: validation.StringIsNotEmpty,
		},
		"project_id": {
			Description:  "Project identifier of the load test. Changing it removes this load test. Recorded runs stay stored under the old unique id and are not listed on a recreated test.",
			Type:         schema.TypeString,
			Required:     true,
			ForceNew:     true,
			ValidateFunc: validation.StringIsNotEmpty,
		},
		"identifier": {
			Description:  "Identifier of the load test. Letters, digits, underscores, and hyphens only. `?`, `#`, spaces, and `/` are rejected so the provider cannot read or delete a different test. Changing it removes this load test. Recorded runs stay stored under the old unique id and are not listed on a recreated test.",
			Type:         schema.TypeString,
			Required:     true,
			ForceNew:     true,
			ValidateFunc: validateLoadTestIdentifier,
		},
		"name": {
			Description:  "Name of the load test.",
			Type:         schema.TypeString,
			Required:     true,
			ValidateFunc: validation.StringIsNotEmpty,
		},
		"description": {
			Description:      "Description of the load test. The API keeps the stored description when this is empty, and that empty change is not planned.",
			Type:             schema.TypeString,
			Optional:         true,
			DiffSuppressFunc: suppressEmptyDescription,
		},
		"tool_type": {
			Description:  "Load generator. `Locust` or `K6`. JMeter is not supported: Load Test Manager rewrites the plan, its variables, and the script on save. K6 requires `target_type` `kubernetes`. Locust may omit it and keeps the server default, `linux`. Changing this removes this load test. Recorded runs stay stored under the old unique id and are not listed on a recreated test.",
			Type:         schema.TypeString,
			Required:     true,
			ForceNew:     true,
			ValidateFunc: validation.StringInSlice([]string{"Locust", "K6"}, false),
		},
		"target_type": {
			Description:  "Infrastructure target. `linux` or `kubernetes`. Required to be `kubernetes` for K6 and managed infrastructure. Locust may omit it; the API stores `linux` and that value stays in state.",
			Type:         schema.TypeString,
			Optional:     true,
			Computed:     true,
			ValidateFunc: validateTargetTypeValue,
		},
		"managed_infra": {
			Description:   "When true, Harness provides Kubernetes infrastructure for every run. `environment_id` and `infrastructure_id` must be omitted.",
			Type:          schema.TypeBool,
			Optional:      true,
			Default:       false,
			ConflictsWith: []string{"environment_id", "infrastructure_id", "service_references"},
		},
		"environment_id": {
			Description:  "Chaos environment identifier. Required with `infrastructure_id` unless `managed_infra` is true.",
			Type:         schema.TypeString,
			Optional:     true,
			RequiredWith: []string{"infrastructure_id"},
		},
		"infrastructure_id": {
			Description:  "Chaos infrastructure identifier. Required with `environment_id` unless `managed_infra` is true.",
			Type:         schema.TypeString,
			Optional:     true,
			RequiredWith: []string{"environment_id"},
		},
		"service_references": {
			Description:   "Chaos service identities linked to this load test. The API stores them only when CHAOS_RT_SERVICES_ENABLED is on, and ignores them for managed infrastructure. An empty set cannot clear stored links. A read-back that does not match this set fails.",
			Type:          schema.TypeSet,
			Optional:      true,
			ConflictsWith: []string{"managed_infra"},
			Elem:          &schema.Schema{Type: schema.TypeString},
		},
		"tags": {
			Description: "Tags for the load test. An empty list clears the stored tags.",
			Type:        schema.TypeSet,
			Optional:    true,
			Elem:        &schema.Schema{Type: schema.TypeString},
		},
		"cleanup_policy": {
			Description:  "What happens to run pods, config maps, and secrets after a run. `delete` or `retain`.",
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "delete",
			ValidateFunc: validation.StringInSlice([]string{"delete", "retain"}, false),
		},
		"max_duration_sec": {
			Description: "Hard cap, in seconds, for any run of this load test. Set to 0 to clear it.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"tool_config": {
			Description:      "JSON object stored as the load test toolConfig. K6 fields sit at the top of the object or under `k6`. Locust is nested under `locust`. `mode` `script` requires `script.content` (base64). `mode` `image` requires `script.image`. k6 `ui` mode is not supported. Tunables (`hostUrl`, `iterations`, `targetUsers`, `durationSeconds`), `envVars`, `thresholds`, `options`, and k6 `metricTags` belong in this object. A `variables` key is rejected; set custom variables on `variable`. The script is stored in Terraform state.",
			Type:             schema.TypeString,
			Required:         true,
			ValidateFunc:     validateToolConfig,
			DiffSuppressFunc: suppressToolConfigDiff,
		},
		"variable": {
			Description: "Custom variables stored on the load test. An empty list clears them.",
			Type:        schema.TypeList,
			Optional:    true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"name": {
						Description:  "Variable name.",
						Type:         schema.TypeString,
						Required:     true,
						ValidateFunc: validation.StringIsNotEmpty,
					},
					"type": {
						Description:  "Variable type. `String` or `Number`.",
						Type:         schema.TypeString,
						Required:     true,
						ValidateFunc: validation.StringInSlice([]string{"String", "Number"}, false),
					},
					"value": {
						Description: "Variable value. A Number literal is sent as a JSON number. `<+input>` and any other `<+...>` expression stay strings, as do String values.",
						Type:        schema.TypeString,
						Required:    true,
					},
					"description": {
						Description: "Variable description.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"required": {
						Description: "Whether a run must supply this variable.",
						Type:        schema.TypeBool,
						Optional:    true,
					},
				},
			},
		},
		"resources": {
			Description: "Per-pod CPU and memory for Kubernetes load test workloads. Remove the block to clear the stored requirements.",
			Type:        schema.TypeList,
			Optional:    true,
			MaxItems:    1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"requests": {
						Description: "Resource requests, for example cpu `500m` and memory `512Mi`.",
						Type:        schema.TypeMap,
						Optional:    true,
						Elem:        &schema.Schema{Type: schema.TypeString},
					},
					"limits": {
						Description: "Resource limits, for example cpu `1` and memory `1Gi`.",
						Type:        schema.TypeMap,
						Optional:    true,
						Elem:        &schema.Schema{Type: schema.TypeString},
					},
				},
			},
		},
		"unique_id": {
			Description: "Server-generated unique id of the load test.",
			Type:        schema.TypeString,
			Computed:    true,
		},
		"latest_revision_id": {
			Description: "Latest script revision identifier.",
			Type:        schema.TypeString,
			Computed:    true,
		},
	}
}

func dataSourceChaosLoadTestSchema() map[string]*schema.Schema {
	src := resourceChaosLoadTestSchema()
	for name, field := range src {
		if name == "org_id" || name == "project_id" || name == "identifier" {
			continue
		}
		field.Required = false
		field.Optional = false
		field.ForceNew = false
		field.Computed = true
		field.Default = nil
		field.ConflictsWith = nil
		field.RequiredWith = nil
		field.ValidateFunc = nil
		field.DiffSuppressFunc = nil
		field.MaxItems = 0
	}
	src["identifier"].Description = "Identifier of the load test. Letters, digits, underscores, and hyphens only."
	src["org_id"].Description = "Organization identifier of the load test."
	src["project_id"].Description = "Project identifier of the load test."
	src["name"].Description = "Name of the load test."
	src["description"].Description = "Description stored on the load test."
	src["tool_type"].Description = "Load generator stored on the test. The API may return `Locust`, `JMeter`, or `K6`. This data source can read a JMeter test. The resource does not manage JMeter."
	src["target_type"].Description = "Infrastructure target stored on the load test. `linux` or `kubernetes`."
	src["managed_infra"].Description = "Whether Harness provides Kubernetes infrastructure for every run."
	src["environment_id"].Description = "Chaos environment identifier stored on the load test."
	src["infrastructure_id"].Description = "Chaos infrastructure identifier stored on the load test."
	src["service_references"].Description = "Chaos service identities linked to this load test."
	src["tags"].Description = "Tags stored on the load test."
	src["cleanup_policy"].Description = "Whether run pods, config maps, and secrets are deleted or retained after a run."
	src["max_duration_sec"].Description = "Hard cap, in seconds, stored for any run of this load test. 0 when no cap is stored."
	src["tool_config"].Description = "JSON object stored as toolConfig. K6 fields are at the top or under `k6`. Locust is nested under `locust`. JMeter may be nested under `jmeter`. `mode` is `script` or `image`. The script or image is included."
	src["variable"].Description = "Custom variables stored on the load test."
	src["resources"].Description = "Per-pod CPU and memory stored for Kubernetes load test workloads."
	return src
}

func validateLoadTestIdentifier(v interface{}, k string) (ws []string, errs []error) {
	value, _ := v.(string)
	if !loadTestIdentifierPattern.MatchString(value) {
		errs = append(errs, fmt.Errorf("%s must contain only letters, digits, underscores, and hyphens", k))
	}
	return
}

func validateTargetTypeValue(v interface{}, k string) (ws []string, errs []error) {
	switch v.(string) {
	case "", "linux", "kubernetes":
	default:
		errs = append(errs, fmt.Errorf("%s must be linux or kubernetes", k))
	}
	return
}

func suppressEmptyDescription(_, oldValue, newValue string, _ *schema.ResourceData) bool {
	return newValue == "" && oldValue != ""
}

func validateToolConfig(v interface{}, k string) (ws []string, errs []error) {
	raw, _ := v.(string)
	if strings.TrimSpace(raw) == "" {
		errs = append(errs, fmt.Errorf("%s must be a JSON object", k))
		return
	}
	var obj map[string]interface{}
	if err := json.Unmarshal([]byte(raw), &obj); err != nil || obj == nil {
		errs = append(errs, fmt.Errorf("%s must be a JSON object", k))
		return
	}
	if err := rejectAuthoredVariables(obj); err != nil {
		errs = append(errs, fmt.Errorf("%s %w", k, err))
	}
	return
}

func suppressToolConfigDiff(_, oldValue, newValue string, _ *schema.ResourceData) bool {
	oldNorm, oldErr := normalizeToolConfig(oldValue)
	newNorm, newErr := normalizeToolConfig(newValue)
	return oldErr == nil && newErr == nil && oldNorm == newNorm
}
