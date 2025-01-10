# har{{classname}}

All URIs are relative to */gateway/har/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteArtifact**](ArtifactsApi.md#DeleteArtifact) | **Delete** /registry/{registry_ref}/+/artifact/{artifact} | Delete Artifact
[**DeleteArtifactVersion**](ArtifactsApi.md#DeleteArtifactVersion) | **Delete** /registry/{registry_ref}/+/artifact/{artifact}/version/{version} | Delete an Artifact Version
[**GetAllArtifactVersions**](ArtifactsApi.md#GetAllArtifactVersions) | **Get** /registry/{registry_ref}/+/artifact/{artifact}/versions | List Artifact Versions
[**GetArtifactDeployments**](ArtifactsApi.md#GetArtifactDeployments) | **Get** /registry/{registry_ref}/+/artifact/{artifact}/version/{version}/deploymentdetails | Describe Artifact Deployments
[**GetArtifactStats**](ArtifactsApi.md#GetArtifactStats) | **Get** /registry/{registry_ref}/+/artifact/{artifact}/stats | Get Artifact Stats
[**GetArtifactStatsForRegistry**](ArtifactsApi.md#GetArtifactStatsForRegistry) | **Get** /registry/{registry_ref}/+/artifact/stats | Get Artifact Stats
[**GetArtifactStatsForSpace**](ArtifactsApi.md#GetArtifactStatsForSpace) | **Get** /spaces/{space_ref}/+/artifact/stats | Get Artifact Stats
[**GetArtifactSummary**](ArtifactsApi.md#GetArtifactSummary) | **Get** /registry/{registry_ref}/+/artifact/{artifact}/summary | Get Artifact Summary
[**GetArtifactVersionSummary**](ArtifactsApi.md#GetArtifactVersionSummary) | **Get** /registry/{registry_ref}/+/artifact/{artifact}/version/{version}/summary | Get Artifact Version Summary
[**ListArtifactLabels**](ArtifactsApi.md#ListArtifactLabels) | **Get** /registry/{registry_ref}/+/artifact/labels | List Artifact Labels
[**RedirectHarnessArtifact**](ArtifactsApi.md#RedirectHarnessArtifact) | **Get** /registry/{registry_identifier}/artifact/{artifact}/redirect | Redirect to Harness Artifact Page
[**UpdateArtifactLabels**](ArtifactsApi.md#UpdateArtifactLabels) | **Put** /registry/{registry_ref}/+/artifact/{artifact}/labels | Update Artifact Labels

# **DeleteArtifact**
> InlineResponse200 DeleteArtifact(ctx, registryRef, artifact)
Delete Artifact

Delete Artifact.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **registryRef** | **string**| Unique registry path. | 
  **artifact** | **string**| Name of artifact. | 

### Return type

[**InlineResponse200**](inline_response_200.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteArtifactVersion**
> InlineResponse200 DeleteArtifactVersion(ctx, registryRef, artifact, version)
Delete an Artifact Version

Delete Artifact Version.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **registryRef** | **string**| Unique registry path. | 
  **artifact** | **string**| Name of artifact. | 
  **version** | **string**| Name of Artifact Version. | 

### Return type

[**InlineResponse200**](inline_response_200.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllArtifactVersions**
> InlineResponse20012 GetAllArtifactVersions(ctx, registryRef, artifact, optional)
List Artifact Versions

Lists all the Artifact Versions.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **registryRef** | **string**| Unique registry path. | 
  **artifact** | **string**| Name of artifact. | 
 **optional** | ***ArtifactsApiGetAllArtifactVersionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ArtifactsApiGetAllArtifactVersionsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **optional.Int64**| Current page number | [default to 1]
 **size** | **optional.Int64**| Number of items per page | [default to 20]
 **sortOrder** | **optional.String**| sortOrder | 
 **sortField** | **optional.String**| sortField | 
 **searchTerm** | **optional.String**| search Term. | 

### Return type

[**InlineResponse20012**](inline_response_200_12.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetArtifactDeployments**
> InlineResponse2003 GetArtifactDeployments(ctx, registryRef, artifact, version, optional)
Describe Artifact Deployments

Get Artifact Deployments

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **registryRef** | **string**| Unique registry path. | 
  **artifact** | **string**| Name of artifact. | 
  **version** | **string**| Name of Artifact Version. | 
 **optional** | ***ArtifactsApiGetArtifactDeploymentsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ArtifactsApiGetArtifactDeploymentsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **envType** | **optional.String**| env type | 
 **page** | **optional.Int64**| Current page number | [default to 1]
 **size** | **optional.Int64**| Number of items per page | [default to 20]
 **sortOrder** | **optional.String**| sortOrder | 
 **sortField** | **optional.String**| sortField | 
 **searchTerm** | **optional.String**| search Term. | 

### Return type

[**InlineResponse2003**](inline_response_200_3.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetArtifactStats**
> InlineResponse2002 GetArtifactStats(ctx, registryRef, artifact, optional)
Get Artifact Stats

Get Artifact Stats.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **registryRef** | **string**| Unique registry path. | 
  **artifact** | **string**| Name of artifact. | 
 **optional** | ***ArtifactsApiGetArtifactStatsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ArtifactsApiGetArtifactStatsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **from** | **optional.String**| Date. Format - MM/DD/YYYY | 
 **to** | **optional.String**| Date. Format - MM/DD/YYYY | 

### Return type

[**InlineResponse2002**](inline_response_200_2.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetArtifactStatsForRegistry**
> InlineResponse2002 GetArtifactStatsForRegistry(ctx, registryRef, optional)
Get Artifact Stats

Get Artifact Stats.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **registryRef** | **string**| Unique registry path. | 
 **optional** | ***ArtifactsApiGetArtifactStatsForRegistryOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ArtifactsApiGetArtifactStatsForRegistryOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **from** | **optional.String**| Date. Format - MM/DD/YYYY | 
 **to** | **optional.String**| Date. Format - MM/DD/YYYY | 

### Return type

[**InlineResponse2002**](inline_response_200_2.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetArtifactStatsForSpace**
> InlineResponse2002 GetArtifactStatsForSpace(ctx, spaceRef, optional)
Get Artifact Stats

Get Artifact Stats.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **spaceRef** | **string**| Unique space path. | 
 **optional** | ***ArtifactsApiGetArtifactStatsForSpaceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ArtifactsApiGetArtifactStatsForSpaceOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **from** | **optional.String**| Date. Format - MM/DD/YYYY | 
 **to** | **optional.String**| Date. Format - MM/DD/YYYY | 

### Return type

[**InlineResponse2002**](inline_response_200_2.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetArtifactSummary**
> InlineResponse2001 GetArtifactSummary(ctx, registryRef, artifact)
Get Artifact Summary

Get Artifact Summary.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **registryRef** | **string**| Unique registry path. | 
  **artifact** | **string**| Name of artifact. | 

### Return type

[**InlineResponse2001**](inline_response_200_1.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetArtifactVersionSummary**
> InlineResponse20011 GetArtifactVersionSummary(ctx, registryRef, artifact, version, optional)
Get Artifact Version Summary

Get Artifact Version Summary.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **registryRef** | **string**| Unique registry path. | 
  **artifact** | **string**| Name of artifact. | 
  **version** | **string**| Name of Artifact Version. | 
 **optional** | ***ArtifactsApiGetArtifactVersionSummaryOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ArtifactsApiGetArtifactVersionSummaryOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **digest** | **optional.String**| Digest. | 

### Return type

[**InlineResponse20011**](inline_response_200_11.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListArtifactLabels**
> InlineResponse20013 ListArtifactLabels(ctx, registryRef, optional)
List Artifact Labels

List Artifact Labels.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **registryRef** | **string**| Unique registry path. | 
 **optional** | ***ArtifactsApiListArtifactLabelsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ArtifactsApiListArtifactLabelsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int64**| Current page number | [default to 1]
 **size** | **optional.Int64**| Number of items per page | [default to 20]
 **searchTerm** | **optional.String**| search Term. | 

### Return type

[**InlineResponse20013**](inline_response_200_13.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RedirectHarnessArtifact**
> RedirectHarnessArtifact(ctx, registryIdentifier, artifact, optional)
Redirect to Harness Artifact Page

Redirect to Harness Artifact Page

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **registryIdentifier** | **string**| Unique registry Identifier in a account. | 
  **artifact** | **string**| Name of artifact. | 
 **optional** | ***ArtifactsApiRedirectHarnessArtifactOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ArtifactsApiRedirectHarnessArtifactOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **accountIdentifier** | **optional.String**| Account Identifier | 
 **version** | **optional.String**| Version | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateArtifactLabels**
> InlineResponse2001 UpdateArtifactLabels(ctx, registryRef, artifact, optional)
Update Artifact Labels

Update Artifact Labels.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **registryRef** | **string**| Unique registry path. | 
  **artifact** | **string**| Name of artifact. | 
 **optional** | ***ArtifactsApiUpdateArtifactLabelsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ArtifactsApiUpdateArtifactLabelsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of ArtifactLabelRequest**](ArtifactLabelRequest.md)| request to update artifact labels | 

### Return type

[**InlineResponse2001**](inline_response_200_1.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

