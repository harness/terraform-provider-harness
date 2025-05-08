# ApplicationsPullRequestGeneratorGitLab

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Project** | **string** | GitLab project to scan. Required. | [optional] [default to null]
**Api** | **string** | The GitLab API URL to talk to. If blank, uses https://gitlab.com/. | [optional] [default to null]
**TokenRef** | [***ApplicationsSecretRef**](applicationsSecretRef.md) |  | [optional] [default to null]
**Labels** | **[]string** |  | [optional] [default to null]
**PullRequestState** | **string** |  | [optional] [default to null]
**Insecure** | **bool** |  | [optional] [default to null]
**CaRef** | [***ApplicationsConfigMapKeyRef**](applicationsConfigMapKeyRef.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

