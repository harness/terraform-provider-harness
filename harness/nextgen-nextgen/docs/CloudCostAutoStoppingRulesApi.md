# nextgen{{classname}}

All URIs are relative to *https://app.harness.io/gateway*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AllAutoStoppingResources**](CloudCostAutoStoppingRulesApi.md#AllAutoStoppingResources) | **Get** /lw/api/accounts/{account_id}/autostopping/rules/{rule_id}/resources | List all the resources for an AutoStopping Rule
[**AutoStoppingRuleDetails**](CloudCostAutoStoppingRulesApi.md#AutoStoppingRuleDetails) | **Get** /lw/api/accounts/{account_id}/autostopping/rules/{rule_id} | Return AutoStopping Rule details
[**CumulativeAutoStoppingSavings**](CloudCostAutoStoppingRulesApi.md#CumulativeAutoStoppingSavings) | **Get** /lw/api/accounts/{account_id}/autostopping/rules/savings/cumulative | Return cumulative savings for all the AutoStopping Rules
[**DeleteAutoStoppingRule**](CloudCostAutoStoppingRulesApi.md#DeleteAutoStoppingRule) | **Delete** /lw/api/accounts/{account_id}/autostopping/rules/{rule_id} | Delete an AutoStopping Rule
[**GetAutoStoppingDiagnostics**](CloudCostAutoStoppingRulesApi.md#GetAutoStoppingDiagnostics) | **Get** /lw/api/accounts/{account_id}/autostopping/rules/{rule_id}/diagnostics | Return diagnostics result of an AutoStopping Rule
[**HealthOfAutoStoppingRule**](CloudCostAutoStoppingRulesApi.md#HealthOfAutoStoppingRule) | **Get** /lw/api/accounts/{account_id}/autostopping/rules/{rule_id}/health | Return health status of an AutoStopping Rule
[**ListAutoStoppingRules**](CloudCostAutoStoppingRulesApi.md#ListAutoStoppingRules) | **Get** /lw/api/accounts/{account_id}/autostopping/rules | List AutoStopping Rules
[**SavingsFromAutoStoppingRule**](CloudCostAutoStoppingRulesApi.md#SavingsFromAutoStoppingRule) | **Get** /lw/api/accounts/{account_id}/autostopping/rules/{rule_id}/savings | Return savings details for an AutoStopping Rule
[**ToggleAutostoppingRule**](CloudCostAutoStoppingRulesApi.md#ToggleAutostoppingRule) | **Put** /lw/api/accounts/{account_id}/autostopping/rules/{rule_id}/toggle_state | Disable/Enable an Autostopping Rule
[**UpdateAutoStoppingRule**](CloudCostAutoStoppingRulesApi.md#UpdateAutoStoppingRule) | **Post** /lw/api/accounts/{account_id}/autostopping/rules | Create an AutoStopping Rule

# **AllAutoStoppingResources**
> AllResourcesOfAccountResponse AllAutoStoppingResources(ctx, accountId, cloudAccountId, region, ruleId, accountIdentifier)
List all the resources for an AutoStopping Rule

Lists all the resources for an AutoStopping Rule for the given identifier.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **string**| Account Identifier for the Entity | 
  **cloudAccountId** | **string**| Connector ID | 
  **region** | **string**| Cloud region where resources belong to | 
  **ruleId** | **float64**| ID of the AutoStopping Rule for which you need to list the resources | 
  **accountIdentifier** | **string**| Account Identifier for the Entity | 

### Return type

[**AllResourcesOfAccountResponse**](AllResourcesOfAccountResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AutoStoppingRuleDetails**
> InlineResponse200 AutoStoppingRuleDetails(ctx, accountId, ruleId, accountIdentifier)
Return AutoStopping Rule details

Returns details of an AutoStopping Rule for the given identifier.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **string**| Account Identifier for the Entity | 
  **ruleId** | **float64**| ID of the AutoStopping Rule for which you need to fetch the details | 
  **accountIdentifier** | **string**| Account Identifier for the Entity | 

### Return type

[**InlineResponse200**](inline_response_200.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CumulativeAutoStoppingSavings**
> CumulativeSavingsResponse CumulativeAutoStoppingSavings(ctx, accountId, accountIdentifier)
Return cumulative savings for all the AutoStopping Rules

Returns cumulative savings for all the AutoStopping Rules.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **string**| Account Identifier for the Entity | 
  **accountIdentifier** | **string**| Account Identifier for the Entity | 

### Return type

[**CumulativeSavingsResponse**](CumulativeSavingsResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteAutoStoppingRule**
> DeleteAutoStoppingRule(ctx, ruleId, accountId, accountIdentifier)
Delete an AutoStopping Rule

Deletes an AutoStopping Rule for the given identifier.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ruleId** | **float64**| ID of the AutoStopping Rule that you want to delete | 
  **accountId** | **string**| Account Identifier for the Entity | 
  **accountIdentifier** | **string**| Account Identifier for the Entity | 

### Return type

 (empty response body)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAutoStoppingDiagnostics**
> ServiceDiagnosticsResponse GetAutoStoppingDiagnostics(ctx, accountId, ruleId, accountIdentifier)
Return diagnostics result of an AutoStopping Rule

Returns the diagnostics result of an AutoStopping rule for the given identifier.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **string**| Account Identifier for the Entity | 
  **ruleId** | **float64**| ID of the AutoStopping rule for which you need to fetch the diagnostics details | 
  **accountIdentifier** | **string**| Account Identifier for the Entity | 

### Return type

[**ServiceDiagnosticsResponse**](ServiceDiagnosticsResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **HealthOfAutoStoppingRule**
> ServiceHealthResponse HealthOfAutoStoppingRule(ctx, accountId, ruleId, accountIdentifier)
Return health status of an AutoStopping Rule

Returns health status of an AutoStopping Rule for the given identifier.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **string**| Account Identifier for the Entity | 
  **ruleId** | **float64**| ID of the AutoStopping Rule for which you need to fetch the health status | 
  **accountIdentifier** | **string**| Account Identifier for the Entity | 

### Return type

[**ServiceHealthResponse**](ServiceHealthResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListAutoStoppingRules**
> ServicesResponse ListAutoStoppingRules(ctx, accountId, accountIdentifier)
List AutoStopping Rules

Lists all the AutoStopping rules separated by comma-separated strings.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **string**| Account Identifier for the Entity | 
  **accountIdentifier** | **string**| Account Identifier for the Entity | 

### Return type

[**ServicesResponse**](ServicesResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SavingsFromAutoStoppingRule**
> interface{} SavingsFromAutoStoppingRule(ctx, accountId, ruleId, accountIdentifier, optional)
Return savings details for an AutoStopping Rule

Returns savings details for an AutoStopping rule for the given identifier and the specified time duration.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **string**| Account Identifier for the Entity | 
  **ruleId** | **float64**| ID of the AutoStopping Rule for which you want to fetch savings detail | 
  **accountIdentifier** | **string**| Account Identifier for the Entity | 
 **optional** | ***CloudCostAutoStoppingRulesApiSavingsFromAutoStoppingRuleOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CloudCostAutoStoppingRulesApiSavingsFromAutoStoppingRuleOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **from** | **optional.String**| Start time for the computation of savings | 
 **to** | **optional.String**| End time for the computation of savings | 
 **groupBy** | **optional.String**|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ToggleAutostoppingRule**
> ServicesResponse ToggleAutostoppingRule(ctx, accountId, ruleId, disable, accountIdentifier)
Disable/Enable an Autostopping Rule

Disables or enables an Autostopping Rule for the given identifier.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **string**| Account Identifier for the Entity | 
  **ruleId** | **string**| ID of the AutoStopping rule to be enabled/disabled | 
  **disable** | **bool**|  | 
  **accountIdentifier** | **string**| Account Identifier for the Entity | 

### Return type

[**ServicesResponse**](ServicesResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateAutoStoppingRule**
> LwServiceResponse UpdateAutoStoppingRule(ctx, body, accountId, accountIdentifier)
Create an AutoStopping Rule

Creates a new AutoStopping Rule.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SaveServiceRequest**](SaveServiceRequest.md)| Service definition of an AutoStopping rule | 
  **accountId** | **string**| Account Identifier for the Entity | 
  **accountIdentifier** | **string**| Account Identifier for the Entity | 

### Return type

[**LwServiceResponse**](LwServiceResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

