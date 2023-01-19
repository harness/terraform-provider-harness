# nextgen{{classname}}

All URIs are relative to *https://app.harness.io/gateway*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateFreeze**](FreezeCRUDApi.md#CreateFreeze) | **Post** /ng/api/freeze | Create a Freeze
[**CreateGlobalFreeze**](FreezeCRUDApi.md#CreateGlobalFreeze) | **Post** /ng/api/freeze/manageGlobalFreeze | Create Global Freeze
[**DeleteFreeze**](FreezeCRUDApi.md#DeleteFreeze) | **Delete** /ng/api/freeze/{freezeIdentifier} | Delete a Freeze
[**DeleteManyFreezes**](FreezeCRUDApi.md#DeleteManyFreezes) | **Post** /ng/api/freeze/delete | Delete many Freezes
[**GetFreeze**](FreezeCRUDApi.md#GetFreeze) | **Get** /ng/api/freeze/{freezeIdentifier} | Get a Freeze
[**GetFreezeList**](FreezeCRUDApi.md#GetFreezeList) | **Post** /ng/api/freeze/list | Gets Freeze list
[**GetGlobalFreeze**](FreezeCRUDApi.md#GetGlobalFreeze) | **Get** /ng/api/freeze/getGlobalFreeze | Get Global Freeze Yaml
[**UpdateFreeze**](FreezeCRUDApi.md#UpdateFreeze) | **Put** /ng/api/freeze/{freezeIdentifier} | Updates a Freeze
[**UpdateFreezeStatus**](FreezeCRUDApi.md#UpdateFreezeStatus) | **Post** /ng/api/freeze/updateFreezeStatus | Update the status of Freeze to active or inactive

# **CreateFreeze**
> ResponseDtoFreezeResponse CreateFreeze(ctx, body, accountIdentifier, optional)
Create a Freeze

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**string**](string.md)| Freeze YAML | 
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 
 **optional** | ***FreezeCRUDApiCreateFreezeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FreezeCRUDApiCreateFreezeOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **orgIdentifier** | **optional.**| Organization Identifier for the Entity. | 
 **projectIdentifier** | **optional.**| Project Identifier for the Entity. | 

### Return type

[**ResponseDtoFreezeResponse**](ResponseDTOFreezeResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateGlobalFreeze**
> ResponseDtoFreezeResponse CreateGlobalFreeze(ctx, body, accountIdentifier, optional)
Create Global Freeze

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**string**](string.md)| Freeze YAML | 
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 
 **optional** | ***FreezeCRUDApiCreateGlobalFreezeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FreezeCRUDApiCreateGlobalFreezeOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **orgIdentifier** | **optional.**| Organization Identifier for the Entity. | 
 **projectIdentifier** | **optional.**| Project Identifier for the Entity. | 

### Return type

[**ResponseDtoFreezeResponse**](ResponseDTOFreezeResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteFreeze**
> DeleteFreeze(ctx, accountIdentifier, freezeIdentifier, optional)
Delete a Freeze

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 
  **freezeIdentifier** | **string**| Freeze Identifier. | 
 **optional** | ***FreezeCRUDApiDeleteFreezeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FreezeCRUDApiDeleteFreezeOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity. | 
 **projectIdentifier** | **optional.String**| Project Identifier for the Entity. | 

### Return type

 (empty response body)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteManyFreezes**
> ResponseDtoFreezeResponseWrapperDto DeleteManyFreezes(ctx, accountIdentifier, optional)
Delete many Freezes

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 
 **optional** | ***FreezeCRUDApiDeleteManyFreezesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FreezeCRUDApiDeleteManyFreezesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of []string**](string.md)| List of Freeze Identifiers | 
 **orgIdentifier** | **optional.**| Organization Identifier for the Entity. | 
 **projectIdentifier** | **optional.**| Project Identifier for the Entity. | 

### Return type

[**ResponseDtoFreezeResponseWrapperDto**](ResponseDTOFreezeResponseWrapperDTO.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFreeze**
> ResponseDtoFreezeDetailedResponse GetFreeze(ctx, accountIdentifier, freezeIdentifier, optional)
Get a Freeze

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 
  **freezeIdentifier** | **string**| Freeze Identifier. | 
 **optional** | ***FreezeCRUDApiGetFreezeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FreezeCRUDApiGetFreezeOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity. | 
 **projectIdentifier** | **optional.String**| Project Identifier for the Entity. | 

### Return type

[**ResponseDtoFreezeDetailedResponse**](ResponseDTOFreezeDetailedResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFreezeList**
> ResponseDtoPageResponseFreezeSummaryResponse GetFreezeList(ctx, accountIdentifier, optional)
Gets Freeze list

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 
 **optional** | ***FreezeCRUDApiGetFreezeListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FreezeCRUDApiGetFreezeListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of FreezeFilterPropertiesDto**](FreezeFilterPropertiesDto.md)| This contains details of Freeze filters | 
 **page** | **optional.**| Page Index of the results to fetch.Default Value: 0 | [default to 0]
 **size** | **optional.**| Results per page | [default to 10]
 **orgIdentifier** | **optional.**| Organization Identifier for the Entity. | 
 **projectIdentifier** | **optional.**| Project Identifier for the Entity. | 

### Return type

[**ResponseDtoPageResponseFreezeSummaryResponse**](ResponseDTOPageResponseFreezeSummaryResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetGlobalFreeze**
> ResponseDtoFreezeDetailedResponse GetGlobalFreeze(ctx, accountIdentifier, optional)
Get Global Freeze Yaml

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 
 **optional** | ***FreezeCRUDApiGetGlobalFreezeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FreezeCRUDApiGetGlobalFreezeOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity. | 
 **projectIdentifier** | **optional.String**| Project Identifier for the Entity. | 

### Return type

[**ResponseDtoFreezeDetailedResponse**](ResponseDTOFreezeDetailedResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateFreeze**
> ResponseDtoFreezeResponse UpdateFreeze(ctx, body, accountIdentifier, freezeIdentifier, optional)
Updates a Freeze

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**string**](string.md)| Freeze YAML | 
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 
  **freezeIdentifier** | **string**| Freeze Identifier. | 
 **optional** | ***FreezeCRUDApiUpdateFreezeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FreezeCRUDApiUpdateFreezeOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **orgIdentifier** | **optional.**| Organization Identifier for the Entity. | 
 **projectIdentifier** | **optional.**| Project Identifier for the Entity. | 

### Return type

[**ResponseDtoFreezeResponse**](ResponseDTOFreezeResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateFreezeStatus**
> ResponseDtoFreezeResponseWrapperDto UpdateFreezeStatus(ctx, accountIdentifier, status, optional)
Update the status of Freeze to active or inactive

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 
  **status** | **string**| Freeze YAML | 
 **optional** | ***FreezeCRUDApiUpdateFreezeStatusOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FreezeCRUDApiUpdateFreezeStatusOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of []string**](string.md)| Comma seperated List of Freeze Identifiers | 
 **orgIdentifier** | **optional.**| Organization Identifier for the Entity. | 
 **projectIdentifier** | **optional.**| Project Identifier for the Entity. | 

### Return type

[**ResponseDtoFreezeResponseWrapperDto**](ResponseDTOFreezeResponseWrapperDTO.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

