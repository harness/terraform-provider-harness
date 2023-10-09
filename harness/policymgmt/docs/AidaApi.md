# {{classname}}

All URIs are relative to *https://app.harness.io/gateway/pm/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AidaAnalyze**](AidaApi.md#AidaAnalyze) | **Post** /api/v1/aida/analyze | 
[**AidaGenerate**](AidaApi.md#AidaGenerate) | **Post** /api/v1/aida/generate | 

# **AidaAnalyze**
> AnalyzeResponse AidaAnalyze(ctx, body, optional)


Describe Policy On Basis of rego

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AnalyzeRequestBody**](AnalyzeRequestBody.md)|  | 
 **optional** | ***AidaApiAidaAnalyzeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AidaApiAidaAnalyzeOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accountIdentifier** | **optional.**| Harness account ID | 
 **orgIdentifier** | **optional.**| Harness organization ID | 
 **projectIdentifier** | **optional.**| Harness project ID | 
 **xApiKey** | **optional.**| Harness PAT key used to perform authorization | 

### Return type

[**AnalyzeResponse**](AnalyzeResponse.md)

### Authorization

[api_key_header_x-api-key](../README.md#api_key_header_x-api-key), [jwt_header_Authorization](../README.md#jwt_header_Authorization)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AidaGenerate**
> PolicySample AidaGenerate(ctx, body, optional)


Generate Policy On Basis of free Text

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**GenerateRequestBody**](GenerateRequestBody.md)|  | 
 **optional** | ***AidaApiAidaGenerateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AidaApiAidaGenerateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accountIdentifier** | **optional.**| Harness account ID | 
 **orgIdentifier** | **optional.**| Harness organization ID | 
 **projectIdentifier** | **optional.**| Harness project ID | 
 **xApiKey** | **optional.**| Harness PAT key used to perform authorization | 

### Return type

[**PolicySample**](PolicySample.md)

### Authorization

[api_key_header_x-api-key](../README.md#api_key_header_x-api-key), [jwt_header_Authorization](../README.md#jwt_header_Authorization)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

