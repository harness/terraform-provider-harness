package service_overrides_v2

import (
	"context"
	"net/http"

	"github.com/antihax/optional"
	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func ResourceServiceOverrides() *schema.Resource {
	resource := &schema.Resource{
		Description: "Resource for creating a Harness service override V2.",

		ReadContext:   resourceServiceOverridesV2Read,
		UpdateContext: resourceServiceOverridesV2CreateOrUpdate,
		DeleteContext: resourceServiceOverridesV2Delete,
		CreateContext: resourceServiceOverridesV2CreateOrUpdate,
		Importer:      helpers.ServiceOverrideV2ResourceImporter,

		Schema: map[string]*schema.Schema{
			"service_id": {
				Description: "The service ID to which the overrides applies.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"env_id": {
				Description: "The environment ID to which the overrides are associated.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"infra_id": {
				Description: "The infrastructure ID to which the overrides are associated.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"cluster_id": {
				Description: "The cluster ID to which the overrides are associated.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"type": {
				Description: "The type of the overrides.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"yaml": {
				Description:      "The yaml of the overrides spec object.",
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: helpers.YamlDiffSuppressFunction,
			},
			"identifier": {
				Description: "The identifier of the override entity.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"is_force_import": {
				Description: "force import override from remote even if same file path already exist",
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
			},
			"import_from_git": {
				Description: "import override from git",
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
			},
			"git_details": {
				Description: "Contains parameters related to creating an Entity for Git Experience.",
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"branch": {
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
							ValidateFunc: validation.StringInSlice([]string{"INLINE", "REMOTE"}, false),
							Computed:     true,
						},
						"repo_name": {
							Description: "Name of the repository.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"last_object_id": {
							Description: "Last object identifier (for Github). To be provided only when updating override.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"last_commit_id": {
							Description: "Last commit identifier (for Git Repositories other than Github). To be provided only when updating override.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"is_harness_code_repo": {
							Description: "If the repo is in harness code",
							Type:        schema.TypeBool,
							Optional:    true,
							Computed:    true,
						},
						"is_new_branch": {
							Description: "If the branch being created is new",
							Type:        schema.TypeBool,
							Optional:    true,
							Computed:    true,
						},
						"load_from_cache": {
							Description: "Load service yaml from catch",
							Type:        schema.TypeBool,
							Optional:    true,
							Computed:    true,
						},
						"load_from_fallback_branch": {
							Description: "Load service yaml from fallback branch",
							Type:        schema.TypeBool,
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
		},
	}

	SetScopedResourceSchemaForServiceOverride(resource.Schema)

	return resource
}

func resourceServiceOverridesV2Read(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)
	identifier := d.Id()
	svcGetParams := getSvcOverrideParams(d)

	resp, httpResp, err := c.ServiceOverridesApi.GetServiceOverridesV2(ctx, identifier, c.AccountId, svcGetParams)
	if err != nil {
		return helpers.HandleReadApiError(err, d, httpResp)
	}

	// GET call for service environment override returns a 200 ok for empty list.
	// Hence specifically marking the resource as new, instead of in L#91
	if &resp == nil || resp.Data == nil {
		d.SetId("")
		d.MarkNewResource()
		return nil
	}

	readServiceOverridesV2(d, resp.Data)

	return nil
}

func resourceServiceOverridesV2CreateOrUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	var err error
	var resp nextgen.ResponseServiceOverridesResponseDtov2
	var importResp nextgen.ResponseServiceOverrideImportResponseDto
	var httpResp *http.Response
	env := buildServiceOverrideV2(d)

	id := d.Id()

	if id == "" {
		if d.Get("import_from_git").(bool) {
			importReq := buildServiceOverrideImportRequest(d)
			svcImportParam := getSvcOverrideImportParams(importReq, d)
			importResp, httpResp, err = c.ServiceOverridesApi.ImportServiceOverrides(ctx, c.AccountId, svcImportParam)
		} else {
			svcCreateParam := svcOverrideCreateParam(env, d)
			resp, httpResp, err = c.ServiceOverridesApi.CreateServiceOverrideV2(ctx, c.AccountId, svcCreateParam)
		}
	} else {
		svcUpdateParam := svcOverrideUpdateParam(env, d)
		resp, httpResp, err = c.ServiceOverridesApi.UpdateServiceOverrideV2(ctx, c.AccountId, svcUpdateParam)
	}

	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	// Soft delete lookup error handling
	// https://harness.atlassian.net/browse/PL-23765
	if (&resp == nil || resp.Data == nil) && !d.Get("import_from_git").(bool) {
		d.SetId("")
		d.MarkNewResource()
		return nil
	}

	if d.Get("import_from_git").(bool) {
		readImportServiceOverridesV2(d, importResp.Data)
	} else {
		readServiceOverridesV2(d, resp.Data)
	}

	return nil
}

func resourceServiceOverridesV2Delete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	_, httpResp, err := c.ServiceOverridesApi.DeleteServiceOverrideV2(ctx, d.Id(), c.AccountId, &nextgen.ServiceOverridesApiDeleteServiceOverrideV2Opts{
		OrgIdentifier:     helpers.BuildField(d, "org_id"),
		ProjectIdentifier: helpers.BuildField(d, "project_id"),
	})

	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	return nil
}

