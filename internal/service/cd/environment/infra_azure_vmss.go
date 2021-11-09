package environment

import (
	"fmt"
	"strings"

	"github.com/harness-io/harness-go-sdk/harness/cd/cac"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func infraDetailsAzureVmss() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"base_name": {
				Description: "Base name.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"cloud_provider_name": {
				Description: "The name of the cloud provider to connect with.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"host_connection_attrs_name": {
				Description: "The name of the host connection attributes to use.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"resource_group_name": {
				Description: "The name of the resource group.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"subscription_id": {
				Description: "The unique id of the azure subscription.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"username": {
				Description: "The username to connect with.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"auth_type": {
				Description: fmt.Sprintf("The type of authentication to use. Valid options are %s.", strings.Join(cac.VmssAuthTypesSlice, ", ")),
				Type:        schema.TypeString,
				Required:    true,
			},
			"deployment_type": {
				Description: fmt.Sprintf("The type of deployment. Valid options are %s", strings.Join(cac.VmssDeploymentTypesSlice, ", ")),
				Type:        schema.TypeString,
				Required:    true,
			},
		},
	}
}
