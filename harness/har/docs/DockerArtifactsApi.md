# har{{classname}}

All URIs are relative to */gateway/har/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDockerArtifactDetails**](DockerArtifactsApi.md#GetDockerArtifactDetails) | **Get** /registry/{registry_ref}/+/artifact/{artifact}/version/{version}/docker/details | Describe Docker Artifact Detail
[**GetDockerArtifactIntegrationDetails**](DockerArtifactsApi.md#GetDockerArtifactIntegrationDetails) | **Get** /registry/{registry_ref}/+/artifact/{artifact}/version/{version}/docker/integrationdetails | Describe Docker Artifact Integration Detail
[**GetDockerArtifactLayers**](DockerArtifactsApi.md#GetDockerArtifactLayers) | **Get** /registry/{registry_ref}/+/artifact/{artifact}/version/{version}/docker/layers | Describe Docker Artifact Layers
[**GetDockerArtifactManifest**](DockerArtifactsApi.md#GetDockerArtifactManifest) | **Get** /registry/{registry_ref}/+/artifact/{artifact}/version/{version}/docker/manifest | Describe Docker Artifact Manifest
[**GetDockerArtifactManifests**](DockerArtifactsApi.md#GetDockerArtifactManifests) | **Get** /registry/{registry_ref}/+/artifact/{artifact}/version/{version}/docker/manifests | Describe Docker Artifact Manifests

# **GetDockerArtifactDetails**
> InlineResponse2004 GetDockerArtifactDetails(ctx, registryRef, artifact, version, digest)
Describe Docker Artifact Detail

Get Docker Artifact Details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **registryRef** | **string**| Unique registry path. | 
  **artifact** | **string**| Name of artifact. | 
  **version** | **string**| Name of Artifact Version. | 
  **digest** | **string**| Digest. | 

### Return type

[**InlineResponse2004**](inline_response_200_4.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDockerArtifactIntegrationDetails**
> InlineResponse2005 GetDockerArtifactIntegrationDetails(ctx, registryRef, artifact, version, digest)
Describe Docker Artifact Integration Detail

Get Docker Artifact Integration Details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **registryRef** | **string**| Unique registry path. | 
  **artifact** | **string**| Name of artifact. | 
  **version** | **string**| Name of Artifact Version. | 
  **digest** | **string**| Digest. | 

### Return type

[**InlineResponse2005**](inline_response_200_5.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDockerArtifactLayers**
> InlineResponse2006 GetDockerArtifactLayers(ctx, registryRef, artifact, version, digest)
Describe Docker Artifact Layers

Get Docker Artifact Layers

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **registryRef** | **string**| Unique registry path. | 
  **artifact** | **string**| Name of artifact. | 
  **version** | **string**| Name of Artifact Version. | 
  **digest** | **string**| Digest. | 

### Return type

[**InlineResponse2006**](inline_response_200_6.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDockerArtifactManifest**
> InlineResponse2007 GetDockerArtifactManifest(ctx, registryRef, artifact, version, digest)
Describe Docker Artifact Manifest

Get Docker Artifact Manifest

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **registryRef** | **string**| Unique registry path. | 
  **artifact** | **string**| Name of artifact. | 
  **version** | **string**| Name of Artifact Version. | 
  **digest** | **string**| Digest. | 

### Return type

[**InlineResponse2007**](inline_response_200_7.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDockerArtifactManifests**
> InlineResponse2008 GetDockerArtifactManifests(ctx, registryRef, artifact, version)
Describe Docker Artifact Manifests

Get Docker Artifact Manifests

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **registryRef** | **string**| Unique registry path. | 
  **artifact** | **string**| Name of artifact. | 
  **version** | **string**| Name of Artifact Version. | 

### Return type

[**InlineResponse2008**](inline_response_200_8.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

