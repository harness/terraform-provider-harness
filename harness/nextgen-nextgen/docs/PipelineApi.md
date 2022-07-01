# nextgen{{classname}}

All URIs are relative to *https://app.harness.io/gateway*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeletePipeline**](PipelineApi.md#DeletePipeline) | **Delete** /pipeline/api/pipelines/{pipelineIdentifier} | Delete a Pipeline
[**GetPipeline**](PipelineApi.md#GetPipeline) | **Get** /pipeline/api/pipelines/{pipelineIdentifier} | Fetch a Pipeline
[**GetPipelineList**](PipelineApi.md#GetPipelineList) | **Post** /pipeline/api/pipelines/list | List Pipelines
[**GetPipelineSummary**](PipelineApi.md#GetPipelineSummary) | **Get** /pipeline/api/pipelines/summary/{pipelineIdentifier} | Fetch Pipeline Summary
[**PostPipeline**](PipelineApi.md#PostPipeline) | **Post** /pipeline/api/pipelines | Create a Pipeline
[**PostPipelineV2**](PipelineApi.md#PostPipelineV2) | **Post** /pipeline/api/pipelines/v2 | Create a Pipeline
[**UpdatePipeline**](PipelineApi.md#UpdatePipeline) | **Put** /pipeline/api/pipelines/{pipelineIdentifier} | Update a Pipeline
[**UpdatePipelineV2**](PipelineApi.md#UpdatePipelineV2) | **Put** /pipeline/api/pipelines/v2/{pipelineIdentifier} | Update a Pipeline

# **DeletePipeline**
> ResponseDtoBoolean DeletePipeline(ctx, accountIdentifier, orgIdentifier, projectIdentifier, pipelineIdentifier, optional)
Delete a Pipeline

Deletes a Pipeline by Identifier

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 
  **orgIdentifier** | **string**| Organization Identifier for the Entity. | 
  **projectIdentifier** | **string**| Project Identifier for the Entity. | 
  **pipelineIdentifier** | **string**| Pipeline Identifier | 
 **optional** | ***PipelineApiDeletePipelineOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PipelineApiDeletePipelineOpts struct
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

# **GetPipeline**
> ResponseDtopmsPipelineResponse GetPipeline(ctx, accountIdentifier, orgIdentifier, projectIdentifier, pipelineIdentifier, optional)
Fetch a Pipeline

Returns a Pipeline by Identifier

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 
  **orgIdentifier** | **string**| Organization Identifier for the Entity. | 
  **projectIdentifier** | **string**| Project Identifier for the Entity. | 
  **pipelineIdentifier** | **string**| Pipeline Identifier | 
 **optional** | ***PipelineApiGetPipelineOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PipelineApiGetPipelineOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **branch** | **optional.String**| Name of the branch. | 
 **repoIdentifier** | **optional.String**| Git Sync Config Id. | 
 **getDefaultFromOtherRepo** | **optional.Bool**| if true, return all the default entities | 
 **getTemplatesResolvedPipeline** | **optional.Bool**| This is a boolean value. If true, returns Templates resolved Pipeline YAML in the response else returns null. | [default to false]

### Return type

[**ResponseDtopmsPipelineResponse**](ResponseDTOPMSPipelineResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPipelineList**
> ResponseDtoPagePmsPipelineSummaryResponse GetPipelineList(ctx, accountIdentifier, orgIdentifier, projectIdentifier, optional)
List Pipelines

Returns List of Pipelines in the Given Project

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 
  **orgIdentifier** | **string**| Organization Identifier for the Entity. | 
  **projectIdentifier** | **string**| Project Identifier for the Entity. | 
 **optional** | ***PipelineApiGetPipelineListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PipelineApiGetPipelineListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | [**optional.Interface of PipelineFilterProperties**](PipelineFilterProperties.md)| This is the body for the filter properties for listing pipelines. | 
 **page** | **optional.**| Page Index of the results to fetch.Default Value: 0 | [default to 0]
 **size** | **optional.**| Results per page | [default to 25]
 **sort** | [**optional.Interface of []string**](string.md)| Sort criteria for the elements. | 
 **searchTerm** | **optional.**| Search term to filter out pipelines based on pipeline name, identifier, tags. | 
 **module** | **optional.**|  | 
 **filterIdentifier** | **optional.**|  | 
 **branch** | **optional.**| Name of the branch. | 
 **repoIdentifier** | **optional.**| Git Sync Config Id. | 
 **getDefaultFromOtherRepo** | **optional.**| if true, return all the default entities | 
 **getDistinctFromBranches** | **optional.**| Boolean flag to get distinct pipelines from all branches. | 

### Return type

[**ResponseDtoPagePmsPipelineSummaryResponse**](ResponseDTOPagePMSPipelineSummaryResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPipelineSummary**
> ResponseDtopmsPipelineSummaryResponse GetPipelineSummary(ctx, accountIdentifier, orgIdentifier, projectIdentifier, pipelineIdentifier, optional)
Fetch Pipeline Summary

Returns Pipeline Summary by Identifier

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 
  **orgIdentifier** | **string**| Organization Identifier for the Entity. | 
  **projectIdentifier** | **string**| Project Identifier for the Entity. | 
  **pipelineIdentifier** | **string**| Pipeline Identifier | 
 **optional** | ***PipelineApiGetPipelineSummaryOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PipelineApiGetPipelineSummaryOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **branch** | **optional.String**| Name of the branch. | 
 **repoIdentifier** | **optional.String**| Git Sync Config Id. | 
 **getDefaultFromOtherRepo** | **optional.Bool**| if true, return all the default entities | 

### Return type

[**ResponseDtopmsPipelineSummaryResponse**](ResponseDTOPMSPipelineSummaryResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostPipeline**
> ResponseDtoString PostPipeline(ctx, body, accountIdentifier, orgIdentifier, projectIdentifier, optional)
Create a Pipeline

Creates a Pipeline

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**string**](string.md)| Pipeline YAML | 
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 
  **orgIdentifier** | **string**| Organization Identifier for the Entity. | 
  **projectIdentifier** | **string**| Project Identifier for the Entity. | 
 **optional** | ***PipelineApiPostPipelineOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PipelineApiPostPipelineOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




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

[**ResponseDtoString**](ResponseDTOString.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostPipelineV2**
> ResponseDtoPipelineSaveResponse PostPipelineV2(ctx, body, accountIdentifier, orgIdentifier, projectIdentifier, optional)
Create a Pipeline

Creates a Pipeline

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**string**](string.md)| Pipeline YAML | 
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 
  **orgIdentifier** | **string**| Organization Identifier for the Entity. | 
  **projectIdentifier** | **string**| Project Identifier for the Entity. | 
 **optional** | ***PipelineApiPostPipelineV2Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PipelineApiPostPipelineV2Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




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

[**ResponseDtoPipelineSaveResponse**](ResponseDTOPipelineSaveResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdatePipeline**
> ResponseDtoString UpdatePipeline(ctx, body, accountIdentifier, orgIdentifier, projectIdentifier, pipelineIdentifier, optional)
Update a Pipeline

Updates a Pipeline by Identifier

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**string**](string.md)| Pipeline YAML to be updated | 
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 
  **orgIdentifier** | **string**| Organization Identifier for the Entity. | 
  **projectIdentifier** | **string**| Project Identifier for the Entity. | 
  **pipelineIdentifier** | **string**| Pipeline Identifier | 
 **optional** | ***PipelineApiUpdatePipelineOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PipelineApiUpdatePipelineOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **ifMatch** | **optional.**| Version of Entity to match | 
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

[**ResponseDtoString**](ResponseDTOString.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdatePipelineV2**
> ResponseDtoPipelineSaveResponse UpdatePipelineV2(ctx, body, accountIdentifier, orgIdentifier, projectIdentifier, pipelineIdentifier, optional)
Update a Pipeline

Updates a Pipeline by Identifier

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**string**](string.md)| Pipeline YAML to be updated | 
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 
  **orgIdentifier** | **string**| Organization Identifier for the Entity. | 
  **projectIdentifier** | **string**| Project Identifier for the Entity. | 
  **pipelineIdentifier** | **string**| Pipeline Identifier | 
 **optional** | ***PipelineApiUpdatePipelineV2Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PipelineApiUpdatePipelineV2Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **ifMatch** | **optional.**| Version of Entity to match | 
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

[**ResponseDtoPipelineSaveResponse**](ResponseDTOPipelineSaveResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

