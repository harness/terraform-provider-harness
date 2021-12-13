# {{classname}}

All URIs are relative to *https://app.harness.io/gateway/authz/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EvaluateCustomFeatureRestriction**](DefaultApi.md#EvaluateCustomFeatureRestriction) | **Put** /enforcement/client/custom/{featureRestrictionName} | 
[**GetFeatureUsage**](DefaultApi.md#GetFeatureUsage) | **Put** /enforcement/client/usage/{featureRestrictionName} | 

# **EvaluateCustomFeatureRestriction**
> ResponseDtoBoolean EvaluateCustomFeatureRestriction(ctx, body, featureRestrictionName, accountIdentifier)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CustomRestrictionEvaluationDto**](CustomRestrictionEvaluationDto.md)|  | 
  **featureRestrictionName** | **string**|  | 
  **accountIdentifier** | **string**|  | 

### Return type

[**ResponseDtoBoolean**](ResponseDTOBoolean.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFeatureUsage**
> ResponseDtoFeatureRestrictionUsageDto GetFeatureUsage(ctx, body, featureRestrictionName, accountIdentifier)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RestrictionMetadata**](RestrictionMetadata.md)|  | 
  **featureRestrictionName** | **string**|  | 
  **accountIdentifier** | **string**|  | 

### Return type

[**ResponseDtoFeatureRestrictionUsageDto**](ResponseDTOFeatureRestrictionUsageDTO.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

