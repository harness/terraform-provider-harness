package secret

import (
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceSecretFile() *schema.Resource {
	resource := &schema.Resource{
		Description: "Datasource for looking up secert file type secret.",
		ReadContext: resourceSecretFileRead,

		Schema: map[string]*schema.Schema{
			"secret_manager_identifier": {
				Description: "Identifier of the Secret Manager used to manage the secret.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"file_path": {
				Description: "Path of the file containing secret value",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"additional_metadata": {
				Description: "Additional Metadata for the Secret",
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"values": {
							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"version": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Version of the secret (for AWS/Azure Secret Manager)",
									},
									"kms_key_id": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "KMS Key ID (for AWS Secret Manager)",
									},
									"regions": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "GCP region for the secret (for GCP Secret Manager)",
									},
									"gcp_project_id": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "GCP Project ID (for GCP Secret Manager)",
									},
								},
							},
						},
					},
				},
			},
		},
	}

	helpers.SetMultiLevelDatasourceSchemaIdentifierRequired(resource.Schema)

	return resource
}
