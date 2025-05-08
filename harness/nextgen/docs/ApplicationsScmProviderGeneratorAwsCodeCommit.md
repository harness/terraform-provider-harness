# ApplicationsScmProviderGeneratorAwsCodeCommit

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TagFilters** | [**[]ApplicationsTagFilter**](applicationsTagFilter.md) |  | [optional] [default to null]
**Role** | **string** | Role provides the AWS IAM role to assume, for cross-account repo discovery if not provided, AppSet controller will use its pod/node identity to discover. | [optional] [default to null]
**Region** | **string** | Region provides the AWS region to discover repos. if not provided, AppSet controller will infer the current region from environment. | [optional] [default to null]
**AllBranches** | **bool** | Scan all branches instead of just the default branch. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

