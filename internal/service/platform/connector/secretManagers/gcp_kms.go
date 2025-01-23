package secretManagers

import (
	"context"

	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceConnectorGcpKms() *schema.Resource {
	resource := &schema.Resource{
		Description:   "Resource for creating a GCP KMS connector.",
		ReadContext:   resourceConnectorGcpKmsRead,
		CreateContext: resourceConnectorGcpKmsCreateOrUpdate,
		UpdateContext: resourceConnectorGcpKmsCreateOrUpdate,
		DeleteContext: resourceConnectorDelete,
		Importer:      helpers.MultiLevelResourceImporter,

		Schema: map[string]*schema.Schema{
			"manual": {
				Description: "Manual credential configuration.",
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				ConflictsWith: []string{
					"oidc_authentication",
				},
				AtLeastOneOf: []string{
					"manual",
					"oidc_authentication",
				},
				ExactlyOneOf: []string{
					"manual",
					"oidc_authentication",
				},
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"credentials": {
							Description: "Reference to the Harness secret containing the secret key." + secret_ref_text,
							Type:        schema.TypeString,
							Required:    true,
						},
						"delegate_selectors": {
							Description: "The delegates to connect with.",
							Type:        schema.TypeSet,
							Optional:    true,
							Elem:        &schema.Schema{Type: schema.TypeString},
						},
					},
				},
			},
			"oidc_authentication": {
				Type:        schema.TypeList,
				Description: "Authentication using harness oidc.",
				Optional:    true,
				ConflictsWith: []string{
					"manual",
				},
				AtLeastOneOf: []string{
					"manual",
					"oidc_authentication",
				},
				ExactlyOneOf: []string{
					"manual",
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
						"delegate_selectors": {
							Description: "The delegates to inherit the credentials from.",
							Type:        schema.TypeSet,
							Optional:    true,
							Elem:        &schema.Schema{Type: schema.TypeString},
						},
					},
				},
			},
			"execute_on_delegate": {
				Description: "Enable this flag to execute on Delegate.",
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
			},
			"default": {
				Description: "Set this flag to set this secret manager as default secret manager.",
				Type:        schema.TypeBool,
				Optional:    false,
			},
		},
	}

	helpers.SetMultiLevelResourceSchema(resource.Schema)

	return resource
}

func resourceConnectorGcpKmsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	conn, err := resourceConnectorReadBase(ctx, d, meta, nextgen.ConnectorTypes.GcpKms)
	if err != nil {
		return err
	}

	if conn == nil {
		return nil
	}

	if err := readConnectorGcpKms(d, conn); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceConnectorGcpKmsCreateOrUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	conn := buildConnectorGcpKms(d)

	newConn, err := resourceConnectorCreateOrUpdateBase(ctx, d, meta, conn)
	if err != nil {
		return err
	}

	if err := readConnectorGcpKms(d, newConn); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func buildConnectorGcpKms(d *schema.ResourceData) *nextgen.ConnectorInfo {
	connector := &nextgen.ConnectorInfo{
		Type_:  nextgen.ConnectorTypes.GcpKms,
		GcpKms: &nextgen.GcpKmsConnector{},
	}

	if attr, ok := d.GetOk("manual"); ok {
		config := attr.([]interface{})[0].(map[string]interface{})

		if attr := config["delegate_selectors"].(*schema.Set).List(); len(attr) > 0 {
			connector.GcpKms.DelegateSelectors = utils.InterfaceSliceToStringSlice(attr)
		}

		if attr := config["credentials"].(string); attr != "" {
			connector.GcpKms.Credentials = attr
		}

	}

	if attr, ok := d.GetOk("oidc_authentication"); ok {
		config := attr.([]interface{})[0].(map[string]interface{})
		connector.GcpKms.OidcDetails = &nextgen.GcpOidcDetails{}

		if attr := config["delegate_selectors"].(*schema.Set).List(); len(attr) > 0 {
			connector.GcpKms.DelegateSelectors = utils.InterfaceSliceToStringSlice(attr)
		}

		if attr := config["workload_pool_id"].(string); attr != "" {
			connector.GcpKms.OidcDetails.WorkloadPoolId = attr
		}

		if attr := config["provider_id"].(string); attr != "" {
			connector.GcpKms.OidcDetails.ProviderId = attr
		}

		if attr := config["service_account_email"].(string); attr != "" {
			connector.GcpKms.OidcDetails.ServiceAccountEmail = attr
		}

		if attr := config["gcp_project_id"].(string); attr != "" {
			connector.GcpKms.OidcDetails.GcpProjectId = attr
		}
	}

	if attr, ok := d.GetOk("execute_on_delegate"); ok {
		connector.GcpKms.ExecuteOnDelegate = attr.(bool)
	}

	if attr, ok := d.GetOk("default"); ok {
		connector.GcpKms.Default_ = attr.(bool)
	}

	if attr, ok := d.GetOk("region"); ok {
		connector.GcpKms.Region = attr.(string)
	}

	if attr, ok := d.GetOk("project_id"); ok {
		connector.GcpKms.ProjectId = attr.(string)
	}

	if attr, ok := d.GetOk("key_ring"); ok {
		connector.GcpKms.KeyRing = attr.(string)
	}

	if attr, ok := d.GetOk("key_name"); ok {
		connector.GcpKms.KeyName = attr.(string)
	}

	return connector
}

func readConnectorGcpKms(d *schema.ResourceData, connector *nextgen.ConnectorInfo) error {
	if connector.GcpKms.Credentials != "" {
		d.Set("manual", []map[string]interface{}{
			{
				"credentials":        connector.GcpKms.Credentials,
				"delegate_selectors": connector.GcpKms.DelegateSelectors,
			},
		})
	}

	if connector.GcpKms.OidcDetails != nil {
		d.Set("oidc_authentication", []map[string]interface{}{
			{
				"workload_pool_id":      connector.GcpKms.OidcDetails.WorkloadPoolId,
				"provider_id":           connector.GcpKms.OidcDetails.ProviderId,
				"service_account_email": connector.GcpKms.OidcDetails.ServiceAccountEmail,
				"gcp_project_id":        connector.GcpKms.OidcDetails.GcpProjectId,
				"delegate_selectors":    connector.GcpKms.DelegateSelectors,
			},
		})
	}

	d.Set("execute_on_delegate", connector.GcpKms.ExecuteOnDelegate)
	d.Set("default", connector.GcpKms.Default_)
	d.Set("region", connector.GcpKms.Region)
	d.Set("project_id", connector.GcpKms.ProjectId)
	d.Set("key_ring", connector.GcpKms.KeyRing)
	d.Set("key_name", connector.GcpKms.KeyName)

	return nil
}
