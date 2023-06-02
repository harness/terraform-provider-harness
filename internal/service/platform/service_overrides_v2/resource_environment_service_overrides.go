package service_overrides_v2

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/antihax/optional"
	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceServiceOverrides() *schema.Resource {
	resource := &schema.Resource{
		Description: "Resource for creating a Harness service override V2.",

		ReadContext:   resourceServiceOverridesV2Read,
		UpdateContext: resourceServiceOverridesV2CreateOrUpdate,
		DeleteContext: resourceServiceOverridesV2Delete,
		CreateContext: resourceServiceOverridesV2CreateOrUpdate,
		Importer:      helpers.ServiceOverrideResourceImporter,

		Schema: map[string]*schema.Schema{
			"service_id": {
				Description: "The service ID to which the overrides applies.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"env_id": {
				Description: "The env ID to which the overrides are associated.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"infra_id": {
				Description: "The infrastructure ID to which the overrides are associated",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"cluster_id": {
				Description: "The cluster ID to which the overrides are associated",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"type": {
				Description: "The type of the overrides",
				Type:        schema.TypeString,
				Required:    true,
			},
			"spec": {
				Description: "spec of the override values",
				Type:        schema.TypeString,
				Required:    true,
			},
		},
	}

	SetScopedResourceSchemaForServiceOverride(resource.Schema)

	return resource
}

func resourceServiceOverridesV2Read(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	resp, httpResp, err := c.ServiceOverridesApi.GetServiceOverrideListV2(ctx, c.AccountId,
		&nextgen.ServiceOverridesApiGetServiceOverrideListV2Opts{
			OrgIdentifier:     helpers.BuildField(d, "org_id"),
			ProjectIdentifier: helpers.BuildField(d, "project_id"),
		})
	if err != nil {
		return helpers.HandleReadApiError(err, d, httpResp)
	}

	readServiceOverridesV2List(d, resp.Data)

	return nil
}

func resourceServiceOverridesV2CreateOrUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	var err error
	var resp nextgen.ResponseServiceOverridesResponseDtov2
	var httpResp *http.Response
	env := buildServiceOverrideV2(d)

	resp, httpResp, err = c.ServiceOverridesApi.UpsertServiceOverrideV2(ctx, c.AccountId, &nextgen.ServiceOverridesApiUpsertServiceOverrideV2Opts{
		Body: optional.NewInterface(env),
	})

	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	// Soft delete lookup error handling
	// https://harness.atlassian.net/browse/PL-23765
	if &resp == nil || resp.Data == nil {
		d.SetId("")
		d.MarkNewResource()
		return nil
	}

	readServiceOverridesV2(d, resp.Data)

	return nil
}

func resourceServiceOverridesV2Delete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	_, httpResp, err := c.ServiceOverridesApi.DeleteServiceOverrideV2(ctx, d.Get("identifier").(string), c.AccountId, &nextgen.ServiceOverridesApiDeleteServiceOverrideV2Opts{
		OrgIdentifier:     helpers.BuildField(d, "org_id"),
		ProjectIdentifier: helpers.BuildField(d, "project_id"),
	})

	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	return nil
}

func buildServiceOverrideV2(d *schema.ResourceData) *nextgen.ServiceOverrideRequestDtov2 {
	// Create a variable of the struct type
	var serviceOverrides *nextgen.ServiceOverridesSpec
	str := d.Get("spec").(string)

	// Unmarshal the string into the struct
	err := json.Unmarshal([]byte(str), &serviceOverrides)
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}

	return &nextgen.ServiceOverrideRequestDtov2{
		Identifier:        d.Get("identifier").(string),
		OrgIdentifier:     d.Get("org_id").(string),
		ProjectIdentifier: d.Get("project_id").(string),
		EnvironmentRef:    d.Get("env_id").(string),
		ServiceRef:        d.Get("service_id").(string),
		InfraIdentifier:   d.Get("infra_id").(string),
		ClusterIdentifier: d.Get("cluster_id").(string),
		Type_:             d.Get("type").(string),
		Spec:              serviceOverrides,
	}
}

func readServiceOverridesV2List(d *schema.ResourceData, ps *nextgen.PageServiceOverridesResponseDtov2) {
	ServiceOverrideList := ps.Content
	for _, value := range ServiceOverrideList {
		readServiceOverridesV2(d, &value)
	}
}

func readServiceOverridesV2(d *schema.ResourceData, so *nextgen.ServiceOverridesResponseDtov2) {

	serviceOverrides := *so.Spec

	str, err := serviceOverridesSpecToString(serviceOverrides)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(str)

	d.SetId(so.Identifier)
	d.Set("identifier", so.Identifier)
	d.Set("org_id", so.OrgIdentifier)
	d.Set("project_id", so.ProjectIdentifier)
	d.Set("env_id", so.EnvironmentRef)
	d.Set("service_id", so.ServiceRef)
	d.Set("account_id", so.AccountId)
	d.Set("infra_id", so.InfraIdentifier)
	d.Set("cluster_id", so.ClusterIdentifier)
	d.Set("type", so.Type_)
	d.Set("spec", str)
	d.Set("newly_created", so.NewlyCreated)
}

func serviceOverridesSpecToString(serviceOverrides nextgen.ServiceOverridesSpec) (string, error) {
	bytes, err := json.Marshal(serviceOverrides)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func SetScopedResourceSchemaForServiceOverride(s map[string]*schema.Schema) {
	s["project_id"] = helpers.GetProjectIdSchema(helpers.SchemaFlagTypes.Optional)
	s["org_id"] = helpers.GetOrgIdSchema(helpers.SchemaFlagTypes.Optional)
	s["identifier"] = helpers.GetIdentifierSchema(helpers.SchemaFlagTypes.Required)
}