func buildServiceOverrideV2(d *schema.ResourceData) *nextgen.ServiceOverrideRequestDtov2 {
	return &nextgen.ServiceOverrideRequestDtov2{
		OrgIdentifier:     d.Get("org_id").(string),
		ProjectIdentifier: d.Get("project_id").(string),
		EnvironmentRef:    d.Get("env_id").(string),
		ServiceRef:        d.Get("service_id").(string),
		InfraIdentifier:   d.Get("infra_id").(string),
		ClusterIdentifier: d.Get("cluster_id").(string),
		Type_:             d.Get("type").(string),
		YamlInternal:      d.Get("yaml").(string),
	}
}

func buildServiceOverrideImportRequest(d *schema.ResourceData) *nextgen.ServiceOverrideImportRequestDto {
	return &nextgen.ServiceOverrideImportRequestDto{
		OrgIdentifier:     d.Get("org_id").(string),
		ProjectIdentifier: d.Get("project_id").(string),
		EnvironmentRef:    d.Get("env_id").(string),
		ServiceRef:        d.Get("service_id").(string),
		InfraIdentifier:   d.Get("infra_id").(string),
		Type_:             d.Get("type").(string),
	}
}

func readServiceOverridesV2(d *schema.ResourceData, so *nextgen.ServiceOverridesResponseDtov2) {
	d.SetId(so.Identifier)
	d.Set("org_id", so.OrgIdentifier)
	d.Set("project_id", so.ProjectIdentifier)
	d.Set("env_id", so.EnvironmentRef)
	d.Set("service_id", so.ServiceRef)
	d.Set("infra_id", so.InfraIdentifier)
	d.Set("cluster_id", so.ClusterIdentifier)
	d.Set("type", so.Type_)
	d.Set("yaml", so.YamlInternal)
	d.Set("identifier", so.Identifier)
}

func readImportServiceOverridesV2(d *schema.ResourceData, so *nextgen.ServiceOverrideImportResponseDto) {
	d.SetId(so.Identifier)
	d.Set("env_id", so.EnvironmentRef)
	d.Set("service_id", so.ServiceRef)
	d.Set("infra_id", so.InfraIdentifier)
	d.Set("type", so.Type_)
	d.Set("identifier", so.Identifier)
}

func SetScopedResourceSchemaForServiceOverride(s map[string]*schema.Schema) {
	s["project_id"] = helpers.GetProjectIdSchema(helpers.SchemaFlagTypes.Optional)
	s["org_id"] = helpers.GetOrgIdSchema(helpers.SchemaFlagTypes.Optional)
}

