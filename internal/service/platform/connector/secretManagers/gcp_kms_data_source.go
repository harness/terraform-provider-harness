package secretManagers

import (
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DatasourceConnectorGcpKms() *schema.Resource {
	resource := &schema.Resource{
		Description: "Datasource for looking up GCP KMS connector.",
		ReadContext: resourceConnectorGcpKmsRead,

		Schema: map[string]*schema.Schema{
			"manual": {
				Description: "Manual credential configuration.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"credentials": {
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
			"oidc_authentication": {
				Description: "Authentication using harness oidc.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"workload_pool_id": {
							Description: "The workload pool ID value created in GCP.",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"provider_id": {
							Description: "The OIDC provider ID value configured in GCP.",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"gcp_project_id": {
							Description: "The project number of the GCP project that is used to create the workload identity..",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"service_account_email": {
							Description: "The service account linked to workload identity pool while setting GCP workload identity provider.",
							Type:        schema.TypeString,
							Computed:    true,
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
			"execute_on_delegate": {
				Description: "Enable this flag to execute on Delegate.",
				Type:        schema.TypeBool,
				Computed:    true,
			},
			"default": {
				Description: "Set this flag to set this secret manager as default secret manager.",
				Type:        schema.TypeBool,
				Computed:    true,
			},
			"region": {
				Description: "The region of the GCP KMS.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"gcp_project_id": {
				Description: "The project ID of the GCP KMS.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"key_ring": {
				Description: "The key ring of the GCP KMS.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"key_name": {
				Description: "The key name of the GCP KMS.",
				Type:        schema.TypeString,
				Computed:    true,
			},
		},
	}

	helpers.SetMultiLevelDatasourceSchemaIdentifierRequired(resource.Schema)

	return resource
}
