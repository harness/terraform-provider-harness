# InputsetInputSet

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountID** | **string** |  | [default to null]
**CreatedAt** | **int32** | creation timestamp of the input set | [optional] [default to null]
**CreatedBy** | **string** | user ID of the user who created the input set | [optional] [default to null]
**Description** | **string** | Description of the input set | [optional] [default to null]
**ExperimentID** | **string** | Foreign key to link with experiment | [optional] [default to null]
**Id** | **string** | Mongo ID (primary key) | [optional] [default to null]
**Identity** | **string** | Human readable ID | [optional] [default to null]
**IsRemoved** | **bool** | TODO: this is not needed, and on delete, input set should be deleted from the DB, makes no sense for storing for audit purpose | [optional] [default to null]
**Name** | **string** | Name of the input set | [optional] [default to null]
**OrgID** | **string** |  | [optional] [default to null]
**ProjectID** | **string** |  | [optional] [default to null]
**Spec** | **string** | Type of input set Type string &#x60;bson:\&quot;type\&quot;&#x60; Foreign key to link with probes TODO: not sure if required ProbeID string &#x60;bson:\&quot;probe_id\&quot;&#x60; For fault level variables, key &#x3D; step | [optional] [default to null]
**UpdatedAt** | **int32** | updation timestamp of the input set | [optional] [default to null]
**UpdatedBy** | **string** | user ID of the user who updated the input set | [optional] [default to null]
**Version** | **string** | Version | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

