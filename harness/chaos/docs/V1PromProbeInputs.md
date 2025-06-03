# V1PromProbeInputs

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Auth** | [***AllOfv1PromProbeInputsAuth**](AllOfv1PromProbeInputsAuth.md) | Auth contains the authentication details for the prometheus probe | [optional] [default to null]
**Comparator** | [***AllOfv1PromProbeInputsComparator**](AllOfv1PromProbeInputsComparator.md) | Comparator check for the correctness of the probe output | [optional] [default to null]
**Endpoint** | **string** | Endpoint for the prometheus probe | [optional] [default to null]
**Query** | **string** | Query to get prometheus metrics | [optional] [default to null]
**QueryPath** | **string** | QueryPath contains filePath, which contains prometheus query | [optional] [default to null]
**TlsConfig** | [***AllOfv1PromProbeInputsTlsConfig**](AllOfv1PromProbeInputsTlsConfig.md) | TLSConfig contains the tls configuration for the prometheus probe | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

