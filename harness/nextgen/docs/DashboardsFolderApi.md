# nextgen{{classname}}

All URIs are relative to *https://app.harness.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateFolder**](DashboardsFoldersApi.md#CreateFolder) | **Post** /dashboard/folders | Create a Dashboard Folder
[**DeleteFolder**](DashboardsFoldersApi.md#DeleteFolder) | **Delete** /dashboard/folders/{folder_id} | Delete a Dashboard Folder
[**GetFolder**](DashboardsFoldersApi.md#GetFolder) | **Get** /dashboard/folders/{folder_id} | Get a Dashboard Folder
[**UpdateFolder**](DashboardsFoldersApi.md#UpdateFolder) | **Patch** /dashboard/folders/{folder_id} | Update a Dashboard Folder

# **CreateFolder**
> GetFolderResponse CreateFolder(ctx, body, optional)


Create a new folder.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateFolderRequestBody**](CreateFolderRequestBody.md)| Create a new folder | 
 **optional** | ***DashboardsFoldersApiCreateFolderOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DashboardsFoldersApiCreateFolderOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accountId** | **optional.**|  | 

### Return type

[**GetFolderResponse**](GetFolderResponse.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteFolder**
> GetFolderResponse DeleteFolder(ctx, folderId, optional)


Delete a folder along with any dashboards it contains.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **folderId** | **string**|  | 
 **optional** | ***DashboardsFoldersApiDeleteFolderOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DashboardsFoldersApiDeleteFolderOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accountId** | **optional.String**|  | 

### Return type

[**GetFolderResponse**](GetFolderResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFolder**
> GetFolderResponse GetFolder(ctx, folderId, optional)


Get a folder by ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **folderId** | **string**|  | 
 **optional** | ***DashboardsFoldersApiGetFolderOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DashboardsFoldersApiGetFolderOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accountId** | **optional.String**|  | 

### Return type

[**GetFolderResponse**](GetFolderResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateFolder**
> GetFolderResponse UpdateFolder(ctx, body, folderId, optional)


Update a folder's name.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateFolderRequestBody**](UpdateFolderRequestBody.md)| Change the name of a folder | 
  **folderId** | **string**|  | 
 **optional** | ***DashboardsFoldersApiUpdateFolderOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DashboardsFoldersApiUpdateFolderOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **accountId** | **optional.**|  | 

### Return type

[**GetFolderResponse**](GetFolderResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

