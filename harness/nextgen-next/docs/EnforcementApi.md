# {{classname}}

All URIs are relative to *https://app.harness.io/gateway/ng/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAllFeatureRestrictionMetadata**](EnforcementApi.md#GetAllFeatureRestrictionMetadata) | **Get** /enforcement/metadata | Fetch All Feature Restriction Metadata
[**GetEnabledFeatureRestrictionDetailByAccountId**](EnforcementApi.md#GetEnabledFeatureRestrictionDetailByAccountId) | **Get** /enforcement/enabled | Fetch List of Enabled Feature Restriction Detail for The Account
[**GetFeatureRestrictionDetail**](EnforcementApi.md#GetFeatureRestrictionDetail) | **Post** /enforcement | Fetch Feature Restriction Detail
[**GetFeatureRestrictionDetails**](EnforcementApi.md#GetFeatureRestrictionDetails) | **Post** /enforcement/details | Fetch List of Feature Restriction Detail

# **GetAllFeatureRestrictionMetadata**
> ResponseDtoListFeatureRestrictionMetadata GetAllFeatureRestrictionMetadata(ctx, )
Fetch All Feature Restriction Metadata

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ResponseDtoListFeatureRestrictionMetadata**](ResponseDTOListFeatureRestrictionMetadata.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEnabledFeatureRestrictionDetailByAccountId**
> ResponseDtoListFeatureRestrictionDetails GetEnabledFeatureRestrictionDetailByAccountId(ctx, accountIdentifier)
Fetch List of Enabled Feature Restriction Detail for The Account

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account id to get the enable features for the account | 

### Return type

[**ResponseDtoListFeatureRestrictionDetails**](ResponseDTOListFeatureRestrictionDetails.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFeatureRestrictionDetail**
> ResponseDtoFeatureRestrictionDetails GetFeatureRestrictionDetail(ctx, body, accountIdentifier)
Fetch Feature Restriction Detail

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**FeatureRestrictionDetailRequest**](FeatureRestrictionDetailRequest.md)|  | 
  **accountIdentifier** | **string**| Account id to get the feature restriction detail. | 

### Return type

[**ResponseDtoFeatureRestrictionDetails**](ResponseDTOFeatureRestrictionDetails.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFeatureRestrictionDetails**
> ResponseDtoListFeatureRestrictionDetails GetFeatureRestrictionDetails(ctx, body, accountIdentifier)
Fetch List of Feature Restriction Detail

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**FeatureRestrictionDetailListRequest**](FeatureRestrictionDetailListRequest.md)|  | 
  **accountIdentifier** | **string**| Account id to get the feature restriction detail. | 

### Return type

[**ResponseDtoListFeatureRestrictionDetails**](ResponseDTOListFeatureRestrictionDetails.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

