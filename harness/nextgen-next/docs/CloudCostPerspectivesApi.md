# {{classname}}

All URIs are relative to *https://app.harness.io/gateway*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePerspective**](CloudCostPerspectivesApi.md#CreatePerspective) | **Post** /ccm/api/perspective | Create a Perspective
[**DeletePerspective**](CloudCostPerspectivesApi.md#DeletePerspective) | **Delete** /ccm/api/perspective | Delete a Perspective by identifier
[**GetForecastCostV2**](CloudCostPerspectivesApi.md#GetForecastCostV2) | **Get** /ccm/api/perspective/forecastCost | Get the forecasted cost of a Perspective
[**GetLastMonthCostV2**](CloudCostPerspectivesApi.md#GetLastMonthCostV2) | **Get** /ccm/api/perspective/lastMonthCost | Get the last month cost for a Perspective
[**GetPerspective**](CloudCostPerspectivesApi.md#GetPerspective) | **Get** /ccm/api/perspective | Get a Perspective by identifier
[**UpdatePerspective**](CloudCostPerspectivesApi.md#UpdatePerspective) | **Put** /ccm/api/perspective | Update an existing Perspective

# **CreatePerspective**
> ResponseDtoceView CreatePerspective(ctx, body, accountIdentifier, clone)
Create a Perspective

Create a Perspective, accepts a url param 'clone' which decides whether the Perspective being created should be a clone of existing Perspective, and a Request Body with the PerspectiveDefinition

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CeView**](CeView.md)| Request body containing Perspective&#x27;s CEView object to create | 
  **accountIdentifier** | **string**| Account Identifier for the entity | 
  **clone** | **bool**| Whether the Perspective being created should be a clone of existing Perspective, if true we will ignore the uuid field in the request body and create a completely new Perspective | 

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
Delete a Perspective by identifier

Deletes a perspective by identifier, it accepts a mandatory CEView's identifier as url param and returns a test response on successful deletion

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the entity | 
  **perspectiveId** | **string**| The identifier of the CEView object to delete | 

### Return type

[**ResponseDtoString**](ResponseDTOString.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetForecastCostV2**
> ResponseDtoDouble GetForecastCostV2(ctx, accountIdentifier, perspectiveId)
Get the forecasted cost of a Perspective

Get the forecasted cost of a Perspective for next 30 days

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the entity | 
  **perspectiveId** | **string**| The Perspective identifier for which we want the forecast cost | 

### Return type

[**ResponseDtoDouble**](ResponseDTODouble.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLastMonthCostV2**
> ResponseDtoDouble GetLastMonthCostV2(ctx, accountIdentifier, perspectiveId)
Get the last month cost for a Perspective

Get last month cost for a Perspective

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the entity | 
  **perspectiveId** | **string**| The Perspective identifier for which we want the last month cost | 

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
Get a Perspective by identifier

Get complete CEView object by Perspective identifier passed as a url param

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the entity | 
  **perspectiveId** | **string**| The identifier of the Perspective to fetch | 

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
Update an existing Perspective

Update an existing Perspective, it accepts a CEView and upserts it using the uuid mentioned in the definition

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CeView**](CeView.md)| Request body containing Perspective&#x27;s CEView object to update | 
  **accountIdentifier** | **string**| Account Identifier for the entity | 

### Return type

[**ResponseDtoceView**](ResponseDTOCEView.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

