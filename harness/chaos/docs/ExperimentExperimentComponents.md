# ExperimentExperimentComponents

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfigMaps** | [**[]ExperimentConfigMap**](experiment.ConfigMap.md) |  | [optional] [default to null]
**Helper** | [***ExperimentHelperConfig**](experiment.HelperConfig.md) |  | [optional] [default to null]
**HostFileVolumes** | [**[]ExperimentHostFile**](experiment.HostFile.md) |  | [optional] [default to null]
**HostPID** | **bool** |  | [optional] [default to null]
**ImagePullSecrets** | [**[]V1LocalObjectReference**](v1.LocalObjectReference.md) |  | [optional] [default to null]
**NodeSelector** | **map[string]string** |  | [optional] [default to null]
**ProjectedVolumes** | [**[]ExperimentProjectedVolumes**](experiment.ProjectedVolumes.md) |  | [optional] [default to null]
**Resources** | [***V1ResourceRequirements**](v1.ResourceRequirements.md) |  | [optional] [default to null]
**Secrets** | [**[]ExperimentSecret**](experiment.Secret.md) |  | [optional] [default to null]
**SecurityContext** | [***ExperimentSecurityContext**](experiment.SecurityContext.md) |  | [optional] [default to null]
**Sidecar** | [**[]ExperimentSidecar**](experiment.Sidecar.md) |  | [optional] [default to null]
**StatusCheckTimeouts** | [***ExperimentStatusCheckTimeout**](experiment.StatusCheckTimeout.md) |  | [optional] [default to null]
**TerminationGracePeriodSeconds** | **int32** |  | [optional] [default to null]
**Tolerations** | [**[]V1Toleration**](v1.Toleration.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

