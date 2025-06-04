# V1Probe

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Exec** | [***V1ExecAction**](v1.ExecAction.md) |  | [optional] [default to null]
**FailureThreshold** | **int32** | Minimum consecutive failures for the probe to be considered failed after having succeeded. Defaults to 3. Minimum value is 1. +optional | [optional] [default to null]
**Grpc** | [***V1GrpcAction**](v1.GRPCAction.md) |  | [optional] [default to null]
**HttpGet** | [***V1HttpGetAction**](v1.HTTPGetAction.md) |  | [optional] [default to null]
**InitialDelaySeconds** | **int32** | Number of seconds after the container has started before liveness probes are initiated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes +optional | [optional] [default to null]
**PeriodSeconds** | **int32** | How often (in seconds) to perform the probe. Default to 10 seconds. Minimum value is 1. +optional | [optional] [default to null]
**SuccessThreshold** | **int32** | Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1. Must be 1 for liveness and startup. Minimum value is 1. +optional | [optional] [default to null]
**TcpSocket** | [***V1TcpSocketAction**](v1.TCPSocketAction.md) |  | [optional] [default to null]
**TerminationGracePeriodSeconds** | **int32** | Optional duration in seconds the pod needs to terminate gracefully upon probe failure. The grace period is the duration in seconds after the processes running in the pod are sent a termination signal and the time when the processes are forcibly halted with a kill signal. Set this value longer than the expected cleanup time for your process. If this value is nil, the pod&#x27;s terminationGracePeriodSeconds will be used. Otherwise, this value overrides the value provided by the pod spec. Value must be non-negative integer. The value zero indicates stop immediately via the kill signal (no opportunity to shut down). This is a beta field and requires enabling ProbeTerminationGracePeriod feature gate. Minimum value is 1. spec.terminationGracePeriodSeconds is used if unset. +optional | [optional] [default to null]
**TimeoutSeconds** | **int32** | Number of seconds after which the probe times out. Defaults to 1 second. Minimum value is 1. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes +optional | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

