# nextgen/x{{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApplicationSetServiceCreate**](ApplicationSetApi.md#ApplicationSetServiceCreate) | **Post** /api/v1/applicationset | Update updates an applicationset
[**ApplicationSetServiceDelete**](ApplicationSetApi.md#ApplicationSetServiceDelete) | **Delete** /api/v1/applicationset/{identifier} | Delete applicationset by id
[**ApplicationSetServiceGenerate**](ApplicationSetApi.md#ApplicationSetServiceGenerate) | **Post** /api/v1/applicationset/generate | Generate child applications from application set
[**ApplicationSetServiceGet**](ApplicationSetApi.md#ApplicationSetServiceGet) | **Get** /api/v1/applicationset/{identifier} | Get applicationset
[**ApplicationSetServiceGetApplicationSetGenerator**](ApplicationSetApi.md#ApplicationSetServiceGetApplicationSetGenerator) | **Get** /api/v1/applicationset/generators/{type} | Get applicationset generator
[**ApplicationSetServiceGetApplicationSetTemplate**](ApplicationSetApi.md#ApplicationSetServiceGetApplicationSetTemplate) | **Get** /api/v1/applicationset/templates/{type} | Get applicationset template
[**ApplicationSetServiceList**](ApplicationSetApi.md#ApplicationSetServiceList) | **Post** /api/v1/applicationsets | List applicationsets
[**ApplicationSetServiceListApplicationSetGenerators**](ApplicationSetApi.md#ApplicationSetServiceListApplicationSetGenerators) | **Get** /api/v1/applicationset/generators | List applicationset generators
[**ApplicationSetServiceResourceTree**](ApplicationSetApi.md#ApplicationSetServiceResourceTree) | **Get** /api/v1/applicationset/{identifier}/resource-tree | ResourceTree returns resource tree
[**ApplicationSetServiceUpdate**](ApplicationSetApi.md#ApplicationSetServiceUpdate) | **Put** /api/v1/applicationset | Update updates an applicationset

# **ApplicationSetServiceCreate**
> Servicev1ApplicationSet ApplicationSetServiceCreate(ctx, body, optional)
Update updates an applicationset

Update applicationset.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ApplicationsApplicationSetCreateRequest**](ApplicationsApplicationSetCreateRequest.md)|  | 
 **optional** | ***ApplicationSetApiApplicationSetServiceCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ApplicationSetApiApplicationSetServiceCreateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accountIdentifier** | **optional.**| Account Identifier for the Entity. | 
 **orgIdentifier** | **optional.**| Organization Identifier for the Entity. | 
 **projectIdentifier** | **optional.**| Project Identifier for the Entity. | 
 **agentIdentifier** | **optional.**| Agent identifier for entity. | 

### Return type

[**Servicev1ApplicationSet**](servicev1ApplicationSet.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApplicationSetServiceDelete**
> ApplicationsApplicationSetResponse ApplicationSetServiceDelete(ctx, identifier, optional)
Delete applicationset by id

Delete applicationset.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **identifier** | **string**| UUID for the Application Set. | 
 **optional** | ***ApplicationSetApiApplicationSetServiceDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ApplicationSetApiApplicationSetServiceDeleteOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accountIdentifier** | **optional.String**| Account Identifier for the Entity. | 
 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity. | 
 **projectIdentifier** | **optional.String**| Project Identifier for the Entity. | 
 **agentIdentifier** | **optional.String**| Agent identifier for entity. | 

### Return type

[**ApplicationsApplicationSetResponse**](applicationsApplicationSetResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApplicationSetServiceGenerate**
> ApplicationsApplicationSetGenerateResponse ApplicationSetServiceGenerate(ctx, body, optional)
Generate child applications from application set

Generate child applications from application set.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ApplicationsApplicationSetGenerateRequest**](ApplicationsApplicationSetGenerateRequest.md)|  | 
 **optional** | ***ApplicationSetApiApplicationSetServiceGenerateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ApplicationSetApiApplicationSetServiceGenerateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accountIdentifier** | **optional.**| Account Identifier for the Entity. | 
 **orgIdentifier** | **optional.**| Organization Identifier for the Entity. | 
 **projectIdentifier** | **optional.**| Project Identifier for the Entity. | 
 **agentIdentifier** | **optional.**| Agent identifier for entity. | 

### Return type

[**ApplicationsApplicationSetGenerateResponse**](applicationsApplicationSetGenerateResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApplicationSetServiceGet**
> Servicev1ApplicationSet ApplicationSetServiceGet(ctx, identifier, optional)
Get applicationset

Returns an applicationset by identifier.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **identifier** | **string**| UUID for the Application Set. | 
 **optional** | ***ApplicationSetApiApplicationSetServiceGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ApplicationSetApiApplicationSetServiceGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accountIdentifier** | **optional.String**| Account Identifier for the Entity. | 
 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity. | 
 **projectIdentifier** | **optional.String**| Project Identifier for the Entity. | 
 **agentIdentifier** | **optional.String**| Agent identifier for entity. | 
 **fetchMode** | **optional.String**| Fetch mode for the entity. | [default to NOT_SET]

### Return type

[**Servicev1ApplicationSet**](servicev1ApplicationSet.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApplicationSetServiceGetApplicationSetGenerator**
> string ApplicationSetServiceGetApplicationSetGenerator(ctx, type_, optional)
Get applicationset generator

Get applicationset generator

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **type_** | **string**|  | 
 **optional** | ***ApplicationSetApiApplicationSetServiceGetApplicationSetGeneratorOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ApplicationSetApiApplicationSetServiceGetApplicationSetGeneratorOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accountIdentifier** | **optional.String**| Account Identifier for the Entity. | 
 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity. | 
 **projectIdentifier** | **optional.String**| Project Identifier for the Entity. | 
 **agentIdentifier** | **optional.String**| Agent identifier for entity. | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApplicationSetServiceGetApplicationSetTemplate**
> string ApplicationSetServiceGetApplicationSetTemplate(ctx, type_, optional)
Get applicationset template

Get applicationset template

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **type_** | **string**|  | 
 **optional** | ***ApplicationSetApiApplicationSetServiceGetApplicationSetTemplateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ApplicationSetApiApplicationSetServiceGetApplicationSetTemplateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accountIdentifier** | **optional.String**| Account Identifier for the Entity. | 
 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity. | 
 **projectIdentifier** | **optional.String**| Project Identifier for the Entity. | 
 **agentIdentifier** | **optional.String**| Agent identifier for entity. | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApplicationSetServiceList**
> Servicev1ApplicationSetList ApplicationSetServiceList(ctx, body)
List applicationsets

List applicationsets

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**V1ApplicationSetQuery**](V1ApplicationSetQuery.md)|  | 

### Return type

[**Servicev1ApplicationSetList**](servicev1ApplicationSetList.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApplicationSetServiceListApplicationSetGenerators**
> V1ApplicationSetGeneratorList ApplicationSetServiceListApplicationSetGenerators(ctx, )
List applicationset generators

List applicationset generators

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**V1ApplicationSetGeneratorList**](v1ApplicationSetGeneratorList.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApplicationSetServiceResourceTree**
> ApplicationsApplicationSetTree ApplicationSetServiceResourceTree(ctx, identifier, optional)
ResourceTree returns resource tree

ResourceTree returns resource tree

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **identifier** | **string**| UUID for the Application Set. | 
 **optional** | ***ApplicationSetApiApplicationSetServiceResourceTreeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ApplicationSetApiApplicationSetServiceResourceTreeOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accountIdentifier** | **optional.String**| Account Identifier for the Entity. | 
 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity. | 
 **projectIdentifier** | **optional.String**| Project Identifier for the Entity. | 
 **agentIdentifier** | **optional.String**| Agent identifier for entity. | 
 **queryName** | **optional.String**|  | 
 **queryAppsetNamespace** | **optional.String**| The application set namespace. Default empty is argocd control plane namespace. | 

### Return type

[**ApplicationsApplicationSetTree**](applicationsApplicationSetTree.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApplicationSetServiceUpdate**
> Servicev1ApplicationSet ApplicationSetServiceUpdate(ctx, body, optional)
Update updates an applicationset

Update existing applicationset.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ApplicationsApplicationSetCreateRequest**](ApplicationsApplicationSetCreateRequest.md)|  | 
 **optional** | ***ApplicationSetApiApplicationSetServiceUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ApplicationSetApiApplicationSetServiceUpdateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accountIdentifier** | **optional.**| Account Identifier for the Entity. | 
 **orgIdentifier** | **optional.**| Organization Identifier for the Entity. | 
 **projectIdentifier** | **optional.**| Project Identifier for the Entity. | 
 **agentIdentifier** | **optional.**| Agent identifier for entity. | 

### Return type

[**Servicev1ApplicationSet**](servicev1ApplicationSet.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

