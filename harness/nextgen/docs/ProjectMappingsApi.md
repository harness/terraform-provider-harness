# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AppProjectMappingServiceCreate**](ProjectMappingsApi.md#AppProjectMappingServiceCreate) | **Post** /api/v1/agents/{agentIdentifier}/appprojectsmapping | CreateAppProjectMapping creates a new mapping between Harness Project and argo project
[**AppProjectMappingServiceCreateV2**](ProjectMappingsApi.md#AppProjectMappingServiceCreateV2) | **Post** /api/v2/agents/{agentIdentifier}/appprojectsmapping | CreateAppProjectMapping creates a new mapping between Harness Project and argo project
[**AppProjectMappingServiceDelete**](ProjectMappingsApi.md#AppProjectMappingServiceDelete) | **Delete** /api/v1/agents/{agentIdentifier}/appprojectsmapping/{name} | Delete an argo project to harness project mapping
[**AppProjectMappingServiceDeleteV2**](ProjectMappingsApi.md#AppProjectMappingServiceDeleteV2) | **Delete** /api/v2/agents/{agentIdentifier}/appprojectsmapping/{identifier} | Delete an argo project to harness project mapping
[**AppProjectMappingServiceGetAppProjectMappingList**](ProjectMappingsApi.md#AppProjectMappingServiceGetAppProjectMappingList) | **Get** /api/v1/appprojectsmapping | 
[**AppProjectMappingServiceGetAppProjectMappingListByAgent**](ProjectMappingsApi.md#AppProjectMappingServiceGetAppProjectMappingListByAgent) | **Get** /api/v1/agents/{agentIdentifier}/appprojectsmapping | 
[**AppProjectMappingServiceGetAppProjectMappingV2**](ProjectMappingsApi.md#AppProjectMappingServiceGetAppProjectMappingV2) | **Get** /api/v2/agents/{agentIdentifier}/appprojectsmapping/{identifier} | 
[**AppProjectMappingServiceGetAppProjectMappingsListByAgentV2**](ProjectMappingsApi.md#AppProjectMappingServiceGetAppProjectMappingsListByAgentV2) | **Get** /api/v2/agents/{agentIdentifier}/appprojectsmappings | V2
[**AppProjectMappingServiceUpdateV2**](ProjectMappingsApi.md#AppProjectMappingServiceUpdateV2) | **Put** /api/v2/agents/{agentIdentifier}/appprojectsmapping/{identifier} | CreateAppProjectMapping creates a new mapping between Harness Project and argo project

# **AppProjectMappingServiceCreate**
> Servicev1Empty AppProjectMappingServiceCreate(ctx, body, agentIdentifier, optional)
CreateAppProjectMapping creates a new mapping between Harness Project and argo project

Creates Harness-Argo project mappings.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Servicev1AppProjectMapping**](Servicev1AppProjectMapping.md)|  | 
  **agentIdentifier** | **string**| Agent identifier for entity. | 
 **optional** | ***ProjectMappingsApiAppProjectMappingServiceCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectMappingsApiAppProjectMappingServiceCreateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **accountIdentifier** | **optional.**| Account Identifier for the Entity. | 
 **orgIdentifier** | **optional.**| Organization Identifier for the Entity. | 
 **projectIdentifier** | **optional.**| Project Identifier for the Entity. | 

### Return type

[**Servicev1Empty**](servicev1Empty.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AppProjectMappingServiceCreateV2**
> V1AppProjectMappingV2 AppProjectMappingServiceCreateV2(ctx, body, agentIdentifier)
CreateAppProjectMapping creates a new mapping between Harness Project and argo project

Creates Harness-Argo project mappings.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**V1AppProjectMappingCreateRequestV2**](V1AppProjectMappingCreateRequestV2.md)|  | 
  **agentIdentifier** | **string**| Agent identifier for entity. | 

### Return type

[**V1AppProjectMappingV2**](v1AppProjectMappingV2.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AppProjectMappingServiceDelete**
> Servicev1Empty AppProjectMappingServiceDelete(ctx, agentIdentifier, name, optional)
Delete an argo project to harness project mapping

Delete Harness-Argo project mappings.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **agentIdentifier** | **string**| Agent identifier for entity. | 
  **name** | **string**|  | 
 **optional** | ***ProjectMappingsApiAppProjectMappingServiceDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectMappingsApiAppProjectMappingServiceDeleteOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **accountIdentifier** | **optional.String**| Account Identifier for the Entity. | 
 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity. | 
 **projectIdentifier** | **optional.String**| Project Identifier for the Entity. | 

### Return type

[**Servicev1Empty**](servicev1Empty.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AppProjectMappingServiceDeleteV2**
> Servicev1Empty AppProjectMappingServiceDeleteV2(ctx, agentIdentifier, identifier, optional)
Delete an argo project to harness project mapping

Delete Harness-Argo project mappings.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **agentIdentifier** | **string**| Agent identifier for entity. | 
  **identifier** | **string**| app project mapping identifier. | 
 **optional** | ***ProjectMappingsApiAppProjectMappingServiceDeleteV2Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectMappingsApiAppProjectMappingServiceDeleteV2Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **accountIdentifier** | **optional.String**| Account Identifier for the Entity. | 
 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity. | 
 **projectIdentifier** | **optional.String**| Project Identifier for the Entity. | 

### Return type

[**Servicev1Empty**](servicev1Empty.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AppProjectMappingServiceGetAppProjectMappingList**
> Servicev1AppProjectMapping AppProjectMappingServiceGetAppProjectMappingList(ctx, optional)


Retrieves Harness-Argo project mappings list.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ProjectMappingsApiAppProjectMappingServiceGetAppProjectMappingListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectMappingsApiAppProjectMappingServiceGetAppProjectMappingListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **agentIdentifier** | **optional.String**| Agent identifier for entity. | 
 **accountIdentifier** | **optional.String**| Account Identifier for the Entity. | 
 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity. | 
 **projectIdentifier** | **optional.String**| Project Identifier for the Entity. | 

### Return type

[**Servicev1AppProjectMapping**](servicev1AppProjectMapping.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AppProjectMappingServiceGetAppProjectMappingListByAgent**
> Servicev1AppProjectMapping AppProjectMappingServiceGetAppProjectMappingListByAgent(ctx, agentIdentifier, optional)


Retrieves Harness-Argo project mappings list by agent.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **agentIdentifier** | **string**| Agent identifier for entity. | 
 **optional** | ***ProjectMappingsApiAppProjectMappingServiceGetAppProjectMappingListByAgentOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectMappingsApiAppProjectMappingServiceGetAppProjectMappingListByAgentOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accountIdentifier** | **optional.String**| Account Identifier for the Entity. | 
 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity. | 
 **projectIdentifier** | **optional.String**| Project Identifier for the Entity. | 

### Return type

[**Servicev1AppProjectMapping**](servicev1AppProjectMapping.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AppProjectMappingServiceGetAppProjectMappingV2**
> V1AppProjectMappingV2 AppProjectMappingServiceGetAppProjectMappingV2(ctx, agentIdentifier, identifier, optional)


Retrieves Harness-Argo project mapping for the given identifier.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **agentIdentifier** | **string**| Agent identifier for entity. | 
  **identifier** | **string**| app project mapping identifier. | 
 **optional** | ***ProjectMappingsApiAppProjectMappingServiceGetAppProjectMappingV2Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectMappingsApiAppProjectMappingServiceGetAppProjectMappingV2Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **accountIdentifier** | **optional.String**| Account Identifier for the Entity. | 
 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity. | 
 **projectIdentifier** | **optional.String**| Project Identifier for the Entity. | 
 **argoProjectName** | **optional.String**|  | 

### Return type

[**V1AppProjectMappingV2**](v1AppProjectMappingV2.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AppProjectMappingServiceGetAppProjectMappingsListByAgentV2**
> V1AppProjectMappingV2List AppProjectMappingServiceGetAppProjectMappingsListByAgentV2(ctx, agentIdentifier, optional)
V2

Retrieves Harness-Argo project mappings list by agent.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **agentIdentifier** | **string**| Agent identifier for entity. | 
 **optional** | ***ProjectMappingsApiAppProjectMappingServiceGetAppProjectMappingsListByAgentV2Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectMappingsApiAppProjectMappingServiceGetAppProjectMappingsListByAgentV2Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **identifier** | **optional.String**| app project mapping identifier. | 
 **accountIdentifier** | **optional.String**| Account Identifier for the Entity. | 
 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity. | 
 **projectIdentifier** | **optional.String**| Project Identifier for the Entity. | 
 **argoProjectName** | **optional.String**|  | 

### Return type

[**V1AppProjectMappingV2List**](v1AppProjectMappingV2List.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AppProjectMappingServiceUpdateV2**
> V1AppProjectMappingV2 AppProjectMappingServiceUpdateV2(ctx, body, agentIdentifier, identifier)
CreateAppProjectMapping creates a new mapping between Harness Project and argo project

Creates Harness-Argo project mappings.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**V1AppProjectMappingQueryV2**](V1AppProjectMappingQueryV2.md)|  | 
  **agentIdentifier** | **string**| Agent identifier for entity. | 
  **identifier** | **string**| app project mapping identifier. | 

### Return type

[**V1AppProjectMappingV2**](v1AppProjectMappingV2.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

