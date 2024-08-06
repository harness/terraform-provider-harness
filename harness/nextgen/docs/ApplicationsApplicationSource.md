# ApplicationsApplicationSource

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RepoURL** | **string** |  | [optional] [default to null]
**Path** | **string** | Path is a directory path within the Git repository, and is only valid for applications sourced from Git. | [optional] [default to null]
**TargetRevision** | **string** | TargetRevision defines the revision of the source to sync the application to. In case of Git, this can be commit, tag, or branch. If omitted, will equal to HEAD. In case of Helm, this is a semver tag for the Chart&#x27;s version. | [optional] [default to null]
**Helm** | [***ApplicationsApplicationSourceHelm**](applicationsApplicationSourceHelm.md) |  | [optional] [default to null]
**Kustomize** | [***ApplicationsApplicationSourceKustomize**](applicationsApplicationSourceKustomize.md) |  | [optional] [default to null]
**Ksonnet** | [***ApplicationsApplicationSourceKsonnet**](applicationsApplicationSourceKsonnet.md) |  | [optional] [default to null]
**Directory** | [***ApplicationsApplicationSourceDirectory**](applicationsApplicationSourceDirectory.md) |  | [optional] [default to null]
**Plugin** | [***ApplicationsApplicationSourcePlugin**](applicationsApplicationSourcePlugin.md) |  | [optional] [default to null]
**Chart** | **string** | Chart is a Helm chart name, and must be specified for applications sourced from a Helm repo. | [optional] [default to null]
**Ref** | **string** | Ref is reference to another source within sources field. This field will not be used if used with a &#x60;source&#x60; tag. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

