package secretManagers

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
			"credentials_ref": {
				Description: "Reference to the secret containing credentials of IAM service account for Google Secret Manager." + secret_ref_text,
				Type:        schema.TypeString,
				Optional:    true,
				ConflictsWith: []string{
					"inherit_from_delegate",
					"oidc_authentication",
				},
				AtLeastOneOf: []string{
					"inherit_from_delegate",
					"credentials_ref",
					"oidc_authentication",
				},
				ExactlyOneOf: []string{
					"credentials_ref",
					"inherit_from_delegate",
					"oidc_authentication",
				},
			},
			"inherit_from_delegate": {
				Type:        schema.TypeBool,
				Description: "Inherit configuration from delegate.",
				Optional:    true,
				ConflictsWith: []string{
					"credentials_ref",
					"oidc_authentication",
				},
				AtLeastOneOf: []string{
					"inherit_from_delegate",
					"credentials_ref",
					"oidc_authentication",
				},
				ExactlyOneOf: []string{
					"credentials_ref",
					"inherit_from_delegate",
					"oidc_authentication",
				},
			},
			"oidc_authentication": {
				Type:        schema.TypeList,
				Description: "Authentication using harness oidc.",
				Optional:    true,
				ConflictsWith: []string{
					"credentials_ref",
					"inherit_from_delegate",
				},
				AtLeastOneOf: []string{
					"inherit_from_delegate",
					"credentials_ref",
					"oidc_authentication",
				},
				ExactlyOneOf: []string{
					"credentials_ref",
					"inherit_from_delegate",
					"oidc_authentication",
				},
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"workload_pool_id": {
							Description: "The workload pool ID value created in GCP.",
							Type:        schema.TypeString,
							Required:    true,
						},
						"provider_id": {
							Description: "The OIDC provider ID value configured in GCP.",
							Type:        schema.TypeString,
							Required:    true,
						},
						"gcp_project_id": {
							Description: "The project number of the GCP project that is used to create the workload identity.",
							Type:        schema.TypeString,
							Required:    true,
						},
						"service_account_email": {
							Description: "The service account linked to workload identity pool while setting GCP workload identity provider.",
							Type:        schema.TypeString,
							Required:    true,
						},
					},
				},
			},
			"delegate_selectors": {
				Description: "The delegates to inherit the credentials from.",
				Type:        schema.TypeSet,
				Optional:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"execute_on_delegate": {
				Description: "Execute on delegate or not.",
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
			},
			"is_default": {
				Description: "Set this flag to set this secret manager as default secret manager.",
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
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

	if conn == nil {
		return nil
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
		Type_: nextgen.ConnectorTypes.GcpSecretManager,
		GcpSecretManager: &nextgen.GcpSecretManager{
			Credential: &nextgen.GcpConnectorCredential{},
		},
	}

	if attr, ok := d.GetOk("delegate_selectors"); ok {
		connector.GcpSecretManager.DelegateSelectors = utils.InterfaceSliceToStringSlice(attr.(*schema.Set).List())
	}

	if attr, ok := d.GetOk("execute_on_delegate"); ok {
		connector.GcpSecretManager.ExecuteOnDelegate = attr.(bool)
	}

	if attr, ok := d.GetOk("is_default"); ok {
		connector.GcpSecretManager.Default_ = attr.(bool)
	}

	if attr, ok := d.GetOk("credentials_ref"); ok {
		connector.GcpSecretManager.Credential.Type_ = nextgen.GcpAuthTypes.ManualConfig
		connector.GcpSecretManager.Credential.ManualConfig = &nextgen.GcpManualDetails{
			SecretKeyRef: attr.(string),
		}
		connector.GcpSecretManager.CredentialsRef = attr.(string)
	}

	if _, ok := d.GetOk("inherit_from_delegate"); ok {
		connector.GcpSecretManager.Credential.Type_ = nextgen.GcpAuthTypes.InheritFromDelegate
		connector.GcpSecretManager.AssumeCredentialsOnDelegate = true
	}

	if attr, ok := d.GetOk("oidc_authentication"); ok {
		config := attr.([]interface{})[0].(map[string]interface{})
		connector.GcpSecretManager.Credential.Type_ = nextgen.GcpAuthTypes.OidcAuthentication
		connector.GcpSecretManager.Credential.OidcConfig = &nextgen.GcpOidcDetails{}

		if attr := config["workload_pool_id"].(string); attr != "" {
			connector.GcpSecretManager.Credential.OidcConfig.WorkloadPoolId = attr
		}

		if attr := config["provider_id"].(string); attr != "" {
			connector.GcpSecretManager.Credential.OidcConfig.ProviderId = attr
		}

		if attr := config["service_account_email"].(string); attr != "" {
			connector.GcpSecretManager.Credential.OidcConfig.ServiceAccountEmail = attr
		}

		if attr := config["gcp_project_id"].(string); attr != "" {
			connector.GcpSecretManager.Credential.OidcConfig.GcpProjectId = attr
		}
	}

	return connector
}

func readConnectorGcpSM(d *schema.ResourceData, connector *nextgen.ConnectorInfo) error {

	switch connector.GcpSecretManager.Credential.Type_ {
	case nextgen.GcpAuthTypes.InheritFromDelegate:
		d.Set("inherit_from_delegate", true)
	case nextgen.GcpAuthTypes.OidcAuthentication:
		d.Set("oidc_authentication", []map[string]interface{}{
			{
				"workload_pool_id":      connector.GcpSecretManager.Credential.OidcConfig.WorkloadPoolId,
				"provider_id":           connector.GcpSecretManager.Credential.OidcConfig.ProviderId,
				"gcp_project_id":        connector.GcpSecretManager.Credential.OidcConfig.GcpProjectId,
				"service_account_email": connector.GcpSecretManager.Credential.OidcConfig.ServiceAccountEmail,
			},
		})
	}
	d.Set("is_default", connector.GcpSecretManager.Default_)
	d.Set("execute_on_delegate", connector.GcpSecretManager.ExecuteOnDelegate)
	d.Set("delegate_selectors", connector.GcpSecretManager.DelegateSelectors)
	d.Set("credentials_ref", connector.GcpSecretManager.CredentialsRef)
	return nil
}
