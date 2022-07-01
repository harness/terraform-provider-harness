# nextgen{{classname}}

All URIs are relative to *https://app.harness.io/gateway*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteInputSet**](PipelineInputSetApi.md#DeleteInputSet) | **Delete** /pipeline/api/inputSets/{inputSetIdentifier} | Delete Input Set
[**GetInputSet**](PipelineInputSetApi.md#GetInputSet) | **Get** /pipeline/api/inputSets/{inputSetIdentifier} | Fetch Input Set
[**ListInputSet**](PipelineInputSetApi.md#ListInputSet) | **Get** /pipeline/api/inputSets | List Input Sets
[**PostInputSet**](PipelineInputSetApi.md#PostInputSet) | **Post** /pipeline/api/inputSets | Create Input Set
[**PutInputSet**](PipelineInputSetApi.md#PutInputSet) | **Put** /pipeline/api/inputSets/{inputSetIdentifier} | Update Input Set
[**RuntimeInputTemplate**](PipelineInputSetApi.md#RuntimeInputTemplate) | **Post** /pipeline/api/inputSets/template | Fetch Runtime Input Template

# **DeleteInputSet**
> ResponseDtoBoolean DeleteInputSet(ctx, inputSetIdentifier, accountIdentifier, orgIdentifier, projectIdentifier, pipelineIdentifier, optional)
Delete Input Set

Deletes the Input Set by Identifier

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **inputSetIdentifier** | **string**| Identifier of the Input Set that should be deleted. | 
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 
  **orgIdentifier** | **string**| Organization Identifier for the Entity. | 
  **projectIdentifier** | **string**| Project Identifier for the Entity. | 
  **pipelineIdentifier** | **string**| Pipeline Identifier for the entity. | 
 **optional** | ***PipelineInputSetApiDeleteInputSetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PipelineInputSetApiDeleteInputSetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **ifMatch** | **optional.String**| Version of Entity to match | 
 **branch** | **optional.String**| Name of the branch. | 
 **repoIdentifier** | **optional.String**| Git Sync Config Id. | 
 **rootFolder** | **optional.String**| Path to the root folder of the Entity. | 
 **filePath** | **optional.String**| File Path of the Entity. | 
 **commitMsg** | **optional.String**| Commit Message to use for the merge commit. | 
 **lastObjectId** | **optional.String**| Last Object Id | 

### Return type

[**ResponseDtoBoolean**](ResponseDTOBoolean.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetInputSet**
> ResponseDtoInputSetResponse GetInputSet(ctx, inputSetIdentifier, accountIdentifier, orgIdentifier, projectIdentifier, pipelineIdentifier, optional)
Fetch Input Set

Returns Input Set for a Given Identifier (Throws an Error if no Input Set Exists)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **inputSetIdentifier** | **string**| Identifier for the Input Set | 
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 
  **orgIdentifier** | **string**| Organization Identifier for the Entity. | 
  **projectIdentifier** | **string**| Project Identifier for the Entity. | 
  **pipelineIdentifier** | **string**| Pipeline Identifier for the entity. | 
 **optional** | ***PipelineInputSetApiGetInputSetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PipelineInputSetApiGetInputSetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **branch** | **optional.String**| Name of the branch. | 
 **repoIdentifier** | **optional.String**| Git Sync Config Id. | 
 **getDefaultFromOtherRepo** | **optional.Bool**| if true, return all the default entities | 

### Return type

[**ResponseDtoInputSetResponse**](ResponseDTOInputSetResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListInputSet**
> ResponseDtoPageResponseInputSetSummaryResponse ListInputSet(ctx, accountIdentifier, orgIdentifier, projectIdentifier, pipelineIdentifier, optional)
List Input Sets

Lists all Input Sets for a Pipeline

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 
  **orgIdentifier** | **string**| Organization Identifier for the Entity. | 
  **projectIdentifier** | **string**| Project Identifier for the Entity. | 
  **pipelineIdentifier** | **string**| Pipeline identifier for which we need the Input Sets list. | 
 **optional** | ***PipelineInputSetApiListInputSetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PipelineInputSetApiListInputSetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **pageIndex** | **optional.Int32**| Page Index of the results to fetch.Default Value: 0 | [default to 0]
 **pageSize** | **optional.Int32**| Results per page | [default to 100]
 **inputSetType** | **optional.String**| Type of Input Set. The default value is ALL. | [default to ALL]
 **searchTerm** | **optional.String**| Search term to filter out Input Sets based on name, identifier, tags. | 
 **sortOrders** | [**optional.Interface of []string**](string.md)| Sort criteria for the elements. | 
 **branch** | **optional.String**| Name of the branch. | 
 **repoIdentifier** | **optional.String**| Git Sync Config Id. | 
 **getDefaultFromOtherRepo** | **optional.Bool**| if true, return all the default entities | 

### Return type

[**ResponseDtoPageResponseInputSetSummaryResponse**](ResponseDTOPageResponseInputSetSummaryResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostInputSet**
> ResponseDtoInputSetResponse PostInputSet(ctx, body, accountIdentifier, orgIdentifier, projectIdentifier, pipelineIdentifier, optional)
Create Input Set

Creates an Input Set for a Pipeline

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**string**](string.md)| Input set YAML to be created. The Account, Org, Project, and Pipeline identifiers inside the YAML should match the query parameters. | 
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 
  **orgIdentifier** | **string**| Organization Identifier for the Entity. | 
  **projectIdentifier** | **string**| Project Identifier for the Entity. | 
  **pipelineIdentifier** | **string**| Pipeline Identifier for the entity. | 
 **optional** | ***PipelineInputSetApiPostInputSetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PipelineInputSetApiPostInputSetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **pipelineBranch** | **optional.**| Github branch of the Pipeline for which the Input Set is to be created | 
 **pipelineRepoID** | **optional.**| Github Repo identifier of the Pipeline for which the Input Set is to be created | 
 **branch** | **optional.**| Name of the branch. | 
 **repoIdentifier** | **optional.**| Git Sync Config Id. | 
 **rootFolder** | **optional.**| Path to the root folder of the Entity. | 
 **filePath** | **optional.**| File Path of the Entity. | 
 **commitMsg** | **optional.**| File Path of the Entity. | 
 **isNewBranch** | **optional.**| Checks the new branch | [default to false]
 **baseBranch** | **optional.**| Name of the default branch. | 
 **connectorRef** | **optional.**| Identifier of Connector needed for CRUD operations on the respective Entity | 
 **storeType** | **optional.**| Tells whether the Entity is to be saved on Git or not | 
 **repoName** | **optional.**| Name of the repository. | 

### Return type

[**ResponseDtoInputSetResponse**](ResponseDTOInputSetResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutInputSet**
> ResponseDtoInputSetResponse PutInputSet(ctx, body, inputSetIdentifier, accountIdentifier, orgIdentifier, projectIdentifier, pipelineIdentifier, optional)
Update Input Set

Updates the Input Set for a Pipeline

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**string**](string.md)| Input set YAML to be updated. The query parameters should match the Account, Org, Project, and Pipeline Ids in the YAML. | 
  **inputSetIdentifier** | **string**| Identifier for the Input Set that needs to be updated. An Input Set corresponding to this identifier should already exist. | 
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 
  **orgIdentifier** | **string**| Organization Identifier for the Entity. | 
  **projectIdentifier** | **string**| Project Identifier for the Entity. | 
  **pipelineIdentifier** | **string**| Pipeline Identifier for the entity. | 
 **optional** | ***PipelineInputSetApiPutInputSetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PipelineInputSetApiPutInputSetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------






 **ifMatch** | **optional.**| Version of Entity to match | 
 **pipelineBranch** | **optional.**| Github branch of the Pipeline for which the Input Set is to be updated | 
 **pipelineRepoID** | **optional.**| Github Repo Id of the Pipeline for which the Input Set is to be updated | 
 **branch** | **optional.**| Name of the branch. | 
 **repoIdentifier** | **optional.**| Git Sync Config Id. | 
 **rootFolder** | **optional.**| Path to the root folder of the Entity. | 
 **filePath** | **optional.**| Path to the root folder of the Entity. | 
 **commitMsg** | **optional.**| Commit Message to use for the merge commit. | 
 **lastObjectId** | **optional.**| Last Object Id | 
 **resolvedConflictCommitId** | **optional.**| If the entity is git-synced, this parameter represents the commit id against which file conflicts are resolved | 
 **baseBranch** | **optional.**| Name of the default branch. | 
 **connectorRef** | **optional.**| Identifier of Connector needed for CRUD operations on the respective Entity | 

### Return type

[**ResponseDtoInputSetResponse**](ResponseDTOInputSetResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RuntimeInputTemplate**
> ResponseDtoInputSetTemplateWithReplacedExpressionsResponse RuntimeInputTemplate(ctx, accountIdentifier, orgIdentifier, projectIdentifier, pipelineIdentifier, optional)
Fetch Runtime Input Template

Returns Runtime Input Template for a Pipeline

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 
  **orgIdentifier** | **string**| Organization Identifier for the Entity. | 
  **projectIdentifier** | **string**| Project Identifier for the Entity. | 
  **pipelineIdentifier** | **string**| Pipeline identifier for which we need the Runtime Input Template. | 
 **optional** | ***PipelineInputSetApiRuntimeInputTemplateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PipelineInputSetApiRuntimeInputTemplateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **body** | [**optional.Interface of InputSetTemplateRequest**](InputSetTemplateRequest.md)|  | 
 **branch** | **optional.**| Name of the branch. | 
 **repoIdentifier** | **optional.**| Git Sync Config Id. | 
 **getDefaultFromOtherRepo** | **optional.**| if true, return all the default entities | 

### Return type

[**ResponseDtoInputSetTemplateWithReplacedExpressionsResponse**](ResponseDTOInputSetTemplateWithReplacedExpressionsResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

