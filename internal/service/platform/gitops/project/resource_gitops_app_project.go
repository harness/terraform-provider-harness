package project

import (
	"context"
	"fmt"

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
				Description: "Agent identifier of the GitOps project. Project is created on agent scope.",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
			},
			"account_id": {
				Description: "Account identifier of the GitOps project/agent.",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
			},
			"org_id": {
				Description: "Org identifier of the GitOps agent for which project is created.",
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
			},
			"project_id": {
				Description: "Project identifier of the GitOps agent for which project is created.",
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
			},
			"query_name": {
				Description: "Identifier for the GitOps project.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"upsert": {
				Description: "Indicates if the GitOps project should be updated if existing and inserted if not.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"project": {
				Type:        schema.TypeList,
				Required:    true,
				Description: "GitOps project configuration.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"metadata": {
							Type:        schema.TypeList,
							Required:    true,
							Description: "Metadata details for the GitOps project.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Name of the GitOps project.",
									},
									"namespace": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Namespace of the GitOps project. This must match agent namespace.",
									},
									"resource_version": {
										Type:        schema.TypeString,
										Computed:    true,
										Optional:    true,
										Description: "Resource Version for the GitOps project",
									},
									"labels": {
										Type:        schema.TypeMap,
										Optional:    true,
										Description: "Labels associated with the GitOps project.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"annotations": {
										Type:        schema.TypeMap,
										Optional:    true,
										Computed:    true,
										Description: "Annotations associated with the GitOps project.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"finalizers": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "Finalizers associated with the GitOps project.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"cluster_name": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Name of the cluster associated with the GitOps project.",
									},
									"managed_fields": {
										Type:        schema.TypeList,
										Optional:    true,
										Computed:    true,
										Description: "Managed fields associated with the GitOps project.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"manager": {
													Type:        schema.TypeString,
													Optional:    true,
													Computed:    true,
													Description: "Manager responsible for the operation.",
												},
												"operation": {
													Type:        schema.TypeString,
													Optional:    true,
													Computed:    true,
													Description: "Operation type performed on the GitOps project.",
												},
												"api_version": {
													Type:        schema.TypeString,
													Optional:    true,
													Computed:    true,
													Description: "API version of the operation performed.",
												},
												"time": {
													Type:        schema.TypeMap,
													Optional:    true,
													Computed:    true,
													Description: "Timestamp of the operation.",
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"fields_type": {
													Type:        schema.TypeString,
													Optional:    true,
													Computed:    true,
													Description: "Type of the fields in the GitOps project.",
												},
												"fields_v1": {
													Type:        schema.TypeMap,
													Optional:    true,
													Computed:    true,
													Description: "Raw fields associated with the GitOps project.",
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"subresource": {
													Type:        schema.TypeString,
													Optional:    true,
													Computed:    true,
													Description: "Subresource associated with the GitOps project.",
												},
											},
										},
									},
								},
							},
						},
						"spec": {
							Type:        schema.TypeList,
							Required:    true,
							Description: "Specification details for the GitOps project.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"source_repos": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "Source repositories for the GitOps project.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"destinations": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "Destinations for deployment of the GitOps project.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"server": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "Server URL of the destination.",
												},
												"namespace": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "Namespace of the destination.",
												},
												"name": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "Name of the destination.",
												},
											},
										},
									},
									"description": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Description of the GitOps project.",
									},
									"roles": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "Roles associated with the GitOps project.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"name": {
													Type:        schema.TypeString,
													Required:    true,
													Description: "Name of the role.",
												},
												"description": {
													Type:        schema.TypeString,
													Required:    true,
													Description: "Description of the role.",
												},
												"policies": {
													Type:        schema.TypeList,
													Optional:    true,
													Description: "Policies associated with the role.",
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"jwt_tokens": {
													Type:        schema.TypeList,
													Optional:    true,
													Description: "JWT tokens associated with the role.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"iat": {
																Type:        schema.TypeString,
																Optional:    true,
																Description: "Issued At time of the JWT token.",
															},
															"exp": {
																Type:        schema.TypeString,
																Optional:    true,
																Description: "Expiration time of the JWT token.",
															},
															"id": {
																Type:        schema.TypeString,
																Optional:    true,
																Description: "ID of the JWT token.",
															},
														},
													},
												},
												"groups": {
													Type:        schema.TypeList,
													Optional:    true,
													Description: "Groups associated with the role.",
													Computed:    true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
											},
										},
									},
									"cluster_resource_whitelist": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "Cluster resource whitelist for the GitOps project.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"group": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "Group of the cluster resource whitelist.",
												},
												"kind": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "Kind of the cluster resource whitelist.",
												},
											},
										},
									},
									"namespace_resource_blacklist": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "Namespace resource blacklist for the GitOps project.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"group": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "Group of the namespace resource blacklist.",
												},
												"kind": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "Kind of the namespace resource blacklist.",
												},
											},
										},
									},
									"orphaned_resources": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "Orphaned resources configuration for the GitOps project.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"warn": {
													Type:        schema.TypeBool,
													Optional:    true,
													Description: "Whether to warn about orphaned resources.",
												},
												"ignore": {
													Type:        schema.TypeList,
													Optional:    true,
													Description: "List of ignored orphaned resources.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"group": {
																Type:        schema.TypeString,
																Optional:    true,
																Description: "Group of the ignored orphaned resource.",
															},
															"kind": {
																Type:        schema.TypeString,
																Optional:    true,
																Description: "Kind of the ignored orphaned resource.",
															},
															"name": {
																Type:        schema.TypeString,
																Optional:    true,
																Description: "Name of the ignored orphaned resource.",
															},
														},
													},
												},
											},
										},
									},
									"sync_windows": {
										Type:        schema.TypeList,
										Optional:    true,
										Computed:    true,
										Description: "Synchronization windows for the GitOps project.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"kind": {
													Type:        schema.TypeString,
													Optional:    true,
													Computed:    true,
													Description: "Kind of synchronization window.",
												},
												"schedule": {
													Type:        schema.TypeString,
													Optional:    true,
													Computed:    true,
													Description: "Schedule of synchronization window.",
												},
												"duration": {
													Type:        schema.TypeString,
													Optional:    true,
													Computed:    true,
													Description: "Duration of synchronization window.",
												},
												"applications": {
													Type:        schema.TypeList,
													Optional:    true,
													Computed:    true,
													Description: "Applications associated with synchronization window.",
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"namespaces": {
													Type:        schema.TypeList,
													Optional:    true,
													Computed:    true,
													Description: "Namespaces associated with synchronization window.",
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"clusters": {
													Type:        schema.TypeList,
													Optional:    true,
													Computed:    true,
													Description: "Clusters associated with synchronization window.",
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"manual_sync": {
													Type:        schema.TypeBool,
													Optional:    true,
													Computed:    true,
													Description: "Whether manual synchronization is enabled.",
												},
												"time_zone": {
													Type:        schema.TypeString,
													Optional:    true,
													Computed:    true,
													Description: "Time zone of synchronization window.",
												},
											},
										},
									},
									"namespace_resource_whitelist": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "Namespace resource whitelist for the GitOps project.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"group": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "Group of the namespace resource whitelist.",
												},
												"kind": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "Kind of the namespace resource whitelist.",
												},
											},
										},
									},
									"signature_keys": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "Signature keys for the GitOps project.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"key_id": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "ID of the signature key.",
												},
											},
										},
									},
									"cluster_resource_blacklist": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "Cluster resource blacklist for the GitOps project.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"group": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "Group of the cluster resource blacklist.",
												},
												"kind": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "Kind of the cluster resource blacklist.",
												},
											},
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

	if attr, ok := d.GetOk("query_name"); ok {
		query_name = attr.(string)
	}

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
		return helpers.HandleReadApiError(err, d, httpResp)
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

	if v, ok := d.GetOk("project"); ok {
		for _, item := range v.([]interface{}) {
			projectData := item.(map[string]interface{})

			if md, ok := projectData["metadata"].([]interface{}); ok {
				mdData := md[0].(map[string]interface{})
				query_name = mdData["name"].(string)
			}
		}
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
	var projectsProjectUpdateRequest nextgen.ProjectsProjectUpdateRequest
	var appprojectsAppProject *nextgen.AppprojectsAppProject
	var approjectsAppProjectSpec *nextgen.AppprojectsAppProjectSpec
	var v1ObjectMeta *nextgen.V1ObjectMeta
	var sourceRepos []string

	if v, ok := d.GetOk("project"); ok {
		for _, item := range v.([]interface{}) {
			projectData := item.(map[string]interface{})

			if md, ok := projectData["metadata"].([]interface{}); ok {
				mdData := md[0].(map[string]interface{})
				var v1ManagedFieldsEntry []nextgen.V1ManagedFieldsEntry
				if managerFieldsList, ok := mdData["managed_fields"].([]interface{}); ok {
					for _, w := range managerFieldsList {
						wData := w.(map[string]interface{})
						v1ManagedFieldsEntry = append(v1ManagedFieldsEntry, nextgen.V1ManagedFieldsEntry{
							Manager:     wData["api_version"].(string),
							Operation:   wData["operation"].(string),
							ApiVersion:  wData["api_version"].(string),
							FieldsType:  wData["fields_type"].(string),
							Subresource: wData["subresource"].(string),
						})
					}
				}

				var namespace, name, clusterName, resource_version string
				var finalizers []interface{}
				var labels, annotations map[string]interface{}
				if nil != mdData["namespace"] {
					namespace = mdData["namespace"].(string)
				}
				if nil != mdData["name"] {
					name = mdData["name"].(string)
				}

				if nil != mdData["resource_version"] {
					resource_version = mdData["resource_version"].(string)
				}

				if nil != mdData["cluster_name"] {
					clusterName = mdData["cluster_name"].(string)
				}

				if f, ok := mdData["finalizers"].([]interface{}); ok {
					for _, r := range f {
						finalizers = append(finalizers, r.(string))
					}
				}

				if nil != mdData["labels"] {
					labels = mdData["labels"].(map[string]interface{})
				}
				if nil != mdData["annotations"] {
					annotations = mdData["annotations"].(map[string]interface{})
				}

				s := make([]string, len(finalizers))
				for i, v := range finalizers {
					s[i] = fmt.Sprint(v)
				}

				var annotationsStr, labelsStr map[string]string
				annotationsStr = make(map[string]string)
				labelsStr = make(map[string]string)

				for key, value := range annotations {
					strKey := fmt.Sprintf("%v", key)
					strValue := fmt.Sprintf("%v", value)
					annotationsStr[strKey] = strValue
				}

				for key, value := range labels {
					strKey := fmt.Sprintf("%v", key)
					strValue := fmt.Sprintf("%v", value)
					labelsStr[strKey] = strValue
				}

				v1ObjectMeta = &nextgen.V1ObjectMeta{
					Name:            name,
					Namespace:       namespace,
					Finalizers:      s,
					ClusterName:     clusterName,
					ResourceVersion: resource_version,
					Labels:          labelsStr,
					Annotations:     annotationsStr,
					ManagedFields:   v1ManagedFieldsEntry,
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

			var v1GroupKindCluster []nextgen.V1GroupKind
			if crw, ok := projectData["spec"].([]interface{}); ok && len(crw) > 0 {
				specData := crw[0].(map[string]interface{})

				if crwList, ok := specData["cluster_resource_blacklist"].([]interface{}); ok {
					for _, w := range crwList {
						wData := w.(map[string]interface{})
						v1GroupKindCluster = append(v1GroupKindCluster, nextgen.V1GroupKind{
							Group: wData["group"].(string),
							Kind:  wData["kind"].(string),
						})
					}
				}
			}

			var v1GroupKindNameSpaceBlacklist []nextgen.V1GroupKind
			if crw, ok := projectData["spec"].([]interface{}); ok && len(crw) > 0 {
				specData := crw[0].(map[string]interface{})

				if crwList, ok := specData["namespace_resource_blacklist"].([]interface{}); ok {
					for _, w := range crwList {
						wData := w.(map[string]interface{})
						v1GroupKindNameSpaceBlacklist = append(v1GroupKindNameSpaceBlacklist, nextgen.V1GroupKind{
							Group: wData["group"].(string),
							Kind:  wData["kind"].(string),
						})
					}
				}
			}

			var v1GroupKindNameSpaceWhitelist []nextgen.V1GroupKind
			if crw, ok := projectData["spec"].([]interface{}); ok && len(crw) > 0 {
				specData := crw[0].(map[string]interface{})

				if crwList, ok := specData["namespace_resource_whitelist"].([]interface{}); ok {
					for _, w := range crwList {
						wData := w.(map[string]interface{})
						v1GroupKindNameSpaceWhitelist = append(v1GroupKindNameSpaceWhitelist, nextgen.V1GroupKind{
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
							Name:      dData["name"].(string),
						})
					}
				}
			}

			if sr, ok := projectData["spec"].([]interface{}); ok && len(sr) > 0 {
				specData := sr[0].(map[string]interface{})

				if repos, ok := specData["source_repos"].([]interface{}); ok {
					for _, repo := range repos {
						sourceRepos = append(sourceRepos, repo.(string))
					}
				}
			}

			var appprojectsProjectRole []nextgen.AppprojectsProjectRole
			if sr, ok := projectData["spec"].([]interface{}); ok && len(sr) > 0 {
				rolesData := sr[0].(map[string]interface{})

				if roles, ok := rolesData["roles"].([]interface{}); ok {
					for _, r := range roles {
						role := r.(map[string]interface{})
						var appprojectsJwtToken []nextgen.AppprojectsJwtToken
						if tokens, ok := role["jwt_tokens"].([]interface{}); ok && len(tokens) > 0 {
							for _, t := range tokens {
								if t != nil {
									token := t.(map[string]interface{})
									appprojectsJwtToken = append(appprojectsJwtToken, nextgen.AppprojectsJwtToken{
										Iat: token["iat"].(string),
										Exp: token["exp"].(string),
										Id:  token["id"].(string),
									})
								}
							}
						}
						p := role["policies"].([]interface{})
						policies := make([]string, len(p))
						if len(p) > 0 {
							for i, v := range p {
								policies[i] = fmt.Sprint(v)
							}
						}
						g := role["groups"].([]interface{})
						groups := make([]string, len(g))
						if len(g) > 0 {
							for i, v := range g {
								groups[i] = fmt.Sprint(v)
							}
						}
						appprojectsProjectRole = append(appprojectsProjectRole, nextgen.AppprojectsProjectRole{
							Name:        role["name"].(string),
							Description: role["description"].(string),
							Policies:    policies,
							Groups:      groups,
							JwtTokens:   appprojectsJwtToken,
						})
					}
				}
			}

			var orphanedResources *nextgen.AppprojectsOrphanedResourcesMonitorSettings
			if sr, ok := projectData["spec"].([]interface{}); ok && len(sr) > 0 {
				resourceData := sr[0].(map[string]interface{})
				var appprojectsOrphanedResourceKey []nextgen.AppprojectsOrphanedResourceKey
				if i, ok := resourceData["ignore"].([]interface{}); ok {
					for _, r := range i {
						s := r.(map[string]interface{})
						appprojectsOrphanedResourceKey = append(appprojectsOrphanedResourceKey, nextgen.AppprojectsOrphanedResourceKey{
							Group: s["group"].(string),
							Kind:  s["kind"].(string),
							Name:  s["name"].(string),
						})
					}
				}
				var warn bool
				if attr, ok := d.GetOk("warn"); ok {
					warn = attr.(bool)
				}
				orphanedResources = &nextgen.AppprojectsOrphanedResourcesMonitorSettings{
					Warn:   warn,
					Ignore: appprojectsOrphanedResourceKey,
				}

			}

			var appprojectsSyncWindow []nextgen.AppprojectsSyncWindow
			if sr, ok := projectData["spec"].([]interface{}); ok && len(sr) > 0 {
				syncData := sr[0].(map[string]interface{})

				if sync, ok := syncData["sync_windows"].([]interface{}); ok && len(sync) > 0 {
					for _, r := range sync {
						s := r.(map[string]interface{})
						var applications []string
						if a, ok := s["applications"].([]interface{}); ok && len(a) > 0 {
							applications = make([]string, len(a))
							if len(a) > 0 {
								for i, v := range a {
									applications[i] = fmt.Sprint(v)
								}
							}
						}

						var namespaces []string
						if name, ok := s["namespaces"].([]interface{}); ok && len(name) > 0 {
							namespaces = make([]string, len(name))
							if len(name) > 0 {
								for i, v := range name {
									namespaces[i] = fmt.Sprint(v)
								}
							}
						}

						var clusters []string
						if c, ok := s["clusters"].([]interface{}); ok && len(c) > 0 {
							clusters = make([]string, len(c))
							if len(c) > 0 {
								for i, v := range c {
									clusters[i] = fmt.Sprint(v)
								}
							}
						}

						appprojectsSyncWindow = append(appprojectsSyncWindow, nextgen.AppprojectsSyncWindow{
							Kind:         s["kind"].(string),
							Schedule:     s["schedule"].(string),
							Duration:     s["duration"].(string),
							Applications: applications,
							Namespaces:   namespaces,
							Clusters:     clusters,
							ManualSync:   s["manual_sync"].(bool),
							TimeZone:     s["time_zone"].(string),
						})
					}
				}
			}

			var appprojectsAppProjectStatus *nextgen.AppprojectsAppProjectStatus
			if rawStatus, ok := d.GetOk("status"); ok {
				if statusData, ok := rawStatus.([]interface{}); ok && len(statusData) > 0 {
					statusMap := statusData[0].(map[string]interface{})
					if jwtTokensByRole, ok := statusMap["jwtTokensByRole"]; ok {
						if jwtTokens, ok := jwtTokensByRole.([]interface{}); ok && len(statusData) > 0 {
							jwt := jwtTokens[0].(map[string]interface{})
							appprojectsAppProjectStatus.JwtTokensByRole = make(map[string]nextgen.AppprojectsJwtTokens)
							for key, value := range jwt {
								appprojectsJwt := nextgen.AppprojectsJwtTokens{}
								if items, ok := value.(map[string]interface{})["items"].([]interface{}); ok {
									for _, item := range items {
										itemMap := item.(map[string]interface{})
										jwtToken := nextgen.AppprojectsJwtToken{
											Iat: itemMap["iat"].(string),
											Exp: itemMap["exp"].(string),
											Id:  itemMap["id"].(string),
										}
										appprojectsJwt.Items = append(appprojectsJwt.Items, jwtToken)
									}
								}
								appprojectsAppProjectStatus.JwtTokensByRole[key] = appprojectsJwt
							}
						}
					}
				}
			}

			var appprojectsSignatureKey []nextgen.AppprojectsSignatureKey
			if crw, ok := projectData["spec"].([]interface{}); ok && len(crw) > 0 {
				specData := crw[0].(map[string]interface{})

				if crwList, ok := specData["signature_keys"].([]interface{}); ok {
					for _, w := range crwList {
						wData := w.(map[string]interface{})
						appprojectsSignatureKey = append(appprojectsSignatureKey, nextgen.AppprojectsSignatureKey{
							KeyID: wData["key_id"].(string),
						})
					}
				}
			}

			approjectsAppProjectSpec = &nextgen.AppprojectsAppProjectSpec{
				Destinations:               appprojectsApplicationDestination,
				ClusterResourceWhitelist:   v1GroupKind,
				ClusterResourceBlacklist:   v1GroupKindCluster,
				NamespaceResourceBlacklist: v1GroupKindNameSpaceBlacklist,
				NamespaceResourceWhitelist: v1GroupKindNameSpaceWhitelist,
				SourceRepos:                sourceRepos,
				Roles:                      appprojectsProjectRole,
				SyncWindows:                appprojectsSyncWindow,
				SignatureKeys:              appprojectsSignatureKey,
				OrphanedResources:          orphanedResources,
			}

			appprojectsAppProject = &nextgen.AppprojectsAppProject{
				Metadata: v1ObjectMeta,
				Spec:     approjectsAppProjectSpec,
				Status:   appprojectsAppProjectStatus,
			}

			projectsProjectUpdateRequest = nextgen.ProjectsProjectUpdateRequest{
				Project: appprojectsAppProject,
			}

		}
	}

	return projectsProjectUpdateRequest
}

// Function to map Terraform schema to internal data structure
func createRequestBody(d *schema.ResourceData) nextgen.ProjectsProjectCreateRequest {
	var projectsProjectCreateRequest nextgen.ProjectsProjectCreateRequest
	var appprojectsAppProject *nextgen.AppprojectsAppProject
	var approjectsAppProjectSpec *nextgen.AppprojectsAppProjectSpec
	var v1ObjectMeta *nextgen.V1ObjectMeta
	var sourceRepos []string

	if v, ok := d.GetOk("project"); ok {
		for _, item := range v.([]interface{}) {
			projectData := item.(map[string]interface{})

			if md, ok := projectData["metadata"].([]interface{}); ok {
				mdData := md[0].(map[string]interface{})
				var v1ManagedFieldsEntry []nextgen.V1ManagedFieldsEntry
				if managerFieldsList, ok := mdData["managed_fields"].([]interface{}); ok {
					for _, w := range managerFieldsList {
						wData := w.(map[string]interface{})
						v1ManagedFieldsEntry = append(v1ManagedFieldsEntry, nextgen.V1ManagedFieldsEntry{
							Manager:     wData["api_version"].(string),
							Operation:   wData["operation"].(string),
							ApiVersion:  wData["api_version"].(string),
							FieldsType:  wData["fields_type"].(string),
							Subresource: wData["subresource"].(string),
						})
					}
				}

				var namespace, name, clusterName string
				var finalizers []interface{}
				var labels, annotations map[string]interface{}
				if nil != mdData["namespace"] {
					namespace = mdData["namespace"].(string)
				}
				if nil != mdData["name"] {
					name = mdData["name"].(string)
				}

				if nil != mdData["cluster_name"] {
					clusterName = mdData["cluster_name"].(string)
				}

				if f, ok := mdData["finalizers"].([]interface{}); ok {
					for _, r := range f {
						finalizers = append(finalizers, r.(string))
					}
				}

				if nil != mdData["labels"] {
					labels = mdData["labels"].(map[string]interface{})
				}
				if nil != mdData["annotations"] {
					annotations = mdData["annotations"].(map[string]interface{})
				}

				s := make([]string, len(finalizers))
				for i, v := range finalizers {
					s[i] = fmt.Sprint(v)
				}

				var annotationsStr, labelsStr map[string]string
				annotationsStr = make(map[string]string)
				labelsStr = make(map[string]string)

				for key, value := range annotations {
					strKey := fmt.Sprintf("%v", key)
					strValue := fmt.Sprintf("%v", value)
					annotationsStr[strKey] = strValue
				}

				for key, value := range labels {
					strKey := fmt.Sprintf("%v", key)
					strValue := fmt.Sprintf("%v", value)
					labelsStr[strKey] = strValue
				}

				v1ObjectMeta = &nextgen.V1ObjectMeta{
					Name:          name,
					Namespace:     namespace,
					Finalizers:    s,
					ClusterName:   clusterName,
					Labels:        labelsStr,
					Annotations:   annotationsStr,
					ManagedFields: v1ManagedFieldsEntry,
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

			var v1GroupKindCluster []nextgen.V1GroupKind
			if crw, ok := projectData["spec"].([]interface{}); ok && len(crw) > 0 {
				specData := crw[0].(map[string]interface{})

				if crwList, ok := specData["cluster_resource_blacklist"].([]interface{}); ok {
					for _, w := range crwList {
						wData := w.(map[string]interface{})
						v1GroupKindCluster = append(v1GroupKindCluster, nextgen.V1GroupKind{
							Group: wData["group"].(string),
							Kind:  wData["kind"].(string),
						})
					}
				}
			}

			var v1GroupKindNameSpaceBlacklist []nextgen.V1GroupKind
			if crw, ok := projectData["spec"].([]interface{}); ok && len(crw) > 0 {
				specData := crw[0].(map[string]interface{})

				if crwList, ok := specData["namespace_resource_blacklist"].([]interface{}); ok {
					for _, w := range crwList {
						wData := w.(map[string]interface{})
						v1GroupKindNameSpaceBlacklist = append(v1GroupKindNameSpaceBlacklist, nextgen.V1GroupKind{
							Group: wData["group"].(string),
							Kind:  wData["kind"].(string),
						})
					}
				}
			}

			var v1GroupKindNameSpaceWhitelist []nextgen.V1GroupKind
			if crw, ok := projectData["spec"].([]interface{}); ok && len(crw) > 0 {
				specData := crw[0].(map[string]interface{})

				if crwList, ok := specData["namespace_resource_whitelist"].([]interface{}); ok {
					for _, w := range crwList {
						wData := w.(map[string]interface{})
						v1GroupKindNameSpaceWhitelist = append(v1GroupKindNameSpaceWhitelist, nextgen.V1GroupKind{
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
							Name:      dData["name"].(string),
						})
					}
				}
			}

			if sr, ok := projectData["spec"].([]interface{}); ok && len(sr) > 0 {
				specData := sr[0].(map[string]interface{})

				if repos, ok := specData["source_repos"].([]interface{}); ok {
					for _, repo := range repos {
						sourceRepos = append(sourceRepos, repo.(string))
					}
				}
			}

			var appprojectsProjectRole []nextgen.AppprojectsProjectRole
			if sr, ok := projectData["spec"].([]interface{}); ok && len(sr) > 0 {
				rolesData := sr[0].(map[string]interface{})

				if roles, ok := rolesData["roles"].([]interface{}); ok {
					for _, r := range roles {
						role := r.(map[string]interface{})
						var appprojectsJwtToken []nextgen.AppprojectsJwtToken
						if tokens, ok := role["jwt_tokens"].([]interface{}); ok && len(tokens) > 0 {
							for _, t := range tokens {
								if t != nil {
									token := t.(map[string]interface{})
									appprojectsJwtToken = append(appprojectsJwtToken, nextgen.AppprojectsJwtToken{
										Iat: token["iat"].(string),
										Exp: token["exp"].(string),
										Id:  token["id"].(string),
									})
								}
							}
						}
						p := role["policies"].([]interface{})
						policies := make([]string, len(p))
						if len(p) > 0 {
							for i, v := range p {
								policies[i] = fmt.Sprint(v)
							}
						}
						g := role["groups"].([]interface{})
						groups := make([]string, len(g))
						if len(g) > 0 {
							for i, v := range g {
								groups[i] = fmt.Sprint(v)
							}
						}
						appprojectsProjectRole = append(appprojectsProjectRole, nextgen.AppprojectsProjectRole{
							Name:        role["name"].(string),
							Description: role["description"].(string),
							Policies:    policies,
							Groups:      groups,
							JwtTokens:   appprojectsJwtToken,
						})
					}
				}
			}

			var orphanedResources *nextgen.AppprojectsOrphanedResourcesMonitorSettings
			if sr, ok := projectData["spec"].([]interface{}); ok && len(sr) > 0 {
				resourceData := sr[0].(map[string]interface{})
				var appprojectsOrphanedResourceKey []nextgen.AppprojectsOrphanedResourceKey
				if i, ok := resourceData["ignore"].([]interface{}); ok {
					for _, r := range i {
						s := r.(map[string]interface{})
						appprojectsOrphanedResourceKey = append(appprojectsOrphanedResourceKey, nextgen.AppprojectsOrphanedResourceKey{
							Group: s["group"].(string),
							Kind:  s["kind"].(string),
							Name:  s["name"].(string),
						})
					}
				}
				var warn bool
				if attr, ok := d.GetOk("warn"); ok {
					warn = attr.(bool)
				}
				orphanedResources = &nextgen.AppprojectsOrphanedResourcesMonitorSettings{
					Warn:   warn,
					Ignore: appprojectsOrphanedResourceKey,
				}

			}

			var appprojectsSyncWindow []nextgen.AppprojectsSyncWindow
			if sr, ok := projectData["spec"].([]interface{}); ok && len(sr) > 0 {
				syncData := sr[0].(map[string]interface{})

				if sync, ok := syncData["sync_windows"].([]interface{}); ok && len(sync) > 0 {
					for _, r := range sync {
						s := r.(map[string]interface{})
						var applications []string
						if a, ok := s["applications"].([]interface{}); ok && len(a) > 0 {
							applications = make([]string, len(a))
							if len(a) > 0 {
								for i, v := range a {
									applications[i] = fmt.Sprint(v)
								}
							}
						}

						var namespaces []string
						if name, ok := s["namespaces"].([]interface{}); ok && len(name) > 0 {
							namespaces = make([]string, len(name))
							if len(name) > 0 {
								for i, v := range name {
									namespaces[i] = fmt.Sprint(v)
								}
							}
						}

						var clusters []string
						if c, ok := s["clusters"].([]interface{}); ok && len(c) > 0 {
							clusters = make([]string, len(c))
							if len(c) > 0 {
								for i, v := range c {
									clusters[i] = fmt.Sprint(v)
								}
							}
						}

						appprojectsSyncWindow = append(appprojectsSyncWindow, nextgen.AppprojectsSyncWindow{
							Kind:         s["kind"].(string),
							Schedule:     s["schedule"].(string),
							Duration:     s["duration"].(string),
							Applications: applications,
							Namespaces:   namespaces,
							Clusters:     clusters,
							ManualSync:   s["manual_sync"].(bool),
							TimeZone:     s["time_zone"].(string),
						})
					}
				}
			}

			var appprojectsSignatureKey []nextgen.AppprojectsSignatureKey
			if crw, ok := projectData["spec"].([]interface{}); ok && len(crw) > 0 {
				specData := crw[0].(map[string]interface{})

				if crwList, ok := specData["signature_keys"].([]interface{}); ok {
					for _, w := range crwList {
						wData := w.(map[string]interface{})
						appprojectsSignatureKey = append(appprojectsSignatureKey, nextgen.AppprojectsSignatureKey{
							KeyID: wData["key_id"].(string),
						})
					}
				}
			}

			approjectsAppProjectSpec = &nextgen.AppprojectsAppProjectSpec{
				Destinations:               appprojectsApplicationDestination,
				ClusterResourceWhitelist:   v1GroupKind,
				ClusterResourceBlacklist:   v1GroupKindCluster,
				NamespaceResourceBlacklist: v1GroupKindNameSpaceBlacklist,
				NamespaceResourceWhitelist: v1GroupKindNameSpaceWhitelist,
				SourceRepos:                sourceRepos,
				Roles:                      appprojectsProjectRole,
				SyncWindows:                appprojectsSyncWindow,
				SignatureKeys:              appprojectsSignatureKey,
				OrphanedResources:          orphanedResources,
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
	d.Set("account_id", "1bvyLackQK-Hapk25-Ry4w")
	projectList := []interface{}{}
	project := map[string]interface{}{}
	if projects.Metadata != nil {
		d.Set("query_name", projects.Metadata.Name)
		metadataList := []interface{}{}
		metadata := map[string]interface{}{}
		var finalizers = projects.Metadata.Finalizers
		metadata["finalizers"] = finalizers
		metadata["resource_version"] = projects.Metadata.ResourceVersion
		metadata["name"] = projects.Metadata.Name
		metadata["namespace"] = projects.Metadata.Namespace
		annotationsStr := make(map[string]string)
		for key, value := range projects.Metadata.Annotations {
			strKey := fmt.Sprintf("%v", key)
			strValue := fmt.Sprintf("%v", value)
			annotationsStr[strKey] = strValue
		}

		labelsStr := make(map[string]string)
		for key, value := range projects.Metadata.Labels {
			strKey := fmt.Sprintf("%v", key)
			strValue := fmt.Sprintf("%v", value)
			labelsStr[strKey] = strValue
		}
		metadata["annotations"] = annotationsStr
		metadata["labels"] = labelsStr
		manageFieldList := []interface{}{}
		manageField := map[string]interface{}{}
		for _, k := range projects.Metadata.ManagedFields {
			manageField["manager"] = k.Manager
			manageField["operation"] = k.Operation
			manageFieldList = append(manageFieldList, manageField)
		}
		metadata["managed_fields"] = manageFieldList
		metadataList = append(metadataList, metadata)
		project["metadata"] = metadataList
	}

	if projects.Spec != nil {
		specdataList := []interface{}{}
		spec := map[string]interface{}{}
		var sourceRepoList = projects.Spec.SourceRepos
		spec["source_repos"] = sourceRepoList
		clusterResourceWhitelist := []interface{}{}
		clusterResourceWhite := map[string]interface{}{}
		clusterResourceWhite["group"] = projects.Spec.ClusterResourceWhitelist[0].Group
		clusterResourceWhite["kind"] = projects.Spec.ClusterResourceWhitelist[0].Kind
		clusterResourceWhitelist = append(clusterResourceWhitelist, clusterResourceWhite)
		spec["cluster_resource_whitelist"] = clusterResourceWhitelist
		destinationList := []interface{}{}
		destination := map[string]interface{}{}
		destination["namespace"] = projects.Spec.Destinations[0].Namespace
		destination["server"] = projects.Spec.Destinations[0].Server
		destination["name"] = projects.Spec.Destinations[0].Name
		destinationList = append(destinationList, destination)
		spec["destinations"] = destinationList
		orphanList := []interface{}{}
		orphan := map[string]interface{}{}
		orphan["warn"] = projects.Spec.OrphanedResources.Warn
		orphanList = append(orphanList, orphan)
		spec["orphaned_resources"] = orphanList
		// Convert 'namespace_resource_blacklist' field
		if len(projects.Spec.NamespaceResourceBlacklist) > 0 {
			nsResourceBlacklist := make([]interface{}, len(projects.Spec.NamespaceResourceBlacklist))
			for i, item := range projects.Spec.NamespaceResourceBlacklist {
				nsResourceBlacklist[i] = map[string]interface{}{
					"group": item.Group,
					"kind":  item.Kind,
				}
			}
			spec["namespace_resource_blacklist"] = nsResourceBlacklist
		}
		if len(projects.Spec.NamespaceResourceWhitelist) > 0 {
			nsResourceWhitelist := make([]interface{}, len(projects.Spec.NamespaceResourceWhitelist))
			for i, item := range projects.Spec.NamespaceResourceWhitelist {
				nsResourceWhitelist[i] = map[string]interface{}{
					"group": item.Group,
					"kind":  item.Kind,
				}
			}
			spec["namespace_resource_whitelist"] = nsResourceWhitelist
		}
		if len(projects.Spec.Roles) > 0 {
			rolesList := make([]interface{}, len(projects.Spec.Roles))
			for i, item := range projects.Spec.Roles {
				jwtTokensList := make([]interface{}, len(item.JwtTokens))
				for j, token := range item.JwtTokens {
					jwtTokensList[j] = map[string]interface{}{
						"iat": token.Iat,
					}
				}
				rolesList[i] = map[string]interface{}{
					"description": item.Description,
					"name":        item.Name,
					"policies":    item.Policies,
					"groups":      item.Groups,
					"jwt_tokens":  jwtTokensList,
				}
			}
			spec["roles"] = rolesList
		}
		if len(projects.Spec.SyncWindows) > 0 {
			syncWindowsList := make([]interface{}, len(projects.Spec.SyncWindows))
			for i, item := range projects.Spec.SyncWindows {
				syncWindowsList[i] = map[string]interface{}{
					"applications": item.Applications,
					"duration":     item.Duration,
					"kind":         item.Kind,
					"manual_sync":  item.ManualSync,
					"schedule":     item.Schedule,
					"namespaces":   item.Namespaces,
					"clusters":     item.Clusters,
				}
			}
			spec["sync_windows"] = syncWindowsList
		}
		specdataList = append(specdataList, spec)
		project["spec"] = specdataList
	}
	projectList = append(projectList, project)
	d.Set("project", projectList)
}
