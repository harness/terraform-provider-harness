package workspace

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/harness/terraform-provider-harness/internal/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceWorkspace() *schema.Resource {
	resource := &schema.Resource{
		Description: "Resource for managing Workspaces",

		ReadContext:   resourceWorkspaceRead,
		DeleteContext: resourceWorkspaceDelete,
		CreateContext: resourceWorkspaceCreate,
		UpdateContext: resourceWorkspaceUpdate,
		Importer:      helpers.ProjectResourceImporter,

		Schema: map[string]*schema.Schema{
			"identifier": {
				Description: "Identifier of the Workspace.",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
			},
			"name": {
				Description: "Name of the Workspace.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"org_id": {
				Description: "Organization identifier of the organization the workspace resides in.",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
			},
			"project_id": {
				Description: "Project identifier of the project the workspace resides in.",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
			},
			"description": {
				Description: "Description of the Workspace.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"provisioner_type": {
				Description: "Provisioner type defines the provisioning tool to use (terraform or opentofu)",
				Type:        schema.TypeString,
				Required:    true,
			},
			"provisioner_version": {
				Description: "Provisioner version defines the provisioner version to use. The latest version of Opentofu should always be supported, Terraform is only supported up to version 1.5.7.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"provider_connector": {
				Description: "Provider connector is the reference to the connector for the infrastructure provider",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"repository": {
				Description: "Repository is the name of the repository to fetch the code from.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"repository_branch": {
				Description:  "Repository branch is the name of the branch to fetch the code from. This cannot be set if repository commit or sha is set.",
				Type:         schema.TypeString,
				Optional:     true,
				ExactlyOneOf: []string{"repository_commit", "repository_branch", "repository_sha"},
			},

			"repository_commit": {
				Description:  "Repository commit is tag to fetch the code from. This cannot be set if repository branch or sha is set.",
				Type:         schema.TypeString,
				Optional:     true,
				ExactlyOneOf: []string{"repository_commit", "repository_branch", "repository_sha"},
			},
			"repository_sha": {
				Description:  "Repository commit is commit SHA to fetch the code from. This cannot be set if repository branch or commit is set.",
				Type:         schema.TypeString,
				Optional:     true,
				ExactlyOneOf: []string{"repository_commit", "repository_branch", "repository_sha"},
			},
			"repository_connector": {
				Description: "Repository connector is the reference to the connector used to fetch the code.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"repository_path": {
				Description: "Repository path is the path in which the code resides.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"cost_estimation_enabled": {
				Description: "Cost estimation enabled determines if cost estimation operations are performed.",
				Type:        schema.TypeBool,
				Required:    true,
			},
			"terraform_variable": {
				Description: "Terraform variables configured on the workspace. Terraform variable keys must be unique within the workspace.",
				Type:        schema.TypeSet,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"key": {
							Description: "Key is the identifier for the variable. Must be unique within the workspace.",
							Type:        schema.TypeString,
							Required:    true,
						},
						"value": {
							Description: "Value is the value of the variable. For string value types this field should contain the value of the variable. For secret value types this should contain a reference to a valid harness secret.",
							Type:        schema.TypeString,
							Required:    true,
						},
						"value_type": {
							Description: "Value type indicates the value type of the variable. Currently we support string and secret.",
							Type:        schema.TypeString,
							Required:    true,
						},
					},
				},
			},
			"environment_variable": {
				Description: "Environment variables configured on the workspace",
				Type:        schema.TypeSet,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"key": {
							Description: "Key is the identifier for the variable. Must be unique within the workspace.",
							Type:        schema.TypeString,
							Required:    true,
						},
						"value": {
							Description: "Value is the value of the variable. For string value types this field should contain the value of the variable. For secret value types this should contain a reference to a valid harness secret.",
							Type:        schema.TypeString,
							Required:    true,
						},
						"value_type": {
							Description: "Value type indicates the value type of the variable. Currently we support string and secret.",
							Type:        schema.TypeString,
							Required:    true,
						},
					},
				},
			},
			"terraform_variable_file": {
				Description: "Terraform variables files configured on the workspace",
				Type:        schema.TypeSet,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"repository": {
							Description: "Repository is the name of the repository to fetch the code from.",
							Type:        schema.TypeString,
							Required:    true,
						},
						"repository_branch": {
							Description: "Repository branch is the name of the branch to fetch the variables from. This cannot be set if repository commit or sha is set",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"repository_commit": {
							Description: "Repository commit is tag to fetch the variables from. This cannot be set if repository branch or sha is set.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"repository_sha": {
							Description: "Repository commit is SHA to fetch the variables from. This cannot be set if repository branch or commit is set.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"repository_connector": {
							Description: "Repository connector is the reference to the connector used to fetch the variables.",
							Type:        schema.TypeString,
							Required:    true,
						},
						"repository_path": {
							Description: "Repository path is the path in which the variables reside.",
							Type:        schema.TypeString,
							Optional:    true,
						},
					},
				},
			},
			"default_pipelines": {
				Description: "Default pipelines associated with this workspace",
				Type:        schema.TypeMap,
				Optional:    true,
			},
			"variable_sets": {
				Description: "Variable sets to use.",
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"connector": {
				Description: "Provider connectors configured on the Workspace. Only one connector of a type is supported",
				Type:        schema.TypeSet,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"connector_ref": {
							Description: "Connector Ref is the reference to the connector",
							Type:        schema.TypeString,
							Required:    true,
						},
						"type": {
							Description:  "Type is the connector type of the connector. Supported types: aws, azure, gcp",
							Type:         schema.TypeString,
							Required:     true,
							ValidateFunc: validation.StringInSlice([]string{"aws", "azure", "gcp"}, false),
						},
					},
				},
			},
		},
	}
	resource.Schema["tags"] = helpers.GetTagsSchema(helpers.SchemaFlagTypes.Optional)
	helpers.SetProjectLevelResourceSchema(resource.Schema)
	return resource
}

func resourceWorkspaceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)
	id := d.Id()
	if id == "" {
		d.MarkNewResource()
	}
	resp, httpResp, err := c.WorkspaceApi.WorkspacesShowWorkspace(
		ctx,
		d.Get("org_id").(string),
		d.Get("project_id").(string),
		d.Get("identifier").(string),
		c.AccountId,
	)
	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	readWorkspace(d, &resp)
	return nil
}

func resourceWorkspaceDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)
	id := d.Id()
	if id == "" {
		return nil
	}

	httpResp, err := c.WorkspaceApi.WorkspacesDestroyWorkspace(
		ctx,
		d.Get("org_id").(string),
		d.Get("project_id").(string),
		d.Get("identifier").(string),
		c.AccountId,
	)
	if err != nil {
		return parseError(err, httpResp)
	}

	return nil
}

func resourceWorkspaceCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)
	id := d.Id()
	if id == "" {
		d.MarkNewResource()
	}

	createWorkspace, err := buildCreateWorkspace(d)
	if err != nil {
		return diag.Errorf(err.Error())
	}

	_, httpResp, err := c.WorkspaceApi.WorkspacesCreateWorkspace(
		ctx,
		createWorkspace,
		c.AccountId,
		d.Get("org_id").(string),
		d.Get("project_id").(string),
	)
	if err != nil {
		return parseError(err, httpResp)
	}

	resourceWorkspaceRead(ctx, d, meta)
	return nil
}

func resourceWorkspaceUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)
	id := d.Id()
	if id == "" {
		d.MarkNewResource()
	}

	updateWorkspace, err := buildUpdateWorkspace(d)
	if err != nil {
		return diag.Errorf(err.Error())
	}

	_, httpResp, err := c.WorkspaceApi.WorkspacesUpdateWorkspace(
		ctx,
		updateWorkspace,
		c.AccountId,
		d.Get("org_id").(string),
		d.Get("project_id").(string),
		d.Get("identifier").(string),
	)
	if err != nil {
		return parseError(err, httpResp)
	}

	resourceWorkspaceRead(ctx, d, meta)
	return nil
}

