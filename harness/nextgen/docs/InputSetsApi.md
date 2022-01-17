# {{classname}}

All URIs are relative to *https://app.harness.io/gateway*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteInputSet**](InputSetsApi.md#DeleteInputSet) | **Delete** /pipeline/api/inputSets/{inputSetIdentifier} | Delete the Input Set by Identifier
[**GetInputSet**](InputSetsApi.md#GetInputSet) | **Get** /pipeline/api/inputSets/{inputSetIdentifier} | Gets Input Set for a given identifier. Throws error if no Input Set exists for the given identifier.
[**GetOverlayInputSet**](InputSetsApi.md#GetOverlayInputSet) | **Get** /pipeline/api/inputSets/overlay/{inputSetIdentifier} | Gets an Overlay Input Set by identifier
[**ListInputSet**](InputSetsApi.md#ListInputSet) | **Get** /pipeline/api/inputSets | List all Input Sets for a pipeline
[**MergeInputSets**](InputSetsApi.md#MergeInputSets) | **Post** /pipeline/api/inputSets/merge | Merge given Input Sets into a single Runtime Input YAML
[**MergeRuntimeInputIntoPipeline**](InputSetsApi.md#MergeRuntimeInputIntoPipeline) | **Post** /pipeline/api/inputSets/mergeWithTemplateYaml | Merge given Runtime Input YAML into the Pipeline
[**PostInputSet**](InputSetsApi.md#PostInputSet) | **Post** /pipeline/api/inputSets | Create an Input Set for a Pipeline
[**PostOverlayInputSet**](InputSetsApi.md#PostOverlayInputSet) | **Post** /pipeline/api/inputSets/overlay | Create an Overlay Input Set for a pipeline
[**PutInputSet**](InputSetsApi.md#PutInputSet) | **Put** /pipeline/api/inputSets/{inputSetIdentifier} | Update Input Set for Pipeline
[**PutOverlayInputSet**](InputSetsApi.md#PutOverlayInputSet) | **Put** /pipeline/api/inputSets/overlay/{inputSetIdentifier} | Update an Overlay Input Set for a pipeline
[**RuntimeInputTemplate**](InputSetsApi.md#RuntimeInputTemplate) | **Post** /pipeline/api/inputSets/template | Fetch Runtime Input Template for a Pipeline

# **DeleteInputSet**
> ResponseDtoBoolean DeleteInputSet(ctx, inputSetIdentifier, accountIdentifier, orgIdentifier, projectIdentifier, pipelineIdentifier, optional)
Delete the Input Set by Identifier

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **inputSetIdentifier** | **string**| This is the Id of the Input Set that needs to be deleted. An error is thrown if no Input Set exists for this Id. | 
  **accountIdentifier** | **string**| Account Identifier for the entity. | 
  **orgIdentifier** | **string**| Organization Identifier for the entity. | 
  **projectIdentifier** | **string**| Project Identifier for the entity. | 
  **pipelineIdentifier** | **string**| Pipeline identifier for the Input Set. Input Set will be deleted for the Pipeline corresponding to this Identifier | 
 **optional** | ***InputSetsApiDeleteInputSetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InputSetsApiDeleteInputSetOpts struct
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

# **GetInputSet**
> ResponseDtoInputSetResponse GetInputSet(ctx, inputSetIdentifier, accountIdentifier, orgIdentifier, projectIdentifier, pipelineIdentifier, optional)
Gets Input Set for a given identifier. Throws error if no Input Set exists for the given identifier.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **inputSetIdentifier** | **string**| Identifier of the Input Set | 
  **accountIdentifier** | **string**| Account Identifier for the entity. | 
  **orgIdentifier** | **string**| Organization Identifier for the entity. | 
  **projectIdentifier** | **string**| Project Identifier for the entity. | 
  **pipelineIdentifier** | **string**| Pipeline ID for the Input Set. The Input Set will work only for the Pipeline corresponding to this identifier. | 
 **optional** | ***InputSetsApiGetInputSetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InputSetsApiGetInputSetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **branch** | **optional.String**| Branch Name | 
 **repoIdentifier** | **optional.String**| Git Sync Config Id | 
 **getDefaultFromOtherRepo** | **optional.Bool**| if true, return all the default entities | 

### Return type

[**ResponseDtoInputSetResponse**](ResponseDTOInputSetResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOverlayInputSet**
> ResponseDtoOverlayInputSetResponse GetOverlayInputSet(ctx, inputSetIdentifier, accountIdentifier, orgIdentifier, projectIdentifier, pipelineIdentifier, optional)
Gets an Overlay Input Set by identifier

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **inputSetIdentifier** | **string**| Identifier of the Overlay Input Set | 
  **accountIdentifier** | **string**| Account Identifier for the entity. | 
  **orgIdentifier** | **string**| Organization Identifier for the entity. | 
  **projectIdentifier** | **string**| Project Identifier for the entity. | 
  **pipelineIdentifier** | **string**| This is the Pipeline Id specific to the Overlay Input Set. Overlay Input Set corresponding to the Pipeline with this Id would be fetched | 
 **optional** | ***InputSetsApiGetOverlayInputSetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InputSetsApiGetOverlayInputSetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **branch** | **optional.String**| Branch Name | 
 **repoIdentifier** | **optional.String**| Git Sync Config Id | 
 **getDefaultFromOtherRepo** | **optional.Bool**| if true, return all the default entities | 

### Return type

[**ResponseDtoOverlayInputSetResponse**](ResponseDTOOverlayInputSetResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListInputSet**
> ResponseDtoPageResponseInputSetSummaryResponse ListInputSet(ctx, accountIdentifier, orgIdentifier, projectIdentifier, pipelineIdentifier, optional)
List all Input Sets for a pipeline

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the entity. | 
  **orgIdentifier** | **string**| Organization Identifier for the entity. | 
  **projectIdentifier** | **string**| Project Identifier for the entity. | 
  **pipelineIdentifier** | **string**| Pipeline identifier for which we need the Input Sets list. | 
 **optional** | ***InputSetsApiListInputSetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InputSetsApiListInputSetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **pageIndex** | **optional.Int32**| The number of the page to fetch | [default to 0]
 **pageSize** | **optional.Int32**| The number of the elements to fetch | [default to 100]
 **inputSetType** | **optional.String**| Type of Input Set needed: \&quot;INPUT_SET\&quot;, or \&quot;OVERLAY_INPUT_SET\&quot;, or \&quot;ALL\&quot;. If nothing is sent, ALL is considered. | [default to ALL]
 **searchTerm** | **optional.String**| Search term to filter out Input Sets based on name, identifier, tags. | 
 **sortOrders** | [**optional.Interface of []string**](string.md)| Sort criteria for the elements. | 
 **branch** | **optional.String**| Branch Name | 
 **repoIdentifier** | **optional.String**| Git Sync Config Id | 
 **getDefaultFromOtherRepo** | **optional.Bool**| if true, return all the default entities | 

### Return type

[**ResponseDtoPageResponseInputSetSummaryResponse**](ResponseDTOPageResponseInputSetSummaryResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MergeInputSets**
> ResponseDtoMergeInputSetResponse MergeInputSets(ctx, body, accountIdentifier, orgIdentifier, projectIdentifier, pipelineIdentifier, optional)
Merge given Input Sets into a single Runtime Input YAML

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MergeInputSetRequest**](MergeInputSetRequest.md)|  | 
  **accountIdentifier** | **string**| Account Identifier for the entity. | 
  **orgIdentifier** | **string**| Organization Identifier for the entity. | 
  **projectIdentifier** | **string**| Project Identifier for the entity. | 
  **pipelineIdentifier** | **string**| Identifier of the Pipeline to which the Input Sets belong | 
 **optional** | ***InputSetsApiMergeInputSetsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InputSetsApiMergeInputSetsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **pipelineBranch** | **optional.**| Github branch of the Pipeline to which the Input Sets belong | 
 **pipelineRepoID** | **optional.**| Github Repo identifier of the Pipeline to which the Input Sets belong | 
 **branch** | **optional.**| Branch Name | 
 **repoIdentifier** | **optional.**| Git Sync Config Id | 
 **getDefaultFromOtherRepo** | **optional.**| if true, return all the default entities | 

### Return type

[**ResponseDtoMergeInputSetResponse**](ResponseDTOMergeInputSetResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MergeRuntimeInputIntoPipeline**
> ResponseDtoMergeInputSetResponse MergeRuntimeInputIntoPipeline(ctx, body, accountIdentifier, orgIdentifier, projectIdentifier, pipelineIdentifier, optional)
Merge given Runtime Input YAML into the Pipeline

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MergeInputSetTemplateRequest**](MergeInputSetTemplateRequest.md)|  | 
  **accountIdentifier** | **string**| Account Identifier for the entity. | 
  **orgIdentifier** | **string**| Organization Identifier for the entity. | 
  **projectIdentifier** | **string**| Project Identifier for the entity. | 
  **pipelineIdentifier** | **string**| Identifier of the Pipeline to which the Input Sets belong | 
 **optional** | ***InputSetsApiMergeRuntimeInputIntoPipelineOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InputSetsApiMergeRuntimeInputIntoPipelineOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **pipelineBranch** | **optional.**| Github branch of the Pipeline to which the Input Sets belong | 
 **pipelineRepoID** | **optional.**| Github Repo identifier of the Pipeline to which the Input Sets belong | 
 **branch** | **optional.**| Branch Name | 
 **repoIdentifier** | **optional.**| Git Sync Config Id | 
 **getDefaultFromOtherRepo** | **optional.**| if true, return all the default entities | 

### Return type

[**ResponseDtoMergeInputSetResponse**](ResponseDTOMergeInputSetResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostInputSet**
> ResponseDtoInputSetResponse PostInputSet(ctx, body, accountIdentifier, orgIdentifier, projectIdentifier, pipelineIdentifier, optional)
Create an Input Set for a Pipeline

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**string**](string.md)| Input set YAML to be created. The account, org, project, and pipeline identifiers inside the YAML should match the query parameters | 
  **accountIdentifier** | **string**| Account Identifier for the entity. | 
  **orgIdentifier** | **string**| Organization Identifier for the entity. | 
  **projectIdentifier** | **string**| Project Identifier for the entity. | 
  **pipelineIdentifier** | **string**| Pipeline identifier for the input set. The input set will work only for the pipeline corresponding to this identifier. | 
 **optional** | ***InputSetsApiPostInputSetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InputSetsApiPostInputSetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **pipelineBranch** | **optional.**| Github branch of the Pipeline for which the Input Set is to be created | 
 **pipelineRepoID** | **optional.**| Github Repo identifier of the Pipeline for which the Input Set is to be created | 
 **branch** | **optional.**| Branch Name | 
 **repoIdentifier** | **optional.**| Git Sync Config Id | 
 **rootFolder** | **optional.**| Default Folder Path | 
 **filePath** | **optional.**| File Path | 
 **commitMsg** | **optional.**| File Path | 
 **isNewBranch** | **optional.**| Checks the new branch | [default to false]
 **baseBranch** | **optional.**| Default Branch | 

### Return type

[**ResponseDtoInputSetResponse**](ResponseDTOInputSetResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostOverlayInputSet**
> ResponseDtoOverlayInputSetResponse PostOverlayInputSet(ctx, body, accountIdentifier, orgIdentifier, projectIdentifier, pipelineIdentifier, optional)
Create an Overlay Input Set for a pipeline

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**string**](string.md)| Overlay Input Set YAML to be created. The Account, Org, Project, and Pipeline identifiers inside the YAML should match the query parameters | 
  **accountIdentifier** | **string**| Account Identifier for the entity. | 
  **orgIdentifier** | **string**| Organization Identifier for the entity. | 
  **projectIdentifier** | **string**| Project Identifier for the entity. | 
  **pipelineIdentifier** | **string**| Pipeline identifier for the overlay input set. The Overlay Input Set will work only for the Pipeline corresponding to this identifier. | 
 **optional** | ***InputSetsApiPostOverlayInputSetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InputSetsApiPostOverlayInputSetOpts struct
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

[**ResponseDtoOverlayInputSetResponse**](ResponseDTOOverlayInputSetResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutInputSet**
> ResponseDtoInputSetResponse PutInputSet(ctx, body, inputSetIdentifier, accountIdentifier, orgIdentifier, projectIdentifier, pipelineIdentifier, optional)
Update Input Set for Pipeline

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**string**](string.md)| Input set YAML to be updated. The Account, Org, Project, and Pipeline Ids inside the YAML should match the query parameters | 
  **inputSetIdentifier** | **string**| Identifier for the Input Set that needs to be updated. An Input Set corresponding to this identifier should already exist. | 
  **accountIdentifier** | **string**| Account Identifier for the entity. | 
  **orgIdentifier** | **string**| Organization Identifier for the entity. | 
  **projectIdentifier** | **string**| Project Identifier for the entity. | 
  **pipelineIdentifier** | **string**| Pipeline Id for the Input Set. The Input Set will work only for the Pipeline corresponding to this Id. | 
 **optional** | ***InputSetsApiPutInputSetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InputSetsApiPutInputSetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------






 **ifMatch** | **optional.**| Version of entity to match | 
 **pipelineBranch** | **optional.**| Github branch of the Pipeline for which the Input Set is to be created | 
 **pipelineRepoID** | **optional.**| Github Repo Id of the Pipeline for which the Input Set is to be created | 
 **branch** | **optional.**| Branch Name | 
 **repoIdentifier** | **optional.**| Git Sync Config Id | 
 **rootFolder** | **optional.**| Default Folder Path | 
 **filePath** | **optional.**| Default Folder Path | 
 **commitMsg** | **optional.**| Commit Message | 
 **lastObjectId** | **optional.**| Last Object Id | 
 **baseBranch** | **optional.**| Default Branch | 

### Return type

[**ResponseDtoInputSetResponse**](ResponseDTOInputSetResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutOverlayInputSet**
> ResponseDtoOverlayInputSetResponse PutOverlayInputSet(ctx, body, inputSetIdentifier, accountIdentifier, orgIdentifier, projectIdentifier, pipelineIdentifier, optional)
Update an Overlay Input Set for a pipeline

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**string**](string.md)| Overlay Input Set YAML to be updated. The Account, Org, Project, and Pipeline identifiers inside the YAML should match the query parameters, and the Overlay Input Set identifier cannot be changed. | 
  **inputSetIdentifier** | **string**| Identifier for the Overlay Input Set that needs to be updated. An Overlay Input Set corresponding to this identifier should already exist. | 
  **accountIdentifier** | **string**| Account Identifier for the entity. | 
  **orgIdentifier** | **string**| Organization Identifier for the entity. | 
  **projectIdentifier** | **string**| Project Identifier for the entity. | 
  **pipelineIdentifier** | **string**| Pipeline identifier for the Overlay Input Set. The Overlay Input Set will work only for the Pipeline corresponding to this identifier. | 
 **optional** | ***InputSetsApiPutOverlayInputSetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InputSetsApiPutOverlayInputSetOpts struct
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

[**ResponseDtoOverlayInputSetResponse**](ResponseDTOOverlayInputSetResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RuntimeInputTemplate**
> ResponseDtoInputSetTemplateWithReplacedExpressionsResponse RuntimeInputTemplate(ctx, accountIdentifier, orgIdentifier, projectIdentifier, pipelineIdentifier, optional)
Fetch Runtime Input Template for a Pipeline

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the entity. | 
  **orgIdentifier** | **string**| Organization Identifier for the entity. | 
  **projectIdentifier** | **string**| Project Identifier for the entity. | 
  **pipelineIdentifier** | **string**| Pipeline identifier for which we need the Runtime Input Template. | 
 **optional** | ***InputSetsApiRuntimeInputTemplateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InputSetsApiRuntimeInputTemplateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **body** | [**optional.Interface of InputSetTemplateRequest**](InputSetTemplateRequest.md)|  | 
 **branch** | **optional.**| Branch Name | 
 **repoIdentifier** | **optional.**| Git Sync Config Id | 
 **getDefaultFromOtherRepo** | **optional.**| if true, return all the default entities | 

### Return type

[**ResponseDtoInputSetTemplateWithReplacedExpressionsResponse**](ResponseDTOInputSetTemplateWithReplacedExpressionsResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

