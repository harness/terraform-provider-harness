# nextgen{{classname}}

All URIs are relative to *https://app.harness.io/gateway*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePR**](SCMApi.md#CreatePR) | **Post** /ng/api/scm/createPR | creates a Pull Request
[**GetFileContent**](SCMApi.md#GetFileContent) | **Get** /ng/api/scm/fileContent | Gets Git File Content
[**GetListOfBranchesByConnector**](SCMApi.md#GetListOfBranchesByConnector) | **Get** /ng/api/scm/listRepoBranches | Lists Branches of given Repo by referenced Connector Identifier
[**GetListOfBranchesByGitConfig**](SCMApi.md#GetListOfBranchesByGitConfig) | **Get** /ng/api/scm/listBranchesByGitConfig | Lists Branches by given Git Sync Config Id

# **CreatePR**
> ResponseDtoprDetails CreatePR(ctx, body)
creates a Pull Request

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**GitPrCreateRequest**](GitPrCreateRequest.md)| Details to create a PR | 

### Return type

[**ResponseDtoprDetails**](ResponseDTOPRDetails.md)

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
  **yamlGitConfigIdentifier** | **string**| Git Sync Config Id. | 
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 
  **filePath** | **string**| File Path | 
 **optional** | ***SCMApiGetFileContentOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SCMApiGetFileContentOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity. | 
 **projectIdentifier** | **optional.String**| Project Identifier for the Entity. | 
 **branch** | **optional.String**| Name of the branch. | 
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
 **optional** | ***SCMApiGetListOfBranchesByConnectorOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SCMApiGetListOfBranchesByConnectorOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **connectorIdentifierRef** | **optional.String**| Connector Identifier Reference | 
 **accountIdentifier** | **optional.String**| Account Identifier for the Entity. | 
 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity. | 
 **projectIdentifier** | **optional.String**| Project Identifier for the Entity. | 
 **repoURL** | **optional.String**| URL of the repository. | 
 **page** | **optional.Int32**| Page Index of the results to fetch.Default Value: 0 | [default to 0]
 **size** | **optional.Int32**| Results per page | [default to 50]
 **searchTerm** | **optional.String**| Search Term. | 

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
 **optional** | ***SCMApiGetListOfBranchesByGitConfigOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SCMApiGetListOfBranchesByGitConfigOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **yamlGitConfigIdentifier** | **optional.String**| Git Sync Config Id. | 
 **accountIdentifier** | **optional.String**| Account Identifier for the Entity. | 
 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity. | 
 **projectIdentifier** | **optional.String**| Project Identifier for the Entity. | 
 **page** | **optional.Int32**| Page Index of the results to fetch.Default Value: 0 | [default to 0]
 **size** | **optional.Int32**| Results per page | [default to 50]
 **searchTerm** | **optional.String**| Search Term. | 

### Return type

[**ResponseDtoListString**](ResponseDTOListString.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml, text/yaml, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

