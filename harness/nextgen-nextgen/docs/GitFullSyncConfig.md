# GitFullSyncConfig

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountIdentifier** | **string** | Account Identifier for the Entity. | [optional] [default to null]
**OrgIdentifier** | **string** | Organization Identifier for the Entity. | [optional] [default to null]
**ProjectIdentifier** | **string** | Project Identifier for the Entity. | [optional] [default to null]
**BaseBranch** | **string** | Name of the branch from which the new branch will be forked out. | [optional] [default to null]
**Branch** | **string** | Name of the branch. Entities were pushed to this branch, and a pull request was made from it. | [optional] [default to null]
**PrTitle** | **string** | Title of the pull request. | [optional] [default to null]
**CreatePullRequest** | **bool** | Determines if pull request was created. | [optional] [default to null]
**RepoIdentifier** | **string** | Git Sync Config Id. | [optional] [default to null]
**IsNewBranch** | **bool** |  | [optional] [default to null]
**TargetBranch** | **string** | Name of the target branch of the pull request. | [optional] [default to null]
**RootFolder** | **string** | Path of the root folder inside which entities were pushed. | [optional] [default to null]
**NewBranch** | **bool** |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

