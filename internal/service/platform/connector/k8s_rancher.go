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

func ResourceConnectorK8sRancher() *schema.Resource {
	resource := &schema.Resource{
		Description:   "Resource for creating a Rancher connector.",
		ReadContext:   resourceConnectorRancherRead,
		CreateContext: resourceConnectorRancherCreateOrUpdate,
		UpdateContext: resourceConnectorRancherCreateOrUpdate,
		DeleteContext: resourceConnectorDelete,
		Importer:      helpers.MultiLevelResourceImporter,

		Schema: map[string]*schema.Schema{
			"delegate_selectors": {
				Description: "Selectors to use for the delegate.",
				Type:        schema.TypeSet,
				Optional:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},

			"bearer_token": {
				Description: "URL and bearer token for the rancher cluster.",
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"rancher_url": {
							Description: "The URL of the Rancher cluster.",
							Type:        schema.TypeString,
							Required:    true,
						},
						"password_ref": {
							Description: "Reference to the secret containing the bearer token for the rancher cluster." + secret_ref_text,
							Type:        schema.TypeString,
							Required:    true,
						},
					},
				},
			},
		},
	}

	helpers.SetMultiLevelResourceSchema(resource.Schema)

	return resource
}

func resourceConnectorRancherRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	conn, err := resourceConnectorReadBase(ctx, d, meta, nextgen.ConnectorTypes.Rancher)
	if err != nil {
		return err
	}

	if conn == nil {
		return nil
	}

	if err := readConnectorRancher(d, conn); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceConnectorRancherCreateOrUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	conn := buildConnectorRancher(d)

	newConn, err := resourceConnectorCreateOrUpdateBase(ctx, d, meta, conn)
	if err != nil {
		return err
	}

	if err := readConnectorRancher(d, newConn); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func buildConnectorRancher(d *schema.ResourceData) *nextgen.ConnectorInfo {
	connector := &nextgen.ConnectorInfo{
		Type_:   nextgen.ConnectorTypes.Rancher,
		Rancher: &nextgen.RancherConnector{},
	}

	if attr, ok := d.GetOk("delegate_selectors"); ok {
		delegate_selectors := attr.(*schema.Set).List()
		if len(delegate_selectors) > 0 {
			connector.Rancher.DelegateSelectors = utils.InterfaceSliceToStringSlice(delegate_selectors)
		}
	}

	connector.Rancher.Credential = &nextgen.RancherConnectorConfig{
		Type_: nextgen.RancherConfigTypes.ManualConfig,
		Spec: &nextgen.RancherConnectorConfigAuth{
			Auth: &nextgen.RancherAuthentication{
				Type_:             nextgen.RancherAuthTypes.BearerToken,
				BearerTokenConfig: &nextgen.RancherConnectorBearerTokenAuthentication{},
			},
		},
	}

	if attr, ok := d.GetOk("bearer_token"); ok {
		config := attr.([]interface{})[0].(map[string]interface{})

		if attr := config["rancher_url"].(string); attr != "" {
			connector.Rancher.Credential.Spec.RancherUrl = attr
		}

		if attr := config["password_ref"].(string); attr != "" {
			connector.Rancher.Credential.Spec.Auth.BearerTokenConfig.PasswordRef = attr
		}
	}

	return connector
}

func readConnectorRancher(d *schema.ResourceData, connector *nextgen.ConnectorInfo) error {
	switch connector.Rancher.Credential.Type_ {
	case nextgen.RancherConfigTypes.ManualConfig:
		auth := connector.Rancher.Credential.Spec.Auth
		switch auth.Type_ {
		case nextgen.RancherAuthTypes.BearerToken:
			d.Set("bearer_token", []map[string]interface{}{
				{
					"rancher_url":  connector.Rancher.Credential.Spec.RancherUrl,
					"password_ref": connector.Rancher.Credential.Spec.Auth.BearerTokenConfig.PasswordRef,
				},
			})
			d.Set("delegate_selectors", connector.Rancher.DelegateSelectors)
		default:
			return fmt.Errorf("unsupported auth method: %s", auth.Type_)
		}
	default:
		return fmt.Errorf("unsupported Rancher.Credential.Type_: %s", connector.Rancher.Credential.Type_)
	}

	return nil
}
