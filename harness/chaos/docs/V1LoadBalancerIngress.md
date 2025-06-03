# V1LoadBalancerIngress

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hostname** | **string** | Hostname is set for load-balancer ingress points that are DNS based (typically AWS load-balancers) +optional | [optional] [default to null]
**Ip** | **string** | IP is set for load-balancer ingress points that are IP based (typically GCE or OpenStack load-balancers) +optional | [optional] [default to null]
**Ports** | [**[]V1PortStatus**](v1.PortStatus.md) | Ports is a list of records of service ports If used, every port defined in the service should have an entry in it +listType&#x3D;atomic +optional | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

