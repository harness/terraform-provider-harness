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
				Required:    true,
			},
			"path": {
				Description: "Harness File Store folder path",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"created_by": {
				Description: "Created by",
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"email": {
							Description: "User email",
							Type:        schema.TypeString,
							Required:    true,
						},
						"name": {
							Description: "User name",
							Type:        schema.TypeString,
							Required:    true,
						},
					},
				},
			},
			"last_modified_by": {
				Description: "Last modified by",
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"email": {
							Description: "User email",
							Type:        schema.TypeString,
							Required:    true,
						},
						"name": {
							Description: "User name",
							Type:        schema.TypeString,
							Required:    true,
						},
					},
				},
			},
			"last_modified_at": {
				Description: "Last modified at",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
		},
	}

	helpers.SetMultiLevelResourceSchema(resource.Schema)

	return resource
}
