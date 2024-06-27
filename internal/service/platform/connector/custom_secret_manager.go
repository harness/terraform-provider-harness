package connector

import (
	"context"
	"log"

	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceConnectorCSM() *schema.Resource {
	resource := &schema.Resource{
		Description:   "Resource for creating a Custom Secrets Manager (CSM) connector.",
		ReadContext:   resourceConnectorCustomSMRead,
		CreateContext: resourceConnectorCustomSMCreateOrUpdate,
		UpdateContext: resourceConnectorCustomSMCreateOrUpdate,
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
				Default:  true,
			},
			"timeout": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"delegate_selectors": {
				Description: "Tags to filter delegates for connection.",
				Type:        schema.TypeSet,
				Optional:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
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
			"target_host": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Host where the custom secrets manager is located, required if 'on_delegate' is false.",
			},
			"ssh_secret_ref": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "SSH secret reference for the custom secrets manager, required if 'on_delegate' is false.",
			},
			"working_directory": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The working directory for operations, required if 'on_delegate' is false.",
			},
			"template_inputs": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"environment_variable": {
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
									"default": {
										Type:     schema.TypeBool,
										Optional: true,
										Default:  false,
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
	conn, err := resourceConnectorReadBase(ctx, d, meta, nextgen.ConnectorTypes.CustomSecretManager)
	if err != nil {
		return err
	}

	if conn == nil {
		return nil
	}

	if err := readConnectorCustomSM(d, conn); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func readConnectorCustomSM(d *schema.ResourceData, connector *nextgen.ConnectorInfo) error {
	csm := connector.CustomSecretManager

	d.Set("delegate_selectors", csm.DelegateSelectors)
	d.Set("on_delegate", csm.OnDelegate)
	d.Set("template_ref", csm.Template.TemplateRef)
	d.Set("version_label", csm.Template.VersionLabel)
	d.Set("timeout", csm.Timeout)

	// Template inputs
	if csm.Template.TemplateInputs != nil {
		envVars := make([]interface{}, len(csm.Template.TemplateInputs["environmentVariables"]))
		for i, v := range csm.Template.TemplateInputs["environmentVariables"] {
			envVars[i] = map[string]interface{}{
				"name":    v.Name,
				"type":    v.Type_,
				"value":   v.Value,
				"default": v.UseAsDefault,
			}
		}
		d.Set("template_inputs", map[string]interface{}{
			"environment_variable": envVars,
		})
	}

	if !csm.OnDelegate {
		d.Set("target_host", csm.Host)
		d.Set("ssh_secret_ref", csm.ConnectorRef)
		d.Set("working_directory", csm.WorkingDirectory)
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

	if err := readConnectorCustomSM(d, newConn); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func buildConnectorCustomSM(d *schema.ResourceData) *nextgen.ConnectorInfo {
	connector := &nextgen.ConnectorInfo{
		Type_:               nextgen.ConnectorTypes.CustomSecretManager,
		CustomSecretManager: &nextgen.CustomSecretManager{},
	}

	if attr, ok := d.GetOk("timeout"); ok {
		log.Printf("timeout exists with value: %v", attr)
		connector.CustomSecretManager.Timeout = attr.(int)
	} else {
		log.Printf("timeout does not exist or is nil")
	}

	if attr, ok := d.GetOk("on_delegate"); ok {
		log.Printf("on_delegate exists with value: %v", attr)
		connector.CustomSecretManager.OnDelegate = attr.(bool)
	} else {
		log.Printf("on_delegate does not exist or is nil")
	}

	connector.CustomSecretManager.Template = &nextgen.TemplateLinkConfigForCustomSecretManager{}
	if attr, ok := d.GetOk("template_ref"); ok {
		connector.CustomSecretManager.Template.TemplateRef = attr.(string)
	}

	if attr, ok := d.GetOk("version_label"); ok {
		connector.CustomSecretManager.Template.VersionLabel = attr.(string)
	}

	if attr, ok := d.GetOk("delegate_selectors"); ok {
		connector.CustomSecretManager.DelegateSelectors = utils.InterfaceSliceToStringSlice(attr.(*schema.Set).List())
	}

	if attr, ok := d.GetOk("default"); ok {
		connector.CustomSecretManager.Default_ = attr.(bool)
	}

	if attr, ok := d.GetOk("template_inputs"); ok {
		templateInputsList := attr.([]interface{})
		environmentVariables := []nextgen.NameValuePairWithDefault{}

		for _, templateInputInterface := range templateInputsList {
			templateInput := templateInputInterface.(map[string]interface{})
			if envVarsInterface, ok := templateInput["environment_variable"]; ok {
				envVarsList := envVarsInterface.([]interface{})

				for _, envVarInterface := range envVarsList {
					envVar := envVarInterface.(map[string]interface{})
					environmentVariables = append(environmentVariables, nextgen.NameValuePairWithDefault{
						Name:         envVar["name"].(string),
						Value:        envVar["value"].(string),
						Type_:        envVar["type"].(string),
						UseAsDefault: envVar["default"].(bool),
					})
				}
			}
		}
		template_set := make(map[string][]nextgen.NameValuePairWithDefault)
		template_set["environmentVariables"] = environmentVariables
		connector.CustomSecretManager.Template.TemplateInputs = template_set
	}
	if onDelegate, ok := d.GetOk("on_delegate"); ok && !onDelegate.(bool) {
		log.Printf("on_delegate is false")
		if attr, ok := d.GetOk("working_directory"); ok {
			connector.CustomSecretManager.WorkingDirectory = attr.(string)
		}

		if attr, ok := d.GetOk("ssh_secret_ref"); ok {
			connector.CustomSecretManager.ConnectorRef = attr.(string)
		}

		if attr, ok := d.GetOk("target_host"); ok {
			connector.CustomSecretManager.Host = attr.(string)
		}
	}

	return connector
}
