package secret

import (
	"context"

	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func ResourceSecretText() *schema.Resource {
	resource := &schema.Resource{
		Description:   "Resource for creating secret of type secret text",
		ReadContext:   resourceSecretTextRead,
		CreateContext: resourceSecretTextCreateOrUpdate,
		UpdateContext: resourceSecretTextCreateOrUpdate,
		DeleteContext: resourceSecretDelete,
		Importer:      helpers.MultiLevelResourceImporter,

		Schema: map[string]*schema.Schema{
			"secret_manager_identifier": {
				Description: "Identifier of the Secret Manager used to manage the secret.",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
			},
			"value_type": {
				Description:  "This has details to specify if the secret value is Inline or Reference.",
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringInSlice([]string{"Reference", "Inline", "CustomSecretManagerValues"}, false),
			},
			"value": {
				Description: "Value of the Secret",
				Sensitive:   true,
				Type:        schema.TypeString,
				Optional:    true,
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
										Type:     schema.TypeString,
										Optional: true,
									},
									// Add other fields for the inner map as needed
								},
							},
						},
						// Add other fields for the outer map as needed
					},
				},
			},
		},
	}
	helpers.SetMultiLevelResourceSchema(resource.Schema)

	return resource
}

func resourceSecretTextRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	secret, err := resourceSecretReadBase(ctx, d, meta, nextgen.SecretTypes.SecretText)
	if err != nil {
		return err
	}

	if err := readSecretText(d, secret); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceSecretTextCreateOrUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	secret := buildSecretText(d)

	newSecret, err := resourceSecretCreateOrUpdateBase(ctx, d, meta, secret)
	if err != nil {
		return err
	}

	if err := readSecretText(d, newSecret); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func buildSecretText(d *schema.ResourceData) *nextgen.Secret {
	secret := &nextgen.Secret{
		Type_: nextgen.SecretTypes.SecretText,
		Text:  &nextgen.SecretTextSpec{},
	}

	if attr, ok := d.GetOk("secret_manager_identifier"); ok {
		secret.Text.SecretManagerIdentifier = attr.(string)
	}

	if attr, ok := d.GetOk("value_type"); ok {
		secret.Text.ValueType = nextgen.SecretTextValueType(attr.(string))
	}

	if attr, ok := d.GetOk("value"); ok {
		secret.Text.Value = attr.(string)
	}

	if attr, ok := d.GetOk("additional_metadata"); ok {
		secret.Text.AdditionalMetadata = readAdditionalMetadata(attr)
	}

	return secret
}

func readSecretText(d *schema.ResourceData, secret *nextgen.Secret) error {
	if secret == nil {
		return nil
	}
	d.Set("secret_manager_identifier", secret.Text.SecretManagerIdentifier)
	d.Set("value_type", secret.Text.ValueType)
	if secret.Text.ValueType == "Reference" {
		d.Set("value", secret.Text.Value)
	}
	if secret.Text.AdditionalMetadata.Values != nil {
		d.Set("additional_metadata", importAdditionalMetadata_2(&secret.Text.AdditionalMetadata))
	}
	return nil
}

func readAdditionalMetadata(metadata interface{}) nextgen.AdditionalMetadata {
	result := nextgen.AdditionalMetadata{}

	// Check if additional_metadata block is present
	metadataList := metadata.([]interface{})
	if len(metadataList) > 0 {
		metadataMap := metadataList[0].(map[string]interface{})

		if valuesSet, ok := metadataMap["values"].(*schema.Set); ok {
			result.Values = make(map[string]string)

			// Loop through "values" set
			for _, v := range valuesSet.List() {
				valueMap := v.(map[string]interface{})
				if version, ok := valueMap["version"].(string); ok {
					result.Values["version"] = version
				}
				// Add other fields as needed
			}
		}
	}

	return result
}

func importAdditionalMetadata(data map[string]string) []map[string]interface{} {
	var result []map[string]interface{}

	for _, value := range data {
		entry := map[string]interface{}{
			"version": value,
			// Add other fields for the inner map as needed
		}

		result = append(result, entry)
	}

	return result
}

func importAdditionalMetadata_2(additionalMetadata *nextgen.AdditionalMetadata) []interface{} {
	response := make([]interface{}, 0)
	data := map[string]interface{}{}
	if additionalMetadata != nil && len(additionalMetadata.Values) > 0 {
		var valuesList []interface{}

		for _, value := range additionalMetadata.Values {
			entry := map[string]string{
				"version": value,
				// Add other fields for the inner map as needed
			}

			valuesList = append(valuesList, entry)
		}

		data["values"] = valuesList
	}
	return append(response, data)
}

// This code uses the Terraform SDK's schema package to decode the additional_metadata parameter from the Terraform configuration into the AdditionalMetadata struct. Adjust the import paths and modify the decoding logic based on your actual requirements and structure.
