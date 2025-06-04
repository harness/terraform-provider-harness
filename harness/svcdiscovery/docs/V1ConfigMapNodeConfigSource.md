# V1ConfigMapNodeConfigSource

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**KubeletConfigKey** | **string** | KubeletConfigKey declares which key of the referenced ConfigMap corresponds to the KubeletConfiguration structure This field is required in all cases. | [optional] [default to null]
**Name** | **string** | Name is the metadata.name of the referenced ConfigMap. This field is required in all cases. | [optional] [default to null]
**Namespace** | **string** | Namespace is the metadata.namespace of the referenced ConfigMap. This field is required in all cases. | [optional] [default to null]
**ResourceVersion** | **string** | ResourceVersion is the metadata.ResourceVersion of the referenced ConfigMap. This field is forbidden in Node.Spec, and required in Node.Status. +optional | [optional] [default to null]
**Uid** | **string** | UID is the metadata.UID of the referenced ConfigMap. This field is forbidden in Node.Spec, and required in Node.Status. +optional | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

