package helpers

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func MergeSchemas(src map[string]*schema.Schema, dest map[string]*schema.Schema) {
	for k, v := range src {
		dest[k] = v
	}
}
