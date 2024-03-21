package connector

import (
	"context"
	"fmt"

	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceConnectorCSM() *schema.Resource {
	resource := &schema.Resource{
		Description:   "Resource for creating a Custom Secrets Manager (CSM) connector.",
		ReadContext:   resourceConnectorCSMRead,
		CreateContext: resourceConnectorCSMCreateOrUpdate,
		UpdateContext: resourceConnectorCSMCreateOrUpdate,
		DeleteContext: resourceConnectorDelete,
		Importer:      helpers.MultiLevelResourceImporter,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"identifier": {
				Type:     schema.TypeString,
				Required: true,
			},
			"type": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "CustomSecretManager",
			},
			"on_delegate": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"timeout": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"tags": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"template_ref": {
				Type:     schema.TypeString,
				Required: true,
			},
			"version_label": {
				Type:     schema.TypeString,
				Required: true,
			},
			"template_inputs": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"environment_variables": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:     schema.TypeString,
										Required: true,
									},
									"type": {
										Type:     schema.TypeString,
										Required: true,
									},
									"value": {
										Type:     schema.TypeString,
										Required: true,
									},
								},
							},
						},
					},
				},
			},
		},
	}

	helpers.SetMultiLevelResourceSchema(resource.Schema)

	return resource
}

func resourceConnectorAwsSMRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	conn, err := resourceConnectorReadBase(ctx, d, meta, nextgen.ConnectorTypes.AwsSecretManager)
	if err != nil {
		return err
	}

	if conn == nil {
		return nil
	}

	if err := readConnectorAwsSM(d, conn); err != nil {
		return diag.FromErr(err)
	}

	return nil
}