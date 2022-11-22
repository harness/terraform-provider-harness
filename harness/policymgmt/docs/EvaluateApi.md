# {{classname}}

All URIs are relative to *https://app.harness.io/gateway/pm/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EvaluateEvaluate**](EvaluateApi.md#EvaluateEvaluate) | **Post** /api/v1/evaluate | 
[**EvaluateEvaluateByIds**](EvaluateApi.md#EvaluateEvaluateByIds) | **Post** /api/v1/evaluate-by-ids | 
[**EvaluateEvaluateByType**](EvaluateApi.md#EvaluateEvaluateByType) | **Post** /api/v1/evaluate-by-type | 

# **EvaluateEvaluate**
> EvaluatedPolicy EvaluateEvaluate(ctx, body, optional)


Evaluate arbitrary rego

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**EvaluateRequestBody**](EvaluateRequestBody.md)|  | 
 **optional** | ***EvaluateApiEvaluateEvaluateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EvaluateApiEvaluateEvaluateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accountIdentifier** | **optional.**| Harness account ID | 
 **orgIdentifier** | **optional.**| Harness organization ID | 
 **projectIdentifier** | **optional.**| Harness project ID | 

### Return type

[**EvaluatedPolicy**](EvaluatedPolicy.md)

### Authorization

[jwt_header_Authorization](../README.md#jwt_header_Authorization)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EvaluateEvaluateByIds**
> Evaluation EvaluateEvaluateByIds(ctx, ids, optional)


Evaluate policy sets by ID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ids** | **string**| Comma-separated list of identifiers for the policy sets that should be evaluated, with account. or org. prefixes if needed | 
 **optional** | ***EvaluateApiEvaluateEvaluateByIdsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EvaluateApiEvaluateEvaluateByIdsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accountIdentifier** | **optional.String**| Harness account ID | 
 **orgIdentifier** | **optional.String**| Harness organization ID | 
 **projectIdentifier** | **optional.String**| Harness project ID | 
 **entity** | **optional.String**| User-supplied global identifier of the entity under evaluation | 
 **entityMetadata** | **optional.String**| User-supplied additional metadata for the entity under evaluation | 
 **principalIdentifier** | **optional.String**| Identifier of the principal that triggered the evaluation - must be specified in conjunction with &#x27;principalType&#x27; | 
 **principalType** | **optional.String**| Type of principal that triggered the evaluation - must be specified in conjunction with &#x27;principalIdentifier&#x27; | 
 **userIdentifier** | **optional.String**| Deprecated: Please use &#x27;principalIdentifier&#x27; and &#x27;principalType&#x27; instead | 

### Return type

[**Evaluation**](Evaluation.md)

### Authorization

[jwt_header_Authorization](../README.md#jwt_header_Authorization)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EvaluateEvaluateByType**
> Evaluation EvaluateEvaluateByType(ctx, type_, action, optional)


Evaluate all policy sets of a specified type

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **type_** | **string**| Type of entity under evaluation | 
  **action** | **string**| Action that triggered the evaluation | 
 **optional** | ***EvaluateApiEvaluateEvaluateByTypeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EvaluateApiEvaluateEvaluateByTypeOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **accountIdentifier** | **optional.String**| Harness account ID | 
 **orgIdentifier** | **optional.String**| Harness organization ID | 
 **projectIdentifier** | **optional.String**| Harness project ID | 
 **entity** | **optional.String**| User-supplied global identifier of the entity under evaluation | 
 **entityMetadata** | **optional.String**| User-supplied additional metadata for the entity under evaluation | 
 **principalIdentifier** | **optional.String**| Identifier of the principal that triggered the evaluation - must be specified in conjunction with &#x27;principalType&#x27; | 
 **principalType** | **optional.String**| Type of principal that triggered the evaluation - must be specified in conjunction with &#x27;principalIdentifier&#x27; | 
 **userIdentifier** | **optional.String**| Deprecated: Please use &#x27;principalIdentifier&#x27; and &#x27;principalType&#x27; instead | 

### Return type

[**Evaluation**](Evaluation.md)

### Authorization

[jwt_header_Authorization](../README.md#jwt_header_Authorization)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

