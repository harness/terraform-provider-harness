package connector

import (
	"context"
	"fmt"

	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal/gitsync"
	"github.com/harness/terraform-provider-harness/internal/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceConnectorArtifactory() *schema.Resource {
	resource := &schema.Resource{
		Description:   "Resource for creating an Artifactory connector.",
		ReadContext:   resourceConnectorArtifactoryRead,
		CreateContext: resourceConnectorArtifactoryCreateOrUpdate,
		UpdateContext: resourceConnectorArtifactoryCreateOrUpdate,
		DeleteContext: resourceConnectorDelete,
		Importer:      helpers.MultiLevelResourceImporter,
		Schema: map[string]*schema.Schema{
			"url": {
				Description: "URL of the Artifactory server.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"delegate_selectors": {
				Description: "Connect using only the delegates which have these tags.",
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
							Description:   "Reference to a secret containing the username to use for authentication.",
							Type:          schema.TypeString,
							Optional:      true,
							ConflictsWith: []string{"credentials.0.username"},
							ExactlyOneOf:  []string{"credentials.0.username", "credentials.0.username_ref"},
						},
						"password_ref": {
							Description: "Reference to a secret containing the password to use for authentication.",
							Type:        schema.TypeString,
							Required:    true,
						},
					},
				},
			},
		},
	}

	helpers.SetMultiLevelResourceSchema(resource.Schema)
	gitsync.SetGitSyncSchema(resource.Schema, false)

	return resource
}

func resourceConnectorArtifactoryRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	conn, err := resourceConnectorReadBase(ctx, d, meta, nextgen.ConnectorTypes.Artifactory)
	if err != nil {
		return err
	}

	if err := readConnectorArtifactory(d, conn); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceConnectorArtifactoryCreateOrUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	conn := buildConnectorArtifactory(d)

	newConn, err := resourceConnectorCreateOrUpdateBase(ctx, d, meta, conn)
	if err != nil {
		return err
	}

	if err := readConnectorArtifactory(d, newConn); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func buildConnectorArtifactory(d *schema.ResourceData) *nextgen.ConnectorInfo {
	connector := &nextgen.ConnectorInfo{
		Type_: nextgen.ConnectorTypes.Artifactory,
		Artifactory: &nextgen.ArtifactoryConnector{
			Auth: &nextgen.ArtifactoryAuthentication{
				Type_: nextgen.ArtifactoryAuthTypes.Anonymous,
			},
		},
	}

	if attr, ok := d.GetOk("url"); ok {
		connector.Artifactory.ArtifactoryServerUrl = attr.(string)
	}

	if attr, ok := d.GetOk("delegate_selectors"); ok {
		connector.Artifactory.DelegateSelectors = utils.InterfaceSliceToStringSlice(attr.(*schema.Set).List())
	}

	if attr, ok := d.GetOk("credentials"); ok {
		config := attr.([]interface{})[0].(map[string]interface{})
		connector.Artifactory.Auth.Type_ = nextgen.ArtifactoryAuthTypes.UsernamePassword
		connector.Artifactory.Auth.UsernamePassword = &nextgen.ArtifactoryUsernamePasswordAuth{}

		if attr, ok := config["username"]; ok {
			connector.Artifactory.Auth.UsernamePassword.Username = attr.(string)
		}

		if attr, ok := config["username_ref"]; ok {
			connector.Artifactory.Auth.UsernamePassword.UsernameRef = attr.(string)
		}

		if attr, ok := config["password_ref"]; ok {
			connector.Artifactory.Auth.UsernamePassword.PasswordRef = attr.(string)
		}
	}

	return connector
}

func readConnectorArtifactory(d *schema.ResourceData, connector *nextgen.ConnectorInfo) error {
	d.Set("url", connector.Artifactory.ArtifactoryServerUrl)
	d.Set("delegate_selectors", connector.Artifactory.DelegateSelectors)

	switch connector.Artifactory.Auth.Type_ {
	case nextgen.ArtifactoryAuthTypes.UsernamePassword:
		d.Set("credentials", []map[string]interface{}{
			{
				"username":     connector.Artifactory.Auth.UsernamePassword.Username,
				"username_ref": connector.Artifactory.Auth.UsernamePassword.UsernameRef,
				"password_ref": connector.Artifactory.Auth.UsernamePassword.PasswordRef,
			},
		})
	case nextgen.ArtifactoryAuthTypes.Anonymous:
		// noop
	default:
		return fmt.Errorf("unsupported artifactory auth type: %s", connector.Artifactory.Auth.Type_)
	}

	return nil
}
