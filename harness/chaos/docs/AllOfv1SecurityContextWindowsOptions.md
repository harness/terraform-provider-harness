# AllOfv1SecurityContextWindowsOptions

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GmsaCredentialSpec** | **string** | GMSACredentialSpec is where the GMSA admission webhook (https://github.com/kubernetes-sigs/windows-gmsa) inlines the contents of the GMSA credential spec named by the GMSACredentialSpecName field. +optional | [optional] [default to null]
**GmsaCredentialSpecName** | **string** | GMSACredentialSpecName is the name of the GMSA credential spec to use. +optional | [optional] [default to null]
**RunAsUserName** | **string** | The UserName in Windows to run the entrypoint of the container process. Defaults to the user specified in image metadata if unspecified. May also be set in PodSecurityContext. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence. +optional | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

