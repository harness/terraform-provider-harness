# nextgen{{classname}}

All URIs are relative to *https://app.harness.io/gateway*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostPipelineExecuteWithInputSetList**](PipelineExecutionApi.md#PostPipelineExecuteWithInputSetList) | **Post** /pipeline/api/pipeline/execute/{identifier}/inputSetList | Execute a Pipeline with Input Set References
[**PostPipelineExecuteWithInputSetYaml**](PipelineExecutionApi.md#PostPipelineExecuteWithInputSetYaml) | **Post** /pipeline/api/pipeline/execute/{identifier} | Execute a Pipeline with Runtime Input YAML
[**PutHandleInterrupt**](PipelineExecutionApi.md#PutHandleInterrupt) | **Put** /pipeline/api/pipeline/execute/interrupt/{planExecutionId} | Execute an Interrupt

# **PostPipelineExecuteWithInputSetList**
> ResponseDtoPlanExecutionResponse PostPipelineExecuteWithInputSetList(ctx, body, accountIdentifier, orgIdentifier, projectIdentifier, moduleType, identifier, optional)
Execute a Pipeline with Input Set References

Execute a Pipeline with Input Set References

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MergeInputSetRequest**](MergeInputSetRequest.md)|  | 
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 
  **orgIdentifier** | **string**| Organization Identifier for the Entity. | 
  **projectIdentifier** | **string**| Project Identifier for the Entity. | 
  **moduleType** | **string**| Module type for the entity. If its from deployments,type will be CD , if its from build type will be CI | 
  **identifier** | **string**| Pipeline identifier for the entity. Identifier of the Pipeline to be executed | 
 **optional** | ***PipelineExecutionApiPostPipelineExecuteWithInputSetListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PipelineExecutionApiPostPipelineExecuteWithInputSetListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------






 **branch** | **optional.**| Name of the branch. | 
 **repoIdentifier** | **optional.**| Git Sync Config Id. | 
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
Execute a Pipeline with Runtime Input YAML

Execute a Pipeline with Runtime Input YAML

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 
  **orgIdentifier** | **string**| Organization Identifier for the Entity. | 
  **projectIdentifier** | **string**| Project Identifier for the Entity. | 
  **moduleType** | **string**| Module type for the entity. If its from deployments,type will be CD , if its from build type will be CI | 
  **identifier** | **string**| Pipeline identifier for the entity. Identifier of the Pipeline to be executed | 
 **optional** | ***PipelineExecutionApiPostPipelineExecuteWithInputSetYamlOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PipelineExecutionApiPostPipelineExecuteWithInputSetYamlOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **body** | [**optional.Interface of string**](string.md)| Enter Runtime Input YAML if the Pipeline contains Runtime Inputs. Please refer to https://ngdocs.harness.io/article/f6yobn7iq0 and https://ngdocs.harness.io/article/1eishcolt3 to see how to generate Runtime Input YAML for a Pipeline. | 
 **branch** | **optional.**| Name of the branch. | 
 **repoIdentifier** | **optional.**| Git Sync Config Id. | 
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
Execute an Interrupt

Executes an Interrupt on a Given Execution

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 
  **orgIdentifier** | **string**| Organization Identifier for the Entity. | 
  **projectIdentifier** | **string**| Project Identifier for the Entity. | 
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