func readWorkspace(d *schema.ResourceData, ws *nextgen.IacmShowWorkspaceResponseBody) {
	d.SetId(ws.Identifier)
	d.Set("identifier", ws.Identifier)
	d.Set("org_id", ws.Org)
	d.Set("project_id", ws.Project)
	d.Set("name", ws.Name)
	d.Set("description", ws.Description)
	d.Set("provisioner_type", ws.Provisioner)
	d.Set("provisioner_version", ws.ProvisionerVersion)
	d.Set("provider_connector", ws.ProviderConnector)
	d.Set("repository", ws.Repository)
	d.Set("repository_branch", ws.RepositoryBranch)
	d.Set("repository_commit", ws.RepositoryCommit)
	d.Set("repository_sha", ws.RepositorySha)
	d.Set("repository_path", ws.RepositoryPath)
	d.Set("repository_connector", ws.RepositoryConnector)
	d.Set("cost_estimation_enabled", ws.CostEstimationEnabled)
	d.Set("variable_sets", ws.VariableSets)
	var environmentVariables []interface{}
	for _, v := range ws.EnvironmentVariables {
		environmentVariables = append(environmentVariables, map[string]string{
			"key":        v.Key,
			"value":      v.Value,
			"value_type": v.ValueType,
		})
	}
	d.Set("environment_variable", environmentVariables)
	var terraformVariables []interface{}
	for _, v := range ws.TerraformVariables {
		terraformVariables = append(terraformVariables, map[string]string{
			"key":        v.Key,
			"value":      v.Value,
			"value_type": v.ValueType,
		})
	}
	d.Set("terraform_variable", terraformVariables)
	var terraformVariableFiles []interface{}
	for _, v := range ws.TerraformVariableFiles {
		terraformVariableFiles = append(terraformVariableFiles, map[string]string{
			"repository":           v.Repository,
			"repository_branch":    v.RepositoryBranch,
			"repository_commit":    v.RepositoryCommit,
			"repository_sha":       v.RepositorySha,
			"repository_path":      v.RepositoryPath,
			"repository_connector": v.RepositoryConnector,
		})
	}
	d.Set("terraform_variable_file", terraformVariableFiles)
	defaultPipelines := map[string]string{}
	for k, v := range ws.DefaultPipelines {
		if v.WorkspacePipeline != "" {
			defaultPipelines[k] = v.WorkspacePipeline
		}
	}
	d.Set("default_pipelines", defaultPipelines)

	var providerConnectors []interface{}
	for _, c := range ws.ProviderConnectors {
		providerConnectors = append(providerConnectors, map[string]string{
			"connector_ref": c.ConnectorRef,
			"type":          c.Type_,
		})
	}
	d.Set("connector", providerConnectors)
	d.Set("tags", helpers.FlattenTags(ws.Tags))
}

