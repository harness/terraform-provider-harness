package provider

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
