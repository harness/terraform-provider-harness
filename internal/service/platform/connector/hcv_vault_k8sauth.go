package connector

import (
	"context"

	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceConnectorHCVK8sAuth() *schema.Resource {
	resource := &schema.Resource{
		Description:   "Resource for creating a Hashicorp Vault connector with K8s authentication.",
		ReadContext:   resourceConnectorK8sAuthRead,
		CreateContext: resourceConnectorK8sAuthCreateOrUpdate,
		UpdateContext: resourceConnectorK8sAuthCreateOrUpdate,
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
			"k8s_auth_endpoint": {
				Description: "This is the path where kubernetes auth is enabled in Vault.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"k8s_sa_token_path": {
				Description: "This is the SA token path where the token is mounted in the K8s Pod.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"k8s_vault_auth_role": {
				Description: "This is the role where K8s auth will happen.",
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

func resourceConnectorK8sAuthRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	conn, err := resourceConnectorReadBase(ctx, d, meta, nextgen.ConnectorTypes.Vault)
	if err != nil {
		return err
	}

	if err := readConnectorK8sAuth(d, conn); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceConnectorK8sAuthCreateOrUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	conn := buildConnectorK8sAuth(d)

	newConn, err := resourceConnectorCreateOrUpdateBase(ctx, d, meta, conn)
	if err != nil {
		return err
	}

	if err := readConnectorK8sAuth(d, newConn); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func buildConnectorK8sAuth(d *schema.ResourceData) *nextgen.ConnectorInfo {
	connector := &nextgen.ConnectorInfo{
		Type_: nextgen.ConnectorTypes.Vault,
		Vault: &nextgen.VaultConnector{},
	}

	connector.Vault.UseK8sAuth = true
	connector.Vault.SecretEngineVersion = 2
	connector.Vault.AccessType = "K8s_AUTH"

	if attr, ok := d.GetOk("url"); ok {
		connector.Vault.VaultUrl = attr.(string)
	}

	if attr, ok := d.GetOk("base_secret_path"); ok {
		connector.Vault.BasePath = attr.(string)
	}

	if attr, ok := d.GetOk("namespace"); ok {
		connector.Vault.Namespace = attr.(string)
	}

	if attr, ok := d.GetOk("k8s_auth_endpoint"); ok {
		connector.Vault.K8sAuthEndpoint = attr.(string)
	}

	if attr, ok := d.GetOk("k8s_sa_token_path"); ok {
		connector.Vault.ServiceAccountTokenPath = attr.(string)
	}

	if attr, ok := d.GetOk("k8s_vault_auth_role"); ok {
		connector.Vault.VaultK8sAuthRole = attr.(string)
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

func readConnectorK8sAuth(d *schema.ResourceData, connector *nextgen.ConnectorInfo) error {
	d.Set("url", connector.Vault.VaultUrl)
	d.Set("base_secret_path", connector.Vault.BasePath)
	d.Set("namespace", connector.Vault.Namespace)
	d.Set("delegate_selectors", connector.Vault.DelegateSelectors)

	return nil
}
