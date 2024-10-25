package file_store

import (
	"fmt"
	"strings"

	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceFileStoreNodeFile() *schema.Resource {
	resource := &schema.Resource{
		Description: "Data source for retrieving files.",

		ReadContext: resourceFileStoreNodeFileRead,

		Schema: map[string]*schema.Schema{
			"parent_identifier": {
				Description: "File parent identifier on Harness File Store",
				Type:        schema.TypeString,
				Optional:    false,
				Computed:    true,
			},
			"file_content_path": {
				Description: "File content path to be upladed on Harness File Store",
				Type:        schema.TypeString,
				Optional:    false,
				Computed:    true,
			},
			"mime_type": {
				Description: "File mime type",
				Type:        schema.TypeString,
				Optional:    false,
				Computed:    true,
			},
			"file_usage": {
				Description: fmt.Sprintf("File usage. Valid options are %s", strings.Join(nextgen.FileUsageValues, ", ")),
				Type:        schema.TypeString,
				Optional:    false,
				Computed:    true,
			},
			"content": {
				Description: "File content stored on Harness File Store",
				Type:        schema.TypeString,
				Optional:    false,
				Computed:    true,
			},
			"path": {
				Description: "Harness File Store file path",
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
