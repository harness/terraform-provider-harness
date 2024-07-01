package connector

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceConnectorGcp() *schema.Resource {
	resource := &schema.Resource{
		Description:   "Resource for creating a Gcp connector.",
		ReadContext:   resourceConnectorGcpRead,
		CreateContext: resourceConnectorGcpCreateOrUpdate,
		UpdateContext: resourceConnectorGcpCreateOrUpdate,
		DeleteContext: resourceConnectorDelete,
		Importer:      helpers.MultiLevelResourceImporter,

		Schema: map[string]*schema.Schema{
			"oidc_configuration": {
				Description:   "OIDC configuration for GCP connector",
				Type:          schema.TypeList,
				MaxItems:      1,
				Optional:      true,
				ConflictsWith: []string{"manual", "inherit_from_delegate"},
				AtLeastOneOf: []string{
					"inherit_from_delegate",
					"manual",
					"oidc_configuration",
				},
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"workload_pool_id": {
							Description: "The workload pool id",
							Type:        schema.TypeString,
							Required:    true,
						},
						"provider_id": {
							Description: "The provider id",
							Type:        schema.TypeString,
							Required:    true,
						},
						"gcp_project_id": {
							Description: "The GCP project id",
							Type:        schema.TypeString,
							Required:    true,
						},
						"service_account_email": {
							Description: "The service account email",
							Type:        schema.TypeString,
							Required:    true,
						},
					},
				},
			},
			"manual": {
				Description:   "Manual credential configuration.",
				Type:          schema.TypeList,
				MaxItems:      1,
				Optional:      true,
				ConflictsWith: []string{"inherit_from_delegate", "oidc_configuration"},
				AtLeastOneOf: []string{
					"inherit_from_delegate",
					"manual",
					"oidc_configuration",
				},
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"secret_key_ref": {
							Description: "Reference to the Harness secret containing the secret key." + secret_ref_text,
							Type:        schema.TypeString,
							Required:    true,
						},
						"delegate_selectors": {
							Description: "The delegates to connect with.",
							Type:        schema.TypeSet,
							Required:    true,
							Elem:        &schema.Schema{Type: schema.TypeString},
						},
					},
				},
			},
			"inherit_from_delegate": {
				Type:          schema.TypeList,
				Description:   "Inherit configuration from delegate.",
				Optional:      true,
				ConflictsWith: []string{"manual", "oidc_configuration"},
				AtLeastOneOf: []string{
					"inherit_from_delegate",
					"manual",
					"oidc_configuration",
				},
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"delegate_selectors": {
							Description: "The delegates to inherit the credentials from.",
							Type:        schema.TypeSet,
							Required:    true,
							Elem:        &schema.Schema{Type: schema.TypeString},
						},
					},
				},
			},
			"force_delete": {
				Description: "Enable this flag for force deletion of connector",
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
			},
			"execute_on_delegate": {
				Description: "Enable this flag to execute on Delegate",
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
			},
		},
	}

	helpers.SetMultiLevelResourceSchema(resource.Schema)

	return resource
}

func resourceConnectorGcpRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	conn, err := resourceConnectorReadBase(ctx, d, meta, nextgen.ConnectorTypes.Gcp)
	if err != nil {
		return err
	}

	if conn == nil {
		return nil
	}

	if err := readConnectorGcp(d, conn); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceConnectorGcpCreateOrUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	conn := buildConnectorGcp(d)

	newConn, err := resourceConnectorCreateOrUpdateBase(ctx, d, meta, conn)
	if err != nil {
		return err
	}

	if err := readConnectorGcp(d, newConn); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func buildConnectorGcp(d *schema.ResourceData) *nextgen.ConnectorInfo {
	connector := &nextgen.ConnectorInfo{
		Type_: nextgen.ConnectorTypes.Gcp,
		Gcp: &nextgen.GcpConnector{
			Credential: &nextgen.GcpConnectorCredential{},
		},
	}

	if attr, ok := d.GetOk("oidc_configuration"); ok {
		config := attr.([]interface{})[0].(map[string]interface{})
		connector.Gcp.Credential.Type_ = "OidcAuthentication"
		oidcConfig := &struct {
			WorkloadPoolId      string `json:"workload_pool_id"`
			ProviderId          string `json:"provider_id"`
			GcpProjectId        string `json:"gcp_project_id"`
			ServiceAccountEmail string `json:"service_account_email"`
		}{
			WorkloadPoolId:      config["workload_pool_id"].(string),
			ProviderId:          config["provider_id"].(string),
			GcpProjectId:        config["gcp_project_id"].(string),
			ServiceAccountEmail: config["service_account_email"].(string),
		}
		b, _ := json.Marshal(oidcConfig)
		connector.Gcp.Credential.Spec = b
		if attr := config["delegate_selectors"].(*schema.Set).List(); len(attr) > 0 {
			connector.Gcp.DelegateSelectors = utils.InterfaceSliceToStringSlice(attr)
		}
	}

	if attr, ok := d.GetOk("manual"); ok {
		config := attr.([]interface{})[0].(map[string]interface{})
		connector.Gcp.Credential.Type_ = nextgen.GcpAuthTypes.ManualConfig
		connector.Gcp.Credential.ManualConfig = &nextgen.GcpManualDetails{}

		if attr := config["delegate_selectors"].(*schema.Set).List(); len(attr) > 0 {
			connector.Gcp.DelegateSelectors = utils.InterfaceSliceToStringSlice(attr)
		}

		if attr, ok := config["secret_key_ref"]; ok {
			connector.Gcp.Credential.ManualConfig.SecretKeyRef = attr.(string)
		}

	}

	if attr, ok := d.GetOk("inherit_from_delegate"); ok {
		config := attr.([]interface{})[0].(map[string]interface{})
		connector.Gcp.Credential.Type_ = nextgen.GcpAuthTypes.InheritFromDelegate

		if attr := config["delegate_selectors"].(*schema.Set).List(); len(attr) > 0 {
			connector.Gcp.DelegateSelectors = utils.InterfaceSliceToStringSlice(attr)
		}
	}

	if attr, ok := d.GetOk("execute_on_delegate"); ok {
		connector.Gcp.ExecuteOnDelegate = attr.(bool)
	}

	return connector
}

func readConnectorGcp(d *schema.ResourceData, connector *nextgen.ConnectorInfo) error {

	switch connector.Gcp.Credential.Type_ {
	// OIDC Type not added yet, pr here https://github.com/harness/harness-go-sdk/pull/519
	case "OidcAuthentication":
		oidcConfig := &struct {
			WorkloadPoolId      string `json:"workload_pool_id"`
			ProviderId          string `json:"provider_id"`
			GcpProjectId        string `json:"gcp_project_id"`
			ServiceAccountEmail string `json:"service_account_email"`
		}{}
		json.Unmarshal(connector.Gcp.Credential.Spec, oidcConfig)
		d.Set("oidc_configuration", []map[string]interface{}{
			{
				"workload_pool_id":      oidcConfig.WorkloadPoolId,
				"provider_id":           oidcConfig.ProviderId,
				"gcp_project_id":        oidcConfig.GcpProjectId,
				"service_account_email": oidcConfig.ServiceAccountEmail,
			},
		})
	case nextgen.GcpAuthTypes.ManualConfig:
		d.Set("manual", []map[string]interface{}{
			{
				"secret_key_ref":      connector.Gcp.Credential.ManualConfig.SecretKeyRef,
				"delegate_selectors":  connector.Gcp.DelegateSelectors,
				"execute_on_delegate": connector.Gcp.ExecuteOnDelegate,
			},
		})
	case nextgen.GcpAuthTypes.InheritFromDelegate:
		d.Set("inherit_from_delegate", []map[string]interface{}{
			{
				"delegate_selectors":  connector.Gcp.DelegateSelectors,
				"execute_on_delegate": true,
			},
		})
	default:
		return fmt.Errorf("invalid gcp credential type: %s", connector.Gcp.Credential.Type_)
	}

	return nil
}
