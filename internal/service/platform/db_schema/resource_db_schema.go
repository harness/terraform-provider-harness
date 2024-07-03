package dbschema

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

func ResourceDBSchema() *schema.Resource {
	resource := &schema.Resource{
		Description: "Resource for creating a Harness DBDevOps Schema.",

		ReadContext:   resourceDBSchemaRead,
		UpdateContext: resourceDBSchemaCreateOrUpdate,
		DeleteContext: resourceDBSchemaDelete,
		CreateContext: resourceDBSchemaCreateOrUpdate,
		Importer:      helpers.ProjectResourceImporter,

		Schema: map[string]*schema.Schema{
			"service": {
				Description: "The service associated with schema",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"schema_source": {
				Description: "Provides a connector and path at which to find the database schema representation",
				Type:        schema.TypeList,
				MaxItems:    1,
				Required:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"connector": {
							Description: "Connector to repository at which to find details about the database schema",
							Type:        schema.TypeString,
							Required:    true,
						},
						"location": {
							Description: "The path within the specified repository at which to find details about the database schema",
							Type:        schema.TypeString,
							Required:    true,
						},
						"repo": {
							Description: "If connector url is of account, which repository to connect to using the connector",
							Type:        schema.TypeString,
							Optional:    true,
						},
					},
				},
			},
		},
	}

	helpers.SetProjectLevelResourceSchema(resource.Schema)

	return resource
}

func resourceDBSchemaRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetDBOpsClientWithContext(ctx)

	localVarOptionals := dbops.DatabaseSchemaApiV1GetProjDbSchemaOpts{
		HarnessAccount: optional.NewString(meta.(*internal.Session).AccountId),
	}

	resp, httpResp, err := c.DatabaseSchemaApi.V1GetProjDbSchema(ctx, d.Get("org_id").(string), d.Get("project_id").(string), d.Id(), &localVarOptionals)

	if err != nil {
		return helpers.HandleDBOpsReadApiError(err, d, httpResp)
	}

	readDBSchema(d, &resp)

	return nil
}

func resourceDBSchemaCreateOrUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetDBOpsClientWithContext(ctx)

	var err error
	var resp dbops.DbSchemaOut
	var httpResp *http.Response
	id := d.Id()
	dbSchema := buildDbSchema(d)

	if id == "" {
		localVarOptionals := dbops.DatabaseSchemaApiV1CreateProjDbSchemaOpts{
			HarnessAccount: optional.NewString(meta.(*internal.Session).AccountId),
		}
		resp, httpResp, err = c.DatabaseSchemaApi.V1CreateProjDbSchema(ctx, *dbSchema, d.Get("org_id").(string), d.Get("project_id").(string), &localVarOptionals)
	} else {
		localVarOptionals := dbops.DatabaseSchemaApiV1UpdateProjDbSchemaOpts{
			HarnessAccount: optional.NewString(meta.(*internal.Session).AccountId),
		}
		resp, httpResp, err = c.DatabaseSchemaApi.V1UpdateProjDbSchema(ctx, dbSchema, d.Get("org_id").(string), d.Get("project_id").(string), d.Id(), &localVarOptionals)
	}

	if err != nil {
		return helpers.HandleDBOpsApiError(err, d, httpResp)
	}

	readDBSchema(d, &resp)

	return nil
}

func resourceDBSchemaDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetDBOpsClientWithContext(ctx)

	localVarOptionals := dbops.DatabaseSchemaApiV1DeleteProjDbSchemaOpts{
		HarnessAccount: optional.NewString(meta.(*internal.Session).AccountId),
	}

	httpResp, err := c.DatabaseSchemaApi.V1DeleteProjDbSchema(ctx, d.Get("org_id").(string), d.Get("project_id").(string), d.Id(), &localVarOptionals)
	if err != nil {
		return helpers.HandleDBOpsApiError(err, d, httpResp)
	}

	return nil
}

func buildDbSchema(d *schema.ResourceData) *dbops.DbSchemaIn {
	return &dbops.DbSchemaIn{
		Identifier: d.Get("identifier").(string),
		Name:       d.Get("name").(string),
		Tags:       helpers.ExpandTags(d.Get("tags").(*schema.Set).List()),
		Service:    d.Get("service").(string),
		Changelog: &dbops.Changelog{
			Repo:      d.Get("schema_source.0.repo").(string),
			Connector: d.Get("schema_source.0.connector").(string),
			Location:  d.Get("schema_source.0.location").(string),
		},
	}
}

func readDBSchema(d *schema.ResourceData, dbSchema *dbops.DbSchemaOut) {
	d.SetId(dbSchema.Identifier)
	d.Set("identifier", dbSchema.Identifier)
	d.Set("name", dbSchema.Name)
	d.Set("tags", helpers.FlattenTags(dbSchema.Tags))
	d.Set("service", dbSchema.Service)
	d.Set("schema_source.0.location", dbSchema.Changelog.Location)
	d.Set("schema_source.0.repo", dbSchema.Changelog.Repo)
	d.Set("schema_source.0.connector", dbSchema.Changelog.Connector)
}
