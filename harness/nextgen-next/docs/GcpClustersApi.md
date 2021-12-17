# {{classname}}

All URIs are relative to *https://app.harness.io/gateway*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetClusterNamesForGcp**](GcpClustersApi.md#GetClusterNamesForGcp) | **Get** /ng/api/gcp/clusters | Gets gcp cluster names

# **GetClusterNamesForGcp**
> ResponseDtoGcpResponse GetClusterNamesForGcp(ctx, accountIdentifier, orgIdentifier, projectIdentifier, optional)
Gets gcp cluster names

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the entity | 
  **orgIdentifier** | **string**| Organization Identifier for the entity | 
  **projectIdentifier** | **string**| Project Identifier for the entity | 
 **optional** | ***GcpClustersApiGetClusterNamesForGcpOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GcpClustersApiGetClusterNamesForGcpOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **connectorRef** | **optional.String**|  | 

### Return type

[**ResponseDtoGcpResponse**](ResponseDTOGcpResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

