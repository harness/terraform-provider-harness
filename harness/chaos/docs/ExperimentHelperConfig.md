# ExperimentHelperConfig

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Annotations** | **map[string]string** |  | [optional] [default to null]
**Args** | **[]string** |  | [optional] [default to null]
**Command** | **[]string** |  | [optional] [default to null]
**ConfigMaps** | [**[]ExperimentConfigMap**](experiment.ConfigMap.md) |  | [optional] [default to null]
**Env** | [**[]V1EnvVar**](v1.EnvVar.md) |  | [optional] [default to null]
**HostFileVolumes** | [**[]ExperimentHostFile**](experiment.HostFile.md) |  | [optional] [default to null]
**HostIPC** | **bool** |  | [optional] [default to null]
**HostNetwork** | **bool** |  | [optional] [default to null]
**HostPID** | **bool** |  | [optional] [default to null]
**Image** | **string** |  | [optional] [default to null]
**ImagePullPolicy** | [***V1PullPolicy**](v1.PullPolicy.md) |  | [optional] [default to null]
**ImagePullSecrets** | [**[]V1LocalObjectReference**](v1.LocalObjectReference.md) |  | [optional] [default to null]
**Labels** | **map[string]string** |  | [optional] [default to null]
**NodeSelector** | **map[string]string** |  | [optional] [default to null]
**Resources** | [***V1ResourceRequirements**](v1.ResourceRequirements.md) |  | [optional] [default to null]
**Secrets** | [**[]ExperimentSecret**](experiment.Secret.md) |  | [optional] [default to null]
**SecurityContext** | [***ExperimentSecurityContext**](experiment.SecurityContext.md) |  | [optional] [default to null]
**Tolerations** | [**[]V1Toleration**](v1.Toleration.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

