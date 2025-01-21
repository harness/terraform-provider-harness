# ApplicationsApplicationSourceKustomize

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NamePrefix** | **string** |  | [optional] [default to null]
**NameSuffix** | **string** |  | [optional] [default to null]
**Images** | **[]string** |  | [optional] [default to null]
**CommonLabels** | **map[string]string** |  | [optional] [default to null]
**Version** | **string** |  | [optional] [default to null]
**CommonAnnotations** | **map[string]string** |  | [optional] [default to null]
**ForceCommonLabels** | **bool** |  | [optional] [default to null]
**ForceCommonAnnotations** | **bool** |  | [optional] [default to null]
**Namespace** | **string** |  | [optional] [default to null]
**Replicas** | [**[]ApplicationsKustomizeReplicas**](applicationsKustomizeReplicas.md) |  | [optional] [default to null]
**Patches** | [**[]ApplicationsKustomizePatch**](applicationsKustomizePatch.md) |  | [optional] [default to null]
**Components** | **[]string** |  | [optional] [default to null]
**LabelWithoutSelector** | **bool** |  | [optional] [default to null]
**KubeVersion** | **string** | KubeVersion specifies the Kubernetes API version to pass to Helm when templating manifests. By default, Argo CD uses the Kubernetes version of the target cluster. | [optional] [default to null]
**ApiVersions** | **[]string** | APIVersions specifies the Kubernetes resource API versions to pass to Helm when templating manifests. By default, Argo CD uses the API versions of the target cluster. The format is [group/]version/kind. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

