package file_store

import (
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceFileStoreNodeFolder() *schema.Resource {
	resource := &schema.Resource{
		Description: "Data source for retrieving folders.",

		ReadContext: resourceFileStoreNodeFolderRead,

		Schema: map[string]*schema.Schema{
			"parent_identifier": {
				Description: "Folder parent identifier on Harness File Store",
				Type:        schema.TypeString,
				Optional:    false,
				Computed:    true,
			},
			"path": {
				Description: "Harness File Store folder path",
				Type:        schema.TypeString,
				Optional:    false,
				Computed:    true,
			},
			"created_by": {
				Description: "Created by",
				Type:        schema.TypeList,
				Optional:    false,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"email": {
							Description: "User email",
							Type:        schema.TypeString,
							Optional:    false,
							Computed:    true,
						},
						"name": {
							Description: "User name",
							Type:        schema.TypeString,
							Optional:    false,
							Computed:    true,
						},
					},
				},
			},
			"last_modified_by": {
				Description: "Last modified by",
				Type:        schema.TypeList,
				Optional:    false,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"email": {
							Description: "User email",
							Type:        schema.TypeString,
							Optional:    false,
							Computed:    true,
						},
						"name": {
							Description: "User name",
							Type:        schema.TypeString,
							Optional:    false,
							Computed:    true,
						},
					},
				},
			},
			"last_modified_at": {
				Description: "Last modified at",
				Type:        schema.TypeInt,
				Optional:    false,
				Computed:    true,
			},
		},
	}

	helpers.SetMultiLevelDatasourceSchemaIdentifierRequired(resource.Schema)

	return resource
}
