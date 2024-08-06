# Servicev1ClusterQuery

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountIdentifier** | **string** | Account Identifier for the Entity. | [optional] [default to null]
**ProjectIdentifier** | **string** | Project Identifier for the Entity. | [optional] [default to null]
**OrgIdentifier** | **string** | Organization Identifier for the Entity. | [optional] [default to null]
**AgentIdentifier** | **string** | Agent identifier for entity. | [optional] [default to null]
**Identifier** | **string** |  | [optional] [default to null]
**SearchTerm** | **string** |  | [optional] [default to null]
**PageSize** | **int32** |  | [optional] [default to null]
**PageIndex** | **int32** |  | [optional] [default to null]
**Filter** | [***interface{}**](interface{}.md) | Filters for Clusters. Eg. \&quot;identifier\&quot;: { \&quot;$in\&quot;: [\&quot;id1\&quot;, \&quot;id2\&quot;] | [optional] [default to null]
**SortBy** | [***ClusterQueryClusterSortByOptions**](ClusterQueryClusterSortByOptions.md) |  | [optional] [default to null]
**SortOrder** | [***V1SortOrderOptions**](v1SortOrderOptions.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

