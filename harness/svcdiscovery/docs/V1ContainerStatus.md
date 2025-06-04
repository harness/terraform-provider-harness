# V1ContainerStatus

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContainerID** | **string** | Container&#x27;s ID in the format &#x27;&lt;type&gt;://&lt;container_id&gt;&#x27;. +optional | [optional] [default to null]
**Image** | **string** | The image the container is running. More info: https://kubernetes.io/docs/concepts/containers/images. | [optional] [default to null]
**ImageID** | **string** | ImageID of the container&#x27;s image. | [optional] [default to null]
**LastState** | [***V1ContainerState**](v1.ContainerState.md) |  | [optional] [default to null]
**Name** | **string** | This must be a DNS_LABEL. Each container in a pod must have a unique name. Cannot be updated. | [optional] [default to null]
**Ready** | **bool** | Specifies whether the container has passed its readiness probe. | [optional] [default to null]
**RestartCount** | **int32** | The number of times the container has been restarted. | [optional] [default to null]
**Started** | **bool** | Specifies whether the container has passed its startup probe. Initialized as false, becomes true after startupProbe is considered successful. Resets to false when the container is restarted, or if kubelet loses state temporarily. Is always true when no startupProbe is defined. +optional | [optional] [default to null]
**State** | [***V1ContainerState**](v1.ContainerState.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

