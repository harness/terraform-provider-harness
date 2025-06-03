# V1PortStatus

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error_** | **string** | Error is to record the problem with the service port The format of the error shall comply with the following rules: - built-in error values shall be specified in this file and those shall use   CamelCase names - cloud provider specific error values must have names that comply with the   format foo.example.com/CamelCase. --- The regex it matches is (dns1123SubdomainFmt/)?(qualifiedNameFmt) +optional +kubebuilder:validation:Required +kubebuilder:validation:Pattern&#x3D;&#x60;^([a-z0-9]([-a-z0-9]*[a-z0-9])?(\\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*_/)?(([A-Za-z0-9][-A-Za-z0-9_.]*)?[A-Za-z0-9])$&#x60; +kubebuilder:validation:MaxLength&#x3D;316 | [optional] [default to null]
**Port** | **int32** | Port is the port number of the service port of which status is recorded here | [optional] [default to null]
**Protocol** | [***AllOfv1PortStatusProtocol**](AllOfv1PortStatusProtocol.md) | Protocol is the protocol of the service port of which status is recorded here The supported values are: \&quot;TCP\&quot;, \&quot;UDP\&quot;, \&quot;SCTP\&quot; | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

