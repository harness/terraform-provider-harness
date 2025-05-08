# ApplicationsPullRequestGeneratorGithub

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Owner** | **string** | GitHub org or user to scan. Required. | [optional] [default to null]
**Repo** | **string** | GitHub repo name to scan. Required. | [optional] [default to null]
**Api** | **string** | The GitHub API URL to talk to. If blank, use https://api.github.com/. | [optional] [default to null]
**TokenRef** | [***ApplicationsSecretRef**](applicationsSecretRef.md) |  | [optional] [default to null]
**AppSecretName** | **string** | AppSecretName is a reference to a GitHub App repo-creds secret with permission to access pull requests. | [optional] [default to null]
**Labels** | **[]string** |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

