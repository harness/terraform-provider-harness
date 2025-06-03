# ChaosactiontemplateChaosActionTemplate

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountID** | **string** |  | [default to null]
**ActionProperties** | [***AllOfchaosactiontemplateChaosActionTemplateActionProperties**](AllOfchaosactiontemplateChaosActionTemplateActionProperties.md) | Needed for API response *not to be stored in DB* | [optional] [default to null]
**CreatedAt** | **int32** |  | [optional] [default to null]
**CreatedBy** | **string** |  | [optional] [default to null]
**Description** | **string** |  | [optional] [default to null]
**HubRef** | **string** |  | [optional] [default to null]
**Id** | **string** |  | [optional] [default to null]
**Identity** | **string** | Unique identifier (human-readable) immutable Initially it will be same as name | [optional] [default to null]
**InfrastructureType** | [***ActionsInfrastructureType**](actions.InfrastructureType.md) |  | [optional] [default to null]
**IsDefault** | **bool** | isDefault indicates if it is the default version for predefined faults, latest should be set as default | [optional] [default to null]
**IsRemoved** | **bool** |  | [default to null]
**Name** | **string** | Fault name to sync the changes from the hub HubRef + Name should be unique | [default to null]
**OrgID** | **string** |  | [optional] [default to null]
**ProjectID** | **string** |  | [optional] [default to null]
**Revision** | **int32** | it increments every time a new version of fault is published | [optional] [default to null]
**RunProperties** | [***ActionActionTemplateRunProperties**](action.ActionTemplateRunProperties.md) |  | [optional] [default to null]
**Tags** | **[]string** |  | [optional] [default to null]
**Template** | **string** |  | [optional] [default to null]
**Type_** | [***ActionsActionType**](actions.ActionType.md) |  | [optional] [default to null]
**UpdatedAt** | **int32** |  | [optional] [default to null]
**UpdatedBy** | **string** |  | [optional] [default to null]
**Variables** | [**[]TemplateVariable**](template.Variable.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

