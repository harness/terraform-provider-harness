package helpers

import "strings"

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
