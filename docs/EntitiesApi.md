# \EntitiesApi

All URIs are relative to *https://vrni.example.com/api/ni*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BulkFetchProblemEvents**](EntitiesApi.md#BulkFetchProblemEvents) | **Post** /entities/problems/fetch | Get details of problem events
[**BulkFetchVendorInfo**](EntitiesApi.md#BulkFetchVendorInfo) | **Post** /entities/vendor-infos/fetch | Get Vendor Information of entities
[**EntitiesFetchPost**](EntitiesApi.md#EntitiesFetchPost) | **Post** /entities/fetch | Get details of entities
[**GetAWSAccountManager**](EntitiesApi.md#GetAWSAccountManager) | **Get** /entities/aws-account-managers/{id} | Show AWS Account manager details
[**GetAzureSubscription**](EntitiesApi.md#GetAzureSubscription) | **Get** /entities/azure-subscriptions/{id} | Show Microsoft Azure Subscription details
[**GetCluster**](EntitiesApi.md#GetCluster) | **Get** /entities/clusters/{id} | Show cluster details
[**GetDatacenter**](EntitiesApi.md#GetDatacenter) | **Get** /entities/vc-datacenters/{id} | Show vCenter datacenter details
[**GetDatastore**](EntitiesApi.md#GetDatastore) | **Get** /entities/datastores/{id} | Show datastore details
[**GetDistributedVirtualPortgroup**](EntitiesApi.md#GetDistributedVirtualPortgroup) | **Get** /entities/distributed-virtual-portgroups/{id} | Show distributed virtual portgroup details
[**GetDistributedVirtualSwitch**](EntitiesApi.md#GetDistributedVirtualSwitch) | **Get** /entities/distributed-virtual-switches/{id} | Show distributed virtual switch details
[**GetFirewall**](EntitiesApi.md#GetFirewall) | **Get** /entities/firewalls/{id} | Show firewall details
[**GetFirewallManager**](EntitiesApi.md#GetFirewallManager) | **Get** /entities/firewall-managers/{id} | Show firewall manager details
[**GetFirewallRule**](EntitiesApi.md#GetFirewallRule) | **Get** /entities/firewall-rules/{id} | Show firewall rule details
[**GetFlow**](EntitiesApi.md#GetFlow) | **Get** /entities/flows/{id} | Show flow details
[**GetFlows**](EntitiesApi.md#GetFlows) | **Get** /entities/flows | List flows
[**GetFolder**](EntitiesApi.md#GetFolder) | **Get** /entities/folders/{id} | Show folder details
[**GetHost**](EntitiesApi.md#GetHost) | **Get** /entities/hosts/{id} | Show host details
[**GetIPSet**](EntitiesApi.md#GetIPSet) | **Get** /entities/ip-sets/{id} | Show NSX IP Set details
[**GetKubernetesServiceById**](EntitiesApi.md#GetKubernetesServiceById) | **Get** /entities/kubernetes-services/{id} | Show Kubernetes service details
[**GetKubernetesServices**](EntitiesApi.md#GetKubernetesServices) | **Get** /entities/kubernetes-services | List Kubernetes Services
[**GetLayer2Network**](EntitiesApi.md#GetLayer2Network) | **Get** /entities/layer2-networks/{id} | Show layer2 network details
[**GetNSXManager**](EntitiesApi.md#GetNSXManager) | **Get** /entities/nsx-managers/{id} | Show NSX manager details
[**GetName**](EntitiesApi.md#GetName) | **Get** /entities/names/{id} | Get name of an entity
[**GetNames**](EntitiesApi.md#GetNames) | **Post** /entities/names | Get names for entities
[**GetProblemEvent**](EntitiesApi.md#GetProblemEvent) | **Get** /entities/problems/{id} | Show problem details
[**GetSecurityGroup**](EntitiesApi.md#GetSecurityGroup) | **Get** /entities/security-groups/{id} | Show security group details
[**GetSecurityTag**](EntitiesApi.md#GetSecurityTag) | **Get** /entities/security-tags/{id} | Show security tag details
[**GetService**](EntitiesApi.md#GetService) | **Get** /entities/services/{id} | Show service details
[**GetServiceGroup**](EntitiesApi.md#GetServiceGroup) | **Get** /entities/service-groups/{id} | Show service group details
[**GetVcenterManager**](EntitiesApi.md#GetVcenterManager) | **Get** /entities/vcenter-managers/{id} | Show vCenter manager details
[**GetVm**](EntitiesApi.md#GetVm) | **Get** /entities/vms/{id} | Show VM details
[**GetVmknic**](EntitiesApi.md#GetVmknic) | **Get** /entities/vmknics/{id} | Show vmknic details
[**GetVnic**](EntitiesApi.md#GetVnic) | **Get** /entities/vnics/{id} | Show vNIC details
[**ListAWSAccountManagers**](EntitiesApi.md#ListAWSAccountManagers) | **Get** /entities/aws-account-managers | List AWS Account managers
[**ListAzureSubscription**](EntitiesApi.md#ListAzureSubscription) | **Get** /entities/azure-subscriptions | List Microsoft Azure Subscriptions
[**ListClusters**](EntitiesApi.md#ListClusters) | **Get** /entities/clusters | List clusters
[**ListDatacenters**](EntitiesApi.md#ListDatacenters) | **Get** /entities/vc-datacenters | List vCenter datacenters
[**ListDatastores**](EntitiesApi.md#ListDatastores) | **Get** /entities/datastores | List datastores
[**ListDistributedVirtualPortgroups**](EntitiesApi.md#ListDistributedVirtualPortgroups) | **Get** /entities/distributed-virtual-portgroups | List distributed virtual portgroups
[**ListDistributedVirtualSwitches**](EntitiesApi.md#ListDistributedVirtualSwitches) | **Get** /entities/distributed-virtual-switches | List distributed virtual switches
[**ListFirewallManagers**](EntitiesApi.md#ListFirewallManagers) | **Get** /entities/firewall-managers | List firewall managers
[**ListFirewallRules**](EntitiesApi.md#ListFirewallRules) | **Get** /entities/firewall-rules | List firewall rules
[**ListFirewalls**](EntitiesApi.md#ListFirewalls) | **Get** /entities/firewalls | List firewalls
[**ListFolders**](EntitiesApi.md#ListFolders) | **Get** /entities/folders | List folders
[**ListHosts**](EntitiesApi.md#ListHosts) | **Get** /entities/hosts | List hosts
[**ListIPSets**](EntitiesApi.md#ListIPSets) | **Get** /entities/ip-sets | List NSX IP Sets
[**ListLayer2Networks**](EntitiesApi.md#ListLayer2Networks) | **Get** /entities/layer2-networks | List layer2 networks
[**ListNSXManagers**](EntitiesApi.md#ListNSXManagers) | **Get** /entities/nsx-managers | List NSX managers
[**ListProblemEvents**](EntitiesApi.md#ListProblemEvents) | **Get** /entities/problems | List problems
[**ListSecurityGroups**](EntitiesApi.md#ListSecurityGroups) | **Get** /entities/security-groups | List security groups
[**ListSecurityTags**](EntitiesApi.md#ListSecurityTags) | **Get** /entities/security-tags | List security tags
[**ListServiceGroups**](EntitiesApi.md#ListServiceGroups) | **Get** /entities/service-groups | List service groups
[**ListServices**](EntitiesApi.md#ListServices) | **Get** /entities/services | List services
[**ListVcenterManagers**](EntitiesApi.md#ListVcenterManagers) | **Get** /entities/vcenter-managers | List vCenter managers
[**ListVmknics**](EntitiesApi.md#ListVmknics) | **Get** /entities/vmknics | List vmknics
[**ListVms**](EntitiesApi.md#ListVms) | **Get** /entities/vms | List vms
[**ListVnics**](EntitiesApi.md#ListVnics) | **Get** /entities/vnics | List vnics



## BulkFetchProblemEvents

> BulkProblemFetchResponse BulkFetchProblemEvents(ctx, optional)

Get details of problem events

Bulk fetch of problems. Max batch size is 1000.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***BulkFetchProblemEventsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a BulkFetchProblemEventsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of FetchRequest**](FetchRequest.md)|  | 

### Return type

[**BulkProblemFetchResponse**](BulkProblemFetchResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BulkFetchVendorInfo

> DetailedVendorInfoResponse BulkFetchVendorInfo(ctx, optional)

Get Vendor Information of entities

Bulk fetch of vendor info. Max batch size is 1000.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***BulkFetchVendorInfoOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a BulkFetchVendorInfoOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of FetchRequest**](FetchRequest.md)|  | 

### Return type

[**DetailedVendorInfoResponse**](DetailedVendorInfoResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EntitiesFetchPost

> BulkFetchResponse EntitiesFetchPost(ctx, optional)

Get details of entities

Bulk fetch of entity details using a list of entity IDs. Max batch size is 1000.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***EntitiesFetchPostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a EntitiesFetchPostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of FetchRequest**](FetchRequest.md)|  | 

### Return type

[**BulkFetchResponse**](BulkFetchResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAWSAccountManager

> AwsAccountManager GetAWSAccountManager(ctx, id, optional)

Show AWS Account manager details

Show AWS Account manager details

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The Entity ID of object requestion information on | 
 **optional** | ***GetAWSAccountManagerOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetAWSAccountManagerOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **time** | **optional.Int64**| Time in epoch seconds | 

### Return type

[**AwsAccountManager**](AWSAccountManager.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, entity_id, name, entity_type, manager, vendor_id

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAzureSubscription

> AzureSubscription GetAzureSubscription(ctx, id, optional)

Show Microsoft Azure Subscription details

Show Microsoft Azure Subscription details

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The Entity ID of object requestion information on | 
 **optional** | ***GetAzureSubscriptionOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetAzureSubscriptionOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **time** | **optional.Int64**| Time in epoch seconds | 

### Return type

[**AzureSubscription**](AzureSubscription.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, entity_id, name, entity_type, manager, vendor_id

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCluster

> Cluster GetCluster(ctx, id, optional)

Show cluster details

Show vSphere cluster details

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The Entity ID of object requestion information on | 
 **optional** | ***GetClusterOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetClusterOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **time** | **optional.Int64**| Time in epoch seconds | 

### Return type

[**Cluster**](Cluster.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, entity_id, name, entity_type, total_cpus, total_memory, num_cpu_cores, vendor_id, num_hosts, num_datastores, nsx_manager, vcenter_manager

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDatacenter

> VcDatacenter GetDatacenter(ctx, id, optional)

Show vCenter datacenter details

Show vCenter datacenter details

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The Entity ID of object requestion information on | 
 **optional** | ***GetDatacenterOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetDatacenterOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **time** | **optional.Int64**| Time in epoch seconds | 

### Return type

[**VcDatacenter**](VCDatacenter.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, entity_id, name, entity_type, vendor_id, vcenter_manager

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDatastore

> Datastore GetDatastore(ctx, id, optional)

Show datastore details

Show vSphere datastore details

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The Entity ID of object requestion information on | 
 **optional** | ***GetDatastoreOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetDatastoreOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **time** | **optional.Int64**| Time in epoch seconds | 

### Return type

[**Datastore**](Datastore.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, entity_id, name, entity_type, vendor_id, vcenter_manager

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDistributedVirtualPortgroup

> DistributedVirtualPortgroup GetDistributedVirtualPortgroup(ctx, id, optional)

Show distributed virtual portgroup details

Show distributed virtual portgroup details

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The Entity ID of object requestion information on | 
 **optional** | ***GetDistributedVirtualPortgroupOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetDistributedVirtualPortgroupOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **time** | **optional.Int64**| Time in epoch seconds | 

### Return type

[**DistributedVirtualPortgroup**](DistributedVirtualPortgroup.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, entity_id, name, entity_type, vendor_id, vcenter_manager, distributed_virtual_switch

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDistributedVirtualSwitch

> DistributedVirtualSwitch GetDistributedVirtualSwitch(ctx, id, optional)

Show distributed virtual switch details

Show distributed virtual switch details

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The Entity ID of object requestion information on | 
 **optional** | ***GetDistributedVirtualSwitchOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetDistributedVirtualSwitchOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **time** | **optional.Int64**| Time in epoch seconds | 

### Return type

[**DistributedVirtualSwitch**](DistributedVirtualSwitch.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, entity_id, name, entity_type, vendor_id, vcenter_manager, hosts

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFirewall

> BaseFirewallRule GetFirewall(ctx, id, optional)

Show firewall details

Show firewall details

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The Entity ID of object requestion information on | 
 **optional** | ***GetFirewallOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetFirewallOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **time** | **optional.Int64**| Time in epoch seconds | 

### Return type

[**BaseFirewallRule**](BaseFirewallRule.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, entity_id, name, entity_type, firewall_rules, exclusions

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFirewallManager

> BaseFirewallManager GetFirewallManager(ctx, id, optional)

Show firewall manager details

Show firewall manager details

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The Entity ID of object requestion information on | 
 **optional** | ***GetFirewallManagerOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetFirewallManagerOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **time** | **optional.Int64**| Time in epoch seconds | 

### Return type

[**BaseFirewallManager**](BaseFirewallManager.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, entity_id, name, entity_type, ip_address

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFirewallRule

> BaseFirewallRule GetFirewallRule(ctx, id, optional)

Show firewall rule details

Show firewall rule details

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The Entity ID of object requestion information on | 
 **optional** | ***GetFirewallRuleOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetFirewallRuleOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **time** | **optional.Int64**| Time in epoch seconds | 

### Return type

[**BaseFirewallRule**](BaseFirewallRule.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, entity_id, name, entity_type, rule_id, section_id, section_name, sequence_number, source_any, destination_any, service_any, sources, destinations, services, action, disabled, source_inversion, destination_inversion, port_ranges, logging_enabled, direction, scope, nsx_managers

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFlow

> Flow GetFlow(ctx, id, optional)

Show flow details

Show network traffic flow details

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The Entity ID of object requestion information on | 
 **optional** | ***GetFlowOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetFlowOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **time** | **optional.Int64**| Time in epoch seconds | 

### Return type

[**Flow**](Flow.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, entity_id, name, entity_type, source_ip, destination_ip, port, source_folders, destination_folders, protocol, source_ip_sets, destination_ip_sets, source_security_groups, destination_security_groups, traffic_type, source_security_tags, destination_security_tags, source_vm_tags, destination_vm_tags, within_host, firewall_action, firewall_rule_id, flow_tag

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFlows

> PagedListResponseWithTime GetFlows(ctx, optional)

List flows

List network traffic flows

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetFlowsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetFlowsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **size** | **optional.Float32**| Page size of results | [default to 10.0]
 **cursor** | **optional.String**| Cursor from previous response | 
 **startTime** | **optional.Float32**| Start time for query in epoch seconds | 
 **endTime** | **optional.Float32**| End time for query in epoch seconds | 

### Return type

[**PagedListResponseWithTime**](PagedListResponseWithTime.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, results, cursor, total_count, start_time, end_time

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFolder

> Folder GetFolder(ctx, id, optional)

Show folder details

Show vCenter folder details

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The Entity ID of object requestion information on | 
 **optional** | ***GetFolderOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetFolderOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **time** | **optional.Int64**| Time in epoch seconds | 

### Return type

[**Folder**](Folder.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, entity_id, name, entity_type, vendor_id, vcenter_manager

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHost

> Host GetHost(ctx, id, optional)

Show host details

Show vSphere host details

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The Entity ID of object requestion information on | 
 **optional** | ***GetHostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetHostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **time** | **optional.Int64**| Time in epoch seconds | 

### Return type

[**Host**](Host.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, entity_id, name, entity_type, vmknics, cluster, vcenter_manager, vm_count, datastores, service_tag, vendor_id, nsx_manager, maintenance_mode, connection_state

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIPSet

> BaseIpSet GetIPSet(ctx, id, optional)

Show NSX IP Set details

Show NSX IP Set details

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The Entity ID of object requestion information on | 
 **optional** | ***GetIPSetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetIPSetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **time** | **optional.Int64**| Time in epoch seconds | 

### Return type

[**BaseIpSet**](BaseIPSet.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, entity_id, name, entity_type, ip_addresses, ip_ranges, ip_numeric_ranges, parent_security_groups, direct_source_rules, direct_destination_rules, indirect_source_rules, indirect_destination_rules, vendor_id, vendor, nsx_managers, scope

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetKubernetesServiceById

> KubernetesService GetKubernetesServiceById(ctx, id, optional)

Show Kubernetes service details

Show Kubernetes service details

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The Entity ID of object requestion information on | 
 **optional** | ***GetKubernetesServiceByIdOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetKubernetesServiceByIdOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **time** | **optional.Int64**| Time in epoch seconds | 

### Return type

[**KubernetesService**](KubernetesService.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, entity_id, vendorId, name, entity_type, type, clusterip

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetKubernetesServices

> PagedListResponseWithTime GetKubernetesServices(ctx, optional)

List Kubernetes Services

List Kubernetes services

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetKubernetesServicesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetKubernetesServicesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **size** | **optional.Float32**| Page size of results | [default to 10.0]
 **cursor** | **optional.String**| Cursor from previous response | 
 **startTime** | **optional.Float32**| Start time for query in epoch seconds | 
 **endTime** | **optional.Float32**| End time for query in epoch seconds | 

### Return type

[**PagedListResponseWithTime**](PagedListResponseWithTime.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, results, cursor, total_count, start_time, end_time

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLayer2Network

> BaseL2Network GetLayer2Network(ctx, id, optional)

Show layer2 network details

Show layer2 network details

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The Entity ID of object requestion information on | 
 **optional** | ***GetLayer2NetworkOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetLayer2NetworkOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **time** | **optional.Int64**| Time in epoch seconds | 

### Return type

[**BaseL2Network**](BaseL2Network.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, entity_id, name, entity_type, network_addresses, gateways, segment_id, vteps, scope, nsx_managers

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNSXManager

> BaseNsxManager GetNSXManager(ctx, id, optional)

Show NSX manager details

Show NSX manager details

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The Entity ID of object requestion information on | 
 **optional** | ***GetNSXManagerOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetNSXManagerOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **time** | **optional.Int64**| Time in epoch seconds | 

### Return type

[**BaseNsxManager**](BaseNSXManager.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, entity_id, name, entity_type, ip_address, version, role

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetName

> EntityName GetName(ctx, id, optional)

Get name of an entity

Get name of an entity

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The Entity ID of object requestion information on | 
 **optional** | ***GetNameOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetNameOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **time** | **optional.Int64**| Time in epoch seconds | 

### Return type

[**EntityName**](EntityName.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNames

> NamesResponse GetNames(ctx, body)

Get names for entities

Get names for entities.Limit of 1000 entities in a single request.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**NamesRequest**](NamesRequest.md)| Names Request | 

### Return type

[**NamesResponse**](NamesResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProblemEvent

> ProblemEvent GetProblemEvent(ctx, id, optional)

Show problem details

Show problem event details.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The Entity ID of object requestion information on | 
 **optional** | ***GetProblemEventOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetProblemEventOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **time** | **optional.Int64**| Time in epoch seconds | 

### Return type

[**ProblemEvent**](ProblemEvent.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, entity_id, name, entity_type, anchor_entities, related_entities, message, event_tags, admin_state, archived, event_time_epoch_ms, severity

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSecurityGroup

> BaseSecurityGroup GetSecurityGroup(ctx, id, optional)

Show security group details

Show security group details

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The Entity ID of object requestion information on | 
 **optional** | ***GetSecurityGroupOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetSecurityGroupOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **time** | **optional.Int64**| Time in epoch seconds | 

### Return type

[**BaseSecurityGroup**](BaseSecurityGroup.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, entity_id, name, entity_type, members, direct_source_rules, direct_destination_rules, indirect_source_rules, indirect_destination_rules, parents, direct_members, vendor_id, excluded_members, nsx_managers, scope, ip_sets, security_tags

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSecurityTag

> SecurityTag GetSecurityTag(ctx, id, optional)

Show security tag details

Show security tag details

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The Entity ID of object requestion information on | 
 **optional** | ***GetSecurityTagOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetSecurityTagOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **time** | **optional.Int64**| Time in epoch seconds | 

### Return type

[**SecurityTag**](SecurityTag.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, entity_id, name, entity_type, description, direct_security_groups, security_groups, vendor_id, nsx_manager

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetService

> BaseService GetService(ctx, id, optional)

Show service details

Show service details

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The Entity ID of object requestion information on | 
 **optional** | ***GetServiceOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetServiceOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **time** | **optional.Int64**| Time in epoch seconds | 

### Return type

[**BaseService**](BaseService.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, entity_id, name, entity_type, protocol, port_ranges, nsx_managers, scope, vendor_id

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServiceGroup

> BaseServiceGroup GetServiceGroup(ctx, id, optional)

Show service group details

Show service group details (used in firewall rules)

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The Entity ID of object requestion information on | 
 **optional** | ***GetServiceGroupOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetServiceGroupOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **time** | **optional.Int64**| Time in epoch seconds | 

### Return type

[**BaseServiceGroup**](BaseServiceGroup.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, entity_id, name, entity_type, members, scope, nsx_managers, vendor_id

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVcenterManager

> VCenterManager GetVcenterManager(ctx, id, optional)

Show vCenter manager details

Show vCenter manager details

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The Entity ID of object requestion information on | 
 **optional** | ***GetVcenterManagerOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetVcenterManagerOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **time** | **optional.Int64**| Time in epoch seconds | 

### Return type

[**VCenterManager**](VCenterManager.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, entity_id, name, entity_type, ip_address

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVm

> BaseVirtualMachine GetVm(ctx, id, optional)

Show VM details

Show VM details

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The Entity ID of object requestion information on | 
 **optional** | ***GetVmOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetVmOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **time** | **optional.Int64**| Time in epoch seconds | 

### Return type

[**BaseVirtualMachine**](BaseVirtualMachine.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, entity_id, name, entity_type, ip_addresses, default_gateway, vnics, security_groups, source_firewall_rules, destination_firewall_rules, ip_sets, cluster, resource_pool, security_tags, layer2_networks, host, vlans, vendor_id, vcenter_manager, folders, datastores, datacenter, nsx_manager, source_inversion_rules, destination_inversion_rules

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVmknic

> Vmknic GetVmknic(ctx, id, optional)

Show vmknic details

Show ESXi host vmknic details

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The Entity ID of object requestion information on | 
 **optional** | ***GetVmknicOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetVmknicOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **time** | **optional.Int64**| Time in epoch seconds | 

### Return type

[**Vmknic**](Vmknic.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, entity_id, name, entity_type, ip_addresses, vlan, host, layer2_network

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVnic

> BaseVnic GetVnic(ctx, id, optional)

Show vNIC details

Show vNIC details

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The Entity ID of object requestion information on | 
 **optional** | ***GetVnicOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetVnicOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **time** | **optional.Int64**| Time in epoch seconds | 

### Return type

[**BaseVnic**](BaseVnic.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, entity_id, name, entity_type, ip_addresses, layer2_network, vlan, vm

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAWSAccountManagers

> PagedListResponseWithTime ListAWSAccountManagers(ctx, optional)

List AWS Account managers

List AWS Account managers

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListAWSAccountManagersOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListAWSAccountManagersOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **size** | **optional.Float32**| Page size of results | [default to 10.0]
 **cursor** | **optional.String**| Cursor from previous response | 
 **startTime** | **optional.Float32**| Start time for query in epoch seconds | 
 **endTime** | **optional.Float32**| End time for query in epoch seconds | 

### Return type

[**PagedListResponseWithTime**](PagedListResponseWithTime.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, results, cursor, total_count, start_time, end_time

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAzureSubscription

> PagedListResponseWithTime ListAzureSubscription(ctx, optional)

List Microsoft Azure Subscriptions

List Microsoft Azure Subscriptions

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListAzureSubscriptionOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListAzureSubscriptionOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **size** | **optional.Float32**| Page size of results | [default to 10.0]
 **cursor** | **optional.String**| Cursor from previous response | 
 **startTime** | **optional.Float32**| Start time for query in epoch seconds | 
 **endTime** | **optional.Float32**| End time for query in epoch seconds | 

### Return type

[**PagedListResponseWithTime**](PagedListResponseWithTime.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, results, cursor, total_count, start_time, end_time

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListClusters

> PagedListResponseWithTime ListClusters(ctx, optional)

List clusters

List vSphere clusters

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListClustersOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListClustersOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **size** | **optional.Float32**| Page size of results | [default to 10.0]
 **cursor** | **optional.String**| Cursor from previous response | 
 **startTime** | **optional.Float32**| Start time for query in epoch seconds | 
 **endTime** | **optional.Float32**| End time for query in epoch seconds | 

### Return type

[**PagedListResponseWithTime**](PagedListResponseWithTime.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, results, cursor, total_count, start_time, end_time

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDatacenters

> PagedListResponseWithTime ListDatacenters(ctx, optional)

List vCenter datacenters

List vCenter datacenters

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListDatacentersOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListDatacentersOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **size** | **optional.Float32**| Page size of results | [default to 10.0]
 **cursor** | **optional.String**| Cursor from previous response | 
 **startTime** | **optional.Float32**| Start time for query in epoch seconds | 
 **endTime** | **optional.Float32**| End time for query in epoch seconds | 

### Return type

[**PagedListResponseWithTime**](PagedListResponseWithTime.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, results, cursor, total_count, start_time, end_time

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDatastores

> PagedListResponseWithTime ListDatastores(ctx, optional)

List datastores

List vSphere datastores

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListDatastoresOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListDatastoresOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **size** | **optional.Float32**| Page size of results | [default to 10.0]
 **cursor** | **optional.String**| Cursor from previous response | 
 **startTime** | **optional.Float32**| Start time for query in epoch seconds | 
 **endTime** | **optional.Float32**| End time for query in epoch seconds | 

### Return type

[**PagedListResponseWithTime**](PagedListResponseWithTime.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, results, cursor, total_count, start_time, end_time

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDistributedVirtualPortgroups

> PagedListResponseWithTime ListDistributedVirtualPortgroups(ctx, optional)

List distributed virtual portgroups

List distributed virtual portgroups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListDistributedVirtualPortgroupsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListDistributedVirtualPortgroupsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **size** | **optional.Float32**| Page size of results | [default to 10.0]
 **cursor** | **optional.String**| Cursor from previous response | 
 **startTime** | **optional.Float32**| Start time for query in epoch seconds | 
 **endTime** | **optional.Float32**| End time for query in epoch seconds | 

### Return type

[**PagedListResponseWithTime**](PagedListResponseWithTime.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, results, cursor, total_count, start_time, end_time

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDistributedVirtualSwitches

> PagedListResponseWithTime ListDistributedVirtualSwitches(ctx, optional)

List distributed virtual switches

List distributed virtual switches

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListDistributedVirtualSwitchesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListDistributedVirtualSwitchesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **size** | **optional.Float32**| Page size of results | [default to 10.0]
 **cursor** | **optional.String**| Cursor from previous response | 
 **startTime** | **optional.Float32**| Start time for query in epoch seconds | 
 **endTime** | **optional.Float32**| End time for query in epoch seconds | 

### Return type

[**PagedListResponseWithTime**](PagedListResponseWithTime.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, results, cursor, total_count, start_time, end_time

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListFirewallManagers

> PagedListResponseWithTime ListFirewallManagers(ctx, optional)

List firewall managers

List firewall managers

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListFirewallManagersOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListFirewallManagersOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **size** | **optional.Float32**| Page size of results | [default to 10.0]
 **cursor** | **optional.String**| Cursor from previous response | 
 **startTime** | **optional.Float32**| Start time for query in epoch seconds | 
 **endTime** | **optional.Float32**| End time for query in epoch seconds | 

### Return type

[**PagedListResponseWithTime**](PagedListResponseWithTime.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, results, cursor, total_count, start_time, end_time

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListFirewallRules

> PagedListResponseWithTime ListFirewallRules(ctx, optional)

List firewall rules

List firewall rules for NSX, AWS, Azure, physical firewalls, anything that has firewall rules

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListFirewallRulesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListFirewallRulesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **size** | **optional.Float32**| Page size of results | [default to 10.0]
 **cursor** | **optional.String**| Cursor from previous response | 
 **startTime** | **optional.Float32**| Start time for query in epoch seconds | 
 **endTime** | **optional.Float32**| End time for query in epoch seconds | 

### Return type

[**PagedListResponseWithTime**](PagedListResponseWithTime.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, results, cursor, total_count, start_time, end_time

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListFirewalls

> PagedListResponseWithTime ListFirewalls(ctx, optional)

List firewalls

List firewalls

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListFirewallsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListFirewallsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **size** | **optional.Float32**| Page size of results | [default to 10.0]
 **cursor** | **optional.String**| Cursor from previous response | 
 **startTime** | **optional.Float32**| Start time for query in epoch seconds | 
 **endTime** | **optional.Float32**| End time for query in epoch seconds | 

### Return type

[**PagedListResponseWithTime**](PagedListResponseWithTime.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, results, cursor, total_count, start_time, end_time

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListFolders

> PagedListResponseWithTime ListFolders(ctx, optional)

List folders

List vCenter folders

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListFoldersOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListFoldersOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **size** | **optional.Float32**| Page size of results | [default to 10.0]
 **cursor** | **optional.String**| Cursor from previous response | 
 **startTime** | **optional.Float32**| Start time for query in epoch seconds | 
 **endTime** | **optional.Float32**| End time for query in epoch seconds | 

### Return type

[**PagedListResponseWithTime**](PagedListResponseWithTime.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, results, cursor, total_count, start_time, end_time

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListHosts

> PagedListResponseWithTime ListHosts(ctx, optional)

List hosts

List vSphere hosts

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListHostsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListHostsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **size** | **optional.Float32**| Page size of results | [default to 10.0]
 **cursor** | **optional.String**| Cursor from previous response | 
 **startTime** | **optional.Float32**| Start time for query in epoch seconds | 
 **endTime** | **optional.Float32**| End time for query in epoch seconds | 

### Return type

[**PagedListResponseWithTime**](PagedListResponseWithTime.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, results, cursor, total_count, start_time, end_time

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListIPSets

> PagedListResponseWithTime ListIPSets(ctx, optional)

List NSX IP Sets

List NSX IP Sets

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListIPSetsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListIPSetsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **size** | **optional.Float32**| Page size of results | [default to 10.0]
 **cursor** | **optional.String**| Cursor from previous response | 
 **startTime** | **optional.Float32**| Start time for query in epoch seconds | 
 **endTime** | **optional.Float32**| End time for query in epoch seconds | 

### Return type

[**PagedListResponseWithTime**](PagedListResponseWithTime.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, results, cursor, total_count, start_time, end_time

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListLayer2Networks

> PagedListResponseWithTime ListLayer2Networks(ctx, optional)

List layer2 networks

List layer2 networks (VLANs, VXLANs)

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListLayer2NetworksOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListLayer2NetworksOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **size** | **optional.Float32**| Page size of results | [default to 10.0]
 **cursor** | **optional.String**| Cursor from previous response | 
 **startTime** | **optional.Float32**| Start time for query in epoch seconds | 
 **endTime** | **optional.Float32**| End time for query in epoch seconds | 

### Return type

[**PagedListResponseWithTime**](PagedListResponseWithTime.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, results, cursor, total_count, start_time, end_time

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListNSXManagers

> PagedListResponseWithTime ListNSXManagers(ctx, optional)

List NSX managers

List NSX managers

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListNSXManagersOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListNSXManagersOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **size** | **optional.Float32**| Page size of results | [default to 10.0]
 **cursor** | **optional.String**| Cursor from previous response | 
 **startTime** | **optional.Float32**| Start time for query in epoch seconds | 
 **endTime** | **optional.Float32**| End time for query in epoch seconds | 

### Return type

[**PagedListResponseWithTime**](PagedListResponseWithTime.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, results, cursor, total_count, start_time, end_time

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListProblemEvents

> PagedListResponseWithTime ListProblemEvents(ctx, optional)

List problems

List problem events.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListProblemEventsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListProblemEventsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **size** | **optional.Float32**| Page size of results | [default to 10.0]
 **cursor** | **optional.String**| Cursor from previous response | 
 **startTime** | **optional.Float32**| Start time for query in epoch seconds | 
 **endTime** | **optional.Float32**| End time for query in epoch seconds | 
 **eventType** | **optional.String**| Type of event, e.g UserDefinedProblemEvent | 
 **eventTags** | [**optional.Interface of []string**](string.md)| Event tags | 
 **eventStatus** | **optional.String**| Status of event open or closed | [default to all]
 **updateTimeFrom** | **optional.Float32**| Events that were created or updated or closed between this time and update_time_to, in seconds | 
 **updateTimeTo** | **optional.Float32**| Events that were created or updated or closed between update_time_from and this time, in seconds | 
 **eventSeverity** | [**optional.Interface of []string**](string.md)| Event severity filter | 
 **managers** | [**optional.Interface of []string**](string.md)| The entity ID of the manager of entity on which event is raised | 

### Return type

[**PagedListResponseWithTime**](PagedListResponseWithTime.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, results, cursor, total_count, start_time, end_time

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSecurityGroups

> PagedListResponseWithTime ListSecurityGroups(ctx, optional)

List security groups

List security groups (NSX, AWS, all security groups)

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListSecurityGroupsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListSecurityGroupsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **size** | **optional.Float32**| Page size of results | [default to 10.0]
 **cursor** | **optional.String**| Cursor from previous response | 
 **startTime** | **optional.Float32**| Start time for query in epoch seconds | 
 **endTime** | **optional.Float32**| End time for query in epoch seconds | 

### Return type

[**PagedListResponseWithTime**](PagedListResponseWithTime.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, results, cursor, total_count, start_time, end_time

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSecurityTags

> PagedListResponseWithTime ListSecurityTags(ctx, optional)

List security tags

List security tags

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListSecurityTagsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListSecurityTagsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **size** | **optional.Float32**| Page size of results | [default to 10.0]
 **cursor** | **optional.String**| Cursor from previous response | 
 **startTime** | **optional.Float32**| Start time for query in epoch seconds | 
 **endTime** | **optional.Float32**| End time for query in epoch seconds | 

### Return type

[**PagedListResponseWithTime**](PagedListResponseWithTime.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, results, cursor, total_count, start_time, end_time

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListServiceGroups

> PagedListResponseWithTime ListServiceGroups(ctx, optional)

List service groups

List service groups (used in firewall rules)

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListServiceGroupsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListServiceGroupsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **size** | **optional.Float32**| Page size of results | [default to 10.0]
 **cursor** | **optional.String**| Cursor from previous response | 
 **startTime** | **optional.Float32**| Start time for query in epoch seconds | 
 **endTime** | **optional.Float32**| End time for query in epoch seconds | 

### Return type

[**PagedListResponseWithTime**](PagedListResponseWithTime.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, results, cursor, total_count, start_time, end_time

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListServices

> PagedListResponseWithTime ListServices(ctx, optional)

List services

List services (used in firewall rules)

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListServicesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListServicesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **size** | **optional.Float32**| Page size of results | [default to 10.0]
 **cursor** | **optional.String**| Cursor from previous response | 
 **startTime** | **optional.Float32**| Start time for query in epoch seconds | 
 **endTime** | **optional.Float32**| End time for query in epoch seconds | 

### Return type

[**PagedListResponseWithTime**](PagedListResponseWithTime.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, results, cursor, total_count, start_time, end_time

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListVcenterManagers

> PagedListResponseWithTime ListVcenterManagers(ctx, optional)

List vCenter managers

List vCenter managers

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListVcenterManagersOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListVcenterManagersOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **size** | **optional.Float32**| Page size of results | [default to 10.0]
 **cursor** | **optional.String**| Cursor from previous response | 
 **startTime** | **optional.Float32**| Start time for query in epoch seconds | 
 **endTime** | **optional.Float32**| End time for query in epoch seconds | 

### Return type

[**PagedListResponseWithTime**](PagedListResponseWithTime.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, results, cursor, total_count, start_time, end_time

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListVmknics

> PagedListResponseWithTime ListVmknics(ctx, optional)

List vmknics

List ESXi host vmknics

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListVmknicsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListVmknicsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **size** | **optional.Float32**| Page size of results | [default to 10.0]
 **cursor** | **optional.String**| Cursor from previous response | 
 **startTime** | **optional.Float32**| Start time for query in epoch seconds | 
 **endTime** | **optional.Float32**| End time for query in epoch seconds | 

### Return type

[**PagedListResponseWithTime**](PagedListResponseWithTime.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, results, cursor, total_count, start_time, end_time

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListVms

> PagedListResponseWithTime ListVms(ctx, optional)

List vms

List vms

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListVmsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListVmsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **size** | **optional.Float32**| Page size of results | [default to 10.0]
 **cursor** | **optional.String**| Cursor from previous response | 
 **startTime** | **optional.Float32**| Start time for query in epoch seconds | 
 **endTime** | **optional.Float32**| End time for query in epoch seconds | 

### Return type

[**PagedListResponseWithTime**](PagedListResponseWithTime.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, results, cursor, total_count, start_time, end_time

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListVnics

> PagedListResponseWithTime ListVnics(ctx, optional)

List vnics

List vnics attached to virtual machines

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListVnicsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListVnicsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **size** | **optional.Float32**| Page size of results | [default to 10.0]
 **cursor** | **optional.String**| Cursor from previous response | 
 **startTime** | **optional.Float32**| Start time for query in epoch seconds | 
 **endTime** | **optional.Float32**| End time for query in epoch seconds | 

### Return type

[**PagedListResponseWithTime**](PagedListResponseWithTime.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, results, cursor, total_count, start_time, end_time

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

