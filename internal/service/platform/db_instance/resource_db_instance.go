package dbinstance

import (
	"context"
	"net/http"

	"github.com/antihax/optional"
	"github.com/harness/harness-go-sdk/harness/dbops"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceDBInstance() *schema.Resource {
	resource := &schema.Resource{
		Description: "Resource for creating a Harness DBDevOps Instance.",

		ReadContext:   resourceDBInstanceRead,
		UpdateContext: resourceDBInstanceCreateOrUpdate,
		DeleteContext: resourceDBInstanceDelete,
		CreateContext: resourceDBInstanceCreateOrUpdate,
		Importer:      helpers.DBInstanceResourceImporter,

		Schema: map[string]*schema.Schema{
			"schema": {
				Description: "The reference to schema",
				Type:        schema.TypeString,
				Required:    true,
			},
			"branch": {
				Description: "The branch of changeSet repository",
				Type:        schema.TypeString,
				Required:    true,
			},
			"connector": {
				Description: "The connector to database",
				Type:        schema.TypeString,
				Required:    true,
			},
			"context": {
				Description: "The liquibase context",
				Type:        schema.TypeString,
				Optional:    true,
			},
		},
	}

	helpers.SetProjectLevelResourceSchema(resource.Schema)

	return resource
}

func resourceDBInstanceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetDBOpsClientWithContext(ctx)

	localVarOptionals := dbops.DatabaseInstanceApiV1GetProjDbSchemaInstanceOpts{
		HarnessAccount: optional.NewString(meta.(*internal.Session).AccountId),
	}

	resp, httpResp, err := c.DatabaseInstanceApi.V1GetProjDbSchemaInstance(ctx, d.Get("org_id").(string), d.Get("project_id").(string), d.Get("schema").(string), d.Id(), &localVarOptionals)

	if err != nil {
		return helpers.HandleDBOpsReadApiError(err, d, httpResp)
	}

	readDBInstance(d, &resp)

	return nil
}

func resourceDBInstanceCreateOrUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetDBOpsClientWithContext(ctx)

	var err error
	var resp dbops.DbInstanceOut
	var httpResp *http.Response
	id := d.Id()
	dbInstance := buildDbInstance(d)

	if id == "" {
		localVarOptionals := dbops.DatabaseInstanceApiV1CreateProjDbSchemaInstanceOpts{
			HarnessAccount: optional.NewString(meta.(*internal.Session).AccountId),
		}
		resp, httpResp, err = c.DatabaseInstanceApi.V1CreateProjDbSchemaInstance(ctx, *dbInstance, d.Get("org_id").(string), d.Get("project_id").(string), d.Get("schema").(string), &localVarOptionals)
	} else {
		localVarOptionals := dbops.DatabaseSchemaApiV1UpdateProjDbSchemaInstanceOpts{
			HarnessAccount: optional.NewString(meta.(*internal.Session).AccountId),
		}
		resp, httpResp, err = c.DatabaseSchemaApi.V1UpdateProjDbSchemaInstance(ctx, dbInstance, d.Get("org_id").(string), d.Get("project_id").(string), d.Get("schema").(string), d.Id(), &localVarOptionals)
	}

	if err != nil {
		return helpers.HandleDBOpsApiError(err, d, httpResp)
	}

	readDBInstance(d, &resp)

	return nil
}

func resourceDBInstanceDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetDBOpsClientWithContext(ctx)

	localVarOptionals := dbops.DatabaseInstanceApiV1DeleteProjDbSchemaInstanceOpts{
		HarnessAccount: optional.NewString(meta.(*internal.Session).AccountId),
	}

	httpResp, err := c.DatabaseInstanceApi.V1DeleteProjDbSchemaInstance(ctx, d.Get("org_id").(string), d.Get("project_id").(string), d.Get("schema").(string), d.Id(), &localVarOptionals)
	if err != nil {
		return helpers.HandleDBOpsApiError(err, d, httpResp)
	}

	return nil
}

func buildDbInstance(d *schema.ResourceData) *dbops.DbInstanceIn {
	return &dbops.DbInstanceIn{
		Identifier: d.Get("identifier").(string),
		Name:       d.Get("name").(string),
		Tags:       helpers.ExpandTags(d.Get("tags").(*schema.Set).List()),
		Branch:     d.Get("branch").(string),
		Connector:  d.Get("connector").(string),
		Context:    d.Get("context").(string),
	}
}

func readDBInstance(d *schema.ResourceData, dbInstance *dbops.DbInstanceOut) {
	d.SetId(dbInstance.Identifier)
	d.Set("identifier", dbInstance.Identifier)
	d.Set("name", dbInstance.Name)
	d.Set("tags", helpers.FlattenTags(dbInstance.Tags))
	d.Set("branch", dbInstance.Branch)
	d.Set("connector", dbInstance.Connector)
	d.Set("context", dbInstance.Context)
}
