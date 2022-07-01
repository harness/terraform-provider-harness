# GitFullSyncConfigRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Branch** | **string** | Name of the branch to which the entities will be pushed and from which pull request will be created. | [default to null]
**RepoIdentifier** | **string** | Git Sync Config Id. | [default to null]
**RootFolder** | **string** | Path of the root folder inside which the entities will be pushed. | [default to null]
**IsNewBranch** | **bool** |  | [optional] [default to null]
**BaseBranch** | **string** | Name of the branch from which new branch will be forked out. | [optional] [default to null]
**CreatePullRequest** | **bool** | If true a pull request will be created from branch to target branch.Default: false. | [optional] [default to null]
**TargetBranch** | **string** | Name of the branch to which pull request will be merged. | [optional] [default to null]
**PrTitle** | **string** | Title of the pull request. | [optional] [default to null]
**NewBranch** | **bool** |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

