package user

import (
	"fmt"
	"strings"

	"github.com/harness/harness-go-sdk/harness/cd/graphql"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var actionsAll = []string{"CREATE", "READ", "UPDATE", "DELETE", "EXECUTE_WORKFLOW", "EXECUTE_PIPELINE", "ROLLBACK_WORKFLOW"}
var standardActions = []string{"CREATE", "READ", "UPDATE", "DELETE"}
var deploymentActions = []string{"READ", "EXECUTE_WORKFLOW", "EXECUTE_PIPELINE", "ROLLBACK_WORKFLOW", "ABORT_WORKFLOW"}

func getUserGroupAccountPermissionsSchema() *schema.Schema {
	return &schema.Schema{
		Description: fmt.Sprintf("The account permissions of the user group. Valid options are %s", strings.Join(graphql.AccountPermissionTypeValues, ", ")),
		Type:        schema.TypeSet,
		Optional:    true,
		MinItems:    1,
		Elem:        &schema.Schema{Type: schema.TypeString},
	}
}

func getUserGroupAppPermissionsSchema() *schema.Schema {
	return &schema.Schema{
		Description: "Application specific permissions",
		Type:        schema.TypeList,
		Optional:    true,
		MaxItems:    1,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"all": {
					Description: "The permission to perform actions against all resources.",
					Type:        schema.TypeSet,
					Optional:    true,
					Elem: &schema.Resource{
						Schema: map[string]*schema.Schema{
							"app_ids": {
								Description: "The application IDs to which the permission applies. Leave empty to apply to all applications.",
								Type:        schema.TypeSet,
								Optional:    true,
								Elem:        &schema.Schema{Type: schema.TypeString},
							},
							"actions": {
								Description: fmt.Sprintf("The actions allowed to be performed. Valid options are %s", strings.Join(actionsAll, ", ")),
								Type:        schema.TypeSet,
								Required:    true,
								MinItems:    1,
								Elem:        &schema.Schema{Type: schema.TypeString},
							},
						},
					},
				},
				"workflow": {
					Description: "Permission configuration to perform actions against workflows.",
					Type:        schema.TypeSet,
					Optional:    true,
					Elem: &schema.Resource{
						Schema: map[string]*schema.Schema{
							"app_ids": {
								Description: "The application IDs to which the permission applies. Leave empty to apply to all applications.",
								Type:        schema.TypeSet,
								Optional:    true,
								Elem:        &schema.Schema{Type: schema.TypeString},
							},
							"filters": {
								Description: fmt.Sprintf("The filters to apply to the action. Valid options are: %s.", strings.Join(graphql.WorkflowPermissionFiltersSlice, ", ")),
								Type:        schema.TypeSet,
								Required:    true,
								MinItems:    1,
								Elem:        &schema.Schema{Type: schema.TypeString},
							},
							"actions": {
								Description: fmt.Sprintf("The actions allowed to be performed. Valid options are %s", strings.Join(standardActions, ", ")),
								Type:        schema.TypeSet,
								Required:    true,
								MinItems:    1,
								Elem:        &schema.Schema{Type: schema.TypeString},
							},
						},
					},
				},
				"deployment": {
					Description: "Permission configuration to perform actions against deployments.",
					Type:        schema.TypeSet,
					Optional:    true,
					Elem: &schema.Resource{
						Schema: map[string]*schema.Schema{
							"app_ids": {
								Description: "The application IDs to which the permission applies. Leave empty to apply to all applications.",
								Type:        schema.TypeSet,
								Optional:    true,
								Elem:        &schema.Schema{Type: schema.TypeString},
							},
							"env_ids": {
								Description: "The environment IDs to which the permission applies. Leave empty to apply to all environments.",
								Type:        schema.TypeSet,
								Optional:    true,
								Elem:        &schema.Schema{Type: schema.TypeString},
							},
							"filters": {
								Description: fmt.Sprintf("The filters to apply to the action. Valid options are: %s.", strings.Join(graphql.DeploymentPermissionFiltersSlice, ", ")),
								Type:        schema.TypeSet,
								Required:    true,
								MinItems:    1,
								Elem:        &schema.Schema{Type: schema.TypeString},
							},
							"actions": {
								Description: fmt.Sprintf("The actions allowed to be performed. Valid options are %s", strings.Join(deploymentActions, ", ")),
								Type:        schema.TypeSet,
								Required:    true,
								MinItems:    1,
								Elem:        &schema.Schema{Type: schema.TypeString},
							},
						},
					},
				},
				"environment": {
					Description: "Permission configuration to perform actions against workflows.",
					Type:        schema.TypeSet,
					Optional:    true,
					Elem: &schema.Resource{
						Schema: map[string]*schema.Schema{
							"app_ids": {
								Description: "The application IDs to which the permission applies. Leave empty to apply to all applications.",
								Type:        schema.TypeSet,
								Optional:    true,
								Elem:        &schema.Schema{Type: schema.TypeString},
							},
							"env_ids": {
								Description: "The environment IDs to which the permission applies. Leave empty to apply to all environments.",
								Type:        schema.TypeSet,
								Optional:    true,
								Elem:        &schema.Schema{Type: schema.TypeString},
							},
							"filters": {
								Description: fmt.Sprintf("The filters to apply to the action. Valid options are: %s.", strings.Join(graphql.EnvFiltersSlice, ", ")),
								Type:        schema.TypeSet,
								Required:    true,
								MinItems:    1,
								Elem:        &schema.Schema{Type: schema.TypeString},
							},
							"actions": {
								Description: fmt.Sprintf("The actions allowed to be performed. Valid options are %s", strings.Join(standardActions, ", ")),
								Type:        schema.TypeSet,
								Required:    true,
								MinItems:    1,
								Elem:        &schema.Schema{Type: schema.TypeString},
							},
						},
					},
				},
				"pipeline": {
					Description: "Permission configuration to perform actions against pipelines.",
					Type:        schema.TypeSet,
					Optional:    true,
					Elem: &schema.Resource{
						Schema: map[string]*schema.Schema{
							"app_ids": {
								Description: "The application IDs to which the permission applies. Leave empty to apply to all applications.",
								Type:        schema.TypeSet,
								Optional:    true,
								Elem:        &schema.Schema{Type: schema.TypeString},
							},
							"env_ids": {
								Description: "The environment IDs to which the permission applies. Leave empty to apply to all environments.",
								Type:        schema.TypeSet,
								Optional:    true,
								Elem:        &schema.Schema{Type: schema.TypeString},
							},
							"filters": {
								Description: fmt.Sprintf("The filters to apply to the action. Valid options are: %s.", strings.Join(graphql.PipelinePermissionFiltersSlice, ", ")),
								Type:        schema.TypeSet,
								Required:    true,
								MinItems:    1,
								Elem:        &schema.Schema{Type: schema.TypeString},
							},
							"actions": {
								Description: fmt.Sprintf("The actions allowed to be performed. Valid options are %s", standardActions),
								Type:        schema.TypeSet,
								Required:    true,
								MinItems:    1,
								Elem:        &schema.Schema{Type: schema.TypeString},
							},
						},
					},
				},
				"provisioner": {
					Description: "Permission configuration to perform actions against provisioners.",
					Type:        schema.TypeSet,
					Optional:    true,
					Elem: &schema.Resource{
						Schema: map[string]*schema.Schema{
							"app_ids": {
								Description: "The application IDs to which the permission applies. Leave empty to apply to all applications.",
								Type:        schema.TypeSet,
								Optional:    true,
								Elem:        &schema.Schema{Type: schema.TypeString},
							},
							"provisioner_ids": {
								Description: "The provisioner IDs to which the permission applies. Leave empty to apply to all provisioners.",
								Type:        schema.TypeSet,
								Optional:    true,
								Elem:        &schema.Schema{Type: schema.TypeString},
							},
							"actions": {
								Description: fmt.Sprintf("The actions allowed to be performed. Valid options are %s", strings.Join(standardActions, ", ")),
								Type:        schema.TypeSet,
								Required:    true,
								MinItems:    1,
								Elem:        &schema.Schema{Type: schema.TypeString},
							},
						},
					},
				},
				"service": {
					Description: "Permission configuration to perform actions against services.",
					Type:        schema.TypeSet,
					Optional:    true,
					Elem: &schema.Resource{
						Schema: map[string]*schema.Schema{
							"app_ids": {
								Description: "The application IDs to which the permission applies. Leave empty to apply to all applications.",
								Type:        schema.TypeSet,
								Optional:    true,
								Elem:        &schema.Schema{Type: schema.TypeString},
							},
							"service_ids": {
								Description: "The service IDs to which the permission applies. Leave empty to apply to all services.",
								Type:        schema.TypeSet,
								Optional:    true,
								Elem:        &schema.Schema{Type: schema.TypeString},
							},
							"actions": {
								Description: fmt.Sprintf("The actions allowed to be performed. Valid options are %s", strings.Join(standardActions, ", ")),
								Type:        schema.TypeSet,
								Required:    true,
								MinItems:    1,
								Elem:        &schema.Schema{Type: schema.TypeString},
							},
						},
					},
				},
				"template": {
					Description: "Permission configuration to perform actions against templates.",
					Type:        schema.TypeSet,
					Optional:    true,
					Elem: &schema.Resource{
						Schema: map[string]*schema.Schema{
							"app_ids": {
								Description: "The application IDs to which the permission applies. Leave empty to apply to all applications.",
								Type:        schema.TypeSet,
								Optional:    true,
								Elem:        &schema.Schema{Type: schema.TypeString},
							},
							"template_ids": {
								Description: "The template IDs to which the permission applies. Leave empty to apply to all environments.",
								Type:        schema.TypeSet,
								Optional:    true,
								Elem:        &schema.Schema{Type: schema.TypeString},
							},
							"actions": {
								Description: fmt.Sprintf("The actions allowed to be performed. Valid options are %s", strings.Join(standardActions, ", ")),
								Type:        schema.TypeSet,
								Required:    true,
								MinItems:    1,
								Elem:        &schema.Schema{Type: schema.TypeString},
							},
						},
					},
				},
			},
		},
	}
}
