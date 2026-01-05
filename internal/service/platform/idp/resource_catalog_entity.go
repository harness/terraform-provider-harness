package idp

import (
	"context"
	"fmt"
	"net/http"
	"strings"

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
				Required:    true,
				ForceNew:    true,
				ValidateFunc: validation.StringInSlice([]string{
					"component", "group", "user", "workflow", "resource", "system",
				}, false),
			},
			"org_id":     helpers.GetOrgIdSchema(helpers.SchemaFlagTypes.Optional),
			"project_id": helpers.GetProjectIdSchema(helpers.SchemaFlagTypes.Optional),
			"yaml": {
				Type:             schema.TypeString,
				Description:      "YAML definition of the catalog entity",
				Required:         true,
				DiffSuppressFunc: helpers.YamlDiffSuppressFunction,
			},
		},
	}
	resource.Schema["project_id"].RequiredWith = []string{"org_id"}

	return resource
}

func resourceCatalogEntityRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetIDPClientWithContext(ctx)

	entityInfo, err := getAndVerifyCatalogEntityInfo(d)
	if err != nil {
		return diag.Errorf("error in validating yaml and inputs: %v", err)
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
		return helpers.HandleReadApiError(err, d, httpResp)
	}

	readCatalogEntity(d, resp)

	return nil
}

func resourceCatalogEntityUpdateOrCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetIDPClientWithContext(ctx)

	var err error
	var resp idp.EntityResponse
	var httpResp *http.Response

	id := d.Id()
	entityInfo, err := getAndVerifyCatalogEntityInfo(d)
	if err != nil {
		return diag.Errorf("failed to parse yaml: %v", err)
	}

	yaml := d.Get("yaml").(string)

	if id == "" {
		resp, httpResp, err = c.EntitiesApi.CreateEntity(ctx, idp.EntityCreateRequest{
			Yaml: yaml,
		},
			&idp.EntitiesApiCreateEntityOpts{
				OrgIdentifier:     entityInfo.OrgId,
				ProjectIdentifier: entityInfo.ProjectId,
				HarnessAccount:    optional.NewString(c.AccountId),
			})
	} else {
		resp, httpResp, err = c.EntitiesApi.UpdateEntity(ctx, idp.EntityUpdateRequest{
			Yaml: yaml,
		},
			entityInfo.Scope, entityInfo.Kind, id, &idp.EntitiesApiUpdateEntityOpts{
				OrgIdentifier:     entityInfo.OrgId,
				ProjectIdentifier: entityInfo.ProjectId,
				HarnessAccount:    optional.NewString(c.AccountId),
			})
	}

	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	readCatalogEntity(d, resp)

	return nil
}

func resourceCatalogEntityDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetIDPClientWithContext(ctx)

	id := d.Id()
	entityInfo, err := getAndVerifyCatalogEntityInfo(d)
	if err != nil {
		return diag.Errorf("failed to parse yaml: %v", err)
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

		return helpers.HandleApiError(err, d, httpResp)
	}

	return nil
}

func readCatalogEntity(d *schema.ResourceData, entity idp.EntityResponse) {
	d.SetId(entity.Identifier)
	d.Set("identifier", entity.Identifier)
	d.Set("kind", entity.Kind)
	d.Set("org_id", entity.OrgIdentifier)
	d.Set("project_id", entity.ProjectIdentifier)
	if v, ok := d.GetOk("yaml"); ok && v.(string) != "" {
		d.Set("yaml", v.(string))
	} else {
		d.Set("yaml", entity.Yaml)
	}
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

	yamlKind := yamlData["kind"].(string)
	if !strings.EqualFold(yamlKind, kind) {
		return catalogEntityInfo{}, fmt.Errorf("kind in YAML (%s) does not match kind parameter (%s)", yamlKind, kind)
	}

	yamlIdentifier := yamlData["identifier"].(string)
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

var entityImporter = &schema.ResourceImporter{
	State: func(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
		// Expected format: <scope>/<kind>/<identifier>
		// If account-level: <kind>/<identifier>
		// Scope examples: "org", "org.project"
		id := d.Id()
		parts := strings.Split(id, "/")

		if len(parts) < 2 || len(parts) > 3 {
			return nil, fmt.Errorf("invalid import ID format: %s. Expected: <scope>/<kind>/<identifier>", id)
		}

		var scope string
		var kind string
		var identifier string
		if len(parts) == 2 {
			scope = "account"
			kind = parts[0]
			identifier = parts[1]
		} else {
			scope = fmt.Sprintf("account.%s", parts[0])
			kind = parts[1]
			identifier = parts[2]
		}

		// Extract org and project from scope if present
		var orgId, projectId optional.String
		scopeParts := strings.Split(scope, ".")
		if len(scopeParts) > 1 {
			orgId = optional.NewString(scopeParts[0])
		}
		if len(scopeParts) > 2 {
			projectId = optional.NewString(scopeParts[1])
		}

		c, ctx := meta.(*internal.Session).GetIDPClientWithContext(context.Background())

		resp, _, err := c.EntitiesApi.GetEntity(ctx, scope, kind, identifier, &idp.EntitiesApiGetEntityOpts{
			OrgIdentifier:     orgId,
			ProjectIdentifier: projectId,
			HarnessAccount:    optional.NewString(c.AccountId),
		})
		if err != nil {
			return nil, fmt.Errorf("failed to fetch entity for import: %w", err)
		}

		readCatalogEntity(d, resp)

		return []*schema.ResourceData{d}, nil
	},
}
