package registry

import (
	"github.com/harness/harness-go-sdk/harness/har"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strings"
)

// the original implementation took two differnet but idential parameters, parent_ref and space_ref
// this function will use the original parameters if they are given, otherwise build it using the standard 
// org_id and project_id parameters found in other standard multi-hierarchy resources
func buildHARPathRef(d *schema.ResourceData, c *har.APIClient) (output *schema.ResourceData) {
	output = d
	if attr, ok := d.GetOk("parent_ref"); ok {
		output = setHarnessHierarchy(d, attr.(string))
	} else {
		if attr, ok := d.GetOk("space_ref"); ok {
			output = setHarnessHierarchy(d, attr.(string))
		} else {
			spaceRef := []string{c.AccountId}
			if org_id, ok := d.GetOk("org_id"); ok {
				spaceRef = append(spaceRef, org_id.(string))
			}
			if project_id, ok := d.GetOk("project_id"); ok {
				spaceRef = append(spaceRef, project_id.(string))
			}
			output = setHarnessHierarchy(d, strings.Join(spaceRef, "/"))
		}
	}
	return
}

// Use the defined space_ref string to generate the correct scoping details.
func setHarnessHierarchy(d *schema.ResourceData, attr string) (output *schema.ResourceData) {
	output = d
	d.Set("parent_ref", attr)
	d.Set("space_ref", attr)
	parts := strings.Split(attr, "/")
	if len(parts) >= 2 {
		d.Set("org_id", parts[1])
	}
	if len(parts) >= 3 {
		d.Set("project_id", parts[2])
	}
	return 
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
