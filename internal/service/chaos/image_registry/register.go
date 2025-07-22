package image_registry

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Resource() *schema.Resource {
	return ResourceChaosImageRegistry()
}

func DataSource() *schema.Resource {
	return DataSourceChaosImageRegistry()
}
