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
)

func DataSourceTemplate() *schema.Resource {
	resource := &schema.Resource{
		Description: "Data source for retrieving a Harness pipeline.",

		ReadContext: dataTemplateRead,

		Schema: map[string]*schema.Schema{
			"template_yaml": {
				Description: "Yaml for creating new Template.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"version": {
				Description: "Version Label for Template.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"branch_name": {
				Description: "Version Label for Template.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"is_stable": {
				Description: "True if given version for template to be set as stable.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"scope": {
				Description: "Scope of template.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"child_type": {
				Description: "Defines child template type.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"connector_ref": {
				Description: "Identifier of the Harness Connector used for CRUD operations on the Entity.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"store_type": {
				Description: "Specifies whether the Entity is to be stored in Git or not. Possible values: INLINE, REMOTE.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"git_details": {
				Description: "Contains parameters related to creating an Entity for Git Experience.",
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"branch_name": {
							Description: "Name of the branch.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"file_path": {
							Description: "File path of the Entity in the repository.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"repo_url": {
							Description: "Repo url of the Entity in the repository.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"file_url": {
							Description: "File url of the Entity in the repository.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"repo_name": {
							Description: "Name of the repository.",
							Type:        schema.TypeString,
							Optional:    true,
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
		},
	}

	helpers.SetMultiLevelDatasourceSchema(resource.Schema)

	return resource
}

func dataTemplateRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetClientWithContext(ctx)

	org_id := d.Get("org_id").(string)
	project_id := d.Get("project_id").(string)
	template_id := d.Get("identifier").(string)
	version := d.Get("version").(string)
	branch_name := helpers.BuildField(d, "branch_name")

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

	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	readDataTemplate(d, resp)

	return nil
}

func readDataTemplate(d *schema.ResourceData, template nextgen.TemplateWithInputsResponse) {
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
	d.Set("description", template.Template.Description)
	d.Set("version", template.Template.VersionLabel)
	d.Set("store_type", template.Template.StoreType)
	d.Set("connector_ref", template.Template.ConnectorRef)
	d.Set("child_type", template.Template.ChildType)
	if template.Template.GitDetails != nil {
		d.Set("git_details", []map[string]interface{}{
			{
				"branch_name":    template.Template.GitDetails.BranchName,
				"file_path":      template.Template.GitDetails.FilePath,
				"repo_name":      template.Template.GitDetails.RepoName,
				"repo_url":       template.Template.GitDetails.RepoUrl,
				"file_url":       template.Template.GitDetails.FileUrl,
				"last_commit_id": template.Template.GitDetails.CommitId,
				"last_object_id": template.Template.GitDetails.ObjectId,
			},
		})
	}
}
