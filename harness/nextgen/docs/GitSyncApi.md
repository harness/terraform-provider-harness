# {{classname}}

All URIs are relative to *https://app.harness.io/gateway*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateGitSyncConfig**](GitSyncApi.md#CreateGitSyncConfig) | **Post** /ng/api/git-sync | Creates Git Sync Config in given scope
[**GetGitSyncConfigList**](GitSyncApi.md#GetGitSyncConfigList) | **Get** /ng/api/git-sync | Lists Git Sync Config for the given scope
[**IsGitSyncEnabled**](GitSyncApi.md#IsGitSyncEnabled) | **Get** /ng/api/git-sync/git-sync-enabled | Check whether Git Sync is enabled for given scope or not
[**UpdateDefaultFolder**](GitSyncApi.md#UpdateDefaultFolder) | **Put** /ng/api/git-sync/{identifier}/folder/{folderIdentifier}/default | Update existing Git Sync Config default root folder by Identifier
[**UpdateGitSyncConfig**](GitSyncApi.md#UpdateGitSyncConfig) | **Put** /ng/api/git-sync | Update existing Git Sync Config by Identifier

# **CreateGitSyncConfig**
> GitSyncConfig CreateGitSyncConfig(ctx, body, optional)
Creates Git Sync Config in given scope

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**GitSyncConfig**](GitSyncConfig.md)| Details of Git Sync Config | 
 **optional** | ***GitSyncApiCreateGitSyncConfigOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GitSyncApiCreateGitSyncConfigOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accountIdentifier** | **optional.**| Account Identifier for the Entity | 

### Return type

[**GitSyncConfig**](GitSyncConfig.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json, text/yaml, text/html, text/plain
 - **Accept**: application/json, application/yaml, text/yaml, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetGitSyncConfigList**
> []GitSyncConfig GetGitSyncConfigList(ctx, optional)
Lists Git Sync Config for the given scope

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GitSyncApiGetGitSyncConfigListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GitSyncApiGetGitSyncConfigListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectIdentifier** | **optional.String**| Project Identifier for the Entity | 
 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity | 
 **accountIdentifier** | **optional.String**| Account Identifier for the Entity | 

### Return type

[**[]GitSyncConfig**](GitSyncConfig.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml, text/yaml, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **IsGitSyncEnabled**
> GitEnabled IsGitSyncEnabled(ctx, optional)
Check whether Git Sync is enabled for given scope or not

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GitSyncApiIsGitSyncEnabledOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GitSyncApiIsGitSyncEnabledOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountIdentifier** | **optional.String**| Account Identifier for the Entity | 
 **projectIdentifier** | **optional.String**| Project Identifier for the Entity | 
 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity | 

### Return type

[**GitEnabled**](GitEnabled.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml, text/yaml, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateDefaultFolder**
> GitSyncConfig UpdateDefaultFolder(ctx, identifier, folderIdentifier, optional)
Update existing Git Sync Config default root folder by Identifier

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **identifier** | **string**| Git Sync Config Id | 
  **folderIdentifier** | **string**| Folder Id | 
 **optional** | ***GitSyncApiUpdateDefaultFolderOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GitSyncApiUpdateDefaultFolderOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **projectId** | **optional.String**| Project Identifier for the Entity | 
 **organizationId** | **optional.String**| Organization Identifier for the Entity | 
 **accountId** | **optional.String**| Account Identifier for the Entity | 

### Return type

[**GitSyncConfig**](GitSyncConfig.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml, text/yaml, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateGitSyncConfig**
> GitSyncConfig UpdateGitSyncConfig(ctx, body, optional)
Update existing Git Sync Config by Identifier

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**GitSyncConfig**](GitSyncConfig.md)| Details of Git Sync Config | 
 **optional** | ***GitSyncApiUpdateGitSyncConfigOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GitSyncApiUpdateGitSyncConfigOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accountIdentifier** | **optional.**| Account Identifier for the Entity | 

### Return type

[**GitSyncConfig**](GitSyncConfig.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json, text/yaml, text/html, text/plain
 - **Accept**: application/json, application/yaml, text/yaml, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

