# nextgen{{classname}}

All URIs are relative to *https://app.harness.io/gateway*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create**](FileStoreApi.md#Create) | **Post** /ng/api/file-store | Creates file or folder
[**CreateViaYAML**](FileStoreApi.md#CreateViaYAML) | **Post** /ng/api/file-store/yaml | Creates file or folder via YAML
[**DeleteFile**](FileStoreApi.md#DeleteFile) | **Delete** /ng/api/file-store/{identifier} | Delete file or folder by identifier
[**DownloadFile**](FileStoreApi.md#DownloadFile) | **Get** /ng/api/file-store/files/{identifier}/download | Download File
[**GetCreatedByList**](FileStoreApi.md#GetCreatedByList) | **Get** /ng/api/file-store/files/createdBy | Get list of created by usernames.
[**GetEntityTypes**](FileStoreApi.md#GetEntityTypes) | **Get** /ng/api/file-store/supported-entity-types | Get entity types.
[**GetFolderNodes**](FileStoreApi.md#GetFolderNodes) | **Post** /ng/api/file-store/folder | Get Folder nodes.
[**GetReferencedBy**](FileStoreApi.md#GetReferencedBy) | **Get** /ng/api/file-store/{identifier}/referenced-by | Get Referenced by Entities.
[**ListFilesAndFolders**](FileStoreApi.md#ListFilesAndFolders) | **Get** /ng/api/file-store | List files and folders
[**ListFilesWithFilter**](FileStoreApi.md#ListFilesWithFilter) | **Post** /ng/api/file-store/files/filter | Get filtered list of files.
[**Update**](FileStoreApi.md#Update) | **Put** /ng/api/file-store/{identifier} | Updates file or folder
[**UpdateViaYAML**](FileStoreApi.md#UpdateViaYAML) | **Put** /ng/api/file-store/yaml/{identifier} | Updates file or folder via YAML

# **Create**
> ResponseDtoFile Create(ctx, optional)
Creates file or folder

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***FileStoreApiCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FileStoreApiCreateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tags** | **optional.**|  | 
 **content** | [**optional.Interface of interface{}**](.md)|  | 
 **identifier** | **optional.**|  | 
 **name** | **optional.**|  | 
 **fileUsage** | **optional.**|  | 
 **type_** | **optional.**|  | 
 **parentIdentifier** | **optional.**|  | 
 **description** | **optional.**|  | 
 **mimeType** | **optional.**|  | 
 **path** | **optional.**|  | 
 **createdBy** | [**optional.Interface of EmbeddedUserDetailsDto**](.md)|  | 
 **lastModifiedBy** | [**optional.Interface of EmbeddedUserDetailsDto**](.md)|  | 
 **lastModifiedAt** | **optional.**|  | 
 **accountIdentifier** | **optional.**| Account Identifier for the Entity. | 
 **orgIdentifier** | **optional.**| Organization Identifier for the Entity. | 
 **projectIdentifier** | **optional.**| Project Identifier for the Entity. | 

### Return type

[**ResponseDtoFile**](ResponseDTOFile.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateViaYAML**
> ResponseDtoFile CreateViaYAML(ctx, body, optional)
Creates file or folder via YAML

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**FileStoreRequest**](FileStoreRequest.md)| YAML definition of file or folder | 
 **optional** | ***FileStoreApiCreateViaYAMLOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FileStoreApiCreateViaYAMLOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accountIdentifier** | **optional.**| Account Identifier for the Entity. | 
 **orgIdentifier** | **optional.**| Organization Identifier for the Entity. | 
 **projectIdentifier** | **optional.**| Project Identifier for the Entity. | 

### Return type

[**ResponseDtoFile**](ResponseDTOFile.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteFile**
> ResponseDtoBoolean DeleteFile(ctx, identifier, optional)
Delete file or folder by identifier

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **identifier** | **string**| The file identifier | 
 **optional** | ***FileStoreApiDeleteFileOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FileStoreApiDeleteFileOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accountIdentifier** | **optional.String**| Account Identifier for the Entity. | 
 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity. | 
 **projectIdentifier** | **optional.String**| Project Identifier for the Entity. | 

### Return type

[**ResponseDtoBoolean**](ResponseDTOBoolean.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DownloadFile**
> DownloadFile(ctx, identifier, optional)
Download File

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **identifier** | **string**| The file identifier | 
 **optional** | ***FileStoreApiDownloadFileOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FileStoreApiDownloadFileOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accountIdentifier** | **optional.String**| Account Identifier for the Entity. | 
 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity. | 
 **projectIdentifier** | **optional.String**| Project Identifier for the Entity. | 

### Return type

 (empty response body)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCreatedByList**
> ResponseDtoSetEmbeddedUserDetailsDto GetCreatedByList(ctx, optional)
Get list of created by usernames.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***FileStoreApiGetCreatedByListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FileStoreApiGetCreatedByListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountIdentifier** | **optional.String**| Account Identifier for the Entity. | 
 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity. | 
 **projectIdentifier** | **optional.String**| Project Identifier for the Entity. | 

### Return type

[**ResponseDtoSetEmbeddedUserDetailsDto**](ResponseDTOSetEmbeddedUserDetailsDTO.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEntityTypes**
> ResponseDtoListEntityType GetEntityTypes(ctx, optional)
Get entity types.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***FileStoreApiGetEntityTypesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FileStoreApiGetEntityTypesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountIdentifier** | **optional.String**| Account Identifier for the Entity. | 

### Return type

[**ResponseDtoListEntityType**](ResponseDTOListEntityType.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFolderNodes**
> ResponseDtoFolderNode GetFolderNodes(ctx, body, optional)
Get Folder nodes.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**FolderNode**](FolderNode.md)| Folder node for which to return the list of nodes | 
 **optional** | ***FileStoreApiGetFolderNodesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FileStoreApiGetFolderNodesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accountIdentifier** | **optional.**| Account Identifier for the Entity. | 
 **orgIdentifier** | **optional.**| Organization Identifier for the Entity. | 
 **projectIdentifier** | **optional.**| Project Identifier for the Entity. | 

### Return type

[**ResponseDtoFolderNode**](ResponseDTOFolderNode.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetReferencedBy**
> ResponseDtoPageEntitySetupUsage GetReferencedBy(ctx, identifier, optional)
Get Referenced by Entities.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **identifier** | **string**| The file identifier | 
 **optional** | ***FileStoreApiGetReferencedByOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FileStoreApiGetReferencedByOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageIndex** | **optional.Int32**| Page number of navigation. The default value is 0 | [default to 0]
 **pageSize** | **optional.Int32**| Number of entries per page. The default value is 100 | [default to 100]
 **accountIdentifier** | **optional.String**| Account Identifier for the Entity. | 
 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity. | 
 **projectIdentifier** | **optional.String**| Project Identifier for the Entity. | 
 **entityType** | **optional.String**| Entity type | 
 **searchTerm** | **optional.String**|  | 

### Return type

[**ResponseDtoPageEntitySetupUsage**](ResponseDTOPageEntitySetupUsage.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListFilesAndFolders**
> ResponseDtoPageFile ListFilesAndFolders(ctx, optional)
List files and folders

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***FileStoreApiListFilesAndFoldersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FileStoreApiListFilesAndFoldersOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountIdentifier** | **optional.String**| Account Identifier for the Entity. | 
 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity. | 
 **projectIdentifier** | **optional.String**| Project Identifier for the Entity. | 
 **identifiers** | [**optional.Interface of []string**](string.md)| This is the list of File IDs. Details specific to these IDs would be fetched. | 
 **searchTerm** | **optional.String**| This would be used to filter Files. Any Files having the specified string in its Name, ID and Tag would be filtered. | 
 **pageIndex** | **optional.Int32**| Page Index of the results to fetch.Default Value: 0 | [default to 0]
 **pageSize** | **optional.Int32**| Results per page(max 100)Default Value: 50 | [default to 50]
 **sortOrders** | [**optional.Interface of []SortOrder**](SortOrder.md)| Sort criteria for the elements. | 

### Return type

[**ResponseDtoPageFile**](ResponseDTOPageFile.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListFilesWithFilter**
> ResponseDtoPageFile ListFilesWithFilter(ctx, optional)
Get filtered list of files.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***FileStoreApiListFilesWithFilterOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FileStoreApiListFilesWithFilterOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of FilesFilterPropertiesDto**](FilesFilterPropertiesDto.md)| Details of the File filter properties to be applied | 
 **pageIndex** | **optional.**| Page Index of the results to fetch.Default Value: 0 | [default to 0]
 **pageSize** | **optional.**| Results per page(max 100)Default Value: 50 | [default to 50]
 **sortOrders** | [**optional.Interface of []SortOrder**](SortOrder.md)| Sort criteria for the elements. | 
 **accountIdentifier** | **optional.**| Account Identifier for the Entity. | 
 **orgIdentifier** | **optional.**| Organization Identifier for the Entity. | 
 **projectIdentifier** | **optional.**| Project Identifier for the Entity. | 
 **filterIdentifier** | **optional.**|  | 
 **searchTerm** | **optional.**|  | 

### Return type

[**ResponseDtoPageFile**](ResponseDTOPageFile.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Update**
> ResponseDtoFile Update(ctx, identifier, optional)
Updates file or folder

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **identifier** | **string**| The file identifier | 
 **optional** | ***FileStoreApiUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FileStoreApiUpdateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tags** | **optional.**|  | 
 **identifier** | **optional.**|  | 
 **name** | **optional.**|  | 
 **fileUsage** | **optional.**|  | 
 **type_** | **optional.**|  | 
 **parentIdentifier** | **optional.**|  | 
 **description** | **optional.**|  | 
 **mimeType** | **optional.**|  | 
 **path** | **optional.**|  | 
 **createdBy** | [**optional.Interface of EmbeddedUserDetailsDto**](.md)|  | 
 **lastModifiedBy** | [**optional.Interface of EmbeddedUserDetailsDto**](.md)|  | 
 **lastModifiedAt** | **optional.**|  | 
 **content** | [**optional.Interface of interface{}**](.md)|  | 
 **accountIdentifier** | **optional.**| Account Identifier for the Entity. | 
 **orgIdentifier** | **optional.**| Organization Identifier for the Entity. | 
 **projectIdentifier** | **optional.**| Project Identifier for the Entity. | 

### Return type

[**ResponseDtoFile**](ResponseDTOFile.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateViaYAML**
> ResponseDtoFile UpdateViaYAML(ctx, body, identifier, optional)
Updates file or folder via YAML

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**FileStoreRequest**](FileStoreRequest.md)| YAML definition of file or folder | 
  **identifier** | **string**| The file identifier | 
 **optional** | ***FileStoreApiUpdateViaYAMLOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FileStoreApiUpdateViaYAMLOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **accountIdentifier** | **optional.**| Account Identifier for the Entity. | 
 **orgIdentifier** | **optional.**| Organization Identifier for the Entity. | 
 **projectIdentifier** | **optional.**| Project Identifier for the Entity. | 

### Return type

[**ResponseDtoFile**](ResponseDTOFile.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

