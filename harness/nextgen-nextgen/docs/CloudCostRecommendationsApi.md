# nextgen{{classname}}

All URIs are relative to *https://app.harness.io/gateway*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListRecommendations**](CloudCostRecommendationsApi.md#ListRecommendations) | **Post** /ccm/api/recommendation/overview/list | Return the list of Recommendations
[**RecommendationFilterValues**](CloudCostRecommendationsApi.md#RecommendationFilterValues) | **Post** /ccm/api/recommendation/overview/filter-values | Return the list of filter values for the Recommendations
[**RecommendationStats**](CloudCostRecommendationsApi.md#RecommendationStats) | **Post** /ccm/api/recommendation/overview/stats | Return Recommendations statistics
[**RecommendationsCount**](CloudCostRecommendationsApi.md#RecommendationsCount) | **Post** /ccm/api/recommendation/overview/count | Return the number of Recommendations

# **ListRecommendations**
> ResponseDtoRecommendations ListRecommendations(ctx, body, accountIdentifier)
Return the list of Recommendations

Returns the list of Cloud Cost Recommendations for the specified filters.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CcmRecommendationFilterProperties**](CcmRecommendationFilterProperties.md)| CCM Recommendations filter body. | 
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 

### Return type

[**ResponseDtoRecommendations**](ResponseDTORecommendations.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RecommendationFilterValues**
> ResponseDtoListFilterStats RecommendationFilterValues(ctx, body, accountIdentifier)
Return the list of filter values for the Recommendations

Returns the list of filter values for all the specified filters.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**FilterValues**](FilterValues.md)| Recommendation Filter Values Body. | 
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 

### Return type

[**ResponseDtoListFilterStats**](ResponseDTOListFilterStats.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RecommendationStats**
> ResponseDtoRecommendationOverviewStats RecommendationStats(ctx, body, accountIdentifier)
Return Recommendations statistics

Returns the Cloud Cost Recommendations statistics for the specified filters.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CcmRecommendationFilterProperties**](CcmRecommendationFilterProperties.md)| CCM Recommendations filter body. | 
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 

### Return type

[**ResponseDtoRecommendationOverviewStats**](ResponseDTORecommendationOverviewStats.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RecommendationsCount**
> ResponseDtoInteger RecommendationsCount(ctx, body, accountIdentifier)
Return the number of Recommendations

Returns the total number of Cloud Cost Recommendations based on the specified filters.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CcmRecommendationFilterProperties**](CcmRecommendationFilterProperties.md)| CCM Recommendations filter body. | 
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 

### Return type

[**ResponseDtoInteger**](ResponseDTOInteger.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

