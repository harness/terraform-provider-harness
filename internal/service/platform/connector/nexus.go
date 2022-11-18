package connector

import (
	"context"
	"fmt"
	"strings"

	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceConnectorNexus() *schema.Resource {
	resource := &schema.Resource{
		Description:   "Resource for creating a Nexus connector.",
		ReadContext:   resourceConnectorNexusRead,
		CreateContext: resourceConnectorNexusCreateOrUpdate,
		UpdateContext: resourceConnectorNexusCreateOrUpdate,
		DeleteContext: resourceConnectorDelete,
		Importer:      helpers.MultiLevelResourceImporter,

		Schema: map[string]*schema.Schema{
			"url": {
				Description: "URL of the Nexus server.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"delegate_selectors": {
				Description: "Tags to filter delegates for connection.",
				Type:        schema.TypeSet,
				Optional:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"version": {
				Description: fmt.Sprintf("Version of the Nexus server. Valid values are %s", strings.Join(nextgen.NexusVersionSlice, ", ")),
				Type:        schema.TypeString,
				Required:    true,
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
		},
	}

	helpers.SetMultiLevelResourceSchema(resource.Schema)

	return resource
}

func resourceConnectorNexusRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	conn, err := resourceConnectorReadBase(ctx, d, meta, nextgen.ConnectorTypes.Nexus)
	if err != nil {
		return err
	}

	if err := readConnectorNexus(d, conn); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceConnectorNexusCreateOrUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	conn := buildConnectorNexus(d)

	newConn, err := resourceConnectorCreateOrUpdateBase(ctx, d, meta, conn)
	if err != nil {
		return err
	}

	if err := readConnectorNexus(d, newConn); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func buildConnectorNexus(d *schema.ResourceData) *nextgen.ConnectorInfo {
	connector := &nextgen.ConnectorInfo{
		Type_: nextgen.ConnectorTypes.Nexus,
		Nexus: &nextgen.NexusConnector{
			Auth: &nextgen.NexusAuthentication{
				Type_: nextgen.NexusAuthTypes.Anonymous,
			},
		},
	}

	if attr, ok := d.GetOk("url"); ok {
		connector.Nexus.NexusServerUrl = attr.(string)
	}

	if attr, ok := d.GetOk("version"); ok {
		connector.Nexus.Version = attr.(string)
	}

	if attr, ok := d.GetOk("delegate_selectors"); ok {
		connector.Nexus.DelegateSelectors = utils.InterfaceSliceToStringSlice(attr.(*schema.Set).List())
	}

	if attr, ok := d.GetOk("credentials"); ok {
		config := attr.([]interface{})[0].(map[string]interface{})
		connector.Nexus.Auth.Type_ = nextgen.NexusAuthTypes.UsernamePassword
		connector.Nexus.Auth.UsernamePassword = &nextgen.NexusUsernamePasswordAuth{}

		if attr := config["username"].(string); attr != "" {
			connector.Nexus.Auth.UsernamePassword.Username = attr
		}

		if attr := config["username_ref"].(string); attr != "" {
			connector.Nexus.Auth.UsernamePassword.UsernameRef = attr
		}

		if attr := config["password_ref"].(string); attr != "" {
			connector.Nexus.Auth.UsernamePassword.PasswordRef = attr
		}
	}

	return connector
}

func readConnectorNexus(d *schema.ResourceData, connector *nextgen.ConnectorInfo) error {
	d.Set("url", connector.Nexus.NexusServerUrl)
	d.Set("delegate_selectors", connector.Nexus.DelegateSelectors)
	d.Set("version", connector.Nexus.Version)

	switch connector.Nexus.Auth.Type_ {
	case nextgen.NexusAuthTypes.UsernamePassword:
		d.Set("credentials", []map[string]interface{}{
			{
				"username":     connector.Nexus.Auth.UsernamePassword.Username,
				"username_ref": connector.Nexus.Auth.UsernamePassword.UsernameRef,
				"password_ref": connector.Nexus.Auth.UsernamePassword.PasswordRef,
			},
		})
	case nextgen.NexusAuthTypes.Anonymous:
		// noop
	default:
		return fmt.Errorf("unsupported nexus auth type: %s", connector.Nexus.Auth.Type_)
	}

	return nil
}
