# DatabaseKubernetesAgentConfiguration

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Annotations** | **map[string]string** |  | [optional] [default to null]
**DisableNamespaceCreation** | **bool** |  | [optional] [default to null]
**ImagePullPolicy** | **string** |  | [optional] [default to null]
**Labels** | **map[string]string** |  | [optional] [default to null]
**Namespace** | **string** |  | [optional] [default to null]
**Namespaced** | **bool** |  | [optional] [default to null]
**NodeSelector** | **map[string]string** |  | [optional] [default to null]
**Resources** | [***DatabaseResourceRequirements**](database.ResourceRequirements.md) |  | [optional] [default to null]
**RunAsGroup** | **int32** |  | [optional] [default to null]
**RunAsUser** | **int32** |  | [optional] [default to null]
**ServiceAccount** | **string** |  | [optional] [default to null]
**Tolerations** | [**[]V1Toleration**](v1.Toleration.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

