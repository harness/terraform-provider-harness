# {{classname}}

All URIs are relative to *https://app.harness.io/gateway*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSourceCodeManager**](SourceCodeManagerApi.md#CreateSourceCodeManager) | **Post** /ng/api/source-code-manager | Creates Source Code Manager
[**DeleteSourceCodeManager**](SourceCodeManagerApi.md#DeleteSourceCodeManager) | **Delete** /ng/api/source-code-manager/{identifier} | Deletes the Source Code Manager corresponding to the specified Source Code Manager Id
[**GetSourceCodeManagers**](SourceCodeManagerApi.md#GetSourceCodeManagers) | **Get** /ng/api/source-code-manager | Lists Source Code Managers for the given account
[**UpdateSourceCodeManager**](SourceCodeManagerApi.md#UpdateSourceCodeManager) | **Put** /ng/api/source-code-manager/{identifier} | Updates Source Code Manager Details with the given Source Code Manager Id

# **CreateSourceCodeManager**
> ResponseDtoSourceCodeManager CreateSourceCodeManager(ctx, optional)
Creates Source Code Manager

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SourceCodeManagerApiCreateSourceCodeManagerOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SourceCodeManagerApiCreateSourceCodeManagerOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of SourceCodeManager**](SourceCodeManager.md)| This contains details of Source Code Manager | 

### Return type

[**ResponseDtoSourceCodeManager**](ResponseDTOSourceCodeManager.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteSourceCodeManager**
> ResponseDtoBoolean DeleteSourceCodeManager(ctx, identifier, accountIdentifier)
Deletes the Source Code Manager corresponding to the specified Source Code Manager Id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **identifier** | **string**| Source Code manager Identifier | 
  **accountIdentifier** | **string**| Account Identifier for the Entity | 

### Return type

[**ResponseDtoBoolean**](ResponseDTOBoolean.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSourceCodeManagers**
> ResponseDtoListSourceCodeManager GetSourceCodeManagers(ctx, accountIdentifier)
Lists Source Code Managers for the given account

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity | 

### Return type

[**ResponseDtoListSourceCodeManager**](ResponseDTOListSourceCodeManager.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateSourceCodeManager**
> ResponseDtoSourceCodeManager UpdateSourceCodeManager(ctx, identifier, optional)
Updates Source Code Manager Details with the given Source Code Manager Id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **identifier** | **string**| Source Code manager Identifier | 
 **optional** | ***SourceCodeManagerApiUpdateSourceCodeManagerOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SourceCodeManagerApiUpdateSourceCodeManagerOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of SourceCodeManager**](SourceCodeManager.md)| This contains details of Source Code Manager | 

### Return type

[**ResponseDtoSourceCodeManager**](ResponseDTOSourceCodeManager.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

