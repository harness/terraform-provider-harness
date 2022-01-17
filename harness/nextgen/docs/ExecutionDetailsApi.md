# {{classname}}

All URIs are relative to *https://app.harness.io/gateway*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetExecutionDetail**](ExecutionDetailsApi.md#GetExecutionDetail) | **Get** /pipeline/api/pipelines/execution/{planExecutionId} | Get the Pipeline Execution details for given PlanExecution Id
[**GetInputsetYaml**](ExecutionDetailsApi.md#GetInputsetYaml) | **Get** /pipeline/api/pipelines/execution/{planExecutionId}/inputset | Get the Input Set YAML used for given Plan Execution
[**GetInputsetYamlV2**](ExecutionDetailsApi.md#GetInputsetYamlV2) | **Get** /pipeline/api/pipelines/execution/{planExecutionId}/inputsetV2 | Get the Input Set YAML used for given Plan Execution
[**GetListOfExecutions**](ExecutionDetailsApi.md#GetListOfExecutions) | **Post** /pipeline/api/pipelines/execution/summary | Gets list of Executions of Pipelines for specific filters.

# **GetExecutionDetail**
> ResponseDtoPipelineExecutionDetail GetExecutionDetail(ctx, accountIdentifier, orgIdentifier, projectIdentifier, planExecutionId, optional)
Get the Pipeline Execution details for given PlanExecution Id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the entity. | 
  **orgIdentifier** | **string**| Organization Identifier for the entity. | 
  **projectIdentifier** | **string**| Project Identifier for the entity. | 
  **planExecutionId** | **string**| Plan Execution Id for which we want to get the Execution details | 
 **optional** | ***ExecutionDetailsApiGetExecutionDetailOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ExecutionDetailsApiGetExecutionDetailOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **stageNodeId** | **optional.String**| Stage Node Identifier to get execution stats. | 

### Return type

[**ResponseDtoPipelineExecutionDetail**](ResponseDTOPipelineExecutionDetail.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetInputsetYaml**
> string GetInputsetYaml(ctx, accountIdentifier, orgIdentifier, projectIdentifier, planExecutionId, optional)
Get the Input Set YAML used for given Plan Execution

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the entity. | 
  **orgIdentifier** | **string**| Organization Identifier for the entity. | 
  **projectIdentifier** | **string**| Project Identifier for the entity. | 
  **planExecutionId** | **string**| Plan Execution Id for which we want to get the Input Set YAML | 
 **optional** | ***ExecutionDetailsApiGetInputsetYamlOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ExecutionDetailsApiGetInputsetYamlOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **resolveExpressions** | **optional.Bool**|  | [default to false]

### Return type

**string**

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetInputsetYamlV2**
> ResponseDtoInputSetTemplateResponse GetInputsetYamlV2(ctx, accountIdentifier, orgIdentifier, projectIdentifier, planExecutionId, optional)
Get the Input Set YAML used for given Plan Execution

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the entity. | 
  **orgIdentifier** | **string**| Organization Identifier for the entity. | 
  **projectIdentifier** | **string**| Project Identifier for the entity. | 
  **planExecutionId** | **string**| Plan Execution Id for which we want to get the Input Set YAML | 
 **optional** | ***ExecutionDetailsApiGetInputsetYamlV2Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ExecutionDetailsApiGetInputsetYamlV2Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **resolveExpressions** | **optional.Bool**|  | [default to false]

### Return type

[**ResponseDtoInputSetTemplateResponse**](ResponseDTOInputSetTemplateResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetListOfExecutions**
> ResponseDtoPagePipelineExecutionSummary GetListOfExecutions(ctx, accountIdentifier, orgIdentifier, projectIdentifier, optional)
Gets list of Executions of Pipelines for specific filters.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the entity. | 
  **orgIdentifier** | **string**| Organization Identifier for the entity. | 
  **projectIdentifier** | **string**| Project Identifier for the entity. | 
 **optional** | ***ExecutionDetailsApiGetListOfExecutionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ExecutionDetailsApiGetListOfExecutionsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | [**optional.Interface of FilterProperties**](FilterProperties.md)|  | 
 **searchTerm** | **optional.**| Search term to filter out pipelines based on pipeline name, identifier, tags. | 
 **pipelineIdentifier** | **optional.**| Pipeline Identifier filter if exact pipelines needs to be filtered. | 
 **page** | **optional.**| The number of the page to fetch | [default to 0]
 **size** | **optional.**| The number of the elements to fetch | [default to 10]
 **sort** | [**optional.Interface of []string**](string.md)| Sort criteria for the elements. | 
 **filterIdentifier** | **optional.**|  | 
 **module** | **optional.**|  | 
 **status** | [**optional.Interface of []string**](string.md)|  | 
 **myDeployments** | **optional.**|  | 
 **branch** | **optional.**| Branch Name | 
 **repoIdentifier** | **optional.**| Git Sync Config Id | 
 **getDefaultFromOtherRepo** | **optional.**| if true, return all the default entities | 

### Return type

[**ResponseDtoPagePipelineExecutionSummary**](ResponseDTOPagePipelineExecutionSummary.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

