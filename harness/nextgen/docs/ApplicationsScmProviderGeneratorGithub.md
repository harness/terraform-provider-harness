# ApplicationsScmProviderGeneratorGithub

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Organization** | **string** | GitHub org to scan. Required. | [optional] [default to null]
**Api** | **string** | The GitHub API URL to talk to. If blank, use https://api.github.com/. | [optional] [default to null]
**TokenRef** | [***ApplicationsSecretRef**](applicationsSecretRef.md) |  | [optional] [default to null]
**AppSecretName** | **string** | AppSecretName is a reference to a GitHub App repo-creds secret. | [optional] [default to null]
**AllBranches** | **bool** | Scan all branches instead of just the default branch. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

