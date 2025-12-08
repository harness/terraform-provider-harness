package module_registry_testing

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceInfraModuleTesting() *schema.Resource {
	resource := &schema.Resource{
		Description:   "Resource for managing Terraform/Tofu Modules.",
		ReadContext:   resourceInfraModuleTestingRead,
		CreateContext: resourceInfraModuleTestingCreate,
		UpdateContext: resourceInfraModuleTestingUpdate,
		DeleteContext: resourceInfraModuleTestingDelete,
		Importer:      helpers.AccountLevelResourceImporter,

		Schema: map[string]*schema.Schema{
			"module_id": {
				Description: "Identifier of the module to enable testing for",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
			},
			// Input fields for enabling/updating testing
			"org": {
				Description: "Organization identifier",
				Type:        schema.TypeString,
				Required:    true,
			},
			"project": {
				Description: "Project identifier",
				Type:        schema.TypeString,
				Required:    true,
			},
			"provider_connector": {
				Description: "Provider connector for testing purposes",
				Type:        schema.TypeString,
				Required:    true,
			},
			"provisioner_type": {
				Description: "Provisioner type for testing purposes (e.g., terraform, tofu)",
				Type:        schema.TypeString,
				Required:    true,
			},
			"provisioner_version": {
				Description: "Provisioner version for testing purposes",
				Type:        schema.TypeString,
				Required:    true,
			},
			"pipelines": {
				Description: "List of pipeline IDs to create webhooks for triggering test executions",
				Type:        schema.TypeList,
				Required:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"release_pipeline": {
				Description: "Pipeline ID to create webhooks for releases",
				Type:        schema.TypeString,
				Optional:    true,
			},
			// Computed output fields from the module
			"description": {
				Description: "Description of the module",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"name": {
				Description: "Name of the module",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"system": {
				Description: "Provider of the module",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"repository": {
				Description: "For account connectors, the repository where the module is stored",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"repository_branch": {
				Description: "Repository Branch in which the module should be accessed",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"repository_commit": {
				Description: "Repository Commit in which the module should be accessed",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"repository_connector": {
				Description: "Repository Connector is the reference to the connector for the repository",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"repository_path": {
				Description: "Repository Path is the path in which the module resides",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"created": {
				Description: "Timestamp when the module was created",
				Type:        schema.TypeInt,
				Computed:    true,
			},
			"repository_url": {
				Description: "URL where the module is stored",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"synced": {
				Description: "Timestamp when the module was last synced",
				Type:        schema.TypeInt,
				Computed:    true,
			},
			"tags": {
				Description: "Tags associated with the module",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"account": {
				Description: "Account that owns the module",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"git_tag_style": {
				Description: "Git Tag Style",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"module_error": {
				Description: "Error while retrieving the module",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"testing_enabled": {
				Description: "Whether testing is enabled for the module",
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
			},
			"testing_metadata": {
				Description: "Testing metadata for the module",
				Type:        schema.TypeList,
				Optional:    true,
				Computed:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"account": {
							Description: "Account is the internal customer account ID",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"org": {
							Description: "Organization identifier",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"pipelines": {
							Description: "Pipelines where the testing is enabled",
							Type:        schema.TypeList,
							Optional:    true,
							Computed:    true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"project": {
							Description: "Project identifier",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"provider_connector": {
							Description: "Provider connector for testing purposes",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"provisioner_type": {
							Description: "Provisioner type for testing purposes",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"provisioner_version": {
							Description: "Provisioner version for testing purposes",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"release_pipeline": {
							Description: "Release pipeline",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"updated": {
				Description: "Timestamp when the module was last modified",
				Type:        schema.TypeInt,
				Computed:    true,
				Optional:    true,
			},
			"versions": {
				Description: "Versions of the module",
				Type:        schema.TypeList,
				Optional:    true,
				Computed:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
	return resource
}

func resourceInfraModuleTestingRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c, ctx := m.(*internal.Session).GetPlatformClientWithContext(ctx)
	id := d.Id()
	if id == "" {
		return nil
	}
	resp, httpRes, err := c.ModuleRegistryApi.ModuleRegistryListModulesById(
		ctx,
		id,
		c.AccountId,
	)
	if err != nil {
		return helpers.HandleApiError(err, d, httpRes)
	}
	readModule(d, &resp)
	return nil
}

func readModule(d *schema.ResourceData, module *nextgen.ModuleResource) {
	d.SetId(module.Id)
	d.Set("account", module.Account)
	d.Set("created", module.Created)
	d.Set("description", module.Description)
	d.Set("git_tag_style", module.GitTagStyle)
	d.Set("module_error", module.ModuleError)
	d.Set("name", module.Name)
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
	
	// Set testing_metadata as a list with one item
	if module.TestingMetadata != nil {
		testingMetadataList := []interface{}{
			map[string]interface{}{
				"account":             module.TestingMetadata.Account,
				"org":                 module.TestingMetadata.Org,
				"project":             module.TestingMetadata.Project,
				"provider_connector":  module.TestingMetadata.ProviderConnector,
				"provisioner_type":    module.TestingMetadata.ProvisionerType,
				"provisioner_version": module.TestingMetadata.ProvisionerVersion,
				"pipelines":           module.TestingMetadata.Pipelines,
				"release_pipeline":    module.TestingMetadata.ReleasePipeline,
			},
		}
		d.Set("testing_metadata", testingMetadataList)
	}
	
	d.Set("updated", module.Updated)
	d.Set("versions", module.Versions)
}

func resourceInfraModuleTestingCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c, ctx := m.(*internal.Session).GetPlatformClientWithContext(ctx)
	moduleId := d.Get("module_id").(string)

	createModule, err := buildCreateModuleRequestTestingBody(d)
	if err != nil {
		return diag.FromErr(err)
	}
	log.Printf("[DEBUG] Enabling module testing with value %v", createModule)
	httpRes, err := c.ModuleRegistryApi.ModuleRegistryEnableTesting(ctx, createModule, c.AccountId, moduleId)

	if err != nil {
		return parseError(err, httpRes)
	}

	// Set the Terraform resource ID
	d.SetId(moduleId)

	return resourceInfraModuleTestingRead(ctx, d, m)
}

func buildCreateModuleRequestTestingBody(d *schema.ResourceData) (nextgen.EnableTestingRequest, error) {
	module := nextgen.EnableTestingRequest{
		Org:                d.Get("org").(string),
		Project:            d.Get("project").(string),
		ProviderConnector:  d.Get("provider_connector").(string),
		ProvisionerType:    d.Get("provisioner_type").(string),
		ProvisionerVersion: d.Get("provisioner_version").(string),
	}

	if releasePipeline, ok := d.GetOk("release_pipeline"); ok {
		module.ReleasePipeline = releasePipeline.(string)
	}

	if pipelines, ok := d.GetOk("pipelines"); ok {
		pipelinesRaw := pipelines.([]interface{})
		pipelinesStr := make([]string, len(pipelinesRaw))
		for i, v := range pipelinesRaw {
			pipelinesStr[i] = v.(string)
		}
		module.PipelineId = pipelinesStr
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

func resourceInfraModuleTestingDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c, ctx := m.(*internal.Session).GetPlatformClientWithContext(ctx)
	id := d.Id()
	if id == "" {
		return nil
	}
	idInt64, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return diag.Errorf("failed to convert module id to int64: %v", err)
	}
	httpRes, err := c.ModuleRegistryApi.ModuleRegistryDisableTesting(ctx, idInt64, c.AccountId)
	if err != nil {
		return parseError(err, httpRes)
	}
	return nil
}

func resourceInfraModuleTestingUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c, ctx := m.(*internal.Session).GetPlatformClientWithContext(ctx)
	moduleId := d.Get("module_id").(string)

	updateModule, err := buildUpdateModuleRequestTestingBody(d)
	if err != nil {
		return diag.FromErr(err)
	}
	log.Printf("[DEBUG] Updating module testing with value %v", updateModule)
	httpRes, err := c.ModuleRegistryApi.ModuleRegistryUpdateModuleTesting(ctx, updateModule, c.AccountId, moduleId)

	if err != nil {
		return parseError(err, httpRes)
	}

	return resourceInfraModuleTestingRead(ctx, d, m)
}

func buildUpdateModuleRequestTestingBody(d *schema.ResourceData) (nextgen.UpdateTestingRequest, error) {
	module := nextgen.UpdateTestingRequest{
		Org:                d.Get("org").(string),
		Project:            d.Get("project").(string),
		ProviderConnector:  d.Get("provider_connector").(string),
		ProvisionerType:    d.Get("provisioner_type").(string),
		ProvisionerVersion: d.Get("provisioner_version").(string),
	}
	if releasePipeline, ok := d.GetOk("release_pipeline"); ok {
		module.ReleasePipeline = releasePipeline.(string)
	}

	if pipelines, ok := d.GetOk("pipelines"); ok {
		pipelinesRaw := pipelines.([]interface{})
		pipelinesStr := make([]string, len(pipelinesRaw))
		for i, v := range pipelinesRaw {
			pipelinesStr[i] = v.(string)
		}
		module.PipelineId = pipelinesStr
	}
	return module, nil
}
