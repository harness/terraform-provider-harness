# ClustersClusterConfig

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | **string** |  | [optional] [default to null]
**Password** | **string** |  | [optional] [default to null]
**BearerToken** | **string** | Server requires Bearer authentication. This client will not attempt to use refresh tokens for an OAuth2 flow. TODO: demonstrate an OAuth2 compatible client. | [optional] [default to null]
**TlsClientConfig** | [***ClustersTlsClientConfig**](clustersTLSClientConfig.md) |  | [optional] [default to null]
**AwsAuthConfig** | [***ClustersAwsAuthConfig**] | (deprecated) | [optional] [default to null]
**RoleARN** | **string** | RoleARN contains optional role ARN. If set then AWS IAM Authenticator assumes a role to perform cluster operations instead of the default AWS credential provider chain. | [optional] [default to null]
**AwsClusterName** | **string** | AWS Cluster name. If set then AWS CLI EKS token command will be used to access cluster. | [optional] [default to null]
**ExecProviderConfig** | [***ClustersExecProviderConfig**](clustersExecProviderConfig.md) |  | [optional] [default to null]
**ClusterConnectionType** | **string** |  | [optional] [default to null]
**DisableCompression** | **bool** | DisableCompression bypasses automatic GZip compression requests to the server. | [optional] [default to null]
**ProxyUrl** | **string** |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

