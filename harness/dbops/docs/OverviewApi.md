# dbops{{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1GetDbOverview**](OverviewApi.md#V1GetDbOverview) | **Get** /v1/orgs/{org}/projects/{project}/dbschema/overview | Get overview

# **V1GetDbOverview**
> InlineResponse2001 V1GetDbOverview(ctx, org, project, optional)
Get overview

retrieves total dbSchemas, dbInstances and the latest 5 instances deployed

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **org** | **string**| Organization identifier | 
  **project** | **string**| Project identifier | 
 **optional** | ***OverviewApiV1GetDbOverviewOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OverviewApiV1GetDbOverviewOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **harnessAccount** | **optional.String**| Identifier field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 

### Return type

[**InlineResponse2001**](inline_response_200_1.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

