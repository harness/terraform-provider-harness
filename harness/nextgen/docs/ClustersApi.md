# nextgen{{classname}}

All URIs are relative to *https://app.harness.io/gateway*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteCluster**](ClustersApi.md#DeleteCluster) | **Delete** /gitops/clusters/{identifier} | Delete a Cluster by identifier
[**GetCluster**](ClustersApi.md#GetCluster) | **Get** /gitops/clusters/{identifier} | Gets a Cluster by identifier
[**GetClusterList**](ClustersApi.md#GetClusterList) | **Get** /gitops/clusters | Gets cluster list
[**LinkCluster**](ClustersApi.md#LinkCluster) | **Post** /gitops/clusters | link a Cluster
[**LinkClusters**](ClustersApi.md#LinkClusters) | **Post** /gitops/clusters/batch | Link Clusters

# **DeleteCluster**
> ResponseDtoBoolean DeleteCluster(ctx, identifier, accountIdentifier, environmentIdentifier, optional)
Delete a Cluster by identifier

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **identifier** | **string**| Cluster Identifier for the entity | 
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 
  **environmentIdentifier** | **string**| environmentIdentifier | 
 **optional** | ***ClustersApiDeleteClusterOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ClustersApiDeleteClusterOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity. | 
 **projectIdentifier** | **optional.String**| Project Identifier for the Entity. | 
 **scope** | **optional.String**| Scope for the gitops cluster | 

### Return type

[**ResponseDtoBoolean**](ResponseDTOBoolean.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCluster**
> ResponseDtoClusterResponse GetCluster(ctx, identifier, accountIdentifier, environmentIdentifier, optional)
Gets a Cluster by identifier

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **identifier** | **string**| Cluster Identifier for the entity | 
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 
  **environmentIdentifier** | **string**| environmentIdentifier | 
 **optional** | ***ClustersApiGetClusterOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ClustersApiGetClusterOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity. | 
 **projectIdentifier** | **optional.String**| Project Identifier for the Entity. | 
 **deleted** | **optional.Bool**| Specify whether cluster is deleted or not | [default to false]

### Return type

[**ResponseDtoClusterResponse**](ResponseDTOClusterResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetClusterList**
> ResponseDtoPageResponseClusterResponse GetClusterList(ctx, accountIdentifier, environmentIdentifier, optional)
Gets cluster list

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 
  **environmentIdentifier** | **string**| Environment Identifier of the clusters | 
 **optional** | ***ClustersApiGetClusterListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ClustersApiGetClusterListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **optional.Int32**| Page Index of the results to fetch.Default Value: 0 | [default to 0]
 **size** | **optional.Int32**| Results per page | [default to 100]
 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity. | 
 **projectIdentifier** | **optional.String**| Project Identifier for the Entity. | 
 **searchTerm** | **optional.String**| The word to be searched and included in the list response | 
 **identifiers** | [**optional.Interface of []string**](string.md)| List of cluster identifiers | 
 **sort** | [**optional.Interface of []string**](string.md)| Specifies the sorting criteria of the list. Like sorting based on the last updated entity, alphabetical sorting in an ascending or descending order | 

### Return type

[**ResponseDtoPageResponseClusterResponse**](ResponseDTOPageResponseClusterResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LinkCluster**
> ResponseDtoClusterResponse LinkCluster(ctx, accountIdentifier, optional)
link a Cluster

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 
 **optional** | ***ClustersApiLinkClusterOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ClustersApiLinkClusterOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of ClusterRequest**](ClusterRequest.md)| Details of the createCluster to be linked | 

### Return type

[**ResponseDtoClusterResponse**](ResponseDTOClusterResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LinkClusters**
> ResponseDtoClusterBatchResponse LinkClusters(ctx, accountIdentifier, optional)
Link Clusters

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 
 **optional** | ***ClustersApiLinkClustersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ClustersApiLinkClustersOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of ClusterBatchRequest**](ClusterBatchRequest.md)| Details of the createCluster to be created | 

### Return type

[**ResponseDtoClusterBatchResponse**](ResponseDTOClusterBatchResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

