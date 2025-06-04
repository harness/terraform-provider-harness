# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApplicationMapAutoGroup**](ApplicationmapApi.md#ApplicationMapAutoGroup) | **Post** /api/v1/agents/{agentIdentity}/autogroupapplicationmaps | auto group applicationmap
[**CreateApplicationMap**](ApplicationmapApi.md#CreateApplicationMap) | **Post** /api/v1/agents/{agentIdentity}/applicationmaps | Create applicationmap
[**DeleteApplicationMap**](ApplicationmapApi.md#DeleteApplicationMap) | **Delete** /api/v1/agents/{agentIdentity}/applicationmaps/{networkMapIdentity} | Delete a applicationmap
[**GetApplicationMap**](ApplicationmapApi.md#GetApplicationMap) | **Get** /api/v1/agents/{agentIdentity}/applicationmaps/{networkMapIdentity} | Get a applicationmap
[**ListApplicationMap**](ApplicationmapApi.md#ListApplicationMap) | **Get** /api/v1/agents/{agentIdentity}/applicationmaps | Get list of applicationmaps
[**ListDiscoveredServiceForApplicationMap**](ApplicationmapApi.md#ListDiscoveredServiceForApplicationMap) | **Get** /api/v1/agents/{agentIdentity}/applicationmaps/{networkMapIdentity}/discoveredservices | Get list of custom services for a given netwrk map
[**RemoveStaleServicesFromApplicationMap**](ApplicationmapApi.md#RemoveStaleServicesFromApplicationMap) | **Post** /api/v1/agents/{agentIdentity}/applicationmaps/{networkMapIdentity}/removestaleservices | Remove stale services from applicationmap
[**SaveApplicationMap**](ApplicationmapApi.md#SaveApplicationMap) | **Post** /api/v1/agents/{agentIdentity}/saveapplicationmaps | Save applicationmap
[**UpdateApplicationMap**](ApplicationmapApi.md#UpdateApplicationMap) | **Put** /api/v1/agents/{agentIdentity}/applicationmaps/{networkMapIdentity} | Update a applicationmap

# **ApplicationMapAutoGroup**
> ApiGetNetworkMapResponse ApplicationMapAutoGroup(ctx, body, agentIdentity, accountIdentifier, environmentIdentifier, optional)
auto group applicationmap

Auto group applicationmap based on given rules

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ApiCreateNetworkMapRequest**](ApiCreateNetworkMapRequest.md)| Create ApplicationMap | 
  **agentIdentity** | **string**| agent identity | 
  **accountIdentifier** | **string**| account id that want to access the resource | 
  **environmentIdentifier** | **string**| environment id is the environment where you want to access the resource | 
 **optional** | ***ApplicationmapApiApplicationMapAutoGroupOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ApplicationmapApiApplicationMapAutoGroupOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **correlationID** | **optional.**| correlation id is used to debug micro svc communication | 
 **organizationIdentifier** | **optional.**| organization id that want to access the resource | 
 **projectIdentifier** | **optional.**| project id that want to access the resource | 

### Return type

[**ApiGetNetworkMapResponse**](api.GetNetworkMapResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateApplicationMap**
> ApiGetNetworkMapResponse CreateApplicationMap(ctx, body, agentIdentity, accountIdentifier, environmentIdentifier, optional)
Create applicationmap

Create applicationmap

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ApiCreateNetworkMapRequest**](ApiCreateNetworkMapRequest.md)| Create ApplicationMap | 
  **agentIdentity** | **string**| agent identity | 
  **accountIdentifier** | **string**| account id that want to access the resource | 
  **environmentIdentifier** | **string**| environment id is the environment where you want to access the resource | 
 **optional** | ***ApplicationmapApiCreateApplicationMapOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ApplicationmapApiCreateApplicationMapOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **correlationID** | **optional.**| correlation id is used to debug micro svc communication | 
 **organizationIdentifier** | **optional.**| organization id that want to access the resource | 
 **projectIdentifier** | **optional.**| project id that want to access the resource | 

### Return type

[**ApiGetNetworkMapResponse**](api.GetNetworkMapResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteApplicationMap**
> ApiEmpty DeleteApplicationMap(ctx, agentIdentity, networkMapIdentity, accountIdentifier, environmentIdentifier, optional)
Delete a applicationmap

Delete a applicationmap

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **agentIdentity** | **string**| agent identity | 
  **networkMapIdentity** | **string**| application map identity | 
  **accountIdentifier** | **string**| account id that want to access the resource | 
  **environmentIdentifier** | **string**| environment id is the environment where you want to access the resource | 
 **optional** | ***ApplicationmapApiDeleteApplicationMapOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ApplicationmapApiDeleteApplicationMapOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **correlationID** | **optional.String**| correlation id is used to debug micro svc communication | 
 **organizationIdentifier** | **optional.String**| organization id that want to access the resource | 
 **projectIdentifier** | **optional.String**| project id that want to access the resource | 

### Return type

[**ApiEmpty**](api.Empty.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApplicationMap**
> ApiGetNetworkMapResponse GetApplicationMap(ctx, agentIdentity, networkMapIdentity, accountIdentifier, environmentIdentifier, optional)
Get a applicationmap

Get a applicationmap

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **agentIdentity** | **string**| agent identity | 
  **networkMapIdentity** | **string**| ApplicationMap map identity | 
  **accountIdentifier** | **string**| account id that want to access the resource | 
  **environmentIdentifier** | **string**| environment id is the environment where you want to access the resource | 
 **optional** | ***ApplicationmapApiGetApplicationMapOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ApplicationmapApiGetApplicationMapOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **correlationID** | **optional.String**| correlation id is used to debug micro svc communication | 
 **organizationIdentifier** | **optional.String**| organization id that want to access the resource | 
 **projectIdentifier** | **optional.String**| project id that want to access the resource | 

### Return type

[**ApiGetNetworkMapResponse**](api.GetNetworkMapResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListApplicationMap**
> ApiListNetworkMapResponse ListApplicationMap(ctx, agentIdentity, accountIdentifier, environmentIdentifier, page, limit, all, optional)
Get list of applicationmaps

Get list of applicationmaps

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **agentIdentity** | **string**| agent identity | 
  **accountIdentifier** | **string**| account id that want to access the resource | 
  **environmentIdentifier** | **string**| environment id is the environment where you want to access the resource | 
  **page** | **int32**| page number | [default to 0]
  **limit** | **int32**| limit per page | [default to 10]
  **all** | **bool**| get all | [default to false]
 **optional** | ***ApplicationmapApiListApplicationMapOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ApplicationmapApiListApplicationMapOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------






 **correlationID** | **optional.String**| correlation id is used to debug micro svc communication | 
 **organizationIdentifier** | **optional.String**| organization id that want to access the resource | 
 **projectIdentifier** | **optional.String**| project id that want to access the resource | 
 **search** | **optional.String**| search based on name | 

### Return type

[**ApiListNetworkMapResponse**](api.ListNetworkMapResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListDiscoveredServiceForApplicationMap**
> ApiListDiscoveredService ListDiscoveredServiceForApplicationMap(ctx, agentIdentity, networkMapIdentity, accountIdentifier, environmentIdentifier, optional)
Get list of custom services for a given netwrk map

Get list of custom services for a given netwrk map

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **agentIdentity** | **string**| agent identity | 
  **networkMapIdentity** | **string**| application map identity | 
  **accountIdentifier** | **string**| account id that want to access the resource | 
  **environmentIdentifier** | **string**| environment id is the environment where you want to access the resource | 
 **optional** | ***ApplicationmapApiListDiscoveredServiceForApplicationMapOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ApplicationmapApiListDiscoveredServiceForApplicationMapOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **correlationID** | **optional.String**| correlation id is used to debug micro svc communication | 
 **organizationIdentifier** | **optional.String**| organization id that want to access the resource | 
 **projectIdentifier** | **optional.String**| project id that want to access the resource | 
 **search** | **optional.String**| search based on name | 

### Return type

[**ApiListDiscoveredService**](api.ListDiscoveredService.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveStaleServicesFromApplicationMap**
> ApiGetNetworkMapResponse RemoveStaleServicesFromApplicationMap(ctx, body, agentIdentity, networkMapIdentity, accountIdentifier, environmentIdentifier, optional)
Remove stale services from applicationmap

Get a applicationmap

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ApiEmpty**](ApiEmpty.md)| Remove stale services | 
  **agentIdentity** | **string**| agent identity | 
  **networkMapIdentity** | **string**| ApplicationMap map identity | 
  **accountIdentifier** | **string**| account id that want to access the resource | 
  **environmentIdentifier** | **string**| environment id is the environment where you want to access the resource | 
 **optional** | ***ApplicationmapApiRemoveStaleServicesFromApplicationMapOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ApplicationmapApiRemoveStaleServicesFromApplicationMapOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **correlationID** | **optional.**| correlation id is used to debug micro svc communication | 
 **organizationIdentifier** | **optional.**| organization id that want to access the resource | 
 **projectIdentifier** | **optional.**| project id that want to access the resource | 

### Return type

[**ApiGetNetworkMapResponse**](api.GetNetworkMapResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SaveApplicationMap**
> ApiGetNetworkMapResponse SaveApplicationMap(ctx, body, agentIdentity, accountIdentifier, environmentIdentifier, optional)
Save applicationmap

Save applicationmap

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ApiSaveNetworkMapRequest**](ApiSaveNetworkMapRequest.md)| Save NetworkMap | 
  **agentIdentity** | **string**| agent identity | 
  **accountIdentifier** | **string**| account id that want to access the resource | 
  **environmentIdentifier** | **string**| environment id is the environment where you want to access the resource | 
 **optional** | ***ApplicationmapApiSaveApplicationMapOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ApplicationmapApiSaveApplicationMapOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **correlationID** | **optional.**| correlation id is used to debug micro svc communication | 
 **organizationIdentifier** | **optional.**| organization id that want to access the resource | 
 **projectIdentifier** | **optional.**| project id that want to access the resource | 

### Return type

[**ApiGetNetworkMapResponse**](api.GetNetworkMapResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateApplicationMap**
> ApiGetNetworkMapResponse UpdateApplicationMap(ctx, body, agentIdentity, networkMapIdentity, accountIdentifier, environmentIdentifier, optional)
Update a applicationmap

Update a applicationmap

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ApiUpdateNetworkMapRequest**](ApiUpdateNetworkMapRequest.md)| Update ApplicationMap | 
  **agentIdentity** | **string**| agent identity | 
  **networkMapIdentity** | **string**| application map identity | 
  **accountIdentifier** | **string**| account id that want to access the resource | 
  **environmentIdentifier** | **string**| environment id is the environment where you want to access the resource | 
 **optional** | ***ApplicationmapApiUpdateApplicationMapOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ApplicationmapApiUpdateApplicationMapOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **correlationID** | **optional.**| correlation id is used to debug micro svc communication | 
 **organizationIdentifier** | **optional.**| organization id that want to access the resource | 
 **projectIdentifier** | **optional.**| project id that want to access the resource | 

### Return type

[**ApiGetNetworkMapResponse**](api.GetNetworkMapResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

