package idp

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/antihax/optional"
	"github.com/harness/harness-go-sdk/harness/idp"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"gopkg.in/yaml.v3"
)

type catalogEntityInfo struct {
	Scope      string
	Kind       string
	Identifier string
	OrgId      optional.String
	ProjectId  optional.String
}

func ResourceCatalogEntity() *schema.Resource {
	resource := &schema.Resource{
		Description:   "Resource for creating IDP catalog entities.",
		ReadContext:   resourceCatalogEntityRead,
		UpdateContext: resourceCatalogEntityUpdateOrCreate,
		CreateContext: resourceCatalogEntityUpdateOrCreate,
		DeleteContext: resourceCatalogEntityDelete,
		Importer:      entityImporter,
		Schema: map[string]*schema.Schema{
			"identifier": helpers.GetIdentifierSchema(helpers.SchemaFlagTypes.Required),
			"kind": {
				Type:        schema.TypeString,
				Description: "Kind of the catalog entity",
				Optional:    true,
				Computed:    true,
				ValidateFunc: validation.StringInSlice([]string{
					"component", "group", "user", "workflow", "resource", "system",
				}, false),
			},
			"org_id":     helpers.GetOrgIdSchema(helpers.SchemaFlagTypes.Optional),
			"project_id": helpers.GetProjectIdSchema(helpers.SchemaFlagTypes.Optional),
			"yaml": {
				Type:             schema.TypeString,
				Description:      "YAML definition of the catalog entity",
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: helpers.YamlDiffSuppressFunction,
			},
			"import_from_git": {
				Description: "Flag to set if importing from Git",
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
			},
			"git_details": {
				Description: "Contains Git Information for importing entities from Git",
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
							Description: "Identifier of the Harness Connector used for importing entity from Git" + helpers.Descriptions.ConnectorRefText.String(),
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
						"is_harness_code_repo": {
							Description: "If the repo is a Harness Code repo",
							Type:        schema.TypeBool,
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
		},
	}
	resource.Schema["project_id"].RequiredWith = []string{"org_id"}

	return resource
}

func resourceCatalogEntityRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetIDPClientWithContext(ctx)

	entityInfo, err := getCatalogEntityInfoFromResourceData(d)
	if err != nil {
		return diag.Errorf("error in reading catalog entity inputs: %v", err)
	}

	id := d.Id()
	if id == "" {
		id = entityInfo.Identifier
	}

	resp, httpResp, err := c.EntitiesApi.GetEntity(ctx, entityInfo.Scope, entityInfo.Kind, id, &idp.EntitiesApiGetEntityOpts{
		OrgIdentifier:     entityInfo.OrgId,
		ProjectIdentifier: entityInfo.ProjectId,
		HarnessAccount:    optional.NewString(c.AccountId),
	})

	if err != nil {
		return handleIDPReadApiError(err, d, httpResp)
	}

	readCatalogEntity(d, resp)

	return nil
}

func resourceCatalogEntityUpdateOrCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetIDPClientWithContext(ctx)

	var err error
	var resp idp.EntityResponse
	var httpResp *http.Response
	var entityInfo catalogEntityInfo

	id := d.Id()
	isNewResource := id == ""
	if id == "" {
		if d.Get("import_from_git").(bool) {
			entityInfo, err = getCatalogEntityInfoFromImportResourceData(d)
			if err != nil {
				return diag.Errorf("failed to get catalog entity info: %v", err)
			}

			importInfo := idp.GitImportDetails{}
			if attr, ok := d.GetOk("git_details"); ok {
				config := attr.([]interface{})[0].(map[string]interface{})
				if attr, ok := config["branch_name"]; ok {
					importInfo.BranchName = attr.(string)
				}
				if attr, ok := config["file_path"]; ok {
					importInfo.FilePath = attr.(string)
				}
				if attr, ok := config["connector_ref"]; ok {
					importInfo.ConnectorRef = attr.(string)
				}
				if attr, ok := config["repo_name"]; ok {
					importInfo.RepoName = attr.(string)
				}
				if attr, ok := config["is_harness_code_repo"]; ok {
					importInfo.IsHarnessCodeRepo = attr.(bool)
				}
			}

			orgId := optional.EmptyString()
			projectId := optional.EmptyString()

			v, ok := d.GetOk("org_id")
			if ok {
				orgId = optional.NewString(v.(string))
			}

			v, ok = d.GetOk("project_id")
			if ok {
				projectId = optional.NewString(v.(string))
			}

			resp, httpResp, err = c.EntitiesApi.ImportEntity(ctx, importInfo, &idp.EntitiesApiImportEntityOpts{
				HarnessAccount:    optional.NewString(c.AccountId),
				OrgIdentifier:     orgId,
				ProjectIdentifier: projectId,
			})
			if err != nil {
				return handleIDPApiError(err, d, httpResp)
			}
			entityInfo = getCatalogEntityInfoFromResponse(resp, entityInfo)
		} else {
			entityInfo, err = getAndVerifyCatalogEntityInfo(d)
			if err != nil {
				return diag.Errorf("failed to get and verify catalog entity info: %v", err)
			}
			gitDetails := buildGitCreateDetails(d)
			yaml := d.Get("yaml").(string)

			resp, httpResp, err = c.EntitiesApi.CreateEntity(ctx, idp.EntityCreateRequest{
				Yaml:       yaml,
				GitDetails: gitDetails,
			},
				&idp.EntitiesApiCreateEntityOpts{
					OrgIdentifier:     entityInfo.OrgId,
					ProjectIdentifier: entityInfo.ProjectId,
					HarnessAccount:    optional.NewString(c.AccountId),
				})
			if err == nil {
				entityInfo = getCatalogEntityInfoFromResponse(resp, entityInfo)
			}
		}
	} else {
		gitDetails := buildGitUpdateDetails(d)

		connectorRefChanged := d.HasChange("git_details.0.connector_ref")
		filePathChanged := d.HasChange("git_details.0.file_path")
		repoNameChanged := d.HasChange("git_details.0.repo_name")
		shouldUpdateGitDetails := connectorRefChanged || filePathChanged || repoNameChanged

		yaml := d.Get("yaml").(string)
		entityInfo, err = getAndVerifyCatalogEntityInfo(d)
		if err != nil {
			return diag.Errorf("failed to get and verify catalog entity info: %v", err)
		}

		resp, httpResp, err = c.EntitiesApi.UpdateEntity(ctx, idp.EntityUpdateRequest{
			Yaml:       yaml,
			GitDetails: gitDetails,
		}, entityInfo.Scope, entityInfo.Kind, id, &idp.EntitiesApiUpdateEntityOpts{
			OrgIdentifier:     entityInfo.OrgId,
			ProjectIdentifier: entityInfo.ProjectId,
			HarnessAccount:    optional.NewString(c.AccountId),
		})
		if err != nil {
			return handleIDPApiError(err, d, httpResp)
		}

		if shouldUpdateGitDetails {
			diags := resourceCatalogEntityUpdateGitMetadata(ctx, c, d, entityInfo)
			if diags.HasError() {
				return diags
			}
		}
	}

	if err != nil {
		return handleIDPApiError(err, d, httpResp)
	}

	if isNewResource {
		resp, httpResp, err = getCatalogEntityWithRetry(ctx, c, entityInfo)
	} else {
		resp, httpResp, err = getCatalogEntity(ctx, c, entityInfo)
	}
	if err != nil {
		return handleIDPApiError(err, d, httpResp)
	}

	readCatalogEntity(d, resp)
	return nil
}

