# K8sfaultChaosSpec

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Annotations** | **map[string]string** |  | [optional] [default to null]
**Args** | **[]string** |  | [optional] [default to null]
**Command** | **[]string** |  | [optional] [default to null]
**ConfigMapVolume** | [**[]K8sfaultConfigMapVolume**](k8sfault.ConfigMapVolume.md) |  | [optional] [default to null]
**ContainerSecurityContext** | [***V1SecurityContext**](v1.SecurityContext.md) |  | [optional] [default to null]
**Env** | [**[]V1EnvVar**](v1.EnvVar.md) |  | [optional] [default to null]
**FaultName** | **string** |  | [optional] [default to null]
**HostIPC** | **bool** |  | [optional] [default to null]
**HostNetwork** | **bool** |  | [optional] [default to null]
**HostPID** | **bool** |  | [optional] [default to null]
**HostPathVolume** | [**[]K8sfaultHostPathVolume**](k8sfault.HostPathVolume.md) |  | [optional] [default to null]
**Image** | **string** |  | [optional] [default to null]
**ImagePullPolicy** | [***V1PullPolicy**](v1.PullPolicy.md) |  | [optional] [default to null]
**ImagePullSecrets** | **[]string** |  | [optional] [default to null]
**Labels** | **map[string]string** |  | [optional] [default to null]
**NodeSelector** | **map[string]string** |  | [optional] [default to null]
**Params** | [**[]K8sfaultChaosParameter**](k8sfault.ChaosParameter.md) |  | [optional] [default to null]
**PodSecurityContext** | [***V1PodSecurityContext**](v1.PodSecurityContext.md) |  | [optional] [default to null]
**ResourceRequirements** | [***K8sfaultResourceRequirements**](k8sfault.ResourceRequirements.md) |  | [optional] [default to null]
**SecretVolume** | [**[]K8sfaultSecretVolume**](k8sfault.SecretVolume.md) |  | [optional] [default to null]
**Toleration** | [***V1Toleration**](v1.Toleration.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

