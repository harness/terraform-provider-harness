# CostDetailsQueryParams

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filters** | [**[]FieldFilter**](FieldFilter.md) | Filters to be applied on the response. | [optional] [default to null]
**GroupBy** | **[]string** | Fields on which the response will be grouped by. | [optional] [default to null]
**TimeResolution** | **string** | Only applicable for Time Series Endpoints, defaults to DAY | [optional] [default to null]
**Limit** | **int32** | Limit on the number of cost values returned, 0 by default. | [optional] [default to null]
**SortOrder** | **string** | Order of sorting on cost, Descending by default. | [optional] [default to null]
**Offset** | **int32** | Offset on the cost values returned, 10 by default. | [optional] [default to null]
**SkipRoundOff** | **bool** | Skip Rounding off the cost values returned, false by default. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

