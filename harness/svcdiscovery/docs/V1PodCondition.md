# V1PodCondition

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastProbeTime** | **string** | Last time we probed the condition. +optional | [optional] [default to null]
**LastTransitionTime** | **string** | Last time the condition transitioned from one status to another. +optional | [optional] [default to null]
**Message** | **string** | Human-readable message indicating details about last transition. +optional | [optional] [default to null]
**Reason** | **string** | Unique, one-word, CamelCase reason for the condition&#x27;s last transition. +optional | [optional] [default to null]
**Status** | **string** | Status is the status of the condition. Can be True, False, Unknown. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#pod-conditions | [optional] [default to null]
**Type_** | **string** | Type is the type of the condition. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#pod-conditions | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

