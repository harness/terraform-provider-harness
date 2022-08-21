package connector

import (
	"context"

	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceConnectorHCVAgent() *schema.Resource {
	resource := &schema.Resource{
		Description:   "Resource for creating a Hashicorp Vault connector with vault agent authentication.",
		ReadContext:   resourceConnectorHCVAgentRead,
		CreateContext: resourceConnectorHCVAgentCreateOrUpdate,
		UpdateContext: resourceConnectorHCVAgentCreateOrUpdate,
		DeleteContext: resourceConnectorDelete,
		Importer:      helpers.MultiLevelResourceImporter,

		Schema: map[string]*schema.Schema{
			"url": {
				Description: "Url of the HashiCorp Vault server.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"base_secret_path": {
				Description: "Base path of the secret engine.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"namespace": {
				Description: "Vault namespace of the secret engine.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"sink_path": {
				Description: "Sink path of the secret data",
				Type:        schema.TypeString,
				Required:    true,
			},
			"delegate_selectors": {
				Description: "Connect using only the delegates which have these tags.",
				Type:        schema.TypeSet,
				Optional:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"renewal_interval": {
				Description: "Renewal interval of the connector, default is 5 minutes.",
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     5,
			},
			"readonly": {
				Description: "Whether the connector is read only, default is false.",
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
			},
			"isdefault": {
				Description: "Whether the connector is the default secret manager, default is false.",
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
			},
		},
	}

	helpers.SetMultiLevelResourceSchema(resource.Schema)

	return resource
}

func resourceConnectorHCVAgentRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	conn, err := resourceConnectorReadBase(ctx, d, meta, nextgen.ConnectorTypes.Vault)
	if err != nil {
		return err
	}

	if err := readConnectorHCVAgent(d, conn); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceConnectorHCVAgentCreateOrUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	conn := buildConnectorHCVAgent(d)

	newConn, err := resourceConnectorCreateOrUpdateBase(ctx, d, meta, conn)
	if err != nil {
		return err
	}

	if err := readConnectorHCVAgent(d, newConn); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func buildConnectorHCVAgent(d *schema.ResourceData) *nextgen.ConnectorInfo {
	connector := &nextgen.ConnectorInfo{
		Type_: nextgen.ConnectorTypes.Vault,
		Vault: &nextgen.VaultConnector{},
	}

	connector.Vault.UseVaultAgent = true
	connector.Vault.SecretEngineVersion = 2
	connector.Vault.AccessType = "VAULT_AGENT"

	if attr, ok := d.GetOk("url"); ok {
		connector.Vault.VaultUrl = attr.(string)
	}

	if attr, ok := d.GetOk("base_secret_path"); ok {
		connector.Vault.BasePath = attr.(string)
	}

	if attr, ok := d.GetOk("namespace"); ok {
		connector.Vault.Namespace = attr.(string)
	}

	if attr, ok := d.GetOk("sink_path"); ok {
		connector.Vault.SinkPath = attr.(string)
	}

	if attr, ok := d.GetOk("delegate_selectors"); ok {
		connector.Vault.DelegateSelectors = utils.InterfaceSliceToStringSlice(attr.(*schema.Set).List())
	}

	if attr, ok := d.GetOk("renewal_interval"); ok {
		connector.Vault.RenewalIntervalMinutes = int64(attr.(int))
	}

	if attr, ok := d.GetOk("readonly"); ok {
		connector.Vault.ReadOnly = attr.(bool)
	}

	if attr, ok := d.GetOk("isdefault"); ok {
		connector.Vault.IsDefault = attr.(bool)
	}

	return connector
}

func readConnectorHCVAgent(d *schema.ResourceData, connector *nextgen.ConnectorInfo) error {
	d.Set("url", connector.Vault.VaultUrl)
	d.Set("base_secret_path", connector.Vault.BasePath)
	d.Set("namespace", connector.Vault.Namespace)
	d.Set("sink_path", connector.Vault.SinkPath)
	d.Set("delegate_selectors", connector.Vault.DelegateSelectors)

	return nil
}
