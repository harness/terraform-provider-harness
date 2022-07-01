# nextgen{{classname}}

All URIs are relative to *https://app.harness.io/gateway*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AccessPointRules**](CloudCostAutoStoppingLoadBalancersApi.md#AccessPointRules) | **Get** /lw/api/accounts/{account_id}/autostopping/loadbalancers/{lb_id}/rules | Return all the AutoStopping Rules in a load balancer
[**CreateLoadBalancer**](CloudCostAutoStoppingLoadBalancersApi.md#CreateLoadBalancer) | **Post** /lw/api/accounts/{account_id}/autostopping/loadbalancers | Create a load balancer
[**DeleteLoadBalancer**](CloudCostAutoStoppingLoadBalancersApi.md#DeleteLoadBalancer) | **Delete** /lw/api/accounts/{account_id}/autostopping/loadbalancers | Delete load balancers and the associated resources
[**DescribeLoadBalancer**](CloudCostAutoStoppingLoadBalancersApi.md#DescribeLoadBalancer) | **Get** /lw/api/accounts/{account_id}/autostopping/loadbalancers/{lb_id} | Return details of a load balancer
[**EditLoadBalancer**](CloudCostAutoStoppingLoadBalancersApi.md#EditLoadBalancer) | **Put** /lw/api/accounts/{account_id}/autostopping/loadbalancers | Update a load balancer
[**ListLoadBalancers**](CloudCostAutoStoppingLoadBalancersApi.md#ListLoadBalancers) | **Get** /lw/api/accounts/{account_id}/autostopping/loadbalancers | Return all the load balancers
[**LoadBalancerActivity**](CloudCostAutoStoppingLoadBalancersApi.md#LoadBalancerActivity) | **Get** /lw/api/accounts/{account_id}/autostopping/loadbalancers/{lb_id}/last_active_at | Return last activity details of a load balancer

# **AccessPointRules**
> ServicesResponse AccessPointRules(ctx, accountId, lbId, accountIdentifier)
Return all the AutoStopping Rules in a load balancer

Returns all the AutoStopping Rules for the given load balancer identifier.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **string**| Account Identifier for the Entity | 
  **lbId** | **string**| ID of the load balancer for which you want to fetch the list of AutoStopping Rules | 
  **accountIdentifier** | **string**| Account Identifier for the Entity | 

### Return type

[**ServicesResponse**](ServicesResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateLoadBalancer**
> CreateAccessPointResponse CreateLoadBalancer(ctx, body, accountId, accountIdentifier)
Create a load balancer

Creates a load balancer.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AccessPoint**](AccessPoint.md)|  | 
  **accountId** | **string**| Account Identifier for the Entity | 
  **accountIdentifier** | **string**| Account Identifier for the Entity | 

### Return type

[**CreateAccessPointResponse**](CreateAccessPointResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteLoadBalancer**
> DeleteLoadBalancer(ctx, body, accountId, accountIdentifier)
Delete load balancers and the associated resources

Deletes load balancers and the associated resources for the given identifier.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DeleteAccessPointPayload**](DeleteAccessPointPayload.md)|  | 
  **accountId** | **string**| Account Identifier for the Entity | 
  **accountIdentifier** | **string**| Account Identifier for the Entity | 

### Return type

 (empty response body)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DescribeLoadBalancer**
> GetAccessPointResponse DescribeLoadBalancer(ctx, accountId, lbId, accountIdentifier)
Return details of a load balancer

Retuns details of a load balancer for the given identifier.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **string**| Account Identifier for the Entity | 
  **lbId** | **string**| ID of the load balancer for which you want to fetch the details | 
  **accountIdentifier** | **string**| Account Identifier for the Entity | 

### Return type

[**GetAccessPointResponse**](GetAccessPointResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EditLoadBalancer**
> CreateAccessPointResponse EditLoadBalancer(ctx, body, accountId, accountIdentifier)
Update a load balancer

Updates a load balancer for the given identifier.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AccessPoint**](AccessPoint.md)|  | 
  **accountId** | **string**| Account Identifier for the Entity | 
  **accountIdentifier** | **string**| Account Identifier for the Entity | 

### Return type

[**CreateAccessPointResponse**](CreateAccessPointResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListLoadBalancers**
> ListAccessPointResponse ListLoadBalancers(ctx, accountId, cloudAccountId, accountIdentifier, optional)
Return all the load balancers

Returns all the load balancers for the given identifier.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **string**| Account Identifier for the Entity | 
  **cloudAccountId** | **string**| Connector ID | 
  **accountIdentifier** | **string**| Account Identifier for the Entity | 
 **optional** | ***CloudCostAutoStoppingLoadBalancersApiListLoadBalancersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CloudCostAutoStoppingLoadBalancersApiListLoadBalancersOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **vpc** | **optional.String**| Virtual Private Cloud (VPC) | 
 **region** | **optional.String**| Cloud region where access point is installed | 

### Return type

[**ListAccessPointResponse**](ListAccessPointResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LoadBalancerActivity**
> AccessPointActivityResponse LoadBalancerActivity(ctx, accountId, lbId, accountIdentifier)
Return last activity details of a load balancer

Returns the last activity details for the given load balancer identifier.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **string**| Account Identifier for the Entity | 
  **lbId** | **string**| ID of the load balancer for which you want to fetch the most recent activity details | 
  **accountIdentifier** | **string**| Account Identifier for the Entity | 

### Return type

[**AccessPointActivityResponse**](AccessPointActivityResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

