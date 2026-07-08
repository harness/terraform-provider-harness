package helpers

import (
	"strconv"
	"strings"

	hcty "github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// WoStringValue returns the string value of a write-only attribute from the current apply config.
// key uses the same dotted-index format as d.GetOk, e.g. "repo.0.password_wo".
// Returns ("", false) when the attribute is absent, null, or during ReadResource (no config context).
func WoStringValue(d *schema.ResourceData, key string) (string, bool) {
	path := parseCtyPath(key)
	v, diags := d.GetRawConfigAt(path)
	if !diags.HasError() && v.IsKnown() && !v.IsNull() && v.Type() == hcty.String {
		return v.AsString(), true
	}
	return "", false
}

// WoActive reports whether the write-only credential path is currently active.
// Returns true if either:
//   - the _wo attribute (woKey) is set in the current apply config, OR
//   - the companion _wo_version integer (versionKey) is non-zero in state (refresh context,
//     where GetRawConfigAt always returns null because there is no config).
func WoActive(d *schema.ResourceData, woKey, versionKey string) bool {
	if _, ok := WoStringValue(d, woKey); ok {
		return true
	}
	_, ok := d.GetOk(versionKey)
	return ok
}

// parseCtyPath converts a dotted state-path string (e.g. "repo.0.password_wo") into a cty.Path.
// Numeric segments become IndexInt steps; all others become GetAttr steps.
func parseCtyPath(key string) hcty.Path {
	parts := strings.Split(key, ".")
	var path hcty.Path
	for _, p := range parts {
		if i, err := strconv.Atoi(p); err == nil {
			path = path.IndexInt(i)
		} else if len(path) == 0 {
			path = hcty.GetAttrPath(p)
		} else {
			path = path.GetAttr(p)
		}
	}
	return path
}