func getCatalogEntity(ctx context.Context, c *idp.APIClient, info catalogEntityInfo) (idp.EntityResponse, *http.Response, error) {
	return c.EntitiesApi.GetEntity(ctx, info.Scope, info.Kind, info.Identifier, &idp.EntitiesApiGetEntityOpts{
		OrgIdentifier:     info.OrgId,
		ProjectIdentifier: info.ProjectId,
		HarnessAccount:    optional.NewString(c.AccountId),
	})
}

func getCatalogEntityWithRetry(ctx context.Context, c *idp.APIClient, info catalogEntityInfo) (idp.EntityResponse, *http.Response, error) {
	const attempts = 6
	const delay = 2 * time.Second

	var resp idp.EntityResponse
	var httpResp *http.Response
	var err error

	for attempt := 0; attempt < attempts; attempt++ {
		resp, httpResp, err = getCatalogEntity(ctx, c, info)
		if err == nil {
			return resp, httpResp, nil
		}

		if !isTransientPostWriteReadError(err, httpResp) || attempt == attempts-1 {
			return resp, httpResp, err
		}

		timer := time.NewTimer(delay)
		select {
		case <-ctx.Done():
			if !timer.Stop() {
				<-timer.C
			}
			return resp, httpResp, ctx.Err()
		case <-timer.C:
		}
	}

	return resp, httpResp, err
}

