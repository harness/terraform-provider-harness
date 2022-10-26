package connector

import (
	"context"

	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceConnectorGCPSecretManager() *schema.Resource {
	resource := &schema.Resource{
		Description:   "Resource for creating a GCP Secret Manager connector.",
		ReadContext:   resourceConnectorGcpSMRead,
		CreateContext: resourceConnectorGcpSMCreateOrUpdate,
		UpdateContext: resourceConnectorGcpSMCreateOrUpdate,
		DeleteContext: resourceConnectorDelete,
		Importer:      helpers.MultiLevelResourceImporter,

		Schema: map[string]*schema.Schema{
			"is_default": {
				Description: "Indicative if this is default Secret manager for secrets.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"delegate_selectors": {
				Description: "Connect using only the delegates which have these tags.",
				Type:        schema.TypeSet,
				Optional:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"credentials_ref": {
				Description: "Reference to the secret containing credentials of IAM service account for Google Secret Manager.",
				Type:        schema.TypeString,
				Required:    true,
			},
		},
	}

	helpers.SetMultiLevelResourceSchema(resource.Schema)

	return resource
}

func resourceConnectorGcpSMRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	conn, err := resourceConnectorReadBase(ctx, d, meta, nextgen.ConnectorTypes.GcpSecretManager)
	if err != nil {
		return err
	}

	if err := readConnectorGcpSM(d, conn); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceConnectorGcpSMCreateOrUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	conn := buildConnectorGcpSM(d)

	newConn, err := resourceConnectorCreateOrUpdateBase(ctx, d, meta, conn)
	if err != nil {
		return err
	}

	if err := readConnectorGcpSM(d, newConn); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func buildConnectorGcpSM(d *schema.ResourceData) *nextgen.ConnectorInfo {
	connector := &nextgen.ConnectorInfo{
		Type_:            nextgen.ConnectorTypes.GcpSecretManager,
		GcpSecretManager: &nextgen.GcpSecretManager{},
	}

	if attr, ok := d.GetOk("is_default"); ok {
		connector.GcpSecretManager.IsDefault = attr.(bool)
	}

	if attr, ok := d.GetOk("credentials_ref"); ok {
		connector.GcpSecretManager.CredentialsRef = attr.(string)
	}

	if attr, ok := d.GetOk("delegate_selectors"); ok {
		connector.GcpSecretManager.DelegateSelectors = utils.InterfaceSliceToStringSlice(attr.(*schema.Set).List())
	}

	return connector
}

func readConnectorGcpSM(d *schema.ResourceData, connector *nextgen.ConnectorInfo) error {
	d.Set("is_default", connector.GcpSecretManager.IsDefault)
	d.Set("credentials_ref", connector.GcpSecretManager.CredentialsRef)
	d.Set("delegate_selectors", connector.GcpSecretManager.DelegateSelectors)
	return nil
}
