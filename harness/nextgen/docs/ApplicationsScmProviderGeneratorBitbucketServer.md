# ApplicationsScmProviderGeneratorBitbucketServer

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Project** | **string** | Project to scan. Required. | [optional] [default to null]
**Api** | **string** | The Bitbucket Server REST API URL to talk to. Required. | [optional] [default to null]
**BasicAuth** | [***ApplicationsBasicAuthBitbucketServer**](applicationsBasicAuthBitbucketServer.md) |  | [optional] [default to null]
**AllBranches** | **bool** | Scan all branches instead of just the default branch. | [optional] [default to null]
**BearerToken** | [***ApplicationsBearerTokenBitbucket**](applicationsBearerTokenBitbucket.md) |  | [optional] [default to null]
**Insecure** | **bool** |  | [optional] [default to null]
**CaRef** | [***ApplicationsConfigMapKeyRef**](applicationsConfigMapKeyRef.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

