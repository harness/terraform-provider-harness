# {{classname}}

All URIs are relative to *https://app.harness.io/gateway*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPreFlightCheckResponse**](ExecuteApi.md#GetPreFlightCheckResponse) | **Get** /pipeline/api/pipeline/execute/getPreflightCheckResponse | Get Preflight Checks Response for a Preflight Id
[**GetRetryStages**](ExecuteApi.md#GetRetryStages) | **Get** /pipeline/api/pipeline/execute/{planExecutionId}/retryStages | Get retry stages for failed pipeline
[**GetStagesExecutionList**](ExecuteApi.md#GetStagesExecutionList) | **Get** /pipeline/api/pipeline/execute/stagesExecutionList | Get list of Stages to select for Stage executions
[**HandleManualInterventionInterrupt**](ExecuteApi.md#HandleManualInterventionInterrupt) | **Put** /pipeline/api/pipeline/execute/manualIntervention/interrupt/{planExecutionId}/{nodeExecutionId} | Handles Ignore,Abort,MarkAsSuccess,Retry on post manual intervention for a given execution with the given planExecutionId
[**HandleStageInterrupt**](ExecuteApi.md#HandleStageInterrupt) | **Put** /pipeline/api/pipeline/execute/interrupt/{planExecutionId}/{nodeExecutionId} | Handles the interrupt for a given stage in a pipeline
[**LatestExecutionId**](ExecuteApi.md#LatestExecutionId) | **Get** /pipeline/api/pipeline/execute/latestExecutionId/{planExecutionId} | Latest ExecutionId from Retry Executions
[**PostExecuteStages**](ExecuteApi.md#PostExecuteStages) | **Post** /pipeline/api/pipeline/execute/{identifier}/stages | Execute given Stages of a Pipeline
[**PostPipelineExecuteWithInputSetList**](ExecuteApi.md#PostPipelineExecuteWithInputSetList) | **Post** /pipeline/api/pipeline/execute/{identifier}/inputSetList | Execute a pipeline with input set references list
[**PostPipelineExecuteWithInputSetYaml**](ExecuteApi.md#PostPipelineExecuteWithInputSetYaml) | **Post** /pipeline/api/pipeline/execute/{identifier} | Execute a pipeline with inputSet pipeline yaml
[**PostPipelineExecuteWithInputSetYamlv2**](ExecuteApi.md#PostPipelineExecuteWithInputSetYamlv2) | **Post** /pipeline/api/pipeline/execute/{identifier}/v2 | Execute a pipeline with inputSet pipeline yaml V2
[**PostReExecuteStages**](ExecuteApi.md#PostReExecuteStages) | **Post** /pipeline/api/pipeline/execute/rerun/{originalExecutionId}/{identifier}/stages | Re-run given Stages of a Pipeline
[**PutHandleInterrupt**](ExecuteApi.md#PutHandleInterrupt) | **Put** /pipeline/api/pipeline/execute/interrupt/{planExecutionId} | Execute an Interrupt on an execution
[**RePostPipelineExecuteWithInputSetYaml**](ExecuteApi.md#RePostPipelineExecuteWithInputSetYaml) | **Post** /pipeline/api/pipeline/execute/rerun/{originalExecutionId}/{identifier} | Re Execute a pipeline with inputSet pipeline yaml
[**RePostPipelineExecuteWithInputSetYamlV2**](ExecuteApi.md#RePostPipelineExecuteWithInputSetYamlV2) | **Post** /pipeline/api/pipeline/execute/rerun/v2/{originalExecutionId}/{identifier} | Re Execute a pipeline with InputSet Pipeline YAML Version 2
[**RerunPipelineWithInputSetIdentifierList**](ExecuteApi.md#RerunPipelineWithInputSetIdentifierList) | **Post** /pipeline/api/pipeline/execute/rerun/{originalExecutionId}/{identifier}/inputSetList | Rerun a pipeline with given inputSet identifiers
[**RetryHistory**](ExecuteApi.md#RetryHistory) | **Get** /pipeline/api/pipeline/execute/retryHistory/{planExecutionId} | Retry History for a given execution
[**RetryPipeline**](ExecuteApi.md#RetryPipeline) | **Post** /pipeline/api/pipeline/execute/retry/{identifier} | Retry a executed pipeline with inputSet pipeline yaml
[**RunASchemaMigration**](ExecuteApi.md#RunASchemaMigration) | **Get** /pipeline/api/pipeline/execute/internal/runSchema | 
[**StartPreFlightCheck**](ExecuteApi.md#StartPreFlightCheck) | **Post** /pipeline/api/pipeline/execute/preflightCheck | Start Preflight Checks for a Pipeline

# **GetPreFlightCheckResponse**
> ResponseDtoPreFlightDto GetPreFlightCheckResponse(ctx, accountIdentifier, orgIdentifier, projectIdentifier, preflightCheckId, optional)
Get Preflight Checks Response for a Preflight Id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the entity. | 
  **orgIdentifier** | **string**| Organization Identifier for the entity. | 
  **projectIdentifier** | **string**| Project Identifier for the entity. | 
  **preflightCheckId** | **string**| Preflight Id from the start Preflight Checks API | 
 **optional** | ***ExecuteApiGetPreFlightCheckResponseOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ExecuteApiGetPreFlightCheckResponseOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **body** | [**optional.Interface of string**](string.md)|  | 

### Return type

[**ResponseDtoPreFlightDto**](ResponseDTOPreFlightDTO.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRetryStages**
> ResponseDtoRetryInfo GetRetryStages(ctx, accountIdentifier, orgIdentifier, projectIdentifier, planExecutionId, optional)
Get retry stages for failed pipeline

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the entity. | 
  **orgIdentifier** | **string**| Organization Identifier for the entity. | 
  **projectIdentifier** | **string**| Project Identifier for the entity. | 
  **planExecutionId** | **string**| planExecutionId of the execution we want to retry | 
 **optional** | ***ExecuteApiGetRetryStagesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ExecuteApiGetRetryStagesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **pipelineIdentifier** | **optional.String**| Pipeline Identifier | 
 **branch** | **optional.String**| Branch Name | 
 **repoIdentifier** | **optional.String**| Git Sync Config Id | 
 **getDefaultFromOtherRepo** | **optional.Bool**| if true, return all the default entities | 

### Return type

[**ResponseDtoRetryInfo**](ResponseDTORetryInfo.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetStagesExecutionList**
> ResponseDtoListStageExecutionResponse GetStagesExecutionList(ctx, accountIdentifier, orgIdentifier, projectIdentifier, pipelineIdentifier, optional)
Get list of Stages to select for Stage executions

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the entity. | 
  **orgIdentifier** | **string**| Organization Identifier for the entity. | 
  **projectIdentifier** | **string**| Project Identifier for the entity. | 
  **pipelineIdentifier** | **string**| Pipeline Identifier | 
 **optional** | ***ExecuteApiGetStagesExecutionListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ExecuteApiGetStagesExecutionListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **branch** | **optional.String**| Branch Name | 
 **repoIdentifier** | **optional.String**| Git Sync Config Id | 
 **getDefaultFromOtherRepo** | **optional.Bool**| if true, return all the default entities | 

### Return type

[**ResponseDtoListStageExecutionResponse**](ResponseDTOListStageExecutionResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **HandleManualInterventionInterrupt**
> ResponseDtoInterruptResponse HandleManualInterventionInterrupt(ctx, accountIdentifier, orgIdentifier, projectIdentifier, interruptType, planExecutionId, nodeExecutionId)
Handles Ignore,Abort,MarkAsSuccess,Retry on post manual intervention for a given execution with the given planExecutionId

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the entity. | 
  **orgIdentifier** | **string**| Organization Identifier for the entity. | 
  **projectIdentifier** | **string**| Project Identifier for the entity. | 
  **interruptType** | **string**| The Interrupt type needed to be applied to the execution. Choose a value from the enum list. | 
  **planExecutionId** | **string**| The Pipeline Execution Id on which the Interrupt needs to be applied. | 
  **nodeExecutionId** | **string**| The runtime Id of the step/stage on which the Interrupt needs to be applied. | 

### Return type

[**ResponseDtoInterruptResponse**](ResponseDTOInterruptResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **HandleStageInterrupt**
> ResponseDtoInterruptResponse HandleStageInterrupt(ctx, accountIdentifier, orgIdentifier, projectIdentifier, interruptType, planExecutionId, nodeExecutionId)
Handles the interrupt for a given stage in a pipeline

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the entity. | 
  **orgIdentifier** | **string**| Organization Identifier for the entity. | 
  **projectIdentifier** | **string**| Project Identifier for the entity. | 
  **interruptType** | **string**| The Interrupt type needed to be applied to the execution. Choose a value from the enum list. | 
  **planExecutionId** | **string**| The Pipeline Execution Id on which the Interrupt needs to be applied. | 
  **nodeExecutionId** | **string**| The runtime Id of the step/stage on which the Interrupt needs to be applied. | 

### Return type

[**ResponseDtoInterruptResponse**](ResponseDTOInterruptResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LatestExecutionId**
> ResponseDtoRetryLatestExecutionResponse LatestExecutionId(ctx, accountIdentifier, orgIdentifier, projectIdentifier, pipelineIdentifier, planExecutionId)
Latest ExecutionId from Retry Executions

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the entity. | 
  **orgIdentifier** | **string**| Organization Identifier for the entity. | 
  **projectIdentifier** | **string**| Project Identifier for the entity. | 
  **pipelineIdentifier** | **string**| Pipeline Identifier | 
  **planExecutionId** | **string**| planExecutionId of the execution of whose we need to find the latest execution planExecutionId | 

### Return type

[**ResponseDtoRetryLatestExecutionResponse**](ResponseDTORetryLatestExecutionResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostExecuteStages**
> ResponseDtoPlanExecutionResponse PostExecuteStages(ctx, accountIdentifier, orgIdentifier, projectIdentifier, moduleType, identifier, optional)
Execute given Stages of a Pipeline

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the entity. | 
  **orgIdentifier** | **string**| Organization Identifier for the entity. | 
  **projectIdentifier** | **string**| Project Identifier for the entity. | 
  **moduleType** | **string**| Module type for the entity. If its from deployments,type will be CD , if its from build type will be CI | 
  **identifier** | **string**| Pipeline Identifier | 
 **optional** | ***ExecuteApiPostExecuteStagesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ExecuteApiPostExecuteStagesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **body** | [**optional.Interface of RunStageRequest**](RunStageRequest.md)|  | 
 **branch** | **optional.**| Branch Name | 
 **repoIdentifier** | **optional.**| Git Sync Config Id | 
 **getDefaultFromOtherRepo** | **optional.**| if true, return all the default entities | 
 **useFQNIfError** | **optional.**|  | [default to false]

### Return type

[**ResponseDtoPlanExecutionResponse**](ResponseDTOPlanExecutionResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostPipelineExecuteWithInputSetList**
> ResponseDtoPlanExecutionResponse PostPipelineExecuteWithInputSetList(ctx, body, accountIdentifier, orgIdentifier, projectIdentifier, moduleType, identifier, optional)
Execute a pipeline with input set references list

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MergeInputSetRequest**](MergeInputSetRequest.md)|  | 
  **accountIdentifier** | **string**| Account Identifier for the entity. | 
  **orgIdentifier** | **string**| Organization Identifier for the entity. | 
  **projectIdentifier** | **string**| Project Identifier for the entity. | 
  **moduleType** | **string**| Module type for the entity. If its from deployments,type will be CD , if its from build type will be CI | 
  **identifier** | **string**| Pipeline identifier for the entity. Identifier of the Pipeline to be executed | 
 **optional** | ***ExecuteApiPostPipelineExecuteWithInputSetListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ExecuteApiPostPipelineExecuteWithInputSetListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------






 **branch** | **optional.**| Branch Name | 
 **repoIdentifier** | **optional.**| Git Sync Config Id | 
 **getDefaultFromOtherRepo** | **optional.**| if true, return all the default entities | 
 **useFQNIfError** | **optional.**|  | [default to false]

### Return type

[**ResponseDtoPlanExecutionResponse**](ResponseDTOPlanExecutionResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostPipelineExecuteWithInputSetYaml**
> ResponseDtoPlanExecutionResponse PostPipelineExecuteWithInputSetYaml(ctx, accountIdentifier, orgIdentifier, projectIdentifier, moduleType, identifier, optional)
Execute a pipeline with inputSet pipeline yaml

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the entity. | 
  **orgIdentifier** | **string**| Organization Identifier for the entity. | 
  **projectIdentifier** | **string**| Project Identifier for the entity. | 
  **moduleType** | **string**| Module type for the entity. If its from deployments,type will be CD , if its from build type will be CI | 
  **identifier** | **string**| Pipeline identifier for the entity. Identifier of the Pipeline to be executed | 
 **optional** | ***ExecuteApiPostPipelineExecuteWithInputSetYamlOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ExecuteApiPostPipelineExecuteWithInputSetYamlOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **body** | [**optional.Interface of string**](string.md)| InputSet YAML if the pipeline contains runtime inputs. This will be empty by default if pipeline does not contains runtime inputs | 
 **branch** | **optional.**| Branch Name | 
 **repoIdentifier** | **optional.**| Git Sync Config Id | 
 **getDefaultFromOtherRepo** | **optional.**| if true, return all the default entities | 
 **useFQNIfError** | **optional.**|  | [default to false]

### Return type

[**ResponseDtoPlanExecutionResponse**](ResponseDTOPlanExecutionResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostPipelineExecuteWithInputSetYamlv2**
> ResponseDtoPlanExecutionResponse PostPipelineExecuteWithInputSetYamlv2(ctx, accountIdentifier, orgIdentifier, projectIdentifier, moduleType, identifier, optional)
Execute a pipeline with inputSet pipeline yaml V2

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the entity. | 
  **orgIdentifier** | **string**| Organization Identifier for the entity. | 
  **projectIdentifier** | **string**| Project Identifier for the entity. | 
  **moduleType** | **string**| Module type for the entity. If its from deployments,type will be CD , if its from build type will be CI | 
  **identifier** | **string**| Pipeline identifier for the entity. Identifier of the Pipeline to be executed | 
 **optional** | ***ExecuteApiPostPipelineExecuteWithInputSetYamlv2Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ExecuteApiPostPipelineExecuteWithInputSetYamlv2Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **body** | [**optional.Interface of string**](string.md)| InputSet YAML if the pipeline contains runtime inputs. This will be empty by default if pipeline does not contains runtime inputs | 
 **branch** | **optional.**| Branch Name | 
 **repoIdentifier** | **optional.**| Git Sync Config Id | 
 **getDefaultFromOtherRepo** | **optional.**| if true, return all the default entities | 
 **useFQNIfError** | **optional.**|  | [default to false]

### Return type

[**ResponseDtoPlanExecutionResponse**](ResponseDTOPlanExecutionResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostReExecuteStages**
> ResponseDtoPlanExecutionResponse PostReExecuteStages(ctx, accountIdentifier, orgIdentifier, projectIdentifier, moduleType, identifier, originalExecutionId, optional)
Re-run given Stages of a Pipeline

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the entity. | 
  **orgIdentifier** | **string**| Organization Identifier for the entity. | 
  **projectIdentifier** | **string**| Project Identifier for the entity. | 
  **moduleType** | **string**| Module type for the entity. If its from deployments,type will be CD , if its from build type will be CI | 
  **identifier** | **string**| Pipeline Identifier | 
  **originalExecutionId** | **string**| This param contains the previous execution execution id. This is basically when we are rerunning a Pipeline. | 
 **optional** | ***ExecuteApiPostReExecuteStagesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ExecuteApiPostReExecuteStagesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------






 **body** | [**optional.Interface of RunStageRequest**](RunStageRequest.md)|  | 
 **branch** | **optional.**| Branch Name | 
 **repoIdentifier** | **optional.**| Git Sync Config Id | 
 **getDefaultFromOtherRepo** | **optional.**| if true, return all the default entities | 
 **useFQNIfError** | **optional.**|  | [default to false]

### Return type

[**ResponseDtoPlanExecutionResponse**](ResponseDTOPlanExecutionResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutHandleInterrupt**
> ResponseDtoInterruptResponse PutHandleInterrupt(ctx, accountIdentifier, orgIdentifier, projectIdentifier, interruptType, planExecutionId)
Execute an Interrupt on an execution

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the entity. | 
  **orgIdentifier** | **string**| Organization Identifier for the entity. | 
  **projectIdentifier** | **string**| Project Identifier for the entity. | 
  **interruptType** | **string**| The Interrupt type needed to be applied to the execution. Choose a value from the enum list. | 
  **planExecutionId** | **string**| The Pipeline Execution Id on which the Interrupt needs to be applied. | 

### Return type

[**ResponseDtoInterruptResponse**](ResponseDTOInterruptResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RePostPipelineExecuteWithInputSetYaml**
> ResponseDtoPlanExecutionResponse RePostPipelineExecuteWithInputSetYaml(ctx, accountIdentifier, orgIdentifier, projectIdentifier, moduleType, originalExecutionId, identifier, optional)
Re Execute a pipeline with inputSet pipeline yaml

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the entity. | 
  **orgIdentifier** | **string**| Organization Identifier for the entity. | 
  **projectIdentifier** | **string**| Project Identifier for the entity. | 
  **moduleType** | **string**| Module type for the entity. If its from deployments,type will be CD , if its from build type will be CI | 
  **originalExecutionId** | **string**| This param contains the previous execution execution id. This is basically when we are rerunning a Pipeline. | 
  **identifier** | **string**| Pipeline identifier for the entity. Identifier of the Pipeline to be executed | 
 **optional** | ***ExecuteApiRePostPipelineExecuteWithInputSetYamlOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ExecuteApiRePostPipelineExecuteWithInputSetYamlOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------






 **body** | [**optional.Interface of string**](string.md)| InputSet YAML if the pipeline contains runtime inputs. This will be empty by default if pipeline does not contains runtime inputs | 
 **branch** | **optional.**| Branch Name | 
 **repoIdentifier** | **optional.**| Git Sync Config Id | 
 **getDefaultFromOtherRepo** | **optional.**| if true, return all the default entities | 
 **useFQNIfError** | **optional.**|  | [default to false]

### Return type

[**ResponseDtoPlanExecutionResponse**](ResponseDTOPlanExecutionResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RePostPipelineExecuteWithInputSetYamlV2**
> ResponseDtoPlanExecutionResponse RePostPipelineExecuteWithInputSetYamlV2(ctx, accountIdentifier, orgIdentifier, projectIdentifier, moduleType, originalExecutionId, identifier, optional)
Re Execute a pipeline with InputSet Pipeline YAML Version 2

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the entity. | 
  **orgIdentifier** | **string**| Organization Identifier for the entity. | 
  **projectIdentifier** | **string**| Project Identifier for the entity. | 
  **moduleType** | **string**| Module type for the entity. If its from deployments,type will be CD , if its from build type will be CI | 
  **originalExecutionId** | **string**| This param contains the previous execution execution id. This is basically when we are rerunning a Pipeline. | 
  **identifier** | **string**| Pipeline identifier for the entity. Identifier of the Pipeline to be executed | 
 **optional** | ***ExecuteApiRePostPipelineExecuteWithInputSetYamlV2Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ExecuteApiRePostPipelineExecuteWithInputSetYamlV2Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------






 **body** | [**optional.Interface of string**](string.md)| InputSet YAML if the pipeline contains runtime inputs. This will be empty by default if pipeline does not contains runtime inputs | 
 **branch** | **optional.**| Branch Name | 
 **repoIdentifier** | **optional.**| Git Sync Config Id | 
 **getDefaultFromOtherRepo** | **optional.**| if true, return all the default entities | 
 **useFQNIfError** | **optional.**|  | [default to false]

### Return type

[**ResponseDtoPlanExecutionResponse**](ResponseDTOPlanExecutionResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RerunPipelineWithInputSetIdentifierList**
> ResponseDtoPlanExecutionResponse RerunPipelineWithInputSetIdentifierList(ctx, body, accountIdentifier, orgIdentifier, projectIdentifier, moduleType, originalExecutionId, identifier, useFQNIfError, optional)
Rerun a pipeline with given inputSet identifiers

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MergeInputSetRequest**](MergeInputSetRequest.md)| InputSet reference details | 
  **accountIdentifier** | **string**| Account Identifier for the entity. | 
  **orgIdentifier** | **string**| Organization Identifier for the entity. | 
  **projectIdentifier** | **string**| Project Identifier for the entity. | 
  **moduleType** | **string**| The module from which execution was triggered. | 
  **originalExecutionId** | **string**| Id of the execution from which we are running | 
  **identifier** | **string**| Pipeline Identifier | 
  **useFQNIfError** | **bool**| Use FQN in error response | [default to false]
 **optional** | ***ExecuteApiRerunPipelineWithInputSetIdentifierListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ExecuteApiRerunPipelineWithInputSetIdentifierListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------








 **branch** | **optional.**| Branch Name | 
 **repoIdentifier** | **optional.**| Git Sync Config Id | 
 **getDefaultFromOtherRepo** | **optional.**| if true, return all the default entities | 

### Return type

[**ResponseDtoPlanExecutionResponse**](ResponseDTOPlanExecutionResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RetryHistory**
> ResponseDtoRetryHistoryResponse RetryHistory(ctx, accountIdentifier, orgIdentifier, projectIdentifier, pipelineIdentifier, planExecutionId)
Retry History for a given execution

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the entity. | 
  **orgIdentifier** | **string**| Organization Identifier for the entity. | 
  **projectIdentifier** | **string**| Project Identifier for the entity. | 
  **pipelineIdentifier** | **string**| Pipeline Identifier | 
  **planExecutionId** | **string**| planExecutionId of the execution of whose we need to find the retry history | 

### Return type

[**ResponseDtoRetryHistoryResponse**](ResponseDTORetryHistoryResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RetryPipeline**
> ResponseDtoPlanExecutionResponse RetryPipeline(ctx, accountIdentifier, orgIdentifier, projectIdentifier, moduleType, planExecutionId, retryStages, identifier, optional)
Retry a executed pipeline with inputSet pipeline yaml

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the entity. | 
  **orgIdentifier** | **string**| Organization Identifier for the entity. | 
  **projectIdentifier** | **string**| Project Identifier for the entity. | 
  **moduleType** | **string**| Module type for the entity. If its from deployments,type will be CD , if its from build type will be CI | 
  **planExecutionId** | **string**| This param contains the previous execution execution id. This is basically when we are rerunning a Pipeline. | 
  **retryStages** | [**[]string**](string.md)| This param contains the identifier of stages from where to resume. It will be a list if we want to retry from parallel group  | 
  **identifier** | **string**| Pipeline Identifier | 
 **optional** | ***ExecuteApiRetryPipelineOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ExecuteApiRetryPipelineOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------







 **body** | [**optional.Interface of string**](string.md)|  | 
 **runAllStages** | **optional.**| This param provides an option to run only the failed stages when Pipeline fails at parallel group. By default, it will run all the stages in the failed parallel group. | [default to true]

### Return type

[**ResponseDtoPlanExecutionResponse**](ResponseDTOPlanExecutionResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RunASchemaMigration**
> RunASchemaMigration(ctx, )


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

# **StartPreFlightCheck**
> ResponseDtoString StartPreFlightCheck(ctx, accountIdentifier, orgIdentifier, projectIdentifier, optional)
Start Preflight Checks for a Pipeline

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the entity. | 
  **orgIdentifier** | **string**| Organization Identifier for the entity. | 
  **projectIdentifier** | **string**| Project Identifier for the entity. | 
 **optional** | ***ExecuteApiStartPreFlightCheckOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ExecuteApiStartPreFlightCheckOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | [**optional.Interface of string**](string.md)| Runtime Input YAML to be sent for Pipeline execution | 
 **pipelineIdentifier** | **optional.**| Pipeline Identifier | 
 **branch** | **optional.**| Branch Name | 
 **repoIdentifier** | **optional.**| Git Sync Config Id | 
 **getDefaultFromOtherRepo** | **optional.**| if true, return all the default entities | 

### Return type

[**ResponseDtoString**](ResponseDTOString.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

