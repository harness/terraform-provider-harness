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

func DataSourceDBInstance() *schema.Resource {
	resource := &schema.Resource{
		Description: "Data source for retrieving a Harness DBDevOps Instance.",

		ReadContext: dataSourceDBInstanceRead,

		Schema: map[string]*schema.Schema{
			"schema": {
				Description: "The reference to schema",
				Type:        schema.TypeString,
				Required:    true,
			},
			"branch": {
				Description: "The branch of changeSet repository",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"connector": {
				Description: "The connector to database",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"context": {
				Description: "The liquibase context",
				Type:        schema.TypeString,
				Computed:    true,
			},
		},
	}

	helpers.SetProjectLevelDataSourceSchemaIdentifierRequired(resource.Schema)

	return resource
}

func dataSourceDBInstanceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetDBOpsClientWithContext(ctx)

	var err error
	var dbInstance dbops.DbInstanceOut
	var httpResp *http.Response

	id := d.Get("identifier").(string)

	localVarOptionals := dbops.DatabaseInstanceApiV1GetProjDbSchemaInstanceOpts{
		HarnessAccount: optional.NewString(meta.(*internal.Session).AccountId),
	}
	dbInstance, httpResp, err = c.DatabaseInstanceApi.V1GetProjDbSchemaInstance(ctx, d.Get("org_id").(string), d.Get("project_id").(string), d.Get("schema").(string), id, &localVarOptionals)

	if err != nil {
		return helpers.HandleDBOpsApiError(err, d, httpResp)
	}

	readDataSourceDBInstance(d, &dbInstance)

	return nil
}

func readDataSourceDBInstance(d *schema.ResourceData, dbInstance *dbops.DbInstanceOut) {
	d.SetId(dbInstance.Identifier)
	d.Set("identifier", dbInstance.Identifier)
	d.Set("name", dbInstance.Name)
	d.Set("tags", helpers.FlattenTags(dbInstance.Tags))
	d.Set("branch", dbInstance.Branch)
	d.Set("connector", dbInstance.Connector)
	d.Set("context", dbInstance.Context)
}
