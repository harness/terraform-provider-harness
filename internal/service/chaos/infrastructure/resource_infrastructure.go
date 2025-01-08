package infrastructure

import (
	"context"
	"log"

	"github.com/harness/harness-go-sdk/harness/chaos"
	hh "github.com/harness/harness-go-sdk/harness/helpers"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceChaosInfrastructure() *schema.Resource {
	resource := &schema.Resource{
		Description: "Resource for creating a Chaos Infrastructure.",

		ReadContext:   resourceChaosInfrastructureRead,
		UpdateContext: resourceChaosInfrastructureUpdate,
		DeleteContext: resourceChaosInfrastructureDelete,
		CreateContext: resourceChaosInfrastructureCreate,
		Importer:      helpers.MultiLevelResourceImporter,

		Schema: map[string]*schema.Schema{
			"org_id": {
				Description: "Identifier of the organization in which the chaos infrastructure is configured.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"project_id": {
				Description: "Identifier of the project in which the chaos infrastructure is configured.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"identifier": {
				Description: "Identifier of the chaos infrastructure.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"name": {
				Description: "Name of the chaos infrastructure.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"description": {
				Description: "Description of the chaos infrastructure.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"environment_id": {
				Description: "Environment ID of the chaos infrastructure.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"namespace": {
				Description: "Namespace of the chaos infrastructure.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"service_account": {
				Description: "Service Account of the chaos infrastructure.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"tags": {
				Description: "Tags of the chaos infrastructure.",
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}

	helpers.SetMultiLevelResourceSchema(resource.Schema)

	return resource
}

func resourceChaosInfrastructureRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetChaosClientWithContext(ctx)
	var accountIdentifier, orgIdentifier, projectIdentifier, identifier, envIdentifier string
	accountIdentifier = c.AccountId
	if attr, ok := d.GetOk("org_id"); ok {
		orgIdentifier = attr.(string)
	}
	if attr, ok := d.GetOk("project_id"); ok {
		projectIdentifier = attr.(string)
	}
	if attr, ok := d.GetOk("identifier"); ok {
		identifier = attr.(string)
	}
	if attr, ok := d.GetOk("environment_id"); ok {
		envIdentifier = attr.(string)
	}
	resp, httpResp, err := c.ChaosSdkApi.GetInfraV2(ctx, identifier, accountIdentifier, orgIdentifier, projectIdentifier, envIdentifier)

	if err != nil {
		if err.Error() == "404 Not Found" {
			d.SetId("")
			d.MarkNewResource()
			return nil
		}
		return helpers.HandleReadApiError(err, d, httpResp)
	}
	readChaosInfrastructure(d, resp)

	return nil
}

func resourceChaosInfrastructureUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetChaosClientWithContext(ctx)
	var accountIdentifier, orgIdentifier, projectIdentifier string
	accountIdentifier = c.AccountId
	if attr, ok := d.GetOk("org_id"); ok {
		orgIdentifier = attr.(string)
	}
	if attr, ok := d.GetOk("project_id"); ok {
		projectIdentifier = attr.(string)
	}

	updateInfraRequest := buildUpdateInfraRequest(d)
	_, httpRespUpdate, errUpdate := c.ChaosSdkApi.UpdateInfraV2(ctx, *updateInfraRequest, accountIdentifier, orgIdentifier, projectIdentifier, &chaos.ChaosSdkApiUpdateInfraV2Opts{})

	if errUpdate != nil {
		return helpers.HandleApiError(errUpdate, d, httpRespUpdate)
	}
	readInfraUpdate(d, updateInfraRequest.Identity, updateInfraRequest.EnvironmentID)
	return nil
}

func buildUpdateInfraRequest(d *schema.ResourceData) *chaos.InfraV2UpdateKubernetesInfrastructureV2Request {
	infraReq := &chaos.InfraV2UpdateKubernetesInfrastructureV2Request{}

	if attr, ok := d.GetOk("identifier"); ok {
		infraReq.Identity = attr.(string)
	}

	infraReq.Name = d.Get("name").(string)
	infraReq.Description = d.Get("description").(string)

	tags := []string{}
	for _, t := range d.Get("tags").(*schema.Set).List() {
		tagStr := t.(string)
		tags = append(tags, tagStr)
	}
	infraReq.Tags = tags
	infraReq.InfraNamespace = d.Get("namespace").(string)
	if attr, ok := d.GetOk("environment_id"); ok {
		infraReq.EnvironmentID = attr.(string)
	}
	if attr, ok := d.GetOk("service_account"); ok {
		infraReq.ServiceAccount = attr.(string)
	}

	return infraReq
}

func readInfraUpdate(d *schema.ResourceData, identifier string, envID string) {

	// can we use corr ID or do we need to return identity?
	d.SetId(identifier)
	d.Set("environment_id", envID)

}

func resourceChaosInfrastructureDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetChaosClientWithContext(ctx)

	var accountIdentifier, orgIdentifier, projectIdentifier, identifier, environmentID string
	accountIdentifier = c.AccountId
	if attr, ok := d.GetOk("org_id"); ok {
		orgIdentifier = attr.(string)
	}
	if attr, ok := d.GetOk("project_id"); ok {
		projectIdentifier = attr.(string)
	}
	if attr, ok := d.GetOk("identifier"); ok {
		identifier = attr.(string)
	}
	if attr, ok := d.GetOk("environment_id"); ok {
		environmentID = attr.(string)
	}
	_, httpResp, err := c.ChaosSdkApi.DeleteInfraV2(ctx, identifier, environmentID, accountIdentifier, orgIdentifier, projectIdentifier)

	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	return nil
}

