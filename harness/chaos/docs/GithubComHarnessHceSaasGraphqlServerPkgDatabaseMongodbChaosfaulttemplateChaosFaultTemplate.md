# GithubComHarnessHceSaasGraphqlServerPkgDatabaseMongodbChaosfaulttemplateChaosFaultTemplate

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountID** | **string** |  | [default to null]
**CreatedAt** | **int32** | creation timestamp of the revision | [optional] [default to null]
**CreatedBy** | **string** | user ID of the user who created the revision | [optional] [default to null]
**HubRef** | **string** | Hub reference to sync the changes whenever there are changes in the faults | [optional] [default to null]
**Id** | **string** | Mongo ID (primary key) | [optional] [default to null]
**Identity** | **string** | Unique identifier (human-readable) immutable Initially it will be same as name | [optional] [default to null]
**IsDefault** | **bool** | isDefault indicates if it is the default version for predefined faults, latest should be set as default | [optional] [default to null]
**IsRemoved** | **bool** | isRemoved indicates if the document is deleted | [optional] [default to null]
**Name** | **string** | Fault name to sync the changes from the hub HubRef + Name should be unique | [optional] [default to null]
**OrgID** | **string** |  | [optional] [default to null]
**ProjectID** | **string** |  | [optional] [default to null]
**Revision** | **int32** | Revision is the version of fault template, it increments every time a new version of fault is published | [optional] [default to null]
**Template** | **string** | template of the fault | [optional] [default to null]
**Variables** | [**[]TemplateVariable**](template.Variable.md) | Variables for template | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

