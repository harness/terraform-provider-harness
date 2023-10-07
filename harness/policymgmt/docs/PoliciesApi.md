# {{classname}}

All URIs are relative to *https://app.harness.io/gateway/pm/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PoliciesCreate**](PoliciesApi.md#PoliciesCreate) | **Post** /api/v1/policies | 
[**PoliciesDelete**](PoliciesApi.md#PoliciesDelete) | **Delete** /api/v1/policies/{identifier} | 
[**PoliciesFind**](PoliciesApi.md#PoliciesFind) | **Get** /api/v1/policies/{identifier} | 
[**PoliciesList**](PoliciesApi.md#PoliciesList) | **Get** /api/v1/policies | 
[**PoliciesUpdate**](PoliciesApi.md#PoliciesUpdate) | **Patch** /api/v1/policies/{identifier} | 

# **PoliciesCreate**
> Policy PoliciesCreate(ctx, body, optional)


Create a policy

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateRequestBody**](CreateRequestBody.md)|  | 
 **optional** | ***PoliciesApiPoliciesCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PoliciesApiPoliciesCreateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accountIdentifier** | **optional.**| Harness account ID | 
 **orgIdentifier** | **optional.**| Harness organization ID | 
 **projectIdentifier** | **optional.**| Harness project ID | 
 **gitCommitMsg** | **optional.**| The commit message used in git when creating the policy | 
 **gitImport** | **optional.**| A flag to determine if the api should try and import and existing policy from git | 
 **gitBranch** | **optional.**| The git branch the policy will be created in | 
 **gitIsNewBranch** | **optional.**| A flag to determine if the api should try and commit to a new branch | 
 **gitBaseBranch** | **optional.**| If committing to a new branch, git_base_branch tells the api which branch to base the new branch from | 
 **xApiKey** | **optional.**| Harness PAT key used to perform authorization | 

### Return type

[**Policy**](Policy.md)

### Authorization

[api_key_header_x-api-key](../README.md#api_key_header_x-api-key), [jwt_header_Authorization](../README.md#jwt_header_Authorization)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PoliciesDelete**
> PoliciesDelete(ctx, identifier, optional)


Delete a policy by identifier

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **identifier** | **string**| Identifier of the policy | 
 **optional** | ***PoliciesApiPoliciesDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PoliciesApiPoliciesDeleteOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accountIdentifier** | **optional.String**| Harness account ID | 
 **orgIdentifier** | **optional.String**| Harness organization ID | 
 **projectIdentifier** | **optional.String**| Harness project ID | 
 **xApiKey** | **optional.String**| Harness PAT key used to perform authorization | 

### Return type

 (empty response body)

### Authorization

[api_key_header_x-api-key](../README.md#api_key_header_x-api-key), [jwt_header_Authorization](../README.md#jwt_header_Authorization)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PoliciesFind**
> Policy PoliciesFind(ctx, identifier, optional)


Find a policy by identifier

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **identifier** | **string**| Identifier of the policy to retrieve | 
 **optional** | ***PoliciesApiPoliciesFindOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PoliciesApiPoliciesFindOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accountIdentifier** | **optional.String**| Harness account ID | 
 **orgIdentifier** | **optional.String**| Harness organization ID | 
 **projectIdentifier** | **optional.String**| Harness project ID | 
 **gitBranch** | **optional.String**| The git branch the policy resides in | 
 **showSummary** | **optional.Bool**| Setting to true returns the metadata about the        requested policy including the information held about the status of this policy in the default branch.        git_branch is ignored as no git operation takes place. | 
 **xApiKey** | **optional.String**| Harness PAT key used to perform authorization | 

### Return type

[**Policy**](Policy.md)

### Authorization

[api_key_header_x-api-key](../README.md#api_key_header_x-api-key), [jwt_header_Authorization](../README.md#jwt_header_Authorization)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PoliciesList**
> []Policy PoliciesList(ctx, optional)


List all policies

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PoliciesApiPoliciesListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PoliciesApiPoliciesListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountIdentifier** | **optional.String**| Harness account ID | 
 **orgIdentifier** | **optional.String**| Harness organization ID | 
 **projectIdentifier** | **optional.String**| Harness project ID | 
 **perPage** | **optional.Int32**| Number of results per page | [default to 50]
 **page** | **optional.Int32**| Page number (starting from 0) | [default to 0]
 **searchTerm** | **optional.String**| Filter results by partial name match | 
 **sort** | **optional.String**| Sort order for results | [default to name,ASC]
 **xApiKey** | **optional.String**| Harness PAT key used to perform authorization | 

### Return type

[**[]Policy**](Policy.md)

### Authorization

[api_key_header_x-api-key](../README.md#api_key_header_x-api-key), [jwt_header_Authorization](../README.md#jwt_header_Authorization)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PoliciesUpdate**
> PoliciesUpdate(ctx, body, identifier, optional)


Update a policy by identifier

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateRequestBody**](UpdateRequestBody.md)|  | 
  **identifier** | **string**| Identifier of the policy | 
 **optional** | ***PoliciesApiPoliciesUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PoliciesApiPoliciesUpdateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **gitCommitMsg** | **optional.**| The commit message used in git when creating the policy | 
 **gitIsNewBranch** | **optional.**| A flag to determine if the api should try and commit to a new branch | 
 **gitBaseBranch** | **optional.**| If committing to a new branch, git_base_branch tells the api which branch to base the new branch from | 
 **gitBranch** | **optional.**| The git branch the policy resides in | 
 **gitCommitSha** | **optional.**| The existing commit sha of the file being updated | 
 **gitFileId** | **optional.**| The existing file id of the file being updated, not required for bitbucket files | 
 **accountIdentifier** | **optional.**| Harness account ID | 
 **orgIdentifier** | **optional.**| Harness organization ID | 
 **projectIdentifier** | **optional.**| Harness project ID | 
 **xApiKey** | **optional.**| Harness PAT key used to perform authorization | 

### Return type

 (empty response body)

### Authorization

[api_key_header_x-api-key](../README.md#api_key_header_x-api-key), [jwt_header_Authorization](../README.md#jwt_header_Authorization)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

