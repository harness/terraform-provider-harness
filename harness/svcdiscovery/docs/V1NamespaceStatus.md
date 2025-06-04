# V1NamespaceStatus

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Conditions** | [**[]V1NamespaceCondition**](v1.NamespaceCondition.md) | Represents the latest available observations of a namespace&#x27;s current state. +optional +patchMergeKey&#x3D;type +patchStrategy&#x3D;merge | [optional] [default to null]
**Phase** | **string** | Phase is the current lifecycle phase of the namespace. More info: https://kubernetes.io/docs/tasks/administer-cluster/namespaces/ +optional | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

