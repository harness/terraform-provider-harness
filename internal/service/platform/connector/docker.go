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

func ResourceConnectorDocker() *schema.Resource {
	resource := &schema.Resource{
		Description:   "Resource for creating a Docker connector.",
		ReadContext:   resourceConnectorDockerRead,
		CreateContext: resourceConnectorDockerCreateOrUpdate,
		UpdateContext: resourceConnectorDockerCreateOrUpdate,
		DeleteContext: resourceConnectorDelete,
		Importer:      helpers.MultiLevelResourceImporter,

		Schema: map[string]*schema.Schema{
			"type": {
				Description: fmt.Sprintf("The type of the docker registry. Valid options are %s", strings.Join(nextgen.DockerRegistryTypesSlice, ", ")),
				Type:        schema.TypeString,
				Required:    true,
			},
			"url": {
				Description: "The URL of the docker registry.",
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
				Description: "The credentials to use for the docker registry. If not specified then the connection is made to the registry anonymously.",
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"username": {
							Description:   "The username to use for the docker registry.",
							Type:          schema.TypeString,
							Optional:      true,
							ConflictsWith: []string{"credentials.0.username_ref"},
							AtLeastOneOf: []string{
								"credentials.0.username",
								"credentials.0.username_ref",
							},
						},
						"username_ref": {
							Description:   "The reference to the Harness secret containing the username to use for the docker registry." + secret_ref_text,
							Type:          schema.TypeString,
							Optional:      true,
							ConflictsWith: []string{"credentials.0.username"},
							AtLeastOneOf: []string{
								"credentials.0.username",
								"credentials.0.username_ref",
							},
						},
						"password_ref": {
							Description: "The reference to the Harness secret containing the password to use for the docker registry." + secret_ref_text,
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

func resourceConnectorDockerRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	conn, err := resourceConnectorReadBase(ctx, d, meta, nextgen.ConnectorTypes.DockerRegistry)
	if err != nil {
		return err
	}

	if err := readConnectorDocker(d, conn); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceConnectorDockerCreateOrUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	conn := buildConnectorDocker(d)

	newConn, err := resourceConnectorCreateOrUpdateBase(ctx, d, meta, conn)
	if err != nil {
		return err
	}

	if err := readConnectorDocker(d, newConn); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func buildConnectorDocker(d *schema.ResourceData) *nextgen.ConnectorInfo {
	connector := &nextgen.ConnectorInfo{
		Type_: nextgen.ConnectorTypes.DockerRegistry,
		DockerRegistry: &nextgen.DockerConnector{
			Auth: &nextgen.DockerAuthentication{
				Type_: nextgen.DockerAuthTypes.Anonymous,
			},
		},
	}

	if attr, ok := d.GetOk("url"); ok {
		connector.DockerRegistry.DockerRegistryUrl = attr.(string)
	}

	if attr, ok := d.GetOk("type"); ok {
		connector.DockerRegistry.ProviderType = attr.(string)
	}

	if attr, ok := d.GetOk("delegate_selectors"); ok {
		connector.DockerRegistry.DelegateSelectors = utils.InterfaceSliceToStringSlice(attr.(*schema.Set).List())
	}

	if attr, ok := d.GetOk("credentials"); ok {
		config := attr.([]interface{})[0].(map[string]interface{})
		connector.DockerRegistry.Auth.Type_ = nextgen.DockerAuthTypes.UsernamePassword
		connector.DockerRegistry.Auth.UsernamePassword = &nextgen.DockerUserNamePassword{}

		if attr, ok := config["username"]; ok {
			connector.DockerRegistry.Auth.UsernamePassword.Username = attr.(string)
		}

		if attr, ok := config["username_ref"]; ok {
			connector.DockerRegistry.Auth.UsernamePassword.UsernameRef = attr.(string)
		}

		if attr, ok := config["password_ref"]; ok {
			connector.DockerRegistry.Auth.UsernamePassword.PasswordRef = attr.(string)
		}
	}

	return connector
}

func readConnectorDocker(d *schema.ResourceData, connector *nextgen.ConnectorInfo) error {
	d.Set("type", connector.DockerRegistry.ProviderType)
	d.Set("url", connector.DockerRegistry.DockerRegistryUrl)
	d.Set("delegate_selectors", connector.DockerRegistry.DelegateSelectors)

	switch connector.DockerRegistry.Auth.Type_ {
	case nextgen.DockerAuthTypes.UsernamePassword:
		d.Set("credentials", []map[string]interface{}{
			{
				"username":     connector.DockerRegistry.Auth.UsernamePassword.Username,
				"username_ref": connector.DockerRegistry.Auth.UsernamePassword.UsernameRef,
				"password_ref": connector.DockerRegistry.Auth.UsernamePassword.PasswordRef,
			},
		})
	case nextgen.DockerAuthTypes.Anonymous:
		// noop
	default:
		return fmt.Errorf("unsupported docker registry auth type: %s", connector.DockerRegistry.Auth.Type_)
	}

	return nil
}
