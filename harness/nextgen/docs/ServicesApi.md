# nextgen{{classname}}

All URIs are relative to *https://app.harness.io/gateway*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateServiceV2**](ServicesApi.md#CreateServiceV2) | **Post** /ng/api/servicesV2 | Create a Service
[**CreateServicesV2**](ServicesApi.md#CreateServicesV2) | **Post** /ng/api/servicesV2/batch | Create Services
[**DeleteServiceV2**](ServicesApi.md#DeleteServiceV2) | **Delete** /ng/api/servicesV2/{serviceIdentifier} | Delete a Service by identifier
[**GetServiceAccessList**](ServicesApi.md#GetServiceAccessList) | **Get** /ng/api/servicesV2/list/access | Gets Service Access list
[**GetServiceList**](ServicesApi.md#GetServiceList) | **Get** /ng/api/servicesV2 | Gets Service list
[**GetServiceV2**](ServicesApi.md#GetServiceV2) | **Get** /ng/api/servicesV2/{serviceIdentifier} | Gets a Service by identifier
[**UpdateServiceV2**](ServicesApi.md#UpdateServiceV2) | **Put** /ng/api/servicesV2 | Update a Service by identifier
[**UpsertServiceV2**](ServicesApi.md#UpsertServiceV2) | **Put** /ng/api/servicesV2/upsert | Upsert a Service by identifier
[**ImportService**](ServicesApi.md#ImportService) | **Post** /servicesV2/import | Get Service YAML from Git Repository

# **CreateServiceV2**
> ResponseDtoServiceResponse CreateServiceV2(ctx, accountIdentifier, optional)
Create a Service

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 
 **optional** | ***ServicesApiCreateServiceV2Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ServicesApiCreateServiceV2Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of ServiceRequest**](ServiceRequest.md)| Details of the Service to be created | 

### Return type

[**ResponseDtoServiceResponse**](ResponseDTOServiceResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateServicesV2**
> ResponseDtoPageResponseServiceResponse CreateServicesV2(ctx, accountIdentifier, optional)
Create Services

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 
 **optional** | ***ServicesApiCreateServicesV2Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ServicesApiCreateServicesV2Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of []ServiceRequest**](ServiceRequest.md)| Details of the Services to be created | 

### Return type

[**ResponseDtoPageResponseServiceResponse**](ResponseDTOPageResponseServiceResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteServiceV2**
> ResponseDtoBoolean DeleteServiceV2(ctx, serviceIdentifier, accountIdentifier, optional)
Delete a Service by identifier

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serviceIdentifier** | **string**| Service Identifier for the entity | 
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 
 **optional** | ***ServicesApiDeleteServiceV2Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ServicesApiDeleteServiceV2Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **ifMatch** | **optional.String**|  | 
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

# **GetServiceAccessList**
> ResponseDtoListServiceResponse GetServiceAccessList(ctx, accountIdentifier, optional)
Gets Service Access list

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 
 **optional** | ***ServicesApiGetServiceAccessListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ServicesApiGetServiceAccessListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| Page Index of the results to fetch.Default Value: 0 | [default to 0]
 **size** | **optional.Int32**| Results per page | [default to 100]
 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity. | 
 **projectIdentifier** | **optional.String**| Project Identifier for the Entity. | 
 **searchTerm** | **optional.String**| The word to be searched and included in the list response | 
 **serviceIdentifiers** | [**optional.Interface of []string**](string.md)| List of ServicesIds | 
 **sort** | [**optional.Interface of []string**](string.md)| Specifies the sorting criteria of the list. Like sorting based on the last updated entity, alphabetical sorting in an ascending or descending order | 
 **type_** | **optional.String**|  | 
 **gitOpsEnabled** | **optional.Bool**|  | 

### Return type

[**ResponseDtoListServiceResponse**](ResponseDTOListServiceResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetServiceList**
> ResponseDtoPageResponseServiceResponse GetServiceList(ctx, accountIdentifier, optional)
Gets Service list

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 
 **optional** | ***ServicesApiGetServiceListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ServicesApiGetServiceListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| Page Index of the results to fetch.Default Value: 0 | [default to 0]
 **size** | **optional.Int32**| Results per page | [default to 100]
 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity. | 
 **projectIdentifier** | **optional.String**| Project Identifier for the Entity. | 
 **searchTerm** | **optional.String**| The word to be searched and included in the list response | 
 **serviceIdentifiers** | [**optional.Interface of []string**](string.md)| List of ServicesIds | 
 **sort** | [**optional.Interface of []string**](string.md)| Specifies the sorting criteria of the list. Like sorting based on the last updated entity, alphabetical sorting in an ascending or descending order | 
 **type_** | **optional.String**|  | 
 **gitOpsEnabled** | **optional.Bool**|  | 
 **deploymentTemplateIdentifier** | **optional.String**|  | 
 **versionLabel** | **optional.String**|  | 

### Return type

[**ResponseDtoPageResponseServiceResponse**](ResponseDTOPageResponseServiceResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetServiceV2**
> ResponseDtoServiceResponse GetServiceV2(ctx, serviceIdentifier, accountIdentifier, optional)
Gets a Service by identifier

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serviceIdentifier** | **string**| Service Identifier for the entity | 
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 
 **optional** | ***ServicesApiGetServiceV2Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ServicesApiGetServiceV2Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity. | 
 **projectIdentifier** | **optional.String**| Project Identifier for the Entity. | 
 **deleted** | **optional.Bool**| Specify whether Service is deleted or not | [default to false]

### Return type

[**ResponseDtoServiceResponse**](ResponseDTOServiceResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateServiceV2**
> ResponseDtoServiceResponse UpdateServiceV2(ctx, accountIdentifier, optional)
Update a Service by identifier

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 
 **optional** | ***ServicesApiUpdateServiceV2Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ServicesApiUpdateServiceV2Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of ServiceRequest**](ServiceRequest.md)| Details of the Service to be updated | 
 **ifMatch** | **optional.**|  | 

### Return type

[**ResponseDtoServiceResponse**](ResponseDTOServiceResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpsertServiceV2**
> ResponseDtoServiceResponse UpsertServiceV2(ctx, accountIdentifier, optional)
Upsert a Service by identifier

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 
 **optional** | ***ServicesApiUpsertServiceV2Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ServicesApiUpsertServiceV2Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of ServiceRequest**](ServiceRequest.md)| Details of the Service to be updated | 
 **ifMatch** | **optional.**|  | 

### Return type

[**ResponseDtoServiceResponse**](ResponseDTOServiceResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ImportService**
> ResponseServiceImportResponseDto ImportService(ctx, accountIdentifier, optional)
Get Service YAML from Git Repository

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**|  | 
 **optional** | ***ServicesV2ApiImportServiceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ServicesV2ApiImportServiceOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **orgIdentifier** | **optional.String**|  | 
 **projectIdentifier** | **optional.String**|  | 
 **serviceIdentifier** | **optional.String**|  | 
 **connectorRef** | **optional.String**|  | 
 **repoName** | **optional.String**|  | 
 **branch** | **optional.String**|  | 
 **filePath** | **optional.String**|  | 
 **isForceImport** | **optional.Bool**|  | [default to false]
 **isHarnessCodeRepo** | **optional.Bool**|  | 

### Return type

[**ResponseServiceImportResponseDto**](ResponseServiceImportResponseDTO.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

