# V1PersistentVolumeClaimSpec

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessModes** | [**[]V1PersistentVolumeAccessMode**](v1.PersistentVolumeAccessMode.md) | AccessModes contains the desired access modes the volume should have. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes-1 +optional | [optional] [default to null]
**DataSource** | [***AllOfv1PersistentVolumeClaimSpecDataSource**](AllOfv1PersistentVolumeClaimSpecDataSource.md) | This field can be used to specify either: * An existing VolumeSnapshot object (snapshot.storage.k8s.io/VolumeSnapshot) * An existing PVC (PersistentVolumeClaim) * An existing custom resource that implements data population (Alpha) In order to use custom resource types that implement data population, the AnyVolumeDataSource feature gate must be enabled. If the provisioner or an external controller can support the specified data source, it will create a new volume based on the contents of the specified data source. +optional | [optional] [default to null]
**Resources** | [***AllOfv1PersistentVolumeClaimSpecResources**](AllOfv1PersistentVolumeClaimSpecResources.md) | Resources represents the minimum resources the volume should have. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#resources +optional | [optional] [default to null]
**Selector** | [***AllOfv1PersistentVolumeClaimSpecSelector**](AllOfv1PersistentVolumeClaimSpecSelector.md) | A label query over volumes to consider for binding. +optional | [optional] [default to null]
**StorageClassName** | **string** | Name of the StorageClass required by the claim. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#class-1 +optional | [optional] [default to null]
**VolumeMode** | [***AllOfv1PersistentVolumeClaimSpecVolumeMode**](AllOfv1PersistentVolumeClaimSpecVolumeMode.md) | volumeMode defines what type of volume is required by the claim. Value of Filesystem is implied when not included in claim spec. +optional | [optional] [default to null]
**VolumeName** | **string** | VolumeName is the binding reference to the PersistentVolume backing this claim. +optional | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