func getSvcOverrideParams(d *schema.ResourceData) *nextgen.ServiceOverridesApiGetServiceOverridesV2Opts {
	svcOverrideParams := &nextgen.ServiceOverridesApiGetServiceOverridesV2Opts{
		OrgIdentifier:          helpers.BuildField(d, "org_id"),
		ProjectIdentifier:      helpers.BuildField(d, "project_id"),
		Branch:                 helpers.BuildField(d, "git_details.0.branch"),
		RepoName:               helpers.BuildField(d, "git_details.0.repo_name"),
		LoadFromCache:          helpers.BuildField(d, "git_details.0.load_from_cache"),
		LoadFromFallbackBranch: helpers.BuildFieldBool(d, "git_details.0.load_from_fallback_branch"),
	}
	return svcOverrideParams
}

func svcOverrideCreateParam(svcOverride *nextgen.ServiceOverrideRequestDtov2, d *schema.ResourceData) *nextgen.ServiceOverridesApiCreateServiceOverrideV2Opts {
	return &nextgen.ServiceOverridesApiCreateServiceOverrideV2Opts{
		Body:              optional.NewInterface(svcOverride),
		Branch:            helpers.BuildField(d, "git_details.0.branch"),
		FilePath:          helpers.BuildField(d, "git_details.0.file_path"),
		CommitMsg:         helpers.BuildField(d, "git_details.0.commit_message"),
		IsNewBranch:       helpers.BuildFieldBool(d, "git_details.0.is_new_branch"),
		BaseBranch:        helpers.BuildField(d, "git_details.0.base_branch"),
		ConnectorRef:      helpers.BuildField(d, "git_details.0.connector_ref"),
		StoreType:         helpers.BuildField(d, "git_details.0.store_type"),
		RepoName:          helpers.BuildField(d, "git_details.0.repo_name"),
		IsHarnessCodeRepo: helpers.BuildFieldBool(d, "git_details.0.is_harness_code_repo"),
	}
}

func svcOverrideUpdateParam(svcOverride *nextgen.ServiceOverrideRequestDtov2, d *schema.ResourceData) *nextgen.ServiceOverridesApiUpdateServiceOverrideV2Opts {
	return &nextgen.ServiceOverridesApiUpdateServiceOverrideV2Opts{
		Body:              optional.NewInterface(svcOverride),
		Branch:            helpers.BuildField(d, "git_details.0.branch"),
		FilePath:          helpers.BuildField(d, "git_details.0.file_path"),
		CommitMsg:         helpers.BuildField(d, "git_details.0.commit_message"),
		IsNewBranch:       helpers.BuildFieldBool(d, "git_details.0.is_new_branch"),
		BaseBranch:        helpers.BuildField(d, "git_details.0.base_branch"),
		ConnectorRef:      helpers.BuildField(d, "git_details.0.connector_ref"),
		StoreType:         helpers.BuildField(d, "git_details.0.store_type"),
		LastObjectId:      helpers.BuildField(d, "git_details.0.last_object_id"),
		LastCommitId:      helpers.BuildField(d, "git_details.0.last_commit_id"),
		IsHarnessCodeRepo: helpers.BuildFieldBool(d, "git_details.0.is_harness_code_repo"),
	}
}

func getSvcOverrideImportParams(svcOverride *nextgen.ServiceOverrideImportRequestDto, d *schema.ResourceData) *nextgen.ServiceOverridesApiImportServiceOverridesOpts {
	svcOverrideParams := &nextgen.ServiceOverridesApiImportServiceOverridesOpts{
		Body:              optional.NewInterface(svcOverride),
		Branch:            helpers.BuildField(d, "git_details.0.branch"),
		RepoName:          helpers.BuildField(d, "git_details.0.repo_name"),
		FilePath:          helpers.BuildField(d, "git_details.0.file_path"),
		ConnectorRef:      helpers.BuildField(d, "git_details.0.connector_ref"),
		IsHarnessCodeRepo: helpers.BuildFieldBool(d, "git_details.0.is_harness_code_repo"),
		IsForceImport:     helpers.BuildFieldBool(d, "is_force_import"),
	}
	return svcOverrideParams
}