func isTransientPostWriteReadError(err error, httpResp *http.Response) bool {
	if err == nil {
		return false
	}
	if httpResp != nil {
		return httpResp.StatusCode == http.StatusNotFound
	}
	return isNotFoundError(err)
}

func handleIDPApiError(err error, d *schema.ResourceData, httpResp *http.Response) diag.Diagnostics {
	if msg := idpAPIErrorMessage(err); msg != "" {
		return diag.Errorf("%s", msg)
	}

	return helpers.HandleApiError(err, d, httpResp)
}

func handleIDPReadApiError(err error, d *schema.ResourceData, httpResp *http.Response) diag.Diagnostics {
	if msg := idpAPIErrorMessage(err); msg != "" && !isIDPNotFoundError(err, httpResp) {
		return diag.Errorf("%s", msg)
	}

	return helpers.HandleReadApiError(err, d, httpResp)
}

func idpAPIErrorMessage(err error) string {
	if err == nil {
		return ""
	}

	swaggerErr, ok := err.(interface {
		Body() []byte
	})
	if !ok {
		return ""
	}

	var body struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	}
	if json.Unmarshal(swaggerErr.Body(), &body) != nil || body.Message == "" {
		return ""
	}
	if body.Code != "" {
		return fmt.Sprintf("%s: %s", body.Code, body.Message)
	}
	return body.Message
}

func isIDPNotFoundError(err error, httpResp *http.Response) bool {
	if httpResp != nil && httpResp.StatusCode == http.StatusNotFound {
		return true
	}

	swaggerErr, ok := err.(interface {
		Body() []byte
	})
	if !ok {
		return false
	}

	var body struct {
		Code string `json:"code"`
	}
	if json.Unmarshal(swaggerErr.Body(), &body) != nil {
		return false
	}

	return body.Code == "ENTITY_NOT_FOUND" || body.Code == "RESOURCE_NOT_FOUND"
}

func resourceCatalogEntityUpdateGitMetadata(ctx context.Context, c *idp.APIClient, d *schema.ResourceData, info catalogEntityInfo) diag.Diagnostics {
	httpResp, err := c.EntitiesApi.UpdateGitMetadata(ctx, idp.GitMetadataUpdateRequest{
		ConnectorRef: d.Get("git_details.0.connector_ref").(string),
		RepoName:     d.Get("git_details.0.repo_name").(string),
		FilePath:     d.Get("git_details.0.file_path").(string),
	}, info.Scope, info.Kind, info.Identifier, &idp.EntitiesApiUpdateGitMetadataOpts{
		HarnessAccount:    optional.NewString(c.AccountId),
		OrgIdentifier:     info.OrgId,
		ProjectIdentifier: info.ProjectId,
	})
	if err != nil {
		return helpers.HandleGitApiErrorWithResourceData(err, d, httpResp)
	}

	return nil
}

func resourceCatalogEntityDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetIDPClientWithContext(ctx)

	id := d.Id()
	entityInfo, err := getCatalogEntityInfoFromResourceData(d)
	if err != nil {
		return diag.Errorf("failed to get catalog entity info: %v", err)
	}

	httpResp, err := c.EntitiesApi.DeleteEntity(ctx, entityInfo.Scope, entityInfo.Kind, id, &idp.EntitiesApiDeleteEntityOpts{
		OrgIdentifier:     entityInfo.OrgId,
		ProjectIdentifier: entityInfo.ProjectId,
		HarnessAccount:    optional.NewString(c.AccountId),
	})
	if err != nil {
		if httpResp != nil && httpResp.StatusCode == 404 {
			d.SetId("")
			return nil
		}

		if isNotFoundError(err) {
			d.SetId("")
			return nil
		}

		return handleIDPApiError(err, d, httpResp)
	}

	return nil
}

func readCatalogEntity(d *schema.ResourceData, entity idp.EntityResponse) {
	d.SetId(entity.Identifier)
	d.Set("identifier", entity.Identifier)
	d.Set("kind", entity.Kind)
	d.Set("org_id", entity.OrgIdentifier)
	d.Set("project_id", entity.ProjectIdentifier)
	d.Set("yaml", entity.Yaml)
	if entity.GitDetails != nil {
		storeType := helpers.BuildField(d, "git_details.0.store_type")
		baseBranch := helpers.BuildField(d, "git_details.0.base_branch")
		commitMessage := helpers.BuildField(d, "git_details.0.commit_message")
		connectorRef := helpers.BuildField(d, "git_details.0.connector_ref")

		d.Set("git_details", []any{readGitDetails(entity, storeType, baseBranch, commitMessage, connectorRef)})
	} else {
		d.Set("git_details", []interface{}{})
	}
}

