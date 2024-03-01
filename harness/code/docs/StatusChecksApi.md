# {{classname}}

All URIs are relative to */gateway/code/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListStatusCheckRecent**](StatusChecksApi.md#ListStatusCheckRecent) | **Get** /repos/{repo_identifier}/checks/recent | List recent status check
[**ListStatusCheckResults**](StatusChecksApi.md#ListStatusCheckResults) | **Get** /repos/{repo_identifier}/checks/commits/{commit_sha} | List status check results
[**ReportStatusCheckResults**](StatusChecksApi.md#ReportStatusCheckResults) | **Put** /repos/{repo_identifier}/checks/commits/{commit_sha} | Report status check results

# **ListStatusCheckRecent**
> []string ListStatusCheckRecent(ctx, accountIdentifier, repoIdentifier, optional)
List recent status check

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity.. | 
  **repoIdentifier** | **string**|  | 
 **optional** | ***StatusChecksApiListStatusCheckRecentOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a StatusChecksApiListStatusCheckRecentOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity.. | 
 **projectIdentifier** | **optional.String**| Project Identifier for the Entity.. | 
 **query** | **optional.String**| The substring which is used to filter the status checks by their UID. | 
 **since** | **optional.Int32**| The timestamp (in Unix time millis) since the status checks have been run. | 

### Return type

**[]string**

### Authorization

[bearerAuth](../README.md#bearerAuth), [x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListStatusCheckResults**
> []TypesCheck ListStatusCheckResults(ctx, accountIdentifier, repoIdentifier, commitSha, optional)
List status check results

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity.. | 
  **repoIdentifier** | **string**|  | 
  **commitSha** | **string**|  | 
 **optional** | ***StatusChecksApiListStatusCheckResultsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a StatusChecksApiListStatusCheckResultsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity.. | 
 **projectIdentifier** | **optional.String**| Project Identifier for the Entity.. | 
 **page** | **optional.Int32**| The page to return. | [default to 1]
 **limit** | **optional.Int32**| The maximum number of results to return. | [default to 30]
 **query** | **optional.String**| The substring which is used to filter the status checks by their UID. | 

### Return type

[**[]TypesCheck**](TypesCheck.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReportStatusCheckResults**
> TypesCheck ReportStatusCheckResults(ctx, accountIdentifier, repoIdentifier, commitSha, optional)
Report status check results

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity.. | 
  **repoIdentifier** | **string**|  | 
  **commitSha** | **string**|  | 
 **optional** | ***StatusChecksApiReportStatusCheckResultsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a StatusChecksApiReportStatusCheckResultsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | [**optional.Interface of CommitsCommitShaBody**](CommitsCommitShaBody.md)|  | 
 **orgIdentifier** | **optional.**| Organization Identifier for the Entity.. | 
 **projectIdentifier** | **optional.**| Project Identifier for the Entity.. | 

### Return type

[**TypesCheck**](TypesCheck.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

