# ClustersCluster

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Server** | **string** |  | [optional] [default to null]
**Name** | **string** |  | [optional] [default to null]
**Config** | [***ClustersClusterConfig**](clustersClusterConfig.md) |  | [optional] [default to null]
**ConnectionState** | [***CommonsConnectionState**](commonsConnectionState.md) |  | [optional] [default to null]
**ServerVersion** | **string** |  | [optional] [default to null]
**Namespaces** | **[]string** | Holds list of namespaces which are accessible in that cluster. Cluster level resources will be ignored if namespace list is not empty. | [optional] [default to null]
**RefreshRequestedAt** | [***V1Time**](v1Time.md) |  | [optional] [default to null]
**Info** | [***ClustersClusterInfo**](clustersClusterInfo.md) |  | [optional] [default to null]
**Shard** | **string** | Shard contains optional shard number. Calculated on the fly by the application controller if not specified. | [optional] [default to null]
**ClusterResources** | **bool** | Indicates if cluster level resources should be managed. This setting is used only if cluster is connected in a namespaced mode. | [optional] [default to null]
**Project** | **string** |  | [optional] [default to null]
**Labels** | **map[string]string** |  | [optional] [default to null]
**Annotations** | **map[string]string** |  | [optional] [default to null]
**Type_** | [***ClustersClusterType**](clustersClusterType.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

