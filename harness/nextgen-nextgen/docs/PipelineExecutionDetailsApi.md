# nextgen{{classname}}

All URIs are relative to *https://app.harness.io/gateway*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetExecutionDetail**](PipelineExecutionDetailsApi.md#GetExecutionDetail) | **Get** /pipeline/api/pipelines/execution/{planExecutionId} | Fetch Pipeline Execution Details
[**GetExecutionDetailV2**](PipelineExecutionDetailsApi.md#GetExecutionDetailV2) | **Get** /pipeline/api/pipelines/execution/v2/{planExecutionId} | Fetch Pipeline Execution Details
[**GetListOfExecutions**](PipelineExecutionDetailsApi.md#GetListOfExecutions) | **Post** /pipeline/api/pipelines/execution/summary | List Executions

# **GetExecutionDetail**
> ResponseDtoPipelineExecutionDetail GetExecutionDetail(ctx, accountIdentifier, orgIdentifier, projectIdentifier, planExecutionId, optional)
Fetch Pipeline Execution Details

Returns the Pipeline Execution Details for a Given PlanExecution ID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 
  **orgIdentifier** | **string**| Organization Identifier for the Entity. | 
  **projectIdentifier** | **string**| Project Identifier for the Entity. | 
  **planExecutionId** | **string**| Plan Execution Id for which we want to get the Execution details | 
 **optional** | ***PipelineExecutionDetailsApiGetExecutionDetailOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PipelineExecutionDetailsApiGetExecutionDetailOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **stageNodeId** | **optional.String**| Stage Node Identifier for which Stage Graph needs to be Rendered | 
 **stageNodeExecutionId** | **optional.String**| Stage Node Execution ID for which Stage Graph needs to be Rendered. (Needed only when there are Multiple Runs for a Given Stage. It can be Extracted from LayoutNodeMap Field) | 

### Return type

[**ResponseDtoPipelineExecutionDetail**](ResponseDTOPipelineExecutionDetail.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetExecutionDetailV2**
> ResponseDtoPipelineExecutionDetail GetExecutionDetailV2(ctx, accountIdentifier, orgIdentifier, projectIdentifier, planExecutionId, optional)
Fetch Pipeline Execution Details

Returns the Pipeline Execution Details for a Given PlanExecution ID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 
  **orgIdentifier** | **string**| Organization Identifier for the Entity. | 
  **projectIdentifier** | **string**| Project Identifier for the Entity. | 
  **planExecutionId** | **string**| Plan Execution Id for which we want to get the Execution details | 
 **optional** | ***PipelineExecutionDetailsApiGetExecutionDetailV2Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PipelineExecutionDetailsApiGetExecutionDetailV2Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **stageNodeId** | **optional.String**| Stage Node Identifier for which Stage Graph needs to be Rendered | 
 **stageNodeExecutionId** | **optional.String**| Stage Node Execution ID for which Stage Graph needs to be Rendered. (Needed only when there are Multiple Runs for a Given Stage. It can be Extracted from LayoutNodeMap Field) | 
 **renderFullBottomGraph** | **optional.Bool**| Generate Graph for all the Stages including Steps in each Stage | 

### Return type

[**ResponseDtoPipelineExecutionDetail**](ResponseDTOPipelineExecutionDetail.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetListOfExecutions**
> ResponseDtoPagePipelineExecutionSummary GetListOfExecutions(ctx, accountIdentifier, orgIdentifier, projectIdentifier, optional)
List Executions

Returns a List of Pipeline Executions with Specific Filters

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 
  **orgIdentifier** | **string**| Organization Identifier for the Entity. | 
  **projectIdentifier** | **string**| Project Identifier for the Entity. | 
 **optional** | ***PipelineExecutionDetailsApiGetListOfExecutionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PipelineExecutionDetailsApiGetListOfExecutionsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | [**optional.Interface of FilterProperties**](FilterProperties.md)|  | 
 **searchTerm** | **optional.**| Search term to filter out pipelines based on pipeline name, identifier, tags. | 
 **pipelineIdentifier** | **optional.**| Pipeline Identifier filter if exact pipelines needs to be filtered. | 
 **page** | **optional.**| Page Index of the results to fetch.Default Value: 0 | [default to 0]
 **size** | **optional.**| Results per page | [default to 10]
 **sort** | [**optional.Interface of []string**](string.md)| Sort criteria for the elements. | 
 **filterIdentifier** | **optional.**|  | 
 **module** | **optional.**|  | 
 **status** | [**optional.Interface of []string**](string.md)|  | 
 **myDeployments** | **optional.**|  | 
 **branch** | **optional.**| Name of the branch. | 
 **repoIdentifier** | **optional.**| Git Sync Config Id. | 
 **getDefaultFromOtherRepo** | **optional.**| if true, return all the default entities | 

### Return type

[**ResponseDtoPagePipelineExecutionSummary**](ResponseDTOPagePipelineExecutionSummary.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

