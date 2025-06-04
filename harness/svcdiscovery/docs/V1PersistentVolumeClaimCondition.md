# V1PersistentVolumeClaimCondition

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastProbeTime** | **string** | lastProbeTime is the time we probed the condition. +optional | [optional] [default to null]
**LastTransitionTime** | **string** | lastTransitionTime is the time the condition transitioned from one status to another. +optional | [optional] [default to null]
**Message** | **string** | message is the human-readable message indicating details about last transition. +optional | [optional] [default to null]
**Reason** | **string** | reason is a unique, this should be a short, machine understandable string that gives the reason for condition&#x27;s last transition. If it reports \&quot;ResizeStarted\&quot; that means the underlying persistent volume is being resized. +optional | [optional] [default to null]
**Status** | **string** |  | [optional] [default to null]
**Type_** | **string** |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

