# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDiscoveredService**](K8sresourceApi.md#GetDiscoveredService) | **Get** /api/v1/agents/{agentIdentity}/discoveredservices/{dsvc_id} | Get discovered service by id
[**GetPodForDiscoveredService**](K8sresourceApi.md#GetPodForDiscoveredService) | **Get** /api/v1/agents/{agentIdentity}/discoveredservices/{dsvc_id}/workloads/{wl_uid}/pods/{pod_uid} | Get pod linked to discovered service for a given workload
[**GetServiceForDiscoveredService**](K8sresourceApi.md#GetServiceForDiscoveredService) | **Get** /api/v1/agents/{agentIdentity}/discoveredservices/{dsvc_id}/service | Get service linked to discovered service
[**ListConnection**](K8sresourceApi.md#ListConnection) | **Get** /api/v1/agents/{agentIdentity}/connections | List connection
[**ListConnectionOfAPod**](K8sresourceApi.md#ListConnectionOfAPod) | **Get** /api/v1/agents/{agentIdentity}/pods/{pod_uid}/connections | List connection of a pod linked to discovered service for a given workload
[**ListContainerOfAPod**](K8sresourceApi.md#ListContainerOfAPod) | **Get** /api/v1/agents/{agentIdentity}/pods/{pod_uid}/containers | List cantainer of a pod linked to discovered service for a given workload
[**ListCronjob**](K8sresourceApi.md#ListCronjob) | **Get** /api/v1/agents/{agentIdentity}/cronjobs | Get list of cronjobs
[**ListDeamonset**](K8sresourceApi.md#ListDeamonset) | **Get** /api/v1/agents/{agentIdentity}/daemonsets | Get list of daemonsets
[**ListDeployment**](K8sresourceApi.md#ListDeployment) | **Get** /api/v1/agents/{agentIdentity}/deployments | Get list of deployments
[**ListDiscoveredService**](K8sresourceApi.md#ListDiscoveredService) | **Get** /api/v1/agents/{agentIdentity}/discoveredservices | Get list of discovered services
[**ListDiscoveredServiceConnection**](K8sresourceApi.md#ListDiscoveredServiceConnection) | **Get** /api/v1/agents/{agentIdentity}/discoveredserviceconnections | List DiscoveredService Connection
[**ListJob**](K8sresourceApi.md#ListJob) | **Get** /api/v1/agents/{agentIdentity}/jobs | Get list of jobs
[**ListNamespace**](K8sresourceApi.md#ListNamespace) | **Get** /api/v1/agents/{agentIdentity}/namespaces | Get list of namespaces
[**ListNode**](K8sresourceApi.md#ListNode) | **Get** /api/v1/agents/{agentIdentity}/nodes | Get list of nodes
[**ListPVCOfAPod**](K8sresourceApi.md#ListPVCOfAPod) | **Get** /api/v1/agents/{agentIdentity}/pods/{pod_uid}/volumes | List pvc of a pod linked to discovered service for a given workload
[**ListPod**](K8sresourceApi.md#ListPod) | **Get** /api/v1/agents/{agentIdentity}/pods | Get list of pods
[**ListPodForDiscoveredService**](K8sresourceApi.md#ListPodForDiscoveredService) | **Get** /api/v1/agents/{agentIdentity}/discoveredservices/{dsvc_id}/workloads/{wl_uid}/pods | List pods linked to discovered service for a given workload
[**ListReplicaSet**](K8sresourceApi.md#ListReplicaSet) | **Get** /api/v1/agents/{agentIdentity}/replicasets | Get list of replicasets
[**ListReplicationController**](K8sresourceApi.md#ListReplicationController) | **Get** /api/v1/agents/{agentIdentity}/replicationcontrollers | Get list of replicationcontrollers
[**ListService**](K8sresourceApi.md#ListService) | **Get** /api/v1/agents/{agentIdentity}/services | Get list of servces
[**ListStatefulSet**](K8sresourceApi.md#ListStatefulSet) | **Get** /api/v1/nfras/{agentIdentity}/statefulsets | Get list of statefulsets

# **GetDiscoveredService**
> ApiGetDiscoveredService GetDiscoveredService(ctx, agentIdentity, dsvcId, accountIdentifier, environmentIdentifier, optional)
Get discovered service by id

Get discovered service by id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **agentIdentity** | **string**| agent identity | 
  **dsvcId** | **string**| discovered service id | 
  **accountIdentifier** | **string**| account id that want to access the resource | 
  **environmentIdentifier** | **string**| environment id is the environment where you want to access the resource | 
 **optional** | ***K8sresourceApiGetDiscoveredServiceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a K8sresourceApiGetDiscoveredServiceOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **correlationID** | **optional.String**| correlation id is used to debug micro svc communication | 
 **organizationIdentifier** | **optional.String**| organization id that want to access the resource | 
 **projectIdentifier** | **optional.String**| project id that want to access the resource | 

### Return type

[**ApiGetDiscoveredService**](api.GetDiscoveredService.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPodForDiscoveredService**
> ApiGetPodResponse GetPodForDiscoveredService(ctx, agentIdentity, dsvcId, wlUid, podUid, accountIdentifier, environmentIdentifier, optional)
Get pod linked to discovered service for a given workload

Get pod linked to discovered service for a given workload

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **agentIdentity** | **string**| agent identity | 
  **dsvcId** | **string**| discovered service id | 
  **wlUid** | **string**| uid of workload | 
  **podUid** | **string**| uid of pod | 
  **accountIdentifier** | **string**| account id that want to access the resource | 
  **environmentIdentifier** | **string**| environment id is the environment where you want to access the resource | 
 **optional** | ***K8sresourceApiGetPodForDiscoveredServiceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a K8sresourceApiGetPodForDiscoveredServiceOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------






 **correlationID** | **optional.String**| correlation id is used to debug micro svc communication | 
 **organizationIdentifier** | **optional.String**| organization id that want to access the resource | 
 **projectIdentifier** | **optional.String**| project id that want to access the resource | 

### Return type

[**ApiGetPodResponse**](api.GetPodResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetServiceForDiscoveredService**
> ApiGetServiceResponse GetServiceForDiscoveredService(ctx, agentIdentity, dsvcId, accountIdentifier, environmentIdentifier, optional)
Get service linked to discovered service

Get service linked to discovered service

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **agentIdentity** | **string**| agent identity | 
  **dsvcId** | **string**| discovered service id | 
  **accountIdentifier** | **string**| account id that want to access the resource | 
  **environmentIdentifier** | **string**| environment id is the environment where you want to access the resource | 
 **optional** | ***K8sresourceApiGetServiceForDiscoveredServiceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a K8sresourceApiGetServiceForDiscoveredServiceOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **correlationID** | **optional.String**| correlation id is used to debug micro svc communication | 
 **organizationIdentifier** | **optional.String**| organization id that want to access the resource | 
 **projectIdentifier** | **optional.String**| project id that want to access the resource | 

### Return type

[**ApiGetServiceResponse**](api.GetServiceResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListConnection**
> ApiListConnection ListConnection(ctx, agentIdentity, accountIdentifier, environmentIdentifier, optional)
List connection

List connection

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **agentIdentity** | **string**| agent identity | 
  **accountIdentifier** | **string**| account id that want to access the resource | 
  **environmentIdentifier** | **string**| environment id is the environment where you want to access the resource | 
 **optional** | ***K8sresourceApiListConnectionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a K8sresourceApiListConnectionOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **correlationID** | **optional.String**| correlation id is used to debug micro svc communication | 
 **organizationIdentifier** | **optional.String**| organization id that want to access the resource | 
 **projectIdentifier** | **optional.String**| project id that want to access the resource | 

### Return type

[**ApiListConnection**](api.ListConnection.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListConnectionOfAPod**
> ApiListConnection ListConnectionOfAPod(ctx, agentIdentity, podUid, accountIdentifier, environmentIdentifier, optional)
List connection of a pod linked to discovered service for a given workload

List connection of a pod linked to discovered service for a given workload

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **agentIdentity** | **string**| agent identity | 
  **podUid** | **string**| uid of pod | 
  **accountIdentifier** | **string**| account id that want to access the resource | 
  **environmentIdentifier** | **string**| environment id is the environment where you want to access the resource | 
 **optional** | ***K8sresourceApiListConnectionOfAPodOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a K8sresourceApiListConnectionOfAPodOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **correlationID** | **optional.String**| correlation id is used to debug micro svc communication | 
 **organizationIdentifier** | **optional.String**| organization id that want to access the resource | 
 **projectIdentifier** | **optional.String**| project id that want to access the resource | 

### Return type

[**ApiListConnection**](api.ListConnection.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListContainerOfAPod**
> ApiListContainer ListContainerOfAPod(ctx, agentIdentity, podUid, accountIdentifier, environmentIdentifier, optional)
List cantainer of a pod linked to discovered service for a given workload

List cantainer of a pod linked to discovered service for a given workload

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **agentIdentity** | **string**| agent identity | 
  **podUid** | **string**| uid of pod | 
  **accountIdentifier** | **string**| account id that want to access the resource | 
  **environmentIdentifier** | **string**| environment id is the environment where you want to access the resource | 
 **optional** | ***K8sresourceApiListContainerOfAPodOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a K8sresourceApiListContainerOfAPodOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **correlationID** | **optional.String**| correlation id is used to debug micro svc communication | 
 **organizationIdentifier** | **optional.String**| organization id that want to access the resource | 
 **projectIdentifier** | **optional.String**| project id that want to access the resource | 

### Return type

[**ApiListContainer**](api.ListContainer.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListCronjob**
> ApiListCronJobResponse ListCronjob(ctx, agentIdentity, accountIdentifier, environmentIdentifier, page, limit, all, optional)
Get list of cronjobs

Get list of cronjobs present in the kubernetes agent, name and namespace can be passed as filter in query param

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **agentIdentity** | **string**| agent identity | 
  **accountIdentifier** | **string**| account id that want to access the resource | 
  **environmentIdentifier** | **string**| environment id is the environment where you want to access the resource | 
  **page** | **int32**| page number | [default to 0]
  **limit** | **int32**| limit per page | [default to 10]
  **all** | **bool**| get all | [default to false]
 **optional** | ***K8sresourceApiListCronjobOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a K8sresourceApiListCronjobOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------






 **correlationID** | **optional.String**| correlation id is used to debug micro svc communication | 
 **organizationIdentifier** | **optional.String**| organization id that want to access the resource | 
 **projectIdentifier** | **optional.String**| project id that want to access the resource | 
 **name** | **optional.String**| name of the cronjob | 
 **namespace** | **optional.String**| namespace of the cronjob | 

### Return type

[**ApiListCronJobResponse**](api.ListCronJobResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListDeamonset**
> ApiListDaemonSetResponse ListDeamonset(ctx, agentIdentity, accountIdentifier, environmentIdentifier, page, limit, all, optional)
Get list of daemonsets

Get list of daemonsets present in the kubernetes agent, name and namespace can be passed as filter in query param

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **agentIdentity** | **string**| agent identity | 
  **accountIdentifier** | **string**| account id that want to access the resource | 
  **environmentIdentifier** | **string**| environment id is the environment where you want to access the resource | 
  **page** | **int32**| page number | [default to 0]
  **limit** | **int32**| limit per page | [default to 10]
  **all** | **bool**| get all | [default to false]
 **optional** | ***K8sresourceApiListDeamonsetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a K8sresourceApiListDeamonsetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------






 **correlationID** | **optional.String**| correlation id is used to debug micro svc communication | 
 **organizationIdentifier** | **optional.String**| organization id that want to access the resource | 
 **projectIdentifier** | **optional.String**| project id that want to access the resource | 
 **name** | **optional.String**| name of the daemonset | 
 **namespace** | **optional.String**| namespace of the daemonset | 

### Return type

[**ApiListDaemonSetResponse**](api.ListDaemonSetResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListDeployment**
> ApiListDeploymentResponse ListDeployment(ctx, agentIdentity, accountIdentifier, environmentIdentifier, page, limit, all, optional)
Get list of deployments

Get list of deployments present in the kubernetes agent, name and namespace can be passed as filter in query param

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **agentIdentity** | **string**| agent identity | 
  **accountIdentifier** | **string**| account id that want to access the resource | 
  **environmentIdentifier** | **string**| environment id is the environment where you want to access the resource | 
  **page** | **int32**| page number | [default to 0]
  **limit** | **int32**| limit per page | [default to 10]
  **all** | **bool**| get all | [default to false]
 **optional** | ***K8sresourceApiListDeploymentOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a K8sresourceApiListDeploymentOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------






 **correlationID** | **optional.String**| correlation id is used to debug micro svc communication | 
 **organizationIdentifier** | **optional.String**| organization id that want to access the resource | 
 **projectIdentifier** | **optional.String**| project id that want to access the resource | 
 **name** | **optional.String**| name of the deployment | 
 **namespace** | **optional.String**| namespace of the deployment | 

### Return type

[**ApiListDeploymentResponse**](api.ListDeploymentResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListDiscoveredService**
> ApiListDiscoveredService ListDiscoveredService(ctx, agentIdentity, accountIdentifier, environmentIdentifier, page, limit, all, optional)
Get list of discovered services

Get list of discovered services

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **agentIdentity** | **string**| agent identity | 
  **accountIdentifier** | **string**| account id that want to access the resource | 
  **environmentIdentifier** | **string**| environment id is the environment where you want to access the resource | 
  **page** | **int32**| page number | [default to 0]
  **limit** | **int32**| limit per page | [default to 10]
  **all** | **bool**| get all | [default to false]
 **optional** | ***K8sresourceApiListDiscoveredServiceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a K8sresourceApiListDiscoveredServiceOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------






 **correlationID** | **optional.String**| correlation id is used to debug micro svc communication | 
 **organizationIdentifier** | **optional.String**| organization id that want to access the resource | 
 **projectIdentifier** | **optional.String**| project id that want to access the resource | 
 **namespace** | **optional.String**| namespace of the discovered service | 
 **search** | **optional.String**| search based on name | 

### Return type

[**ApiListDiscoveredService**](api.ListDiscoveredService.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListDiscoveredServiceConnection**
> ApiListDiscoveredServiceConnection ListDiscoveredServiceConnection(ctx, agentIdentity, accountIdentifier, environmentIdentifier, optional)
List DiscoveredService Connection

List connections in the context of DiscoveredService

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **agentIdentity** | **string**| agent identity | 
  **accountIdentifier** | **string**| account id that want to access the resource | 
  **environmentIdentifier** | **string**| environment id is the environment where you want to access the resource | 
 **optional** | ***K8sresourceApiListDiscoveredServiceConnectionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a K8sresourceApiListDiscoveredServiceConnectionOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **correlationID** | **optional.String**| correlation id is used to debug micro svc communication | 
 **organizationIdentifier** | **optional.String**| organization id that want to access the resource | 
 **projectIdentifier** | **optional.String**| project id that want to access the resource | 

### Return type

[**ApiListDiscoveredServiceConnection**](api.ListDiscoveredServiceConnection.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListJob**
> ApiListJobResponse ListJob(ctx, agentIdentity, accountIdentifier, environmentIdentifier, page, limit, all, optional)
Get list of jobs

Get list of jobs present in the kubernetes agent, name and namespace can be passed as filter in query param

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **agentIdentity** | **string**| agent identity | 
  **accountIdentifier** | **string**| account id that want to access the resource | 
  **environmentIdentifier** | **string**| environment id is the environment where you want to access the resource | 
  **page** | **int32**| page number | [default to 0]
  **limit** | **int32**| limit per page | [default to 10]
  **all** | **bool**| get all | [default to false]
 **optional** | ***K8sresourceApiListJobOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a K8sresourceApiListJobOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------






 **correlationID** | **optional.String**| correlation id is used to debug micro svc communication | 
 **organizationIdentifier** | **optional.String**| organization id that want to access the resource | 
 **projectIdentifier** | **optional.String**| project id that want to access the resource | 
 **name** | **optional.String**| name of the job | 
 **namespace** | **optional.String**| namespace of the job | 

### Return type

[**ApiListJobResponse**](api.ListJobResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListNamespace**
> ApiListNamespaceResponse ListNamespace(ctx, accountIdentifier, environmentIdentifier, agentIdentity, page, limit, all, optional)
Get list of namespaces

Get list of namespaces present in the kubernetes agent, name can be passed as filter in query param

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| account id that want to access the resource | 
  **environmentIdentifier** | **string**| environment id is the environment where you want to access the resource | 
  **agentIdentity** | **string**| agent identity | 
  **page** | **int32**| page number | [default to 0]
  **limit** | **int32**| limit per page | [default to 10]
  **all** | **bool**| get all | [default to false]
 **optional** | ***K8sresourceApiListNamespaceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a K8sresourceApiListNamespaceOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------






 **correlationID** | **optional.String**| correlation id is used to debug micro svc communication | 
 **organizationIdentifier** | **optional.String**| organization id that want to access the resource | 
 **projectIdentifier** | **optional.String**| project id that want to access the resource | 
 **name** | **optional.String**| name of the namespace | 

### Return type

[**ApiListNamespaceResponse**](api.ListNamespaceResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListNode**
> ApiListNodeResponse ListNode(ctx, agentIdentity, accountIdentifier, environmentIdentifier, page, limit, all, optional)
Get list of nodes

Get list of nodes present in the kubernetes agent, name can be passed as filter in query param

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **agentIdentity** | **string**| agent identity | 
  **accountIdentifier** | **string**| account id that want to access the resource | 
  **environmentIdentifier** | **string**| environment id is the environment where you want to access the resource | 
  **page** | **int32**| page number | [default to 0]
  **limit** | **int32**| limit per page | [default to 10]
  **all** | **bool**| get all | [default to false]
 **optional** | ***K8sresourceApiListNodeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a K8sresourceApiListNodeOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------






 **correlationID** | **optional.String**| correlation id is used to debug micro svc communication | 
 **organizationIdentifier** | **optional.String**| organization id that want to access the resource | 
 **projectIdentifier** | **optional.String**| project id that want to access the resource | 
 **name** | **optional.String**| name of the node | 

### Return type

[**ApiListNodeResponse**](api.ListNodeResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListPVCOfAPod**
> ApiListContainerVolume ListPVCOfAPod(ctx, agentIdentity, podUid, accountIdentifier, environmentIdentifier, optional)
List pvc of a pod linked to discovered service for a given workload

List pvc of a pod linked to discovered service for a given workload

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **agentIdentity** | **string**| agent identity | 
  **podUid** | **string**| uid of pod | 
  **accountIdentifier** | **string**| account id that want to access the resource | 
  **environmentIdentifier** | **string**| environment id is the environment where you want to access the resource | 
 **optional** | ***K8sresourceApiListPVCOfAPodOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a K8sresourceApiListPVCOfAPodOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **correlationID** | **optional.String**| correlation id is used to debug micro svc communication | 
 **organizationIdentifier** | **optional.String**| organization id that want to access the resource | 
 **projectIdentifier** | **optional.String**| project id that want to access the resource | 

### Return type

[**ApiListContainerVolume**](api.ListContainerVolume.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListPod**
> ApiListPodResponse ListPod(ctx, agentIdentity, accountIdentifier, environmentIdentifier, page, limit, all, optional)
Get list of pods

Get list of pods present in the kubernetes agent, name and namespace can be passed as filter in query param

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **agentIdentity** | **string**| agent identity | 
  **accountIdentifier** | **string**| account id that want to access the resource | 
  **environmentIdentifier** | **string**| environment id is the environment where you want to access the resource | 
  **page** | **int32**| page number | [default to 0]
  **limit** | **int32**| limit per page | [default to 10]
  **all** | **bool**| get all | [default to false]
 **optional** | ***K8sresourceApiListPodOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a K8sresourceApiListPodOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------






 **correlationID** | **optional.String**| correlation id is used to debug micro svc communication | 
 **organizationIdentifier** | **optional.String**| organization id that want to access the resource | 
 **projectIdentifier** | **optional.String**| project id that want to access the resource | 
 **name** | **optional.String**| name of the pod | 
 **namespace** | **optional.String**| namespace of the pod | 

### Return type

[**ApiListPodResponse**](api.ListPodResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListPodForDiscoveredService**
> ApiListPodResponse ListPodForDiscoveredService(ctx, agentIdentity, dsvcId, wlUid, accountIdentifier, environmentIdentifier, optional)
List pods linked to discovered service for a given workload

List pods linked to discovered service for a given workload

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **agentIdentity** | **string**| agent identity | 
  **dsvcId** | **string**| discovered service id | 
  **wlUid** | **string**| uid of workload | 
  **accountIdentifier** | **string**| account id that want to access the resource | 
  **environmentIdentifier** | **string**| environment id is the environment where you want to access the resource | 
 **optional** | ***K8sresourceApiListPodForDiscoveredServiceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a K8sresourceApiListPodForDiscoveredServiceOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **correlationID** | **optional.String**| correlation id is used to debug micro svc communication | 
 **organizationIdentifier** | **optional.String**| organization id that want to access the resource | 
 **projectIdentifier** | **optional.String**| project id that want to access the resource | 

### Return type

[**ApiListPodResponse**](api.ListPodResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListReplicaSet**
> ApiListReplicaSetResponse ListReplicaSet(ctx, agentIdentity, accountIdentifier, environmentIdentifier, page, limit, all, optional)
Get list of replicasets

Get list of replicasets present in the kubernetes agent, name and namespace can be passed as filter in query param

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **agentIdentity** | **string**| agent identity | 
  **accountIdentifier** | **string**| account id that want to access the resource | 
  **environmentIdentifier** | **string**| environment id is the environment where you want to access the resource | 
  **page** | **int32**| page number | [default to 0]
  **limit** | **int32**| limit per page | [default to 10]
  **all** | **bool**| get all | [default to false]
 **optional** | ***K8sresourceApiListReplicaSetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a K8sresourceApiListReplicaSetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------






 **correlationID** | **optional.String**| correlation id is used to debug micro svc communication | 
 **organizationIdentifier** | **optional.String**| organization id that want to access the resource | 
 **projectIdentifier** | **optional.String**| project id that want to access the resource | 
 **name** | **optional.String**| name of the replicaset | 
 **namespace** | **optional.String**| namespace of the replicaset | 

### Return type

[**ApiListReplicaSetResponse**](api.ListReplicaSetResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListReplicationController**
> ApiListReplicationControllerResponse ListReplicationController(ctx, agentIdentity, accountIdentifier, environmentIdentifier, page, limit, all, optional)
Get list of replicationcontrollers

Get list of replicationcontrollers present in the kubernetes agent, name and namespace can be passed as filter in query param

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **agentIdentity** | **string**| agent identity | 
  **accountIdentifier** | **string**| account id that want to access the resource | 
  **environmentIdentifier** | **string**| environment id is the environment where you want to access the resource | 
  **page** | **int32**| page number | [default to 0]
  **limit** | **int32**| limit per page | [default to 10]
  **all** | **bool**| get all | [default to false]
 **optional** | ***K8sresourceApiListReplicationControllerOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a K8sresourceApiListReplicationControllerOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------






 **correlationID** | **optional.String**| correlation id is used to debug micro svc communication | 
 **organizationIdentifier** | **optional.String**| organization id that want to access the resource | 
 **projectIdentifier** | **optional.String**| project id that want to access the resource | 
 **name** | **optional.String**| name of the replicationcontroller | 
 **namespace** | **optional.String**| namespace of the replicationcontroller | 

### Return type

[**ApiListReplicationControllerResponse**](api.ListReplicationControllerResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListService**
> ApiListServiceResponse ListService(ctx, agentIdentity, accountIdentifier, environmentIdentifier, page, limit, all, optional)
Get list of servces

Get list of services present in the kubernetes agent, name and namespace can be passed as filter in query param

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **agentIdentity** | **string**| agent identity | 
  **accountIdentifier** | **string**| account id that want to access the resource | 
  **environmentIdentifier** | **string**| environment id is the environment where you want to access the resource | 
  **page** | **int32**| page number | [default to 0]
  **limit** | **int32**| limit per page | [default to 10]
  **all** | **bool**| get all | [default to false]
 **optional** | ***K8sresourceApiListServiceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a K8sresourceApiListServiceOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------






 **correlationID** | **optional.String**| correlation id is used to debug micro svc communication | 
 **organizationIdentifier** | **optional.String**| organization id that want to access the resource | 
 **projectIdentifier** | **optional.String**| project id that want to access the resource | 
 **name** | **optional.String**| name of the service | 
 **namespace** | **optional.String**| namespace of the service | 

### Return type

[**ApiListServiceResponse**](api.ListServiceResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListStatefulSet**
> ApiListStatefulSetResponse ListStatefulSet(ctx, agentIdentity, accountIdentifier, environmentIdentifier, page, limit, all, optional)
Get list of statefulsets

Get list of statefulsets present in the kubernetes agent, name and namespace can be passed as filter in query param

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **agentIdentity** | **string**| agent identity | 
  **accountIdentifier** | **string**| account id that want to access the resource | 
  **environmentIdentifier** | **string**| environment id is the environment where you want to access the resource | 
  **page** | **int32**| page number | [default to 0]
  **limit** | **int32**| limit per page | [default to 10]
  **all** | **bool**| get all | [default to false]
 **optional** | ***K8sresourceApiListStatefulSetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a K8sresourceApiListStatefulSetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------






 **correlationID** | **optional.String**| correlation id is used to debug micro svc communication | 
 **organizationIdentifier** | **optional.String**| organization id that want to access the resource | 
 **projectIdentifier** | **optional.String**| project id that want to access the resource | 
 **name** | **optional.String**| name of the statefulset | 
 **namespace** | **optional.String**| namespace of the statefulset | 

### Return type

[**ApiListStatefulSetResponse**](api.ListStatefulSetResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

