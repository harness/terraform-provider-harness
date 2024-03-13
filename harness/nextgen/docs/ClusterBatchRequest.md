# ClusterBatchRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrgIdentifier** | **string** | organization identifier of the cluster | [optional] [default to null]
**ProjectIdentifier** | **string** | project identifier of the cluster | [optional] [default to null]
**EnvRef** | **string** | environment identifier of the cluster | [default to null]
**LinkAllClusters** | **bool** | link all clusters | [optional] [default to null]
**UnlinkAllClusters** | **bool** | unlink all clusters | [optional] [default to null]
**SearchTerm** | **string** | search term if applicable. only valid if linking all clusters | [optional] [default to null]
**Clusters** | [**[]ClusterBasicDto**](ClusterBasicDTO.md) | list of cluster identifiers and names | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

