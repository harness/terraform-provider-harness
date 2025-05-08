# ApplicationsScmProviderGeneratorAzureDevOps

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Organization** | **string** | Azure Devops organization. Required. E.g. \&quot;my-organization\&quot;. | [optional] [default to null]
**Api** | **string** | The URL to Azure DevOps. If blank, use https://dev.azure.com. | [optional] [default to null]
**TeamProject** | **string** | Azure Devops team project. Required. E.g. \&quot;my-team\&quot;. | [optional] [default to null]
**AccessTokenRef** | [***ApplicationsSecretRef**](applicationsSecretRef.md) |  | [optional] [default to null]
**AllBranches** | **bool** | Scan all branches instead of just the default branch. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

