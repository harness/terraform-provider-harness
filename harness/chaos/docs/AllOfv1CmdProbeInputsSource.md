# AllOfv1CmdProbeInputsSource

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Annotations** | **map[string]string** | Annotations for the source pod | [optional] [default to null]
**Args** | **[]string** | Args for the source pod | [optional] [default to null]
**Command** | **[]string** | Command for the source pod | [optional] [default to null]
**Env** | [**[]V1EnvVar**](v1.EnvVar.md) | ENVList contains ENV passed to the source pod | [optional] [default to null]
**HostNetwork** | **bool** | HostNetwork define the hostNetwork of the external pod it supports boolean values and default value is false | [optional] [default to null]
**Image** | **string** | Image for the source pod | [optional] [default to null]
**ImagePullPolicy** | [***interface{}**](interface{}.md) | ImagePullPolicy for the source pod | [optional] [default to null]
**ImagePullSecrets** | [**[]V1LocalObjectReference**](v1.LocalObjectReference.md) | ImagePullSecrets for source pod | [optional] [default to null]
**InheritInputs** | **bool** | InheritInputs defined to inherit experiment pod attributes(ENV, volumes, and volumeMounts) into probe pod it supports boolean values and default value is false | [optional] [default to null]
**Labels** | **map[string]string** | Labels for the source pod | [optional] [default to null]
**NodeSelector** | **map[string]string** | NodeSelector for the source pod | [optional] [default to null]
**Privileged** | **bool** | Privileged for the source pod | [optional] [default to null]
**Tolerations** | [**[]V1Toleration**](v1.Toleration.md) |  | [optional] [default to null]
**VolumeMount** | [**[]V1VolumeMount**](v1.VolumeMount.md) | VolumesMount for the source pod | [optional] [default to null]
**Volumes** | [**[]V1Volume**](v1.Volume.md) | Volumes for the source pod | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

