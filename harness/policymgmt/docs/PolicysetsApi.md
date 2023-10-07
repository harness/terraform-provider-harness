# {{classname}}

All URIs are relative to *https://app.harness.io/gateway/pm/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PolicysetsCreate**](PolicysetsApi.md#PolicysetsCreate) | **Post** /api/v1/policysets | 
[**PolicysetsDelete**](PolicysetsApi.md#PolicysetsDelete) | **Delete** /api/v1/policysets/{identifier} | 
[**PolicysetsFind**](PolicysetsApi.md#PolicysetsFind) | **Get** /api/v1/policysets/{identifier} | 
[**PolicysetsList**](PolicysetsApi.md#PolicysetsList) | **Get** /api/v1/policysets | 
[**PolicysetsUpdate**](PolicysetsApi.md#PolicysetsUpdate) | **Patch** /api/v1/policysets/{identifier} | 

# **PolicysetsCreate**
> PolicySet PolicysetsCreate(ctx, body, optional)


Create a policy set

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateRequestBody2**](CreateRequestBody2.md)|  | 
 **optional** | ***PolicysetsApiPolicysetsCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicysetsApiPolicysetsCreateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accountIdentifier** | **optional.**| Harness account ID | 
 **orgIdentifier** | **optional.**| Harness organization ID | 
 **projectIdentifier** | **optional.**| Harness project ID | 
 **xApiKey** | **optional.**| Harness PAT key used to perform authorization | 

### Return type

[**PolicySet**](PolicySet.md)

### Authorization

[api_key_header_x-api-key](../README.md#api_key_header_x-api-key), [jwt_header_Authorization](../README.md#jwt_header_Authorization)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PolicysetsDelete**
> PolicysetsDelete(ctx, identifier, optional)


Delete a policy set by identifier

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **identifier** | **string**| Identifier of the policy set | 
 **optional** | ***PolicysetsApiPolicysetsDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicysetsApiPolicysetsDeleteOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accountIdentifier** | **optional.String**| Harness account ID | 
 **orgIdentifier** | **optional.String**| Harness organization ID | 
 **projectIdentifier** | **optional.String**| Harness project ID | 
 **xApiKey** | **optional.String**| Harness PAT key used to perform authorization | 

### Return type

 (empty response body)

### Authorization

[api_key_header_x-api-key](../README.md#api_key_header_x-api-key), [jwt_header_Authorization](../README.md#jwt_header_Authorization)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PolicysetsFind**
> PolicySet PolicysetsFind(ctx, identifier, optional)


Find a policy set by identifier

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **identifier** | **string**| Identifier of the policy set to retrieve | 
 **optional** | ***PolicysetsApiPolicysetsFindOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicysetsApiPolicysetsFindOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accountIdentifier** | **optional.String**| Harness account ID | 
 **orgIdentifier** | **optional.String**| Harness organization ID | 
 **projectIdentifier** | **optional.String**| Harness project ID | 
 **xApiKey** | **optional.String**| Harness PAT key used to perform authorization | 

### Return type

[**PolicySet**](PolicySet.md)

### Authorization

[api_key_header_x-api-key](../README.md#api_key_header_x-api-key), [jwt_header_Authorization](../README.md#jwt_header_Authorization)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PolicysetsList**
> []PolicySet PolicysetsList(ctx, optional)


List all policy sets

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PolicysetsApiPolicysetsListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicysetsApiPolicysetsListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountIdentifier** | **optional.String**| Harness account ID | 
 **orgIdentifier** | **optional.String**| Harness organization ID | 
 **projectIdentifier** | **optional.String**| Harness project ID | 
 **perPage** | **optional.Int32**| Number of results per page | [default to 50]
 **page** | **optional.Int32**| Page number (starting from 0) | [default to 0]
 **searchTerm** | **optional.String**| Filter results by partial name match | 
 **sort** | **optional.String**| Sort order for results | [default to name,ASC]
 **type_** | **optional.String**| Filter results by type | 
 **action** | **optional.String**| Filter results by action | 
 **xApiKey** | **optional.String**| Harness PAT key used to perform authorization | 

### Return type

[**[]PolicySet**](PolicySet.md)

### Authorization

[api_key_header_x-api-key](../README.md#api_key_header_x-api-key), [jwt_header_Authorization](../README.md#jwt_header_Authorization)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PolicysetsUpdate**
> PolicysetsUpdate(ctx, body, identifier, optional)


Update a policy set by identifier

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateRequestBody2**](UpdateRequestBody2.md)|  | 
  **identifier** | **string**| Identifier of the policy set | 
 **optional** | ***PolicysetsApiPolicysetsUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicysetsApiPolicysetsUpdateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **accountIdentifier** | **optional.**| Harness account ID | 
 **orgIdentifier** | **optional.**| Harness organization ID | 
 **projectIdentifier** | **optional.**| Harness project ID | 
 **xApiKey** | **optional.**| Harness PAT key used to perform authorization | 

### Return type

 (empty response body)

### Authorization

[api_key_header_x-api-key](../README.md#api_key_header_x-api-key), [jwt_header_Authorization](../README.md#jwt_header_Authorization)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

