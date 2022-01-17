# {{classname}}

All URIs are relative to *https://app.harness.io/gateway*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePR**](ScmApi.md#CreatePR) | **Post** /ng/api/scm/createPR | creates a Pull Request
[**GetFileContent**](ScmApi.md#GetFileContent) | **Get** /ng/api/scm/fileContent | Gets Git File Content
[**GetListOfBranchesByConnector**](ScmApi.md#GetListOfBranchesByConnector) | **Get** /ng/api/scm/listRepoBranches | Lists Branches of given Repo by referenced Connector Identifier
[**GetListOfBranchesByGitConfig**](ScmApi.md#GetListOfBranchesByGitConfig) | **Get** /ng/api/scm/listBranchesByGitConfig | Lists Branches by given Git Sync Config Id
[**IsSaasGit**](ScmApi.md#IsSaasGit) | **Post** /ng/api/scm/isSaasGit | Checks if Saas is possible for given Repo Url

# **CreatePR**
> ResponseDtoCreatePr CreatePR(ctx, body)
creates a Pull Request

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**GitPrCreateRequest**](GitPrCreateRequest.md)| Details to create a PR | 

### Return type

[**ResponseDtoCreatePr**](ResponseDTOCreatePR.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json, text/yaml, text/html
 - **Accept**: application/json, application/yaml, text/yaml, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFileContent**
> ResponseDtoGitFileContent GetFileContent(ctx, yamlGitConfigIdentifier, accountIdentifier, filePath, optional)
Gets Git File Content

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **yamlGitConfigIdentifier** | **string**| Git Sync Config Id | 
  **accountIdentifier** | **string**| Account Identifier for the Entity | 
  **filePath** | **string**| File Path | 
 **optional** | ***ScmApiGetFileContentOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ScmApiGetFileContentOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity | 
 **projectIdentifier** | **optional.String**| Project Identifier for the Entity | 
 **branch** | **optional.String**| Branch Name | 
 **commitId** | **optional.String**| Commit Id | 

### Return type

[**ResponseDtoGitFileContent**](ResponseDTOGitFileContent.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml, text/yaml, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetListOfBranchesByConnector**
> ResponseDtoListString GetListOfBranchesByConnector(ctx, optional)
Lists Branches of given Repo by referenced Connector Identifier

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ScmApiGetListOfBranchesByConnectorOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ScmApiGetListOfBranchesByConnectorOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **connectorIdentifierRef** | **optional.String**| Connector Identifier Reference | 
 **accountIdentifier** | **optional.String**| Account Identifier for the Entity | 
 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity | 
 **projectIdentifier** | **optional.String**| Project Identifier for the Entity | 
 **repoURL** | **optional.String**| Repo Url | 
 **page** | **optional.Int32**| Indicates the number of pages. Results for these pages will be retrieved. | [default to 0]
 **size** | **optional.Int32**| The number of the elements to fetch | [default to 50]
 **searchTerm** | **optional.String**| Search Term | 

### Return type

[**ResponseDtoListString**](ResponseDTOListString.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml, text/yaml, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetListOfBranchesByGitConfig**
> ResponseDtoListString GetListOfBranchesByGitConfig(ctx, optional)
Lists Branches by given Git Sync Config Id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ScmApiGetListOfBranchesByGitConfigOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ScmApiGetListOfBranchesByGitConfigOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **yamlGitConfigIdentifier** | **optional.String**| Git Sync Config Id | 
 **accountIdentifier** | **optional.String**| Account Identifier for the Entity | 
 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity | 
 **projectIdentifier** | **optional.String**| Project Identifier for the Entity | 
 **page** | **optional.Int32**| Indicates the number of pages. Results for these pages will be retrieved. | [default to 0]
 **size** | **optional.Int32**| The number of the elements to fetch | [default to 50]
 **searchTerm** | **optional.String**| Search Term | 

### Return type

[**ResponseDtoListString**](ResponseDTOListString.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml, text/yaml, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **IsSaasGit**
> ResponseDtoSaasGit IsSaasGit(ctx, optional)
Checks if Saas is possible for given Repo Url

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ScmApiIsSaasGitOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ScmApiIsSaasGitOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **repoURL** | **optional.String**| Repo Url | 

### Return type

[**ResponseDtoSaasGit**](ResponseDTOSaasGit.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml, text/yaml, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

