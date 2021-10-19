package provider

import "fmt"

const cacDescription string = "This object uses the config-as-code API's. When updating the `name` or `path` of this resource you should typically also set the `create_before_destroy = true` lifecycle setting."

func configAsCodeDescription(descripton string) string {
	return fmt.Sprintf("%s %s", descripton, cacDescription)
}

func expandDelegateSelectors(ds []interface{}) []string {
	selectors := make([]string, 0)

	for _, v := range ds {
		selectors = append(selectors, v.(string))
	}

	return selectors
}

func flattenDelgateSelectors(ds []string) []interface{} {
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

// func ActionSliceToStringSlice