func readInfraCreate(d *schema.ResourceData, infraResponse *chaos.InfraV2RegisterInfrastructureV2Response, envID string) {
	d.SetId(infraResponse.Identity)
	d.Set("name", infraResponse.Name)
	d.Set("environment_id", envID)
}

func resourceChaosInfrastructureCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetChaosClientWithContext(ctx)
	log.Printf("Debugging message: %v", c)

	ctx = context.WithValue(ctx, chaos.ContextAccessToken, hh.EnvVars.BearerToken.Get()) // do we need it
	var accountIdentifier, orgIdentifier, projectIdentifier string
	createInfraRequest := buildInfraCreateRequest(d)
	identifiers := chaos.InfraV2Identifiers{}
	createInfraRequest.Identifier = &identifiers
	createInfraRequest.Identifier.AccountIdentifier = c.AccountId
	accountIdentifier = c.AccountId
	if attr, ok := d.GetOk("org_id"); ok {
		orgIdentifier = attr.(string)
		createInfraRequest.Identifier.OrgIdentifier = attr.(string)
	}
	if attr, ok := d.GetOk("project_id"); ok {
		projectIdentifier = attr.(string)
		createInfraRequest.Identifier.ProjectIdentifier = attr.(string)
	}

	resp, httpRespUpdate, errUpdate := c.ChaosSdkApi.RegisterInfraV2(ctx, createInfraRequest, accountIdentifier, orgIdentifier, projectIdentifier, &chaos.ChaosSdkApiRegisterInfraV2Opts{})
	if errUpdate != nil {
		return helpers.HandleApiError(errUpdate, d, httpRespUpdate)
	}
	readInfraCreate(d, &resp, createInfraRequest.EnvironmentID)
	defer httpRespUpdate.Body.Close()
	return nil
}

func buildInfraCreateRequest(d *schema.ResourceData) chaos.InfraV2RegisterInfrastructureV2Request {
	infraReq := chaos.InfraV2RegisterInfrastructureV2Request{}

	if attr, ok := d.GetOk("identifier"); ok {
		infraReq.Identity = attr.(string)
		infraReq.InfraID = attr.(string)
	}

	infraReq.Name = d.Get("name").(string)
	infraReq.Description = d.Get("description").(string)

	tags := []string{}
	for _, t := range d.Get("tags").(*schema.Set).List() {
		tagStr := t.(string)
		tags = append(tags, tagStr)
	}
	infraReq.Tags = tags

	infraReq.InfraNamespace = d.Get("namespace").(string)
	if attr, ok := d.GetOk("environment_id"); ok {
		infraReq.EnvironmentID = attr.(string)
	}
	if attr, ok := d.GetOk("service_account"); ok {
		infraReq.ServiceAccount = attr.(string)
	}
	scope := chaos.CLUSTER_InfraV2InfraScope
	infraReq.InfraScope = &scope
	infraType := chaos.KUBERNETES_InfraV2InfraType
	infraReq.InfraType = &infraType

	return infraReq
}

func readChaosInfrastructure(d *schema.ResourceData, infra chaos.InfraV2KubernetesInfrastructureV2Details) {
	d.SetId(infra.Identity)
	d.Set("org_id", infra.Identifier.OrgIdentifier)
	d.Set("project_id", infra.Identifier.ProjectIdentifier)
	d.Set("identifier", infra.Identity)
	d.Set("name", infra.Name)
	d.Set("description", infra.Description)
	d.Set("infra_id", infra.InfraID)
	d.Set("environment_id", infra.EnvironmentID)
}
