# V1PersistentVolumeClaimSpec

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessModes** | **[]string** | accessModes contains the desired access modes the volume should have. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes-1 +optional | [optional] [default to null]
**DataSource** | [***V1TypedLocalObjectReference**](v1.TypedLocalObjectReference.md) |  | [optional] [default to null]
**DataSourceRef** | [***V1TypedLocalObjectReference**](v1.TypedLocalObjectReference.md) |  | [optional] [default to null]
**Resources** | [***V1ResourceRequirements**](v1.ResourceRequirements.md) |  | [optional] [default to null]
**Selector** | [***V1LabelSelector**](v1.LabelSelector.md) |  | [optional] [default to null]
**StorageClassName** | **string** | storageClassName is the name of the StorageClass required by the claim. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#class-1 +optional | [optional] [default to null]
**VolumeMode** | **string** | volumeMode defines what type of volume is required by the claim. Value of Filesystem is implied when not included in claim spec. +optional | [optional] [default to null]
**VolumeName** | **string** | volumeName is the binding reference to the PersistentVolume backing this claim. +optional | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

