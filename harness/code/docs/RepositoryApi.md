# {{classname}}

All URIs are relative to */gateway/code/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CalculateCommitDivergence**](RepositoryApi.md#CalculateCommitDivergence) | **Post** /repos/{repo_identifier}/commits/calculate-divergence | Get commit divergence
[**CodeOwnersValidate**](RepositoryApi.md#CodeOwnersValidate) | **Get** /repos/{repo_identifier}/codeowners/validate | Validate code owners file
[**CommitFiles**](RepositoryApi.md#CommitFiles) | **Post** /repos/{repo_identifier}/commits | Commit files
[**CreateBranch**](RepositoryApi.md#CreateBranch) | **Post** /repos/{repo_identifier}/branches | Create branch
[**CreateRepository**](RepositoryApi.md#CreateRepository) | **Post** /repos | Create repository
[**CreateTag**](RepositoryApi.md#CreateTag) | **Post** /repos/{repo_identifier}/tags | Create tag
[**DeleteBranch**](RepositoryApi.md#DeleteBranch) | **Delete** /repos/{repo_identifier}/branches/{branch_name} | Delete branch
[**DeleteRepository**](RepositoryApi.md#DeleteRepository) | **Delete** /repos/{repo_identifier} | Soft delete repository
[**DeleteTag**](RepositoryApi.md#DeleteTag) | **Delete** /repos/{repo_identifier}/tags/{tag_name} | Delete tag
[**DiffStats**](RepositoryApi.md#DiffStats) | **Get** /repos/{repo_identifier}/diff-stats/{range} | Get diff stats
[**GetBlame**](RepositoryApi.md#GetBlame) | **Get** /repos/{repo_identifier}/blame/{path} | Get git blame
[**GetBranch**](RepositoryApi.md#GetBranch) | **Get** /repos/{repo_identifier}/branches/{branch_name} | Get branch
[**GetCommit**](RepositoryApi.md#GetCommit) | **Get** /repos/{repo_identifier}/commits/{commit_sha} | Get commit
[**GetCommitDiff**](RepositoryApi.md#GetCommitDiff) | **Get** /repos/{repo_identifier}/commits/{commit_sha}/diff | Get raw git diff of a commit
[**GetContent**](RepositoryApi.md#GetContent) | **Get** /repos/{repo_identifier}/content/{path} | Get content of a file
[**GetRaw**](RepositoryApi.md#GetRaw) | **Get** /repos/{repo_identifier}/raw/{path} | Get raw file content
[**GetRepository**](RepositoryApi.md#GetRepository) | **Get** /repos/{repo_identifier} | Get repository
[**ImportRepository**](RepositoryApi.md#ImportRepository) | **Post** /repos/import | Import repository
[**ListBranches**](RepositoryApi.md#ListBranches) | **Get** /repos/{repo_identifier}/branches | List branches
[**ListCommits**](RepositoryApi.md#ListCommits) | **Get** /repos/{repo_identifier}/commits | List commits
[**ListRepos**](RepositoryApi.md#ListRepos) | **Get** /repos | List repositories
[**ListTags**](RepositoryApi.md#ListTags) | **Get** /repos/{repo_identifier}/tags | List tags
[**MergeCheck**](RepositoryApi.md#MergeCheck) | **Post** /repos/{repo_identifier}/merge-check/{range} | Check mergeability
[**MoveRepository**](RepositoryApi.md#MoveRepository) | **Post** /repos/{repo_identifier}/move | Move repository
[**PathDetails**](RepositoryApi.md#PathDetails) | **Post** /repos/{repo_identifier}/path-details | Get commit details
[**PurgeRepository**](RepositoryApi.md#PurgeRepository) | **Post** /repos/{repo_identifier}/purge | Purge repository
[**RawDiff**](RepositoryApi.md#RawDiff) | **Get** /repos/{repo_identifier}/diff/{range} | Get raw diff
[**RawDiffPost**](RepositoryApi.md#RawDiffPost) | **Post** /repos/{repo_identifier}/diff/{range} | Get raw diff
[**RestoreRepository**](RepositoryApi.md#RestoreRepository) | **Post** /repos/{repo_identifier}/restore | Restore repository
[**RuleAdd**](RepositoryApi.md#RuleAdd) | **Post** /repos/{repo_identifier}/rules | Add protection rule
[**RuleDelete**](RepositoryApi.md#RuleDelete) | **Delete** /repos/{repo_identifier}/rules/{rule_uid} | Delete protection rule
[**RuleGet**](RepositoryApi.md#RuleGet) | **Get** /repos/{repo_identifier}/rules/{rule_uid} | Get protection rule
[**RuleList**](RepositoryApi.md#RuleList) | **Get** /repos/{repo_identifier}/rules | List protection rules
[**RuleUpdate**](RepositoryApi.md#RuleUpdate) | **Patch** /repos/{repo_identifier}/rules/{rule_uid} | Update protection rule
[**UpdateRepository**](RepositoryApi.md#UpdateRepository) | **Patch** /repos/{repo_identifier} | Update repository

# **CalculateCommitDivergence**
> []RepoCommitDivergence CalculateCommitDivergence(ctx, accountIdentifier, repoIdentifier, optional)
Get commit divergence

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity.. | 
  **repoIdentifier** | **string**|  | 
 **optional** | ***RepositoryApiCalculateCommitDivergenceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RepositoryApiCalculateCommitDivergenceOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of OpenapiCalculateCommitDivergenceRequest**](OpenapiCalculateCommitDivergenceRequest.md)|  | 
 **orgIdentifier** | **optional.**| Organization Identifier for the Entity.. | 
 **projectIdentifier** | **optional.**| Project Identifier for the Entity.. | 

### Return type

[**[]RepoCommitDivergence**](RepoCommitDivergence.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CodeOwnersValidate**
> CodeOwnersValidate(ctx, accountIdentifier, repoIdentifier, optional)
Validate code owners file

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity.. | 
  **repoIdentifier** | **string**|  | 
 **optional** | ***RepositoryApiCodeOwnersValidateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RepositoryApiCodeOwnersValidateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity.. | 
 **projectIdentifier** | **optional.String**| Project Identifier for the Entity.. | 
 **gitRef** | **optional.String**| The git reference (branch / tag / commitID) that will be used to retrieve the data. If no value is provided the default branch of the repository is used. | [default to {Repository Default Branch}]

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth), [x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CommitFiles**
> TypesCommitFilesResponse CommitFiles(ctx, accountIdentifier, repoIdentifier, optional)
Commit files

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity.. | 
  **repoIdentifier** | **string**|  | 
 **optional** | ***RepositoryApiCommitFilesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RepositoryApiCommitFilesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of OpenapiCommitFilesRequest**](OpenapiCommitFilesRequest.md)|  | 
 **orgIdentifier** | **optional.**| Organization Identifier for the Entity.. | 
 **projectIdentifier** | **optional.**| Project Identifier for the Entity.. | 

### Return type

[**TypesCommitFilesResponse**](TypesCommitFilesResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateBranch**
> RepoBranch CreateBranch(ctx, accountIdentifier, repoIdentifier, optional)
Create branch

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity.. | 
  **repoIdentifier** | **string**|  | 
 **optional** | ***RepositoryApiCreateBranchOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RepositoryApiCreateBranchOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of OpenapiCreateBranchRequest**](OpenapiCreateBranchRequest.md)|  | 
 **orgIdentifier** | **optional.**| Organization Identifier for the Entity.. | 
 **projectIdentifier** | **optional.**| Project Identifier for the Entity.. | 

### Return type

[**RepoBranch**](RepoBranch.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateRepository**
> TypesRepository CreateRepository(ctx, accountIdentifier, optional)
Create repository

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity.. | 
 **optional** | ***RepositoryApiCreateRepositoryOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RepositoryApiCreateRepositoryOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of OpenapiCreateRepositoryRequest**](OpenapiCreateRepositoryRequest.md)|  | 
 **orgIdentifier** | **optional.**| Organization Identifier for the Entity.. | 
 **projectIdentifier** | **optional.**| Project Identifier for the Entity.. | 

### Return type

[**TypesRepository**](TypesRepository.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateTag**
> RepoCommitTag CreateTag(ctx, accountIdentifier, repoIdentifier, optional)
Create tag

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity.. | 
  **repoIdentifier** | **string**|  | 
 **optional** | ***RepositoryApiCreateTagOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RepositoryApiCreateTagOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of OpenapiCreateTagRequest**](OpenapiCreateTagRequest.md)|  | 
 **orgIdentifier** | **optional.**| Organization Identifier for the Entity.. | 
 **projectIdentifier** | **optional.**| Project Identifier for the Entity.. | 

### Return type

[**RepoCommitTag**](RepoCommitTag.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteBranch**
> DeleteBranch(ctx, accountIdentifier, repoIdentifier, branchName, optional)
Delete branch

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity.. | 
  **repoIdentifier** | **string**|  | 
  **branchName** | **string**|  | 
 **optional** | ***RepositoryApiDeleteBranchOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RepositoryApiDeleteBranchOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity.. | 
 **projectIdentifier** | **optional.String**| Project Identifier for the Entity.. | 
 **bypassRules** | **optional.Bool**| Bypass rule violations if possible. | [default to false]

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth), [x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteRepository**
> RepoSoftDeleteResponse DeleteRepository(ctx, accountIdentifier, repoIdentifier, optional)
Soft delete repository

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity.. | 
  **repoIdentifier** | **string**|  | 
 **optional** | ***RepositoryApiDeleteRepositoryOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RepositoryApiDeleteRepositoryOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity.. | 
 **projectIdentifier** | **optional.String**| Project Identifier for the Entity.. | 

### Return type

[**RepoSoftDeleteResponse**](RepoSoftDeleteResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteTag**
> DeleteTag(ctx, accountIdentifier, repoIdentifier, tagName, optional)
Delete tag

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity.. | 
  **repoIdentifier** | **string**|  | 
  **tagName** | **string**|  | 
 **optional** | ***RepositoryApiDeleteTagOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RepositoryApiDeleteTagOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity.. | 
 **projectIdentifier** | **optional.String**| Project Identifier for the Entity.. | 
 **bypassRules** | **optional.Bool**| Bypass rule violations if possible. | [default to false]

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth), [x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DiffStats**
> TypesDiffStats DiffStats(ctx, accountIdentifier, repoIdentifier, range_, optional)
Get diff stats

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity.. | 
  **repoIdentifier** | **string**|  | 
  **range_** | **string**|  | 
 **optional** | ***RepositoryApiDiffStatsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RepositoryApiDiffStatsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity.. | 
 **projectIdentifier** | **optional.String**| Project Identifier for the Entity.. | 
 **path** | [**optional.Interface of []string**](string.md)| provide path for diff operation | 

### Return type

[**TypesDiffStats**](TypesDiffStats.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBlame**
> []GitBlamePart GetBlame(ctx, accountIdentifier, repoIdentifier, path, optional)
Get git blame

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity.. | 
  **repoIdentifier** | **string**|  | 
  **path** | **string**|  | 
 **optional** | ***RepositoryApiGetBlameOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RepositoryApiGetBlameOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity.. | 
 **projectIdentifier** | **optional.String**| Project Identifier for the Entity.. | 
 **gitRef** | **optional.String**| The git reference (branch / tag / commitID) that will be used to retrieve the data. If no value is provided the default branch of the repository is used. | [default to {Repository Default Branch}]
 **lineFrom** | **optional.Int32**| Line number from which the file data is considered | [default to 0]
 **lineTo** | **optional.Int32**| Line number to which the file data is considered | [default to 0]

### Return type

[**[]GitBlamePart**](GitBlamePart.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBranch**
> RepoBranch GetBranch(ctx, accountIdentifier, repoIdentifier, branchName, optional)
Get branch

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity.. | 
  **repoIdentifier** | **string**|  | 
  **branchName** | **string**|  | 
 **optional** | ***RepositoryApiGetBranchOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RepositoryApiGetBranchOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity.. | 
 **projectIdentifier** | **optional.String**| Project Identifier for the Entity.. | 

### Return type

[**RepoBranch**](RepoBranch.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCommit**
> TypesCommit GetCommit(ctx, accountIdentifier, repoIdentifier, commitSha, optional)
Get commit

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity.. | 
  **repoIdentifier** | **string**|  | 
  **commitSha** | **string**|  | 
 **optional** | ***RepositoryApiGetCommitOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RepositoryApiGetCommitOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity.. | 
 **projectIdentifier** | **optional.String**| Project Identifier for the Entity.. | 

### Return type

[**TypesCommit**](TypesCommit.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCommitDiff**
> string GetCommitDiff(ctx, accountIdentifier, repoIdentifier, commitSha, optional)
Get raw git diff of a commit

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity.. | 
  **repoIdentifier** | **string**|  | 
  **commitSha** | **string**|  | 
 **optional** | ***RepositoryApiGetCommitDiffOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RepositoryApiGetCommitDiffOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity.. | 
 **projectIdentifier** | **optional.String**| Project Identifier for the Entity.. | 

### Return type

**string**

### Authorization

[bearerAuth](../README.md#bearerAuth), [x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetContent**
> OpenapiGetContentOutput GetContent(ctx, accountIdentifier, repoIdentifier, path, optional)
Get content of a file

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity.. | 
  **repoIdentifier** | **string**|  | 
  **path** | **string**|  | 
 **optional** | ***RepositoryApiGetContentOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RepositoryApiGetContentOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity.. | 
 **projectIdentifier** | **optional.String**| Project Identifier for the Entity.. | 

### Return type

[**OpenapiGetContentOutput**](OpenapiGetContentOutput.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRaw**
> GetRaw(ctx, accountIdentifier, repoIdentifier, path, optional)
Get raw file content

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity.. | 
  **repoIdentifier** | **string**|  | 
  **path** | **string**|  | 
 **optional** | ***RepositoryApiGetRawOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RepositoryApiGetRawOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity.. | 
 **projectIdentifier** | **optional.String**| Project Identifier for the Entity.. | 
 **gitRef** | **optional.String**| The git reference (branch / tag / commitID) that will be used to retrieve the data. If no value is provided the default branch of the repository is used. | [default to {Repository Default Branch}]

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth), [x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRepository**
> TypesRepository GetRepository(ctx, accountIdentifier, repoIdentifier, optional)
Get repository

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity.. | 
  **repoIdentifier** | **string**|  | 
 **optional** | ***RepositoryApiGetRepositoryOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RepositoryApiGetRepositoryOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity.. | 
 **projectIdentifier** | **optional.String**| Project Identifier for the Entity.. | 

### Return type

[**TypesRepository**](TypesRepository.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ImportRepository**
> TypesRepository ImportRepository(ctx, accountIdentifier, optional)
Import repository

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity.. | 
 **optional** | ***RepositoryApiImportRepositoryOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RepositoryApiImportRepositoryOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of ReposImportBody**](ReposImportBody.md)|  | 
 **orgIdentifier** | **optional.**| Organization Identifier for the Entity.. | 
 **projectIdentifier** | **optional.**| Project Identifier for the Entity.. | 

### Return type

[**TypesRepository**](TypesRepository.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListBranches**
> []RepoBranch ListBranches(ctx, accountIdentifier, repoIdentifier, optional)
List branches

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity.. | 
  **repoIdentifier** | **string**|  | 
 **optional** | ***RepositoryApiListBranchesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RepositoryApiListBranchesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity.. | 
 **projectIdentifier** | **optional.String**| Project Identifier for the Entity.. | 
 **includeCommit** | **optional.Bool**| Indicates whether optional commit information should be included in the response. | [default to false]
 **query** | **optional.String**| The substring by which the branches are filtered. | 
 **order** | **optional.String**| The order of the output. | [default to asc]
 **sort** | **optional.String**| The data by which the branches are sorted. | [default to name]
 **page** | **optional.Int32**| The page to return. | [default to 1]
 **limit** | **optional.Int32**| The maximum number of results to return. | [default to 30]

### Return type

[**[]RepoBranch**](RepoBranch.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListCommits**
> []TypesListCommitResponse ListCommits(ctx, accountIdentifier, repoIdentifier, optional)
List commits

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity.. | 
  **repoIdentifier** | **string**|  | 
 **optional** | ***RepositoryApiListCommitsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RepositoryApiListCommitsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity.. | 
 **projectIdentifier** | **optional.String**| Project Identifier for the Entity.. | 
 **gitRef** | **optional.String**| The git reference (branch / tag / commitID) that will be used to retrieve the data. If no value is provided the default branch of the repository is used. | [default to {Repository Default Branch}]
 **after** | **optional.String**| The result should only contain commits that occurred after the provided reference. | 
 **path** | **optional.String**| Path for which commit information should be retrieved | 
 **since** | **optional.Int32**| Epoch since when commit information should be retrieved. | 
 **until** | **optional.Int32**| Epoch until when commit information should be retrieved. | 
 **committer** | **optional.String**| Committer pattern for which commit information should be retrieved. | 
 **page** | **optional.Int32**| The page to return. | [default to 1]
 **limit** | **optional.Int32**| The maximum number of results to return. | [default to 30]

### Return type

[**[]TypesListCommitResponse**](TypesListCommitResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListRepos**
> []TypesRepository ListRepos(ctx, accountIdentifier, optional)
List repositories

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity.. | 
 **optional** | ***RepositoryApiListReposOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RepositoryApiListReposOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity.. | 
 **projectIdentifier** | **optional.String**| Project Identifier for the Entity.. | 
 **query** | **optional.String**| The substring which is used to filter the repositories by their path name. | 
 **sort** | **optional.String**| The data by which the repositories are sorted. | [default to identifier]
 **order** | **optional.String**| The order of the output. | [default to asc]
 **page** | **optional.Int32**| The page to return. | [default to 1]
 **limit** | **optional.Int32**| The maximum number of results to return. | [default to 30]

### Return type

[**[]TypesRepository**](TypesRepository.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListTags**
> []RepoCommitTag ListTags(ctx, accountIdentifier, repoIdentifier, optional)
List tags

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity.. | 
  **repoIdentifier** | **string**|  | 
 **optional** | ***RepositoryApiListTagsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RepositoryApiListTagsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity.. | 
 **projectIdentifier** | **optional.String**| Project Identifier for the Entity.. | 
 **includeCommit** | **optional.Bool**| Indicates whether optional commit information should be included in the response. | [default to false]
 **query** | **optional.String**| The substring by which the tags are filtered. | 
 **order** | **optional.String**| The order of the output. | [default to asc]
 **sort** | **optional.String**| The data by which the tags are sorted. | [default to name]
 **page** | **optional.Int32**| The page to return. | [default to 1]
 **limit** | **optional.Int32**| The maximum number of results to return. | [default to 30]

### Return type

[**[]RepoCommitTag**](RepoCommitTag.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MergeCheck**
> RepoMergeCheck MergeCheck(ctx, accountIdentifier, repoIdentifier, range_, optional)
Check mergeability

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity.. | 
  **repoIdentifier** | **string**|  | 
  **range_** | **string**|  | 
 **optional** | ***RepositoryApiMergeCheckOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RepositoryApiMergeCheckOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity.. | 
 **projectIdentifier** | **optional.String**| Project Identifier for the Entity.. | 
 **path** | [**optional.Interface of []string**](string.md)| provide path for diff operation | 

### Return type

[**RepoMergeCheck**](RepoMergeCheck.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MoveRepository**
> TypesRepository MoveRepository(ctx, accountIdentifier, repoIdentifier, optional)
Move repository

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity.. | 
  **repoIdentifier** | **string**|  | 
 **optional** | ***RepositoryApiMoveRepositoryOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RepositoryApiMoveRepositoryOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of OpenapiMoveRepoRequest**](OpenapiMoveRepoRequest.md)|  | 
 **orgIdentifier** | **optional.**| Organization Identifier for the Entity.. | 
 **projectIdentifier** | **optional.**| Project Identifier for the Entity.. | 

### Return type

[**TypesRepository**](TypesRepository.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PathDetails**
> RepoPathsDetailsOutput PathDetails(ctx, accountIdentifier, repoIdentifier, optional)
Get commit details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity.. | 
  **repoIdentifier** | **string**|  | 
 **optional** | ***RepositoryApiPathDetailsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RepositoryApiPathDetailsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of OpenapiPathsDetailsRequest**](OpenapiPathsDetailsRequest.md)|  | 
 **orgIdentifier** | **optional.**| Organization Identifier for the Entity.. | 
 **projectIdentifier** | **optional.**| Project Identifier for the Entity.. | 
 **gitRef** | **optional.**| The git reference (branch / tag / commitID) that will be used to retrieve the data. If no value is provided the default branch of the repository is used. | [default to {Repository Default Branch}]

### Return type

[**RepoPathsDetailsOutput**](RepoPathsDetailsOutput.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PurgeRepository**
> PurgeRepository(ctx, accountIdentifier, deletedAt, repoIdentifier, optional)
Purge repository

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity.. | 
  **deletedAt** | **int32**| The exact time the resource was delete at in epoch format. | 
  **repoIdentifier** | **string**|  | 
 **optional** | ***RepositoryApiPurgeRepositoryOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RepositoryApiPurgeRepositoryOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity.. | 
 **projectIdentifier** | **optional.String**| Project Identifier for the Entity.. | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth), [x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RawDiff**
> []GitFileDiff RawDiff(ctx, accountIdentifier, repoIdentifier, range_, optional)
Get raw diff

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity.. | 
  **repoIdentifier** | **string**|  | 
  **range_** | **string**|  | 
 **optional** | ***RepositoryApiRawDiffOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RepositoryApiRawDiffOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity.. | 
 **projectIdentifier** | **optional.String**| Project Identifier for the Entity.. | 
 **path** | [**optional.Interface of []string**](string.md)| provide path for diff operation | 

### Return type

[**[]GitFileDiff**](GitFileDiff.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RawDiffPost**
> []GitFileDiff RawDiffPost(ctx, accountIdentifier, repoIdentifier, range_, optional)
Get raw diff

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity.. | 
  **repoIdentifier** | **string**|  | 
  **range_** | **string**|  | 
 **optional** | ***RepositoryApiRawDiffPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RepositoryApiRawDiffPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | [**optional.Interface of []TypesFileDiffRequest**](TypesFileDiffRequest.md)|  | 
 **orgIdentifier** | **optional.**| Organization Identifier for the Entity.. | 
 **projectIdentifier** | **optional.**| Project Identifier for the Entity.. | 

### Return type

[**[]GitFileDiff**](GitFileDiff.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RestoreRepository**
> TypesRepository RestoreRepository(ctx, accountIdentifier, deletedAt, repoIdentifier, optional)
Restore repository

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity.. | 
  **deletedAt** | **int32**| The exact time the resource was delete at in epoch format. | 
  **repoIdentifier** | **string**|  | 
 **optional** | ***RepositoryApiRestoreRepositoryOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RepositoryApiRestoreRepositoryOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | [**optional.Interface of OpenapiRestoreRequest**](OpenapiRestoreRequest.md)|  | 
 **orgIdentifier** | **optional.**| Organization Identifier for the Entity.. | 
 **projectIdentifier** | **optional.**| Project Identifier for the Entity.. | 

### Return type

[**TypesRepository**](TypesRepository.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RuleAdd**
> OpenapiRule RuleAdd(ctx, accountIdentifier, repoIdentifier, optional)
Add protection rule

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity.. | 
  **repoIdentifier** | **string**|  | 
 **optional** | ***RepositoryApiRuleAddOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RepositoryApiRuleAddOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of RepoIdentifierRulesBody**](RepoIdentifierRulesBody.md)|  | 
 **orgIdentifier** | **optional.**| Organization Identifier for the Entity.. | 
 **projectIdentifier** | **optional.**| Project Identifier for the Entity.. | 

### Return type

[**OpenapiRule**](OpenapiRule.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RuleDelete**
> RuleDelete(ctx, accountIdentifier, repoIdentifier, ruleUid, optional)
Delete protection rule

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity.. | 
  **repoIdentifier** | **string**|  | 
  **ruleUid** | **string**|  | 
 **optional** | ***RepositoryApiRuleDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RepositoryApiRuleDeleteOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity.. | 
 **projectIdentifier** | **optional.String**| Project Identifier for the Entity.. | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth), [x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RuleGet**
> OpenapiRule RuleGet(ctx, accountIdentifier, repoIdentifier, ruleUid, optional)
Get protection rule

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity.. | 
  **repoIdentifier** | **string**|  | 
  **ruleUid** | **string**|  | 
 **optional** | ***RepositoryApiRuleGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RepositoryApiRuleGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity.. | 
 **projectIdentifier** | **optional.String**| Project Identifier for the Entity.. | 

### Return type

[**OpenapiRule**](OpenapiRule.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RuleList**
> []OpenapiRule RuleList(ctx, accountIdentifier, repoIdentifier, optional)
List protection rules

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity.. | 
  **repoIdentifier** | **string**|  | 
 **optional** | ***RepositoryApiRuleListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RepositoryApiRuleListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity.. | 
 **projectIdentifier** | **optional.String**| Project Identifier for the Entity.. | 
 **query** | **optional.String**| The substring by which the repository protection rules are filtered. | 
 **order** | **optional.String**| The order of the output. | [default to asc]
 **sort** | **optional.String**| The field by which the protection rules are sorted. | [default to created_at]
 **page** | **optional.Int32**| The page to return. | [default to 1]
 **limit** | **optional.Int32**| The maximum number of results to return. | [default to 30]

### Return type

[**[]OpenapiRule**](OpenapiRule.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RuleUpdate**
> OpenapiRule RuleUpdate(ctx, accountIdentifier, repoIdentifier, ruleUid, optional)
Update protection rule

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity.. | 
  **repoIdentifier** | **string**|  | 
  **ruleUid** | **string**|  | 
 **optional** | ***RepositoryApiRuleUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RepositoryApiRuleUpdateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | [**optional.Interface of RulesRuleUidBody**](RulesRuleUidBody.md)|  | 
 **orgIdentifier** | **optional.**| Organization Identifier for the Entity.. | 
 **projectIdentifier** | **optional.**| Project Identifier for the Entity.. | 

### Return type

[**OpenapiRule**](OpenapiRule.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateRepository**
> TypesRepository UpdateRepository(ctx, accountIdentifier, repoIdentifier, optional)
Update repository

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity.. | 
  **repoIdentifier** | **string**|  | 
 **optional** | ***RepositoryApiUpdateRepositoryOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RepositoryApiUpdateRepositoryOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of OpenapiUpdateRepoRequest**](OpenapiUpdateRepoRequest.md)|  | 
 **orgIdentifier** | **optional.**| Organization Identifier for the Entity.. | 
 **projectIdentifier** | **optional.**| Project Identifier for the Entity.. | 

### Return type

[**TypesRepository**](TypesRepository.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

