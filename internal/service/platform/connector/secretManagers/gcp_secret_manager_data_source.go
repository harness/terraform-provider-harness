package secretManagers

import (
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DatasourceConnectorGcpSM() *schema.Resource {
	resource := &schema.Resource{
		Description: "Datasource for looking up GCP Secret Manager connector.",
		ReadContext: resourceConnectorGcpSMRead,

		Schema: map[string]*schema.Schema{
			"execute_on_delegate": {
				Description: "Execute on delegate or not.",
				Type:        schema.TypeBool,
				Computed:    true,
			},
			"is_default": {
				Description: "Set this flag to set this secret manager as default secret manager.",
				Type:        schema.TypeBool,
				Computed:    true,
			},
			"delegate_selectors": {
				Description: "The delegates to inherit the credentials from.",
				Type:        schema.TypeSet,
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"credentials_ref": {
				Description: "Reference to the secret containing credentials of IAM service account for Google Secret Manager." + secret_ref_text,
				Type:        schema.TypeString,
				Computed:    true,
			},
			"inherit_from_delegate": {
				Type:        schema.TypeBool,
				Description: "Inherit configuration from delegate.",
				Computed:    true,
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
					},
				},
			},
		},
	}

	helpers.SetMultiLevelDatasourceSchemaIdentifierRequired(resource.Schema)

	return resource
}
