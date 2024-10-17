package connector

import (
	"context"

	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func ResourceConnectorAzureKeyVault() *schema.Resource {
	resource := &schema.Resource{
		Description:   "Resource for creating an Azure key vault in Harness.",
		ReadContext:   resourceConnectorAzureKeyVaultRead,
		CreateContext: resourceConnectorAzureKeyVaultCreateOrUpdate,
		UpdateContext: resourceConnectorAzureKeyVaultCreateOrUpdate,
		DeleteContext: resourceConnectorDelete,
		Importer:      helpers.MultiLevelResourceImporter,

		Schema: map[string]*schema.Schema{
			"client_id": {
				Description: "Application ID of the Azure App.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"secret_key": {
				Description: "The Harness text secret with the Azure authentication key as its value.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"tenant_id": {
				Description: "The Azure Active Directory (Azure AD) directory ID where you created your application.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"vault_name": {
				Description: "Name of the vault.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"subscription": {
				Description: "Azure subscription ID.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"is_default": {
				Description: "Specifies whether or not is the default value.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"azure_environment_type": {
				Description:  "Azure environment type. Possible values: AZURE or AZURE_US_GOVERNMENT. Default value: AZURE",
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validation.StringInSlice([]string{"AZURE", "AZURE_US_GOVERNMENT"}, false),
			},
			"delegate_selectors": {
				Description: "Tags to filter delegates for connection.",
				Type:        schema.TypeSet,
				Optional:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
		},
	}

	helpers.SetMultiLevelResourceSchema(resource.Schema)

	return resource
}

func resourceConnectorAzureKeyVaultRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	conn, err := resourceConnectorReadBase(ctx, d, meta, nextgen.ConnectorTypes.AzureKeyVault)
	if err != nil {
		return err
	}

	if conn == nil {
		return nil
	}

	if err := readConnectorAzureKeyVault(d, conn); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceConnectorAzureKeyVaultCreateOrUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	conn := buildConnectorAzureKeyVault(d)

	newConn, err := resourceConnectorCreateOrUpdateBase(ctx, d, meta, conn)
	if err != nil {
		return err
	}

	if err := readConnectorAzureKeyVault(d, newConn); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func buildConnectorAzureKeyVault(d *schema.ResourceData) *nextgen.ConnectorInfo {
	connector := &nextgen.ConnectorInfo{
		Type_:         nextgen.ConnectorTypes.AzureKeyVault,
		AzureKeyVault: &nextgen.AzureKeyVaultConnector{},
	}

	if attr, ok := d.GetOk("tenant_id"); ok {
		connector.AzureKeyVault.TenantId = attr.(string)
	}

	if attr, ok := d.GetOk("client_id"); ok {
		connector.AzureKeyVault.ClientId = attr.(string)
	}

	if attr, ok := d.GetOk("secret_key"); ok {
		connector.AzureKeyVault.SecretKey = attr.(string)
	}

	if attr, ok := d.GetOk("vault_name"); ok {
		connector.AzureKeyVault.VaultName = attr.(string)
	}

	if attr, ok := d.GetOk("subscription"); ok {
		connector.AzureKeyVault.Subscription = attr.(string)
	}

	if attr, ok := d.GetOk("is_default"); ok {
		connector.AzureKeyVault.IsDefault = attr.(bool)
	}

	if attr, ok := d.GetOk("azure_environment_type"); ok {
		connector.AzureKeyVault.AzureEnvironmentType = attr.(string)
	}

	if attr, ok := d.GetOk("delegate_selectors"); ok {
		delegate_selectors := attr.(*schema.Set).List()
		if len(delegate_selectors) > 0 {
			connector.Azure.DelegateSelectors = utils.InterfaceSliceToStringSlice(delegate_selectors)
		}
	}

	return connector
}

func readConnectorAzureKeyVault(d *schema.ResourceData, connector *nextgen.ConnectorInfo) error {
	d.Set("client_id", connector.AzureKeyVault.ClientId)
	d.Set("secret_key", connector.AzureKeyVault.SecretKey)
	d.Set("tenant_id", connector.AzureKeyVault.TenantId)
	d.Set("vault_name", connector.AzureKeyVault.VaultName)
	d.Set("subscription", connector.AzureKeyVault.Subscription)
	d.Set("is_default", connector.AzureKeyVault.IsDefault)

	d.Set("delegate_selectors", connector.AzureKeyVault.DelegateSelectors)
	d.Set("azure_environment_type", connector.AzureKeyVault.AzureEnvironmentType)

	return nil
}
