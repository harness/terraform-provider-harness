# ApplicationsApplicationSourceHelm

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ValueFiles** | **[]string** |  | [optional] [default to null]
**Parameters** | [**[]ApplicationsHelmParameter**](applicationsHelmParameter.md) |  | [optional] [default to null]
**ReleaseName** | **string** |  | [optional] [default to null]
**Values** | **string** |  | [optional] [default to null]
**FileParameters** | [**[]ApplicationsHelmFileParameter**](applicationsHelmFileParameter.md) |  | [optional] [default to null]
**Version** | **string** |  | [optional] [default to null]
**PassCredentials** | **bool** |  | [optional] [default to null]
**IgnoreMissingValueFiles** | **bool** |  | [optional] [default to null]
**SkipCrds** | **bool** |  | [optional] [default to null]
**ValuesObject** | [***interface{}**](interface{}.md) |  | [optional] [default to null]
**Namespace** | **string** | Namespace is an optional namespace to template with. If left empty, defaults to the app&#x27;s destination namespace. | [optional] [default to null]
**KubeVersion** | **string** | KubeVersion specifies the Kubernetes API version to pass to Helm when templating manifests. By default, Argo CD uses the Kubernetes version of the target cluster. | [optional] [default to null]
**ApiVersions** | **[]string** | APIVersions specifies the Kubernetes resource API versions to pass to Helm when templating manifests. By default, Argo CD uses the API versions of the target cluster. The format is [group/]version/kind. | [optional] [default to null]
**SkipTests** | **bool** | SkipTests skips test manifest installation step (Helm&#x27;s --skip-tests). | [optional] [default to null]
**SkipSchemaValidation** | **bool** |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