func buildUpdateWorkspace(d *schema.ResourceData) (nextgen.IacmUpdateWorkspaceRequestBody, error) {
	ws := nextgen.IacmUpdateWorkspaceRequestBody{
		Name:                  d.Get("name").(string),
		Provisioner:           d.Get("provisioner_type").(string),
		ProvisionerVersion:    d.Get("provisioner_version").(string),
		Repository:            d.Get("repository").(string),
		RepositoryPath:        d.Get("repository_path").(string),
		RepositoryConnector:   d.Get("repository_connector").(string),
		CostEstimationEnabled: d.Get("cost_estimation_enabled").(bool),
	}

	if providerConnector, ok := d.GetOk("provider_connector"); ok {
		ws.ProviderConnector = providerConnector.(string)
	}

	if desc, ok := d.GetOk("description"); ok {
		ws.Description = desc.(string)
	}

	if desc, ok := d.GetOk("repository_branch"); ok {
		ws.RepositoryBranch = desc.(string)
	}

	if desc, ok := d.GetOk("repository_commit"); ok {
		ws.RepositoryCommit = desc.(string)
	}

	if desc, ok := d.GetOk("repository_sha"); ok {
		ws.RepositorySha = desc.(string)
	}

	if desc, ok := d.GetOk("repository_path"); ok {
		ws.RepositoryPath = desc.(string)
	}

	environmentVariables, err := buildVariables(d, "environment_variable")
	if err != nil {
		return nextgen.IacmUpdateWorkspaceRequestBody{}, err
	}
	ws.EnvironmentVariables = environmentVariables

	terraformVariables, err := buildVariables(d, "terraform_variable")
	if err != nil {
		return nextgen.IacmUpdateWorkspaceRequestBody{}, err
	}
	ws.TerraformVariables = terraformVariables

	ws.TerraformVariableFiles = buildTerraformVariableFiles(d)

	defaultPipelines, err := buildDefaultPipelines(d)
	if err != nil {
		return nextgen.IacmUpdateWorkspaceRequestBody{}, err
	}
	ws.DefaultPipelines = defaultPipelines

	if varSets := d.Get("variable_sets").([]interface{}); len(varSets) > 0 {
		ws.VariableSets = utils.InterfaceSliceToStringSlice(varSets)
	}

	providerConnectors, err := buildProviderConnectors(d)
	if err != nil {
		return nextgen.IacmUpdateWorkspaceRequestBody{}, err
	}
	ws.ProviderConnectors = providerConnectors

	if attr := d.Get("tags").(*schema.Set).List(); len(attr) > 0 {
		ws.Tags = helpers.ExpandTags(attr)
	}

	return ws, nil
}

func buildCreateWorkspace(d *schema.ResourceData) (nextgen.IacmCreateWorkspaceRequestBody, error) {
	ws := nextgen.IacmCreateWorkspaceRequestBody{
		Identifier:            d.Get("identifier").(string),
		Name:                  d.Get("name").(string),
		Provisioner:           d.Get("provisioner_type").(string),
		ProvisionerVersion:    d.Get("provisioner_version").(string),
		Repository:            d.Get("repository").(string),
		RepositoryPath:        d.Get("repository_path").(string),
		RepositoryConnector:   d.Get("repository_connector").(string),
		CostEstimationEnabled: d.Get("cost_estimation_enabled").(bool),
	}

	if providerConnector, ok := d.GetOk("provider_connector"); ok {
		ws.ProviderConnector = providerConnector.(string)
	}

	if desc, ok := d.GetOk("description"); ok {
		ws.Description = desc.(string)
	}

	if desc, ok := d.GetOk("repository_branch"); ok {
		ws.RepositoryBranch = desc.(string)
	}

	if desc, ok := d.GetOk("repository_commit"); ok {
		ws.RepositoryCommit = desc.(string)
	}

	if desc, ok := d.GetOk("repository_sha"); ok {
		ws.RepositorySha = desc.(string)
	}

	if desc, ok := d.GetOk("repository_path"); ok {
		ws.RepositoryPath = desc.(string)
	}

	environmentVariables, err := buildVariables(d, "environment_variable")
	if err != nil {
		return nextgen.IacmCreateWorkspaceRequestBody{}, err
	}
	ws.EnvironmentVariables = environmentVariables

	terraformVariables, err := buildVariables(d, "terraform_variable")
	if err != nil {
		return nextgen.IacmCreateWorkspaceRequestBody{}, err
	}
	ws.TerraformVariables = terraformVariables

	ws.TerraformVariableFiles = buildTerraformVariableFiles(d)

	defaultPipelines, err := buildDefaultPipelines(d)
	if err != nil {
		return nextgen.IacmCreateWorkspaceRequestBody{}, err
	}
	ws.DefaultPipelines = defaultPipelines

	if varSets := d.Get("variable_sets").([]interface{}); len(varSets) > 0 {
		ws.VariableSets = utils.InterfaceSliceToStringSlice(varSets)
	}

	providerConnectors, err := buildProviderConnectors(d)
	if err != nil {
		return nextgen.IacmCreateWorkspaceRequestBody{}, err
	}
	ws.ProviderConnectors = providerConnectors

	if attr := d.Get("tags").(*schema.Set).List(); len(attr) > 0 {
		ws.Tags = helpers.ExpandTags(attr)
	}

	return ws, nil
}

