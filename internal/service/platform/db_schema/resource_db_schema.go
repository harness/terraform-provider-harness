package dbschema

import (
	"context"
	"net/http"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

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
			"type": {
				Description:  "Type of the database schema. Valid values are: SCRIPT, REPOSITORY",
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
				Description: "If percona-toolkit is to be enabled for the database schema",
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
			},
			"schema_source": {
				Description:   "Provides a connector and path at which to find the database schema representation",
				Type:          schema.TypeList,
				MaxItems:      1,
				Optional:      true,
				ConflictsWith: []string{"changelog_script"},
				AtLeastOneOf:  []string{"schema_source", "changelog_script"},
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
						"archive_path": {
							Description: "If connector type is artifactory, path to the archive file which contains the changeLog",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"toml": {
							Description: "Config file, to define various settings and properties for managing database schema change",
							Type:        schema.TypeString,
							Optional:    true,
						},
					},
				},
			},
			"changelog_script": {
				Description:   "Configuration to clone changeSets using script",
				Type:          schema.TypeList,
				Optional:      true,
				MaxItems:      1,
				ConflictsWith: []string{"schema_source"},
				AtLeastOneOf:  []string{"schema_source", "changelog_script"},
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"image": {
							Description: "The fully-qualified name (FQN) of the image",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"command": {
							Description: "Script to clone changeSets",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"shell": {
							Description: "Type of the shell. For example Sh or Bash",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"location": {
							Description: "Path to changeLog file",
							Type:        schema.TypeString,
							Optional:    true,
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
	schemaIn := &dbops.DbSchemaIn{
		Identifier: d.Get("identifier").(string),
		Name:       d.Get("name").(string),
		Tags:       helpers.ExpandTags(d.Get("tags").(*schema.Set).List()),
		Service:    d.Get("service").(string),
	}
	if v, ok := d.GetOk("type"); ok {
		dbSchemaType := dbops.DbSchemaType(v.(string))
		schemaIn.Type_ = &dbSchemaType
	}

	if v, ok := d.GetOk("migration_type"); ok {
		migrationType := dbops.MigrationType(v.(string))
		schemaIn.MigrationType = &migrationType
	}

	if v, ok := d.GetOk("use_percona"); ok {
		schemaIn.UsePercona = v.(bool)
	}

	if _, ok := d.GetOk("changelog_script"); ok {
		changelogScript := &dbops.ChangeLogScript{
			Image:    d.Get("changelog_script.0.image").(string),
			Command:  d.Get("changelog_script.0.command").(string),
			Shell:    d.Get("changelog_script.0.shell").(string),
			Location: d.Get("changelog_script.0.location").(string),
			Toml:     d.Get("changelog_script.0.toml").(string),
		}
		schemaIn.ChangeLogScript = changelogScript
	}

	if _, ok := d.GetOk("schema_source"); ok {
		Changelog := &dbops.Changelog{
			Repo:        d.Get("schema_source.0.repo").(string),
			Connector:   d.Get("schema_source.0.connector").(string),
			Location:    d.Get("schema_source.0.location").(string),
			ArchivePath: d.Get("schema_source.0.archive_path").(string),
			Toml:        d.Get("schema_source.0.toml").(string),
		}
		schemaIn.Changelog = Changelog
	}

	return schemaIn
}

func readDBSchema(d *schema.ResourceData, dbSchema *dbops.DbSchemaOut) {
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
}
