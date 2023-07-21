package file_store

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/antihax/optional"
	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	orgId            = "org_id"
	projectId        = "project_id"
	identifier       = "identifier"
	name             = "name"
	parentIdentifier = "parent_identifier"
	path             = "path"
	content          = "content"
	description      = "description"
	mimeType         = "mime_type"
	tags             = "tags"
	type_            = "type"
	fileUsage        = "file_usage"
	createdBy        = "created_by"
	lastModifiedBy   = "last_modified_by"
	lastModifiedAt   = "last_modified_at"
	draft            = "draft"
)

func ResourceFileStoreNode() *schema.Resource {
	resource := &schema.Resource{
		Description: "Resource for creating files and folders in Harness.",

		ReadContext:   resourceFileStoreNodeRead,
		UpdateContext: resourceFileStoreNodeCreateOrUpdate,
		CreateContext: resourceFileStoreNodeCreateOrUpdate,
		DeleteContext: resourceFileStoreNodeDelete,
		Importer:      helpers.MultiLevelResourceImporter,

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

func resourceFileStoreNodeRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)
	id := d.Get(identifier).(string)

	var contentStr string

	resp, httpResp, err := c.FileStoreApi.GetFile(ctx, id, c.AccountId, &nextgen.FileStoreApiGetFileOpts{
		OrgIdentifier:     helpers.BuildField(d, orgId),
		ProjectIdentifier: helpers.BuildField(d, projectId),
	})

	if err != nil {
		return helpers.HandleReadApiError(err, d, httpResp)
	}

	if resp.Data == nil {
		return nil
	}

	if resp.Data.Type_ == nextgen.NGFileTypes.File.String() {
		resp, err := c.FileStoreApi.DownloadFile(ctx, id, c.AccountId, &nextgen.FileStoreApiDownloadFileOpts{
			OrgIdentifier:     helpers.BuildField(d, orgId),
			ProjectIdentifier: helpers.BuildField(d, projectId),
		})

		if err != nil {
			return helpers.HandleReadApiError(err, d, httpResp)
		}

		content, err := io.ReadAll(resp.Body)

		if err != nil {
			return helpers.HandleReadApiError(err, d, httpResp)
		}

		contentStr = string(content)
	}

	readFileNode(d, resp.Data, contentStr)

	return nil
}

func resourceFileStoreNodeCreateOrUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)
	id := d.Id()

	var resp nextgen.ResponseDtoFile
	var httpResp *http.Response
	var err error

	if id == "" {
		resp, httpResp, err = c.FileStoreApi.Create(ctx, c.AccountId, buildFileStoreApiCreateRequest(d))
	} else {
		resp, httpResp, err = c.FileStoreApi.Update(ctx, c.AccountId, id, buildFileStoreApiUpdateRequest(d))
	}

	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	readFileNode(d, resp.Data, d.Get(content).(string))

	return nil
}

func resourceFileStoreNodeDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	_, httpResp, err := c.FileStoreApi.DeleteFile(ctx, c.AccountId, d.Id(), &nextgen.FileStoreApiDeleteFileOpts{
		OrgIdentifier:     helpers.BuildField(d, orgId),
		ProjectIdentifier: helpers.BuildField(d, projectId),
	})

	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	return nil
}

func buildFileStoreApiCreateRequest(d *schema.ResourceData) *nextgen.FileStoreApiCreateOpts {
	create := &nextgen.FileStoreApiCreateOpts{
		OrgIdentifier:     getOptionalString(d.Get(orgId)),
		ProjectIdentifier: getOptionalString(d.Get(projectId)),
		Identifier:        getOptionalString(d.Get(identifier)),
		//Content:           d.Get(content).(optional.Interface),
		Name:             getOptionalString(d.Get(name)),
		FileUsage:        getOptionalString(d.Get(fileUsage)),
		Type_:            getOptionalString(d.Get(type_)),
		ParentIdentifier: getOptionalString(d.Get(parentIdentifier)),
		Description:      getOptionalString(d.Get(description)),
		MimeType:         getOptionalString(d.Get(mimeType)),
		Tags:             getOptionalString(d.Get(tags)),
	}

	return create
}

func getOptionalString(i interface{}) optional.String {
	v, ok := i.(string)
	if !ok {
		return optional.String{}
	}

	if len(v) == 0 {
		return optional.EmptyString()
	}

	return optional.NewString(v)
}

func buildFileStoreApiUpdateRequest(d *schema.ResourceData) *nextgen.FileStoreApiUpdateOpts {
	return &nextgen.FileStoreApiUpdateOpts{
		OrgIdentifier:     getOptionalString(d.Get(orgId)),
		ProjectIdentifier: getOptionalString(d.Get(projectId)),
		Identifier:        getOptionalString(d.Get(identifier)),
		//Content:           d.Get(content).(optional.Interface),
		Name:             getOptionalString(d.Get(name)),
		FileUsage:        getOptionalString(d.Get(fileUsage)),
		Type_:            getOptionalString(d.Get(type_)),
		ParentIdentifier: getOptionalString(d.Get(parentIdentifier)),
		Description:      getOptionalString(d.Get(description)),
		MimeType:         getOptionalString(d.Get(mimeType)),
		Tags:             getOptionalString(d.Get(tags)),
	}
}

func readFileNode(d *schema.ResourceData, file *nextgen.File, content string) {
	d.SetId(file.Identifier)
	d.Set(identifier, file.Identifier)
	d.Set(description, file.Description)
	d.Set(name, file.Name)
	d.Set(orgId, file.OrgIdentifier)
	d.Set(projectId, file.ProjectIdentifier)
	d.Set(fileUsage, file.FileUsage)
	d.Set(type_, file.Type_)
	d.Set(parentIdentifier, file.ParentIdentifier)
	d.Set(mimeType, file.MimeType)
	d.Set(path, file.Path)
	d.Set(draft, file.Draft)
	d.Set(createdBy, []interface{}{
		map[string]interface{}{
			"email": file.CreatedBy.Email,
			"name":  file.CreatedBy.Name,
		},
	})
	d.Set(lastModifiedBy, []interface{}{
		map[string]interface{}{
			"email": file.LastModifiedBy.Email,
			"name":  file.LastModifiedBy.Name,
		},
	})
	d.Set(lastModifiedAt, file.LastModifiedAt)
	d.Set(tags, FlattenTags(file.Tags))
	d.Set(content, content)
}

func FlattenTags(tags []nextgen.NgTag) []string {
	var result []string
	for _, tag := range tags {
		result = append(result, tag.Key+":"+tag.Value)
	}
	return result
}
