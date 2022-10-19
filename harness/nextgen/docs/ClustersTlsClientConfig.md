# ClustersTlsClientConfig

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Insecure** | **bool** | Insecure specifies that the server should be accessed without verifying the TLS certificate. For testing only. | [optional] [default to null]
**ServerName** | **string** | ServerName is passed to the server for SNI and is used in the client to check server certificates against. If ServerName is empty, the hostname used to contact the server is used. | [optional] [default to null]
**CertData** | **string** |  | [optional] [default to null]
**KeyData** | **string** |  | [optional] [default to null]
**CaData** | **string** |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

