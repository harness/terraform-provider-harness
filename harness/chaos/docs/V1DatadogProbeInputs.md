# V1DatadogProbeInputs

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DatadogCredentialsSecretName** | **string** | name of the kubernetes secret containing the Datadog credntials | [optional] [default to null]
**DatadogSite** | **string** | datadog site URL identifier | [optional] [default to null]
**Metrics** | [***AllOfv1DatadogProbeInputsMetrics**](AllOfv1DatadogProbeInputsMetrics.md) | metrics parameters | [optional] [default to null]
**SyntheticsTest** | [***AllOfv1DatadogProbeInputsSyntheticsTest**](AllOfv1DatadogProbeInputsSyntheticsTest.md) | synthetics test parameters | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

