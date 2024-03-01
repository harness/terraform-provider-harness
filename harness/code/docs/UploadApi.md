# {{classname}}

All URIs are relative to */gateway/code/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RepoArtifactDownload**](UploadApi.md#RepoArtifactDownload) | **Get** /repos/{repo_identifier}/uploads/{file_ref} | Repo artifact download
[**RepoArtifactUpload**](UploadApi.md#RepoArtifactUpload) | **Post** /repos/{repo_identifier}/uploads | Repo artifact upload

# **RepoArtifactDownload**
> RepoArtifactDownload(ctx, accountIdentifier, repoIdentifier, fileRef, optional)
Repo artifact download

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity.. | 
  **repoIdentifier** | **string**|  | 
  **fileRef** | **string**|  | 
 **optional** | ***UploadApiRepoArtifactDownloadOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UploadApiRepoArtifactDownloadOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity.. | 
 **projectIdentifier** | **optional.String**| Project Identifier for the Entity.. | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth), [x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RepoArtifactUpload**
> UploadResult RepoArtifactUpload(ctx, accountIdentifier, repoIdentifier, optional)
Repo artifact upload

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity.. | 
  **repoIdentifier** | **string**|  | 
 **optional** | ***UploadApiRepoArtifactUploadOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UploadApiRepoArtifactUploadOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity.. | 
 **projectIdentifier** | **optional.String**| Project Identifier for the Entity.. | 

### Return type

[**UploadResult**](UploadResult.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

