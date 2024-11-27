# dbops{{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1GetDbinstanceLog**](LogIngestApi.md#V1GetDbinstanceLog) | **Get** /v1/orgs/{org}/projects/{project}/dbschema/{dbschema}/instance/{dbinstance}/logs/{log} | 
[**V1IngestLogs**](LogIngestApi.md#V1IngestLogs) | **Post** /v1/orgs/{org}/projects/{project}/dbschema/{dbschema}/instance/{dbinstance}/logs | Ingest database migration logs

# **V1GetDbinstanceLog**
> ParsedLogOut V1GetDbinstanceLog(ctx, org, project, dbschema, dbinstance, log, optional)


Retrieves the specified log event

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **org** | **string**| Organization identifier | 
  **project** | **string**| Project identifier | 
  **dbschema** | **string**| Identifier of the database schema | 
  **dbinstance** | **string**| database instance unique id | 
  **log** | **string**|  | 
 **optional** | ***LogIngestApiV1GetDbinstanceLogOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LogIngestApiV1GetDbinstanceLogOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **harnessAccount** | **optional.String**| Identifier field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 

### Return type

[**ParsedLogOut**](ParsedLogOut.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1IngestLogs**
> ParsedLogOut V1IngestLogs(ctx, body, org, project, dbschema, dbinstance, optional)
Ingest database migration logs

Ingest database migration logs to update the state of the database

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**string**](string.md)| Request body for log ingestion | 
  **org** | **string**| Organization identifier | 
  **project** | **string**| Project identifier | 
  **dbschema** | **string**| Identifier of the database schema | 
  **dbinstance** | **string**| database instance unique id | 
 **optional** | ***LogIngestApiV1IngestLogsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LogIngestApiV1IngestLogsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **harnessAccount** | **optional.**| Identifier field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 

### Return type

[**ParsedLogOut**](ParsedLogOut.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

