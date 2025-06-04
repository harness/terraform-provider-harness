# V1RollingUpdateStatefulSetStrategy

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxUnavailable** | [***IntstrIntOrString**](intstr.IntOrString.md) |  | [optional] [default to null]
**Partition** | **int32** | Partition indicates the ordinal at which the StatefulSet should be partitioned for updates. During a rolling update, all pods from ordinal Replicas-1 to Partition are updated. All pods from ordinal Partition-1 to 0 remain untouched. This is helpful in being able to do a canary based deployment. The default value is 0. +optional | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

