package service_account

import (
	"context"
	"net/http"

	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceServiceAccount() *schema.Resource {
	resource := &schema.Resource{
		Description:   "Resource for creating service account.",
		ReadContext:   resourceServiceAccountRead,
		UpdateContext: resourceServiceAccountCreateOrUpdate,
		CreateContext: resourceServiceAccountCreateOrUpdate,
		DeleteContext: resourceServiceAccountDelete,
		Importer:      helpers.MultiLevelResourceImporter,

		Schema: map[string]*schema.Schema{
			"email": {
				Description: "Email of the Service Account.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"account_id": {
				Description: "Account Identifier for the Entity.",
				Type:        schema.TypeString,
				Optional:    true,
			},
		},
	}

	helpers.SetMultiLevelResourceSchema(resource.Schema)

	return resource
}

func resourceServiceAccountRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	id := d.Get("identifier").(string)

	resp, httpResp, err := c.ServiceAccountApi.GetAggregatedServiceAccount(ctx, c.AccountId, id, &nextgen.ServiceAccountApiGetAggregatedServiceAccountOpts{
		OrgIdentifier:     helpers.BuildField(d, "org_id"),
		ProjectIdentifier: helpers.BuildField(d, "project_id"),
	})

	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	if resp.Data.ServiceAccount == nil {
		return nil
	}

	readServiceAccount(d, resp.Data.ServiceAccount)

	return nil
}

func resourceServiceAccountCreateOrUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	var err error
	var resp nextgen.ResponseDtoServiceAccount
	var httpResp *http.Response
	id := d.Id()

	serviceAccount := buildServiceAccount(d)

	if id == "" {
		resp, httpResp, err = c.ServiceAccountApi.CreateServiceAccount(ctx, *serviceAccount, c.AccountId, &nextgen.ServiceAccountApiCreateServiceAccountOpts{
			OrgIdentifier:     helpers.BuildField(d, "org_id"),
			ProjectIdentifier: helpers.BuildField(d, "project_id"),
		})
	} else {
		resp, httpResp, err = c.ServiceAccountApi.UpdateServiceAccount(ctx, *serviceAccount, c.AccountId, id, &nextgen.ServiceAccountApiUpdateServiceAccountOpts{
			OrgIdentifier:     helpers.BuildField(d, "org_id"),
			ProjectIdentifier: helpers.BuildField(d, "project_id"),
		})
	}

	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}
	readServiceAccount(d, resp.Data)
	return nil
}

func resourceServiceAccountDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	_, httpResp, err := c.ServiceAccountApi.DeleteServiceAccount(ctx, c.AccountId, d.Id(), &nextgen.ServiceAccountApiDeleteServiceAccountOpts{
		OrgIdentifier:     helpers.BuildField(d, "org_id"),
		ProjectIdentifier: helpers.BuildField(d, "project_id"),
	})

	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	return nil
}

func buildServiceAccount(d *schema.ResourceData) *nextgen.ServiceAccount {
	return &nextgen.ServiceAccount{
		Identifier:        d.Get("identifier").(string),
		Name:              d.Get("name").(string),
		Description:       d.Get("description").(string),
		Tags:              helpers.ExpandTags(d.Get("tags").(*schema.Set).List()),
		Email:             d.Get("email").(string),
		AccountIdentifier: d.Get("account_id").(string),
		OrgIdentifier:     d.Get("org_id").(string),
		ProjectIdentifier: d.Get("project_id").(string),
	}
}

func readServiceAccount(d *schema.ResourceData, serviceAccount *nextgen.ServiceAccount) {
	d.SetId(serviceAccount.Identifier)
	d.Set("identifier", serviceAccount.Identifier)
	d.Set("name", serviceAccount.Name)
	d.Set("description", serviceAccount.Description)
	d.Set("tags", helpers.FlattenTags(serviceAccount.Tags))
	d.Set("email", serviceAccount.Email)
	d.Set("account_id", serviceAccount.AccountIdentifier)
	d.Set("org_id", serviceAccount.OrgIdentifier)
	d.Set("project_id", serviceAccount.ProjectIdentifier)
}
