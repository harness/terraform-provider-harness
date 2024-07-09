package connector

import (
	"context"

	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceConnectorJDBC() *schema.Resource {
	resource := &schema.Resource{
		Description:   "Resource for creating a JDBC connector.",
		ReadContext:   resourceConnectorJDBCRead,
		CreateContext: resourceConnectorJDBCCreateOrUpdate,
		UpdateContext: resourceConnectorJDBCCreateOrUpdate,
		DeleteContext: resourceConnectorDelete,
		Importer:      helpers.MultiLevelResourceImporter,

		Schema: map[string]*schema.Schema{
			"url": {
				Description: "The URL of the database server.",
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
				Description: "The credentials to use for the database server.",
				Type:        schema.TypeList,
				MaxItems:    1,
				Required:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"username": {
							Description:   "The username to use for the database server.",
							Type:          schema.TypeString,
							Optional:      true,
							ConflictsWith: []string{"credentials.0.username_ref"},
							AtLeastOneOf: []string{
								"credentials.0.username",
								"credentials.0.username_ref",
							},
						},
						"username_ref": {
							Description:   "The reference to the Harness secret containing the username to use for the database server." + secret_ref_text,
							Type:          schema.TypeString,
							Optional:      true,
							ConflictsWith: []string{"credentials.0.username"},
							AtLeastOneOf: []string{
								"credentials.0.username",
								"credentials.0.username_ref",
							},
						},
						"password_ref": {
							Description: "The reference to the Harness secret containing the password to use for the database server." + secret_ref_text,
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

func resourceConnectorJDBCRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	conn, err := resourceConnectorReadBase(ctx, d, meta, nextgen.ConnectorTypes.JDBC)
	if err != nil {
		return err
	}

	if conn == nil {
		return nil
	}

	if err := readConnectorJDBC(d, conn); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceConnectorJDBCCreateOrUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	conn := buildConnectorJDBC(d)

	newConn, err := resourceConnectorCreateOrUpdateBase(ctx, d, meta, conn)
	if err != nil {
		return err
	}

	if err := readConnectorJDBC(d, newConn); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func buildConnectorJDBC(d *schema.ResourceData) *nextgen.ConnectorInfo {
	connector := &nextgen.ConnectorInfo{
		Type_: nextgen.ConnectorTypes.JDBC,
		JDBC: &nextgen.JdbcConnector{
			Url: d.Get("url").(string),
			Auth: &nextgen.JdbcAuthenticationDto{
				Type_:            nextgen.JDBCAuthTypes.UsernamePassword,
				UsernamePassword: &nextgen.JdbcUserNamePasswordDto{},
			},
			// As currently we support through delegate only
			ExecuteOnDelegate: true,
		},
	}

	config := d.Get("credentials").([]interface{})[0].(map[string]interface{})

	if attr, ok := config["username"]; ok {
		connector.JDBC.Auth.UsernamePassword.Username = attr.(string)
	}

	if attr, ok := config["username_ref"]; ok {
		connector.JDBC.Auth.UsernamePassword.UsernameRef = attr.(string)
	}

	if attr, ok := config["password_ref"]; ok {
		connector.JDBC.Auth.UsernamePassword.PasswordRef = attr.(string)
	}

	if attr, ok := d.GetOk("delegate_selectors"); ok {
		connector.JDBC.DelegateSelectors = utils.InterfaceSliceToStringSlice(attr.(*schema.Set).List())
	}

	return connector
}

func readConnectorJDBC(d *schema.ResourceData, connector *nextgen.ConnectorInfo) error {
	d.Set("url", connector.JDBC.Url)
	d.Set("delegate_selectors", connector.JDBC.DelegateSelectors)

	d.Set("credentials", []map[string]interface{}{
		{
			"username":     connector.JDBC.Auth.UsernamePassword.Username,
			"username_ref": connector.JDBC.Auth.UsernamePassword.UsernameRef,
			"password_ref": connector.JDBC.Auth.UsernamePassword.PasswordRef,
		},
	})

	return nil
}
