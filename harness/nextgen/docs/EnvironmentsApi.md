# {{classname}}

All URIs are relative to *https://app.harness.io/gateway*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateEnvironmentV2**](EnvironmentsApi.md#CreateEnvironmentV2) | **Post** /ng/api/environmentsV2 | Create an Environment
[**DeleteEnvironmentV2**](EnvironmentsApi.md#DeleteEnvironmentV2) | **Delete** /ng/api/environmentsV2/{environmentIdentifier} | Delete an Environment by identifier
[**GetEnvironmentAccessList**](EnvironmentsApi.md#GetEnvironmentAccessList) | **Get** /ng/api/environmentsV2/list/access | Gets Environment Access list
[**GetEnvironmentList**](EnvironmentsApi.md#GetEnvironmentList) | **Get** /ng/api/environmentsV2 | Gets Environment list for a project
[**GetEnvironmentV2**](EnvironmentsApi.md#GetEnvironmentV2) | **Get** /ng/api/environmentsV2/{environmentIdentifier} | Gets an Environment by identifier
[**UpdateEnvironmentV2**](EnvironmentsApi.md#UpdateEnvironmentV2) | **Put** /ng/api/environmentsV2 | Update an Environment by identifier
[**UpsertEnvironmentV2**](EnvironmentsApi.md#UpsertEnvironmentV2) | **Put** /ng/api/environmentsV2/upsert | Upsert an Environment by identifier

# **CreateEnvironmentV2**
> ResponseDtoEnvironmentResponse CreateEnvironmentV2(ctx, accountIdentifier, optional)
Create an Environment

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity | 
 **optional** | ***EnvironmentsApiCreateEnvironmentV2Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EnvironmentsApiCreateEnvironmentV2Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of EnvironmentRequest**](EnvironmentRequest.md)| Details of the Environment to be created | 

### Return type

[**ResponseDtoEnvironmentResponse**](ResponseDTOEnvironmentResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteEnvironmentV2**
> ResponseDtoBoolean DeleteEnvironmentV2(ctx, environmentIdentifier, accountIdentifier, optional)
Delete an Environment by identifier

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **environmentIdentifier** | **string**| Environment Identifier for the entity | 
  **accountIdentifier** | **string**| Account Identifier for the Entity | 
 **optional** | ***EnvironmentsApiDeleteEnvironmentV2Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EnvironmentsApiDeleteEnvironmentV2Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **ifMatch** | **optional.String**|  | 
 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity | 
 **projectIdentifier** | **optional.String**| Project Identifier for the Entity | 

### Return type

[**ResponseDtoBoolean**](ResponseDTOBoolean.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEnvironmentAccessList**
> ResponseDtoListEnvironmentResponse GetEnvironmentAccessList(ctx, accountIdentifier, optional)
Gets Environment Access list

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity | 
 **optional** | ***EnvironmentsApiGetEnvironmentAccessListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EnvironmentsApiGetEnvironmentAccessListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| page | [default to 0]
 **size** | **optional.Int32**| size | [default to 100]
 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity | 
 **projectIdentifier** | **optional.String**| Project Identifier for the Entity | 
 **searchTerm** | **optional.String**| The word to be searched and included in the list response | 
 **envIdentifiers** | [**optional.Interface of []string**](string.md)| List of EnvironmentIds | 
 **sort** | [**optional.Interface of []string**](string.md)| Specifies sorting criteria of the list. Like sorting based on the last updated entity, alphabetical sorting in an ascending or descending order | 

### Return type

[**ResponseDtoListEnvironmentResponse**](ResponseDTOListEnvironmentResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEnvironmentList**
> ResponseDtoPageResponseEnvironmentResponse GetEnvironmentList(ctx, accountIdentifier, optional)
Gets Environment list for a project

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity | 
 **optional** | ***EnvironmentsApiGetEnvironmentListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EnvironmentsApiGetEnvironmentListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| Indicates the number of pages. Results for these pages will be retrieved. | [default to 0]
 **size** | **optional.Int32**| The number of the elements to fetch | [default to 100]
 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity | 
 **projectIdentifier** | **optional.String**| Project Identifier for the Entity | 
 **searchTerm** | **optional.String**| The word to be searched and included in the list response | 
 **envIdentifiers** | [**optional.Interface of []string**](string.md)| List of EnvironmentIds | 
 **sort** | [**optional.Interface of []string**](string.md)| Specifies sorting criteria of the list. Like sorting based on the last updated entity, alphabetical sorting in an ascending or descending order | 

### Return type

[**ResponseDtoPageResponseEnvironmentResponse**](ResponseDTOPageResponseEnvironmentResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEnvironmentV2**
> ResponseDtoEnvironmentResponse GetEnvironmentV2(ctx, environmentIdentifier, accountIdentifier, optional)
Gets an Environment by identifier

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **environmentIdentifier** | **string**| Environment Identifier for the entity | 
  **accountIdentifier** | **string**| Account Identifier for the Entity | 
 **optional** | ***EnvironmentsApiGetEnvironmentV2Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EnvironmentsApiGetEnvironmentV2Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity | 
 **projectIdentifier** | **optional.String**| Project Identifier for the Entity | 
 **deleted** | **optional.Bool**| Specify whether Environment is deleted or not | [default to false]

### Return type

[**ResponseDtoEnvironmentResponse**](ResponseDTOEnvironmentResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateEnvironmentV2**
> ResponseDtoEnvironmentResponse UpdateEnvironmentV2(ctx, accountIdentifier, optional)
Update an Environment by identifier

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity | 
 **optional** | ***EnvironmentsApiUpdateEnvironmentV2Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EnvironmentsApiUpdateEnvironmentV2Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of EnvironmentRequest**](EnvironmentRequest.md)| Details of the Environment to be updated | 
 **ifMatch** | **optional.**|  | 

### Return type

[**ResponseDtoEnvironmentResponse**](ResponseDTOEnvironmentResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpsertEnvironmentV2**
> ResponseDtoEnvironmentResponse UpsertEnvironmentV2(ctx, accountIdentifier, optional)
Upsert an Environment by identifier

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity | 
 **optional** | ***EnvironmentsApiUpsertEnvironmentV2Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EnvironmentsApiUpsertEnvironmentV2Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of EnvironmentRequest**](EnvironmentRequest.md)| Details of the Environment to be updated | 
 **ifMatch** | **optional.**|  | 

### Return type

[**ResponseDtoEnvironmentResponse**](ResponseDTOEnvironmentResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

