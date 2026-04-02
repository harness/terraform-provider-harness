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
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func DataSourceDBSchema() *schema.Resource {
	resource := &schema.Resource{
		Description: "Data source for retrieving a Harness DBDevOps Schema.",

		ReadContext: dataSourceDBSchemaRead,

		Schema: map[string]*schema.Schema{
			"service": {
				Description: "The service associated with schema",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"type": {
				Description:  "Type of the database schema. Valid values are: Repository, Script",
				Type:         schema.TypeString,
				Default:      string(dbops.REPOSITORY_DbSchemaType),
				ValidateFunc: validation.StringInSlice([]string{string(dbops.REPOSITORY_DbSchemaType), string(dbops.SCRIPT_DbSchemaType)}, false),
				Optional:     true,
			},
			"migration_type": {
				Description:  "DB Migration tool type. Valid values are: Liquibase, Flyway",
				Type:         schema.TypeString,
				Default:      string(dbops.LIQUIBASE_MigrationType),
				ValidateFunc: validation.StringInSlice([]string{string(dbops.LIQUIBASE_MigrationType), string(dbops.FLYWAY_MigrationType)}, false),
				Optional:     true,
			},
			"use_percona": {
				Description: "If percona-toolkit is enabled for the database schema",
				Type:        schema.TypeBool,
				Default:     false,
				Optional:    true,
			},
			"changelog_script": {
				Description: "Configuration to clone changeSets using script",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"image": {
							Description: "The fully-qualified name (FQN) of the image",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"command": {
							Description: "Script to clone changeSets",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"shell": {
							Description: "Type of the shell. For example Sh or Bash",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"location": {
							Description: "Path to changeLog file",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"toml": {
							Description: "Config file, to define various settings and properties for managing database schema change",
							Type:        schema.TypeString,
							Optional:    true,
						},
					},
				},
			},
			"schema_source": {
				Description: "Provides a connector and path at which to find the database schema representation",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"connector": {
							Description: "Connector to repository at which to find details about the database schema",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"location": {
							Description: "The path within the specified repository at which to find details about the database schema",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"repo": {
							Description: "If connector url is of account, which repository to connect to using the connector",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"archive_path": {
							Description: "If connector type is artifactory, path to the archive file which contains the changeLog",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"toml": {
							Description: "Config file, to define various settings and properties for managing database schema change",
							Type:        schema.TypeString,
							Optional:    true,
						},
					},
				},
			},
		},
	}

	helpers.SetProjectLevelDataSourceSchemaIdentifierRequired(resource.Schema)

	return resource
}

func dataSourceDBSchemaRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetDBOpsClientWithContext(ctx)

	var err error
	var dbSchema dbops.DbSchemaOut
	var httpResp *http.Response

	id := d.Get("identifier").(string)

	localVarOptionals := dbops.DatabaseSchemaApiV1GetProjDbSchemaOpts{
		HarnessAccount: optional.NewString(meta.(*internal.Session).AccountId),
	}
	dbSchema, httpResp, err = c.DatabaseSchemaApi.V1GetProjDbSchema(ctx, d.Get("org_id").(string), d.Get("project_id").(string), id, &localVarOptionals)

	if err != nil {
		return helpers.HandleDBOpsApiError(err, d, httpResp)
	}

	readDataSourceDBSchema(d, &dbSchema)

	return nil
}

func readDataSourceDBSchema(d *schema.ResourceData, dbSchema *dbops.DbSchemaOut) {
	d.SetId(dbSchema.Identifier)
	d.Set("identifier", dbSchema.Identifier)
	d.Set("name", dbSchema.Name)
	d.Set("tags", helpers.FlattenTags(dbSchema.Tags))
	d.Set("service", dbSchema.Service)

	if dbSchema.Type_ != nil {
		d.Set("type", string(*dbSchema.Type_))
	} else {
		d.Set("type", nil)
	}

	if dbSchema.MigrationType != nil {
		d.Set("migration_type", string(*dbSchema.MigrationType))
	} else {
		d.Set("migration_type", nil)
	}

	d.Set("use_percona", dbSchema.UsePercona)

	if dbSchema.ChangeLogScript != nil {
		d.Set("changelog_script.0.image", dbSchema.ChangeLogScript.Image)
		d.Set("changelog_script.0.command", dbSchema.ChangeLogScript.Command)
		d.Set("changelog_script.0.shell", dbSchema.ChangeLogScript.Shell)
		d.Set("changelog_script.0.location", dbSchema.ChangeLogScript.Location)

		if dbSchema.MigrationType != nil && *dbSchema.MigrationType == dbops.FLYWAY_MigrationType {
			d.Set("changelog_script.0.toml", dbSchema.ChangeLogScript.Toml)
		}

		d.Set("schema_source", nil)
	}

	if dbSchema.Changelog != nil {
		d.Set("schema_source.0.location", dbSchema.Changelog.Location)
		d.Set("schema_source.0.repo", dbSchema.Changelog.Repo)
		d.Set("schema_source.0.connector", dbSchema.Changelog.Connector)
		d.Set("schema_source.0.archive_path", dbSchema.Changelog.ArchivePath)

		if dbSchema.MigrationType != nil && *dbSchema.MigrationType == dbops.FLYWAY_MigrationType {
			d.Set("schema_source.0.toml", dbSchema.Changelog.Toml)
		}

		d.Set("changelog_script", nil)
	}
}
