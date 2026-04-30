package ansible_playbook

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceAnsiblePlaybook() *schema.Resource {
	resource := &schema.Resource{
		Description: "Resource for managing Harness IaCM Ansible Playbooks.",

		ReadContext:   resourceAnsiblePlaybookRead,
		DeleteContext: resourceAnsiblePlaybookDelete,
		CreateContext: resourceAnsiblePlaybookCreate,
		UpdateContext: resourceAnsiblePlaybookUpdate,
		Importer:      helpers.ProjectResourceImporter,

		Schema: playbookSchema(false),
	}
	resource.Schema["tags"] = helpers.GetTagsSchema(helpers.SchemaFlagTypes.Optional)
	helpers.SetProjectLevelResourceSchema(resource.Schema)
	return resource
}

func playbookSchema(dataSource bool) map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"repository": {
			Description: "Repository name for the playbook.",
			Type:        schema.TypeString,
			Optional:    !dataSource,
			Computed:    dataSource,
		},
		"repository_branch": {
			Description: "Repository branch for the playbook.",
			Type:        schema.TypeString,
			Optional:    !dataSource,
			Computed:    dataSource,
		},
		"repository_commit": {
			Description: "Repository commit or tag for the playbook.",
			Type:        schema.TypeString,
			Optional:    !dataSource,
			Computed:    dataSource,
		},
		"repository_connector": {
			Description: "Repository connector reference for the playbook.",
			Type:        schema.TypeString,
			Optional:    !dataSource,
			Computed:    dataSource,
		},
		"repository_path": {
			Description: "Path within the repository where the playbook resides.",
			Type:        schema.TypeString,
			Required:    !dataSource,
			Computed:    dataSource,
		},
		"ansible_galaxy": {
			Description: "Install Ansible Galaxy dependencies.",
			Type:        schema.TypeBool,
			Optional:    !dataSource,
			Computed:    dataSource,
		},
		"ansible_galaxy_requirements_file": {
			Description: "Path to the Ansible Galaxy requirements file.",
			Type:        schema.TypeString,
			Optional:    !dataSource,
			Computed:    dataSource,
		},
		"vars": {
			Description: "Variables configured on the playbook.",
			Type:        schema.TypeSet,
			Optional:    !dataSource,
			Computed:    dataSource,
			Elem:        ansibleVariableSchema(dataSource),
		},
		"env_vars": {
			Description: "Environment variables configured on the playbook.",
			Type:        schema.TypeSet,
			Optional:    !dataSource,
			Computed:    dataSource,
			Elem:        ansibleVariableSchema(dataSource),
		},
	}
}

func ansibleVariableSchema(dataSource bool) *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"key": {
				Description: "Key is the identifier for the variable.",
				Type:        schema.TypeString,
				Required:    !dataSource,
				Computed:    dataSource,
			},
			"value": {
				Description: "Value of the variable. For secret value types this must be a Harness secret reference.",
				Type:        schema.TypeString,
				Required:    !dataSource,
				Computed:    dataSource,
			},
			"value_type": {
				Description: "Value type. One of: string, secret.",
				Type:        schema.TypeString,
				Required:    !dataSource,
				Computed:    dataSource,
			},
			"file_name": {
				Description: "Filename to store the value in (used for file-backed variables).",
				Type:        schema.TypeString,
				Optional:    !dataSource,
				Computed:    dataSource,
			},
		},
	}
}

func resourceAnsiblePlaybookRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	orgID := d.Get("org_id").(string)
	projectID := d.Get("project_id").(string)
	identifier := d.Get("identifier").(string)

	resp, httpResp, err := c.AnsibleApi.AnsibleShowPlaybook(ctx, orgID, projectID, identifier, c.AccountId)
	if err != nil {
		return helpers.HandleReadApiError(err, d, httpResp)
	}

	return readPlaybook(d, &resp)
}

func resourceAnsiblePlaybookDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)
	if d.Id() == "" {
		return nil
	}

	httpResp, err := c.AnsibleApi.AnsibleDeletePlaybook(
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

func resourceAnsiblePlaybookCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	body, err := buildCreatePlaybook(d)
	if err != nil {
		return diag.Errorf("%s", err.Error())
	}

	_, httpResp, err := c.AnsibleApi.AnsibleCreatePlaybook(
		ctx,
		body,
		c.AccountId,
		d.Get("org_id").(string),
		d.Get("project_id").(string),
	)
	if err != nil {
		return parseError(err, httpResp)
	}

	d.SetId(body.Identifier)
	return resourceAnsiblePlaybookRead(ctx, d, meta)
}

func resourceAnsiblePlaybookUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	body, err := buildUpdatePlaybook(d)
	if err != nil {
		return diag.Errorf("%s", err.Error())
	}

	httpResp, err := c.AnsibleApi.AnsibleUpdatePlaybook(
		ctx,
		body,
		c.AccountId,
		d.Get("org_id").(string),
		d.Get("project_id").(string),
		d.Get("identifier").(string),
	)
	if err != nil {
		return parseError(err, httpResp)
	}

	return resourceAnsiblePlaybookRead(ctx, d, meta)
}

func buildCreatePlaybook(d *schema.ResourceData) (nextgen.CreatePlaybookRequest, error) {
	req := nextgen.CreatePlaybookRequest{
		Identifier:                    d.Get("identifier").(string),
		Name:                          d.Get("name").(string),
		Repository:                    d.Get("repository").(string),
		RepositoryBranch:              d.Get("repository_branch").(string),
		RepositoryCommit:              d.Get("repository_commit").(string),
		RepositoryConnector:           d.Get("repository_connector").(string),
		RepositoryPath:                d.Get("repository_path").(string),
		AnsibleGalaxy:                 d.Get("ansible_galaxy").(bool),
		AnsibleGalaxyRequirementsFile: d.Get("ansible_galaxy_requirements_file").(string),
		Vars:                          variablesFromSet(d.Get("vars")),
		EnvVars:                       variablesFromSet(d.Get("env_vars")),
	}
	if tagAttr, ok := d.GetOk("tags"); ok {
		tags, err := marshalTags(tagAttr)
		if err != nil {
			return req, err
		}
		req.Tags = tags
	}
	return req, nil
}

func buildUpdatePlaybook(d *schema.ResourceData) (nextgen.UpdatePlaybookRequest, error) {
	req := nextgen.UpdatePlaybookRequest{
		Name:                          d.Get("name").(string),
		Repository:                    d.Get("repository").(string),
		RepositoryBranch:              d.Get("repository_branch").(string),
		RepositoryCommit:              d.Get("repository_commit").(string),
		RepositoryConnector:           d.Get("repository_connector").(string),
		RepositoryPath:                d.Get("repository_path").(string),
		AnsibleGalaxy:                 d.Get("ansible_galaxy").(bool),
		AnsibleGalaxyRequirementsFile: d.Get("ansible_galaxy_requirements_file").(string),
		Vars:                          variablesFromSet(d.Get("vars")),
		EnvVars:                       variablesFromSet(d.Get("env_vars")),
	}
	if tagAttr, ok := d.GetOk("tags"); ok {
		tags, err := marshalTags(tagAttr)
		if err != nil {
			return req, err
		}
		req.Tags = tags
	}
	return req, nil
}

func variablesFromSet(v interface{}) map[string]nextgen.AnsibleVariable {
	if v == nil {
		return nil
	}
	set, ok := v.(*schema.Set)
	if !ok {
		return nil
	}
	out := map[string]nextgen.AnsibleVariable{}
	for _, item := range set.List() {
		m := item.(map[string]interface{})
		key := m["key"].(string)
		out[key] = nextgen.AnsibleVariable{
			Key:       key,
			Value:     m["value"].(string),
			ValueType: m["value_type"].(string),
			FileName:  m["file_name"].(string),
		}
	}
	if len(out) == 0 {
		return nil
	}
	return out
}

