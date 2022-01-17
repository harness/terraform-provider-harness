# {{classname}}

All URIs are relative to *https://app.harness.io/gateway*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteFilter**](FilterApi.md#DeleteFilter) | **Delete** /ng/api/filters/{identifier} | Delete a Filter by identifier
[**GetConnectorListV21**](FilterApi.md#GetConnectorListV21) | **Get** /ng/api/filters | Get the list of Filters satisfying the criteria (if any) in the request
[**GetFilter**](FilterApi.md#GetFilter) | **Get** /ng/api/filters/{identifier} | Gets a Filter by identifier
[**PipelinedeleteFilter**](FilterApi.md#PipelinedeleteFilter) | **Delete** /pipeline/api/filters/{identifier} | Delete a Filter by identifier
[**PipelinegetConnectorListV2**](FilterApi.md#PipelinegetConnectorListV2) | **Get** /pipeline/api/filters | Get the list of Filters satisfying the criteria (if any) in the request
[**PipelinegetFilter**](FilterApi.md#PipelinegetFilter) | **Get** /pipeline/api/filters/{identifier} | Gets a Filter by identifier
[**PipelinepostFilter**](FilterApi.md#PipelinepostFilter) | **Post** /pipeline/api/filters | Creates a Filter
[**PipelineupdateFilter**](FilterApi.md#PipelineupdateFilter) | **Put** /pipeline/api/filters | Updates the Filter
[**PostFilter**](FilterApi.md#PostFilter) | **Post** /ng/api/filters | Creates a Filter
[**UpdateFilter**](FilterApi.md#UpdateFilter) | **Put** /ng/api/filters | Updates the Filter

# **DeleteFilter**
> ResponseDtoBoolean DeleteFilter(ctx, accountIdentifier, identifier, type_, optional)
Delete a Filter by identifier

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity | 
  **identifier** | **string**| Filter Identifier | 
  **type_** | **string**| Type of Filter | 
 **optional** | ***FilterApiDeleteFilterOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FilterApiDeleteFilterOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity | 
 **projectIdentifier** | **optional.String**| Project Identifier for the Entity | 

### Return type

[**ResponseDtoBoolean**](ResponseDTOBoolean.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml, text/yaml, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetConnectorListV21**
> ResponseDtoPageResponseFilter GetConnectorListV21(ctx, accountIdentifier, type_, optional)
Get the list of Filters satisfying the criteria (if any) in the request

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity | 
  **type_** | **string**| Type of Filter | 
 **optional** | ***FilterApiGetConnectorListV21Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FilterApiGetConnectorListV21Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pageIndex** | **optional.Int32**| Page number of navigation. If left empty, default value of 0 is assumed | [default to 0]
 **pageSize** | **optional.Int32**| Number of entries per page. If left empty, default value of 100 is assumed | [default to 100]
 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity | 
 **projectIdentifier** | **optional.String**| Project Identifier for the Entity | 

### Return type

[**ResponseDtoPageResponseFilter**](ResponseDTOPageResponseFilter.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml, text/yaml, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFilter**
> ResponseDtoFilter GetFilter(ctx, accountIdentifier, identifier, type_, optional)
Gets a Filter by identifier

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity | 
  **identifier** | **string**| Filter Identifier | 
  **type_** | **string**| Type of Filter | 
 **optional** | ***FilterApiGetFilterOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FilterApiGetFilterOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity | 
 **projectIdentifier** | **optional.String**| Project Identifier for the Entity | 

### Return type

[**ResponseDtoFilter**](ResponseDTOFilter.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml, text/yaml, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PipelinedeleteFilter**
> ResponseDtoBoolean PipelinedeleteFilter(ctx, accountIdentifier, identifier, type_, optional)
Delete a Filter by identifier

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the entity | 
  **identifier** | **string**| Filter Identifier | 
  **type_** | **string**| Type of Filter | 
 **optional** | ***FilterApiPipelinedeleteFilterOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FilterApiPipelinedeleteFilterOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **orgIdentifier** | **optional.String**| Organization Identifier for the entity | 
 **projectIdentifier** | **optional.String**| Project Identifier for the entity | 

### Return type

[**ResponseDtoBoolean**](ResponseDTOBoolean.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml, text/yaml, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PipelinegetConnectorListV2**
> ResponseDtoPageResponseFilter PipelinegetConnectorListV2(ctx, accountIdentifier, type_, optional)
Get the list of Filters satisfying the criteria (if any) in the request

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the entity | 
  **type_** | **string**| Type of Filter | 
 **optional** | ***FilterApiPipelinegetConnectorListV2Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FilterApiPipelinegetConnectorListV2Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pageIndex** | **optional.Int32**| Page number of navigation. If left empty, default value of 0 is assumed | [default to 0]
 **pageSize** | **optional.Int32**| Number of entries per page. If left empty, default value of 100 is assumed | [default to 100]
 **orgIdentifier** | **optional.String**| Organization Identifier for the entity | 
 **projectIdentifier** | **optional.String**| Project Identifier for the entity | 

### Return type

[**ResponseDtoPageResponseFilter**](ResponseDTOPageResponseFilter.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml, text/yaml, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PipelinegetFilter**
> ResponseDtoFilter PipelinegetFilter(ctx, accountIdentifier, identifier, type_, optional)
Gets a Filter by identifier

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the entity | 
  **identifier** | **string**| Filter Identifier | 
  **type_** | **string**| Type of Filter | 
 **optional** | ***FilterApiPipelinegetFilterOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FilterApiPipelinegetFilterOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **orgIdentifier** | **optional.String**| Organization Identifier for the entity | 
 **projectIdentifier** | **optional.String**| Project Identifier for the entity | 

### Return type

[**ResponseDtoFilter**](ResponseDTOFilter.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml, text/yaml, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PipelinepostFilter**
> ResponseDtoFilter PipelinepostFilter(ctx, body, accountIdentifier)
Creates a Filter

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Filter**](Filter.md)| Details of the Connector to create | 
  **accountIdentifier** | **string**| Account Identifier for the entity | 

### Return type

[**ResponseDtoFilter**](ResponseDTOFilter.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json, text/yaml, text/html, text/plain
 - **Accept**: application/json, application/yaml, text/yaml, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PipelineupdateFilter**
> ResponseDtoFilter PipelineupdateFilter(ctx, body, accountIdentifier)
Updates the Filter

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Filter**](Filter.md)| This is the updated Filter. This should have all the fields not just the updated ones | 
  **accountIdentifier** | **string**| Account Identifier for the entity | 

### Return type

[**ResponseDtoFilter**](ResponseDTOFilter.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json, text/yaml, text/html, text/plain
 - **Accept**: application/json, application/yaml, text/yaml, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostFilter**
> ResponseDtoFilter PostFilter(ctx, body, accountIdentifier)
Creates a Filter

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Filter**](Filter.md)| Details of the Connector to create | 
  **accountIdentifier** | **string**| Account Identifier for the Entity | 

### Return type

[**ResponseDtoFilter**](ResponseDTOFilter.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json, text/yaml, text/html, text/plain
 - **Accept**: application/json, application/yaml, text/yaml, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateFilter**
> ResponseDtoFilter UpdateFilter(ctx, body, accountIdentifier)
Updates the Filter

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Filter**](Filter.md)| This is the updated Filter. This should have all the fields not just the updated ones | 
  **accountIdentifier** | **string**| Account Identifier for the Entity | 

### Return type

[**ResponseDtoFilter**](ResponseDTOFilter.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json, text/yaml, text/html, text/plain
 - **Accept**: application/json, application/yaml, text/yaml, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

