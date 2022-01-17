# {{classname}}

All URIs are relative to *https://app.harness.io/gateway*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateGitSyncSetting**](GitSyncSettingsApi.md#CreateGitSyncSetting) | **Post** /ng/api/git-sync-settings | Creates Git Sync Setting in a scope
[**GetGitSyncSettings**](GitSyncSettingsApi.md#GetGitSyncSettings) | **Get** /ng/api/git-sync-settings | Get Git Sync Setting for the given scope
[**UpdateGitSyncSetting**](GitSyncSettingsApi.md#UpdateGitSyncSetting) | **Put** /ng/api/git-sync-settings | This updates the existing Git Sync settings within the scope. Only changing Connectivity Mode is allowed

# **CreateGitSyncSetting**
> ResponseDtoGitSyncSettings CreateGitSyncSetting(ctx, body)
Creates Git Sync Setting in a scope

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**GitSyncSettings**](GitSyncSettings.md)| This contains details of Git Sync settings like - (scope, executionOnDelegate) | 

### Return type

[**ResponseDtoGitSyncSettings**](ResponseDTOGitSyncSettings.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json, text/yaml, text/html
 - **Accept**: application/json, application/yaml, text/yaml, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetGitSyncSettings**
> ResponseDtoGitSyncSettings GetGitSyncSettings(ctx, optional)
Get Git Sync Setting for the given scope

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GitSyncSettingsApiGetGitSyncSettingsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GitSyncSettingsApiGetGitSyncSettingsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectIdentifier** | **optional.String**| Project Identifier for the Entity | 
 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity | 
 **accountIdentifier** | **optional.String**| Account Identifier for the Entity | 

### Return type

[**ResponseDtoGitSyncSettings**](ResponseDTOGitSyncSettings.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml, text/yaml, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateGitSyncSetting**
> ResponseDtoGitSyncSettings UpdateGitSyncSetting(ctx, body)
This updates the existing Git Sync settings within the scope. Only changing Connectivity Mode is allowed

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**GitSyncSettings**](GitSyncSettings.md)| This contains details of Git Sync Settings | 

### Return type

[**ResponseDtoGitSyncSettings**](ResponseDTOGitSyncSettings.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json, text/yaml, text/html
 - **Accept**: application/json, application/yaml, text/yaml, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

