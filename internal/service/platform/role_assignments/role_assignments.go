package role_assignments

import (
	"context"
	"net/http"

	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func ResourceRoleAssignments() *schema.Resource {
	resource := &schema.Resource{
		Description: "Resource for creating role assignments in Harness.",

		ReadContext:   resourceRoleAssignmentsRead,
		UpdateContext: resourceRoleAssignmentsCreateorUpdate,
		CreateContext: resourceRoleAssignmentsCreateorUpdate,
		DeleteContext: resourceRoleAssignmentsDelete,
		Importer:      helpers.MultiLevelResourceImporter,

		Schema: map[string]*schema.Schema{
			"identifier": {
				Description: "Identifier for role assignment.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"resource_group_identifier": {
				Description: "Resource group identifier.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"role_identifier": {
				Description: "Role identifier.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"principal": {
				Description: "Principal.",
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"scope_level": {
							Description: "Scope level.",
							Type:        schema.TypeString,
							Computed:    true,
							Optional:    true,
						},
						"identifier": {
							Description: "Identifier.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"type": {
							Description:  "Type.",
							Type:         schema.TypeString,
							Optional:     true,
							ValidateFunc: validation.StringInSlice([]string{"USER", "USER_GROUP", "SERVICE", "API_KEY", "SERVICE_ACCOUNT"}, false),
						},
					},
				},
			},
			"disabled": {
				Description: "Disabled or not.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"managed": {
				Description: "Managed or not.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"project_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Project Identifier",
			},
			"org_id": {
				Description: "Org identifier.",
				Type:        schema.TypeString,
				Optional:    true,
			},
		},
	}
	return resource
}

func resourceRoleAssignmentsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)
	id := d.Get("identifier").(string)

	resp, httpResp, err := c.RoleAssignmentsApi.GetRoleAssignment(ctx, c.AccountId, id, &nextgen.RoleAssignmentsApiGetRoleAssignmentOpts{
		OrgIdentifier:     helpers.BuildField(d, "org_id"),
		ProjectIdentifier: helpers.BuildField(d, "project_id"),
	})

	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	if resp.Data == nil {
		return nil
	}

	readRoleAssignments(d, resp.Data.RoleAssignment)

	return nil
}

func resourceRoleAssignmentsCreateorUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	var err error
	var resp nextgen.ResponseDtoRoleAssignmentResponse
	var httpResp *http.Response

	id := d.Id()
	roleAssignment := buildRoleAssignment(d)

	if id == "" {
		resp, httpResp, err = c.RoleAssignmentsApi.PostRoleAssignment(ctx, *roleAssignment, c.AccountId, &nextgen.RoleAssignmentsApiPostRoleAssignmentOpts{
			OrgIdentifier:     helpers.BuildField(d, "org_id"),
			ProjectIdentifier: helpers.BuildField(d, "project_id"),
		})
	} else {
		resp, httpResp, err = c.RoleAssignmentsApi.DeleteRoleAssignment(ctx, c.AccountId, id, &nextgen.RoleAssignmentsApiDeleteRoleAssignmentOpts{
			OrgIdentifier:     helpers.BuildField(d, "org_id"),
			ProjectIdentifier: helpers.BuildField(d, "project_id"),
		})

		if err != nil {
			return helpers.HandleApiError(err, d, httpResp)
		}

		resp, httpResp, err = c.RoleAssignmentsApi.PostRoleAssignment(ctx, *roleAssignment, c.AccountId, &nextgen.RoleAssignmentsApiPostRoleAssignmentOpts{
			OrgIdentifier:     helpers.BuildField(d, "org_id"),
			ProjectIdentifier: helpers.BuildField(d, "project_id"),
		})
	}

	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	readRoleAssignments(d, resp.Data.RoleAssignment

	return nil
}

func resourceRoleAssignmentsDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	_, httpResp, err := c.RoleAssignmentsApi.DeleteRoleAssignment(ctx, c.AccountId, d.Id(), &nextgen.RoleAssignmentsApiDeleteRoleAssignmentOpts{
		OrgIdentifier:     helpers.BuildField(d, "org_id"),
		ProjectIdentifier: helpers.BuildField(d, "project_id"),
	})

	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	return nil
}

func buildRoleAssignment(d *schema.ResourceData) *nextgen.RoleAssignment {
	roleAssignment := &nextgen.RoleAssignment{
		Principal: &nextgen.AuthzPrincipal{},
	}

	if attr, ok := d.GetOk("disabled"); ok {
		roleAssignment.Disabled = attr.(bool)
	}

	if attr, ok := d.GetOk("managed"); ok {
		roleAssignment.Managed = attr.(bool)
	}

	if attr, ok := d.GetOk("identifier"); ok {
		roleAssignment.Identifier = attr.(string)
	}

	if attr, ok := d.GetOk("resource_group_identifier"); ok {
		roleAssignment.ResourceGroupIdentifier = attr.(string)
	}

	if attr, ok := d.GetOk("role_identifier"); ok {
		roleAssignment.RoleIdentifier = attr.(string)
	}

	if attr, ok := d.GetOk("principal"); ok {
		config := attr.([]interface{})[0].(map[string]interface{})
		if attr, ok := config["scope_level"]; ok {
			roleAssignment.Principal.ScopeLevel = attr.(string)
		}

		if attr, ok := config["identifier"]; ok {
			roleAssignment.Principal.Identifier = attr.(string)
		}

		if attr, ok := config["type"]; ok {
			roleAssignment.Principal.Type_ = attr.(string)
		}
	}
	return roleAssignment
}

func readRoleAssignments(d *schema.ResourceData, roleAssignments *nextgen.RoleAssignment) {
	d.SetId(roleAssignments.Identifier)
	d.Set("identifier", roleAssignments.Identifier)
	d.Set("disabled", roleAssignments.Disabled)
	d.Set("managed", roleAssignments.Managed)
	d.Set("resource_group_identifier", roleAssignments.ResourceGroupIdentifier)
	d.Set("role_identifier", roleAssignments.RoleIdentifier)
	d.Set("principal", []interface{}{
		map[string]interface{}{
			"scope_level": roleAssignments.Principal.ScopeLevel,
			"identifier":  roleAssignments.Principal.Identifier,
			"type":        roleAssignments.Principal.Type_,
		},
	})
}
