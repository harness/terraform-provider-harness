# GitSyncErrorAggregateByCommit

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GitCommitId** | **string** | Commit Id | [optional] [default to null]
**FailedCount** | **int32** | The number of active errors in a commit | [optional] [default to null]
**RepoId** | **string** | Git Sync Config Id. | [optional] [default to null]
**BranchName** | **string** | Name of the branch. | [optional] [default to null]
**CommitMessage** | **string** | Commit Message to use for the merge commit. | [optional] [default to null]
**CreatedAt** | **int64** | This is the time at which the Git Sync error was logged | [optional] [default to null]
**ErrorsForSummaryView** | [**[]GitSyncError**](GitSyncError.md) | This has the list of Git Sync errors corresponding to a specific Commit Id | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

