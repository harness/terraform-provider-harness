# V1VolumeMount

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MountPath** | **string** | Path within the container at which the volume should be mounted.  Must not contain &#x27;:&#x27;. | [optional] [default to null]
**MountPropagation** | **string** | mountPropagation determines how mounts are propagated from the host to container and the other way around. When not set, MountPropagationNone is used. This field is beta in 1.10. +optional | [optional] [default to null]
**Name** | **string** | This must match the Name of a Volume. | [optional] [default to null]
**ReadOnly** | **bool** | Mounted read-only if true, read-write otherwise (false or unspecified). Defaults to false. +optional | [optional] [default to null]
**SubPath** | **string** | Path within the volume from which the container&#x27;s volume should be mounted. Defaults to \&quot;\&quot; (volume&#x27;s root). +optional | [optional] [default to null]
**SubPathExpr** | **string** | Expanded path within the volume from which the container&#x27;s volume should be mounted. Behaves similarly to SubPath but environment variable references $(VAR_NAME) are expanded using the container&#x27;s environment. Defaults to \&quot;\&quot; (volume&#x27;s root). SubPathExpr and SubPath are mutually exclusive. +optional | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

