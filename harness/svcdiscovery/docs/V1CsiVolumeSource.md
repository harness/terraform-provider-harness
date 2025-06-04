# V1CsiVolumeSource

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Driver** | **string** | driver is the name of the CSI driver that handles this volume. Consult with your admin for the correct name as registered in the cluster. | [optional] [default to null]
**FsType** | **string** | fsType to mount. Ex. \&quot;ext4\&quot;, \&quot;xfs\&quot;, \&quot;ntfs\&quot;. If not provided, the empty value is passed to the associated CSI driver which will determine the default filesystem to apply. +optional | [optional] [default to null]
**NodePublishSecretRef** | [***V1LocalObjectReference**](v1.LocalObjectReference.md) |  | [optional] [default to null]
**ReadOnly** | **bool** | readOnly specifies a read-only configuration for the volume. Defaults to false (read/write). +optional | [optional] [default to null]
**VolumeAttributes** | **map[string]string** | volumeAttributes stores driver-specific properties that are passed to the CSI driver. Consult your driver&#x27;s documentation for supported values. +optional | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

