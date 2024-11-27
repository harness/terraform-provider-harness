# dbops{{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1GetDeployedState**](DeployedStateApi.md#V1GetDeployedState) | **Post** /v1/orgs/{org}/projects/{project}/dbschema/{dbschema}/instance/{dbinstance}/deployedState | Get Deployed State

# **V1GetDeployedState**
> []DeployedStateOutput V1GetDeployedState(ctx, org, project, dbschema, dbinstance, optional)
Get Deployed State

Status of changeset deployment as part of execution with comparison to earlier state.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **org** | **string**| Organization identifier | 
  **project** | **string**| Project identifier | 
  **dbschema** | **string**| Identifier of the database schema | 
  **dbinstance** | **string**| database instance unique id | 
 **optional** | ***DeployedStateApiV1GetDeployedStateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeployedStateApiV1GetDeployedStateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **body** | [**optional.Interface of DeployedStateInput**](DeployedStateInput.md)|  | 
 **harnessAccount** | **optional.**| Identifier field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 
 **page** | **optional.**| Pagination page number strategy: Specify the page number within the paginated collection related to the number of items on each page. | [default to 0]
 **limit** | **optional.**| Pagination: Number of items to return. | [default to 10]

### Return type

[**[]DeployedStateOutput**](DeployedStateOutput.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

