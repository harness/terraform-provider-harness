# CcmRecommendationFilterProperties

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**K8sRecommendationFilterPropertiesDTO** | [***K8sRecommendationFilterProperties**](K8sRecommendationFilterProperties.md) |  | [optional] [default to null]
**PerspectiveFilters** | [**[]QlceViewFilterWrapper**](QLCEViewFilterWrapper.md) | Get Recommendations for a perspective | [optional] [default to null]
**MinSaving** | **float64** | Fetch recommendations with Saving more than minSaving | [optional] [default to null]
**MinCost** | **float64** | Fetch recommendations with Cost more than minCost | [optional] [default to null]
**Offset** | **int64** | Query Offset | [optional] [default to null]
**Limit** | **int64** | Query Limit | [optional] [default to null]
**Tags** | **map[string]string** | Filter tags as a key-value pair. | [optional] [default to null]
**FilterType** | **string** | This specifies the corresponding Entity of the filter. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

