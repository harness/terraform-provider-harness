# Policy

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** | Harness account ID associated with this policy | 
**Created** | **int64** | Time the policy was created | [default to null]
**GitCommitSha** | **string** | The commit sha of the commit that last effected the file | [optional] [default to null]
**GitConnectorRef** | **string** | The harness connector used for authenticating on the git provider | [optional] [default to null]
**GitDefaultBranch** | **string** | The default branch, the service pulls in changes from from this branch for policy evaluation | [optional] [default to null]
**GitDefaultBranchUpdateError** | [***GitErrorResult**](GitErrorResult.md) |  | [optional] [default to null]
**GitDefaultBranchUpdated** | **int64** | The last time the service successfully pulled in changes from the default branch | [optional] [default to null]
**GitFileId** | **string** | The file if od the bile being updated | [optional] [default to null]
**GitFileUrl** | **string** | The url of the file on the fit provider | [optional] [default to null]
**GitPath** | **string** | The path to the file in the git repo | [optional] [default to null]
**GitRepo** | **string** | The git repo the policy resides in | [optional] [default to null]
**Identifier** | **string** | identifier of the policy | [default to null]
**Name** | **string** | Name of the policy | [default to null]
**OrgId** | **string** | Harness organization ID associated with this policy | 
**ProjectId** | **string** | Harness project ID associated with this policy | 
**Rego** | **string** | Rego that defines the policy | [default to null]
**Updated** | **int64** | Time the policy was last updated | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

