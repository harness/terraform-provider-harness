package iacm

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceIacmDefaultPipeline() *schema.Resource {
	resource := &schema.Resource{
		Description: "Data source for retrieving IACM default pipelines.",

		ReadContext: resourceIacmDefaultPipelineRead,

		Schema: map[string]*schema.Schema{
			"org_id": {
				Description: "Organization identifier of the organization the default pipelines resides in.",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
			},
			"project_id": {
				Description: "Project identifier of the project the default pipelines resides in.",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
			},
			"provisioner_type": {
				Description: "The provisioner associated with this default.",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    false,
			},
			"operation": {
				Description: "The operation associated with this default.",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    false,
			},
			"pipeline": {
				Description: "The pipeline associated with this default.",
				Type:        schema.TypeString,
				Computed:    true,
				ForceNew:    false,
			},
		},
	}
	return resource
}
