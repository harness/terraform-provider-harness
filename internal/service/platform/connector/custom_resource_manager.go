package connector

import (
	"context"

	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
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
			"template_script": {
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
									"useAsDefault": {
										Type:     schema.TypeBool,
										Required: false,
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

func resourceConnectorCustomSMRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

func resourceConnectorCustomSMCreateOrUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	conn := buildConnectorCustomSM(d)

	// Use the base function to create or update the connector
	newConn, err := resourceConnectorCreateOrUpdateBase(ctx, d, meta, conn)
	if err != nil {
		return err
	}

	// Read the connector's current state and update the Terraform state accordingly
	if err := readConnectorCustomSM(d, newConn); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceConnectorDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	_, httpResp, err := c.ConnectorsApi.DeleteConnector(ctx, c.AccountId, d.Id(), &nextgen.ConnectorsApiDeleteConnectorOpts{
		OrgIdentifier:     helpers.BuildField(d, "org_id"),
		ProjectIdentifier: helpers.BuildField(d, "project_id"),
		ForceDelete:       helpers.BuildFieldBool(d, "force_delete")})

	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	return nil
}

func readConnectorCustomSM(d *schema.ResourceData, connector *nextgen.ConnectorInfo) error {
	csm := connector.CustomSecretManager

	// Set common fields
	d.Set("template_script", csm.TemplateScript)
	d.Set("version_label", csm.VersionLabel)
	d.Set("on_delegate", csm.OnDelegate)

	// Template inputs
	if csm.TemplateInputs != nil {
		envVars := make([]interface{}, len(csm.TemplateInputs.EnvironmentVariables))
		for i, v := range csm.TemplateInputs.EnvironmentVariables {
			envVars[i] = map[string]interface{}{
				"name":  v.Name,
				"type":  v.Type,
				"value": v.Value,
			}
		}
		d.Set("template_inputs", map[string]interface{}{
			"environment_variables": envVars,
		})
	}

	// Fields when on_delegate is false
	if !csm.OnDelegate {
		d.Set("target_host", csm.TargetHost)
		d.Set("ssh_secret_ref", csm.SshSecretRef)
		d.Set("working_directory", csm.WorkingDirectory)
	}

	return nil
}
