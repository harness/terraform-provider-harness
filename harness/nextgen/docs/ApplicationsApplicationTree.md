# ApplicationsApplicationTree

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Nodes** | [**[]ApplicationsResourceNode**](applicationsResourceNode.md) | Nodes contains list of nodes which either directly managed by the application and children of directly managed nodes. | [optional] [default to null]
**OrphanedNodes** | [**[]ApplicationsResourceNode**](applicationsResourceNode.md) | OrphanedNodes contains if or orphaned nodes: nodes which are not managed by the app but in the same namespace. List is populated only if orphaned resources enabled in app project. | [optional] [default to null]
**Hosts** | [**[]ApplicationsHostInfo**](applicationsHostInfo.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

