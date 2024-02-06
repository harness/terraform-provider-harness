# code{{classname}}

All URIs are relative to */api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CodeownersPullReq**](PullreqApi.md#CodeownersPullReq) | **Get** /repos/{repo_identifier}/pullreq/{pullreq_number}/codeowners | 
[**CommentCreatePullReq**](PullreqApi.md#CommentCreatePullReq) | **Post** /repos/{repo_identifier}/pullreq/{pullreq_number}/comments | 
[**CommentDeletePullReq**](PullreqApi.md#CommentDeletePullReq) | **Delete** /repos/{repo_identifier}/pullreq/{pullreq_number}/comments/{pullreq_comment_id} | 
[**CommentStatusPullReq**](PullreqApi.md#CommentStatusPullReq) | **Put** /repos/{repo_identifier}/pullreq/{pullreq_number}/comments/{pullreq_comment_id}/status | 
[**CommentUpdatePullReq**](PullreqApi.md#CommentUpdatePullReq) | **Patch** /repos/{repo_identifier}/pullreq/{pullreq_number}/comments/{pullreq_comment_id} | 
[**CreatePullReq**](PullreqApi.md#CreatePullReq) | **Post** /repos/{repo_identifier}/pullreq | 
[**DiffPullReq**](PullreqApi.md#DiffPullReq) | **Get** /repos/{repo_identifier}/pullreq/{pullreq_number}/diff | 
[**DiffPullReqPost**](PullreqApi.md#DiffPullReqPost) | **Post** /repos/{repo_identifier}/pullreq/{pullreq_number}/diff | 
[**FileViewAddPullReq**](PullreqApi.md#FileViewAddPullReq) | **Put** /repos/{repo_identifier}/pullreq/{pullreq_number}/file-views | 
[**FileViewDeletePullReq**](PullreqApi.md#FileViewDeletePullReq) | **Delete** /repos/{repo_identifier}/pullreq/{pullreq_number}/file-views/{file_path} | 
[**FileViewListPullReq**](PullreqApi.md#FileViewListPullReq) | **Get** /repos/{repo_identifier}/pullreq/{pullreq_number}/file-views | 
[**GetPullReq**](PullreqApi.md#GetPullReq) | **Get** /repos/{repo_identifier}/pullreq/{pullreq_number} | 
[**ListPullReq**](PullreqApi.md#ListPullReq) | **Get** /repos/{repo_identifier}/pullreq | 
[**ListPullReqActivities**](PullreqApi.md#ListPullReqActivities) | **Get** /repos/{repo_identifier}/pullreq/{pullreq_number}/activities | 
[**ListPullReqCommits**](PullreqApi.md#ListPullReqCommits) | **Get** /repos/{repo_identifier}/pullreq/{pullreq_number}/commits | 
[**MergePullReqOp**](PullreqApi.md#MergePullReqOp) | **Post** /repos/{repo_identifier}/pullreq/{pullreq_number}/merge | 
[**PullReqMetaData**](PullreqApi.md#PullReqMetaData) | **Get** /repos/{repo_identifier}/pullreq/{pullreq_number}/metadata | 
[**ReviewSubmitPullReq**](PullreqApi.md#ReviewSubmitPullReq) | **Post** /repos/{repo_identifier}/pullreq/{pullreq_number}/reviews | 
[**ReviewerAddPullReq**](PullreqApi.md#ReviewerAddPullReq) | **Put** /repos/{repo_identifier}/pullreq/{pullreq_number}/reviewers | 
[**ReviewerDeletePullReq**](PullreqApi.md#ReviewerDeletePullReq) | **Delete** /repos/{repo_identifier}/pullreq/{pullreq_number}/reviewers/{pullreq_reviewer_id} | 
[**ReviewerListPullReq**](PullreqApi.md#ReviewerListPullReq) | **Get** /repos/{repo_identifier}/pullreq/{pullreq_number}/reviewers | 
[**StatePullReq**](PullreqApi.md#StatePullReq) | **Post** /repos/{repo_identifier}/pullreq/{pullreq_number}/state | 
[**UpdatePullReq**](PullreqApi.md#UpdatePullReq) | **Patch** /repos/{repo_identifier}/pullreq/{pullreq_number} | 

# **CodeownersPullReq**
> TypesCodeOwnerEvaluation CodeownersPullReq(ctx, accountIdentifier, repoIdentifier, pullreqNumber, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity.. | 
  **repoIdentifier** | **string**|  | 
  **pullreqNumber** | **int32**|  | 
 **optional** | ***PullreqApiCodeownersPullReqOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PullreqApiCodeownersPullReqOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity.. | 
 **projectIdentifier** | **optional.String**| Project Identifier for the Entity.. | 

### Return type

[**TypesCodeOwnerEvaluation**](TypesCodeOwnerEvaluation.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CommentCreatePullReq**
> TypesPullReqActivity CommentCreatePullReq(ctx, accountIdentifier, repoIdentifier, pullreqNumber, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity.. | 
  **repoIdentifier** | **string**|  | 
  **pullreqNumber** | **int32**|  | 
 **optional** | ***PullreqApiCommentCreatePullReqOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PullreqApiCommentCreatePullReqOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | [**optional.Interface of OpenapiCommentCreatePullReqRequest**](OpenapiCommentCreatePullReqRequest.md)|  | 
 **orgIdentifier** | **optional.**| Organization Identifier for the Entity.. | 
 **projectIdentifier** | **optional.**| Project Identifier for the Entity.. | 

### Return type

[**TypesPullReqActivity**](TypesPullReqActivity.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CommentDeletePullReq**
> CommentDeletePullReq(ctx, accountIdentifier, repoIdentifier, pullreqNumber, pullreqCommentId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity.. | 
  **repoIdentifier** | **string**|  | 
  **pullreqNumber** | **int32**|  | 
  **pullreqCommentId** | **int32**|  | 
 **optional** | ***PullreqApiCommentDeletePullReqOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PullreqApiCommentDeletePullReqOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity.. | 
 **projectIdentifier** | **optional.String**| Project Identifier for the Entity.. | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CommentStatusPullReq**
> TypesPullReqActivity CommentStatusPullReq(ctx, accountIdentifier, repoIdentifier, pullreqNumber, pullreqCommentId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity.. | 
  **repoIdentifier** | **string**|  | 
  **pullreqNumber** | **int32**|  | 
  **pullreqCommentId** | **int32**|  | 
 **optional** | ***PullreqApiCommentStatusPullReqOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PullreqApiCommentStatusPullReqOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **body** | [**optional.Interface of OpenapiCommentStatusPullReqRequest**](OpenapiCommentStatusPullReqRequest.md)|  | 
 **orgIdentifier** | **optional.**| Organization Identifier for the Entity.. | 
 **projectIdentifier** | **optional.**| Project Identifier for the Entity.. | 

### Return type

[**TypesPullReqActivity**](TypesPullReqActivity.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CommentUpdatePullReq**
> TypesPullReqActivity CommentUpdatePullReq(ctx, accountIdentifier, repoIdentifier, pullreqNumber, pullreqCommentId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity.. | 
  **repoIdentifier** | **string**|  | 
  **pullreqNumber** | **int32**|  | 
  **pullreqCommentId** | **int32**|  | 
 **optional** | ***PullreqApiCommentUpdatePullReqOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PullreqApiCommentUpdatePullReqOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **body** | [**optional.Interface of OpenapiCommentUpdatePullReqRequest**](OpenapiCommentUpdatePullReqRequest.md)|  | 
 **orgIdentifier** | **optional.**| Organization Identifier for the Entity.. | 
 **projectIdentifier** | **optional.**| Project Identifier for the Entity.. | 

### Return type

[**TypesPullReqActivity**](TypesPullReqActivity.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreatePullReq**
> TypesPullReq CreatePullReq(ctx, accountIdentifier, repoIdentifier, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity.. | 
  **repoIdentifier** | **string**|  | 
 **optional** | ***PullreqApiCreatePullReqOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PullreqApiCreatePullReqOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of OpenapiCreatePullReqRequest**](OpenapiCreatePullReqRequest.md)|  | 
 **orgIdentifier** | **optional.**| Organization Identifier for the Entity.. | 
 **projectIdentifier** | **optional.**| Project Identifier for the Entity.. | 

### Return type

[**TypesPullReq**](TypesPullReq.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DiffPullReq**
> []GitFileDiff DiffPullReq(ctx, accountIdentifier, repoIdentifier, pullreqNumber, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity.. | 
  **repoIdentifier** | **string**|  | 
  **pullreqNumber** | **int32**|  | 
 **optional** | ***PullreqApiDiffPullReqOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PullreqApiDiffPullReqOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity.. | 
 **projectIdentifier** | **optional.String**| Project Identifier for the Entity.. | 
 **path** | [**optional.Interface of []string**](string.md)| provide path for diff operation | 

### Return type

[**[]GitFileDiff**](GitFileDiff.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DiffPullReqPost**
> []GitFileDiff DiffPullReqPost(ctx, accountIdentifier, repoIdentifier, pullreqNumber, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity.. | 
  **repoIdentifier** | **string**|  | 
  **pullreqNumber** | **int32**|  | 
 **optional** | ***PullreqApiDiffPullReqPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PullreqApiDiffPullReqPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | [**optional.Interface of []TypesFileDiffRequest**](TypesFileDiffRequest.md)|  | 
 **orgIdentifier** | **optional.**| Organization Identifier for the Entity.. | 
 **projectIdentifier** | **optional.**| Project Identifier for the Entity.. | 

### Return type

[**[]GitFileDiff**](GitFileDiff.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FileViewAddPullReq**
> TypesPullReqFileView FileViewAddPullReq(ctx, accountIdentifier, repoIdentifier, pullreqNumber, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity.. | 
  **repoIdentifier** | **string**|  | 
  **pullreqNumber** | **int32**|  | 
 **optional** | ***PullreqApiFileViewAddPullReqOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PullreqApiFileViewAddPullReqOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | [**optional.Interface of OpenapiFileViewAddPullReqRequest**](OpenapiFileViewAddPullReqRequest.md)|  | 
 **orgIdentifier** | **optional.**| Organization Identifier for the Entity.. | 
 **projectIdentifier** | **optional.**| Project Identifier for the Entity.. | 

### Return type

[**TypesPullReqFileView**](TypesPullReqFileView.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FileViewDeletePullReq**
> FileViewDeletePullReq(ctx, accountIdentifier, repoIdentifier, pullreqNumber, filePath, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity.. | 
  **repoIdentifier** | **string**|  | 
  **pullreqNumber** | **int32**|  | 
  **filePath** | **string**|  | 
 **optional** | ***PullreqApiFileViewDeletePullReqOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PullreqApiFileViewDeletePullReqOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity.. | 
 **projectIdentifier** | **optional.String**| Project Identifier for the Entity.. | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FileViewListPullReq**
> []TypesPullReqFileView FileViewListPullReq(ctx, accountIdentifier, repoIdentifier, pullreqNumber, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity.. | 
  **repoIdentifier** | **string**|  | 
  **pullreqNumber** | **int32**|  | 
 **optional** | ***PullreqApiFileViewListPullReqOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PullreqApiFileViewListPullReqOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity.. | 
 **projectIdentifier** | **optional.String**| Project Identifier for the Entity.. | 

### Return type

[**[]TypesPullReqFileView**](TypesPullReqFileView.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPullReq**
> TypesPullReq GetPullReq(ctx, accountIdentifier, repoIdentifier, pullreqNumber, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity.. | 
  **repoIdentifier** | **string**|  | 
  **pullreqNumber** | **int32**|  | 
 **optional** | ***PullreqApiGetPullReqOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PullreqApiGetPullReqOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity.. | 
 **projectIdentifier** | **optional.String**| Project Identifier for the Entity.. | 

### Return type

[**TypesPullReq**](TypesPullReq.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListPullReq**
> []TypesPullReq ListPullReq(ctx, accountIdentifier, repoIdentifier, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity.. | 
  **repoIdentifier** | **string**|  | 
 **optional** | ***PullreqApiListPullReqOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PullreqApiListPullReqOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity.. | 
 **projectIdentifier** | **optional.String**| Project Identifier for the Entity.. | 
 **state** | [**optional.Interface of []string**](string.md)| The state of the pull requests to include in the result. | 
 **sourceRepoIdentifier** | **optional.String**| Source repository ref of the pull requests. | 
 **sourceBranch** | **optional.String**| Source branch of the pull requests. | 
 **targetBranch** | **optional.String**| Target branch of the pull requests. | 
 **query** | **optional.String**| The substring by which the pull requests are filtered. | 
 **createdBy** | **optional.Int32**| The principal ID who created pull requests. | 
 **order** | **optional.String**| The order of the output. | [default to asc]
 **sort** | **optional.String**| The data by which the pull requests are sorted. | [default to number]
 **page** | **optional.Int32**| The page to return. | [default to 1]
 **limit** | **optional.Int32**| The maximum number of results to return. | [default to 30]

### Return type

[**[]TypesPullReq**](TypesPullReq.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListPullReqActivities**
> []TypesPullReqActivity ListPullReqActivities(ctx, accountIdentifier, repoIdentifier, pullreqNumber, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity.. | 
  **repoIdentifier** | **string**|  | 
  **pullreqNumber** | **int32**|  | 
 **optional** | ***PullreqApiListPullReqActivitiesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PullreqApiListPullReqActivitiesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity.. | 
 **projectIdentifier** | **optional.String**| Project Identifier for the Entity.. | 
 **kind** | [**optional.Interface of []string**](string.md)| The kind of the pull request activity to include in the result. | 
 **type_** | [**optional.Interface of []string**](string.md)| The type of the pull request activity to include in the result. | 
 **after** | **optional.Int32**| The result should contain only entries created at and after this timestamp (unix millis). | 
 **before** | **optional.Int32**| The result should contain only entries created before this timestamp (unix millis). | 
 **limit** | **optional.Int32**| The maximum number of results to return. | [default to 30]

### Return type

[**[]TypesPullReqActivity**](TypesPullReqActivity.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListPullReqCommits**
> []TypesCommit ListPullReqCommits(ctx, accountIdentifier, repoIdentifier, pullreqNumber, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity.. | 
  **repoIdentifier** | **string**|  | 
  **pullreqNumber** | **int32**|  | 
 **optional** | ***PullreqApiListPullReqCommitsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PullreqApiListPullReqCommitsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity.. | 
 **projectIdentifier** | **optional.String**| Project Identifier for the Entity.. | 
 **page** | **optional.Int32**| The page to return. | [default to 1]
 **limit** | **optional.Int32**| The maximum number of results to return. | [default to 30]

### Return type

[**[]TypesCommit**](TypesCommit.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MergePullReqOp**
> TypesMergeResponse MergePullReqOp(ctx, accountIdentifier, repoIdentifier, pullreqNumber, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity.. | 
  **repoIdentifier** | **string**|  | 
  **pullreqNumber** | **int32**|  | 
 **optional** | ***PullreqApiMergePullReqOpOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PullreqApiMergePullReqOpOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | [**optional.Interface of OpenapiMergePullReq**](OpenapiMergePullReq.md)|  | 
 **orgIdentifier** | **optional.**| Organization Identifier for the Entity.. | 
 **projectIdentifier** | **optional.**| Project Identifier for the Entity.. | 

### Return type

[**TypesMergeResponse**](TypesMergeResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PullReqMetaData**
> TypesPullReqStats PullReqMetaData(ctx, accountIdentifier, repoIdentifier, pullreqNumber, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity.. | 
  **repoIdentifier** | **string**|  | 
  **pullreqNumber** | **int32**|  | 
 **optional** | ***PullreqApiPullReqMetaDataOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PullreqApiPullReqMetaDataOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity.. | 
 **projectIdentifier** | **optional.String**| Project Identifier for the Entity.. | 

### Return type

[**TypesPullReqStats**](TypesPullReqStats.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReviewSubmitPullReq**
> ReviewSubmitPullReq(ctx, accountIdentifier, repoIdentifier, pullreqNumber, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity.. | 
  **repoIdentifier** | **string**|  | 
  **pullreqNumber** | **int32**|  | 
 **optional** | ***PullreqApiReviewSubmitPullReqOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PullreqApiReviewSubmitPullReqOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | [**optional.Interface of OpenapiReviewSubmitPullReqRequest**](OpenapiReviewSubmitPullReqRequest.md)|  | 
 **orgIdentifier** | **optional.**| Organization Identifier for the Entity.. | 
 **projectIdentifier** | **optional.**| Project Identifier for the Entity.. | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReviewerAddPullReq**
> TypesPullReqReviewer ReviewerAddPullReq(ctx, accountIdentifier, repoIdentifier, pullreqNumber, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity.. | 
  **repoIdentifier** | **string**|  | 
  **pullreqNumber** | **int32**|  | 
 **optional** | ***PullreqApiReviewerAddPullReqOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PullreqApiReviewerAddPullReqOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | [**optional.Interface of OpenapiReviewerAddPullReqRequest**](OpenapiReviewerAddPullReqRequest.md)|  | 
 **orgIdentifier** | **optional.**| Organization Identifier for the Entity.. | 
 **projectIdentifier** | **optional.**| Project Identifier for the Entity.. | 

### Return type

[**TypesPullReqReviewer**](TypesPullReqReviewer.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReviewerDeletePullReq**
> ReviewerDeletePullReq(ctx, accountIdentifier, repoIdentifier, pullreqNumber, pullreqReviewerId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity.. | 
  **repoIdentifier** | **string**|  | 
  **pullreqNumber** | **int32**|  | 
  **pullreqReviewerId** | **int32**|  | 
 **optional** | ***PullreqApiReviewerDeletePullReqOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PullreqApiReviewerDeletePullReqOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity.. | 
 **projectIdentifier** | **optional.String**| Project Identifier for the Entity.. | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReviewerListPullReq**
> []TypesPullReqReviewer ReviewerListPullReq(ctx, accountIdentifier, repoIdentifier, pullreqNumber, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity.. | 
  **repoIdentifier** | **string**|  | 
  **pullreqNumber** | **int32**|  | 
 **optional** | ***PullreqApiReviewerListPullReqOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PullreqApiReviewerListPullReqOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity.. | 
 **projectIdentifier** | **optional.String**| Project Identifier for the Entity.. | 

### Return type

[**[]TypesPullReqReviewer**](TypesPullReqReviewer.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StatePullReq**
> TypesPullReq StatePullReq(ctx, accountIdentifier, repoIdentifier, pullreqNumber, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity.. | 
  **repoIdentifier** | **string**|  | 
  **pullreqNumber** | **int32**|  | 
 **optional** | ***PullreqApiStatePullReqOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PullreqApiStatePullReqOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | [**optional.Interface of OpenapiStatePullReqRequest**](OpenapiStatePullReqRequest.md)|  | 
 **orgIdentifier** | **optional.**| Organization Identifier for the Entity.. | 
 **projectIdentifier** | **optional.**| Project Identifier for the Entity.. | 

### Return type

[**TypesPullReq**](TypesPullReq.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdatePullReq**
> TypesPullReq UpdatePullReq(ctx, accountIdentifier, repoIdentifier, pullreqNumber, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity.. | 
  **repoIdentifier** | **string**|  | 
  **pullreqNumber** | **int32**|  | 
 **optional** | ***PullreqApiUpdatePullReqOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PullreqApiUpdatePullReqOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | [**optional.Interface of OpenapiUpdatePullReqRequest**](OpenapiUpdatePullReqRequest.md)|  | 
 **orgIdentifier** | **optional.**| Organization Identifier for the Entity.. | 
 **projectIdentifier** | **optional.**| Project Identifier for the Entity.. | 

### Return type

[**TypesPullReq**](TypesPullReq.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

