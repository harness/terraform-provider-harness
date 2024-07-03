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

func ResourceProject() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceProjectCreate,
		ReadContext:   resourceProjectRead,
		UpdateContext: resourceProjectUpdate,
		DeleteContext: resourceProjectDelete,
		Importer:      helpers.GitopsAgentProjectImporter,
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
			"query_name": {
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
				Required: true,
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

func resourceProjectCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)
	var orgIdentifier, projectIdentifier, agentIdentifier string
	var upsert bool
	if attr, ok := d.GetOk("org_id"); ok {
		orgIdentifier = attr.(string)
	}
	if attr, ok := d.GetOk("project_id"); ok {
		projectIdentifier = attr.(string)
	}
	if attr, ok := d.GetOk("agent_id"); ok {
		agentIdentifier = attr.(string)
	}

	if attr, ok := d.GetOk("upsert"); ok {
		upsert = attr.(bool)
	}

	projectData := createRequestBody(d)
	projectData.Upsert = upsert

	resp, httpResp, err := c.ProjectGitOpsApi.AgentProjectServiceCreate(ctx, projectData, c.AccountId, agentIdentifier, &nextgen.ProjectsApiAgentProjectServiceCreateOpts{
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

func resourceProjectRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	// if attr, ok := d.GetOk("query_name"); ok {
	// 	query_name = attr.(string)
	// }

	if v, ok := d.GetOk("project"); ok {
		for _, item := range v.([]interface{}) {
			projectData := item.(map[string]interface{})

			if md, ok := projectData["metadata"].([]interface{}); ok {
				mdData := md[0].(map[string]interface{})
				query_name = mdData["name"].(string)
			}
		}
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

func resourceProjectUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)
	var orgIdentifier, projectIdentifier, agentIdentifier string
	if attr, ok := d.GetOk("org_id"); ok {
		orgIdentifier = attr.(string)
	}
	if attr, ok := d.GetOk("project_id"); ok {
		projectIdentifier = attr.(string)
	}
	if attr, ok := d.GetOk("agent_id"); ok {
		agentIdentifier = attr.(string)
	}

	projectData := updateRequestBody(d)

	resp, httpResp, err := c.ProjectGitOpsApi.AgentProjectServiceUpdate(ctx, projectData, c.AccountId, agentIdentifier, projectData.Project.Metadata.Name, &nextgen.ProjectsApiAgentProjectServiceUpdateOpts{
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

func resourceProjectDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	_, httpResp, err := c.ProjectGitOpsApi.AgentProjectServiceDelete(ctx, agentIdentifier, query_name, c.AccountId, orgIdentifier, &nextgen.ProjectsApiAgentProjectServiceDeleteOpts{
		ProjectIdentifier: optional.NewString(projectIdentifier),
	})

	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	return nil
}

// Function to map Terraform schema to internal data structure
func updateRequestBody(d *schema.ResourceData) nextgen.ProjectsProjectUpdateRequest {
	var projectsProjectCreateRequest nextgen.ProjectsProjectUpdateRequest
	var appprojectsAppProject *nextgen.AppprojectsAppProject
	var approjectsAppProjectSpec *nextgen.AppprojectsAppProjectSpec
	var v1ObjectMeta *nextgen.V1ObjectMeta

	if v, ok := d.GetOk("project"); ok {
		for _, item := range v.([]interface{}) {
			projectData := item.(map[string]interface{})

			if md, ok := projectData["metadata"].([]interface{}); ok {
				mdData := md[0].(map[string]interface{})
				v1ObjectMeta = &nextgen.V1ObjectMeta{
					Generation: mdData["generation"].(string),
					Name:       mdData["name"].(string),
					Namespace:  mdData["namespace"].(string),
				}
			}

			var v1GroupKind []nextgen.V1GroupKind
			if crw, ok := projectData["spec"].([]interface{}); ok && len(crw) > 0 {
				specData := crw[0].(map[string]interface{})

				if crwList, ok := specData["cluster_resource_whitelist"].([]interface{}); ok {
					for _, w := range crwList {
						wData := w.(map[string]interface{})
						v1GroupKind = append(v1GroupKind, nextgen.V1GroupKind{
							Group: wData["group"].(string),
							Kind:  wData["kind"].(string),
						})
					}
				}
			}

			var appprojectsApplicationDestination []nextgen.AppprojectsApplicationDestination
			if dest, ok := projectData["spec"].([]interface{}); ok && len(dest) > 0 {
				specData := dest[0].(map[string]interface{})

				if destList, ok := specData["destinations"].([]interface{}); ok {
					for _, d := range destList {
						dData := d.(map[string]interface{})
						appprojectsApplicationDestination = append(appprojectsApplicationDestination, nextgen.AppprojectsApplicationDestination{
							Namespace: dData["namespace"].(string),
							Server:    dData["server"].(string),
						})
					}
				}
			}

			var sourceRepos []string
			if sr, ok := projectData["spec"].([]interface{}); ok && len(sr) > 0 {
				specData := sr[0].(map[string]interface{})

				if repos, ok := specData["source_repos"].([]interface{}); ok {
					for _, repo := range repos {
						sourceRepos = append(sourceRepos, repo.(string))
					}
				}
			}

			approjectsAppProjectSpec = &nextgen.AppprojectsAppProjectSpec{
				Destinations:             appprojectsApplicationDestination,
				ClusterResourceWhitelist: v1GroupKind,
				SourceRepos:              sourceRepos,
			}

			appprojectsAppProject = &nextgen.AppprojectsAppProject{
				Metadata: v1ObjectMeta,
				Spec:     approjectsAppProjectSpec,
			}

			projectsProjectCreateRequest = nextgen.ProjectsProjectUpdateRequest{
				Project: appprojectsAppProject,
			}

		}
	}

	return projectsProjectCreateRequest
}

// Function to map Terraform schema to internal data structure
func createRequestBody(d *schema.ResourceData) nextgen.ProjectsProjectCreateRequest {
	var projectsProjectCreateRequest nextgen.ProjectsProjectCreateRequest
	var appprojectsAppProject *nextgen.AppprojectsAppProject
	var approjectsAppProjectSpec *nextgen.AppprojectsAppProjectSpec
	var v1ObjectMeta *nextgen.V1ObjectMeta

	if v, ok := d.GetOk("project"); ok {
		for _, item := range v.([]interface{}) {
			projectData := item.(map[string]interface{})

			if md, ok := projectData["metadata"].([]interface{}); ok {
				mdData := md[0].(map[string]interface{})
				v1ObjectMeta = &nextgen.V1ObjectMeta{
					Generation: mdData["generation"].(string),
					Name:       mdData["name"].(string),
					Namespace:  mdData["namespace"].(string),
				}
			}

			var v1GroupKind []nextgen.V1GroupKind
			if crw, ok := projectData["spec"].([]interface{}); ok && len(crw) > 0 {
				specData := crw[0].(map[string]interface{})

				if crwList, ok := specData["cluster_resource_whitelist"].([]interface{}); ok {
					for _, w := range crwList {
						wData := w.(map[string]interface{})
						v1GroupKind = append(v1GroupKind, nextgen.V1GroupKind{
							Group: wData["group"].(string),
							Kind:  wData["kind"].(string),
						})
					}
				}
			}

			var appprojectsApplicationDestination []nextgen.AppprojectsApplicationDestination
			if dest, ok := projectData["spec"].([]interface{}); ok && len(dest) > 0 {
				specData := dest[0].(map[string]interface{})

				if destList, ok := specData["destinations"].([]interface{}); ok {
					for _, d := range destList {
						dData := d.(map[string]interface{})
						appprojectsApplicationDestination = append(appprojectsApplicationDestination, nextgen.AppprojectsApplicationDestination{
							Namespace: dData["namespace"].(string),
							Server:    dData["server"].(string),
						})
					}
				}
			}

			var sourceRepos []string
			if sr, ok := projectData["spec"].([]interface{}); ok && len(sr) > 0 {
				specData := sr[0].(map[string]interface{})

				if repos, ok := specData["source_repos"].([]interface{}); ok {
					for _, repo := range repos {
						sourceRepos = append(sourceRepos, repo.(string))
					}
				}
			}

			approjectsAppProjectSpec = &nextgen.AppprojectsAppProjectSpec{
				Destinations:             appprojectsApplicationDestination,
				ClusterResourceWhitelist: v1GroupKind,
				SourceRepos:              sourceRepos,
			}

			appprojectsAppProject = &nextgen.AppprojectsAppProject{
				Metadata: v1ObjectMeta,
				Spec:     approjectsAppProjectSpec,
			}

			projectsProjectCreateRequest = nextgen.ProjectsProjectCreateRequest{
				Project: appprojectsAppProject,
			}

		}
	}

	return projectsProjectCreateRequest
}

func setProjectDetails(d *schema.ResourceData, projects *nextgen.AppprojectsAppProject) {
	d.SetId(projects.Metadata.Name)
	// d.Set("query_name", projects.Metadata.Name)
	d.Set("account_id", "1bvyLackQK-Hapk25-Ry4w")
	if projects.Metadata != nil {
		metadataList := []interface{}{}
		metadata := map[string]interface{}{}
		metadata["name"] = projects.Metadata.Name
		metadata["namespace"] = projects.Metadata.Namespace
		metadata["generation"] = projects.Metadata.Generation
		metadata["uid"] = projects.Metadata.Uid
		metadataList = append(metadataList, metadata)
		specdataList := []interface{}{}
		spec := map[string]interface{}{}
		spec["sourceRepos"] = projects.Spec.SourceRepos
		spec["destinations"] = projects.Spec.Destinations
		spec["cluster_resource_whitelist"] = projects.Spec.ClusterResourceWhitelist
		specdataList = append(specdataList, spec)
		projectList := []interface{}{}
		project := map[string]interface{}{}
		project["metadata"] = metadataList
		project["spec"] = specdataList
		// project["status"] = projects.Status
		projectList = append(projectList, project)
		d.Set("project", projectList)
	}
}
