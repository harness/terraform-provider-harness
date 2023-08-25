package file_store

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/antihax/optional"
	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceFileStoreNodeFile() *schema.Resource {
	resource := &schema.Resource{
		Description: "Resource for creating files in Harness.",

		ReadContext:   resourceFileStoreNodeFileRead,
		UpdateContext: resourceFileStoreNodeFileCreateOrUpdate,
		CreateContext: resourceFileStoreNodeFileCreateOrUpdate,
		DeleteContext: resourceFileStoreNodeFileDelete,
		Importer:      helpers.MultiLevelResourceImporter,

		Schema: map[string]*schema.Schema{
			"parent_identifier": {
				Description: "File parent identifier on Harness File Store",
				Type:        schema.TypeString,
				Required:    true,
			},
			"file_content_path": {
				Description: "File content path to be upladed on Harness File Store",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"mime_type": {
				Description: "File mime type",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"file_usage": {
				Description: fmt.Sprintf("File usage. Valid options are %s", strings.Join(nextgen.FileUsageValues, ", ")),
				Type:        schema.TypeString,
				Optional:    true,
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

	helpers.SetMultiLevelResourceSchema(resource.Schema)

	return resource
}

func resourceFileStoreNodeFileRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)
	id := d.Get(identifier).(string)
	var contentStr optional.Interface

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

	// download content
	downloadResp, bodyContent, downloadErr := c.FileStoreApi.DownloadFile(ctx, id, c.AccountId, &nextgen.FileStoreApiDownloadFileOpts{
		OrgIdentifier:     helpers.BuildField(d, orgId),
		ProjectIdentifier: helpers.BuildField(d, projectId),
	})
	if downloadErr != nil {
		return helpers.HandleReadApiError(downloadErr, d, downloadResp)
	}

	contentStr = optional.NewInterface(bodyContent)

	readFileNode(d, resp.Data, contentStr)
	return nil
}

func resourceFileStoreNodeFileCreateOrUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)
	id := d.Id()

	var resp nextgen.ResponseDtoFile
	var httpResp *http.Response
	var err error
	var fileContent optional.Interface

	if id == "" {
		createRequest, internalErr := buildFileStoreApiFileCreateRequest(d)
		if internalErr != nil {
			return helpers.HandleApiError(internalErr, d, httpResp)
		}
		fileContent = createRequest.Content
		resp, httpResp, err = c.FileStoreApi.Create(ctx, c.AccountId, createRequest)
	} else {
		updateRequest, internalErr := buildFileStoreApiFileUpdateRequest(d)
		if internalErr != nil {
			return helpers.HandleApiError(internalErr, d, httpResp)
		}
		fileContent = updateRequest.Content
		resp, httpResp, err = c.FileStoreApi.Update(ctx, c.AccountId, id, updateRequest)
	}

	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	readFileNode(d, resp.Data, fileContent)

	return nil
}

func resourceFileStoreNodeFileDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

func buildFileStoreApiFileCreateRequest(d *schema.ResourceData) (*nextgen.FileStoreApiCreateOpts, error) {
	fileContent, err := getFileContent(d.Get(fileContentPath))
	if err != nil {
		return nil, err
	}

	var tagsJson string
	if attr, ok := d.GetOk(tags); ok {
		tags := attr.(*schema.Set)
		tagsJson = buildTagsJson(tags)
	}

	create := &nextgen.FileStoreApiCreateOpts{
		OrgIdentifier:     getOptionalString(d.Get(orgId)),
		ProjectIdentifier: getOptionalString(d.Get(projectId)),
		Identifier:        getOptionalString(d.Get(identifier)),
		Content:           fileContent,
		Name:              getOptionalString(d.Get(name)),
		FileUsage:         getOptionalString(d.Get(fileUsage)),
		Type_:             getOptionalString(nextgen.NGFileTypes.File.String()),
		ParentIdentifier:  getOptionalString(d.Get(parentIdentifier)),
		Description:       getOptionalString(d.Get(description)),
		MimeType:          getOptionalString(d.Get(mimeType)),
		Tags:              getOptionalString(tagsJson),
	}

	return create, nil
}

func buildFileStoreApiFileUpdateRequest(d *schema.ResourceData) (*nextgen.FileStoreApiUpdateOpts, error) {
	fileContent, err := getFileContent(d.Get(fileContentPath))
	if err != nil {
		return nil, err
	}

	var tagsJson string
	if attr, ok := d.GetOk(tags); ok {
		tags := attr.(*schema.Set)
		tagsJson = buildTagsJson(tags)
	}

	update := &nextgen.FileStoreApiUpdateOpts{
		OrgIdentifier:     getOptionalString(d.Get(orgId)),
		ProjectIdentifier: getOptionalString(d.Get(projectId)),
		Identifier:        getOptionalString(d.Get(identifier)),
		Content:           fileContent,
		Name:              getOptionalString(d.Get(name)),
		FileUsage:         getOptionalString(d.Get(fileUsage)),
		Type_:             getOptionalString(nextgen.NGFileTypes.File.String()),
		ParentIdentifier:  getOptionalString(d.Get(parentIdentifier)),
		Description:       getOptionalString(d.Get(description)),
		MimeType:          getOptionalString(d.Get(mimeType)),
		Tags:              getOptionalString(tagsJson),
	}

	return update, nil
}

func readFileNode(d *schema.ResourceData, file *nextgen.File, fileContentOpt optional.Interface) {
	d.SetId(file.Identifier)
	d.Set(identifier, file.Identifier)
	d.Set(name, file.Name)
	d.Set(orgId, file.OrgIdentifier)
	d.Set(projectId, file.ProjectIdentifier)
	d.Set(parentIdentifier, file.ParentIdentifier)
	d.Set(path, file.Path)
	d.Set(tags, FlattenTags(file.Tags))
	d.Set(createdBy, []interface{}{
		map[string]interface{}{
			"email": getEmail(file.CreatedBy),
			"name":  getName(file.CreatedBy),
		},
	})
	d.Set(lastModifiedBy, []interface{}{
		map[string]interface{}{
			"email": getEmail(file.LastModifiedBy),
			"name":  getName(file.LastModifiedBy),
		},
	})
	d.Set(lastModifiedAt, file.LastModifiedAt)
	d.Set(description, file.Description)
	d.Set(fileUsage, file.FileUsage)
	d.Set(mimeType, file.MimeType)
	//content
	var fileContent string
	if fileContentOpt.IsSet() {
		fileContent = string(fileContentOpt.Value().([]byte))
	} else {
		fileContent = ""
	}
	d.Set(content, fileContent)
}

func getFileContent(filePath interface{}) (optional.Interface, error) {
	filePathStr, ok := filePath.(string)
	if !ok {
		return optional.Interface{}, nil
	}

	if len(filePathStr) == 0 {
		return optional.EmptyInterface(), nil
	}

	fileContent, err := ioutil.ReadFile(filePathStr)

	if err != nil {
		return optional.EmptyInterface(), err
	}

	return optional.NewInterface(fileContent), nil
}
