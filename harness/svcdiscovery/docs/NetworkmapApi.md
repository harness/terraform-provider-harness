# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNetworkMap**](NetworkmapApi.md#CreateNetworkMap) | **Post** /api/v1/agents/{agentIdentity}/networkmaps | Create networkmap
[**DeleteNetworkMap**](NetworkmapApi.md#DeleteNetworkMap) | **Delete** /api/v1/agents/{agentIdentity}/networkmaps/{networkMapIdentity} | Delete a networkmap
[**GetNetworkMap**](NetworkmapApi.md#GetNetworkMap) | **Get** /api/v1/agents/{agentIdentity}/networkmaps/{networkMapIdentity} | Get a networkmap
[**ListDiscoveredServiceForANetworkMap**](NetworkmapApi.md#ListDiscoveredServiceForANetworkMap) | **Get** /api/v1/agents/{agentIdentity}/networkmaps/{networkMapIdentity}/discoveredservices | Get list of custom services for a given netwrk map
[**ListNetworkMap**](NetworkmapApi.md#ListNetworkMap) | **Get** /api/v1/agents/{agentIdentity}/networkmaps | Get list of networkmaps
[**NetworkMapAutoGroup**](NetworkmapApi.md#NetworkMapAutoGroup) | **Post** /api/v1/agents/{agentIdentity}/autogroupnetworkmaps | auto group networkmap
[**SaveNetworkMap**](NetworkmapApi.md#SaveNetworkMap) | **Post** /api/v1/agents/{agentIdentity}/savenetworkmaps | Save networkmap
[**UpdateNetworkMap**](NetworkmapApi.md#UpdateNetworkMap) | **Put** /api/v1/agents/{agentIdentity}/networkmaps/{networkMapIdentity} | Update a networkmap

# **CreateNetworkMap**
> ApiGetNetworkMapResponse CreateNetworkMap(ctx, body, agentIdentity, accountIdentifier, environmentIdentifier, optional)
Create networkmap

Create networkmap

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ApiCreateNetworkMapRequest**](ApiCreateNetworkMapRequest.md)| Create NetworkMap | 
  **agentIdentity** | **string**| agent identity | 
  **accountIdentifier** | **string**| account id that want to access the resource | 
  **environmentIdentifier** | **string**| environment id is the environment where you want to access the resource | 
 **optional** | ***NetworkmapApiCreateNetworkMapOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NetworkmapApiCreateNetworkMapOpts struct
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

# **DeleteNetworkMap**
> ApiEmpty DeleteNetworkMap(ctx, agentIdentity, networkMapIdentity, accountIdentifier, environmentIdentifier, optional)
Delete a networkmap

Delete a networkmap

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **agentIdentity** | **string**| agent identity | 
  **networkMapIdentity** | **string**| network map identity | 
  **accountIdentifier** | **string**| account id that want to access the resource | 
  **environmentIdentifier** | **string**| environment id is the environment where you want to access the resource | 
 **optional** | ***NetworkmapApiDeleteNetworkMapOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NetworkmapApiDeleteNetworkMapOpts struct
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

# **GetNetworkMap**
> ApiGetNetworkMapResponse GetNetworkMap(ctx, agentIdentity, networkMapIdentity, accountIdentifier, environmentIdentifier, optional)
Get a networkmap

Get a networkmap

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **agentIdentity** | **string**| agent identity | 
  **networkMapIdentity** | **string**| network map identity | 
  **accountIdentifier** | **string**| account id that want to access the resource | 
  **environmentIdentifier** | **string**| environment id is the environment where you want to access the resource | 
 **optional** | ***NetworkmapApiGetNetworkMapOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NetworkmapApiGetNetworkMapOpts struct
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

# **ListDiscoveredServiceForANetworkMap**
> ApiListDiscoveredService ListDiscoveredServiceForANetworkMap(ctx, agentIdentity, networkMapIdentity, accountIdentifier, environmentIdentifier, optional)
Get list of custom services for a given netwrk map

Get list of custom services for a given netwrk map

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **agentIdentity** | **string**| agent identity | 
  **networkMapIdentity** | **string**| network map identity | 
  **accountIdentifier** | **string**| account id that want to access the resource | 
  **environmentIdentifier** | **string**| environment id is the environment where you want to access the resource | 
 **optional** | ***NetworkmapApiListDiscoveredServiceForANetworkMapOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NetworkmapApiListDiscoveredServiceForANetworkMapOpts struct
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

# **ListNetworkMap**
> ApiListNetworkMapResponse ListNetworkMap(ctx, agentIdentity, accountIdentifier, environmentIdentifier, page, limit, all, optional)
Get list of networkmaps

Get list of networkmaps

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
 **optional** | ***NetworkmapApiListNetworkMapOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NetworkmapApiListNetworkMapOpts struct
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

# **NetworkMapAutoGroup**
> ApiGetNetworkMapResponse NetworkMapAutoGroup(ctx, body, agentIdentity, accountIdentifier, environmentIdentifier, optional)
auto group networkmap

Auto group networkmap based on given rules

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ApiCreateNetworkMapRequest**](ApiCreateNetworkMapRequest.md)| Create NetworkMap | 
  **agentIdentity** | **string**| agent identity | 
  **accountIdentifier** | **string**| account id that want to access the resource | 
  **environmentIdentifier** | **string**| environment id is the environment where you want to access the resource | 
 **optional** | ***NetworkmapApiNetworkMapAutoGroupOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NetworkmapApiNetworkMapAutoGroupOpts struct
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

# **SaveNetworkMap**
> ApiGetNetworkMapResponse SaveNetworkMap(ctx, body, agentIdentity, accountIdentifier, environmentIdentifier, optional)
Save networkmap

Save networkmap

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ApiSaveNetworkMapRequest**](ApiSaveNetworkMapRequest.md)| Save NetworkMap | 
  **agentIdentity** | **string**| agent identity | 
  **accountIdentifier** | **string**| account id that want to access the resource | 
  **environmentIdentifier** | **string**| environment id is the environment where you want to access the resource | 
 **optional** | ***NetworkmapApiSaveNetworkMapOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NetworkmapApiSaveNetworkMapOpts struct
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

# **UpdateNetworkMap**
> ApiGetNetworkMapResponse UpdateNetworkMap(ctx, body, agentIdentity, networkMapIdentity, accountIdentifier, environmentIdentifier, optional)
Update a networkmap

Update a networkmap

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ApiUpdateNetworkMapRequest**](ApiUpdateNetworkMapRequest.md)| Update NetworkMap | 
  **agentIdentity** | **string**| agent identity | 
  **networkMapIdentity** | **string**| network map identity | 
  **accountIdentifier** | **string**| account id that want to access the resource | 
  **environmentIdentifier** | **string**| environment id is the environment where you want to access the resource | 
 **optional** | ***NetworkmapApiUpdateNetworkMapOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NetworkmapApiUpdateNetworkMapOpts struct
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

