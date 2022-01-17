# {{classname}}

All URIs are relative to *https://app.harness.io/gateway*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetListOfBranchesWithStatus**](GitBranchesApi.md#GetListOfBranchesWithStatus) | **Get** /ng/api/git-sync-branch/listBranchesWithStatus | Lists branches with their status(Synced, Unsynced) by Git Sync Config Id for the given scope
[**SyncGitBranch**](GitBranchesApi.md#SyncGitBranch) | **Post** /ng/api/git-sync-branch/sync | Sync the content of new Git Branch into harness with Git Sync Config Id

# **GetListOfBranchesWithStatus**
> ResponseDtoGitBranchList GetListOfBranchesWithStatus(ctx, yamlGitConfigIdentifier, optional)
Lists branches with their status(Synced, Unsynced) by Git Sync Config Id for the given scope

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **yamlGitConfigIdentifier** | **string**| Git Sync Config Id | 
 **optional** | ***GitBranchesApiGetListOfBranchesWithStatusOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GitBranchesApiGetListOfBranchesWithStatusOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accountIdentifier** | **optional.String**| Account Identifier for the Entity | 
 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity | 
 **projectIdentifier** | **optional.String**| Project Identifier for the Entity | 
 **page** | **optional.Int32**| Indicates the number of pages. Results for these pages will be retrieved. | [default to 0]
 **size** | **optional.Int32**| The number of the elements to fetch | 
 **searchTerm** | **optional.String**| Search Term | 
 **branchSyncStatus** | **optional.String**| Used to filter out Synced and Unsynced branches | 

### Return type

[**ResponseDtoGitBranchList**](ResponseDTOGitBranchList.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml, text/yaml, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SyncGitBranch**
> ResponseDtoBoolean SyncGitBranch(ctx, repoIdentifier, optional)
Sync the content of new Git Branch into harness with Git Sync Config Id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repoIdentifier** | **string**| Git Sync Config Id | 
 **optional** | ***GitBranchesApiSyncGitBranchOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GitBranchesApiSyncGitBranchOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accountIdentifier** | **optional.String**| Account Identifier for the Entity | 
 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity | 
 **projectIdentifier** | **optional.String**| Project Identifier for the Entity | 
 **branch** | **optional.String**| Branch Name | 

### Return type

[**ResponseDtoBoolean**](ResponseDTOBoolean.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml, text/yaml, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

