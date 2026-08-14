package helpers

import (
	"strings"

	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ExpandTags(tags []interface{}) map[string]string {
	result := map[string]string{}

	for _, tag := range tags {
		parts := strings.Split(tag.(string), ":")
		if len(parts) == 1 {
			parts = append(parts, "")
		}
		result[parts[0]] = parts[1]
	}

	return result
}

func FlattenTags(tags map[string]string) []string {
	var result []string
	for k, v := range tags {
		if v == "" {
			result = append(result, k)
		} else {
			result = append(result, k+":"+v)
		}
	}
	return result
}

// SetTags writes resource tags into state, leaving the attribute untouched when the API reports no
// tags and prior state also had none.
//
// d.Set always writes the flatmap count key "tags.#", and the SDK's flatmap-to-cty shim treats that
// key's presence as the difference between null and an empty set. Since the null/empty
// reconciliation in normalizeNullValues only runs during apply, an apply settles on null (the config
// wins) while a refresh settles on [], and Terraform reports that disagreement as drift on every
// run even though nothing changed remotely (PL-73759).
//
// GetRawState is non-null only on the resource refresh path: ReadResource populates it from the
// prior state, whereas ReadDataSource reads with no prior state at all. That keeps the guard off the
// data source paths, where tags must stay [] so expressions like length(data.x.tags) keep working.
func SetTags(d *schema.ResourceData, tags map[string]string) error {
	if len(tags) == 0 {
		raw := d.GetRawState()
		if !raw.IsNull() && raw.Type().IsObjectType() && raw.Type().HasAttribute("tags") &&
			raw.GetAttr("tags").IsNull() {
			return nil
		}
	}
	return d.Set("tags", FlattenTags(tags))
}

func ExpandScopeSelector(scopeSelectors []interface{}) []nextgen.ScopeSelector {
	var result []nextgen.ScopeSelector
	for _, scopeSelector := range scopeSelectors {
		v := scopeSelector.(map[string]interface{})

		var resultScopeSelector nextgen.ScopeSelector
		resultScopeSelector.Filter = v["filter"].(string)
		resultScopeSelector.AccountIdentifier = v["account_id"].(string)
		resultScopeSelector.OrgIdentifier = v["org_id"].(string)
		resultScopeSelector.ProjectIdentifier = v["project_id"].(string)
		result = append(result, resultScopeSelector)
	}
	return result
}

func ExpandField(permissions []interface{}) []string {
	var result []string
	for _, permission := range permissions {
		result = append(result, permission.(string))
	}
	return result
}

// ExpandPipelineTags converts Terraform tag strings to a map for Pipeline API.
// Unlike ExpandTags(), this function splits on the FIRST colon only, preserving
// any additional colons in the tag value. This is critical for pipeline tags that
// contain Harness expressions, URLs, timestamps, or other colon-separated values.
//
// Examples:
//   Input:  "ImagePush:<+condition?value1:value2>"
//   Output: {ImagePush: "<+condition?value1:value2>"}
//
//   Input:  "registry:https://example.com:5000/repo"
//   Output: {registry: "https://example.com:5000/repo"}
//
// Fixes: PIPE-30810 - Pipeline tags with colons were being truncated
func ExpandPipelineTags(tags []interface{}) map[string]string {
	result := make(map[string]string)

	for _, tag := range tags {
		tagStr := tag.(string)
		// Split on first colon only - everything after first : is the value
		parts := strings.SplitN(tagStr, ":", 2)

		if len(parts) == 1 {
			// Tag has no colon, treat as key with empty value
			result[parts[0]] = ""
		} else {
			// Tag has "key:value" format where value may contain colons
			result[parts[0]] = parts[1]
		}
	}

	return result
}

// func ExpandKeyValueTags(tags []interface{}) map[string]string {
// 	result := map[string]string{}

// 	for _, tag := range tags {
// 		parts := strings.Split(tag.(string), ":")
// 		result[parts[0]] = parts[1]
// 	}

// 	return result
// }
