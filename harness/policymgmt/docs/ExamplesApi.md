# {{classname}}

All URIs are relative to *https://app.harness.io/gateway/pm/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ExamplesList**](ExamplesApi.md#ExamplesList) | **Get** /api/v1/examples | 

# **ExamplesList**
> []PolicyExample ExamplesList(ctx, optional)


list examples

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ExamplesApiExamplesListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ExamplesApiExamplesListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountIdentifier** | **optional.String**| Harness account ID | 
 **orgIdentifier** | **optional.String**| Harness organization ID | 
 **projectIdentifier** | **optional.String**| Harness project ID | 
 **xApiKey** | **optional.String**| Harness PAT key used to perform authorization | 

### Return type

[**[]PolicyExample**](PolicyExample.md)

### Authorization

[api_key_header_x-api-key](../README.md#api_key_header_x-api-key), [jwt_header_Authorization](../README.md#jwt_header_Authorization)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

