package module_registry

import (
	"context"
	"encoding/json"
	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
	"net/http"
)

func ResourceInfraModule() *schema.Resource {
	resource := &schema.Resource{
		Description:   "Resource for managing Terraform/Tofu Modules.",
		ReadContext:   resourceInfraModuleRead,
		CreateContext: resourceInfraModuleCreate,
		UpdateContext: resourceInfraModuleUpdate,
		DeleteContext: resourceInfraModuleDelete,
		Importer:      helpers.AccountLevelResourceImporter,

		Schema: map[string]*schema.Schema{
			"description": {
				Description: "Description of the module.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"name": {
				Description: "Name of the module.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"system": {
				Description: "Provider of the module.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"repository": {
				Description: "For account connectors, the repository where the module can be found",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"repository_branch": {
				Description:  "Name of the branch to fetch the code from. This cannot be set if repository commit is set.",
				Type:         schema.TypeString,
				Optional:     true,
				ExactlyOneOf: []string{"repository_commit", "repository_branch"},
			},
			"repository_commit": {
				Description:  "Tag to fetch the code from. This cannot be set if repository branch is set.",
				Type:         schema.TypeString,
				Optional:     true,
				ExactlyOneOf: []string{"repository_commit", "repository_branch"},
			},
			"repository_connector": {
				Description: "Reference to the connector to be used to fetch the code.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"repository_path": {
				Description: "Path to the module within the repository.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"created": {
				Description: "Timestamp when the module was created.",
				Type:        schema.TypeInt,
				Computed:    true,
				Optional:    true,
			},
			"id": {
				Description: "Unique identifier of the module.",
				Type:        schema.TypeString,
				Computed:    true,
				Optional:    true,
			},
			"repository_url": {
				Description: "URL of the repository where the module is stored.",
				Type:        schema.TypeString,
				Computed:    true,
				Optional:    true,
			},
			"synced": {
				Description: "Timestamp when the module was last synced.",
				Type:        schema.TypeInt,
				Computed:    true,
				Optional:    true,
			},
			"tags": {
				Description: "Git tags associated with the module.",
				Type:        schema.TypeString,
				Computed:    true,
				Optional:    true,
			},
			"versions": {
				Description: "List of versions of the module.",
				Type:        schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Computed: true,
				Optional: true,
			},
		},
	}
	return resource
}

func resourceInfraModuleRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c, ctx := m.(*internal.Session).GetPlatformClientWithContext(ctx)
	id := d.Id()
	if id == "" {
		d.MarkNewResource()
	}
	resp, httpRes, err := c.ModuleRegistryApi.ModuleRegistryListModulesById(
		ctx,
		d.Get("id").(string),
		c.AccountId,
	)
	if err != nil {
		return helpers.HandleApiError(err, d, httpRes)
	}
	readModule(d, &resp)
	return nil
}

func resourceInfraModuleCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c, ctx := m.(*internal.Session).GetPlatformClientWithContext(ctx)
	id := d.Id()
	if id == "" {
		d.MarkNewResource()
	}
	createModule, err := buildCreateModuleRequestBody(d)
	if err != nil {
		return diag.FromErr(err)
	}
	log.Printf("[DEBUG] Creating module with value %v", createModule)
	modRes, httpRes, err := c.ModuleRegistryApi.ModuleRegistryCreateModule(ctx, createModule, c.AccountId)

	if err != nil {
		return parseError(err, httpRes)
	}
	setModuleId(d, &modRes)
	resourceInfraModuleRead(ctx, d, m)
	return nil
}

func resourceInfraModuleDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c, ctx := m.(*internal.Session).GetPlatformClientWithContext(ctx)
	id := d.Id()
	if id == "" {
		return nil
	}
	httpRes, err := c.ModuleRegistryApi.ModuleRegistryDeleteModule(ctx, id, c.AccountId)
	if err != nil {
		return parseError(err, httpRes)
	}
	return nil
}

func resourceInfraModuleUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c, ctx := m.(*internal.Session).GetPlatformClientWithContext(ctx)
	id := d.Id()
	if id == "" {
		d.MarkNewResource()
	}
	updateModule, err := buildUpdateModuleRequestBody(d)
	if err != nil {
		return diag.FromErr(err)
	}
	httpRes, err := c.ModuleRegistryApi.ModuleRegistryUpdateModule(ctx, updateModule, c.AccountId, id)
	if err != nil {
		return parseError(err, httpRes)
	}
	resourceInfraModuleRead(ctx, d, m)
	return nil
}

