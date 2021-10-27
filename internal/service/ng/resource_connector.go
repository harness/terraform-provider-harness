package ng

import (
	"context"
	"fmt"
	"strings"

	"github.com/antihax/optional"
	"github.com/harness-io/harness-go-sdk/harness/api"
	"github.com/harness-io/harness-go-sdk/harness/api/nextgen"
	"github.com/harness-io/terraform-provider-harness/internal/service/ng/connectors"
	"github.com/harness-io/terraform-provider-harness/internal/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceConnector() *schema.Resource {
	return &schema.Resource{
		Description:   utils.GetNextgenDescription("Resource for creating a connector."),
		CreateContext: resourceConnectorCreate,
		ReadContext:   resourceConnectorRead,
		UpdateContext: resourceConnectorUpdate,
		DeleteContext: resourceConnectorDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceConnectorImport,
		},

		Schema: map[string]*schema.Schema{
			"identifier": {
				Description: "The unique identifier for the connector.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"name": {
				Description: "The name of the connector.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"description": {
				Description: "The description of the connector.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"org_id": {
				Description: "The unique identifier for the organization.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"project_id": {
				Description:  "The unique identifier for the project.",
				Type:         schema.TypeString,
				Optional:     true,
				RequiredWith: []string{"org_id"},
			},
			"branch": {
				Description: "The branch to use for the connector.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"repo_id": {
				Description: "The unique identifier for the repository.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"tags": {
				Description: "Tags associated with the connector.",
				Type:        schema.TypeSet,
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"aws":             connectors.GetAwsSchema(),
			"docker_registry": connectors.GetDockerRegistrySchema(),
			"gcp":             connectors.GetGcpSchema(),
			"k8s_cluster":     connectors.GetK8sClusterSchema(),
		},
	}
}

func resourceConnectorRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*api.Client)

	id := d.Id()
	if id == "" {
		id = d.Get("identifier").(string)
	}

	options := &nextgen.ConnectorsApiGetConnectorOpts{
		AccountIdentifier: optional.NewString(c.AccountId),
	}

	if attr := d.Get("org_id").(string); attr != "" {
		options.OrgIdentifier = optional.NewString(attr)
	}

	if attr := d.Get("project_id").(string); attr != "" {
		options.ProjectIdentifier = optional.NewString(attr)
	}

	if attr := d.Get("branch").(string); attr != "" {
		options.Branch = optional.NewString(attr)
	}

	if attr := d.Get("repo_id").(string); attr != "" {
		options.RepoIdentifier = optional.NewString(attr)
	}

	resp, _, err := c.NGClient.ConnectorsApi.GetConnector(ctx, id, options)
	if err != nil {
		e := err.(nextgen.GenericSwaggerError)
		if e.Code() == nextgen.ErrorCodes.ResourceNotFound {
			d.SetId("")
			d.MarkNewResource()
			return nil
		}
		return diag.FromErr(err)
	}

	return readConnector(d, resp.Data.Connector)
}

func resourceConnectorCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*api.Client)

	connector := buildConnector(d)
	options := &nextgen.ConnectorsApiCreateConnectorOpts{AccountIdentifier: optional.NewString(c.AccountId)}

	resp, _, err := c.NGClient.ConnectorsApi.CreateConnector(ctx, nextgen.Connector{Connector: connector}, options)
	if err != nil {
		return diag.Errorf(err.(nextgen.GenericSwaggerError).Error())
	}

	return readConnector(d, resp.Data.Connector)
}

func buildConnector(d *schema.ResourceData) *nextgen.ConnectorInfoDto {
	connector := &nextgen.ConnectorInfoDto{}

	if attr := d.Get("name").(string); attr != "" {
		connector.Name = attr
	}

	if attr := d.Get("identifier").(string); attr != "" {
		connector.Identifier = attr
	}

	if attr := d.Get("description").(string); attr != "" {
		connector.Description = attr
	}

	if attr := d.Get("org_id").(string); attr != "" {
		connector.OrgIdentifier = attr
	}

	if attr := d.Get("project_id").(string); attr != "" {
		connector.ProjectIdentifier = attr
	}

	if attr := d.Get("tags").(*schema.Set).List(); len(attr) > 0 {
		connector.Tags = utils.ExpandTags(attr)
	}

	if attr, ok := d.GetOk("aws"); ok {
		connectors.ExpandAwsConfig(attr.([]interface{}), connector)
	}

	if attr, ok := d.GetOk("docker_registry"); ok {
		connectors.ExpandDockerRegistry(attr.([]interface{}), connector)
	}

	if attr, ok := d.GetOk("gcp"); ok {
		connectors.ExpandGcpConfig(attr.([]interface{}), connector)
	}

	if attr, ok := d.GetOk("k8s_cluster"); ok {
		connectors.ExpandK8sCluster(attr.([]interface{}), connector)
	}

	return connector
}

