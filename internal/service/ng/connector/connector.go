package connector

import (
	"context"
	"fmt"
	"strings"

	"github.com/antihax/optional"
	"github.com/harness-io/harness-go-sdk/harness/api"
	"github.com/harness-io/harness-go-sdk/harness/nextgen"
	"github.com/harness-io/terraform-provider-harness/internal/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var connectorConfigNames = []string{
	"app_dynamics",
	"artifactory",
	"aws",
	"aws_cloudcost",
	"aws_kms",
	"aws_secret_manager",
	"bitbucket",
	"datadog",
	"docker_registry",
	"dynatrace",
	"gcp",
	"git",
	"github",
	"gitlab",
	"http_helm",
	"jira",
	"k8s_cluster",
	"newrelic",
	"nexus",
	"pagerduty",
	"prometheus",
	"splunk",
	"sumologic",
}

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
			"app_dynamics":       getAppDynamicsSchema(),
			"artifactory":        getArtifactorySchema(),
			"aws":                getAwsSchema(),
			"aws_cloudcost":      getAwsCCSchema(),
			"aws_kms":            getAwsKmsSchema(),
			"aws_secret_manager": getAwsSecretManagerSchema(),
			"bitbucket":          getBitBucketSchema(),
			"datadog":            getDatadogSchema(),
			"docker_registry":    getDockerRegistrySchema(),
			"dynatrace":          getDynatraceSchema(),
			"gcp":                getGcpSchema(),
			"git":                getGitSchema(),
			"github":             getGithubSchema(),
			"gitlab":             getGitlabSchema(),
			"http_helm":          getHttpHelmSchema(),
			"jira":               getJiraSchema(),
			"k8s_cluster":        getK8sClusterSchema(),
			"newrelic":           getNewRelicSchema(),
			"nexus":              getNexusSchema(),
			"pagerduty":          getPagerDutySchema(),
			"prometheus":         getPrometheusSchema(),
			"splunk":             getSplunkSchema(),
			"sumologic":          getSumoLogicSchema(),
		},
	}
}

func resourceConnectorRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*api.Client)

	id := d.Id()
	if id == "" {
		id = d.Get("identifier").(string)
	}

	options := &nextgen.ConnectorsApiGetConnectorOpts{}

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

	resp, _, err := c.NGClient.ConnectorsApi.GetConnector(ctx, c.AccountId, id, options)
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

	connector := nextgen.Connector{Connector: buildConnector(d)}
	options := &nextgen.ConnectorsApiCreateConnectorOpts{}

	resp, _, err := c.NGClient.ConnectorsApi.CreateConnector(ctx, connector, c.AccountId, options)
	if err != nil {
		return diag.Errorf(err.(nextgen.GenericSwaggerError).Error())
	}

	return readConnector(d, resp.Data.Connector)
}

func buildConnector(d *schema.ResourceData) *nextgen.ConnectorInfo {
	connector := &nextgen.ConnectorInfo{}

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

	if attr, ok := d.GetOk("app_dynamics"); ok {
		expandAppDynamicsConfig(attr.([]interface{}), connector)
	}

	if attr, ok := d.GetOk("artifactory"); ok {
		expandArtifactoryConfig(attr.([]interface{}), connector)
	}

	if attr, ok := d.GetOk("aws"); ok {
		expandAwsConfig(attr.([]interface{}), connector)
	}

	if attr, ok := d.GetOk("aws_cloudcost"); ok {
		expandAwsCCConfig(attr.([]interface{}), connector)
	}

	if attr, ok := d.GetOk("aws_kms"); ok {
		expandAwsKmsConfig(attr.([]interface{}), connector)
	}

	if attr, ok := d.GetOk("aws_secret_manager"); ok {
		expandAwsSecretManagerConfig(attr.([]interface{}), connector)
	}

	if attr, ok := d.GetOk("bitbucket"); ok {
		expandBitBucketConfig(attr.([]interface{}), connector)
	}

	if attr, ok := d.GetOk("datadog"); ok {
		expandDatadogConfig(attr.([]interface{}), connector)
	}

	if attr, ok := d.GetOk("docker_registry"); ok {
		expandDockerRegistryConfig(attr.([]interface{}), connector)
	}

	if attr, ok := d.GetOk("dynatrace"); ok {
		expandDynatraceConfig(attr.([]interface{}), connector)
	}

	if attr, ok := d.GetOk("gcp"); ok {
		expandGcpConfig(attr.([]interface{}), connector)
	}

	if attr, ok := d.GetOk("git"); ok {
		expandGitConfig(attr.([]interface{}), connector)
	}

	if attr, ok := d.GetOk("github"); ok {
		expandGithubConfig(attr.([]interface{}), connector)
	}

	if attr, ok := d.GetOk("gitlab"); ok {
		expandGitlabConfig(attr.([]interface{}), connector)
	}

	if attr, ok := d.GetOk("http_helm"); ok {
		expandHttpHelmConfig(attr.([]interface{}), connector)
	}

	if attr, ok := d.GetOk("jira"); ok {
		expandJiraConfig(attr.([]interface{}), connector)
	}

	if attr, ok := d.GetOk("k8s_cluster"); ok {
		expandK8sCluster(attr.([]interface{}), connector)
	}

	if attr, ok := d.GetOk("newrelic"); ok {
		expandNewRelicConfig(attr.([]interface{}), connector)
	}

	if attr, ok := d.GetOk("nexus"); ok {
		expandNexusConfig(attr.([]interface{}), connector)
	}

	if attr, ok := d.GetOk("pagerduty"); ok {
		expandPagerDutyConfig(attr.([]interface{}), connector)
	}

	if attr, ok := d.GetOk("prometheus"); ok {
		expandPrometheusConfig(attr.([]interface{}), connector)
	}

	if attr, ok := d.GetOk("splunk"); ok {
		expandSplunkConfig(attr.([]interface{}), connector)
	}

	if attr, ok := d.GetOk("sumologic"); ok {
		expandSumoLogicConfig(attr.([]interface{}), connector)
	}

	return connector
}