func setModuleId(d *schema.ResourceData, module *nextgen.CreateModuleResponseBody) {
	d.SetId(module.Id)
}

func readModule(d *schema.ResourceData, module *nextgen.ModuleResource) {
	d.SetId(module.Id)
	d.Set("account", module.Account)
	d.Set("created", module.Created)
	d.Set("description", module.Description)
	d.Set("git_tag_style", module.GitTagStyle)
	d.Set("id", module.Id)
	d.Set("module_error", module.ModuleError)
	d.Set("name", module.Name)
	d.Set("org", module.Org)
	d.Set("project", module.Project)
	d.Set("repository", module.Repository)
	d.Set("repository_branch", module.RepositoryBranch)
	d.Set("repository_commit", module.RepositoryCommit)
	d.Set("repository_connector", module.RepositoryConnector)
	d.Set("repository_path", module.RepositoryPath)
	d.Set("repository_url", module.RepositoryUrl)
	d.Set("synced", module.Synced)
	d.Set("system", module.System)
	d.Set("tags", module.Tags)
	d.Set("testing_enabled", module.TestingEnabled)
	d.Set("testing_metadata", module.TestingMetadata)
	d.Set("updated", module.Updated)
	d.Set("versions", module.Versions)
}

func buildCreateModuleRequestBody(d *schema.ResourceData) (nextgen.CreateModuleRequestBody, error) {
	module := nextgen.CreateModuleRequestBody{
		Name:   d.Get("name").(string),
		System: d.Get("system").(string),
	}

	if desc, ok := d.GetOk("description"); ok {
		module.Description = desc.(string)
	}
	if repo, ok := d.GetOk("repository"); ok {
		module.Repository = repo.(string)
	}
	if repoBranch, ok := d.GetOk("repository_branch"); ok {
		module.RepositoryBranch = repoBranch.(string)
	}
	if repoCommit, ok := d.GetOk("repository_commit"); ok {
		module.RepositoryCommit = repoCommit.(string)
	}
	if repoConnector, ok := d.GetOk("repository_connector"); ok {
		module.RepositoryConnector = repoConnector.(string)
	}
	if repoPath, ok := d.GetOk("repository_path"); ok {
		module.RepositoryPath = repoPath.(string)
	}
	return module, nil
}

func buildUpdateModuleRequestBody(d *schema.ResourceData) (nextgen.CreateModuleRequestBody, error) {
	module := nextgen.CreateModuleRequestBody{
		Name:   d.Get("name").(string),
		System: d.Get("system").(string),
	}

	if desc, ok := d.GetOk("description"); ok {
		module.Description = desc.(string)
	}
	if repo, ok := d.GetOk("repository"); ok {
		module.Repository = repo.(string)
	}
	if repoBranch, ok := d.GetOk("repository_branch"); ok {
		module.RepositoryBranch = repoBranch.(string)
	}
	if repoCommit, ok := d.GetOk("repository_commit"); ok {
		module.RepositoryCommit = repoCommit.(string)
	}
	if repoConnector, ok := d.GetOk("repository_connector"); ok {
		module.RepositoryConnector = repoConnector.(string)
	}
	if repoPath, ok := d.GetOk("repository_path"); ok {
		module.RepositoryPath = repoPath.(string)
	}
	return module, nil
}

func parseError(err error, httpResp *http.Response) diag.Diagnostics {
	if httpResp != nil && httpResp.StatusCode == 401 {
		return diag.Errorf(httpResp.Status + "\n" + "Hint:\n" +
			"1) Please check if token has expired or is wrong.\n" +
			"2) Harness Provider is misconfigured. For firstgen resources please give the correct api_key and for nextgen resources please give the correct platform_api_key.")
	}
	if httpResp != nil && httpResp.StatusCode == 403 {
		return diag.Errorf(httpResp.Status + "\n" + "Hint:\n" +
			"1) Please check if the token has required permission for this operation.\n" +
			"2) Please check if the token has expired or is wrong.")
	}

	se, ok := err.(nextgen.GenericSwaggerError)
	if !ok {
		diag.FromErr(err)
	}

	iacmErrBody := se.Body()
	iacmErr := nextgen.IacmError{}
	jsonErr := json.Unmarshal(iacmErrBody, &iacmErr)
	if jsonErr != nil {
		return diag.Errorf(err.Error())
	}

	return diag.Errorf(httpResp.Status + "\n" + "Hint:\n" +
		"1) " + iacmErr.Message)
}
