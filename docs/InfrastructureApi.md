# \InfrastructureApi

All URIs are relative to *https://vrni.example.com/api/ni*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetNode**](InfrastructureApi.md#GetNode) | **Get** /infra/nodes/{id} | Show node details
[**ListNodes**](InfrastructureApi.md#ListNodes) | **Get** /infra/nodes | List nodes



## GetNode

> Node GetNode(ctx, id)

Show node details

Get details of infrastructure nodes. Only admin users can get this information. The proxy id is required for adding a data source for selecting appropriate proxy node to add the data source.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The Entity ID of object requestion information on | 

### Return type

[**Node**](Node.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, id, entity_type, node_type, node_id, ip_address

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListNodes

> NodeListResult ListNodes(ctx, )

List nodes

Get list of infrastructure nodes. Only admin users can retrieve this information.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**NodeListResult**](NodeListResult.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, results, total_count

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

