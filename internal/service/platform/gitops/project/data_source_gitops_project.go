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
				Computed:    true,
			},
			"upsert": {
				Description: "Indicates if the GitOps repository should be updated if existing and inserted if not.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"project": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "GitOps project configuration.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"metadata": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "Metadata details for the GitOps project.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Name of the GitOps project.",
									},
									"generate_name": {
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
										Description: "Generated name of the GitOps project.",
									},
									"namespace": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Namespace of the GitOps project.",
									},
									"self_link": {
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
										Description: "Self link of the GitOps project.",
									},
									"uid": {
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
										Description: "UID of the GitOps project.",
									},
									"resource_version": {
										Type:        schema.TypeString,
										Computed:    true,
										Optional:    true,
										Description: "Resource version of the GitOps project.",
									},
									"generation": {
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
										Description: "Generation of the GitOps project.",
									},
									"deletion_grace_period_seconds": {
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
										Description: "Deletion grace period in seconds of the GitOps project.",
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
									"owner_references": {
										Type:        schema.TypeList,
										Optional:    true,
										Computed:    true,
										Description: "Owner references associated with the GitOps project.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"api_version": {
													Type:        schema.TypeString,
													Optional:    true,
													Computed:    true,
													Description: "API version of the owner reference.",
												},
												"kind": {
													Type:        schema.TypeString,
													Optional:    true,
													Computed:    true,
													Description: "Kind of the owner reference.",
												},
												"name": {
													Type:        schema.TypeString,
													Optional:    true,
													Computed:    true,
													Description: "Name of the owner reference.",
												},
												"uid": {
													Type:        schema.TypeString,
													Optional:    true,
													Computed:    true,
													Description: "UID of the owner reference.",
												},
												"controller": {
													Type:        schema.TypeBool,
													Optional:    true,
													Computed:    true,
													Description: "Specifies whether the owner reference is a controller.",
												},
												"block_owner_deletion": {
													Type:        schema.TypeBool,
													Optional:    true,
													Computed:    true,
													Description: "Specifies whether to block owner deletion.",
												},
											},
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
							Optional:    true,
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
						"status": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "Status details for the GitOps project.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"jwt_tokens_by_role": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "JWT tokens by role status for the GitOps project.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"items": {
													Type:        schema.TypeList,
													Optional:    true,
													Description: "List of JWT tokens by role.",
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
