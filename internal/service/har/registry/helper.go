package registry

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strings"
)

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

func expandStringSet(s *schema.Set) []string {
	if s == nil {
		return nil
	}
	out := make([]string, 0, s.Len())
	for _, v := range s.List() {
		out = append(out, v.(string))
	}
	return out
}
