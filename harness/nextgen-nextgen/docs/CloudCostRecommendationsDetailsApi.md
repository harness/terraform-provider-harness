# nextgen{{classname}}

All URIs are relative to *https://app.harness.io/gateway*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EcsRecommendationDetail**](CloudCostRecommendationsDetailsApi.md#EcsRecommendationDetail) | **Get** /ccm/api/recommendation/details/ecs-service | Return ECS Recommendation
[**NodeRecommendationDetail**](CloudCostRecommendationsDetailsApi.md#NodeRecommendationDetail) | **Get** /ccm/api/recommendation/details/node-pool | Return node pool Recommendation
[**WorkloadRecommendationDetail**](CloudCostRecommendationsDetailsApi.md#WorkloadRecommendationDetail) | **Get** /ccm/api/recommendation/details/workload | Return workload Recommendation

# **EcsRecommendationDetail**
> ResponseDtoecsRecommendationDto EcsRecommendationDetail(ctx, accountIdentifier, id, optional)
Return ECS Recommendation

Returns ECS Recommendation details for the given Recommendation identifier.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 
  **id** | **string**| ECS Recommendation identifier. | 
 **optional** | ***CloudCostRecommendationsDetailsApiEcsRecommendationDetailOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CloudCostRecommendationsDetailsApiEcsRecommendationDetailOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **from** | **optional.String**| Should use org.joda.time.DateTime parsable format. Example, &#x27;2022-01-31&#x27;, &#x27;2022-01-31T07:54Z&#x27; or &#x27;2022-01-31T07:54:51.264Z&#x27; Defaults to Today-7days | 
 **to** | **optional.String**| Should use org.joda.time.DateTime parsable format. Example, &#x27;2022-01-31&#x27;, &#x27;2022-01-31T07:54Z&#x27; or &#x27;2022-01-31T07:54:51.264Z&#x27; Defaults to Today | 

### Return type

[**ResponseDtoecsRecommendationDto**](ResponseDTOECSRecommendationDTO.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **NodeRecommendationDetail**
> ResponseDtoNodeRecommendationDto NodeRecommendationDetail(ctx, accountIdentifier, id)
Return node pool Recommendation

Returns node pool Recommendation details for the given identifier.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 
  **id** | **string**| Node pool Recommendation identifier | 

### Return type

[**ResponseDtoNodeRecommendationDto**](ResponseDTONodeRecommendationDTO.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WorkloadRecommendationDetail**
> ResponseDtoWorkloadRecommendationDto WorkloadRecommendationDetail(ctx, accountIdentifier, id, optional)
Return workload Recommendation

Returns workload Recommendation details for the given Recommendation identifier.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 
  **id** | **string**| Workload Recommendation identifier. | 
 **optional** | ***CloudCostRecommendationsDetailsApiWorkloadRecommendationDetailOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CloudCostRecommendationsDetailsApiWorkloadRecommendationDetailOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **from** | **optional.String**| Should use org.joda.time.DateTime parsable format. Example, &#x27;2022-01-31&#x27;, &#x27;2022-01-31T07:54Z&#x27; or &#x27;2022-01-31T07:54:51.264Z&#x27; Defaults to Today-7days | 
 **to** | **optional.String**| Should use org.joda.time.DateTime parsable format. Example, &#x27;2022-01-31&#x27;, &#x27;2022-01-31T07:54Z&#x27; or &#x27;2022-01-31T07:54:51.264Z&#x27; Defaults to Today | 

### Return type

[**ResponseDtoWorkloadRecommendationDto**](ResponseDTOWorkloadRecommendationDTO.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

