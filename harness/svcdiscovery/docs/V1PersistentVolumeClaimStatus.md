# V1PersistentVolumeClaimStatus

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessModes** | **[]string** | accessModes contains the actual access modes the volume backing the PVC has. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes-1 +optional | [optional] [default to null]
**AllocatedResources** | [***map[string]ResourceQuantity**](map.md) |  | [optional] [default to null]
**Capacity** | [***map[string]ResourceQuantity**](map.md) |  | [optional] [default to null]
**Conditions** | [**[]V1PersistentVolumeClaimCondition**](v1.PersistentVolumeClaimCondition.md) | conditions is the current Condition of persistent volume claim. If underlying persistent volume is being resized then the Condition will be set to &#x27;ResizeStarted&#x27;. +optional +patchMergeKey&#x3D;type +patchStrategy&#x3D;merge | [optional] [default to null]
**Phase** | **string** | phase represents the current phase of PersistentVolumeClaim. +optional | [optional] [default to null]
**ResizeStatus** | **string** | resizeStatus stores status of resize operation. ResizeStatus is not set by default but when expansion is complete resizeStatus is set to empty string by resize controller or kubelet. This is an alpha field and requires enabling RecoverVolumeExpansionFailure feature. +featureGate&#x3D;RecoverVolumeExpansionFailure +optional | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

