# {{classname}}

All URIs are relative to *https://app.harness.io/gateway*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateVariables**](PipelinesApi.md#CreateVariables) | **Post** /pipeline/api/pipelines/variables | Get all the Variables which can be used as expression in the Pipeline.
[**DeletePipeline**](PipelinesApi.md#DeletePipeline) | **Delete** /pipeline/api/pipelines/{pipelineIdentifier} | Deletes a Pipeline
[**GetExecutionNode**](PipelinesApi.md#GetExecutionNode) | **Get** /pipeline/api/pipelines/getExecutionNode | Get the Execution Node by Execution Id
[**GetExpandedPipelineJSON**](PipelinesApi.md#GetExpandedPipelineJSON) | **Get** /pipeline/api/pipelines/expandedJSON/{pipelineIdentifier} | Gets Pipeline JSON with extra info for some fields
[**GetNotificationSchema**](PipelinesApi.md#GetNotificationSchema) | **Get** /pipeline/api/pipelines/notification | 
[**GetPipeline**](PipelinesApi.md#GetPipeline) | **Get** /pipeline/api/pipelines/{pipelineIdentifier} | Gets a Pipeline by identifier
[**GetPipelineList**](PipelinesApi.md#GetPipelineList) | **Post** /pipeline/api/pipelines/list | List of pipelines
[**GetPipelineSummary**](PipelinesApi.md#GetPipelineSummary) | **Get** /pipeline/api/pipelines/summary/{pipelineIdentifier} | Gets pipeline summary by pipeline identifier
[**GetPipelinesCount**](PipelinesApi.md#GetPipelinesCount) | **Post** /pipeline/api/landingDashboards/pipelinesCount | 
[**GetPmsStepNodes**](PipelinesApi.md#GetPmsStepNodes) | **Get** /pipeline/api/pipelines/dummy-pmsSteps-api | 
[**GetSteps**](PipelinesApi.md#GetSteps) | **Get** /pipeline/api/pipelines/steps | Gets all the Steps for given Category
[**GetStepsV2**](PipelinesApi.md#GetStepsV2) | **Post** /pipeline/api/pipelines/v2/steps | Gets all the Steps for given Category (V2 Version)
[**GetTemplateStepNode**](PipelinesApi.md#GetTemplateStepNode) | **Get** /pipeline/api/pipelines/dummy-templateStep-api | 
[**PostPipeline**](PipelinesApi.md#PostPipeline) | **Post** /pipeline/api/pipelines | Create a Pipeline
[**PostPipelineV2**](PipelinesApi.md#PostPipelineV2) | **Post** /pipeline/api/pipelines/v2 | Create a Pipeline API (V2 Version)
[**RefreshFFCache**](PipelinesApi.md#RefreshFFCache) | **Get** /pipeline/api/pipelines/ffCache/refresh | Refresh the feature flag cache
[**UpdatePipeline**](PipelinesApi.md#UpdatePipeline) | **Put** /pipeline/api/pipelines/{pipelineIdentifier} | Update a Pipeline by identifier
[**UpdatePipelineV2**](PipelinesApi.md#UpdatePipelineV2) | **Put** /pipeline/api/pipelines/v2/{pipelineIdentifier} | Updates a Pipeline by identifier (V2 Version)

# **CreateVariables**
> ResponseDtoVariableMergeServiceResponse CreateVariables(ctx, body, accountIdentifier, orgIdentifier, projectIdentifier)
Get all the Variables which can be used as expression in the Pipeline.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**string**](string.md)| Pipeline YAML | 
  **accountIdentifier** | **string**| Account Identifier for the entity. | 
  **orgIdentifier** | **string**| Organization Identifier for the entity. | 
  **projectIdentifier** | **string**| Project Identifier for the entity. | 

### Return type

[**ResponseDtoVariableMergeServiceResponse**](ResponseDTOVariableMergeServiceResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeletePipeline**
> ResponseDtoBoolean DeletePipeline(ctx, accountIdentifier, orgIdentifier, projectIdentifier, pipelineIdentifier, optional)
Deletes a Pipeline

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the entity. | 
  **orgIdentifier** | **string**| Organization Identifier for the entity. | 
  **projectIdentifier** | **string**| Project Identifier for the entity. | 
  **pipelineIdentifier** | **string**| Pipeline Identifier | 
 **optional** | ***PipelinesApiDeletePipelineOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PipelinesApiDeletePipelineOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **ifMatch** | **optional.String**| Version of entity to match | 
 **branch** | **optional.String**| Branch Name | 
 **repoIdentifier** | **optional.String**| Git Sync Config Id | 
 **rootFolder** | **optional.String**| Default Folder Path | 
 **filePath** | **optional.String**| File Path | 
 **commitMsg** | **optional.String**| Commit Message | 
 **lastObjectId** | **optional.String**| Last Object Id | 

### Return type

[**ResponseDtoBoolean**](ResponseDTOBoolean.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetExecutionNode**
> ResponseDtoExecutionNode GetExecutionNode(ctx, accountIdentifier, orgIdentifier, projectIdentifier, nodeExecutionId)
Get the Execution Node by Execution Id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the entity. | 
  **orgIdentifier** | **string**| Organization Identifier for the entity. | 
  **projectIdentifier** | **string**| Project Identifier for the entity. | 
  **nodeExecutionId** | **string**| Id for the corresponding Node Execution | 

### Return type

[**ResponseDtoExecutionNode**](ResponseDTOExecutionNode.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetExpandedPipelineJSON**
> ResponseDtoPipelineExpandedJson GetExpandedPipelineJSON(ctx, accountIdentifier, orgIdentifier, projectIdentifier, pipelineIdentifier, optional)
Gets Pipeline JSON with extra info for some fields

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the entity. | 
  **orgIdentifier** | **string**| Organization Identifier for the entity. | 
  **projectIdentifier** | **string**| Project Identifier for the entity. | 
  **pipelineIdentifier** | **string**| Pipeline Identifier | 
 **optional** | ***PipelinesApiGetExpandedPipelineJSONOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PipelinesApiGetExpandedPipelineJSONOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **branch** | **optional.String**| Branch Name | 
 **repoIdentifier** | **optional.String**| Git Sync Config Id | 
 **getDefaultFromOtherRepo** | **optional.Bool**| if true, return all the default entities | 

### Return type

[**ResponseDtoPipelineExpandedJson**](ResponseDTOPipelineExpandedJson.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNotificationSchema**
> GetNotificationSchema(ctx, )


### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPipeline**
> ResponseDtopmsPipelineResponse GetPipeline(ctx, accountIdentifier, orgIdentifier, projectIdentifier, pipelineIdentifier, optional)
Gets a Pipeline by identifier

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the entity. | 
  **orgIdentifier** | **string**| Organization Identifier for the entity. | 
  **projectIdentifier** | **string**| Project Identifier for the entity. | 
  **pipelineIdentifier** | **string**| Pipeline Identifier | 
 **optional** | ***PipelinesApiGetPipelineOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PipelinesApiGetPipelineOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **branch** | **optional.String**| Branch Name | 
 **repoIdentifier** | **optional.String**| Git Sync Config Id | 
 **getDefaultFromOtherRepo** | **optional.Bool**| if true, return all the default entities | 

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
List of pipelines

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the entity. | 
  **orgIdentifier** | **string**| Organization Identifier for the entity. | 
  **projectIdentifier** | **string**| Project Identifier for the entity. | 
 **optional** | ***PipelinesApiGetPipelineListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PipelinesApiGetPipelineListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | [**optional.Interface of FilterProperties**](FilterProperties.md)| This is the body for the filter properties for listing pipelines. | 
 **page** | **optional.**| The number of the page to fetch | [default to 0]
 **size** | **optional.**| The number of the elements to fetch | [default to 25]
 **sort** | [**optional.Interface of []string**](string.md)| Sort criteria for the elements. | 
 **searchTerm** | **optional.**| Search term to filter out pipelines based on pipeline name, identifier, tags. | 
 **module** | **optional.**|  | 
 **filterIdentifier** | **optional.**|  | 
 **branch** | **optional.**| Branch Name | 
 **repoIdentifier** | **optional.**| Git Sync Config Id | 
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
Gets pipeline summary by pipeline identifier

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the entity. | 
  **orgIdentifier** | **string**| Organization Identifier for the entity. | 
  **projectIdentifier** | **string**| Project Identifier for the entity. | 
  **pipelineIdentifier** | **string**| Pipeline Identifier | 
 **optional** | ***PipelinesApiGetPipelineSummaryOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PipelinesApiGetPipelineSummaryOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **branch** | **optional.String**| Branch Name | 
 **repoIdentifier** | **optional.String**| Git Sync Config Id | 
 **getDefaultFromOtherRepo** | **optional.Bool**| if true, return all the default entities | 

### Return type

[**ResponseDtopmsPipelineSummaryResponse**](ResponseDTOPMSPipelineSummaryResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPipelinesCount**
> ResponseDtoPipelinesCount GetPipelinesCount(ctx, body, accountIdentifier, startTime, endTime)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**LandingDashboardRequestPms**](LandingDashboardRequestPms.md)|  | 
  **accountIdentifier** | **string**|  | 
  **startTime** | **int64**|  | 
  **endTime** | **int64**|  | 

### Return type

[**ResponseDtoPipelinesCount**](ResponseDTOPipelinesCount.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPmsStepNodes**
> GetPmsStepNodes(ctx, )


### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSteps**
> ResponseDtoStepCategory GetSteps(ctx, category, module, optional)
Gets all the Steps for given Category

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **category** | **string**| Step Category for which you needs all its steps | 
  **module** | **string**| Module of the step to which it belongs | 
 **optional** | ***PipelinesApiGetStepsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PipelinesApiGetStepsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **accountId** | **optional.String**| Account Identifier for the entity. | 

### Return type

[**ResponseDtoStepCategory**](ResponseDTOStepCategory.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetStepsV2**
> ResponseDtoStepCategory GetStepsV2(ctx, body, accountId)
Gets all the Steps for given Category (V2 Version)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**StepPalleteFilterWrapper**](StepPalleteFilterWrapper.md)| Step Pallete Filter request body | 
  **accountId** | **string**| Account Identifier for the entity. | 

### Return type

[**ResponseDtoStepCategory**](ResponseDTOStepCategory.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTemplateStepNode**
> GetTemplateStepNode(ctx, )


### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostPipeline**
> ResponseDtoString PostPipeline(ctx, body, accountIdentifier, orgIdentifier, projectIdentifier, optional)
Create a Pipeline

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**string**](string.md)| Pipeline YAML | 
  **accountIdentifier** | **string**| Account Identifier for the entity. | 
  **orgIdentifier** | **string**| Organization Identifier for the entity. | 
  **projectIdentifier** | **string**| Project Identifier for the entity. | 
 **optional** | ***PipelinesApiPostPipelineOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PipelinesApiPostPipelineOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **branch** | **optional.**| Branch Name | 
 **repoIdentifier** | **optional.**| Git Sync Config Id | 
 **rootFolder** | **optional.**| Default Folder Path | 
 **filePath** | **optional.**| File Path | 
 **commitMsg** | **optional.**| File Path | 
 **isNewBranch** | **optional.**| Checks the new branch | [default to false]
 **baseBranch** | **optional.**| Default Branch | 

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
Create a Pipeline API (V2 Version)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**string**](string.md)| Pipeline YAML | 
  **accountIdentifier** | **string**| Account Identifier for the entity. | 
  **orgIdentifier** | **string**| Organization Identifier for the entity. | 
  **projectIdentifier** | **string**| Project Identifier for the entity. | 
 **optional** | ***PipelinesApiPostPipelineV2Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PipelinesApiPostPipelineV2Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **branch** | **optional.**| Branch Name | 
 **repoIdentifier** | **optional.**| Git Sync Config Id | 
 **rootFolder** | **optional.**| Default Folder Path | 
 **filePath** | **optional.**| File Path | 
 **commitMsg** | **optional.**| File Path | 
 **isNewBranch** | **optional.**| Checks the new branch | [default to false]
 **baseBranch** | **optional.**| Default Branch | 

### Return type

[**ResponseDtoPipelineSaveResponse**](ResponseDTOPipelineSaveResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RefreshFFCache**
> ResponseDtoBoolean RefreshFFCache(ctx, accountIdentifier)
Refresh the feature flag cache

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the entity. | 

### Return type

[**ResponseDtoBoolean**](ResponseDTOBoolean.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdatePipeline**
> ResponseDtoString UpdatePipeline(ctx, body, accountIdentifier, orgIdentifier, projectIdentifier, pipelineIdentifier, optional)
Update a Pipeline by identifier

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**string**](string.md)| Pipeline YAML to be updated | 
  **accountIdentifier** | **string**| Account Identifier for the entity. | 
  **orgIdentifier** | **string**| Organization Identifier for the entity. | 
  **projectIdentifier** | **string**| Project Identifier for the entity. | 
  **pipelineIdentifier** | **string**| Pipeline Identifier | 
 **optional** | ***PipelinesApiUpdatePipelineOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PipelinesApiUpdatePipelineOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **ifMatch** | **optional.**| Version of entity to match | 
 **branch** | **optional.**| Branch Name | 
 **repoIdentifier** | **optional.**| Git Sync Config Id | 
 **rootFolder** | **optional.**| Default Folder Path | 
 **filePath** | **optional.**| Default Folder Path | 
 **commitMsg** | **optional.**| Commit Message | 
 **lastObjectId** | **optional.**| Last Object Id | 
 **baseBranch** | **optional.**| Default Branch | 

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
Updates a Pipeline by identifier (V2 Version)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**string**](string.md)| Pipeline YAML to be updated | 
  **accountIdentifier** | **string**| Account Identifier for the entity. | 
  **orgIdentifier** | **string**| Organization Identifier for the entity. | 
  **projectIdentifier** | **string**| Project Identifier for the entity. | 
  **pipelineIdentifier** | **string**| Pipeline Identifier | 
 **optional** | ***PipelinesApiUpdatePipelineV2Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PipelinesApiUpdatePipelineV2Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **ifMatch** | **optional.**| Version of entity to match | 
 **branch** | **optional.**| Branch Name | 
 **repoIdentifier** | **optional.**| Git Sync Config Id | 
 **rootFolder** | **optional.**| Default Folder Path | 
 **filePath** | **optional.**| Default Folder Path | 
 **commitMsg** | **optional.**| Commit Message | 
 **lastObjectId** | **optional.**| Last Object Id | 
 **baseBranch** | **optional.**| Default Branch | 

### Return type

[**ResponseDtoPipelineSaveResponse**](ResponseDTOPipelineSaveResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

