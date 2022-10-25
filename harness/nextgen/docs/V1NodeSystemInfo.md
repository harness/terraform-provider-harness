# V1NodeSystemInfo

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MachineID** | **string** |  | [optional] [default to null]
**SystemUUID** | **string** |  | [optional] [default to null]
**BootID** | **string** | Boot ID reported by the node. | [optional] [default to null]
**KernelVersion** | **string** | Kernel Version reported by the node from &#x27;uname -r&#x27; (e.g. 3.16.0-0.bpo.4-amd64). | [optional] [default to null]
**OsImage** | **string** | OS Image reported by the node from /etc/os-release (e.g. Debian GNU/Linux 7 (wheezy)). | [optional] [default to null]
**ContainerRuntimeVersion** | **string** | ContainerRuntime Version reported by the node through runtime remote API (e.g. docker://1.5.0). | [optional] [default to null]
**KubeletVersion** | **string** | Kubelet Version reported by the node. | [optional] [default to null]
**KubeProxyVersion** | **string** | KubeProxy Version reported by the node. | [optional] [default to null]
**OperatingSystem** | **string** |  | [optional] [default to null]
**Architecture** | **string** |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

