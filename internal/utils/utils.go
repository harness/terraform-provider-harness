package utils

import (
	"fmt"
	"hash/crc32"
	"strings"
)

const CacDescription string = "This resource uses the config-as-code API's. When updating the `name` or `path` of this resource you should typically also set the `create_before_destroy = true` lifecycle setting."
const NextgenDescription string = "This resource is part of the Harness nextgen platform."
const CDClientAPIKeyError = "Please provide value for FirstGen API Key in api_key."

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

func GetConflictsWithSlice(source []string, self string) []string {
	tmp := make([]string, len(source))
	copy(tmp, source)

	for i, v := range tmp {
		if v == self {
			return append(tmp[:i], tmp[i+1:]...)
		}
	}
	return tmp
}

// Borrowed from https://github.com/hashicorp/terraform-provider-aws/blob/main/internal/create/hashcode.go
// StringHashcode hashes a string to a unique hashcode.
//
// crc32 returns a uint32, but for our use we need
// and non negative integer. Here we cast to an integer
// and invert it if the result is negative.
func StringHashcode(s string) int {
	v := int(crc32.ChecksumIEEE([]byte(s)))
	if v >= 0 {
		return v
	}
	if -v >= 0 {
		return -v
	}
	// v == MinInt
	return 0
}
