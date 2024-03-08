# nextgen{{classname}}

All URIs are relative to *https://app.harness.io/gateway*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateInfrastructure**](InfrastructuresApi.md#CreateInfrastructure) | **Post** /ng/api/infrastructures | Create an Infrastructure in an Environment
[**DeleteInfrastructure**](InfrastructuresApi.md#DeleteInfrastructure) | **Delete** /ng/api/infrastructures/{infraIdentifier} | Delete an Infrastructure by identifier
[**GetInfrastructure**](InfrastructuresApi.md#GetInfrastructure) | **Get** /ng/api/infrastructures/{infraIdentifier} | Gets an Infrastructure by identifier
[**GetInfrastructureList**](InfrastructuresApi.md#GetInfrastructureList) | **Get** /ng/api/infrastructures | Gets Infrastructure list
[**UpdateInfrastructure**](InfrastructuresApi.md#UpdateInfrastructure) | **Put** /ng/api/infrastructures | Update an Infrastructure by identifier
[**ImportInfrastructure**](InfrastructuresApi.md#ImportInfrastructure) | **Post** ng/api/infrastructures/import | Get Infrastructure YAML from Git Repository

# **CreateInfrastructure**
> ResponseDtoInfrastructureResponse CreateInfrastructure(ctx, accountIdentifier, optional)
Create an Infrastructure in an Environment

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 
 **optional** | ***InfrastructuresApiCreateInfrastructureOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InfrastructuresApiCreateInfrastructureOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of InfrastructureRequest**](InfrastructureRequest.md)| Details of the Infrastructure to be created | 

### Return type

[**ResponseDtoInfrastructureResponse**](ResponseDTOInfrastructureResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteInfrastructure**
> ResponseDtoBoolean DeleteInfrastructure(ctx, infraIdentifier, accountIdentifier, environmentIdentifier, optional)
Delete an Infrastructure by identifier

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **infraIdentifier** | **string**| Infrastructure Identifier for the entity | 
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 
  **environmentIdentifier** | **string**| Environment Identifier for the Entity. | 
 **optional** | ***InfrastructuresApiDeleteInfrastructureOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InfrastructuresApiDeleteInfrastructureOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity. | 
 **projectIdentifier** | **optional.String**| Project Identifier for the Entity. | 

### Return type

[**ResponseDtoBoolean**](ResponseDTOBoolean.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetInfrastructure**
> ResponseDtoInfrastructureResponse GetInfrastructure(ctx, infraIdentifier, accountIdentifier, environmentIdentifier, optional)
Gets an Infrastructure by identifier

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **infraIdentifier** | **string**| Infrastructure Identifier for the entity | 
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 
  **environmentIdentifier** | **string**| envId | 
 **optional** | ***InfrastructuresApiGetInfrastructureOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InfrastructuresApiGetInfrastructureOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity. | 
 **projectIdentifier** | **optional.String**| Project Identifier for the Entity. | 
 **deleted** | **optional.Bool**| Specify whether Infrastructure is deleted or not | [default to false]

### Return type

[**ResponseDtoInfrastructureResponse**](ResponseDTOInfrastructureResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetInfrastructureList**
> ResponseDtoPageResponseInfrastructureResponse GetInfrastructureList(ctx, accountIdentifier, environmentIdentifier, optional)
Gets Infrastructure list

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 
  **environmentIdentifier** | **string**| Environment Identifier for the Entity. | 
 **optional** | ***InfrastructuresApiGetInfrastructureListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InfrastructuresApiGetInfrastructureListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **optional.Int32**| Page Index of the results to fetch.Default Value: 0 | [default to 0]
 **size** | **optional.Int32**| Results per page | [default to 100]
 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity. | 
 **projectIdentifier** | **optional.String**| Project Identifier for the Entity. | 
 **searchTerm** | **optional.String**| The word to be searched and included in the list response | 
 **infraIdentifiers** | [**optional.Interface of []string**](string.md)| List of InfrastructureIds | 
 **deploymentType** | **optional.String**|  | 
 **deploymentTemplateIdentifier** | **optional.String**|  | 
 **versionLabel** | **optional.String**|  | 
 **sort** | [**optional.Interface of []string**](string.md)| Specifies the sorting criteria of the list. Like sorting based on the last updated entity, alphabetical sorting in an ascending or descending order | 

### Return type

[**ResponseDtoPageResponseInfrastructureResponse**](ResponseDTOPageResponseInfrastructureResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateInfrastructure**
> ResponseDtoInfrastructureResponse UpdateInfrastructure(ctx, accountIdentifier, optional)
Update an Infrastructure by identifier

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 
 **optional** | ***InfrastructuresApiUpdateInfrastructureOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InfrastructuresApiUpdateInfrastructureOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of InfrastructureRequest**](InfrastructureRequest.md)| Details of the Infrastructure to be updated | 

### Return type

[**ResponseDtoInfrastructureResponse**](ResponseDTOInfrastructureResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ImportInfrastructure**
> ResponseInfrastructureImportResponse ImportInfrastructure(ctx, accountIdentifier, optional)
Get Infrastructure YAML from Git Repository

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**|  | 
 **optional** | ***InfrastructuresApiImportInfrastructureOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InfrastructuresApiImportInfrastructureOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **orgIdentifier** | **optional.String**|  | 
 **projectIdentifier** | **optional.String**|  | 
 **infraIdentifier** | **optional.String**|  | 
 **connectorRef** | **optional.String**|  | 
 **repoName** | **optional.String**|  | 
 **branch** | **optional.String**|  | 
 **filePath** | **optional.String**|  | 
 **isForceImport** | **optional.Bool**|  | [default to false]
 **isHarnessCodeRepo** | **optional.Bool**|  | 

### Return type

[**ResponseInfrastructureImportResponse**](ResponseInfrastructureImportResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)