# ApplicationsResourceNetworkingInfo

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TargetLabels** | **map[string]string** |  | [optional] [default to null]
**TargetRefs** | [**[]ApplicationsResourceRef**](applicationsResourceRef.md) |  | [optional] [default to null]
**Labels** | **map[string]string** |  | [optional] [default to null]
**Ingress** | [**[]V1LoadBalancerIngress**](v1LoadBalancerIngress.md) |  | [optional] [default to null]
**ExternalURLs** | **[]string** | ExternalURLs holds list of URLs which should be available externally. List is populated for ingress resources using rules hostnames. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

