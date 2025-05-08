# ApplicationsScmProviderGeneratorGitea

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Owner** | **string** | Gitea organization or user to scan. Required. | [optional] [default to null]
**Api** | **string** | The Gitea URL to talk to. For example https://gitea.mydomain.com/. | [optional] [default to null]
**TokenRef** | [***ApplicationsSecretRef**](applicationsSecretRef.md) |  | [optional] [default to null]
**AllBranches** | **bool** | Scan all branches instead of just the default branch. | [optional] [default to null]
**Insecure** | **bool** |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

