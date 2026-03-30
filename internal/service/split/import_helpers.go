package split

import (
	"fmt"
	"strings"
)

// ParseImportID3 splits an import id as org_id/project_id/<third>.
func ParseImportID3(id string) (orgID, projectID, third string, err error) {
	parts := strings.Split(id, "/")
	if len(parts) != 3 {
		return "", "", "", fmt.Errorf("import id must be org_id/project_id/<id>, got %q", id)
	}
	return parts[0], parts[1], parts[2], nil
}

// ParseImportID4 splits an import id as org_id/project_id/<third>/<fourth>.
func ParseImportID4(id string) (a, b, c, d string, err error) {
	parts := strings.Split(id, "/")
	if len(parts) != 4 {
		return "", "", "", "", fmt.Errorf("import id must have 4 slash-separated segments, got %q", id)
	}
	return parts[0], parts[1], parts[2], parts[3], nil
}

// ParseImportID6 splits org_id/project_id/environment_id/api_key_type/name/key_id (six segments; name must not contain '/').
func ParseImportID6(id string) (orgID, projectID, environmentID, apiKeyType, name, keyID string, err error) {
	parts := strings.Split(id, "/")
	if len(parts) != 6 {
		return "", "", "", "", "", "", fmt.Errorf("import id must be org_id/project_id/environment_id/api_key_type/name/key_id (6 segments), got %q", id)
	}
	return parts[0], parts[1], parts[2], parts[3], parts[4], parts[5], nil
}
