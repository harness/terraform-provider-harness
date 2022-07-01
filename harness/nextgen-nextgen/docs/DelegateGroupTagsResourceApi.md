# nextgen{{classname}}

All URIs are relative to *https://app.harness.io/gateway*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddTagsToDelegateGroup**](DelegateGroupTagsResourceApi.md#AddTagsToDelegateGroup) | **Post** /ng/api/delegate-group-tags/{groupIdentifier} | Add given list of tags to the Delegate group
[**DeleteTagsFromDelegateGroup**](DelegateGroupTagsResourceApi.md#DeleteTagsFromDelegateGroup) | **Delete** /ng/api/delegate-group-tags/{groupIdentifier} | Deletes all tags from the Delegate group
[**ListTagsForDelegateGroup**](DelegateGroupTagsResourceApi.md#ListTagsForDelegateGroup) | **Get** /ng/api/delegate-group-tags/{groupIdentifier} | Retrieves list of tags attached with Delegate group
[**UpdateTagsOfDelegateGroup**](DelegateGroupTagsResourceApi.md#UpdateTagsOfDelegateGroup) | **Put** /ng/api/delegate-group-tags/{groupIdentifier} | Clears all existing tags with delegate group and attach given set of tags to delegate group.

# **AddTagsToDelegateGroup**
> RestResponseDelegateGroupDto AddTagsToDelegateGroup(ctx, body, groupIdentifier, optional)
Add given list of tags to the Delegate group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DelegateGroupTags**](DelegateGroupTags.md)| Set of tags | 
  **groupIdentifier** | **string**| Delegate Group Identifier | 
 **optional** | ***DelegateGroupTagsResourceApiAddTagsToDelegateGroupOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DelegateGroupTagsResourceApiAddTagsToDelegateGroupOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **accountIdentifier** | **optional.**| Account Identifier for the Entity. | 
 **orgIdentifier** | **optional.**| Organization Identifier for the Entity. | 
 **projectIdentifier** | **optional.**| Project Identifier for the Entity. | 

### Return type

[**RestResponseDelegateGroupDto**](RestResponseDelegateGroupDTO.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteTagsFromDelegateGroup**
> RestResponseDelegateGroupDto DeleteTagsFromDelegateGroup(ctx, groupIdentifier, optional)
Deletes all tags from the Delegate group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **groupIdentifier** | **string**| Delegate Group Identifier | 
 **optional** | ***DelegateGroupTagsResourceApiDeleteTagsFromDelegateGroupOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DelegateGroupTagsResourceApiDeleteTagsFromDelegateGroupOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accountIdentifier** | **optional.String**| Account Identifier for the Entity. | 
 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity. | 
 **projectIdentifier** | **optional.String**| Project Identifier for the Entity. | 

### Return type

[**RestResponseDelegateGroupDto**](RestResponseDelegateGroupDTO.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListTagsForDelegateGroup**
> RestResponseDelegateGroupDto ListTagsForDelegateGroup(ctx, groupIdentifier, optional)
Retrieves list of tags attached with Delegate group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **groupIdentifier** | **string**| Delegate Group Identifier | 
 **optional** | ***DelegateGroupTagsResourceApiListTagsForDelegateGroupOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DelegateGroupTagsResourceApiListTagsForDelegateGroupOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accountIdentifier** | **optional.String**| Account Identifier for the Entity. | 
 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity. | 
 **projectIdentifier** | **optional.String**| Project Identifier for the Entity. | 

### Return type

[**RestResponseDelegateGroupDto**](RestResponseDelegateGroupDTO.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateTagsOfDelegateGroup**
> RestResponseDelegateGroupDto UpdateTagsOfDelegateGroup(ctx, body, groupIdentifier, optional)
Clears all existing tags with delegate group and attach given set of tags to delegate group.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DelegateGroupTags**](DelegateGroupTags.md)| Set of tags | 
  **groupIdentifier** | **string**| Delegate Group Identifier | 
 **optional** | ***DelegateGroupTagsResourceApiUpdateTagsOfDelegateGroupOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DelegateGroupTagsResourceApiUpdateTagsOfDelegateGroupOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **accountIdentifier** | **optional.**| Account Identifier for the Entity. | 
 **orgIdentifier** | **optional.**| Organization Identifier for the Entity. | 
 **projectIdentifier** | **optional.**| Project Identifier for the Entity. | 

### Return type

[**RestResponseDelegateGroupDto**](RestResponseDelegateGroupDTO.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

