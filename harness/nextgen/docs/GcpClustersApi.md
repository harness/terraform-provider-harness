# {{classname}}

All URIs are relative to *https://app.harness.io/gateway*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetClusterNamesForGcp**](GcpClustersApi.md#GetClusterNamesForGcp) | **Get** /ng/api/gcp/clusters | Gets gcp cluster names

# **GetClusterNamesForGcp**
> ResponseDtoGcpResponse GetClusterNamesForGcp(ctx, connectorRef, accountIdentifier, orgIdentifier, projectIdentifier)
Gets gcp cluster names

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **connectorRef** | **string**| GCP Connector Identifier | 
  **accountIdentifier** | **string**| Account Identifier for the Entity | 
  **orgIdentifier** | **string**| Organization Identifier for the Entity | 
  **projectIdentifier** | **string**| Project Identifier for the Entity | 

### Return type

[**ResponseDtoGcpResponse**](ResponseDTOGcpResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

