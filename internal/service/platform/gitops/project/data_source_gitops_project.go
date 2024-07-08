package project

import (
	"context"

	"github.com/antihax/optional"
	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceGitOpsProject() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceProjectRead,
		Schema: map[string]*schema.Schema{
			"agent_id": {
				Description: "Agent identifier of the GitOps project.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"account_id": {
				Description: "Account identifier of the GitOps project.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"org_id": {
				Description: "Org identifier of the GitOps project.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"project_id": {
				Description: "Project identifier of the GitOps repository.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"query_name": {
				Description: "Identifier for the GitOps project.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"upsert": {
				Description: "Indicates if the GitOps repository should be updated if existing and inserted if not.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"project": {
				Description: "App projects configuration details.",
				Type:        schema.TypeList,
				Required:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"metadata": {
							Description: "Metadata details that all persisted resources must have.",
							Type:        schema.TypeList,
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"generation": {
										Description: "A sequence number representing a specific generation of the desired state.",
										Type:        schema.TypeString,
										Optional:    true,
									},
									"name": {
										Description: "Name must be unique within a namespace. Is required when creating resources, although some resources may allow a client to request the generation of an appropriate name automatically.",
										Type:        schema.TypeString,
										Optional:    true,
									},
									"namespace": {
										Description: "The namespace where the GitOps project should be created.",
										Type:        schema.TypeString,
										Optional:    true,
									},
								},
							},
						},
						"spec": {
							Description: "Spec is the specification of an AppProject.",
							Type:        schema.TypeList,
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"cluster_resource_whitelist": {
										Description: "ClusterResourceWhitelist contains list of whitelisted cluster level resources.",
										Type:        schema.TypeList,
										Optional:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"group": {
													Description: "Cluster group name.",
													Type:        schema.TypeString,
													Optional:    true,
												},
												"kind": {
													Description: "Cluster kind.",
													Type:        schema.TypeString,
													Optional:    true,
												},
											},
										},
									},
									"destinations": {
										Description: "Destinations contains list of destinations available for deployment.",
										Type:        schema.TypeList,
										Optional:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"namespace": {
													Description: "Namespace specifies the target namespace for the application's resources.",
													Type:        schema.TypeString,
													Optional:    true,
												},
												"server": {
													Description: "Server specifies the URL of the target cluster and must be set to the Kubernetes control plane API.",
													Type:        schema.TypeString,
													Optional:    true,
												},
											},
										},
									},
									"source_repos": {
										Description: "SourceRepos contains list of repository URLs which can be used for deployment.",
										Type:        schema.TypeList,
										Optional:    true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func dataSourceProjectRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)
	var orgIdentifier, projectIdentifier, agentIdentifier, query_name string
	if attr, ok := d.GetOk("org_id"); ok {
		orgIdentifier = attr.(string)
	}
	if attr, ok := d.GetOk("project_id"); ok {
		projectIdentifier = attr.(string)
	}
	if attr, ok := d.GetOk("agent_id"); ok {
		agentIdentifier = attr.(string)
	}

	if attr, ok := d.GetOk("query_name"); ok {
		query_name = attr.(string)
	}

	resp, httpResp, err := c.ProjectGitOpsApi.AgentProjectServiceGet(ctx, agentIdentifier, query_name, c.AccountId, &nextgen.ProjectsApiAgentProjectServiceGetOpts{
		OrgIdentifier:     optional.NewString(orgIdentifier),
		ProjectIdentifier: optional.NewString(projectIdentifier),
	})

	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}
	// Soft delete lookup error handling
	// https://harness.atlassian.net/browse/PL-23765
	if resp.Metadata == nil {
		d.SetId("")
		d.MarkNewResource()
		return nil
	}
	setProjectDetails(d, &resp)

	return nil
}
