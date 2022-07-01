# nextgen{{classname}}

All URIs are relative to *https://app.harness.io/gateway*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAutoStoppingSchedules**](CloudCostAutoStoppingFixedSchedulesApi.md#CreateAutoStoppingSchedules) | **Post** /lw/api/accounts/{account_id}/schedules | Create a fixed schedule for an AutoStopping Rule
[**DeleteAutoStoppingSchedule**](CloudCostAutoStoppingFixedSchedulesApi.md#DeleteAutoStoppingSchedule) | **Delete** /lw/api/accounts/{account_id}/schedules/{schedule_id} | Delete a fixed schedule for AutoStopping Rule.
[**ListAutoStoppingSchedules**](CloudCostAutoStoppingFixedSchedulesApi.md#ListAutoStoppingSchedules) | **Get** /lw/api/accounts/{account_id}/schedules | Return all the AutoStopping Rule fixed schedules

# **CreateAutoStoppingSchedules**
> FixedSchedule CreateAutoStoppingSchedules(ctx, body, accountId, cloudAccountId, accountIdentifier)
Create a fixed schedule for an AutoStopping Rule

Creates an AutoStopping rule to run resources based on the schedule.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SaveStaticSchedulesRequest**](SaveStaticSchedulesRequest.md)| Fixed schedule payload | 
  **accountId** | **string**| Account Identifier for the Entity | 
  **cloudAccountId** | **string**| Connector ID | 
  **accountIdentifier** | **string**| Account Identifier for the Entity | 

### Return type

[**FixedSchedule**](FixedSchedule.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteAutoStoppingSchedule**
> InlineResponse2001 DeleteAutoStoppingSchedule(ctx, accountId, scheduleId, accountIdentifier)
Delete a fixed schedule for AutoStopping Rule.

Deletes a fixed schedule for the given AutoStopping Rule.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **string**| Account Identifier for the Entity | 
  **scheduleId** | **float64**| ID of a fixed schedule added to an AutoStopping rule | 
  **accountIdentifier** | **string**| Account Identifier for the Entity | 

### Return type

[**InlineResponse2001**](inline_response_200_1.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListAutoStoppingSchedules**
> FixedSchedulesListResponse ListAutoStoppingSchedules(ctx, accountId, cloudAccountId, accountIdentifier, resId, resType)
Return all the AutoStopping Rule fixed schedules

Returns all the AutoStopping Rule fixed schedules for the given identifier.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **string**| Account Identifier for the Entity | 
  **cloudAccountId** | **string**| Connector ID | 
  **accountIdentifier** | **string**| Account Identifier for the Entity | 
  **resId** | **string**| IDs of resources whose fixed schedules are to be fetched. This can be an AutoStopping rule ID if the res_type is \&quot;autostop_rule\&quot; | 
  **resType** | **string**| Type of resource to which schedules are attached | 

### Return type

[**FixedSchedulesListResponse**](FixedSchedulesListResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

