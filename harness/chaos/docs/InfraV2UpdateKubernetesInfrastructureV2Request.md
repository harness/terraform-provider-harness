# InfraV2UpdateKubernetesInfrastructureV2Request

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AiEnabled** | **bool** |  | [optional] [default to null]
**Annotation** | **map[string]string** |  | [optional] [default to null]
**Containers** | **string** |  | [optional] [default to null]
**CorrelationId** | **string** |  | [optional] [default to null]
**Description** | **string** |  | [optional] [default to null]
**Env** | [**[]InfraV2Env**](infra_v2.Env.md) |  | [optional] [default to null]
**EnvironmentID** | **string** |  | [optional] [default to null]
**Identity** | **string** |  | [optional] [default to null]
**ImageRegistry** | [***ImageRegistryImageRegistryV2**](image_registry.ImageRegistryV2.md) |  | [optional] [default to null]
**InfraNamespace** | **string** |  | [optional] [default to null]
**InsecureSkipVerify** | **bool** |  | [optional] [default to null]
**Label** | **map[string]string** |  | [optional] [default to null]
**Mtls** | [***InfraV2MtlsConfiguration**](infra_v2.MTLSConfiguration.md) |  | [optional] [default to null]
**Name** | **string** |  | [optional] [default to null]
**NodeSelector** | **map[string]string** |  | [optional] [default to null]
**Proxy** | [***InfraV2ProxyConfiguration**](infra_v2.ProxyConfiguration.md) |  | [optional] [default to null]
**RunAsGroup** | **int32** |  | [optional] [default to null]
**RunAsUser** | **int32** |  | [optional] [default to null]
**ServiceAccount** | **string** |  | [optional] [default to null]
**Tags** | **[]string** |  | [optional] [default to null]
**Tolerations** | [**[]V1Toleration**](v1.Toleration.md) |  | [optional] [default to null]
**VolumeMounts** | [**[]V1VolumeMount**](v1.VolumeMount.md) |  | [optional] [default to null]
**Volumes** | [**[]InfraV2Volumes**](infra_v2.Volumes.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

