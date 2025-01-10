# har{{classname}}

All URIs are relative to */gateway/har/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetHelmArtifactDetails**](HelmArtifactsApi.md#GetHelmArtifactDetails) | **Get** /registry/{registry_ref}/+/artifact/{artifact}/version/{version}/helm/details | Describe Helm Artifact Detail
[**GetHelmArtifactManifest**](HelmArtifactsApi.md#GetHelmArtifactManifest) | **Get** /registry/{registry_ref}/+/artifact/{artifact}/version/{version}/helm/manifest | Describe Helm Artifact Manifest

# **GetHelmArtifactDetails**
> InlineResponse2009 GetHelmArtifactDetails(ctx, registryRef, artifact, version)
Describe Helm Artifact Detail

Get Helm Artifact Details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **registryRef** | **string**| Unique registry path. | 
  **artifact** | **string**| Name of artifact. | 
  **version** | **string**| Name of Artifact Version. | 

### Return type

[**InlineResponse2009**](inline_response_200_9.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetHelmArtifactManifest**
> InlineResponse20010 GetHelmArtifactManifest(ctx, registryRef, artifact, version)
Describe Helm Artifact Manifest

Get Helm Artifact Manifest

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **registryRef** | **string**| Unique registry path. | 
  **artifact** | **string**| Name of artifact. | 
  **version** | **string**| Name of Artifact Version. | 

### Return type

[**InlineResponse20010**](inline_response_200_10.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

