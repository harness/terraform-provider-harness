package utils

import (
	"fmt"
	"strings"
)

const CacDescription string = "This resource uses the config-as-code API's. When updating the `name` or `path` of this resource you should typically also set the `create_before_destroy = true` lifecycle setting."
const NextgenDescription string = "This resource is part of the Harness nextgen platform."

func ConfigAsCodeDescription(descripton string) string {
	return fmt.Sprintf("%s %s", descripton, CacDescription)
}

func GetNextgenDescription(description string) string {
	return fmt.Sprintf("[NG] - %s %s", description, NextgenDescription)
}

func ExpandDelegateSelectors(ds []interface{}) []string {
	selectors := make([]string, 0)

	for _, v := range ds {
		selectors = append(selectors, v.(string))
	}

	return selectors
}

func FlattenDelgateSelectors(ds []string) []interface{} {
	selectors := make([]interface{}, 0)

	for _, v := range ds {
		selectors = append(selectors, v)
	}

	return selectors
}

func InterfaceSliceToStringSlice(slice []interface{}) []string {
	ss := make([]string, 0)

	for _, v := range slice {
		ss = append(ss, v.(string))
	}

	return ss
}

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
