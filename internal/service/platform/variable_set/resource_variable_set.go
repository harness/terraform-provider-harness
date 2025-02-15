package variable_set

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceVariableSet() *schema.Resource {
	resource := &schema.Resource{
		Description: "Resource for managing Variable Sets",

		ReadContext:   resourceVarsetRead,
		DeleteContext: resourceVarsetDelete,
		CreateContext: resourceVarsetCreate,
		UpdateContext: resourceVarsetUpdate,
		Importer:      helpers.MultiLevelResourceImporter,

		Schema: map[string]*schema.Schema{
			"identifier": {
				Description: "Identifier of the Variable Set.",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
			},
			"org_id": {
				Description: "Organization identifier of the organization the Variable Set resides in.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"project_id": {
				Description: "Project identifier of the project the Variable Set resides in.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"name": {
				Description: "Name of the Variable Set.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"description": {
				Description: "Description of the Variable Set.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"connector": {
				Description: "Provider connectors configured on the Variable Set. Only one connector of a type is supported",
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
			"environment_variable": {
				Description: "Environment variables configured on the Variable Set",
				Type:        schema.TypeSet,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"key": {
							Description: "Key is the identifier for the variable. Must be unique within the Variable Set.",
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
			"terraform_variable": {
				Description: "Terraform variables configured on the Variable Set. Terraform variable keys must be unique within the Variable Set.",
				Type:        schema.TypeSet,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"key": {
							Description: "Key is the identifier for the variable. Must be unique within the Variable Set.",
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
				Description: "Terraform variables files configured on the Variable Set",
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
		},
	}

	helpers.SetMultiLevelResourceSchema(resource.Schema)

	return resource
}

func resourceVarsetRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	identifier := d.Get("identifier").(string)
	scope := getScope(d)
	var resp nextgen.VariableSetsGetVariableSetAccountLevelResponseBody
	var httpResp *http.Response
	var err error

	switch scope.scope {
	case Project:
		resp, httpResp, err = c.VariableSetsApi.VariableSetsGetVariableSetProjLevel(
			ctx,
			scope.org,
			scope.project,
			identifier,
			c.AccountId,
		)
	case Org:
		resp, httpResp, err = c.VariableSetsApi.VariableSetsGetVariableSetOrgLevel(
			ctx,
			scope.org,
			identifier,
			c.AccountId,
		)
	default:
		resp, httpResp, err = c.VariableSetsApi.VariableSetsGetVariableSetAccountLevel(
			ctx,
			identifier,
			c.AccountId,
		)
	}

	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	readVariableSet(d, &resp)
	return nil
}

func resourceVarsetDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	identifier := d.Get("identifier").(string)
	scope := getScope(d)
	var httpResp *http.Response
	var err error

	switch scope.scope {
	case Project:
		httpResp, err = c.VariableSetsApi.VariableSetsDeleteVariableSetProjLevel(
			ctx,
			scope.org,
			scope.project,
			identifier,
			c.AccountId,
		)
	case Org:
		httpResp, err = c.VariableSetsApi.VariableSetsDeleteVariableSetOrgLevel(
			ctx,
			scope.org,
			identifier,
			c.AccountId,
		)
	default:
		httpResp, err = c.VariableSetsApi.VariableSetsDeleteVariableSetAccountLevel(
			ctx,
			identifier,
			c.AccountId,
		)
	}

	if err != nil {
		return parseError(err, httpResp)
	}

	return nil
}

func resourceVarsetCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)
	id := d.Id()
	if id == "" {
		d.MarkNewResource()
	}

	scope := getScope(d)
	var httpResp *http.Response
	var err error

	switch scope.scope {
	case Project:
		req, err2 := buildCreateVarSetReqProjLevel(d)
		if err2 != nil {
			return diag.Errorf(err2.Error())
		}
		_, httpResp, err = c.VariableSetsApi.VariableSetsCreateVariableSetProjLevel(
			ctx,
			*req,
			c.AccountId,
			scope.org,
			scope.project,
		)
	case Org:
		req, err2 := buildCreateVarSetReqOrgLevel(d)
		if err2 != nil {
			return diag.Errorf(err2.Error())
		}
		_, httpResp, err = c.VariableSetsApi.VariableSetsCreateVariableSetOrgLevel(
			ctx,
			*req,
			c.AccountId,
			scope.org,
		)
	default:
		req, err2 := buildCreateVarSetReqAccountLevel(d)
		if err2 != nil {
			return diag.Errorf(err2.Error())
		}
		_, httpResp, err = c.VariableSetsApi.VariableSetsCreateVariableSetAccountLevel(
			ctx,
			*req,
			c.AccountId,
		)
	}

	if err != nil {
		return parseError(err, httpResp)
	}

	resourceVarsetRead(ctx, d, meta)
	return nil
}

func resourceVarsetUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	identifier := d.Get("identifier").(string)
	scope := getScope(d)
	var httpResp *http.Response
	var err error

	switch scope.scope {
	case Project:
		req, err2 := buildUpdateVarSetReqProjLevel(d)
		if err2 != nil {
			return diag.Errorf(err2.Error())
		}
		_, httpResp, err = c.VariableSetsApi.VariableSetsUpdateVariableSetProjLevel(
			ctx,
			*req,
			c.AccountId,
			scope.org,
			scope.project,
			identifier,
		)
	case Org:
		req, err2 := buildUpdateVarSetReqOrgLevel(d)
		if err2 != nil {
			return diag.Errorf(err2.Error())
		}
		_, httpResp, err = c.VariableSetsApi.VariableSetsUpdateVariableSetOrgLevel(
			ctx,
			*req,
			c.AccountId,
			scope.org,
			identifier,
		)
	default:
		req, err2 := buildUpdateVarSetReqAccountLevel(d)
		if err2 != nil {
			return diag.Errorf(err2.Error())
		}
		_, httpResp, err = c.VariableSetsApi.VariableSetsUpdateVariableSetAccountLevel(
			ctx,
			*req,
			c.AccountId,
			identifier,
		)
	}

	if err != nil {
		return parseError(err, httpResp)
	}

	resourceVarsetRead(ctx, d, meta)
	return nil
}