func readConnector(d *schema.ResourceData, connector *nextgen.ConnectorInfo) diag.Diagnostics {
	d.SetId(connector.Identifier)
	d.Set("identifier", connector.Identifier)
	d.Set("description", connector.Description)
	d.Set("name", connector.Name)
	d.Set("org_id", connector.OrgIdentifier)
	d.Set("project_id", connector.ProjectIdentifier)
	d.Set("tags", utils.FlattenTags(connector.Tags))

	if err := flattenAppDynamicsConfig(d, connector); err != nil {
		return diag.FromErr(err)
	}

	if err := flattenArtifactoryConfig(d, connector); err != nil {
		return diag.FromErr(err)
	}

	if err := flattenAwsConfig(d, connector); err != nil {
		return diag.FromErr(err)
	}

	if err := flattenAwsCCConfig(d, connector); err != nil {
		return diag.FromErr(err)
	}

	if err := flattenAwsKmsConfig(d, connector); err != nil {
		return diag.FromErr(err)
	}

	if err := flattenAwsSecretManagerConfig(d, connector); err != nil {
		return diag.FromErr(err)
	}

	if err := flattenBitBucketConfig(d, connector); err != nil {
		return diag.FromErr(err)
	}

	if err := flattenDatadogConfig(d, connector); err != nil {
		return diag.FromErr(err)
	}

	if err := flattenDynatraceConfig(d, connector); err != nil {
		return diag.FromErr(err)
	}

	if err := flattenDockerRegistryConfig(d, connector); err != nil {
		return diag.FromErr(err)
	}

	if err := flattenGcpConfig(d, connector); err != nil {
		return diag.FromErr(err)
	}

	if err := flattenGitConfig(d, connector); err != nil {
		return diag.FromErr(err)
	}

	if err := flattenGithubConfig(d, connector); err != nil {
		return diag.FromErr(err)
	}

	if err := flattenGitlabConfig(d, connector); err != nil {
		return diag.FromErr(err)
	}

	if err := flattenHttpHelmConfig(d, connector); err != nil {
		return diag.FromErr(err)
	}

	if err := flattenJiraConfig(d, connector); err != nil {
		return diag.FromErr(err)
	}

	if err := flattenK8sCluster(d, connector); err != nil {
		return diag.FromErr(err)
	}

	if err := flattenNewRelicConfig(d, connector); err != nil {
		return diag.FromErr(err)
	}

	if err := flattenNexusConfig(d, connector); err != nil {
		return diag.FromErr(err)
	}

	if err := flattenPagerDutyConfig(d, connector); err != nil {
		return diag.FromErr(err)
	}

	if err := flattenPrometheusConfig(d, connector); err != nil {
		return diag.FromErr(err)
	}

	if err := flattenSplunkConfig(d, connector); err != nil {
		return diag.FromErr(err)
	}

	if err := flattenSumoLogicConfig(d, connector); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceConnectorUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*api.Client)

	connector := buildConnector(d)
	options := &nextgen.ConnectorsApiUpdateConnectorOpts{}

	if attr := d.Get("branch").(string); attr != "" {
		options.Branch = optional.NewString(attr)
	}

	if attr := d.Get("repo_id").(string); attr != "" {
		options.RepoIdentifier = optional.NewString(attr)
	}

	resp, _, err := c.NGClient.ConnectorsApi.UpdateConnector(ctx, nextgen.Connector{Connector: connector}, c.AccountId, options)
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
