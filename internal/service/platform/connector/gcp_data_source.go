package connector

import (
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DatasourceConnectorGcp() *schema.Resource {
	resource := &schema.Resource{
		Description: "Datasource for looking up a Gcp connector.",
		ReadContext: resourceConnectorGcpRead,

		Schema: map[string]*schema.Schema{
			"manual": {
				Description: "Manual credential configuration.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"secret_key_ref": {
							Description: "Reference to the Harness secret containing the secret key." + secret_ref_text,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"delegate_selectors": {
							Description: "The delegates to connect with.",
							Type:        schema.TypeSet,
							Computed:    true,
							Elem:        &schema.Schema{Type: schema.TypeString},
						},
					},
				},
			},
			"inherit_from_delegate": {
				Type:        schema.TypeList,
				Description: "Inherit configuration from delegate.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"delegate_selectors": {
							Description: "The delegates to inherit the credentials from.",
							Type:        schema.TypeSet,
							Computed:    true,
							Elem:        &schema.Schema{Type: schema.TypeString},
						},
					},
				},
			},
			"oidc_authentication": {
				Description: "Authentication using harness oidc.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"workload_pool_id": {
							Description: "The workload pool ID value created in GCP.",
							Type:        schema.TypeString,
							Required:    true,
						},
						"provider_id": {
							Description: "The OIDC provider ID value configured in GCP.",
							Type:        schema.TypeString,
							Required:    true,
						},
						"gcp_project_id": {
							Description: "The project number of the GCP project that is used to create the workload identity..",
							Type:        schema.TypeString,
							Required:    true,
						},
						"service_account_email": {
							Description: "The service account linked to workload identity pool while setting GCP workload identity provider.",
							Type:        schema.TypeString,
							Required:    true,
						},
						"delegate_selectors": {
							Description: "The delegates to inherit the credentials from.",
							Type:        schema.TypeSet,
							Computed:    true,
							Elem:        &schema.Schema{Type: schema.TypeString},
						},
					},
				},
			},
		},
	}

	helpers.SetMultiLevelDatasourceSchemaIdentifierRequired(resource.Schema)

	return resource
}
