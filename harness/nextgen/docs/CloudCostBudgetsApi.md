# {{classname}}

All URIs are relative to *https://app.harness.io/gateway*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloneBudget**](CloudCostBudgetsApi.md#CloneBudget) | **Post** /ccm/api/budgets/{id} | Clone an existing Budget
[**CreateBudget**](CloudCostBudgetsApi.md#CreateBudget) | **Post** /ccm/api/budgets | Create a Budget
[**DeleteBudget**](CloudCostBudgetsApi.md#DeleteBudget) | **Delete** /ccm/api/budgets/{id} | Delete an existing Budget
[**GetBudget**](CloudCostBudgetsApi.md#GetBudget) | **Get** /ccm/api/budgets/{id} | Get a Budget
[**GetCostDetails**](CloudCostBudgetsApi.md#GetCostDetails) | **Get** /ccm/api/budgets/{id}/costDetails | Get the cost details associated with a Budget
[**ListBudgets**](CloudCostBudgetsApi.md#ListBudgets) | **Get** /ccm/api/budgets | List all the Budgets
[**ListBudgetsForPerspective**](CloudCostBudgetsApi.md#ListBudgetsForPerspective) | **Get** /ccm/api/budgets/perspectiveBudgets | List all the Budgets associated with a Perspective
[**UpdateBudget**](CloudCostBudgetsApi.md#UpdateBudget) | **Put** /ccm/api/budgets/{id} | Update an existing Budget

# **CloneBudget**
> ResponseDtoString CloneBudget(ctx, accountIdentifier, id, cloneName)
Clone an existing Budget

Clone an existing Budget using an existing Budget identifier

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the entity | 
  **id** | **string**| The identifier of the Budget | 
  **cloneName** | **string**| The name of the new Budget created after cloning operation | 

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

Creates a Budget from the Budget object passed as a request body

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Budget**](Budget.md)| The Budget definition | 
  **accountIdentifier** | **string**| Account Identifier for the entity | 

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
Delete an existing Budget

Delete an existing Cloud Cost Budget by identifier

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the entity | 
  **id** | **string**| The identifier of the Budget | 

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
Get a Budget

Get a Cloud Cost Budget by an identifier

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the entity | 
  **id** | **string**| The identifier of an existing Budget | 

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
Get the cost details associated with a Budget

Get the cost details associated with a Cloud Cost Budget identifier

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the entity | 
  **id** | **string**| The identifier of the Budget | 

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

List all the Cloud Cost Budgets

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the entity | 

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

List all the Cloud Cost Budgets associated with a Cloud Cost Perspective identifier

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the entity | 
  **perspectiveId** | **string**| The identifier of an existing Perspective | 

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
Update an existing Budget

Update an existing Budget using the identifier passed as a path param

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Budget**](Budget.md)| The Budget object as a request body | 
  **accountIdentifier** | **string**| Account Identifier for the entity | 
  **id** | **string**| The identifier of an existing Budget | 

### Return type

[**ResponseDtoString**](ResponseDTOString.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

