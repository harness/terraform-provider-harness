package helpers

import (
	"strings"

	"github.com/harness/harness-go-sdk/harness/nextgen"
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
