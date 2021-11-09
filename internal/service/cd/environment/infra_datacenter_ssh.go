package environment

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func infraDetailsDatacenterSSH() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cloud_provider_name": {
				Description: "The name of the cloud provider to connect with.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"hostnames": {
				Description: "A list of hosts to deploy to.",
				Type:        schema.TypeSet,
				Required:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"host_connection_attributes_name": {
				Description: "The name of the SSH connection attributes to use.",
				Type:        schema.TypeString,
				Required:    true,
			},
		},
	}
}
