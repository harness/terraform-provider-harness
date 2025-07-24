package registry

import "strings"

func getParentRef(accountID, orgID, projectID string, parentRef string) string {
	if parentRef != "" {
		return parentRef
	}
	return getRef(accountID, orgID, projectID)
}

func getRef(params ...string) string {
	var result []string
	for _, param := range params {
		if param == "" {
			break
		}
		result = append(result, param)
	}
	return strings.Join(result, "/")
}
