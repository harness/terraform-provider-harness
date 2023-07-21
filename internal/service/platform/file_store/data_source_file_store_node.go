package file_store

import (
	"fmt"
	"strings"

	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceFileStoreNode() *schema.Resource {
	resource := &schema.Resource{
		Description: "Data source for retrieving files and folders.",

		ReadContext: resourceFileStoreNodeRead,

		Schema: map[string]*schema.Schema{
			"parent_identifier": {
				Description: "File or folder parent idnetifier",
				Type:        schema.TypeString,
				Required:    true,
			},
			"content": {
				Description: "File or folder content",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"mime_type": {
				Description: "File or folder mime type",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"type": {
				Description: fmt.Sprintf("The type of file. Valid options are %s", strings.Join(nextgen.NGFileTypeValues, ", ")),
				Type:        schema.TypeString,
				Required:    true,
			},
			"file_usage": {
				Description: fmt.Sprintf("File usage. Valid options are %s", strings.Join(nextgen.FileUsageValues, ", ")),
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
		},
	}

	helpers.SetMultiLevelResourceSchema(resource.Schema)

	return resource
}
