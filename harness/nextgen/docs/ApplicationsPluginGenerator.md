# ApplicationsPluginGenerator

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfigMapRef** | [***ApplicationsPluginConfigMapRef**](applicationsPluginConfigMapRef.md) |  | [optional] [default to null]
**Input** | [***ApplicationsPluginInput**](applicationsPluginInput.md) |  | [optional] [default to null]
**RequeueAfterSeconds** | **string** | RequeueAfterSeconds determines how long the ApplicationSet controller will wait before reconciling the ApplicationSet again. | [optional] [default to null]
**Template** | [***ApplicationsApplicationSetTemplate**](applicationsApplicationSetTemplate.md) |  | [optional] [default to null]
**Values** | **map[string]string** | Values contains key/value pairs which are passed directly as parameters to the template. These values will not be sent as parameters to the plugin. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

