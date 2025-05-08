# ApplicationsScmProviderGeneratorGitlab

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Group** | **string** | Gitlab group to scan. Required.  You can use either the project id (recommended) or the full namespaced path. | [optional] [default to null]
**IncludeSubgroups** | **bool** |  | [optional] [default to null]
**Api** | **string** | The Gitlab API URL to talk to. | [optional] [default to null]
**TokenRef** | [***ApplicationsSecretRef**](applicationsSecretRef.md) |  | [optional] [default to null]
**AllBranches** | **bool** | Scan all branches instead of just the default branch. | [optional] [default to null]
**Insecure** | **bool** |  | [optional] [default to null]
**IncludeSharedProjects** | **bool** |  | [optional] [default to null]
**Topic** | **string** | Filter repos list based on Gitlab Topic. | [optional] [default to null]
**CaRef** | [***ApplicationsConfigMapKeyRef**](applicationsConfigMapKeyRef.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

