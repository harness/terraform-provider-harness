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
				Type:     schema.TypeString,
				Required: true,
			},
			"account_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"org_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"project_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"upsert": {
				Description: "Indicates if the GitOps repository should be updated if existing and inserted if not.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"project": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"metadata": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"generation": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"name": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"namespace": {
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
						"spec": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"cluster_resource_whitelist": {
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"group": {
													Type:     schema.TypeString,
													Optional: true,
												},
												"kind": {
													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},
									"destinations": {
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"namespace": {
													Type:     schema.TypeString,
													Optional: true,
												},
												"server": {
													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},
									"source_repos": {
										Type:     schema.TypeList,
										Optional: true,
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