func readVariableSet(d *schema.ResourceData, vs *nextgen.VariableSetsGetVariableSetAccountLevelResponseBody) {
	d.SetId(vs.Identifier)
	d.Set("identifier", vs.Identifier)
	d.Set("org_id", vs.Org)
	d.Set("project_id", vs.Project)
	d.Set("name", vs.Name)
	d.Set("description", vs.Description)

	var environmentVariables []interface{}
	for _, v := range vs.EnvironmentVariables {
		environmentVariables = append(environmentVariables, map[string]string{
			"key":        v.Key,
			"value":      v.Value,
			"value_type": v.ValueType,
		})
	}
	d.Set("environment_variable", environmentVariables)

	var terraformVariables []interface{}
	for _, v := range vs.TerraformVariables {
		terraformVariables = append(terraformVariables, map[string]string{
			"key":        v.Key,
			"value":      v.Value,
			"value_type": v.ValueType,
		})
	}
	d.Set("terraform_variable", terraformVariables)

	var terraformVariableFiles []interface{}
	for _, v := range vs.TerraformVariableFiles {
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

	var connectors []interface{}
	for _, c := range vs.Connectors {
		connectors = append(connectors, map[string]string{
			"connector_ref": c.ConnectorRef,
			"type":          c.Type_,
		})
	}
	d.Set("connector", connectors)
}

func buildCreateVarSetReqProjLevel(d *schema.ResourceData) (*nextgen.CreateVariableSetRequestProjScope, error) {
	vs := &nextgen.CreateVariableSetRequestProjScope{
		Identifier: d.Get("identifier").(string),
		Name:       d.Get("name").(string),
	}

	if desc, ok := d.GetOk("description"); ok {
		vs.Description = desc.(string)
	}

	connectors, err := buildConnectors(d)
	if err != nil {
		return nil, err
	}
	vs.Connectors = connectors

	environmentVariables, err := buildVariables(d, "environment_variable")
	if err != nil {
		return nil, err
	}
	vs.EnvironmentVariables = environmentVariables

	terraformVariables, err := buildVariables(d, "terraform_variable")
	if err != nil {
		return nil, err
	}
	vs.TerraformVariables = terraformVariables

	vs.TerraformVariableFiles = buildTerraformVariableFiles(d)

	return vs, nil
}

func buildCreateVarSetReqOrgLevel(d *schema.ResourceData) (*nextgen.CreateVariableSetRequestOrgScope, error) {
	vs, err := buildCreateVarSetReqProjLevel(d)
	if err != nil {
		return nil, err
	}

	return &nextgen.CreateVariableSetRequestOrgScope{
		Identifier:             vs.Identifier,
		Name:                   vs.Name,
		Description:            vs.Description,
		EnvironmentVariables:   vs.EnvironmentVariables,
		TerraformVariables:     vs.TerraformVariables,
		TerraformVariableFiles: vs.TerraformVariableFiles,
		Connectors:             vs.Connectors,
	}, nil
}

func buildCreateVarSetReqAccountLevel(d *schema.ResourceData) (*nextgen.CreateVariableSetRequestAccScope, error) {
	vs, err := buildCreateVarSetReqProjLevel(d)
	if err != nil {
		return nil, err
	}

	return &nextgen.CreateVariableSetRequestAccScope{
		Identifier:             vs.Identifier,
		Name:                   vs.Name,
		Description:            vs.Description,
		EnvironmentVariables:   vs.EnvironmentVariables,
		TerraformVariables:     vs.TerraformVariables,
		TerraformVariableFiles: vs.TerraformVariableFiles,
		Connectors:             vs.Connectors,
	}, nil
}

func buildUpdateVarSetReqProjLevel(d *schema.ResourceData) (*nextgen.UpdateVariableSetRequestProjScope, error) {
	vs, err := buildCreateVarSetReqProjLevel(d)
	if err != nil {
		return nil, err
	}

	return &nextgen.UpdateVariableSetRequestProjScope{
		Name:                   vs.Name,
		Description:            vs.Description,
		EnvironmentVariables:   vs.EnvironmentVariables,
		TerraformVariables:     vs.TerraformVariables,
		TerraformVariableFiles: vs.TerraformVariableFiles,
		Connectors:             vs.Connectors,
	}, nil
}

func buildUpdateVarSetReqOrgLevel(d *schema.ResourceData) (*nextgen.UpdateVariableSetRequestOrgScope, error) {
	vs, err := buildCreateVarSetReqProjLevel(d)
	if err != nil {
		return nil, err
	}

	return &nextgen.UpdateVariableSetRequestOrgScope{
		Name:                   vs.Name,
		Description:            vs.Description,
		EnvironmentVariables:   vs.EnvironmentVariables,
		TerraformVariables:     vs.TerraformVariables,
		TerraformVariableFiles: vs.TerraformVariableFiles,
		Connectors:             vs.Connectors,
	}, nil
}

func buildUpdateVarSetReqAccountLevel(d *schema.ResourceData) (*nextgen.UpdateVariableSetRequestAccountScope, error) {
	vs, err := buildCreateVarSetReqProjLevel(d)
	if err != nil {
		return nil, err
	}

	return &nextgen.UpdateVariableSetRequestAccountScope{
		Name:                   vs.Name,
		Description:            vs.Description,
		EnvironmentVariables:   vs.EnvironmentVariables,
		TerraformVariables:     vs.TerraformVariables,
		TerraformVariableFiles: vs.TerraformVariableFiles,
		Connectors:             vs.Connectors,
	}, nil
}

func buildTerraformVariableFiles(d *schema.ResourceData) []nextgen.VariableSetVarFile {
	terraformVariableFiles := []nextgen.VariableSetVarFile{}
	if _, ok := d.GetOk("terraform_variable_file"); ok {
		for _, v := range d.Get("terraform_variable_file").(*schema.Set).List() {
			if tfv, ok := v.(map[string]interface{}); ok {
				terraformVariableFiles = append(terraformVariableFiles, nextgen.VariableSetVarFile{
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

func buildConnectors(d *schema.ResourceData) ([]nextgen.VariableSetConnector, error) {
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

func buildVariables(d *schema.ResourceData, attribute string) (map[string]nextgen.VariableSetVar, error) {
	variables := map[string]nextgen.VariableSetVar{}
	if _, ok := d.GetOk(attribute); ok {
		for _, v := range d.Get(attribute).(*schema.Set).List() {
			if ev, ok := v.(map[string]interface{}); ok {
				if _, ok = variables[ev["key"].(string)]; ok {
					return variables, fmt.Errorf("%s keys must be unique", attribute)
				}
				variables[ev["key"].(string)] = nextgen.VariableSetVar{
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

type Scope struct {
	org     string
	project string
	scope   ScopeLevel
}

type ScopeLevel string

const (
	Account ScopeLevel = "account"
	Org     ScopeLevel = "org"
	Project ScopeLevel = "project"
)

func getScope(d *schema.ResourceData) *Scope {
	org := ""
	project := ""

	if attr, ok := d.GetOk("org_id"); ok {
		org = (attr.(string))
	}

	if attr, ok := d.GetOk("project_id"); ok {
		project = (attr.(string))
	}

	var scope ScopeLevel
	if org == "" {
		scope = Account
	} else if project == "" {
		scope = Org
	} else {
		scope = Project
	}

	return &Scope{
		org:     org,
		project: project,
		scope:   scope,
	}
}
