# {{classname}}

All URIs are relative to *https://app.harness.io/gateway*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetGitSyncErrorsCount**](GitSyncErrorsApi.md#GetGitSyncErrorsCount) | **Get** /ng/api/git-sync-errors/count | Get Errors Count for the given scope, Repo and Branch
[**ListGitSyncErrors**](GitSyncErrorsApi.md#ListGitSyncErrors) | **Get** /ng/api/git-sync-errors | Lists Git to Harness Errors by file or connectivity errors for the given scope, Repo and Branch
[**ListGitToHarnessErrorForCommit**](GitSyncErrorsApi.md#ListGitToHarnessErrorForCommit) | **Get** /ng/api/git-sync-errors/commits/{commitId} | Lists Git to Harness Errors for the given Commit Id
[**ListGitToHarnessErrorsGroupedByCommits**](GitSyncErrorsApi.md#ListGitToHarnessErrorsGroupedByCommits) | **Get** /ng/api/git-sync-errors/aggregate | Lists Git to Harness Errors grouped by Commits for the given scope, Repo and Branch

# **GetGitSyncErrorsCount**
> ResponseDtoGitSyncErrorCount GetGitSyncErrorsCount(ctx, optional)
Get Errors Count for the given scope, Repo and Branch

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GitSyncErrorsApiGetGitSyncErrorsCountOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GitSyncErrorsApiGetGitSyncErrorsCountOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountIdentifier** | **optional.String**| Account Identifier for the Entity | 
 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity | 
 **projectIdentifier** | **optional.String**| Project Identifier for the Entity | 
 **searchTerm** | **optional.String**| Search Term | 
 **branch** | **optional.String**| Branch Name | 
 **repoIdentifier** | **optional.String**| Git Sync Config Id | 
 **getDefaultFromOtherRepo** | **optional.Bool**| if true, return all the default entities | 

### Return type

[**ResponseDtoGitSyncErrorCount**](ResponseDTOGitSyncErrorCount.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml, text/yaml, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListGitSyncErrors**
> ResponseDtoPageResponseGitSyncError ListGitSyncErrors(ctx, optional)
Lists Git to Harness Errors by file or connectivity errors for the given scope, Repo and Branch

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GitSyncErrorsApiListGitSyncErrorsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GitSyncErrorsApiListGitSyncErrorsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageIndex** | **optional.Int32**| Indicates the number of pages. Results for these pages will be retrieved. | [default to 0]
 **pageSize** | **optional.Int32**| The number of the elements to fetch | [default to 50]
 **sortOrders** | [**optional.Interface of []SortOrder**](SortOrder.md)| Sort criteria for the elements. | 
 **accountIdentifier** | **optional.String**| Account Identifier for the Entity | 
 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity | 
 **projectIdentifier** | **optional.String**| Project Identifier for the Entity | 
 **searchTerm** | **optional.String**| Search Term | 
 **branch** | **optional.String**| Branch Name | 
 **repoIdentifier** | **optional.String**| Git Sync Config Id | 
 **getDefaultFromOtherRepo** | **optional.Bool**| if true, return all the default entities | 
 **gitToHarness** | **optional.Bool**| This specifies which errors to show - (Git to Harness or Connectivity), Put true to show Git to Harness Errors | [default to true]

### Return type

[**ResponseDtoPageResponseGitSyncError**](ResponseDTOPageResponseGitSyncError.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml, text/yaml, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListGitToHarnessErrorForCommit**
> ResponseDtoPageResponseGitSyncError ListGitToHarnessErrorForCommit(ctx, commitId, optional)
Lists Git to Harness Errors for the given Commit Id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **commitId** | **string**| Commit Id | 
 **optional** | ***GitSyncErrorsApiListGitToHarnessErrorForCommitOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GitSyncErrorsApiListGitToHarnessErrorForCommitOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageIndex** | **optional.Int32**| Indicates the number of pages. Results for these pages will be retrieved. | [default to 0]
 **pageSize** | **optional.Int32**| The number of the elements to fetch | [default to 50]
 **sortOrders** | [**optional.Interface of []SortOrder**](SortOrder.md)| Sort criteria for the elements. | 
 **accountIdentifier** | **optional.String**| Account Identifier for the Entity | 
 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity | 
 **projectIdentifier** | **optional.String**| Project Identifier for the Entity | 
 **branch** | **optional.String**| Branch Name | 
 **repoIdentifier** | **optional.String**| Git Sync Config Id | 
 **getDefaultFromOtherRepo** | **optional.Bool**| if true, return all the default entities | 

### Return type

[**ResponseDtoPageResponseGitSyncError**](ResponseDTOPageResponseGitSyncError.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml, text/yaml, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListGitToHarnessErrorsGroupedByCommits**
> ResponseDtoPageResponseGitSyncErrorAggregateByCommit ListGitToHarnessErrorsGroupedByCommits(ctx, optional)
Lists Git to Harness Errors grouped by Commits for the given scope, Repo and Branch

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GitSyncErrorsApiListGitToHarnessErrorsGroupedByCommitsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GitSyncErrorsApiListGitToHarnessErrorsGroupedByCommitsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageIndex** | **optional.Int32**| Indicates the number of pages. Results for these pages will be retrieved. | [default to 0]
 **pageSize** | **optional.Int32**| The number of the elements to fetch | [default to 50]
 **sortOrders** | [**optional.Interface of []SortOrder**](SortOrder.md)| Sort criteria for the elements. | 
 **accountIdentifier** | **optional.String**| Account Identifier for the Entity | 
 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity | 
 **projectIdentifier** | **optional.String**| Project Identifier for the Entity | 
 **searchTerm** | **optional.String**| Search Term | 
 **branch** | **optional.String**| Branch Name | 
 **repoIdentifier** | **optional.String**| Git Sync Config Id | 
 **getDefaultFromOtherRepo** | **optional.Bool**| if true, return all the default entities | 
 **numberOfErrorsInSummary** | **optional.Int32**| Number of errors that will be displayed in the summary | [default to 5]

### Return type

[**ResponseDtoPageResponseGitSyncErrorAggregateByCommit**](ResponseDTOPageResponseGitSyncErrorAggregateByCommit.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml, text/yaml, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

