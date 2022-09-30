package helpers

import (
	"strings"

	"github.com/harness/harness-go-sdk/harness/nextgen"
)

func ExpandTags(tags []interface{}) map[string]string {
	result := map[string]string{}

	for _, tag := range tags {
		parts := strings.Split(tag.(string), ":")
		result[parts[0]] = parts[1]
	}

	return result
}

func FlattenTags(tags map[string]string) []string {
	var result []string
	for k, v := range tags {
		result = append(result, k+":"+v)
	}
	return result
}

func ExpandClusters(clusterBasicDTO []interface{}) []nextgen.ClusterBasicDto {
	var result []nextgen.ClusterBasicDto
	for _, cluster := range clusterBasicDTO {
		v := cluster.(map[string]interface{})

		var resultcluster nextgen.ClusterBasicDto
		resultcluster.Identifier = v["identifier"].(string)
		resultcluster.Name = v["name"].(string)
		resultcluster.Scope = v["scope"].(string)
		result = append(result, resultcluster)
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

// func ExpandKeyValueTags(tags []interface{}) map[string]string {
// 	result := map[string]string{}

// 	for _, tag := range tags {
// 		parts := strings.Split(tag.(string), ":")
// 		result[parts[0]] = parts[1]
// 	}

// 	return result
// }
