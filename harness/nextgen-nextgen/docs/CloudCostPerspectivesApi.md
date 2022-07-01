# nextgen{{classname}}

All URIs are relative to *https://app.harness.io/gateway*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePerspective**](CloudCostPerspectivesApi.md#CreatePerspective) | **Post** /ccm/api/perspective | Create a Perspective
[**DeletePerspective**](CloudCostPerspectivesApi.md#DeletePerspective) | **Delete** /ccm/api/perspective | Delete a Perspective
[**GetAllPerspectives**](CloudCostPerspectivesApi.md#GetAllPerspectives) | **Get** /ccm/api/perspective/getAllPerspectives | Return details of all the Perspectives
[**GetLastPeriodCost**](CloudCostPerspectivesApi.md#GetLastPeriodCost) | **Get** /ccm/api/perspective/lastPeriodCost | Get the last period cost for a Perspective
[**GetPerspective**](CloudCostPerspectivesApi.md#GetPerspective) | **Get** /ccm/api/perspective | Fetch details of a Perspective
[**UpdatePerspective**](CloudCostPerspectivesApi.md#UpdatePerspective) | **Put** /ccm/api/perspective | Update a Perspective

# **CreatePerspective**
> ResponseDtoceView CreatePerspective(ctx, body, accountIdentifier, clone)
Create a Perspective

Create a Perspective. You can set the clone parameter as true to clone a Perspective.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CeView**](CeView.md)| Request body containing Perspective&#x27;s CEView object | 
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 
  **clone** | **bool**| Set the clone parameter as true to clone a Perspective. | 

### Return type

[**ResponseDtoceView**](ResponseDTOCEView.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeletePerspective**
> ResponseDtoString DeletePerspective(ctx, accountIdentifier, perspectiveId)
Delete a Perspective

Delete a Perspective for the given Perspective ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 
  **perspectiveId** | **string**| Unique identifier for the Perspective | 

### Return type

[**ResponseDtoString**](ResponseDTOString.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllPerspectives**
> ResponseDtoListPerspective GetAllPerspectives(ctx, accountIdentifier)
Return details of all the Perspectives

Return details of all the Perspectives for the given account ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 

### Return type

[**ResponseDtoListPerspective**](ResponseDTOListPerspective.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLastPeriodCost**
> ResponseDtoDouble GetLastPeriodCost(ctx, accountIdentifier, perspectiveId, startTime, period)
Get the last period cost for a Perspective

Get last period cost for a Perspective

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 
  **perspectiveId** | **string**| The Perspective identifier for which we want the cost | 
  **startTime** | **int64**| The Start time (timestamp in millis) for the current period | 
  **period** | **string**| The period (DAILY, WEEKLY, MONTHLY, QUARTERLY, YEARLY) for which we want the cost | 

### Return type

[**ResponseDtoDouble**](ResponseDTODouble.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPerspective**
> ResponseDtoceView GetPerspective(ctx, accountIdentifier, perspectiveId)
Fetch details of a Perspective

Fetch details of a Perspective for the given Perspective ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 
  **perspectiveId** | **string**| Unique identifier for the Perspective | 

### Return type

[**ResponseDtoceView**](ResponseDTOCEView.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdatePerspective**
> ResponseDtoceView UpdatePerspective(ctx, body, accountIdentifier)
Update a Perspective

Update a Perspective. It accepts a CEView object and upserts it using the uuid mentioned in the definition.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CeView**](CeView.md)| Perspective&#x27;s CEView object | 
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 

### Return type

[**ResponseDtoceView**](ResponseDTOCEView.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

