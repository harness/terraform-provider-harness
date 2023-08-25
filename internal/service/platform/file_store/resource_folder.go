package file_store

import (
	"context"
	"net/http"

	"github.com/antihax/optional"
	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceFileStoreNodeFolder() *schema.Resource {
	resource := &schema.Resource{
		Description: "Resource for creating folders in Harness.",

		ReadContext:   resourceFileStoreNodeFolderRead,
		UpdateContext: resourceFileStoreNodeFolderCreateOrUpdate,
		CreateContext: resourceFileStoreNodeFolderCreateOrUpdate,
		DeleteContext: resourceFileStoreNodeFolderDelete,
		Importer:      helpers.MultiLevelResourceImporter,

		Schema: map[string]*schema.Schema{
			"parent_identifier": {
				Description: "Folder parent identifier on Harness File Store",
				Type:        schema.TypeString,
				Required:    true,
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

	helpers.SetMultiLevelResourceSchema(resource.Schema)

	return resource
}

func resourceFileStoreNodeFolderRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)
	id := d.Get(identifier).(string)

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

	readFolderNode(d, resp.Data, optional.EmptyInterface())

	return nil
}

func resourceFileStoreNodeFolderCreateOrUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)
	id := d.Id()

	var resp nextgen.ResponseDtoFile
	var httpResp *http.Response
	var err error
	var fileContent optional.Interface

	if id == "" {
		createRequest, internalErr := buildFileStoreApiFolderCreateRequest(d)
		if internalErr != nil {
			return helpers.HandleApiError(internalErr, d, httpResp)
		}
		fileContent = createRequest.Content
		resp, httpResp, err = c.FileStoreApi.Create(ctx, c.AccountId, createRequest)
	} else {
		updateRequest, internalErr := buildFileStoreApiFolderUpdateRequest(d)
		if internalErr != nil {
			return helpers.HandleApiError(internalErr, d, httpResp)
		}
		fileContent = updateRequest.Content
		resp, httpResp, err = c.FileStoreApi.Update(ctx, c.AccountId, id, updateRequest)
	}

	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	readFolderNode(d, resp.Data, fileContent)

	return nil
}

func resourceFileStoreNodeFolderDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

func buildFileStoreApiFolderCreateRequest(d *schema.ResourceData) (*nextgen.FileStoreApiCreateOpts, error) {
	var tagsJson string
	if attr, ok := d.GetOk(tags); ok {
		tags := attr.(*schema.Set)
		tagsJson = buildTagsJson(tags)
	}

	create := &nextgen.FileStoreApiCreateOpts{
		OrgIdentifier:     getOptionalString(d.Get(orgId)),
		ProjectIdentifier: getOptionalString(d.Get(projectId)),
		Identifier:        getOptionalString(d.Get(identifier)),
		Name:              getOptionalString(d.Get(name)),
		Description:       getOptionalString(d.Get(description)),
		Type_:             getOptionalString(nextgen.NGFileTypes.Folder.String()),
		Tags:              getOptionalString(tagsJson),
		ParentIdentifier:  getOptionalString(d.Get(parentIdentifier)),
	}

	return create, nil
}

func buildFileStoreApiFolderUpdateRequest(d *schema.ResourceData) (*nextgen.FileStoreApiUpdateOpts, error) {
	var tagsJson string
	if attr, ok := d.GetOk(tags); ok {
		tags := attr.(*schema.Set)
		tagsJson = buildTagsJson(tags)
	}

	update := &nextgen.FileStoreApiUpdateOpts{
		OrgIdentifier:     getOptionalString(d.Get(orgId)),
		ProjectIdentifier: getOptionalString(d.Get(projectId)),
		Identifier:        getOptionalString(d.Get(identifier)),
		Name:              getOptionalString(d.Get(name)),
		Description:       getOptionalString(d.Get(description)),
		Type_:             getOptionalString(nextgen.NGFileTypes.Folder.String()),
		Tags:              getOptionalString(tagsJson),
		ParentIdentifier:  getOptionalString(d.Get(parentIdentifier)),
	}

	return update, nil
}

func readFolderNode(d *schema.ResourceData, file *nextgen.File, fileContentOpt optional.Interface) {
	d.SetId(file.Identifier)
	d.Set(identifier, file.Identifier)
	d.Set(orgId, file.OrgIdentifier)
	d.Set(projectId, file.ProjectIdentifier)
	d.Set(name, file.Name)
	d.Set(description, file.Description)
	d.Set(tags, FlattenTags(file.Tags))
	d.Set(parentIdentifier, file.ParentIdentifier)
	d.Set(path, file.Path)
	d.Set(createdBy, []interface{}{
		map[string]interface{}{
			"email": getSafeEmail(file.CreatedBy),
			"name":  getSafeName(file.CreatedBy),
		},
	})
	d.Set(lastModifiedBy, []interface{}{
		map[string]interface{}{
			"email": getSafeEmail(file.LastModifiedBy),
			"name":  getSafeName(file.LastModifiedBy),
		},
	})
	d.Set(lastModifiedAt, file.LastModifiedAt)
}
