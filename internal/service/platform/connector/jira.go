package connector

import (
	"context"

	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceConnectorJira() *schema.Resource {
	resource := &schema.Resource{
		Description:   "Resource for creating a Jira connector.",
		ReadContext:   resourceConnectorJiraRead,
		CreateContext: resourceConnectorJiraCreateOrUpdate,
		UpdateContext: resourceConnectorJiraCreateOrUpdate,
		DeleteContext: resourceConnectorDelete,
		Importer:      helpers.MultiLevelResourceImporter,

		Schema: map[string]*schema.Schema{
			"url": {
				Description: "Url of the Jira server.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"username": {
				Description:   "Username to use for authentication.",
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"username_ref"},
				ExactlyOneOf:  []string{"username", "username_ref"},
			},
			"username_ref": {
				Description:   "Reference to a secret containing the username to use for authentication.",
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"username"},
				ExactlyOneOf:  []string{"username", "username_ref"},
			},
			"password_ref": {
				Description: "Reference to a secret containing the password to use for authentication.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"delegate_selectors": {
				Description: "Connect using only the delegates which have these tags.",
				Type:        schema.TypeSet,
				Optional:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
		},
	}

	helpers.SetMultiLevelResourceSchema(resource.Schema)

	return resource
}

func resourceConnectorJiraRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	conn, err := resourceConnectorReadBase(ctx, d, meta, nextgen.ConnectorTypes.Jira)
	if err != nil {
		return err
	}

	if err := readConnectorJira(d, conn); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceConnectorJiraCreateOrUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	conn := buildConnectorJira(d)

	newConn, err := resourceConnectorCreateOrUpdateBase(ctx, d, meta, conn)
	if err != nil {
		return err
	}

	if err := readConnectorJira(d, newConn); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func buildConnectorJira(d *schema.ResourceData) *nextgen.ConnectorInfo {
	connector := &nextgen.ConnectorInfo{
		Type_: nextgen.ConnectorTypes.Jira,
		Jira:  &nextgen.JiraConnector{},
	}

	if attr, ok := d.GetOk("url"); ok {
		connector.Jira.JiraUrl = attr.(string)
	}

	if attr, ok := d.GetOk("username"); ok {
		connector.Jira.Username = attr.(string)
	}

	if attr, ok := d.GetOk("username_ref"); ok {
		connector.Jira.UsernameRef = attr.(string)
	}

	if attr, ok := d.GetOk("password_ref"); ok {
		connector.Jira.PasswordRef = attr.(string)
	}

	if attr, ok := d.GetOk("delegate_selectors"); ok {
		connector.Jira.DelegateSelectors = utils.InterfaceSliceToStringSlice(attr.(*schema.Set).List())
	}

	return connector
}

func readConnectorJira(d *schema.ResourceData, connector *nextgen.ConnectorInfo) error {

	d.Set("url", connector.Jira.JiraUrl)
	d.Set("username", connector.Jira.Username)
	d.Set("username_ref", connector.Jira.UsernameRef)
	d.Set("password_ref", connector.Jira.PasswordRef)
	d.Set("delegate_selectors", connector.Jira.DelegateSelectors)

	return nil
}