func readPlaybook(d *schema.ResourceData, resp *nextgen.ShowPlaybookResponse) diag.Diagnostics {
	d.SetId(resp.Identifier)
	d.Set("identifier", resp.Identifier)
	d.Set("org_id", resp.Org)
	d.Set("project_id", resp.Project)
	d.Set("name", resp.Name)
	d.Set("repository", resp.Repository)
	d.Set("repository_branch", resp.RepositoryBranch)
	d.Set("repository_commit", resp.RepositoryCommit)
	d.Set("repository_connector", resp.RepositoryConnector)
	d.Set("repository_path", resp.RepositoryPath)
	d.Set("ansible_galaxy", resp.AnsibleGalaxy)
	d.Set("ansible_galaxy_requirements_file", resp.AnsibleGalaxyRequirementsFile)

	if resp.Tags != "" {
		tags, err := unmarshalTags(resp.Tags)
		if err != nil {
			return diag.Errorf("failed to decode playbook tags: %s", err.Error())
		}
		d.Set("tags", tags)
	}

	vars, err := unmarshalVariables(resp.Vars)
	if err != nil {
		return diag.Errorf("failed to decode playbook vars: %s", err.Error())
	}
	d.Set("vars", flattenVariables(vars))

	envVars, err := unmarshalVariables(resp.EnvVars)
	if err != nil {
		return diag.Errorf("failed to decode playbook env_vars: %s", err.Error())
	}
	d.Set("env_vars", flattenVariables(envVars))

	return nil
}

func unmarshalVariables(raw string) (map[string]nextgen.AnsibleVariable, error) {
	if raw == "" {
		return nil, nil
	}
	out := map[string]nextgen.AnsibleVariable{}
	if err := json.Unmarshal([]byte(raw), &out); err != nil {
		return nil, err
	}
	return out, nil
}

func flattenVariables(vars map[string]nextgen.AnsibleVariable) []interface{} {
	if len(vars) == 0 {
		return nil
	}
	out := make([]interface{}, 0, len(vars))
	for _, v := range vars {
		out = append(out, map[string]interface{}{
			"key":        v.Key,
			"value":      v.Value,
			"value_type": v.ValueType,
			"file_name":  v.FileName,
		})
	}
	return out
}

func marshalTags(v interface{}) (string, error) {
	set, ok := v.(*schema.Set)
	if !ok {
		return "", nil
	}
	tags := helpers.ExpandTags(set.List())
	if len(tags) == 0 {
		return "", nil
	}
	b, err := json.Marshal(tags)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func unmarshalTags(data string) ([]interface{}, error) {
	if data == "" {
		return nil, nil
	}
	var m map[string]string
	if err := json.Unmarshal([]byte(data), &m); err == nil {
		out := make([]interface{}, 0, len(m))
		for _, t := range helpers.FlattenTags(m) {
			out = append(out, t)
		}
		return out, nil
	}
	var list []string
	if err := json.Unmarshal([]byte(data), &list); err == nil {
		out := make([]interface{}, 0, len(list))
		for _, t := range list {
			out = append(out, t)
		}
		return out, nil
	}
	return nil, nil
}

func parseError(err error, httpResp *http.Response) diag.Diagnostics {
	if httpResp != nil && httpResp.StatusCode == 401 {
		return diag.Errorf("%s\nHint:\n1) Please check if token has expired or is wrong.\n2) Harness Provider is misconfigured.", httpResp.Status)
	}
	if httpResp != nil && httpResp.StatusCode == 403 {
		return diag.Errorf("%s\nHint:\n1) Please check if the token has required permission for this operation.\n2) Please check if the token has expired or is wrong.", httpResp.Status)
	}

	se, ok := err.(nextgen.GenericSwaggerError)
	if !ok {
		return diag.FromErr(err)
	}

	iacmErr := nextgen.IacmError{}
	if jsonErr := json.Unmarshal(se.Body(), &iacmErr); jsonErr == nil && iacmErr.Message != "" {
		if httpResp != nil {
			return diag.Errorf("%s\nHint:\n1) %s", httpResp.Status, iacmErr.Message)
		}
		return diag.Errorf("%s", iacmErr.Message)
	}
	return diag.Errorf("%s", err.Error())
}
