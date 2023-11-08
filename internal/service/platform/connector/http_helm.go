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

func ResourceConnectorHelm() *schema.Resource {
	resource := &schema.Resource{
		Description:   "Resource for creating a HTTP Helm connector.",
		ReadContext:   resourceConnectorHelmRead,
		CreateContext: resourceConnectorHelmCreateOrUpdate,
		UpdateContext: resourceConnectorHelmCreateOrUpdate,
		DeleteContext: resourceConnectorDelete,
		Importer:      helpers.MultiLevelResourceImporter,

		Schema: map[string]*schema.Schema{
			"url": {
				Description: "URL of the helm server.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"delegate_selectors": {
				Description: "Tags to filter delegates for connection.",
				Type:        schema.TypeSet,
				Optional:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"credentials": {
				Description: "Credentials to use for authentication.",
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"username": {
							Description:   "Username to use for authentication.",
							Type:          schema.TypeString,
							Optional:      true,
							ConflictsWith: []string{"credentials.0.username_ref"},
							ExactlyOneOf:  []string{"credentials.0.username", "credentials.0.username_ref"},
						},
						"username_ref": {
							Description:   "Reference to a secret containing the username to use for authentication." + secret_ref_text,
							Type:          schema.TypeString,
							Optional:      true,
							ConflictsWith: []string{"credentials.0.username"},
							ExactlyOneOf:  []string{"credentials.0.username", "credentials.0.username_ref"},
						},
						"password_ref": {
							Description: "Reference to a secret containing the password to use for authentication." + secret_ref_text,
							Type:        schema.TypeString,
							Required:    true,
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
		},
	}

	helpers.SetMultiLevelResourceSchema(resource.Schema)

	return resource
}

func resourceConnectorHelmRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	conn, err := resourceConnectorReadBase(ctx, d, meta, nextgen.ConnectorTypes.HttpHelmRepo)
	if err != nil {
		return err
	}

	if conn == nil {
		return nil
	}

	if err := readConnectorHelm(d, conn); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceConnectorHelmCreateOrUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	conn := buildConnectorHelm(d)

	newConn, err := resourceConnectorCreateOrUpdateBase(ctx, d, meta, conn)
	if err != nil {
		return err
	}

	if err := readConnectorHelm(d, newConn); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func buildConnectorHelm(d *schema.ResourceData) *nextgen.ConnectorInfo {
	connector := &nextgen.ConnectorInfo{
		Type_:    nextgen.ConnectorTypes.HttpHelmRepo,
		HttpHelm: &nextgen.HttpHelmConnector{},
	}

	if attr, ok := d.GetOk("url"); ok {
		connector.HttpHelm.HelmRepoUrl = attr.(string)
	}

	if attr, ok := d.GetOk("delegate_selectors"); ok {
		connector.HttpHelm.DelegateSelectors = utils.InterfaceSliceToStringSlice(attr.(*schema.Set).List())
	}

	connector.HttpHelm.Auth = &nextgen.HttpHelmAuthentication{
		Type_: nextgen.HttpHelmAuthTypes.Anonymous,
	}

	if attr, ok := d.GetOk("credentials"); ok {
		config := attr.([]interface{})[0].(map[string]interface{})
		connector.HttpHelm.Auth.Type_ = nextgen.HttpHelmAuthTypes.UsernamePassword
		connector.HttpHelm.Auth.UsernamePassword = &nextgen.HttpHelmUsernamePassword{}

		if attr, ok := d.GetOk("credentials.0.username"); ok {
			connector.HttpHelm.Auth.UsernamePassword.Username = attr.(string)
		}

		if attr, ok := config["credentials.0.username_ref"]; ok {
			connector.HttpHelm.Auth.UsernamePassword.UsernameRef = attr.(string)
		}

		if attr, ok := d.GetOk("credentials.0.password_ref"); ok {
			connector.HttpHelm.Auth.UsernamePassword.PasswordRef = attr.(string)
		}
	}

	return connector
}

func readConnectorHelm(d *schema.ResourceData, connector *nextgen.ConnectorInfo) error {
	d.Set("url", connector.HttpHelm.HelmRepoUrl)
	d.Set("delegate_selectors", connector.HttpHelm.DelegateSelectors)

	switch connector.HttpHelm.Auth.Type_ {
	case nextgen.HttpHelmAuthTypes.UsernamePassword:
		d.Set("credentials", []map[string]interface{}{
			{
				"username":     connector.HttpHelm.Auth.UsernamePassword.Username,
				"username_ref": connector.HttpHelm.Auth.UsernamePassword.UsernameRef,
				"password_ref": connector.HttpHelm.Auth.UsernamePassword.PasswordRef,
			},
		})
	case nextgen.HttpHelmAuthTypes.Anonymous:
		// noop
	default:
		return fmt.Errorf("unsupported http helm auth type: %s", connector.HttpHelm.Auth.Type_)
	}

	return nil
}
