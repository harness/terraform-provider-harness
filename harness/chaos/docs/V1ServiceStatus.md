# V1ServiceStatus

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Conditions** | [**[]V1Condition**](v1.Condition.md) | Current service state +optional +patchMergeKey&#x3D;type +patchStrategy&#x3D;merge +listType&#x3D;map +listMapKey&#x3D;type | [optional] [default to null]
**LoadBalancer** | [***AllOfv1ServiceStatusLoadBalancer**](AllOfv1ServiceStatusLoadBalancer.md) | LoadBalancer contains the current status of the load-balancer, if one is present. +optional | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

