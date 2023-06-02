package template

import (
	"context"
	"net/http"

	"github.com/antihax/optional"
	"github.com/harness/harness-openapi-go-client/nextgen"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func ResourceTemplate() *schema.Resource {
	resource := &schema.Resource{
		Description: "Resource for creating a Template. Description field is deprecated",

		ReadContext:   resourceTemplateRead,
		UpdateContext: resourceTemplateCreateOrUpdate,
		DeleteContext: resourceTemplateDelete,
		CreateContext: resourceTemplateCreateOrUpdate,
		Importer:      helpers.MultiLevelResourceImporter,

		Schema: map[string]*schema.Schema{
			"template_yaml": {
				Description: "Yaml for creating new Template." + helpers.Descriptions.YamlText.String(),
				Type:        schema.TypeString,
				Required:    true,
			},
			"version": {
				Description: "Version Label for Template.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"is_stable": {
				Description: "True if given version for template to be set as stable.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"comments": {
				Description: "Specify comment with respect to changes.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"git_details": {
				Description: "Contains parameters related to creating an Entity for Git Experience.",
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"branch_name": {
							Description: "Name of the branch.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"file_path": {
							Description: "File path of the Entity in the repository.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"commit_message": {
							Description: "Commit message used for the merge commit.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"base_branch": {
							Description: "Name of the default branch (this checks out a new branch titled by branch_name).",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"connector_ref": {
							Description: "Identifier of the Harness Connector used for CRUD operations on the Entity." + helpers.Descriptions.ConnectorRefText.String(),
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"store_type": {
							Description:  "Specifies whether the Entity is to be stored in Git or not. Possible values: INLINE, REMOTE.",
							Type:         schema.TypeString,
							Optional:     true,
							Computed:     true,
							ValidateFunc: validation.StringInSlice([]string{"INLINE", "REMOTE"}, false),
						},
						"repo_name": {
							Description: "Name of the repository.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"last_object_id": {
							Description: "Last object identifier (for Github). To be provided only when updating Pipeline.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"last_commit_id": {
							Description: "Last commit identifier (for Git Repositories other than Github). To be provided only when updating Pipeline.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"force_delete": {
				Description: "Enable this flag for force deletion of template",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"identifier": {
				Description: "Unique identifier of the resource",
				Type:        schema.TypeString,
				Required:    true,
			},
			"name": {
				Description: "Name of the Variable",
				Type:        schema.TypeString,
				Required:    true,
			},
			"description": {
				Description: "Description of the entity. Description field is deprecated",
				Type:        schema.TypeString,
				Optional:    true,
				Deprecated:  "description field is deprecated",
			},
			"org_id": {
				Description: "Organization Identifier for the Entity",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"project_id": {
				Description: "Project Identifier for the Entity",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"tags": {
				Description: "Tags to associate with the resource.",
				Type:        schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional: true,
			},
		},
	}

	return resource
}

func resourceTemplateRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetClientWithContext(ctx)

	template_id := d.Id()
	org_id := d.Get("org_id").(string)
	project_id := d.Get("project_id").(string)
	var comments = helpers.BuildField(d, "comments")
	var store_type = helpers.BuildField(d, "git_details.0.store_type")
	var base_branch = helpers.BuildField(d, "git_details.0.base_branch")
	var commit_message = helpers.BuildField(d, "git_details.0.commit_message")
	var connector_ref = helpers.BuildField(d, "git_details.0.connector_ref")
	var branch_name optional.String
	branch_name = helpers.BuildField(d, "git_details.0.branch_name")
	version := d.Get("version").(string)

	var err error
	var resp nextgen.TemplateWithInputsResponse
	var httpResp *http.Response

	if project_id != "" {
		if version == "" {
			resp, httpResp, err = c.ProjectTemplateApi.GetTemplateStableProject(ctx, org_id, project_id, template_id, &nextgen.ProjectTemplateApiGetTemplateStableProjectOpts{
				HarnessAccount: optional.NewString(c.AccountId),
				BranchName:     branch_name})
		} else {
			resp, httpResp, err = c.ProjectTemplateApi.GetTemplateProject(ctx, project_id, template_id, org_id, version, &nextgen.ProjectTemplateApiGetTemplateProjectOpts{
				HarnessAccount: optional.NewString(c.AccountId),
				BranchName:     branch_name})
		}
	} else if org_id != "" && project_id == "" {
		if version == "" {
			resp, httpResp, err = c.OrgTemplateApi.GetTemplateStableOrg(ctx, org_id, template_id, &nextgen.OrgTemplateApiGetTemplateStableOrgOpts{
				HarnessAccount: optional.NewString(c.AccountId),
				BranchName:     branch_name,
			})
		} else {
			resp, httpResp, err = c.OrgTemplateApi.GetTemplateOrg(ctx, template_id, org_id, version, &nextgen.OrgTemplateApiGetTemplateOrgOpts{
				HarnessAccount: optional.NewString(c.AccountId),
				BranchName:     branch_name,
			})
		}
	} else {
		if version == "" {
			resp, httpResp, err = c.AccountTemplateApi.GetTemplateStableAcc(ctx, template_id, &nextgen.AccountTemplateApiGetTemplateStableAccOpts{
				HarnessAccount: optional.NewString(c.AccountId),
				BranchName:     branch_name,
			})
		} else {
			resp, httpResp, err = c.AccountTemplateApi.GetTemplateAcc(ctx, template_id, version, &nextgen.AccountTemplateApiGetTemplateAccOpts{
				HarnessAccount: optional.NewString(c.AccountId),
				BranchName:     branch_name,
			})
		}
	}

	if httpResp.StatusCode == 404 {
		d.SetId("")
		d.MarkNewResource()
		return nil
	}

	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	readTemplate(d, resp, comments.Value(), store_type, base_branch, commit_message, connector_ref)

	return nil
}

func resourceTemplateCreateOrUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetClientWithContext(ctx)

	var err error
	var template_id string
	var branch_name string
	var store_type optional.String
	var base_branch optional.String
	var commit_message optional.String
	var connector_ref optional.String
	var resp nextgen.TemplateResponse
	var httpResp *http.Response
	id := d.Id()
	org_id := d.Get("org_id").(string)
	project_id := d.Get("project_id").(string)
	comments := d.Get("comments").(string)
	version := d.Get("version").(string)

	if id == "" {
		template := buildCreateTemplate(d)
		if template.GitDetails != nil {
			base_branch = optional.NewString(template.GitDetails.BaseBranch)
			store_type = optional.NewString(template.GitDetails.StoreType)
			commit_message = optional.NewString(template.GitDetails.CommitMessage)
			connector_ref = optional.NewString(template.GitDetails.ConnectorRef)
			branch_name = template.GitDetails.BranchName
		}
		if project_id != "" {
			resp, httpResp, err = c.ProjectTemplateApi.CreateTemplatesProject(ctx, org_id, project_id, &nextgen.ProjectTemplateApiCreateTemplatesProjectOpts{
				Body:           optional.NewInterface(template),
				HarnessAccount: optional.NewString(c.AccountId),
			})
			if resp.Identifier != "" {
				template_id = resp.Identifier
			} else {
				template_id = resp.Slug
			}
		} else if org_id != "" && project_id == "" {
			resp, httpResp, err = c.OrgTemplateApi.CreateTemplatesOrg(ctx, org_id, &nextgen.OrgTemplateApiCreateTemplatesOrgOpts{
				HarnessAccount: optional.NewString(c.AccountId),
				Body:           optional.NewInterface(template),
			})
			if resp.Identifier != "" {
				template_id = resp.Identifier
			} else {
				template_id = resp.Slug
			}
		} else {
			resp, httpResp, err = c.AccountTemplateApi.CreateTemplatesAcc(ctx, &nextgen.AccountTemplateApiCreateTemplatesAccOpts{
				HarnessAccount: optional.NewString(c.AccountId),
				Body:           optional.NewInterface(template),
			})
			if resp.Identifier != "" {
				template_id = resp.Identifier
			} else {
				template_id = resp.Slug
			}
		}
	} else {
		template := buildUpdateTemplate(d)
		if template.GitDetails != nil {
			base_branch = optional.NewString(template.GitDetails.BaseBranch)
			branch_name = template.GitDetails.BranchName
			store_type = optional.NewString(template.GitDetails.StoreType)
			commit_message = optional.NewString(template.GitDetails.CommitMessage)
			connector_ref = optional.NewString(template.GitDetails.ConnectorRef)
		}
		if project_id != "" {
			resp, httpResp, err = c.ProjectTemplateApi.UpdateTemplateProject(ctx, project_id, id, org_id, version, &nextgen.ProjectTemplateApiUpdateTemplateProjectOpts{
				Body:           optional.NewInterface(template),
				HarnessAccount: optional.NewString(c.AccountId),
			})
			if resp.Identifier != "" {
				template_id = resp.Identifier
			} else {
				template_id = resp.Slug
			}
		} else if org_id != "" && project_id == "" {
			resp, httpResp, err = c.OrgTemplateApi.UpdateTemplateOrg(ctx, id, org_id, version, &nextgen.OrgTemplateApiUpdateTemplateOrgOpts{
				HarnessAccount: optional.NewString(c.AccountId),
				Body:           optional.NewInterface(template),
			})
			if resp.Identifier != "" {
				template_id = resp.Identifier
			} else {
				template_id = resp.Slug
			}
		} else {
			resp, httpResp, err = c.AccountTemplateApi.UpdateTemplateAcc(ctx, id, version, &nextgen.AccountTemplateApiUpdateTemplateAccOpts{
				HarnessAccount: optional.NewString(c.AccountId),
				Body:           optional.NewInterface(template),
			})
			if resp.Identifier != "" {
				template_id = resp.Identifier
			} else {
				template_id = resp.Slug
			}
		}
	}

	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	var respGet nextgen.TemplateWithInputsResponse

	if project_id != "" {
		if version == "" {
			respGet, httpResp, err = c.ProjectTemplateApi.GetTemplateStableProject(ctx, org_id, project_id, template_id, &nextgen.ProjectTemplateApiGetTemplateStableProjectOpts{
				HarnessAccount: optional.NewString(c.AccountId),
				BranchName:     optional.NewString(branch_name)})
		} else {
			respGet, httpResp, err = c.ProjectTemplateApi.GetTemplateProject(ctx, project_id, template_id, org_id, version, &nextgen.ProjectTemplateApiGetTemplateProjectOpts{
				HarnessAccount: optional.NewString(c.AccountId),
				BranchName:     optional.NewString(branch_name)})
		}
	} else if org_id != "" && project_id == "" {
		if version == "" {
			respGet, httpResp, err = c.OrgTemplateApi.GetTemplateStableOrg(ctx, org_id, template_id, &nextgen.OrgTemplateApiGetTemplateStableOrgOpts{
				HarnessAccount: optional.NewString(c.AccountId),
				BranchName:     optional.NewString(branch_name),
			})
		} else {
			respGet, httpResp, err = c.OrgTemplateApi.GetTemplateOrg(ctx, template_id, org_id, version, &nextgen.OrgTemplateApiGetTemplateOrgOpts{
				HarnessAccount: optional.NewString(c.AccountId),
				BranchName:     optional.NewString(branch_name),
			})
		}
	} else {
		if version == "" {
			respGet, httpResp, err = c.AccountTemplateApi.GetTemplateStableAcc(ctx, template_id, &nextgen.AccountTemplateApiGetTemplateStableAccOpts{
				HarnessAccount: optional.NewString(c.AccountId),
				BranchName:     optional.NewString(branch_name),
			})
		} else {
			respGet, httpResp, err = c.AccountTemplateApi.GetTemplateAcc(ctx, template_id, version, &nextgen.AccountTemplateApiGetTemplateAccOpts{
				HarnessAccount: optional.NewString(c.AccountId),
				BranchName:     optional.NewString(branch_name),
			})
		}
	}

	readTemplate(d, respGet, comments, store_type, base_branch, commit_message, connector_ref)

	return nil
}

func resourceTemplateDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetClientWithContext(ctx)

	id := d.Get("identifier").(string)
	org_id := d.Get("org_id").(string)
	project_id := d.Get("project_id").(string)
	version := d.Get("version").(string)
	var httpResp *http.Response
	var err error

	if project_id != "" {
		httpResp, err = c.ProjectTemplateApi.DeleteTemplateProject(ctx, project_id, id, org_id, version, &nextgen.ProjectTemplateApiDeleteTemplateProjectOpts{
			HarnessAccount: optional.NewString(c.AccountId),
			Comments:       helpers.BuildField(d, "comments"),
			ForceDelete:    helpers.BuildFieldForBoolean(d, "force_delete"),
		})
	} else if org_id != "" && project_id == "" {
		httpResp, err = c.OrgTemplateApi.DeleteTemplateOrg(ctx, id, org_id, version, &nextgen.OrgTemplateApiDeleteTemplateOrgOpts{
			HarnessAccount: optional.NewString(c.AccountId),
			Comments:       helpers.BuildField(d, "comments"),
			ForceDelete:    helpers.BuildFieldForBoolean(d, "force_delete"),
		})
	} else {
		httpResp, err = c.AccountTemplateApi.DeleteTemplateAcc(ctx, id, version, &nextgen.AccountTemplateApiDeleteTemplateAccOpts{
			HarnessAccount: optional.NewString(c.AccountId),
			Comments:       helpers.BuildField(d, "comments"),
			ForceDelete:    helpers.BuildFieldForBoolean(d, "force_delete"),
		})

	}
	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	return nil
}

func buildUpdateTemplate(d *schema.ResourceData) nextgen.TemplateUpdateRequestBody {
	template := nextgen.TemplateUpdateRequestBody{
		TemplateYaml: d.Get("template_yaml").(string),
		Comments:     d.Get("comments").(string),
	}

	if attr, ok := d.GetOk("git_details"); ok {
		config := attr.([]interface{})[0].(map[string]interface{})

		if attr, ok := config["store_type"]; ok {
			if attr != "" {
				template.GitDetails = &nextgen.GitUpdateDetails1{}

				if attr, ok := config["branch_name"]; ok {
					template.GitDetails.BranchName = attr.(string)
				}

				if attr, ok := config["file_path"]; ok {
					template.GitDetails.FilePath = attr.(string)
				}

				if attr, ok := config["last_object_id"]; ok {
					template.GitDetails.LastObjectId = attr.(string)
				}

				if attr, ok := config["commit_message"]; ok {
					template.GitDetails.CommitMessage = attr.(string)
				}

				if attr, ok := config["base_branch"]; ok {
					template.GitDetails.BaseBranch = attr.(string)
				}

				if attr, ok := config["connector_ref"]; ok {
					template.GitDetails.ConnectorRef = attr.(string)
				}

				if attr, ok := config["store_type"]; ok {
					template.GitDetails.StoreType = attr.(string)
				}

				if attr, ok := config["repo_name"]; ok {
					template.GitDetails.RepoName = attr.(string)
				}
			}
		}
	}

	return template
}

func buildCreateTemplate(d *schema.ResourceData) nextgen.TemplateCreateRequestBody {
	template := nextgen.TemplateCreateRequestBody{
		TemplateYaml: d.Get("template_yaml").(string),
		IsStable:     d.Get("is_stable").(bool),
		Comments:     d.Get("comments").(string),
	}

	if attr, ok := d.GetOk("git_details"); ok {
		config := attr.([]interface{})[0].(map[string]interface{})
		template.GitDetails = &nextgen.GitCreateDetails1{}

		if attr, ok := config["branch_name"]; ok {
			template.GitDetails.BranchName = attr.(string)
		}

		if attr, ok := config["file_path"]; ok {
			template.GitDetails.FilePath = attr.(string)
		}

		if attr, ok := config["commit_message"]; ok {
			template.GitDetails.CommitMessage = attr.(string)
		}

		if attr, ok := config["base_branch"]; ok {
			template.GitDetails.BaseBranch = attr.(string)
		}

		if attr, ok := config["connector_ref"]; ok {
			template.GitDetails.ConnectorRef = attr.(string)
		}

		if attr, ok := config["store_type"]; ok {
			template.GitDetails.StoreType = attr.(string)
		}

		if attr, ok := config["repo_name"]; ok {
			template.GitDetails.RepoName = attr.(string)
		}

	}

	return template
}

func readTemplate(d *schema.ResourceData, template nextgen.TemplateWithInputsResponse, comments string, store_type optional.String, base_branch optional.String, commit_message optional.String, connector_ref optional.String) {
	if template.Template.Identifier != "" {
		d.SetId(template.Template.Identifier)
	} else {
		d.SetId(template.Template.Slug)
	}
	if template.Template.Identifier != "" {
		d.Set("identifier", template.Template.Identifier)
	} else {
		d.Set("identifier", template.Template.Slug)
	}
	d.Set("name", template.Template.Name)
	d.Set("org_id", template.Template.Org)
	d.Set("project_id", template.Template.Project)
	d.Set("template_yaml", template.Template.Yaml)
	d.Set("is_stable", template.Template.StableTemplate)
	d.Set("version", template.Template.VersionLabel)
	d.Set("comments", comments)
	if template.Template.GitDetails != nil {
		d.Set("git_details", []interface{}{readGitDetails(template, store_type, base_branch, commit_message, connector_ref)})
	}
}

func readGitDetails(template nextgen.TemplateWithInputsResponse, store_type optional.String, base_branch optional.String, commit_message optional.String, connector_ref optional.String) map[string]interface{} {
	git_details := map[string]interface{}{
		"branch_name":    template.Template.GitDetails.BranchName,
		"file_path":      template.Template.GitDetails.FilePath,
		"repo_name":      template.Template.GitDetails.RepoName,
		"last_commit_id": template.Template.GitDetails.CommitId,
		"last_object_id": template.Template.GitDetails.ObjectId,
	}
	if store_type.IsSet() {
		git_details["store_type"] = store_type.Value()
	}
	if base_branch.IsSet() {
		git_details["base_branch"] = base_branch.Value()
	}
	if commit_message.IsSet() {
		git_details["commit_message"] = commit_message.Value()
	}
	if connector_ref.IsSet() {
		git_details["connector_ref"] = connector_ref.Value()
	}
	return git_details
}
