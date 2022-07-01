# nextgen{{classname}}

All URIs are relative to *https://app.harness.io/gateway*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloneBudget**](CloudCostBudgetsApi.md#CloneBudget) | **Post** /ccm/api/budgets/{id} | Clone a budget
[**CreateBudget**](CloudCostBudgetsApi.md#CreateBudget) | **Post** /ccm/api/budgets | Create a Budget
[**DeleteBudget**](CloudCostBudgetsApi.md#DeleteBudget) | **Delete** /ccm/api/budgets/{id} | Delete a budget
[**GetBudget**](CloudCostBudgetsApi.md#GetBudget) | **Get** /ccm/api/budgets/{id} | Fetch Budget details
[**GetCostDetails**](CloudCostBudgetsApi.md#GetCostDetails) | **Get** /ccm/api/budgets/{id}/costDetails | Fetch the cost details of a Budget
[**ListBudgets**](CloudCostBudgetsApi.md#ListBudgets) | **Get** /ccm/api/budgets | List all the Budgets
[**ListBudgetsForPerspective**](CloudCostBudgetsApi.md#ListBudgetsForPerspective) | **Get** /ccm/api/budgets/perspectiveBudgets | List all the Budgets associated with a Perspective
[**UpdateBudget**](CloudCostBudgetsApi.md#UpdateBudget) | **Put** /ccm/api/budgets/{id} | Update an existing budget

# **CloneBudget**
> ResponseDtoString CloneBudget(ctx, accountIdentifier, id, cloneName)
Clone a budget

Clone a Cloud Cost Budget using the given Budget ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 
  **id** | **string**| Unique identifier for the budget | 
  **cloneName** | **string**| Name of the new budget | 

### Return type

[**ResponseDtoString**](ResponseDTOString.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateBudget**
> ResponseDtoString CreateBudget(ctx, body, accountIdentifier)
Create a Budget

Create a Budget to set and receive alerts when your costs exceed (or are forecasted to exceed) your budget amount.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Budget**](Budget.md)| Budget definition | 
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 

### Return type

[**ResponseDtoString**](ResponseDTOString.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteBudget**
> ResponseDtoString DeleteBudget(ctx, accountIdentifier, id)
Delete a budget

Delete a Cloud Cost Budget for the given Budget ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 
  **id** | **string**| Unique identifier for the budget | 

### Return type

[**ResponseDtoString**](ResponseDTOString.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBudget**
> ResponseDtoBudget GetBudget(ctx, accountIdentifier, id)
Fetch Budget details

Fetch details of a Cloud Cost Budget for the given Budget ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 
  **id** | **string**| Unique identifier for the budget | 

### Return type

[**ResponseDtoBudget**](ResponseDTOBudget.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCostDetails**
> ResponseDtoBudgetData GetCostDetails(ctx, accountIdentifier, id)
Fetch the cost details of a Budget

Fetch the cost details of a Cloud Cost Budget for the given Budget ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 
  **id** | **string**| Unique identifier for the Budget | 

### Return type

[**ResponseDtoBudgetData**](ResponseDTOBudgetData.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListBudgets**
> ResponseDtoListBudget ListBudgets(ctx, accountIdentifier)
List all the Budgets

List all the Cloud Cost Budgets.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 

### Return type

[**ResponseDtoListBudget**](ResponseDTOListBudget.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListBudgetsForPerspective**
> ResponseDtoListBudget ListBudgetsForPerspective(ctx, accountIdentifier, perspectiveId)
List all the Budgets associated with a Perspective

List all the Cloud Cost Budgets associated for the given Perspective ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 
  **perspectiveId** | **string**| Unique identifier for the Perspective | 

### Return type

[**ResponseDtoListBudget**](ResponseDTOListBudget.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateBudget**
> ResponseDtoString UpdateBudget(ctx, body, accountIdentifier, id)
Update an existing budget

Update an existing Cloud Cost Budget for the given Budget ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Budget**](Budget.md)| The Budget object | 
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 
  **id** | **string**| Unique identifier for the budget | 

### Return type

[**ResponseDtoString**](ResponseDTOString.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

