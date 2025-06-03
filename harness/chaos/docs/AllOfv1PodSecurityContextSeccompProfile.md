# AllOfv1PodSecurityContextSeccompProfile

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LocalhostProfile** | **string** | localhostProfile indicates a profile defined in a file on the node should be used. The profile must be preconfigured on the node to work. Must be a descending path, relative to the kubelet&#x27;s configured seccomp profile location. Must only be set if type is \&quot;Localhost\&quot;. +optional | [optional] [default to null]
**Type_** | [***interface{}**](interface{}.md) | type indicates which kind of seccomp profile will be applied. Valid options are:  Localhost - a profile defined in a file on the node should be used. RuntimeDefault - the container runtime default profile should be used. Unconfined - no profile should be applied. +unionDiscriminator | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