func readConnector(d *schema.ResourceData, connector *nextgen.ConnectorInfoDto) diag.Diagnostics {
	d.SetId(connector.Identifier)
	d.Set("identifier", connector.Identifier)
	d.Set("description", connector.Description)
	d.Set("name", connector.Name)
	d.Set("org_id", connector.OrgIdentifier)
	d.Set("project_id", connector.ProjectIdentifier)
	d.Set("tags", utils.FlattenTags(connector.Tags))

	if err := connectors.FlattenAwsConfig(d, connector); err != nil {
		return diag.FromErr(err)
	}

	if err := connectors.FlattenDockerRegistry(d, connector); err != nil {
		return diag.FromErr(err)
	}

	if err := connectors.FlattenGcpConfig(d, connector); err != nil {
		return diag.FromErr(err)
	}

	if err := connectors.FlattenK8sCluster(d, connector); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceConnectorUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*api.Client)

	connector := buildConnector(d)
	options := &nextgen.ConnectorsApiPutConnectorOpts{AccountIdentifier: optional.NewString(c.AccountId)}

	if attr := d.Get("branch").(string); attr != "" {
		options.Branch = optional.NewString(attr)
	}

	if attr := d.Get("repo_id").(string); attr != "" {
		options.RepoIdentifier = optional.NewString(attr)
	}

	resp, _, err := c.NGClient.ConnectorsApi.PutConnector(ctx, nextgen.Connector{Connector: connector}, options)
	if err != nil {
		return diag.Errorf(err.(nextgen.GenericSwaggerError).Error())
	}

	return readConnector(d, resp.Data.Connector)
}

func resourceConnectorDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*api.Client)

	options := nextgen.ConnectorsApiDeleteConnectorOpts{}

	if attr := d.Get("org_id").(string); attr != "" {
		options.OrgIdentifier = optional.NewString(attr)
	}

	if attr := d.Get("project_id").(string); attr != "" {
		options.ProjectIdentifier = optional.NewString(attr)
	}

	if attr := d.Get("branch").(string); attr != "" {
		options.Branch = optional.NewString(attr)
	}

	if attr := d.Get("repo_id").(string); attr != "" {
		options.RepoIdentifier = optional.NewString(attr)
	}

	_, _, err := c.NGClient.ConnectorsApi.DeleteConnector(ctx, c.AccountId, d.Id(), &options)
	if err != nil {
		return diag.Errorf(err.(nextgen.GenericSwaggerError).Error())
	}

	return nil
}

func resourceConnectorImport(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	// org_id/project_id/connector_id <- Project connector
	// org_id/connector_id <- Organization connector
	// connector_id <- Account connector

	parts := strings.Split(d.Id(), "/")

	partCount := len(parts)
	isAccountConnector := partCount == 1
	isOrgConnector := partCount == 2
	isProjectConnector := partCount == 3

	if isAccountConnector {
		d.SetId(parts[0])
		d.Set("identifier", parts[0])
		return []*schema.ResourceData{d}, nil
	}

	if isOrgConnector {
		d.SetId(parts[1])
		d.Set("identifier", parts[1])
		d.Set("org_id", parts[0])
		return []*schema.ResourceData{d}, nil
	}

	if isProjectConnector {
		d.SetId(parts[2])
		d.Set("identifier", parts[2])
		d.Set("project_id", parts[1])
		d.Set("org_id", parts[0])
		return []*schema.ResourceData{d}, nil
	}

	return nil, fmt.Errorf("invalid connector identifier: %s", d.Id())
}
