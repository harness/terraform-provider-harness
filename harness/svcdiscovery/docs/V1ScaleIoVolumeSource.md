# V1ScaleIoVolumeSource

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FsType** | **string** | fsType is the filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. \&quot;ext4\&quot;, \&quot;xfs\&quot;, \&quot;ntfs\&quot;. Default is \&quot;xfs\&quot;. +optional | [optional] [default to null]
**Gateway** | **string** | gateway is the host address of the ScaleIO API Gateway. | [optional] [default to null]
**ProtectionDomain** | **string** | protectionDomain is the name of the ScaleIO Protection Domain for the configured storage. +optional | [optional] [default to null]
**ReadOnly** | **bool** | readOnly Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts. +optional | [optional] [default to null]
**SecretRef** | [***V1LocalObjectReference**](v1.LocalObjectReference.md) |  | [optional] [default to null]
**SslEnabled** | **bool** | sslEnabled Flag enable/disable SSL communication with Gateway, default false +optional | [optional] [default to null]
**StorageMode** | **string** | storageMode indicates whether the storage for a volume should be ThickProvisioned or ThinProvisioned. Default is ThinProvisioned. +optional | [optional] [default to null]
**StoragePool** | **string** | storagePool is the ScaleIO Storage Pool associated with the protection domain. +optional | [optional] [default to null]
**System** | **string** | system is the name of the storage system as configured in ScaleIO. | [optional] [default to null]
**VolumeName** | **string** | volumeName is the name of a volume already created in the ScaleIO system that is associated with this volume source. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