func buildTerraformVariableFiles(d *schema.ResourceData) []nextgen.IacmWorkspaceTerraformVariableFiles {
	terraformVariableFiles := []nextgen.IacmWorkspaceTerraformVariableFiles{}
	if _, ok := d.GetOk("terraform_variable_file"); ok {
		for _, v := range d.Get("terraform_variable_file").(*schema.Set).List() {
			if tfv, ok := v.(map[string]interface{}); ok {
				terraformVariableFiles = append(terraformVariableFiles, nextgen.IacmWorkspaceTerraformVariableFiles{
					Repository:          tfv["repository"].(string),
					RepositoryConnector: tfv["repository_connector"].(string),
					RepositoryBranch:    tfv["repository_branch"].(string),
					RepositoryCommit:    tfv["repository_commit"].(string),
					RepositorySha:       tfv["repository_sha"].(string),
					RepositoryPath:      tfv["repository_path"].(string),
				})
			}
		}
	}
	return terraformVariableFiles
}

func buildVariables(d *schema.ResourceData, attribute string) (map[string]nextgen.IacmVariable, error) {
	variables := map[string]nextgen.IacmVariable{}
	if _, ok := d.GetOk(attribute); ok {
		for _, v := range d.Get(attribute).(*schema.Set).List() {
			if ev, ok := v.(map[string]interface{}); ok {
				if _, ok = variables[ev["key"].(string)]; ok {
					return variables, fmt.Errorf("%s keys must be unique", attribute)
				}
				variables[ev["key"].(string)] = nextgen.IacmVariable{
					Key:       ev["key"].(string),
					Value:     ev["value"].(string),
					ValueType: ev["value_type"].(string),
				}
			}
		}
	}
	return variables, nil
}

func buildDefaultPipelines(d *schema.ResourceData) (map[string]nextgen.IacmDefaultPipelineOverride, error) {
	defaultPipelines := map[string]nextgen.IacmDefaultPipelineOverride{}
	if pipelines, ok := d.GetOk("default_pipelines"); ok {
		if pipelinesMap, ok := pipelines.(map[string]interface{}); ok {
			for operation, pipeline := range pipelinesMap {
				if pipelineStr, ok := pipeline.(string); ok {
					defaultPipelines[operation] = nextgen.IacmDefaultPipelineOverride{WorkspacePipeline: pipelineStr}
				}
			}
		}
	}
	return defaultPipelines, nil
}

func buildProviderConnectors(d *schema.ResourceData) ([]nextgen.VariableSetConnector, error) {
	connectors := []nextgen.VariableSetConnector{}
	seen := make(map[string]struct{}) // Map to track seen combinations

	if _, ok := d.GetOk("connector"); ok {
		for _, v := range d.Get("connector").(*schema.Set).List() {
			if con, ok := v.(map[string]interface{}); ok {
				connectorType := con["type"].(string)

				// Check if the type already exists
				if _, exists := seen[connectorType]; exists {
					return connectors, fmt.Errorf("%s types must be unique for connectors", connectorType)
				}

				// Mark the type as seen
				seen[connectorType] = struct{}{}

				connectors = append(connectors, nextgen.VariableSetConnector{
					ConnectorRef: con["connector_ref"].(string),
					Type_:        connectorType,
				})
			}
		}
	}
	return connectors, nil
}

// iacm errors are in a different format from other harness services
// this function parses iacm errors and attempts to return them in way
// that is consistent with the provider.
func parseError(err error, httpResp *http.Response) diag.Diagnostics {
	// copied from helpers/errors.go
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
