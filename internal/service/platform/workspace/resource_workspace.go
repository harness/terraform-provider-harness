package workspace

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

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
				Description: "Provisioner type defines the provisioning tool to use. Currently only terraform is supported.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"provisioner_version": {
				Description: "Provisioner version defines the tool version to use. Currently we support versions of terraform less than or equal 1.5.6",
				Type:        schema.TypeString,
				Required:    true,
			},
			"provider_connector": {
				Description: "Provider connector is the reference to the connector for the infrastructure provider",
				Type:        schema.TypeString,
				Required:    true,
			},
			"repository": {
				Description: "Repository is the name of the repository to fetch the code from.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"repository_branch": {
				Description:  "Repository branch is the name of the branch to fetch the code from. This cannot be set if repository commit is set.",
				Type:         schema.TypeString,
				Optional:     true,
				ExactlyOneOf: []string{"repository_commit", "repository_branch"},
			},

			"repository_commit": {
				Description:  "Repository commit is commit or tag to fetch the code from. This cannot be set if repository branch is set.",
				Type:         schema.TypeString,
				Optional:     true,
				ExactlyOneOf: []string{"repository_commit", "repository_branch"},
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
							Description: "Repository branch is the name of the branch to fetch the variables from. This cannot be set if repository commit is set",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"repository_commit": {
							Description: "Repository commit is commit or tag to fetch the variables from. This cannot be set if repository branch is set.",
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
		},
	}

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
	d.Set("repository_path", ws.RepositoryPath)
	d.Set("repository_connector", ws.RepositoryConnector)
	d.Set("cost_estimation_enabled", ws.CostEstimationEnabled)
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
			"repository_path":      v.RepositoryPath,
			"repository_connector": v.RepositoryConnector,
		})
	}
	d.Set("terraform_variable_file", terraformVariableFiles)
}

func buildUpdateWorkspace(d *schema.ResourceData) (nextgen.IacmUpdateWorkspaceRequestBody, error) {
	ws := nextgen.IacmUpdateWorkspaceRequestBody{
		Name:                  d.Get("name").(string),
		Provisioner:           d.Get("provisioner_type").(string),
		ProvisionerVersion:    d.Get("provisioner_version").(string),
		Repository:            d.Get("repository").(string),
		RepositoryPath:        d.Get("repository_path").(string),
		RepositoryConnector:   d.Get("repository_connector").(string),
		ProviderConnector:     d.Get("provider_connector").(string),
		CostEstimationEnabled: d.Get("cost_estimation_enabled").(bool),
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
		ProviderConnector:     d.Get("provider_connector").(string),
		CostEstimationEnabled: d.Get("cost_estimation_enabled").(bool),
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
