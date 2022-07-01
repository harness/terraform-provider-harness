# AnomalyFilterProperties

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**K8sClusterNames** | **[]string** | This is the list of Cluster Names on which filter will be applied. | [optional] [default to null]
**K8sNamespaces** | **[]string** | This is the list of Namespaces on which filter will be applied. | [optional] [default to null]
**K8sWorkloadNames** | **[]string** | This is the list of Workload Names on which filter will be applied. | [optional] [default to null]
**GcpProjects** | **[]string** | This is the list of GCP Projects on which filter will be applied. | [optional] [default to null]
**GcpProducts** | **[]string** | This is the list of GCP Products on which filter will be applied. | [optional] [default to null]
**GcpSKUDescriptions** | **[]string** | This is the list of GCP SKU Descriptions on which filter will be applied. | [optional] [default to null]
**AwsAccounts** | **[]string** | This is the list of AWS Accounts on which filter will be applied. | [optional] [default to null]
**AwsServices** | **[]string** | This is the list of AWS Services on which filter will be applied. | [optional] [default to null]
**AwsUsageTypes** | **[]string** | This is the list of AWS Usage Types on which filter will be applied. | [optional] [default to null]
**AzureSubscriptionGuids** | **[]string** | This is the list of Azure Subscription Guids on which filter will be applied. | [optional] [default to null]
**AzureResourceGroups** | **[]string** | This is the list of Azure Resource Groups on which filter will be applied. | [optional] [default to null]
**AzureMeterCategories** | **[]string** | This is the list of Azure Meter Categories on which filter will be applied. | [optional] [default to null]
**MinActualAmount** | **float64** | Fetch anomalies with Actual Amount greater-than or equal-to minActualAmount | [optional] [default to null]
**MinAnomalousSpend** | **float64** | Fetch anomalies with Anomalous Spend greater-than or equal-to minAnomalousSpend | [optional] [default to null]
**TimeFilters** | [**[]CcmTimeFilter**](CCMTimeFilter.md) | List of filters to be applied on Anomaly Time | [optional] [default to null]
**OrderBy** | [**[]CcmSort**](CCMSort.md) | The order by condition for anomaly query | [optional] [default to null]
**GroupBy** | [**[]CcmGroupBy**](CCMGroupBy.md) | The group by clause for anomaly query | [optional] [default to null]
**Aggregations** | [**[]CcmAggregation**](CCMAggregation.md) | The aggregations for anomaly query | [optional] [default to null]
**SearchText** | **[]string** | The search text entered to filter out rows | [optional] [default to null]
**Offset** | **int32** | Query Offset | [optional] [default to null]
**Limit** | **int32** | Query Limit | [optional] [default to null]
**Tags** | **map[string]string** | Filter tags as a key-value pair. | [optional] [default to null]
**FilterType** | **string** | This specifies the corresponding Entity of the filter. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