func readGitDetails(entity idp.EntityResponse, store_type optional.String, base_branch optional.String, commit_message optional.String, connector_ref optional.String) map[string]interface{} {
	git_details := map[string]interface{}{
		"branch_name":    entity.GitDetails.BranchName,
		"file_path":      entity.GitDetails.FilePath,
		"repo_name":      entity.GitDetails.RepoName,
		"last_commit_id": entity.GitDetails.CommitId,
		"last_object_id": entity.GitDetails.ObjectId,
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
	if !connector_ref.IsSet() || connector_ref.Value() == "" {
		git_details["is_harness_code_repo"] = true
	}

	return git_details
}

func getAndVerifyCatalogEntityInfo(d *schema.ResourceData) (catalogEntityInfo, error) {
	yamlString := d.Get("yaml").(string)
	kind := d.Get("kind").(string)
	identifier := d.Get("identifier").(string)
	orgId := d.Get("org_id").(string)
	projectId := d.Get("project_id").(string)

	var yamlData map[string]any
	if err := yaml.Unmarshal([]byte(yamlString), &yamlData); err != nil {
		return catalogEntityInfo{}, err
	}

	yamlKind, ok := yamlData["kind"].(string)
	if !ok || yamlKind == "" {
		return catalogEntityInfo{}, fmt.Errorf("kind is missing from YAML")
	}
	if !strings.EqualFold(yamlKind, kind) {
		return catalogEntityInfo{}, fmt.Errorf("kind in YAML (%s) does not match kind parameter (%s)", yamlKind, kind)
	}

	yamlIdentifier, ok := yamlData["identifier"].(string)
	if !ok || yamlIdentifier == "" {
		return catalogEntityInfo{}, fmt.Errorf("identifier is missing from YAML")
	}
	if yamlIdentifier != identifier {
		return catalogEntityInfo{}, fmt.Errorf("identifier in YAML (%s) does not match identifier parameter (%s)", yamlIdentifier, identifier)
	}

	yamlProject := ""
	if project, ok := yamlData["projectIdentifier"].(string); ok && project != "" {
		yamlProject = project
	}

	if yamlProject != projectId {
		return catalogEntityInfo{}, fmt.Errorf("projectIdentifier in YAML (%s) does not match project_id parameter (%s)", yamlProject, projectId)
	}

	yamlOrg := ""
	if org, ok := yamlData["orgIdentifier"].(string); ok && org != "" {
		yamlOrg = org
	}

	if yamlOrg != orgId {
		return catalogEntityInfo{}, fmt.Errorf("orgIdentifier in YAML (%s) does not match org_id parameter (%s)", yamlOrg, orgId)
	}

	catalogInfo := catalogEntityInfo{
		Kind:       kind,
		Scope:      "account",
		Identifier: identifier,
	}

	if yamlOrg != "" {
		catalogInfo.OrgId = optional.NewString(yamlOrg)
		catalogInfo.Scope = fmt.Sprintf("%s.%s", catalogInfo.Scope, yamlOrg)
	} else {
		catalogInfo.OrgId = optional.EmptyString()
	}

	if yamlProject != "" {
		catalogInfo.ProjectId = optional.NewString(yamlProject)
		catalogInfo.Scope = fmt.Sprintf("%s.%s", catalogInfo.Scope, yamlProject)
	} else {
		catalogInfo.ProjectId = optional.EmptyString()
	}

	return catalogInfo, nil
}

func getCatalogEntityInfoFromResourceData(d *schema.ResourceData) (catalogEntityInfo, error) {
	kind := d.Get("kind").(string)
	identifier := d.Get("identifier").(string)
	orgId := d.Get("org_id").(string)
	projectId := d.Get("project_id").(string)

	if identifier == "" {
		identifier = d.Id()
	}
	if identifier == "" {
		return catalogEntityInfo{}, fmt.Errorf("identifier is required")
	}
	if kind == "" {
		return catalogEntityInfo{}, fmt.Errorf("kind is required")
	}

	catalogInfo := catalogEntityInfo{
		Kind:       kind,
		Scope:      "account",
		Identifier: identifier,
	}

	if orgId != "" {
		catalogInfo.OrgId = optional.NewString(orgId)
		catalogInfo.Scope = fmt.Sprintf("%s.%s", catalogInfo.Scope, orgId)
	} else {
		catalogInfo.OrgId = optional.EmptyString()
	}

	if projectId != "" {
		catalogInfo.ProjectId = optional.NewString(projectId)
		catalogInfo.Scope = fmt.Sprintf("%s.%s", catalogInfo.Scope, projectId)
	} else {
		catalogInfo.ProjectId = optional.EmptyString()
	}

	return catalogInfo, nil
}

func getCatalogEntityInfoFromImportResourceData(d *schema.ResourceData) (catalogEntityInfo, error) {
	identifier := d.Get("identifier").(string)
	if identifier == "" {
		identifier = d.Id()
	}
	if identifier == "" {
		return catalogEntityInfo{}, fmt.Errorf("identifier is required")
	}

	orgId := optional.EmptyString()
	if v, ok := d.GetOk("org_id"); ok && v.(string) != "" {
		orgId = optional.NewString(v.(string))
	}

	projectId := optional.EmptyString()
	if v, ok := d.GetOk("project_id"); ok && v.(string) != "" {
		projectId = optional.NewString(v.(string))
	}

	return catalogEntityInfo{
		Kind:       d.Get("kind").(string),
		Scope:      catalogEntityScope(orgId, projectId),
		Identifier: identifier,
		OrgId:      orgId,
		ProjectId:  projectId,
	}, nil
}

func getCatalogEntityInfoFromResponse(entity idp.EntityResponse, fallback catalogEntityInfo) catalogEntityInfo {
	catalogInfo := fallback

	if entity.Identifier != "" {
		catalogInfo.Identifier = entity.Identifier
	}
	if entity.Kind != "" {
		catalogInfo.Kind = entity.Kind
	}
	if catalogInfo.Scope == "" {
		catalogInfo.Scope = "account"
	}

	responseHasScope := hasUsableResponseScope(entity, fallback)
	if responseHasScope {
		catalogInfo.OrgId = optional.NewString(entity.OrgIdentifier)
		if entity.ProjectIdentifier != "" {
			catalogInfo.ProjectId = optional.NewString(entity.ProjectIdentifier)
		} else {
			catalogInfo.ProjectId = optional.EmptyString()
		}
		catalogInfo.Scope = catalogEntityScope(catalogInfo.OrgId, catalogInfo.ProjectId)
	}

	return catalogInfo
}

func hasUsableResponseScope(entity idp.EntityResponse, fallback catalogEntityInfo) bool {
	if entity.OrgIdentifier == "" {
		return false
	}

	return !fallback.ProjectId.IsSet() || fallback.ProjectId.Value() == "" || entity.ProjectIdentifier != ""
}

func catalogEntityScope(orgId optional.String, projectId optional.String) string {
	scope := "account"
	if orgId.IsSet() && orgId.Value() != "" {
		scope = fmt.Sprintf("%s.%s", scope, orgId.Value())
	}
	if projectId.IsSet() && projectId.Value() != "" {
		scope = fmt.Sprintf("%s.%s", scope, projectId.Value())
	}

	return scope
}

func catalogEntityInfoFromImportID(id string) (catalogEntityInfo, error) {
	parts := strings.Split(id, "/")

	if len(parts) < 2 || len(parts) > 4 {
		return catalogEntityInfo{}, fmt.Errorf("invalid import ID format: %s. Expected: <kind>/<identifier>, <org>/<kind>/<identifier>, or <org>/<project>/<kind>/<identifier>", id)
	}

	catalogInfo := catalogEntityInfo{
		Scope: "account",
	}

	if len(parts) == 2 {
		catalogInfo.Kind = parts[0]
		catalogInfo.Identifier = parts[1]
	} else if len(parts) == 3 {
		scopePrefix, err := catalogEntityScopeFromImportPrefix(parts[0])
		if err != nil {
			return catalogEntityInfo{}, err
		}
		catalogInfo.Scope = scopePrefix.Scope
		catalogInfo.OrgId = scopePrefix.OrgId
		catalogInfo.ProjectId = scopePrefix.ProjectId
		catalogInfo.Kind = parts[1]
		catalogInfo.Identifier = parts[2]
	} else {
		catalogInfo.OrgId = optional.NewString(parts[0])
		catalogInfo.ProjectId = optional.NewString(parts[1])
		catalogInfo.Scope = fmt.Sprintf("%s.%s.%s", catalogInfo.Scope, parts[0], parts[1])
		catalogInfo.Kind = parts[2]
		catalogInfo.Identifier = parts[3]
	}

	return catalogInfo, nil
}

func catalogEntityScopeFromImportPrefix(prefix string) (catalogEntityInfo, error) {
	scopeParts := strings.Split(prefix, ".")
	if len(scopeParts) < 1 || len(scopeParts) > 2 || scopeParts[0] == "" {
		return catalogEntityInfo{}, fmt.Errorf("invalid import scope: %s. Expected: <org> or <org>.<project>", prefix)
	}

	catalogInfo := catalogEntityInfo{
		Scope: "account",
		OrgId: optional.NewString(scopeParts[0]),
	}
	catalogInfo.Scope = fmt.Sprintf("%s.%s", catalogInfo.Scope, scopeParts[0])

	if len(scopeParts) == 2 {
		if scopeParts[1] == "" {
			return catalogEntityInfo{}, fmt.Errorf("invalid import scope: %s. Expected: <org> or <org>.<project>", prefix)
		}
		catalogInfo.ProjectId = optional.NewString(scopeParts[1])
		catalogInfo.Scope = fmt.Sprintf("%s.%s", catalogInfo.Scope, scopeParts[1])
	} else {
		catalogInfo.ProjectId = optional.EmptyString()
	}

	return catalogInfo, nil
}

func buildGitCreateDetails(d *schema.ResourceData) *idp.GitCreateDetails {
	if _, ok := d.GetOk("git_details"); !ok {
		return nil
	}

	config := d.Get("git_details").([]interface{})[0].(map[string]interface{})
	details := &idp.GitCreateDetails{}
	if attr, ok := config["branch_name"]; ok {
		details.BranchName = attr.(string)
	}
	if attr, ok := config["file_path"]; ok {
		details.FilePath = attr.(string)
	}
	if attr, ok := config["commit_message"]; ok {
		details.CommitMessage = attr.(string)
	}
	if attr, ok := config["base_branch"]; ok {
		details.BaseBranch = attr.(string)
	}
	if attr, ok := config["connector_ref"]; ok {
		details.ConnectorRef = attr.(string)
	}
	if attr, ok := config["store_type"]; ok {
		details.StoreType = attr.(string)
	}
	if attr, ok := config["repo_name"]; ok {
		details.RepoName = attr.(string)
	}
	if attr, ok := config["is_harness_code_repo"]; ok {
		details.IsHarnessCodeRepo = attr.(bool)
	}

	return details
}

func buildGitUpdateDetails(d *schema.ResourceData) *idp.GitUpdateDetails {
	if _, ok := d.GetOk("git_details"); !ok {
		return nil
	}

	config := d.Get("git_details").([]interface{})[0].(map[string]interface{})
	details := &idp.GitUpdateDetails{}

	if attr, ok := config["branch_name"]; ok {
		details.BranchName = attr.(string)
	}
	if attr, ok := config["commit_message"]; ok {
		details.CommitMessage = attr.(string)
	}
	if attr, ok := config["base_branch"]; ok {
		details.BaseBranch = attr.(string)
	}
	if attr, ok := config["last_object_id"]; ok {
		details.LastObjectId = attr.(string)
	}
	if attr, ok := config["last_commit_id"]; ok {
		details.LastCommitId = attr.(string)
	}
	if attr, ok := config["is_harness_code_repo"]; ok {
		details.IsHarnessCodeRepo = attr.(bool)
	}
	if attr, ok := config["store_type"]; ok {
		details.StoreType = attr.(string)
	}
	if attr, ok := config["connector_ref"]; ok {
		details.ConnectorRef = attr.(string)
	}
	if attr, ok := config["repo_name"]; ok {
		details.RepoName = attr.(string)
	}
	if attr, ok := config["file_path"]; ok {
		details.FilePath = attr.(string)
	}

	return details
}

var entityImporter = &schema.ResourceImporter{
	State: func(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
		// Expected format: <kind>/<identifier>, <org>/<kind>/<identifier>, or <org>/<project>/<kind>/<identifier>
		// If account-level: <kind>/<identifier>
		info, err := catalogEntityInfoFromImportID(d.Id())
		if err != nil {
			return nil, err
		}

		c, ctx := meta.(*internal.Session).GetIDPClientWithContext(context.Background())

		resp, _, err := c.EntitiesApi.GetEntity(ctx, info.Scope, info.Kind, info.Identifier, &idp.EntitiesApiGetEntityOpts{
			OrgIdentifier:     info.OrgId,
			ProjectIdentifier: info.ProjectId,
			HarnessAccount:    optional.NewString(c.AccountId),
		})
		if err != nil {
			return nil, fmt.Errorf("failed to fetch entity for import: %w", err)
		}

		readCatalogEntity(d, resp)

		return []*schema.ResourceData{d}, nil
	},
}
