package loadtest

import (
	"encoding/json"
	"fmt"
	"math"
	"strconv"
	"strings"

	sdkloadtest "github.com/harness/harness-go-sdk/harness/loadtest"
	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func scopeFrom(d *schema.ResourceData, accountID string) sdkloadtest.Scope {
	return sdkloadtest.Scope{
		AccountID: accountID,
		OrgID:     d.Get("org_id").(string),
		ProjectID: d.Get("project_id").(string),
	}
}

func loadTestCreateRequest(d *schema.ResourceData) (sdkloadtest.LoadTestCreateRequest, error) {
	toolConfig, err := parseToolConfig(d.Get("tool_config").(string))
	if err != nil {
		return sdkloadtest.LoadTestCreateRequest{}, err
	}
	managed := d.Get("managed_infra").(bool)
	environmentID := d.Get("environment_id").(string)
	infrastructureID := d.Get("infrastructure_id").(string)
	if err := validateInfra(managed, environmentID, infrastructureID); err != nil {
		return sdkloadtest.LoadTestCreateRequest{}, err
	}
	if err := validateTargetType(d.Get("tool_type").(string), d.Get("target_type").(string), managed); err != nil {
		return sdkloadtest.LoadTestCreateRequest{}, err
	}
	if err := rejectAuthoredVariables(toolConfig); err != nil {
		return sdkloadtest.LoadTestCreateRequest{}, err
	}
	if err := validateToolMode(d.Get("tool_type").(string), toolConfig); err != nil {
		return sdkloadtest.LoadTestCreateRequest{}, err
	}
	req := sdkloadtest.LoadTestCreateRequest{
		Identity:          d.Get("identifier").(string),
		Name:              d.Get("name").(string),
		Description:       d.Get("description").(string),
		ToolType:          d.Get("tool_type").(string),
		TargetType:        d.Get("target_type").(string),
		ManagedInfra:      managed,
		Tags:              expandStringSet(d.Get("tags")),
		ServiceReferences: expandStringSet(d.Get("service_references")),
		Variables:         expandVariables(d.Get("variable")),
		ToolConfig:        toolConfig,
		CleanupPolicy:     d.Get("cleanup_policy").(string),
		Resources:         expandResources(d.Get("resources")),
	}
	if !managed {
		req.EnvironmentIdentifier = environmentID
		req.InfraIdentifier = infrastructureID
	}
	if v, ok := d.GetOk("max_duration_sec"); ok {
		seconds := v.(int)
		req.MaxDurationSec = &seconds
	}
	return req, nil
}

func loadTestUpdateRequest(d *schema.ResourceData) (sdkloadtest.LoadTestUpdateRequest, error) {
	toolConfig, err := parseToolConfig(d.Get("tool_config").(string))
	if err != nil {
		return sdkloadtest.LoadTestUpdateRequest{}, err
	}
	managed := d.Get("managed_infra").(bool)
	environmentID := d.Get("environment_id").(string)
	infrastructureID := d.Get("infrastructure_id").(string)
	if err := validateInfra(managed, environmentID, infrastructureID); err != nil {
		return sdkloadtest.LoadTestUpdateRequest{}, err
	}
	if err := validateTargetType(d.Get("tool_type").(string), d.Get("target_type").(string), managed); err != nil {
		return sdkloadtest.LoadTestUpdateRequest{}, err
	}
	if err := rejectAuthoredVariables(toolConfig); err != nil {
		return sdkloadtest.LoadTestUpdateRequest{}, err
	}
	if err := validateToolMode(d.Get("tool_type").(string), toolConfig); err != nil {
		return sdkloadtest.LoadTestUpdateRequest{}, err
	}
	req := sdkloadtest.LoadTestUpdateRequest{
		Name:              d.Get("name").(string),
		Description:       d.Get("description").(string),
		Tags:              expandStringSet(d.Get("tags")),
		ServiceReferences: expandStringSet(d.Get("service_references")),
		ManagedInfra:      &managed,
		TargetType:        d.Get("target_type").(string),
		Variables:         expandVariables(d.Get("variable")),
		ToolConfig:        toolConfig,
		CleanupPolicy:     d.Get("cleanup_policy").(string),
	}
	if !managed {
		req.EnvironmentIdentifier = environmentID
		req.InfraIdentifier = infrastructureID
	}
	seconds := d.Get("max_duration_sec").(int)
	req.MaxDurationSec = &seconds
	if _, ok := d.GetOk("resources"); ok {
		req.Resources = expandResources(d.Get("resources"))
	} else {
		req.Resources = &sdkloadtest.LoadTestResources{}
	}
	return req, nil
}

// validateLoadTestPlan reads the configuration, not the planned value. target_type is computed,
// so an omitted value is unknown in the plan and Get would skip the check.
func validateLoadTestPlan(d *schema.ResourceDiff) error {
	toolType, toolKnown := rawString(d, "tool_type")
	targetType, targetKnown := rawString(d, "target_type")
	managed, managedKnown := rawBool(d, "managed_infra")
	if toolKnown && targetKnown && managedKnown {
		if err := validateTargetType(toolType, targetType, managed); err != nil {
			return err
		}
	}
	environmentID, envKnown := rawString(d, "environment_id")
	infrastructureID, infraKnown := rawString(d, "infrastructure_id")
	if managedKnown && envKnown && infraKnown {
		if err := validateInfra(managed, environmentID, infrastructureID); err != nil {
			return err
		}
	}
	toolConfig, configKnown := rawString(d, "tool_config")
	if toolKnown && configKnown && strings.TrimSpace(toolConfig) != "" {
		cfg, err := parseToolConfig(toolConfig)
		if err != nil {
			return err
		}
		if err := validateToolMode(toolType, cfg); err != nil {
			return err
		}
	}
	return nil
}

func rawAttr(d *schema.ResourceDiff, name string) (cty.Value, bool) {
	cfg := d.GetRawConfig()
	if cfg.IsNull() || !cfg.IsKnown() || !cfg.Type().IsObjectType() || !cfg.Type().HasAttribute(name) {
		return cty.NilVal, false
	}
	attr := cfg.GetAttr(name)
	if !attr.IsKnown() {
		return cty.NilVal, false
	}
	return attr, true
}

func rawString(d *schema.ResourceDiff, name string) (string, bool) {
	attr, ok := rawAttr(d, name)
	if !ok {
		return "", false
	}
	if attr.IsNull() {
		return "", true
	}
	if attr.Type() != cty.String {
		return "", false
	}
	return attr.AsString(), true
}

func rawBool(d *schema.ResourceDiff, name string) (bool, bool) {
	attr, ok := rawAttr(d, name)
	if !ok {
		return false, false
	}
	if attr.IsNull() {
		return false, true
	}
	if attr.Type() != cty.Bool {
		return false, false
	}
	return attr.True(), true
}

func validateTargetType(toolType, targetType string, managed bool) error {
	needsKubernetes := managed || toolType == sdkloadtest.LoadTestToolK6
	if needsKubernetes && targetType != sdkloadtest.LoadTestTargetKubernetes {
		return fmt.Errorf("target_type must be kubernetes for %s load tests", toolType)
	}
	return nil
}

func validateInfra(managed bool, environmentID, infrastructureID string) error {
	if managed {
		if environmentID != "" || infrastructureID != "" {
			return fmt.Errorf("environment_id and infrastructure_id must be omitted when managed_infra is true")
		}
		return nil
	}
	if environmentID == "" || infrastructureID == "" {
		return fmt.Errorf("environment_id and infrastructure_id are required unless managed_infra is true")
	}
	return nil
}

func setLoadTestState(d *schema.ResourceData, test *sdkloadtest.LoadTest, variables []sdkloadtest.LoadTestVariable) error {
	d.SetId(test.Identity)
	setters := []struct {
		key   string
		value interface{}
	}{
		{"identifier", test.Identity},
		{"name", test.Name},
		{"description", test.Description},
		{"tool_type", test.ToolType},
		{"target_type", test.TargetType},
		{"managed_infra", test.ManagedInfra},
		{"environment_id", test.EnvironmentIdentifier},
		{"infrastructure_id", test.InfraIdentifier},
		{"service_references", nonNilStrings(test.ServiceReferences)},
		{"tags", nonNilStrings(test.Tags)},
		{"cleanup_policy", test.CleanupPolicy},
		{"unique_id", test.UniqueID},
		{"latest_revision_id", test.LatestRevisionIdentifier},
	}
	for _, setter := range setters {
		if err := d.Set(setter.key, setter.value); err != nil {
			return fmt.Errorf("set %s: %w", setter.key, err)
		}
	}
	seconds := 0
	if test.MaxDurationSec != nil {
		seconds = *test.MaxDurationSec
	}
	if err := d.Set("max_duration_sec", seconds); err != nil {
		return err
	}
	stripServerOwnedToolConfig(test.ToolConfig)
	config, err := normalizeToolConfigValue(test.ToolConfig)
	if err != nil {
		return err
	}
	if err := d.Set("tool_config", config); err != nil {
		return err
	}
	if err := d.Set("variable", flattenVariables(variables)); err != nil {
		return err
	}
	return d.Set("resources", flattenResources(test.Resources))
}

func parseToolConfig(raw string) (map[string]interface{}, error) {
	var obj map[string]interface{}
	if err := json.Unmarshal([]byte(raw), &obj); err != nil || obj == nil {
		return nil, fmt.Errorf("tool_config must be a JSON object")
	}
	return obj, nil
}

func normalizeToolConfig(raw string) (string, error) {
	if strings.TrimSpace(raw) == "" {
		return "", nil
	}
	obj, err := parseToolConfig(raw)
	if err != nil {
		return "", err
	}
	stripServerOwnedToolConfig(obj)
	return normalizeToolConfigValue(obj)
}

// validateToolMode checks the mode Load Test Manager accepts for Locust and K6.
// Script mode needs the base64 script. Image mode needs the container image.
// k6 UI mode is not supported.
func validateToolMode(toolType string, cfg map[string]interface{}) error {
	section := cfg
	label := "tool_config"
	switch toolType {
	case sdkloadtest.LoadTestToolLocust:
		nested, ok := cfg["locust"].(map[string]interface{})
		if !ok {
			return fmt.Errorf("tool_config.locust is required for a Locust load test")
		}
		section = nested
		label = "tool_config.locust"
	case sdkloadtest.LoadTestToolK6:
		// Load Test Manager accepts the k6 spec at the top of toolConfig or under k6.
		if _, exists := cfg["k6"]; exists {
			nested, ok := cfg["k6"].(map[string]interface{})
			if !ok {
				return fmt.Errorf("tool_config.k6 must be an object")
			}
			section = nested
			label = "tool_config.k6"
		}
	}
	mode, _ := section["mode"].(string)
	script, _ := section["script"].(map[string]interface{})
	switch mode {
	case "script":
		content, _ := script["content"].(string)
		if strings.TrimSpace(content) == "" {
			return fmt.Errorf("%s mode script requires script.content", label)
		}
	case "image":
		image, _ := script["image"].(string)
		if strings.TrimSpace(image) == "" {
			return fmt.Errorf("%s mode image requires script.image", label)
		}
	case "ui":
		if toolType == sdkloadtest.LoadTestToolK6 {
			return fmt.Errorf("k6 ui mode is not supported; use script or image")
		}
		return fmt.Errorf("%s mode %q is not supported; use script or image", label, mode)
	default:
		return fmt.Errorf("%s mode must be script or image", label)
	}
	return nil
}

// rejectAuthoredVariables fails when tool_config carries variables. The variable block is the only copy Terraform sends.
// Stripping the key would drop the values and still produce an empty plan.
func rejectAuthoredVariables(cfg map[string]interface{}) error {
	for _, section := range toolConfigSections(cfg) {
		if _, ok := section["variables"]; ok {
			return fmt.Errorf("must not contain variables; set them on the variable block")
		}
	}
	return nil
}

// stripServerOwnedToolConfig removes the variables list Load Test Manager copies into toolConfig.
// K6 and Locust tunables are stored as sent. JMeter is not managed here.
func stripServerOwnedToolConfig(cfg map[string]interface{}) {
	for _, section := range toolConfigSections(cfg) {
		delete(section, "variables")
	}
}

func toolConfigSections(cfg map[string]interface{}) []map[string]interface{} {
	if cfg == nil {
		return nil
	}
	sections := []map[string]interface{}{cfg}
	for _, key := range []string{"k6", "locust", "jmeter"} {
		if nested, ok := cfg[key].(map[string]interface{}); ok && nested != nil {
			sections = append(sections, nested)
		}
	}
	return sections
}

func normalizeToolConfigValue(obj map[string]interface{}) (string, error) {
	if obj == nil {
		return "", nil
	}
	buf, err := json.Marshal(obj)
	if err != nil {
		return "", err
	}
	return string(buf), nil
}

func nonNilStrings(values []string) []string {
	if values == nil {
		return []string{}
	}
	return values
}

func expandStringSet(value interface{}) []string {
	set, ok := value.(*schema.Set)
	if !ok || set == nil || set.Len() == 0 {
		return []string{}
	}
	out := make([]string, 0, set.Len())
	for _, item := range set.List() {
		if text, ok := item.(string); ok && text != "" {
			out = append(out, text)
		}
	}
	return out
}

func expandVariables(value interface{}) []sdkloadtest.LoadTestVariable {
	raw, ok := value.([]interface{})
	if !ok || len(raw) == 0 {
		return []sdkloadtest.LoadTestVariable{}
	}
	out := make([]sdkloadtest.LoadTestVariable, 0, len(raw))
	for _, item := range raw {
		block, ok := item.(map[string]interface{})
		if !ok {
			continue
		}
		varType := block["type"].(string)
		out = append(out, sdkloadtest.LoadTestVariable{
			Name:        block["name"].(string),
			Type:        varType,
			Value:       variableWireValue(varType, block["value"].(string)),
			Description: block["description"].(string),
			Required:    block["required"].(bool),
		})
	}
	return out
}

func flattenVariables(variables []sdkloadtest.LoadTestVariable) []interface{} {
	if len(variables) == 0 {
		return []interface{}{}
	}
	out := make([]interface{}, 0, len(variables))
	for _, variable := range variables {
		value := variableValueString(variable.Value)
		out = append(out, map[string]interface{}{
			"name":        variable.Name,
			"type":        variable.Type,
			"value":       value,
			"description": variable.Description,
			"required":    variable.Required,
		})
	}
	return out
}

// variableWireValue sends a Number literal as a JSON number. Load Test Manager stores a
// string as a string, and a numeric toolConfig leaf cannot decode "100". Expressions stay strings.
func variableWireValue(varType, raw string) interface{} {
	if varType != "Number" || strings.Contains(raw, "<+") {
		return raw
	}
	if n, err := strconv.ParseInt(raw, 10, 64); err == nil && strconv.FormatInt(n, 10) == raw {
		return n
	}
	if n, err := strconv.ParseFloat(raw, 64); err == nil && !math.IsNaN(n) && !math.IsInf(n, 0) {
		return n
	}
	return raw
}

// variableValueString keeps JSON numbers in the form Terraform stored.
// fmt.Sprint turns a large whole number into scientific notation and the next plan never settles.
func variableValueString(value interface{}) string {
	switch typed := value.(type) {
	case nil:
		return ""
	case string:
		return typed
	case float64:
		if math.IsNaN(typed) || math.IsInf(typed, 0) {
			return ""
		}
		if typed == math.Trunc(typed) {
			return strconv.FormatFloat(typed, 'f', 0, 64)
		}
		return strconv.FormatFloat(typed, 'f', -1, 64)
	case json.Number:
		return typed.String()
	default:
		return fmt.Sprint(typed)
	}
}

func expandResources(value interface{}) *sdkloadtest.LoadTestResources {
	raw, ok := value.([]interface{})
	if !ok || len(raw) == 0 || raw[0] == nil {
		return &sdkloadtest.LoadTestResources{}
	}
	block, ok := raw[0].(map[string]interface{})
	if !ok {
		return &sdkloadtest.LoadTestResources{}
	}
	return &sdkloadtest.LoadTestResources{
		Requests: expandStringMap(block["requests"]),
		Limits:   expandStringMap(block["limits"]),
	}
}

func expandStringMap(value interface{}) map[string]string {
	raw, ok := value.(map[string]interface{})
	if !ok || len(raw) == 0 {
		return nil
	}
	out := make(map[string]string, len(raw))
	for key, item := range raw {
		if text, ok := item.(string); ok {
			out[key] = text
		}
	}
	if len(out) == 0 {
		return nil
	}
	return out
}

func flattenResources(resources *sdkloadtest.LoadTestResources) []interface{} {
	if resources == nil || (len(resources.Requests) == 0 && len(resources.Limits) == 0) {
		return []interface{}{}
	}
	return []interface{}{map[string]interface{}{
		"requests": resources.Requests,
		"limits":   resources.Limits,
	}}
}

func serviceReferencesDrift(d *schema.ResourceData, stored []string) error {
	configured, ok := configuredServiceReferences(d)
	if !ok || len(configured) == 0 {
		return nil
	}
	if sameStrings(configured, stored) {
		return nil
	}
	if d.Get("managed_infra").(bool) {
		return fmt.Errorf("load test %s is stored, but service_references are not stored when managed_infra is true. Remove service_references from the configuration", d.Id())
	}
	reason := fmt.Sprintf("service_references read back as %v, config has %v", stored, configured)
	if len(stored) == 0 {
		reason = "service_references were not stored; CHAOS_RT_SERVICES_ENABLED is off or the API dropped the list"
	}
	return fmt.Errorf("load test %s is stored, but %s. Remove service_references from the configuration or enable the flag", d.Id(), reason)
}

func configuredServiceReferences(d *schema.ResourceData) ([]string, bool) {
	cfg := d.GetRawConfig()
	if cfg.IsNull() || !cfg.IsKnown() || !cfg.Type().IsObjectType() || !cfg.Type().HasAttribute("service_references") {
		return nil, false
	}
	attr := cfg.GetAttr("service_references")
	if attr.IsNull() || !attr.IsKnown() {
		return nil, true
	}
	var out []string
	attr.ForEachElement(func(_ cty.Value, val cty.Value) bool {
		if val.IsKnown() && val.Type() == cty.String && val.AsString() != "" {
			out = append(out, val.AsString())
		}
		return false
	})
	return out, true
}

func sameStrings(left, right []string) bool {
	if len(left) != len(right) {
		return false
	}
	seen := map[string]int{}
	for _, value := range left {
		seen[value]++
	}
	for _, value := range right {
		seen[value]--
		if seen[value] < 0 {
			return false
		}
	}
	return true
}

// scriptFindingSummary lists every script problem when Load Test Manager returns more than the first one in message.
func scriptFindingSummary(err error) string {
	bodyErr, ok := err.(interface{ Body() []byte })
	if !ok || err == nil {
		return ""
	}
	var parsed struct {
		Validation struct {
			Findings []struct {
				Severity string `json:"severity"`
				Message  string `json:"message"`
				File     string `json:"file"`
				Line     int    `json:"line"`
				Column   int    `json:"column"`
			} `json:"findings"`
		} `json:"validation"`
	}
	if json.Unmarshal(bodyErr.Body(), &parsed) != nil || len(parsed.Validation.Findings) < 2 {
		return ""
	}
	lines := make([]string, 0, len(parsed.Validation.Findings))
	for _, finding := range parsed.Validation.Findings {
		if finding.Severity != "" && finding.Severity != "error" {
			continue
		}
		line := finding.Message
		if finding.File != "" {
			line = finding.File + ": " + line
		}
		if finding.Line > 0 {
			line = fmt.Sprintf("%s (%d:%d)", line, finding.Line, finding.Column)
		}
		if line != "" {
			lines = append(lines, line)
		}
	}
	if len(lines) < 2 {
		return ""
	}
	return "Script validation found " + strings.Join(lines, "; ")
}

func parseLoadTestImportID(id string) (orgID, projectID, identifier string, err error) {
	parts := strings.Split(id, "/")
	if len(parts) != 3 || parts[0] == "" || parts[1] == "" || !loadTestIdentifierPattern.MatchString(parts[2]) {
		return "", "", "", fmt.Errorf("invalid import ID %q, expected org_id/project_id/identifier with an identifier of letters, digits, underscores, and hyphens", id)
	}
	return parts[0], parts[1], parts[2], nil
}
