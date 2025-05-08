# ApplicationsPullRequestGeneratorAzureDevOps

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Organization** | **string** | Azure DevOps org to scan. Required. | [optional] [default to null]
**Project** | **string** | Azure DevOps project name to scan. Required. | [optional] [default to null]
**Repo** | **string** | Azure DevOps repo name to scan. Required. | [optional] [default to null]
**Api** | **string** | The Azure DevOps API URL to talk to. If blank, use https://dev.azure.com/. | [optional] [default to null]
**TokenRef** | [***ApplicationsSecretRef**](applicationsSecretRef.md) |  | [optional] [default to null]
**Labels** | **[]string** |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

