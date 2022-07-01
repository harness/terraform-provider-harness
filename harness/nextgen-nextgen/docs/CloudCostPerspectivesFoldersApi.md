# nextgen{{classname}}

All URIs are relative to *https://app.harness.io/gateway*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePerspectiveFolder**](CloudCostPerspectivesFoldersApi.md#CreatePerspectiveFolder) | **Post** /ccm/api/perspectiveFolders/create | Create a Perspective folder
[**DeleteFolder**](CloudCostPerspectivesFoldersApi.md#DeleteFolder) | **Delete** /ccm/api/perspectiveFolders/{folderId} | Delete a folder
[**GetAllFolderPerspectives**](CloudCostPerspectivesFoldersApi.md#GetAllFolderPerspectives) | **Get** /ccm/api/perspectiveFolders/{folderId}/perspectives | Return details of all the Perspectives
[**GetFolders**](CloudCostPerspectivesFoldersApi.md#GetFolders) | **Get** /ccm/api/perspectiveFolders | Fetch folders for an account
[**MovePerspectives**](CloudCostPerspectivesFoldersApi.md#MovePerspectives) | **Post** /ccm/api/perspectiveFolders/movePerspectives | Move a Perspective
[**UpdateFolder**](CloudCostPerspectivesFoldersApi.md#UpdateFolder) | **Put** /ccm/api/perspectiveFolders | Update a folder

# **CreatePerspectiveFolder**
> ResponseDtoceViewFolder CreatePerspectiveFolder(ctx, body, accountIdentifier)
Create a Perspective folder

Create a Perspective Folder.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreatePerspectiveFolderDto**](CreatePerspectiveFolderDto.md)| Request body containing Perspective&#x27;s CEViewFolder object | 
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 

### Return type

[**ResponseDtoceViewFolder**](ResponseDTOCEViewFolder.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteFolder**
> ResponseDtoBoolean DeleteFolder(ctx, accountIdentifier, folderId)
Delete a folder

Delete a Folder for the given Folder ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 
  **folderId** | **string**| Unique identifier for the Perspective folder | 

### Return type

[**ResponseDtoBoolean**](ResponseDTOBoolean.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllFolderPerspectives**
> ResponseDtoListPerspective GetAllFolderPerspectives(ctx, accountIdentifier, folderId)
Return details of all the Perspectives

Return details of all the Perspectives for the given account ID and folder

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 
  **folderId** | **string**| Unique identifier for folder | 

### Return type

[**ResponseDtoListPerspective**](ResponseDTOListPerspective.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFolders**
> ResponseDtoListCeViewFolder GetFolders(ctx, accountIdentifier)
Fetch folders for an account

Fetch folders given an accountId

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 

### Return type

[**ResponseDtoListCeViewFolder**](ResponseDTOListCEViewFolder.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MovePerspectives**
> ResponseDtoListCeView MovePerspectives(ctx, body, accountIdentifier)
Move a Perspective

Move a perspective from a folder to another.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MovePerspectiveDto**](MovePerspectiveDto.md)| Request body containing perspectiveIds to be moved and newFolderId | 
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 

### Return type

[**ResponseDtoListCeView**](ResponseDTOListCEView.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateFolder**
> ResponseDtoceViewFolder UpdateFolder(ctx, body, accountIdentifier)
Update a folder

Update a folder

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CeViewFolder**](CeViewFolder.md)| Request body containing ceViewFolder object | 
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 

### Return type

[**ResponseDtoceViewFolder**](ResponseDTOCEViewFolder.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

