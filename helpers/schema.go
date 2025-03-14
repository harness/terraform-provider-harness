package helpers

import (
	"fmt"
	"strings"

	"github.com/antihax/optional"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func MergeSchemas(src map[string]*schema.Schema, dest map[string]*schema.Schema) {
	for k, v := range src {
		dest[k] = v
	}
}

func SetSchemaFlagType(s *schema.Schema, flag SchemaFlagType) {
	switch flag {
	case SchemaFlagTypes.Computed:
		s.Computed = true
	case SchemaFlagTypes.Optional:
		s.Optional = true
	case SchemaFlagTypes.Required:
		s.Required = true
	}
}

func GetTagsSchema(flag SchemaFlagType) *schema.Schema {
	s := &schema.Schema{
		Description: "Tags to associate with the resource.",
		Type:        schema.TypeSet,
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
	}
	SetSchemaFlagType(s, flag)
	return s
}

func GetIdentifierSchema(flag SchemaFlagType) *schema.Schema {
	s := &schema.Schema{
		Description: "Unique identifier of the resource.",
		Type:        schema.TypeString,
	}

	if flag == SchemaFlagTypes.Required {
		s.ForceNew = true
	}

	SetSchemaFlagType(s, flag)

	return s
}

func GetProjectIdSchema(flag SchemaFlagType) *schema.Schema {
	s := &schema.Schema{
		Description: "Unique identifier of the project.",
		Type:        schema.TypeString,
	}

	if flag == SchemaFlagTypes.Required {
		s.ForceNew = true
	}

	SetSchemaFlagType(s, flag)
	return s
}

func GetOrgIdSchema(flag SchemaFlagType) *schema.Schema {
	s := &schema.Schema{
		Description: "Unique identifier of the organization.",
		Type:        schema.TypeString,
	}

	if flag == SchemaFlagTypes.Required {
		s.ForceNew = true
	}

	SetSchemaFlagType(s, flag)
	return s
}

func GetNameSchema(flag SchemaFlagType) *schema.Schema {
	s := &schema.Schema{
		Description: "Name of the resource.",
		Type:        schema.TypeString,
	}

	SetSchemaFlagType(s, flag)
	return s
}

func GetDescriptionSchema(flag SchemaFlagType) *schema.Schema {
	s := &schema.Schema{
		Description: "Description of the resource.",
		Type:        schema.TypeString,
	}
	SetSchemaFlagType(s, flag)
	return s
}

// SetCommonResourceSchema sets the default schema objects used for most resources.
func SetCommonResourceSchema(s map[string]*schema.Schema) {
	s["identifier"] = GetIdentifierSchema(SchemaFlagTypes.Required)
	s["description"] = GetDescriptionSchema(SchemaFlagTypes.Optional)
	s["name"] = GetNameSchema(SchemaFlagTypes.Required)
	s["tags"] = GetTagsSchema(SchemaFlagTypes.Optional)
}

// SetCommonDataSourceSchema sets the default schema objects used for most data sources.
func SetCommonDataSourceSchema(s map[string]*schema.Schema) {
	s["identifier"] = GetIdentifierSchema(SchemaFlagTypes.Optional)
	s["description"] = GetDescriptionSchema(SchemaFlagTypes.Computed)
	s["name"] = GetNameSchema(SchemaFlagTypes.Optional)
	s["tags"] = GetTagsSchema(SchemaFlagTypes.Computed)
}

func SetCommonDataSourceSchemaWRequired(s map[string]*schema.Schema) {
	s["identifier"] = GetIdentifierSchema(SchemaFlagTypes.Required)
	s["description"] = GetDescriptionSchema(SchemaFlagTypes.Computed)
	s["name"] = GetNameSchema(SchemaFlagTypes.Required)
	s["tags"] = GetTagsSchema(SchemaFlagTypes.Computed)
}

func SetCommonDataSourceSchemaIdentifierRequired(s map[string]*schema.Schema) {
	s["identifier"] = GetIdentifierSchema(SchemaFlagTypes.Required)
	s["description"] = GetDescriptionSchema(SchemaFlagTypes.Computed)
	s["name"] = GetNameSchema(SchemaFlagTypes.Optional)
	s["tags"] = GetTagsSchema(SchemaFlagTypes.Computed)
}

func SetOrgLevelDataSourceSchema(s map[string]*schema.Schema) {
	SetCommonDataSourceSchema(s)
	s["org_id"] = GetOrgIdSchema(SchemaFlagTypes.Required)
}

func SetOrgLevelDataSourceSchemaIdentifierRequired(s map[string]*schema.Schema) {
	SetCommonDataSourceSchemaIdentifierRequired(s)
	s["org_id"] = GetOrgIdSchema(SchemaFlagTypes.Required)
}

func SetProjectLevelDataSourceSchema(s map[string]*schema.Schema) {
	SetOrgLevelDataSourceSchema(s)
	s["project_id"] = GetProjectIdSchema(SchemaFlagTypes.Required)
}

func SetProjectLevelDataSourceSchemaIdentifierRequired(s map[string]*schema.Schema) {
	SetOrgLevelDataSourceSchemaIdentifierRequired(s)
	s["project_id"] = GetProjectIdSchema(SchemaFlagTypes.Required)
}

func SetOptionalOrgAndProjectLevelDataSourceSchema(s map[string]*schema.Schema) {
	s["org_id"] = GetOrgIdSchema(SchemaFlagTypes.Optional)
	s["project_id"] = GetProjectIdSchema(SchemaFlagTypes.Optional)
}

// SetOrgLevelResourceSchema sets the default schema objects used for org level resources.
func SetOrgLevelResourceSchema(s map[string]*schema.Schema) {
	SetCommonResourceSchema(s)
	s["org_id"] = GetOrgIdSchema(SchemaFlagTypes.Required)
}

// SetProjectLevelResourceSchema sets the default schema objects used for project level resources.
func SetProjectLevelResourceSchema(s map[string]*schema.Schema) {
	SetOrgLevelResourceSchema(s)
	s["project_id"] = GetProjectIdSchema(SchemaFlagTypes.Required)
}

func SetMultiLevelResourceSchema(s map[string]*schema.Schema) {
	SetCommonResourceSchema(s)
	s["org_id"] = GetOrgIdSchema(SchemaFlagTypes.Optional)
	s["project_id"] = GetProjectIdSchema(SchemaFlagTypes.Optional)
	s["project_id"].RequiredWith = []string{"org_id"}
}

func SetMultiLevelDatasourceSchema(s map[string]*schema.Schema) {
	SetCommonDataSourceSchema(s)
	s["org_id"] = GetOrgIdSchema(SchemaFlagTypes.Optional)
	s["project_id"] = GetProjectIdSchema(SchemaFlagTypes.Optional)
	s["project_id"].RequiredWith = []string{"org_id"}
}

func SetMultiLevelDatasourceSchemaIdentifierRequired(s map[string]*schema.Schema) {
	SetCommonDataSourceSchemaIdentifierRequired(s)
	s["org_id"] = GetOrgIdSchema(SchemaFlagTypes.Optional)
	s["project_id"] = GetProjectIdSchema(SchemaFlagTypes.Optional)
	s["project_id"].RequiredWith = []string{"org_id"}
}

func SetMultiLevelDatasourceSchemaWithoutCommonFields(s map[string]*schema.Schema) {
	s["org_id"] = GetOrgIdSchema(SchemaFlagTypes.Optional)
	s["project_id"] = GetProjectIdSchema(SchemaFlagTypes.Optional)
	s["project_id"].RequiredWith = []string{"org_id"}
}

func BuildField(d *schema.ResourceData, field string) optional.String {
	if arr, ok := d.GetOk(field); ok {
		return optional.NewString(arr.(string))
	}
	return optional.EmptyString()
}

func BuildFieldInt32(d *schema.ResourceData, field string) optional.Int32 {
	if arr, ok := d.GetOk(field); ok {
		return optional.NewInt32(int32(arr.(int)))
	}
	return optional.EmptyInt32()
}

func BuildFieldForBoolean(d *schema.ResourceData, field string) optional.Bool {
	if arr, ok := d.GetOk(field); ok {

		return optional.NewBool(arr.(string) == "true")
	}
	return optional.EmptyBool()
}

func BuildFieldBool(d *schema.ResourceData, field string) optional.Bool {
	if b, ok := d.GetOk(field); ok {
		return optional.NewBool(b.(bool))
	}

	return optional.EmptyBool()
}

// PipelineResourceImporter defines the importer configuration for all pipeline level resources.
var PipelineResourceImporter = &schema.ResourceImporter{
	State: func(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
		parts := strings.Split(d.Id(), "/")
		d.Set("org_id", parts[0])
		d.Set("project_id", parts[1])
		d.Set("pipeline_id", parts[2])
		d.Set("identifier", parts[3])
		d.SetId(parts[3])

		return []*schema.ResourceData{d}, nil
	},
}

var DBInstanceResourceImporter = &schema.ResourceImporter{
	State: func(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
		parts := strings.Split(d.Id(), "/")
		d.Set("org_id", parts[0])
		d.Set("project_id", parts[1])
		d.Set("schema", parts[2])
		d.Set("identifier", parts[3])
		d.SetId(parts[3])

		return []*schema.ResourceData{d}, nil
	},
}

var TriggerResourceImporter = &schema.ResourceImporter{
	State: func(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
		parts := strings.Split(d.Id(), "/")
		d.Set("org_id", parts[0])
		d.Set("project_id", parts[1])
		d.Set("target_id", parts[2])
		d.Set("identifier", parts[3])
		d.SetId(parts[3])

		return []*schema.ResourceData{d}, nil
	},
}

var EnvRelatedResourceImporter = &schema.ResourceImporter{
	State: func(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
		parts := strings.Split(d.Id(), "/")
		partCount := len(parts)
		isAccountEntity := partCount == 2
		isOrgEntity := partCount == 3
		isProjectEntity := partCount == 4
		if isAccountEntity {
			d.Set("env_id", parts[0])
			d.Set("identifier", parts[1])
			d.SetId(parts[1])
			return []*schema.ResourceData{d}, nil
		}
		if isOrgEntity {
			d.Set("org_id", parts[0])
			d.Set("env_id", parts[1])
			d.Set("identifier", parts[2])
			d.SetId(parts[2])
			return []*schema.ResourceData{d}, nil
		}
		if isProjectEntity {
			d.Set("org_id", parts[0])
			d.Set("project_id", parts[1])
			d.Set("env_id", parts[2])
			d.Set("identifier", parts[3])
			d.SetId(parts[3])
			return []*schema.ResourceData{d}, nil
		}

		return nil, fmt.Errorf("invalid identifier: %s", d.Id())
	},
}

var ServiceOverrideResourceImporter = &schema.ResourceImporter{
	State: func(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
		parts := strings.Split(d.Id(), "/")
		partCount := len(parts)
		isAccountEntity := partCount == 1
		isOrgEntity := partCount == 2
		isProjectEntity := partCount == 3
		if isAccountEntity {
			d.Set("env_id", parts[0])
			return []*schema.ResourceData{d}, nil
		}
		if isOrgEntity {
			d.Set("org_id", parts[0])
			d.Set("env_id", parts[1])
			return []*schema.ResourceData{d}, nil
		}
		if isProjectEntity {
			d.Set("org_id", parts[0])
			d.Set("project_id", parts[1])
			d.Set("env_id", parts[2])
			return []*schema.ResourceData{d}, nil
		}

		return nil, fmt.Errorf("invalid identifier: %s", d.Id())
	},
}

var ServiceOverrideV2ResourceImporter = &schema.ResourceImporter{
	State: func(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
		parts := strings.Split(d.Id(), "/")
		partCount := len(parts)
		isAccountEntity := partCount == 1
		isOrgEntity := partCount == 2
		isProjectEntity := partCount == 3
		if isAccountEntity {
			d.SetId(parts[0])
			return []*schema.ResourceData{d}, nil
		}
		if isOrgEntity {
			d.Set("org_id", parts[0])
			d.SetId(parts[1])
			return []*schema.ResourceData{d}, nil
		}
		if isProjectEntity {
			d.Set("org_id", parts[0])
			d.Set("project_id", parts[1])
			d.SetId(parts[2])
			return []*schema.ResourceData{d}, nil
		}

		return nil, fmt.Errorf("invalid identifier: %s", d.Id())
	},
}

var UserResourceImporter = &schema.ResourceImporter{
	State: func(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
		parts := strings.Split(d.Id(), "/")
		d.Set("email", parts[0])
		if len(parts) > 1 {
			d.Set("org_id", parts[1])
		}
		if len(parts) > 2 {
			d.Set("project_id", parts[2])
		}
		return []*schema.ResourceData{d}, nil
	},
}

// ProjectResourceImporter defines the importer configuration for all project level resources.
// The id used for the import should be in the format <org_id>/<project_id>/<identifier>
// The id used for the import should be in the format <org_id>/<project_id>/<identifier>/<branch>
var ProjectResourceImporter = &schema.ResourceImporter{
	State: func(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
		parts := strings.Split(d.Id(), "/")
		d.Set("org_id", parts[0])
		d.Set("project_id", parts[1])
		d.Set("identifier", parts[2])
		d.SetId(parts[2])

		if len(parts) == 4 {
			d.Set("git_details", []interface{}{map[string]interface{}{"branch_name": parts[3]}})
		}

		return []*schema.ResourceData{d}, nil
	},
}

// GitopsAgentResourceImporter defines the importer configuration for all project level gitops agent resources.
// The id used for the import should be in the format <org_id>/<project_id>/<identifier>/<agentId>
var GitopsAgentResourceImporter = &schema.ResourceImporter{
	State: func(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
		parts := strings.Split(d.Id(), "/")
		if len(parts) == 2 { //Account level
			d.Set("agent_id", parts[0])
			d.Set("identifier", parts[1])
			d.SetId(parts[1])
			return []*schema.ResourceData{d}, nil
		}

		if len(parts) == 3 { //Org level
			d.Set("org_id", parts[0])
			d.Set("agent_id", parts[1])
			d.Set("identifier", parts[2])
			d.SetId(parts[2])
			return []*schema.ResourceData{d}, nil
		}

		if len(parts) == 4 { //Project level
			d.Set("org_id", parts[0])
			d.Set("project_id", parts[1])
			d.Set("agent_id", parts[2])
			d.Set("identifier", parts[3])
			d.SetId(parts[3])
			return []*schema.ResourceData{d}, nil
		}

		return nil, fmt.Errorf("invalid identifier: %s", d.Id())
	},
}

var GitopsAgentApplicationImporter = &schema.ResourceImporter{
	State: func(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
		parts := strings.Split(d.Id(), "/")
		if len(parts) == 4 { //Project level
			d.Set("org_id", parts[0])
			d.Set("project_id", parts[1])
			d.Set("agent_id", parts[2])
			d.Set("name", parts[3])
			d.SetId(parts[3])
			return []*schema.ResourceData{d}, nil
		}

		return nil, fmt.Errorf("invalid identifier: %s", d.Id())
	},
}

// GitopsAppProjectMappingImporter defines the importer configuration for app project mapping.
// The id used for the import should be in the format <org_id>/<project_id>/<identifier>/<argo_project_name>\
// It is used always at project level.
var GitopsAppProjectMappingImporter = &schema.ResourceImporter{
	State: func(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
		parts := strings.Split(d.Id(), "/")

		// Approject mapping always has 4 parts
		if len(parts) == 4 { //Project level
			d.Set("org_id", parts[0])
			d.Set("project_id", parts[1])
			d.Set("agent_id", parts[2])
			d.Set("argo_project_name", parts[3])
			// During import we are using argo_project_name as identifier not the actual identifier which is mongo id
			// that way we are not fetching mapping by mongo id but by argo_project_name, agent_id, account_id, org_id and project_id.
			d.SetId(parts[3])
			return []*schema.ResourceData{d}, nil
		}

		return nil, fmt.Errorf("invalid identifier: %s", d.Id())
	},
}

// GitopsAgentResourceImporter defines the importer configuration for all project level gitops agent resources.
// The id used for the import should be in the format <agent_id>/<query_name>
var GitopsAgentProjectImporter = &schema.ResourceImporter{
	State: func(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
		parts := strings.Split(d.Id(), "/")
		if len(parts) == 2 { //Account level
			d.Set("agent_id", parts[0])
			d.Set("query_name", parts[1])
			d.SetId(parts[1])
			return []*schema.ResourceData{d}, nil
		}

		if len(parts) == 3 { //Org level
			d.Set("org_id", parts[0])
			d.Set("agent_id", parts[1])
			d.Set("query_name", parts[2])
			d.SetId(parts[2])
			return []*schema.ResourceData{d}, nil
		}

		if len(parts) == 4 { //Project level
			d.Set("org_id", parts[0])
			d.Set("project_id", parts[1])
			d.Set("agent_id", parts[2])
			d.Set("query_name", parts[3])
			d.SetId(parts[3])
			return []*schema.ResourceData{d}, nil
		}

		return nil, fmt.Errorf("invalid identifier: %s", d.Id())
	},
}

// The id used for the import should be in the format <org_id>/<project_id>/<repoIdentifier>/<identifier>
var RepoRuleResourceImporter = &schema.ResourceImporter{
	State: func(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
		parts := strings.Split(d.Id(), "/")
		if len(parts) == 2 { //Account level
			d.Set("repo_identifier", parts[0])
			d.Set("identifier", parts[1])
			d.SetId(parts[1])
			return []*schema.ResourceData{d}, nil
		}

		if len(parts) == 3 { //Org level
			d.Set("org_id", parts[0])
			d.Set("repo_identifier", parts[1])
			d.Set("identifier", parts[2])
			d.SetId(parts[2])
			return []*schema.ResourceData{d}, nil
		}

		if len(parts) == 4 { //Project level
			d.Set("org_id", parts[0])
			d.Set("project_id", parts[1])
			d.Set("repo_identifier", parts[2])
			d.Set("identifier", parts[3])
			d.SetId(parts[3])
			return []*schema.ResourceData{d}, nil
		}

		return nil, fmt.Errorf("invalid identifier: %s", d.Id())
	},
}

var GitopsRepoCertResourceImporter = &schema.ResourceImporter{
	State: func(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
		parts := strings.Split(d.Id(), "/")
		d.Set("agent_id", parts[0])
		d.SetId(parts[1])
		return []*schema.ResourceData{d}, nil
	},
}

// OrgResourceImporter defines the importer configuration for all organization level resources.
// The id used for the import should be in the format <org_id>/<identifier>
var OrgResourceImporter = &schema.ResourceImporter{
	State: func(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
		parts := strings.Split(d.Id(), "/")
		d.Set("org_id", parts[0])
		d.Set("identifier", parts[1])
		d.SetId(parts[1])

		return []*schema.ResourceData{d}, nil
	},
}

// MultiLevelResourceImporter defines the importer configuration for all multi level resources.
// The format used for the id is as follows:
//   - Account Level: <identifier>
//   - Org Level: <org_id>/<identifier>
//   - Project Level: <org_id>/<project_id>/<identifier>
var MultiLevelResourceImporter = &schema.ResourceImporter{
	State: func(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
		parts := strings.Split(d.Id(), "/")

		partCount := len(parts)
		isAccountConnector := partCount == 1
		isOrgConnector := partCount == 2
		isProjectConnector := partCount == 3

		if isAccountConnector {
			d.SetId(parts[0])
			d.Set("identifier", parts[0])
			return []*schema.ResourceData{d}, nil
		}

		if isOrgConnector {
			d.SetId(parts[1])
			d.Set("identifier", parts[1])
			d.Set("org_id", parts[0])
			return []*schema.ResourceData{d}, nil
		}

		if isProjectConnector {
			d.SetId(parts[2])
			d.Set("identifier", parts[2])
			d.Set("project_id", parts[1])
			d.Set("org_id", parts[0])
			return []*schema.ResourceData{d}, nil
		}

		return nil, fmt.Errorf("invalid identifier: %s", d.Id())
	},
}

var MultiLevelFilterImporter = &schema.ResourceImporter{
	State: func(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
		parts := strings.Split(d.Id(), "/")

		partCount := len(parts)
		isAccountConnector := partCount == 2
		isOrgConnector := partCount == 3
		isProjectConnector := partCount == 4

		if isAccountConnector {
			d.SetId(parts[0])
			d.Set("identifier", parts[0])
			d.Set("type", parts[1])
			return []*schema.ResourceData{d}, nil
		}

		if isOrgConnector {
			d.SetId(parts[1])
			d.Set("identifier", parts[1])
			d.Set("org_id", parts[0])
			d.Set("type", parts[2])
			return []*schema.ResourceData{d}, nil
		}

		if isProjectConnector {
			d.SetId(parts[2])
			d.Set("identifier", parts[2])
			d.Set("project_id", parts[1])
			d.Set("org_id", parts[0])
			d.Set("type", parts[3])
			return []*schema.ResourceData{d}, nil
		}

		return nil, fmt.Errorf("invalid identifier: %s", d.Id())
	},
}

var GitWebhookResourceImporter = &schema.ResourceImporter{
	State: func(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
		parts := strings.Split(d.Id(), "/")
		d.Set("identifier", parts[0])
		if len(parts) > 1 {
			d.Set("org_id", parts[1])
		}
		if len(parts) > 2 {
			d.Set("project_id", parts[2])
		}
		d.SetId(parts[0])

		return []*schema.ResourceData{d}, nil
	},
}

var AccountLevelResourceImporter = &schema.ResourceImporter{
	State: func(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
		if d.Id() != "" {
			return []*schema.ResourceData{d}, nil
		}

		return nil, fmt.Errorf("invalid identifier: %s", d.Id())
	},
}
