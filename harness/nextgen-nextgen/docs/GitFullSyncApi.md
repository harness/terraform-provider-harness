# nextgen{{classname}}

All URIs are relative to *https://app.harness.io/gateway*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateGitFullSyncConfig**](GitFullSyncApi.md#CreateGitFullSyncConfig) | **Post** /ng/api/git-full-sync/config | Create Configuration for Git Full Sync for the provided scope
[**GetGitFullSyncConfig**](GitFullSyncApi.md#GetGitFullSyncConfig) | **Get** /ng/api/git-full-sync/config | Fetch Configuration for Git Full Sync for the provided scope
[**ListFullSyncFiles**](GitFullSyncApi.md#ListFullSyncFiles) | **Post** /ng/api/git-full-sync/files | List files in full sync along with their status
[**TriggerFullSync**](GitFullSyncApi.md#TriggerFullSync) | **Post** /ng/api/git-full-sync | Trigger Full Sync
[**UpdateGitFullSyncConfig**](GitFullSyncApi.md#UpdateGitFullSyncConfig) | **Put** /ng/api/git-full-sync/config | Update Configuration for Git Full Sync for the provided scope

# **CreateGitFullSyncConfig**
> ResponseDtoGitFullSyncConfig CreateGitFullSyncConfig(ctx, accountIdentifier, optional)
Create Configuration for Git Full Sync for the provided scope

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 
 **optional** | ***GitFullSyncApiCreateGitFullSyncConfigOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GitFullSyncApiCreateGitFullSyncConfigOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of GitFullSyncConfigRequest**](GitFullSyncConfigRequest.md)| Details of the Git Full sync Configuration | 
 **orgIdentifier** | **optional.**| Organization Identifier for the Entity. | 
 **projectIdentifier** | **optional.**| Project Identifier for the Entity. | 

### Return type

[**ResponseDtoGitFullSyncConfig**](ResponseDTOGitFullSyncConfig.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json, text/yaml, text/html
 - **Accept**: application/json, application/yaml, text/yaml, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetGitFullSyncConfig**
> ResponseDtoGitFullSyncConfig GetGitFullSyncConfig(ctx, accountIdentifier, optional)
Fetch Configuration for Git Full Sync for the provided scope

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 
 **optional** | ***GitFullSyncApiGetGitFullSyncConfigOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GitFullSyncApiGetGitFullSyncConfigOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity. | 
 **projectIdentifier** | **optional.String**| Project Identifier for the Entity. | 

### Return type

[**ResponseDtoGitFullSyncConfig**](ResponseDTOGitFullSyncConfig.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml, text/yaml, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListFullSyncFiles**
> ResponseDtoPageResponseGitFullSyncEntityInfo ListFullSyncFiles(ctx, accountIdentifier, optional)
List files in full sync along with their status

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 
 **optional** | ***GitFullSyncApiListFullSyncFilesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GitFullSyncApiListFullSyncFilesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of GitFullSyncEntityInfoFilter**](GitFullSyncEntityInfoFilter.md)| Entity Type and Sync Status | 
 **orgIdentifier** | **optional.**| Organization Identifier for the Entity. | 
 **projectIdentifier** | **optional.**| Project Identifier for the Entity. | 
 **pageIndex** | **optional.**| Page Index of the results to fetch.Default Value: 0 | [default to 0]
 **pageSize** | **optional.**| Results per page(max 100)Default Value: 50 | [default to 50]
 **sortOrders** | [**optional.Interface of []SortOrder**](SortOrder.md)| Sort criteria for the elements. | 
 **searchTerm** | **optional.**| Search Term. | 

### Return type

[**ResponseDtoPageResponseGitFullSyncEntityInfo**](ResponseDTOPageResponseGitFullSyncEntityInfo.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json, text/yaml, text/html
 - **Accept**: application/json, application/yaml, text/yaml, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TriggerFullSync**
> ResponseDtoTriggerGitFullSyncResponse TriggerFullSync(ctx, accountIdentifier, optional)
Trigger Full Sync

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 
 **optional** | ***GitFullSyncApiTriggerFullSyncOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GitFullSyncApiTriggerFullSyncOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity. | 
 **projectIdentifier** | **optional.String**| Project Identifier for the Entity. | 

### Return type

[**ResponseDtoTriggerGitFullSyncResponse**](ResponseDTOTriggerGitFullSyncResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml, text/yaml, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateGitFullSyncConfig**
> ResponseDtoGitFullSyncConfig UpdateGitFullSyncConfig(ctx, accountIdentifier, optional)
Update Configuration for Git Full Sync for the provided scope

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 
 **optional** | ***GitFullSyncApiUpdateGitFullSyncConfigOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GitFullSyncApiUpdateGitFullSyncConfigOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of GitFullSyncConfigRequest**](GitFullSyncConfigRequest.md)| Details of the Git Full sync Configuration | 
 **orgIdentifier** | **optional.**| Organization Identifier for the Entity. | 
 **projectIdentifier** | **optional.**| Project Identifier for the Entity. | 

### Return type

[**ResponseDtoGitFullSyncConfig**](ResponseDTOGitFullSyncConfig.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json, text/yaml, text/html
 - **Accept**: application/json, application/yaml, text/yaml, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

