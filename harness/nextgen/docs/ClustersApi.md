# {{classname}}

All URIs are relative to *https://app.harness.io/gateway*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AgentClusterServiceCreate**](ClustersApi.md#AgentClusterServiceCreate) | **Post** /gitops/api/v1/agents/{agentIdentifier}/clusters | Create creates a cluster
[**AgentClusterServiceDelete**](ClustersApi.md#AgentClusterServiceDelete) | **Delete** /gitops/api/v1/agents/{agentIdentifier}/clusters/{identifier} | Delete deletes a cluster
[**AgentClusterServiceGet**](ClustersApi.md#AgentClusterServiceGet) | **Get** /gitops/api/v1/agents/{agentIdentifier}/clusters/{identifier} | Get returns a cluster by identifier
[**AgentClusterServiceList**](ClustersApi.md#AgentClusterServiceList) | **Get** /gitops/api/v1/agents/{agentIdentifier}/clusters | List returns list of clusters
[**AgentClusterServiceUpdate**](ClustersApi.md#AgentClusterServiceUpdate) | **Put** /gitops/api/v1/agents/{agentIdentifier}/clusters/{identifier} | Update updates a cluster
[**AgentGPGKeyServiceList**](ClustersApi.md#AgentGPGKeyServiceList) | **Get** /gitops/api/v1/agents/{agentIdentifier}/gpgkeys | List all available repository certificates
[**ClusterServiceExists**](ClustersApi.md#ClusterServiceExists) | **Get** /gitops/api/v1/clusters/exists | Checks for whether the cluster exists
[**ClusterServiceListClusters**](ClustersApi.md#ClusterServiceListClusters) | **Post** /gitops/api/v1/clusters | List returns list of Clusters
[**DeleteCluster**](ClustersApi.md#DeleteCluster) | **Delete** /ng/api/gitops/clusters/{identifier} | Delete a Cluster by identifier
[**GetCluster**](ClustersApi.md#GetCluster) | **Get** /ng/api/gitops/clusters/{identifier} | Gets a Cluster by identifier
[**GetClusterList**](ClustersApi.md#GetClusterList) | **Get** /ng/api/gitops/clusters | Gets cluster list
[**LinkCluster**](ClustersApi.md#LinkCluster) | **Post** /ng/api/gitops/clusters | link a Cluster
[**LinkClusters**](ClustersApi.md#LinkClusters) | **Post** /ng/api/gitops/clusters/batch | Link Clusters
[**UnlinkClustersInBatch**](ClustersApi.md#UnlinkClustersInBatch) | **Post** /ng/api/gitops/clusters/batchunlink | Unlink Clusters

# **AgentClusterServiceCreate**
> Servicev1Cluster AgentClusterServiceCreate(ctx, body, agentIdentifier, optional)
Create creates a cluster

Create clusters.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ClustersClusterCreateRequest**](ClustersClusterCreateRequest.md)|  | 
  **agentIdentifier** | **string**| Agent identifier for entity. | 
 **optional** | ***ClustersApiAgentClusterServiceCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ClustersApiAgentClusterServiceCreateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **accountIdentifier** | **optional.**| Account Identifier for the Entity. | 
 **orgIdentifier** | **optional.**| Organization Identifier for the Entity. | 
 **projectIdentifier** | **optional.**| Project Identifier for the Entity. | 
 **identifier** | **optional.**|  | 

### Return type

[**Servicev1Cluster**](servicev1Cluster.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AgentClusterServiceDelete**
> ClustersClusterResponse AgentClusterServiceDelete(ctx, agentIdentifier, identifier, optional)
Delete deletes a cluster

Delete cluster.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **agentIdentifier** | **string**| Agent identifier for entity. | 
  **identifier** | **string**|  | 
 **optional** | ***ClustersApiAgentClusterServiceDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ClustersApiAgentClusterServiceDeleteOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **accountIdentifier** | **optional.String**| Account Identifier for the Entity. | 
 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity. | 
 **projectIdentifier** | **optional.String**| Project Identifier for the Entity. | 
 **queryServer** | **optional.String**|  | 
 **queryName** | **optional.String**|  | 
 **queryIdType** | **optional.String**| type is the type of the specified cluster identifier ( \&quot;server\&quot; - default, \&quot;name\&quot; ). | 
 **queryIdValue** | **optional.String**| value holds the cluster server URL or cluster name. | 
 **queryProject** | **optional.String**|  | 

### Return type

[**ClustersClusterResponse**](clustersClusterResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AgentClusterServiceGet**
> Servicev1Cluster AgentClusterServiceGet(ctx, agentIdentifier, identifier, accountIdentifier, optional)
Get returns a cluster by identifier

Get cluster.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **agentIdentifier** | **string**| Agent identifier for entity. | 
  **identifier** | **string**|  | 
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 
 **optional** | ***ClustersApiAgentClusterServiceGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ClustersApiAgentClusterServiceGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity. | 
 **projectIdentifier** | **optional.String**| Project Identifier for the Entity. | 
 **queryServer** | **optional.String**|  | 
 **queryName** | **optional.String**|  | 
 **queryIdType** | **optional.String**| type is the type of the specified cluster identifier ( \&quot;server\&quot; - default, \&quot;name\&quot; ). | 
 **queryIdValue** | **optional.String**| value holds the cluster server URL or cluster name. | 
 **queryProject** | **optional.String**|  | 

### Return type

[**Servicev1Cluster**](servicev1Cluster.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AgentClusterServiceList**
> ClustersClusterList AgentClusterServiceList(ctx, agentIdentifier, accountIdentifier, optional)
List returns list of clusters

List clusters.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **agentIdentifier** | **string**| Agent identifier for entity. | 
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 
 **optional** | ***ClustersApiAgentClusterServiceListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ClustersApiAgentClusterServiceListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity. | 
 **projectIdentifier** | **optional.String**| Project Identifier for the Entity. | 
 **identifier** | **optional.String**|  | 
 **queryServer** | **optional.String**|  | 
 **queryName** | **optional.String**|  | 
 **queryIdType** | **optional.String**| type is the type of the specified cluster identifier ( \&quot;server\&quot; - default, \&quot;name\&quot; ). | 
 **queryIdValue** | **optional.String**| value holds the cluster server URL or cluster name. | 
 **queryProject** | **optional.String**|  | 

### Return type

[**ClustersClusterList**](clustersClusterList.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AgentClusterServiceUpdate**
> Servicev1Cluster AgentClusterServiceUpdate(ctx, body, agentIdentifier, identifier, optional)
Update updates a cluster

Update cluster.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ClustersClusterUpdateRequest**](ClustersClusterUpdateRequest.md)|  | 
  **agentIdentifier** | **string**| Agent identifier for entity. | 
  **identifier** | **string**|  | 
 **optional** | ***ClustersApiAgentClusterServiceUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ClustersApiAgentClusterServiceUpdateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **accountIdentifier** | **optional.**| Account Identifier for the Entity. | 
 **orgIdentifier** | **optional.**| Organization Identifier for the Entity. | 
 **projectIdentifier** | **optional.**| Project Identifier for the Entity. | 

### Return type

[**Servicev1Cluster**](servicev1Cluster.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AgentGPGKeyServiceList**
> GpgkeysGnuPgPublicKeyList AgentGPGKeyServiceList(ctx, agentIdentifier, accountIdentifier, optional)
List all available repository certificates

List all available repository certificates.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **agentIdentifier** | **string**| Agent identifier for entity. | 
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 
 **optional** | ***ClustersApiAgentGPGKeyServiceListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ClustersApiAgentGPGKeyServiceListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity. | 
 **projectIdentifier** | **optional.String**| Project Identifier for the Entity. | 
 **queryKeyID** | **optional.String**| The GPG key ID to query for. | 

### Return type

[**GpgkeysGnuPgPublicKeyList**](gpgkeysGnuPGPublicKeyList.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ClusterServiceExists**
> bool ClusterServiceExists(ctx, accountIdentifier, optional)
Checks for whether the cluster exists

Checks for whether the cluster exists

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 
 **optional** | ***ClustersApiClusterServiceExistsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ClustersApiClusterServiceExistsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity. | 
 **projectIdentifier** | **optional.String**| Project Identifier for the Entity. | 
 **agentIdentifier** | **optional.String**| Agent identifier for entity. | 
 **server** | **optional.String**|  | 

### Return type

**bool**

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ClusterServiceListClusters**
> V1Clusterlist ClusterServiceListClusters(ctx, body)
List returns list of Clusters

List returns list of Clusters

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Servicev1ClusterQuery**](Servicev1ClusterQuery.md)|  | 

### Return type

[**V1Clusterlist**](v1Clusterlist.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

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
 **agentIdentifier** | **optional.String**| agentIdentifier | 
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
 **agentIdentifier** | **optional.String**| agentIdentifier | 
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

# **UnlinkClustersInBatch**
> ResponseDtoClusterBatchResponse UnlinkClustersInBatch(ctx, accountIdentifier, optional)
Unlink Clusters

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 
 **optional** | ***ClustersApiUnlinkClustersInBatchOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ClustersApiUnlinkClustersInBatchOpts struct
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

